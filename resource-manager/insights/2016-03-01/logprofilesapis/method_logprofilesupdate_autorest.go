package logprofilesapis

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogProfilesUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *LogProfileResource
}

// LogProfilesUpdate ...
func (c LogProfilesAPIsClient) LogProfilesUpdate(ctx context.Context, id LogProfileId, input LogProfileResourcePatch) (result LogProfilesUpdateOperationResponse, err error) {
	req, err := c.preparerForLogProfilesUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logprofilesapis.LogProfilesAPIsClient", "LogProfilesUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "logprofilesapis.LogProfilesAPIsClient", "LogProfilesUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForLogProfilesUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logprofilesapis.LogProfilesAPIsClient", "LogProfilesUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForLogProfilesUpdate prepares the LogProfilesUpdate request.
func (c LogProfilesAPIsClient) preparerForLogProfilesUpdate(ctx context.Context, id LogProfileId, input LogProfileResourcePatch) (*http.Request, error) {
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

// responderForLogProfilesUpdate handles the response to the LogProfilesUpdate request. The method always
// closes the http.Response Body.
func (c LogProfilesAPIsClient) responderForLogProfilesUpdate(resp *http.Response) (result LogProfilesUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
