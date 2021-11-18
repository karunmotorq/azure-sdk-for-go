//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ListTenantConfigurationViolationsClient contains the methods for the ListTenantConfigurationViolations group.
// Don't use this type directly, use NewListTenantConfigurationViolationsClient() instead.
type ListTenantConfigurationViolationsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewListTenantConfigurationViolationsClient creates a new instance of ListTenantConfigurationViolationsClient with the specified values.
func NewListTenantConfigurationViolationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ListTenantConfigurationViolationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ListTenantConfigurationViolationsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Gets list of items that violate tenant's configuration.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ListTenantConfigurationViolationsClient) List(options *ListTenantConfigurationViolationsListOptions) *ListTenantConfigurationViolationsListPager {
	return &ListTenantConfigurationViolationsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ListTenantConfigurationViolationsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ViolationsList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ListTenantConfigurationViolationsClient) listCreateRequest(ctx context.Context, options *ListTenantConfigurationViolationsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Portal/listTenantConfigurationViolations"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ListTenantConfigurationViolationsClient) listHandleResponse(resp *http.Response) (ListTenantConfigurationViolationsListResponse, error) {
	result := ListTenantConfigurationViolationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ViolationsList); err != nil {
		return ListTenantConfigurationViolationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ListTenantConfigurationViolationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}