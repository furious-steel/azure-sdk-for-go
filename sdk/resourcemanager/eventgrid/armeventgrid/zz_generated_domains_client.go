//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DomainsClient contains the methods for the Domains group.
// Don't use this type directly, use NewDomainsClient() instead.
type DomainsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDomainsClient creates a new instance of DomainsClient with the specified values.
func NewDomainsClient(con *arm.Connection, subscriptionID string) *DomainsClient {
	return &DomainsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Asynchronously creates or updates a new domain with the specified parameters.
// If the operation fails it returns a generic error.
func (client *DomainsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainInfo Domain, options *DomainsBeginCreateOrUpdateOptions) (DomainsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, domainName, domainInfo, options)
	if err != nil {
		return DomainsCreateOrUpdatePollerResponse{}, err
	}
	result := DomainsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return DomainsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &DomainsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Asynchronously creates or updates a new domain with the specified parameters.
// If the operation fails it returns a generic error.
func (client *DomainsClient) createOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainInfo Domain, options *DomainsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, domainName, domainInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DomainsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainInfo Domain, options *DomainsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainInfo)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DomainsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete existing domain.
// If the operation fails it returns a generic error.
func (client *DomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, domainName string, options *DomainsBeginDeleteOptions) (DomainsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, domainName, options)
	if err != nil {
		return DomainsDeletePollerResponse{}, err
	}
	result := DomainsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return DomainsDeletePollerResponse{}, err
	}
	result.Poller = &DomainsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete existing domain.
// If the operation fails it returns a generic error.
func (client *DomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, domainName string, options *DomainsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DomainsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get properties of a domain.
// If the operation fails it returns a generic error.
func (client *DomainsClient) Get(ctx context.Context, resourceGroupName string, domainName string, options *DomainsGetOptions) (DomainsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainName, options)
	if err != nil {
		return DomainsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DomainsClient) getHandleResponse(resp *http.Response) (DomainsGetResponse, error) {
	result := DomainsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Domain); err != nil {
		return DomainsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DomainsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByResourceGroup - List all the domains under a resource group.
// If the operation fails it returns a generic error.
func (client *DomainsClient) ListByResourceGroup(resourceGroupName string, options *DomainsListByResourceGroupOptions) *DomainsListByResourceGroupPager {
	return &DomainsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DomainsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DomainsListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DomainsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DomainsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DomainsClient) listByResourceGroupHandleResponse(resp *http.Response) (DomainsListByResourceGroupResponse, error) {
	result := DomainsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainsListResult); err != nil {
		return DomainsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DomainsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListBySubscription - List all the domains under an Azure subscription.
// If the operation fails it returns a generic error.
func (client *DomainsClient) ListBySubscription(options *DomainsListBySubscriptionOptions) *DomainsListBySubscriptionPager {
	return &DomainsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DomainsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DomainsListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DomainsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DomainsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/domains"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DomainsClient) listBySubscriptionHandleResponse(resp *http.Response) (DomainsListBySubscriptionResponse, error) {
	result := DomainsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainsListResult); err != nil {
		return DomainsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *DomainsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListSharedAccessKeys - List the two keys used to publish to a domain.
// If the operation fails it returns a generic error.
func (client *DomainsClient) ListSharedAccessKeys(ctx context.Context, resourceGroupName string, domainName string, options *DomainsListSharedAccessKeysOptions) (DomainsListSharedAccessKeysResponse, error) {
	req, err := client.listSharedAccessKeysCreateRequest(ctx, resourceGroupName, domainName, options)
	if err != nil {
		return DomainsListSharedAccessKeysResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainsListSharedAccessKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainsListSharedAccessKeysResponse{}, client.listSharedAccessKeysHandleError(resp)
	}
	return client.listSharedAccessKeysHandleResponse(resp)
}

// listSharedAccessKeysCreateRequest creates the ListSharedAccessKeys request.
func (client *DomainsClient) listSharedAccessKeysCreateRequest(ctx context.Context, resourceGroupName string, domainName string, options *DomainsListSharedAccessKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSharedAccessKeysHandleResponse handles the ListSharedAccessKeys response.
func (client *DomainsClient) listSharedAccessKeysHandleResponse(resp *http.Response) (DomainsListSharedAccessKeysResponse, error) {
	result := DomainsListSharedAccessKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainSharedAccessKeys); err != nil {
		return DomainsListSharedAccessKeysResponse{}, err
	}
	return result, nil
}

// listSharedAccessKeysHandleError handles the ListSharedAccessKeys error response.
func (client *DomainsClient) listSharedAccessKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// RegenerateKey - Regenerate a shared access key for a domain.
// If the operation fails it returns a generic error.
func (client *DomainsClient) RegenerateKey(ctx context.Context, resourceGroupName string, domainName string, regenerateKeyRequest DomainRegenerateKeyRequest, options *DomainsRegenerateKeyOptions) (DomainsRegenerateKeyResponse, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, domainName, regenerateKeyRequest, options)
	if err != nil {
		return DomainsRegenerateKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DomainsRegenerateKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DomainsRegenerateKeyResponse{}, client.regenerateKeyHandleError(resp)
	}
	return client.regenerateKeyHandleResponse(resp)
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *DomainsClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, domainName string, regenerateKeyRequest DomainRegenerateKeyRequest, options *DomainsRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, regenerateKeyRequest)
}

// regenerateKeyHandleResponse handles the RegenerateKey response.
func (client *DomainsClient) regenerateKeyHandleResponse(resp *http.Response) (DomainsRegenerateKeyResponse, error) {
	result := DomainsRegenerateKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainSharedAccessKeys); err != nil {
		return DomainsRegenerateKeyResponse{}, err
	}
	return result, nil
}

// regenerateKeyHandleError handles the RegenerateKey error response.
func (client *DomainsClient) regenerateKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Asynchronously updates a domain with the specified parameters.
// If the operation fails it returns a generic error.
func (client *DomainsClient) BeginUpdate(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters DomainUpdateParameters, options *DomainsBeginUpdateOptions) (DomainsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, domainName, domainUpdateParameters, options)
	if err != nil {
		return DomainsUpdatePollerResponse{}, err
	}
	result := DomainsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DomainsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return DomainsUpdatePollerResponse{}, err
	}
	result.Poller = &DomainsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Asynchronously updates a domain with the specified parameters.
// If the operation fails it returns a generic error.
func (client *DomainsClient) update(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters DomainUpdateParameters, options *DomainsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, domainName, domainUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DomainsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters DomainUpdateParameters, options *DomainsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/domains/{domainName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainName == "" {
		return nil, errors.New("parameter domainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainName}", url.PathEscape(domainName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, domainUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *DomainsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}