//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// LogicAppsClient contains the methods for the LogicApps group.
// Don't use this type directly, use NewLogicAppsClient() instead.
type LogicAppsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLogicAppsClient creates a new instance of LogicAppsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLogicAppsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LogicAppsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LogicAppsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a Logic App extension resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the Logic App, the extension resource.
//   - resource - Logic app resource properties.
//   - options - LogicAppsClientCreateOrUpdateOptions contains the optional parameters for the LogicAppsClient.CreateOrUpdate
//     method.
func (client *LogicAppsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, resource LogicApp, options *LogicAppsClientCreateOrUpdateOptions) (LogicAppsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "LogicAppsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, resource, options)
	if err != nil {
		return LogicAppsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicAppsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return LogicAppsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LogicAppsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, resource LogicApp, options *LogicAppsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LogicAppsClient) createOrUpdateHandleResponse(resp *http.Response) (LogicAppsClientCreateOrUpdateResponse, error) {
	result := LogicAppsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicApp); err != nil {
		return LogicAppsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Logic App extension resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the Logic App, the extension resource.
//   - options - LogicAppsClientDeleteOptions contains the optional parameters for the LogicAppsClient.Delete method.
func (client *LogicAppsClient) Delete(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientDeleteOptions) (LogicAppsClientDeleteResponse, error) {
	var err error
	const operationName = "LogicAppsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, options)
	if err != nil {
		return LogicAppsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicAppsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return LogicAppsClientDeleteResponse{}, err
	}
	return LogicAppsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LogicAppsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeployWorkflowArtifacts - Creates or updates the artifacts for the logic app
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the Logic App, the extension resource.
//   - options - LogicAppsClientDeployWorkflowArtifactsOptions contains the optional parameters for the LogicAppsClient.DeployWorkflowArtifacts
//     method.
func (client *LogicAppsClient) DeployWorkflowArtifacts(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientDeployWorkflowArtifactsOptions) (LogicAppsClientDeployWorkflowArtifactsResponse, error) {
	var err error
	const operationName = "LogicAppsClient.DeployWorkflowArtifacts"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deployWorkflowArtifactsCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, options)
	if err != nil {
		return LogicAppsClientDeployWorkflowArtifactsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicAppsClientDeployWorkflowArtifactsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogicAppsClientDeployWorkflowArtifactsResponse{}, err
	}
	return LogicAppsClientDeployWorkflowArtifactsResponse{}, nil
}

// deployWorkflowArtifactsCreateRequest creates the DeployWorkflowArtifacts request.
func (client *LogicAppsClient) deployWorkflowArtifactsCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientDeployWorkflowArtifactsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}/deployWorkflowArtifacts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.WorkflowArtifacts != nil {
		if err := runtime.MarshalAsJSON(req, *options.WorkflowArtifacts); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Get - Gets a logic app extension resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the Logic App, the extension resource.
//   - options - LogicAppsClientGetOptions contains the optional parameters for the LogicAppsClient.Get method.
func (client *LogicAppsClient) Get(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientGetOptions) (LogicAppsClientGetResponse, error) {
	var err error
	const operationName = "LogicAppsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, options)
	if err != nil {
		return LogicAppsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicAppsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogicAppsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LogicAppsClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LogicAppsClient) getHandleResponse(resp *http.Response) (LogicAppsClientGetResponse, error) {
	result := LogicAppsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicApp); err != nil {
		return LogicAppsClientGetResponse{}, err
	}
	return result, nil
}

// GetWorkflow - Get workflow information by its name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the Logic App, the extension resource.
//   - workflowName - Workflow name.
//   - options - LogicAppsClientGetWorkflowOptions contains the optional parameters for the LogicAppsClient.GetWorkflow method.
func (client *LogicAppsClient) GetWorkflow(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, workflowName string, options *LogicAppsClientGetWorkflowOptions) (LogicAppsClientGetWorkflowResponse, error) {
	var err error
	const operationName = "LogicAppsClient.GetWorkflow"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getWorkflowCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, workflowName, options)
	if err != nil {
		return LogicAppsClientGetWorkflowResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicAppsClientGetWorkflowResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogicAppsClientGetWorkflowResponse{}, err
	}
	resp, err := client.getWorkflowHandleResponse(httpResp)
	return resp, err
}

// getWorkflowCreateRequest creates the GetWorkflow request.
func (client *LogicAppsClient) getWorkflowCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, workflowName string, options *LogicAppsClientGetWorkflowOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}/workflows/{workflowName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getWorkflowHandleResponse handles the GetWorkflow response.
func (client *LogicAppsClient) getWorkflowHandleResponse(resp *http.Response) (LogicAppsClientGetWorkflowResponse, error) {
	result := LogicAppsClientGetWorkflowResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowEnvelope); err != nil {
		return LogicAppsClientGetWorkflowResponse{}, err
	}
	return result, nil
}

// Invoke - Proxies a the API call to the logic app backed by the container app.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the LogicApp App, the extension resource.
//   - xmsLogicAppsProxyPath - The proxy path for the API call
//   - xmsLogicAppsProxyMethod - The proxy method for the API call
//   - options - LogicAppsClientInvokeOptions contains the optional parameters for the LogicAppsClient.Invoke method.
func (client *LogicAppsClient) Invoke(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, xmsLogicAppsProxyPath string, xmsLogicAppsProxyMethod LogicAppsProxyMethod, options *LogicAppsClientInvokeOptions) (LogicAppsClientInvokeResponse, error) {
	var err error
	const operationName = "LogicAppsClient.Invoke"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.invokeCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, xmsLogicAppsProxyPath, xmsLogicAppsProxyMethod, options)
	if err != nil {
		return LogicAppsClientInvokeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicAppsClientInvokeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogicAppsClientInvokeResponse{}, err
	}
	resp, err := client.invokeHandleResponse(httpResp)
	return resp, err
}

// invokeCreateRequest creates the Invoke request.
func (client *LogicAppsClient) invokeCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, xmsLogicAppsProxyPath string, xmsLogicAppsProxyMethod LogicAppsProxyMethod, options *LogicAppsClientInvokeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}/invoke"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["x-ms-logicApps-proxy-method"] = []string{string(xmsLogicAppsProxyMethod)}
	req.Raw().Header["x-ms-logicApps-proxy-path"] = []string{xmsLogicAppsProxyPath}
	return req, nil
}

// invokeHandleResponse handles the Invoke response.
func (client *LogicAppsClient) invokeHandleResponse(resp *http.Response) (LogicAppsClientInvokeResponse, error) {
	result := LogicAppsClientInvokeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return LogicAppsClientInvokeResponse{}, err
	}
	return result, nil
}

// NewListWorkflowsPager - List the workflows for a logic app.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the Logic App, the extension resource.
//   - options - LogicAppsClientListWorkflowsOptions contains the optional parameters for the LogicAppsClient.NewListWorkflowsPager
//     method.
func (client *LogicAppsClient) NewListWorkflowsPager(resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientListWorkflowsOptions) *runtime.Pager[LogicAppsClientListWorkflowsResponse] {
	return runtime.NewPager(runtime.PagingHandler[LogicAppsClientListWorkflowsResponse]{
		More: func(page LogicAppsClientListWorkflowsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LogicAppsClientListWorkflowsResponse) (LogicAppsClientListWorkflowsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LogicAppsClient.NewListWorkflowsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listWorkflowsCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, options)
			}, nil)
			if err != nil {
				return LogicAppsClientListWorkflowsResponse{}, err
			}
			return client.listWorkflowsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listWorkflowsCreateRequest creates the ListWorkflows request.
func (client *LogicAppsClient) listWorkflowsCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientListWorkflowsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}/workflows"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWorkflowsHandleResponse handles the ListWorkflows response.
func (client *LogicAppsClient) listWorkflowsHandleResponse(resp *http.Response) (LogicAppsClientListWorkflowsResponse, error) {
	result := LogicAppsClientListWorkflowsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowEnvelopeCollection); err != nil {
		return LogicAppsClientListWorkflowsResponse{}, err
	}
	return result, nil
}

// ListWorkflowsConnections - Gets logic app's connections.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerAppName - Name of the Container App.
//   - logicAppName - Name of the Logic App, the extension resource.
//   - options - LogicAppsClientListWorkflowsConnectionsOptions contains the optional parameters for the LogicAppsClient.ListWorkflowsConnections
//     method.
func (client *LogicAppsClient) ListWorkflowsConnections(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientListWorkflowsConnectionsOptions) (LogicAppsClientListWorkflowsConnectionsResponse, error) {
	var err error
	const operationName = "LogicAppsClient.ListWorkflowsConnections"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listWorkflowsConnectionsCreateRequest(ctx, resourceGroupName, containerAppName, logicAppName, options)
	if err != nil {
		return LogicAppsClientListWorkflowsConnectionsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LogicAppsClientListWorkflowsConnectionsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LogicAppsClientListWorkflowsConnectionsResponse{}, err
	}
	resp, err := client.listWorkflowsConnectionsHandleResponse(httpResp)
	return resp, err
}

// listWorkflowsConnectionsCreateRequest creates the ListWorkflowsConnections request.
func (client *LogicAppsClient) listWorkflowsConnectionsCreateRequest(ctx context.Context, resourceGroupName string, containerAppName string, logicAppName string, options *LogicAppsClientListWorkflowsConnectionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{containerAppName}/providers/Microsoft.App/logicApps/{logicAppName}/listWorkflowsConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerAppName == "" {
		return nil, errors.New("parameter containerAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerAppName}", url.PathEscape(containerAppName))
	if logicAppName == "" {
		return nil, errors.New("parameter logicAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{logicAppName}", url.PathEscape(logicAppName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listWorkflowsConnectionsHandleResponse handles the ListWorkflowsConnections response.
func (client *LogicAppsClient) listWorkflowsConnectionsHandleResponse(resp *http.Response) (LogicAppsClientListWorkflowsConnectionsResponse, error) {
	result := LogicAppsClientListWorkflowsConnectionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowEnvelope); err != nil {
		return LogicAppsClientListWorkflowsConnectionsResponse{}, err
	}
	return result, nil
}
