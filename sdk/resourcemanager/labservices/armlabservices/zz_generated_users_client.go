//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

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

// UsersClient contains the methods for the Users group.
// Don't use this type directly, use NewUsersClient() instead.
type UsersClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewUsersClient creates a new instance of UsersClient with the specified values.
func NewUsersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *UsersClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &UsersClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Operation to create or update a lab user.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, body User, options *UsersBeginCreateOrUpdateOptions) (UsersCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return UsersCreateOrUpdatePollerResponse{}, err
	}
	result := UsersCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.CreateOrUpdate", "original-uri", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return UsersCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &UsersCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Operation to create or update a lab user.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, body User, options *UsersBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *UsersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, body User, options *UsersBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *UsersClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Operation to delete a user resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersBeginDeleteOptions) (UsersDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, labName, userName, options)
	if err != nil {
		return UsersDeletePollerResponse{}, err
	}
	result := UsersDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return UsersDeletePollerResponse{}, err
	}
	result.Poller = &UsersDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Operation to delete a user resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, options)
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
func (client *UsersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *UsersClient) deleteHandleError(resp *http.Response) error {
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

// Get - Returns the properties of a lab user.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersGetOptions) (UsersGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, options)
	if err != nil {
		return UsersGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UsersGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UsersGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *UsersClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *UsersGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UsersClient) getHandleResponse(resp *http.Response) (UsersGetResponse, error) {
	result := UsersGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return UsersGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *UsersClient) getHandleError(resp *http.Response) error {
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

// BeginInvite - Operation to invite a user to a lab.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) BeginInvite(ctx context.Context, resourceGroupName string, labName string, userName string, body InviteBody, options *UsersBeginInviteOptions) (UsersInvitePollerResponse, error) {
	resp, err := client.invite(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return UsersInvitePollerResponse{}, err
	}
	result := UsersInvitePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.Invite", "location", resp, client.pl, client.inviteHandleError)
	if err != nil {
		return UsersInvitePollerResponse{}, err
	}
	result.Poller = &UsersInvitePoller{
		pt: pt,
	}
	return result, nil
}

// Invite - Operation to invite a user to a lab.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) invite(ctx context.Context, resourceGroupName string, labName string, userName string, body InviteBody, options *UsersBeginInviteOptions) (*http.Response, error) {
	req, err := client.inviteCreateRequest(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.inviteHandleError(resp)
	}
	return resp, nil
}

// inviteCreateRequest creates the Invite request.
func (client *UsersClient) inviteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, body InviteBody, options *UsersBeginInviteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}/invite"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// inviteHandleError handles the Invite error response.
func (client *UsersClient) inviteHandleError(resp *http.Response) error {
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

// ListByLab - Returns a list of all users for a lab.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) ListByLab(resourceGroupName string, labName string, options *UsersListByLabOptions) *UsersListByLabPager {
	return &UsersListByLabPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLabCreateRequest(ctx, resourceGroupName, labName, options)
		},
		advancer: func(ctx context.Context, resp UsersListByLabResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PagedUsers.NextLink)
		},
	}
}

// listByLabCreateRequest creates the ListByLab request.
func (client *UsersClient) listByLabCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *UsersListByLabOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLabHandleResponse handles the ListByLab response.
func (client *UsersClient) listByLabHandleResponse(resp *http.Response) (UsersListByLabResponse, error) {
	result := UsersListByLabResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PagedUsers); err != nil {
		return UsersListByLabResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByLabHandleError handles the ListByLab error response.
func (client *UsersClient) listByLabHandleError(resp *http.Response) error {
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

// BeginUpdate - Operation to update a lab user.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) BeginUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, body UserUpdate, options *UsersBeginUpdateOptions) (UsersUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return UsersUpdatePollerResponse{}, err
	}
	result := UsersUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("UsersClient.Update", "location", resp, client.pl, client.updateHandleError)
	if err != nil {
		return UsersUpdatePollerResponse{}, err
	}
	result.Poller = &UsersUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Operation to update a lab user.
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsersClient) update(ctx context.Context, resourceGroupName string, labName string, userName string, body UserUpdate, options *UsersBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *UsersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, body UserUpdate, options *UsersBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labs/{labName}/users/{userName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// updateHandleError handles the Update error response.
func (client *UsersClient) updateHandleError(resp *http.Response) error {
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