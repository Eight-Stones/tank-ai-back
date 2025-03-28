// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: action/v1/action.proto

package actionv1

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

// ObjectType enum for types of game object.
type ObjectType int32

const (
	ObjectType_UNDEFINED_OBJECT ObjectType = 0
	ObjectType_TANK             ObjectType = 1
	ObjectType_BULLET           ObjectType = 2
)

// Enum value maps for ObjectType.
var (
	ObjectType_name = map[int32]string{
		0: "UNDEFINED_OBJECT",
		1: "TANK",
		2: "BULLET",
	}
	ObjectType_value = map[string]int32{
		"UNDEFINED_OBJECT": 0,
		"TANK":             1,
		"BULLET":           2,
	}
)

func (x ObjectType) Enum() *ObjectType {
	p := new(ObjectType)
	*p = x
	return p
}

func (x ObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_action_v1_action_proto_enumTypes[0].Descriptor()
}

func (ObjectType) Type() protoreflect.EnumType {
	return &file_action_v1_action_proto_enumTypes[0]
}

func (x ObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectType.Descriptor instead.
func (ObjectType) EnumDescriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{0}
}

// ActionType enum for mark any action.
type ActionType int32

const (
	ActionType_UNDEFINED_ACTION           ActionType = 0       // 0b0
	ActionType_FAIL                       ActionType = 1       // 0b1
	ActionType_OK                         ActionType = 2       // 0b10
	ActionType_BAN                        ActionType = 4       // 0b100
	ActionType_OK_BORDER                  ActionType = 8       // 0b1000
	ActionType_FAIL_BORDER                ActionType = 16      // 0b10000
	ActionType_OK_STEP                    ActionType = 32      // 0b100000
	ActionType_FAIL_STEP                  ActionType = 64      // 0b1000000
	ActionType_OK_ROTATE                  ActionType = 128     // 0b10000000
	ActionType_FAIL_ROTATE                ActionType = 256     // 0b100000000
	ActionType_OK_COLLISION               ActionType = 512     // 0b1000000000
	ActionType_NO_COLLISION               ActionType = 1024    // 0b10000000000
	ActionType_NOT_INTERRUPT_OK_COLLISION ActionType = 8192    // 0b1000000000000
	ActionType_DAMAGED                    ActionType = 16384   // 0b10000000000000
	ActionType_BOTH_DAMAGED               ActionType = 32768   // 0b100000000000000
	ActionType_OK_SHOT                    ActionType = 65536   // 0b10000000000000000
	ActionType_FAIL_SHOT                  ActionType = 131072  // 0b100000000000000000
	ActionType_DISAPPEAR                  ActionType = 262144  // 0b1000000000000000000
	ActionType_FOUND                      ActionType = 524288  // 0b10000000000000000000
	ActionType_NOT_FOUND                  ActionType = 1048576 // 0b100000000000000000000
)

// Enum value maps for ActionType.
var (
	ActionType_name = map[int32]string{
		0:       "UNDEFINED_ACTION",
		1:       "FAIL",
		2:       "OK",
		4:       "BAN",
		8:       "OK_BORDER",
		16:      "FAIL_BORDER",
		32:      "OK_STEP",
		64:      "FAIL_STEP",
		128:     "OK_ROTATE",
		256:     "FAIL_ROTATE",
		512:     "OK_COLLISION",
		1024:    "NO_COLLISION",
		8192:    "NOT_INTERRUPT_OK_COLLISION",
		16384:   "DAMAGED",
		32768:   "BOTH_DAMAGED",
		65536:   "OK_SHOT",
		131072:  "FAIL_SHOT",
		262144:  "DISAPPEAR",
		524288:  "FOUND",
		1048576: "NOT_FOUND",
	}
	ActionType_value = map[string]int32{
		"UNDEFINED_ACTION":           0,
		"FAIL":                       1,
		"OK":                         2,
		"BAN":                        4,
		"OK_BORDER":                  8,
		"FAIL_BORDER":                16,
		"OK_STEP":                    32,
		"FAIL_STEP":                  64,
		"OK_ROTATE":                  128,
		"FAIL_ROTATE":                256,
		"OK_COLLISION":               512,
		"NO_COLLISION":               1024,
		"NOT_INTERRUPT_OK_COLLISION": 8192,
		"DAMAGED":                    16384,
		"BOTH_DAMAGED":               32768,
		"OK_SHOT":                    65536,
		"FAIL_SHOT":                  131072,
		"DISAPPEAR":                  262144,
		"FOUND":                      524288,
		"NOT_FOUND":                  1048576,
	}
)

func (x ActionType) Enum() *ActionType {
	p := new(ActionType)
	*p = x
	return p
}

func (x ActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_action_v1_action_proto_enumTypes[1].Descriptor()
}

func (ActionType) Type() protoreflect.EnumType {
	return &file_action_v1_action_proto_enumTypes[1]
}

func (x ActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionType.Descriptor instead.
func (ActionType) EnumDescriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{1}
}

// Info describe player.
type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	X    int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y    int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Hp   int32  `protobuf:"varint,4,opt,name=hp,proto3" json:"hp,omitempty"`
	Ammo int32  `protobuf:"varint,5,opt,name=ammo,proto3" json:"ammo,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	mi := &file_action_v1_action_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Info) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Info) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Info) GetHp() int32 {
	if x != nil {
		return x.Hp
	}
	return 0
}

func (x *Info) GetAmmo() int32 {
	if x != nil {
		return x.Ammo
	}
	return 0
}

// InfoReq describe player info request.
type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // gamer id.
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	mi := &file_action_v1_action_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReq.ProtoReflect.Descriptor instead.
func (*InfoReq) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{1}
}

func (x *InfoReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// InfoReq describe player info response.
type InfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"` // координаты, направление, hp, ammo
}

func (x *InfoResp) Reset() {
	*x = InfoResp{}
	mi := &file_action_v1_action_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResp) ProtoMessage() {}

func (x *InfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResp.ProtoReflect.Descriptor instead.
func (*InfoResp) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{2}
}

func (x *InfoResp) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

// RotateReq describe how player can rotate his object.
type RotateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                // gamer id.
	Direction uint32 `protobuf:"varint,2,opt,name=direction,proto3" json:"direction,omitempty"` // direction up/down/left/right.
}

func (x *RotateReq) Reset() {
	*x = RotateReq{}
	mi := &file_action_v1_action_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RotateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateReq) ProtoMessage() {}

func (x *RotateReq) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateReq.ProtoReflect.Descriptor instead.
func (*RotateReq) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{3}
}

func (x *RotateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RotateReq) GetDirection() uint32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

// RotateResp describe result of rotation.
type RotateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ActionType `protobuf:"varint,1,opt,name=code,proto3,enum=action.v1.ActionType" json:"code,omitempty"` // result info about rotate.
}

func (x *RotateResp) Reset() {
	*x = RotateResp{}
	mi := &file_action_v1_action_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RotateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateResp) ProtoMessage() {}

func (x *RotateResp) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateResp.ProtoReflect.Descriptor instead.
func (*RotateResp) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{4}
}

func (x *RotateResp) GetCode() ActionType {
	if x != nil {
		return x.Code
	}
	return ActionType_UNDEFINED_ACTION
}

// MoveReq describe how player can move his object.
type MoveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                // gamer id.
	Direction uint32 `protobuf:"varint,2,opt,name=direction,proto3" json:"direction,omitempty"` // direction up/down/left/right.
}

func (x *MoveReq) Reset() {
	*x = MoveReq{}
	mi := &file_action_v1_action_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveReq) ProtoMessage() {}

func (x *MoveReq) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveReq.ProtoReflect.Descriptor instead.
func (*MoveReq) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{5}
}

func (x *MoveReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MoveReq) GetDirection() uint32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

// MoveResp describe result of moving.
type MoveResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ActionType `protobuf:"varint,1,opt,name=code,proto3,enum=action.v1.ActionType" json:"code,omitempty"` // result info about move.
}

func (x *MoveResp) Reset() {
	*x = MoveResp{}
	mi := &file_action_v1_action_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResp) ProtoMessage() {}

func (x *MoveResp) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResp.ProtoReflect.Descriptor instead.
func (*MoveResp) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{6}
}

func (x *MoveResp) GetCode() ActionType {
	if x != nil {
		return x.Code
	}
	return ActionType_UNDEFINED_ACTION
}

// ShootReq describe how player can shoot by his object.
type ShootReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // gamer id.
}

func (x *ShootReq) Reset() {
	*x = ShootReq{}
	mi := &file_action_v1_action_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShootReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShootReq) ProtoMessage() {}

func (x *ShootReq) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShootReq.ProtoReflect.Descriptor instead.
func (*ShootReq) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{7}
}

func (x *ShootReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// ShootResp describe result of shooting.
type ShootResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ActionType `protobuf:"varint,1,opt,name=code,proto3,enum=action.v1.ActionType" json:"code,omitempty"` // result info about shoot.
}

func (x *ShootResp) Reset() {
	*x = ShootResp{}
	mi := &file_action_v1_action_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShootResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShootResp) ProtoMessage() {}

func (x *ShootResp) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShootResp.ProtoReflect.Descriptor instead.
func (*ShootResp) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{8}
}

func (x *ShootResp) GetCode() ActionType {
	if x != nil {
		return x.Code
	}
	return ActionType_UNDEFINED_ACTION
}

// Cell describe cell of field.
type Cell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X    int64      `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y    int64      `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Type ObjectType `protobuf:"varint,3,opt,name=type,proto3,enum=action.v1.ObjectType" json:"type,omitempty"`
}

func (x *Cell) Reset() {
	*x = Cell{}
	mi := &file_action_v1_action_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cell) ProtoMessage() {}

func (x *Cell) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cell.ProtoReflect.Descriptor instead.
func (*Cell) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{9}
}

func (x *Cell) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Cell) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Cell) GetType() ObjectType {
	if x != nil {
		return x.Type
	}
	return ObjectType_UNDEFINED_OBJECT
}

// VisionReq describe how player can vision by his object.
type VisionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // gamer id.
}

func (x *VisionReq) Reset() {
	*x = VisionReq{}
	mi := &file_action_v1_action_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisionReq) ProtoMessage() {}

func (x *VisionReq) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisionReq.ProtoReflect.Descriptor instead.
func (*VisionReq) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{10}
}

func (x *VisionReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// VisionResp describe result of vision.
type VisionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  ActionType `protobuf:"varint,1,opt,name=code,proto3,enum=action.v1.ActionType" json:"code,omitempty"` // result info about vision.
	Cells []*Cell    `protobuf:"bytes,2,rep,name=cells,proto3" json:"cells,omitempty"`
}

func (x *VisionResp) Reset() {
	*x = VisionResp{}
	mi := &file_action_v1_action_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisionResp) ProtoMessage() {}

func (x *VisionResp) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisionResp.ProtoReflect.Descriptor instead.
func (*VisionResp) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{11}
}

func (x *VisionResp) GetCode() ActionType {
	if x != nil {
		return x.Code
	}
	return ActionType_UNDEFINED_ACTION
}

func (x *VisionResp) GetCells() []*Cell {
	if x != nil {
		return x.Cells
	}
	return nil
}

// RadarReq describe how player can radar by his object.
type RadarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // gamer id.
}

func (x *RadarReq) Reset() {
	*x = RadarReq{}
	mi := &file_action_v1_action_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RadarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadarReq) ProtoMessage() {}

func (x *RadarReq) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadarReq.ProtoReflect.Descriptor instead.
func (*RadarReq) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{12}
}

func (x *RadarReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// VisionResp describe result of radar.
type RadarResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  ActionType `protobuf:"varint,1,opt,name=code,proto3,enum=action.v1.ActionType" json:"code,omitempty"` // result info about radar.
	Cells []*Cell    `protobuf:"bytes,2,rep,name=cells,proto3" json:"cells,omitempty"`
}

func (x *RadarResp) Reset() {
	*x = RadarResp{}
	mi := &file_action_v1_action_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RadarResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadarResp) ProtoMessage() {}

func (x *RadarResp) ProtoReflect() protoreflect.Message {
	mi := &file_action_v1_action_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadarResp.ProtoReflect.Descriptor instead.
func (*RadarResp) Descriptor() ([]byte, []int) {
	return file_action_v1_action_proto_rawDescGZIP(), []int{13}
}

func (x *RadarResp) GetCode() ActionType {
	if x != nil {
		return x.Code
	}
	return ActionType_UNDEFINED_ACTION
}

func (x *RadarResp) GetCells() []*Cell {
	if x != nil {
		return x.Cells
	}
	return nil
}

var File_action_v1_action_proto protoreflect.FileDescriptor

var file_action_v1_action_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x68, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x6d, 0x6d, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x61, 0x6d, 0x6d, 0x6f, 0x22, 0x19, 0x0a, 0x07, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x08, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x39, 0x0a, 0x09, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0a, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x37, 0x0a, 0x07, 0x4d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x08, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x53,
	0x68, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x6f, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x4d, 0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x79, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1b,
	0x0a, 0x09, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x0a, 0x56,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x65, 0x6c, 0x6c, 0x52, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x22, 0x1a, 0x0a, 0x08, 0x52,
	0x61, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x09, 0x52, 0x61, 0x64, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x25, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52,
	0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x2a, 0x38, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x41,
	0x4e, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x55, 0x4c, 0x4c, 0x45, 0x54, 0x10, 0x02,
	0x2a, 0xd4, 0x02, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x41, 0x4e, 0x10, 0x04,
	0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x4b, 0x5f, 0x42, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x08, 0x12,
	0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x42, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x10,
	0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4b, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x10, 0x20, 0x12, 0x0d, 0x0a,
	0x09, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x10, 0x40, 0x12, 0x0e, 0x0a, 0x09,
	0x4f, 0x4b, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x45, 0x10, 0x80, 0x01, 0x12, 0x10, 0x0a, 0x0b,
	0x46, 0x41, 0x49, 0x4c, 0x5f, 0x52, 0x4f, 0x54, 0x41, 0x54, 0x45, 0x10, 0x80, 0x02, 0x12, 0x11,
	0x0a, 0x0c, 0x4f, 0x4b, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x80,
	0x04, 0x12, 0x11, 0x0a, 0x0c, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x49, 0x53, 0x49, 0x4f,
	0x4e, 0x10, 0x80, 0x08, 0x12, 0x1f, 0x0a, 0x1a, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x52, 0x55, 0x50, 0x54, 0x5f, 0x4f, 0x4b, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x49, 0x53, 0x49,
	0x4f, 0x4e, 0x10, 0x80, 0x40, 0x12, 0x0d, 0x0a, 0x07, 0x44, 0x41, 0x4d, 0x41, 0x47, 0x45, 0x44,
	0x10, 0x80, 0x80, 0x01, 0x12, 0x12, 0x0a, 0x0c, 0x42, 0x4f, 0x54, 0x48, 0x5f, 0x44, 0x41, 0x4d,
	0x41, 0x47, 0x45, 0x44, 0x10, 0x80, 0x80, 0x02, 0x12, 0x0d, 0x0a, 0x07, 0x4f, 0x4b, 0x5f, 0x53,
	0x48, 0x4f, 0x54, 0x10, 0x80, 0x80, 0x04, 0x12, 0x0f, 0x0a, 0x09, 0x46, 0x41, 0x49, 0x4c, 0x5f,
	0x53, 0x48, 0x4f, 0x54, 0x10, 0x80, 0x80, 0x08, 0x12, 0x0f, 0x0a, 0x09, 0x44, 0x49, 0x53, 0x41,
	0x50, 0x50, 0x45, 0x41, 0x52, 0x10, 0x80, 0x80, 0x10, 0x12, 0x0b, 0x0a, 0x05, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x80, 0x80, 0x20, 0x12, 0x0f, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x80, 0x80, 0x40, 0x32, 0xd3, 0x02, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06,
	0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x04, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x12, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x6f,
	0x74, 0x12, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x06, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x52, 0x61, 0x64, 0x61, 0x72,
	0x12, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x64,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x61, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x9c, 0x01,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa,
	0x02, 0x09, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_action_v1_action_proto_rawDescOnce sync.Once
	file_action_v1_action_proto_rawDescData = file_action_v1_action_proto_rawDesc
)

func file_action_v1_action_proto_rawDescGZIP() []byte {
	file_action_v1_action_proto_rawDescOnce.Do(func() {
		file_action_v1_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_action_v1_action_proto_rawDescData)
	})
	return file_action_v1_action_proto_rawDescData
}

var file_action_v1_action_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_action_v1_action_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_action_v1_action_proto_goTypes = []any{
	(ObjectType)(0),    // 0: action.v1.ObjectType
	(ActionType)(0),    // 1: action.v1.ActionType
	(*Info)(nil),       // 2: action.v1.Info
	(*InfoReq)(nil),    // 3: action.v1.InfoReq
	(*InfoResp)(nil),   // 4: action.v1.InfoResp
	(*RotateReq)(nil),  // 5: action.v1.RotateReq
	(*RotateResp)(nil), // 6: action.v1.RotateResp
	(*MoveReq)(nil),    // 7: action.v1.MoveReq
	(*MoveResp)(nil),   // 8: action.v1.MoveResp
	(*ShootReq)(nil),   // 9: action.v1.ShootReq
	(*ShootResp)(nil),  // 10: action.v1.ShootResp
	(*Cell)(nil),       // 11: action.v1.Cell
	(*VisionReq)(nil),  // 12: action.v1.VisionReq
	(*VisionResp)(nil), // 13: action.v1.VisionResp
	(*RadarReq)(nil),   // 14: action.v1.RadarReq
	(*RadarResp)(nil),  // 15: action.v1.RadarResp
}
var file_action_v1_action_proto_depIdxs = []int32{
	2,  // 0: action.v1.InfoResp.info:type_name -> action.v1.Info
	1,  // 1: action.v1.RotateResp.code:type_name -> action.v1.ActionType
	1,  // 2: action.v1.MoveResp.code:type_name -> action.v1.ActionType
	1,  // 3: action.v1.ShootResp.code:type_name -> action.v1.ActionType
	0,  // 4: action.v1.Cell.type:type_name -> action.v1.ObjectType
	1,  // 5: action.v1.VisionResp.code:type_name -> action.v1.ActionType
	11, // 6: action.v1.VisionResp.cells:type_name -> action.v1.Cell
	1,  // 7: action.v1.RadarResp.code:type_name -> action.v1.ActionType
	11, // 8: action.v1.RadarResp.cells:type_name -> action.v1.Cell
	3,  // 9: action.v1.ActionService.Info:input_type -> action.v1.InfoReq
	5,  // 10: action.v1.ActionService.Rotate:input_type -> action.v1.RotateReq
	7,  // 11: action.v1.ActionService.Move:input_type -> action.v1.MoveReq
	9,  // 12: action.v1.ActionService.Shoot:input_type -> action.v1.ShootReq
	12, // 13: action.v1.ActionService.Vision:input_type -> action.v1.VisionReq
	14, // 14: action.v1.ActionService.Radar:input_type -> action.v1.RadarReq
	4,  // 15: action.v1.ActionService.Info:output_type -> action.v1.InfoResp
	6,  // 16: action.v1.ActionService.Rotate:output_type -> action.v1.RotateResp
	8,  // 17: action.v1.ActionService.Move:output_type -> action.v1.MoveResp
	10, // 18: action.v1.ActionService.Shoot:output_type -> action.v1.ShootResp
	13, // 19: action.v1.ActionService.Vision:output_type -> action.v1.VisionResp
	15, // 20: action.v1.ActionService.Radar:output_type -> action.v1.RadarResp
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_action_v1_action_proto_init() }
func file_action_v1_action_proto_init() {
	if File_action_v1_action_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_action_v1_action_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_action_v1_action_proto_goTypes,
		DependencyIndexes: file_action_v1_action_proto_depIdxs,
		EnumInfos:         file_action_v1_action_proto_enumTypes,
		MessageInfos:      file_action_v1_action_proto_msgTypes,
	}.Build()
	File_action_v1_action_proto = out.File
	file_action_v1_action_proto_rawDesc = nil
	file_action_v1_action_proto_goTypes = nil
	file_action_v1_action_proto_depIdxs = nil
}
