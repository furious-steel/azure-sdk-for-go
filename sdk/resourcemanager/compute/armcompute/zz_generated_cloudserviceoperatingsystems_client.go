//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// CloudServiceOperatingSystemsClient contains the methods for the CloudServiceOperatingSystems group.
// Don't use this type directly, use NewCloudServiceOperatingSystemsClient() instead.
type CloudServiceOperatingSystemsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCloudServiceOperatingSystemsClient creates a new instance of CloudServiceOperatingSystemsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCloudServiceOperatingSystemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CloudServiceOperatingSystemsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &CloudServiceOperatingSystemsClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// GetOSFamily - Gets properties of a guest operating system family that can be specified in the XML service configuration
// (.cscfg) for a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Name of the location that the OS family pertains to.
// osFamilyName - Name of the OS family.
// options - CloudServiceOperatingSystemsClientGetOSFamilyOptions contains the optional parameters for the CloudServiceOperatingSystemsClient.GetOSFamily
// method.
func (client *CloudServiceOperatingSystemsClient) GetOSFamily(ctx context.Context, location string, osFamilyName string, options *CloudServiceOperatingSystemsClientGetOSFamilyOptions) (CloudServiceOperatingSystemsClientGetOSFamilyResponse, error) {
	req, err := client.getOSFamilyCreateRequest(ctx, location, osFamilyName, options)
	if err != nil {
		return CloudServiceOperatingSystemsClientGetOSFamilyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceOperatingSystemsClientGetOSFamilyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceOperatingSystemsClientGetOSFamilyResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOSFamilyHandleResponse(resp)
}

// getOSFamilyCreateRequest creates the GetOSFamily request.
func (client *CloudServiceOperatingSystemsClient) getOSFamilyCreateRequest(ctx context.Context, location string, osFamilyName string, options *CloudServiceOperatingSystemsClientGetOSFamilyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsFamilies/{osFamilyName}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if osFamilyName == "" {
		return nil, errors.New("parameter osFamilyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{osFamilyName}", url.PathEscape(osFamilyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOSFamilyHandleResponse handles the GetOSFamily response.
func (client *CloudServiceOperatingSystemsClient) getOSFamilyHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsClientGetOSFamilyResponse, error) {
	result := CloudServiceOperatingSystemsClientGetOSFamilyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSFamily); err != nil {
		return CloudServiceOperatingSystemsClientGetOSFamilyResponse{}, err
	}
	return result, nil
}

// GetOSVersion - Gets properties of a guest operating system version that can be specified in the XML service configuration
// (.cscfg) for a cloud service.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Name of the location that the OS version pertains to.
// osVersionName - Name of the OS version.
// options - CloudServiceOperatingSystemsClientGetOSVersionOptions contains the optional parameters for the CloudServiceOperatingSystemsClient.GetOSVersion
// method.
func (client *CloudServiceOperatingSystemsClient) GetOSVersion(ctx context.Context, location string, osVersionName string, options *CloudServiceOperatingSystemsClientGetOSVersionOptions) (CloudServiceOperatingSystemsClientGetOSVersionResponse, error) {
	req, err := client.getOSVersionCreateRequest(ctx, location, osVersionName, options)
	if err != nil {
		return CloudServiceOperatingSystemsClientGetOSVersionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceOperatingSystemsClientGetOSVersionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceOperatingSystemsClientGetOSVersionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOSVersionHandleResponse(resp)
}

// getOSVersionCreateRequest creates the GetOSVersion request.
func (client *CloudServiceOperatingSystemsClient) getOSVersionCreateRequest(ctx context.Context, location string, osVersionName string, options *CloudServiceOperatingSystemsClientGetOSVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsVersions/{osVersionName}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if osVersionName == "" {
		return nil, errors.New("parameter osVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{osVersionName}", url.PathEscape(osVersionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOSVersionHandleResponse handles the GetOSVersion response.
func (client *CloudServiceOperatingSystemsClient) getOSVersionHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsClientGetOSVersionResponse, error) {
	result := CloudServiceOperatingSystemsClientGetOSVersionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSVersion); err != nil {
		return CloudServiceOperatingSystemsClientGetOSVersionResponse{}, err
	}
	return result, nil
}

// ListOSFamilies - Gets a list of all guest operating system families available to be specified in the XML service configuration
// (.cscfg) for a cloud service. Use nextLink property in the response to get the next page
// of OS Families. Do this till nextLink is null to fetch all the OS Families.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Name of the location that the OS families pertain to.
// options - CloudServiceOperatingSystemsClientListOSFamiliesOptions contains the optional parameters for the CloudServiceOperatingSystemsClient.ListOSFamilies
// method.
func (client *CloudServiceOperatingSystemsClient) ListOSFamilies(location string, options *CloudServiceOperatingSystemsClientListOSFamiliesOptions) *CloudServiceOperatingSystemsClientListOSFamiliesPager {
	return &CloudServiceOperatingSystemsClientListOSFamiliesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOSFamiliesCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp CloudServiceOperatingSystemsClientListOSFamiliesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OSFamilyListResult.NextLink)
		},
	}
}

// listOSFamiliesCreateRequest creates the ListOSFamilies request.
func (client *CloudServiceOperatingSystemsClient) listOSFamiliesCreateRequest(ctx context.Context, location string, options *CloudServiceOperatingSystemsClientListOSFamiliesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsFamilies"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOSFamiliesHandleResponse handles the ListOSFamilies response.
func (client *CloudServiceOperatingSystemsClient) listOSFamiliesHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsClientListOSFamiliesResponse, error) {
	result := CloudServiceOperatingSystemsClientListOSFamiliesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSFamilyListResult); err != nil {
		return CloudServiceOperatingSystemsClientListOSFamiliesResponse{}, err
	}
	return result, nil
}

// ListOSVersions - Gets a list of all guest operating system versions available to be specified in the XML service configuration
// (.cscfg) for a cloud service. Use nextLink property in the response to get the next page
// of OS versions. Do this till nextLink is null to fetch all the OS versions.
// If the operation fails it returns an *azcore.ResponseError type.
// location - Name of the location that the OS versions pertain to.
// options - CloudServiceOperatingSystemsClientListOSVersionsOptions contains the optional parameters for the CloudServiceOperatingSystemsClient.ListOSVersions
// method.
func (client *CloudServiceOperatingSystemsClient) ListOSVersions(location string, options *CloudServiceOperatingSystemsClientListOSVersionsOptions) *CloudServiceOperatingSystemsClientListOSVersionsPager {
	return &CloudServiceOperatingSystemsClientListOSVersionsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOSVersionsCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp CloudServiceOperatingSystemsClientListOSVersionsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OSVersionListResult.NextLink)
		},
	}
}

// listOSVersionsCreateRequest creates the ListOSVersions request.
func (client *CloudServiceOperatingSystemsClient) listOSVersionsCreateRequest(ctx context.Context, location string, options *CloudServiceOperatingSystemsClientListOSVersionsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/cloudServiceOsVersions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOSVersionsHandleResponse handles the ListOSVersions response.
func (client *CloudServiceOperatingSystemsClient) listOSVersionsHandleResponse(resp *http.Response) (CloudServiceOperatingSystemsClientListOSVersionsResponse, error) {
	result := CloudServiceOperatingSystemsClientListOSVersionsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSVersionListResult); err != nil {
		return CloudServiceOperatingSystemsClientListOSVersionsResponse{}, err
	}
	return result, nil
}
