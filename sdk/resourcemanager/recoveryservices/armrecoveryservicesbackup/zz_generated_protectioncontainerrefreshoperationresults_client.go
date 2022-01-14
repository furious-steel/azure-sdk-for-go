//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// ProtectionContainerRefreshOperationResultsClient contains the methods for the ProtectionContainerRefreshOperationResults group.
// Don't use this type directly, use NewProtectionContainerRefreshOperationResultsClient() instead.
type ProtectionContainerRefreshOperationResultsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProtectionContainerRefreshOperationResultsClient creates a new instance of ProtectionContainerRefreshOperationResultsClient with the specified values.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProtectionContainerRefreshOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ProtectionContainerRefreshOperationResultsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &ProtectionContainerRefreshOperationResultsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Provides the result of the refresh operation triggered by the BeginRefresh operation.
// If the operation fails it returns an *azcore.ResponseError type.
// vaultName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// fabricName - Fabric name associated with the container.
// operationID - Operation ID associated with the operation whose result needs to be fetched.
// options - ProtectionContainerRefreshOperationResultsClientGetOptions contains the optional parameters for the ProtectionContainerRefreshOperationResultsClient.Get
// method.
func (client *ProtectionContainerRefreshOperationResultsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string, options *ProtectionContainerRefreshOperationResultsClientGetOptions) (ProtectionContainerRefreshOperationResultsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, fabricName, operationID, options)
	if err != nil {
		return ProtectionContainerRefreshOperationResultsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProtectionContainerRefreshOperationResultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return ProtectionContainerRefreshOperationResultsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return ProtectionContainerRefreshOperationResultsClientGetResponse{RawResponse: resp}, nil
}

// getCreateRequest creates the Get request.
func (client *ProtectionContainerRefreshOperationResultsClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string, options *ProtectionContainerRefreshOperationResultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/operationResults/{operationId}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}