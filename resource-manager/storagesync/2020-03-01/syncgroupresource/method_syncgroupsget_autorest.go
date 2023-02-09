package syncgroupresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SyncGroup
}

// SyncGroupsGet ...
func (c SyncGroupResourceClient) SyncGroupsGet(ctx context.Context, id SyncGroupId) (result SyncGroupsGetOperationResponse, err error) {
	req, err := c.preparerForSyncGroupsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSyncGroupsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "syncgroupresource.SyncGroupResourceClient", "SyncGroupsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSyncGroupsGet prepares the SyncGroupsGet request.
func (c SyncGroupResourceClient) preparerForSyncGroupsGet(ctx context.Context, id SyncGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSyncGroupsGet handles the response to the SyncGroupsGet request. The method always
// closes the http.Response Body.
func (c SyncGroupResourceClient) responderForSyncGroupsGet(resp *http.Response) (result SyncGroupsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
