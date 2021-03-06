// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2020 Intel Corporation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: updateapp.proto

package updateapp

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UpdateAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateFromAppContext string `protobuf:"bytes,1,opt,name=update_from_appContext,json=updateFromAppContext,proto3" json:"update_from_appContext,omitempty"`
	UpdateToAppContext   string `protobuf:"bytes,2,opt,name=update_to_appContext,json=updateToAppContext,proto3" json:"update_to_appContext,omitempty"`
}

func (x *UpdateAppRequest) Reset() {
	*x = UpdateAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updateapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppRequest) ProtoMessage() {}

func (x *UpdateAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_updateapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppRequest.ProtoReflect.Descriptor instead.
func (*UpdateAppRequest) Descriptor() ([]byte, []int) {
	return file_updateapp_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAppRequest) GetUpdateFromAppContext() string {
	if x != nil {
		return x.UpdateFromAppContext
	}
	return ""
}

func (x *UpdateAppRequest) GetUpdateToAppContext() string {
	if x != nil {
		return x.UpdateToAppContext
	}
	return ""
}

type UpdateAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppContextUpdated       bool   `protobuf:"varint,1,opt,name=app_context_updated,json=appContextUpdated,proto3" json:"app_context_updated,omitempty"`
	AppContextUpdateMessage string `protobuf:"bytes,2,opt,name=app_context_update_message,json=appContextUpdateMessage,proto3" json:"app_context_update_message,omitempty"`
}

func (x *UpdateAppResponse) Reset() {
	*x = UpdateAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updateapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppResponse) ProtoMessage() {}

func (x *UpdateAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_updateapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppResponse.ProtoReflect.Descriptor instead.
func (*UpdateAppResponse) Descriptor() ([]byte, []int) {
	return file_updateapp_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAppResponse) GetAppContextUpdated() bool {
	if x != nil {
		return x.AppContextUpdated
	}
	return false
}

func (x *UpdateAppResponse) GetAppContextUpdateMessage() string {
	if x != nil {
		return x.AppContextUpdateMessage
	}
	return ""
}

type RollbackAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RollbackFromAppContext string `protobuf:"bytes,1,opt,name=rollback_from_appContext,json=rollbackFromAppContext,proto3" json:"rollback_from_appContext,omitempty"`
	RollbackToAppContext   string `protobuf:"bytes,2,opt,name=rollback_to_appContext,json=rollbackToAppContext,proto3" json:"rollback_to_appContext,omitempty"`
}

func (x *RollbackAppRequest) Reset() {
	*x = RollbackAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updateapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackAppRequest) ProtoMessage() {}

func (x *RollbackAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_updateapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackAppRequest.ProtoReflect.Descriptor instead.
func (*RollbackAppRequest) Descriptor() ([]byte, []int) {
	return file_updateapp_proto_rawDescGZIP(), []int{2}
}

func (x *RollbackAppRequest) GetRollbackFromAppContext() string {
	if x != nil {
		return x.RollbackFromAppContext
	}
	return ""
}

func (x *RollbackAppRequest) GetRollbackToAppContext() string {
	if x != nil {
		return x.RollbackToAppContext
	}
	return ""
}

type RollbackAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppContextRolledback      bool   `protobuf:"varint,1,opt,name=app_context_rolledback,json=appContextRolledback,proto3" json:"app_context_rolledback,omitempty"`
	AppContextRollbackMessage string `protobuf:"bytes,2,opt,name=app_context_rollback_message,json=appContextRollbackMessage,proto3" json:"app_context_rollback_message,omitempty"`
}

func (x *RollbackAppResponse) Reset() {
	*x = RollbackAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_updateapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RollbackAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RollbackAppResponse) ProtoMessage() {}

func (x *RollbackAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_updateapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RollbackAppResponse.ProtoReflect.Descriptor instead.
func (*RollbackAppResponse) Descriptor() ([]byte, []int) {
	return file_updateapp_proto_rawDescGZIP(), []int{3}
}

func (x *RollbackAppResponse) GetAppContextRolledback() bool {
	if x != nil {
		return x.AppContextRolledback
	}
	return false
}

func (x *RollbackAppResponse) GetAppContextRollbackMessage() string {
	if x != nil {
		return x.AppContextRollbackMessage
	}
	return ""
}

var File_updateapp_proto protoreflect.FileDescriptor

var file_updateapp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7a, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x80, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x84, 0x01, 0x0a, 0x12, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x70, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x72, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x6f,
	0x5f, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x6f, 0x41, 0x70, 0x70,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x13, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x16, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x14, 0x61, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x6f, 0x6c, 0x6c, 0x65,
	0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x1c, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x61, 0x70, 0x70,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x7d, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x61, 0x70, 0x70, 0x12, 0x34, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x12, 0x11, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x52, 0x6f, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x70, 0x70, 0x12, 0x13, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_updateapp_proto_rawDescOnce sync.Once
	file_updateapp_proto_rawDescData = file_updateapp_proto_rawDesc
)

func file_updateapp_proto_rawDescGZIP() []byte {
	file_updateapp_proto_rawDescOnce.Do(func() {
		file_updateapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_updateapp_proto_rawDescData)
	})
	return file_updateapp_proto_rawDescData
}

var file_updateapp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_updateapp_proto_goTypes = []interface{}{
	(*UpdateAppRequest)(nil),    // 0: UpdateAppRequest
	(*UpdateAppResponse)(nil),   // 1: UpdateAppResponse
	(*RollbackAppRequest)(nil),  // 2: RollbackAppRequest
	(*RollbackAppResponse)(nil), // 3: RollbackAppResponse
}
var file_updateapp_proto_depIdxs = []int32{
	0, // 0: updateapp.UpdateApp:input_type -> UpdateAppRequest
	2, // 1: updateapp.RollbackApp:input_type -> RollbackAppRequest
	1, // 2: updateapp.UpdateApp:output_type -> UpdateAppResponse
	3, // 3: updateapp.RollbackApp:output_type -> RollbackAppResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_updateapp_proto_init() }
func file_updateapp_proto_init() {
	if File_updateapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_updateapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppRequest); i {
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
		file_updateapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppResponse); i {
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
		file_updateapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackAppRequest); i {
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
		file_updateapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RollbackAppResponse); i {
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
			RawDescriptor: file_updateapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_updateapp_proto_goTypes,
		DependencyIndexes: file_updateapp_proto_depIdxs,
		MessageInfos:      file_updateapp_proto_msgTypes,
	}.Build()
	File_updateapp_proto = out.File
	file_updateapp_proto_rawDesc = nil
	file_updateapp_proto_goTypes = nil
	file_updateapp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UpdateappClient is the client API for Updateapp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdateappClient interface {
	// Sync
	UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error)
	RollbackApp(ctx context.Context, in *RollbackAppRequest, opts ...grpc.CallOption) (*RollbackAppResponse, error)
}

type updateappClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateappClient(cc grpc.ClientConnInterface) UpdateappClient {
	return &updateappClient{cc}
}

func (c *updateappClient) UpdateApp(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error) {
	out := new(UpdateAppResponse)
	err := c.cc.Invoke(ctx, "/updateapp/UpdateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateappClient) RollbackApp(ctx context.Context, in *RollbackAppRequest, opts ...grpc.CallOption) (*RollbackAppResponse, error) {
	out := new(RollbackAppResponse)
	err := c.cc.Invoke(ctx, "/updateapp/RollbackApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateappServer is the server API for Updateapp service.
type UpdateappServer interface {
	// Sync
	UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error)
	RollbackApp(context.Context, *RollbackAppRequest) (*RollbackAppResponse, error)
}

// UnimplementedUpdateappServer can be embedded to have forward compatible implementations.
type UnimplementedUpdateappServer struct {
}

func (*UnimplementedUpdateappServer) UpdateApp(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApp not implemented")
}
func (*UnimplementedUpdateappServer) RollbackApp(context.Context, *RollbackAppRequest) (*RollbackAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackApp not implemented")
}

func RegisterUpdateappServer(s *grpc.Server, srv UpdateappServer) {
	s.RegisterService(&_Updateapp_serviceDesc, srv)
}

func _Updateapp_UpdateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateappServer).UpdateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateapp/UpdateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateappServer).UpdateApp(ctx, req.(*UpdateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Updateapp_RollbackApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateappServer).RollbackApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/updateapp/RollbackApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateappServer).RollbackApp(ctx, req.(*RollbackAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Updateapp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "updateapp",
	HandlerType: (*UpdateappServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateApp",
			Handler:    _Updateapp_UpdateApp_Handler,
		},
		{
			MethodName: "RollbackApp",
			Handler:    _Updateapp_RollbackApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "updateapp.proto",
}
