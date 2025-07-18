//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/TopicTypes_List.json
func ExampleTopicTypesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicTypesClient().NewListPager(nil)
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
		// page.TopicTypesListResult = armeventgrid.TopicTypesListResult{
		// 	Value: []*armeventgrid.TopicTypeInfo{
		// 		{
		// 			Name: to.Ptr("Microsoft.Eventhub.Namespaces"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Eventhub.Namespaces"),
		// 			Properties: &armeventgrid.TopicTypeProperties{
		// 				Description: to.Ptr("Microsoft EventHubs service events."),
		// 				DisplayName: to.Ptr("EventHubs Namespace"),
		// 				Provider: to.Ptr("Microsoft.Eventhub"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicTypeProvisioningStateSucceeded),
		// 				ResourceRegionType: to.Ptr(armeventgrid.ResourceRegionTypeRegionalResource),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.StorageAccounts"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts"),
		// 			Properties: &armeventgrid.TopicTypeProperties{
		// 				Description: to.Ptr("Microsoft Storage service events."),
		// 				DisplayName: to.Ptr("Storage Accounts"),
		// 				Provider: to.Ptr("Microsoft.Storage"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicTypeProvisioningStateSucceeded),
		// 				ResourceRegionType: to.Ptr(armeventgrid.ResourceRegionTypeRegionalResource),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/TopicTypes_Get.json
func ExampleTopicTypesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicTypesClient().Get(ctx, "Microsoft.Storage.StorageAccounts", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TopicTypeInfo = armeventgrid.TopicTypeInfo{
	// 	Name: to.Ptr("Microsoft.Storage.StorageAccounts"),
	// 	Type: to.Ptr("Microsoft.EventGrid/topicTypes"),
	// 	ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts"),
	// 	Properties: &armeventgrid.TopicTypeProperties{
	// 		Description: to.Ptr("Microsoft Storage service events."),
	// 		DisplayName: to.Ptr("Storage Accounts"),
	// 		Provider: to.Ptr("Microsoft.Storage"),
	// 		ProvisioningState: to.Ptr(armeventgrid.TopicTypeProvisioningStateSucceeded),
	// 		ResourceRegionType: to.Ptr(armeventgrid.ResourceRegionTypeRegionalResource),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/TopicTypes_ListEventTypes.json
func ExampleTopicTypesClient_NewListEventTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicTypesClient().NewListEventTypesPager("Microsoft.Storage.StorageAccounts", nil)
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
		// page.EventTypesListResult = armeventgrid.EventTypesListResult{
		// 	Value: []*armeventgrid.EventType{
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.BlobCreated"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes/eventTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts/eventTypes/Microsoft.Storage.BlobCreated"),
		// 			Properties: &armeventgrid.EventTypeProperties{
		// 				Description: to.Ptr("Raised when a blob is created."),
		// 				DisplayName: to.Ptr("Blob Created"),
		// 				SchemaURL: to.Ptr("tbd"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Storage.BlobDeleted"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topicTypes/eventTypes"),
		// 			ID: to.Ptr("providers/Microsoft.EventGrid/topicTypes/Microsoft.Storage.StorageAccounts/eventTypes/Microsoft.Storage.BlobDeleted"),
		// 			Properties: &armeventgrid.EventTypeProperties{
		// 				Description: to.Ptr("Raised when a blob is deleted."),
		// 				DisplayName: to.Ptr("Blob Deleted"),
		// 				SchemaURL: to.Ptr("tbd"),
		// 			},
		// 	}},
		// }
	}
}
