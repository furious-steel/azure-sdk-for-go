//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsaas

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// AccessTokenResult - the ISV access token result response.
type AccessTokenResult struct {
	// The Publisher Offer Base Uri
	PublisherOfferBaseURI *string `json:"publisherOfferBaseUri,omitempty"`

	// The generated token
	Token *string `json:"token,omitempty"`
}

// App - the saasApp resource.
type App struct {
	// the resource location.
	Location *string `json:"location,omitempty"`

	// the resource name.
	Name *string `json:"name,omitempty"`

	// the resource properties.
	Properties *AppProperties `json:"properties,omitempty"`

	// the resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// the resource type.
	Type *string `json:"type,omitempty"`

	// READ-ONLY; the resource Id.
	ID *string `json:"id,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type App.
func (a App) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// AppOperation - saas app operations
type AppOperation struct {
	// the operation display
	Display *AppOperationDisplay `json:"display,omitempty"`

	// whether the operation is a data action or not.
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// the operation name
	Name *string `json:"name,omitempty"`

	// the operation origin
	Origin *string `json:"origin,omitempty"`
}

// AppOperationDisplay - Saas app operation display
type AppOperationDisplay struct {
	// Description of the operation for display purposes
	Description *string `json:"description,omitempty"`

	// Name of the operation for display purposes
	Operation *string `json:"operation,omitempty"`

	// Name of the provider for display purposes
	Provider *string `json:"provider,omitempty"`

	// Name of the resource type for display purposes
	Resource *string `json:"resource,omitempty"`
}

// AppOperationsResponseWithContinuation - saas app operation response with continuation.
type AppOperationsResponseWithContinuation struct {
	// the next link to query to get the remaining results.
	NextLink *string `json:"nextLink,omitempty"`

	// the value of response.
	Value []*AppOperation `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type AppOperationsResponseWithContinuation.
func (a AppOperationsResponseWithContinuation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// AppPlan - Saas resource plan.
type AppPlan struct {
	// the plan id.
	Name *string `json:"name,omitempty"`

	// the offer id.
	Product *string `json:"product,omitempty"`

	// the publisher id.
	Publisher *string `json:"publisher,omitempty"`
}

// AppProperties - Saas resource properties.
type AppProperties struct {
	// the resource plan details.
	SaasAppPlan *AppPlan `json:"saasAppPlan,omitempty"`

	// the Saas resource status.
	Status *SaasAppStatus `json:"status,omitempty"`
}

// AppResponseWithContinuation - saas app response with continuation.
type AppResponseWithContinuation struct {
	// the next link to query to get the remaining results.
	NextLink *string `json:"nextLink,omitempty"`

	// the value of response.
	Value []*App `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type AppResponseWithContinuation.
func (a AppResponseWithContinuation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", a.NextLink)
	populate(objectMap, "value", a.Value)
	return json.Marshal(objectMap)
}

// ApplicationsClientListOptions contains the optional parameters for the ApplicationsClient.List method.
type ApplicationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ClientBeginCreateResourceOptions contains the optional parameters for the Client.BeginCreateResource method.
type ClientBeginCreateResourceOptions struct {
	// placeholder for future optional parameters
}

// ClientBeginDeleteOptions contains the optional parameters for the Client.BeginDelete method.
type ClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// ClientBeginUpdateResourceOptions contains the optional parameters for the Client.BeginUpdateResource method.
type ClientBeginUpdateResourceOptions struct {
	// placeholder for future optional parameters
}

// ClientGetResourceOptions contains the optional parameters for the Client.GetResource method.
type ClientGetResourceOptions struct {
	// placeholder for future optional parameters
}

// CreationProperties - properties for creation saas
type CreationProperties struct {
	// Whether the SaaS subscription will auto renew upon term end.
	AutoRenew *bool `json:"autoRenew,omitempty"`

	// The offer id.
	OfferID *string `json:"offerId,omitempty"`

	// The metadata about the SaaS subscription such as the AzureSubscriptionId and ResourceUri.
	PaymentChannelMetadata map[string]*string `json:"paymentChannelMetadata,omitempty"`

	// The Payment channel for the SaasSubscription.
	PaymentChannelType *PaymentChannelType `json:"paymentChannelType,omitempty"`

	// The publisher id.
	PublisherID *string `json:"publisherId,omitempty"`

	// The environment in the publisher side for this resource.
	PublisherTestEnvironment *string `json:"publisherTestEnvironment,omitempty"`

	// The seat count.
	Quantity *float32 `json:"quantity,omitempty"`

	// The plan id.
	SKUID *string `json:"skuId,omitempty"`

	// The SaaS resource name.
	SaasResourceName *string `json:"saasResourceName,omitempty"`

	// The saas session id used for dev service migration request.
	SaasSessionID *string `json:"saasSessionId,omitempty"`

	// The saas subscription id used for tenant to subscription level migration request.
	SaasSubscriptionID *string `json:"saasSubscriptionId,omitempty"`

	// The current Term id.
	TermID *string `json:"termId,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CreationProperties.
func (c CreationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoRenew", c.AutoRenew)
	populate(objectMap, "offerId", c.OfferID)
	populate(objectMap, "paymentChannelMetadata", c.PaymentChannelMetadata)
	populate(objectMap, "paymentChannelType", c.PaymentChannelType)
	populate(objectMap, "publisherId", c.PublisherID)
	populate(objectMap, "publisherTestEnvironment", c.PublisherTestEnvironment)
	populate(objectMap, "quantity", c.Quantity)
	populate(objectMap, "skuId", c.SKUID)
	populate(objectMap, "saasResourceName", c.SaasResourceName)
	populate(objectMap, "saasSessionId", c.SaasSessionID)
	populate(objectMap, "saasSubscriptionId", c.SaasSubscriptionID)
	populate(objectMap, "termId", c.TermID)
	return json.Marshal(objectMap)
}

// DeleteOptions - delete Options
type DeleteOptions struct {
	// the feedback
	Feedback *string `json:"feedback,omitempty"`

	// The reasonCode
	ReasonCode *float32 `json:"reasonCode,omitempty"`

	// whether it is unsubscribeOnly
	UnsubscribeOnly *bool `json:"unsubscribeOnly,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetail.
func (e ErrorDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// MoveResource - Resource Move Options
type MoveResource struct {
	// The resource uris to move
	Resources []*string `json:"resources,omitempty"`

	// The target resource group uri for the move
	TargetResourceGroup *string `json:"targetResourceGroup,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type MoveResource.
func (m MoveResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "resources", m.Resources)
	populate(objectMap, "targetResourceGroup", m.TargetResourceGroup)
	return json.Marshal(objectMap)
}

// OperationClientBeginGetOptions contains the optional parameters for the OperationClient.BeginGet method.
type OperationClientBeginGetOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Properties - saas resource properties
type Properties struct {
	// Whether the current term is a Free Trial term
	IsFreeTrial *bool `json:"isFreeTrial,omitempty"`

	// The last modifier date if this resource.
	LastModified *string `json:"lastModified,omitempty"`

	// The SaaS Subscription Status.
	Status *SaasResourceStatus `json:"status,omitempty"`

	// The current Term object.
	Term *PropertiesTerm `json:"term,omitempty"`

	// READ-ONLY; The created date of this resource.
	Created *string `json:"created,omitempty" azure:"ro"`
}

// PropertiesTerm - The current Term object.
type PropertiesTerm struct {
	// The end date of the current term
	EndDate *string `json:"endDate,omitempty"`

	// The start date of the current term
	StartDate *string `json:"startDate,omitempty"`

	// The unit indicating Monthly / Yearly
	TermUnit *string `json:"termUnit,omitempty"`
}

// Resource - SaaS REST API resource definition.
type Resource struct {
	// saas properties
	Properties *ResourceProperties `json:"properties,omitempty"`

	// the resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource uri
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// ResourceCreation - SaaS REST API resource definition for creation.
type ResourceCreation struct {
	// Resource location. Only value allowed for SaaS is 'global'
	Location *string `json:"location,omitempty"`

	// The resource name
	Name *string `json:"name,omitempty"`

	// Properties of the SaaS resource that are relevant for creation.
	Properties *CreationProperties `json:"properties,omitempty"`

	// the resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; The resource uri
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceCreation.
func (r ResourceCreation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// ResourceProperties - saas properties
type ResourceProperties struct {
	// Whether the SaaS subscription will auto renew upon term end.
	AutoRenew *bool `json:"autoRenew,omitempty"`

	// Whether the current term is a Free Trial term
	IsFreeTrial *bool `json:"isFreeTrial,omitempty"`

	// The last modifier date if this resource.
	LastModified *string `json:"lastModified,omitempty"`

	// The offer id.
	OfferID *string `json:"offerId,omitempty"`

	// The metadata about the SaaS subscription such as the AzureSubscriptionId and ResourceUri.
	PaymentChannelMetadata map[string]*string `json:"paymentChannelMetadata,omitempty"`

	// The Payment channel for the SaasSubscription.
	PaymentChannelType *PaymentChannelType `json:"paymentChannelType,omitempty"`

	// The publisher id.
	PublisherID *string `json:"publisherId,omitempty"`

	// The environment in the publisher side for this resource.
	PublisherTestEnvironment *string `json:"publisherTestEnvironment,omitempty"`

	// The seat count.
	Quantity *float32 `json:"quantity,omitempty"`

	// The plan id.
	SKUID *string `json:"skuId,omitempty"`

	// The SaaS resource name.
	SaasResourceName *string `json:"saasResourceName,omitempty"`

	// The saas session id used for dev service migration request.
	SaasSessionID *string `json:"saasSessionId,omitempty"`

	// The saas subscription id used for tenant to subscription level migration request.
	SaasSubscriptionID *string `json:"saasSubscriptionId,omitempty"`

	// The SaaS Subscription Status.
	Status *SaasResourceStatus `json:"status,omitempty"`

	// The current Term object.
	Term *PropertiesTerm `json:"term,omitempty"`

	// The current Term id.
	TermID *string `json:"termId,omitempty"`

	// READ-ONLY; The created date of this resource.
	Created *string `json:"created,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProperties.
func (r ResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoRenew", r.AutoRenew)
	populate(objectMap, "created", r.Created)
	populate(objectMap, "isFreeTrial", r.IsFreeTrial)
	populate(objectMap, "lastModified", r.LastModified)
	populate(objectMap, "offerId", r.OfferID)
	populate(objectMap, "paymentChannelMetadata", r.PaymentChannelMetadata)
	populate(objectMap, "paymentChannelType", r.PaymentChannelType)
	populate(objectMap, "publisherId", r.PublisherID)
	populate(objectMap, "publisherTestEnvironment", r.PublisherTestEnvironment)
	populate(objectMap, "quantity", r.Quantity)
	populate(objectMap, "skuId", r.SKUID)
	populate(objectMap, "saasResourceName", r.SaasResourceName)
	populate(objectMap, "saasSessionId", r.SaasSessionID)
	populate(objectMap, "saasSubscriptionId", r.SaasSubscriptionID)
	populate(objectMap, "status", r.Status)
	populate(objectMap, "term", r.Term)
	populate(objectMap, "termId", r.TermID)
	return json.Marshal(objectMap)
}

// ResourceResponseWithContinuation - saas resources response with continuation.
type ResourceResponseWithContinuation struct {
	// the next link to query to get the remaining results.
	NextLink *string `json:"nextLink,omitempty"`

	// the value of response.
	Value []*Resource `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourceResponseWithContinuation.
func (r ResourceResponseWithContinuation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// ResourcesClientListAccessTokenOptions contains the optional parameters for the ResourcesClient.ListAccessToken method.
type ResourcesClientListAccessTokenOptions struct {
	// placeholder for future optional parameters
}

// ResourcesClientListOptions contains the optional parameters for the ResourcesClient.List method.
type ResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// Result - Sample result definition
type Result struct {
	// Sample property of type string
	SampleProperty *string `json:"sampleProperty,omitempty"`
}

// SubscriptionLevelClientBeginCreateOrUpdateOptions contains the optional parameters for the SubscriptionLevelClient.BeginCreateOrUpdate
// method.
type SubscriptionLevelClientBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientBeginDeleteOptions contains the optional parameters for the SubscriptionLevelClient.BeginDelete
// method.
type SubscriptionLevelClientBeginDeleteOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientBeginMoveResourcesOptions contains the optional parameters for the SubscriptionLevelClient.BeginMoveResources
// method.
type SubscriptionLevelClientBeginMoveResourcesOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientBeginUpdateOptions contains the optional parameters for the SubscriptionLevelClient.BeginUpdate
// method.
type SubscriptionLevelClientBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientBeginUpdateToUnsubscribedOptions contains the optional parameters for the SubscriptionLevelClient.BeginUpdateToUnsubscribed
// method.
type SubscriptionLevelClientBeginUpdateToUnsubscribedOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientGetOptions contains the optional parameters for the SubscriptionLevelClient.Get method.
type SubscriptionLevelClientGetOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientListAccessTokenOptions contains the optional parameters for the SubscriptionLevelClient.ListAccessToken
// method.
type SubscriptionLevelClientListAccessTokenOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientListByAzureSubscriptionOptions contains the optional parameters for the SubscriptionLevelClient.ListByAzureSubscription
// method.
type SubscriptionLevelClientListByAzureSubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientListByResourceGroupOptions contains the optional parameters for the SubscriptionLevelClient.ListByResourceGroup
// method.
type SubscriptionLevelClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionLevelClientValidateMoveResourcesOptions contains the optional parameters for the SubscriptionLevelClient.ValidateMoveResources
// method.
type SubscriptionLevelClientValidateMoveResourcesOptions struct {
	// placeholder for future optional parameters
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}