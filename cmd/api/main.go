package main

import (
	"context"

	"github.com/sorohimm/uacs-store-back/internal/service/product"
)

var version, buildTime string

func main() {
	app := product.NewService()
	app.Init(context.Background(), "uacs-store", version, buildTime)
}
