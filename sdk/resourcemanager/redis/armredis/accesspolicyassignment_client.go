//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AccessPolicyAssignmentClient contains the methods for the AccessPolicyAssignment group.
// Don't use this type directly, use NewAccessPolicyAssignmentClient() instead.
type AccessPolicyAssignmentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessPolicyAssignmentClient creates a new instance of AccessPolicyAssignmentClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessPolicyAssignmentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessPolicyAssignmentClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessPolicyAssignmentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateUpdate - Adds the access policy assignment to the specified users
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - accessPolicyAssignmentName - The name of the access policy assignment.
//   - parameters - Parameters supplied to the Create Update Access Policy Assignment operation.
//   - options - AccessPolicyAssignmentClientBeginCreateUpdateOptions contains the optional parameters for the AccessPolicyAssignmentClient.BeginCreateUpdate
//     method.
func (client *AccessPolicyAssignmentClient) BeginCreateUpdate(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, parameters CacheAccessPolicyAssignment, options *AccessPolicyAssignmentClientBeginCreateUpdateOptions) (*runtime.Poller[AccessPolicyAssignmentClientCreateUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createUpdate(ctx, resourceGroupName, cacheName, accessPolicyAssignmentName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccessPolicyAssignmentClientCreateUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccessPolicyAssignmentClientCreateUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateUpdate - Adds the access policy assignment to the specified users
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *AccessPolicyAssignmentClient) createUpdate(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, parameters CacheAccessPolicyAssignment, options *AccessPolicyAssignmentClientBeginCreateUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AccessPolicyAssignmentClient.BeginCreateUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createUpdateCreateRequest(ctx, resourceGroupName, cacheName, accessPolicyAssignmentName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createUpdateCreateRequest creates the CreateUpdate request.
func (client *AccessPolicyAssignmentClient) createUpdateCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, parameters CacheAccessPolicyAssignment, options *AccessPolicyAssignmentClientBeginCreateUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicyAssignments/{accessPolicyAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if accessPolicyAssignmentName == "" {
		return nil, errors.New("parameter accessPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyAssignmentName}", url.PathEscape(accessPolicyAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the access policy assignment from a redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - accessPolicyAssignmentName - The name of the access policy assignment.
//   - options - AccessPolicyAssignmentClientBeginDeleteOptions contains the optional parameters for the AccessPolicyAssignmentClient.BeginDelete
//     method.
func (client *AccessPolicyAssignmentClient) BeginDelete(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, options *AccessPolicyAssignmentClientBeginDeleteOptions) (*runtime.Poller[AccessPolicyAssignmentClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, cacheName, accessPolicyAssignmentName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccessPolicyAssignmentClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccessPolicyAssignmentClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the access policy assignment from a redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
func (client *AccessPolicyAssignmentClient) deleteOperation(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, options *AccessPolicyAssignmentClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AccessPolicyAssignmentClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cacheName, accessPolicyAssignmentName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccessPolicyAssignmentClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, options *AccessPolicyAssignmentClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicyAssignments/{accessPolicyAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if accessPolicyAssignmentName == "" {
		return nil, errors.New("parameter accessPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyAssignmentName}", url.PathEscape(accessPolicyAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the list of assignments for an access policy of a redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - accessPolicyAssignmentName - The name of the access policy assignment.
//   - options - AccessPolicyAssignmentClientGetOptions contains the optional parameters for the AccessPolicyAssignmentClient.Get
//     method.
func (client *AccessPolicyAssignmentClient) Get(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, options *AccessPolicyAssignmentClientGetOptions) (AccessPolicyAssignmentClientGetResponse, error) {
	var err error
	const operationName = "AccessPolicyAssignmentClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, cacheName, accessPolicyAssignmentName, options)
	if err != nil {
		return AccessPolicyAssignmentClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessPolicyAssignmentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessPolicyAssignmentClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AccessPolicyAssignmentClient) getCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyAssignmentName string, options *AccessPolicyAssignmentClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicyAssignments/{accessPolicyAssignmentName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if accessPolicyAssignmentName == "" {
		return nil, errors.New("parameter accessPolicyAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyAssignmentName}", url.PathEscape(accessPolicyAssignmentName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessPolicyAssignmentClient) getHandleResponse(resp *http.Response) (AccessPolicyAssignmentClientGetResponse, error) {
	result := AccessPolicyAssignmentClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CacheAccessPolicyAssignment); err != nil {
		return AccessPolicyAssignmentClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of access policy assignments associated with this redis cache
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - options - AccessPolicyAssignmentClientListOptions contains the optional parameters for the AccessPolicyAssignmentClient.NewListPager
//     method.
func (client *AccessPolicyAssignmentClient) NewListPager(resourceGroupName string, cacheName string, options *AccessPolicyAssignmentClientListOptions) *runtime.Pager[AccessPolicyAssignmentClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessPolicyAssignmentClientListResponse]{
		More: func(page AccessPolicyAssignmentClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessPolicyAssignmentClientListResponse) (AccessPolicyAssignmentClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccessPolicyAssignmentClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, cacheName, options)
			}, nil)
			if err != nil {
				return AccessPolicyAssignmentClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AccessPolicyAssignmentClient) listCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *AccessPolicyAssignmentClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicyAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessPolicyAssignmentClient) listHandleResponse(resp *http.Response) (AccessPolicyAssignmentClientListResponse, error) {
	result := AccessPolicyAssignmentClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CacheAccessPolicyAssignmentList); err != nil {
		return AccessPolicyAssignmentClientListResponse{}, err
	}
	return result, nil
}
