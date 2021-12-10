//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkubernetesconfiguration

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

// ClusterExtensionTypeClient contains the methods for the ClusterExtensionType group.
// Don't use this type directly, use NewClusterExtensionTypeClient() instead.
type ClusterExtensionTypeClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewClusterExtensionTypeClient creates a new instance of ClusterExtensionTypeClient with the specified values.
func NewClusterExtensionTypeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ClusterExtensionTypeClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ClusterExtensionTypeClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get Extension Type details
// If the operation fails it returns the *ErrorResponse error type.
func (client *ClusterExtensionTypeClient) Get(ctx context.Context, resourceGroupName string, clusterRp Enum0, clusterResourceName Enum1, clusterName string, extensionTypeName string, options *ClusterExtensionTypeGetOptions) (ClusterExtensionTypeGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterRp, clusterResourceName, clusterName, extensionTypeName, options)
	if err != nil {
		return ClusterExtensionTypeGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ClusterExtensionTypeGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClusterExtensionTypeGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ClusterExtensionTypeClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterRp Enum0, clusterResourceName Enum1, clusterName string, extensionTypeName string, options *ClusterExtensionTypeGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{clusterRp}/{clusterResourceName}/{clusterName}/providers/Microsoft.KubernetesConfiguration/extensionTypes/{extensionTypeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterRp == "" {
		return nil, errors.New("parameter clusterRp cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterRp}", url.PathEscape(string(clusterRp)))
	if clusterResourceName == "" {
		return nil, errors.New("parameter clusterResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterResourceName}", url.PathEscape(string(clusterResourceName)))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if extensionTypeName == "" {
		return nil, errors.New("parameter extensionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extensionTypeName}", url.PathEscape(extensionTypeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClusterExtensionTypeClient) getHandleResponse(resp *http.Response) (ClusterExtensionTypeGetResponse, error) {
	result := ClusterExtensionTypeGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtensionType); err != nil {
		return ClusterExtensionTypeGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ClusterExtensionTypeClient) getHandleError(resp *http.Response) error {
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