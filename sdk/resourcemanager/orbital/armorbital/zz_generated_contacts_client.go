//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

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

// ContactsClient contains the methods for the Contacts group.
// Don't use this type directly, use NewContactsClient() instead.
type ContactsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewContactsClient creates a new instance of ContactsClient with the specified values.
func NewContactsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ContactsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ContactsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Creates a contact.
// If the operation fails it returns the *CloudError error type.
func (client *ContactsClient) BeginCreate(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, parameters Contact, options *ContactsBeginCreateOptions) (ContactsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, spacecraftName, contactName, parameters, options)
	if err != nil {
		return ContactsCreatePollerResponse{}, err
	}
	result := ContactsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ContactsClient.Create", "azure-async-operation", resp, client.pl, client.createHandleError)
	if err != nil {
		return ContactsCreatePollerResponse{}, err
	}
	result.Poller = &ContactsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a contact.
// If the operation fails it returns the *CloudError error type.
func (client *ContactsClient) create(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, parameters Contact, options *ContactsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, spacecraftName, contactName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ContactsClient) createCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, parameters Contact, options *ContactsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts/{contactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	if contactName == "" {
		return nil, errors.New("parameter contactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactName}", url.PathEscape(contactName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleError handles the Create error response.
func (client *ContactsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a specified contact
// If the operation fails it returns the *CloudError error type.
func (client *ContactsClient) BeginDelete(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsBeginDeleteOptions) (ContactsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, spacecraftName, contactName, options)
	if err != nil {
		return ContactsDeletePollerResponse{}, err
	}
	result := ContactsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ContactsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ContactsDeletePollerResponse{}, err
	}
	result.Poller = &ContactsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a specified contact
// If the operation fails it returns the *CloudError error type.
func (client *ContactsClient) deleteOperation(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, spacecraftName, contactName, options)
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
func (client *ContactsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts/{contactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	if contactName == "" {
		return nil, errors.New("parameter contactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactName}", url.PathEscape(contactName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ContactsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified contact in a specified resource group
// If the operation fails it returns the *CloudError error type.
func (client *ContactsClient) Get(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsGetOptions) (ContactsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, spacecraftName, contactName, options)
	if err != nil {
		return ContactsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContactsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContactsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ContactsClient) getCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, contactName string, options *ContactsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts/{contactName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	if contactName == "" {
		return nil, errors.New("parameter contactName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contactName}", url.PathEscape(contactName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContactsClient) getHandleResponse(resp *http.Response) (ContactsGetResponse, error) {
	result := ContactsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Contact); err != nil {
		return ContactsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ContactsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Returns list of contacts by spacecraftName
// If the operation fails it returns the *CloudError error type.
func (client *ContactsClient) List(ctx context.Context, resourceGroupName string, spacecraftName string, options *ContactsListOptions) (ContactsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, spacecraftName, options)
	if err != nil {
		return ContactsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ContactsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ContactsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ContactsClient) listCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, options *ContactsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/contacts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContactsClient) listHandleResponse(resp *http.Response) (ContactsListResponse, error) {
	result := ContactsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContactListResult); err != nil {
		return ContactsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ContactsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
