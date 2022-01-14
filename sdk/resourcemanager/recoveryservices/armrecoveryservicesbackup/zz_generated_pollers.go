//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ClientBMSPrepareDataMovePoller provides polling facilities until the operation reaches a terminal state.
type ClientBMSPrepareDataMovePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClientBMSPrepareDataMovePoller) Done() bool {
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
func (p *ClientBMSPrepareDataMovePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClientBMSPrepareDataMoveResponse will be returned.
func (p *ClientBMSPrepareDataMovePoller) FinalResponse(ctx context.Context) (ClientBMSPrepareDataMoveResponse, error) {
	respType := ClientBMSPrepareDataMoveResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ClientBMSPrepareDataMoveResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClientBMSPrepareDataMovePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ClientBMSTriggerDataMovePoller provides polling facilities until the operation reaches a terminal state.
type ClientBMSTriggerDataMovePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClientBMSTriggerDataMovePoller) Done() bool {
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
func (p *ClientBMSTriggerDataMovePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClientBMSTriggerDataMoveResponse will be returned.
func (p *ClientBMSTriggerDataMovePoller) FinalResponse(ctx context.Context) (ClientBMSTriggerDataMoveResponse, error) {
	respType := ClientBMSTriggerDataMoveResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ClientBMSTriggerDataMoveResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClientBMSTriggerDataMovePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ClientMoveRecoveryPointPoller provides polling facilities until the operation reaches a terminal state.
type ClientMoveRecoveryPointPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ClientMoveRecoveryPointPoller) Done() bool {
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
func (p *ClientMoveRecoveryPointPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ClientMoveRecoveryPointResponse will be returned.
func (p *ClientMoveRecoveryPointPoller) FinalResponse(ctx context.Context) (ClientMoveRecoveryPointResponse, error) {
	respType := ClientMoveRecoveryPointResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ClientMoveRecoveryPointResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ClientMoveRecoveryPointPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionClientDeletePoller) Done() bool {
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
func (p *PrivateEndpointConnectionClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionClientDeleteResponse will be returned.
func (p *PrivateEndpointConnectionClientDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionClientDeleteResponse, error) {
	respType := PrivateEndpointConnectionClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return PrivateEndpointConnectionClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionClientPutPoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionClientPutPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionClientPutPoller) Done() bool {
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
func (p *PrivateEndpointConnectionClientPutPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionClientPutResponse will be returned.
func (p *PrivateEndpointConnectionClientPutPoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionClientPutResponse, error) {
	respType := PrivateEndpointConnectionClientPutResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnectionResource)
	if err != nil {
		return PrivateEndpointConnectionClientPutResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionClientPutPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ProtectionPoliciesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type ProtectionPoliciesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ProtectionPoliciesClientDeletePoller) Done() bool {
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
func (p *ProtectionPoliciesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ProtectionPoliciesClientDeleteResponse will be returned.
func (p *ProtectionPoliciesClientDeletePoller) FinalResponse(ctx context.Context) (ProtectionPoliciesClientDeleteResponse, error) {
	respType := ProtectionPoliciesClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ProtectionPoliciesClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ProtectionPoliciesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// RestoresClientTriggerPoller provides polling facilities until the operation reaches a terminal state.
type RestoresClientTriggerPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *RestoresClientTriggerPoller) Done() bool {
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
func (p *RestoresClientTriggerPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final RestoresClientTriggerResponse will be returned.
func (p *RestoresClientTriggerPoller) FinalResponse(ctx context.Context) (RestoresClientTriggerResponse, error) {
	respType := RestoresClientTriggerResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RestoresClientTriggerResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *RestoresClientTriggerPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ValidateOperationClientTriggerPoller provides polling facilities until the operation reaches a terminal state.
type ValidateOperationClientTriggerPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ValidateOperationClientTriggerPoller) Done() bool {
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
func (p *ValidateOperationClientTriggerPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ValidateOperationClientTriggerResponse will be returned.
func (p *ValidateOperationClientTriggerPoller) FinalResponse(ctx context.Context) (ValidateOperationClientTriggerResponse, error) {
	respType := ValidateOperationClientTriggerResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ValidateOperationClientTriggerResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ValidateOperationClientTriggerPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}