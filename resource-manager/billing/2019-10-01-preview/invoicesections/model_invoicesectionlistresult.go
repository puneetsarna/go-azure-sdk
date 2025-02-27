package invoicesections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceSectionListResult struct {
	NextLink *string           `json:"nextLink,omitempty"`
	Value    *[]InvoiceSection `json:"value,omitempty"`
}
