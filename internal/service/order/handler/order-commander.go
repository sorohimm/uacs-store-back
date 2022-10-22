package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sorohimm/uacs-store-back/pkg/api"
)

type OrderCommanderHandler struct {
	api.UnimplementedOrderServiceCommanderServer
}

func (o *OrderCommanderHandler) SubmitOrder(ctx context.Context, req *api.Order) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
