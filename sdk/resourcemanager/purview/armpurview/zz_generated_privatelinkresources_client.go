//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpurview

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
// subscriptionID - The subscription identifier
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkResourcesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// GetByGroupID - Gets a privately linkable resources for an account with given group identifier
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - The name of the account.
// groupID - The group identifier.
// options - PrivateLinkResourcesClientGetByGroupIDOptions contains the optional parameters for the PrivateLinkResourcesClient.GetByGroupID
// method.
func (client *PrivateLinkResourcesClient) GetByGroupID(ctx context.Context, resourceGroupName string, accountName string, groupID string, options *PrivateLinkResourcesClientGetByGroupIDOptions) (PrivateLinkResourcesClientGetByGroupIDResponse, error) {
	req, err := client.getByGroupIDCreateRequest(ctx, resourceGroupName, accountName, groupID, options)
	if err != nil {
		return PrivateLinkResourcesClientGetByGroupIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkResourcesClientGetByGroupIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkResourcesClientGetByGroupIDResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByGroupIDHandleResponse(resp)
}

// getByGroupIDCreateRequest creates the GetByGroupID request.
func (client *PrivateLinkResourcesClient) getByGroupIDCreateRequest(ctx context.Context, resourceGroupName string, accountName string, groupID string, options *PrivateLinkResourcesClientGetByGroupIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/privateLinkResources/{groupId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if groupID == "" {
		return nil, errors.New("parameter groupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupId}", url.PathEscape(groupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByGroupIDHandleResponse handles the GetByGroupID response.
func (client *PrivateLinkResourcesClient) getByGroupIDHandleResponse(resp *http.Response) (PrivateLinkResourcesClientGetByGroupIDResponse, error) {
	result := PrivateLinkResourcesClientGetByGroupIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesClientGetByGroupIDResponse{}, err
	}
	return result, nil
}

// ListByAccount - Gets a list of privately linkable resources for an account
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// accountName - The name of the account.
// options - PrivateLinkResourcesClientListByAccountOptions contains the optional parameters for the PrivateLinkResourcesClient.ListByAccount
// method.
func (client *PrivateLinkResourcesClient) ListByAccount(resourceGroupName string, accountName string, options *PrivateLinkResourcesClientListByAccountOptions) *PrivateLinkResourcesClientListByAccountPager {
	return &PrivateLinkResourcesClientListByAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkResourcesClientListByAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateLinkResourceList.NextLink)
		},
	}
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *PrivateLinkResourcesClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *PrivateLinkResourcesClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Purview/accounts/{accountName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *PrivateLinkResourcesClient) listByAccountHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByAccountResponse, error) {
	result := PrivateLinkResourcesClientListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceList); err != nil {
		return PrivateLinkResourcesClientListByAccountResponse{}, err
	}
	return result, nil
}