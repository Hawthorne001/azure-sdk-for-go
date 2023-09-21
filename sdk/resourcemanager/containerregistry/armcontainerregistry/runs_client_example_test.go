//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsList.json
func ExampleRunsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRunsClient().NewListPager("myResourceGroup", "myRegistry", &armcontainerregistry.RunsClientListOptions{Filter: to.Ptr(""),
		Top: to.Ptr[int32](10),
	})
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
		// page.RunListResult = armcontainerregistry.RunListResult{
		// 	Value: []*armcontainerregistry.Run{
		// 		{
		// 			Name: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
		// 			Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/0accec26-d6de-4757-8e74-d080f38eaaab"),
		// 			Properties: &armcontainerregistry.RunProperties{
		// 				AgentConfiguration: &armcontainerregistry.AgentProperties{
		// 					CPU: to.Ptr[int32](2),
		// 				},
		// 				CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
		// 				FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T06:13:51.618Z"); return t}()),
		// 				ImageUpdateTrigger: &armcontainerregistry.ImageUpdateTrigger{
		// 					ID: to.Ptr("c0c43143-da5d-41ef-b9e1-e7d749272e88"),
		// 					Images: []*armcontainerregistry.ImageDescriptor{
		// 						{
		// 							Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
		// 							Registry: to.Ptr("registry.hub.docker.com"),
		// 							Repository: to.Ptr("mybaseimage"),
		// 							Tag: to.Ptr("latest"),
		// 					}},
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
		// 				},
		// 				IsArchiveEnabled: to.Ptr(true),
		// 				LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.617Z"); return t}()),
		// 				LogArtifact: &armcontainerregistry.ImageDescriptor{
		// 					Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
		// 					Registry: to.Ptr("myregistry"),
		// 					Repository: to.Ptr("acr/tasks"),
		// 					Tag: to.Ptr("mytask-0accec26-d6de-4757-8e74-d080f38eaaab-log"),
		// 				},
		// 				OutputImages: []*armcontainerregistry.ImageDescriptor{
		// 					{
		// 						Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
		// 						Registry: to.Ptr("myregistry.azurecr.io"),
		// 						Repository: to.Ptr("myimage"),
		// 						Tag: to.Ptr("latest"),
		// 				}},
		// 				Platform: &armcontainerregistry.PlatformProperties{
		// 					Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
		// 					OS: to.Ptr(armcontainerregistry.OSLinux),
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				RunID: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
		// 				RunType: to.Ptr(armcontainerregistry.RunTypeAutoBuild),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:50:51.618Z"); return t}()),
		// 				Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
		// 				Task: to.Ptr("myTask"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsGet.json
func ExampleRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRunsClient().Get(ctx, "myResourceGroup", "myRegistry", "0accec26-d6de-4757-8e74-d080f38eaaab", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Run = armcontainerregistry.Run{
	// 	Name: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Properties: &armcontainerregistry.RunProperties{
	// 		AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 			CPU: to.Ptr[int32](2),
	// 		},
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
	// 		FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T06:13:51.618Z"); return t}()),
	// 		ImageUpdateTrigger: &armcontainerregistry.ImageUpdateTrigger{
	// 			ID: to.Ptr("c0c43143-da5d-41ef-b9e1-e7d749272e88"),
	// 			Images: []*armcontainerregistry.ImageDescriptor{
	// 				{
	// 					Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 					Registry: to.Ptr("registry.hub.docker.com"),
	// 					Repository: to.Ptr("mybaseimage"),
	// 					Tag: to.Ptr("latest"),
	// 			}},
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
	// 		},
	// 		IsArchiveEnabled: to.Ptr(true),
	// 		LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.617Z"); return t}()),
	// 		LogArtifact: &armcontainerregistry.ImageDescriptor{
	// 			Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 			Registry: to.Ptr("myregistry"),
	// 			Repository: to.Ptr("acr/tasks"),
	// 			Tag: to.Ptr("mytask-0accec26-d6de-4757-8e74-d080f38eaaab-log"),
	// 		},
	// 		OutputImages: []*armcontainerregistry.ImageDescriptor{
	// 			{
	// 				Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 				Registry: to.Ptr("myregistry.azurecr.io"),
	// 				Repository: to.Ptr("myimage"),
	// 				Tag: to.Ptr("latest"),
	// 		}},
	// 		Platform: &armcontainerregistry.PlatformProperties{
	// 			Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 			OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunID: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 		RunType: to.Ptr(armcontainerregistry.RunTypeAutoBuild),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:50:51.618Z"); return t}()),
	// 		Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
	// 		Task: to.Ptr("myTask"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsUpdate.json
func ExampleRunsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRunsClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "0accec26-d6de-4757-8e74-d080f38eaaab", armcontainerregistry.RunUpdateParameters{
		IsArchiveEnabled: to.Ptr(true),
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
	// res.Run = armcontainerregistry.Run{
	// 	Name: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/runs"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/runs/0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 	Properties: &armcontainerregistry.RunProperties{
	// 		AgentConfiguration: &armcontainerregistry.AgentProperties{
	// 			CPU: to.Ptr[int32](2),
	// 		},
	// 		CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
	// 		FinishTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T06:13:51.618Z"); return t}()),
	// 		ImageUpdateTrigger: &armcontainerregistry.ImageUpdateTrigger{
	// 			ID: to.Ptr("c0c43143-da5d-41ef-b9e1-e7d749272e88"),
	// 			Images: []*armcontainerregistry.ImageDescriptor{
	// 				{
	// 					Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 					Registry: to.Ptr("registry.hub.docker.com"),
	// 					Repository: to.Ptr("mybaseimage"),
	// 					Tag: to.Ptr("latest"),
	// 			}},
	// 			Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.618Z"); return t}()),
	// 		},
	// 		IsArchiveEnabled: to.Ptr(true),
	// 		LastUpdatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:13:51.617Z"); return t}()),
	// 		LogArtifact: &armcontainerregistry.ImageDescriptor{
	// 			Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 			Registry: to.Ptr("myregistry"),
	// 			Repository: to.Ptr("acr/tasks"),
	// 			Tag: to.Ptr("mytask-0accec26-d6de-4757-8e74-d080f38eaaab-log"),
	// 		},
	// 		OutputImages: []*armcontainerregistry.ImageDescriptor{
	// 			{
	// 				Digest: to.Ptr("sha256:cbbf2f9a99b47fc460d422812b6a5adff7dfee951d8fa2e4a98caa0"),
	// 				Registry: to.Ptr("myregistry.azurecr.io"),
	// 				Repository: to.Ptr("myimage"),
	// 				Tag: to.Ptr("latest"),
	// 		}},
	// 		Platform: &armcontainerregistry.PlatformProperties{
	// 			Architecture: to.Ptr(armcontainerregistry.ArchitectureAmd64),
	// 			OS: to.Ptr(armcontainerregistry.OSLinux),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		RunID: to.Ptr("0accec26-d6de-4757-8e74-d080f38eaaab"),
	// 		RunType: to.Ptr(armcontainerregistry.RunTypeAutoBuild),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-25T05:50:51.618Z"); return t}()),
	// 		Status: to.Ptr(armcontainerregistry.RunStatusSucceeded),
	// 		Task: to.Ptr("myTask"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsGetLogSasUrl.json
func ExampleRunsClient_GetLogSasURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRunsClient().GetLogSasURL(ctx, "myResourceGroup", "myRegistry", "0accec26-d6de-4757-8e74-d080f38eaaab", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RunGetLogResult = armcontainerregistry.RunGetLogResult{
	// 	LogLink: to.Ptr("https://registrystorageaccount.blob.core.windows.net/sascontainer/logs/0accec26-d6de-4757-8e74-d080f38eaaab/rawtext.log?sv=2015-04-05&st=2015-04-29T22%3A18%3A26Z&se=2015-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=Z%2FRHIX5Xcg0Mq2rqI3OlWTjEg2tYkboXr1P9ZUXDtkk%3D"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/68d03f91ea7c30e1ab28fb9d35c13f81bc85b724/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsCancel.json
func ExampleRunsClient_BeginCancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRunsClient().BeginCancel(ctx, "myResourceGroup", "myRegistry", "0accec26-d6de-4757-8e74-d080f38eaaab", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
