package handler

import (
	"context"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderCommanderHandler struct {
	api.UnimplementedOrderServiceCommanderServer
}

func (o *OrderCommanderHandler) NewOrder(ctx context.Context, req *api.Order) (*emptypb.Empty, error) {
	return nil, nil
}
