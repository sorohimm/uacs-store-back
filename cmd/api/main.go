package main

import (
	"context"

	"github.com/sorohimm/uacs-store-back/internal/service/api"
)

var version, buildTime string

func main() {
	app := api.NewService()
	app.Init(context.Background(), "uacs-store", version, buildTime)
}
