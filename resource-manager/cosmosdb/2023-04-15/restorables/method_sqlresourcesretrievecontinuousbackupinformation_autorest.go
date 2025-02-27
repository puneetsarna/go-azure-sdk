package restorables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlResourcesRetrieveContinuousBackupInformationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *BackupInformation
}

// SqlResourcesRetrieveContinuousBackupInformation ...
func (c RestorablesClient) SqlResourcesRetrieveContinuousBackupInformation(ctx context.Context, id ContainerId, input ContinuousBackupRestoreLocation) (result SqlResourcesRetrieveContinuousBackupInformationOperationResponse, err error) {
	req, err := c.preparerForSqlResourcesRetrieveContinuousBackupInformation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "SqlResourcesRetrieveContinuousBackupInformation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForSqlResourcesRetrieveContinuousBackupInformation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "SqlResourcesRetrieveContinuousBackupInformation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// SqlResourcesRetrieveContinuousBackupInformationThenPoll performs SqlResourcesRetrieveContinuousBackupInformation then polls until it's completed
func (c RestorablesClient) SqlResourcesRetrieveContinuousBackupInformationThenPoll(ctx context.Context, id ContainerId, input ContinuousBackupRestoreLocation) error {
	result, err := c.SqlResourcesRetrieveContinuousBackupInformation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SqlResourcesRetrieveContinuousBackupInformation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after SqlResourcesRetrieveContinuousBackupInformation: %+v", err)
	}

	return nil
}

// preparerForSqlResourcesRetrieveContinuousBackupInformation prepares the SqlResourcesRetrieveContinuousBackupInformation request.
func (c RestorablesClient) preparerForSqlResourcesRetrieveContinuousBackupInformation(ctx context.Context, id ContainerId, input ContinuousBackupRestoreLocation) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/retrieveContinuousBackupInformation", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForSqlResourcesRetrieveContinuousBackupInformation sends the SqlResourcesRetrieveContinuousBackupInformation request. The method will close the
// http.Response Body if it receives an error.
func (c RestorablesClient) senderForSqlResourcesRetrieveContinuousBackupInformation(ctx context.Context, req *http.Request) (future SqlResourcesRetrieveContinuousBackupInformationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
