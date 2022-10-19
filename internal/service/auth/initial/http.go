// Package initial ...
package initial

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/sorohimm/uacs-store-back/internal/log"
	"github.com/sorohimm/uacs-store-back/internal/service/auth/config"
)

type HTTPRegistrar func(ctx context.Context, mux *runtime.ServeMux, grpcAddr string, opts []grpc.DialOption) error

func HTTP(ctx context.Context, cnf *config.Config, registrar HTTPRegistrar) (func() error, func(error), error) {
	var (
		httpServer *HTTPServer
		err        error
		logger     = log.FromContext(ctx).Sugar()
	)

	httpc := &HTTPConf{
		Timeout: struct {
			Idle       time.Duration
			Read       time.Duration
			ReadHeader time.Duration
			Write      time.Duration
			MustShutIn time.Duration
		}{
			Idle:       cnf.HTTP.Timeout.Idle,
			Read:       cnf.HTTP.Timeout.Read,
			ReadHeader: cnf.HTTP.Timeout.ReadHeader,
			Write:      cnf.HTTP.Timeout.Write,
			MustShutIn: cnf.HTTP.Timeout.MustShutIn,
		},
		Host: cnf.HTTP.Host,
		Port: cnf.HTTP.Port,
		TLS: struct {
			Cert string
			Key  string
		}{
			Cert: cnf.HTTP.TLS.Cert,
			Key:  cnf.HTTP.TLS.Key,
		},
	}

	httpMux := http.NewServeMux()
	gwMux := runtime.NewServeMux(grpcGatewayOptions()...)
	httpMux.Handle("/", gwMux)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
	})
	handler := c.Handler(httpMux)

	if httpServer, err = NewHTTPServer(httpc); err != nil {
		return nil, nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)),
	}
	//
	grpcAddr := net.JoinHostPort(cnf.GRPC.Host, strconv.Itoa(cnf.GRPC.Port))

	if registrar != nil {
		if err = registrar(ctx, gwMux, grpcAddr, opts); err != nil {
			return nil, nil, err
		}
	}

	httpServer.Handler = handler

	exec := func() error {
		httpAddr := net.JoinHostPort(cnf.HTTP.Host, strconv.Itoa(cnf.HTTP.Port))
		logger.With("listen http ", httpAddr).Info("starting http ...")
		return httpServer.Start()
	}

	interr := func(err error) {
		logger.Info("stopping http ...")
		if err != nil {
			logger.Errorf("inerrupt http: %v", err)
		}
		if err = httpServer.Shutdown(); err != nil {
			logger.Errorf("http shutdown: %v", err)
		}
	}

	return exec, interr, nil
}

type HTTPConf struct {
	Timeout struct {
		Idle       time.Duration
		Read       time.Duration
		ReadHeader time.Duration
		Write      time.Duration
		MustShutIn time.Duration
	}

	Host string
	Port int
	TLS  struct {
		Cert string
		Key  string
	}
}

func NewHTTPServer(c *HTTPConf) (*HTTPServer, error) {
	if c == nil {
		return nil, fmt.Errorf("empty config")
	}

	srv := &http.Server{
		IdleTimeout:       c.Timeout.Idle,
		ReadTimeout:       c.Timeout.Read,
		WriteTimeout:      c.Timeout.Write,
		ReadHeaderTimeout: c.Timeout.ReadHeader,
		Addr:              net.JoinHostPort(c.Host, strconv.Itoa(c.Port)),
	}

	if err := http2.ConfigureServer(srv, &http2.Server{
		IdleTimeout: c.Timeout.Idle,
	}); err != nil {
		return nil, err
	}

	return &HTTPServer{
		Server:          srv,
		cert:            c.TLS.Cert,
		key:             c.TLS.Key,
		shutdownTimeout: c.Timeout.MustShutIn,
	}, nil
}

type HTTPServer struct {
	*http.Server
	cert, key       string
	shutdownTimeout time.Duration
}

func (o *HTTPServer) Start() error {
	if o.cert != "" && o.key != "" {
		return o.Server.ListenAndServeTLS(o.cert, o.key)
	}
	return o.Server.ListenAndServe()
}

func (o *HTTPServer) Shutdown() error {
	stopCtx, cancel := context.WithTimeout(context.Background(), o.shutdownTimeout)
	defer cancel()
	return o.Server.Shutdown(stopCtx)
}

func grpcGatewayOptions() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		runtime.WithMarshalerOption("application/json", &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseEnumNumbers:  true,
				EmitUnpopulated: true,
				UseProtoNames:   true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		// runtime.WithMarshalerOption("application/octet-stream", new(runtime.ProtoMarshaller)),
		// runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		// 	switch key {
		// 	case "X-Session", "X-Request-ID":
		// 		return key, true
		// 	default:
		// 		return runtime.DefaultHeaderMatcher(key)
		// 	}
		// }),

		// func(context.Context, http.ResponseWriter, proto.Message) error
		// runtime.WithForwardResponseOption(func(ctx context.Context, w http.ResponseWriter, message proto.Message) error {
		//	md, ok := runtime.ServerMetadataFromContext(ctx)
		//	if !ok {
		//		return nil
		//	}
		//	val := strings.Join(md.HeaderMD.Get("x-http-code"), "")
		//	if val != "" {
		//		code, err := strconv.ParseInt(val, 10, 64)
		//		if err == nil {
		//			w.WriteHeader(int(code))
		//		}
		//		delete(md.HeaderMD, "x-http-code")
		//		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		//	}
		//	return nil
		// }),
	}
}
