// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package fotune_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = timestamp.Timestamp

type GetUserMasterReq struct {
	InviteUid            string   `protobuf:"bytes,1,opt,name=InviteUid,proto3" json:"InviteUid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserMasterReq) Reset()         { *m = GetUserMasterReq{} }
func (m *GetUserMasterReq) String() string { return proto.CompactTextString(m) }
func (*GetUserMasterReq) ProtoMessage()    {}
func (*GetUserMasterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *GetUserMasterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserMasterReq.Unmarshal(m, b)
}
func (m *GetUserMasterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserMasterReq.Marshal(b, m, deterministic)
}
func (m *GetUserMasterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserMasterReq.Merge(m, src)
}
func (m *GetUserMasterReq) XXX_Size() int {
	return xxx_messageInfo_GetUserMasterReq.Size(m)
}
func (m *GetUserMasterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserMasterReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserMasterReq proto.InternalMessageInfo

func (m *GetUserMasterReq) GetInviteUid() string {
	if m != nil {
		return m.InviteUid
	}
	return ""
}

type UserMasterResp struct {
	UserMasterId         string   `protobuf:"bytes,1,opt,name=UserMasterId,proto3" json:"UserMasterId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMasterResp) Reset()         { *m = UserMasterResp{} }
func (m *UserMasterResp) String() string { return proto.CompactTextString(m) }
func (*UserMasterResp) ProtoMessage()    {}
func (*UserMasterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserMasterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMasterResp.Unmarshal(m, b)
}
func (m *UserMasterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMasterResp.Marshal(b, m, deterministic)
}
func (m *UserMasterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMasterResp.Merge(m, src)
}
func (m *UserMasterResp) XXX_Size() int {
	return xxx_messageInfo_UserMasterResp.Size(m)
}
func (m *UserMasterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMasterResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserMasterResp proto.InternalMessageInfo

func (m *UserMasterResp) GetUserMasterId() string {
	if m != nil {
		return m.UserMasterId
	}
	return ""
}

type AllUserInfoResp struct {
	UserInfo             []*LoginResp `protobuf:"bytes,1,rep,name=UserInfo,proto3" json:"UserInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AllUserInfoResp) Reset()         { *m = AllUserInfoResp{} }
func (m *AllUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*AllUserInfoResp) ProtoMessage()    {}
func (*AllUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *AllUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllUserInfoResp.Unmarshal(m, b)
}
func (m *AllUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllUserInfoResp.Marshal(b, m, deterministic)
}
func (m *AllUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllUserInfoResp.Merge(m, src)
}
func (m *AllUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_AllUserInfoResp.Size(m)
}
func (m *AllUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AllUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_AllUserInfoResp proto.InternalMessageInfo

func (m *AllUserInfoResp) GetUserInfo() []*LoginResp {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type UserInfoReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReq) Reset()         { *m = UserInfoReq{} }
func (m *UserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoReq) ProtoMessage()    {}
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReq.Unmarshal(m, b)
}
func (m *UserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReq.Marshal(b, m, deterministic)
}
func (m *UserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReq.Merge(m, src)
}
func (m *UserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoReq.Size(m)
}
func (m *UserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReq proto.InternalMessageInfo

func (m *UserInfoReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type LoginReq struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResp struct {
	UserID               string               `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	InvitationCode       string               `protobuf:"bytes,2,opt,name=invitationCode,proto3" json:"invitationCode,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string               `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Phone                string               `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	LastLogin            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=lastLogin,proto3" json:"lastLogin,omitempty"`
	LoginCount           int32                `protobuf:"varint,7,opt,name=loginCount,proto3" json:"loginCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *LoginResp) GetInvitationCode() string {
	if m != nil {
		return m.InvitationCode
	}
	return ""
}

func (m *LoginResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginResp) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *LoginResp) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *LoginResp) GetLastLogin() *timestamp.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

func (m *LoginResp) GetLoginCount() int32 {
	if m != nil {
		return m.LoginCount
	}
	return 0
}

type ValidateCodeReq struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateCodeReq) Reset()         { *m = ValidateCodeReq{} }
func (m *ValidateCodeReq) String() string { return proto.CompactTextString(m) }
func (*ValidateCodeReq) ProtoMessage()    {}
func (*ValidateCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *ValidateCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateCodeReq.Unmarshal(m, b)
}
func (m *ValidateCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateCodeReq.Marshal(b, m, deterministic)
}
func (m *ValidateCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateCodeReq.Merge(m, src)
}
func (m *ValidateCodeReq) XXX_Size() int {
	return xxx_messageInfo_ValidateCodeReq.Size(m)
}
func (m *ValidateCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateCodeReq proto.InternalMessageInfo

func (m *ValidateCodeReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type ValidateCodeResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateCodeResp) Reset()         { *m = ValidateCodeResp{} }
func (m *ValidateCodeResp) String() string { return proto.CompactTextString(m) }
func (*ValidateCodeResp) ProtoMessage()    {}
func (*ValidateCodeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *ValidateCodeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateCodeResp.Unmarshal(m, b)
}
func (m *ValidateCodeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateCodeResp.Marshal(b, m, deterministic)
}
func (m *ValidateCodeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateCodeResp.Merge(m, src)
}
func (m *ValidateCodeResp) XXX_Size() int {
	return xxx_messageInfo_ValidateCodeResp.Size(m)
}
func (m *ValidateCodeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateCodeResp.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateCodeResp proto.InternalMessageInfo

func (m *ValidateCodeResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type RegisterReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,4,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	InvitationCode       string   `protobuf:"bytes,5,opt,name=invitationCode,proto3" json:"invitationCode,omitempty"`
	ValidateCode         string   `protobuf:"bytes,6,opt,name=validateCode,proto3" json:"validateCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterReq) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *RegisterReq) GetInvitationCode() string {
	if m != nil {
		return m.InvitationCode
	}
	return ""
}

func (m *RegisterReq) GetValidateCode() string {
	if m != nil {
		return m.ValidateCode
	}
	return ""
}

type ChangePasswordReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,3,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordReq) Reset()         { *m = ChangePasswordReq{} }
func (m *ChangePasswordReq) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordReq) ProtoMessage()    {}
func (*ChangePasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *ChangePasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordReq.Unmarshal(m, b)
}
func (m *ChangePasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordReq.Marshal(b, m, deterministic)
}
func (m *ChangePasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordReq.Merge(m, src)
}
func (m *ChangePasswordReq) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordReq.Size(m)
}
func (m *ChangePasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordReq proto.InternalMessageInfo

func (m *ChangePasswordReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ChangePasswordReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ChangePasswordReq) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

type ForgetPasswordReq struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string   `protobuf:"bytes,2,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	ValidateCode         string   `protobuf:"bytes,4,opt,name=validateCode,proto3" json:"validateCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgetPasswordReq) Reset()         { *m = ForgetPasswordReq{} }
func (m *ForgetPasswordReq) String() string { return proto.CompactTextString(m) }
func (*ForgetPasswordReq) ProtoMessage()    {}
func (*ForgetPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *ForgetPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgetPasswordReq.Unmarshal(m, b)
}
func (m *ForgetPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgetPasswordReq.Marshal(b, m, deterministic)
}
func (m *ForgetPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgetPasswordReq.Merge(m, src)
}
func (m *ForgetPasswordReq) XXX_Size() int {
	return xxx_messageInfo_ForgetPasswordReq.Size(m)
}
func (m *ForgetPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgetPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ForgetPasswordReq proto.InternalMessageInfo

func (m *ForgetPasswordReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ForgetPasswordReq) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *ForgetPasswordReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ForgetPasswordReq) GetValidateCode() string {
	if m != nil {
		return m.ValidateCode
	}
	return ""
}

type UpdateUserReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserReq) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UpdateUserReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Member struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Price                int32    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	OldPrice             int32    `protobuf:"varint,5,opt,name=old_price,json=oldPrice,proto3" json:"old_price,omitempty"`
	Duration             int32    `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	State                int32    `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{12}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Member) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Member) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Member) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Member) GetOldPrice() int32 {
	if m != nil {
		return m.OldPrice
	}
	return 0
}

func (m *Member) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Member) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Member) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Member) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type GetMembersResp struct {
	Members              []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetMembersResp) Reset()         { *m = GetMembersResp{} }
func (m *GetMembersResp) String() string { return proto.CompactTextString(m) }
func (*GetMembersResp) ProtoMessage()    {}
func (*GetMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{13}
}

func (m *GetMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMembersResp.Unmarshal(m, b)
}
func (m *GetMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMembersResp.Marshal(b, m, deterministic)
}
func (m *GetMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMembersResp.Merge(m, src)
}
func (m *GetMembersResp) XXX_Size() int {
	return xxx_messageInfo_GetMembersResp.Size(m)
}
func (m *GetMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetMembersResp proto.InternalMessageInfo

func (m *GetMembersResp) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type Payment struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	BitAddr              string   `protobuf:"bytes,4,opt,name=BitAddr,proto3" json:"BitAddr,omitempty"`
	BitCode              string   `protobuf:"bytes,5,opt,name=BitCode,proto3" json:"BitCode,omitempty"`
	State                int32    `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{14}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Payment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Payment) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Payment) GetBitAddr() string {
	if m != nil {
		return m.BitAddr
	}
	return ""
}

func (m *Payment) GetBitCode() string {
	if m != nil {
		return m.BitCode
	}
	return ""
}

func (m *Payment) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

type GetPaymentMethodResp struct {
	Payments             []*Payment `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetPaymentMethodResp) Reset()         { *m = GetPaymentMethodResp{} }
func (m *GetPaymentMethodResp) String() string { return proto.CompactTextString(m) }
func (*GetPaymentMethodResp) ProtoMessage()    {}
func (*GetPaymentMethodResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{15}
}

func (m *GetPaymentMethodResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentMethodResp.Unmarshal(m, b)
}
func (m *GetPaymentMethodResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentMethodResp.Marshal(b, m, deterministic)
}
func (m *GetPaymentMethodResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentMethodResp.Merge(m, src)
}
func (m *GetPaymentMethodResp) XXX_Size() int {
	return xxx_messageInfo_GetPaymentMethodResp.Size(m)
}
func (m *GetPaymentMethodResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentMethodResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentMethodResp proto.InternalMessageInfo

func (m *GetPaymentMethodResp) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserMasterReq)(nil), "GetUserMasterReq")
	proto.RegisterType((*UserMasterResp)(nil), "UserMasterResp")
	proto.RegisterType((*AllUserInfoResp)(nil), "AllUserInfoResp")
	proto.RegisterType((*UserInfoReq)(nil), "userInfoReq")
	proto.RegisterType((*LoginReq)(nil), "LoginReq")
	proto.RegisterType((*LoginResp)(nil), "LoginResp")
	proto.RegisterType((*ValidateCodeReq)(nil), "ValidateCodeReq")
	proto.RegisterType((*ValidateCodeResp)(nil), "ValidateCodeResp")
	proto.RegisterType((*RegisterReq)(nil), "RegisterReq")
	proto.RegisterType((*ChangePasswordReq)(nil), "ChangePasswordReq")
	proto.RegisterType((*ForgetPasswordReq)(nil), "ForgetPasswordReq")
	proto.RegisterType((*UpdateUserReq)(nil), "UpdateUserReq")
	proto.RegisterType((*Member)(nil), "Member")
	proto.RegisterType((*GetMembersResp)(nil), "GetMembersResp")
	proto.RegisterType((*Payment)(nil), "Payment")
	proto.RegisterType((*GetPaymentMethodResp)(nil), "GetPaymentMethodResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 895 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0x8e, 0x93, 0x38, 0x97, 0x93, 0x6e, 0x2e, 0xc3, 0xb2, 0x32, 0x5e, 0x2e, 0x61, 0x04, 0x25,
	0xbc, 0xcc, 0xa2, 0x2e, 0x02, 0x56, 0x14, 0xa4, 0xb4, 0x40, 0x55, 0x89, 0x4a, 0x95, 0x97, 0xf6,
	0x81, 0x97, 0x6a, 0x5a, 0x4f, 0x53, 0x8b, 0xd8, 0xe3, 0xd8, 0x93, 0xa0, 0xfe, 0x0a, 0x1e, 0xf8,
	0x55, 0xfc, 0x12, 0xfe, 0x02, 0x8f, 0x68, 0x66, 0x3c, 0xbe, 0x35, 0x09, 0x0f, 0xbc, 0xf9, 0x3b,
	0x67, 0xce, 0x99, 0x73, 0xbe, 0x73, 0xf1, 0x00, 0xac, 0x53, 0x96, 0x90, 0x38, 0xe1, 0x82, 0xbb,
	0x2f, 0x17, 0x9c, 0x2f, 0x96, 0xec, 0x95, 0x42, 0xb7, 0xeb, 0xfb, 0x57, 0x2c, 0x8c, 0xc5, 0x63,
	0xa6, 0xfc, 0xa8, 0xae, 0x14, 0x41, 0xc8, 0x52, 0x41, 0xc3, 0x58, 0x1f, 0xc0, 0x5f, 0xc0, 0xf8,
	0x8c, 0x89, 0xab, 0x94, 0x25, 0x17, 0x34, 0x15, 0x2c, 0xf1, 0xd8, 0x0a, 0xbd, 0x0f, 0xfd, 0xf3,
	0x68, 0x13, 0x08, 0x76, 0x15, 0xf8, 0x8e, 0x35, 0xb5, 0x66, 0x7d, 0xaf, 0x10, 0xe0, 0x2f, 0x61,
	0x58, 0x3e, 0x9e, 0xc6, 0x08, 0xc3, 0x41, 0x21, 0x39, 0x37, 0x26, 0x15, 0x19, 0x7e, 0x03, 0xa3,
	0xf9, 0x72, 0x29, 0x45, 0xe7, 0xd1, 0x3d, 0x57, 0x66, 0x87, 0xd0, 0x33, 0xd8, 0xb1, 0xa6, 0xad,
	0xd9, 0xe0, 0x08, 0xc8, 0xcf, 0x7c, 0x11, 0x44, 0x52, 0xeb, 0xe5, 0x3a, 0xfc, 0x29, 0x0c, 0xd6,
	0xb9, 0xdd, 0x0a, 0xbd, 0x80, 0x8e, 0x82, 0x3f, 0x64, 0xf7, 0x64, 0x08, 0x1f, 0x43, 0x2f, 0xb3,
	0x5e, 0xa1, 0xe7, 0x60, 0xc7, 0x0f, 0x3c, 0x62, 0xd9, 0x11, 0x0d, 0x90, 0x0b, 0xbd, 0x98, 0xa6,
	0xe9, 0xef, 0x3c, 0xf1, 0x9d, 0xa6, 0x52, 0xe4, 0x18, 0xff, 0x6d, 0x41, 0x3f, 0xbf, 0x7c, 0xd7,
	0x1d, 0xe8, 0x10, 0x86, 0x81, 0x24, 0x82, 0x8a, 0x80, 0x47, 0xa7, 0xdc, 0x67, 0x99, 0x9f, 0x9a,
	0x14, 0x21, 0x68, 0x47, 0x34, 0x64, 0x4e, 0x4b, 0x69, 0xd5, 0xb7, 0xf4, 0x49, 0x37, 0x54, 0xd0,
	0xc4, 0x69, 0x6b, 0x9f, 0x1a, 0x15, 0xb1, 0xda, 0xe5, 0x58, 0xbf, 0x81, 0xfe, 0x92, 0xa6, 0x42,
	0x85, 0xe4, 0x74, 0xa6, 0xd6, 0x6c, 0x70, 0xe4, 0x12, 0x5d, 0x4c, 0x62, 0x8a, 0x49, 0x7e, 0x31,
	0xc5, 0xf4, 0x8a, 0xc3, 0xe8, 0x43, 0x80, 0xa5, 0xfc, 0x38, 0xe5, 0xeb, 0x48, 0x38, 0xdd, 0xa9,
	0x35, 0xb3, 0xbd, 0x92, 0x04, 0x7f, 0x06, 0xa3, 0x6b, 0xba, 0x0c, 0x7c, 0x2a, 0x98, 0x8c, 0x75,
	0x27, 0x5d, 0xf8, 0x10, 0xc6, 0xd5, 0x83, 0x69, 0x2c, 0x13, 0xbb, 0x93, 0x69, 0xeb, 0x83, 0xea,
	0x1b, 0xff, 0x65, 0xc1, 0xc0, 0x63, 0x8b, 0xc0, 0xb4, 0x8f, 0x49, 0xde, 0x2a, 0x25, 0x9f, 0xdf,
	0xd0, 0xdc, 0x55, 0x90, 0x56, 0xb5, 0x20, 0x68, 0x06, 0xa3, 0x3b, 0x1e, 0xdd, 0x07, 0x49, 0x78,
	0x69, 0x8e, 0x68, 0xde, 0xea, 0xe2, 0x2d, 0x45, 0xb1, 0xb7, 0x16, 0x05, 0xc3, 0xc1, 0xa6, 0x94,
	0x8f, 0x62, 0xb5, 0xef, 0x55, 0x64, 0x78, 0x05, 0x93, 0xd3, 0x07, 0x1a, 0x2d, 0x98, 0xf1, 0xbe,
	0xa7, 0xe3, 0xf6, 0xf5, 0xd3, 0xb6, 0xf0, 0x5b, 0x5b, 0xc3, 0xc7, 0x7f, 0x5a, 0x30, 0xf9, 0x89,
	0x27, 0x0b, 0x26, 0xca, 0x77, 0x96, 0x7d, 0x5b, 0xff, 0xed, 0xbb, 0xb9, 0x9d, 0x9a, 0x9c, 0xf6,
	0x56, 0x99, 0xf6, 0x3a, 0x11, 0xed, 0x2d, 0x44, 0xbc, 0x85, 0x67, 0x57, 0xb1, 0x44, 0x72, 0x0c,
	0x77, 0x55, 0xb5, 0x68, 0xe9, 0x66, 0xa5, 0xa5, 0x0d, 0x61, 0x26, 0xe7, 0x0c, 0xc9, 0x21, 0xeb,
	0x5c, 0xb0, 0xf0, 0x96, 0x25, 0x68, 0x08, 0xcd, 0x6c, 0xb9, 0xd8, 0x5e, 0x33, 0xf0, 0x73, 0xf7,
	0xcd, 0xaa, 0xfb, 0x84, 0x85, 0x34, 0xf9, 0xcd, 0xb8, 0xd1, 0x48, 0x65, 0x95, 0x04, 0x77, 0x3a,
	0x70, 0xdb, 0xd3, 0x00, 0xbd, 0x84, 0x3e, 0x5f, 0xfa, 0x37, 0x5a, 0x63, 0x2b, 0x4d, 0x8f, 0x2f,
	0xfd, 0x4b, 0xa5, 0x74, 0xa1, 0xe7, 0xaf, 0x13, 0xd5, 0x0b, 0xaa, 0xee, 0xb6, 0x97, 0x63, 0xe9,
	0x2e, 0x15, 0x54, 0xb0, 0x6c, 0x56, 0x34, 0x40, 0x1f, 0x00, 0xdc, 0x25, 0x8c, 0x0a, 0xe6, 0xdf,
	0x50, 0xe1, 0xf4, 0xf4, 0x16, 0xcc, 0x24, 0x73, 0x21, 0xd5, 0x6b, 0xc5, 0x8f, 0x52, 0xf7, 0xb5,
	0x3a, 0x93, 0xcc, 0x05, 0x7e, 0x0d, 0xc3, 0x33, 0x26, 0x74, 0xae, 0xa9, 0x9a, 0x9c, 0x8f, 0xa1,
	0x1b, 0x6a, 0x98, 0x2d, 0xbb, 0x2e, 0xd1, 0x6a, 0xcf, 0xc8, 0xf1, 0x1f, 0x16, 0x74, 0x2f, 0xe9,
	0x63, 0xc8, 0x22, 0xf1, 0xbf, 0xf8, 0x71, 0xa0, 0x7b, 0x12, 0x88, 0xb9, 0xef, 0x9b, 0x55, 0x63,
	0x60, 0xa6, 0x29, 0xcd, 0x88, 0x81, 0x05, 0x09, 0x9d, 0x12, 0x09, 0xf8, 0x18, 0x9e, 0x9f, 0xc9,
	0xbe, 0x54, 0x31, 0x5d, 0x30, 0xf1, 0xc0, 0x7d, 0x95, 0xcc, 0x27, 0xb2, 0x3b, 0x95, 0xd0, 0x64,
	0xd3, 0x23, 0xd9, 0x29, 0x2f, 0xd7, 0x1c, 0xfd, 0xd3, 0x86, 0xb6, 0x6c, 0x1f, 0x34, 0x05, 0x5b,
	0xef, 0xa6, 0xbe, 0x59, 0xf0, 0x2b, 0xb7, 0xb4, 0xeb, 0x71, 0x03, 0xbd, 0x81, 0xf1, 0x5b, 0x16,
	0xf9, 0xe5, 0x7d, 0x83, 0xc6, 0xa4, 0xb6, 0xa7, 0xdc, 0x09, 0xa9, 0x2f, 0x24, 0xdc, 0x40, 0x47,
	0xd0, 0x33, 0xdb, 0x07, 0x1d, 0x90, 0xd2, 0x22, 0x72, 0x5f, 0x3c, 0x59, 0x98, 0x3f, 0xca, 0x5f,
	0x23, 0x6e, 0xa0, 0xef, 0xe0, 0x99, 0xc7, 0xd2, 0x62, 0xe2, 0x10, 0x22, 0x4f, 0xc6, 0x7e, 0x8f,
	0xf9, 0xf7, 0x30, 0xac, 0x4e, 0x2c, 0x42, 0xe4, 0xc9, 0x08, 0xef, 0xb1, 0xff, 0x0a, 0xa0, 0x18,
	0x2e, 0x34, 0x24, 0x95, 0x49, 0xdb, 0x63, 0xf7, 0x39, 0x0c, 0xb2, 0x9f, 0xb5, 0xfc, 0x19, 0xa2,
	0x03, 0x52, 0xfa, 0x2f, 0xd6, 0x08, 0xfd, 0x1a, 0xa0, 0x68, 0x40, 0xb4, 0xc3, 0xa5, 0x3b, 0x22,
	0xd5, 0x2e, 0xc5, 0x0d, 0x74, 0xaa, 0x1e, 0x04, 0x95, 0x92, 0xef, 0x34, 0x7f, 0x97, 0x6c, 0xeb,
	0x0e, 0xdc, 0x40, 0xc7, 0xaa, 0xfd, 0x4b, 0x3f, 0xfc, 0x9d, 0x2e, 0xc6, 0xa4, 0xf6, 0x2c, 0xc0,
	0x0d, 0x34, 0x87, 0xf7, 0x2a, 0x6f, 0x92, 0x93, 0xc7, 0xf3, 0xe8, 0x3a, 0xc8, 0xd8, 0x9a, 0x90,
	0xfa, 0x7b, 0xc5, 0x1d, 0x91, 0xea, 0x83, 0x04, 0x37, 0x4e, 0xde, 0xf9, 0x75, 0x42, 0xbe, 0xbd,
	0xe7, 0x62, 0x1d, 0xb1, 0x9b, 0x34, 0xd9, 0xdc, 0x48, 0xa2, 0x2e, 0xad, 0xdb, 0x8e, 0xba, 0xfd,
	0xf5, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0x06, 0xf1, 0x15, 0x40, 0x09, 0x00, 0x00,
}
