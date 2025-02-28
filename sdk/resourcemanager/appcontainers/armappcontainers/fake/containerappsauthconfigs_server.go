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
	"net/http"
	"net/url"
	"regexp"

	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// ContainerAppsAuthConfigsServer is a fake server for instances of the armappcontainers.ContainerAppsAuthConfigsClient type.
type ContainerAppsAuthConfigsServer struct {
	// CreateOrUpdate is the fake for method ContainerAppsAuthConfigsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, authConfigEnvelope armappcontainers.AuthConfig, options *armappcontainers.ContainerAppsAuthConfigsClientCreateOrUpdateOptions) (resp azfake.Responder[armappcontainers.ContainerAppsAuthConfigsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ContainerAppsAuthConfigsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, options *armappcontainers.ContainerAppsAuthConfigsClientDeleteOptions) (resp azfake.Responder[armappcontainers.ContainerAppsAuthConfigsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ContainerAppsAuthConfigsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, containerAppName string, authConfigName string, options *armappcontainers.ContainerAppsAuthConfigsClientGetOptions) (resp azfake.Responder[armappcontainers.ContainerAppsAuthConfigsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByContainerAppPager is the fake for method ContainerAppsAuthConfigsClient.NewListByContainerAppPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByContainerAppPager func(resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsAuthConfigsClientListByContainerAppOptions) (resp azfake.PagerResponder[armappcontainers.ContainerAppsAuthConfigsClientListByContainerAppResponse])
}

// NewContainerAppsAuthConfigsServerTransport creates a new instance of ContainerAppsAuthConfigsServerTransport with the provided implementation.
// The returned ContainerAppsAuthConfigsServerTransport instance is connected to an instance of armappcontainers.ContainerAppsAuthConfigsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerAppsAuthConfigsServerTransport(srv *ContainerAppsAuthConfigsServer) *ContainerAppsAuthConfigsServerTransport {
	return &ContainerAppsAuthConfigsServerTransport{
		srv:                        srv,
		newListByContainerAppPager: newTracker[azfake.PagerResponder[armappcontainers.ContainerAppsAuthConfigsClientListByContainerAppResponse]](),
	}
}

// ContainerAppsAuthConfigsServerTransport connects instances of armappcontainers.ContainerAppsAuthConfigsClient to instances of ContainerAppsAuthConfigsServer.
// Don't use this type directly, use NewContainerAppsAuthConfigsServerTransport instead.
type ContainerAppsAuthConfigsServerTransport struct {
	srv                        *ContainerAppsAuthConfigsServer
	newListByContainerAppPager *tracker[azfake.PagerResponder[armappcontainers.ContainerAppsAuthConfigsClientListByContainerAppResponse]]
}

// Do implements the policy.Transporter interface for ContainerAppsAuthConfigsServerTransport.
func (c *ContainerAppsAuthConfigsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerAppsAuthConfigsClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ContainerAppsAuthConfigsClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ContainerAppsAuthConfigsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ContainerAppsAuthConfigsClient.NewListByContainerAppPager":
		resp, err = c.dispatchNewListByContainerAppPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerAppsAuthConfigsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authConfigs/(?P<authConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcontainers.AuthConfig](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	authConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authConfigName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, containerAppNameParam, authConfigNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AuthConfig, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsAuthConfigsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authConfigs/(?P<authConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	authConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authConfigName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, containerAppNameParam, authConfigNameParam, nil)
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

func (c *ContainerAppsAuthConfigsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authConfigs/(?P<authConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	authConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authConfigName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, containerAppNameParam, authConfigNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AuthConfig, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsAuthConfigsServerTransport) dispatchNewListByContainerAppPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByContainerAppPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByContainerAppPager not implemented")}
	}
	newListByContainerAppPager := c.newListByContainerAppPager.get(req)
	if newListByContainerAppPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authConfigs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByContainerAppPager(resourceGroupNameParam, containerAppNameParam, nil)
		newListByContainerAppPager = &resp
		c.newListByContainerAppPager.add(req, newListByContainerAppPager)
		server.PagerResponderInjectNextLinks(newListByContainerAppPager, req, func(page *armappcontainers.ContainerAppsAuthConfigsClientListByContainerAppResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByContainerAppPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByContainerAppPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByContainerAppPager) {
		c.newListByContainerAppPager.remove(req)
	}
	return resp, nil
}
