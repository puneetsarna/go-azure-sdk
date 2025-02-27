package enrichment

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrichmentClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEnrichmentClientWithBaseURI(endpoint string) EnrichmentClient {
	return EnrichmentClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
