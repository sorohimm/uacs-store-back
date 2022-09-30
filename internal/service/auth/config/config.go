package config

import (
	"context"
	"time"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres"
)

type GRPCConf struct {
	Host string `long:"host" default:"0.0.0.0" env:"HOST" description:"grpc host to listen to"`
	Port int    `long:"port" default:"9001" env:"PORT" description:"grpc port to listen to"`
}

type LoggerConf struct {
	Level   string `short:"l" long:"level" env:"LEVEL" description:"logging level" default:"DEBUG"`
	EncType string `long:"enctype" env:"ENCTYPE" description:"log as json or not (console|json)" default:"json" `
}

type JwtConf struct {
	Secret                     string        `long:"secret" env:"SECRET" description:"jwt secret encryption key"`
	AccessTokenExpireDuration  time.Duration `long:"host" env:"ACCESS_EXPIRE_DURATION" default:"24h" description:"access token expire duration"`
	RefreshTokenExpireDuration time.Duration `long:"host" env:"REFRESH_EXPIRE_DURATION" default:"96h" description:"refresh token expire duration"`
}

type HTTPConfig struct {
	Host    string `long:"host" env:"HOST" default:"127.0.0.1" description:"host to listen to"`
	Port    int    `long:"port" env:"PORT" default:"2104" description:"port to listen to"`
	Timeout struct {
		Idle       time.Duration `long:"idle" env:"IDLE" description:"the maximum amount of time to wait for the next request when keep-alives are enabled."`
		Read       time.Duration `long:"read" env:"READ" description:"the maximum duration for reading the entire request, including the body"`
		ReadHeader time.Duration `long:"read-header" env:"READ_HEADER" description:"the maximum duration for reading the request's header"`
		Write      time.Duration `long:"write" env:"ENV" description:"the maximum duration before timing out writes of the response."`
		MustShutIn time.Duration `long:"shut" env:"SHUT" default:"30s" description:"the maximum duration before timing out the graceful shutdown"`
	} `group:"timeout" namespace:"timeout" env-namespace:"TIMEOUT"`
	TLS struct {
		Cert string `long:"cert" env:"CERT" description:"cert file"`
		Key  string `long:"key" env:"KEY"  description:"key file"`
	} `group:"tls opts" namespace:"tls" env-namespace:"TLS"`
}

type Config struct {
	Log      *LoggerConf      `group:"logger option" namespace:"log" env-namespace:"LOG"`
	HTTP     *HTTPConfig      `group:"http grpc gateway option" namespace:"http" env-namespace:"HTTP"`
	GRPC     *GRPCConf        `group:"grpc option" namespace:"grpc" env-namespace:"GRPC"`
	JWT      *JwtConf         `group:"jwt option" namespace:"jwt" env-namespace:"JWT"`
	Postgres *postgres.Config `group:"pg" namespace:"pg" env-namespace:"PG"`
}

type confKey struct{} // or exported to use outside the package

func WithContext(ctx context.Context, c *Config) context.Context {
	return context.WithValue(ctx, confKey{}, c)
}

func FromContext(ctx context.Context) *Config {
	if cc, ok := ctx.Value(confKey{}).(*Config); ok {
		return cc
	}
	return NewDefaultConfig()
}

func NewDefaultConfig() *Config {
	return &Config{}
}
