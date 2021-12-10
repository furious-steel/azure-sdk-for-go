//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RecommendationTypesClient contains the methods for the RecommendationTypes group.
// Don't use this type directly, use NewRecommendationTypesClient() instead.
type RecommendationTypesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRecommendationTypesClient creates a new instance of RecommendationTypesClient with the specified values.
func NewRecommendationTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RecommendationTypesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RecommendationTypesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get IoT recommendation type
// If the operation fails it returns the *ErrorResponse error type.
func (client *RecommendationTypesClient) Get(ctx context.Context, recommendationTypeName string, options *RecommendationTypesGetOptions) (RecommendationTypesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, recommendationTypeName, options)
	if err != nil {
		return RecommendationTypesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendationTypesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecommendationTypesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecommendationTypesClient) getCreateRequest(ctx context.Context, recommendationTypeName string, options *RecommendationTypesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/recommendationTypes/{recommendationTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if recommendationTypeName == "" {
		return nil, errors.New("parameter recommendationTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recommendationTypeName}", url.PathEscape(recommendationTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecommendationTypesClient) getHandleResponse(resp *http.Response) (RecommendationTypesGetResponse, error) {
	result := RecommendationTypesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendationType); err != nil {
		return RecommendationTypesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RecommendationTypesClient) getHandleError(resp *http.Response) error {
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

// List - List IoT recommendation types
// If the operation fails it returns the *ErrorResponse error type.
func (client *RecommendationTypesClient) List(ctx context.Context, options *RecommendationTypesListOptions) (RecommendationTypesListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return RecommendationTypesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RecommendationTypesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecommendationTypesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RecommendationTypesClient) listCreateRequest(ctx context.Context, options *RecommendationTypesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/recommendationTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecommendationTypesClient) listHandleResponse(resp *http.Response) (RecommendationTypesListResponse, error) {
	result := RecommendationTypesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecommendationTypeList); err != nil {
		return RecommendationTypesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RecommendationTypesClient) listHandleError(resp *http.Response) error {
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