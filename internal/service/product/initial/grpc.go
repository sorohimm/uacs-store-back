package initial

import (
	"context"
	"github.com/sorohimm/uacs-store-back/pkg/log"
	"net"
	"strconv"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sorohimm/uacs-store-back/internal/service/product/config"
)

func Grpc(ctx context.Context, registrar func(s *grpc.Server)) (func() error, func(error)) {
	cnf := config.FromContext(ctx).GRPC
	logger := log.FromContext(ctx).Sugar()

	recoveryHandler := func(p interface{}) error {
		return status.Errorf(codes.Internal, "caught panic: %v", p)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcPrometheus.UnaryServerInterceptor,
				grpcRecovery.UnaryServerInterceptor(
					grpcRecovery.WithRecoveryHandler(recoveryHandler),
				),
				// add our logger into every request context
				func(reqCtx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
					return handler(log.CtxWithLogger(reqCtx, logger.Desugar()), req)
				},
			),
		),
	)

	if registrar != nil {
		registrar(grpcServer)
	}

	grpcAddr := net.JoinHostPort(cnf.Host, strconv.Itoa(cnf.Port))

	exec := func() error {
		logger.With("listen", grpcAddr).Info("starting grpc ...")
		lis, err := net.Listen("tcp", grpcAddr)
		defer func() {
			if lis != nil {
				_ = lis.Close()
			}
		}()
		if err != nil {
			return err
		}
		return grpcServer.Serve(lis)
	}

	interr := func(err error) {
		if err != nil {
			logger.Errorf("inerrupt grpc: %v", err)
		}
		logger.Info("stopping grpc ...")
		grpcServer.GracefulStop()
	}

	return exec, interr
}
