// Package auth TODO
package auth

import (
	"context"
	"errors"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"github.com/sorohimm/uacs-store-back/pkg/conf"
	"github.com/sorohimm/uacs-store-back/pkg/log"
	stdl "log"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/sorohimm/uacs-store-back/internal"
	"github.com/sorohimm/uacs-store-back/internal/service/auth/config"
	"github.com/sorohimm/uacs-store-back/internal/service/auth/handler"
	"github.com/sorohimm/uacs-store-back/internal/service/auth/initial"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
)

func NewService() *Service {
	return &Service{
		internal.NewRunGroup(),
	}
}

type Service struct {
	*internal.RunGroup
}

func (o *Service) initConfigs(ctx context.Context) context.Context {
	appConf := &config.Config{}
	if err := conf.New(appConf); err != nil {
		if errors.Is(err, conf.ErrHelp) {
			os.Exit(0)
		}
		stdl.Fatalf("failed to read app config: %v", err)
	}
	return config.WithContext(ctx, appConf)
}

func (o *Service) initLogger(ctx context.Context, version, built, appName string) context.Context {
	appConf := config.FromContext(ctx)
	// init logger
	l, err := log.NewZap(
		appConf.Log.Level,
		appConf.Log.EncType)
	if err != nil {
		stdl.Fatalf("failed to init logger: %v", err)
	}
	logger := l.Sugar().With("v", version, "built", built, "app", appName)
	return log.CtxWithLogger(ctx, logger.Desugar())
}

func (o *Service) Init(ctx context.Context, appName, version, built string) {
	var (
		err  error
		pool *pgxpool.Pool
	)

	logger := log.FromContext(ctx).Sugar()

	ctx = o.initConfigs(ctx)
	ctx = o.initLogger(ctx, version, built, appName)

	if pool, err = postgres.NewPGXPool(ctx, config.FromContext(ctx).Postgres); err != nil {
		logger.Fatalf("failed to init ruleset.RepoRuleset from postgres: %v", err)
	}

	cfg := config.FromContext(ctx)

	storeReqHandler := handler.NewAuthHandler(cfg.Postgres.SchemaName, pool).
		SetSigningKey(cfg.JWT.Secret).
		SetAccessExpireDuration(cfg.JWT.AccessTokenExpireDuration).
		SetRefreshExpireDuration(cfg.JWT.RefreshTokenExpireDuration)

	o.Add(initial.Grpc(ctx, func(s *grpc.Server) {
		api.RegisterAuthServiceServer(s, storeReqHandler)
	}))

	exec, inter, err := initial.HTTP(ctx, cfg,
		func(ctx context.Context, mux *runtime.ServeMux, grpcAddr string, opts []grpc.DialOption) error {
			if err = api.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts); err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		logger.Fatalf("failed to init http: %v", err)
	}

	o.Add(exec, inter)
	o.run(ctx)
}

func (o *Service) run(ctx context.Context) {
	logger := log.FromContext(ctx).Sugar()
	// running application
	if err := o.RunGroup.Run(func(err error) {
		if err != nil {
			logger.Error(err)
		}
	}); err != nil {
		logger.Error("unexpected error", zap.Error(err))
	}
}
