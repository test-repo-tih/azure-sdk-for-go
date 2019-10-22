package blockchainapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/blockchain/mgmt/2018-06-01-preview/blockchain/v2"
)

// MembersClientAPI contains the set of methods on the MembersClient type.
type MembersClientAPI interface {
	Create(ctx context.Context, blockchainMemberName string, resourceGroupName string, blockchainMember *blockchain.Member) (result blockchain.MembersCreateFuture, err error)
	Delete(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result blockchain.MembersDeleteFuture, err error)
	Get(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result blockchain.Member, err error)
	List(ctx context.Context, resourceGroupName string) (result blockchain.MemberCollectionPage, err error)
	ListAll(ctx context.Context) (result blockchain.MemberCollectionPage, err error)
	ListAPIKeys(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result blockchain.APIKeyCollection, err error)
	ListConsortiumMembers(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result blockchain.ConsortiumMemberCollectionPage, err error)
	RegenerateAPIKeys(ctx context.Context, blockchainMemberName string, resourceGroupName string, APIKey *blockchain.APIKey) (result blockchain.APIKeyCollection, err error)
	Update(ctx context.Context, blockchainMemberName string, resourceGroupName string, blockchainMember *blockchain.MemberUpdate) (result blockchain.Member, err error)
}

var _ MembersClientAPI = (*blockchain.MembersClient)(nil)

// MemberOperationResultsClientAPI contains the set of methods on the MemberOperationResultsClient type.
type MemberOperationResultsClientAPI interface {
	Get(ctx context.Context, locationName string, operationID string) (result blockchain.OperationResult, err error)
}

var _ MemberOperationResultsClientAPI = (*blockchain.MemberOperationResultsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	CheckNameAvailability(ctx context.Context, locationName string, nameAvailabilityRequest *blockchain.NameAvailabilityRequest) (result blockchain.NameAvailability, err error)
	ListConsortiums(ctx context.Context, locationName string) (result blockchain.ConsortiumCollection, err error)
}

var _ LocationsClientAPI = (*blockchain.LocationsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result blockchain.ResourceProviderOperationCollectionPage, err error)
}

var _ OperationsClientAPI = (*blockchain.OperationsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context) (result blockchain.ResourceTypeSkuCollection, err error)
}

var _ SkusClientAPI = (*blockchain.SkusClient)(nil)

// TransactionNodesClientAPI contains the set of methods on the TransactionNodesClient type.
type TransactionNodesClientAPI interface {
	Create(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, transactionNode *blockchain.TransactionNode) (result blockchain.TransactionNodesCreateFuture, err error)
	Delete(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (result blockchain.TransactionNodesDeleteFuture, err error)
	Get(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (result blockchain.TransactionNode, err error)
	List(ctx context.Context, blockchainMemberName string, resourceGroupName string) (result blockchain.TransactionNodeCollectionPage, err error)
	ListAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string) (result blockchain.APIKeyCollection, err error)
	RegenerateAPIKeys(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, APIKey *blockchain.APIKey) (result blockchain.APIKeyCollection, err error)
	Update(ctx context.Context, blockchainMemberName string, transactionNodeName string, resourceGroupName string, transactionNode *blockchain.TransactionNodeUpdate) (result blockchain.TransactionNode, err error)
}

var _ TransactionNodesClientAPI = (*blockchain.TransactionNodesClient)(nil)
