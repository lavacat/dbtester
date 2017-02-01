// Code generated by protoc-gen-gogo.
// source: agent/agentpb/message.proto
// DO NOT EDIT!

/*
	Package agentpb is a generated protocol buffer package.

	It is generated from these files:
		agent/agentpb/message.proto

	It has these top-level messages:
		Request
		Response
*/
package agentpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request_Operation int32

const (
	Request_Start     Request_Operation = 0
	Request_Stop      Request_Operation = 1
	Request_Heartbeat Request_Operation = 2
)

var Request_Operation_name = map[int32]string{
	0: "Start",
	1: "Stop",
	2: "Heartbeat",
}
var Request_Operation_value = map[string]int32{
	"Start":     0,
	"Stop":      1,
	"Heartbeat": 2,
}

func (x Request_Operation) String() string {
	return proto.EnumName(Request_Operation_name, int32(x))
}
func (Request_Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 0} }

type Request_Database int32

const (
	Request_etcdv2    Request_Database = 0
	Request_etcdv3    Request_Database = 1
	Request_ZooKeeper Request_Database = 2
	Request_Consul    Request_Database = 3
	Request_zetcd     Request_Database = 4
	Request_cetcd     Request_Database = 5
)

var Request_Database_name = map[int32]string{
	0: "etcdv2",
	1: "etcdv3",
	2: "ZooKeeper",
	3: "Consul",
	4: "zetcd",
	5: "cetcd",
}
var Request_Database_value = map[string]int32{
	"etcdv2":    0,
	"etcdv3":    1,
	"ZooKeeper": 2,
	"Consul":    3,
	"zetcd":     4,
	"cetcd":     5,
}

func (x Request_Database) String() string {
	return proto.EnumName(Request_Database_name, int32(x))
}
func (Request_Database) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 1} }

type Request struct {
	Operation    Request_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=agentpb.Request_Operation" json:"operation,omitempty"`
	Database     Request_Database  `protobuf:"varint,2,opt,name=database,proto3,enum=agentpb.Request_Database" json:"database,omitempty"`
	PeerIPString string            `protobuf:"bytes,3,opt,name=peerIPString,proto3" json:"peerIPString,omitempty"`
	// ServerIPIndex is the index in peerIPs that points to the
	// corresponding remote IP.
	ServerIndex uint32 `protobuf:"varint,4,opt,name=serverIndex,proto3" json:"serverIndex,omitempty"`
	// TestName prefixes all logs to be generated in agent.
	TestName string `protobuf:"bytes,5,opt,name=testName,proto3" json:"testName,omitempty"`
	// GoogleCloudProjectName is the project name to use
	// to upload logs.
	GoogleCloudProjectName string `protobuf:"bytes,6,opt,name=googleCloudProjectName,proto3" json:"googleCloudProjectName,omitempty"`
	// GoogleCloudStorageKey is the key to be used to upload
	// data and logs to Google Cloud Storage and others.
	GoogleCloudStorageKey string `protobuf:"bytes,7,opt,name=googleCloudStorageKey,proto3" json:"googleCloudStorageKey,omitempty"`
	// GoogleCloudStorageBucketName is the bucket name to store all data and logs.
	GoogleCloudStorageBucketName string `protobuf:"bytes,8,opt,name=googleCloudStorageBucketName,proto3" json:"googleCloudStorageBucketName,omitempty"`
	// GoogleCloudStorageSubDirectory is the sub-directory name to store data.
	GoogleCloudStorageSubDirectory string `protobuf:"bytes,9,opt,name=googleCloudStorageSubDirectory,proto3" json:"googleCloudStorageSubDirectory,omitempty"`
	// ZookeeperMyID is myid that needs to be stored as a file in the remote machine.
	ZookeeperMyID uint32 `protobuf:"varint,10,opt,name=zookeeperMyID,proto3" json:"zookeeperMyID,omitempty"`
	// EtcdSnapCount is 100,000 by default.
	EtcdSnapCount int64 `protobuf:"varint,11,opt,name=etcdSnapCount,proto3" json:"etcdSnapCount,omitempty"`
	// EtcdQuotaSizeBytes is the backend size limit in bytes.
	// 0 defaults to low space quota (2 GB).
	EtcdQuotaSizeBytes int64 `protobuf:"varint,12,opt,name=etcdQuotaSizeBytes,proto3" json:"etcdQuotaSizeBytes,omitempty"`
	// ZookeeperSnapCount is 100,000 by default.
	ZookeeperSnapCount int64 `protobuf:"varint,13,opt,name=zookeeperSnapCount,proto3" json:"zookeeperSnapCount,omitempty"`
	// ZookeeperMaxClientCnxns limits the number of concurrent connections
	// (at the socket level) that a single client, identified by IP address.
	ZookeeperMaxClientCnxns int64 `protobuf:"varint,14,opt,name=zookeeperMaxClientCnxns,proto3" json:"zookeeperMaxClientCnxns,omitempty"`
	// ClientNum is current number of clients.
	ClientNum int64 `protobuf:"varint,15,opt,name=clientNum,proto3" json:"clientNum,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

type Response struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Datasize is the data size of the database on disk.
	// It measures after database is requested to stop.
	Datasize int64 `protobuf:"varint,2,opt,name=datasize,proto3" json:"datasize,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "agentpb.Request")
	proto.RegisterType((*Response)(nil), "agentpb.Response")
	proto.RegisterEnum("agentpb.Request_Operation", Request_Operation_name, Request_Operation_value)
	proto.RegisterEnum("agentpb.Request_Database", Request_Database_name, Request_Database_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Transporter service

type TransporterClient interface {
	Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type transporterClient struct {
	cc *grpc.ClientConn
}

func NewTransporterClient(cc *grpc.ClientConn) TransporterClient {
	return &transporterClient{cc}
}

func (c *transporterClient) Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/agentpb.Transporter/Transfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Transporter service

type TransporterServer interface {
	Transfer(context.Context, *Request) (*Response, error)
}

func RegisterTransporterServer(s *grpc.Server, srv TransporterServer) {
	s.RegisterService(&_Transporter_serviceDesc, srv)
}

func _Transporter_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransporterServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agentpb.Transporter/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransporterServer).Transfer(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agentpb.Transporter",
	HandlerType: (*TransporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Transporter_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent/agentpb/message.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Operation != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Operation))
	}
	if m.Database != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Database))
	}
	if len(m.PeerIPString) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.PeerIPString)))
		i += copy(dAtA[i:], m.PeerIPString)
	}
	if m.ServerIndex != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ServerIndex))
	}
	if len(m.TestName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.TestName)))
		i += copy(dAtA[i:], m.TestName)
	}
	if len(m.GoogleCloudProjectName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudProjectName)))
		i += copy(dAtA[i:], m.GoogleCloudProjectName)
	}
	if len(m.GoogleCloudStorageKey) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudStorageKey)))
		i += copy(dAtA[i:], m.GoogleCloudStorageKey)
	}
	if len(m.GoogleCloudStorageBucketName) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudStorageBucketName)))
		i += copy(dAtA[i:], m.GoogleCloudStorageBucketName)
	}
	if len(m.GoogleCloudStorageSubDirectory) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.GoogleCloudStorageSubDirectory)))
		i += copy(dAtA[i:], m.GoogleCloudStorageSubDirectory)
	}
	if m.ZookeeperMyID != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ZookeeperMyID))
	}
	if m.EtcdSnapCount != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.EtcdSnapCount))
	}
	if m.EtcdQuotaSizeBytes != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.EtcdQuotaSizeBytes))
	}
	if m.ZookeeperSnapCount != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ZookeeperSnapCount))
	}
	if m.ZookeeperMaxClientCnxns != 0 {
		dAtA[i] = 0x70
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ZookeeperMaxClientCnxns))
	}
	if m.ClientNum != 0 {
		dAtA[i] = 0x78
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ClientNum))
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Datasize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Datasize))
	}
	return i, nil
}

func encodeFixed64Message(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Message(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Operation != 0 {
		n += 1 + sovMessage(uint64(m.Operation))
	}
	if m.Database != 0 {
		n += 1 + sovMessage(uint64(m.Database))
	}
	l = len(m.PeerIPString)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ServerIndex != 0 {
		n += 1 + sovMessage(uint64(m.ServerIndex))
	}
	l = len(m.TestName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudProjectName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudStorageKey)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudStorageBucketName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.GoogleCloudStorageSubDirectory)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ZookeeperMyID != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperMyID))
	}
	if m.EtcdSnapCount != 0 {
		n += 1 + sovMessage(uint64(m.EtcdSnapCount))
	}
	if m.EtcdQuotaSizeBytes != 0 {
		n += 1 + sovMessage(uint64(m.EtcdQuotaSizeBytes))
	}
	if m.ZookeeperSnapCount != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperSnapCount))
	}
	if m.ZookeeperMaxClientCnxns != 0 {
		n += 1 + sovMessage(uint64(m.ZookeeperMaxClientCnxns))
	}
	if m.ClientNum != 0 {
		n += 1 + sovMessage(uint64(m.ClientNum))
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.Datasize != 0 {
		n += 1 + sovMessage(uint64(m.Datasize))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operation |= (Request_Operation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			m.Database = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Database |= (Request_Database(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerIPString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerIPString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerIndex", wireType)
			}
			m.ServerIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerIndex |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudProjectName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudProjectName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudStorageKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudStorageKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudStorageBucketName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudStorageBucketName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoogleCloudStorageSubDirectory", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoogleCloudStorageSubDirectory = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperMyID", wireType)
			}
			m.ZookeeperMyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZookeeperMyID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdSnapCount", wireType)
			}
			m.EtcdSnapCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EtcdSnapCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdQuotaSizeBytes", wireType)
			}
			m.EtcdQuotaSizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EtcdQuotaSizeBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperSnapCount", wireType)
			}
			m.ZookeeperSnapCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZookeeperSnapCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZookeeperMaxClientCnxns", wireType)
			}
			m.ZookeeperMaxClientCnxns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ZookeeperMaxClientCnxns |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientNum", wireType)
			}
			m.ClientNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientNum |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Datasize", wireType)
			}
			m.Datasize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Datasize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMessage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("agent/agentpb/message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0xa6, 0x4d, 0xec, 0x69, 0x53, 0xcc, 0x8a, 0x3f, 0x4b, 0xa8, 0xa2, 0xc8, 0xe2,
	0x90, 0x0b, 0x8e, 0x68, 0x01, 0xf5, 0x58, 0x25, 0x15, 0xa2, 0xaa, 0x28, 0xc5, 0xe6, 0xc4, 0x6d,
	0xed, 0x4c, 0x8d, 0x69, 0xb2, 0x6b, 0x76, 0xd7, 0x55, 0x93, 0x27, 0xe1, 0x91, 0x7a, 0xe4, 0x11,
	0x4a, 0x79, 0x11, 0xe4, 0x75, 0xea, 0xb4, 0xa4, 0x85, 0x8b, 0x35, 0xf3, 0x7d, 0xbf, 0x99, 0x59,
	0x5b, 0x3b, 0x86, 0xe7, 0x2c, 0x41, 0xae, 0xfb, 0xe6, 0x99, 0x45, 0xfd, 0x09, 0x2a, 0xc5, 0x12,
	0xf4, 0x33, 0x29, 0xb4, 0x20, 0xcd, 0xb9, 0xdc, 0x7e, 0x99, 0xa4, 0xfa, 0x6b, 0x1e, 0xf9, 0xb1,
	0x98, 0xf4, 0x13, 0x91, 0x88, 0xbe, 0xf1, 0xa3, 0xfc, 0xc4, 0x64, 0x26, 0x31, 0x51, 0x59, 0xe7,
	0x5d, 0x36, 0xa0, 0x19, 0xe0, 0xf7, 0x1c, 0x95, 0x26, 0xbb, 0xe0, 0x88, 0x0c, 0x25, 0xd3, 0xa9,
	0xe0, 0xd4, 0xea, 0x5a, 0xbd, 0xcd, 0xed, 0xb6, 0x3f, 0xef, 0xeb, 0xcf, 0x21, 0xff, 0xe3, 0x35,
	0x11, 0x2c, 0x60, 0xf2, 0x06, 0xec, 0x11, 0xd3, 0x2c, 0x62, 0x0a, 0xe9, 0x8a, 0x29, 0x7c, 0xb6,
	0x54, 0xb8, 0x3f, 0x07, 0x82, 0x0a, 0x25, 0x1e, 0x6c, 0x64, 0x88, 0xf2, 0xe0, 0x38, 0xd4, 0x32,
	0xe5, 0x09, 0xad, 0x77, 0xad, 0x9e, 0x13, 0xdc, 0xd2, 0x48, 0x17, 0xd6, 0x15, 0xca, 0x33, 0x94,
	0x07, 0x7c, 0x84, 0xe7, 0x74, 0xb5, 0x6b, 0xf5, 0x5a, 0xc1, 0x4d, 0x89, 0xb4, 0xc1, 0xd6, 0xa8,
	0xf4, 0x11, 0x9b, 0x20, 0x5d, 0x33, 0x1d, 0xaa, 0x9c, 0xbc, 0x85, 0x27, 0x89, 0x10, 0xc9, 0x18,
	0x87, 0x63, 0x91, 0x8f, 0x8e, 0xa5, 0xf8, 0x86, 0x71, 0x49, 0x36, 0x0c, 0x79, 0x8f, 0x4b, 0x5e,
	0xc3, 0xe3, 0x1b, 0x4e, 0xa8, 0x85, 0x64, 0x09, 0x1e, 0xe2, 0x94, 0x36, 0x4d, 0xd9, 0xdd, 0x26,
	0x19, 0xc0, 0xd6, 0xb2, 0x31, 0xc8, 0xe3, 0x53, 0x2c, 0x67, 0xda, 0xa6, 0xf8, 0x9f, 0x0c, 0x79,
	0x07, 0x9d, 0x65, 0x3f, 0xcc, 0xa3, 0xfd, 0x54, 0x62, 0xac, 0x85, 0x9c, 0x52, 0xc7, 0x74, 0xf9,
	0x0f, 0x45, 0x5e, 0x40, 0x6b, 0x26, 0xc4, 0x29, 0x62, 0x86, 0xf2, 0xc3, 0xf4, 0x60, 0x9f, 0x82,
	0xf9, 0x72, 0xb7, 0xc5, 0x82, 0x42, 0x1d, 0x8f, 0x42, 0xce, 0xb2, 0xa1, 0xc8, 0xb9, 0xa6, 0xeb,
	0x5d, 0xab, 0x57, 0x0f, 0x6e, 0x8b, 0xc4, 0x07, 0x52, 0x08, 0x9f, 0x72, 0xa1, 0x59, 0x98, 0xce,
	0x70, 0x30, 0xd5, 0xa8, 0xe8, 0x86, 0x41, 0xef, 0x70, 0x0a, 0xbe, 0x1a, 0xb3, 0x68, 0xdd, 0x2a,
	0xf9, 0x65, 0x87, 0xec, 0xc2, 0xd3, 0xc5, 0xb1, 0xd8, 0xf9, 0x70, 0x9c, 0x22, 0xd7, 0x43, 0x7e,
	0xce, 0x15, 0xdd, 0x34, 0x45, 0xf7, 0xd9, 0x64, 0x0b, 0x9c, 0xd8, 0xa4, 0x47, 0xf9, 0x84, 0x3e,
	0x30, 0xec, 0x42, 0xf0, 0xfa, 0xe0, 0x54, 0xd7, 0x95, 0x38, 0xb0, 0x16, 0x6a, 0x26, 0xb5, 0x5b,
	0x23, 0x36, 0xac, 0x86, 0x5a, 0x64, 0xae, 0x45, 0x5a, 0xe0, 0xbc, 0x47, 0x26, 0x75, 0x84, 0x4c,
	0xbb, 0x2b, 0x5e, 0x08, 0xf6, 0xf5, 0x35, 0x25, 0x00, 0x8d, 0xe2, 0xd5, 0xce, 0xb6, 0xdd, 0x5a,
	0x15, 0xef, 0x94, 0x25, 0x5f, 0x84, 0x38, 0x34, 0xa7, 0x71, 0x57, 0x0a, 0x6b, 0x28, 0xb8, 0xca,
	0xc7, 0x6e, 0xbd, 0x18, 0x31, 0x2b, 0x38, 0x77, 0xb5, 0x08, 0x63, 0x13, 0xae, 0x79, 0x7b, 0x60,
	0x07, 0xa8, 0x32, 0xc1, 0x15, 0x12, 0x0a, 0x4d, 0x95, 0xc7, 0x31, 0x2a, 0x65, 0x16, 0xcc, 0x0e,
	0xae, 0xd3, 0xe2, 0x16, 0x17, 0x7b, 0xa1, 0xd2, 0x59, 0xb9, 0x42, 0xf5, 0xa0, 0xca, 0xb7, 0xf7,
	0x60, 0xfd, 0xb3, 0x64, 0x5c, 0x65, 0x42, 0x6a, 0x94, 0xe4, 0x15, 0xd8, 0x26, 0x3d, 0x41, 0x49,
	0xdc, 0xbf, 0xf7, 0xac, 0xfd, 0xf0, 0x86, 0x52, 0x4e, 0xf5, 0x6a, 0x83, 0x47, 0x17, 0xbf, 0x3a,
	0xb5, 0x8b, 0xab, 0x8e, 0xf5, 0xf3, 0xaa, 0x63, 0x5d, 0x5e, 0x75, 0xac, 0x1f, 0xbf, 0x3b, 0xb5,
	0xa8, 0x61, 0xfe, 0x01, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x31, 0x68, 0xb7, 0xb3, 0x5a,
	0x04, 0x00, 0x00,
}
