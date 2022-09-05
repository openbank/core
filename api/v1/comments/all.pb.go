// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: bnk.to/core/api/v1/comments/all.proto

// Package comments allows creating and retrieving comments.

package comments

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/xo/ecosystem/proto/xo"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Owner is the type of an entity owning the comment.
type Owner int32

const (
	Owner_Client         Owner = 0
	Owner_Group          Owner = 1
	Owner_LoanProduct    Owner = 2
	Owner_SavingsProduct Owner = 3
	Owner_Centre         Owner = 4
	Owner_Branch         Owner = 5
	Owner_User           Owner = 6
	Owner_LoanAccount    Owner = 7
	Owner_DepositAccount Owner = 8
	Owner_IDDocument     Owner = 9
	Owner_LineOfCredit   Owner = 10
	Owner_GLJournalEntry Owner = 11
)

// Enum value maps for Owner.
var (
	Owner_name = map[int32]string{
		0:  "Client",
		1:  "Group",
		2:  "LoanProduct",
		3:  "SavingsProduct",
		4:  "Centre",
		5:  "Branch",
		6:  "User",
		7:  "LoanAccount",
		8:  "DepositAccount",
		9:  "IDDocument",
		10: "LineOfCredit",
		11: "GLJournalEntry",
	}
	Owner_value = map[string]int32{
		"Client":         0,
		"Group":          1,
		"LoanProduct":    2,
		"SavingsProduct": 3,
		"Centre":         4,
		"Branch":         5,
		"User":           6,
		"LoanAccount":    7,
		"DepositAccount": 8,
		"IDDocument":     9,
		"LineOfCredit":   10,
		"GLJournalEntry": 11,
	}
)

func (x Owner) Enum() *Owner {
	p := new(Owner)
	*p = x
	return p
}

func (x Owner) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Owner) Descriptor() protoreflect.EnumDescriptor {
	return file_bnk_to_core_api_v1_comments_all_proto_enumTypes[0].Descriptor()
}

func (Owner) Type() protoreflect.EnumType {
	return &file_bnk_to_core_api_v1_comments_all_proto_enumTypes[0]
}

func (x Owner) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Owner.Descriptor instead.
func (Owner) EnumDescriptor() ([]byte, []int) {
	return file_bnk_to_core_api_v1_comments_all_proto_rawDescGZIP(), []int{0}
}

// Comment is a comment added to an entity.
type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommentID is the ID of the comment.
	CommentID string `protobuf:"bytes,1,opt,name=CommentID,json=comment_id,proto3" json:"comment_id,omitempty"`
	// UserID is the ID of the user.
	UserID string `protobuf:"bytes,2,opt,name=UserID,json=user_id,proto3" json:"user_id,omitempty"`
	// OwnerID is the ID of the holder of the comment.
	OwnerID string `protobuf:"bytes,3,opt,name=OwnerID,json=owner_id,proto3" json:"owner_id,omitempty"`
	// OwnerType is the type of holder of the document.
	OwnerType Owner `protobuf:"varint,4,opt,name=OwnerType,json=owner_type,proto3,enum=openbank.core.v1.comments.Owner" json:"owner_type,omitempty"`
	// Text is the content of the comment.
	Text string `protobuf:"bytes,5,opt,name=Text,json=text,proto3" json:"text,omitempty"`
	// CreateTime is when the comment was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreateTime,json=create_time,proto3" json:"create_time,omitempty"`
	// UpdateTime is the time at which the comment was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=UpdateTime,json=update_time,proto3" json:"update_time,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_bnk_to_core_api_v1_comments_all_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetCommentID() string {
	if x != nil {
		return x.CommentID
	}
	return ""
}

func (x *Comment) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Comment) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

func (x *Comment) GetOwnerType() Owner {
	if x != nil {
		return x.OwnerType
	}
	return Owner_Client
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Comment) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// CreateCommentRequest is the create comment request.
type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Body is the comment to be created.
	Body *Comment `protobuf:"bytes,1,opt,name=Body,json=body,proto3" json:"body,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_bnk_to_core_api_v1_comments_all_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCommentRequest) GetBody() *Comment {
	if x != nil {
		return x.Body
	}
	return nil
}

// ListCommentsRequest is the request to list comments.
type ListCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PageToken is the token of the page to be returned.
	PageToken string `protobuf:"bytes,1,opt,name=PageToken,json=page_token,proto3" json:"page_token,omitempty"`
	// PageSize is the number of comments to be returned.
	PageSize int32 `protobuf:"varint,2,opt,name=PageSize,json=page_size,proto3" json:"page_size,omitempty"`
	// OrderBy is the field to order the comments by.
	OrderBy string `protobuf:"bytes,3,opt,name=OrderBy,json=order_by,proto3" json:"order_by,omitempty"`
	// Filter is the filter to apply to the comments.
	Filter string `protobuf:"bytes,4,opt,name=Filter,json=filter,proto3" json:"filter,omitempty"`
}

func (x *ListCommentsRequest) Reset() {
	*x = ListCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsRequest) ProtoMessage() {}

func (x *ListCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListCommentsRequest) Descriptor() ([]byte, []int) {
	return file_bnk_to_core_api_v1_comments_all_proto_rawDescGZIP(), []int{2}
}

func (x *ListCommentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCommentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCommentsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ListCommentsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// ListCommentsResponse is the response for listing comments.
type ListCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total is the total number of comments matching the filter.
	Total int32 `protobuf:"varint,1,opt,name=Total,json=total,proto3" json:"total,omitempty"`
	// Remaining is the number of comments remaining in the list, including the
	// comments in the response.
	Remaining int32 `protobuf:"varint,2,opt,name=Remaining,json=remaining,proto3" json:"remaining,omitempty"`
	// Comments is the list of comments.
	Comments []*Comment `protobuf:"bytes,3,rep,name=Comments,json=comments,proto3" json:"comments,omitempty"`
	// NextPageToken is the token of the next page.
	NextPageToken string `protobuf:"bytes,4,opt,name=NextPageToken,json=next_page_token,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCommentsResponse) Reset() {
	*x = ListCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsResponse) ProtoMessage() {}

func (x *ListCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListCommentsResponse) Descriptor() ([]byte, []int) {
	return file_bnk_to_core_api_v1_comments_all_proto_rawDescGZIP(), []int{3}
}

func (x *ListCommentsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListCommentsResponse) GetRemaining() int32 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

func (x *ListCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListCommentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// GetCommentRequest is the request to retrieve a comment.
type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CommentID is the ID of the comment to retrieve.
	CommentID string `protobuf:"bytes,1,opt,name=CommentID,json=comment_id,proto3" json:"comment_id,omitempty"`
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bnk_to_core_api_v1_comments_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_bnk_to_core_api_v1_comments_all_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentRequest) GetCommentID() string {
	if x != nil {
		return x.CommentID
	}
	return ""
}

var File_bnk_to_core_api_v1_comments_all_proto protoreflect.FileDescriptor

var file_bnk_to_core_api_v1_comments_all_proto_rawDesc = []byte{
	0x0a, 0x25, 0x62, 0x6e, 0x6b, 0x2e, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x6c,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0x0b, 0x78, 0x6f, 0x2f, 0x78, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5,
	0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0xda, 0x47, 0x02, 0x08, 0x02, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0xda, 0x47, 0x16, 0x2a, 0x14, 0x0a, 0x0a, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x07, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12,
	0x4b, 0x0a, 0x09, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x54, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x47, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x3a, 0x06,
	0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x62, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42,
	0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0xbc, 0x01, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a,
	0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x12, 0x22, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0xea, 0x01, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x4a,
	0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x0d, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3a, 0x06,
	0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x46, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x2a, 0xf4,
	0x01, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x0d, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x13, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x6e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x10, 0x02, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x16, 0x0a, 0x0e,
	0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x10, 0x03,
	0x1a, 0x02, 0x08, 0x00, 0x12, 0x0e, 0x0a, 0x06, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x10, 0x04,
	0x1a, 0x02, 0x08, 0x00, 0x12, 0x0e, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x10, 0x05,
	0x1a, 0x02, 0x08, 0x00, 0x12, 0x0c, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x10, 0x06, 0x1a, 0x02,
	0x08, 0x00, 0x12, 0x13, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x10, 0x07, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x16, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x08, 0x1a, 0x02, 0x08, 0x00, 0x12,
	0x12, 0x0a, 0x0a, 0x49, 0x44, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x09, 0x1a,
	0x02, 0x08, 0x00, 0x12, 0x14, 0x0a, 0x0c, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x66, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x10, 0x0a, 0x1a, 0x02, 0x08, 0x00, 0x12, 0x16, 0x0a, 0x0e, 0x47, 0x4c, 0x4a,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x10, 0x0b, 0x1a, 0x02, 0x08,
	0x00, 0x1a, 0x02, 0x18, 0x00, 0x32, 0xc2, 0x03, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x20, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x04,
	0x42, 0x6f, 0x64, 0x79, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x28, 0x00, 0x30, 0x00, 0x12, 0x8f, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x28, 0x00, 0x30, 0x00, 0x12, 0x8a, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x88, 0x02, 0x00, 0x90, 0x02,
	0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x7d, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42, 0x98, 0x01, 0x0a, 0x19, 0x6f,
	0x70, 0x65, 0x6e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x01, 0x50, 0x01, 0x5a, 0x24, 0x62, 0x6e,
	0x6b, 0x2e, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01,
	0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x92, 0x41, 0x3b, 0x0a, 0x03, 0x32, 0x2e, 0x30, 0x12,
	0x0d, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x02, 0x76, 0x31, 0x2a, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bnk_to_core_api_v1_comments_all_proto_rawDescOnce sync.Once
	file_bnk_to_core_api_v1_comments_all_proto_rawDescData = file_bnk_to_core_api_v1_comments_all_proto_rawDesc
)

func file_bnk_to_core_api_v1_comments_all_proto_rawDescGZIP() []byte {
	file_bnk_to_core_api_v1_comments_all_proto_rawDescOnce.Do(func() {
		file_bnk_to_core_api_v1_comments_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_bnk_to_core_api_v1_comments_all_proto_rawDescData)
	})
	return file_bnk_to_core_api_v1_comments_all_proto_rawDescData
}

var (
	file_bnk_to_core_api_v1_comments_all_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_bnk_to_core_api_v1_comments_all_proto_msgTypes  = make([]protoimpl.MessageInfo, 5)
	file_bnk_to_core_api_v1_comments_all_proto_goTypes   = []interface{}{
		(Owner)(0),                    // 0: openbank.core.v1.comments.Owner
		(*Comment)(nil),               // 1: openbank.core.v1.comments.Comment
		(*CreateCommentRequest)(nil),  // 2: openbank.core.v1.comments.CreateCommentRequest
		(*ListCommentsRequest)(nil),   // 3: openbank.core.v1.comments.ListCommentsRequest
		(*ListCommentsResponse)(nil),  // 4: openbank.core.v1.comments.ListCommentsResponse
		(*GetCommentRequest)(nil),     // 5: openbank.core.v1.comments.GetCommentRequest
		(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	}
)

var file_bnk_to_core_api_v1_comments_all_proto_depIdxs = []int32{
	0, // 0: openbank.core.v1.comments.Comment.OwnerType:type_name -> openbank.core.v1.comments.Owner
	6, // 1: openbank.core.v1.comments.Comment.CreateTime:type_name -> google.protobuf.Timestamp
	6, // 2: openbank.core.v1.comments.Comment.UpdateTime:type_name -> google.protobuf.Timestamp
	1, // 3: openbank.core.v1.comments.CreateCommentRequest.Body:type_name -> openbank.core.v1.comments.Comment
	1, // 4: openbank.core.v1.comments.ListCommentsResponse.Comments:type_name -> openbank.core.v1.comments.Comment
	2, // 5: openbank.core.v1.comments.CommentsService.CreateComment:input_type -> openbank.core.v1.comments.CreateCommentRequest
	3, // 6: openbank.core.v1.comments.CommentsService.ListComments:input_type -> openbank.core.v1.comments.ListCommentsRequest
	5, // 7: openbank.core.v1.comments.CommentsService.GetComment:input_type -> openbank.core.v1.comments.GetCommentRequest
	1, // 8: openbank.core.v1.comments.CommentsService.CreateComment:output_type -> openbank.core.v1.comments.Comment
	4, // 9: openbank.core.v1.comments.CommentsService.ListComments:output_type -> openbank.core.v1.comments.ListCommentsResponse
	1, // 10: openbank.core.v1.comments.CommentsService.GetComment:output_type -> openbank.core.v1.comments.Comment
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bnk_to_core_api_v1_comments_all_proto_init() }
func file_bnk_to_core_api_v1_comments_all_proto_init() {
	if File_bnk_to_core_api_v1_comments_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bnk_to_core_api_v1_comments_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_bnk_to_core_api_v1_comments_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
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
		file_bnk_to_core_api_v1_comments_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentsRequest); i {
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
		file_bnk_to_core_api_v1_comments_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentsResponse); i {
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
		file_bnk_to_core_api_v1_comments_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentRequest); i {
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
			RawDescriptor: file_bnk_to_core_api_v1_comments_all_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bnk_to_core_api_v1_comments_all_proto_goTypes,
		DependencyIndexes: file_bnk_to_core_api_v1_comments_all_proto_depIdxs,
		EnumInfos:         file_bnk_to_core_api_v1_comments_all_proto_enumTypes,
		MessageInfos:      file_bnk_to_core_api_v1_comments_all_proto_msgTypes,
	}.Build()
	File_bnk_to_core_api_v1_comments_all_proto = out.File
	file_bnk_to_core_api_v1_comments_all_proto_rawDesc = nil
	file_bnk_to_core_api_v1_comments_all_proto_goTypes = nil
	file_bnk_to_core_api_v1_comments_all_proto_depIdxs = nil
}