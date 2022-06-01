// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bikes/bikes.proto

/*
Package bikes is a generated protocol buffer package.

It is generated from these files:
	proto/bikes/bikes.proto
	proto/bikes/bikes_messages.proto

It has these top-level messages:
	ListBikesRequest
	ListBikesResponse
	GetBikeRequest
	GetBikeResponse
	GetBikesRequest
	GetBikesResponse
	GetBikesByTYPERequest
	GetBikesByTYPEResponse
	GetBikesByMAKERequest
	GetBikesByMAKEResponse
	GetBikesByOWNERRequest
	GetBikesByOWNERResponse
	AddBikeRequest
	AddBikeResponse
	DeleteBikeRequest
	DeleteBikeResponse
	Bike
*/
package bikes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ListBikesRequest struct {
}

func (m *ListBikesRequest) Reset()                    { *m = ListBikesRequest{} }
func (m *ListBikesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBikesRequest) ProtoMessage()               {}
func (*ListBikesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListBikesResponse struct {
	Bikes []*Bike `protobuf:"bytes,1,rep,name=bikes" json:"bikes,omitempty"`
}

func (m *ListBikesResponse) Reset()                    { *m = ListBikesResponse{} }
func (m *ListBikesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBikesResponse) ProtoMessage()               {}
func (*ListBikesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListBikesResponse) GetBikes() []*Bike {
	if m != nil {
		return m.Bikes
	}
	return nil
}

type GetBikeRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetBikeRequest) Reset()                    { *m = GetBikeRequest{} }
func (m *GetBikeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBikeRequest) ProtoMessage()               {}
func (*GetBikeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetBikeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBikeResponse struct {
	Bike *Bike `protobuf:"bytes,1,opt,name=bike" json:"bike,omitempty"`
}

func (m *GetBikeResponse) Reset()                    { *m = GetBikeResponse{} }
func (m *GetBikeResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBikeResponse) ProtoMessage()               {}
func (*GetBikeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetBikeResponse) GetBike() *Bike {
	if m != nil {
		return m.Bike
	}
	return nil
}

type GetBikesRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *GetBikesRequest) Reset()                    { *m = GetBikesRequest{} }
func (m *GetBikesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBikesRequest) ProtoMessage()               {}
func (*GetBikesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetBikesRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type GetBikesResponse struct {
	Bikes []*Bike `protobuf:"bytes,1,rep,name=bikes" json:"bikes,omitempty"`
}

func (m *GetBikesResponse) Reset()                    { *m = GetBikesResponse{} }
func (m *GetBikesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBikesResponse) ProtoMessage()               {}
func (*GetBikesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetBikesResponse) GetBikes() []*Bike {
	if m != nil {
		return m.Bikes
	}
	return nil
}

type GetBikesByTYPERequest struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *GetBikesByTYPERequest) Reset()                    { *m = GetBikesByTYPERequest{} }
func (m *GetBikesByTYPERequest) String() string            { return proto.CompactTextString(m) }
func (*GetBikesByTYPERequest) ProtoMessage()               {}
func (*GetBikesByTYPERequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetBikesByTYPERequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetBikesByTYPEResponse struct {
	Bikes []*Bike `protobuf:"bytes,1,rep,name=bikes" json:"bikes,omitempty"`
}

func (m *GetBikesByTYPEResponse) Reset()                    { *m = GetBikesByTYPEResponse{} }
func (m *GetBikesByTYPEResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBikesByTYPEResponse) ProtoMessage()               {}
func (*GetBikesByTYPEResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetBikesByTYPEResponse) GetBikes() []*Bike {
	if m != nil {
		return m.Bikes
	}
	return nil
}

type GetBikesByMAKERequest struct {
	Make string `protobuf:"bytes,1,opt,name=make" json:"make,omitempty"`
}

func (m *GetBikesByMAKERequest) Reset()                    { *m = GetBikesByMAKERequest{} }
func (m *GetBikesByMAKERequest) String() string            { return proto.CompactTextString(m) }
func (*GetBikesByMAKERequest) ProtoMessage()               {}
func (*GetBikesByMAKERequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetBikesByMAKERequest) GetMake() string {
	if m != nil {
		return m.Make
	}
	return ""
}

type GetBikesByMAKEResponse struct {
	Bikes []*Bike `protobuf:"bytes,1,rep,name=bikes" json:"bikes,omitempty"`
}

func (m *GetBikesByMAKEResponse) Reset()                    { *m = GetBikesByMAKEResponse{} }
func (m *GetBikesByMAKEResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBikesByMAKEResponse) ProtoMessage()               {}
func (*GetBikesByMAKEResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetBikesByMAKEResponse) GetBikes() []*Bike {
	if m != nil {
		return m.Bikes
	}
	return nil
}

type GetBikesByOWNERRequest struct {
	OwnerName string `protobuf:"bytes,1,opt,name=owner_name,json=ownerName" json:"owner_name,omitempty"`
}

func (m *GetBikesByOWNERRequest) Reset()                    { *m = GetBikesByOWNERRequest{} }
func (m *GetBikesByOWNERRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBikesByOWNERRequest) ProtoMessage()               {}
func (*GetBikesByOWNERRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetBikesByOWNERRequest) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

type GetBikesByOWNERResponse struct {
	Bikes []*Bike `protobuf:"bytes,1,rep,name=bikes" json:"bikes,omitempty"`
}

func (m *GetBikesByOWNERResponse) Reset()                    { *m = GetBikesByOWNERResponse{} }
func (m *GetBikesByOWNERResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBikesByOWNERResponse) ProtoMessage()               {}
func (*GetBikesByOWNERResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetBikesByOWNERResponse) GetBikes() []*Bike {
	if m != nil {
		return m.Bikes
	}
	return nil
}

type AddBikeRequest struct {
	Bike *Bike `protobuf:"bytes,1,opt,name=bike" json:"bike,omitempty"`
}

func (m *AddBikeRequest) Reset()                    { *m = AddBikeRequest{} }
func (m *AddBikeRequest) String() string            { return proto.CompactTextString(m) }
func (*AddBikeRequest) ProtoMessage()               {}
func (*AddBikeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AddBikeRequest) GetBike() *Bike {
	if m != nil {
		return m.Bike
	}
	return nil
}

type AddBikeResponse struct {
	Bike *Bike `protobuf:"bytes,1,opt,name=bike" json:"bike,omitempty"`
}

func (m *AddBikeResponse) Reset()                    { *m = AddBikeResponse{} }
func (m *AddBikeResponse) String() string            { return proto.CompactTextString(m) }
func (*AddBikeResponse) ProtoMessage()               {}
func (*AddBikeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AddBikeResponse) GetBike() *Bike {
	if m != nil {
		return m.Bike
	}
	return nil
}

type DeleteBikeRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteBikeRequest) Reset()                    { *m = DeleteBikeRequest{} }
func (m *DeleteBikeRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteBikeRequest) ProtoMessage()               {}
func (*DeleteBikeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *DeleteBikeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBikeResponse struct {
}

func (m *DeleteBikeResponse) Reset()                    { *m = DeleteBikeResponse{} }
func (m *DeleteBikeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteBikeResponse) ProtoMessage()               {}
func (*DeleteBikeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func init() {
	proto.RegisterType((*ListBikesRequest)(nil), "bikerenting.grpc.bikes.v1.ListBikesRequest")
	proto.RegisterType((*ListBikesResponse)(nil), "bikerenting.grpc.bikes.v1.ListBikesResponse")
	proto.RegisterType((*GetBikeRequest)(nil), "bikerenting.grpc.bikes.v1.GetBikeRequest")
	proto.RegisterType((*GetBikeResponse)(nil), "bikerenting.grpc.bikes.v1.GetBikeResponse")
	proto.RegisterType((*GetBikesRequest)(nil), "bikerenting.grpc.bikes.v1.GetBikesRequest")
	proto.RegisterType((*GetBikesResponse)(nil), "bikerenting.grpc.bikes.v1.GetBikesResponse")
	proto.RegisterType((*GetBikesByTYPERequest)(nil), "bikerenting.grpc.bikes.v1.GetBikesByTYPERequest")
	proto.RegisterType((*GetBikesByTYPEResponse)(nil), "bikerenting.grpc.bikes.v1.GetBikesByTYPEResponse")
	proto.RegisterType((*GetBikesByMAKERequest)(nil), "bikerenting.grpc.bikes.v1.GetBikesByMAKERequest")
	proto.RegisterType((*GetBikesByMAKEResponse)(nil), "bikerenting.grpc.bikes.v1.GetBikesByMAKEResponse")
	proto.RegisterType((*GetBikesByOWNERRequest)(nil), "bikerenting.grpc.bikes.v1.GetBikesByOWNERRequest")
	proto.RegisterType((*GetBikesByOWNERResponse)(nil), "bikerenting.grpc.bikes.v1.GetBikesByOWNERResponse")
	proto.RegisterType((*AddBikeRequest)(nil), "bikerenting.grpc.bikes.v1.AddBikeRequest")
	proto.RegisterType((*AddBikeResponse)(nil), "bikerenting.grpc.bikes.v1.AddBikeResponse")
	proto.RegisterType((*DeleteBikeRequest)(nil), "bikerenting.grpc.bikes.v1.DeleteBikeRequest")
	proto.RegisterType((*DeleteBikeResponse)(nil), "bikerenting.grpc.bikes.v1.DeleteBikeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BikesAPI service

type BikesAPIClient interface {
	// Get all bikes
	ListBikes(ctx context.Context, in *ListBikesRequest, opts ...grpc.CallOption) (*ListBikesResponse, error)
	// Get bike by id
	GetBike(ctx context.Context, in *GetBikeRequest, opts ...grpc.CallOption) (*GetBikeResponse, error)
	// Get bikes by ids
	GetBikes(ctx context.Context, in *GetBikesRequest, opts ...grpc.CallOption) (*GetBikesResponse, error)
	// Get bikes by type
	GetBikesByTYPE(ctx context.Context, in *GetBikesByTYPERequest, opts ...grpc.CallOption) (*GetBikesByTYPEResponse, error)
	// Get bikes by make
	GetBikesByMAKE(ctx context.Context, in *GetBikesByMAKERequest, opts ...grpc.CallOption) (*GetBikesByMAKEResponse, error)
	// Get bikes by owner_name
	GetBikesByOWNER(ctx context.Context, in *GetBikesByOWNERRequest, opts ...grpc.CallOption) (*GetBikesByOWNERResponse, error)
	// Add new bike
	AddBike(ctx context.Context, in *AddBikeRequest, opts ...grpc.CallOption) (*AddBikeResponse, error)
	// Delete bike
	DeleteBike(ctx context.Context, in *DeleteBikeRequest, opts ...grpc.CallOption) (*DeleteBikeResponse, error)
}

type bikesAPIClient struct {
	cc *grpc.ClientConn
}

func NewBikesAPIClient(cc *grpc.ClientConn) BikesAPIClient {
	return &bikesAPIClient{cc}
}

func (c *bikesAPIClient) ListBikes(ctx context.Context, in *ListBikesRequest, opts ...grpc.CallOption) (*ListBikesResponse, error) {
	out := new(ListBikesResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/ListBikes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesAPIClient) GetBike(ctx context.Context, in *GetBikeRequest, opts ...grpc.CallOption) (*GetBikeResponse, error) {
	out := new(GetBikeResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/GetBike", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesAPIClient) GetBikes(ctx context.Context, in *GetBikesRequest, opts ...grpc.CallOption) (*GetBikesResponse, error) {
	out := new(GetBikesResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesAPIClient) GetBikesByTYPE(ctx context.Context, in *GetBikesByTYPERequest, opts ...grpc.CallOption) (*GetBikesByTYPEResponse, error) {
	out := new(GetBikesByTYPEResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikesByTYPE", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesAPIClient) GetBikesByMAKE(ctx context.Context, in *GetBikesByMAKERequest, opts ...grpc.CallOption) (*GetBikesByMAKEResponse, error) {
	out := new(GetBikesByMAKEResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikesByMAKE", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesAPIClient) GetBikesByOWNER(ctx context.Context, in *GetBikesByOWNERRequest, opts ...grpc.CallOption) (*GetBikesByOWNERResponse, error) {
	out := new(GetBikesByOWNERResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikesByOWNER", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesAPIClient) AddBike(ctx context.Context, in *AddBikeRequest, opts ...grpc.CallOption) (*AddBikeResponse, error) {
	out := new(AddBikeResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/AddBike", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bikesAPIClient) DeleteBike(ctx context.Context, in *DeleteBikeRequest, opts ...grpc.CallOption) (*DeleteBikeResponse, error) {
	out := new(DeleteBikeResponse)
	err := grpc.Invoke(ctx, "/bikerenting.grpc.bikes.v1.BikesAPI/DeleteBike", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BikesAPI service

type BikesAPIServer interface {
	// Get all bikes
	ListBikes(context.Context, *ListBikesRequest) (*ListBikesResponse, error)
	// Get bike by id
	GetBike(context.Context, *GetBikeRequest) (*GetBikeResponse, error)
	// Get bikes by ids
	GetBikes(context.Context, *GetBikesRequest) (*GetBikesResponse, error)
	// Get bikes by type
	GetBikesByTYPE(context.Context, *GetBikesByTYPERequest) (*GetBikesByTYPEResponse, error)
	// Get bikes by make
	GetBikesByMAKE(context.Context, *GetBikesByMAKERequest) (*GetBikesByMAKEResponse, error)
	// Get bikes by owner_name
	GetBikesByOWNER(context.Context, *GetBikesByOWNERRequest) (*GetBikesByOWNERResponse, error)
	// Add new bike
	AddBike(context.Context, *AddBikeRequest) (*AddBikeResponse, error)
	// Delete bike
	DeleteBike(context.Context, *DeleteBikeRequest) (*DeleteBikeResponse, error)
}

func RegisterBikesAPIServer(s *grpc.Server, srv BikesAPIServer) {
	s.RegisterService(&_BikesAPI_serviceDesc, srv)
}

func _BikesAPI_ListBikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).ListBikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/ListBikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).ListBikes(ctx, req.(*ListBikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BikesAPI_GetBike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).GetBike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/GetBike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).GetBike(ctx, req.(*GetBikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BikesAPI_GetBikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).GetBikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).GetBikes(ctx, req.(*GetBikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BikesAPI_GetBikesByTYPE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBikesByTYPERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).GetBikesByTYPE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikesByTYPE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).GetBikesByTYPE(ctx, req.(*GetBikesByTYPERequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BikesAPI_GetBikesByMAKE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBikesByMAKERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).GetBikesByMAKE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikesByMAKE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).GetBikesByMAKE(ctx, req.(*GetBikesByMAKERequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BikesAPI_GetBikesByOWNER_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBikesByOWNERRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).GetBikesByOWNER(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/GetBikesByOWNER",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).GetBikesByOWNER(ctx, req.(*GetBikesByOWNERRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BikesAPI_AddBike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).AddBike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/AddBike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).AddBike(ctx, req.(*AddBikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BikesAPI_DeleteBike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BikesAPIServer).DeleteBike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bikerenting.grpc.bikes.v1.BikesAPI/DeleteBike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BikesAPIServer).DeleteBike(ctx, req.(*DeleteBikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BikesAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bikerenting.grpc.bikes.v1.BikesAPI",
	HandlerType: (*BikesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBikes",
			Handler:    _BikesAPI_ListBikes_Handler,
		},
		{
			MethodName: "GetBike",
			Handler:    _BikesAPI_GetBike_Handler,
		},
		{
			MethodName: "GetBikes",
			Handler:    _BikesAPI_GetBikes_Handler,
		},
		{
			MethodName: "GetBikesByTYPE",
			Handler:    _BikesAPI_GetBikesByTYPE_Handler,
		},
		{
			MethodName: "GetBikesByMAKE",
			Handler:    _BikesAPI_GetBikesByMAKE_Handler,
		},
		{
			MethodName: "GetBikesByOWNER",
			Handler:    _BikesAPI_GetBikesByOWNER_Handler,
		},
		{
			MethodName: "AddBike",
			Handler:    _BikesAPI_AddBike_Handler,
		},
		{
			MethodName: "DeleteBike",
			Handler:    _BikesAPI_DeleteBike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/bikes/bikes.proto",
}

func init() { proto.RegisterFile("proto/bikes/bikes.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdb, 0x0f, 0xd2, 0x30,
	0x14, 0xc6, 0xc3, 0x45, 0x60, 0xc7, 0x84, 0x4b, 0xa3, 0x82, 0x4b, 0x8c, 0x4b, 0x7d, 0x41, 0xd1,
	0x29, 0x10, 0xe3, 0x33, 0x44, 0x30, 0x78, 0x01, 0xb2, 0x98, 0x18, 0x7d, 0xc1, 0xc1, 0x2a, 0x59,
	0x70, 0x17, 0xb7, 0xa1, 0xf2, 0x4f, 0xfb, 0x37, 0x98, 0xb5, 0xdd, 0x0d, 0x74, 0x29, 0xce, 0x17,
	0x52, 0xba, 0xef, 0x3b, 0xbf, 0x1e, 0x76, 0xbe, 0x02, 0x5d, 0xd7, 0x73, 0x02, 0xe7, 0xe9, 0xd6,
	0x3c, 0x10, 0x9f, 0x7d, 0xaa, 0x74, 0x07, 0xdd, 0x0d, 0xbf, 0x78, 0xc4, 0x0e, 0x4c, 0x7b, 0xaf,
	0xee, 0x3d, 0x77, 0xa7, 0xb2, 0xa7, 0xdf, 0x87, 0xb2, 0x72, 0xe1, 0xd9, 0x58, 0xc4, 0xf7, 0xf5,
	0x7d, 0x64, 0xc6, 0x08, 0xda, 0x6f, 0x4d, 0x3f, 0x98, 0x86, 0xcf, 0x34, 0xf2, 0xed, 0x48, 0xfc,
	0x00, 0xbf, 0x86, 0x4e, 0x6a, 0xcf, 0x77, 0x1d, 0xdb, 0x27, 0xe8, 0x39, 0xdc, 0xa0, 0x05, 0x7a,
	0x25, 0xa5, 0xd2, 0xbf, 0x39, 0xba, 0xaf, 0xfe, 0x95, 0xaa, 0x86, 0x46, 0x8d, 0xa9, 0xb1, 0x02,
	0xcd, 0x57, 0x84, 0x96, 0xe2, 0xd5, 0x51, 0x13, 0xca, 0xa6, 0xd1, 0x2b, 0x29, 0xa5, 0xbe, 0xa4,
	0x95, 0x4d, 0x03, 0xcf, 0xa1, 0x15, 0x2b, 0x38, 0x6b, 0x0c, 0xd5, 0xd0, 0x4d, 0x45, 0x02, 0x28,
	0x2a, 0xc6, 0x0f, 0xe2, 0x3a, 0x51, 0x23, 0xa8, 0x0d, 0x15, 0xd3, 0x60, 0x27, 0x96, 0xb4, 0x70,
	0x89, 0x17, 0xd0, 0x4e, 0x44, 0xc5, 0x3a, 0x1b, 0xc0, 0xed, 0xa8, 0xd4, 0xf4, 0xf4, 0xfe, 0xe3,
	0x7a, 0x16, 0x51, 0x11, 0x54, 0x83, 0x93, 0x4b, 0x78, 0x8b, 0x74, 0x8d, 0x57, 0x70, 0xe7, 0x5c,
	0xfc, 0x1f, 0xe9, 0xef, 0x26, 0x6f, 0xd2, 0x74, 0x4b, 0x3f, 0xc4, 0xf4, 0x70, 0x9d, 0xa5, 0x33,
	0x71, 0x31, 0xfa, 0x8b, 0x74, 0xc1, 0xd5, 0x87, 0xe5, 0x4c, 0x8b, 0xf0, 0xf7, 0x00, 0x9c, 0x1f,
	0x36, 0xf1, 0x36, 0xb6, 0x6e, 0x45, 0x87, 0x90, 0xe8, 0xce, 0x52, 0xb7, 0x08, 0x5e, 0x43, 0xf7,
	0xc2, 0x58, 0xec, 0x28, 0x33, 0x68, 0x4e, 0x0c, 0x23, 0x3d, 0x60, 0xff, 0x34, 0x3d, 0x73, 0x68,
	0xc5, 0x65, 0x8a, 0x4d, 0x61, 0xe7, 0x25, 0xf9, 0x4a, 0x02, 0x92, 0x37, 0xf2, 0xb7, 0x00, 0xa5,
	0x45, 0x8c, 0x37, 0xfa, 0x55, 0x83, 0x06, 0xfd, 0x65, 0x26, 0xeb, 0x05, 0xfa, 0x02, 0x52, 0x9c,
	0x41, 0x34, 0xc8, 0x61, 0x9f, 0xa7, 0x57, 0x7e, 0x2c, 0x26, 0xe6, 0x4d, 0x7e, 0x86, 0x3a, 0x7f,
	0x21, 0xe8, 0x61, 0x8e, 0x31, 0x9b, 0x61, 0xf9, 0x91, 0x88, 0x94, 0x13, 0x76, 0xd0, 0x88, 0x5e,
	0x39, 0x12, 0xf0, 0xc5, 0x7d, 0x0c, 0x84, 0xb4, 0x1c, 0x72, 0x8c, 0xaf, 0x19, 0x9e, 0x2f, 0xf4,
	0x4c, 0xc0, 0x9e, 0xc9, 0xad, 0x3c, 0xbc, 0xc2, 0xf1, 0x27, 0x6c, 0x18, 0x2c, 0x41, 0x6c, 0x2a,
	0xb0, 0x82, 0xd8, 0x4c, 0x6a, 0x7f, 0x26, 0x57, 0x1d, 0x4f, 0x11, 0x12, 0xab, 0x92, 0x8e, 0xaa,
	0x3c, 0xba, 0xc6, 0x92, 0x8c, 0x0b, 0x8f, 0x49, 0xee, 0xb8, 0x64, 0x13, 0x99, 0x3b, 0x2e, 0xe7,
	0xa9, 0x33, 0x01, 0x92, 0x6c, 0xa0, 0xbc, 0x61, 0xbe, 0xc8, 0x99, 0xfc, 0x44, 0x50, 0xcd, 0x50,
	0xd3, 0xfa, 0x27, 0x76, 0x87, 0x6c, 0x6b, 0xf4, 0xbf, 0x70, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff,
	0xad, 0x21, 0x2a, 0x1e, 0x63, 0x07, 0x00, 0x00,
}