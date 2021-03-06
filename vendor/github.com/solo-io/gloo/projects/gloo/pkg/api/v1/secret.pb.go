// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Certain features such as the AWS Lambda option require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.
//
//Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:
//
//- Kubernetes Secrets
//- Hashicorp Vault
//- Plaintext files (recommended only for testing)
//- Secrets must adhere to a structure, specified by the option that requires them.
//
//Gloo's secret backend can be configured in Gloo's bootstrap options
type Secret struct {
	// Types that are valid to be assigned to Kind:
	//	*Secret_Aws
	//	*Secret_Azure
	//	*Secret_Tls
	//	*Secret_Oauth
	//	*Secret_ApiKey
	//	*Secret_Extensions
	Kind isSecret_Kind `protobuf_oneof:"kind"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

type isSecret_Kind interface {
	isSecret_Kind()
	Equal(interface{}) bool
}

type Secret_Aws struct {
	Aws *AwsSecret `protobuf:"bytes,1,opt,name=aws,proto3,oneof" json:"aws,omitempty"`
}
type Secret_Azure struct {
	Azure *AzureSecret `protobuf:"bytes,2,opt,name=azure,proto3,oneof" json:"azure,omitempty"`
}
type Secret_Tls struct {
	Tls *TlsSecret `protobuf:"bytes,3,opt,name=tls,proto3,oneof" json:"tls,omitempty"`
}
type Secret_Oauth struct {
	Oauth *v1.OauthSecret `protobuf:"bytes,5,opt,name=oauth,proto3,oneof" json:"oauth,omitempty"`
}
type Secret_ApiKey struct {
	ApiKey *v1.ApiKeySecret `protobuf:"bytes,6,opt,name=api_key,json=apiKey,proto3,oneof" json:"api_key,omitempty"`
}
type Secret_Extensions struct {
	Extensions *Extensions `protobuf:"bytes,4,opt,name=extensions,proto3,oneof" json:"extensions,omitempty"`
}

func (*Secret_Aws) isSecret_Kind()        {}
func (*Secret_Azure) isSecret_Kind()      {}
func (*Secret_Tls) isSecret_Kind()        {}
func (*Secret_Oauth) isSecret_Kind()      {}
func (*Secret_ApiKey) isSecret_Kind()     {}
func (*Secret_Extensions) isSecret_Kind() {}

func (m *Secret) GetKind() isSecret_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Secret) GetAws() *AwsSecret {
	if x, ok := m.GetKind().(*Secret_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Secret) GetAzure() *AzureSecret {
	if x, ok := m.GetKind().(*Secret_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Secret) GetTls() *TlsSecret {
	if x, ok := m.GetKind().(*Secret_Tls); ok {
		return x.Tls
	}
	return nil
}

func (m *Secret) GetOauth() *v1.OauthSecret {
	if x, ok := m.GetKind().(*Secret_Oauth); ok {
		return x.Oauth
	}
	return nil
}

func (m *Secret) GetApiKey() *v1.ApiKeySecret {
	if x, ok := m.GetKind().(*Secret_ApiKey); ok {
		return x.ApiKey
	}
	return nil
}

func (m *Secret) GetExtensions() *Extensions {
	if x, ok := m.GetKind().(*Secret_Extensions); ok {
		return x.Extensions
	}
	return nil
}

func (m *Secret) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Secret) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Secret_Aws)(nil),
		(*Secret_Azure)(nil),
		(*Secret_Tls)(nil),
		(*Secret_Oauth)(nil),
		(*Secret_ApiKey)(nil),
		(*Secret_Extensions)(nil),
	}
}

//
//
//There are two ways of providing AWS secrets:
//
//- Method 1: `glooctl create secret aws`
//
// ```
// glooctl create secret aws --name aws-secret-from-glooctl \
//     --namespace default \
//     --access-key $ACC \
//     --secret-key $SEC
// ```
//
//will produce a Kubernetes resource similar to this (note the `aws` field and `resource_kind` annotation):
//
// ```
// apiVersion: v1
// data:
//   aws: base64EncodedStringForMachineConsumption
// kind: Secret
// metadata:
//   annotations:
//     resource_kind: '*v1.Secret'
//   creationTimestamp: "2019-08-23T15:10:20Z"
//   name: aws-secret-from-glooctl
//   namespace: default
//   resourceVersion: "592637"
//   selfLink: /api/v1/namespaces/default/secrets/secret-e2e
//   uid: 1f8c147f-c5b8-11e9-bbf3-42010a8001bc
// type: Opaque
// ```
//
// - Method 2: `kubectl apply -f resource-file.yaml`
//   - If using a git-ops flow, or otherwise creating secrets from yaml files, you may prefer to provide AWS credentials
//   using the format below, with `aws_access_key_id` and `aws_secret_access_key` fields.
//   - This circumvents the need for the annotation, which are not supported by some tools such as
//   [godaddy/kubernetes-external-secrets](https://github.com/godaddy/kubernetes-external-secrets)
//
// ```yaml
// # a sample aws secret resource-file.yaml
// apiVersion: v1
// data:
//   aws_access_key_id: some-id
//   aws_secret_access_key: some-secret
// kind: Secret
// metadata:
//   name: aws-secret-abcd
//   namespace: default
// ```
//
type AwsSecret struct {
	// provided by `glooctl create secret aws`
	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	// provided by `glooctl create secret aws`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsSecret) Reset()         { *m = AwsSecret{} }
func (m *AwsSecret) String() string { return proto.CompactTextString(m) }
func (*AwsSecret) ProtoMessage()    {}
func (*AwsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{1}
}
func (m *AwsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsSecret.Unmarshal(m, b)
}
func (m *AwsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsSecret.Marshal(b, m, deterministic)
}
func (m *AwsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsSecret.Merge(m, src)
}
func (m *AwsSecret) XXX_Size() int {
	return xxx_messageInfo_AwsSecret.Size(m)
}
func (m *AwsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AwsSecret proto.InternalMessageInfo

func (m *AwsSecret) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AwsSecret) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type AzureSecret struct {
	// provided by `glooctl create secret azure`
	ApiKeys              map[string]string `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AzureSecret) Reset()         { *m = AzureSecret{} }
func (m *AzureSecret) String() string { return proto.CompactTextString(m) }
func (*AzureSecret) ProtoMessage()    {}
func (*AzureSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{2}
}
func (m *AzureSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSecret.Unmarshal(m, b)
}
func (m *AzureSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSecret.Marshal(b, m, deterministic)
}
func (m *AzureSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSecret.Merge(m, src)
}
func (m *AzureSecret) XXX_Size() int {
	return xxx_messageInfo_AzureSecret.Size(m)
}
func (m *AzureSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSecret proto.InternalMessageInfo

func (m *AzureSecret) GetApiKeys() map[string]string {
	if m != nil {
		return m.ApiKeys
	}
	return nil
}

//
//Note that the annotation `resource_kind: '*v1.Secret'` is needed for Gloo to find this secret.
//Glooctl adds it by default when the tls secret is created via `glooctl create secret tls`.
type TlsSecret struct {
	// provided by `glooctl create secret tls`
	CertChain string `protobuf:"bytes,1,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	// provided by `glooctl create secret tls`
	PrivateKey string `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// provided by `glooctl create secret tls`
	RootCa               string   `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsSecret) Reset()         { *m = TlsSecret{} }
func (m *TlsSecret) String() string { return proto.CompactTextString(m) }
func (*TlsSecret) ProtoMessage()    {}
func (*TlsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{3}
}
func (m *TlsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSecret.Unmarshal(m, b)
}
func (m *TlsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSecret.Marshal(b, m, deterministic)
}
func (m *TlsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSecret.Merge(m, src)
}
func (m *TlsSecret) XXX_Size() int {
	return xxx_messageInfo_TlsSecret.Size(m)
}
func (m *TlsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSecret proto.InternalMessageInfo

func (m *TlsSecret) GetCertChain() string {
	if m != nil {
		return m.CertChain
	}
	return ""
}

func (m *TlsSecret) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TlsSecret) GetRootCa() string {
	if m != nil {
		return m.RootCa
	}
	return ""
}

func init() {
	proto.RegisterType((*Secret)(nil), "gloo.solo.io.Secret")
	proto.RegisterType((*AwsSecret)(nil), "gloo.solo.io.AwsSecret")
	proto.RegisterType((*AzureSecret)(nil), "gloo.solo.io.AzureSecret")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.AzureSecret.ApiKeysEntry")
	proto.RegisterType((*TlsSecret)(nil), "gloo.solo.io.TlsSecret")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto", fileDescriptor_c2f79c35f1213791)
}

var fileDescriptor_c2f79c35f1213791 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x8e, 0x63, 0xc7, 0xa9, 0x4f, 0xba, 0xf8, 0x35, 0xaa, 0x5a, 0xff, 0x91, 0xda, 0xa2, 0x08,
	0x50, 0x05, 0xc2, 0xa6, 0x65, 0x53, 0x22, 0x16, 0x24, 0x55, 0xa4, 0xa2, 0x0a, 0x21, 0x19, 0x56,
	0x6c, 0xa2, 0xa9, 0x33, 0x4a, 0x86, 0xb8, 0x1e, 0x6b, 0x66, 0x92, 0x26, 0x2c, 0xbb, 0xe6, 0x41,
	0x90, 0x78, 0x01, 0x1e, 0x81, 0xa7, 0x60, 0xc1, 0x1b, 0x74, 0xc1, 0x1e, 0xcd, 0x25, 0x8e, 0x8b,
	0x08, 0x62, 0x37, 0xe7, 0x7c, 0x97, 0xcc, 0x77, 0xce, 0x38, 0xf0, 0x7c, 0x4c, 0xe5, 0x64, 0x76,
	0x19, 0xa5, 0xec, 0x2a, 0x16, 0x2c, 0x63, 0x4f, 0x28, 0x8b, 0xc7, 0x19, 0x63, 0x71, 0xc1, 0xd9,
	0x07, 0x92, 0x4a, 0x61, 0x2a, 0x5c, 0xd0, 0x78, 0x7e, 0x1c, 0x0b, 0x92, 0x72, 0x22, 0xa3, 0x82,
	0x33, 0xc9, 0xd0, 0xb6, 0x42, 0x22, 0x25, 0x8a, 0x28, 0x6b, 0xef, 0x8c, 0xd9, 0x98, 0x69, 0x20,
	0x56, 0x27, 0xc3, 0x69, 0x23, 0xb2, 0x90, 0xa6, 0x49, 0x16, 0x56, 0xd7, 0x7e, 0xb4, 0xd9, 0x9f,
	0x2c, 0x24, 0xc9, 0x05, 0x65, 0xb9, 0xb0, 0xdc, 0xc1, 0x5f, 0xb8, 0xb9, 0x24, 0xbc, 0xe0, 0x54,
	0x90, 0x98, 0x15, 0x52, 0x69, 0x94, 0x1c, 0xcf, 0xe4, 0xc4, 0x3a, 0xa9, 0xa3, 0xb5, 0x39, 0xd0,
	0xd1, 0xa6, 0x54, 0xae, 0xc4, 0x57, 0x44, 0xe2, 0x11, 0x96, 0x78, 0x13, 0xbe, 0xaa, 0x0d, 0xde,
	0xf9, 0xe2, 0x82, 0xff, 0x56, 0x67, 0x47, 0x8f, 0xc1, 0xc5, 0xd7, 0x22, 0x74, 0xee, 0x39, 0x47,
	0xad, 0x93, 0xbd, 0xa8, 0x3a, 0x83, 0xa8, 0x77, 0x2d, 0x0c, 0xeb, 0xbc, 0x96, 0x28, 0x16, 0x3a,
	0x86, 0x06, 0xfe, 0x38, 0xe3, 0x24, 0xac, 0x6b, 0xfa, 0xff, 0xbf, 0xd1, 0x15, 0x54, 0x0a, 0x0c,
	0x53, 0xf9, 0xcb, 0x4c, 0x84, 0xee, 0x9f, 0xfc, 0xdf, 0x65, 0x15, 0x7f, 0x99, 0x09, 0xf4, 0x02,
	0x1a, 0x4c, 0xc5, 0x0c, 0x1b, 0x9a, 0x7e, 0x3f, 0x5a, 0x0f, 0xe5, 0xae, 0xf2, 0x8d, 0x62, 0xad,
	0x7f, 0x4a, 0x8b, 0xd0, 0x4b, 0x68, 0xe2, 0x82, 0x0e, 0xa7, 0x64, 0x19, 0xfa, 0x5a, 0xff, 0x60,
	0xa3, 0xbe, 0x57, 0xd0, 0x0b, 0xb2, 0x2c, 0x0d, 0x7c, 0xac, 0x6b, 0xd4, 0x05, 0x58, 0xaf, 0x2c,
	0xf4, 0xb4, 0x49, 0x78, 0x57, 0x39, 0x28, 0xf1, 0xf3, 0x5a, 0x52, 0x61, 0xa3, 0x53, 0xd8, 0x5a,
	0x6d, 0x21, 0x6c, 0x6a, 0xe5, 0x6e, 0x94, 0x32, 0x4e, 0x4a, 0xe5, 0x6b, 0x8b, 0xf6, 0xbd, 0x6f,
	0xdf, 0x0f, 0x6b, 0x49, 0xc9, 0xee, 0xee, 0xde, 0xdc, 0x7a, 0x0d, 0x70, 0x05, 0x49, 0x6f, 0x6e,
	0xbd, 0x00, 0x35, 0xcd, 0xab, 0x14, 0x7d, 0x1f, 0xbc, 0x29, 0xcd, 0x47, 0x9d, 0x57, 0x10, 0x94,
	0x9b, 0x40, 0xfb, 0x00, 0x38, 0x4d, 0x89, 0x10, 0x3a, 0xa7, 0x5a, 0x5b, 0x90, 0x04, 0xa6, 0xa3,
	0x12, 0xec, 0x03, 0x18, 0xb9, 0x86, 0xeb, 0x06, 0x36, 0x9d, 0x0b, 0xb2, 0xec, 0x7c, 0x72, 0xa0,
	0x55, 0x59, 0x13, 0xea, 0xc1, 0x96, 0x1d, 0x99, 0x7a, 0x02, 0xee, 0x51, 0xeb, 0xe4, 0xe1, 0xc6,
	0x9d, 0xda, 0xa1, 0x89, 0x41, 0x2e, 0xf9, 0x32, 0x69, 0x9a, 0x91, 0x89, 0x76, 0x17, 0xb6, 0xab,
	0x00, 0xfa, 0x0f, 0xdc, 0xf5, 0xcd, 0xd4, 0x11, 0xed, 0x40, 0x63, 0x8e, 0xb3, 0x19, 0xb1, 0xd7,
	0x31, 0x45, 0xb7, 0x7e, 0xea, 0x74, 0x46, 0x10, 0x94, 0x6f, 0x40, 0x5d, 0x3d, 0x25, 0x5c, 0x0e,
	0xd3, 0x09, 0xa6, 0xf9, 0x2a, 0x99, 0xea, 0x9c, 0xa9, 0x06, 0x3a, 0x84, 0x56, 0xc1, 0xe9, 0x1c,
	0x4b, 0x52, 0x89, 0x06, 0xb6, 0xa5, 0xa2, 0xef, 0x41, 0x93, 0x33, 0x26, 0x87, 0x29, 0xd6, 0xaf,
	0x2d, 0x48, 0x7c, 0x55, 0x9e, 0xe1, 0x7e, 0xf7, 0xeb, 0x4f, 0xcf, 0xf9, 0xfc, 0xe3, 0xc0, 0x79,
	0xff, 0xf4, 0xdf, 0xfe, 0x1d, 0x8a, 0xe9, 0xd8, 0x7e, 0x38, 0x97, 0xbe, 0xfe, 0x60, 0x9e, 0xfd,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x53, 0xe8, 0xd1, 0x63, 0x58, 0x04, 0x00, 0x00,
}

func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
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
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Secret_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Aws)
	if !ok {
		that2, ok := that.(Secret_Aws)
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
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *Secret_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Azure)
	if !ok {
		that2, ok := that.(Secret_Azure)
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
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *Secret_Tls) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Tls)
	if !ok {
		that2, ok := that.(Secret_Tls)
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
	if !this.Tls.Equal(that1.Tls) {
		return false
	}
	return true
}
func (this *Secret_Oauth) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Oauth)
	if !ok {
		that2, ok := that.(Secret_Oauth)
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
	if !this.Oauth.Equal(that1.Oauth) {
		return false
	}
	return true
}
func (this *Secret_ApiKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_ApiKey)
	if !ok {
		that2, ok := that.(Secret_ApiKey)
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
	if !this.ApiKey.Equal(that1.ApiKey) {
		return false
	}
	return true
}
func (this *Secret_Extensions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Extensions)
	if !ok {
		that2, ok := that.(Secret_Extensions)
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
	if !this.Extensions.Equal(that1.Extensions) {
		return false
	}
	return true
}
func (this *AwsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsSecret)
	if !ok {
		that2, ok := that.(AwsSecret)
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
	if this.AccessKey != that1.AccessKey {
		return false
	}
	if this.SecretKey != that1.SecretKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AzureSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AzureSecret)
	if !ok {
		that2, ok := that.(AzureSecret)
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
	if len(this.ApiKeys) != len(that1.ApiKeys) {
		return false
	}
	for i := range this.ApiKeys {
		if this.ApiKeys[i] != that1.ApiKeys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TlsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TlsSecret)
	if !ok {
		that2, ok := that.(TlsSecret)
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
	if this.CertChain != that1.CertChain {
		return false
	}
	if this.PrivateKey != that1.PrivateKey {
		return false
	}
	if this.RootCa != that1.RootCa {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
