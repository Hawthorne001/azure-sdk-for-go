//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armprivatelinkscopes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armprivatelinkscopes"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateEndpointConnectionGet.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatelinkscopes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "myResourceGroup", "myPrivateLinkScope", "private-endpoint-connection-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armprivatelinkscopes.PrivateEndpointConnection{
	// 	Name: to.Ptr("private-endpoint-connection-name"),
	// 	Type: to.Ptr("Microsoft.KubernetesConfiguration/privateLinkScopes/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers//privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 	Properties: &armprivatelinkscopes.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armprivatelinkscopes.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armprivatelinkscopes.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armprivatelinkscopes.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armprivatelinkscopes.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateEndpointConnectionUpdate.json
func ExamplePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatelinkscopes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myPrivateLinkScope", "private-endpoint-connection-name", armprivatelinkscopes.PrivateEndpointConnection{
		Properties: &armprivatelinkscopes.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armprivatelinkscopes.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Approved by johndoe@contoso.com"),
				Status:      to.Ptr(armprivatelinkscopes.PrivateEndpointServiceConnectionStatusApproved),
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
	// res.PrivateEndpointConnection = armprivatelinkscopes.PrivateEndpointConnection{
	// 	Name: to.Ptr("private-endpoint-connection-name"),
	// 	Type: to.Ptr("Microsoft.KubernetesConfiguration/privateLinkScopes/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.KubernetesConfiguration/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
	// 	Properties: &armprivatelinkscopes.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armprivatelinkscopes.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armprivatelinkscopes.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Approved by johndoe@contoso.com"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armprivatelinkscopes.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armprivatelinkscopes.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateEndpointConnectionDelete.json
func ExamplePrivateEndpointConnectionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatelinkscopes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionsClient().Delete(ctx, "myResourceGroup", "myPrivateLinkScope", "private-endpoint-connection-name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateEndpointConnectionList.json
func ExamplePrivateEndpointConnectionsClient_ListByPrivateLinkScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatelinkscopes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().ListByPrivateLinkScope(ctx, "myResourceGroup", "myPrivateLinkScope", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionListResult = armprivatelinkscopes.PrivateEndpointConnectionListResult{
	// 	Value: []*armprivatelinkscopes.PrivateEndpointConnection{
	// 		{
	// 			Name: to.Ptr("private-endpoint-connection-name"),
	// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/privateLinkScopes/privateEndpointConnections"),
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.KubernetesConfiguration/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name-2"),
	// 			Properties: &armprivatelinkscopes.PrivateEndpointConnectionProperties{
	// 				PrivateEndpoint: &armprivatelinkscopes.PrivateEndpoint{
	// 					ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 				},
	// 				PrivateLinkServiceConnectionState: &armprivatelinkscopes.PrivateLinkServiceConnectionState{
	// 					Description: to.Ptr("Auto-approved"),
	// 					ActionsRequired: to.Ptr("None"),
	// 					Status: to.Ptr(armprivatelinkscopes.PrivateEndpointServiceConnectionStatusApproved),
	// 				},
	// 				ProvisioningState: to.Ptr(armprivatelinkscopes.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("private-endpoint-connection-name-2"),
	// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/privateLinkScopes/privateEndpointConnections"),
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.KubernetesConfiguration/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name-2"),
	// 			Properties: &armprivatelinkscopes.PrivateEndpointConnectionProperties{
	// 				PrivateEndpoint: &armprivatelinkscopes.PrivateEndpoint{
	// 					ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name-2"),
	// 				},
	// 				PrivateLinkServiceConnectionState: &armprivatelinkscopes.PrivateLinkServiceConnectionState{
	// 					Description: to.Ptr("Please approve my connection."),
	// 					ActionsRequired: to.Ptr("None"),
	// 					Status: to.Ptr(armprivatelinkscopes.PrivateEndpointServiceConnectionStatusPending),
	// 				},
	// 				ProvisioningState: to.Ptr(armprivatelinkscopes.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 			},
	// 	}},
	// }
}
