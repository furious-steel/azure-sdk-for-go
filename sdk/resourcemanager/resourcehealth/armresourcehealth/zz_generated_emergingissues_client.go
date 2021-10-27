//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

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

// EmergingIssuesClient contains the methods for the EmergingIssues group.
// Don't use this type directly, use NewEmergingIssuesClient() instead.
type EmergingIssuesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewEmergingIssuesClient creates a new instance of EmergingIssuesClient with the specified values.
func NewEmergingIssuesClient(con *arm.Connection) *EmergingIssuesClient {
	return &EmergingIssuesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Gets Azure services' emerging issues.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmergingIssuesClient) Get(ctx context.Context, issueName Enum0, options *EmergingIssuesGetOptions) (EmergingIssuesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, issueName, options)
	if err != nil {
		return EmergingIssuesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EmergingIssuesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmergingIssuesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EmergingIssuesClient) getCreateRequest(ctx context.Context, issueName Enum0, options *EmergingIssuesGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/emergingIssues/{issueName}"
	if issueName == "" {
		return nil, errors.New("parameter issueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{issueName}", url.PathEscape(string(issueName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EmergingIssuesClient) getHandleResponse(resp *http.Response) (EmergingIssuesGetResponse, error) {
	result := EmergingIssuesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmergingIssuesGetResult); err != nil {
		return EmergingIssuesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *EmergingIssuesClient) getHandleError(resp *http.Response) error {
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

// List - Lists Azure services' emerging issues.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EmergingIssuesClient) List(options *EmergingIssuesListOptions) *EmergingIssuesListPager {
	return &EmergingIssuesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp EmergingIssuesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EmergingIssueListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *EmergingIssuesClient) listCreateRequest(ctx context.Context, options *EmergingIssuesListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/emergingIssues"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EmergingIssuesClient) listHandleResponse(resp *http.Response) (EmergingIssuesListResponse, error) {
	result := EmergingIssuesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EmergingIssueListResult); err != nil {
		return EmergingIssuesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *EmergingIssuesClient) listHandleError(resp *http.Response) error {
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