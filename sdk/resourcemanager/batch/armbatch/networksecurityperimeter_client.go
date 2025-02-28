//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

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

// NetworkSecurityPerimeterClient contains the methods for the NetworkSecurityPerimeter group.
// Don't use this type directly, use NewNetworkSecurityPerimeterClient() instead.
type NetworkSecurityPerimeterClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkSecurityPerimeterClient creates a new instance of NetworkSecurityPerimeterClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkSecurityPerimeterClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkSecurityPerimeterClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkSecurityPerimeterClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetConfiguration - Gets information about the specified NSP configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - networkSecurityPerimeterConfigurationName - The name for Network Security Perimeter configuration
//   - options - NetworkSecurityPerimeterClientGetConfigurationOptions contains the optional parameters for the NetworkSecurityPerimeterClient.GetConfiguration
//     method.
func (client *NetworkSecurityPerimeterClient) GetConfiguration(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterClientGetConfigurationOptions) (NetworkSecurityPerimeterClientGetConfigurationResponse, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterClient.GetConfiguration"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getConfigurationCreateRequest(ctx, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName, options)
	if err != nil {
		return NetworkSecurityPerimeterClientGetConfigurationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkSecurityPerimeterClientGetConfigurationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkSecurityPerimeterClientGetConfigurationResponse{}, err
	}
	resp, err := client.getConfigurationHandleResponse(httpResp)
	return resp, err
}

// getConfigurationCreateRequest creates the GetConfiguration request.
func (client *NetworkSecurityPerimeterClient) getConfigurationCreateRequest(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterClientGetConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/networkSecurityPerimeterConfigurations/{networkSecurityPerimeterConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if networkSecurityPerimeterConfigurationName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterConfigurationName}", url.PathEscape(networkSecurityPerimeterConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getConfigurationHandleResponse handles the GetConfiguration response.
func (client *NetworkSecurityPerimeterClient) getConfigurationHandleResponse(resp *http.Response) (NetworkSecurityPerimeterClientGetConfigurationResponse, error) {
	result := NetworkSecurityPerimeterClientGetConfigurationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfiguration); err != nil {
		return NetworkSecurityPerimeterClientGetConfigurationResponse{}, err
	}
	return result, nil
}

// NewListConfigurationsPager - Lists all of the NSP configurations in the specified account.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - options - NetworkSecurityPerimeterClientListConfigurationsOptions contains the optional parameters for the NetworkSecurityPerimeterClient.NewListConfigurationsPager
//     method.
func (client *NetworkSecurityPerimeterClient) NewListConfigurationsPager(resourceGroupName string, accountName string, options *NetworkSecurityPerimeterClientListConfigurationsOptions) *runtime.Pager[NetworkSecurityPerimeterClientListConfigurationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkSecurityPerimeterClientListConfigurationsResponse]{
		More: func(page NetworkSecurityPerimeterClientListConfigurationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkSecurityPerimeterClientListConfigurationsResponse) (NetworkSecurityPerimeterClientListConfigurationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkSecurityPerimeterClient.NewListConfigurationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listConfigurationsCreateRequest(ctx, resourceGroupName, accountName, options)
			}, nil)
			if err != nil {
				return NetworkSecurityPerimeterClientListConfigurationsResponse{}, err
			}
			return client.listConfigurationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listConfigurationsCreateRequest creates the ListConfigurations request.
func (client *NetworkSecurityPerimeterClient) listConfigurationsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *NetworkSecurityPerimeterClientListConfigurationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/networkSecurityPerimeterConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listConfigurationsHandleResponse handles the ListConfigurations response.
func (client *NetworkSecurityPerimeterClient) listConfigurationsHandleResponse(resp *http.Response) (NetworkSecurityPerimeterClientListConfigurationsResponse, error) {
	result := NetworkSecurityPerimeterClientListConfigurationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfigurationListResult); err != nil {
		return NetworkSecurityPerimeterClientListConfigurationsResponse{}, err
	}
	return result, nil
}

// BeginReconcileConfiguration - Reconciles the specified NSP configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group that contains the Batch account.
//   - accountName - The name of the Batch account.
//   - networkSecurityPerimeterConfigurationName - The name for Network Security Perimeter configuration
//   - options - NetworkSecurityPerimeterClientBeginReconcileConfigurationOptions contains the optional parameters for the NetworkSecurityPerimeterClient.BeginReconcileConfiguration
//     method.
func (client *NetworkSecurityPerimeterClient) BeginReconcileConfiguration(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterClientBeginReconcileConfigurationOptions) (*runtime.Poller[NetworkSecurityPerimeterClientReconcileConfigurationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reconcileConfiguration(ctx, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkSecurityPerimeterClientReconcileConfigurationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkSecurityPerimeterClientReconcileConfigurationResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ReconcileConfiguration - Reconciles the specified NSP configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-07-01
func (client *NetworkSecurityPerimeterClient) reconcileConfiguration(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterClientBeginReconcileConfigurationOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkSecurityPerimeterClient.BeginReconcileConfiguration"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.reconcileConfigurationCreateRequest(ctx, resourceGroupName, accountName, networkSecurityPerimeterConfigurationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// reconcileConfigurationCreateRequest creates the ReconcileConfiguration request.
func (client *NetworkSecurityPerimeterClient) reconcileConfigurationCreateRequest(ctx context.Context, resourceGroupName string, accountName string, networkSecurityPerimeterConfigurationName string, options *NetworkSecurityPerimeterClientBeginReconcileConfigurationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/networkSecurityPerimeterConfigurations/{networkSecurityPerimeterConfigurationName}/reconcile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if networkSecurityPerimeterConfigurationName == "" {
		return nil, errors.New("parameter networkSecurityPerimeterConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityPerimeterConfigurationName}", url.PathEscape(networkSecurityPerimeterConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
