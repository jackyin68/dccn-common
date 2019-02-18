// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Task Events operation code
type DCOperation int32

const (
	DCOperation_TASK_CREATE DCOperation = 0
	DCOperation_TASK_CANCEL DCOperation = 1
	DCOperation_TASK_UPDATE DCOperation = 2
	DCOperation_HEARTBEAT   DCOperation = 3
)

var DCOperation_name = map[int32]string{
	0: "TASK_CREATE",
	1: "TASK_CANCEL",
	2: "TASK_UPDATE",
	3: "HEARTBEAT",
}
var DCOperation_value = map[string]int32{
	"TASK_CREATE": 0,
	"TASK_CANCEL": 1,
	"TASK_UPDATE": 2,
	"HEARTBEAT":   3,
}

func (x DCOperation) String() string {
	return proto.EnumName(DCOperation_name, int32(x))
}
func (DCOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{0}
}

// Hub task status
type TaskStatus int32

const (
	TaskStatus_STARTING       TaskStatus = 0
	TaskStatus_START_SUCCESS  TaskStatus = 1
	TaskStatus_START_FAILED   TaskStatus = 2
	TaskStatus_RUNNING        TaskStatus = 3
	TaskStatus_UPDATING       TaskStatus = 4
	TaskStatus_UPDATE_SUCCESS TaskStatus = 5
	TaskStatus_UPDATE_FAILED  TaskStatus = 6
	TaskStatus_CANCELLING     TaskStatus = 7
	TaskStatus_CANCELLED      TaskStatus = 8
	TaskStatus_CANCEL_FAILED  TaskStatus = 9
	TaskStatus_DONE           TaskStatus = 10
)

var TaskStatus_name = map[int32]string{
	0:  "STARTING",
	1:  "START_SUCCESS",
	2:  "START_FAILED",
	3:  "RUNNING",
	4:  "UPDATING",
	5:  "UPDATE_SUCCESS",
	6:  "UPDATE_FAILED",
	7:  "CANCELLING",
	8:  "CANCELLED",
	9:  "CANCEL_FAILED",
	10: "DONE",
}
var TaskStatus_value = map[string]int32{
	"STARTING":       0,
	"START_SUCCESS":  1,
	"START_FAILED":   2,
	"RUNNING":        3,
	"UPDATING":       4,
	"UPDATE_SUCCESS": 5,
	"UPDATE_FAILED":  6,
	"CANCELLING":     7,
	"CANCELLED":      8,
	"CANCEL_FAILED":  9,
	"DONE":           10,
}

func (x TaskStatus) String() string {
	return proto.EnumName(TaskStatus_name, int32(x))
}
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{1}
}

type TaskType int32

const (
	TaskType_DEFAULT    TaskType = 0
	TaskType_DEPLOYMENT TaskType = 1
	TaskType_JOB        TaskType = 2
	TaskType_CRONJOB    TaskType = 3
)

var TaskType_name = map[int32]string{
	0: "DEFAULT",
	1: "DEPLOYMENT",
	2: "JOB",
	3: "CRONJOB",
}
var TaskType_value = map[string]int32{
	"DEFAULT":    0,
	"DEPLOYMENT": 1,
	"JOB":        2,
	"CRONJOB":    3,
}

func (x TaskType) String() string {
	return proto.EnumName(TaskType_name, int32(x))
}
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{2}
}

// Data center status
type DCStatus int32

const (
	DCStatus_AVAILABLE   DCStatus = 0
	DCStatus_UNAVAILABLE DCStatus = 1
)

var DCStatus_name = map[int32]string{
	0: "AVAILABLE",
	1: "UNAVAILABLE",
}
var DCStatus_value = map[string]int32{
	"AVAILABLE":   0,
	"UNAVAILABLE": 1,
}

func (x DCStatus) String() string {
	return proto.EnumName(DCStatus_name, int32(x))
}
func (DCStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{3}
}

// Emtpy Message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// Task Data structure
type Task struct {
	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type TaskType `protobuf:"varint,3,opt,name=type,proto3,enum=common.proto.TaskType" json:"type,omitempty"`
	// Types that are valid to be assigned to TypeData:
	//	*Task_TypeDeployment
	//	*Task_TypeJob
	//	*Task_TypeCronJob
	TypeData             isTask_TypeData `protobuf_oneof:"type_data"`
	DataCenterName       string          `protobuf:"bytes,7,opt,name=data_center_name,json=dataCenterName,proto3" json:"data_center_name,omitempty"`
	Status               TaskStatus      `protobuf:"varint,8,opt,name=status,proto3,enum=common.proto.TaskStatus" json:"status,omitempty"`
	Attributes           *TaskAttributes `protobuf:"bytes,9,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Environment          *Environment    `protobuf:"bytes,10,opt,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{1}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetType() TaskType {
	if m != nil {
		return m.Type
	}
	return TaskType_DEFAULT
}

type isTask_TypeData interface {
	isTask_TypeData()
}

type Task_TypeDeployment struct {
	TypeDeployment *TaskTypeDeployment `protobuf:"bytes,4,opt,name=type_deployment,json=typeDeployment,proto3,oneof"`
}

type Task_TypeJob struct {
	TypeJob *TaskTypeJob `protobuf:"bytes,5,opt,name=type_job,json=typeJob,proto3,oneof"`
}

type Task_TypeCronJob struct {
	TypeCronJob *TaskTypeCronJob `protobuf:"bytes,6,opt,name=type_cron_job,json=typeCronJob,proto3,oneof"`
}

func (*Task_TypeDeployment) isTask_TypeData() {}

func (*Task_TypeJob) isTask_TypeData() {}

func (*Task_TypeCronJob) isTask_TypeData() {}

func (m *Task) GetTypeData() isTask_TypeData {
	if m != nil {
		return m.TypeData
	}
	return nil
}

func (m *Task) GetTypeDeployment() *TaskTypeDeployment {
	if x, ok := m.GetTypeData().(*Task_TypeDeployment); ok {
		return x.TypeDeployment
	}
	return nil
}

func (m *Task) GetTypeJob() *TaskTypeJob {
	if x, ok := m.GetTypeData().(*Task_TypeJob); ok {
		return x.TypeJob
	}
	return nil
}

func (m *Task) GetTypeCronJob() *TaskTypeCronJob {
	if x, ok := m.GetTypeData().(*Task_TypeCronJob); ok {
		return x.TypeCronJob
	}
	return nil
}

func (m *Task) GetDataCenterName() string {
	if m != nil {
		return m.DataCenterName
	}
	return ""
}

func (m *Task) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_STARTING
}

func (m *Task) GetAttributes() *TaskAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Task) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Task) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Task_OneofMarshaler, _Task_OneofUnmarshaler, _Task_OneofSizer, []interface{}{
		(*Task_TypeDeployment)(nil),
		(*Task_TypeJob)(nil),
		(*Task_TypeCronJob)(nil),
	}
}

func _Task_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Task)
	// type_data
	switch x := m.TypeData.(type) {
	case *Task_TypeDeployment:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypeDeployment); err != nil {
			return err
		}
	case *Task_TypeJob:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypeJob); err != nil {
			return err
		}
	case *Task_TypeCronJob:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypeCronJob); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Task.TypeData has unexpected type %T", x)
	}
	return nil
}

func _Task_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Task)
	switch tag {
	case 4: // type_data.type_deployment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TaskTypeDeployment)
		err := b.DecodeMessage(msg)
		m.TypeData = &Task_TypeDeployment{msg}
		return true, err
	case 5: // type_data.type_job
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TaskTypeJob)
		err := b.DecodeMessage(msg)
		m.TypeData = &Task_TypeJob{msg}
		return true, err
	case 6: // type_data.type_cron_job
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TaskTypeCronJob)
		err := b.DecodeMessage(msg)
		m.TypeData = &Task_TypeCronJob{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Task_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Task)
	// type_data
	switch x := m.TypeData.(type) {
	case *Task_TypeDeployment:
		s := proto.Size(x.TypeDeployment)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Task_TypeJob:
		s := proto.Size(x.TypeJob)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Task_TypeCronJob:
		s := proto.Size(x.TypeCronJob)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TaskTypeDeployment struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskTypeDeployment) Reset()         { *m = TaskTypeDeployment{} }
func (m *TaskTypeDeployment) String() string { return proto.CompactTextString(m) }
func (*TaskTypeDeployment) ProtoMessage()    {}
func (*TaskTypeDeployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{2}
}
func (m *TaskTypeDeployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskTypeDeployment.Unmarshal(m, b)
}
func (m *TaskTypeDeployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskTypeDeployment.Marshal(b, m, deterministic)
}
func (dst *TaskTypeDeployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskTypeDeployment.Merge(dst, src)
}
func (m *TaskTypeDeployment) XXX_Size() int {
	return xxx_messageInfo_TaskTypeDeployment.Size(m)
}
func (m *TaskTypeDeployment) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskTypeDeployment.DiscardUnknown(m)
}

var xxx_messageInfo_TaskTypeDeployment proto.InternalMessageInfo

func (m *TaskTypeDeployment) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type TaskTypeJob struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskTypeJob) Reset()         { *m = TaskTypeJob{} }
func (m *TaskTypeJob) String() string { return proto.CompactTextString(m) }
func (*TaskTypeJob) ProtoMessage()    {}
func (*TaskTypeJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{3}
}
func (m *TaskTypeJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskTypeJob.Unmarshal(m, b)
}
func (m *TaskTypeJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskTypeJob.Marshal(b, m, deterministic)
}
func (dst *TaskTypeJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskTypeJob.Merge(dst, src)
}
func (m *TaskTypeJob) XXX_Size() int {
	return xxx_messageInfo_TaskTypeJob.Size(m)
}
func (m *TaskTypeJob) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskTypeJob.DiscardUnknown(m)
}

var xxx_messageInfo_TaskTypeJob proto.InternalMessageInfo

func (m *TaskTypeJob) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type TaskTypeCronJob struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Schedule             string   `protobuf:"bytes,2,opt,name=schedule,proto3" json:"schedule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskTypeCronJob) Reset()         { *m = TaskTypeCronJob{} }
func (m *TaskTypeCronJob) String() string { return proto.CompactTextString(m) }
func (*TaskTypeCronJob) ProtoMessage()    {}
func (*TaskTypeCronJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{4}
}
func (m *TaskTypeCronJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskTypeCronJob.Unmarshal(m, b)
}
func (m *TaskTypeCronJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskTypeCronJob.Marshal(b, m, deterministic)
}
func (dst *TaskTypeCronJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskTypeCronJob.Merge(dst, src)
}
func (m *TaskTypeCronJob) XXX_Size() int {
	return xxx_messageInfo_TaskTypeCronJob.Size(m)
}
func (m *TaskTypeCronJob) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskTypeCronJob.DiscardUnknown(m)
}

var xxx_messageInfo_TaskTypeCronJob proto.InternalMessageInfo

func (m *TaskTypeCronJob) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *TaskTypeCronJob) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

type TaskAttributes struct {
	Replica              int32    `protobuf:"varint,1,opt,name=replica,proto3" json:"replica,omitempty"`
	Hidden               bool     `protobuf:"varint,2,opt,name=hidden,proto3" json:"hidden,omitempty"`
	CreationDate         uint64   `protobuf:"varint,3,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	LastModifiedDate     uint64   `protobuf:"varint,4,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskAttributes) Reset()         { *m = TaskAttributes{} }
func (m *TaskAttributes) String() string { return proto.CompactTextString(m) }
func (*TaskAttributes) ProtoMessage()    {}
func (*TaskAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{5}
}
func (m *TaskAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskAttributes.Unmarshal(m, b)
}
func (m *TaskAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskAttributes.Marshal(b, m, deterministic)
}
func (dst *TaskAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskAttributes.Merge(dst, src)
}
func (m *TaskAttributes) XXX_Size() int {
	return xxx_messageInfo_TaskAttributes.Size(m)
}
func (m *TaskAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_TaskAttributes proto.InternalMessageInfo

func (m *TaskAttributes) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *TaskAttributes) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *TaskAttributes) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *TaskAttributes) GetLastModifiedDate() uint64 {
	if m != nil {
		return m.LastModifiedDate
	}
	return 0
}

type Environment struct {
	Cpu                  int32    `protobuf:"varint,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               int32    `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk                 int32    `protobuf:"varint,3,opt,name=disk,proto3" json:"disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{6}
}
func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (dst *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(dst, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *Environment) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *Environment) GetDisk() int32 {
	if m != nil {
		return m.Disk
	}
	return 0
}

type GeoLocation struct {
	Lat                  string   `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  string   `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoLocation) Reset()         { *m = GeoLocation{} }
func (m *GeoLocation) String() string { return proto.CompactTextString(m) }
func (*GeoLocation) ProtoMessage()    {}
func (*GeoLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{7}
}
func (m *GeoLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoLocation.Unmarshal(m, b)
}
func (m *GeoLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoLocation.Marshal(b, m, deterministic)
}
func (dst *GeoLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoLocation.Merge(dst, src)
}
func (m *GeoLocation) XXX_Size() int {
	return xxx_messageInfo_GeoLocation.Size(m)
}
func (m *GeoLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoLocation.DiscardUnknown(m)
}

var xxx_messageInfo_GeoLocation proto.InternalMessageInfo

func (m *GeoLocation) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *GeoLocation) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

// Data Center structure
type DataCenter struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GeoLocation          *GeoLocation          `protobuf:"bytes,3,opt,name=geo_location,json=geoLocation,proto3" json:"geo_location,omitempty"`
	Status               DCStatus              `protobuf:"varint,4,opt,name=status,proto3,enum=common.proto.DCStatus" json:"status,omitempty"`
	DcAttributes         *DataCenterAttributes `protobuf:"bytes,5,opt,name=dc_attributes,json=dcAttributes,proto3" json:"dc_attributes,omitempty"`
	DcHeartbeatReport    *DCHeartbeatReport    `protobuf:"bytes,6,opt,name=dc_heartbeat_report,json=dcHeartbeatReport,proto3" json:"dc_heartbeat_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DataCenter) Reset()         { *m = DataCenter{} }
func (m *DataCenter) String() string { return proto.CompactTextString(m) }
func (*DataCenter) ProtoMessage()    {}
func (*DataCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{8}
}
func (m *DataCenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenter.Unmarshal(m, b)
}
func (m *DataCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenter.Marshal(b, m, deterministic)
}
func (dst *DataCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenter.Merge(dst, src)
}
func (m *DataCenter) XXX_Size() int {
	return xxx_messageInfo_DataCenter.Size(m)
}
func (m *DataCenter) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenter.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenter proto.InternalMessageInfo

func (m *DataCenter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataCenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenter) GetGeoLocation() *GeoLocation {
	if m != nil {
		return m.GeoLocation
	}
	return nil
}

func (m *DataCenter) GetStatus() DCStatus {
	if m != nil {
		return m.Status
	}
	return DCStatus_AVAILABLE
}

func (m *DataCenter) GetDcAttributes() *DataCenterAttributes {
	if m != nil {
		return m.DcAttributes
	}
	return nil
}

func (m *DataCenter) GetDcHeartbeatReport() *DCHeartbeatReport {
	if m != nil {
		return m.DcHeartbeatReport
	}
	return nil
}

type DataCenterAttributes struct {
	WalletAddress        string   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	CreationDate         uint64   `protobuf:"varint,2,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	LastModifiedDate     uint64   `protobuf:"varint,3,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenterAttributes) Reset()         { *m = DataCenterAttributes{} }
func (m *DataCenterAttributes) String() string { return proto.CompactTextString(m) }
func (*DataCenterAttributes) ProtoMessage()    {}
func (*DataCenterAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{9}
}
func (m *DataCenterAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterAttributes.Unmarshal(m, b)
}
func (m *DataCenterAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterAttributes.Marshal(b, m, deterministic)
}
func (dst *DataCenterAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterAttributes.Merge(dst, src)
}
func (m *DataCenterAttributes) XXX_Size() int {
	return xxx_messageInfo_DataCenterAttributes.Size(m)
}
func (m *DataCenterAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterAttributes proto.InternalMessageInfo

func (m *DataCenterAttributes) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *DataCenterAttributes) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *DataCenterAttributes) GetLastModifiedDate() uint64 {
	if m != nil {
		return m.LastModifiedDate
	}
	return 0
}

type DCHeartbeatReport struct {
	Metrics              string   `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Report               string   `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
	ReportTime           uint64   `protobuf:"varint,3,opt,name=report_time,json=reportTime,proto3" json:"report_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DCHeartbeatReport) Reset()         { *m = DCHeartbeatReport{} }
func (m *DCHeartbeatReport) String() string { return proto.CompactTextString(m) }
func (*DCHeartbeatReport) ProtoMessage()    {}
func (*DCHeartbeatReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{10}
}
func (m *DCHeartbeatReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCHeartbeatReport.Unmarshal(m, b)
}
func (m *DCHeartbeatReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCHeartbeatReport.Marshal(b, m, deterministic)
}
func (dst *DCHeartbeatReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCHeartbeatReport.Merge(dst, src)
}
func (m *DCHeartbeatReport) XXX_Size() int {
	return xxx_messageInfo_DCHeartbeatReport.Size(m)
}
func (m *DCHeartbeatReport) XXX_DiscardUnknown() {
	xxx_messageInfo_DCHeartbeatReport.DiscardUnknown(m)
}

var xxx_messageInfo_DCHeartbeatReport proto.InternalMessageInfo

func (m *DCHeartbeatReport) GetMetrics() string {
	if m != nil {
		return m.Metrics
	}
	return ""
}

func (m *DCHeartbeatReport) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *DCHeartbeatReport) GetReportTime() uint64 {
	if m != nil {
		return m.ReportTime
	}
	return 0
}

// data center communicate with dc manager
type DCStream struct {
	OpType DCOperation `protobuf:"varint,1,opt,name=op_type,json=opType,proto3,enum=common.proto.DCOperation" json:"op_type,omitempty"`
	// Types that are valid to be assigned to OpPayload:
	//	*DCStream_Task
	//	*DCStream_DataCenter
	OpPayload            isDCStream_OpPayload `protobuf_oneof:"op_payload"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DCStream) Reset()         { *m = DCStream{} }
func (m *DCStream) String() string { return proto.CompactTextString(m) }
func (*DCStream) ProtoMessage()    {}
func (*DCStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{11}
}
func (m *DCStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCStream.Unmarshal(m, b)
}
func (m *DCStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCStream.Marshal(b, m, deterministic)
}
func (dst *DCStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCStream.Merge(dst, src)
}
func (m *DCStream) XXX_Size() int {
	return xxx_messageInfo_DCStream.Size(m)
}
func (m *DCStream) XXX_DiscardUnknown() {
	xxx_messageInfo_DCStream.DiscardUnknown(m)
}

var xxx_messageInfo_DCStream proto.InternalMessageInfo

func (m *DCStream) GetOpType() DCOperation {
	if m != nil {
		return m.OpType
	}
	return DCOperation_TASK_CREATE
}

type isDCStream_OpPayload interface {
	isDCStream_OpPayload()
}

type DCStream_Task struct {
	Task *Task `protobuf:"bytes,2,opt,name=task,proto3,oneof"`
}

type DCStream_DataCenter struct {
	DataCenter *DataCenter `protobuf:"bytes,3,opt,name=data_center,json=dataCenter,proto3,oneof"`
}

func (*DCStream_Task) isDCStream_OpPayload() {}

func (*DCStream_DataCenter) isDCStream_OpPayload() {}

func (m *DCStream) GetOpPayload() isDCStream_OpPayload {
	if m != nil {
		return m.OpPayload
	}
	return nil
}

func (m *DCStream) GetTask() *Task {
	if x, ok := m.GetOpPayload().(*DCStream_Task); ok {
		return x.Task
	}
	return nil
}

func (m *DCStream) GetDataCenter() *DataCenter {
	if x, ok := m.GetOpPayload().(*DCStream_DataCenter); ok {
		return x.DataCenter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DCStream) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DCStream_OneofMarshaler, _DCStream_OneofUnmarshaler, _DCStream_OneofSizer, []interface{}{
		(*DCStream_Task)(nil),
		(*DCStream_DataCenter)(nil),
	}
}

func _DCStream_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DCStream)
	// op_payload
	switch x := m.OpPayload.(type) {
	case *DCStream_Task:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Task); err != nil {
			return err
		}
	case *DCStream_DataCenter:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DataCenter); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DCStream.OpPayload has unexpected type %T", x)
	}
	return nil
}

func _DCStream_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DCStream)
	switch tag {
	case 2: // op_payload.task
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Task)
		err := b.DecodeMessage(msg)
		m.OpPayload = &DCStream_Task{msg}
		return true, err
	case 3: // op_payload.data_center
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DataCenter)
		err := b.DecodeMessage(msg)
		m.OpPayload = &DCStream_DataCenter{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DCStream_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DCStream)
	// op_payload
	switch x := m.OpPayload.(type) {
	case *DCStream_Task:
		s := proto.Size(x.Task)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DCStream_DataCenter:
		s := proto.Size(x.DataCenter)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Mail For the future
type Mail struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mail) Reset()         { *m = Mail{} }
func (m *Mail) String() string { return proto.CompactTextString(m) }
func (*Mail) ProtoMessage()    {}
func (*Mail) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_35578ab15f2c9754, []int{12}
}
func (m *Mail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mail.Unmarshal(m, b)
}
func (m *Mail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mail.Marshal(b, m, deterministic)
}
func (dst *Mail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mail.Merge(dst, src)
}
func (m *Mail) XXX_Size() int {
	return xxx_messageInfo_Mail.Size(m)
}
func (m *Mail) XXX_DiscardUnknown() {
	xxx_messageInfo_Mail.DiscardUnknown(m)
}

var xxx_messageInfo_Mail proto.InternalMessageInfo

func (m *Mail) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *Mail) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Mail) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Mail) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "common.proto.Empty")
	proto.RegisterType((*Task)(nil), "common.proto.Task")
	proto.RegisterType((*TaskTypeDeployment)(nil), "common.proto.TaskTypeDeployment")
	proto.RegisterType((*TaskTypeJob)(nil), "common.proto.TaskTypeJob")
	proto.RegisterType((*TaskTypeCronJob)(nil), "common.proto.TaskTypeCronJob")
	proto.RegisterType((*TaskAttributes)(nil), "common.proto.TaskAttributes")
	proto.RegisterType((*Environment)(nil), "common.proto.Environment")
	proto.RegisterType((*GeoLocation)(nil), "common.proto.GeoLocation")
	proto.RegisterType((*DataCenter)(nil), "common.proto.DataCenter")
	proto.RegisterType((*DataCenterAttributes)(nil), "common.proto.DataCenterAttributes")
	proto.RegisterType((*DCHeartbeatReport)(nil), "common.proto.DCHeartbeatReport")
	proto.RegisterType((*DCStream)(nil), "common.proto.DCStream")
	proto.RegisterType((*Mail)(nil), "common.proto.Mail")
	proto.RegisterEnum("common.proto.DCOperation", DCOperation_name, DCOperation_value)
	proto.RegisterEnum("common.proto.TaskStatus", TaskStatus_name, TaskStatus_value)
	proto.RegisterEnum("common.proto.TaskType", TaskType_name, TaskType_value)
	proto.RegisterEnum("common.proto.DCStatus", DCStatus_name, DCStatus_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_35578ab15f2c9754) }

var fileDescriptor_common_35578ab15f2c9754 = []byte{
	// 1003 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xe3, 0x36,
	0x10, 0xb5, 0x6c, 0xf9, 0x6b, 0x64, 0x3b, 0x0a, 0xbb, 0x58, 0x78, 0x8b, 0x16, 0x1b, 0x68, 0x51,
	0x20, 0x30, 0x8a, 0xa0, 0x4d, 0x81, 0x5e, 0xba, 0x3d, 0x28, 0xb6, 0x36, 0xce, 0xc6, 0xb1, 0x17,
	0xb4, 0x52, 0xa0, 0x27, 0x81, 0x96, 0x98, 0x44, 0x8d, 0x24, 0x0a, 0x32, 0xd3, 0xc2, 0xf7, 0xfe,
	0x80, 0x5e, 0xfa, 0x4b, 0x7a, 0xe9, 0xa5, 0xff, 0xad, 0xe0, 0x87, 0x6d, 0xc5, 0x76, 0x80, 0x3d,
	0x69, 0x66, 0xf4, 0x66, 0xf8, 0xf8, 0xf8, 0x48, 0xe8, 0x84, 0x2c, 0x4d, 0x59, 0x76, 0x96, 0x17,
	0x8c, 0x33, 0xf4, 0x2c, 0x73, 0x9a, 0x50, 0xf7, 0xd2, 0x9c, 0xaf, 0x9c, 0x3f, 0x4d, 0x30, 0x7d,
	0xb2, 0x7c, 0x44, 0x3d, 0xa8, 0xc6, 0x51, 0xdf, 0x38, 0x31, 0x4e, 0xdb, 0xb8, 0x1a, 0x47, 0x08,
	0x81, 0x99, 0x91, 0x94, 0xf6, 0xab, 0xb2, 0x22, 0x63, 0x34, 0x00, 0x93, 0xaf, 0x72, 0xda, 0xaf,
	0x9d, 0x18, 0xa7, 0xbd, 0xf3, 0xd7, 0x67, 0xe5, 0x91, 0x67, 0x62, 0x8a, 0xbf, 0xca, 0x29, 0x96,
	0x18, 0x74, 0x0d, 0x47, 0xe2, 0x1b, 0x44, 0x34, 0x4f, 0xd8, 0x2a, 0xa5, 0x19, 0xef, 0x9b, 0x27,
	0xc6, 0xa9, 0x75, 0x7e, 0x72, 0xb8, 0x6d, 0xb4, 0xc1, 0x8d, 0x2b, 0xb8, 0xc7, 0x9f, 0x55, 0xd0,
	0x8f, 0xd0, 0x92, 0xc3, 0x7e, 0x63, 0x8b, 0x7e, 0x5d, 0x4e, 0x79, 0x73, 0x78, 0xca, 0x47, 0xb6,
	0x18, 0x57, 0x70, 0x93, 0xab, 0x10, 0x0d, 0xa1, 0x2b, 0xfb, 0xc2, 0x82, 0x65, 0xb2, 0xb9, 0x21,
	0x9b, 0xbf, 0x3e, 0xdc, 0x3c, 0x2c, 0x58, 0xa6, 0x06, 0x58, 0x7c, 0x9b, 0xa2, 0x53, 0xb0, 0x23,
	0xc2, 0x49, 0x10, 0xd2, 0x8c, 0xd3, 0x22, 0x90, 0xaa, 0x34, 0xa5, 0x2a, 0x3d, 0x51, 0x1f, 0xca,
	0xf2, 0x54, 0xe8, 0xf3, 0x1d, 0x34, 0x96, 0x9c, 0xf0, 0xa7, 0x65, 0xbf, 0x25, 0x15, 0xea, 0xef,
	0xaf, 0x33, 0x97, 0xff, 0xb1, 0xc6, 0xa1, 0xf7, 0x00, 0x84, 0xf3, 0x22, 0x5e, 0x3c, 0x71, 0xba,
	0xec, 0xb7, 0x25, 0xbb, 0xaf, 0xf6, 0xbb, 0xdc, 0x0d, 0x06, 0x97, 0xf0, 0xe8, 0x27, 0xb0, 0x68,
	0xf6, 0x7b, 0x5c, 0xb0, 0x4c, 0xea, 0x0b, 0x87, 0x94, 0xf1, 0xb6, 0x00, 0x5c, 0x46, 0x5f, 0x58,
	0xd0, 0x56, 0x07, 0x44, 0x38, 0x71, 0x06, 0x80, 0xf6, 0x0f, 0x02, 0xbd, 0x82, 0x7a, 0x9c, 0x92,
	0x7b, 0xaa, 0x6d, 0xa1, 0x12, 0xe7, 0x1d, 0x58, 0x25, 0xb9, 0x5f, 0x00, 0x0d, 0xe1, 0x68, 0x47,
	0xd6, 0xc3, 0x40, 0xf4, 0x25, 0xb4, 0x96, 0xe1, 0x03, 0x8d, 0x9e, 0x92, 0xb5, 0xd7, 0x36, 0xb9,
	0xf3, 0xb7, 0x01, 0xbd, 0xe7, 0xdb, 0x47, 0x7d, 0x68, 0x16, 0x34, 0x4f, 0xe2, 0x90, 0xc8, 0x31,
	0x75, 0xbc, 0x4e, 0xd1, 0x6b, 0x68, 0x3c, 0xc4, 0x51, 0x44, 0x33, 0x39, 0xa6, 0x85, 0x75, 0x86,
	0xde, 0x41, 0x37, 0x2c, 0x28, 0xe1, 0x31, 0xcb, 0xc4, 0x5e, 0x95, 0x7b, 0x4d, 0xdc, 0x59, 0x17,
	0x47, 0x84, 0x53, 0xf4, 0x2d, 0xa0, 0x84, 0x2c, 0x79, 0x90, 0xb2, 0x28, 0xbe, 0x8b, 0x69, 0xa4,
	0x90, 0xa6, 0x44, 0xda, 0xe2, 0xcf, 0x8d, 0xfe, 0x21, 0xd0, 0xce, 0x35, 0x58, 0x25, 0x59, 0x91,
	0x0d, 0xb5, 0x30, 0x7f, 0xd2, 0x7c, 0x44, 0x28, 0xb8, 0xa4, 0x34, 0x65, 0xc5, 0x4a, 0x72, 0xa9,
	0x63, 0x9d, 0x89, 0x4b, 0x15, 0xc5, 0xcb, 0x47, 0x49, 0xa1, 0x8e, 0x65, 0xec, 0x7c, 0x0f, 0xd6,
	0x25, 0x65, 0x13, 0x16, 0x4a, 0x36, 0x62, 0x58, 0x42, 0xb8, 0xd6, 0x48, 0x84, 0xb2, 0x92, 0xdd,
	0x6b, 0x71, 0x44, 0xe8, 0xfc, 0x5b, 0x05, 0x18, 0x6d, 0xac, 0xf7, 0x59, 0x57, 0xf7, 0x3d, 0x74,
	0xee, 0x29, 0x0b, 0x12, 0xbd, 0x8c, 0x64, 0xb0, 0xe7, 0x95, 0x12, 0x0f, 0x6c, 0xdd, 0x97, 0x48,
	0x9d, 0x6d, 0x8c, 0x6d, 0x1e, 0xba, 0xfa, 0xa3, 0xe1, 0x8e, 0xad, 0x2f, 0xa1, 0x1b, 0x85, 0x41,
	0xc9, 0xd9, 0xea, 0xd2, 0x3a, 0x3b, 0x6d, 0x9b, 0x2d, 0x94, 0xfc, 0xdd, 0x89, 0xc2, 0xd2, 0x71,
	0xcf, 0xe0, 0x8b, 0x28, 0x0c, 0x1e, 0x28, 0x29, 0xf8, 0x82, 0x12, 0x1e, 0x14, 0x34, 0x67, 0x05,
	0xd7, 0xd7, 0xf8, 0xed, 0x2e, 0x8b, 0xf1, 0x1a, 0x87, 0x25, 0x0c, 0x1f, 0x47, 0xe1, 0x4e, 0xc9,
	0xf9, 0xcb, 0x80, 0x57, 0x87, 0xd6, 0x45, 0xdf, 0x40, 0xef, 0x0f, 0x92, 0x24, 0x94, 0x07, 0x24,
	0x8a, 0x0a, 0xba, 0x5c, 0x6a, 0x41, 0xbb, 0xaa, 0xea, 0xaa, 0xe2, 0xbe, 0x9b, 0xaa, 0x9f, 0xed,
	0xa6, 0xda, 0x0b, 0x6e, 0xba, 0x83, 0xe3, 0x3d, 0xea, 0xc2, 0xe7, 0x29, 0xe5, 0x45, 0x1c, 0xae,
	0x79, 0xac, 0x53, 0xe1, 0x2d, 0xad, 0x82, 0x3a, 0x5f, 0x9d, 0xa1, 0xb7, 0x60, 0xa9, 0x28, 0xe0,
	0x71, 0xba, 0x5e, 0x0d, 0x54, 0xc9, 0x8f, 0x53, 0xea, 0xfc, 0x63, 0x40, 0x4b, 0x9c, 0x54, 0x41,
	0x49, 0x8a, 0xce, 0xa1, 0xc9, 0xf2, 0x40, 0xbe, 0xe6, 0x86, 0x3c, 0xd2, 0x37, 0xbb, 0x62, 0xce,
	0x72, 0x5a, 0x28, 0x2b, 0x34, 0x58, 0x2e, 0xee, 0x31, 0x3a, 0x05, 0x93, 0x93, 0xe5, 0xa3, 0x5c,
	0xd7, 0x3a, 0x47, 0xfb, 0xcf, 0xd4, 0xb8, 0x82, 0x25, 0x42, 0x3c, 0x4c, 0xa5, 0x27, 0x53, 0x9b,
	0xad, 0xff, 0xd2, 0xe9, 0x8f, 0x2b, 0x18, 0xb6, 0x2f, 0xe9, 0x45, 0x07, 0x80, 0xe5, 0x41, 0x4e,
	0x56, 0x09, 0x23, 0x91, 0xe3, 0x83, 0x79, 0x43, 0xe2, 0x44, 0x98, 0xfa, 0xae, 0x60, 0xa9, 0x56,
	0x43, 0xc6, 0xc2, 0xf8, 0x9c, 0x69, 0x19, 0xaa, 0x9c, 0x89, 0x17, 0x86, 0xc7, 0x3c, 0x51, 0x9b,
	0x6f, 0x63, 0x95, 0x88, 0xce, 0x05, 0x8b, 0x56, 0xd2, 0xba, 0x6d, 0x2c, 0xe3, 0xc1, 0x0c, 0xac,
	0xd2, 0x0e, 0xd1, 0x11, 0x58, 0xbe, 0x3b, 0xbf, 0x0e, 0x86, 0xd8, 0x73, 0x7d, 0xcf, 0xae, 0x6c,
	0x0b, 0xee, 0x74, 0xe8, 0x4d, 0x6c, 0x63, 0x53, 0xb8, 0xfd, 0x34, 0x12, 0x88, 0x2a, 0xea, 0x42,
	0x7b, 0xec, 0xb9, 0xd8, 0xbf, 0xf0, 0x5c, 0xdf, 0xae, 0x0d, 0xfe, 0x33, 0x00, 0xb6, 0xef, 0x3b,
	0xea, 0x40, 0x6b, 0xee, 0xbb, 0xd8, 0xbf, 0x9a, 0x5e, 0xda, 0x15, 0x74, 0x0c, 0x5d, 0x99, 0x05,
	0xf3, 0xdb, 0xe1, 0xd0, 0x9b, 0xcf, 0x6d, 0x03, 0xd9, 0xd0, 0x51, 0xa5, 0x0f, 0xee, 0xd5, 0xc4,
	0x1b, 0xd9, 0x55, 0x64, 0x41, 0x13, 0xdf, 0x4e, 0xa7, 0xa2, 0xa3, 0x26, 0xfa, 0xe5, 0x4a, 0x22,
	0x33, 0x11, 0x82, 0x9e, 0x5a, 0x77, 0x33, 0xa0, 0x2e, 0x66, 0xea, 0x9a, 0x9e, 0xd0, 0x40, 0x3d,
	0x00, 0xc5, 0x77, 0x22, 0xda, 0x9a, 0x82, 0xa2, 0xce, 0xbd, 0x91, 0xdd, 0x12, 0x1d, 0x2a, 0x5d,
	0x77, 0xb4, 0x51, 0x0b, 0xcc, 0xd1, 0x6c, 0xea, 0xd9, 0x30, 0xf8, 0x19, 0x5a, 0xeb, 0xf7, 0x5a,
	0x30, 0x19, 0x79, 0x1f, 0xdc, 0xdb, 0x89, 0x6f, 0x57, 0xc4, 0xd0, 0x91, 0xf7, 0x69, 0x32, 0xfb,
	0xf5, 0xc6, 0x9b, 0xfa, 0xb6, 0x81, 0x9a, 0x50, 0xfb, 0x38, 0xbb, 0x50, 0x7c, 0x87, 0x78, 0x36,
	0x15, 0x49, 0x6d, 0x30, 0x50, 0xd6, 0x92, 0x7b, 0xef, 0x42, 0xdb, 0xfd, 0xc5, 0xbd, 0x9a, 0xb8,
	0x17, 0x13, 0x2d, 0xe5, 0xed, 0x74, 0x5b, 0x30, 0x16, 0x0d, 0x79, 0xfe, 0x3f, 0xfc, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0xa6, 0xb6, 0xe3, 0xa0, 0x08, 0x00, 0x00,
}
