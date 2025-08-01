// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// SnapshotsClient contains the methods for the Snapshots group.
// Don't use this type directly, use NewSnapshotsClient() instead.
type SnapshotsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSnapshotsClient creates a new instance of SnapshotsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSnapshotsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SnapshotsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SnapshotsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - parameters - The snapshot to create or update.
//   - options - SnapshotsClientCreateOrUpdateOptions contains the optional parameters for the SnapshotsClient.CreateOrUpdate
//     method.
func (client *SnapshotsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters Snapshot, options *SnapshotsClientCreateOrUpdateOptions) (SnapshotsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "SnapshotsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SnapshotsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SnapshotsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SnapshotsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SnapshotsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters Snapshot, _ *SnapshotsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/snapshots/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SnapshotsClient) createOrUpdateHandleResponse(resp *http.Response) (SnapshotsClientCreateOrUpdateResponse, error) {
	result := SnapshotsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Snapshot); err != nil {
		return SnapshotsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - SnapshotsClientDeleteOptions contains the optional parameters for the SnapshotsClient.Delete method.
func (client *SnapshotsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, options *SnapshotsClientDeleteOptions) (SnapshotsClientDeleteResponse, error) {
	var err error
	const operationName = "SnapshotsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SnapshotsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SnapshotsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SnapshotsClientDeleteResponse{}, err
	}
	return SnapshotsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SnapshotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, _ *SnapshotsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/snapshots/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - SnapshotsClientGetOptions contains the optional parameters for the SnapshotsClient.Get method.
func (client *SnapshotsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *SnapshotsClientGetOptions) (SnapshotsClientGetResponse, error) {
	var err error
	const operationName = "SnapshotsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SnapshotsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SnapshotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, _ *SnapshotsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/snapshots/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SnapshotsClient) getHandleResponse(resp *http.Response) (SnapshotsClientGetResponse, error) {
	result := SnapshotsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Snapshot); err != nil {
		return SnapshotsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of snapshots in the specified subscription.
//
// Generated from API version 2025-05-02-preview
//   - options - SnapshotsClientListOptions contains the optional parameters for the SnapshotsClient.NewListPager method.
func (client *SnapshotsClient) NewListPager(options *SnapshotsClientListOptions) *runtime.Pager[SnapshotsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SnapshotsClientListResponse]{
		More: func(page SnapshotsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SnapshotsClientListResponse) (SnapshotsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SnapshotsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SnapshotsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SnapshotsClient) listCreateRequest(ctx context.Context, _ *SnapshotsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/snapshots"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SnapshotsClient) listHandleResponse(resp *http.Response) (SnapshotsClientListResponse, error) {
	result := SnapshotsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotListResult); err != nil {
		return SnapshotsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists snapshots in the specified subscription and resource group.
//
// Generated from API version 2025-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SnapshotsClientListByResourceGroupOptions contains the optional parameters for the SnapshotsClient.NewListByResourceGroupPager
//     method.
func (client *SnapshotsClient) NewListByResourceGroupPager(resourceGroupName string, options *SnapshotsClientListByResourceGroupOptions) *runtime.Pager[SnapshotsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SnapshotsClientListByResourceGroupResponse]{
		More: func(page SnapshotsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SnapshotsClientListByResourceGroupResponse) (SnapshotsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SnapshotsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SnapshotsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SnapshotsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, _ *SnapshotsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/snapshots"
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
	reqQP.Set("api-version", "2025-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SnapshotsClient) listByResourceGroupHandleResponse(resp *http.Response) (SnapshotsClientListByResourceGroupResponse, error) {
	result := SnapshotsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SnapshotListResult); err != nil {
		return SnapshotsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates tags on a snapshot.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - parameters - Parameters supplied to the Update snapshot Tags operation.
//   - options - SnapshotsClientUpdateTagsOptions contains the optional parameters for the SnapshotsClient.UpdateTags method.
func (client *SnapshotsClient) UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters TagsObject, options *SnapshotsClientUpdateTagsOptions) (SnapshotsClientUpdateTagsResponse, error) {
	var err error
	const operationName = "SnapshotsClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return SnapshotsClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SnapshotsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SnapshotsClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SnapshotsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters TagsObject, _ *SnapshotsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/snapshots/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SnapshotsClient) updateTagsHandleResponse(resp *http.Response) (SnapshotsClientUpdateTagsResponse, error) {
	result := SnapshotsClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Snapshot); err != nil {
		return SnapshotsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
