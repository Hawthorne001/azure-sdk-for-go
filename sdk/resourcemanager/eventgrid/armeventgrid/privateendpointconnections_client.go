// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Delete a specific private endpoint connection under a topic, domain, or partner namespace or namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentType - The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or
//     namespace name).
//   - privateEndpointConnectionName - The name of the private endpoint connection connection.
//   - options - PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
//     method.
func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateEndpointConnectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateEndpointConnectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a specific private endpoint connection under a topic, domain, or partner namespace or namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
func (client *PrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, _ *PrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentType == "" {
		return nil, errors.New("parameter parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific private endpoint connection under a topic, domain, or partner namespace or namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentType - The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or
//     namespace name).
//   - privateEndpointConnectionName - The name of the private endpoint connection connection.
//   - options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
//     method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, _ *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentType == "" {
		return nil, errors.New("parameter parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourcePager - Get all private endpoint connections under a topic, domain, or partner namespace or namespace.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentType - The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or
//     namespace name).
//   - options - PrivateEndpointConnectionsClientListByResourceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByResourcePager
//     method.
func (client *PrivateEndpointConnectionsClient) NewListByResourcePager(resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, options *PrivateEndpointConnectionsClientListByResourceOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsClientListByResourceResponse]{
		More: func(page PrivateEndpointConnectionsClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByResourceResponse) (PrivateEndpointConnectionsClientListByResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateEndpointConnectionsClient.NewListByResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceCreateRequest(ctx, resourceGroupName, parentType, parentName, options)
			}, nil)
			if err != nil {
				return PrivateEndpointConnectionsClientListByResourceResponse{}, err
			}
			return client.listByResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *PrivateEndpointConnectionsClient) listByResourceCreateRequest(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, options *PrivateEndpointConnectionsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentType == "" {
		return nil, errors.New("parameter parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *PrivateEndpointConnectionsClient) listByResourceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByResourceResponse, error) {
	result := PrivateEndpointConnectionsClientListByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByResourceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a specific private endpoint connection under a topic, domain or partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentType - The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or
//     namespace name).
//   - privateEndpointConnectionName - The name of the private endpoint connection connection.
//   - privateEndpointConnection - The private endpoint connection object to update.
//   - options - PrivateEndpointConnectionsClientBeginUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginUpdate
//     method.
func (client *PrivateEndpointConnectionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginUpdateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName, privateEndpointConnection, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateEndpointConnectionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateEndpointConnectionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a specific private endpoint connection under a topic, domain or partner namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
func (client *PrivateEndpointConnectionsClient) update(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, parentType, parentName, privateEndpointConnectionName, privateEndpointConnection, options)
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

// updateCreateRequest creates the Update request.
func (client *PrivateEndpointConnectionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, parentType PrivateEndpointConnectionsParentType, parentName string, privateEndpointConnectionName string, privateEndpointConnection PrivateEndpointConnection, _ *PrivateEndpointConnectionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentType == "" {
		return nil, errors.New("parameter parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(string(parentType)))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, privateEndpointConnection); err != nil {
		return nil, err
	}
	return req, nil
}
