package flowlogs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = NetworkWatcherId{}

// NetworkWatcherId is a struct representing the Resource ID for a Network Watcher
type NetworkWatcherId struct {
	SubscriptionId     string
	ResourceGroupName  string
	NetworkWatcherName string
}

// NewNetworkWatcherID returns a new NetworkWatcherId struct
func NewNetworkWatcherID(subscriptionId string, resourceGroupName string, networkWatcherName string) NetworkWatcherId {
	return NetworkWatcherId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		NetworkWatcherName: networkWatcherName,
	}
}

// ParseNetworkWatcherID parses 'input' into a NetworkWatcherId
func ParseNetworkWatcherID(input string) (*NetworkWatcherId, error) {
	parser := resourceids.NewParserFromResourceIdType(NetworkWatcherId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NetworkWatcherId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NetworkWatcherName, ok = parsed.Parsed["networkWatcherName"]; !ok {
		return nil, fmt.Errorf("the segment 'networkWatcherName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseNetworkWatcherIDInsensitively parses 'input' case-insensitively into a NetworkWatcherId
// note: this method should only be used for API response data and not user input
func ParseNetworkWatcherIDInsensitively(input string) (*NetworkWatcherId, error) {
	parser := resourceids.NewParserFromResourceIdType(NetworkWatcherId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NetworkWatcherId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.NetworkWatcherName, ok = parsed.Parsed["networkWatcherName"]; !ok {
		return nil, fmt.Errorf("the segment 'networkWatcherName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateNetworkWatcherID checks that 'input' can be parsed as a Network Watcher ID
func ValidateNetworkWatcherID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNetworkWatcherID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Network Watcher ID
func (id NetworkWatcherId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/networkWatchers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetworkWatcherName)
}

// Segments returns a slice of Resource ID Segments which comprise this Network Watcher ID
func (id NetworkWatcherId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticNetworkWatchers", "networkWatchers", "networkWatchers"),
		resourceids.UserSpecifiedSegment("networkWatcherName", "networkWatcherValue"),
	}
}

// String returns a human-readable description of this Network Watcher ID
func (id NetworkWatcherId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Network Watcher Name: %q", id.NetworkWatcherName),
	}
	return fmt.Sprintf("Network Watcher (%s)", strings.Join(components, "\n"))
}
