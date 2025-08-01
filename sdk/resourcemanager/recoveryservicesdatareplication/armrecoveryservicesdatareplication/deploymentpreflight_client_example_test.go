// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armrecoveryservicesdatareplication_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
	"log"
)

// Generated from example definition: 2024-09-01/DeploymentPreflight_Post.json
func ExampleDeploymentPreflightClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeploymentPreflightClient().Post(ctx, "rgswagger_2024-09-01", "lnfcwsmlowbwkndkztzvaj", &armrecoveryservicesdatareplication.DeploymentPreflightClientPostOptions{
		Body: &armrecoveryservicesdatareplication.DeploymentPreflightModel{
			Resources: []*armrecoveryservicesdatareplication.DeploymentPreflightResource{
				{
					Name:       to.Ptr("xtgugoflfc"),
					Type:       to.Ptr("nsnaptduolqcxsikrewvgjbxqpt"),
					Location:   to.Ptr("cbsgtxkjdzwbyp"),
					APIVersion: to.Ptr("otihymhvzblycdoxo"),
				},
			},
		}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesdatareplication.DeploymentPreflightClientPostResponse{
	// 	DeploymentPreflightModel: &armrecoveryservicesdatareplication.DeploymentPreflightModel{
	// 		Resources: []*armrecoveryservicesdatareplication.DeploymentPreflightResource{
	// 			{
	// 				Name: to.Ptr("xtgugoflfc"),
	// 				Type: to.Ptr("nsnaptduolqcxsikrewvgjbxqpt"),
	// 				Location: to.Ptr("cbsgtxkjdzwbyp"),
	// 				APIVersion: to.Ptr("otihymhvzblycdoxo"),
	// 			},
	// 		},
	// 	},
	// }
}
