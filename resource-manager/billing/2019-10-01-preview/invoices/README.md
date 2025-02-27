
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/invoices` Documentation

The `invoices` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/invoices"
```


### Client Initialization

```go
client := invoices.NewInvoicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InvoicesClient.DownloadMultipleBillingProfileInvoices`

```go
ctx := context.TODO()
id := invoices.NewBillingProfileID("billingAccountValue", "billingProfileValue")
var payload []string

if err := client.DownloadMultipleBillingProfileInvoicesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoicesClient.DownloadMultipleBillingSubscriptionInvoices`

```go
ctx := context.TODO()
id := invoices.NewBillingSubscriptionID("subscriptionIdValue")
var payload []string

if err := client.DownloadMultipleBillingSubscriptionInvoicesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoicesClient.DownloadMultipleEAInvoices`

```go
ctx := context.TODO()
id := invoices.NewBillingAccountID("billingAccountValue")
var payload []string

if err := client.DownloadMultipleEAInvoicesThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InvoicesClient.Get`

```go
ctx := context.TODO()
id := invoices.NewBillingProfileInvoiceID("billingAccountValue", "billingProfileValue", "invoiceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.GetBillingAccountInvoice`

```go
ctx := context.TODO()
id := invoices.NewInvoiceID("billingAccountValue", "invoiceValue")

read, err := client.GetBillingAccountInvoice(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.GetById`

```go
ctx := context.TODO()
id := invoices.NewBillingSubscriptionInvoiceID("billingAccountValue", "billingSubscriptionValue", "invoiceValue")

read, err := client.GetById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := invoices.NewBillingAccountID("billingAccountValue")

read, err := client.ListByBillingAccount(ctx, id, invoices.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := invoices.NewBillingProfileID("billingAccountValue", "billingProfileValue")

read, err := client.ListByBillingProfile(ctx, id, invoices.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvoicesClient.ListByBillingSubscription`

```go
ctx := context.TODO()
id := invoices.NewBillingAccountBillingSubscriptionID("billingAccountValue", "billingSubscriptionValue")

// alternatively `client.ListByBillingSubscription(ctx, id, invoices.DefaultListByBillingSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingSubscriptionComplete(ctx, id, invoices.DefaultListByBillingSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
