// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// ServiceTasksClient contains the methods for the ServiceTasks group.
// Don't use this type directly, use NewServiceTasksClient() instead.
type ServiceTasksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceTasksClient creates a new instance of ServiceTasksClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceTasksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceTasksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceTasksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - The service tasks resource is a nested, proxy-only resource representing work performed by a DMS (classic) instance.
// This method cancels a service task if it's currently queued or running.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - taskName - Name of the Task
//   - options - ServiceTasksClientCancelOptions contains the optional parameters for the ServiceTasksClient.Cancel method.
func (client *ServiceTasksClient) Cancel(ctx context.Context, groupName string, serviceName string, taskName string, options *ServiceTasksClientCancelOptions) (ServiceTasksClientCancelResponse, error) {
	var err error
	const operationName = "ServiceTasksClient.Cancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, groupName, serviceName, taskName, options)
	if err != nil {
		return ServiceTasksClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceTasksClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceTasksClientCancelResponse{}, err
	}
	resp, err := client.cancelHandleResponse(httpResp)
	return resp, err
}

// cancelCreateRequest creates the Cancel request.
func (client *ServiceTasksClient) cancelCreateRequest(ctx context.Context, groupName string, serviceName string, taskName string, _ *ServiceTasksClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/serviceTasks/{taskName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *ServiceTasksClient) cancelHandleResponse(resp *http.Response) (ServiceTasksClientCancelResponse, error) {
	result := ServiceTasksClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return ServiceTasksClientCancelResponse{}, err
	}
	return result, nil
}

// CreateOrUpdate - The service tasks resource is a nested, proxy-only resource representing work performed by a DMS (classic)
// instance. The PUT method creates a new service task or updates an existing one, although
// since service tasks have no mutable custom properties, there is little reason to update an existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - taskName - Name of the Task
//   - parameters - Information about the task
//   - options - ServiceTasksClientCreateOrUpdateOptions contains the optional parameters for the ServiceTasksClient.CreateOrUpdate
//     method.
func (client *ServiceTasksClient) CreateOrUpdate(ctx context.Context, groupName string, serviceName string, taskName string, parameters ProjectTask, options *ServiceTasksClientCreateOrUpdateOptions) (ServiceTasksClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ServiceTasksClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, groupName, serviceName, taskName, parameters, options)
	if err != nil {
		return ServiceTasksClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceTasksClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ServiceTasksClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceTasksClient) createOrUpdateCreateRequest(ctx context.Context, groupName string, serviceName string, taskName string, parameters ProjectTask, _ *ServiceTasksClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/serviceTasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServiceTasksClient) createOrUpdateHandleResponse(resp *http.Response) (ServiceTasksClientCreateOrUpdateResponse, error) {
	result := ServiceTasksClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return ServiceTasksClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The service tasks resource is a nested, proxy-only resource representing work performed by a DMS (classic) instance.
// The DELETE method deletes a service task, canceling it first if it's running.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - taskName - Name of the Task
//   - options - ServiceTasksClientDeleteOptions contains the optional parameters for the ServiceTasksClient.Delete method.
func (client *ServiceTasksClient) Delete(ctx context.Context, groupName string, serviceName string, taskName string, options *ServiceTasksClientDeleteOptions) (ServiceTasksClientDeleteResponse, error) {
	var err error
	const operationName = "ServiceTasksClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, groupName, serviceName, taskName, options)
	if err != nil {
		return ServiceTasksClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceTasksClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceTasksClientDeleteResponse{}, err
	}
	return ServiceTasksClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceTasksClient) deleteCreateRequest(ctx context.Context, groupName string, serviceName string, taskName string, options *ServiceTasksClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/serviceTasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-15-preview")
	if options != nil && options.DeleteRunningTasks != nil {
		reqQP.Set("deleteRunningTasks", strconv.FormatBool(*options.DeleteRunningTasks))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The service tasks resource is a nested, proxy-only resource representing work performed by a DMS (classic) instance.
// The GET method retrieves information about a service task.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - taskName - Name of the Task
//   - options - ServiceTasksClientGetOptions contains the optional parameters for the ServiceTasksClient.Get method.
func (client *ServiceTasksClient) Get(ctx context.Context, groupName string, serviceName string, taskName string, options *ServiceTasksClientGetOptions) (ServiceTasksClientGetResponse, error) {
	var err error
	const operationName = "ServiceTasksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, groupName, serviceName, taskName, options)
	if err != nil {
		return ServiceTasksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceTasksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceTasksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServiceTasksClient) getCreateRequest(ctx context.Context, groupName string, serviceName string, taskName string, options *ServiceTasksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/serviceTasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceTasksClient) getHandleResponse(resp *http.Response) (ServiceTasksClientGetResponse, error) {
	result := ServiceTasksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return ServiceTasksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The services resource is the top-level resource that represents the Azure Database Migration Service (classic).
// This method returns a list of service level tasks owned by a service resource. Some
// tasks may have a status of Unknown, which indicates that an error occurred while querying the status of that task.
//
// Generated from API version 2025-03-15-preview
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - options - ServiceTasksClientListOptions contains the optional parameters for the ServiceTasksClient.NewListPager method.
func (client *ServiceTasksClient) NewListPager(groupName string, serviceName string, options *ServiceTasksClientListOptions) *runtime.Pager[ServiceTasksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceTasksClientListResponse]{
		More: func(page ServiceTasksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceTasksClientListResponse) (ServiceTasksClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServiceTasksClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, groupName, serviceName, options)
			}, nil)
			if err != nil {
				return ServiceTasksClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ServiceTasksClient) listCreateRequest(ctx context.Context, groupName string, serviceName string, options *ServiceTasksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/serviceTasks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-15-preview")
	if options != nil && options.TaskType != nil {
		reqQP.Set("taskType", *options.TaskType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceTasksClient) listHandleResponse(resp *http.Response) (ServiceTasksClientListResponse, error) {
	result := ServiceTasksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TaskList); err != nil {
		return ServiceTasksClientListResponse{}, err
	}
	return result, nil
}

// Update - The service tasks resource is a nested, proxy-only resource representing work performed by a DMS (classic) instance.
// The PATCH method updates an existing service task, but since service tasks have no
// mutable custom properties, there is little reason to do so.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - groupName - Name of the resource group
//   - serviceName - Name of the service
//   - taskName - Name of the Task
//   - parameters - Information about the task
//   - options - ServiceTasksClientUpdateOptions contains the optional parameters for the ServiceTasksClient.Update method.
func (client *ServiceTasksClient) Update(ctx context.Context, groupName string, serviceName string, taskName string, parameters ProjectTask, options *ServiceTasksClientUpdateOptions) (ServiceTasksClientUpdateResponse, error) {
	var err error
	const operationName = "ServiceTasksClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, groupName, serviceName, taskName, parameters, options)
	if err != nil {
		return ServiceTasksClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceTasksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceTasksClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ServiceTasksClient) updateCreateRequest(ctx context.Context, groupName string, serviceName string, taskName string, parameters ProjectTask, _ *ServiceTasksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/serviceTasks/{taskName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if taskName == "" {
		return nil, errors.New("parameter taskName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{taskName}", url.PathEscape(taskName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ServiceTasksClient) updateHandleResponse(resp *http.Response) (ServiceTasksClientUpdateResponse, error) {
	result := ServiceTasksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProjectTask); err != nil {
		return ServiceTasksClientUpdateResponse{}, err
	}
	return result, nil
}
