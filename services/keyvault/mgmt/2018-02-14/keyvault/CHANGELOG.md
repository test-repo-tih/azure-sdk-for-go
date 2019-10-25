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
1. NewPrivateEndpointConnectionsClient(string) PrivateEndpointConnectionsClient
1. NewPrivateEndpointConnectionsClientWithBaseURI(string,string) PrivateEndpointConnectionsClient
1. NewPrivateLinkResourcesClient(string) PrivateLinkResourcesClient
1. NewPrivateLinkResourcesClientWithBaseURI(string,string) PrivateLinkResourcesClient
1. PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState
1. PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus
1. PrivateEndpointConnection.MarshalJSON() ([]byte,error)
1. PrivateEndpointConnectionsClient.Delete(context.Context,string,string,string) (autorest.Response,error)
1. PrivateEndpointConnectionsClient.DeletePreparer(context.Context,string,string,string) (*http.Request,error)
1. PrivateEndpointConnectionsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. PrivateEndpointConnectionsClient.DeleteSender(*http.Request) (*http.Response,error)
1. PrivateEndpointConnectionsClient.Get(context.Context,string,string,string) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. PrivateEndpointConnectionsClient.GetResponder(*http.Response) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.GetSender(*http.Request) (*http.Response,error)
1. PrivateEndpointConnectionsClient.Put(context.Context,string,string,string,PrivateEndpointConnection) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.PutPreparer(context.Context,string,string,string,PrivateEndpointConnection) (*http.Request,error)
1. PrivateEndpointConnectionsClient.PutResponder(*http.Response) (PrivateEndpointConnection,error)
1. PrivateEndpointConnectionsClient.PutSender(*http.Request) (*http.Response,error)
1. PrivateLinkResource.MarshalJSON() ([]byte,error)
1. PrivateLinkResourcesClient.ListByVault(context.Context,string,string) (PrivateLinkResourceListResult,error)
1. PrivateLinkResourcesClient.ListByVaultPreparer(context.Context,string,string) (*http.Request,error)
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

1. VaultProperties.PrivateEndpointConnections
