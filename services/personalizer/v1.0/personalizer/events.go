package personalizer

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// EventsClient is the personalizer Service is an Azure Cognitive Service that makes it easy to target content and
// experiences without complex pre-analysis or cleanup of past data. Given a context and featurized content, the
// Personalizer Service returns which content item to show to users in rewardActionId. As rewards are sent in response
// to the use of rewardActionId, the reinforcement learning algorithm will improve the model and improve performance of
// future rank calls.
type EventsClient struct {
	BaseClient
}

// NewEventsClient creates an instance of the EventsClient client.
func NewEventsClient(endpoint string) EventsClient {
	return EventsClient{New(endpoint)}
}

// Activate report that the specified event was actually displayed to the user and a reward should be expected for it
// Parameters:
// eventID - the event ID this activation applies to.
func (client EventsClient) Activate(ctx context.Context, eventID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventsClient.Activate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: eventID,
			Constraints: []validation.Constraint{{Target: "eventID", Name: validation.MaxLength, Rule: 256, Chain: nil}}}}); err != nil {
		return result, validation.NewError("personalizer.EventsClient", "Activate", err.Error())
	}

	req, err := client.ActivatePreparer(ctx, eventID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.EventsClient", "Activate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ActivateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "personalizer.EventsClient", "Activate", resp, "Failure sending request")
		return
	}

	result, err = client.ActivateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.EventsClient", "Activate", resp, "Failure responding to request")
	}

	return
}

// ActivatePreparer prepares the Activate request.
func (client EventsClient) ActivatePreparer(ctx context.Context, eventID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"eventId": autorest.Encode("path", eventID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/personalizer/v1.0", urlParameters),
		autorest.WithPathParameters("/events/{eventId}/activate", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ActivateSender sends the Activate request. The method will close the
// http.Response Body if it receives an error.
func (client EventsClient) ActivateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ActivateResponder handles the response to the Activate request. The method always
// closes the http.Response Body.
func (client EventsClient) ActivateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Reward report reward that resulted from using the action specified in rewardActionId for the specified event.
// Parameters:
// eventID - the event id this reward applies to.
// reward - the reward should be a floating point number, typically between 0 and 1.
func (client EventsClient) Reward(ctx context.Context, eventID string, reward RewardRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventsClient.Reward")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: eventID,
			Constraints: []validation.Constraint{{Target: "eventID", Name: validation.MaxLength, Rule: 256, Chain: nil}}},
		{TargetValue: reward,
			Constraints: []validation.Constraint{{Target: "reward.Value", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("personalizer.EventsClient", "Reward", err.Error())
	}

	req, err := client.RewardPreparer(ctx, eventID, reward)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.EventsClient", "Reward", nil, "Failure preparing request")
		return
	}

	resp, err := client.RewardSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "personalizer.EventsClient", "Reward", resp, "Failure sending request")
		return
	}

	result, err = client.RewardResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "personalizer.EventsClient", "Reward", resp, "Failure responding to request")
	}

	return
}

// RewardPreparer prepares the Reward request.
func (client EventsClient) RewardPreparer(ctx context.Context, eventID string, reward RewardRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"eventId": autorest.Encode("path", eventID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{Endpoint}/personalizer/v1.0", urlParameters),
		autorest.WithPathParameters("/events/{eventId}/reward", pathParameters),
		autorest.WithJSON(reward))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RewardSender sends the Reward request. The method will close the
// http.Response Body if it receives an error.
func (client EventsClient) RewardSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RewardResponder handles the response to the Reward request. The method always
// closes the http.Response Body.
func (client EventsClient) RewardResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}