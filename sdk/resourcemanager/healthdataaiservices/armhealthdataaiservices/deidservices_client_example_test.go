// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhealthdataaiservices_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
	"log"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/DeidServices_Create_MaximumSet_Gen.json
func ExampleDeidServicesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeidServicesClient().BeginCreate(ctx, "rgopenapi", "deidTest", armhealthdataaiservices.DeidService{
		Properties: &armhealthdataaiservices.DeidServiceProperties{
			PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
		},
		Identity: &armhealthdataaiservices.ManagedServiceIdentity{
			Type:                   to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{},
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("qwyhvdwcsjulggagdqxlmazcl"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhealthdataaiservices.DeidServicesClientCreateResponse{
	// 	DeidService: &armhealthdataaiservices.DeidService{
	// 		Properties: &armhealthdataaiservices.DeidServiceProperties{
	// 			ProvisioningState: to.Ptr(armhealthdataaiservices.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
	// 			PrivateEndpointConnections: []*armhealthdataaiservices.PrivateEndpointConnection{
	// 				{
	// 					Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
	// 							ID: to.Ptr("gpnxxbbtsysdhhclm"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
	// 							ActionsRequired: to.Ptr("ulb"),
	// 							Description: to.Ptr("ddxuoved"),
	// 						},
	// 						GroupIDs: []*string{
	// 							to.Ptr("xbdyjqg"),
	// 						},
	// 						ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 					ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih/privateEndpointConnections/mdwvqjtwcjcvrh"),
	// 					Name: to.Ptr("mdwvqjtwcjcvrh"),
	// 					Type: to.Ptr("bzxabjlpbwreez"),
	// 					SystemData: &armhealthdataaiservices.SystemData{
	// 						CreatedBy: to.Ptr("p"),
	// 						CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 						LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			ServiceURL: to.Ptr("woponotsxaippkvhwmibffywnqcfru"),
	// 		},
	// 		Identity: &armhealthdataaiservices.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("a82361f4-5320-4a26-8d1b-45832d2164dd"),
	// 			TenantID: to.Ptr("53a6a686-ae15-4a1d-badf-3e7947918893"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("qwyhvdwcsjulggagdqxlmazcl"),
	// 		ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih"),
	// 		Name: to.Ptr("nlrthrxaukih"),
	// 		Type: to.Ptr("slyfiibvwlhfdpzjynsywhbfauexk"),
	// 		SystemData: &armhealthdataaiservices.SystemData{
	// 			CreatedBy: to.Ptr("p"),
	// 			CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 			LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/DeidServices_Delete_MaximumSet_Gen.json
func ExampleDeidServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeidServicesClient().BeginDelete(ctx, "rgopenapi", "deidTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/DeidServices_Get_MaximumSet_Gen.json
func ExampleDeidServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeidServicesClient().Get(ctx, "rgopenapi", "deidTest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhealthdataaiservices.DeidServicesClientGetResponse{
	// 	DeidService: &armhealthdataaiservices.DeidService{
	// 		Properties: &armhealthdataaiservices.DeidServiceProperties{
	// 			ProvisioningState: to.Ptr(armhealthdataaiservices.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
	// 			PrivateEndpointConnections: []*armhealthdataaiservices.PrivateEndpointConnection{
	// 				{
	// 					Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
	// 							ID: to.Ptr("gpnxxbbtsysdhhclm"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
	// 							ActionsRequired: to.Ptr("ulb"),
	// 							Description: to.Ptr("ddxuoved"),
	// 						},
	// 						GroupIDs: []*string{
	// 							to.Ptr("xbdyjqg"),
	// 						},
	// 						ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 					ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih/privateEndpointConnections/mdwvqjtwcjcvrh"),
	// 					Name: to.Ptr("mdwvqjtwcjcvrh"),
	// 					Type: to.Ptr("bzxabjlpbwreez"),
	// 					SystemData: &armhealthdataaiservices.SystemData{
	// 						CreatedBy: to.Ptr("p"),
	// 						CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 						LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			ServiceURL: to.Ptr("woponotsxaippkvhwmibffywnqcfru"),
	// 		},
	// 		Identity: &armhealthdataaiservices.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("a82361f4-5320-4a26-8d1b-45832d2164dd"),
	// 			TenantID: to.Ptr("53a6a686-ae15-4a1d-badf-3e7947918893"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("qwyhvdwcsjulggagdqxlmazcl"),
	// 		ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih"),
	// 		Name: to.Ptr("nlrthrxaukih"),
	// 		Type: to.Ptr("slyfiibvwlhfdpzjynsywhbfauexk"),
	// 		SystemData: &armhealthdataaiservices.SystemData{
	// 			CreatedBy: to.Ptr("p"),
	// 			CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 			LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/DeidServices_ListByResourceGroup_MaximumSet_Gen.json
func ExampleDeidServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeidServicesClient().NewListByResourceGroupPager("rgopenapi", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armhealthdataaiservices.DeidServicesClientListByResourceGroupResponse{
		// 	DeidServiceListResult: armhealthdataaiservices.DeidServiceListResult{
		// 		Value: []*armhealthdataaiservices.DeidService{
		// 			{
		// 				Properties: &armhealthdataaiservices.DeidServiceProperties{
		// 					ProvisioningState: to.Ptr(armhealthdataaiservices.ProvisioningStateSucceeded),
		// 					ServiceURL: to.Ptr("woponotsxaippkvhwmibffywnqcfru"),
		// 					PrivateEndpointConnections: []*armhealthdataaiservices.PrivateEndpointConnection{
		// 						{
		// 							Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("xbdyjqg"),
		// 								},
		// 								PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
		// 									ID: to.Ptr("gpnxxbbtsysdhhclm"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
		// 									Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
		// 									ActionsRequired: to.Ptr("ulb"),
		// 									Description: to.Ptr("ro"),
		// 								},
		// 								ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 							},
		// 							ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih/privateEndpointConnections/mdwvqjtwcjcvrh"),
		// 							Name: to.Ptr("mdwvqjtwcjcvrh"),
		// 							Type: to.Ptr("bzxabjlpbwreez"),
		// 							SystemData: &armhealthdataaiservices.SystemData{
		// 								CreatedBy: to.Ptr("p"),
		// 								CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
		// 								LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
		// 							},
		// 						},
		// 					},
		// 					PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
		// 				},
		// 				Identity: &armhealthdataaiservices.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("a82361f4-5320-4a26-8d1b-45832d2164dd"),
		// 					TenantID: to.Ptr("53a6a686-ae15-4a1d-badf-3e7947918893"),
		// 					Type: to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("qwyhvdwcsjulggagdqxlmazcl"),
		// 				ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih"),
		// 				Name: to.Ptr("nlrthrxaukih"),
		// 				Type: to.Ptr("slyfiibvwlhfdpzjynsywhbfauexk"),
		// 				SystemData: &armhealthdataaiservices.SystemData{
		// 					CreatedBy: to.Ptr("p"),
		// 					CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
		// 					LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/DeidServices_ListBySubscription_MaximumSet_Gen.json
func ExampleDeidServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeidServicesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armhealthdataaiservices.DeidServicesClientListBySubscriptionResponse{
		// 	DeidServiceListResult: armhealthdataaiservices.DeidServiceListResult{
		// 		Value: []*armhealthdataaiservices.DeidService{
		// 			{
		// 				Properties: &armhealthdataaiservices.DeidServiceProperties{
		// 					ProvisioningState: to.Ptr(armhealthdataaiservices.ProvisioningStateSucceeded),
		// 					ServiceURL: to.Ptr("woponotsxaippkvhwmibffywnqcfru"),
		// 					PrivateEndpointConnections: []*armhealthdataaiservices.PrivateEndpointConnection{
		// 						{
		// 							Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("xbdyjqg"),
		// 								},
		// 								PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
		// 									ID: to.Ptr("gpnxxbbtsysdhhclm"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
		// 									Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
		// 									ActionsRequired: to.Ptr("ulb"),
		// 									Description: to.Ptr("ro"),
		// 								},
		// 								ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 							},
		// 							ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih/privateEndpointConnections/mdwvqjtwcjcvrh"),
		// 							Name: to.Ptr("mdwvqjtwcjcvrh"),
		// 							Type: to.Ptr("bzxabjlpbwreez"),
		// 							SystemData: &armhealthdataaiservices.SystemData{
		// 								CreatedBy: to.Ptr("p"),
		// 								CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
		// 								LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
		// 							},
		// 						},
		// 					},
		// 					PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
		// 				},
		// 				Identity: &armhealthdataaiservices.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("a82361f4-5320-4a26-8d1b-45832d2164dd"),
		// 					TenantID: to.Ptr("53a6a686-ae15-4a1d-badf-3e7947918893"),
		// 					Type: to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("qwyhvdwcsjulggagdqxlmazcl"),
		// 				ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih"),
		// 				Name: to.Ptr("nlrthrxaukih"),
		// 				Type: to.Ptr("slyfiibvwlhfdpzjynsywhbfauexk"),
		// 				SystemData: &armhealthdataaiservices.SystemData{
		// 					CreatedBy: to.Ptr("p"),
		// 					CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
		// 					LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/DeidServices_Update_MaximumSet_Gen.json
func ExampleDeidServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeidServicesClient().BeginUpdate(ctx, "rgopenapi", "deidTest", armhealthdataaiservices.DeidUpdate{
		Identity: &armhealthdataaiservices.ManagedServiceIdentityUpdate{
			Type:                   to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{},
		},
		Tags: map[string]*string{},
		Properties: &armhealthdataaiservices.DeidPropertiesUpdate{
			PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhealthdataaiservices.DeidServicesClientUpdateResponse{
	// 	DeidService: &armhealthdataaiservices.DeidService{
	// 		Properties: &armhealthdataaiservices.DeidServiceProperties{
	// 			ProvisioningState: to.Ptr(armhealthdataaiservices.ProvisioningStateSucceeded),
	// 			PublicNetworkAccess: to.Ptr(armhealthdataaiservices.PublicNetworkAccessEnabled),
	// 			PrivateEndpointConnections: []*armhealthdataaiservices.PrivateEndpointConnection{
	// 				{
	// 					Properties: &armhealthdataaiservices.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armhealthdataaiservices.PrivateEndpoint{
	// 							ID: to.Ptr("gpnxxbbtsysdhhclm"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armhealthdataaiservices.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armhealthdataaiservices.PrivateEndpointServiceConnectionStatusPending),
	// 							ActionsRequired: to.Ptr("ulb"),
	// 							Description: to.Ptr("ddxuoved"),
	// 						},
	// 						GroupIDs: []*string{
	// 							to.Ptr("xbdyjqg"),
	// 						},
	// 						ProvisioningState: to.Ptr(armhealthdataaiservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 					ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih/privateEndpointConnections/mdwvqjtwcjcvrh"),
	// 					Name: to.Ptr("mdwvqjtwcjcvrh"),
	// 					Type: to.Ptr("bzxabjlpbwreez"),
	// 					SystemData: &armhealthdataaiservices.SystemData{
	// 						CreatedBy: to.Ptr("p"),
	// 						CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 						LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			ServiceURL: to.Ptr("woponotsxaippkvhwmibffywnqcfru"),
	// 		},
	// 		Identity: &armhealthdataaiservices.ManagedServiceIdentity{
	// 			Type: to.Ptr(armhealthdataaiservices.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armhealthdataaiservices.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("a82361f4-5320-4a26-8d1b-45832d2164dd"),
	// 			TenantID: to.Ptr("53a6a686-ae15-4a1d-badf-3e7947918893"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("qwyhvdwcsjulggagdqxlmazcl"),
	// 		ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/nlrthrxaukih"),
	// 		Name: to.Ptr("nlrthrxaukih"),
	// 		Type: to.Ptr("slyfiibvwlhfdpzjynsywhbfauexk"),
	// 		SystemData: &armhealthdataaiservices.SystemData{
	// 			CreatedBy: to.Ptr("p"),
	// 			CreatedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.985Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("pmbozfvwrblbknedeb"),
	// 			LastModifiedByType: to.Ptr(armhealthdataaiservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-13T19:47:24.986Z"); return t}()),
	// 		},
	// 	},
	// }
}
