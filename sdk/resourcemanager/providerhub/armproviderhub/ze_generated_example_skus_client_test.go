//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_Get.json
func ExampleSKUsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientGetResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdate.json
func ExampleSKUsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<sku>",
		armproviderhub.SKUResource{
			Properties: &armproviderhub.SKUResourceProperties{
				SKUSettings: []*armproviderhub.SKUSetting{
					{
						Name: to.StringPtr("<name>"),
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					},
					{
						Name: to.StringPtr("<name>"),
						Costs: []*armproviderhub.SKUCost{
							{
								MeterID: to.StringPtr("<meter-id>"),
							}},
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_Delete.json
func ExampleSKUsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_GetNestedResourceTypeFirst.json
func ExampleSKUsClient_GetNestedResourceTypeFirst() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.GetNestedResourceTypeFirst(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientGetNestedResourceTypeFirstResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdateNestedResourceTypeFirst.json
func ExampleSKUsClient_CreateOrUpdateNestedResourceTypeFirst() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateNestedResourceTypeFirst(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<sku>",
		armproviderhub.SKUResource{
			Properties: &armproviderhub.SKUResourceProperties{
				SKUSettings: []*armproviderhub.SKUSetting{
					{
						Name: to.StringPtr("<name>"),
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					},
					{
						Name: to.StringPtr("<name>"),
						Costs: []*armproviderhub.SKUCost{
							{
								MeterID: to.StringPtr("<meter-id>"),
							}},
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientCreateOrUpdateNestedResourceTypeFirstResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_DeleteNestedResourceTypeFirst.json
func ExampleSKUsClient_DeleteNestedResourceTypeFirst() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	_, err = client.DeleteNestedResourceTypeFirst(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_GetNestedResourceTypeSecond.json
func ExampleSKUsClient_GetNestedResourceTypeSecond() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.GetNestedResourceTypeSecond(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientGetNestedResourceTypeSecondResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdateNestedResourceTypeSecond.json
func ExampleSKUsClient_CreateOrUpdateNestedResourceTypeSecond() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateNestedResourceTypeSecond(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<sku>",
		armproviderhub.SKUResource{
			Properties: &armproviderhub.SKUResourceProperties{
				SKUSettings: []*armproviderhub.SKUSetting{
					{
						Name: to.StringPtr("<name>"),
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					},
					{
						Name: to.StringPtr("<name>"),
						Costs: []*armproviderhub.SKUCost{
							{
								MeterID: to.StringPtr("<meter-id>"),
							}},
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientCreateOrUpdateNestedResourceTypeSecondResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_DeleteNestedResourceTypeSecond.json
func ExampleSKUsClient_DeleteNestedResourceTypeSecond() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	_, err = client.DeleteNestedResourceTypeSecond(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_GetNestedResourceTypeThird.json
func ExampleSKUsClient_GetNestedResourceTypeThird() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.GetNestedResourceTypeThird(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<nested-resource-type-third>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientGetNestedResourceTypeThirdResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_CreateOrUpdateNestedResourceTypeThird.json
func ExampleSKUsClient_CreateOrUpdateNestedResourceTypeThird() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdateNestedResourceTypeThird(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<nested-resource-type-third>",
		"<sku>",
		armproviderhub.SKUResource{
			Properties: &armproviderhub.SKUResourceProperties{
				SKUSettings: []*armproviderhub.SKUSetting{
					{
						Name: to.StringPtr("<name>"),
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					},
					{
						Name: to.StringPtr("<name>"),
						Costs: []*armproviderhub.SKUCost{
							{
								MeterID: to.StringPtr("<meter-id>"),
							}},
						Kind: to.StringPtr("<kind>"),
						Tier: to.StringPtr("<tier>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SKUsClientCreateOrUpdateNestedResourceTypeThirdResult)
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_DeleteNestedResourceTypeThird.json
func ExampleSKUsClient_DeleteNestedResourceTypeThird() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	_, err = client.DeleteNestedResourceTypeThird(ctx,
		"<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<nested-resource-type-third>",
		"<sku>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrations.json
func ExampleSKUsClient_ListByResourceTypeRegistrations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceTypeRegistrations("<provider-namespace>",
		"<resource-type>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrationsNestedResourceTypeFirst.json
func ExampleSKUsClient_ListByResourceTypeRegistrationsNestedResourceTypeFirst() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceTypeRegistrationsNestedResourceTypeFirst("<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrationsNestedResourceTypeSecond.json
func ExampleSKUsClient_ListByResourceTypeRegistrationsNestedResourceTypeSecond() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceTypeRegistrationsNestedResourceTypeSecond("<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Skus_ListByResourceTypeRegistrationsNestedResourceTypeThird.json
func ExampleSKUsClient_ListByResourceTypeRegistrationsNestedResourceTypeThird() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armproviderhub.NewSKUsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceTypeRegistrationsNestedResourceTypeThird("<provider-namespace>",
		"<resource-type>",
		"<nested-resource-type-first>",
		"<nested-resource-type-second>",
		"<nested-resource-type-third>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}