//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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
)

// HostPoolsClient contains the methods for the HostPools group.
// Don't use this type directly, use NewHostPoolsClient() instead.
type HostPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHostPoolsClient creates a new instance of HostPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHostPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HostPoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HostPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a host pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - hostPool - Object containing HostPool definitions.
//   - options - HostPoolsClientCreateOrUpdateOptions contains the optional parameters for the HostPoolsClient.CreateOrUpdate
//     method.
func (client *HostPoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hostPoolName string, hostPool HostPool, options *HostPoolsClientCreateOrUpdateOptions) (HostPoolsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "HostPoolsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hostPoolName, hostPool, options)
	if err != nil {
		return HostPoolsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HostPoolsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return HostPoolsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HostPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, hostPool HostPool, options *HostPoolsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, hostPool); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HostPoolsClient) createOrUpdateHandleResponse(resp *http.Response) (HostPoolsClientCreateOrUpdateResponse, error) {
	result := HostPoolsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPool); err != nil {
		return HostPoolsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Remove a host pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - HostPoolsClientDeleteOptions contains the optional parameters for the HostPoolsClient.Delete method.
func (client *HostPoolsClient) Delete(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientDeleteOptions) (HostPoolsClientDeleteResponse, error) {
	var err error
	const operationName = "HostPoolsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HostPoolsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HostPoolsClientDeleteResponse{}, err
	}
	return HostPoolsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HostPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a host pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - HostPoolsClientGetOptions contains the optional parameters for the HostPoolsClient.Get method.
func (client *HostPoolsClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientGetOptions) (HostPoolsClientGetResponse, error) {
	var err error
	const operationName = "HostPoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HostPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HostPoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HostPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HostPoolsClient) getHandleResponse(resp *http.Response) (HostPoolsClientGetResponse, error) {
	result := HostPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPool); err != nil {
		return HostPoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List hostPools in subscription.
//
// Generated from API version 2024-04-03
//   - options - HostPoolsClientListOptions contains the optional parameters for the HostPoolsClient.NewListPager method.
func (client *HostPoolsClient) NewListPager(options *HostPoolsClientListOptions) *runtime.Pager[HostPoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HostPoolsClientListResponse]{
		More: func(page HostPoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HostPoolsClientListResponse) (HostPoolsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HostPoolsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return HostPoolsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *HostPoolsClient) listCreateRequest(ctx context.Context, options *HostPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DesktopVirtualization/hostPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HostPoolsClient) listHandleResponse(resp *http.Response) (HostPoolsClientListResponse, error) {
	result := HostPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPoolList); err != nil {
		return HostPoolsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List hostPools.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - HostPoolsClientListByResourceGroupOptions contains the optional parameters for the HostPoolsClient.NewListByResourceGroupPager
//     method.
func (client *HostPoolsClient) NewListByResourceGroupPager(resourceGroupName string, options *HostPoolsClientListByResourceGroupOptions) *runtime.Pager[HostPoolsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[HostPoolsClientListByResourceGroupResponse]{
		More: func(page HostPoolsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HostPoolsClientListByResourceGroupResponse) (HostPoolsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HostPoolsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return HostPoolsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *HostPoolsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *HostPoolsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *HostPoolsClient) listByResourceGroupHandleResponse(resp *http.Response) (HostPoolsClientListByResourceGroupResponse, error) {
	result := HostPoolsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPoolList); err != nil {
		return HostPoolsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListRegistrationTokens - Operation to list the RegistrationTokens associated with the HostPool
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - HostPoolsClientListRegistrationTokensOptions contains the optional parameters for the HostPoolsClient.ListRegistrationTokens
//     method.
func (client *HostPoolsClient) ListRegistrationTokens(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientListRegistrationTokensOptions) (HostPoolsClientListRegistrationTokensResponse, error) {
	var err error
	const operationName = "HostPoolsClient.ListRegistrationTokens"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listRegistrationTokensCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsClientListRegistrationTokensResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HostPoolsClientListRegistrationTokensResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HostPoolsClientListRegistrationTokensResponse{}, err
	}
	resp, err := client.listRegistrationTokensHandleResponse(httpResp)
	return resp, err
}

// listRegistrationTokensCreateRequest creates the ListRegistrationTokens request.
func (client *HostPoolsClient) listRegistrationTokensCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientListRegistrationTokensOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/listRegistrationTokens"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRegistrationTokensHandleResponse handles the ListRegistrationTokens response.
func (client *HostPoolsClient) listRegistrationTokensHandleResponse(resp *http.Response) (HostPoolsClientListRegistrationTokensResponse, error) {
	result := HostPoolsClientListRegistrationTokensResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationTokenList); err != nil {
		return HostPoolsClientListRegistrationTokensResponse{}, err
	}
	return result, nil
}

// RetrieveRegistrationToken - Registration token of the host pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - HostPoolsClientRetrieveRegistrationTokenOptions contains the optional parameters for the HostPoolsClient.RetrieveRegistrationToken
//     method.
func (client *HostPoolsClient) RetrieveRegistrationToken(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientRetrieveRegistrationTokenOptions) (HostPoolsClientRetrieveRegistrationTokenResponse, error) {
	var err error
	const operationName = "HostPoolsClient.RetrieveRegistrationToken"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.retrieveRegistrationTokenCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsClientRetrieveRegistrationTokenResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HostPoolsClientRetrieveRegistrationTokenResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HostPoolsClientRetrieveRegistrationTokenResponse{}, err
	}
	resp, err := client.retrieveRegistrationTokenHandleResponse(httpResp)
	return resp, err
}

// retrieveRegistrationTokenCreateRequest creates the RetrieveRegistrationToken request.
func (client *HostPoolsClient) retrieveRegistrationTokenCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientRetrieveRegistrationTokenOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/retrieveRegistrationToken"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// retrieveRegistrationTokenHandleResponse handles the RetrieveRegistrationToken response.
func (client *HostPoolsClient) retrieveRegistrationTokenHandleResponse(resp *http.Response) (HostPoolsClientRetrieveRegistrationTokenResponse, error) {
	result := HostPoolsClientRetrieveRegistrationTokenResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationInfo); err != nil {
		return HostPoolsClientRetrieveRegistrationTokenResponse{}, err
	}
	return result, nil
}

// Update - Update a host pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - HostPoolsClientUpdateOptions contains the optional parameters for the HostPoolsClient.Update method.
func (client *HostPoolsClient) Update(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientUpdateOptions) (HostPoolsClientUpdateResponse, error) {
	var err error
	const operationName = "HostPoolsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return HostPoolsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HostPoolsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HostPoolsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *HostPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *HostPoolsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.HostPool != nil {
		if err := runtime.MarshalAsJSON(req, *options.HostPool); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *HostPoolsClient) updateHandleResponse(resp *http.Response) (HostPoolsClientUpdateResponse, error) {
	result := HostPoolsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HostPool); err != nil {
		return HostPoolsClientUpdateResponse{}, err
	}
	return result, nil
}
