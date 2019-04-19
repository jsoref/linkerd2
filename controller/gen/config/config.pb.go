// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config // import "github.com/linkerd/linkerd2/controller/gen/config"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type All struct {
	Global               *Global  `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	Proxy                *Proxy   `protobuf:"bytes,2,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Install              *Install `protobuf:"bytes,3,opt,name=install,proto3" json:"install,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{0}
}
func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (dst *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(dst, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetGlobal() *Global {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *All) GetProxy() *Proxy {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *All) GetInstall() *Install {
	if m != nil {
		return m.Install
	}
	return nil
}

type Global struct {
	LinkerdNamespace string `protobuf:"bytes,1,opt,name=linkerd_namespace,json=linkerdNamespace,proto3" json:"linkerd_namespace,omitempty"`
	CniEnabled       bool   `protobuf:"varint,2,opt,name=cni_enabled,json=cniEnabled,proto3" json:"cni_enabled,omitempty"`
	// Control plane and proxy-init version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// If present, configures identity.
	IdentityContext *IdentityContext `protobuf:"bytes,4,opt,name=identity_context,json=identityContext,proto3" json:"identity_context,omitempty"`
	// If present, indicates that the Mutating Webhook Admission Controller should
	// be configured to automatically inject proxies.
	AutoInjectContext    *AutoInjectContext `protobuf:"bytes,6,opt,name=auto_inject_context,json=autoInjectContext,proto3" json:"auto_inject_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{1}
}
func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (dst *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(dst, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetLinkerdNamespace() string {
	if m != nil {
		return m.LinkerdNamespace
	}
	return ""
}

func (m *Global) GetCniEnabled() bool {
	if m != nil {
		return m.CniEnabled
	}
	return false
}

func (m *Global) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Global) GetIdentityContext() *IdentityContext {
	if m != nil {
		return m.IdentityContext
	}
	return nil
}

func (m *Global) GetAutoInjectContext() *AutoInjectContext {
	if m != nil {
		return m.AutoInjectContext
	}
	return nil
}

type Proxy struct {
	ProxyImage              *Image                `protobuf:"bytes,1,opt,name=proxy_image,json=proxyImage,proto3" json:"proxy_image,omitempty"`
	ProxyInitImage          *Image                `protobuf:"bytes,2,opt,name=proxy_init_image,json=proxyInitImage,proto3" json:"proxy_init_image,omitempty"`
	ControlPort             *Port                 `protobuf:"bytes,3,opt,name=control_port,json=controlPort,proto3" json:"control_port,omitempty"`
	IgnoreInboundPorts      []*Port               `protobuf:"bytes,4,rep,name=ignore_inbound_ports,json=ignoreInboundPorts,proto3" json:"ignore_inbound_ports,omitempty"`
	IgnoreOutboundPorts     []*Port               `protobuf:"bytes,5,rep,name=ignore_outbound_ports,json=ignoreOutboundPorts,proto3" json:"ignore_outbound_ports,omitempty"`
	InboundPort             *Port                 `protobuf:"bytes,6,opt,name=inbound_port,json=inboundPort,proto3" json:"inbound_port,omitempty"`
	AdminPort               *Port                 `protobuf:"bytes,7,opt,name=admin_port,json=adminPort,proto3" json:"admin_port,omitempty"`
	OutboundPort            *Port                 `protobuf:"bytes,8,opt,name=outbound_port,json=outboundPort,proto3" json:"outbound_port,omitempty"`
	Resource                *ResourceRequirements `protobuf:"bytes,9,opt,name=resource,proto3" json:"resource,omitempty"`
	ProxyUid                int64                 `protobuf:"varint,10,opt,name=proxy_uid,json=proxyUid,proto3" json:"proxy_uid,omitempty"`
	LogLevel                *LogLevel             `protobuf:"bytes,11,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	DisableExternalProfiles bool                  `protobuf:"varint,12,opt,name=disable_external_profiles,json=disableExternalProfiles,proto3" json:"disable_external_profiles,omitempty"`
	ProxyVersion            string                `protobuf:"bytes,13,opt,name=proxy_version,json=proxyVersion,proto3" json:"proxy_version,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{2}
}
func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxy.Unmarshal(m, b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
}
func (dst *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(dst, src)
}
func (m *Proxy) XXX_Size() int {
	return xxx_messageInfo_Proxy.Size(m)
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

func (m *Proxy) GetProxyImage() *Image {
	if m != nil {
		return m.ProxyImage
	}
	return nil
}

func (m *Proxy) GetProxyInitImage() *Image {
	if m != nil {
		return m.ProxyInitImage
	}
	return nil
}

func (m *Proxy) GetControlPort() *Port {
	if m != nil {
		return m.ControlPort
	}
	return nil
}

func (m *Proxy) GetIgnoreInboundPorts() []*Port {
	if m != nil {
		return m.IgnoreInboundPorts
	}
	return nil
}

func (m *Proxy) GetIgnoreOutboundPorts() []*Port {
	if m != nil {
		return m.IgnoreOutboundPorts
	}
	return nil
}

func (m *Proxy) GetInboundPort() *Port {
	if m != nil {
		return m.InboundPort
	}
	return nil
}

func (m *Proxy) GetAdminPort() *Port {
	if m != nil {
		return m.AdminPort
	}
	return nil
}

func (m *Proxy) GetOutboundPort() *Port {
	if m != nil {
		return m.OutboundPort
	}
	return nil
}

func (m *Proxy) GetResource() *ResourceRequirements {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Proxy) GetProxyUid() int64 {
	if m != nil {
		return m.ProxyUid
	}
	return 0
}

func (m *Proxy) GetLogLevel() *LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func (m *Proxy) GetDisableExternalProfiles() bool {
	if m != nil {
		return m.DisableExternalProfiles
	}
	return false
}

func (m *Proxy) GetProxyVersion() string {
	if m != nil {
		return m.ProxyVersion
	}
	return ""
}

type Image struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	PullPolicy           string   `protobuf:"bytes,2,opt,name=pull_policy,json=pullPolicy,proto3" json:"pull_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{3}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Image) GetPullPolicy() string {
	if m != nil {
		return m.PullPolicy
	}
	return ""
}

type Port struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{4}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (dst *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(dst, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ResourceRequirements struct {
	RequestCpu           string   `protobuf:"bytes,1,opt,name=request_cpu,json=requestCpu,proto3" json:"request_cpu,omitempty"`
	RequestMemory        string   `protobuf:"bytes,2,opt,name=request_memory,json=requestMemory,proto3" json:"request_memory,omitempty"`
	LimitCpu             string   `protobuf:"bytes,3,opt,name=limit_cpu,json=limitCpu,proto3" json:"limit_cpu,omitempty"`
	LimitMemory          string   `protobuf:"bytes,4,opt,name=limit_memory,json=limitMemory,proto3" json:"limit_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{5}
}
func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (dst *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(dst, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetRequestCpu() string {
	if m != nil {
		return m.RequestCpu
	}
	return ""
}

func (m *ResourceRequirements) GetRequestMemory() string {
	if m != nil {
		return m.RequestMemory
	}
	return ""
}

func (m *ResourceRequirements) GetLimitCpu() string {
	if m != nil {
		return m.LimitCpu
	}
	return ""
}

func (m *ResourceRequirements) GetLimitMemory() string {
	if m != nil {
		return m.LimitMemory
	}
	return ""
}

// Currently, this is basically a boolean.
type AutoInjectContext struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoInjectContext) Reset()         { *m = AutoInjectContext{} }
func (m *AutoInjectContext) String() string { return proto.CompactTextString(m) }
func (*AutoInjectContext) ProtoMessage()    {}
func (*AutoInjectContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{6}
}
func (m *AutoInjectContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoInjectContext.Unmarshal(m, b)
}
func (m *AutoInjectContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoInjectContext.Marshal(b, m, deterministic)
}
func (dst *AutoInjectContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoInjectContext.Merge(dst, src)
}
func (m *AutoInjectContext) XXX_Size() int {
	return xxx_messageInfo_AutoInjectContext.Size(m)
}
func (m *AutoInjectContext) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoInjectContext.DiscardUnknown(m)
}

var xxx_messageInfo_AutoInjectContext proto.InternalMessageInfo

type IdentityContext struct {
	TrustDomain          string             `protobuf:"bytes,1,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	TrustAnchorsPem      string             `protobuf:"bytes,2,opt,name=trust_anchors_pem,json=trustAnchorsPem,proto3" json:"trust_anchors_pem,omitempty"`
	IssuanceLifetime     *duration.Duration `protobuf:"bytes,3,opt,name=issuance_lifetime,json=issuanceLifetime,proto3" json:"issuance_lifetime,omitempty"`
	ClockSkewAllowance   *duration.Duration `protobuf:"bytes,4,opt,name=clock_skew_allowance,json=clockSkewAllowance,proto3" json:"clock_skew_allowance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IdentityContext) Reset()         { *m = IdentityContext{} }
func (m *IdentityContext) String() string { return proto.CompactTextString(m) }
func (*IdentityContext) ProtoMessage()    {}
func (*IdentityContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{7}
}
func (m *IdentityContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityContext.Unmarshal(m, b)
}
func (m *IdentityContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityContext.Marshal(b, m, deterministic)
}
func (dst *IdentityContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityContext.Merge(dst, src)
}
func (m *IdentityContext) XXX_Size() int {
	return xxx_messageInfo_IdentityContext.Size(m)
}
func (m *IdentityContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityContext.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityContext proto.InternalMessageInfo

func (m *IdentityContext) GetTrustDomain() string {
	if m != nil {
		return m.TrustDomain
	}
	return ""
}

func (m *IdentityContext) GetTrustAnchorsPem() string {
	if m != nil {
		return m.TrustAnchorsPem
	}
	return ""
}

func (m *IdentityContext) GetIssuanceLifetime() *duration.Duration {
	if m != nil {
		return m.IssuanceLifetime
	}
	return nil
}

func (m *IdentityContext) GetClockSkewAllowance() *duration.Duration {
	if m != nil {
		return m.ClockSkewAllowance
	}
	return nil
}

type LogLevel struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{8}
}
func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (dst *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(dst, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

func (m *LogLevel) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

// Stores information about the last installation/upgrade.
//
// Useful for driving upgrades.
type Install struct {
	// The unique ID for this installation. Does not change on upgrade.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// The CLI version that drove the last install or upgrade.
	CliVersion string `protobuf:"bytes,2,opt,name=cli_version,json=cliVersion,proto3" json:"cli_version,omitempty"`
	// The CLI arguments to the install (or upgrade) command, indicating the
	// installer's intent.
	Flags                []*Install_Flag `protobuf:"bytes,3,rep,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{9}
}
func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (dst *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(dst, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

func (m *Install) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Install) GetCliVersion() string {
	if m != nil {
		return m.CliVersion
	}
	return ""
}

func (m *Install) GetFlags() []*Install_Flag {
	if m != nil {
		return m.Flags
	}
	return nil
}

type Install_Flag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install_Flag) Reset()         { *m = Install_Flag{} }
func (m *Install_Flag) String() string { return proto.CompactTextString(m) }
func (*Install_Flag) ProtoMessage()    {}
func (*Install_Flag) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_1dc3723f04bb94b9, []int{9, 0}
}
func (m *Install_Flag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install_Flag.Unmarshal(m, b)
}
func (m *Install_Flag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install_Flag.Marshal(b, m, deterministic)
}
func (dst *Install_Flag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install_Flag.Merge(dst, src)
}
func (m *Install_Flag) XXX_Size() int {
	return xxx_messageInfo_Install_Flag.Size(m)
}
func (m *Install_Flag) XXX_DiscardUnknown() {
	xxx_messageInfo_Install_Flag.DiscardUnknown(m)
}

var xxx_messageInfo_Install_Flag proto.InternalMessageInfo

func (m *Install_Flag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Install_Flag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*All)(nil), "linkerd2.config.All")
	proto.RegisterType((*Global)(nil), "linkerd2.config.Global")
	proto.RegisterType((*Proxy)(nil), "linkerd2.config.Proxy")
	proto.RegisterType((*Image)(nil), "linkerd2.config.Image")
	proto.RegisterType((*Port)(nil), "linkerd2.config.Port")
	proto.RegisterType((*ResourceRequirements)(nil), "linkerd2.config.ResourceRequirements")
	proto.RegisterType((*AutoInjectContext)(nil), "linkerd2.config.AutoInjectContext")
	proto.RegisterType((*IdentityContext)(nil), "linkerd2.config.IdentityContext")
	proto.RegisterType((*LogLevel)(nil), "linkerd2.config.LogLevel")
	proto.RegisterType((*Install)(nil), "linkerd2.config.Install")
	proto.RegisterType((*Install_Flag)(nil), "linkerd2.config.Install.Flag")
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_config_1dc3723f04bb94b9) }

var fileDescriptor_config_1dc3723f04bb94b9 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xcf, 0x73, 0x1b, 0x35,
	0x14, 0x1e, 0xd7, 0x76, 0x62, 0x3f, 0xdb, 0x4d, 0xac, 0xa4, 0x74, 0x13, 0xa6, 0x60, 0x96, 0xe9,
	0x4c, 0x07, 0x18, 0x1b, 0x12, 0x06, 0x3a, 0x3d, 0x61, 0xfa, 0x23, 0xe3, 0x69, 0x80, 0x8c, 0x18,
	0x38, 0x70, 0xd9, 0x59, 0xef, 0xca, 0x5b, 0x11, 0xad, 0xe4, 0x6a, 0xa5, 0x24, 0xfd, 0x33, 0xb8,
	0x71, 0xe2, 0xc6, 0x9f, 0xc8, 0x9d, 0xd1, 0x93, 0x36, 0xa4, 0x59, 0xe2, 0xd3, 0x4a, 0xdf, 0xfb,
	0xbe, 0x4f, 0x6f, 0xa5, 0xa7, 0x27, 0xd8, 0xcb, 0x94, 0x5c, 0xf1, 0x62, 0xe6, 0x3f, 0xd3, 0xb5,
	0x56, 0x46, 0x91, 0x1d, 0xc1, 0xe5, 0x39, 0xd3, 0xf9, 0xd1, 0xd4, 0xc3, 0x87, 0x1f, 0x15, 0x4a,
	0x15, 0x82, 0xcd, 0x30, 0xbc, 0xb4, 0xab, 0x59, 0x6e, 0x75, 0x6a, 0xb8, 0x92, 0x5e, 0x10, 0xff,
	0xd9, 0x82, 0xf6, 0x5c, 0x08, 0x32, 0x83, 0xad, 0x42, 0xa8, 0x65, 0x2a, 0xa2, 0xd6, 0xa4, 0xf5,
	0x64, 0x70, 0xf4, 0x70, 0x7a, 0xcb, 0x69, 0x7a, 0x82, 0x61, 0x1a, 0x68, 0xe4, 0x0b, 0xe8, 0xae,
	0xb5, 0xba, 0x7a, 0x17, 0xdd, 0x43, 0xfe, 0x07, 0x0d, 0xfe, 0x99, 0x8b, 0x52, 0x4f, 0x22, 0x47,
	0xb0, 0xcd, 0x65, 0x65, 0x52, 0x21, 0xa2, 0x36, 0xf2, 0xa3, 0x06, 0x7f, 0xe1, 0xe3, 0xb4, 0x26,
	0xc6, 0x7f, 0xdc, 0x83, 0x2d, 0xbf, 0x28, 0xf9, 0x1c, 0xc6, 0x81, 0x9e, 0xc8, 0xb4, 0x64, 0xd5,
	0x3a, 0xcd, 0x18, 0x26, 0xda, 0xa7, 0xbb, 0x21, 0xf0, 0x63, 0x8d, 0x93, 0x8f, 0x61, 0x90, 0x49,
	0x9e, 0x30, 0x99, 0x2e, 0x05, 0xcb, 0x31, 0xbf, 0x1e, 0x85, 0x4c, 0xf2, 0x97, 0x1e, 0x21, 0x11,
	0x6c, 0x5f, 0x30, 0x5d, 0x71, 0x25, 0x31, 0x99, 0x3e, 0xad, 0xa7, 0xe4, 0x35, 0xec, 0xf2, 0x9c,
	0x49, 0xc3, 0xcd, 0xbb, 0x24, 0x53, 0xd2, 0xb0, 0x2b, 0x13, 0x75, 0x30, 0xdf, 0x49, 0x33, 0xdf,
	0x40, 0x7c, 0xee, 0x79, 0x74, 0x87, 0xbf, 0x0f, 0x10, 0x0a, 0x7b, 0xa9, 0x35, 0x2a, 0xe1, 0xf2,
	0x77, 0x96, 0x99, 0x6b, 0xbf, 0x2d, 0xf4, 0x8b, 0x1b, 0x7e, 0x73, 0x6b, 0xd4, 0x02, 0xa9, 0xb5,
	0xe3, 0x38, 0xbd, 0x0d, 0xc5, 0xff, 0x74, 0xa1, 0x8b, 0x1b, 0x4b, 0xbe, 0x85, 0x01, 0x6e, 0x6d,
	0xc2, 0xcb, 0xb4, 0x60, 0xe1, 0xd4, 0x9a, 0xa7, 0xb0, 0x70, 0x51, 0x0a, 0x48, 0xc5, 0x31, 0xf9,
	0x0e, 0x76, 0x83, 0x50, 0x72, 0x13, 0xd4, 0xf7, 0x36, 0xaa, 0xef, 0x7b, 0xb5, 0xe4, 0xc6, 0x3b,
	0x3c, 0x85, 0xa1, 0xfb, 0x19, 0xad, 0x44, 0xb2, 0x56, 0xda, 0x84, 0x13, 0x7d, 0xd0, 0xac, 0x00,
	0xa5, 0x0d, 0x1d, 0x04, 0xaa, 0x9b, 0x90, 0x13, 0xd8, 0xe7, 0x85, 0x54, 0x9a, 0x25, 0x5c, 0x2e,
	0x95, 0x95, 0x39, 0x1a, 0x54, 0x51, 0x67, 0xd2, 0xbe, 0xdb, 0x81, 0x78, 0xc9, 0xc2, 0x2b, 0x1c,
	0x54, 0x91, 0x05, 0x3c, 0x08, 0x46, 0xca, 0x9a, 0x9b, 0x4e, 0xdd, 0x4d, 0x4e, 0x7b, 0x5e, 0xf3,
	0x53, 0x90, 0x78, 0xab, 0xa7, 0x30, 0xbc, 0x99, 0x4c, 0x38, 0x9f, 0xbb, 0xfe, 0x86, 0xff, 0x97,
	0x05, 0xf9, 0x1a, 0x20, 0xcd, 0x4b, 0x2e, 0xbd, 0x6e, 0x7b, 0x93, 0xae, 0x8f, 0x44, 0x54, 0x3d,
	0x83, 0xd1, 0x7b, 0x39, 0x47, 0xbd, 0x4d, 0xc2, 0xa1, 0xba, 0x91, 0x2c, 0x99, 0x43, 0x4f, 0xb3,
	0x4a, 0x59, 0x9d, 0xb1, 0xa8, 0x8f, 0xb2, 0xc7, 0x0d, 0x19, 0x0d, 0x04, 0xca, 0xde, 0x5a, 0xae,
	0x59, 0xc9, 0xa4, 0xa9, 0xe8, 0xb5, 0x8c, 0x7c, 0x08, 0x7d, 0x7f, 0xfc, 0x96, 0xe7, 0x11, 0x4c,
	0x5a, 0x4f, 0xda, 0xb4, 0x87, 0xc0, 0x2f, 0x3c, 0x27, 0xdf, 0x40, 0x5f, 0xa8, 0x22, 0x11, 0xec,
	0x82, 0x89, 0x68, 0x80, 0x0b, 0x1c, 0x34, 0x16, 0x38, 0x55, 0xc5, 0xa9, 0x23, 0xd0, 0x9e, 0x08,
	0x23, 0xf2, 0x0c, 0x0e, 0x72, 0x5e, 0xb9, 0xdb, 0x95, 0xb0, 0x2b, 0xc3, 0xb4, 0x4c, 0x45, 0xb2,
	0xd6, 0x6a, 0xc5, 0x05, 0xab, 0xa2, 0x21, 0x5e, 0xc0, 0x87, 0x81, 0xf0, 0x32, 0xc4, 0xcf, 0x42,
	0x98, 0x7c, 0x0a, 0x23, 0x9f, 0x50, 0x7d, 0x27, 0x47, 0x78, 0x27, 0x87, 0x08, 0xfe, 0xea, 0xb1,
	0xf8, 0x04, 0xba, 0xbe, 0xf6, 0x1e, 0x01, 0x60, 0xc9, 0x62, 0x1f, 0x08, 0x2d, 0xa0, 0x8f, 0x88,
	0x6b, 0x00, 0xee, 0xee, 0xaf, 0xad, 0x70, 0x75, 0x29, 0x78, 0xe6, 0x7b, 0x53, 0x9f, 0x82, 0x83,
	0xce, 0x10, 0x89, 0x0f, 0xa1, 0x83, 0x3b, 0x49, 0xa0, 0x83, 0x9b, 0xef, 0x1c, 0x46, 0x14, 0xc7,
	0xf1, 0x5f, 0x2d, 0xd8, 0xff, 0xbf, 0xdd, 0x73, 0xae, 0x9a, 0xbd, 0xb5, 0xac, 0x32, 0x49, 0xb6,
	0xb6, 0x61, 0x55, 0x08, 0xd0, 0xf3, 0xb5, 0x25, 0x8f, 0xe1, 0x7e, 0x4d, 0x28, 0x59, 0xa9, 0x74,
	0xbd, 0xf2, 0x28, 0xa0, 0x3f, 0x20, 0xe8, 0xf6, 0x5e, 0xf0, 0x92, 0x7b, 0x17, 0xdf, 0x7a, 0x7a,
	0x08, 0x38, 0x8f, 0x4f, 0x60, 0xe8, 0x83, 0xc1, 0xa1, 0x83, 0xf1, 0x01, 0x62, 0x5e, 0x1f, 0xef,
	0xc1, 0x78, 0xde, 0x6c, 0x09, 0x2d, 0xd8, 0xb9, 0xd5, 0x8b, 0x9c, 0x97, 0xd1, 0xb6, 0x32, 0x49,
	0xae, 0xca, 0x94, 0xcb, 0x90, 0xf1, 0x00, 0xb1, 0x17, 0x08, 0x91, 0xcf, 0x60, 0xec, 0x29, 0xa9,
	0xcc, 0xde, 0x28, 0x5d, 0x25, 0x6b, 0x56, 0x86, 0xac, 0x77, 0x30, 0x30, 0xf7, 0xf8, 0x19, 0x2b,
	0xc9, 0x2b, 0x18, 0xf3, 0xaa, 0xb2, 0xa9, 0xcc, 0x58, 0x22, 0xf8, 0x8a, 0x19, 0x5e, 0xb2, 0x70,
	0xeb, 0x0f, 0xa6, 0xfe, 0x81, 0x99, 0xd6, 0x0f, 0xcc, 0xf4, 0x45, 0x78, 0x60, 0xe8, 0x6e, 0xad,
	0x39, 0x0d, 0x12, 0xf2, 0x1a, 0xf6, 0x33, 0xa1, 0xb2, 0xf3, 0xa4, 0x3a, 0x67, 0x97, 0x49, 0x2a,
	0x84, 0xba, 0x74, 0xf1, 0xd0, 0x62, 0x37, 0x58, 0x11, 0x94, 0xfd, 0x7c, 0xce, 0x2e, 0xe7, 0xb5,
	0x28, 0x9e, 0x40, 0xaf, 0xae, 0x44, 0xb2, 0x0f, 0x5d, 0x5f, 0xb3, 0xfe, 0x47, 0xfd, 0x24, 0xfe,
	0xbb, 0x05, 0xdb, 0xe1, 0x55, 0x71, 0xe7, 0x6d, 0x5d, 0xc5, 0x7b, 0x02, 0x8e, 0xf1, 0xa1, 0x10,
	0xfc, 0xba, 0xee, 0x42, 0xb1, 0x64, 0x82, 0x87, 0xaa, 0x23, 0xc7, 0xd0, 0x5d, 0x89, 0xb4, 0xa8,
	0xa2, 0x36, 0x76, 0x95, 0x47, 0x77, 0xbd, 0x59, 0xd3, 0x57, 0x22, 0x2d, 0xa8, 0xe7, 0x1e, 0x7e,
	0x09, 0x1d, 0x37, 0x75, 0x2b, 0xde, 0xa8, 0x51, 0x1c, 0xbb, 0x3c, 0x2f, 0x52, 0x61, 0x59, 0x58,
	0xcb, 0x4f, 0xbe, 0x3f, 0xfe, 0xed, 0xab, 0x82, 0x9b, 0x37, 0x76, 0x39, 0xcd, 0x54, 0x39, 0x0b,
	0x6b, 0xd4, 0xdf, 0xa3, 0x59, 0x68, 0xa0, 0x82, 0xe9, 0x59, 0xc1, 0x64, 0x78, 0xef, 0x97, 0x5b,
	0xb8, 0x4b, 0xc7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x59, 0xbd, 0x38, 0x07, 0x08, 0x00,
	0x00,
}
