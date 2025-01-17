//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RestoresClient contains the methods for the Restores group.
// Don't use this type directly, use NewRestoresClient() instead.
type RestoresClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestoresClient creates a new instance of RestoresClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestoresClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &RestoresClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginTrigger - Restores the specified backed up data. This is an asynchronous operation. To know the status of this API
// call, use GetProtectedItemOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the backed up items.
//   - containerName - Container name associated with the backed up items.
//   - protectedItemName - Backed up item to be restored.
//   - recoveryPointID - Recovery point ID which represents the backed up data to be restored.
//   - parameters - resource restore request
//   - options - RestoresClientBeginTriggerOptions contains the optional parameters for the RestoresClient.BeginTrigger method.
func (client *RestoresClient) BeginTrigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters RestoreRequestResource, options *RestoresClientBeginTriggerOptions) (*runtime.Poller[RestoresClientTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.trigger(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RestoresClientTriggerResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[RestoresClientTriggerResponse](options.ResumeToken, client.pl, nil)
	}
}

// Trigger - Restores the specified backed up data. This is an asynchronous operation. To know the status of this API call,
// use GetProtectedItemOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
func (client *RestoresClient) trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters RestoreRequestResource, options *RestoresClientBeginTriggerOptions) (*http.Response, error) {
	req, err := client.triggerCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// triggerCreateRequest creates the Trigger request.
func (client *RestoresClient) triggerCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters RestoreRequestResource, options *RestoresClientBeginTriggerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/restore"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	if recoveryPointID == "" {
		return nil, errors.New("parameter recoveryPointID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointId}", url.PathEscape(recoveryPointID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
