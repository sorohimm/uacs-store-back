package api

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sorohimm/shop/internal/storage/postgres"
	"os"

	"github.com/sorohimm/shop/internal"
	"github.com/sorohimm/shop/internal/conf"
	"github.com/sorohimm/shop/internal/log"
	"github.com/sorohimm/shop/internal/service/api/config"

	stdl "log"
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
		stdl.Fatal(err)
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

	logger.Debug(pool)
}
