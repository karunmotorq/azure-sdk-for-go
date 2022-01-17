//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// PrefixesClient contains the methods for the Prefixes group.
// Don't use this type directly, use NewPrefixesClient() instead.
type PrefixesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPrefixesClient creates a new instance of PrefixesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrefixesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrefixesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ListByPeeringService - Lists the peerings prefix in the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name.
// peeringServiceName - The peering service name.
// options - PrefixesClientListByPeeringServiceOptions contains the optional parameters for the PrefixesClient.ListByPeeringService
// method.
func (client *PrefixesClient) ListByPeeringService(resourceGroupName string, peeringServiceName string, options *PrefixesClientListByPeeringServiceOptions) *PrefixesClientListByPeeringServicePager {
	return &PrefixesClientListByPeeringServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByPeeringServiceCreateRequest(ctx, resourceGroupName, peeringServiceName, options)
		},
		advancer: func(ctx context.Context, resp PrefixesClientListByPeeringServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServicePrefixListResult.NextLink)
		},
	}
}

// listByPeeringServiceCreateRequest creates the ListByPeeringService request.
func (client *PrefixesClient) listByPeeringServiceCreateRequest(ctx context.Context, resourceGroupName string, peeringServiceName string, options *PrefixesClientListByPeeringServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringServiceName == "" {
		return nil, errors.New("parameter peeringServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringServiceName}", url.PathEscape(peeringServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPeeringServiceHandleResponse handles the ListByPeeringService response.
func (client *PrefixesClient) listByPeeringServiceHandleResponse(resp *http.Response) (PrefixesClientListByPeeringServiceResponse, error) {
	result := PrefixesClientListByPeeringServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServicePrefixListResult); err != nil {
		return PrefixesClientListByPeeringServiceResponse{}, err
	}
	return result, nil
}