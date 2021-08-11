// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: transport_grpc.proto

package grpc

import (
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

type AddFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *GRPCForm `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *AddFormRequest) Reset() {
	*x = AddFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFormRequest) ProtoMessage() {}

func (x *AddFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFormRequest.ProtoReflect.Descriptor instead.
func (*AddFormRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *AddFormRequest) GetForm() *GRPCForm {
	if x != nil {
		return x.Form
	}
	return nil
}

type AddFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *GRPCForm `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *AddFormResponse) Reset() {
	*x = AddFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFormResponse) ProtoMessage() {}

func (x *AddFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFormResponse.ProtoReflect.Descriptor instead.
func (*AddFormResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *AddFormResponse) GetForm() *GRPCForm {
	if x != nil {
		return x.Form
	}
	return nil
}

type DeleteFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formid string `protobuf:"bytes,1,opt,name=formid,proto3" json:"formid,omitempty"`
}

func (x *DeleteFormRequest) Reset() {
	*x = DeleteFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFormRequest) ProtoMessage() {}

func (x *DeleteFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFormRequest.ProtoReflect.Descriptor instead.
func (*DeleteFormRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFormRequest) GetFormid() string {
	if x != nil {
		return x.Formid
	}
	return ""
}

type DeleteFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFormResponse) Reset() {
	*x = DeleteFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFormResponse) ProtoMessage() {}

func (x *DeleteFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFormResponse.ProtoReflect.Descriptor instead.
func (*DeleteFormResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{3}
}

type GetFormsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFormsRequest) Reset() {
	*x = GetFormsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormsRequest) ProtoMessage() {}

func (x *GetFormsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormsRequest.ProtoReflect.Descriptor instead.
func (*GetFormsRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{4}
}

type GetFormsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forms []*GRPCForm `protobuf:"bytes,1,rep,name=forms,proto3" json:"forms,omitempty"`
}

func (x *GetFormsResponse) Reset() {
	*x = GetFormsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormsResponse) ProtoMessage() {}

func (x *GetFormsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormsResponse.ProtoReflect.Descriptor instead.
func (*GetFormsResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetFormsResponse) GetForms() []*GRPCForm {
	if x != nil {
		return x.Forms
	}
	return nil
}

type GetFormRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formid string `protobuf:"bytes,1,opt,name=formid,proto3" json:"formid,omitempty"`
}

func (x *GetFormRequest) Reset() {
	*x = GetFormRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormRequest) ProtoMessage() {}

func (x *GetFormRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormRequest.ProtoReflect.Descriptor instead.
func (*GetFormRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *GetFormRequest) GetFormid() string {
	if x != nil {
		return x.Formid
	}
	return ""
}

type GetFormResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Form *GRPCForm `protobuf:"bytes,1,opt,name=form,proto3" json:"form,omitempty"`
}

func (x *GetFormResponse) Reset() {
	*x = GetFormResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormResponse) ProtoMessage() {}

func (x *GetFormResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormResponse.ProtoReflect.Descriptor instead.
func (*GetFormResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *GetFormResponse) GetForm() *GRPCForm {
	if x != nil {
		return x.Form
	}
	return nil
}

type SubmitResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formid   string `protobuf:"bytes,1,opt,name=formid,proto3" json:"formid,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SubmitResponseRequest) Reset() {
	*x = SubmitResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponseRequest) ProtoMessage() {}

func (x *SubmitResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponseRequest.ProtoReflect.Descriptor instead.
func (*SubmitResponseRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *SubmitResponseRequest) GetFormid() string {
	if x != nil {
		return x.Formid
	}
	return ""
}

func (x *SubmitResponseRequest) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type SubmitResponseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitResponseResponse) Reset() {
	*x = SubmitResponseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponseResponse) ProtoMessage() {}

func (x *SubmitResponseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponseResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponseResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{9}
}

type GetFormResponsesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formid string `protobuf:"bytes,1,opt,name=formid,proto3" json:"formid,omitempty"`
}

func (x *GetFormResponsesRequest) Reset() {
	*x = GetFormResponsesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormResponsesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormResponsesRequest) ProtoMessage() {}

func (x *GetFormResponsesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormResponsesRequest.ProtoReflect.Descriptor instead.
func (*GetFormResponsesRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{10}
}

func (x *GetFormResponsesRequest) GetFormid() string {
	if x != nil {
		return x.Formid
	}
	return ""
}

type GetFormResponsesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []string `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *GetFormResponsesResponse) Reset() {
	*x = GetFormResponsesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormResponsesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormResponsesResponse) ProtoMessage() {}

func (x *GetFormResponsesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormResponsesResponse.ProtoReflect.Descriptor instead.
func (*GetFormResponsesResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{11}
}

func (x *GetFormResponsesResponse) GetResponses() []string {
	if x != nil {
		return x.Responses
	}
	return nil
}

type AddFormBackendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formid  string       `protobuf:"bytes,1,opt,name=formid,proto3" json:"formid,omitempty"`
	Backend *GRPCBackend `protobuf:"bytes,2,opt,name=backend,proto3" json:"backend,omitempty"`
}

func (x *AddFormBackendRequest) Reset() {
	*x = AddFormBackendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFormBackendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFormBackendRequest) ProtoMessage() {}

func (x *AddFormBackendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFormBackendRequest.ProtoReflect.Descriptor instead.
func (*AddFormBackendRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{12}
}

func (x *AddFormBackendRequest) GetFormid() string {
	if x != nil {
		return x.Formid
	}
	return ""
}

func (x *AddFormBackendRequest) GetBackend() *GRPCBackend {
	if x != nil {
		return x.Backend
	}
	return nil
}

type AddFormBackendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddFormBackendResponse) Reset() {
	*x = AddFormBackendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFormBackendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFormBackendResponse) ProtoMessage() {}

func (x *AddFormBackendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFormBackendResponse.ProtoReflect.Descriptor instead.
func (*AddFormBackendResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{13}
}

type GetFormBackendsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Formid string `protobuf:"bytes,1,opt,name=formid,proto3" json:"formid,omitempty"`
}

func (x *GetFormBackendsRequest) Reset() {
	*x = GetFormBackendsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormBackendsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormBackendsRequest) ProtoMessage() {}

func (x *GetFormBackendsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormBackendsRequest.ProtoReflect.Descriptor instead.
func (*GetFormBackendsRequest) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{14}
}

func (x *GetFormBackendsRequest) GetFormid() string {
	if x != nil {
		return x.Formid
	}
	return ""
}

type GetFormBackendsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backends []*GRPCBackend `protobuf:"bytes,1,rep,name=backends,proto3" json:"backends,omitempty"`
}

func (x *GetFormBackendsResponse) Reset() {
	*x = GetFormBackendsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_grpc_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFormBackendsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFormBackendsResponse) ProtoMessage() {}

func (x *GetFormBackendsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_grpc_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFormBackendsResponse.ProtoReflect.Descriptor instead.
func (*GetFormBackendsResponse) Descriptor() ([]byte, []int) {
	return file_transport_grpc_proto_rawDescGZIP(), []int{15}
}

func (x *GetFormBackendsResponse) GetBackends() []*GRPCBackend {
	if x != nil {
		return x.Backends
	}
	return nil
}

var File_transport_grpc_proto protoreflect.FileDescriptor

var file_transport_grpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04,
	0x66, 0x6f, 0x72, 0x6d, 0x22, 0x30, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x46, 0x6f, 0x72, 0x6d,
	0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x6d,
	0x73, 0x22, 0x28, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x4b, 0x0a,
	0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x73, 0x22, 0x57, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x41, 0x64,
	0x64, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72,
	0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x52, 0x50, 0x43, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x32, 0xf7, 0x03, 0x0a, 0x0c,
	0x46, 0x6f, 0x72, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f,
	0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d,
	0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x72, 0x6d, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x73, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46,
	0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x41, 0x64, 0x64,
	0x46, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73,
	0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x72, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_grpc_proto_rawDescOnce sync.Once
	file_transport_grpc_proto_rawDescData = file_transport_grpc_proto_rawDesc
)

func file_transport_grpc_proto_rawDescGZIP() []byte {
	file_transport_grpc_proto_rawDescOnce.Do(func() {
		file_transport_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_grpc_proto_rawDescData)
	})
	return file_transport_grpc_proto_rawDescData
}

var file_transport_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_transport_grpc_proto_goTypes = []interface{}{
	(*AddFormRequest)(nil),           // 0: AddFormRequest
	(*AddFormResponse)(nil),          // 1: AddFormResponse
	(*DeleteFormRequest)(nil),        // 2: DeleteFormRequest
	(*DeleteFormResponse)(nil),       // 3: DeleteFormResponse
	(*GetFormsRequest)(nil),          // 4: GetFormsRequest
	(*GetFormsResponse)(nil),         // 5: GetFormsResponse
	(*GetFormRequest)(nil),           // 6: GetFormRequest
	(*GetFormResponse)(nil),          // 7: GetFormResponse
	(*SubmitResponseRequest)(nil),    // 8: SubmitResponseRequest
	(*SubmitResponseResponse)(nil),   // 9: SubmitResponseResponse
	(*GetFormResponsesRequest)(nil),  // 10: GetFormResponsesRequest
	(*GetFormResponsesResponse)(nil), // 11: GetFormResponsesResponse
	(*AddFormBackendRequest)(nil),    // 12: AddFormBackendRequest
	(*AddFormBackendResponse)(nil),   // 13: AddFormBackendResponse
	(*GetFormBackendsRequest)(nil),   // 14: GetFormBackendsRequest
	(*GetFormBackendsResponse)(nil),  // 15: GetFormBackendsResponse
	(*GRPCForm)(nil),                 // 16: GRPCForm
	(*GRPCBackend)(nil),              // 17: GRPCBackend
}
var file_transport_grpc_proto_depIdxs = []int32{
	16, // 0: AddFormRequest.form:type_name -> GRPCForm
	16, // 1: AddFormResponse.form:type_name -> GRPCForm
	16, // 2: GetFormsResponse.forms:type_name -> GRPCForm
	16, // 3: GetFormResponse.form:type_name -> GRPCForm
	17, // 4: AddFormBackendRequest.backend:type_name -> GRPCBackend
	17, // 5: GetFormBackendsResponse.backends:type_name -> GRPCBackend
	0,  // 6: FormsService.AddForm:input_type -> AddFormRequest
	2,  // 7: FormsService.DeleteForm:input_type -> DeleteFormRequest
	4,  // 8: FormsService.GetForms:input_type -> GetFormsRequest
	6,  // 9: FormsService.GetForm:input_type -> GetFormRequest
	8,  // 10: FormsService.SubmitResponse:input_type -> SubmitResponseRequest
	10, // 11: FormsService.GetFormResponses:input_type -> GetFormResponsesRequest
	12, // 12: FormsService.AddFormBackend:input_type -> AddFormBackendRequest
	14, // 13: FormsService.GetFormBackends:input_type -> GetFormBackendsRequest
	1,  // 14: FormsService.AddForm:output_type -> AddFormResponse
	3,  // 15: FormsService.DeleteForm:output_type -> DeleteFormResponse
	5,  // 16: FormsService.GetForms:output_type -> GetFormsResponse
	7,  // 17: FormsService.GetForm:output_type -> GetFormResponse
	9,  // 18: FormsService.SubmitResponse:output_type -> SubmitResponseResponse
	11, // 19: FormsService.GetFormResponses:output_type -> GetFormResponsesResponse
	13, // 20: FormsService.AddFormBackend:output_type -> AddFormBackendResponse
	15, // 21: FormsService.GetFormBackends:output_type -> GetFormBackendsResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_transport_grpc_proto_init() }
func file_transport_grpc_proto_init() {
	if File_transport_grpc_proto != nil {
		return
	}
	file_forms_proto_init()
	file_backends_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transport_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFormRequest); i {
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
		file_transport_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFormResponse); i {
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
		file_transport_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFormRequest); i {
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
		file_transport_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFormResponse); i {
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
		file_transport_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormsRequest); i {
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
		file_transport_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormsResponse); i {
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
		file_transport_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormRequest); i {
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
		file_transport_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormResponse); i {
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
		file_transport_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponseRequest); i {
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
		file_transport_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponseResponse); i {
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
		file_transport_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormResponsesRequest); i {
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
		file_transport_grpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormResponsesResponse); i {
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
		file_transport_grpc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFormBackendRequest); i {
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
		file_transport_grpc_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFormBackendResponse); i {
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
		file_transport_grpc_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormBackendsRequest); i {
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
		file_transport_grpc_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFormBackendsResponse); i {
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
			RawDescriptor: file_transport_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_grpc_proto_goTypes,
		DependencyIndexes: file_transport_grpc_proto_depIdxs,
		MessageInfos:      file_transport_grpc_proto_msgTypes,
	}.Build()
	File_transport_grpc_proto = out.File
	file_transport_grpc_proto_rawDesc = nil
	file_transport_grpc_proto_goTypes = nil
	file_transport_grpc_proto_depIdxs = nil
}
