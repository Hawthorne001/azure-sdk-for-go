// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armstorageactions.ClientFactory type.
type ServerFactory struct {
	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// StorageTaskAssignmentServer contains the fakes for client StorageTaskAssignmentClient
	StorageTaskAssignmentServer StorageTaskAssignmentServer

	// StorageTasksServer contains the fakes for client StorageTasksClient
	StorageTasksServer StorageTasksServer

	// StorageTasksReportServer contains the fakes for client StorageTasksReportClient
	StorageTasksReportServer StorageTasksReportServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armstorageactions.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armstorageactions.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                           *ServerFactory
	trMu                          sync.Mutex
	trOperationsServer            *OperationsServerTransport
	trStorageTaskAssignmentServer *StorageTaskAssignmentServerTransport
	trStorageTasksServer          *StorageTasksServerTransport
	trStorageTasksReportServer    *StorageTasksReportServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "StorageTaskAssignmentClient":
		initServer(s, &s.trStorageTaskAssignmentServer, func() *StorageTaskAssignmentServerTransport {
			return NewStorageTaskAssignmentServerTransport(&s.srv.StorageTaskAssignmentServer)
		})
		resp, err = s.trStorageTaskAssignmentServer.Do(req)
	case "StorageTasksClient":
		initServer(s, &s.trStorageTasksServer, func() *StorageTasksServerTransport { return NewStorageTasksServerTransport(&s.srv.StorageTasksServer) })
		resp, err = s.trStorageTasksServer.Do(req)
	case "StorageTasksReportClient":
		initServer(s, &s.trStorageTasksReportServer, func() *StorageTasksReportServerTransport {
			return NewStorageTasksReportServerTransport(&s.srv.StorageTasksReportServer)
		})
		resp, err = s.trStorageTasksReportServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
