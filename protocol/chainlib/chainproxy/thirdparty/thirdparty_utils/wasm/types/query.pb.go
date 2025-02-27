package types

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_tendermint_tendermint_libs_bytes "github.com/tendermint/tendermint/libs/bytes"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryContractInfoRequest is the request type for the Query/ContractInfo RPC
// method
type QueryContractInfoRequest struct {
	// address is the address of the contract to query
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryContractInfoRequest) Reset()         { *m = QueryContractInfoRequest{} }
func (m *QueryContractInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractInfoRequest) ProtoMessage()    {}
func (*QueryContractInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{0}
}

func (m *QueryContractInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractInfoRequest.Merge(m, src)
}

func (m *QueryContractInfoRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractInfoRequest proto.InternalMessageInfo

// QueryContractInfoResponse is the response type for the Query/ContractInfo RPC
// method
type QueryContractInfoResponse struct {
	// address is the address of the contract
	Address      string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ContractInfo `protobuf:"bytes,2,opt,name=contract_info,json=contractInfo,proto3,embedded=contract_info" json:""`
}

func (m *QueryContractInfoResponse) Reset()         { *m = QueryContractInfoResponse{} }
func (m *QueryContractInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractInfoResponse) ProtoMessage()    {}
func (*QueryContractInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{1}
}

func (m *QueryContractInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractInfoResponse.Merge(m, src)
}

func (m *QueryContractInfoResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractInfoResponse proto.InternalMessageInfo

// QueryContractHistoryRequest is the request type for the Query/ContractHistory
// RPC method
type QueryContractHistoryRequest struct {
	// address is the address of the contract to query
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContractHistoryRequest) Reset()         { *m = QueryContractHistoryRequest{} }
func (m *QueryContractHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractHistoryRequest) ProtoMessage()    {}
func (*QueryContractHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{2}
}

func (m *QueryContractHistoryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractHistoryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractHistoryRequest.Merge(m, src)
}

func (m *QueryContractHistoryRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractHistoryRequest proto.InternalMessageInfo

// QueryContractHistoryResponse is the response type for the
// Query/ContractHistory RPC method
type QueryContractHistoryResponse struct {
	Entries []ContractCodeHistoryEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContractHistoryResponse) Reset()         { *m = QueryContractHistoryResponse{} }
func (m *QueryContractHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractHistoryResponse) ProtoMessage()    {}
func (*QueryContractHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{3}
}

func (m *QueryContractHistoryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractHistoryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractHistoryResponse.Merge(m, src)
}

func (m *QueryContractHistoryResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractHistoryResponse proto.InternalMessageInfo

// QueryContractsByCodeRequest is the request type for the Query/ContractsByCode
// RPC method
type QueryContractsByCodeRequest struct {
	CodeId uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContractsByCodeRequest) Reset()         { *m = QueryContractsByCodeRequest{} }
func (m *QueryContractsByCodeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractsByCodeRequest) ProtoMessage()    {}
func (*QueryContractsByCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{4}
}

func (m *QueryContractsByCodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractsByCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractsByCodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractsByCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractsByCodeRequest.Merge(m, src)
}

func (m *QueryContractsByCodeRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractsByCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractsByCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractsByCodeRequest proto.InternalMessageInfo

// QueryContractsByCodeResponse is the response type for the
// Query/ContractsByCode RPC method
type QueryContractsByCodeResponse struct {
	// contracts are a set of contract addresses
	Contracts []string `protobuf:"bytes,1,rep,name=contracts,proto3" json:"contracts,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContractsByCodeResponse) Reset()         { *m = QueryContractsByCodeResponse{} }
func (m *QueryContractsByCodeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractsByCodeResponse) ProtoMessage()    {}
func (*QueryContractsByCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{5}
}

func (m *QueryContractsByCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractsByCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractsByCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractsByCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractsByCodeResponse.Merge(m, src)
}

func (m *QueryContractsByCodeResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractsByCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractsByCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractsByCodeResponse proto.InternalMessageInfo

// QueryAllContractStateRequest is the request type for the
// Query/AllContractState RPC method
type QueryAllContractStateRequest struct {
	// address is the address of the contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllContractStateRequest) Reset()         { *m = QueryAllContractStateRequest{} }
func (m *QueryAllContractStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllContractStateRequest) ProtoMessage()    {}
func (*QueryAllContractStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{6}
}

func (m *QueryAllContractStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryAllContractStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllContractStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryAllContractStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllContractStateRequest.Merge(m, src)
}

func (m *QueryAllContractStateRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryAllContractStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllContractStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllContractStateRequest proto.InternalMessageInfo

// QueryAllContractStateResponse is the response type for the
// Query/AllContractState RPC method
type QueryAllContractStateResponse struct {
	Models []Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllContractStateResponse) Reset()         { *m = QueryAllContractStateResponse{} }
func (m *QueryAllContractStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllContractStateResponse) ProtoMessage()    {}
func (*QueryAllContractStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{7}
}

func (m *QueryAllContractStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryAllContractStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllContractStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryAllContractStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllContractStateResponse.Merge(m, src)
}

func (m *QueryAllContractStateResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryAllContractStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllContractStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllContractStateResponse proto.InternalMessageInfo

// QueryRawContractStateRequest is the request type for the
// Query/RawContractState RPC method
type QueryRawContractStateRequest struct {
	// address is the address of the contract
	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	QueryData []byte `protobuf:"bytes,2,opt,name=query_data,json=queryData,proto3" json:"query_data,omitempty"`
}

func (m *QueryRawContractStateRequest) Reset()         { *m = QueryRawContractStateRequest{} }
func (m *QueryRawContractStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRawContractStateRequest) ProtoMessage()    {}
func (*QueryRawContractStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{8}
}

func (m *QueryRawContractStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryRawContractStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRawContractStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryRawContractStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRawContractStateRequest.Merge(m, src)
}

func (m *QueryRawContractStateRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryRawContractStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRawContractStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRawContractStateRequest proto.InternalMessageInfo

// QueryRawContractStateResponse is the response type for the
// Query/RawContractState RPC method
type QueryRawContractStateResponse struct {
	// Data contains the raw store data
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *QueryRawContractStateResponse) Reset()         { *m = QueryRawContractStateResponse{} }
func (m *QueryRawContractStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRawContractStateResponse) ProtoMessage()    {}
func (*QueryRawContractStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{9}
}

func (m *QueryRawContractStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryRawContractStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRawContractStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryRawContractStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRawContractStateResponse.Merge(m, src)
}

func (m *QueryRawContractStateResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryRawContractStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRawContractStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRawContractStateResponse proto.InternalMessageInfo

// QuerySmartContractStateRequest is the request type for the
// Query/SmartContractState RPC method
type QuerySmartContractStateRequest struct {
	// address is the address of the contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// QueryData contains the query data passed to the contract
	QueryData RawContractMessage `protobuf:"bytes,2,opt,name=query_data,json=queryData,proto3,casttype=RawContractMessage" json:"query_data,omitempty"`
}

func (m *QuerySmartContractStateRequest) Reset()         { *m = QuerySmartContractStateRequest{} }
func (m *QuerySmartContractStateRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySmartContractStateRequest) ProtoMessage()    {}
func (*QuerySmartContractStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{10}
}

func (m *QuerySmartContractStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QuerySmartContractStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySmartContractStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QuerySmartContractStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySmartContractStateRequest.Merge(m, src)
}

func (m *QuerySmartContractStateRequest) XXX_Size() int {
	return m.Size()
}

func (m *QuerySmartContractStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySmartContractStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySmartContractStateRequest proto.InternalMessageInfo

// QuerySmartContractStateResponse is the response type for the
// Query/SmartContractState RPC method
type QuerySmartContractStateResponse struct {
	// Data contains the json data returned from the smart contract
	Data RawContractMessage `protobuf:"bytes,1,opt,name=data,proto3,casttype=RawContractMessage" json:"data,omitempty"`
}

func (m *QuerySmartContractStateResponse) Reset()         { *m = QuerySmartContractStateResponse{} }
func (m *QuerySmartContractStateResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySmartContractStateResponse) ProtoMessage()    {}
func (*QuerySmartContractStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{11}
}

func (m *QuerySmartContractStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QuerySmartContractStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySmartContractStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QuerySmartContractStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySmartContractStateResponse.Merge(m, src)
}

func (m *QuerySmartContractStateResponse) XXX_Size() int {
	return m.Size()
}

func (m *QuerySmartContractStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySmartContractStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySmartContractStateResponse proto.InternalMessageInfo

// QueryCodeRequest is the request type for the Query/Code RPC method
type QueryCodeRequest struct {
	CodeId uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
}

func (m *QueryCodeRequest) Reset()         { *m = QueryCodeRequest{} }
func (m *QueryCodeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCodeRequest) ProtoMessage()    {}
func (*QueryCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{12}
}

func (m *QueryCodeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodeRequest.Merge(m, src)
}

func (m *QueryCodeRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodeRequest proto.InternalMessageInfo

// CodeInfoResponse contains code meta data from CodeInfo
type CodeInfoResponse struct {
	CodeID                uint64                                               `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"id"`
	Creator               string                                               `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	DataHash              github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"data_hash,omitempty"`
	InstantiatePermission AccessConfig                                         `protobuf:"bytes,6,opt,name=instantiate_permission,json=instantiatePermission,proto3" json:"instantiate_permission"`
}

func (m *CodeInfoResponse) Reset()         { *m = CodeInfoResponse{} }
func (m *CodeInfoResponse) String() string { return proto.CompactTextString(m) }
func (*CodeInfoResponse) ProtoMessage()    {}
func (*CodeInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{13}
}

func (m *CodeInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *CodeInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *CodeInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeInfoResponse.Merge(m, src)
}

func (m *CodeInfoResponse) XXX_Size() int {
	return m.Size()
}

func (m *CodeInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CodeInfoResponse proto.InternalMessageInfo

// QueryCodeResponse is the response type for the Query/Code RPC method
type QueryCodeResponse struct {
	*CodeInfoResponse `protobuf:"bytes,1,opt,name=code_info,json=codeInfo,proto3,embedded=code_info" json:""`
	Data              []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (m *QueryCodeResponse) Reset()         { *m = QueryCodeResponse{} }
func (m *QueryCodeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCodeResponse) ProtoMessage()    {}
func (*QueryCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{14}
}

func (m *QueryCodeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodeResponse.Merge(m, src)
}

func (m *QueryCodeResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodeResponse proto.InternalMessageInfo

// QueryCodesRequest is the request type for the Query/Codes RPC method
type QueryCodesRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCodesRequest) Reset()         { *m = QueryCodesRequest{} }
func (m *QueryCodesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCodesRequest) ProtoMessage()    {}
func (*QueryCodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{15}
}

func (m *QueryCodesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryCodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryCodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodesRequest.Merge(m, src)
}

func (m *QueryCodesRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryCodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodesRequest proto.InternalMessageInfo

// QueryCodesResponse is the response type for the Query/Codes RPC method
type QueryCodesResponse struct {
	CodeInfos []CodeInfoResponse `protobuf:"bytes,1,rep,name=code_infos,json=codeInfos,proto3" json:"code_infos"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCodesResponse) Reset()         { *m = QueryCodesResponse{} }
func (m *QueryCodesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCodesResponse) ProtoMessage()    {}
func (*QueryCodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{16}
}

func (m *QueryCodesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryCodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCodesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryCodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCodesResponse.Merge(m, src)
}

func (m *QueryCodesResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryCodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCodesResponse proto.InternalMessageInfo

// QueryPinnedCodesRequest is the request type for the Query/PinnedCodes
// RPC method
type QueryPinnedCodesRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryPinnedCodesRequest) Reset()         { *m = QueryPinnedCodesRequest{} }
func (m *QueryPinnedCodesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPinnedCodesRequest) ProtoMessage()    {}
func (*QueryPinnedCodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{17}
}

func (m *QueryPinnedCodesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryPinnedCodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPinnedCodesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryPinnedCodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPinnedCodesRequest.Merge(m, src)
}

func (m *QueryPinnedCodesRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryPinnedCodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPinnedCodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPinnedCodesRequest proto.InternalMessageInfo

// QueryPinnedCodesResponse is the response type for the
// Query/PinnedCodes RPC method
type QueryPinnedCodesResponse struct {
	CodeIDs []uint64 `protobuf:"varint,1,rep,packed,name=code_ids,json=codeIds,proto3" json:"code_ids,omitempty"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryPinnedCodesResponse) Reset()         { *m = QueryPinnedCodesResponse{} }
func (m *QueryPinnedCodesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPinnedCodesResponse) ProtoMessage()    {}
func (*QueryPinnedCodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{18}
}

func (m *QueryPinnedCodesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryPinnedCodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPinnedCodesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryPinnedCodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPinnedCodesResponse.Merge(m, src)
}

func (m *QueryPinnedCodesResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryPinnedCodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPinnedCodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPinnedCodesResponse proto.InternalMessageInfo

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct{}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{19}
}

func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}

func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{20}
}

func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}

func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

// QueryContractsByCreatorRequest is the request type for the
// Query/ContractsByCreator RPC method.
type QueryContractsByCreatorRequest struct {
	// CreatorAddress is the address of contract creator
	CreatorAddress string `protobuf:"bytes,1,opt,name=creator_address,json=creatorAddress,proto3" json:"creator_address,omitempty"`
	// Pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContractsByCreatorRequest) Reset()         { *m = QueryContractsByCreatorRequest{} }
func (m *QueryContractsByCreatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContractsByCreatorRequest) ProtoMessage()    {}
func (*QueryContractsByCreatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{21}
}

func (m *QueryContractsByCreatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractsByCreatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractsByCreatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractsByCreatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractsByCreatorRequest.Merge(m, src)
}

func (m *QueryContractsByCreatorRequest) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractsByCreatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractsByCreatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractsByCreatorRequest proto.InternalMessageInfo

// QueryContractsByCreatorResponse is the response type for the
// Query/ContractsByCreator RPC method.
type QueryContractsByCreatorResponse struct {
	// ContractAddresses result set
	ContractAddresses []string `protobuf:"bytes,1,rep,name=contract_addresses,json=contractAddresses,proto3" json:"contract_addresses,omitempty"`
	// Pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryContractsByCreatorResponse) Reset()         { *m = QueryContractsByCreatorResponse{} }
func (m *QueryContractsByCreatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContractsByCreatorResponse) ProtoMessage()    {}
func (*QueryContractsByCreatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9677c207036b9f2b, []int{22}
}

func (m *QueryContractsByCreatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *QueryContractsByCreatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContractsByCreatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *QueryContractsByCreatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContractsByCreatorResponse.Merge(m, src)
}

func (m *QueryContractsByCreatorResponse) XXX_Size() int {
	return m.Size()
}

func (m *QueryContractsByCreatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContractsByCreatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContractsByCreatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryContractInfoRequest)(nil), "cosmwasm.wasm.v1.QueryContractInfoRequest")
	proto.RegisterType((*QueryContractInfoResponse)(nil), "cosmwasm.wasm.v1.QueryContractInfoResponse")
	proto.RegisterType((*QueryContractHistoryRequest)(nil), "cosmwasm.wasm.v1.QueryContractHistoryRequest")
	proto.RegisterType((*QueryContractHistoryResponse)(nil), "cosmwasm.wasm.v1.QueryContractHistoryResponse")
	proto.RegisterType((*QueryContractsByCodeRequest)(nil), "cosmwasm.wasm.v1.QueryContractsByCodeRequest")
	proto.RegisterType((*QueryContractsByCodeResponse)(nil), "cosmwasm.wasm.v1.QueryContractsByCodeResponse")
	proto.RegisterType((*QueryAllContractStateRequest)(nil), "cosmwasm.wasm.v1.QueryAllContractStateRequest")
	proto.RegisterType((*QueryAllContractStateResponse)(nil), "cosmwasm.wasm.v1.QueryAllContractStateResponse")
	proto.RegisterType((*QueryRawContractStateRequest)(nil), "cosmwasm.wasm.v1.QueryRawContractStateRequest")
	proto.RegisterType((*QueryRawContractStateResponse)(nil), "cosmwasm.wasm.v1.QueryRawContractStateResponse")
	proto.RegisterType((*QuerySmartContractStateRequest)(nil), "cosmwasm.wasm.v1.QuerySmartContractStateRequest")
	proto.RegisterType((*QuerySmartContractStateResponse)(nil), "cosmwasm.wasm.v1.QuerySmartContractStateResponse")
	proto.RegisterType((*QueryCodeRequest)(nil), "cosmwasm.wasm.v1.QueryCodeRequest")
	proto.RegisterType((*CodeInfoResponse)(nil), "cosmwasm.wasm.v1.CodeInfoResponse")
	proto.RegisterType((*QueryCodeResponse)(nil), "cosmwasm.wasm.v1.QueryCodeResponse")
	proto.RegisterType((*QueryCodesRequest)(nil), "cosmwasm.wasm.v1.QueryCodesRequest")
	proto.RegisterType((*QueryCodesResponse)(nil), "cosmwasm.wasm.v1.QueryCodesResponse")
	proto.RegisterType((*QueryPinnedCodesRequest)(nil), "cosmwasm.wasm.v1.QueryPinnedCodesRequest")
	proto.RegisterType((*QueryPinnedCodesResponse)(nil), "cosmwasm.wasm.v1.QueryPinnedCodesResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmwasm.wasm.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmwasm.wasm.v1.QueryParamsResponse")
	proto.RegisterType((*QueryContractsByCreatorRequest)(nil), "cosmwasm.wasm.v1.QueryContractsByCreatorRequest")
	proto.RegisterType((*QueryContractsByCreatorResponse)(nil), "cosmwasm.wasm.v1.QueryContractsByCreatorResponse")
}

func init() { proto.RegisterFile("cosmwasm/wasm/v1/query.proto", fileDescriptor_9677c207036b9f2b) }

var fileDescriptor_9677c207036b9f2b = []byte{
	// 1331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x98, 0xcd, 0x6f, 0x1b, 0xc5,
	0x1b, 0xc7, 0x3d, 0xa9, 0xe3, 0x97, 0x69, 0xfa, 0xab, 0x3b, 0xbf, 0xd2, 0x18, 0x93, 0xae, 0xa3,
	0xa5, 0xa4, 0x69, 0x9a, 0xee, 0x92, 0x34, 0xa1, 0x80, 0x84, 0x50, 0x9c, 0x42, 0x93, 0x48, 0x91,
	0xd2, 0xad, 0x50, 0x25, 0x7a, 0xb0, 0xc6, 0xde, 0x89, 0xb3, 0x52, 0xbc, 0xeb, 0xec, 0x4c, 0x92,
	0x5a, 0x51, 0x00, 0x55, 0xe2, 0x86, 0x78, 0x11, 0xe2, 0xc0, 0x01, 0xc1, 0x01, 0x15, 0xce, 0x70,
	0x41, 0x5c, 0xb9, 0xe4, 0x18, 0x89, 0x0b, 0x27, 0x0b, 0x1c, 0x0e, 0x28, 0x7f, 0x42, 0x4f, 0x68,
	0x67, 0x67, 0x9c, 0x5d, 0xdb, 0x1b, 0x3b, 0x95, 0xc5, 0xc5, 0xda, 0xdd, 0x79, 0x66, 0x9e, 0xcf,
	0xf3, 0x9d, 0x67, 0xe6, 0x79, 0x64, 0x38, 0x56, 0x76, 0x68, 0x75, 0x17, 0xd3, 0xaa, 0xce, 0x7f,
	0x76, 0x66, 0xf4, 0xad, 0x6d, 0xe2, 0xd6, 0xb5, 0x9a, 0xeb, 0x30, 0x07, 0x65, 0xe4, 0xa8, 0xc6,
	0x7f, 0x76, 0x66, 0x72, 0x97, 0x2b, 0x4e, 0xc5, 0xe1, 0x83, 0xba, 0xf7, 0xe4, 0xdb, 0xe5, 0x3a,
	0x57, 0x61, 0xf5, 0x1a, 0xa1, 0x72, 0xb4, 0xe2, 0x38, 0x95, 0x4d, 0xa2, 0xe3, 0x9a, 0xa5, 0x63,
	0xdb, 0x76, 0x18, 0x66, 0x96, 0x63, 0xcb, 0xd1, 0x29, 0x6f, 0xae, 0x43, 0xf5, 0x12, 0xa6, 0xc4,
	0x77, 0xae, 0xef, 0xcc, 0x94, 0x08, 0xc3, 0x33, 0x7a, 0x0d, 0x57, 0x2c, 0x9b, 0x1b, 0xfb, 0xb6,
	0xea, 0x1c, 0xcc, 0xde, 0xf7, 0x2c, 0x16, 0x1d, 0x9b, 0xb9, 0xb8, 0xcc, 0x96, 0xed, 0x75, 0xc7,
	0x20, 0x5b, 0xdb, 0x84, 0x32, 0x94, 0x85, 0x49, 0x6c, 0x9a, 0x2e, 0xa1, 0x34, 0x0b, 0xc6, 0xc1,
	0x64, 0xda, 0x90, 0xaf, 0xea, 0xa7, 0x00, 0xbe, 0xd8, 0x65, 0x1a, 0xad, 0x39, 0x36, 0x25, 0xd1,
	0xf3, 0xd0, 0x7d, 0x78, 0xa1, 0x2c, 0x66, 0x14, 0x2d, 0x7b, 0xdd, 0xc9, 0x0e, 0x8d, 0x83, 0xc9,
	0xf3, 0xb3, 0x8a, 0xd6, 0xae, 0x8a, 0x16, 0x5c, 0xb8, 0x30, 0x72, 0xd0, 0xc8, 0xc7, 0x0e, 0x1b,
	0x79, 0x70, 0xdc, 0xc8, 0xc7, 0x8c, 0x91, 0x72, 0x60, 0xec, 0xcd, 0xf8, 0x3f, 0xdf, 0xe5, 0x81,
	0xfa, 0x21, 0x7c, 0x29, 0xc4, 0xb3, 0x64, 0x51, 0xe6, 0xb8, 0xf5, 0x9e, 0x91, 0xa0, 0x77, 0x21,
	0x3c, 0xd1, 0x44, 0xe0, 0x4c, 0x68, 0xbe, 0x80, 0x9a, 0x27, 0xa0, 0xe6, 0xef, 0x9e, 0x10, 0x50,
	0x5b, 0xc3, 0x15, 0x22, 0x56, 0x35, 0x02, 0x33, 0xd5, 0x9f, 0x01, 0x1c, 0xeb, 0x4e, 0x20, 0x44,
	0x59, 0x81, 0x49, 0x62, 0x33, 0xd7, 0x22, 0x1e, 0xc2, 0xb9, 0xc9, 0xf3, 0xb3, 0x53, 0xd1, 0x41,
	0x2f, 0x3a, 0x26, 0x11, 0xf3, 0xdf, 0xb1, 0x99, 0x5b, 0x2f, 0xc4, 0x3d, 0x01, 0x0c, 0xb9, 0x00,
	0xba, 0xd7, 0x05, 0xfa, 0x7a, 0x4f, 0x68, 0x1f, 0x24, 0x44, 0xfd, 0x41, 0x9b, 0x6c, 0xb4, 0x50,
	0xf7, 0x7c, 0x4b, 0xd9, 0x46, 0x61, 0xb2, 0xec, 0x98, 0xa4, 0x68, 0x99, 0x5c, 0xb6, 0xb8, 0x91,
	0xf0, 0x5e, 0x97, 0xcd, 0x81, 0xa9, 0xf6, 0x71, 0xbb, 0x6a, 0x2d, 0x00, 0xa1, 0xda, 0x18, 0x4c,
	0xcb, 0xdd, 0xf6, 0x75, 0x4b, 0x1b, 0x27, 0x1f, 0x06, 0xa7, 0xc3, 0x47, 0x92, 0x63, 0x61, 0x73,
	0x53, 0xa2, 0x3c, 0x60, 0x98, 0x91, 0xff, 0x2e, 0x81, 0xbe, 0x05, 0xf0, 0x6a, 0x04, 0x82, 0xd0,
	0x62, 0x1e, 0x26, 0xaa, 0x8e, 0x49, 0x36, 0x65, 0x02, 0x8d, 0x76, 0x26, 0xd0, 0xaa, 0x37, 0x2e,
	0xb2, 0x45, 0x18, 0x0f, 0x4e, 0xa4, 0x87, 0x42, 0x23, 0x03, 0xef, 0x9e, 0x51, 0xa3, 0xab, 0x10,
	0x72, 0x1f, 0x45, 0x13, 0x33, 0xcc, 0x11, 0x46, 0x8c, 0x34, 0xff, 0x72, 0x17, 0x33, 0xac, 0xde,
	0x16, 0x91, 0x77, 0x2e, 0x2c, 0x22, 0x47, 0x30, 0xce, 0x67, 0x02, 0x3e, 0x93, 0x3f, 0xab, 0x5b,
	0x50, 0xe1, 0x93, 0x1e, 0x54, 0xb1, 0xcb, 0xce, 0xc8, 0x33, 0xdf, 0xc9, 0x53, 0xb8, 0xf2, 0xac,
	0x91, 0x47, 0x01, 0x82, 0x55, 0x42, 0xa9, 0xa7, 0x44, 0x80, 0x73, 0x15, 0xe6, 0x23, 0x5d, 0x0a,
	0xd2, 0xa9, 0x20, 0x69, 0xe4, 0x9a, 0x7e, 0x04, 0x37, 0x61, 0x46, 0xe4, 0x7e, 0xef, 0x13, 0xa7,
	0x7e, 0x33, 0x04, 0x33, 0x9e, 0x61, 0xe8, 0xa2, 0xbd, 0xd1, 0x66, 0x5d, 0xc8, 0x34, 0x1b, 0xf9,
	0x04, 0x37, 0xbb, 0x7b, 0xdc, 0xc8, 0x0f, 0x59, 0x66, 0xeb, 0xc4, 0x66, 0x61, 0xb2, 0xec, 0x12,
	0xcc, 0x1c, 0x97, 0xc7, 0x9b, 0x36, 0xe4, 0x2b, 0x7a, 0x0f, 0xa6, 0x3d, 0x9c, 0xe2, 0x06, 0xa6,
	0x1b, 0xd9, 0x73, 0x9c, 0xfb, 0xf5, 0x67, 0x8d, 0xfc, 0x5c, 0xc5, 0x62, 0x1b, 0xdb, 0x25, 0xad,
	0xec, 0x54, 0x75, 0x46, 0x6c, 0x93, 0xb8, 0x55, 0xcb, 0x66, 0xc1, 0xc7, 0x4d, 0xab, 0x44, 0xf5,
	0x52, 0x9d, 0x11, 0xaa, 0x2d, 0x91, 0xc7, 0x05, 0xef, 0xc1, 0x48, 0x79, 0x4b, 0x2d, 0x61, 0xba,
	0x81, 0x1e, 0xc1, 0x2b, 0x96, 0x4d, 0x19, 0xb6, 0x99, 0x85, 0x19, 0x29, 0xd6, 0xbc, 0x49, 0x94,
	0x7a, 0x29, 0x98, 0x88, 0xba, 0xf3, 0x17, 0xca, 0x65, 0x42, 0xe9, 0xa2, 0x63, 0xaf, 0x5b, 0x15,
	0x91, 0xc4, 0x2f, 0x04, 0xd6, 0x58, 0x6b, 0x2d, 0xe1, 0x5f, 0xfa, 0x2b, 0xf1, 0x54, 0x3c, 0x33,
	0xbc, 0x12, 0x4f, 0x0d, 0x67, 0x12, 0xea, 0x13, 0x00, 0x2f, 0x05, 0xd4, 0x14, 0x02, 0x2d, 0x7b,
	0xd7, 0x87, 0x27, 0x90, 0x57, 0x6b, 0x00, 0xf7, 0xab, 0x76, 0xbb, 0x76, 0xc3, 0xba, 0x16, 0x52,
	0xad, 0x5a, 0x93, 0x2a, 0x8b, 0x31, 0x34, 0x26, 0x76, 0xd6, 0xcf, 0x96, 0xd4, 0x71, 0x23, 0xcf,
	0xdf, 0xfd, 0xbd, 0x14, 0x55, 0xe8, 0x51, 0x80, 0x81, 0xca, 0x2d, 0x0d, 0x5f, 0x10, 0xe0, 0xb9,
	0x2f, 0x88, 0xa7, 0x00, 0xa2, 0xe0, 0xea, 0x22, 0xc4, 0x7b, 0x10, 0xb6, 0x42, 0x94, 0x37, 0x43,
	0x3f, 0x31, 0xfa, 0xfa, 0xa6, 0x65, 0x7c, 0x03, 0xbc, 0x27, 0x30, 0x1c, 0xe5, 0x9c, 0x6b, 0x96,
	0x6d, 0x13, 0xf3, 0x14, 0x2d, 0x9e, 0xff, 0xb2, 0xfc, 0x0c, 0x88, 0xb6, 0x25, 0xe4, 0xa3, 0x75,
	0x06, 0x53, 0xe2, 0x54, 0xf8, 0x7a, 0xc4, 0x0b, 0x17, 0xbd, 0x58, 0x9b, 0x8d, 0x7c, 0xd2, 0x3f,
	0x1a, 0xd4, 0x48, 0xfa, 0xa7, 0x62, 0x80, 0x41, 0x5f, 0x16, 0x9b, 0xb3, 0x86, 0x5d, 0x5c, 0x95,
	0xf1, 0xaa, 0xab, 0xf0, 0xff, 0xa1, 0xaf, 0x82, 0xf0, 0x35, 0x98, 0xa8, 0xf1, 0x2f, 0x22, 0x1d,
	0xb2, 0x9d, 0xfb, 0xe5, 0xcf, 0x90, 0x57, 0xb9, 0x6f, 0xad, 0x7e, 0x01, 0xc4, 0xa5, 0x17, 0x2c,
	0x97, 0xfe, 0x31, 0x96, 0x0a, 0x5f, 0x87, 0x17, 0xc5, 0xc1, 0x2e, 0x86, 0x2f, 0xbf, 0xff, 0x89,
	0xcf, 0x0b, 0x03, 0xae, 0x5b, 0x5f, 0x03, 0x71, 0x2b, 0x76, 0x63, 0x12, 0xf1, 0xde, 0x82, 0xa8,
	0xd5, 0xf6, 0x09, 0x2a, 0x22, 0xcb, 0xf9, 0x25, 0x39, 0xb2, 0x20, 0x07, 0x06, 0xb6, 0x29, 0xb3,
	0xbf, 0x5d, 0x80, 0xc3, 0x9c, 0x0d, 0x7d, 0x05, 0xe0, 0x48, 0xb0, 0xa5, 0x44, 0x5d, 0xba, 0xaf,
	0xa8, 0x3e, 0x38, 0x77, 0xb3, 0x2f, 0x5b, 0xdf, 0xbf, 0x3a, 0xfd, 0xe4, 0xf7, 0xbf, 0xbf, 0x1c,
	0x9a, 0x40, 0xd7, 0xf4, 0x8e, 0x0e, 0x5e, 0x46, 0xaa, 0xef, 0x09, 0x11, 0xf6, 0xd1, 0x53, 0x00,
	0x2f, 0xb6, 0x75, 0x8c, 0xe8, 0x56, 0x0f, 0x77, 0xe1, 0xde, 0x36, 0xa7, 0xf5, 0x6b, 0x2e, 0x00,
	0xe7, 0x38, 0xa0, 0x86, 0xa6, 0xfb, 0x01, 0xd4, 0x37, 0x04, 0xd4, 0xf7, 0x01, 0x50, 0xd1, 0xa4,
	0xf5, 0x04, 0x0d, 0x77, 0x93, 0x3d, 0x41, 0xdb, 0x7a, 0x3f, 0x75, 0x96, 0x83, 0x4e, 0xa3, 0xa9,
	0x6e, 0xa0, 0x26, 0xd1, 0xf7, 0xc4, 0x29, 0xdf, 0xd7, 0x4f, 0x3a, 0xc2, 0x1f, 0x00, 0xcc, 0xb4,
	0x37, 0x50, 0x28, 0xca, 0x71, 0x44, 0xb3, 0x97, 0xd3, 0xfb, 0xb6, 0xef, 0x87, 0xb4, 0x43, 0x52,
	0xca, 0xa1, 0x7e, 0x02, 0x30, 0xd3, 0xde, 0xf0, 0x44, 0x92, 0x46, 0xb4, 0x5c, 0x91, 0xa4, 0x51,
	0x9d, 0x94, 0xfa, 0x16, 0x27, 0xbd, 0x83, 0xe6, 0xfb, 0x22, 0x75, 0xf1, 0xae, 0xbe, 0x77, 0xd2,
	0x29, 0xed, 0xa3, 0x5f, 0x01, 0x44, 0x9d, 0xdd, 0x0f, 0x7a, 0x35, 0x02, 0x23, 0xb2, 0x37, 0xcb,
	0xcd, 0x9c, 0x61, 0x86, 0x40, 0x7f, 0x9b, 0xa3, 0xbf, 0x81, 0xee, 0xf4, 0x27, 0xb2, 0xb7, 0x50,
	0x18, 0xbe, 0x0e, 0xe3, 0x3c, 0x6d, 0xd5, 0xc8, 0x3c, 0x3c, 0xc9, 0xd5, 0x97, 0x4f, 0xb5, 0x11,
	0x44, 0x93, 0x9c, 0x48, 0x45, 0xe3, 0xbd, 0x12, 0x14, 0xb9, 0x70, 0x98, 0xd7, 0x28, 0x74, 0xda,
	0xba, 0xb2, 0x6a, 0xe4, 0xae, 0x9d, 0x6e, 0x24, 0xbc, 0x2b, 0xdc, 0x7b, 0x16, 0x5d, 0xe9, 0xee,
	0x1d, 0x7d, 0x02, 0xe0, 0xf9, 0x40, 0x79, 0x44, 0x37, 0x22, 0x56, 0xed, 0x2c, 0xd3, 0xb9, 0xa9,
	0x7e, 0x4c, 0x05, 0xc6, 0x04, 0xc7, 0x18, 0x47, 0x4a, 0x77, 0x0c, 0xaa, 0xd7, 0xf8, 0x24, 0xb4,
	0x0f, 0x13, 0x7e, 0x4d, 0x43, 0x51, 0xe1, 0x85, 0x4a, 0x67, 0xee, 0x95, 0x1e, 0x56, 0x7d, 0xbb,
	0xf7, 0x9d, 0xfe, 0x02, 0x20, 0xea, 0xac, 0x50, 0x91, 0x99, 0x1b, 0x59, 0x60, 0x23, 0x33, 0x37,
	0xba, 0xfc, 0xf5, 0x73, 0xe8, 0xa8, 0x2e, 0xca, 0xb3, 0xbe, 0xd7, 0x56, 0xbe, 0xf7, 0x0b, 0x4b,
	0x07, 0x7f, 0x29, 0xb1, 0x1f, 0x9b, 0x4a, 0xec, 0xa0, 0xa9, 0x80, 0xc3, 0xa6, 0x02, 0xfe, 0x6c,
	0x2a, 0xe0, 0xf3, 0x23, 0x25, 0x76, 0x78, 0xa4, 0xc4, 0xfe, 0x38, 0x52, 0x62, 0xef, 0x4f, 0x04,
	0x7a, 0xf5, 0x45, 0x87, 0x56, 0x1f, 0x4a, 0x17, 0xa6, 0xfe, 0xd8, 0x77, 0xc5, 0xff, 0x3c, 0x2a,
	0x25, 0xf8, 0x7f, 0x3e, 0xb7, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x78, 0x46, 0xe1, 0xa8, 0xa3,
	0x12, 0x00, 0x00,
}

func (this *QueryContractInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryContractInfoResponse)
	if !ok {
		that2, ok := that.(QueryContractInfoResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if !this.ContractInfo.Equal(&that1.ContractInfo) {
		return false
	}
	return true
}

func (this *CodeInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CodeInfoResponse)
	if !ok {
		that2, ok := that.(CodeInfoResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CodeID != that1.CodeID {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if !bytes.Equal(this.DataHash, that1.DataHash) {
		return false
	}
	if !this.InstantiatePermission.Equal(&that1.InstantiatePermission) {
		return false
	}
	return true
}

func (this *QueryCodeResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryCodeResponse)
	if !ok {
		that2, ok := that.(QueryCodeResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CodeInfoResponse.Equal(that1.CodeInfoResponse) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ContractInfo gets the contract meta data
	ContractInfo(ctx context.Context, in *QueryContractInfoRequest, opts ...grpc.CallOption) (*QueryContractInfoResponse, error)
	// ContractHistory gets the contract code history
	ContractHistory(ctx context.Context, in *QueryContractHistoryRequest, opts ...grpc.CallOption) (*QueryContractHistoryResponse, error)
	// ContractsByCode lists all smart contracts for a code id
	ContractsByCode(ctx context.Context, in *QueryContractsByCodeRequest, opts ...grpc.CallOption) (*QueryContractsByCodeResponse, error)
	// AllContractState gets all raw store data for a single contract
	AllContractState(ctx context.Context, in *QueryAllContractStateRequest, opts ...grpc.CallOption) (*QueryAllContractStateResponse, error)
	// RawContractState gets single key from the raw store data of a contract
	RawContractState(ctx context.Context, in *QueryRawContractStateRequest, opts ...grpc.CallOption) (*QueryRawContractStateResponse, error)
	// SmartContractState get smart query result from the contract
	SmartContractState(ctx context.Context, in *QuerySmartContractStateRequest, opts ...grpc.CallOption) (*QuerySmartContractStateResponse, error)
	// Code gets the binary code and metadata for a singe wasm code
	Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error)
	// Codes gets the metadata for all stored wasm codes
	Codes(ctx context.Context, in *QueryCodesRequest, opts ...grpc.CallOption) (*QueryCodesResponse, error)
	// PinnedCodes gets the pinned code ids
	PinnedCodes(ctx context.Context, in *QueryPinnedCodesRequest, opts ...grpc.CallOption) (*QueryPinnedCodesResponse, error)
	// Params gets the module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ContractsByCreator gets the contracts by creator
	ContractsByCreator(ctx context.Context, in *QueryContractsByCreatorRequest, opts ...grpc.CallOption) (*QueryContractsByCreatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ContractInfo(ctx context.Context, in *QueryContractInfoRequest, opts ...grpc.CallOption) (*QueryContractInfoResponse, error) {
	out := new(QueryContractInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/ContractInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractHistory(ctx context.Context, in *QueryContractHistoryRequest, opts ...grpc.CallOption) (*QueryContractHistoryResponse, error) {
	out := new(QueryContractHistoryResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/ContractHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractsByCode(ctx context.Context, in *QueryContractsByCodeRequest, opts ...grpc.CallOption) (*QueryContractsByCodeResponse, error) {
	out := new(QueryContractsByCodeResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/ContractsByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllContractState(ctx context.Context, in *QueryAllContractStateRequest, opts ...grpc.CallOption) (*QueryAllContractStateResponse, error) {
	out := new(QueryAllContractStateResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/AllContractState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RawContractState(ctx context.Context, in *QueryRawContractStateRequest, opts ...grpc.CallOption) (*QueryRawContractStateResponse, error) {
	out := new(QueryRawContractStateResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/RawContractState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SmartContractState(ctx context.Context, in *QuerySmartContractStateRequest, opts ...grpc.CallOption) (*QuerySmartContractStateResponse, error) {
	out := new(QuerySmartContractStateResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/SmartContractState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error) {
	out := new(QueryCodeResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/Code", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Codes(ctx context.Context, in *QueryCodesRequest, opts ...grpc.CallOption) (*QueryCodesResponse, error) {
	out := new(QueryCodesResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/Codes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PinnedCodes(ctx context.Context, in *QueryPinnedCodesRequest, opts ...grpc.CallOption) (*QueryPinnedCodesResponse, error) {
	out := new(QueryPinnedCodesResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/PinnedCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractsByCreator(ctx context.Context, in *QueryContractsByCreatorRequest, opts ...grpc.CallOption) (*QueryContractsByCreatorResponse, error) {
	out := new(QueryContractsByCreatorResponse)
	err := c.cc.Invoke(ctx, "/cosmwasm.wasm.v1.Query/ContractsByCreator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ContractInfo gets the contract meta data
	ContractInfo(context.Context, *QueryContractInfoRequest) (*QueryContractInfoResponse, error)
	// ContractHistory gets the contract code history
	ContractHistory(context.Context, *QueryContractHistoryRequest) (*QueryContractHistoryResponse, error)
	// ContractsByCode lists all smart contracts for a code id
	ContractsByCode(context.Context, *QueryContractsByCodeRequest) (*QueryContractsByCodeResponse, error)
	// AllContractState gets all raw store data for a single contract
	AllContractState(context.Context, *QueryAllContractStateRequest) (*QueryAllContractStateResponse, error)
	// RawContractState gets single key from the raw store data of a contract
	RawContractState(context.Context, *QueryRawContractStateRequest) (*QueryRawContractStateResponse, error)
	// SmartContractState get smart query result from the contract
	SmartContractState(context.Context, *QuerySmartContractStateRequest) (*QuerySmartContractStateResponse, error)
	// Code gets the binary code and metadata for a singe wasm code
	Code(context.Context, *QueryCodeRequest) (*QueryCodeResponse, error)
	// Codes gets the metadata for all stored wasm codes
	Codes(context.Context, *QueryCodesRequest) (*QueryCodesResponse, error)
	// PinnedCodes gets the pinned code ids
	PinnedCodes(context.Context, *QueryPinnedCodesRequest) (*QueryPinnedCodesResponse, error)
	// Params gets the module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// ContractsByCreator gets the contracts by creator
	ContractsByCreator(context.Context, *QueryContractsByCreatorRequest) (*QueryContractsByCreatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct{}

func (*UnimplementedQueryServer) ContractInfo(ctx context.Context, req *QueryContractInfoRequest) (*QueryContractInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractInfo not implemented")
}

func (*UnimplementedQueryServer) ContractHistory(ctx context.Context, req *QueryContractHistoryRequest) (*QueryContractHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractHistory not implemented")
}

func (*UnimplementedQueryServer) ContractsByCode(ctx context.Context, req *QueryContractsByCodeRequest) (*QueryContractsByCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractsByCode not implemented")
}

func (*UnimplementedQueryServer) AllContractState(ctx context.Context, req *QueryAllContractStateRequest) (*QueryAllContractStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllContractState not implemented")
}

func (*UnimplementedQueryServer) RawContractState(ctx context.Context, req *QueryRawContractStateRequest) (*QueryRawContractStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawContractState not implemented")
}

func (*UnimplementedQueryServer) SmartContractState(ctx context.Context, req *QuerySmartContractStateRequest) (*QuerySmartContractStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SmartContractState not implemented")
}

func (*UnimplementedQueryServer) Code(ctx context.Context, req *QueryCodeRequest) (*QueryCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code not implemented")
}

func (*UnimplementedQueryServer) Codes(ctx context.Context, req *QueryCodesRequest) (*QueryCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Codes not implemented")
}

func (*UnimplementedQueryServer) PinnedCodes(ctx context.Context, req *QueryPinnedCodesRequest) (*QueryPinnedCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PinnedCodes not implemented")
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func (*UnimplementedQueryServer) ContractsByCreator(ctx context.Context, req *QueryContractsByCreatorRequest) (*QueryContractsByCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractsByCreator not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ContractInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/ContractInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractInfo(ctx, req.(*QueryContractInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/ContractHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractHistory(ctx, req.(*QueryContractHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractsByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractsByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractsByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/ContractsByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractsByCode(ctx, req.(*QueryContractsByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllContractState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllContractStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllContractState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/AllContractState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllContractState(ctx, req.(*QueryAllContractStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RawContractState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRawContractStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RawContractState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/RawContractState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RawContractState(ctx, req.(*QueryRawContractStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SmartContractState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySmartContractStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SmartContractState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/SmartContractState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SmartContractState(ctx, req.(*QuerySmartContractStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Code_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Code(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/Code",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Code(ctx, req.(*QueryCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Codes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Codes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/Codes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Codes(ctx, req.(*QueryCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PinnedCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPinnedCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PinnedCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/PinnedCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PinnedCodes(ctx, req.(*QueryPinnedCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractsByCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractsByCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractsByCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmwasm.wasm.v1.Query/ContractsByCreator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractsByCreator(ctx, req.(*QueryContractsByCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmwasm.wasm.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContractInfo",
			Handler:    _Query_ContractInfo_Handler,
		},
		{
			MethodName: "ContractHistory",
			Handler:    _Query_ContractHistory_Handler,
		},
		{
			MethodName: "ContractsByCode",
			Handler:    _Query_ContractsByCode_Handler,
		},
		{
			MethodName: "AllContractState",
			Handler:    _Query_AllContractState_Handler,
		},
		{
			MethodName: "RawContractState",
			Handler:    _Query_RawContractState_Handler,
		},
		{
			MethodName: "SmartContractState",
			Handler:    _Query_SmartContractState_Handler,
		},
		{
			MethodName: "Code",
			Handler:    _Query_Code_Handler,
		},
		{
			MethodName: "Codes",
			Handler:    _Query_Codes_Handler,
		},
		{
			MethodName: "PinnedCodes",
			Handler:    _Query_PinnedCodes_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ContractsByCreator",
			Handler:    _Query_ContractsByCreator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmwasm/wasm/v1/query.proto",
}

func (m *QueryContractInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ContractInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractHistoryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractHistoryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractHistoryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractHistoryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractHistoryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractHistoryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractsByCodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractsByCodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractsByCodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CodeId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractsByCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractsByCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractsByCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contracts[iNdEx])
			copy(dAtA[i:], m.Contracts[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Contracts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllContractStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllContractStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllContractStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllContractStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllContractStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllContractStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Models) > 0 {
		for iNdEx := len(m.Models) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Models[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryRawContractStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRawContractStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRawContractStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryData) > 0 {
		i -= len(m.QueryData)
		copy(dAtA[i:], m.QueryData)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.QueryData)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRawContractStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRawContractStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRawContractStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySmartContractStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySmartContractStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySmartContractStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueryData) > 0 {
		i -= len(m.QueryData)
		copy(dAtA[i:], m.QueryData)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.QueryData)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySmartContractStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySmartContractStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySmartContractStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCodeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CodeId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CodeInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.InstantiatePermission.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.DataHash) > 0 {
		i -= len(m.DataHash)
		copy(dAtA[i:], m.DataHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.DataHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCodeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeInfoResponse != nil {
		{
			size, err := m.CodeInfoResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCodesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCodesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCodesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCodesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CodeInfos) > 0 {
		for iNdEx := len(m.CodeInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CodeInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryPinnedCodesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPinnedCodesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPinnedCodesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *QueryPinnedCodesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPinnedCodesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPinnedCodesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CodeIDs) > 0 {
		dAtA15 := make([]byte, len(m.CodeIDs)*10)
		var j14 int
		for _, num := range m.CodeIDs {
			for num >= 1<<7 {
				dAtA15[j14] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j14++
			}
			dAtA15[j14] = uint8(num)
			j14++
		}
		i -= j14
		copy(dAtA[i:], dAtA15[:j14])
		i = encodeVarintQuery(dAtA, i, uint64(j14))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryContractsByCreatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractsByCreatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractsByCreatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CreatorAddress) > 0 {
		i -= len(m.CreatorAddress)
		copy(dAtA[i:], m.CreatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CreatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryContractsByCreatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContractsByCreatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContractsByCreatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddresses) > 0 {
		for iNdEx := len(m.ContractAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ContractAddresses[iNdEx])
			copy(dAtA[i:], m.ContractAddresses[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *QueryContractInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.ContractInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryContractHistoryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractHistoryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractsByCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovQuery(uint64(m.CodeId))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractsByCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Contracts) > 0 {
		for _, s := range m.Contracts {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllContractStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllContractStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Models) > 0 {
		for _, e := range m.Models {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRawContractStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.QueryData)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRawContractStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySmartContractStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.QueryData)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySmartContractStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCodeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovQuery(uint64(m.CodeId))
	}
	return n
}

func (m *CodeInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovQuery(uint64(m.CodeID))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.DataHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = m.InstantiatePermission.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryCodeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeInfoResponse != nil {
		l = m.CodeInfoResponse.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCodesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCodesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CodeInfos) > 0 {
		for _, e := range m.CodeInfos {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPinnedCodesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryPinnedCodesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CodeIDs) > 0 {
		l = 0
		for _, e := range m.CodeIDs {
			l += sovQuery(uint64(e))
		}
		n += 1 + sovQuery(uint64(l)) + l
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryContractsByCreatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CreatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryContractsByCreatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ContractAddresses) > 0 {
		for _, s := range m.ContractAddresses {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *QueryContractInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryContractInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryContractHistoryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractHistoryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractHistoryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryContractHistoryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractHistoryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractHistoryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, ContractCodeHistoryEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryContractsByCodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractsByCodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractsByCodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryContractsByCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractsByCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractsByCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryAllContractStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllContractStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllContractStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryAllContractStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllContractStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllContractStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Models", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Models = append(m.Models, Model{})
			if err := m.Models[len(m.Models)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryRawContractStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRawContractStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRawContractStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryData = append(m.QueryData[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryData == nil {
				m.QueryData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryRawContractStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRawContractStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRawContractStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QuerySmartContractStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySmartContractStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySmartContractStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryData = append(m.QueryData[:0], dAtA[iNdEx:postIndex]...)
			if m.QueryData == nil {
				m.QueryData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QuerySmartContractStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySmartContractStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySmartContractStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryCodeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCodeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *CodeInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CodeInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHash = append(m.DataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.DataHash == nil {
				m.DataHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstantiatePermission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InstantiatePermission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryCodeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCodeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeInfoResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CodeInfoResponse == nil {
				m.CodeInfoResponse = &CodeInfoResponse{}
			}
			if err := m.CodeInfoResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryCodesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCodesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryCodesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCodesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCodesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeInfos = append(m.CodeInfos, CodeInfoResponse{})
			if err := m.CodeInfos[len(m.CodeInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryPinnedCodesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPinnedCodesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPinnedCodesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryPinnedCodesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPinnedCodesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPinnedCodesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CodeIDs = append(m.CodeIDs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthQuery
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthQuery
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CodeIDs) == 0 {
					m.CodeIDs = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CodeIDs = append(m.CodeIDs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeIDs", wireType)
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryContractsByCreatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractsByCreatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractsByCreatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func (m *QueryContractsByCreatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContractsByCreatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContractsByCreatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddresses = append(m.ContractAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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

func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
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
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
