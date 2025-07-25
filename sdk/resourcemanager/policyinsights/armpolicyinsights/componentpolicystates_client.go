// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicyinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ComponentPolicyStatesClient contains the methods for the ComponentPolicyStates group.
// Don't use this type directly, use NewComponentPolicyStatesClient() instead.
type ComponentPolicyStatesClient struct {
	internal *arm.Client
}

// NewComponentPolicyStatesClient creates a new instance of ComponentPolicyStatesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewComponentPolicyStatesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ComponentPolicyStatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ComponentPolicyStatesClient{
		internal: cl,
	}
	return client, nil
}

// ListQueryResultsForPolicyDefinition - Queries component policy states for the subscription level policy definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - subscriptionID - Microsoft Azure subscription ID.
//   - policyDefinitionName - Policy definition name.
//   - componentPolicyStatesResource - The virtual resource under ComponentPolicyStates resource type. In a given time range,
//     'latest' represents the latest component policy state(s).
//   - options - ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionOptions contains the optional parameters for the
//     ComponentPolicyStatesClient.ListQueryResultsForPolicyDefinition method.
func (client *ComponentPolicyStatesClient) ListQueryResultsForPolicyDefinition(ctx context.Context, subscriptionID string, policyDefinitionName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionOptions) (ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse, error) {
	var err error
	const operationName = "ComponentPolicyStatesClient.ListQueryResultsForPolicyDefinition"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listQueryResultsForPolicyDefinitionCreateRequest(ctx, subscriptionID, policyDefinitionName, componentPolicyStatesResource, options)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse{}, err
	}
	resp, err := client.listQueryResultsForPolicyDefinitionHandleResponse(httpResp)
	return resp, err
}

// listQueryResultsForPolicyDefinitionCreateRequest creates the ListQueryResultsForPolicyDefinition request.
func (client *ComponentPolicyStatesClient) listQueryResultsForPolicyDefinitionCreateRequest(ctx context.Context, subscriptionID string, policyDefinitionName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{authorizationNamespace}/policyDefinitions/{policyDefinitionName}/providers/Microsoft.PolicyInsights/componentPolicyStates/{componentPolicyStatesResource}/queryResults"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationNamespace}", url.PathEscape("Microsoft.Authorization"))
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if componentPolicyStatesResource == "" {
		return nil, errors.New("parameter componentPolicyStatesResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentPolicyStatesResource}", url.PathEscape(string(componentPolicyStatesResource)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Apply != nil {
		reqQP.Set("$apply", *options.Apply)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.From != nil {
		reqQP.Set("$from", options.From.Format(time.RFC3339Nano))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderby", *options.OrderBy)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.To != nil {
		reqQP.Set("$to", options.To.Format(time.RFC3339Nano))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQueryResultsForPolicyDefinitionHandleResponse handles the ListQueryResultsForPolicyDefinition response.
func (client *ComponentPolicyStatesClient) listQueryResultsForPolicyDefinitionHandleResponse(resp *http.Response) (ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse, error) {
	result := ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPolicyStatesQueryResults); err != nil {
		return ComponentPolicyStatesClientListQueryResultsForPolicyDefinitionResponse{}, err
	}
	return result, nil
}

// ListQueryResultsForResource - Queries component policy states for the resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceID - Resource ID.
//   - componentPolicyStatesResource - The virtual resource under ComponentPolicyStates resource type. In a given time range,
//     'latest' represents the latest component policy state(s).
//   - options - ComponentPolicyStatesClientListQueryResultsForResourceOptions contains the optional parameters for the ComponentPolicyStatesClient.ListQueryResultsForResource
//     method.
func (client *ComponentPolicyStatesClient) ListQueryResultsForResource(ctx context.Context, resourceID string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForResourceOptions) (ComponentPolicyStatesClientListQueryResultsForResourceResponse, error) {
	var err error
	const operationName = "ComponentPolicyStatesClient.ListQueryResultsForResource"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listQueryResultsForResourceCreateRequest(ctx, resourceID, componentPolicyStatesResource, options)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComponentPolicyStatesClientListQueryResultsForResourceResponse{}, err
	}
	resp, err := client.listQueryResultsForResourceHandleResponse(httpResp)
	return resp, err
}

// listQueryResultsForResourceCreateRequest creates the ListQueryResultsForResource request.
func (client *ComponentPolicyStatesClient) listQueryResultsForResourceCreateRequest(ctx context.Context, resourceID string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceId}/providers/Microsoft.PolicyInsights/componentPolicyStates/{componentPolicyStatesResource}/queryResults"
	urlPath = strings.ReplaceAll(urlPath, "{resourceId}", resourceID)
	if componentPolicyStatesResource == "" {
		return nil, errors.New("parameter componentPolicyStatesResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentPolicyStatesResource}", url.PathEscape(string(componentPolicyStatesResource)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Apply != nil {
		reqQP.Set("$apply", *options.Apply)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.From != nil {
		reqQP.Set("$from", options.From.Format(time.RFC3339Nano))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderby", *options.OrderBy)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.To != nil {
		reqQP.Set("$to", options.To.Format(time.RFC3339Nano))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQueryResultsForResourceHandleResponse handles the ListQueryResultsForResource response.
func (client *ComponentPolicyStatesClient) listQueryResultsForResourceHandleResponse(resp *http.Response) (ComponentPolicyStatesClientListQueryResultsForResourceResponse, error) {
	result := ComponentPolicyStatesClientListQueryResultsForResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPolicyStatesQueryResults); err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceResponse{}, err
	}
	return result, nil
}

// ListQueryResultsForResourceGroup - Queries component policy states under resource group scope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - subscriptionID - Microsoft Azure subscription ID.
//   - resourceGroupName - Resource group name.
//   - componentPolicyStatesResource - The virtual resource under ComponentPolicyStates resource type. In a given time range,
//     'latest' represents the latest component policy state(s).
//   - options - ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions contains the optional parameters for the ComponentPolicyStatesClient.ListQueryResultsForResourceGroup
//     method.
func (client *ComponentPolicyStatesClient) ListQueryResultsForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions) (ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse, error) {
	var err error
	const operationName = "ComponentPolicyStatesClient.ListQueryResultsForResourceGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listQueryResultsForResourceGroupCreateRequest(ctx, subscriptionID, resourceGroupName, componentPolicyStatesResource, options)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse{}, err
	}
	resp, err := client.listQueryResultsForResourceGroupHandleResponse(httpResp)
	return resp, err
}

// listQueryResultsForResourceGroupCreateRequest creates the ListQueryResultsForResourceGroup request.
func (client *ComponentPolicyStatesClient) listQueryResultsForResourceGroupCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PolicyInsights/componentPolicyStates/{componentPolicyStatesResource}/queryResults"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if componentPolicyStatesResource == "" {
		return nil, errors.New("parameter componentPolicyStatesResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentPolicyStatesResource}", url.PathEscape(string(componentPolicyStatesResource)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Apply != nil {
		reqQP.Set("$apply", *options.Apply)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.From != nil {
		reqQP.Set("$from", options.From.Format(time.RFC3339Nano))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderby", *options.OrderBy)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.To != nil {
		reqQP.Set("$to", options.To.Format(time.RFC3339Nano))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQueryResultsForResourceGroupHandleResponse handles the ListQueryResultsForResourceGroup response.
func (client *ComponentPolicyStatesClient) listQueryResultsForResourceGroupHandleResponse(resp *http.Response) (ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse, error) {
	result := ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPolicyStatesQueryResults); err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupResponse{}, err
	}
	return result, nil
}

// ListQueryResultsForResourceGroupLevelPolicyAssignment - Queries component policy states for the resource group level policy
// assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - subscriptionID - Microsoft Azure subscription ID.
//   - resourceGroupName - Resource group name.
//   - policyAssignmentName - Policy assignment name.
//   - componentPolicyStatesResource - The virtual resource under ComponentPolicyStates resource type. In a given time range,
//     'latest' represents the latest component policy state(s).
//   - options - ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions contains the optional
//     parameters for the ComponentPolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment method.
func (client *ComponentPolicyStatesClient) ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions) (ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse, error) {
	var err error
	const operationName = "ComponentPolicyStatesClient.ListQueryResultsForResourceGroupLevelPolicyAssignment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listQueryResultsForResourceGroupLevelPolicyAssignmentCreateRequest(ctx, subscriptionID, resourceGroupName, policyAssignmentName, componentPolicyStatesResource, options)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse{}, err
	}
	resp, err := client.listQueryResultsForResourceGroupLevelPolicyAssignmentHandleResponse(httpResp)
	return resp, err
}

// listQueryResultsForResourceGroupLevelPolicyAssignmentCreateRequest creates the ListQueryResultsForResourceGroupLevelPolicyAssignment request.
func (client *ComponentPolicyStatesClient) listQueryResultsForResourceGroupLevelPolicyAssignmentCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{authorizationNamespace}/policyAssignments/{policyAssignmentName}/providers/Microsoft.PolicyInsights/componentPolicyStates/{componentPolicyStatesResource}/queryResults"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationNamespace}", url.PathEscape("Microsoft.Authorization"))
	if policyAssignmentName == "" {
		return nil, errors.New("parameter policyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyAssignmentName}", url.PathEscape(policyAssignmentName))
	if componentPolicyStatesResource == "" {
		return nil, errors.New("parameter componentPolicyStatesResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentPolicyStatesResource}", url.PathEscape(string(componentPolicyStatesResource)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Apply != nil {
		reqQP.Set("$apply", *options.Apply)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.From != nil {
		reqQP.Set("$from", options.From.Format(time.RFC3339Nano))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderby", *options.OrderBy)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.To != nil {
		reqQP.Set("$to", options.To.Format(time.RFC3339Nano))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQueryResultsForResourceGroupLevelPolicyAssignmentHandleResponse handles the ListQueryResultsForResourceGroupLevelPolicyAssignment response.
func (client *ComponentPolicyStatesClient) listQueryResultsForResourceGroupLevelPolicyAssignmentHandleResponse(resp *http.Response) (ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse, error) {
	result := ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPolicyStatesQueryResults); err != nil {
		return ComponentPolicyStatesClientListQueryResultsForResourceGroupLevelPolicyAssignmentResponse{}, err
	}
	return result, nil
}

// ListQueryResultsForSubscription - Queries component policy states under subscription scope.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - subscriptionID - Microsoft Azure subscription ID.
//   - componentPolicyStatesResource - The virtual resource under ComponentPolicyStates resource type. In a given time range,
//     'latest' represents the latest component policy state(s).
//   - options - ComponentPolicyStatesClientListQueryResultsForSubscriptionOptions contains the optional parameters for the ComponentPolicyStatesClient.ListQueryResultsForSubscription
//     method.
func (client *ComponentPolicyStatesClient) ListQueryResultsForSubscription(ctx context.Context, subscriptionID string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForSubscriptionOptions) (ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse, error) {
	var err error
	const operationName = "ComponentPolicyStatesClient.ListQueryResultsForSubscription"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listQueryResultsForSubscriptionCreateRequest(ctx, subscriptionID, componentPolicyStatesResource, options)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse{}, err
	}
	resp, err := client.listQueryResultsForSubscriptionHandleResponse(httpResp)
	return resp, err
}

// listQueryResultsForSubscriptionCreateRequest creates the ListQueryResultsForSubscription request.
func (client *ComponentPolicyStatesClient) listQueryResultsForSubscriptionCreateRequest(ctx context.Context, subscriptionID string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PolicyInsights/componentPolicyStates/{componentPolicyStatesResource}/queryResults"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if componentPolicyStatesResource == "" {
		return nil, errors.New("parameter componentPolicyStatesResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentPolicyStatesResource}", url.PathEscape(string(componentPolicyStatesResource)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Apply != nil {
		reqQP.Set("$apply", *options.Apply)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.From != nil {
		reqQP.Set("$from", options.From.Format(time.RFC3339Nano))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderby", *options.OrderBy)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.To != nil {
		reqQP.Set("$to", options.To.Format(time.RFC3339Nano))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQueryResultsForSubscriptionHandleResponse handles the ListQueryResultsForSubscription response.
func (client *ComponentPolicyStatesClient) listQueryResultsForSubscriptionHandleResponse(resp *http.Response) (ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse, error) {
	result := ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPolicyStatesQueryResults); err != nil {
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionResponse{}, err
	}
	return result, nil
}

// ListQueryResultsForSubscriptionLevelPolicyAssignment - Queries component policy states for the subscription level policy
// assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - subscriptionID - Microsoft Azure subscription ID.
//   - policyAssignmentName - Policy assignment name.
//   - componentPolicyStatesResource - The virtual resource under ComponentPolicyStates resource type. In a given time range,
//     'latest' represents the latest component policy state(s).
//   - options - ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions contains the optional
//     parameters for the ComponentPolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment method.
func (client *ComponentPolicyStatesClient) ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx context.Context, subscriptionID string, policyAssignmentName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions) (ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse, error) {
	var err error
	const operationName = "ComponentPolicyStatesClient.ListQueryResultsForSubscriptionLevelPolicyAssignment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listQueryResultsForSubscriptionLevelPolicyAssignmentCreateRequest(ctx, subscriptionID, policyAssignmentName, componentPolicyStatesResource, options)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse{}, err
	}
	resp, err := client.listQueryResultsForSubscriptionLevelPolicyAssignmentHandleResponse(httpResp)
	return resp, err
}

// listQueryResultsForSubscriptionLevelPolicyAssignmentCreateRequest creates the ListQueryResultsForSubscriptionLevelPolicyAssignment request.
func (client *ComponentPolicyStatesClient) listQueryResultsForSubscriptionLevelPolicyAssignmentCreateRequest(ctx context.Context, subscriptionID string, policyAssignmentName string, componentPolicyStatesResource ComponentPolicyStatesResource, options *ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{authorizationNamespace}/policyAssignments/{policyAssignmentName}/providers/Microsoft.PolicyInsights/componentPolicyStates/{componentPolicyStatesResource}/queryResults"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{authorizationNamespace}", url.PathEscape("Microsoft.Authorization"))
	if policyAssignmentName == "" {
		return nil, errors.New("parameter policyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyAssignmentName}", url.PathEscape(policyAssignmentName))
	if componentPolicyStatesResource == "" {
		return nil, errors.New("parameter componentPolicyStatesResource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentPolicyStatesResource}", url.PathEscape(string(componentPolicyStatesResource)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Apply != nil {
		reqQP.Set("$apply", *options.Apply)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.From != nil {
		reqQP.Set("$from", options.From.Format(time.RFC3339Nano))
	}
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderby", *options.OrderBy)
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.To != nil {
		reqQP.Set("$to", options.To.Format(time.RFC3339Nano))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listQueryResultsForSubscriptionLevelPolicyAssignmentHandleResponse handles the ListQueryResultsForSubscriptionLevelPolicyAssignment response.
func (client *ComponentPolicyStatesClient) listQueryResultsForSubscriptionLevelPolicyAssignmentHandleResponse(resp *http.Response) (ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse, error) {
	result := ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentPolicyStatesQueryResults); err != nil {
		return ComponentPolicyStatesClientListQueryResultsForSubscriptionLevelPolicyAssignmentResponse{}, err
	}
	return result, nil
}
