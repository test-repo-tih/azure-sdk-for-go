// Package cognitiveservices implements the Azure ARM Cognitiveservices service API version 2017-04-18.
//
// Cognitive Services Management Client
package cognitiveservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Cognitiveservices
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Cognitiveservices.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckDomainAvailability check whether a domain is available.
// Parameters:
// parameters - check Domain Availability parameter.
func (client BaseClient) CheckDomainAvailability(ctx context.Context, parameters CheckDomainAvailabilityParameter) (result CheckDomainAvailabilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckDomainAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SubdomainName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.BaseClient", "CheckDomainAvailability", err.Error())
	}

	req, err := client.CheckDomainAvailabilityPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckDomainAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckDomainAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckDomainAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckDomainAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckDomainAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckDomainAvailabilityPreparer prepares the CheckDomainAvailability request.
func (client BaseClient) CheckDomainAvailabilityPreparer(ctx context.Context, parameters CheckDomainAvailabilityParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/checkDomainAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckDomainAvailabilitySender sends the CheckDomainAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckDomainAvailabilitySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckDomainAvailabilityResponder handles the response to the CheckDomainAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckDomainAvailabilityResponder(resp *http.Response) (result CheckDomainAvailabilityResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CheckSkuAvailability check available SKUs.
// Parameters:
// location - resource location.
// parameters - check SKU Availability POST body.
func (client BaseClient) CheckSkuAvailability(ctx context.Context, location string, parameters CheckSkuAvailabilityParameter) (result CheckSkuAvailabilityResultList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckSkuAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Skus", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Kind", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cognitiveservices.BaseClient", "CheckSkuAvailability", err.Error())
	}

	req, err := client.CheckSkuAvailabilityPreparer(ctx, location, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckSkuAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckSkuAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckSkuAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckSkuAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservices.BaseClient", "CheckSkuAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckSkuAvailabilityPreparer prepares the CheckSkuAvailability request.
func (client BaseClient) CheckSkuAvailabilityPreparer(ctx context.Context, location string, parameters CheckSkuAvailabilityParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-18"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/checkSkuAvailability", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckSkuAvailabilitySender sends the CheckSkuAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckSkuAvailabilitySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckSkuAvailabilityResponder handles the response to the CheckSkuAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckSkuAvailabilityResponder(resp *http.Response) (result CheckSkuAvailabilityResultList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
