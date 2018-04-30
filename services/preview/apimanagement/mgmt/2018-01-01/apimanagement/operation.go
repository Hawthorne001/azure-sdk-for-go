package apimanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// OperationClient is the apiManagement Client
type OperationClient struct {
	BaseClient
}

// NewOperationClient creates an instance of the OperationClient client.
func NewOperationClient(subscriptionID string) OperationClient {
	return NewOperationClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationClientWithBaseURI creates an instance of the OperationClient client.
func NewOperationClientWithBaseURI(baseURI string, subscriptionID string) OperationClient {
	return OperationClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByTags lists a collection of operations associated with tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API revision identifier. Must be unique in the current API Management service instance. Non-current
// revision has ;rev=n as a suffix where n is the revision number.
// filter - | Field       | Supported operators    | Supported functions                         |
// |-------------|------------------------|---------------------------------------------|
// | id          | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | name        | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | apiName     | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | description | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | method      | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// | urlTemplate | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client OperationClient) ListByTags(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result TagResourceCollectionPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.OperationClient", "ListByTags", err.Error())
	}

	result.fn = client.listByTagsNextResults
	req, err := client.ListByTagsPreparer(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OperationClient", "ListByTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByTagsSender(req)
	if err != nil {
		result.trc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.OperationClient", "ListByTags", resp, "Failure sending request")
		return
	}

	result.trc, err = client.ListByTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OperationClient", "ListByTags", resp, "Failure responding to request")
	}

	return
}

// ListByTagsPreparer prepares the ListByTags request.
func (client OperationClient) ListByTagsPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/operationsByTags", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByTagsSender sends the ListByTags request. The method will close the
// http.Response Body if it receives an error.
func (client OperationClient) ListByTagsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByTagsResponder handles the response to the ListByTags request. The method always
// closes the http.Response Body.
func (client OperationClient) ListByTagsResponder(resp *http.Response) (result TagResourceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByTagsNextResults retrieves the next set of results, if any.
func (client OperationClient) listByTagsNextResults(lastResults TagResourceCollection) (result TagResourceCollection, err error) {
	req, err := lastResults.tagResourceCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.OperationClient", "listByTagsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.OperationClient", "listByTagsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.OperationClient", "listByTagsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByTagsComplete enumerates all values, automatically crossing page boundaries as required.
func (client OperationClient) ListByTagsComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result TagResourceCollectionIterator, err error) {
	result.page, err = client.ListByTags(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
	return
}
