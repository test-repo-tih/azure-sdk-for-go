## Breaking Changes

## Signature Changes

### Funcs

1. DatabasesClient.Failover
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,ReplicaType
1. DatabasesClient.FailoverPreparer
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,ReplicaType

## New Content

### New Constants

1. ReplicaType.Primary
1. ReplicaType.ReadableSecondary

### New Funcs

1. *AdministratorListResultIterator.Next() error
1. *AdministratorListResultIterator.NextWithContext(context.Context) error
1. *AdministratorListResultPage.Next() error
1. *AdministratorListResultPage.NextWithContext(context.Context) error
1. *ServerAzureADAdministrator.UnmarshalJSON([]byte) error
1. *ServerAzureADAdministratorsCreateOrUpdateFuture.Result(ServerAzureADAdministratorsClient) (ServerAzureADAdministrator,error)
1. *ServerAzureADAdministratorsDeleteFuture.Result(ServerAzureADAdministratorsClient) (autorest.Response,error)
1. AdministratorListResult.IsEmpty() bool
1. AdministratorListResultIterator.NotDone() bool
1. AdministratorListResultIterator.Response() AdministratorListResult
1. AdministratorListResultIterator.Value() ServerAzureADAdministrator
1. AdministratorListResultPage.NotDone() bool
1. AdministratorListResultPage.Response() AdministratorListResult
1. AdministratorListResultPage.Values() []ServerAzureADAdministrator
1. NewAdministratorListResultIterator(AdministratorListResultPage) AdministratorListResultIterator
1. NewAdministratorListResultPage(func(context.Context, AdministratorListResult) (AdministratorListResult, error)) AdministratorListResultPage
1. NewServerAzureADAdministratorsClient(string) ServerAzureADAdministratorsClient
1. NewServerAzureADAdministratorsClientWithBaseURI(string,string) ServerAzureADAdministratorsClient
1. PossibleReplicaTypeValues() []ReplicaType
1. ServerAzureADAdministrator.MarshalJSON() ([]byte,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdate(context.Context,string,string,string,ServerAzureADAdministrator) (ServerAzureADAdministratorsCreateOrUpdateFuture,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdatePreparer(context.Context,string,string,string,ServerAzureADAdministrator) (*http.Request,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdateResponder(*http.Response) (ServerAzureADAdministrator,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdateSender(*http.Request) (ServerAzureADAdministratorsCreateOrUpdateFuture,error)
1. ServerAzureADAdministratorsClient.Delete(context.Context,string,string,string) (ServerAzureADAdministratorsDeleteFuture,error)
1. ServerAzureADAdministratorsClient.DeletePreparer(context.Context,string,string,string) (*http.Request,error)
1. ServerAzureADAdministratorsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. ServerAzureADAdministratorsClient.DeleteSender(*http.Request) (ServerAzureADAdministratorsDeleteFuture,error)
1. ServerAzureADAdministratorsClient.Get(context.Context,string,string,string) (ServerAzureADAdministrator,error)
1. ServerAzureADAdministratorsClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. ServerAzureADAdministratorsClient.GetResponder(*http.Response) (ServerAzureADAdministrator,error)
1. ServerAzureADAdministratorsClient.GetSender(*http.Request) (*http.Response,error)
1. ServerAzureADAdministratorsClient.ListByServer(context.Context,string,string) (AdministratorListResultPage,error)
1. ServerAzureADAdministratorsClient.ListByServerComplete(context.Context,string,string) (AdministratorListResultIterator,error)
1. ServerAzureADAdministratorsClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. ServerAzureADAdministratorsClient.ListByServerResponder(*http.Response) (AdministratorListResult,error)
1. ServerAzureADAdministratorsClient.ListByServerSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. AdministratorListResult
1. AdministratorListResultIterator
1. AdministratorListResultPage
1. AdministratorProperties
1. ServerAzureADAdministrator
1. ServerAzureADAdministratorsClient
1. ServerAzureADAdministratorsCreateOrUpdateFuture
1. ServerAzureADAdministratorsDeleteFuture
