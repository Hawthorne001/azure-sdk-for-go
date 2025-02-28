//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/InboundEndpoint_Put.json
func ExampleInboundEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInboundEndpointsClient().BeginCreateOrUpdate(ctx, "sampleResourceGroup", "sampleDnsResolver", "sampleInboundEndpoint", armdnsresolver.InboundEndpoint{
		Location: to.Ptr("westus2"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armdnsresolver.InboundEndpointProperties{
			IPConfigurations: []*armdnsresolver.IPConfiguration{
				{
					PrivateIPAllocationMethod: to.Ptr(armdnsresolver.IPAllocationMethodDynamic),
					Subnet: &armdnsresolver.SubResource{
						ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet"),
					},
				}},
		},
	}, &armdnsresolver.InboundEndpointsClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.InboundEndpoint = armdnsresolver.InboundEndpoint{
	// 	Name: to.Ptr("sampleInboundEndpoint"),
	// 	Type: to.Ptr("Microsoft.Network/dnsResolvers/inboundEndpoints"),
	// 	ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/inboundEndpoints/sampleInboundEndpoint"),
	// 	SystemData: &armdnsresolver.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
	// 		CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdnsresolver.InboundEndpointProperties{
	// 		IPConfigurations: []*armdnsresolver.IPConfiguration{
	// 			{
	// 				PrivateIPAddress: to.Ptr("255.255.255.255"),
	// 				PrivateIPAllocationMethod: to.Ptr(armdnsresolver.IPAllocationMethodDynamic),
	// 				Subnet: &armdnsresolver.SubResource{
	// 					ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("87b3e20a-5833-4c40-8ad7-c5160bb1c5bd"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/InboundEndpoint_Patch.json
func ExampleInboundEndpointsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInboundEndpointsClient().BeginUpdate(ctx, "sampleResourceGroup", "sampleDnsResolver", "sampleInboundEndpoint", armdnsresolver.InboundEndpointPatch{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
	}, &armdnsresolver.InboundEndpointsClientBeginUpdateOptions{IfMatch: nil})
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
	// res.InboundEndpoint = armdnsresolver.InboundEndpoint{
	// 	Name: to.Ptr("sampleInboundEndpoint"),
	// 	Type: to.Ptr("Microsoft.Network/dnsResolvers/inboundEndpoints"),
	// 	ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/inboundEndpoints/sampleInboundEndpoint"),
	// 	SystemData: &armdnsresolver.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
	// 		CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdnsresolver.InboundEndpointProperties{
	// 		IPConfigurations: []*armdnsresolver.IPConfiguration{
	// 			{
	// 				PrivateIPAddress: to.Ptr("255.255.255.255"),
	// 				PrivateIPAllocationMethod: to.Ptr(armdnsresolver.IPAllocationMethodDynamic),
	// 				Subnet: &armdnsresolver.SubResource{
	// 					ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("87b3e20a-5833-4c40-8ad7-c5160bb1c5bd"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/InboundEndpoint_Delete.json
func ExampleInboundEndpointsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInboundEndpointsClient().BeginDelete(ctx, "sampleResourceGroup", "sampleDnsResolver", "sampleInboundEndpoint", &armdnsresolver.InboundEndpointsClientBeginDeleteOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/InboundEndpoint_Get.json
func ExampleInboundEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInboundEndpointsClient().Get(ctx, "sampleResourceGroup", "sampleDnsResolver", "sampleInboundEndpoint", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InboundEndpoint = armdnsresolver.InboundEndpoint{
	// 	Name: to.Ptr("sampleInboundEndpoint"),
	// 	Type: to.Ptr("Microsoft.Network/dnsResolvers/inboundEndpoints"),
	// 	ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/inboundEndpoints/sampleInboundEndpoint"),
	// 	SystemData: &armdnsresolver.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
	// 		CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
	// 		LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("westus2"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armdnsresolver.InboundEndpointProperties{
	// 		IPConfigurations: []*armdnsresolver.IPConfiguration{
	// 			{
	// 				PrivateIPAddress: to.Ptr("255.255.255.255"),
	// 				PrivateIPAllocationMethod: to.Ptr(armdnsresolver.IPAllocationMethodDynamic),
	// 				Subnet: &armdnsresolver.SubResource{
	// 					ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("b6b2d964-8588-4e3a-a7fe-8a5b7fe8eca5"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/InboundEndpoint_List.json
func ExampleInboundEndpointsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdnsresolver.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInboundEndpointsClient().NewListPager("sampleResourceGroup", "sampleDnsResolver", &armdnsresolver.InboundEndpointsClientListOptions{Top: nil})
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
		// page.InboundEndpointListResult = armdnsresolver.InboundEndpointListResult{
		// 	Value: []*armdnsresolver.InboundEndpoint{
		// 		{
		// 			Name: to.Ptr("sampleInboundEndpoint1"),
		// 			Type: to.Ptr("Microsoft.Network/dnsResolvers/inboundEndpoints"),
		// 			ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/inboundEndpoints/sampleInboundEndpoint1"),
		// 			SystemData: &armdnsresolver.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T01:01:01.107Z"); return t}()),
		// 				CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T02:03:01.197Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armdnsresolver.InboundEndpointProperties{
		// 				IPConfigurations: []*armdnsresolver.IPConfiguration{
		// 					{
		// 						PrivateIPAddress: to.Ptr("255.1.255.1"),
		// 						PrivateIPAllocationMethod: to.Ptr(armdnsresolver.IPAllocationMethodDynamic),
		// 						Subnet: &armdnsresolver.SubResource{
		// 							ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet1"),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("b6b2d964-8588-4e3a-a7fe-8a5b7fe8eca5"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sampleInboundEndpoint2"),
		// 			Type: to.Ptr("Microsoft.Network/dnsResolvers/inboundEndpoints"),
		// 			ID: to.Ptr("/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/inboundEndpoints/sampleInboundEndpoint2"),
		// 			SystemData: &armdnsresolver.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-02T01:01:01.107Z"); return t}()),
		// 				CreatedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-03T02:03:01.197Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armdnsresolver.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 			},
		// 			Etag: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armdnsresolver.InboundEndpointProperties{
		// 				IPConfigurations: []*armdnsresolver.IPConfiguration{
		// 					{
		// 						PrivateIPAddress: to.Ptr("1.1.255.1"),
		// 						PrivateIPAllocationMethod: to.Ptr(armdnsresolver.IPAllocationMethodDynamic),
		// 						Subnet: &armdnsresolver.SubResource{
		// 							ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet1"),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armdnsresolver.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("87b3e20a-5833-4c40-8ad7-c5160bb1c5bd"),
		// 			},
		// 	}},
		// }
	}
}
