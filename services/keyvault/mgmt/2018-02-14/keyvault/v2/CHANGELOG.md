## Breaking Changes

## Signature Changes

### Funcs

1. New
	- Params
		- From: string
		- To: string,string,string,string
1. NewOperationsClient
	- Params
		- From: string
		- To: string,string,string,string
1. NewOperationsClientWithBaseURI
	- Params
		- From: string,string
		- To: string,string,string,string,string
1. NewVaultsClient
	- Params
		- From: string
		- To: string,string,string,string
1. NewVaultsClientWithBaseURI
	- Params
		- From: string,string
		- To: string,string,string,string,string
1. NewWithBaseURI
	- Params
		- From: string,string
		- To: string,string,string,string,string

## New Content

### New Constants

1. PrivateEndpointConnectionProvisioningState.Creating
1. PrivateEndpointConnectionProvisioningState.Deleting
1. PrivateEndpointConnectionProvisioningState.Failed
1. PrivateEndpointConnectionProvisioningState.Succeeded
1. PrivateEndpointServiceConnectionStatus.Approved
1. PrivateEndpointServiceConnectionStatus.Disconnected
1. PrivateEndpointServiceConnectionStatus.Pending
1. PrivateEndpointServiceConnectionStatus.Rejected

### New Funcs

1. *PrivateEndpointConnection.UnmarshalJSON([]byte) error
1. *PrivateLinkResource.UnmarshalJSON([]byte) error
1. NewPrivateEndpointConnectionsClient(string,string,string,string) PrivateEndpointConnectionsClient
1. NewPrivateEndpointConnectionsClientWithBaseURI(string,string,string,string,string) PrivateEndpointConnectionsClient
1. NewPrivateLinkResourcesClient(string,string,string,string) PrivateLinkResourcesClient
1. NewPrivateLinkResourcesClientWithBaseURI(string,string,string,string,string) PrivateLinkResourcesClient
1. PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState
1. PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus
1. PrivateEndpointConnection.MarshalJSON() ([]byte,error)
1. PrivateEndpointConnectionsClient.Delete(context.Context) (autorest.Response,error)
1. PrivateEndpointConnectionsClient.DeletePreparer(context.Context) (*http.Request,error)
1. PrivateEndpointConnectionsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. PrivateEndpointConnectionsClient.DeleteSender(*http.Request) (*http.Response,error)
1. PrivateEndpointConnectionsClient.Get(context.Context) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.GetPreparer(context.Context) (*http.Request,error)
1. PrivateEndpointConnectionsClient.GetResponder(*http.Response) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.GetSender(*http.Request) (*http.Response,error)
1. PrivateEndpointConnectionsClient.Put(context.Context,PrivateEndpointConnection) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.PutPreparer(context.Context,PrivateEndpointConnection) (*http.Request,error)
1. PrivateEndpointConnectionsClient.PutResponder(*http.Response) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.PutSender(*http.Request) (*http.Response,error)
1. PrivateLinkResource.MarshalJSON() ([]byte,error)
1. PrivateLinkResourcesClient.ListByVault(context.Context) (PrivateLinkResourceListResult,error)
1. PrivateLinkResourcesClient.ListByVaultPreparer(context.Context) (*http.Request,error)
1. PrivateLinkResourcesClient.ListByVaultResponder(*http.Response) (PrivateLinkResourceListResult,error)
1. PrivateLinkResourcesClient.ListByVaultSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. PrivateEndpoint
1. PrivateEndpointConnection
1. PrivateEndpointConnectionProperties
1. PrivateEndpointConnectionsClient
1. PrivateLinkResource
1. PrivateLinkResourceListResult
1. PrivateLinkResourceProperties
1. PrivateLinkResourcesClient
1. PrivateLinkServiceConnectionState

### New Struct Fields

1. BaseClient.PrivateEndpointConnectionName
1. BaseClient.ResourceGroupName
1. BaseClient.VaultName
