package iotsecuritysolutions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotSecuritySolutionUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecuritySolutionModel
}

// IotSecuritySolutionUpdate ...
func (c IotSecuritySolutionsClient) IotSecuritySolutionUpdate(ctx context.Context, id IotSecuritySolutionId, input UpdateIotSecuritySolutionData) (result IotSecuritySolutionUpdateOperationResponse, err error) {
	req, err := c.preparerForIotSecuritySolutionUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotSecuritySolutionUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotSecuritySolutionUpdate prepares the IotSecuritySolutionUpdate request.
func (c IotSecuritySolutionsClient) preparerForIotSecuritySolutionUpdate(ctx context.Context, id IotSecuritySolutionId, input UpdateIotSecuritySolutionData) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIotSecuritySolutionUpdate handles the response to the IotSecuritySolutionUpdate request. The method always
// closes the http.Response Body.
func (c IotSecuritySolutionsClient) responderForIotSecuritySolutionUpdate(resp *http.Response) (result IotSecuritySolutionUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
