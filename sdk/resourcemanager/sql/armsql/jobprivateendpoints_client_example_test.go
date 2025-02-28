//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ListJobPrivateEndpointsByAgent.json
func ExampleJobPrivateEndpointsClient_NewListByAgentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobPrivateEndpointsClient().NewListByAgentPager("group1", "server1", "agent1", nil)
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
		// page.JobPrivateEndpointListResult = armsql.JobPrivateEndpointListResult{
		// 	Value: []*armsql.JobPrivateEndpoint{
		// 		{
		// 			Name: to.Ptr("endpoint1"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/privateEndpoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/privateEndpoints/endpoint1"),
		// 			Properties: &armsql.JobPrivateEndpointProperties{
		// 				PrivateEndpointID: to.Ptr("EJ_47e33188-85ff-4705-8d78-444444444444_endpoint1"),
		// 				TargetServerAzureResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("endpoint2"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/jobAgents/privateEndpoints"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/privateEndpoints/endpoint2"),
		// 			Properties: &armsql.JobPrivateEndpointProperties{
		// 				PrivateEndpointID: to.Ptr("EJ_00000000-85ff-222-8d78-444444444444_endpoint2"),
		// 				TargetServerAzureResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver2"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/GetJobPrivateEndpoint.json
func ExampleJobPrivateEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewJobPrivateEndpointsClient().Get(ctx, "group1", "server1", "agent1", "endpoint1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobPrivateEndpoint = armsql.JobPrivateEndpoint{
	// 	Name: to.Ptr("endpoint1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAgents/privateEndpoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/privateEndpoints/endpoint1"),
	// 	Properties: &armsql.JobPrivateEndpointProperties{
	// 		PrivateEndpointID: to.Ptr("EJ_47e33188-85ff-4705-8d78-444444444444_endpoint1"),
	// 		TargetServerAzureResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/CreateOrUpdateJobPrivateEndpoint.json
func ExampleJobPrivateEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobPrivateEndpointsClient().BeginCreateOrUpdate(ctx, "group1", "server1", "agent1", "endpoint1", armsql.JobPrivateEndpoint{
		Properties: &armsql.JobPrivateEndpointProperties{
			TargetServerAzureResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver1"),
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
	// res.JobPrivateEndpoint = armsql.JobPrivateEndpoint{
	// 	Name: to.Ptr("endpoint1"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/jobAgents/privateEndpoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/privateEndpoints/endpoint1"),
	// 	Properties: &armsql.JobPrivateEndpointProperties{
	// 		PrivateEndpointID: to.Ptr("EJ_47e33188-85ff-4705-8d78-444444444444_endpoint1"),
	// 		TargetServerAzureResourceID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/targetserver1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/DeleteJobPrivateEndpoint.json
func ExampleJobPrivateEndpointsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewJobPrivateEndpointsClient().BeginDelete(ctx, "group1", "server1", "agent1", "endpoint1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
