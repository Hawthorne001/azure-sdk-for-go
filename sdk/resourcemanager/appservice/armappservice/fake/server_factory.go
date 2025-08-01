// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armappservice.ClientFactory type.
type ServerFactory struct {
	// CertificateOrdersServer contains the fakes for client CertificateOrdersClient
	CertificateOrdersServer CertificateOrdersServer

	// CertificateOrdersDiagnosticsServer contains the fakes for client CertificateOrdersDiagnosticsClient
	CertificateOrdersDiagnosticsServer CertificateOrdersDiagnosticsServer

	// CertificateRegistrationProviderServer contains the fakes for client CertificateRegistrationProviderClient
	CertificateRegistrationProviderServer CertificateRegistrationProviderServer

	// CertificatesServer contains the fakes for client CertificatesClient
	CertificatesServer CertificatesServer

	// DeletedWebAppsServer contains the fakes for client DeletedWebAppsClient
	DeletedWebAppsServer DeletedWebAppsServer

	// DiagnosticsServer contains the fakes for client DiagnosticsClient
	DiagnosticsServer DiagnosticsServer

	// DomainRegistrationProviderServer contains the fakes for client DomainRegistrationProviderClient
	DomainRegistrationProviderServer DomainRegistrationProviderServer

	// DomainsServer contains the fakes for client DomainsClient
	DomainsServer DomainsServer

	// EnvironmentsServer contains the fakes for client EnvironmentsClient
	EnvironmentsServer EnvironmentsServer

	// GetUsagesInLocationServer contains the fakes for client GetUsagesInLocationClient
	GetUsagesInLocationServer GetUsagesInLocationServer

	// GlobalServer contains the fakes for client GlobalClient
	GlobalServer GlobalServer

	// KubeEnvironmentsServer contains the fakes for client KubeEnvironmentsClient
	KubeEnvironmentsServer KubeEnvironmentsServer

	// PlansServer contains the fakes for client PlansClient
	PlansServer PlansServer

	// ProviderServer contains the fakes for client ProviderClient
	ProviderServer ProviderServer

	// RecommendationsServer contains the fakes for client RecommendationsClient
	RecommendationsServer RecommendationsServer

	// ResourceHealthMetadataServer contains the fakes for client ResourceHealthMetadataClient
	ResourceHealthMetadataServer ResourceHealthMetadataServer

	// SiteCertificatesServer contains the fakes for client SiteCertificatesClient
	SiteCertificatesServer SiteCertificatesServer

	// StaticSitesServer contains the fakes for client StaticSitesClient
	StaticSitesServer StaticSitesServer

	// TopLevelDomainsServer contains the fakes for client TopLevelDomainsClient
	TopLevelDomainsServer TopLevelDomainsServer

	// WebAppsServer contains the fakes for client WebAppsClient
	WebAppsServer WebAppsServer

	// WebSiteManagementServer contains the fakes for client WebSiteManagementClient
	WebSiteManagementServer WebSiteManagementServer

	// WorkflowRunActionRepetitionsServer contains the fakes for client WorkflowRunActionRepetitionsClient
	WorkflowRunActionRepetitionsServer WorkflowRunActionRepetitionsServer

	// WorkflowRunActionRepetitionsRequestHistoriesServer contains the fakes for client WorkflowRunActionRepetitionsRequestHistoriesClient
	WorkflowRunActionRepetitionsRequestHistoriesServer WorkflowRunActionRepetitionsRequestHistoriesServer

	// WorkflowRunActionScopeRepetitionsServer contains the fakes for client WorkflowRunActionScopeRepetitionsClient
	WorkflowRunActionScopeRepetitionsServer WorkflowRunActionScopeRepetitionsServer

	// WorkflowRunActionsServer contains the fakes for client WorkflowRunActionsClient
	WorkflowRunActionsServer WorkflowRunActionsServer

	// WorkflowRunsServer contains the fakes for client WorkflowRunsClient
	WorkflowRunsServer WorkflowRunsServer

	// WorkflowTriggerHistoriesServer contains the fakes for client WorkflowTriggerHistoriesClient
	WorkflowTriggerHistoriesServer WorkflowTriggerHistoriesServer

	// WorkflowTriggersServer contains the fakes for client WorkflowTriggersClient
	WorkflowTriggersServer WorkflowTriggersServer

	// WorkflowVersionsServer contains the fakes for client WorkflowVersionsClient
	WorkflowVersionsServer WorkflowVersionsServer

	// WorkflowsServer contains the fakes for client WorkflowsClient
	WorkflowsServer WorkflowsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armappservice.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armappservice.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                                  *ServerFactory
	trMu                                                 sync.Mutex
	trCertificateOrdersServer                            *CertificateOrdersServerTransport
	trCertificateOrdersDiagnosticsServer                 *CertificateOrdersDiagnosticsServerTransport
	trCertificateRegistrationProviderServer              *CertificateRegistrationProviderServerTransport
	trCertificatesServer                                 *CertificatesServerTransport
	trDeletedWebAppsServer                               *DeletedWebAppsServerTransport
	trDiagnosticsServer                                  *DiagnosticsServerTransport
	trDomainRegistrationProviderServer                   *DomainRegistrationProviderServerTransport
	trDomainsServer                                      *DomainsServerTransport
	trEnvironmentsServer                                 *EnvironmentsServerTransport
	trGetUsagesInLocationServer                          *GetUsagesInLocationServerTransport
	trGlobalServer                                       *GlobalServerTransport
	trKubeEnvironmentsServer                             *KubeEnvironmentsServerTransport
	trPlansServer                                        *PlansServerTransport
	trProviderServer                                     *ProviderServerTransport
	trRecommendationsServer                              *RecommendationsServerTransport
	trResourceHealthMetadataServer                       *ResourceHealthMetadataServerTransport
	trSiteCertificatesServer                             *SiteCertificatesServerTransport
	trStaticSitesServer                                  *StaticSitesServerTransport
	trTopLevelDomainsServer                              *TopLevelDomainsServerTransport
	trWebAppsServer                                      *WebAppsServerTransport
	trWebSiteManagementServer                            *WebSiteManagementServerTransport
	trWorkflowRunActionRepetitionsServer                 *WorkflowRunActionRepetitionsServerTransport
	trWorkflowRunActionRepetitionsRequestHistoriesServer *WorkflowRunActionRepetitionsRequestHistoriesServerTransport
	trWorkflowRunActionScopeRepetitionsServer            *WorkflowRunActionScopeRepetitionsServerTransport
	trWorkflowRunActionsServer                           *WorkflowRunActionsServerTransport
	trWorkflowRunsServer                                 *WorkflowRunsServerTransport
	trWorkflowTriggerHistoriesServer                     *WorkflowTriggerHistoriesServerTransport
	trWorkflowTriggersServer                             *WorkflowTriggersServerTransport
	trWorkflowVersionsServer                             *WorkflowVersionsServerTransport
	trWorkflowsServer                                    *WorkflowsServerTransport
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
	case "CertificateOrdersClient":
		initServer(s, &s.trCertificateOrdersServer, func() *CertificateOrdersServerTransport {
			return NewCertificateOrdersServerTransport(&s.srv.CertificateOrdersServer)
		})
		resp, err = s.trCertificateOrdersServer.Do(req)
	case "CertificateOrdersDiagnosticsClient":
		initServer(s, &s.trCertificateOrdersDiagnosticsServer, func() *CertificateOrdersDiagnosticsServerTransport {
			return NewCertificateOrdersDiagnosticsServerTransport(&s.srv.CertificateOrdersDiagnosticsServer)
		})
		resp, err = s.trCertificateOrdersDiagnosticsServer.Do(req)
	case "CertificateRegistrationProviderClient":
		initServer(s, &s.trCertificateRegistrationProviderServer, func() *CertificateRegistrationProviderServerTransport {
			return NewCertificateRegistrationProviderServerTransport(&s.srv.CertificateRegistrationProviderServer)
		})
		resp, err = s.trCertificateRegistrationProviderServer.Do(req)
	case "CertificatesClient":
		initServer(s, &s.trCertificatesServer, func() *CertificatesServerTransport { return NewCertificatesServerTransport(&s.srv.CertificatesServer) })
		resp, err = s.trCertificatesServer.Do(req)
	case "DeletedWebAppsClient":
		initServer(s, &s.trDeletedWebAppsServer, func() *DeletedWebAppsServerTransport {
			return NewDeletedWebAppsServerTransport(&s.srv.DeletedWebAppsServer)
		})
		resp, err = s.trDeletedWebAppsServer.Do(req)
	case "DiagnosticsClient":
		initServer(s, &s.trDiagnosticsServer, func() *DiagnosticsServerTransport { return NewDiagnosticsServerTransport(&s.srv.DiagnosticsServer) })
		resp, err = s.trDiagnosticsServer.Do(req)
	case "DomainRegistrationProviderClient":
		initServer(s, &s.trDomainRegistrationProviderServer, func() *DomainRegistrationProviderServerTransport {
			return NewDomainRegistrationProviderServerTransport(&s.srv.DomainRegistrationProviderServer)
		})
		resp, err = s.trDomainRegistrationProviderServer.Do(req)
	case "DomainsClient":
		initServer(s, &s.trDomainsServer, func() *DomainsServerTransport { return NewDomainsServerTransport(&s.srv.DomainsServer) })
		resp, err = s.trDomainsServer.Do(req)
	case "EnvironmentsClient":
		initServer(s, &s.trEnvironmentsServer, func() *EnvironmentsServerTransport { return NewEnvironmentsServerTransport(&s.srv.EnvironmentsServer) })
		resp, err = s.trEnvironmentsServer.Do(req)
	case "GetUsagesInLocationClient":
		initServer(s, &s.trGetUsagesInLocationServer, func() *GetUsagesInLocationServerTransport {
			return NewGetUsagesInLocationServerTransport(&s.srv.GetUsagesInLocationServer)
		})
		resp, err = s.trGetUsagesInLocationServer.Do(req)
	case "GlobalClient":
		initServer(s, &s.trGlobalServer, func() *GlobalServerTransport { return NewGlobalServerTransport(&s.srv.GlobalServer) })
		resp, err = s.trGlobalServer.Do(req)
	case "KubeEnvironmentsClient":
		initServer(s, &s.trKubeEnvironmentsServer, func() *KubeEnvironmentsServerTransport {
			return NewKubeEnvironmentsServerTransport(&s.srv.KubeEnvironmentsServer)
		})
		resp, err = s.trKubeEnvironmentsServer.Do(req)
	case "PlansClient":
		initServer(s, &s.trPlansServer, func() *PlansServerTransport { return NewPlansServerTransport(&s.srv.PlansServer) })
		resp, err = s.trPlansServer.Do(req)
	case "ProviderClient":
		initServer(s, &s.trProviderServer, func() *ProviderServerTransport { return NewProviderServerTransport(&s.srv.ProviderServer) })
		resp, err = s.trProviderServer.Do(req)
	case "RecommendationsClient":
		initServer(s, &s.trRecommendationsServer, func() *RecommendationsServerTransport {
			return NewRecommendationsServerTransport(&s.srv.RecommendationsServer)
		})
		resp, err = s.trRecommendationsServer.Do(req)
	case "ResourceHealthMetadataClient":
		initServer(s, &s.trResourceHealthMetadataServer, func() *ResourceHealthMetadataServerTransport {
			return NewResourceHealthMetadataServerTransport(&s.srv.ResourceHealthMetadataServer)
		})
		resp, err = s.trResourceHealthMetadataServer.Do(req)
	case "SiteCertificatesClient":
		initServer(s, &s.trSiteCertificatesServer, func() *SiteCertificatesServerTransport {
			return NewSiteCertificatesServerTransport(&s.srv.SiteCertificatesServer)
		})
		resp, err = s.trSiteCertificatesServer.Do(req)
	case "StaticSitesClient":
		initServer(s, &s.trStaticSitesServer, func() *StaticSitesServerTransport { return NewStaticSitesServerTransport(&s.srv.StaticSitesServer) })
		resp, err = s.trStaticSitesServer.Do(req)
	case "TopLevelDomainsClient":
		initServer(s, &s.trTopLevelDomainsServer, func() *TopLevelDomainsServerTransport {
			return NewTopLevelDomainsServerTransport(&s.srv.TopLevelDomainsServer)
		})
		resp, err = s.trTopLevelDomainsServer.Do(req)
	case "WebAppsClient":
		initServer(s, &s.trWebAppsServer, func() *WebAppsServerTransport { return NewWebAppsServerTransport(&s.srv.WebAppsServer) })
		resp, err = s.trWebAppsServer.Do(req)
	case "WebSiteManagementClient":
		initServer(s, &s.trWebSiteManagementServer, func() *WebSiteManagementServerTransport {
			return NewWebSiteManagementServerTransport(&s.srv.WebSiteManagementServer)
		})
		resp, err = s.trWebSiteManagementServer.Do(req)
	case "WorkflowRunActionRepetitionsClient":
		initServer(s, &s.trWorkflowRunActionRepetitionsServer, func() *WorkflowRunActionRepetitionsServerTransport {
			return NewWorkflowRunActionRepetitionsServerTransport(&s.srv.WorkflowRunActionRepetitionsServer)
		})
		resp, err = s.trWorkflowRunActionRepetitionsServer.Do(req)
	case "WorkflowRunActionRepetitionsRequestHistoriesClient":
		initServer(s, &s.trWorkflowRunActionRepetitionsRequestHistoriesServer, func() *WorkflowRunActionRepetitionsRequestHistoriesServerTransport {
			return NewWorkflowRunActionRepetitionsRequestHistoriesServerTransport(&s.srv.WorkflowRunActionRepetitionsRequestHistoriesServer)
		})
		resp, err = s.trWorkflowRunActionRepetitionsRequestHistoriesServer.Do(req)
	case "WorkflowRunActionScopeRepetitionsClient":
		initServer(s, &s.trWorkflowRunActionScopeRepetitionsServer, func() *WorkflowRunActionScopeRepetitionsServerTransport {
			return NewWorkflowRunActionScopeRepetitionsServerTransport(&s.srv.WorkflowRunActionScopeRepetitionsServer)
		})
		resp, err = s.trWorkflowRunActionScopeRepetitionsServer.Do(req)
	case "WorkflowRunActionsClient":
		initServer(s, &s.trWorkflowRunActionsServer, func() *WorkflowRunActionsServerTransport {
			return NewWorkflowRunActionsServerTransport(&s.srv.WorkflowRunActionsServer)
		})
		resp, err = s.trWorkflowRunActionsServer.Do(req)
	case "WorkflowRunsClient":
		initServer(s, &s.trWorkflowRunsServer, func() *WorkflowRunsServerTransport { return NewWorkflowRunsServerTransport(&s.srv.WorkflowRunsServer) })
		resp, err = s.trWorkflowRunsServer.Do(req)
	case "WorkflowTriggerHistoriesClient":
		initServer(s, &s.trWorkflowTriggerHistoriesServer, func() *WorkflowTriggerHistoriesServerTransport {
			return NewWorkflowTriggerHistoriesServerTransport(&s.srv.WorkflowTriggerHistoriesServer)
		})
		resp, err = s.trWorkflowTriggerHistoriesServer.Do(req)
	case "WorkflowTriggersClient":
		initServer(s, &s.trWorkflowTriggersServer, func() *WorkflowTriggersServerTransport {
			return NewWorkflowTriggersServerTransport(&s.srv.WorkflowTriggersServer)
		})
		resp, err = s.trWorkflowTriggersServer.Do(req)
	case "WorkflowVersionsClient":
		initServer(s, &s.trWorkflowVersionsServer, func() *WorkflowVersionsServerTransport {
			return NewWorkflowVersionsServerTransport(&s.srv.WorkflowVersionsServer)
		})
		resp, err = s.trWorkflowVersionsServer.Do(req)
	case "WorkflowsClient":
		initServer(s, &s.trWorkflowsServer, func() *WorkflowsServerTransport { return NewWorkflowsServerTransport(&s.srv.WorkflowsServer) })
		resp, err = s.trWorkflowsServer.Do(req)
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
