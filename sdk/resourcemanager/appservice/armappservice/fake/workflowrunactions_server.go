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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WorkflowRunActionsServer is a fake server for instances of the armappservice.WorkflowRunActionsClient type.
type WorkflowRunActionsServer struct {
	// Get is the fake for method WorkflowRunActionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, name string, workflowName string, runName string, actionName string, options *armappservice.WorkflowRunActionsClientGetOptions) (resp azfake.Responder[armappservice.WorkflowRunActionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowRunActionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, name string, workflowName string, runName string, options *armappservice.WorkflowRunActionsClientListOptions) (resp azfake.PagerResponder[armappservice.WorkflowRunActionsClientListResponse])

	// NewListExpressionTracesPager is the fake for method WorkflowRunActionsClient.NewListExpressionTracesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListExpressionTracesPager func(resourceGroupName string, name string, workflowName string, runName string, actionName string, options *armappservice.WorkflowRunActionsClientListExpressionTracesOptions) (resp azfake.PagerResponder[armappservice.WorkflowRunActionsClientListExpressionTracesResponse])
}

// NewWorkflowRunActionsServerTransport creates a new instance of WorkflowRunActionsServerTransport with the provided implementation.
// The returned WorkflowRunActionsServerTransport instance is connected to an instance of armappservice.WorkflowRunActionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowRunActionsServerTransport(srv *WorkflowRunActionsServer) *WorkflowRunActionsServerTransport {
	return &WorkflowRunActionsServerTransport{
		srv:                          srv,
		newListPager:                 newTracker[azfake.PagerResponder[armappservice.WorkflowRunActionsClientListResponse]](),
		newListExpressionTracesPager: newTracker[azfake.PagerResponder[armappservice.WorkflowRunActionsClientListExpressionTracesResponse]](),
	}
}

// WorkflowRunActionsServerTransport connects instances of armappservice.WorkflowRunActionsClient to instances of WorkflowRunActionsServer.
// Don't use this type directly, use NewWorkflowRunActionsServerTransport instead.
type WorkflowRunActionsServerTransport struct {
	srv                          *WorkflowRunActionsServer
	newListPager                 *tracker[azfake.PagerResponder[armappservice.WorkflowRunActionsClientListResponse]]
	newListExpressionTracesPager *tracker[azfake.PagerResponder[armappservice.WorkflowRunActionsClientListExpressionTracesResponse]]
}

// Do implements the policy.Transporter interface for WorkflowRunActionsServerTransport.
func (w *WorkflowRunActionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkflowRunActionsClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkflowRunActionsClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkflowRunActionsClient.NewListExpressionTracesPager":
		resp, err = w.dispatchNewListExpressionTracesPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkflowRunActionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
	if err != nil {
		return nil, err
	}
	actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, actionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkflowRunAction, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowRunActionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armappservice.WorkflowRunActionsClientListOptions
		if topParam != nil || filterParam != nil {
			options = &armappservice.WorkflowRunActionsClientListOptions{
				Top:    topParam,
				Filter: filterParam,
			}
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappservice.WorkflowRunActionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WorkflowRunActionsServerTransport) dispatchNewListExpressionTracesPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListExpressionTracesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListExpressionTracesPager not implemented")}
	}
	newListExpressionTracesPager := w.newListExpressionTracesPager.get(req)
	if newListExpressionTracesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/sites/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hostruntime/runtime/webhooks/workflow/api/management/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listExpressionTraces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
		if err != nil {
			return nil, err
		}
		actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListExpressionTracesPager(resourceGroupNameParam, nameParam, workflowNameParam, runNameParam, actionNameParam, nil)
		newListExpressionTracesPager = &resp
		w.newListExpressionTracesPager.add(req, newListExpressionTracesPager)
		server.PagerResponderInjectNextLinks(newListExpressionTracesPager, req, func(page *armappservice.WorkflowRunActionsClientListExpressionTracesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListExpressionTracesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListExpressionTracesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListExpressionTracesPager) {
		w.newListExpressionTracesPager.remove(req)
	}
	return resp, nil
}
