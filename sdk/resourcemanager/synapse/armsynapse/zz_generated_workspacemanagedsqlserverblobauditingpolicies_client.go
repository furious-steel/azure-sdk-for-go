//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// WorkspaceManagedSQLServerBlobAuditingPoliciesClient contains the methods for the WorkspaceManagedSQLServerBlobAuditingPolicies group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient() instead.
type WorkspaceManagedSQLServerBlobAuditingPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient creates a new instance of WorkspaceManagedSQLServerBlobAuditingPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerBlobAuditingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspaceManagedSQLServerBlobAuditingPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WorkspaceManagedSQLServerBlobAuditingPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or Update a workspace managed sql server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// blobAuditingPolicyName - The name of the blob auditing policy.
// parameters - Properties of extended blob auditing policy.
// options - WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters
// for the WorkspaceManagedSQLServerBlobAuditingPoliciesClient.BeginCreateOrUpdate method.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspaceManagedSQLServerBlobAuditingPoliciesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspaceManagedSQLServerBlobAuditingPoliciesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or Update a workspace managed sql server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings/{blobAuditingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Get a workspace managed sql server's blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// blobAuditingPolicyName - The name of the blob auditing policy.
// options - WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetOptions contains the optional parameters for the WorkspaceManagedSQLServerBlobAuditingPoliciesClient.Get
// method.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetOptions) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, options)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings/{blobAuditingPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse, error) {
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicy); err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - List workspace managed sql server's blob auditing policies.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceOptions contains the optional parameters for
// the WorkspaceManagedSQLServerBlobAuditingPoliciesClient.ListByWorkspace method.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) ListByWorkspace(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceOptions) *WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspacePager {
	return &WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
		},
		advancer: func(ctx context.Context, resp WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerBlobAuditingPolicyListResult.NextLink)
		},
	}
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/auditingSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *WorkspaceManagedSQLServerBlobAuditingPoliciesClient) listByWorkspaceHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse, error) {
	result := WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerBlobAuditingPolicyListResult); err != nil {
		return WorkspaceManagedSQLServerBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}