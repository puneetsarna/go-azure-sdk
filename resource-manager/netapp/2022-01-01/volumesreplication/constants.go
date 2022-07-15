package volumesreplication

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointType string

const (
	EndpointTypeDst EndpointType = "dst"
	EndpointTypeSrc EndpointType = "src"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeDst),
		string(EndpointTypeSrc),
	}
}

func parseEndpointType(input string) (*EndpointType, error) {
	vals := map[string]EndpointType{
		"dst": EndpointTypeDst,
		"src": EndpointTypeSrc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointType(input)
	return &out, nil
}

type MirrorState string

const (
	MirrorStateBroken        MirrorState = "Broken"
	MirrorStateMirrored      MirrorState = "Mirrored"
	MirrorStateUninitialized MirrorState = "Uninitialized"
)

func PossibleValuesForMirrorState() []string {
	return []string{
		string(MirrorStateBroken),
		string(MirrorStateMirrored),
		string(MirrorStateUninitialized),
	}
}

func parseMirrorState(input string) (*MirrorState, error) {
	vals := map[string]MirrorState{
		"broken":        MirrorStateBroken,
		"mirrored":      MirrorStateMirrored,
		"uninitialized": MirrorStateUninitialized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MirrorState(input)
	return &out, nil
}

type RelationshipStatus string

const (
	RelationshipStatusIdle         RelationshipStatus = "Idle"
	RelationshipStatusTransferring RelationshipStatus = "Transferring"
)

func PossibleValuesForRelationshipStatus() []string {
	return []string{
		string(RelationshipStatusIdle),
		string(RelationshipStatusTransferring),
	}
}

func parseRelationshipStatus(input string) (*RelationshipStatus, error) {
	vals := map[string]RelationshipStatus{
		"idle":         RelationshipStatusIdle,
		"transferring": RelationshipStatusTransferring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelationshipStatus(input)
	return &out, nil
}

type ReplicationSchedule string

const (
	ReplicationScheduleDaily           ReplicationSchedule = "daily"
	ReplicationScheduleHourly          ReplicationSchedule = "hourly"
	ReplicationScheduleOneZerominutely ReplicationSchedule = "_10minutely"
)

func PossibleValuesForReplicationSchedule() []string {
	return []string{
		string(ReplicationScheduleDaily),
		string(ReplicationScheduleHourly),
		string(ReplicationScheduleOneZerominutely),
	}
}

func parseReplicationSchedule(input string) (*ReplicationSchedule, error) {
	vals := map[string]ReplicationSchedule{
		"daily":       ReplicationScheduleDaily,
		"hourly":      ReplicationScheduleHourly,
		"_10minutely": ReplicationScheduleOneZerominutely,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationSchedule(input)
	return &out, nil
}
