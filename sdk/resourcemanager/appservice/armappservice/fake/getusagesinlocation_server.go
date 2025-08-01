// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
	"net/http"
	"net/url"
	"regexp"
)

// GetUsagesInLocationServer is a fake server for instances of the armappservice.GetUsagesInLocationClient type.
type GetUsagesInLocationServer struct {
	// NewListPager is the fake for method GetUsagesInLocationClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, options *armappservice.GetUsagesInLocationClientListOptions) (resp azfake.PagerResponder[armappservice.GetUsagesInLocationClientListResponse])
}

// NewGetUsagesInLocationServerTransport creates a new instance of GetUsagesInLocationServerTransport with the provided implementation.
// The returned GetUsagesInLocationServerTransport instance is connected to an instance of armappservice.GetUsagesInLocationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGetUsagesInLocationServerTransport(srv *GetUsagesInLocationServer) *GetUsagesInLocationServerTransport {
	return &GetUsagesInLocationServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappservice.GetUsagesInLocationClientListResponse]](),
	}
}

// GetUsagesInLocationServerTransport connects instances of armappservice.GetUsagesInLocationClient to instances of GetUsagesInLocationServer.
// Don't use this type directly, use NewGetUsagesInLocationServerTransport instead.
type GetUsagesInLocationServerTransport struct {
	srv          *GetUsagesInLocationServer
	newListPager *tracker[azfake.PagerResponder[armappservice.GetUsagesInLocationClientListResponse]]
}

// Do implements the policy.Transporter interface for GetUsagesInLocationServerTransport.
func (g *GetUsagesInLocationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GetUsagesInLocationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if getUsagesInLocationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = getUsagesInLocationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GetUsagesInLocationClient.NewListPager":
				res.resp, res.err = g.dispatchNewListPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (g *GetUsagesInLocationServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := g.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Web/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListPager(locationParam, nil)
		newListPager = &resp
		g.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappservice.GetUsagesInLocationClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		g.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to GetUsagesInLocationServerTransport
var getUsagesInLocationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
