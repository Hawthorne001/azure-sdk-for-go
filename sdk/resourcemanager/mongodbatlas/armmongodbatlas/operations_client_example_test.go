// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armmongodbatlas_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
	"log"
)

// Generated from example definition: 2025-06-01/Operations_List_MaximumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page = armmongodbatlas.OperationsClientListResponse{
		// 	OperationListResult: armmongodbatlas.OperationListResult{
		// 		Value: []*armmongodbatlas.Operation{
		// 			{
		// 				Name: to.Ptr("glurvmlhuxhlxhmjoiycpicdrjmsd"),
		// 				IsDataAction: to.Ptr(true),
		// 				Display: &armmongodbatlas.OperationDisplay{
		// 					Provider: to.Ptr("ybkvitxcbahxljn"),
		// 					Resource: to.Ptr("wsdbqubnkmfbclp"),
		// 					Operation: to.Ptr("vsix"),
		// 					Description: to.Ptr("ichocwjfmgzbiliknwwnasvdsnowu"),
		// 				},
		// 				Origin: to.Ptr(armmongodbatlas.OriginUser),
		// 				ActionType: to.Ptr(armmongodbatlas.ActionTypeInternal),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

// Generated from example definition: 2025-06-01/Operations_List_MinimumSet_Gen.json
func ExampleOperationsClient_NewListPager_operationsListMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page = armmongodbatlas.OperationsClientListResponse{
		// 	OperationListResult: armmongodbatlas.OperationListResult{
		// 	},
		// }
	}
}
