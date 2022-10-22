package handler

import (
	"context"
	"errors"
	"github.com/sorohimm/uacs-store-back/internal/model/product"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/brand"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/category"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/product/dto"
	"github.com/sorohimm/uacs-store-back/pkg/product"
)

func TestProductCommanderHandler_CreateBrand(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockBrandCmdr := product.NewMockBrandCommanderHandler(ctrl)
	c := &ProductCommanderHandler{
		brandCommander: mockBrandCmdr,
	}

	t.Run("create brand no err", func(t *testing.T) {
		ctx := context.Background()
		req := &product.CreateBrandRequest{Name: "testBrandName"}
		expResp := &brand.Brand{
			ID:   1,
			Name: "testBrandName",
		}

		mockBrandCmdr.EXPECT().CreateBrand(ctx, req).Return(expResp, nil)

		resp, err := c.CreateBrand(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, expResp.ID, resp.ID)
		require.Equal(t, expResp.Name, resp.Name)
	})

	t.Run("create brand internal err", func(t *testing.T) {
		ctx := context.Background()
		req := &product.CreateBrandRequest{Name: "testBrandName"}

		expErr := errors.New("some internal err")
		mockBrandCmdr.EXPECT().CreateBrand(ctx, req).Return(nil, expErr)

		resp, err := c.CreateBrand(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestProductCommanderHandler_CreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockCategoryCmdr := product.NewMockCategoryCommanderHandler(ctrl)
	c := &ProductCommanderHandler{
		categoryCommander: mockCategoryCmdr,
	}

	t.Run("create category no err", func(t *testing.T) {
		ctx := context.Background()

		req := &product.CreateCategoryRequest{Name: "someTestCategory"}
		expResp := &category.Category{
			ID:   1,
			Name: "someTestCategory",
		}

		mockCategoryCmdr.EXPECT().CreateCategory(ctx, req).Return(expResp, nil)

		resp, err := c.CreateCategory(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, expResp.ID, resp.Id)
		require.Equal(t, expResp.Name, resp.Name)
	})

	t.Run("create category internal err", func(t *testing.T) {
		ctx := context.Background()

		req := &product.CreateCategoryRequest{Name: "someTestCategory"}

		err := errors.New("some internal create category err")
		mockCategoryCmdr.EXPECT().CreateCategory(ctx, req).Return(nil, err)

		resp, err := c.CreateCategory(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestProductCommanderHandler_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockCategoryCmdr := product.NewMockProductCommanderHandler(ctrl)
	c := &ProductCommanderHandler{
		productCommander: mockCategoryCmdr,
	}

	t.Run("create product no err", func(t *testing.T) {
		ctx := context.Background()

		req := &product.CreateProductRequest{
			Name:    "test product name",
			Price:   100,
			BrandId: 10,
			TypeId:  1,
			Info:    nil,
		}
		expResp := &dto.Product{
			ID:      1,
			Name:    "test product name",
			Price:   100,
			BrandID: 10,
			TypeID:  1,
			Info:    nil,
		}

		mockCategoryCmdr.EXPECT().CreateProduct(ctx, req).Return(expResp, nil)

		resp, err := c.CreateProduct(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, codes.OK, status.Code(err))

		require.Equal(t, expResp.ID, resp.Id)
		require.Equal(t, expResp.Name, resp.Name)
		require.Equal(t, expResp.Price, resp.Price)
		// todo: add info, img compare
	})

	t.Run("create product internal err", func(t *testing.T) {
		ctx := context.Background()

		req := &product.CreateProductRequest{
			Name:    "test product name",
			Price:   100,
			BrandId: 10,
			TypeId:  1,
			Info:    nil,
		}

		expErr := errors.New("some internal err")
		mockCategoryCmdr.EXPECT().CreateProduct(ctx, req).Return(nil, expErr)

		resp, err := c.CreateProduct(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})
}
