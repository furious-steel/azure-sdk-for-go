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

// SQLPoolTableColumnsClient contains the methods for the SQLPoolTableColumns group.
// Don't use this type directly, use NewSQLPoolTableColumnsClient() instead.
type SQLPoolTableColumnsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLPoolTableColumnsClient creates a new instance of SQLPoolTableColumnsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLPoolTableColumnsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SQLPoolTableColumnsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SQLPoolTableColumnsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ListByTableName - Gets columns in a given table in a SQL pool.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// schemaName - The name of the schema.
// tableName - The name of the table.
// options - SQLPoolTableColumnsClientListByTableNameOptions contains the optional parameters for the SQLPoolTableColumnsClient.ListByTableName
// method.
func (client *SQLPoolTableColumnsClient) ListByTableName(resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, options *SQLPoolTableColumnsClientListByTableNameOptions) *SQLPoolTableColumnsClientListByTableNamePager {
	return &SQLPoolTableColumnsClientListByTableNamePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByTableNameCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, schemaName, tableName, options)
		},
		advancer: func(ctx context.Context, resp SQLPoolTableColumnsClientListByTableNameResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SQLPoolColumnListResult.NextLink)
		},
	}
}

// listByTableNameCreateRequest creates the ListByTableName request.
func (client *SQLPoolTableColumnsClient) listByTableNameCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, options *SQLPoolTableColumnsClientListByTableNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/schemas/{schemaName}/tables/{tableName}/columns"
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
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByTableNameHandleResponse handles the ListByTableName response.
func (client *SQLPoolTableColumnsClient) listByTableNameHandleResponse(resp *http.Response) (SQLPoolTableColumnsClientListByTableNameResponse, error) {
	result := SQLPoolTableColumnsClientListByTableNameResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLPoolColumnListResult); err != nil {
		return SQLPoolTableColumnsClientListByTableNameResponse{}, err
	}
	return result, nil
}