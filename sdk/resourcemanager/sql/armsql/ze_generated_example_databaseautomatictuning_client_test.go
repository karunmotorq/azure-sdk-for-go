//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseAutomaticTuningGet.json
func ExampleDatabaseAutomaticTuningClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewDatabaseAutomaticTuningClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatabaseAutomaticTuningClientGetResult)
}

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseAutomaticTuningUpdateMax.json
func ExampleDatabaseAutomaticTuningClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewDatabaseAutomaticTuningClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<server-name>",
		"<database-name>",
		armsql.DatabaseAutomaticTuning{
			Properties: &armsql.DatabaseAutomaticTuningProperties{
				DesiredState: armsql.AutomaticTuningModeAuto.ToPtr(),
				Options: map[string]*armsql.AutomaticTuningOptions{
					"createIndex": {
						DesiredState: armsql.AutomaticTuningOptionModeDesiredOff.ToPtr(),
					},
					"dropIndex": {
						DesiredState: armsql.AutomaticTuningOptionModeDesiredOn.ToPtr(),
					},
					"forceLastGoodPlan": {
						DesiredState: armsql.AutomaticTuningOptionModeDesiredDefault.ToPtr(),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DatabaseAutomaticTuningClientUpdateResult)
}