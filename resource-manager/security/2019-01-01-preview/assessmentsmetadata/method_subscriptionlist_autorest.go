package assessmentsmetadata

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SecurityAssessmentMetadata

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SubscriptionListOperationResponse, error)
}

type SubscriptionListCompleteResult struct {
	Items []SecurityAssessmentMetadata
}

func (r SubscriptionListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SubscriptionListOperationResponse) LoadMore(ctx context.Context) (resp SubscriptionListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SubscriptionList ...
func (c AssessmentsMetadataClient) SubscriptionList(ctx context.Context, id commonids.SubscriptionId) (resp SubscriptionListOperationResponse, err error) {
	req, err := c.preparerForSubscriptionList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSubscriptionList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSubscriptionList prepares the SubscriptionList request.
func (c AssessmentsMetadataClient) preparerForSubscriptionList(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/assessmentMetadata", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSubscriptionListWithNextLink prepares the SubscriptionList request with the given nextLink token.
func (c AssessmentsMetadataClient) preparerForSubscriptionListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSubscriptionList handles the response to the SubscriptionList request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForSubscriptionList(resp *http.Response) (result SubscriptionListOperationResponse, err error) {
	type page struct {
		Values   []SecurityAssessmentMetadata `json:"value"`
		NextLink *string                      `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SubscriptionListOperationResponse, err error) {
			req, err := c.preparerForSubscriptionListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSubscriptionList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SubscriptionListComplete retrieves all of the results into a single object
func (c AssessmentsMetadataClient) SubscriptionListComplete(ctx context.Context, id commonids.SubscriptionId) (SubscriptionListCompleteResult, error) {
	return c.SubscriptionListCompleteMatchingPredicate(ctx, id, SecurityAssessmentMetadataOperationPredicate{})
}

// SubscriptionListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AssessmentsMetadataClient) SubscriptionListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate SecurityAssessmentMetadataOperationPredicate) (resp SubscriptionListCompleteResult, err error) {
	items := make([]SecurityAssessmentMetadata, 0)

	page, err := c.SubscriptionList(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := SubscriptionListCompleteResult{
		Items: items,
	}
	return out, nil
}
