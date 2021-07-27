// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.2
// source: user_role_service.proto

package auth_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type UserRoleId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRoleId) Reset() {
	*x = UserRoleId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_role_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleId) ProtoMessage() {}

func (x *UserRoleId) ProtoReflect() protoreflect.Message {
	mi := &file_user_role_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleId.ProtoReflect.Descriptor instead.
func (*UserRoleId) Descriptor() ([]byte, []int) {
	return file_user_role_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserRoleId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllUserRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId  string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	UserTypeId string `protobuf:"bytes,2,opt,name=user_type_id,json=userTypeId,proto3" json:"user_type_id,omitempty"`
	Page       uint64 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit      uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllUserRolesRequest) Reset() {
	*x = GetAllUserRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_role_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserRolesRequest) ProtoMessage() {}

func (x *GetAllUserRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_role_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserRolesRequest.ProtoReflect.Descriptor instead.
func (*GetAllUserRolesRequest) Descriptor() ([]byte, []int) {
	return file_user_role_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllUserRolesRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *GetAllUserRolesRequest) GetUserTypeId() string {
	if x != nil {
		return x.UserTypeId
	}
	return ""
}

func (x *GetAllUserRolesRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllUserRolesRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllUserRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRoles []*UserRole `protobuf:"bytes,1,rep,name=user_roles,json=userRoles,proto3" json:"user_roles,omitempty"`
	Count     uint64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllUserRolesResponse) Reset() {
	*x = GetAllUserRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_role_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserRolesResponse) ProtoMessage() {}

func (x *GetAllUserRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_role_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserRolesResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserRolesResponse) Descriptor() ([]byte, []int) {
	return file_user_role_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllUserRolesResponse) GetUserRoles() []*UserRole {
	if x != nil {
		return x.UserRoles
	}
	return nil
}

func (x *GetAllUserRolesResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SavePermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRoleId  string                               `protobuf:"bytes,1,opt,name=user_role_id,json=userRoleId,proto3" json:"user_role_id,omitempty"`
	Permissions []*SavePermissionsRequest_Permission `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *SavePermissionsRequest) Reset() {
	*x = SavePermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_role_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePermissionsRequest) ProtoMessage() {}

func (x *SavePermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_role_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePermissionsRequest.ProtoReflect.Descriptor instead.
func (*SavePermissionsRequest) Descriptor() ([]byte, []int) {
	return file_user_role_service_proto_rawDescGZIP(), []int{3}
}

func (x *SavePermissionsRequest) GetUserRoleId() string {
	if x != nil {
		return x.UserRoleId
	}
	return ""
}

func (x *SavePermissionsRequest) GetPermissions() []*SavePermissionsRequest_Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type GetUserRolePermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *GetUserRolePermissionsResponse) Reset() {
	*x = GetUserRolePermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_role_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRolePermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRolePermissionsResponse) ProtoMessage() {}

func (x *GetUserRolePermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_role_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRolePermissionsResponse.ProtoReflect.Descriptor instead.
func (*GetUserRolePermissionsResponse) Descriptor() ([]byte, []int) {
	return file_user_role_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserRolePermissionsResponse) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type SavePermissionsRequest_Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ActionIds []string `protobuf:"bytes,2,rep,name=action_ids,json=actionIds,proto3" json:"action_ids,omitempty"`
}

func (x *SavePermissionsRequest_Permission) Reset() {
	*x = SavePermissionsRequest_Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_role_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePermissionsRequest_Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePermissionsRequest_Permission) ProtoMessage() {}

func (x *SavePermissionsRequest_Permission) ProtoReflect() protoreflect.Message {
	mi := &file_user_role_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePermissionsRequest_Permission.ProtoReflect.Descriptor instead.
func (*SavePermissionsRequest_Permission) Descriptor() ([]byte, []int) {
	return file_user_role_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SavePermissionsRequest_Permission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SavePermissionsRequest_Permission) GetActionIds() []string {
	if x != nil {
		return x.ActionIds
	}
	return nil
}

var File_user_role_service_proto protoreflect.FileDescriptor

var file_user_role_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x0a,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x62, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x4d, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x3b, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x58, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe0, 0x03, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x1a, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x12,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x20,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x1a, 0x28,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_role_service_proto_rawDescOnce sync.Once
	file_user_role_service_proto_rawDescData = file_user_role_service_proto_rawDesc
)

func file_user_role_service_proto_rawDescGZIP() []byte {
	file_user_role_service_proto_rawDescOnce.Do(func() {
		file_user_role_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_role_service_proto_rawDescData)
	})
	return file_user_role_service_proto_rawDescData
}

var file_user_role_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_role_service_proto_goTypes = []interface{}{
	(*UserRoleId)(nil),                        // 0: genproto.UserRoleId
	(*GetAllUserRolesRequest)(nil),            // 1: genproto.GetAllUserRolesRequest
	(*GetAllUserRolesResponse)(nil),           // 2: genproto.GetAllUserRolesResponse
	(*SavePermissionsRequest)(nil),            // 3: genproto.SavePermissionsRequest
	(*GetUserRolePermissionsResponse)(nil),    // 4: genproto.GetUserRolePermissionsResponse
	(*SavePermissionsRequest_Permission)(nil), // 5: genproto.SavePermissionsRequest.Permission
	(*UserRole)(nil),                          // 6: genproto.UserRole
	(*Permission)(nil),                        // 7: genproto.Permission
	(*empty.Empty)(nil),                       // 8: google.protobuf.Empty
}
var file_user_role_service_proto_depIdxs = []int32{
	6,  // 0: genproto.GetAllUserRolesResponse.user_roles:type_name -> genproto.UserRole
	5,  // 1: genproto.SavePermissionsRequest.permissions:type_name -> genproto.SavePermissionsRequest.Permission
	7,  // 2: genproto.GetUserRolePermissionsResponse.permissions:type_name -> genproto.Permission
	6,  // 3: genproto.UserRoleService.Create:input_type -> genproto.UserRole
	0,  // 4: genproto.UserRoleService.Get:input_type -> genproto.UserRoleId
	1,  // 5: genproto.UserRoleService.GetAll:input_type -> genproto.GetAllUserRolesRequest
	6,  // 6: genproto.UserRoleService.Update:input_type -> genproto.UserRole
	0,  // 7: genproto.UserRoleService.Delete:input_type -> genproto.UserRoleId
	3,  // 8: genproto.UserRoleService.SavePermissions:input_type -> genproto.SavePermissionsRequest
	0,  // 9: genproto.UserRoleService.GetPermissions:input_type -> genproto.UserRoleId
	0,  // 10: genproto.UserRoleService.Create:output_type -> genproto.UserRoleId
	6,  // 11: genproto.UserRoleService.Get:output_type -> genproto.UserRole
	2,  // 12: genproto.UserRoleService.GetAll:output_type -> genproto.GetAllUserRolesResponse
	8,  // 13: genproto.UserRoleService.Update:output_type -> google.protobuf.Empty
	8,  // 14: genproto.UserRoleService.Delete:output_type -> google.protobuf.Empty
	8,  // 15: genproto.UserRoleService.SavePermissions:output_type -> google.protobuf.Empty
	4,  // 16: genproto.UserRoleService.GetPermissions:output_type -> genproto.GetUserRolePermissionsResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_user_role_service_proto_init() }
func file_user_role_service_proto_init() {
	if File_user_role_service_proto != nil {
		return
	}
	file_auth_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_role_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleId); i {
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
		file_user_role_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserRolesRequest); i {
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
		file_user_role_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserRolesResponse); i {
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
		file_user_role_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePermissionsRequest); i {
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
		file_user_role_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRolePermissionsResponse); i {
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
		file_user_role_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePermissionsRequest_Permission); i {
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
			RawDescriptor: file_user_role_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_role_service_proto_goTypes,
		DependencyIndexes: file_user_role_service_proto_depIdxs,
		MessageInfos:      file_user_role_service_proto_msgTypes,
	}.Build()
	File_user_role_service_proto = out.File
	file_user_role_service_proto_rawDesc = nil
	file_user_role_service_proto_goTypes = nil
	file_user_role_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserRoleServiceClient is the client API for UserRoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserRoleServiceClient interface {
	Create(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*UserRoleId, error)
	Get(ctx context.Context, in *UserRoleId, opts ...grpc.CallOption) (*UserRole, error)
	GetAll(ctx context.Context, in *GetAllUserRolesRequest, opts ...grpc.CallOption) (*GetAllUserRolesResponse, error)
	Update(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *UserRoleId, opts ...grpc.CallOption) (*empty.Empty, error)
	SavePermissions(ctx context.Context, in *SavePermissionsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetPermissions(ctx context.Context, in *UserRoleId, opts ...grpc.CallOption) (*GetUserRolePermissionsResponse, error)
}

type userRoleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRoleServiceClient(cc grpc.ClientConnInterface) UserRoleServiceClient {
	return &userRoleServiceClient{cc}
}

func (c *userRoleServiceClient) Create(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*UserRoleId, error) {
	out := new(UserRoleId)
	err := c.cc.Invoke(ctx, "/genproto.UserRoleService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) Get(ctx context.Context, in *UserRoleId, opts ...grpc.CallOption) (*UserRole, error) {
	out := new(UserRole)
	err := c.cc.Invoke(ctx, "/genproto.UserRoleService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) GetAll(ctx context.Context, in *GetAllUserRolesRequest, opts ...grpc.CallOption) (*GetAllUserRolesResponse, error) {
	out := new(GetAllUserRolesResponse)
	err := c.cc.Invoke(ctx, "/genproto.UserRoleService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) Update(ctx context.Context, in *UserRole, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.UserRoleService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) Delete(ctx context.Context, in *UserRoleId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.UserRoleService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) SavePermissions(ctx context.Context, in *SavePermissionsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.UserRoleService/SavePermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) GetPermissions(ctx context.Context, in *UserRoleId, opts ...grpc.CallOption) (*GetUserRolePermissionsResponse, error) {
	out := new(GetUserRolePermissionsResponse)
	err := c.cc.Invoke(ctx, "/genproto.UserRoleService/GetPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRoleServiceServer is the server API for UserRoleService service.
type UserRoleServiceServer interface {
	Create(context.Context, *UserRole) (*UserRoleId, error)
	Get(context.Context, *UserRoleId) (*UserRole, error)
	GetAll(context.Context, *GetAllUserRolesRequest) (*GetAllUserRolesResponse, error)
	Update(context.Context, *UserRole) (*empty.Empty, error)
	Delete(context.Context, *UserRoleId) (*empty.Empty, error)
	SavePermissions(context.Context, *SavePermissionsRequest) (*empty.Empty, error)
	GetPermissions(context.Context, *UserRoleId) (*GetUserRolePermissionsResponse, error)
}

// UnimplementedUserRoleServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserRoleServiceServer struct {
}

func (*UnimplementedUserRoleServiceServer) Create(context.Context, *UserRole) (*UserRoleId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserRoleServiceServer) Get(context.Context, *UserRoleId) (*UserRole, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserRoleServiceServer) GetAll(context.Context, *GetAllUserRolesRequest) (*GetAllUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedUserRoleServiceServer) Update(context.Context, *UserRole) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserRoleServiceServer) Delete(context.Context, *UserRoleId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedUserRoleServiceServer) SavePermissions(context.Context, *SavePermissionsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePermissions not implemented")
}
func (*UnimplementedUserRoleServiceServer) GetPermissions(context.Context, *UserRoleId) (*GetUserRolePermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissions not implemented")
}

func RegisterUserRoleServiceServer(s *grpc.Server, srv UserRoleServiceServer) {
	s.RegisterService(&_UserRoleService_serviceDesc, srv)
}

func _UserRoleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserRoleService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).Create(ctx, req.(*UserRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserRoleService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).Get(ctx, req.(*UserRoleId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserRoleService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).GetAll(ctx, req.(*GetAllUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserRoleService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).Update(ctx, req.(*UserRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserRoleService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).Delete(ctx, req.(*UserRoleId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_SavePermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).SavePermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserRoleService/SavePermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).SavePermissions(ctx, req.(*SavePermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_GetPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRoleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).GetPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.UserRoleService/GetPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).GetPermissions(ctx, req.(*UserRoleId))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.UserRoleService",
	HandlerType: (*UserRoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserRoleService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserRoleService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserRoleService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserRoleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserRoleService_Delete_Handler,
		},
		{
			MethodName: "SavePermissions",
			Handler:    _UserRoleService_SavePermissions_Handler,
		},
		{
			MethodName: "GetPermissions",
			Handler:    _UserRoleService_GetPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_role_service.proto",
}