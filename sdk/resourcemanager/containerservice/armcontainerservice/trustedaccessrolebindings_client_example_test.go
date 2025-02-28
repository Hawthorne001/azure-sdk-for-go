//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f2826dd1e08d4625b29867c01a232d2dad423521/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-10-01/examples/TrustedAccessRoleBindings_List.json
func ExampleTrustedAccessRoleBindingsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTrustedAccessRoleBindingsClient().NewListPager("rg1", "clustername1", nil)
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
		// page.TrustedAccessRoleBindingListResult = armcontainerservice.TrustedAccessRoleBindingListResult{
		// 	Value: []*armcontainerservice.TrustedAccessRoleBinding{
		// 		{
		// 			Name: to.Ptr("binding1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/trustedAccessRoleBindings/binding1"),
		// 			Properties: &armcontainerservice.TrustedAccessRoleBindingProperties{
		// 				Roles: []*string{
		// 					to.Ptr("Microsoft.MachineLearningServices/workspaces/reader"),
		// 					to.Ptr("Microsoft.MachineLearningServices/workspaces/writer")},
		// 					SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f2826dd1e08d4625b29867c01a232d2dad423521/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-10-01/examples/TrustedAccessRoleBindings_Get.json
func ExampleTrustedAccessRoleBindingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTrustedAccessRoleBindingsClient().Get(ctx, "rg1", "clustername1", "binding1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TrustedAccessRoleBinding = armcontainerservice.TrustedAccessRoleBinding{
	// 	Name: to.Ptr("binding1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/trustedAccessRoleBindings/binding1"),
	// 	Properties: &armcontainerservice.TrustedAccessRoleBindingProperties{
	// 		Roles: []*string{
	// 			to.Ptr("Microsoft.MachineLearningServices/workspaces/reader"),
	// 			to.Ptr("Microsoft.MachineLearningServices/workspaces/writer")},
	// 			SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f2826dd1e08d4625b29867c01a232d2dad423521/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-10-01/examples/TrustedAccessRoleBindings_CreateOrUpdate.json
func ExampleTrustedAccessRoleBindingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTrustedAccessRoleBindingsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "binding1", armcontainerservice.TrustedAccessRoleBinding{
		Properties: &armcontainerservice.TrustedAccessRoleBindingProperties{
			Roles: []*string{
				to.Ptr("Microsoft.MachineLearningServices/workspaces/reader"),
				to.Ptr("Microsoft.MachineLearningServices/workspaces/writer")},
			SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"),
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
	// res.TrustedAccessRoleBinding = armcontainerservice.TrustedAccessRoleBinding{
	// 	Name: to.Ptr("binding1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/trustedAccessRoleBindings"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/trustedAccessRoleBindings/binding1"),
	// 	Properties: &armcontainerservice.TrustedAccessRoleBindingProperties{
	// 		Roles: []*string{
	// 			to.Ptr("Microsoft.MachineLearningServices/workspaces/reader"),
	// 			to.Ptr("Microsoft.MachineLearningServices/workspaces/writer")},
	// 			SourceResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f2826dd1e08d4625b29867c01a232d2dad423521/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-10-01/examples/TrustedAccessRoleBindings_Delete.json
func ExampleTrustedAccessRoleBindingsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTrustedAccessRoleBindingsClient().BeginDelete(ctx, "rg1", "clustername1", "binding1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
