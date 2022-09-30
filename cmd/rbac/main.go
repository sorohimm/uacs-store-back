package main

import (
	"context"
	"github.com/sorohimm/uacs-store-back/internal/service/auth"
)

var version, buildTime string

func main() {
	app := auth.NewService()
	app.Init(context.Background(), "uacs-jwt-service", version, buildTime)
}
