//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

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

// ResourceTypeRegistrationsClient contains the methods for the ResourceTypeRegistrations group.
// Don't use this type directly, use NewResourceTypeRegistrationsClient() instead.
type ResourceTypeRegistrationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewResourceTypeRegistrationsClient creates a new instance of ResourceTypeRegistrationsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceTypeRegistrationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ResourceTypeRegistrationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ResourceTypeRegistrationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a resource type.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// resourceType - The resource type.
// properties - The required request body parameters supplied to the resource type registration CreateOrUpdate operation.
// options - ResourceTypeRegistrationsClientBeginCreateOrUpdateOptions contains the optional parameters for the ResourceTypeRegistrationsClient.BeginCreateOrUpdate
// method.
func (client *ResourceTypeRegistrationsClient) BeginCreateOrUpdate(ctx context.Context, providerNamespace string, resourceType string, properties ResourceTypeRegistration, options *ResourceTypeRegistrationsClientBeginCreateOrUpdateOptions) (ResourceTypeRegistrationsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, providerNamespace, resourceType, properties, options)
	if err != nil {
		return ResourceTypeRegistrationsClientCreateOrUpdatePollerResponse{}, err
	}
	result := ResourceTypeRegistrationsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ResourceTypeRegistrationsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return ResourceTypeRegistrationsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ResourceTypeRegistrationsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a resource type.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ResourceTypeRegistrationsClient) createOrUpdate(ctx context.Context, providerNamespace string, resourceType string, properties ResourceTypeRegistration, options *ResourceTypeRegistrationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, providerNamespace, resourceType, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ResourceTypeRegistrationsClient) createOrUpdateCreateRequest(ctx context.Context, providerNamespace string, resourceType string, properties ResourceTypeRegistration, options *ResourceTypeRegistrationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/resourcetypeRegistrations/{resourceType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// Delete - Deletes a resource type
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// resourceType - The resource type.
// options - ResourceTypeRegistrationsClientDeleteOptions contains the optional parameters for the ResourceTypeRegistrationsClient.Delete
// method.
func (client *ResourceTypeRegistrationsClient) Delete(ctx context.Context, providerNamespace string, resourceType string, options *ResourceTypeRegistrationsClientDeleteOptions) (ResourceTypeRegistrationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, providerNamespace, resourceType, options)
	if err != nil {
		return ResourceTypeRegistrationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceTypeRegistrationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ResourceTypeRegistrationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ResourceTypeRegistrationsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ResourceTypeRegistrationsClient) deleteCreateRequest(ctx context.Context, providerNamespace string, resourceType string, options *ResourceTypeRegistrationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/resourcetypeRegistrations/{resourceType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets a resource type details in the given subscription and provider.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// resourceType - The resource type.
// options - ResourceTypeRegistrationsClientGetOptions contains the optional parameters for the ResourceTypeRegistrationsClient.Get
// method.
func (client *ResourceTypeRegistrationsClient) Get(ctx context.Context, providerNamespace string, resourceType string, options *ResourceTypeRegistrationsClientGetOptions) (ResourceTypeRegistrationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, providerNamespace, resourceType, options)
	if err != nil {
		return ResourceTypeRegistrationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceTypeRegistrationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceTypeRegistrationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourceTypeRegistrationsClient) getCreateRequest(ctx context.Context, providerNamespace string, resourceType string, options *ResourceTypeRegistrationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/resourcetypeRegistrations/{resourceType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceTypeRegistrationsClient) getHandleResponse(resp *http.Response) (ResourceTypeRegistrationsClientGetResponse, error) {
	result := ResourceTypeRegistrationsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceTypeRegistration); err != nil {
		return ResourceTypeRegistrationsClientGetResponse{}, err
	}
	return result, nil
}

// ListByProviderRegistration - Gets the list of the resource types for the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
// providerNamespace - The name of the resource provider hosted within ProviderHub.
// options - ResourceTypeRegistrationsClientListByProviderRegistrationOptions contains the optional parameters for the ResourceTypeRegistrationsClient.ListByProviderRegistration
// method.
func (client *ResourceTypeRegistrationsClient) ListByProviderRegistration(providerNamespace string, options *ResourceTypeRegistrationsClientListByProviderRegistrationOptions) *ResourceTypeRegistrationsClientListByProviderRegistrationPager {
	return &ResourceTypeRegistrationsClientListByProviderRegistrationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByProviderRegistrationCreateRequest(ctx, providerNamespace, options)
		},
		advancer: func(ctx context.Context, resp ResourceTypeRegistrationsClientListByProviderRegistrationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ResourceTypeRegistrationArrayResponseWithContinuation.NextLink)
		},
	}
}

// listByProviderRegistrationCreateRequest creates the ListByProviderRegistration request.
func (client *ResourceTypeRegistrationsClient) listByProviderRegistrationCreateRequest(ctx context.Context, providerNamespace string, options *ResourceTypeRegistrationsClientListByProviderRegistrationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/resourcetypeRegistrations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByProviderRegistrationHandleResponse handles the ListByProviderRegistration response.
func (client *ResourceTypeRegistrationsClient) listByProviderRegistrationHandleResponse(resp *http.Response) (ResourceTypeRegistrationsClientListByProviderRegistrationResponse, error) {
	result := ResourceTypeRegistrationsClientListByProviderRegistrationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceTypeRegistrationArrayResponseWithContinuation); err != nil {
		return ResourceTypeRegistrationsClientListByProviderRegistrationResponse{}, err
	}
	return result, nil
}