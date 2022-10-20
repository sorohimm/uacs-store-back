package handler

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/sorohimm/uacs-store-back/internal/model"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestStoreRequesterHandler_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockProdReq := model.NewMockProductRequesterHandler(ctrl)

	requester := StoreRequesterHandler{
		productRequester: mockProdReq,
	}

	t.Run("get product no err", func(t *testing.T) {
		ctx := context.Background()

		req := &api.ProductRequest{
			Id: 1,
		}
		mockResp := &dto.Product{
			ID:      1,
			Name:    "someTestProductName",
			Price:   100,
			BrandID: 10,
			TypeID:  1,
			Img:     "12345",
			Info:    nil,
		}

		mockProdReq.EXPECT().GetProductByID(ctx, req.GetId()).Return(mockResp, nil)

		resp, err := requester.GetProduct(ctx, req)

		require.NoError(t, err)
		require.NotNil(t, mockResp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, resp.Id, req.Id)
	})
}
