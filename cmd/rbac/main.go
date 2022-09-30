package main

import (
	"context"
	"github.com/sorohimm/uacs-store-back/internal/service/rbac"
)

var version, buildTime string

func main() {
	app := rbac.NewService()
	app.Init(context.Background(), "uacs-auth-service", version, buildTime)
}
