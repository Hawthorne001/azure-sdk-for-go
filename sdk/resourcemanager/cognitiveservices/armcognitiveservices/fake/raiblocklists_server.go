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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"net/http"
	"net/url"
	"regexp"
)

// RaiBlocklistsServer is a fake server for instances of the armcognitiveservices.RaiBlocklistsClient type.
type RaiBlocklistsServer struct {
	// CreateOrUpdate is the fake for method RaiBlocklistsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, raiBlocklist armcognitiveservices.RaiBlocklist, options *armcognitiveservices.RaiBlocklistsClientCreateOrUpdateOptions) (resp azfake.Responder[armcognitiveservices.RaiBlocklistsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RaiBlocklistsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, options *armcognitiveservices.RaiBlocklistsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcognitiveservices.RaiBlocklistsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RaiBlocklistsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, raiBlocklistName string, options *armcognitiveservices.RaiBlocklistsClientGetOptions) (resp azfake.Responder[armcognitiveservices.RaiBlocklistsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RaiBlocklistsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, options *armcognitiveservices.RaiBlocklistsClientListOptions) (resp azfake.PagerResponder[armcognitiveservices.RaiBlocklistsClientListResponse])
}

// NewRaiBlocklistsServerTransport creates a new instance of RaiBlocklistsServerTransport with the provided implementation.
// The returned RaiBlocklistsServerTransport instance is connected to an instance of armcognitiveservices.RaiBlocklistsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRaiBlocklistsServerTransport(srv *RaiBlocklistsServer) *RaiBlocklistsServerTransport {
	return &RaiBlocklistsServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armcognitiveservices.RaiBlocklistsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armcognitiveservices.RaiBlocklistsClientListResponse]](),
	}
}

// RaiBlocklistsServerTransport connects instances of armcognitiveservices.RaiBlocklistsClient to instances of RaiBlocklistsServer.
// Don't use this type directly, use NewRaiBlocklistsServerTransport instead.
type RaiBlocklistsServerTransport struct {
	srv          *RaiBlocklistsServer
	beginDelete  *tracker[azfake.PollerResponder[armcognitiveservices.RaiBlocklistsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armcognitiveservices.RaiBlocklistsClientListResponse]]
}

// Do implements the policy.Transporter interface for RaiBlocklistsServerTransport.
func (r *RaiBlocklistsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RaiBlocklistsClient.CreateOrUpdate":
		resp, err = r.dispatchCreateOrUpdate(req)
	case "RaiBlocklistsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RaiBlocklistsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RaiBlocklistsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RaiBlocklistsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiBlocklists/(?P<raiBlocklistName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcognitiveservices.RaiBlocklist](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	raiBlocklistNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("raiBlocklistName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, accountNameParam, raiBlocklistNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RaiBlocklist, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RaiBlocklistsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiBlocklists/(?P<raiBlocklistName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		raiBlocklistNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("raiBlocklistName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, raiBlocklistNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RaiBlocklistsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiBlocklists/(?P<raiBlocklistName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	raiBlocklistNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("raiBlocklistName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, raiBlocklistNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RaiBlocklist, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RaiBlocklistsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/raiBlocklists`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, accountNameParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcognitiveservices.RaiBlocklistsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}
