// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: ticket.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32    `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	AssignedTo  int32    `protobuf:"varint,3,opt,name=AssignedTo,proto3" json:"AssignedTo,omitempty"`
	Priority    int32    `protobuf:"varint,4,opt,name=Priority,proto3" json:"Priority,omitempty"`
	Status      string   `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Timestamp   string   `protobuf:"bytes,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	CreatedBy   int32    `protobuf:"varint,9,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Labels      []string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *TicketRequest) Reset() {
	*x = TicketRequest{}
	mi := &file_ticket_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketRequest) ProtoMessage() {}

func (x *TicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketRequest.ProtoReflect.Descriptor instead.
func (*TicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *TicketRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TicketRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TicketRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TicketRequest) GetAssignedTo() int32 {
	if x != nil {
		return x.AssignedTo
	}
	return 0
}

func (x *TicketRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *TicketRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TicketRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *TicketRequest) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *TicketRequest) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type TicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	AssignedTo  int32    `protobuf:"varint,3,opt,name=AssignedTo,proto3" json:"AssignedTo,omitempty"`
	Priority    int32    `protobuf:"varint,4,opt,name=Priority,proto3" json:"Priority,omitempty"`
	Status      string   `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Timestamp   string   `protobuf:"bytes,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	ID          int32    `protobuf:"varint,7,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedBy   int32    `protobuf:"varint,9,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Labels      []string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *TicketResponse) Reset() {
	*x = TicketResponse{}
	mi := &file_ticket_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketResponse) ProtoMessage() {}

func (x *TicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketResponse.ProtoReflect.Descriptor instead.
func (*TicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *TicketResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TicketResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TicketResponse) GetAssignedTo() int32 {
	if x != nil {
		return x.AssignedTo
	}
	return 0
}

func (x *TicketResponse) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *TicketResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TicketResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *TicketResponse) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TicketResponse) GetCreatedBy() int32 {
	if x != nil {
		return x.CreatedBy
	}
	return 0
}

func (x *TicketResponse) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type TicketID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *TicketID) Reset() {
	*x = TicketID{}
	mi := &file_ticket_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketID) ProtoMessage() {}

func (x *TicketID) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketID.ProtoReflect.Descriptor instead.
func (*TicketID) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *TicketID) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type TicketLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Label  string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *TicketLabel) Reset() {
	*x = TicketLabel{}
	mi := &file_ticket_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TicketLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketLabel) ProtoMessage() {}

func (x *TicketLabel) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketLabel.ProtoReflect.Descriptor instead.
func (*TicketLabel) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *TicketLabel) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TicketLabel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type AssignTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	ID     int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AssignTicket) Reset() {
	*x = AssignTicket{}
	mi := &file_ticket_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignTicket) ProtoMessage() {}

func (x *AssignTicket) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignTicket.ProtoReflect.Descriptor instead.
func (*AssignTicket) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *AssignTicket) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *AssignTicket) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ListTicketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tickets []*TicketResponse `protobuf:"bytes,1,rep,name=tickets,proto3" json:"tickets,omitempty"`
}

func (x *ListTicketResponse) Reset() {
	*x = ListTicketResponse{}
	mi := &file_ticket_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTicketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTicketResponse) ProtoMessage() {}

func (x *ListTicketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTicketResponse.ProtoReflect.Descriptor instead.
func (*ListTicketResponse) Descriptor() ([]byte, []int) {
	return file_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *ListTicketResponse) GetTickets() []*TicketResponse {
	if x != nil {
		return x.Tickets
	}
	return nil
}

var File_ticket_proto protoreflect.FileDescriptor

var file_ticket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a, 0x0d, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x54, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x1a, 0x0a, 0x08, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x0b, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x22, 0x36, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x32, 0xf2, 0x04, 0x0a, 0x0d,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1c, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44,
	0x1a, 0x1d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44,
	0x1a, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52, 0x0a, 0x11, 0x67, 0x65,
	0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x21, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55,
	0x0a, 0x14, 0x67, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x42, 0x79, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x1a, 0x21, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x15, 0x5a, 0x13, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_proto_rawDescOnce sync.Once
	file_ticket_proto_rawDescData = file_ticket_proto_rawDesc
)

func file_ticket_proto_rawDescGZIP() []byte {
	file_ticket_proto_rawDescOnce.Do(func() {
		file_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_proto_rawDescData)
	})
	return file_ticket_proto_rawDescData
}

var file_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ticket_proto_goTypes = []any{
	(*TicketRequest)(nil),      // 0: ticketService.TicketRequest
	(*TicketResponse)(nil),     // 1: ticketService.TicketResponse
	(*TicketID)(nil),           // 2: ticketService.TicketID
	(*TicketLabel)(nil),        // 3: ticketService.TicketLabel
	(*AssignTicket)(nil),       // 4: ticketService.AssignTicket
	(*ListTicketResponse)(nil), // 5: ticketService.ListTicketResponse
	(*emptypb.Empty)(nil),      // 6: google.protobuf.Empty
}
var file_ticket_proto_depIdxs = []int32{
	1, // 0: ticketService.ListTicketResponse.tickets:type_name -> ticketService.TicketResponse
	0, // 1: ticketService.ticketService.createTicket:input_type -> ticketService.TicketRequest
	2, // 2: ticketService.ticketService.GetTicket:input_type -> ticketService.TicketID
	2, // 3: ticketService.ticketService.GetAllTicketsForUser:input_type -> ticketService.TicketID
	0, // 4: ticketService.ticketService.updateTicket:input_type -> ticketService.TicketRequest
	2, // 5: ticketService.ticketService.deleteTicket:input_type -> ticketService.TicketID
	3, // 6: ticketService.ticketService.getTicketsByLabel:input_type -> ticketService.TicketLabel
	3, // 7: ticketService.ticketService.getTicketsByPriority:input_type -> ticketService.TicketLabel
	4, // 8: ticketService.ticketService.assignTicketToUser:input_type -> ticketService.AssignTicket
	1, // 9: ticketService.ticketService.createTicket:output_type -> ticketService.TicketResponse
	1, // 10: ticketService.ticketService.GetTicket:output_type -> ticketService.TicketResponse
	5, // 11: ticketService.ticketService.GetAllTicketsForUser:output_type -> ticketService.ListTicketResponse
	6, // 12: ticketService.ticketService.updateTicket:output_type -> google.protobuf.Empty
	6, // 13: ticketService.ticketService.deleteTicket:output_type -> google.protobuf.Empty
	5, // 14: ticketService.ticketService.getTicketsByLabel:output_type -> ticketService.ListTicketResponse
	5, // 15: ticketService.ticketService.getTicketsByPriority:output_type -> ticketService.ListTicketResponse
	6, // 16: ticketService.ticketService.assignTicketToUser:output_type -> google.protobuf.Empty
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ticket_proto_init() }
func file_ticket_proto_init() {
	if File_ticket_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_proto_goTypes,
		DependencyIndexes: file_ticket_proto_depIdxs,
		MessageInfos:      file_ticket_proto_msgTypes,
	}.Build()
	File_ticket_proto = out.File
	file_ticket_proto_rawDesc = nil
	file_ticket_proto_goTypes = nil
	file_ticket_proto_depIdxs = nil
}
