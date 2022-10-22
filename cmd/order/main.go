package main

import (
	"context"

	"github.com/sorohimm/uacs-store-back/internal/service/order"
)

var version, buildTime string

func main() {
	app := order.NewService()
	app.Init(context.Background(), "uacs-order", version, buildTime)
}
