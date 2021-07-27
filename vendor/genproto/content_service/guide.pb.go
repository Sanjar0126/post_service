// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.2
// source: guide.proto

package content_service

import (
	context "context"
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

type GetAllGuidesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guides []*Guide `protobuf:"bytes,1,rep,name=guides,proto3" json:"guides,omitempty"`
}

func (x *GetAllGuidesResponse) Reset() {
	*x = GetAllGuidesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllGuidesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllGuidesResponse) ProtoMessage() {}

func (x *GetAllGuidesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllGuidesResponse.ProtoReflect.Descriptor instead.
func (*GetAllGuidesResponse) Descriptor() ([]byte, []int) {
	return file_guide_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllGuidesResponse) GetGuides() []*Guide {
	if x != nil {
		return x.Guides
	}
	return nil
}

type GuideId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GuideId) Reset() {
	*x = GuideId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuideId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuideId) ProtoMessage() {}

func (x *GuideId) ProtoReflect() protoreflect.Message {
	mi := &file_guide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuideId.ProtoReflect.Descriptor instead.
func (*GuideId) Descriptor() ([]byte, []int) {
	return file_guide_proto_rawDescGZIP(), []int{1}
}

func (x *GuideId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Guide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        *Language            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Instructions []*Guide_Instruction `protobuf:"bytes,3,rep,name=instructions,proto3" json:"instructions,omitempty"`
	CreatedAt    string               `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string               `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Guide) Reset() {
	*x = Guide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guide) ProtoMessage() {}

func (x *Guide) ProtoReflect() protoreflect.Message {
	mi := &file_guide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guide.ProtoReflect.Descriptor instead.
func (*Guide) Descriptor() ([]byte, []int) {
	return file_guide_proto_rawDescGZIP(), []int{2}
}

func (x *Guide) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Guide) GetTitle() *Language {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Guide) GetInstructions() []*Guide_Instruction {
	if x != nil {
		return x.Instructions
	}
	return nil
}

func (x *Guide) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Guide) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Guide_Instruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     *Language `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body      *Language `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	GuideId   string    `protobuf:"bytes,4,opt,name=guide_id,json=guideId,proto3" json:"guide_id,omitempty"`
	CreatedAt string    `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string    `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Guide_Instruction) Reset() {
	*x = Guide_Instruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guide_Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guide_Instruction) ProtoMessage() {}

func (x *Guide_Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_guide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guide_Instruction.ProtoReflect.Descriptor instead.
func (*Guide_Instruction) Descriptor() ([]byte, []int) {
	return file_guide_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Guide_Instruction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Guide_Instruction) GetTitle() *Language {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Guide_Instruction) GetBody() *Language {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Guide_Instruction) GetGuideId() string {
	if x != nil {
		return x.GuideId
	}
	return ""
}

func (x *Guide_Instruction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Guide_Instruction) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_guide_proto protoreflect.FileDescriptor

var file_guide_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47,
	0x75, 0x69, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x06, 0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x75, 0x69, 0x64, 0x65, 0x52, 0x06,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x47, 0x75, 0x69, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x8b, 0x03, 0x0a, 0x05, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x1a, 0xc8, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x75, 0x69, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32,
	0x86, 0x02, 0x0a, 0x0c, 0x47, 0x75, 0x69, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x75, 0x69, 0x64, 0x65, 0x1a, 0x11, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x75, 0x69, 0x64, 0x65, 0x49, 0x64, 0x22, 0x00,
	0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x75, 0x69, 0x64, 0x65, 0x1a, 0x0f, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x47, 0x75, 0x69, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x75,
	0x69, 0x64, 0x65, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x75, 0x69, 0x64, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x75,
	0x69, 0x64, 0x65, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_guide_proto_rawDescOnce sync.Once
	file_guide_proto_rawDescData = file_guide_proto_rawDesc
)

func file_guide_proto_rawDescGZIP() []byte {
	file_guide_proto_rawDescOnce.Do(func() {
		file_guide_proto_rawDescData = protoimpl.X.CompressGZIP(file_guide_proto_rawDescData)
	})
	return file_guide_proto_rawDescData
}

var file_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_guide_proto_goTypes = []interface{}{
	(*GetAllGuidesResponse)(nil), // 0: genproto.GetAllGuidesResponse
	(*GuideId)(nil),              // 1: genproto.GuideId
	(*Guide)(nil),                // 2: genproto.Guide
	(*Guide_Instruction)(nil),    // 3: genproto.Guide.Instruction
	(*Language)(nil),             // 4: genproto.Language
	(*Empty)(nil),                // 5: genproto.Empty
}
var file_guide_proto_depIdxs = []int32{
	2,  // 0: genproto.GetAllGuidesResponse.guides:type_name -> genproto.Guide
	4,  // 1: genproto.Guide.title:type_name -> genproto.Language
	3,  // 2: genproto.Guide.instructions:type_name -> genproto.Guide.Instruction
	4,  // 3: genproto.Guide.Instruction.title:type_name -> genproto.Language
	4,  // 4: genproto.Guide.Instruction.body:type_name -> genproto.Language
	2,  // 5: genproto.GuideService.Create:input_type -> genproto.Guide
	2,  // 6: genproto.GuideService.Update:input_type -> genproto.Guide
	5,  // 7: genproto.GuideService.GetAll:input_type -> genproto.Empty
	1,  // 8: genproto.GuideService.Get:input_type -> genproto.GuideId
	1,  // 9: genproto.GuideService.Delete:input_type -> genproto.GuideId
	1,  // 10: genproto.GuideService.Create:output_type -> genproto.GuideId
	5,  // 11: genproto.GuideService.Update:output_type -> genproto.Empty
	0,  // 12: genproto.GuideService.GetAll:output_type -> genproto.GetAllGuidesResponse
	2,  // 13: genproto.GuideService.Get:output_type -> genproto.Guide
	5,  // 14: genproto.GuideService.Delete:output_type -> genproto.Empty
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_guide_proto_init() }
func file_guide_proto_init() {
	if File_guide_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_guide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllGuidesResponse); i {
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
		file_guide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuideId); i {
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
		file_guide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guide); i {
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
		file_guide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guide_Instruction); i {
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
			RawDescriptor: file_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_guide_proto_goTypes,
		DependencyIndexes: file_guide_proto_depIdxs,
		MessageInfos:      file_guide_proto_msgTypes,
	}.Build()
	File_guide_proto = out.File
	file_guide_proto_rawDesc = nil
	file_guide_proto_goTypes = nil
	file_guide_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GuideServiceClient is the client API for GuideService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuideServiceClient interface {
	Create(ctx context.Context, in *Guide, opts ...grpc.CallOption) (*GuideId, error)
	Update(ctx context.Context, in *Guide, opts ...grpc.CallOption) (*Empty, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllGuidesResponse, error)
	Get(ctx context.Context, in *GuideId, opts ...grpc.CallOption) (*Guide, error)
	Delete(ctx context.Context, in *GuideId, opts ...grpc.CallOption) (*Empty, error)
}

type guideServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuideServiceClient(cc grpc.ClientConnInterface) GuideServiceClient {
	return &guideServiceClient{cc}
}

func (c *guideServiceClient) Create(ctx context.Context, in *Guide, opts ...grpc.CallOption) (*GuideId, error) {
	out := new(GuideId)
	err := c.cc.Invoke(ctx, "/genproto.GuideService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) Update(ctx context.Context, in *Guide, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/genproto.GuideService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllGuidesResponse, error) {
	out := new(GetAllGuidesResponse)
	err := c.cc.Invoke(ctx, "/genproto.GuideService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) Get(ctx context.Context, in *GuideId, opts ...grpc.CallOption) (*Guide, error) {
	out := new(Guide)
	err := c.cc.Invoke(ctx, "/genproto.GuideService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) Delete(ctx context.Context, in *GuideId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/genproto.GuideService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuideServiceServer is the server API for GuideService service.
type GuideServiceServer interface {
	Create(context.Context, *Guide) (*GuideId, error)
	Update(context.Context, *Guide) (*Empty, error)
	GetAll(context.Context, *Empty) (*GetAllGuidesResponse, error)
	Get(context.Context, *GuideId) (*Guide, error)
	Delete(context.Context, *GuideId) (*Empty, error)
}

// UnimplementedGuideServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGuideServiceServer struct {
}

func (*UnimplementedGuideServiceServer) Create(context.Context, *Guide) (*GuideId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGuideServiceServer) Update(context.Context, *Guide) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedGuideServiceServer) GetAll(context.Context, *Empty) (*GetAllGuidesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedGuideServiceServer) Get(context.Context, *GuideId) (*Guide, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedGuideServiceServer) Delete(context.Context, *GuideId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterGuideServiceServer(s *grpc.Server, srv GuideServiceServer) {
	s.RegisterService(&_GuideService_serviceDesc, srv)
}

func _GuideService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Guide)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GuideService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).Create(ctx, req.(*Guide))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Guide)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GuideService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).Update(ctx, req.(*Guide))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GuideService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuideId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GuideService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).Get(ctx, req.(*GuideId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuideId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.GuideService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).Delete(ctx, req.(*GuideId))
	}
	return interceptor(ctx, in, info, handler)
}

var _GuideService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.GuideService",
	HandlerType: (*GuideServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GuideService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GuideService_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _GuideService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GuideService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GuideService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guide.proto",
}
