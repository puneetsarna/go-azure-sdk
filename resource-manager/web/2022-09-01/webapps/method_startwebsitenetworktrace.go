package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartWebSiteNetworkTraceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *string
}

type StartWebSiteNetworkTraceOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartWebSiteNetworkTraceOperationOptions() StartWebSiteNetworkTraceOperationOptions {
	return StartWebSiteNetworkTraceOperationOptions{}
}

func (o StartWebSiteNetworkTraceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StartWebSiteNetworkTraceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o StartWebSiteNetworkTraceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DurationInSeconds != nil {
		out.Append("durationInSeconds", fmt.Sprintf("%v", *o.DurationInSeconds))
	}
	if o.MaxFrameLength != nil {
		out.Append("maxFrameLength", fmt.Sprintf("%v", *o.MaxFrameLength))
	}
	if o.SasUrl != nil {
		out.Append("sasUrl", fmt.Sprintf("%v", *o.SasUrl))
	}
	return &out
}

// StartWebSiteNetworkTrace ...
func (c WebAppsClient) StartWebSiteNetworkTrace(ctx context.Context, id commonids.AppServiceId, options StartWebSiteNetworkTraceOperationOptions) (result StartWebSiteNetworkTraceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/networkTrace/start", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
