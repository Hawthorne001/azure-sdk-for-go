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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	"net/http"
	"net/url"
	"regexp"
)

// SystemAssignedIdentitiesServer is a fake server for instances of the armmsi.SystemAssignedIdentitiesClient type.
type SystemAssignedIdentitiesServer struct {
	// GetByScope is the fake for method SystemAssignedIdentitiesClient.GetByScope
	// HTTP status codes to indicate success: http.StatusOK
	GetByScope func(ctx context.Context, scope string, options *armmsi.SystemAssignedIdentitiesClientGetByScopeOptions) (resp azfake.Responder[armmsi.SystemAssignedIdentitiesClientGetByScopeResponse], errResp azfake.ErrorResponder)
}

// NewSystemAssignedIdentitiesServerTransport creates a new instance of SystemAssignedIdentitiesServerTransport with the provided implementation.
// The returned SystemAssignedIdentitiesServerTransport instance is connected to an instance of armmsi.SystemAssignedIdentitiesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSystemAssignedIdentitiesServerTransport(srv *SystemAssignedIdentitiesServer) *SystemAssignedIdentitiesServerTransport {
	return &SystemAssignedIdentitiesServerTransport{srv: srv}
}

// SystemAssignedIdentitiesServerTransport connects instances of armmsi.SystemAssignedIdentitiesClient to instances of SystemAssignedIdentitiesServer.
// Don't use this type directly, use NewSystemAssignedIdentitiesServerTransport instead.
type SystemAssignedIdentitiesServerTransport struct {
	srv *SystemAssignedIdentitiesServer
}

// Do implements the policy.Transporter interface for SystemAssignedIdentitiesServerTransport.
func (s *SystemAssignedIdentitiesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SystemAssignedIdentitiesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if systemAssignedIdentitiesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = systemAssignedIdentitiesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SystemAssignedIdentitiesClient.GetByScope":
				res.resp, res.err = s.dispatchGetByScope(req)
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

func (s *SystemAssignedIdentitiesServerTransport) dispatchGetByScope(req *http.Request) (*http.Response, error) {
	if s.srv.GetByScope == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByScope not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedIdentity/identities/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetByScope(req.Context(), scopeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SystemAssignedIdentity, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SystemAssignedIdentitiesServerTransport
var systemAssignedIdentitiesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
