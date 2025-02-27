package allowedconnections

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedConnectionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAllowedConnectionsClientWithBaseURI(endpoint string) AllowedConnectionsClient {
	return AllowedConnectionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
