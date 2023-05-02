package virtualnetworks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DdosSettingsProtectionMode string

const (
	DdosSettingsProtectionModeDisabled                DdosSettingsProtectionMode = "Disabled"
	DdosSettingsProtectionModeEnabled                 DdosSettingsProtectionMode = "Enabled"
	DdosSettingsProtectionModeVirtualNetworkInherited DdosSettingsProtectionMode = "VirtualNetworkInherited"
)

func PossibleValuesForDdosSettingsProtectionMode() []string {
	return []string{
		string(DdosSettingsProtectionModeDisabled),
		string(DdosSettingsProtectionModeEnabled),
		string(DdosSettingsProtectionModeVirtualNetworkInherited),
	}
}

func (s *DdosSettingsProtectionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDdosSettingsProtectionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDdosSettingsProtectionMode(input string) (*DdosSettingsProtectionMode, error) {
	vals := map[string]DdosSettingsProtectionMode{
		"disabled":                DdosSettingsProtectionModeDisabled,
		"enabled":                 DdosSettingsProtectionModeEnabled,
		"virtualnetworkinherited": DdosSettingsProtectionModeVirtualNetworkInherited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DdosSettingsProtectionMode(input)
	return &out, nil
}

type DeleteOptions string

const (
	DeleteOptionsDelete DeleteOptions = "Delete"
	DeleteOptionsDetach DeleteOptions = "Detach"
)

func PossibleValuesForDeleteOptions() []string {
	return []string{
		string(DeleteOptionsDelete),
		string(DeleteOptionsDetach),
	}
}

func (s *DeleteOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeleteOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeleteOptions(input string) (*DeleteOptions, error) {
	vals := map[string]DeleteOptions{
		"delete": DeleteOptionsDelete,
		"detach": DeleteOptionsDetach,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeleteOptions(input)
	return &out, nil
}

type FlowLogFormatType string

const (
	FlowLogFormatTypeJSON FlowLogFormatType = "JSON"
)

func PossibleValuesForFlowLogFormatType() []string {
	return []string{
		string(FlowLogFormatTypeJSON),
	}
}

func (s *FlowLogFormatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFlowLogFormatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFlowLogFormatType(input string) (*FlowLogFormatType, error) {
	vals := map[string]FlowLogFormatType{
		"json": FlowLogFormatTypeJSON,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FlowLogFormatType(input)
	return &out, nil
}

type GatewayLoadBalancerTunnelInterfaceType string

const (
	GatewayLoadBalancerTunnelInterfaceTypeExternal GatewayLoadBalancerTunnelInterfaceType = "External"
	GatewayLoadBalancerTunnelInterfaceTypeInternal GatewayLoadBalancerTunnelInterfaceType = "Internal"
	GatewayLoadBalancerTunnelInterfaceTypeNone     GatewayLoadBalancerTunnelInterfaceType = "None"
)

func PossibleValuesForGatewayLoadBalancerTunnelInterfaceType() []string {
	return []string{
		string(GatewayLoadBalancerTunnelInterfaceTypeExternal),
		string(GatewayLoadBalancerTunnelInterfaceTypeInternal),
		string(GatewayLoadBalancerTunnelInterfaceTypeNone),
	}
}

func (s *GatewayLoadBalancerTunnelInterfaceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGatewayLoadBalancerTunnelInterfaceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGatewayLoadBalancerTunnelInterfaceType(input string) (*GatewayLoadBalancerTunnelInterfaceType, error) {
	vals := map[string]GatewayLoadBalancerTunnelInterfaceType{
		"external": GatewayLoadBalancerTunnelInterfaceTypeExternal,
		"internal": GatewayLoadBalancerTunnelInterfaceTypeInternal,
		"none":     GatewayLoadBalancerTunnelInterfaceTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GatewayLoadBalancerTunnelInterfaceType(input)
	return &out, nil
}

type GatewayLoadBalancerTunnelProtocol string

const (
	GatewayLoadBalancerTunnelProtocolNative GatewayLoadBalancerTunnelProtocol = "Native"
	GatewayLoadBalancerTunnelProtocolNone   GatewayLoadBalancerTunnelProtocol = "None"
	GatewayLoadBalancerTunnelProtocolVXLAN  GatewayLoadBalancerTunnelProtocol = "VXLAN"
)

func PossibleValuesForGatewayLoadBalancerTunnelProtocol() []string {
	return []string{
		string(GatewayLoadBalancerTunnelProtocolNative),
		string(GatewayLoadBalancerTunnelProtocolNone),
		string(GatewayLoadBalancerTunnelProtocolVXLAN),
	}
}

func (s *GatewayLoadBalancerTunnelProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGatewayLoadBalancerTunnelProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGatewayLoadBalancerTunnelProtocol(input string) (*GatewayLoadBalancerTunnelProtocol, error) {
	vals := map[string]GatewayLoadBalancerTunnelProtocol{
		"native": GatewayLoadBalancerTunnelProtocolNative,
		"none":   GatewayLoadBalancerTunnelProtocolNone,
		"vxlan":  GatewayLoadBalancerTunnelProtocolVXLAN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GatewayLoadBalancerTunnelProtocol(input)
	return &out, nil
}

type IPAllocationMethod string

const (
	IPAllocationMethodDynamic IPAllocationMethod = "Dynamic"
	IPAllocationMethodStatic  IPAllocationMethod = "Static"
)

func PossibleValuesForIPAllocationMethod() []string {
	return []string{
		string(IPAllocationMethodDynamic),
		string(IPAllocationMethodStatic),
	}
}

func (s *IPAllocationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPAllocationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPAllocationMethod(input string) (*IPAllocationMethod, error) {
	vals := map[string]IPAllocationMethod{
		"dynamic": IPAllocationMethodDynamic,
		"static":  IPAllocationMethodStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPAllocationMethod(input)
	return &out, nil
}

type IPVersion string

const (
	IPVersionIPvFour IPVersion = "IPv4"
	IPVersionIPvSix  IPVersion = "IPv6"
)

func PossibleValuesForIPVersion() []string {
	return []string{
		string(IPVersionIPvFour),
		string(IPVersionIPvSix),
	}
}

func (s *IPVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIPVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIPVersion(input string) (*IPVersion, error) {
	vals := map[string]IPVersion{
		"ipv4": IPVersionIPvFour,
		"ipv6": IPVersionIPvSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPVersion(input)
	return &out, nil
}

type IsWorkloadProtected string

const (
	IsWorkloadProtectedFalse IsWorkloadProtected = "False"
	IsWorkloadProtectedTrue  IsWorkloadProtected = "True"
)

func PossibleValuesForIsWorkloadProtected() []string {
	return []string{
		string(IsWorkloadProtectedFalse),
		string(IsWorkloadProtectedTrue),
	}
}

func (s *IsWorkloadProtected) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIsWorkloadProtected(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIsWorkloadProtected(input string) (*IsWorkloadProtected, error) {
	vals := map[string]IsWorkloadProtected{
		"false": IsWorkloadProtectedFalse,
		"true":  IsWorkloadProtectedTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IsWorkloadProtected(input)
	return &out, nil
}

type LoadBalancerBackendAddressAdminState string

const (
	LoadBalancerBackendAddressAdminStateDown  LoadBalancerBackendAddressAdminState = "Down"
	LoadBalancerBackendAddressAdminStateDrain LoadBalancerBackendAddressAdminState = "Drain"
	LoadBalancerBackendAddressAdminStateNone  LoadBalancerBackendAddressAdminState = "None"
	LoadBalancerBackendAddressAdminStateUp    LoadBalancerBackendAddressAdminState = "Up"
)

func PossibleValuesForLoadBalancerBackendAddressAdminState() []string {
	return []string{
		string(LoadBalancerBackendAddressAdminStateDown),
		string(LoadBalancerBackendAddressAdminStateDrain),
		string(LoadBalancerBackendAddressAdminStateNone),
		string(LoadBalancerBackendAddressAdminStateUp),
	}
}

func (s *LoadBalancerBackendAddressAdminState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLoadBalancerBackendAddressAdminState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLoadBalancerBackendAddressAdminState(input string) (*LoadBalancerBackendAddressAdminState, error) {
	vals := map[string]LoadBalancerBackendAddressAdminState{
		"down":  LoadBalancerBackendAddressAdminStateDown,
		"drain": LoadBalancerBackendAddressAdminStateDrain,
		"none":  LoadBalancerBackendAddressAdminStateNone,
		"up":    LoadBalancerBackendAddressAdminStateUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoadBalancerBackendAddressAdminState(input)
	return &out, nil
}

type NatGatewaySkuName string

const (
	NatGatewaySkuNameStandard NatGatewaySkuName = "Standard"
)

func PossibleValuesForNatGatewaySkuName() []string {
	return []string{
		string(NatGatewaySkuNameStandard),
	}
}

func (s *NatGatewaySkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNatGatewaySkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNatGatewaySkuName(input string) (*NatGatewaySkuName, error) {
	vals := map[string]NatGatewaySkuName{
		"standard": NatGatewaySkuNameStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NatGatewaySkuName(input)
	return &out, nil
}

type NetworkInterfaceAuxiliaryMode string

const (
	NetworkInterfaceAuxiliaryModeFloating       NetworkInterfaceAuxiliaryMode = "Floating"
	NetworkInterfaceAuxiliaryModeMaxConnections NetworkInterfaceAuxiliaryMode = "MaxConnections"
	NetworkInterfaceAuxiliaryModeNone           NetworkInterfaceAuxiliaryMode = "None"
)

func PossibleValuesForNetworkInterfaceAuxiliaryMode() []string {
	return []string{
		string(NetworkInterfaceAuxiliaryModeFloating),
		string(NetworkInterfaceAuxiliaryModeMaxConnections),
		string(NetworkInterfaceAuxiliaryModeNone),
	}
}

func (s *NetworkInterfaceAuxiliaryMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkInterfaceAuxiliaryMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkInterfaceAuxiliaryMode(input string) (*NetworkInterfaceAuxiliaryMode, error) {
	vals := map[string]NetworkInterfaceAuxiliaryMode{
		"floating":       NetworkInterfaceAuxiliaryModeFloating,
		"maxconnections": NetworkInterfaceAuxiliaryModeMaxConnections,
		"none":           NetworkInterfaceAuxiliaryModeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkInterfaceAuxiliaryMode(input)
	return &out, nil
}

type NetworkInterfaceMigrationPhase string

const (
	NetworkInterfaceMigrationPhaseAbort     NetworkInterfaceMigrationPhase = "Abort"
	NetworkInterfaceMigrationPhaseCommit    NetworkInterfaceMigrationPhase = "Commit"
	NetworkInterfaceMigrationPhaseCommitted NetworkInterfaceMigrationPhase = "Committed"
	NetworkInterfaceMigrationPhaseNone      NetworkInterfaceMigrationPhase = "None"
	NetworkInterfaceMigrationPhasePrepare   NetworkInterfaceMigrationPhase = "Prepare"
)

func PossibleValuesForNetworkInterfaceMigrationPhase() []string {
	return []string{
		string(NetworkInterfaceMigrationPhaseAbort),
		string(NetworkInterfaceMigrationPhaseCommit),
		string(NetworkInterfaceMigrationPhaseCommitted),
		string(NetworkInterfaceMigrationPhaseNone),
		string(NetworkInterfaceMigrationPhasePrepare),
	}
}

func (s *NetworkInterfaceMigrationPhase) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkInterfaceMigrationPhase(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkInterfaceMigrationPhase(input string) (*NetworkInterfaceMigrationPhase, error) {
	vals := map[string]NetworkInterfaceMigrationPhase{
		"abort":     NetworkInterfaceMigrationPhaseAbort,
		"commit":    NetworkInterfaceMigrationPhaseCommit,
		"committed": NetworkInterfaceMigrationPhaseCommitted,
		"none":      NetworkInterfaceMigrationPhaseNone,
		"prepare":   NetworkInterfaceMigrationPhasePrepare,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkInterfaceMigrationPhase(input)
	return &out, nil
}

type NetworkInterfaceNicType string

const (
	NetworkInterfaceNicTypeElastic  NetworkInterfaceNicType = "Elastic"
	NetworkInterfaceNicTypeStandard NetworkInterfaceNicType = "Standard"
)

func PossibleValuesForNetworkInterfaceNicType() []string {
	return []string{
		string(NetworkInterfaceNicTypeElastic),
		string(NetworkInterfaceNicTypeStandard),
	}
}

func (s *NetworkInterfaceNicType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkInterfaceNicType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkInterfaceNicType(input string) (*NetworkInterfaceNicType, error) {
	vals := map[string]NetworkInterfaceNicType{
		"elastic":  NetworkInterfaceNicTypeElastic,
		"standard": NetworkInterfaceNicTypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkInterfaceNicType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type PublicIPAddressMigrationPhase string

const (
	PublicIPAddressMigrationPhaseAbort     PublicIPAddressMigrationPhase = "Abort"
	PublicIPAddressMigrationPhaseCommit    PublicIPAddressMigrationPhase = "Commit"
	PublicIPAddressMigrationPhaseCommitted PublicIPAddressMigrationPhase = "Committed"
	PublicIPAddressMigrationPhaseNone      PublicIPAddressMigrationPhase = "None"
	PublicIPAddressMigrationPhasePrepare   PublicIPAddressMigrationPhase = "Prepare"
)

func PossibleValuesForPublicIPAddressMigrationPhase() []string {
	return []string{
		string(PublicIPAddressMigrationPhaseAbort),
		string(PublicIPAddressMigrationPhaseCommit),
		string(PublicIPAddressMigrationPhaseCommitted),
		string(PublicIPAddressMigrationPhaseNone),
		string(PublicIPAddressMigrationPhasePrepare),
	}
}

func (s *PublicIPAddressMigrationPhase) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePublicIPAddressMigrationPhase(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePublicIPAddressMigrationPhase(input string) (*PublicIPAddressMigrationPhase, error) {
	vals := map[string]PublicIPAddressMigrationPhase{
		"abort":     PublicIPAddressMigrationPhaseAbort,
		"commit":    PublicIPAddressMigrationPhaseCommit,
		"committed": PublicIPAddressMigrationPhaseCommitted,
		"none":      PublicIPAddressMigrationPhaseNone,
		"prepare":   PublicIPAddressMigrationPhasePrepare,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicIPAddressMigrationPhase(input)
	return &out, nil
}

type PublicIPAddressSkuName string

const (
	PublicIPAddressSkuNameBasic    PublicIPAddressSkuName = "Basic"
	PublicIPAddressSkuNameStandard PublicIPAddressSkuName = "Standard"
)

func PossibleValuesForPublicIPAddressSkuName() []string {
	return []string{
		string(PublicIPAddressSkuNameBasic),
		string(PublicIPAddressSkuNameStandard),
	}
}

func (s *PublicIPAddressSkuName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePublicIPAddressSkuName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePublicIPAddressSkuName(input string) (*PublicIPAddressSkuName, error) {
	vals := map[string]PublicIPAddressSkuName{
		"basic":    PublicIPAddressSkuNameBasic,
		"standard": PublicIPAddressSkuNameStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicIPAddressSkuName(input)
	return &out, nil
}

type PublicIPAddressSkuTier string

const (
	PublicIPAddressSkuTierGlobal   PublicIPAddressSkuTier = "Global"
	PublicIPAddressSkuTierRegional PublicIPAddressSkuTier = "Regional"
)

func PossibleValuesForPublicIPAddressSkuTier() []string {
	return []string{
		string(PublicIPAddressSkuTierGlobal),
		string(PublicIPAddressSkuTierRegional),
	}
}

func (s *PublicIPAddressSkuTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePublicIPAddressSkuTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePublicIPAddressSkuTier(input string) (*PublicIPAddressSkuTier, error) {
	vals := map[string]PublicIPAddressSkuTier{
		"global":   PublicIPAddressSkuTierGlobal,
		"regional": PublicIPAddressSkuTierRegional,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicIPAddressSkuTier(input)
	return &out, nil
}

type RouteNextHopType string

const (
	RouteNextHopTypeInternet              RouteNextHopType = "Internet"
	RouteNextHopTypeNone                  RouteNextHopType = "None"
	RouteNextHopTypeVirtualAppliance      RouteNextHopType = "VirtualAppliance"
	RouteNextHopTypeVirtualNetworkGateway RouteNextHopType = "VirtualNetworkGateway"
	RouteNextHopTypeVnetLocal             RouteNextHopType = "VnetLocal"
)

func PossibleValuesForRouteNextHopType() []string {
	return []string{
		string(RouteNextHopTypeInternet),
		string(RouteNextHopTypeNone),
		string(RouteNextHopTypeVirtualAppliance),
		string(RouteNextHopTypeVirtualNetworkGateway),
		string(RouteNextHopTypeVnetLocal),
	}
}

func (s *RouteNextHopType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRouteNextHopType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRouteNextHopType(input string) (*RouteNextHopType, error) {
	vals := map[string]RouteNextHopType{
		"internet":              RouteNextHopTypeInternet,
		"none":                  RouteNextHopTypeNone,
		"virtualappliance":      RouteNextHopTypeVirtualAppliance,
		"virtualnetworkgateway": RouteNextHopTypeVirtualNetworkGateway,
		"vnetlocal":             RouteNextHopTypeVnetLocal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RouteNextHopType(input)
	return &out, nil
}

type SecurityRuleAccess string

const (
	SecurityRuleAccessAllow SecurityRuleAccess = "Allow"
	SecurityRuleAccessDeny  SecurityRuleAccess = "Deny"
)

func PossibleValuesForSecurityRuleAccess() []string {
	return []string{
		string(SecurityRuleAccessAllow),
		string(SecurityRuleAccessDeny),
	}
}

func (s *SecurityRuleAccess) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRuleAccess(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRuleAccess(input string) (*SecurityRuleAccess, error) {
	vals := map[string]SecurityRuleAccess{
		"allow": SecurityRuleAccessAllow,
		"deny":  SecurityRuleAccessDeny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRuleAccess(input)
	return &out, nil
}

type SecurityRuleDirection string

const (
	SecurityRuleDirectionInbound  SecurityRuleDirection = "Inbound"
	SecurityRuleDirectionOutbound SecurityRuleDirection = "Outbound"
)

func PossibleValuesForSecurityRuleDirection() []string {
	return []string{
		string(SecurityRuleDirectionInbound),
		string(SecurityRuleDirectionOutbound),
	}
}

func (s *SecurityRuleDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRuleDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRuleDirection(input string) (*SecurityRuleDirection, error) {
	vals := map[string]SecurityRuleDirection{
		"inbound":  SecurityRuleDirectionInbound,
		"outbound": SecurityRuleDirectionOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRuleDirection(input)
	return &out, nil
}

type SecurityRuleProtocol string

const (
	SecurityRuleProtocolAh   SecurityRuleProtocol = "Ah"
	SecurityRuleProtocolAny  SecurityRuleProtocol = "*"
	SecurityRuleProtocolEsp  SecurityRuleProtocol = "Esp"
	SecurityRuleProtocolIcmp SecurityRuleProtocol = "Icmp"
	SecurityRuleProtocolTcp  SecurityRuleProtocol = "Tcp"
	SecurityRuleProtocolUdp  SecurityRuleProtocol = "Udp"
)

func PossibleValuesForSecurityRuleProtocol() []string {
	return []string{
		string(SecurityRuleProtocolAh),
		string(SecurityRuleProtocolAny),
		string(SecurityRuleProtocolEsp),
		string(SecurityRuleProtocolIcmp),
		string(SecurityRuleProtocolTcp),
		string(SecurityRuleProtocolUdp),
	}
}

func (s *SecurityRuleProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRuleProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRuleProtocol(input string) (*SecurityRuleProtocol, error) {
	vals := map[string]SecurityRuleProtocol{
		"ah":   SecurityRuleProtocolAh,
		"*":    SecurityRuleProtocolAny,
		"esp":  SecurityRuleProtocolEsp,
		"icmp": SecurityRuleProtocolIcmp,
		"tcp":  SecurityRuleProtocolTcp,
		"udp":  SecurityRuleProtocolUdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRuleProtocol(input)
	return &out, nil
}

type TransportProtocol string

const (
	TransportProtocolAll TransportProtocol = "All"
	TransportProtocolTcp TransportProtocol = "Tcp"
	TransportProtocolUdp TransportProtocol = "Udp"
)

func PossibleValuesForTransportProtocol() []string {
	return []string{
		string(TransportProtocolAll),
		string(TransportProtocolTcp),
		string(TransportProtocolUdp),
	}
}

func (s *TransportProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransportProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransportProtocol(input string) (*TransportProtocol, error) {
	vals := map[string]TransportProtocol{
		"all": TransportProtocolAll,
		"tcp": TransportProtocolTcp,
		"udp": TransportProtocolUdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransportProtocol(input)
	return &out, nil
}

type VirtualNetworkEncryptionEnforcement string

const (
	VirtualNetworkEncryptionEnforcementAllowUnencrypted VirtualNetworkEncryptionEnforcement = "AllowUnencrypted"
	VirtualNetworkEncryptionEnforcementDropUnencrypted  VirtualNetworkEncryptionEnforcement = "DropUnencrypted"
)

func PossibleValuesForVirtualNetworkEncryptionEnforcement() []string {
	return []string{
		string(VirtualNetworkEncryptionEnforcementAllowUnencrypted),
		string(VirtualNetworkEncryptionEnforcementDropUnencrypted),
	}
}

func (s *VirtualNetworkEncryptionEnforcement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualNetworkEncryptionEnforcement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualNetworkEncryptionEnforcement(input string) (*VirtualNetworkEncryptionEnforcement, error) {
	vals := map[string]VirtualNetworkEncryptionEnforcement{
		"allowunencrypted": VirtualNetworkEncryptionEnforcementAllowUnencrypted,
		"dropunencrypted":  VirtualNetworkEncryptionEnforcementDropUnencrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkEncryptionEnforcement(input)
	return &out, nil
}

type VirtualNetworkPeeringLevel string

const (
	VirtualNetworkPeeringLevelFullyInSync             VirtualNetworkPeeringLevel = "FullyInSync"
	VirtualNetworkPeeringLevelLocalAndRemoteNotInSync VirtualNetworkPeeringLevel = "LocalAndRemoteNotInSync"
	VirtualNetworkPeeringLevelLocalNotInSync          VirtualNetworkPeeringLevel = "LocalNotInSync"
	VirtualNetworkPeeringLevelRemoteNotInSync         VirtualNetworkPeeringLevel = "RemoteNotInSync"
)

func PossibleValuesForVirtualNetworkPeeringLevel() []string {
	return []string{
		string(VirtualNetworkPeeringLevelFullyInSync),
		string(VirtualNetworkPeeringLevelLocalAndRemoteNotInSync),
		string(VirtualNetworkPeeringLevelLocalNotInSync),
		string(VirtualNetworkPeeringLevelRemoteNotInSync),
	}
}

func (s *VirtualNetworkPeeringLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualNetworkPeeringLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualNetworkPeeringLevel(input string) (*VirtualNetworkPeeringLevel, error) {
	vals := map[string]VirtualNetworkPeeringLevel{
		"fullyinsync":             VirtualNetworkPeeringLevelFullyInSync,
		"localandremotenotinsync": VirtualNetworkPeeringLevelLocalAndRemoteNotInSync,
		"localnotinsync":          VirtualNetworkPeeringLevelLocalNotInSync,
		"remotenotinsync":         VirtualNetworkPeeringLevelRemoteNotInSync,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkPeeringLevel(input)
	return &out, nil
}

type VirtualNetworkPeeringState string

const (
	VirtualNetworkPeeringStateConnected    VirtualNetworkPeeringState = "Connected"
	VirtualNetworkPeeringStateDisconnected VirtualNetworkPeeringState = "Disconnected"
	VirtualNetworkPeeringStateInitiated    VirtualNetworkPeeringState = "Initiated"
)

func PossibleValuesForVirtualNetworkPeeringState() []string {
	return []string{
		string(VirtualNetworkPeeringStateConnected),
		string(VirtualNetworkPeeringStateDisconnected),
		string(VirtualNetworkPeeringStateInitiated),
	}
}

func (s *VirtualNetworkPeeringState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualNetworkPeeringState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualNetworkPeeringState(input string) (*VirtualNetworkPeeringState, error) {
	vals := map[string]VirtualNetworkPeeringState{
		"connected":    VirtualNetworkPeeringStateConnected,
		"disconnected": VirtualNetworkPeeringStateDisconnected,
		"initiated":    VirtualNetworkPeeringStateInitiated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkPeeringState(input)
	return &out, nil
}

type VirtualNetworkPrivateEndpointNetworkPolicies string

const (
	VirtualNetworkPrivateEndpointNetworkPoliciesDisabled VirtualNetworkPrivateEndpointNetworkPolicies = "Disabled"
	VirtualNetworkPrivateEndpointNetworkPoliciesEnabled  VirtualNetworkPrivateEndpointNetworkPolicies = "Enabled"
)

func PossibleValuesForVirtualNetworkPrivateEndpointNetworkPolicies() []string {
	return []string{
		string(VirtualNetworkPrivateEndpointNetworkPoliciesDisabled),
		string(VirtualNetworkPrivateEndpointNetworkPoliciesEnabled),
	}
}

func (s *VirtualNetworkPrivateEndpointNetworkPolicies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualNetworkPrivateEndpointNetworkPolicies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualNetworkPrivateEndpointNetworkPolicies(input string) (*VirtualNetworkPrivateEndpointNetworkPolicies, error) {
	vals := map[string]VirtualNetworkPrivateEndpointNetworkPolicies{
		"disabled": VirtualNetworkPrivateEndpointNetworkPoliciesDisabled,
		"enabled":  VirtualNetworkPrivateEndpointNetworkPoliciesEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkPrivateEndpointNetworkPolicies(input)
	return &out, nil
}

type VirtualNetworkPrivateLinkServiceNetworkPolicies string

const (
	VirtualNetworkPrivateLinkServiceNetworkPoliciesDisabled VirtualNetworkPrivateLinkServiceNetworkPolicies = "Disabled"
	VirtualNetworkPrivateLinkServiceNetworkPoliciesEnabled  VirtualNetworkPrivateLinkServiceNetworkPolicies = "Enabled"
)

func PossibleValuesForVirtualNetworkPrivateLinkServiceNetworkPolicies() []string {
	return []string{
		string(VirtualNetworkPrivateLinkServiceNetworkPoliciesDisabled),
		string(VirtualNetworkPrivateLinkServiceNetworkPoliciesEnabled),
	}
}

func (s *VirtualNetworkPrivateLinkServiceNetworkPolicies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualNetworkPrivateLinkServiceNetworkPolicies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualNetworkPrivateLinkServiceNetworkPolicies(input string) (*VirtualNetworkPrivateLinkServiceNetworkPolicies, error) {
	vals := map[string]VirtualNetworkPrivateLinkServiceNetworkPolicies{
		"disabled": VirtualNetworkPrivateLinkServiceNetworkPoliciesDisabled,
		"enabled":  VirtualNetworkPrivateLinkServiceNetworkPoliciesEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualNetworkPrivateLinkServiceNetworkPolicies(input)
	return &out, nil
}
