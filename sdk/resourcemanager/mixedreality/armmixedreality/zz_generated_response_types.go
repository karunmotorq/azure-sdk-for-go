//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality

import "net/http"

// ClientCheckNameAvailabilityLocalResponse contains the response from method Client.CheckNameAvailabilityLocal.
type ClientCheckNameAvailabilityLocalResponse struct {
	ClientCheckNameAvailabilityLocalResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientCheckNameAvailabilityLocalResult contains the result from method Client.CheckNameAvailabilityLocal.
type ClientCheckNameAvailabilityLocalResult struct {
	CheckNameAvailabilityResponse
}

// ObjectAnchorsAccountsClientCreateResponse contains the response from method ObjectAnchorsAccountsClient.Create.
type ObjectAnchorsAccountsClientCreateResponse struct {
	ObjectAnchorsAccountsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientCreateResult contains the result from method ObjectAnchorsAccountsClient.Create.
type ObjectAnchorsAccountsClientCreateResult struct {
	ObjectAnchorsAccount
}

// ObjectAnchorsAccountsClientDeleteResponse contains the response from method ObjectAnchorsAccountsClient.Delete.
type ObjectAnchorsAccountsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientGetResponse contains the response from method ObjectAnchorsAccountsClient.Get.
type ObjectAnchorsAccountsClientGetResponse struct {
	ObjectAnchorsAccountsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientGetResult contains the result from method ObjectAnchorsAccountsClient.Get.
type ObjectAnchorsAccountsClientGetResult struct {
	ObjectAnchorsAccount
}

// ObjectAnchorsAccountsClientListByResourceGroupResponse contains the response from method ObjectAnchorsAccountsClient.ListByResourceGroup.
type ObjectAnchorsAccountsClientListByResourceGroupResponse struct {
	ObjectAnchorsAccountsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientListByResourceGroupResult contains the result from method ObjectAnchorsAccountsClient.ListByResourceGroup.
type ObjectAnchorsAccountsClientListByResourceGroupResult struct {
	ObjectAnchorsAccountPage
}

// ObjectAnchorsAccountsClientListBySubscriptionResponse contains the response from method ObjectAnchorsAccountsClient.ListBySubscription.
type ObjectAnchorsAccountsClientListBySubscriptionResponse struct {
	ObjectAnchorsAccountsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientListBySubscriptionResult contains the result from method ObjectAnchorsAccountsClient.ListBySubscription.
type ObjectAnchorsAccountsClientListBySubscriptionResult struct {
	ObjectAnchorsAccountPage
}

// ObjectAnchorsAccountsClientListKeysResponse contains the response from method ObjectAnchorsAccountsClient.ListKeys.
type ObjectAnchorsAccountsClientListKeysResponse struct {
	ObjectAnchorsAccountsClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientListKeysResult contains the result from method ObjectAnchorsAccountsClient.ListKeys.
type ObjectAnchorsAccountsClientListKeysResult struct {
	AccountKeys
}

// ObjectAnchorsAccountsClientRegenerateKeysResponse contains the response from method ObjectAnchorsAccountsClient.RegenerateKeys.
type ObjectAnchorsAccountsClientRegenerateKeysResponse struct {
	ObjectAnchorsAccountsClientRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientRegenerateKeysResult contains the result from method ObjectAnchorsAccountsClient.RegenerateKeys.
type ObjectAnchorsAccountsClientRegenerateKeysResult struct {
	AccountKeys
}

// ObjectAnchorsAccountsClientUpdateResponse contains the response from method ObjectAnchorsAccountsClient.Update.
type ObjectAnchorsAccountsClientUpdateResponse struct {
	ObjectAnchorsAccountsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectAnchorsAccountsClientUpdateResult contains the result from method ObjectAnchorsAccountsClient.Update.
type ObjectAnchorsAccountsClientUpdateResult struct {
	ObjectAnchorsAccount
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationPage
}

// RemoteRenderingAccountsClientCreateResponse contains the response from method RemoteRenderingAccountsClient.Create.
type RemoteRenderingAccountsClientCreateResponse struct {
	RemoteRenderingAccountsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientCreateResult contains the result from method RemoteRenderingAccountsClient.Create.
type RemoteRenderingAccountsClientCreateResult struct {
	RemoteRenderingAccount
}

// RemoteRenderingAccountsClientDeleteResponse contains the response from method RemoteRenderingAccountsClient.Delete.
type RemoteRenderingAccountsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientGetResponse contains the response from method RemoteRenderingAccountsClient.Get.
type RemoteRenderingAccountsClientGetResponse struct {
	RemoteRenderingAccountsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientGetResult contains the result from method RemoteRenderingAccountsClient.Get.
type RemoteRenderingAccountsClientGetResult struct {
	RemoteRenderingAccount
}

// RemoteRenderingAccountsClientListByResourceGroupResponse contains the response from method RemoteRenderingAccountsClient.ListByResourceGroup.
type RemoteRenderingAccountsClientListByResourceGroupResponse struct {
	RemoteRenderingAccountsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientListByResourceGroupResult contains the result from method RemoteRenderingAccountsClient.ListByResourceGroup.
type RemoteRenderingAccountsClientListByResourceGroupResult struct {
	RemoteRenderingAccountPage
}

// RemoteRenderingAccountsClientListBySubscriptionResponse contains the response from method RemoteRenderingAccountsClient.ListBySubscription.
type RemoteRenderingAccountsClientListBySubscriptionResponse struct {
	RemoteRenderingAccountsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientListBySubscriptionResult contains the result from method RemoteRenderingAccountsClient.ListBySubscription.
type RemoteRenderingAccountsClientListBySubscriptionResult struct {
	RemoteRenderingAccountPage
}

// RemoteRenderingAccountsClientListKeysResponse contains the response from method RemoteRenderingAccountsClient.ListKeys.
type RemoteRenderingAccountsClientListKeysResponse struct {
	RemoteRenderingAccountsClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientListKeysResult contains the result from method RemoteRenderingAccountsClient.ListKeys.
type RemoteRenderingAccountsClientListKeysResult struct {
	AccountKeys
}

// RemoteRenderingAccountsClientRegenerateKeysResponse contains the response from method RemoteRenderingAccountsClient.RegenerateKeys.
type RemoteRenderingAccountsClientRegenerateKeysResponse struct {
	RemoteRenderingAccountsClientRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientRegenerateKeysResult contains the result from method RemoteRenderingAccountsClient.RegenerateKeys.
type RemoteRenderingAccountsClientRegenerateKeysResult struct {
	AccountKeys
}

// RemoteRenderingAccountsClientUpdateResponse contains the response from method RemoteRenderingAccountsClient.Update.
type RemoteRenderingAccountsClientUpdateResponse struct {
	RemoteRenderingAccountsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RemoteRenderingAccountsClientUpdateResult contains the result from method RemoteRenderingAccountsClient.Update.
type RemoteRenderingAccountsClientUpdateResult struct {
	RemoteRenderingAccount
}

// SpatialAnchorsAccountsClientCreateResponse contains the response from method SpatialAnchorsAccountsClient.Create.
type SpatialAnchorsAccountsClientCreateResponse struct {
	SpatialAnchorsAccountsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientCreateResult contains the result from method SpatialAnchorsAccountsClient.Create.
type SpatialAnchorsAccountsClientCreateResult struct {
	SpatialAnchorsAccount
}

// SpatialAnchorsAccountsClientDeleteResponse contains the response from method SpatialAnchorsAccountsClient.Delete.
type SpatialAnchorsAccountsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientGetResponse contains the response from method SpatialAnchorsAccountsClient.Get.
type SpatialAnchorsAccountsClientGetResponse struct {
	SpatialAnchorsAccountsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientGetResult contains the result from method SpatialAnchorsAccountsClient.Get.
type SpatialAnchorsAccountsClientGetResult struct {
	SpatialAnchorsAccount
}

// SpatialAnchorsAccountsClientListByResourceGroupResponse contains the response from method SpatialAnchorsAccountsClient.ListByResourceGroup.
type SpatialAnchorsAccountsClientListByResourceGroupResponse struct {
	SpatialAnchorsAccountsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientListByResourceGroupResult contains the result from method SpatialAnchorsAccountsClient.ListByResourceGroup.
type SpatialAnchorsAccountsClientListByResourceGroupResult struct {
	SpatialAnchorsAccountPage
}

// SpatialAnchorsAccountsClientListBySubscriptionResponse contains the response from method SpatialAnchorsAccountsClient.ListBySubscription.
type SpatialAnchorsAccountsClientListBySubscriptionResponse struct {
	SpatialAnchorsAccountsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientListBySubscriptionResult contains the result from method SpatialAnchorsAccountsClient.ListBySubscription.
type SpatialAnchorsAccountsClientListBySubscriptionResult struct {
	SpatialAnchorsAccountPage
}

// SpatialAnchorsAccountsClientListKeysResponse contains the response from method SpatialAnchorsAccountsClient.ListKeys.
type SpatialAnchorsAccountsClientListKeysResponse struct {
	SpatialAnchorsAccountsClientListKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientListKeysResult contains the result from method SpatialAnchorsAccountsClient.ListKeys.
type SpatialAnchorsAccountsClientListKeysResult struct {
	AccountKeys
}

// SpatialAnchorsAccountsClientRegenerateKeysResponse contains the response from method SpatialAnchorsAccountsClient.RegenerateKeys.
type SpatialAnchorsAccountsClientRegenerateKeysResponse struct {
	SpatialAnchorsAccountsClientRegenerateKeysResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientRegenerateKeysResult contains the result from method SpatialAnchorsAccountsClient.RegenerateKeys.
type SpatialAnchorsAccountsClientRegenerateKeysResult struct {
	AccountKeys
}

// SpatialAnchorsAccountsClientUpdateResponse contains the response from method SpatialAnchorsAccountsClient.Update.
type SpatialAnchorsAccountsClientUpdateResponse struct {
	SpatialAnchorsAccountsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SpatialAnchorsAccountsClientUpdateResult contains the result from method SpatialAnchorsAccountsClient.Update.
type SpatialAnchorsAccountsClientUpdateResult struct {
	SpatialAnchorsAccount
}