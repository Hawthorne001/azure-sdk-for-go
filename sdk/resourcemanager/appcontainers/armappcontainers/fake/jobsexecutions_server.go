//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
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

// JobsExecutionsServer is a fake server for instances of the armappcontainers.JobsExecutionsClient type.
type JobsExecutionsServer struct {
	// NewListPager is the fake for method JobsExecutionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, jobName string, options *armappcontainers.JobsExecutionsClientListOptions) (resp azfake.PagerResponder[armappcontainers.JobsExecutionsClientListResponse])
}

// NewJobsExecutionsServerTransport creates a new instance of JobsExecutionsServerTransport with the provided implementation.
// The returned JobsExecutionsServerTransport instance is connected to an instance of armappcontainers.JobsExecutionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobsExecutionsServerTransport(srv *JobsExecutionsServer) *JobsExecutionsServerTransport {
	return &JobsExecutionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappcontainers.JobsExecutionsClientListResponse]](),
	}
}

// JobsExecutionsServerTransport connects instances of armappcontainers.JobsExecutionsClient to instances of JobsExecutionsServer.
// Don't use this type directly, use NewJobsExecutionsServerTransport instead.
type JobsExecutionsServerTransport struct {
	srv          *JobsExecutionsServer
	newListPager *tracker[azfake.PagerResponder[armappcontainers.JobsExecutionsClientListResponse]]
}

// Do implements the policy.Transporter interface for JobsExecutionsServerTransport.
func (j *JobsExecutionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "JobsExecutionsClient.NewListPager":
		resp, err = j.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *JobsExecutionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := j.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armappcontainers.JobsExecutionsClientListOptions
		if filterParam != nil {
			options = &armappcontainers.JobsExecutionsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := j.srv.NewListPager(resourceGroupNameParam, jobNameParam, options)
		newListPager = &resp
		j.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappcontainers.JobsExecutionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		j.newListPager.remove(req)
	}
	return resp, nil
}
