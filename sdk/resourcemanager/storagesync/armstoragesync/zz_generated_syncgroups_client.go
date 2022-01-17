//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SyncGroupsClient contains the methods for the SyncGroups group.
// Don't use this type directly, use NewSyncGroupsClient() instead.
type SyncGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSyncGroupsClient creates a new instance of SyncGroupsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSyncGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SyncGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SyncGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Create - Create a new SyncGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// parameters - Sync Group Body
// options - SyncGroupsClientCreateOptions contains the optional parameters for the SyncGroupsClient.Create method.
func (client *SyncGroupsClient) Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters SyncGroupCreateParameters, options *SyncGroupsClientCreateOptions) (SyncGroupsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, parameters, options)
	if err != nil {
		return SyncGroupsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncGroupsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SyncGroupsClient) createCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters SyncGroupCreateParameters, options *SyncGroupsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *SyncGroupsClient) createHandleResponse(resp *http.Response) (SyncGroupsClientCreateResponse, error) {
	result := SyncGroupsClientCreateResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncGroup); err != nil {
		return SyncGroupsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a given SyncGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// options - SyncGroupsClientDeleteOptions contains the optional parameters for the SyncGroupsClient.Delete method.
func (client *SyncGroupsClient) Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsClientDeleteOptions) (SyncGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, options)
	if err != nil {
		return SyncGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SyncGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *SyncGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *SyncGroupsClient) deleteHandleResponse(resp *http.Response) (SyncGroupsClientDeleteResponse, error) {
	result := SyncGroupsClientDeleteResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	return result, nil
}

// Get - Get a given SyncGroup.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// syncGroupName - Name of Sync Group resource.
// options - SyncGroupsClientGetOptions contains the optional parameters for the SyncGroupsClient.Get method.
func (client *SyncGroupsClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsClientGetOptions) (SyncGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, options)
	if err != nil {
		return SyncGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SyncGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, options *SyncGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	if syncGroupName == "" {
		return nil, errors.New("parameter syncGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{syncGroupName}", url.PathEscape(syncGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SyncGroupsClient) getHandleResponse(resp *http.Response) (SyncGroupsClientGetResponse, error) {
	result := SyncGroupsClientGetResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncGroup); err != nil {
		return SyncGroupsClientGetResponse{}, err
	}
	return result, nil
}

// ListByStorageSyncService - Get a SyncGroup List.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// storageSyncServiceName - Name of Storage Sync Service resource.
// options - SyncGroupsClientListByStorageSyncServiceOptions contains the optional parameters for the SyncGroupsClient.ListByStorageSyncService
// method.
func (client *SyncGroupsClient) ListByStorageSyncService(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *SyncGroupsClientListByStorageSyncServiceOptions) (SyncGroupsClientListByStorageSyncServiceResponse, error) {
	req, err := client.listByStorageSyncServiceCreateRequest(ctx, resourceGroupName, storageSyncServiceName, options)
	if err != nil {
		return SyncGroupsClientListByStorageSyncServiceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SyncGroupsClientListByStorageSyncServiceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SyncGroupsClientListByStorageSyncServiceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByStorageSyncServiceHandleResponse(resp)
}

// listByStorageSyncServiceCreateRequest creates the ListByStorageSyncService request.
func (client *SyncGroupsClient) listByStorageSyncServiceCreateRequest(ctx context.Context, resourceGroupName string, storageSyncServiceName string, options *SyncGroupsClientListByStorageSyncServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageSyncServiceName == "" {
		return nil, errors.New("parameter storageSyncServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageSyncServiceName}", url.PathEscape(storageSyncServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByStorageSyncServiceHandleResponse handles the ListByStorageSyncService response.
func (client *SyncGroupsClient) listByStorageSyncServiceHandleResponse(resp *http.Response) (SyncGroupsClientListByStorageSyncServiceResponse, error) {
	result := SyncGroupsClientListByStorageSyncServiceResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if val := resp.Header.Get("x-ms-correlation-request-id"); val != "" {
		result.XMSCorrelationRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SyncGroupArray); err != nil {
		return SyncGroupsClientListByStorageSyncServiceResponse{}, err
	}
	return result, nil
}