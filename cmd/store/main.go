package main

import (
	"context"
	"github.com/sorohimm/shop/internal/service/store"
)

var version, buildTime string

func main() {
	app := store.NewService()
	app.Init(context.Background(), "uacs_store", version, buildTime)
}
