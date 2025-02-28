//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// StorageInsightConfigsServer is a fake server for instances of the armoperationalinsights.StorageInsightConfigsClient type.
type StorageInsightConfigsServer struct {
	// CreateOrUpdate is the fake for method StorageInsightConfigsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, parameters armoperationalinsights.StorageInsight, options *armoperationalinsights.StorageInsightConfigsClientCreateOrUpdateOptions) (resp azfake.Responder[armoperationalinsights.StorageInsightConfigsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method StorageInsightConfigsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *armoperationalinsights.StorageInsightConfigsClientDeleteOptions) (resp azfake.Responder[armoperationalinsights.StorageInsightConfigsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method StorageInsightConfigsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *armoperationalinsights.StorageInsightConfigsClientGetOptions) (resp azfake.Responder[armoperationalinsights.StorageInsightConfigsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByWorkspacePager is the fake for method StorageInsightConfigsClient.NewListByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWorkspacePager func(resourceGroupName string, workspaceName string, options *armoperationalinsights.StorageInsightConfigsClientListByWorkspaceOptions) (resp azfake.PagerResponder[armoperationalinsights.StorageInsightConfigsClientListByWorkspaceResponse])
}

// NewStorageInsightConfigsServerTransport creates a new instance of StorageInsightConfigsServerTransport with the provided implementation.
// The returned StorageInsightConfigsServerTransport instance is connected to an instance of armoperationalinsights.StorageInsightConfigsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewStorageInsightConfigsServerTransport(srv *StorageInsightConfigsServer) *StorageInsightConfigsServerTransport {
	return &StorageInsightConfigsServerTransport{
		srv:                     srv,
		newListByWorkspacePager: newTracker[azfake.PagerResponder[armoperationalinsights.StorageInsightConfigsClientListByWorkspaceResponse]](),
	}
}

// StorageInsightConfigsServerTransport connects instances of armoperationalinsights.StorageInsightConfigsClient to instances of StorageInsightConfigsServer.
// Don't use this type directly, use NewStorageInsightConfigsServerTransport instead.
type StorageInsightConfigsServerTransport struct {
	srv                     *StorageInsightConfigsServer
	newListByWorkspacePager *tracker[azfake.PagerResponder[armoperationalinsights.StorageInsightConfigsClientListByWorkspaceResponse]]
}

// Do implements the policy.Transporter interface for StorageInsightConfigsServerTransport.
func (s *StorageInsightConfigsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "StorageInsightConfigsClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "StorageInsightConfigsClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "StorageInsightConfigsClient.Get":
		resp, err = s.dispatchGet(req)
	case "StorageInsightConfigsClient.NewListByWorkspacePager":
		resp, err = s.dispatchNewListByWorkspacePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *StorageInsightConfigsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageInsightConfigs/(?P<storageInsightName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoperationalinsights.StorageInsight](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	storageInsightNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageInsightName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, storageInsightNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StorageInsight, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StorageInsightConfigsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageInsightConfigs/(?P<storageInsightName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	storageInsightNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageInsightName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, storageInsightNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StorageInsightConfigsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageInsightConfigs/(?P<storageInsightName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	storageInsightNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageInsightName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, storageInsightNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StorageInsight, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StorageInsightConfigsServerTransport) dispatchNewListByWorkspacePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWorkspacePager not implemented")}
	}
	newListByWorkspacePager := s.newListByWorkspacePager.get(req)
	if newListByWorkspacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageInsightConfigs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByWorkspacePager(resourceGroupNameParam, workspaceNameParam, nil)
		newListByWorkspacePager = &resp
		s.newListByWorkspacePager.add(req, newListByWorkspacePager)
		server.PagerResponderInjectNextLinks(newListByWorkspacePager, req, func(page *armoperationalinsights.StorageInsightConfigsClientListByWorkspaceResponse, createLink func() string) {
			page.ODataNextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByWorkspacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWorkspacePager) {
		s.newListByWorkspacePager.remove(req)
	}
	return resp, nil
}
