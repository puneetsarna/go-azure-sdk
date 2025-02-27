package trustedaccess

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleBindingsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *TrustedAccessRoleBinding
}

// RoleBindingsCreateOrUpdate ...
func (c TrustedAccessClient) RoleBindingsCreateOrUpdate(ctx context.Context, id TrustedAccessRoleBindingId, input TrustedAccessRoleBinding) (result RoleBindingsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRoleBindingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trustedaccess.TrustedAccessClient", "RoleBindingsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "trustedaccess.TrustedAccessClient", "RoleBindingsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRoleBindingsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trustedaccess.TrustedAccessClient", "RoleBindingsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRoleBindingsCreateOrUpdate prepares the RoleBindingsCreateOrUpdate request.
func (c TrustedAccessClient) preparerForRoleBindingsCreateOrUpdate(ctx context.Context, id TrustedAccessRoleBindingId, input TrustedAccessRoleBinding) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRoleBindingsCreateOrUpdate handles the response to the RoleBindingsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c TrustedAccessClient) responderForRoleBindingsCreateOrUpdate(resp *http.Response) (result RoleBindingsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
