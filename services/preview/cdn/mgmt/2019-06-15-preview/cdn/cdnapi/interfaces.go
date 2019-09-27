package cdnapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/cdn/mgmt/2019-06-15-preview/cdn"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput cdn.CheckNameAvailabilityInput) (result cdn.CheckNameAvailabilityOutput, err error)
	CheckNameAvailabilityWithSubscription(ctx context.Context, checkNameAvailabilityInput cdn.CheckNameAvailabilityInput) (result cdn.CheckNameAvailabilityOutput, err error)
	ValidateProbe(ctx context.Context, validateProbeInput cdn.ValidateProbeInput) (result cdn.ValidateProbeOutput, err error)
}

var _ BaseClientAPI = (*cdn.BaseClient)(nil)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, profile cdn.Profile) (result cdn.ProfilesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ProfilesDeleteFuture, err error)
	GenerateSsoURI(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SsoURI, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string) (result cdn.Profile, err error)
	List(ctx context.Context) (result cdn.ProfileListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result cdn.ProfileListResultPage, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string) (result cdn.ResourceUsageListResultPage, err error)
	ListSupportedOptimizationTypes(ctx context.Context, resourceGroupName string, profileName string) (result cdn.SupportedOptimizationTypesListResult, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, profileUpdateParameters cdn.ProfileUpdateParameters) (result cdn.ProfilesUpdateFuture, err error)
}

var _ ProfilesClientAPI = (*cdn.ProfilesClient)(nil)

// EndpointsClientAPI contains the set of methods on the EndpointsClient type.
type EndpointsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint cdn.Endpoint) (result cdn.EndpointsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointNamee string) (result cdn.Endpoint, err error)
	ListByProfile(ctx context.Context, resourceGroupName string, profileName string) (result cdn.EndpointListResultPage, err error)
	ListResourceUsage(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.ResourceUsageListResultPage, err error)
	LoadContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths cdn.LoadParameters) (result cdn.EndpointsLoadContentFuture, err error)
	PurgeContent(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths cdn.PurgeParameters) (result cdn.EndpointsPurgeContentFuture, err error)
	Start(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.EndpointsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties cdn.EndpointUpdateParameters) (result cdn.EndpointsUpdateFuture, err error)
	ValidateCustomDomain(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties cdn.ValidateCustomDomainInput) (result cdn.ValidateCustomDomainOutput, err error)
}

var _ EndpointsClientAPI = (*cdn.EndpointsClient)(nil)

// OriginsClientAPI contains the set of methods on the OriginsClient type.
type OriginsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string) (result cdn.Origin, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.OriginListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties cdn.OriginUpdateParameters) (result cdn.OriginsUpdateFuture, err error)
}

var _ OriginsClientAPI = (*cdn.OriginsClient)(nil)

// CustomDomainsClientAPI contains the set of methods on the CustomDomainsClient type.
type CustomDomainsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties cdn.CustomDomainParameters) (result cdn.CustomDomainsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomainsDeleteFuture, err error)
	DisableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
	EnableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainHTTPSParameters *cdn.BasicCustomDomainHTTPSParameters) (result cdn.CustomDomain, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result cdn.CustomDomain, err error)
	ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result cdn.CustomDomainListResultPage, err error)
}

var _ CustomDomainsClientAPI = (*cdn.CustomDomainsClient)(nil)

// ResourceUsageClientAPI contains the set of methods on the ResourceUsageClient type.
type ResourceUsageClientAPI interface {
	List(ctx context.Context) (result cdn.ResourceUsageListResultPage, err error)
}

var _ ResourceUsageClientAPI = (*cdn.ResourceUsageClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result cdn.OperationsListResultPage, err error)
}

var _ OperationsClientAPI = (*cdn.OperationsClient)(nil)

// EdgeNodesClientAPI contains the set of methods on the EdgeNodesClient type.
type EdgeNodesClientAPI interface {
	List(ctx context.Context) (result cdn.EdgenodeResultPage, err error)
}

var _ EdgeNodesClientAPI = (*cdn.EdgeNodesClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, cdnWebApplicationFirewallPolicy cdn.WebApplicationFirewallPolicy) (result cdn.PoliciesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, policyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, policyName string) (result cdn.WebApplicationFirewallPolicy, err error)
	List(ctx context.Context, resourceGroupName string) (result cdn.WebApplicationFirewallPolicyListPage, err error)
	Update(ctx context.Context, resourceGroupName string, policyName string, cdnWebApplicationFirewallPolicyPatchParameters cdn.WebApplicationFirewallPolicyPatchParameters) (result cdn.PoliciesUpdateFuture, err error)
}

var _ PoliciesClientAPI = (*cdn.PoliciesClient)(nil)

// ManagedRuleSetsClientAPI contains the set of methods on the ManagedRuleSetsClient type.
type ManagedRuleSetsClientAPI interface {
	List(ctx context.Context) (result cdn.ManagedRuleSetDefinitionListPage, err error)
}

var _ ManagedRuleSetsClientAPI = (*cdn.ManagedRuleSetsClient)(nil)
