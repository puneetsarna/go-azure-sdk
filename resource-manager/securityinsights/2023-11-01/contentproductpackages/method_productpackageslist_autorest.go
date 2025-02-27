package contentproductpackages

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductPackagesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]ProductPackageModel

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ProductPackagesListOperationResponse, error)
}

type ProductPackagesListCompleteResult struct {
	Items []ProductPackageModel
}

func (r ProductPackagesListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ProductPackagesListOperationResponse) LoadMore(ctx context.Context) (resp ProductPackagesListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ProductPackagesListOperationOptions struct {
	Filter  *string
	Orderby *string
	Top     *int64
}

func DefaultProductPackagesListOperationOptions() ProductPackagesListOperationOptions {
	return ProductPackagesListOperationOptions{}
}

func (o ProductPackagesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ProductPackagesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Orderby != nil {
		out["$orderby"] = *o.Orderby
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ProductPackagesList ...
func (c ContentProductPackagesClient) ProductPackagesList(ctx context.Context, id WorkspaceId, options ProductPackagesListOperationOptions) (resp ProductPackagesListOperationResponse, err error) {
	req, err := c.preparerForProductPackagesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackagesList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackagesList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForProductPackagesList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackagesList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForProductPackagesList prepares the ProductPackagesList request.
func (c ContentProductPackagesClient) preparerForProductPackagesList(ctx context.Context, id WorkspaceId, options ProductPackagesListOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/contentProductPackages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForProductPackagesListWithNextLink prepares the ProductPackagesList request with the given nextLink token.
func (c ContentProductPackagesClient) preparerForProductPackagesListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForProductPackagesList handles the response to the ProductPackagesList request. The method always
// closes the http.Response Body.
func (c ContentProductPackagesClient) responderForProductPackagesList(resp *http.Response) (result ProductPackagesListOperationResponse, err error) {
	type page struct {
		Values   []ProductPackageModel `json:"value"`
		NextLink *string               `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ProductPackagesListOperationResponse, err error) {
			req, err := c.preparerForProductPackagesListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackagesList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackagesList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForProductPackagesList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackagesList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ProductPackagesListComplete retrieves all of the results into a single object
func (c ContentProductPackagesClient) ProductPackagesListComplete(ctx context.Context, id WorkspaceId, options ProductPackagesListOperationOptions) (ProductPackagesListCompleteResult, error) {
	return c.ProductPackagesListCompleteMatchingPredicate(ctx, id, options, ProductPackageModelOperationPredicate{})
}

// ProductPackagesListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ContentProductPackagesClient) ProductPackagesListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options ProductPackagesListOperationOptions, predicate ProductPackageModelOperationPredicate) (resp ProductPackagesListCompleteResult, err error) {
	items := make([]ProductPackageModel, 0)

	page, err := c.ProductPackagesList(ctx, id, options)
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

	out := ProductPackagesListCompleteResult{
		Items: items,
	}
	return out, nil
}
