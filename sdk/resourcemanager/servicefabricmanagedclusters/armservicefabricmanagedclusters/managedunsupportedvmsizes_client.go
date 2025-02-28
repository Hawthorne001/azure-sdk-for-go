//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmanagedclusters

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

// ManagedUnsupportedVMSizesClient contains the methods for the ManagedUnsupportedVMSizes group.
// Don't use this type directly, use NewManagedUnsupportedVMSizesClient() instead.
type ManagedUnsupportedVMSizesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedUnsupportedVMSizesClient creates a new instance of ManagedUnsupportedVMSizesClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedUnsupportedVMSizesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedUnsupportedVMSizesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedUnsupportedVMSizesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get unsupported vm size for Service Fabric Managed Clusters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01-preview
//   - location - The location for the cluster code versions. This is different from cluster location.
//   - vmSize - VM Size name.
//   - options - ManagedUnsupportedVMSizesClientGetOptions contains the optional parameters for the ManagedUnsupportedVMSizesClient.Get
//     method.
func (client *ManagedUnsupportedVMSizesClient) Get(ctx context.Context, location string, vmSize string, options *ManagedUnsupportedVMSizesClientGetOptions) (ManagedUnsupportedVMSizesClientGetResponse, error) {
	var err error
	const operationName = "ManagedUnsupportedVMSizesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, vmSize, options)
	if err != nil {
		return ManagedUnsupportedVMSizesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedUnsupportedVMSizesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedUnsupportedVMSizesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedUnsupportedVMSizesClient) getCreateRequest(ctx context.Context, location string, vmSize string, options *ManagedUnsupportedVMSizesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedUnsupportedVMSizes/{vmSize}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmSize == "" {
		return nil, errors.New("parameter vmSize cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmSize}", url.PathEscape(vmSize))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedUnsupportedVMSizesClient) getHandleResponse(resp *http.Response) (ManagedUnsupportedVMSizesClientGetResponse, error) {
	result := ManagedUnsupportedVMSizesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedVMSize); err != nil {
		return ManagedUnsupportedVMSizesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the lists of unsupported vm sizes for Service Fabric Managed Clusters.
//
// Generated from API version 2024-09-01-preview
//   - location - The location for the cluster code versions. This is different from cluster location.
//   - options - ManagedUnsupportedVMSizesClientListOptions contains the optional parameters for the ManagedUnsupportedVMSizesClient.NewListPager
//     method.
func (client *ManagedUnsupportedVMSizesClient) NewListPager(location string, options *ManagedUnsupportedVMSizesClientListOptions) *runtime.Pager[ManagedUnsupportedVMSizesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedUnsupportedVMSizesClientListResponse]{
		More: func(page ManagedUnsupportedVMSizesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedUnsupportedVMSizesClientListResponse) (ManagedUnsupportedVMSizesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedUnsupportedVMSizesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return ManagedUnsupportedVMSizesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ManagedUnsupportedVMSizesClient) listCreateRequest(ctx context.Context, location string, options *ManagedUnsupportedVMSizesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedUnsupportedVMSizes"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ManagedUnsupportedVMSizesClient) listHandleResponse(resp *http.Response) (ManagedUnsupportedVMSizesClientListResponse, error) {
	result := ManagedUnsupportedVMSizesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedVMSizesResult); err != nil {
		return ManagedUnsupportedVMSizesClientListResponse{}, err
	}
	return result, nil
}
