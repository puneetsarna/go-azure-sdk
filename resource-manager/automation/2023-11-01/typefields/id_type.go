package typefields

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TypeId{}

// TypeId is a struct representing the Resource ID for a Type
type TypeId struct {
	SubscriptionId        string
	ResourceGroupName     string
	AutomationAccountName string
	ModuleName            string
	TypeName              string
}

// NewTypeID returns a new TypeId struct
func NewTypeID(subscriptionId string, resourceGroupName string, automationAccountName string, moduleName string, typeName string) TypeId {
	return TypeId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		AutomationAccountName: automationAccountName,
		ModuleName:            moduleName,
		TypeName:              typeName,
	}
}

// ParseTypeID parses 'input' into a TypeId
func ParseTypeID(input string) (*TypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(TypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "automationAccountName", *parsed)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "moduleName", *parsed)
	}

	if id.TypeName, ok = parsed.Parsed["typeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "typeName", *parsed)
	}

	return &id, nil
}

// ParseTypeIDInsensitively parses 'input' case-insensitively into a TypeId
// note: this method should only be used for API response data and not user input
func ParseTypeIDInsensitively(input string) (*TypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(TypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TypeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.AutomationAccountName, ok = parsed.Parsed["automationAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "automationAccountName", *parsed)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "moduleName", *parsed)
	}

	if id.TypeName, ok = parsed.Parsed["typeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "typeName", *parsed)
	}

	return &id, nil
}

// ValidateTypeID checks that 'input' can be parsed as a Type ID
func ValidateTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Type ID
func (id TypeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Automation/automationAccounts/%s/modules/%s/types/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutomationAccountName, id.ModuleName, id.TypeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Type ID
func (id TypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutomation", "Microsoft.Automation", "Microsoft.Automation"),
		resourceids.StaticSegment("staticAutomationAccounts", "automationAccounts", "automationAccounts"),
		resourceids.UserSpecifiedSegment("automationAccountName", "automationAccountValue"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleValue"),
		resourceids.StaticSegment("staticTypes", "types", "types"),
		resourceids.UserSpecifiedSegment("typeName", "typeValue"),
	}
}

// String returns a human-readable description of this Type ID
func (id TypeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Automation Account Name: %q", id.AutomationAccountName),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
		fmt.Sprintf("Type Name: %q", id.TypeName),
	}
	return fmt.Sprintf("Type (%s)", strings.Join(components, "\n"))
}
