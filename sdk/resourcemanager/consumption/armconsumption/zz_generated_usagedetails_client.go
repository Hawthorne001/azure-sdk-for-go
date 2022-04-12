//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// UsageDetailsClient contains the methods for the UsageDetails group.
// Don't use this type directly, use NewUsageDetailsClient() instead.
type UsageDetailsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewUsageDetailsClient creates a new instance of UsageDetailsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewUsageDetailsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*UsageDetailsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &UsageDetailsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// List - Lists the usage details for the defined scope. Usage details are available via this API only for May 1, 2014 or
// later.
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope associated with usage details operations. This includes '/subscriptions/{subscriptionId}/' for subscription
// scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for
// Billing Account scope, '/providers/Microsoft.Billing/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/enrollmentAccounts/{enrollmentAccountId}'
// for EnrollmentAccount
// scope and '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope. For subscription,
// billing account, department, enrollment account and management group, you
// can also add billing period to the scope using '/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'. For e.g.
// to specify billing period at department scope use
// '/providers/Microsoft.Billing/departments/{departmentId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}'.
// Also, Modern Commerce Account scopes are
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for billingAccount scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}'
// for
// billingProfile scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/invoiceSections/{invoiceSectionId}'
// for invoiceSection scope, and
// 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/customers/{customerId}' specific for partners.
// options - UsageDetailsClientListOptions contains the optional parameters for the UsageDetailsClient.List method.
func (client *UsageDetailsClient) List(scope string, options *UsageDetailsClientListOptions) *runtime.Pager[UsageDetailsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[UsageDetailsClientListResponse]{
		More: func(page UsageDetailsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UsageDetailsClientListResponse) (UsageDetailsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return UsageDetailsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return UsageDetailsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UsageDetailsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *UsageDetailsClient) listCreateRequest(ctx context.Context, scope string, options *UsageDetailsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/usageDetails"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2021-10-01")
	if options != nil && options.Metric != nil {
		reqQP.Set("metric", string(*options.Metric))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsageDetailsClient) listHandleResponse(resp *http.Response) (UsageDetailsClientListResponse, error) {
	result := UsageDetailsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageDetailsListResult); err != nil {
		return UsageDetailsClientListResponse{}, err
	}
	return result, nil
}
