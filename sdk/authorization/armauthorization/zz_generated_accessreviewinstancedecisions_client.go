//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AccessReviewInstanceDecisionsClient contains the methods for the AccessReviewInstanceDecisions group.
// Don't use this type directly, use NewAccessReviewInstanceDecisionsClient() instead.
type AccessReviewInstanceDecisionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAccessReviewInstanceDecisionsClient creates a new instance of AccessReviewInstanceDecisionsClient with the specified values.
func NewAccessReviewInstanceDecisionsClient(con *arm.Connection, subscriptionID string) *AccessReviewInstanceDecisionsClient {
	return &AccessReviewInstanceDecisionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Get access review instance decisions
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewInstanceDecisionsClient) List(scheduleDefinitionID string, id string, options *AccessReviewInstanceDecisionsListOptions) *AccessReviewInstanceDecisionsListPager {
	return &AccessReviewInstanceDecisionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scheduleDefinitionID, id, options)
		},
		advancer: func(ctx context.Context, resp AccessReviewInstanceDecisionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccessReviewDecisionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstanceDecisionsClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceDecisionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/decisions"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewInstanceDecisionsClient) listHandleResponse(resp *http.Response) (AccessReviewInstanceDecisionsListResponse, error) {
	result := AccessReviewInstanceDecisionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDecisionListResult); err != nil {
		return AccessReviewInstanceDecisionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AccessReviewInstanceDecisionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}