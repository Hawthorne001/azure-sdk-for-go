// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhealthdataaiservices_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthdataaiservices/armhealthdataaiservices"
	"log"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/925b1febaaebc3e1d602c765168e8ddabc7153a5/specification/healthdataaiservices/HealthDataAIServices.Management/examples/2024-02-28-preview/PrivateLinks_ListByDeidService_MaximumSet_Gen.json
func ExamplePrivateLinksClient_NewListByDeidServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthdataaiservices.NewClientFactory("F21BB31B-C214-42C0-ACF0-DACCA05D3011", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinksClient().NewListByDeidServicePager("rgopenapi", "deidTest", nil)
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
		// page = armhealthdataaiservices.PrivateLinksClientListByDeidServiceResponse{
		// 	PrivateLinkResourceListResult: armhealthdataaiservices.PrivateLinkResourceListResult{
		// 		Value: []*armhealthdataaiservices.PrivateLinkResource{
		// 			{
		// 				Properties: &armhealthdataaiservices.PrivateLinkResourceProperties{
		// 					GroupID: to.Ptr("flaloakbdcqjisfhfnkrxmt"),
		// 					RequiredMembers: []*string{
		// 						to.Ptr("wbraaugftjqa"),
		// 					},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("ndsmc"),
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/F21BB31B-C214-42C0-ACF0-DACCA05D3011/resourceGroups/rgopenapi/providers/Microsoft.HealthDataAIServices/deidServices/deidTest/privateLinkResources/rmptbzejhgxhlfpdizsekhsq"),
		// 				Name: to.Ptr("rmptbzejhgxhlfpdizsekhsq"),
		// 				Type: to.Ptr("lplbvyjkvsujegzqr"),
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
		// 		NextLink: to.Ptr("https://microsoft.com/arzijt"),
		// 	},
		// }
	}
}
