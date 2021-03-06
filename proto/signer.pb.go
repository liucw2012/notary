// Code generated by protoc-gen-go.
// source: proto/signer.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/signer.proto

It has these top-level messages:
	KeyInfo
	KeyID
	Algorithm
	PublicKey
	Signature
	SignatureRequest
	Void
	HealthStatus
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// KeyInfo holds a KeyID that is used to reference the key and it's algorithm
type KeyInfo struct {
	KeyID     *KeyID     `protobuf:"bytes,1,opt,name=keyID" json:"keyID,omitempty"`
	Algorithm *Algorithm `protobuf:"bytes,2,opt,name=algorithm" json:"algorithm,omitempty"`
}

func (m *KeyInfo) Reset()         { *m = KeyInfo{} }
func (m *KeyInfo) String() string { return proto1.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()    {}

func (m *KeyInfo) GetKeyID() *KeyID {
	if m != nil {
		return m.KeyID
	}
	return nil
}

func (m *KeyInfo) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

// KeyID holds an ID that is used to reference the key
type KeyID struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *KeyID) Reset()         { *m = KeyID{} }
func (m *KeyID) String() string { return proto1.CompactTextString(m) }
func (*KeyID) ProtoMessage()    {}

// Type holds the type of crypto algorithm used
type Algorithm struct {
	Algorithm string `protobuf:"bytes,1,opt,name=algorithm" json:"algorithm,omitempty"`
}

func (m *Algorithm) Reset()         { *m = Algorithm{} }
func (m *Algorithm) String() string { return proto1.CompactTextString(m) }
func (*Algorithm) ProtoMessage()    {}

// PublicKey has a KeyInfo that is used to reference the key, and opaque bytes of a publicKey
type PublicKey struct {
	KeyInfo   *KeyInfo `protobuf:"bytes,1,opt,name=keyInfo" json:"keyInfo,omitempty"`
	PublicKey []byte   `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto1.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}

func (m *PublicKey) GetKeyInfo() *KeyInfo {
	if m != nil {
		return m.KeyInfo
	}
	return nil
}

// Signature specifies a KeyInfo that was used for signing and signed content
type Signature struct {
	KeyInfo   *KeyInfo   `protobuf:"bytes,1,opt,name=keyInfo" json:"keyInfo,omitempty"`
	Algorithm *Algorithm `protobuf:"bytes,2,opt,name=algorithm" json:"algorithm,omitempty"`
	Content   []byte     `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto1.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}

func (m *Signature) GetKeyInfo() *KeyInfo {
	if m != nil {
		return m.KeyInfo
	}
	return nil
}

func (m *Signature) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

// SignatureRequests specifies a KeyInfo, and content to be signed
type SignatureRequest struct {
	KeyID   *KeyID `protobuf:"bytes,1,opt,name=keyID" json:"keyID,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *SignatureRequest) Reset()         { *m = SignatureRequest{} }
func (m *SignatureRequest) String() string { return proto1.CompactTextString(m) }
func (*SignatureRequest) ProtoMessage()    {}

func (m *SignatureRequest) GetKeyID() *KeyID {
	if m != nil {
		return m.KeyID
	}
	return nil
}

// Void represents an empty message type
type Void struct {
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto1.CompactTextString(m) }
func (*Void) ProtoMessage()    {}

// A mapping of health check name to the check result message
type HealthStatus struct {
	Status map[string]string `protobuf:"bytes,1,rep,name=status" json:"status,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HealthStatus) Reset()         { *m = HealthStatus{} }
func (m *HealthStatus) String() string { return proto1.CompactTextString(m) }
func (*HealthStatus) ProtoMessage()    {}

func (m *HealthStatus) GetStatus() map[string]string {
	if m != nil {
		return m.Status
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for KeyManagement service

type KeyManagementClient interface {
	// CreateKey creates as asymmetric key pair and returns the PublicKey
	CreateKey(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*PublicKey, error)
	// DeleteKey deletes the key associated with a KeyID
	DeleteKey(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*Void, error)
	// GetKeyInfo returns the PublicKey associated with a KeyID
	GetKeyInfo(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*PublicKey, error)
	// CheckHealth returns the HealthStatus with the service
	CheckHealth(ctx context.Context, in *Void, opts ...grpc.CallOption) (*HealthStatus, error)
}

type keyManagementClient struct {
	cc *grpc.ClientConn
}

func NewKeyManagementClient(cc *grpc.ClientConn) KeyManagementClient {
	return &keyManagementClient{cc}
}

func (c *keyManagementClient) CreateKey(ctx context.Context, in *Algorithm, opts ...grpc.CallOption) (*PublicKey, error) {
	out := new(PublicKey)
	err := grpc.Invoke(ctx, "/proto.KeyManagement/CreateKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementClient) DeleteKey(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/proto.KeyManagement/DeleteKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementClient) GetKeyInfo(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*PublicKey, error) {
	out := new(PublicKey)
	err := grpc.Invoke(ctx, "/proto.KeyManagement/GetKeyInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementClient) CheckHealth(ctx context.Context, in *Void, opts ...grpc.CallOption) (*HealthStatus, error) {
	out := new(HealthStatus)
	err := grpc.Invoke(ctx, "/proto.KeyManagement/CheckHealth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyManagement service

type KeyManagementServer interface {
	// CreateKey creates as asymmetric key pair and returns the PublicKey
	CreateKey(context.Context, *Algorithm) (*PublicKey, error)
	// DeleteKey deletes the key associated with a KeyID
	DeleteKey(context.Context, *KeyID) (*Void, error)
	// GetKeyInfo returns the PublicKey associated with a KeyID
	GetKeyInfo(context.Context, *KeyID) (*PublicKey, error)
	// CheckHealth returns the HealthStatus with the service
	CheckHealth(context.Context, *Void) (*HealthStatus, error)
}

func RegisterKeyManagementServer(s *grpc.Server, srv KeyManagementServer) {
	s.RegisterService(&_KeyManagement_serviceDesc, srv)
}

func _KeyManagement_CreateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Algorithm)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(KeyManagementServer).CreateKey(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _KeyManagement_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(KeyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(KeyManagementServer).DeleteKey(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _KeyManagement_GetKeyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(KeyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(KeyManagementServer).GetKeyInfo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _KeyManagement_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(KeyManagementServer).CheckHealth(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _KeyManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KeyManagement",
	HandlerType: (*KeyManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKey",
			Handler:    _KeyManagement_CreateKey_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _KeyManagement_DeleteKey_Handler,
		},
		{
			MethodName: "GetKeyInfo",
			Handler:    _KeyManagement_GetKeyInfo_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _KeyManagement_CheckHealth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Signer service

type SignerClient interface {
	// Sign calculates a cryptographic signature using the Key associated with a KeyID and returns the signature
	Sign(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*Signature, error)
}

type signerClient struct {
	cc *grpc.ClientConn
}

func NewSignerClient(cc *grpc.ClientConn) SignerClient {
	return &signerClient{cc}
}

func (c *signerClient) Sign(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*Signature, error) {
	out := new(Signature)
	err := grpc.Invoke(ctx, "/proto.Signer/Sign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Signer service

type SignerServer interface {
	// Sign calculates a cryptographic signature using the Key associated with a KeyID and returns the signature
	Sign(context.Context, *SignatureRequest) (*Signature, error)
}

func RegisterSignerServer(s *grpc.Server, srv SignerServer) {
	s.RegisterService(&_Signer_serviceDesc, srv)
}

func _Signer_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SignerServer).Sign(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Signer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Signer",
	HandlerType: (*SignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _Signer_Sign_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
