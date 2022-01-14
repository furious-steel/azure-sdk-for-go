//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// MonitorsClientCreatePoller provides polling facilities until the operation reaches a terminal state.
type MonitorsClientCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MonitorsClientCreatePoller) Done() bool {
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
func (p *MonitorsClientCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MonitorsClientCreateResponse will be returned.
func (p *MonitorsClientCreatePoller) FinalResponse(ctx context.Context) (MonitorsClientCreateResponse, error) {
	respType := MonitorsClientCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.MonitorResource)
	if err != nil {
		return MonitorsClientCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MonitorsClientCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// MonitorsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type MonitorsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *MonitorsClientDeletePoller) Done() bool {
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
func (p *MonitorsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final MonitorsClientDeleteResponse will be returned.
func (p *MonitorsClientDeletePoller) FinalResponse(ctx context.Context) (MonitorsClientDeleteResponse, error) {
	respType := MonitorsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return MonitorsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *MonitorsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SingleSignOnClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type SingleSignOnClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SingleSignOnClientCreateOrUpdatePoller) Done() bool {
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
func (p *SingleSignOnClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SingleSignOnClientCreateOrUpdateResponse will be returned.
func (p *SingleSignOnClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (SingleSignOnClientCreateOrUpdateResponse, error) {
	respType := SingleSignOnClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.SingleSignOnResource)
	if err != nil {
		return SingleSignOnClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SingleSignOnClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubAccountClientCreatePoller provides polling facilities until the operation reaches a terminal state.
type SubAccountClientCreatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubAccountClientCreatePoller) Done() bool {
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
func (p *SubAccountClientCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubAccountClientCreateResponse will be returned.
func (p *SubAccountClientCreatePoller) FinalResponse(ctx context.Context) (SubAccountClientCreateResponse, error) {
	respType := SubAccountClientCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.MonitorResource)
	if err != nil {
		return SubAccountClientCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubAccountClientCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// SubAccountClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type SubAccountClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *SubAccountClientDeletePoller) Done() bool {
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
func (p *SubAccountClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final SubAccountClientDeleteResponse will be returned.
func (p *SubAccountClientDeletePoller) FinalResponse(ctx context.Context) (SubAccountClientDeleteResponse, error) {
	respType := SubAccountClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return SubAccountClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *SubAccountClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}