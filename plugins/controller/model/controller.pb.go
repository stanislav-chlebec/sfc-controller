// Code generated by protoc-gen-gogo.
// source: controller.proto
// DO NOT EDIT!

/*
Package controller is a generated protocol buffer package.

It is generated from these files:
	controller.proto

It has these top-level messages:
	L3VRFRoute
	L3ArpEntry
	L2FIBEntry
	BDParms
	L2BD
	IPAMPoolStatus
	IPAMPoolSpec
	IPAMPool
	SystemParameters
	NetworkPodToNodeMap
	RenderedVppAgentEntry
	InterfaceStatus
	InterfaceSpec
	Interface
	NetworkPodSpec
	NetworkPod
	Connection
	NetworkServiceStatus
	NetworkServiceSpec
	NetworkService
	NetworkNodeOverlayStatus
	NetworkNodeOverlaySpec
	NetworkNodeOverlay
	NetworkNodeSpec
	NetworkNodeStatus
	NetworkNode
	MetaDataType
*/
package controller

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type L3VRFRoute struct {
	VrfId             uint32 `protobuf:"varint,1,opt,name=vrf_id,proto3" json:"vrf_id,omitempty"`
	Description       string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DstIpAddr         string `protobuf:"bytes,3,opt,name=dst_ip_addr,proto3" json:"dst_ip_addr,omitempty"`
	NextHopAddr       string `protobuf:"bytes,4,opt,name=next_hop_addr,proto3" json:"next_hop_addr,omitempty"`
	OutgoingInterface string `protobuf:"bytes,5,opt,name=outgoing_interface,proto3" json:"outgoing_interface,omitempty"`
	Weight            uint32 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Preference        uint32 `protobuf:"varint,7,opt,name=preference,proto3" json:"preference,omitempty"`
}

func (m *L3VRFRoute) Reset()         { *m = L3VRFRoute{} }
func (m *L3VRFRoute) String() string { return proto.CompactTextString(m) }
func (*L3VRFRoute) ProtoMessage()    {}

type L3ArpEntry struct {
	IpAddress   string `protobuf:"bytes,1,opt,name=ip_address,proto3" json:"ip_address,omitempty"`
	PhysAddress string `protobuf:"bytes,2,opt,name=phys_address,proto3" json:"phys_address,omitempty"`
}

func (m *L3ArpEntry) Reset()         { *m = L3ArpEntry{} }
func (m *L3ArpEntry) String() string { return proto.CompactTextString(m) }
func (*L3ArpEntry) ProtoMessage()    {}

type L2FIBEntry struct {
	DestMacAddress string `protobuf:"bytes,1,opt,name=dest_mac_address,proto3" json:"dest_mac_address,omitempty"`
	BdName         string `protobuf:"bytes,2,opt,name=bd_name,proto3" json:"bd_name,omitempty"`
	OutgoingIf     string `protobuf:"bytes,3,opt,name=outgoing_if,proto3" json:"outgoing_if,omitempty"`
	Action         string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	Bvi            bool   `protobuf:"varint,5,opt,name=bvi,proto3" json:"bvi,omitempty"`
}

func (m *L2FIBEntry) Reset()         { *m = L2FIBEntry{} }
func (m *L2FIBEntry) String() string { return proto.CompactTextString(m) }
func (*L2FIBEntry) ProtoMessage()    {}

type BDParms struct {
	Name                string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Flood               bool   `protobuf:"varint,2,opt,name=flood,proto3" json:"flood,omitempty"`
	UnknownUnicastFlood bool   `protobuf:"varint,3,opt,name=unknown_unicast_flood,proto3" json:"unknown_unicast_flood,omitempty"`
	Forward             bool   `protobuf:"varint,4,opt,name=forward,proto3" json:"forward,omitempty"`
	Learn               bool   `protobuf:"varint,5,opt,name=learn,proto3" json:"learn,omitempty"`
	ArpTermination      bool   `protobuf:"varint,6,opt,name=arp_termination,proto3" json:"arp_termination,omitempty"`
	MacAgeMinutes       uint32 `protobuf:"varint,7,opt,name=mac_age_minutes,proto3" json:"mac_age_minutes,omitempty"`
}

func (m *BDParms) Reset()         { *m = BDParms{} }
func (m *BDParms) String() string { return proto.CompactTextString(m) }
func (*BDParms) ProtoMessage()    {}

type L2BD struct {
	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BdParms      *BDParms `protobuf:"bytes,2,opt,name=bd_parms" json:"bd_parms,omitempty"`
	L2BdTemplate string   `protobuf:"bytes,3,opt,name=l2bd_template,proto3" json:"l2bd_template,omitempty"`
}

func (m *L2BD) Reset()         { *m = L2BD{} }
func (m *L2BD) String() string { return proto.CompactTextString(m) }
func (*L2BD) ProtoMessage()    {}

func (m *L2BD) GetBdParms() *BDParms {
	if m != nil {
		return m.BdParms
	}
	return nil
}

// IPAM will choose an address using the prefix.  In the example, there is an 8 bit address space [0..255]
// if 0..255 are not desired then a start and end address can be used.  See comments below.
type IPAMPoolStatus struct {
	Addresses string   `protobuf:"bytes,1,opt,name=addresses,proto3" json:"addresses,omitempty"`
	Status    string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Msg       []string `protobuf:"bytes,3,rep,name=msg" json:"msg,omitempty"`
}

func (m *IPAMPoolStatus) Reset()         { *m = IPAMPoolStatus{} }
func (m *IPAMPoolStatus) String() string { return proto.CompactTextString(m) }
func (*IPAMPoolStatus) ProtoMessage()    {}

type IPAMPoolSpec struct {
	Scope      string `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	Network    string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	StartRange uint32 `protobuf:"varint,4,opt,name=start_range,proto3" json:"start_range,omitempty"`
	EndRange   uint32 `protobuf:"varint,5,opt,name=end_range,proto3" json:"end_range,omitempty"`
}

func (m *IPAMPoolSpec) Reset()         { *m = IPAMPoolSpec{} }
func (m *IPAMPoolSpec) String() string { return proto.CompactTextString(m) }
func (*IPAMPoolSpec) ProtoMessage()    {}

type IPAMPool struct {
	Metadata *MetaDataType   `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *IPAMPoolSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *IPAMPoolStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *IPAMPool) Reset()         { *m = IPAMPool{} }
func (m *IPAMPool) String() string { return proto.CompactTextString(m) }
func (*IPAMPool) ProtoMessage()    {}

func (m *IPAMPool) GetMetadata() *MetaDataType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *IPAMPool) GetSpec() *IPAMPoolSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *IPAMPool) GetStatus() *IPAMPoolStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type SystemParameters struct {
	Mtu                          uint32     `protobuf:"varint,1,opt,name=mtu,proto3" json:"mtu,omitempty"`
	DefaultStaticRouteWeight     uint32     `protobuf:"varint,3,opt,name=default_static_route_weight,proto3" json:"default_static_route_weight,omitempty"`
	DefaultStaticRoutePreference uint32     `protobuf:"varint,4,opt,name=default_static_route_preference,proto3" json:"default_static_route_preference,omitempty"`
	L2BdTemplates                []*BDParms `protobuf:"bytes,5,rep,name=l2bd_templates" json:"l2bd_templates,omitempty"`
	RxMode                       string     `protobuf:"bytes,7,opt,name=rx_mode,proto3" json:"rx_mode,omitempty"`
}

func (m *SystemParameters) Reset()         { *m = SystemParameters{} }
func (m *SystemParameters) String() string { return proto.CompactTextString(m) }
func (*SystemParameters) ProtoMessage()    {}

func (m *SystemParameters) GetL2BdTemplates() []*BDParms {
	if m != nil {
		return m.L2BdTemplates
	}
	return nil
}

type NetworkPodToNodeMap struct {
	Pod  string `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	Node string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *NetworkPodToNodeMap) Reset()         { *m = NetworkPodToNodeMap{} }
func (m *NetworkPodToNodeMap) String() string { return proto.CompactTextString(m) }
func (*NetworkPodToNodeMap) ProtoMessage()    {}

// The rendered vpp agent entry is identified by the key, and type, which are used
// to read the entry from etcd
type RenderedVppAgentEntry struct {
	VppAgentKey  string `protobuf:"bytes,1,opt,name=vpp_agent_key,proto3" json:"vpp_agent_key,omitempty"`
	VppAgentType string `protobuf:"bytes,3,opt,name=vpp_agent_type,proto3" json:"vpp_agent_type,omitempty"`
}

func (m *RenderedVppAgentEntry) Reset()         { *m = RenderedVppAgentEntry{} }
func (m *RenderedVppAgentEntry) String() string { return proto.CompactTextString(m) }
func (*RenderedVppAgentEntry) ProtoMessage()    {}

//
// Interface ... channel?
//
type InterfaceStatus struct {
	PodInterfaceName string   `protobuf:"bytes,1,opt,name=pod_interface_name,proto3" json:"pod_interface_name,omitempty"`
	Status           string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Msg              []string `protobuf:"bytes,4,rep,name=msg" json:"msg,omitempty"`
	MacAddress       string   `protobuf:"bytes,5,opt,name=mac_address,proto3" json:"mac_address,omitempty"`
	IpAddresses      []string `protobuf:"bytes,7,rep,name=ip_addresses" json:"ip_addresses,omitempty"`
	MemifID          uint32   `protobuf:"varint,8,opt,name=memifID,proto3" json:"memifID,omitempty"`
	VrfID            uint32   `protobuf:"varint,9,opt,name=vrfID,proto3" json:"vrfID,omitempty"`
	Node             string   `protobuf:"bytes,10,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *InterfaceStatus) Reset()         { *m = InterfaceStatus{} }
func (m *InterfaceStatus) String() string { return proto.CompactTextString(m) }
func (*InterfaceStatus) ProtoMessage()    {}

type InterfaceSpec struct {
	IfType     string `protobuf:"bytes,2,opt,name=if_type,proto3" json:"if_type,omitempty"`
	MacAddress string `protobuf:"bytes,3,opt,name=mac_address,proto3" json:"mac_address,omitempty"`
	// 02:00:00:00:00:00 as base, last octet increments
	Mtu          uint32                    `protobuf:"varint,4,opt,name=mtu,proto3" json:"mtu,omitempty"`
	RxMode       string                    `protobuf:"bytes,5,opt,name=rx_mode,proto3" json:"rx_mode,omitempty"`
	IpAddresses  []string                  `protobuf:"bytes,6,rep,name=ip_addresses" json:"ip_addresses,omitempty"`
	VrfId        uint32                    `protobuf:"varint,7,opt,name=vrf_id,proto3" json:"vrf_id,omitempty"`
	IpamPoolName string                    `protobuf:"bytes,8,opt,name=ipam_pool_name,proto3" json:"ipam_pool_name,omitempty"`
	AdminStatus  string                    `protobuf:"bytes,9,opt,name=admin_status,proto3" json:"admin_status,omitempty"`
	MemifParms   *InterfaceSpec_MemIFParms `protobuf:"bytes,11,opt,name=memif_parms" json:"memif_parms,omitempty"`
}

func (m *InterfaceSpec) Reset()         { *m = InterfaceSpec{} }
func (m *InterfaceSpec) String() string { return proto.CompactTextString(m) }
func (*InterfaceSpec) ProtoMessage()    {}

func (m *InterfaceSpec) GetMemifParms() *InterfaceSpec_MemIFParms {
	if m != nil {
		return m.MemifParms
	}
	return nil
}

type InterfaceSpec_MemIFParms struct {
	Mode         string `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	InterPodConn string `protobuf:"bytes,2,opt,name=inter_pod_conn,proto3" json:"inter_pod_conn,omitempty"`
}

func (m *InterfaceSpec_MemIFParms) Reset()         { *m = InterfaceSpec_MemIFParms{} }
func (m *InterfaceSpec_MemIFParms) String() string { return proto.CompactTextString(m) }
func (*InterfaceSpec_MemIFParms) ProtoMessage()    {}

type Interface struct {
	Metadata *MetaDataType  `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *InterfaceSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
}

func (m *Interface) Reset()         { *m = Interface{} }
func (m *Interface) String() string { return proto.CompactTextString(m) }
func (*Interface) ProtoMessage()    {}

func (m *Interface) GetMetadata() *MetaDataType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Interface) GetSpec() *InterfaceSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type NetworkPodSpec struct {
	PodType    string       `protobuf:"bytes,2,opt,name=pod_type,proto3" json:"pod_type,omitempty"`
	Interfaces []*Interface `protobuf:"bytes,3,rep,name=interfaces" json:"interfaces,omitempty"`
	L2Bds      []*L2BD      `protobuf:"bytes,4,rep,name=l2bds" json:"l2bds,omitempty"`
}

func (m *NetworkPodSpec) Reset()         { *m = NetworkPodSpec{} }
func (m *NetworkPodSpec) String() string { return proto.CompactTextString(m) }
func (*NetworkPodSpec) ProtoMessage()    {}

func (m *NetworkPodSpec) GetInterfaces() []*Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *NetworkPodSpec) GetL2Bds() []*L2BD {
	if m != nil {
		return m.L2Bds
	}
	return nil
}

type NetworkPod struct {
	Metadata *MetaDataType   `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *NetworkPodSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
}

func (m *NetworkPod) Reset()         { *m = NetworkPod{} }
func (m *NetworkPod) String() string { return proto.CompactTextString(m) }
func (*NetworkPod) ProtoMessage()    {}

func (m *NetworkPod) GetMetadata() *MetaDataType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkPod) GetSpec() *NetworkPodSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type Connection struct {
	ConnType               string   `protobuf:"bytes,1,opt,name=conn_type,proto3" json:"conn_type,omitempty"`
	NetworkNodeOverlayName string   `protobuf:"bytes,2,opt,name=network_node_overlay_name,proto3" json:"network_node_overlay_name,omitempty"`
	PodInterfaces          []string `protobuf:"bytes,3,rep,name=pod_interfaces" json:"pod_interfaces,omitempty"`
	UseNodeL2Bd            string   `protobuf:"bytes,4,opt,name=use_node_l2bd,proto3" json:"use_node_l2bd,omitempty"`
	// only for l2mp connections
	L2Bd *L2BD `protobuf:"bytes,5,opt,name=l2bd" json:"l2bd,omitempty"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}

func (m *Connection) GetL2Bd() *L2BD {
	if m != nil {
		return m.L2Bd
	}
	return nil
}

type NetworkServiceStatus struct {
	OperStatus              string                   `protobuf:"bytes,2,opt,name=oper_status,proto3" json:"oper_status,omitempty"`
	Msg                     []string                 `protobuf:"bytes,3,rep,name=msg" json:"msg,omitempty"`
	RenderedVppAgentEntries []*RenderedVppAgentEntry `protobuf:"bytes,4,rep,name=rendered_vpp_agent_entries" json:"rendered_vpp_agent_entries,omitempty"`
	Interfaces              []*InterfaceStatus       `protobuf:"bytes,5,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *NetworkServiceStatus) Reset()         { *m = NetworkServiceStatus{} }
func (m *NetworkServiceStatus) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceStatus) ProtoMessage()    {}

func (m *NetworkServiceStatus) GetRenderedVppAgentEntries() []*RenderedVppAgentEntry {
	if m != nil {
		return m.RenderedVppAgentEntries
	}
	return nil
}

func (m *NetworkServiceStatus) GetInterfaces() []*InterfaceStatus {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type NetworkServiceSpec struct {
	NetworkPods []*NetworkPod `protobuf:"bytes,3,rep,name=network_pods" json:"network_pods,omitempty"`
	Connections []*Connection `protobuf:"bytes,4,rep,name=connections" json:"connections,omitempty"`
}

func (m *NetworkServiceSpec) Reset()         { *m = NetworkServiceSpec{} }
func (m *NetworkServiceSpec) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceSpec) ProtoMessage()    {}

func (m *NetworkServiceSpec) GetNetworkPods() []*NetworkPod {
	if m != nil {
		return m.NetworkPods
	}
	return nil
}

func (m *NetworkServiceSpec) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

type NetworkService struct {
	Metadata *MetaDataType         `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *NetworkServiceSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *NetworkServiceStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *NetworkService) Reset()         { *m = NetworkService{} }
func (m *NetworkService) String() string { return proto.CompactTextString(m) }
func (*NetworkService) ProtoMessage()    {}

func (m *NetworkService) GetMetadata() *MetaDataType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkService) GetSpec() *NetworkServiceSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *NetworkService) GetStatus() *NetworkServiceStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type NetworkNodeOverlayStatus struct {
	Status                  string                   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                     []string                 `protobuf:"bytes,2,rep,name=msg" json:"msg,omitempty"`
	RenderedVppAgentEntries []*RenderedVppAgentEntry `protobuf:"bytes,3,rep,name=rendered_vpp_agent_entries" json:"rendered_vpp_agent_entries,omitempty"`
}

func (m *NetworkNodeOverlayStatus) Reset()         { *m = NetworkNodeOverlayStatus{} }
func (m *NetworkNodeOverlayStatus) String() string { return proto.CompactTextString(m) }
func (*NetworkNodeOverlayStatus) ProtoMessage()    {}

func (m *NetworkNodeOverlayStatus) GetRenderedVppAgentEntries() []*RenderedVppAgentEntry {
	if m != nil {
		return m.RenderedVppAgentEntries
	}
	return nil
}

type NetworkNodeOverlaySpec struct {
	ServiceMeshType       string                                        `protobuf:"bytes,2,opt,name=service_mesh_type,proto3" json:"service_mesh_type,omitempty"`
	ConnectionType        string                                        `protobuf:"bytes,3,opt,name=connection_type,proto3" json:"connection_type,omitempty"`
	VxlanHubAndSpokeParms *NetworkNodeOverlaySpec_VxlanHubAndSpokeParms `protobuf:"bytes,8,opt,name=vxlan_hub_and_spoke_parms" json:"vxlan_hub_and_spoke_parms,omitempty"`
	VxlanMeshParms        *NetworkNodeOverlaySpec_VxlanMeshParms        `protobuf:"bytes,9,opt,name=vxlan_mesh_parms" json:"vxlan_mesh_parms,omitempty"`
}

func (m *NetworkNodeOverlaySpec) Reset()         { *m = NetworkNodeOverlaySpec{} }
func (m *NetworkNodeOverlaySpec) String() string { return proto.CompactTextString(m) }
func (*NetworkNodeOverlaySpec) ProtoMessage()    {}

func (m *NetworkNodeOverlaySpec) GetVxlanHubAndSpokeParms() *NetworkNodeOverlaySpec_VxlanHubAndSpokeParms {
	if m != nil {
		return m.VxlanHubAndSpokeParms
	}
	return nil
}

func (m *NetworkNodeOverlaySpec) GetVxlanMeshParms() *NetworkNodeOverlaySpec_VxlanMeshParms {
	if m != nil {
		return m.VxlanMeshParms
	}
	return nil
}

type NetworkNodeOverlaySpec_VxlanHubAndSpokeParms struct {
	HubNodeName               string `protobuf:"bytes,1,opt,name=hub_node_name,proto3" json:"hub_node_name,omitempty"`
	Vni                       uint32 `protobuf:"varint,2,opt,name=vni,proto3" json:"vni,omitempty"`
	LoopbackIpamPoolName      string `protobuf:"bytes,3,opt,name=loopback_ipam_pool_name,proto3" json:"loopback_ipam_pool_name,omitempty"`
	NetworkNodeInterfaceLabel string `protobuf:"bytes,4,opt,name=network_node_interface_label,proto3" json:"network_node_interface_label,omitempty"`
}

func (m *NetworkNodeOverlaySpec_VxlanHubAndSpokeParms) Reset() {
	*m = NetworkNodeOverlaySpec_VxlanHubAndSpokeParms{}
}
func (m *NetworkNodeOverlaySpec_VxlanHubAndSpokeParms) String() string {
	return proto.CompactTextString(m)
}
func (*NetworkNodeOverlaySpec_VxlanHubAndSpokeParms) ProtoMessage() {}

type NetworkNodeOverlaySpec_VxlanMeshParms struct {
	VniRangeStart             uint32 `protobuf:"varint,1,opt,name=vni_range_start,proto3" json:"vni_range_start,omitempty"`
	VniRangeEnd               uint32 `protobuf:"varint,2,opt,name=vni_range_end,proto3" json:"vni_range_end,omitempty"`
	LoopbackIpamPoolName      string `protobuf:"bytes,3,opt,name=loopback_ipam_pool_name,proto3" json:"loopback_ipam_pool_name,omitempty"`
	NetworkNodeInterfaceLabel string `protobuf:"bytes,4,opt,name=network_node_interface_label,proto3" json:"network_node_interface_label,omitempty"`
}

func (m *NetworkNodeOverlaySpec_VxlanMeshParms) Reset()         { *m = NetworkNodeOverlaySpec_VxlanMeshParms{} }
func (m *NetworkNodeOverlaySpec_VxlanMeshParms) String() string { return proto.CompactTextString(m) }
func (*NetworkNodeOverlaySpec_VxlanMeshParms) ProtoMessage()    {}

type NetworkNodeOverlay struct {
	Metadata *MetaDataType             `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *NetworkNodeOverlaySpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *NetworkNodeOverlayStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *NetworkNodeOverlay) Reset()         { *m = NetworkNodeOverlay{} }
func (m *NetworkNodeOverlay) String() string { return proto.CompactTextString(m) }
func (*NetworkNodeOverlay) ProtoMessage()    {}

func (m *NetworkNodeOverlay) GetMetadata() *MetaDataType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkNodeOverlay) GetSpec() *NetworkNodeOverlaySpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *NetworkNodeOverlay) GetStatus() *NetworkNodeOverlayStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// NetworkNodeSpec extends the k8s node
type NetworkNodeSpec struct {
	K8SNodeName string       `protobuf:"bytes,1,opt,name=k8s_node_name,proto3" json:"k8s_node_name,omitempty"`
	NodeType    string       `protobuf:"bytes,2,opt,name=node_type,proto3" json:"node_type,omitempty"`
	Interfaces  []*Interface `protobuf:"bytes,4,rep,name=interfaces" json:"interfaces,omitempty"`
	L2Bds       []*L2BD      `protobuf:"bytes,5,rep,name=l2bds" json:"l2bds,omitempty"`
}

func (m *NetworkNodeSpec) Reset()         { *m = NetworkNodeSpec{} }
func (m *NetworkNodeSpec) String() string { return proto.CompactTextString(m) }
func (*NetworkNodeSpec) ProtoMessage()    {}

func (m *NetworkNodeSpec) GetInterfaces() []*Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *NetworkNodeSpec) GetL2Bds() []*L2BD {
	if m != nil {
		return m.L2Bds
	}
	return nil
}

type NetworkNodeStatus struct {
	Status                  string                   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                     []string                 `protobuf:"bytes,2,rep,name=msg" json:"msg,omitempty"`
	RenderedVppAgentEntries []*RenderedVppAgentEntry `protobuf:"bytes,3,rep,name=rendered_vpp_agent_entries" json:"rendered_vpp_agent_entries,omitempty"`
	// map<string,RenderedVppAgentEntry> rendered_vpp_agent_entries = 3; // key: modeltype/name
	Interfaces []*InterfaceStatus `protobuf:"bytes,5,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *NetworkNodeStatus) Reset()         { *m = NetworkNodeStatus{} }
func (m *NetworkNodeStatus) String() string { return proto.CompactTextString(m) }
func (*NetworkNodeStatus) ProtoMessage()    {}

func (m *NetworkNodeStatus) GetRenderedVppAgentEntries() []*RenderedVppAgentEntry {
	if m != nil {
		return m.RenderedVppAgentEntries
	}
	return nil
}

func (m *NetworkNodeStatus) GetInterfaces() []*InterfaceStatus {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type NetworkNode struct {
	Metadata *MetaDataType      `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *NetworkNodeSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *NetworkNodeStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *NetworkNode) Reset()         { *m = NetworkNode{} }
func (m *NetworkNode) String() string { return proto.CompactTextString(m) }
func (*NetworkNode) ProtoMessage()    {}

func (m *NetworkNode) GetMetadata() *MetaDataType {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkNode) GetSpec() *NetworkNodeSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *NetworkNode) GetStatus() *NetworkNodeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// MetaDataType generic parms for all controller high level objects
type MetaDataType struct {
	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels []string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
}

func (m *MetaDataType) Reset()         { *m = MetaDataType{} }
func (m *MetaDataType) String() string { return proto.CompactTextString(m) }
func (*MetaDataType) ProtoMessage()    {}
