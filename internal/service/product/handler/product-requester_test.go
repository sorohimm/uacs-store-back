package handler

import (
	"context"
	"errors"
	product2 "github.com/sorohimm/uacs-store-back/internal/model/product"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
)

func TestProductRequesterHandler_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockProdReq := product2.NewMockProductRequesterHandler(ctrl)

	requester := ProductRequesterHandler{
		productRequester: mockProdReq,
	}

	t.Run("get api no err", func(t *testing.T) {
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
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, resp, mockResp.ToAPIResponse())
	})

	t.Run("get api err not found", func(t *testing.T) {
		ctx := context.Background()

		req := &api.ProductRequest{
			Id: 1,
		}

		expErr := product.ErrNotFound
		mockProdReq.EXPECT().GetProductByID(ctx, req.GetId()).Return(nil, expErr)

		resp, err := requester.GetProduct(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("get api err internal", func(t *testing.T) {
		ctx := context.Background()

		req := &api.ProductRequest{
			Id: 0,
		}

		expErr := errors.New("some internal err")
		mockProdReq.EXPECT().GetProductByID(ctx, req.GetId()).Return(nil, expErr)

		resp, err := requester.GetProduct(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestProductRequesterHandler_GetAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockProdReq := product2.NewMockProductRequesterHandler(ctrl)

	requester := ProductRequesterHandler{
		productRequester: mockProdReq,
	}

	t.Run("request product without specific type and brand (no err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			Limit: 10,
			Page:  1,
		}
		mockResp := &dto.Products{
			{
				ID:      1,
				Name:    "someTestProductName",
				Price:   100,
				BrandID: 10,
				TypeID:  1,
				Img:     "12345",
				Info:    nil,
			}, {
				ID:      2,
				Name:    "someTestProductName1",
				Price:   213,
				BrandID: 32424,
				TypeID:  12312,
				Img:     "gqfwc4vt5345vq",
				Info:    nil,
			},
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		mockProdReq.EXPECT().GetAllProducts(ctx, limit, offset).Return(mockResp, nil)

		resp, err := requester.GetAllProducts(ctx, req)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, resp, mockResp.ToAPIResponse())
	})

	t.Run("request product without specific type and brand (Internal err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			Limit: 10,
			Page:  1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := errors.New("some internal err")
		mockProdReq.EXPECT().GetAllProducts(ctx, limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("request product without specific type and brand (NotFound err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			Limit: 10,
			Page:  1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := product.ErrNotFound
		mockProdReq.EXPECT().GetAllProducts(ctx, limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("request product with brandId and with typeId (no err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			BrandId: 115,
			TypeId:  200,
			Limit:   10,
			Page:    1,
		}
		mockResp := &dto.Products{
			{
				ID:      1,
				Name:    "some TOYOTA Brand api",
				Price:   100,
				BrandID: 115,
				TypeID:  200,
				Img:     "12345",
				Info:    nil,
			}, {
				ID:      2,
				Name:    "another TOYOTA Brand api",
				Price:   213,
				BrandID: 115,
				TypeID:  200,
				Img:     "gqfwc4vt5345vq",
				Info:    nil,
			},
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		mockProdReq.EXPECT().GetAllProductsWithBrandAndType(ctx, req.TypeId, req.GetBrandId(), limit, offset).Return(mockResp, nil)

		resp, err := requester.GetAllProducts(ctx, req)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, resp, mockResp.ToAPIResponse())
	})

	t.Run("request product with brandId and with typeId (Internal err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			BrandId: 115,
			TypeId:  200,
			Limit:   10,
			Page:    1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := errors.New("some internal err")
		mockProdReq.EXPECT().GetAllProductsWithBrandAndType(ctx, req.TypeId, req.GetBrandId(), limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("request product with brandId and with typeId (Internal err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			BrandId: 115,
			TypeId:  200,
			Limit:   10,
			Page:    1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := product.ErrNotFound
		mockProdReq.EXPECT().GetAllProductsWithBrandAndType(ctx, req.TypeId, req.GetBrandId(), limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("request product with brandId and without typeId (no err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			BrandId: 115,
			Limit:   10,
			Page:    1,
		}
		mockResp := &dto.Products{
			{
				ID:      1,
				Name:    "some TOYOTA Brand api",
				Price:   100,
				BrandID: 115,
				TypeID:  200,
				Img:     "12345",
				Info:    nil,
			}, {
				ID:      2,
				Name:    "another TOYOTA Brand api",
				Price:   213,
				BrandID: 115,
				TypeID:  200,
				Img:     "gqfwc4vt5345vq",
				Info:    nil,
			},
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		mockProdReq.EXPECT().GetAllProductsWithBrand(ctx, req.GetBrandId(), limit, offset).Return(mockResp, nil)

		resp, err := requester.GetAllProducts(ctx, req)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, resp, mockResp.ToAPIResponse())
	})

	t.Run("request product with brandId and without typeId (Internal err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			BrandId: 115,
			Limit:   10,
			Page:    1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := errors.New("some internal err")
		mockProdReq.EXPECT().GetAllProductsWithBrand(ctx, req.GetBrandId(), limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("request product with brandId and without typeId (NotFoundErr err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			BrandId: 115,
			Limit:   10,
			Page:    1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := product.ErrNotFound
		mockProdReq.EXPECT().GetAllProductsWithBrand(ctx, req.GetBrandId(), limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("request product with typeId and without brandId (no err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			TypeId: 200,
			Limit:  10,
			Page:   1,
		}
		mockResp := &dto.Products{
			{
				ID:      1,
				Name:    "some TOYOTA Brand api",
				Price:   100,
				BrandID: 115,
				TypeID:  200,
				Img:     "12345",
				Info:    nil,
			}, {
				ID:      2,
				Name:    "another TOYOTA Brand api",
				Price:   213,
				BrandID: 115,
				TypeID:  200,
				Img:     "gqfwc4vt5345vq",
				Info:    nil,
			},
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		mockProdReq.EXPECT().GetAllProductsWithType(ctx, req.GetTypeId(), limit, offset).Return(mockResp, nil)

		resp, err := requester.GetAllProducts(ctx, req)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, resp, mockResp.ToAPIResponse())
	})

	t.Run("request product with typeId and without brandId (Internal err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			TypeId: 200,
			Limit:  10,
			Page:   1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := errors.New("some internal err")
		mockProdReq.EXPECT().GetAllProductsWithType(ctx, req.GetTypeId(), limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})

	t.Run("request product with typeId and without brandId (NotFound err)", func(t *testing.T) {
		ctx := context.Background()

		req := &api.AllProductsRequest{
			TypeId: 200,
			Limit:  10,
			Page:   1,
		}

		limit := req.GetLimit()
		offset := req.GetPage()*limit - limit

		expErr := errors.New("some internal err")
		mockProdReq.EXPECT().GetAllProductsWithType(ctx, req.GetTypeId(), limit, offset).Return(nil, expErr)

		resp, err := requester.GetAllProducts(ctx, req)

		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})
}
