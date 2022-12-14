// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pkg/api/store.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string         `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Price   float32        `protobuf:"fixed32,20,opt,name=price,proto3" json:"price,omitempty"`
	BrandId int64          `protobuf:"varint,30,opt,name=brandId,proto3" json:"brandId,omitempty"`
	TypeId  int64          `protobuf:"varint,40,opt,name=typeId,proto3" json:"typeId,omitempty"`
	Info    []*ProductInfo `protobuf:"bytes,50,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetBrandId() int64 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *CreateProductRequest) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *CreateProductRequest) GetInfo() []*ProductInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ProductInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId   int64  `protobuf:"varint,10,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Title       string `protobuf:"bytes,20,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,30,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProductInfo) Reset() {
	*x = ProductInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductInfo) ProtoMessage() {}

func (x *ProductInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductInfo.ProtoReflect.Descriptor instead.
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{1}
}

func (x *ProductInfo) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *ProductInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProductInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{2}
}

func (x *ProductRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64          `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	Name  string         `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	Price float32        `protobuf:"fixed32,30,opt,name=price,proto3" json:"price,omitempty"`
	Img   string         `protobuf:"bytes,40,opt,name=img,proto3" json:"img,omitempty"`
	Info  []*ProductInfo `protobuf:"bytes,50,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{3}
}

func (x *ProductResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *ProductResponse) GetInfo() []*ProductInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type AllProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId int64 `protobuf:"varint,10,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	TypeId  int64 `protobuf:"varint,20,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Limit   int64 `protobuf:"varint,30,opt,name=limit,proto3" json:"limit,omitempty"`
	Page    int64 `protobuf:"varint,40,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *AllProductsRequest) Reset() {
	*x = AllProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllProductsRequest) ProtoMessage() {}

func (x *AllProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllProductsRequest.ProtoReflect.Descriptor instead.
func (*AllProductsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{4}
}

func (x *AllProductsRequest) GetBrandId() int64 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *AllProductsRequest) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *AllProductsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *AllProductsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type AllProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductResponse `protobuf:"bytes,10,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *AllProductsResponse) Reset() {
	*x = AllProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllProductsResponse) ProtoMessage() {}

func (x *AllProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllProductsResponse.ProtoReflect.Descriptor instead.
func (*AllProductsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{5}
}

func (x *AllProductsResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

type CreateBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBrandRequest) Reset() {
	*x = CreateBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandRequest) ProtoMessage() {}

func (x *CreateBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrandRequest.ProtoReflect.Descriptor instead.
func (*CreateBrandRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBrandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateBrandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int64  `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateBrandResponse) Reset() {
	*x = CreateBrandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandResponse) ProtoMessage() {}

func (x *CreateBrandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrandResponse.ProtoReflect.Descriptor instead.
func (*CreateBrandResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{7}
}

func (x *CreateBrandResponse) GetId() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CreateBrandResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{8}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateCategoryResponse) Reset() {
	*x = CreateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_store_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResponse) ProtoMessage() {}

func (x *CreateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_store_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_store_proto_rawDescGZIP(), []int{9}
}

func (x *CreateCategoryResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCategoryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_pkg_api_store_proto protoreflect.FileDescriptor

var file_pkg_api_store_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49,
	0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63,
	0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x64, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x20, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x9e, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d,
	0x67, 0x12, 0x3f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72,
	0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x72, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x62, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72,
	0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xb2, 0x02, 0x0a, 0x15, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x2e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x8e,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x32, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32,
	0xd2, 0x03, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x34, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d,
	0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x8c, 0x01, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x32, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d,
	0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72, 0x6f,
	0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x98, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x35, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69,
	0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2e, 0x75, 0x61, 0x63, 0x73, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x3a, 0x01, 0x2a, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x72, 0x6f, 0x68, 0x69, 0x6d, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x70,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_store_proto_rawDescOnce sync.Once
	file_pkg_api_store_proto_rawDescData = file_pkg_api_store_proto_rawDesc
)

func file_pkg_api_store_proto_rawDescGZIP() []byte {
	file_pkg_api_store_proto_rawDescOnce.Do(func() {
		file_pkg_api_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_store_proto_rawDescData)
	})
	return file_pkg_api_store_proto_rawDescData
}

var file_pkg_api_store_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_api_store_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),   // 0: github.com.sorohimm.uacs_store.CreateProductRequest
	(*ProductInfo)(nil),            // 1: github.com.sorohimm.uacs_store.ProductInfo
	(*ProductRequest)(nil),         // 2: github.com.sorohimm.uacs_store.ProductRequest
	(*ProductResponse)(nil),        // 3: github.com.sorohimm.uacs_store.ProductResponse
	(*AllProductsRequest)(nil),     // 4: github.com.sorohimm.uacs_store.AllProductsRequest
	(*AllProductsResponse)(nil),    // 5: github.com.sorohimm.uacs_store.AllProductsResponse
	(*CreateBrandRequest)(nil),     // 6: github.com.sorohimm.uacs_store.CreateBrandRequest
	(*CreateBrandResponse)(nil),    // 7: github.com.sorohimm.uacs_store.CreateBrandResponse
	(*CreateCategoryRequest)(nil),  // 8: github.com.sorohimm.uacs_store.CreateCategoryRequest
	(*CreateCategoryResponse)(nil), // 9: github.com.sorohimm.uacs_store.CreateCategoryResponse
}
var file_pkg_api_store_proto_depIdxs = []int32{
	1, // 0: github.com.sorohimm.uacs_store.CreateProductRequest.info:type_name -> github.com.sorohimm.uacs_store.ProductInfo
	1, // 1: github.com.sorohimm.uacs_store.ProductResponse.info:type_name -> github.com.sorohimm.uacs_store.ProductInfo
	3, // 2: github.com.sorohimm.uacs_store.AllProductsResponse.products:type_name -> github.com.sorohimm.uacs_store.ProductResponse
	2, // 3: github.com.sorohimm.uacs_store.StoreServiceRequester.GetProduct:input_type -> github.com.sorohimm.uacs_store.ProductRequest
	4, // 4: github.com.sorohimm.uacs_store.StoreServiceRequester.GetAllProducts:input_type -> github.com.sorohimm.uacs_store.AllProductsRequest
	0, // 5: github.com.sorohimm.uacs_store.StoreServiceCommander.CreateProduct:input_type -> github.com.sorohimm.uacs_store.CreateProductRequest
	6, // 6: github.com.sorohimm.uacs_store.StoreServiceCommander.CreateBrand:input_type -> github.com.sorohimm.uacs_store.CreateBrandRequest
	8, // 7: github.com.sorohimm.uacs_store.StoreServiceCommander.CreateCategory:input_type -> github.com.sorohimm.uacs_store.CreateCategoryRequest
	3, // 8: github.com.sorohimm.uacs_store.StoreServiceRequester.GetProduct:output_type -> github.com.sorohimm.uacs_store.ProductResponse
	5, // 9: github.com.sorohimm.uacs_store.StoreServiceRequester.GetAllProducts:output_type -> github.com.sorohimm.uacs_store.AllProductsResponse
	3, // 10: github.com.sorohimm.uacs_store.StoreServiceCommander.CreateProduct:output_type -> github.com.sorohimm.uacs_store.ProductResponse
	7, // 11: github.com.sorohimm.uacs_store.StoreServiceCommander.CreateBrand:output_type -> github.com.sorohimm.uacs_store.CreateBrandResponse
	9, // 12: github.com.sorohimm.uacs_store.StoreServiceCommander.CreateCategory:output_type -> github.com.sorohimm.uacs_store.CreateCategoryResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_api_store_proto_init() }
func file_pkg_api_store_proto_init() {
	if File_pkg_api_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllProductsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllProductsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrandRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrandResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_api_store_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_api_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_api_store_proto_goTypes,
		DependencyIndexes: file_pkg_api_store_proto_depIdxs,
		MessageInfos:      file_pkg_api_store_proto_msgTypes,
	}.Build()
	File_pkg_api_store_proto = out.File
	file_pkg_api_store_proto_rawDesc = nil
	file_pkg_api_store_proto_goTypes = nil
	file_pkg_api_store_proto_depIdxs = nil
}
