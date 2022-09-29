// Code generated by MockGen. DO NOT EDIT.
// Source: product-handler.go

// Package model is a generated GoMock package.
package model

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/sorohimm/uacs-store-back/pkg/api"
)

// MockProductRequesterHandler is a mock of ProductRequesterHandler interface.
type MockProductRequesterHandler struct {
	ctrl     *gomock.Controller
	recorder *MockProductRequesterHandlerMockRecorder
}

// MockProductRequesterHandlerMockRecorder is the mock recorder for MockProductRequesterHandler.
type MockProductRequesterHandlerMockRecorder struct {
	mock *MockProductRequesterHandler
}

// NewMockProductRequesterHandler creates a new mock instance.
func NewMockProductRequesterHandler(ctrl *gomock.Controller) *MockProductRequesterHandler {
	mock := &MockProductRequesterHandler{ctrl: ctrl}
	mock.recorder = &MockProductRequesterHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRequesterHandler) EXPECT() *MockProductRequesterHandlerMockRecorder {
	return m.recorder
}

// GetAllProducts mocks base method.
func (m *MockProductRequesterHandler) GetAllProducts(ctx context.Context, req *api.AllProductsRequest) (*api.AllProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts", ctx, req)
	ret0, _ := ret[0].(*api.AllProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductRequesterHandlerMockRecorder) GetAllProducts(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProductRequesterHandler)(nil).GetAllProducts), ctx, req)
}

// GetProduct mocks base method.
func (m *MockProductRequesterHandler) GetProduct(ctx context.Context, req *api.ProductRequest) (*api.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, req)
	ret0, _ := ret[0].(*api.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductRequesterHandlerMockRecorder) GetProduct(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductRequesterHandler)(nil).GetProduct), ctx, req)
}

// MockProductCommanderHandler is a mock of ProductCommanderHandler interface.
type MockProductCommanderHandler struct {
	ctrl     *gomock.Controller
	recorder *MockProductCommanderHandlerMockRecorder
}

// MockProductCommanderHandlerMockRecorder is the mock recorder for MockProductCommanderHandler.
type MockProductCommanderHandlerMockRecorder struct {
	mock *MockProductCommanderHandler
}

// NewMockProductCommanderHandler creates a new mock instance.
func NewMockProductCommanderHandler(ctrl *gomock.Controller) *MockProductCommanderHandler {
	mock := &MockProductCommanderHandler{ctrl: ctrl}
	mock.recorder = &MockProductCommanderHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductCommanderHandler) EXPECT() *MockProductCommanderHandlerMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductCommanderHandler) CreateProduct(ctx context.Context, req *api.CreateProductRequest) (*api.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, req)
	ret0, _ := ret[0].(*api.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductCommanderHandlerMockRecorder) CreateProduct(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductCommanderHandler)(nil).CreateProduct), ctx, req)
}
