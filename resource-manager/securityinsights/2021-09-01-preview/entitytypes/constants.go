package entitytypes

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertSeverity string

const (
	AlertSeverityHigh          AlertSeverity = "High"
	AlertSeverityInformational AlertSeverity = "Informational"
	AlertSeverityLow           AlertSeverity = "Low"
	AlertSeverityMedium        AlertSeverity = "Medium"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeverityHigh),
		string(AlertSeverityInformational),
		string(AlertSeverityLow),
		string(AlertSeverityMedium),
	}
}

func parseAlertSeverity(input string) (*AlertSeverity, error) {
	vals := map[string]AlertSeverity{
		"high":          AlertSeverityHigh,
		"informational": AlertSeverityInformational,
		"low":           AlertSeverityLow,
		"medium":        AlertSeverityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertSeverity(input)
	return &out, nil
}

type AlertStatus string

const (
	AlertStatusDismissed  AlertStatus = "Dismissed"
	AlertStatusInProgress AlertStatus = "InProgress"
	AlertStatusNew        AlertStatus = "New"
	AlertStatusResolved   AlertStatus = "Resolved"
	AlertStatusUnknown    AlertStatus = "Unknown"
)

func PossibleValuesForAlertStatus() []string {
	return []string{
		string(AlertStatusDismissed),
		string(AlertStatusInProgress),
		string(AlertStatusNew),
		string(AlertStatusResolved),
		string(AlertStatusUnknown),
	}
}

func parseAlertStatus(input string) (*AlertStatus, error) {
	vals := map[string]AlertStatus{
		"dismissed":  AlertStatusDismissed,
		"inprogress": AlertStatusInProgress,
		"new":        AlertStatusNew,
		"resolved":   AlertStatusResolved,
		"unknown":    AlertStatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertStatus(input)
	return &out, nil
}

type AntispamMailDirection string

const (
	AntispamMailDirectionInbound  AntispamMailDirection = "Inbound"
	AntispamMailDirectionIntraorg AntispamMailDirection = "Intraorg"
	AntispamMailDirectionOutbound AntispamMailDirection = "Outbound"
	AntispamMailDirectionUnknown  AntispamMailDirection = "Unknown"
)

func PossibleValuesForAntispamMailDirection() []string {
	return []string{
		string(AntispamMailDirectionInbound),
		string(AntispamMailDirectionIntraorg),
		string(AntispamMailDirectionOutbound),
		string(AntispamMailDirectionUnknown),
	}
}

func parseAntispamMailDirection(input string) (*AntispamMailDirection, error) {
	vals := map[string]AntispamMailDirection{
		"inbound":  AntispamMailDirectionInbound,
		"intraorg": AntispamMailDirectionIntraorg,
		"outbound": AntispamMailDirectionOutbound,
		"unknown":  AntispamMailDirectionUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AntispamMailDirection(input)
	return &out, nil
}

type AttackTactic string

const (
	AttackTacticCollection          AttackTactic = "Collection"
	AttackTacticCommandAndControl   AttackTactic = "CommandAndControl"
	AttackTacticCredentialAccess    AttackTactic = "CredentialAccess"
	AttackTacticDefenseEvasion      AttackTactic = "DefenseEvasion"
	AttackTacticDiscovery           AttackTactic = "Discovery"
	AttackTacticExecution           AttackTactic = "Execution"
	AttackTacticExfiltration        AttackTactic = "Exfiltration"
	AttackTacticImpact              AttackTactic = "Impact"
	AttackTacticInitialAccess       AttackTactic = "InitialAccess"
	AttackTacticLateralMovement     AttackTactic = "LateralMovement"
	AttackTacticPersistence         AttackTactic = "Persistence"
	AttackTacticPreAttack           AttackTactic = "PreAttack"
	AttackTacticPrivilegeEscalation AttackTactic = "PrivilegeEscalation"
)

func PossibleValuesForAttackTactic() []string {
	return []string{
		string(AttackTacticCollection),
		string(AttackTacticCommandAndControl),
		string(AttackTacticCredentialAccess),
		string(AttackTacticDefenseEvasion),
		string(AttackTacticDiscovery),
		string(AttackTacticExecution),
		string(AttackTacticExfiltration),
		string(AttackTacticImpact),
		string(AttackTacticInitialAccess),
		string(AttackTacticLateralMovement),
		string(AttackTacticPersistence),
		string(AttackTacticPreAttack),
		string(AttackTacticPrivilegeEscalation),
	}
}

func parseAttackTactic(input string) (*AttackTactic, error) {
	vals := map[string]AttackTactic{
		"collection":          AttackTacticCollection,
		"commandandcontrol":   AttackTacticCommandAndControl,
		"credentialaccess":    AttackTacticCredentialAccess,
		"defenseevasion":      AttackTacticDefenseEvasion,
		"discovery":           AttackTacticDiscovery,
		"execution":           AttackTacticExecution,
		"exfiltration":        AttackTacticExfiltration,
		"impact":              AttackTacticImpact,
		"initialaccess":       AttackTacticInitialAccess,
		"lateralmovement":     AttackTacticLateralMovement,
		"persistence":         AttackTacticPersistence,
		"preattack":           AttackTacticPreAttack,
		"privilegeescalation": AttackTacticPrivilegeEscalation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttackTactic(input)
	return &out, nil
}

type ConfidenceLevel string

const (
	ConfidenceLevelHigh    ConfidenceLevel = "High"
	ConfidenceLevelLow     ConfidenceLevel = "Low"
	ConfidenceLevelUnknown ConfidenceLevel = "Unknown"
)

func PossibleValuesForConfidenceLevel() []string {
	return []string{
		string(ConfidenceLevelHigh),
		string(ConfidenceLevelLow),
		string(ConfidenceLevelUnknown),
	}
}

func parseConfidenceLevel(input string) (*ConfidenceLevel, error) {
	vals := map[string]ConfidenceLevel{
		"high":    ConfidenceLevelHigh,
		"low":     ConfidenceLevelLow,
		"unknown": ConfidenceLevelUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfidenceLevel(input)
	return &out, nil
}

type ConfidenceScoreStatus string

const (
	ConfidenceScoreStatusFinal         ConfidenceScoreStatus = "Final"
	ConfidenceScoreStatusInProcess     ConfidenceScoreStatus = "InProcess"
	ConfidenceScoreStatusNotApplicable ConfidenceScoreStatus = "NotApplicable"
	ConfidenceScoreStatusNotFinal      ConfidenceScoreStatus = "NotFinal"
)

func PossibleValuesForConfidenceScoreStatus() []string {
	return []string{
		string(ConfidenceScoreStatusFinal),
		string(ConfidenceScoreStatusInProcess),
		string(ConfidenceScoreStatusNotApplicable),
		string(ConfidenceScoreStatusNotFinal),
	}
}

func parseConfidenceScoreStatus(input string) (*ConfidenceScoreStatus, error) {
	vals := map[string]ConfidenceScoreStatus{
		"final":         ConfidenceScoreStatusFinal,
		"inprocess":     ConfidenceScoreStatusInProcess,
		"notapplicable": ConfidenceScoreStatusNotApplicable,
		"notfinal":      ConfidenceScoreStatusNotFinal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfidenceScoreStatus(input)
	return &out, nil
}

type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

func PossibleValuesForCreatedByType() []string {
	return []string{
		string(CreatedByTypeApplication),
		string(CreatedByTypeKey),
		string(CreatedByTypeManagedIdentity),
		string(CreatedByTypeUser),
	}
}

func parseCreatedByType(input string) (*CreatedByType, error) {
	vals := map[string]CreatedByType{
		"application":     CreatedByTypeApplication,
		"key":             CreatedByTypeKey,
		"managedidentity": CreatedByTypeManagedIdentity,
		"user":            CreatedByTypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreatedByType(input)
	return &out, nil
}

type DeliveryAction string

const (
	DeliveryActionBlocked         DeliveryAction = "Blocked"
	DeliveryActionDelivered       DeliveryAction = "Delivered"
	DeliveryActionDeliveredAsSpam DeliveryAction = "DeliveredAsSpam"
	DeliveryActionReplaced        DeliveryAction = "Replaced"
	DeliveryActionUnknown         DeliveryAction = "Unknown"
)

func PossibleValuesForDeliveryAction() []string {
	return []string{
		string(DeliveryActionBlocked),
		string(DeliveryActionDelivered),
		string(DeliveryActionDeliveredAsSpam),
		string(DeliveryActionReplaced),
		string(DeliveryActionUnknown),
	}
}

func parseDeliveryAction(input string) (*DeliveryAction, error) {
	vals := map[string]DeliveryAction{
		"blocked":         DeliveryActionBlocked,
		"delivered":       DeliveryActionDelivered,
		"deliveredasspam": DeliveryActionDeliveredAsSpam,
		"replaced":        DeliveryActionReplaced,
		"unknown":         DeliveryActionUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryAction(input)
	return &out, nil
}

type DeliveryLocation string

const (
	DeliveryLocationDeletedFolder DeliveryLocation = "DeletedFolder"
	DeliveryLocationDropped       DeliveryLocation = "Dropped"
	DeliveryLocationExternal      DeliveryLocation = "External"
	DeliveryLocationFailed        DeliveryLocation = "Failed"
	DeliveryLocationForwarded     DeliveryLocation = "Forwarded"
	DeliveryLocationInbox         DeliveryLocation = "Inbox"
	DeliveryLocationJunkFolder    DeliveryLocation = "JunkFolder"
	DeliveryLocationQuarantine    DeliveryLocation = "Quarantine"
	DeliveryLocationUnknown       DeliveryLocation = "Unknown"
)

func PossibleValuesForDeliveryLocation() []string {
	return []string{
		string(DeliveryLocationDeletedFolder),
		string(DeliveryLocationDropped),
		string(DeliveryLocationExternal),
		string(DeliveryLocationFailed),
		string(DeliveryLocationForwarded),
		string(DeliveryLocationInbox),
		string(DeliveryLocationJunkFolder),
		string(DeliveryLocationQuarantine),
		string(DeliveryLocationUnknown),
	}
}

func parseDeliveryLocation(input string) (*DeliveryLocation, error) {
	vals := map[string]DeliveryLocation{
		"deletedfolder": DeliveryLocationDeletedFolder,
		"dropped":       DeliveryLocationDropped,
		"external":      DeliveryLocationExternal,
		"failed":        DeliveryLocationFailed,
		"forwarded":     DeliveryLocationForwarded,
		"inbox":         DeliveryLocationInbox,
		"junkfolder":    DeliveryLocationJunkFolder,
		"quarantine":    DeliveryLocationQuarantine,
		"unknown":       DeliveryLocationUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryLocation(input)
	return &out, nil
}

type ElevationToken string

const (
	ElevationTokenDefault ElevationToken = "Default"
	ElevationTokenFull    ElevationToken = "Full"
	ElevationTokenLimited ElevationToken = "Limited"
)

func PossibleValuesForElevationToken() []string {
	return []string{
		string(ElevationTokenDefault),
		string(ElevationTokenFull),
		string(ElevationTokenLimited),
	}
}

func parseElevationToken(input string) (*ElevationToken, error) {
	vals := map[string]ElevationToken{
		"default": ElevationTokenDefault,
		"full":    ElevationTokenFull,
		"limited": ElevationTokenLimited,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElevationToken(input)
	return &out, nil
}

type EntityKind string

const (
	EntityKindAccount          EntityKind = "Account"
	EntityKindAzureResource    EntityKind = "AzureResource"
	EntityKindBookmark         EntityKind = "Bookmark"
	EntityKindCloudApplication EntityKind = "CloudApplication"
	EntityKindDnsResolution    EntityKind = "DnsResolution"
	EntityKindFile             EntityKind = "File"
	EntityKindFileHash         EntityKind = "FileHash"
	EntityKindHost             EntityKind = "Host"
	EntityKindIP               EntityKind = "Ip"
	EntityKindIoTDevice        EntityKind = "IoTDevice"
	EntityKindMailCluster      EntityKind = "MailCluster"
	EntityKindMailMessage      EntityKind = "MailMessage"
	EntityKindMailbox          EntityKind = "Mailbox"
	EntityKindMalware          EntityKind = "Malware"
	EntityKindProcess          EntityKind = "Process"
	EntityKindRegistryKey      EntityKind = "RegistryKey"
	EntityKindRegistryValue    EntityKind = "RegistryValue"
	EntityKindSecurityAlert    EntityKind = "SecurityAlert"
	EntityKindSecurityGroup    EntityKind = "SecurityGroup"
	EntityKindSubmissionMail   EntityKind = "SubmissionMail"
	EntityKindUrl              EntityKind = "Url"
)

func PossibleValuesForEntityKind() []string {
	return []string{
		string(EntityKindAccount),
		string(EntityKindAzureResource),
		string(EntityKindBookmark),
		string(EntityKindCloudApplication),
		string(EntityKindDnsResolution),
		string(EntityKindFile),
		string(EntityKindFileHash),
		string(EntityKindHost),
		string(EntityKindIP),
		string(EntityKindIoTDevice),
		string(EntityKindMailCluster),
		string(EntityKindMailMessage),
		string(EntityKindMailbox),
		string(EntityKindMalware),
		string(EntityKindProcess),
		string(EntityKindRegistryKey),
		string(EntityKindRegistryValue),
		string(EntityKindSecurityAlert),
		string(EntityKindSecurityGroup),
		string(EntityKindSubmissionMail),
		string(EntityKindUrl),
	}
}

func parseEntityKind(input string) (*EntityKind, error) {
	vals := map[string]EntityKind{
		"account":          EntityKindAccount,
		"azureresource":    EntityKindAzureResource,
		"bookmark":         EntityKindBookmark,
		"cloudapplication": EntityKindCloudApplication,
		"dnsresolution":    EntityKindDnsResolution,
		"file":             EntityKindFile,
		"filehash":         EntityKindFileHash,
		"host":             EntityKindHost,
		"ip":               EntityKindIP,
		"iotdevice":        EntityKindIoTDevice,
		"mailcluster":      EntityKindMailCluster,
		"mailmessage":      EntityKindMailMessage,
		"mailbox":          EntityKindMailbox,
		"malware":          EntityKindMalware,
		"process":          EntityKindProcess,
		"registrykey":      EntityKindRegistryKey,
		"registryvalue":    EntityKindRegistryValue,
		"securityalert":    EntityKindSecurityAlert,
		"securitygroup":    EntityKindSecurityGroup,
		"submissionmail":   EntityKindSubmissionMail,
		"url":              EntityKindUrl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityKind(input)
	return &out, nil
}

type FileHashAlgorithm string

const (
	FileHashAlgorithmMDFive          FileHashAlgorithm = "MD5"
	FileHashAlgorithmSHAOne          FileHashAlgorithm = "SHA1"
	FileHashAlgorithmSHATwoFiveSix   FileHashAlgorithm = "SHA256"
	FileHashAlgorithmSHATwoFiveSixAC FileHashAlgorithm = "SHA256AC"
	FileHashAlgorithmUnknown         FileHashAlgorithm = "Unknown"
)

func PossibleValuesForFileHashAlgorithm() []string {
	return []string{
		string(FileHashAlgorithmMDFive),
		string(FileHashAlgorithmSHAOne),
		string(FileHashAlgorithmSHATwoFiveSix),
		string(FileHashAlgorithmSHATwoFiveSixAC),
		string(FileHashAlgorithmUnknown),
	}
}

func parseFileHashAlgorithm(input string) (*FileHashAlgorithm, error) {
	vals := map[string]FileHashAlgorithm{
		"md5":      FileHashAlgorithmMDFive,
		"sha1":     FileHashAlgorithmSHAOne,
		"sha256":   FileHashAlgorithmSHATwoFiveSix,
		"sha256ac": FileHashAlgorithmSHATwoFiveSixAC,
		"unknown":  FileHashAlgorithmUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileHashAlgorithm(input)
	return &out, nil
}

type IncidentSeverity string

const (
	IncidentSeverityHigh          IncidentSeverity = "High"
	IncidentSeverityInformational IncidentSeverity = "Informational"
	IncidentSeverityLow           IncidentSeverity = "Low"
	IncidentSeverityMedium        IncidentSeverity = "Medium"
)

func PossibleValuesForIncidentSeverity() []string {
	return []string{
		string(IncidentSeverityHigh),
		string(IncidentSeverityInformational),
		string(IncidentSeverityLow),
		string(IncidentSeverityMedium),
	}
}

func parseIncidentSeverity(input string) (*IncidentSeverity, error) {
	vals := map[string]IncidentSeverity{
		"high":          IncidentSeverityHigh,
		"informational": IncidentSeverityInformational,
		"low":           IncidentSeverityLow,
		"medium":        IncidentSeverityMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncidentSeverity(input)
	return &out, nil
}

type KillChainIntent string

const (
	KillChainIntentCollection          KillChainIntent = "Collection"
	KillChainIntentCommandAndControl   KillChainIntent = "CommandAndControl"
	KillChainIntentCredentialAccess    KillChainIntent = "CredentialAccess"
	KillChainIntentDefenseEvasion      KillChainIntent = "DefenseEvasion"
	KillChainIntentDiscovery           KillChainIntent = "Discovery"
	KillChainIntentExecution           KillChainIntent = "Execution"
	KillChainIntentExfiltration        KillChainIntent = "Exfiltration"
	KillChainIntentExploitation        KillChainIntent = "Exploitation"
	KillChainIntentImpact              KillChainIntent = "Impact"
	KillChainIntentLateralMovement     KillChainIntent = "LateralMovement"
	KillChainIntentPersistence         KillChainIntent = "Persistence"
	KillChainIntentPrivilegeEscalation KillChainIntent = "PrivilegeEscalation"
	KillChainIntentProbing             KillChainIntent = "Probing"
	KillChainIntentUnknown             KillChainIntent = "Unknown"
)

func PossibleValuesForKillChainIntent() []string {
	return []string{
		string(KillChainIntentCollection),
		string(KillChainIntentCommandAndControl),
		string(KillChainIntentCredentialAccess),
		string(KillChainIntentDefenseEvasion),
		string(KillChainIntentDiscovery),
		string(KillChainIntentExecution),
		string(KillChainIntentExfiltration),
		string(KillChainIntentExploitation),
		string(KillChainIntentImpact),
		string(KillChainIntentLateralMovement),
		string(KillChainIntentPersistence),
		string(KillChainIntentPrivilegeEscalation),
		string(KillChainIntentProbing),
		string(KillChainIntentUnknown),
	}
}

func parseKillChainIntent(input string) (*KillChainIntent, error) {
	vals := map[string]KillChainIntent{
		"collection":          KillChainIntentCollection,
		"commandandcontrol":   KillChainIntentCommandAndControl,
		"credentialaccess":    KillChainIntentCredentialAccess,
		"defenseevasion":      KillChainIntentDefenseEvasion,
		"discovery":           KillChainIntentDiscovery,
		"execution":           KillChainIntentExecution,
		"exfiltration":        KillChainIntentExfiltration,
		"exploitation":        KillChainIntentExploitation,
		"impact":              KillChainIntentImpact,
		"lateralmovement":     KillChainIntentLateralMovement,
		"persistence":         KillChainIntentPersistence,
		"privilegeescalation": KillChainIntentPrivilegeEscalation,
		"probing":             KillChainIntentProbing,
		"unknown":             KillChainIntentUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KillChainIntent(input)
	return &out, nil
}

type OSFamily string

const (
	OSFamilyAndroid OSFamily = "Android"
	OSFamilyIOS     OSFamily = "IOS"
	OSFamilyLinux   OSFamily = "Linux"
	OSFamilyUnknown OSFamily = "Unknown"
	OSFamilyWindows OSFamily = "Windows"
)

func PossibleValuesForOSFamily() []string {
	return []string{
		string(OSFamilyAndroid),
		string(OSFamilyIOS),
		string(OSFamilyLinux),
		string(OSFamilyUnknown),
		string(OSFamilyWindows),
	}
}

func parseOSFamily(input string) (*OSFamily, error) {
	vals := map[string]OSFamily{
		"android": OSFamilyAndroid,
		"ios":     OSFamilyIOS,
		"linux":   OSFamilyLinux,
		"unknown": OSFamilyUnknown,
		"windows": OSFamilyWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSFamily(input)
	return &out, nil
}

type RegistryHive string

const (
	RegistryHiveHKEYA                        RegistryHive = "HKEY_A"
	RegistryHiveHKEYCLASSESROOT              RegistryHive = "HKEY_CLASSES_ROOT"
	RegistryHiveHKEYCURRENTCONFIG            RegistryHive = "HKEY_CURRENT_CONFIG"
	RegistryHiveHKEYCURRENTUSER              RegistryHive = "HKEY_CURRENT_USER"
	RegistryHiveHKEYCURRENTUSERLOCALSETTINGS RegistryHive = "HKEY_CURRENT_USER_LOCAL_SETTINGS"
	RegistryHiveHKEYLOCALMACHINE             RegistryHive = "HKEY_LOCAL_MACHINE"
	RegistryHiveHKEYPERFORMANCEDATA          RegistryHive = "HKEY_PERFORMANCE_DATA"
	RegistryHiveHKEYPERFORMANCENLSTEXT       RegistryHive = "HKEY_PERFORMANCE_NLSTEXT"
	RegistryHiveHKEYPERFORMANCETEXT          RegistryHive = "HKEY_PERFORMANCE_TEXT"
	RegistryHiveHKEYUSERS                    RegistryHive = "HKEY_USERS"
)

func PossibleValuesForRegistryHive() []string {
	return []string{
		string(RegistryHiveHKEYA),
		string(RegistryHiveHKEYCLASSESROOT),
		string(RegistryHiveHKEYCURRENTCONFIG),
		string(RegistryHiveHKEYCURRENTUSER),
		string(RegistryHiveHKEYCURRENTUSERLOCALSETTINGS),
		string(RegistryHiveHKEYLOCALMACHINE),
		string(RegistryHiveHKEYPERFORMANCEDATA),
		string(RegistryHiveHKEYPERFORMANCENLSTEXT),
		string(RegistryHiveHKEYPERFORMANCETEXT),
		string(RegistryHiveHKEYUSERS),
	}
}

func parseRegistryHive(input string) (*RegistryHive, error) {
	vals := map[string]RegistryHive{
		"hkey_a":                           RegistryHiveHKEYA,
		"hkey_classes_root":                RegistryHiveHKEYCLASSESROOT,
		"hkey_current_config":              RegistryHiveHKEYCURRENTCONFIG,
		"hkey_current_user":                RegistryHiveHKEYCURRENTUSER,
		"hkey_current_user_local_settings": RegistryHiveHKEYCURRENTUSERLOCALSETTINGS,
		"hkey_local_machine":               RegistryHiveHKEYLOCALMACHINE,
		"hkey_performance_data":            RegistryHiveHKEYPERFORMANCEDATA,
		"hkey_performance_nlstext":         RegistryHiveHKEYPERFORMANCENLSTEXT,
		"hkey_performance_text":            RegistryHiveHKEYPERFORMANCETEXT,
		"hkey_users":                       RegistryHiveHKEYUSERS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryHive(input)
	return &out, nil
}

type RegistryValueKind string

const (
	RegistryValueKindBinary       RegistryValueKind = "Binary"
	RegistryValueKindDWord        RegistryValueKind = "DWord"
	RegistryValueKindExpandString RegistryValueKind = "ExpandString"
	RegistryValueKindMultiString  RegistryValueKind = "MultiString"
	RegistryValueKindNone         RegistryValueKind = "None"
	RegistryValueKindQWord        RegistryValueKind = "QWord"
	RegistryValueKindString       RegistryValueKind = "String"
	RegistryValueKindUnknown      RegistryValueKind = "Unknown"
)

func PossibleValuesForRegistryValueKind() []string {
	return []string{
		string(RegistryValueKindBinary),
		string(RegistryValueKindDWord),
		string(RegistryValueKindExpandString),
		string(RegistryValueKindMultiString),
		string(RegistryValueKindNone),
		string(RegistryValueKindQWord),
		string(RegistryValueKindString),
		string(RegistryValueKindUnknown),
	}
}

func parseRegistryValueKind(input string) (*RegistryValueKind, error) {
	vals := map[string]RegistryValueKind{
		"binary":       RegistryValueKindBinary,
		"dword":        RegistryValueKindDWord,
		"expandstring": RegistryValueKindExpandString,
		"multistring":  RegistryValueKindMultiString,
		"none":         RegistryValueKindNone,
		"qword":        RegistryValueKindQWord,
		"string":       RegistryValueKindString,
		"unknown":      RegistryValueKindUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistryValueKind(input)
	return &out, nil
}
