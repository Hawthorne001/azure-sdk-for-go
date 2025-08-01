//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/AuthorizedApplications_Get.json
func ExampleAuthorizedApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAuthorizedApplicationsClient().Get(ctx, "Microsoft.Contoso", "760505bf-dcfa-4311-b890-18da392a00b2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizedApplication = armproviderhub.AuthorizedApplication{
	// 	Name: to.Ptr("Microsoft.Contoso/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/authorizedApplications"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/authorizedApplications/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 	SystemData: &armproviderhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armproviderhub.AuthorizedApplicationProperties{
	// 		DataAuthorizations: []*armproviderhub.ApplicationDataAuthorization{
	// 			{
	// 				ResourceTypes: []*string{
	// 					to.Ptr("*")},
	// 					Role: to.Ptr(armproviderhub.RoleServiceOwner),
	// 			}},
	// 			ProviderAuthorization: &armproviderhub.ApplicationProviderAuthorization{
	// 				ManagedByRoleDefinitionID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 				RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 			},
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/AuthorizedApplications_CreateOrUpdate.json
func ExampleAuthorizedApplicationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAuthorizedApplicationsClient().BeginCreateOrUpdate(ctx, "Microsoft.Contoso", "760505bf-dcfa-4311-b890-18da392a00b2", armproviderhub.AuthorizedApplication{
		Properties: &armproviderhub.AuthorizedApplicationProperties{
			DataAuthorizations: []*armproviderhub.ApplicationDataAuthorization{
				{
					ResourceTypes: []*string{
						to.Ptr("*")},
					Role: to.Ptr(armproviderhub.RoleServiceOwner),
				}},
			ProviderAuthorization: &armproviderhub.ApplicationProviderAuthorization{
				ManagedByRoleDefinitionID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
				RoleDefinitionID:          to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
			},
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
	// res.AuthorizedApplication = armproviderhub.AuthorizedApplication{
	// 	Name: to.Ptr("Microsoft.Contoso/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 	Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/authorizedApplications"),
	// 	ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/authorizedApplications/760505bf-dcfa-4311-b890-18da392a00b2"),
	// 	SystemData: &armproviderhub.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
	// 	},
	// 	Properties: &armproviderhub.AuthorizedApplicationProperties{
	// 		DataAuthorizations: []*armproviderhub.ApplicationDataAuthorization{
	// 			{
	// 				ResourceTypes: []*string{
	// 					to.Ptr("*")},
	// 					Role: to.Ptr(armproviderhub.RoleServiceOwner),
	// 			}},
	// 			ProviderAuthorization: &armproviderhub.ApplicationProviderAuthorization{
	// 				ManagedByRoleDefinitionID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
	// 				RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
	// 			},
	// 			ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/AuthorizedApplications_Delete.json
func ExampleAuthorizedApplicationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAuthorizedApplicationsClient().Delete(ctx, "Microsoft.Contoso", "760505bf-dcfa-4311-b890-18da392a00b2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7611bb6c9bad11244f4351eecfc50b2c46a86fde/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/AuthorizedApplications_List.json
func ExampleAuthorizedApplicationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAuthorizedApplicationsClient().NewListPager("Microsoft.Contoso", nil)
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
		// page.AuthorizedApplicationArrayResponseWithContinuation = armproviderhub.AuthorizedApplicationArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.AuthorizedApplication{
		// 		{
		// 			Name: to.Ptr("Microsoft.Contoso/760505bf-dcfa-4311-b890-18da392a00b2"),
		// 			Type: to.Ptr("Microsoft.ProviderHub/providerRegistrations/authorizedApplications"),
		// 			ID: to.Ptr("/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.ProviderHub/providerRegistrations/Microsoft.Contoso/authorizedApplications/760505bf-dcfa-4311-b890-18da392a00b2"),
		// 			SystemData: &armproviderhub.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armproviderhub.CreatedByTypeUser),
		// 			},
		// 			Properties: &armproviderhub.AuthorizedApplicationProperties{
		// 				DataAuthorizations: []*armproviderhub.ApplicationDataAuthorization{
		// 					{
		// 						ResourceTypes: []*string{
		// 							to.Ptr("*")},
		// 							Role: to.Ptr(armproviderhub.RoleServiceOwner),
		// 					}},
		// 					ProviderAuthorization: &armproviderhub.ApplicationProviderAuthorization{
		// 						ManagedByRoleDefinitionID: to.Ptr("1a3b5c7d-8e9f-10g1-1h12-i13j14k1"),
		// 						RoleDefinitionID: to.Ptr("123456bf-gkur-2098-b890-98da392a00b2"),
		// 					},
		// 					ProvisioningState: to.Ptr(armproviderhub.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
