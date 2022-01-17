//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// DedicatedCloudNodesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DedicatedCloudNodesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DedicatedCloudNodesClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *DedicatedCloudNodesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DedicatedCloudNodesClientCreateOrUpdateResponse will be returned.
func (p *DedicatedCloudNodesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (DedicatedCloudNodesClientCreateOrUpdateResponse, error) {
	respType := DedicatedCloudNodesClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.DedicatedCloudNode)
	if err != nil {
		return DedicatedCloudNodesClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DedicatedCloudNodesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DedicatedCloudServicesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type DedicatedCloudServicesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DedicatedCloudServicesClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *DedicatedCloudServicesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DedicatedCloudServicesClientDeleteResponse will be returned.
func (p *DedicatedCloudServicesClientDeletePoller) FinalResponse(ctx context.Context) (DedicatedCloudServicesClientDeleteResponse, error) {
	respType := DedicatedCloudServicesClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return DedicatedCloudServicesClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DedicatedCloudServicesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachinesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachinesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachinesClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachinesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachinesClientCreateOrUpdateResponse will be returned.
func (p *VirtualMachinesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VirtualMachinesClientCreateOrUpdateResponse, error) {
	respType := VirtualMachinesClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VirtualMachine)
	if err != nil {
		return VirtualMachinesClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachinesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachinesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachinesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachinesClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachinesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachinesClientDeleteResponse will be returned.
func (p *VirtualMachinesClientDeletePoller) FinalResponse(ctx context.Context) (VirtualMachinesClientDeleteResponse, error) {
	respType := VirtualMachinesClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachinesClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachinesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachinesClientStartPoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachinesClientStartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachinesClientStartPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachinesClientStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachinesClientStartResponse will be returned.
func (p *VirtualMachinesClientStartPoller) FinalResponse(ctx context.Context) (VirtualMachinesClientStartResponse, error) {
	respType := VirtualMachinesClientStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachinesClientStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachinesClientStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachinesClientStopPoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachinesClientStopPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachinesClientStopPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachinesClientStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachinesClientStopResponse will be returned.
func (p *VirtualMachinesClientStopPoller) FinalResponse(ctx context.Context) (VirtualMachinesClientStopResponse, error) {
	respType := VirtualMachinesClientStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VirtualMachinesClientStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachinesClientStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VirtualMachinesClientUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VirtualMachinesClientUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VirtualMachinesClientUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VirtualMachinesClientUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VirtualMachinesClientUpdateResponse will be returned.
func (p *VirtualMachinesClientUpdatePoller) FinalResponse(ctx context.Context) (VirtualMachinesClientUpdateResponse, error) {
	respType := VirtualMachinesClientUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VirtualMachine)
	if err != nil {
		return VirtualMachinesClientUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VirtualMachinesClientUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}