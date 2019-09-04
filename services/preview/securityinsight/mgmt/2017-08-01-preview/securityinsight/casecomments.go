package securityinsight

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

// CaseCommentsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type CaseCommentsClient struct {
	BaseClient
}

// NewCaseCommentsClient creates an instance of the CaseCommentsClient client.
func NewCaseCommentsClient(subscriptionID string) CaseCommentsClient {
	return NewCaseCommentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCaseCommentsClientWithBaseURI creates an instance of the CaseCommentsClient client.
func NewCaseCommentsClientWithBaseURI(baseURI string, subscriptionID string) CaseCommentsClient {
	return CaseCommentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateComment creates the case comment.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// caseID - case ID
// caseCommentID - case comment ID
// caseComment - the case comment
func (client CaseCommentsClient) CreateComment(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseCommentID string, caseComment CaseComment) (result CaseComment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CaseCommentsClient.CreateComment")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: caseComment,
			Constraints: []validation.Constraint{{Target: "caseComment.CaseCommentProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "caseComment.CaseCommentProperties.Message", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "caseComment.CaseCommentProperties.UserInfo", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "caseComment.CaseCommentProperties.UserInfo.ObjectID", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("securityinsight.CaseCommentsClient", "CreateComment", err.Error())
	}

	req, err := client.CreateCommentPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, caseCommentID, caseComment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseCommentsClient", "CreateComment", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateCommentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.CaseCommentsClient", "CreateComment", resp, "Failure sending request")
		return
	}

	result, err = client.CreateCommentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.CaseCommentsClient", "CreateComment", resp, "Failure responding to request")
	}

	return
}

// CreateCommentPreparer prepares the CreateComment request.
func (client CaseCommentsClient) CreateCommentPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseCommentID string, caseComment CaseComment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"caseCommentId": autorest.Encode("path", caseCommentID),
		"caseId":        autorest.Encode("path", caseID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}/comments/{caseCommentId}", pathParameters),
		autorest.WithJSON(caseComment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateCommentSender sends the CreateComment request. The method will close the
// http.Response Body if it receives an error.
func (client CaseCommentsClient) CreateCommentSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateCommentResponder handles the response to the CreateComment request. The method always
// closes the http.Response Body.
func (client CaseCommentsClient) CreateCommentResponder(resp *http.Response) (result CaseComment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
