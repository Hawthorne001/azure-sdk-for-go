//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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

// ResourceSKUsClient contains the methods for the ResourceSKUs group.
// Don't use this type directly, use NewResourceSKUsClient() instead.
type ResourceSKUsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceSKUsClient creates a new instance of ResourceSKUsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceSKUsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceSKUsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets the list of Microsoft.CognitiveServices SKUs available for your Subscription.
//
// Generated from API version 2024-10-01
//   - options - ResourceSKUsClientListOptions contains the optional parameters for the ResourceSKUsClient.NewListPager method.
func (client *ResourceSKUsClient) NewListPager(options *ResourceSKUsClientListOptions) *runtime.Pager[ResourceSKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceSKUsClientListResponse]{
		More: func(page ResourceSKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceSKUsClientListResponse) (ResourceSKUsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ResourceSKUsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ResourceSKUsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ResourceSKUsClient) listCreateRequest(ctx context.Context, options *ResourceSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceSKUsClient) listHandleResponse(resp *http.Response) (ResourceSKUsClientListResponse, error) {
	result := ResourceSKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceSKUListResult); err != nil {
		return ResourceSKUsClientListResponse{}, err
	}
	return result, nil
}
