package main

import (
	"context"

	"github.com/sorohimm/shop/internal/service/api"
)

var version, buildTime string

func main() {
	app := api.NewService()
	app.Init(context.Background(), "uacs_store", version, buildTime)
}
