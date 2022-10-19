package handler

import (
	"context"
	"errors"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/category"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sorohimm/uacs-store-back/internal/model"
	"github.com/sorohimm/uacs-store-back/internal/storage/postgres/api/brand"
	"github.com/sorohimm/uacs-store-back/pkg/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewStoreCommanderHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockBrandCmdr := model.NewMockBrandCommanderHandler(ctrl)
	c := &StoreCommanderHandler{
		brandCommander: mockBrandCmdr,
	}

	require.Equal(t, c.brandCommander, mockBrandCmdr)
}

func TestStoreCommanderHandler_CreateBrand(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockBrandCmdr := model.NewMockBrandCommanderHandler(ctrl)
	c := &StoreCommanderHandler{
		brandCommander: mockBrandCmdr,
	}

	t.Run("create brand no err", func(t *testing.T) {
		ctx := context.Background()
		req := &api.CreateBrandRequest{Name: "testBrandName"}
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
		req := &api.CreateBrandRequest{Name: "testBrandName"}

		expErr := errors.New("some internal err")
		mockBrandCmdr.EXPECT().CreateBrand(ctx, req).Return(nil, expErr)

		resp, err := c.CreateBrand(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})
}

func TestStoreCommanderHandler_CreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mockCategoryCmdr := model.NewMockCategoryCommanderHandler(ctrl)
	c := &StoreCommanderHandler{
		categoryCommander: mockCategoryCmdr,
	}

	t.Run("create category no err", func(t *testing.T) {
		ctx := context.Background()

		req := &api.CreateCategoryRequest{Name: "someTestCategory"}
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

	t.Run("create category no err", func(t *testing.T) {
		ctx := context.Background()

		req := &api.CreateCategoryRequest{Name: "someTestCategory"}

		err := errors.New("some internal create category err")
		mockCategoryCmdr.EXPECT().CreateCategory(ctx, req).Return(nil, err)

		resp, err := c.CreateCategory(ctx, req)
		require.Error(t, err)
		require.Nil(t, resp)
		require.Equal(t, codes.Internal, status.Code(err))
	})

}
