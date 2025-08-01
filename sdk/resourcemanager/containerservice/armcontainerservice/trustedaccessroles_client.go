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

// TrustedAccessRolesClient contains the methods for the TrustedAccessRoles group.
// Don't use this type directly, use NewTrustedAccessRolesClient() instead.
type TrustedAccessRolesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTrustedAccessRolesClient creates a new instance of TrustedAccessRolesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTrustedAccessRolesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TrustedAccessRolesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TrustedAccessRolesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List supported trusted access roles.
//
// Generated from API version 2025-05-02-preview
//   - location - The name of the Azure region.
//   - options - TrustedAccessRolesClientListOptions contains the optional parameters for the TrustedAccessRolesClient.NewListPager
//     method.
func (client *TrustedAccessRolesClient) NewListPager(location string, options *TrustedAccessRolesClientListOptions) *runtime.Pager[TrustedAccessRolesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TrustedAccessRolesClientListResponse]{
		More: func(page TrustedAccessRolesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TrustedAccessRolesClientListResponse) (TrustedAccessRolesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TrustedAccessRolesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return TrustedAccessRolesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TrustedAccessRolesClient) listCreateRequest(ctx context.Context, location string, _ *TrustedAccessRolesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/trustedAccessRoles"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *TrustedAccessRolesClient) listHandleResponse(resp *http.Response) (TrustedAccessRolesClientListResponse, error) {
	result := TrustedAccessRolesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrustedAccessRoleListResult); err != nil {
		return TrustedAccessRolesClientListResponse{}, err
	}
	return result, nil
}
