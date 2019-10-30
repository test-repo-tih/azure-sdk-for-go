## Breaking Changes

## Struct Changes

### Removed Struct Fields

1. BaseClient.SubscriptionID

## Signature Changes

### Funcs

1. DataSourcesClient.CreateOrUpdate
	- Params
		- From: context.Context,string,string,string,DataSource
		- To: context.Context,string,string,string,string,DataSource
1. DataSourcesClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context,string,string,string,DataSource
		- To: context.Context,string,string,string,string,DataSource
1. DataSourcesClient.Delete
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. DataSourcesClient.DeletePreparer
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. DataSourcesClient.Get
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. DataSourcesClient.GetPreparer
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. DataSourcesClient.ListByWorkspace
	- Params
		- From: context.Context,string,string,string,string
		- To: context.Context,string,string,string,string,string
1. DataSourcesClient.ListByWorkspaceComplete
	- Params
		- From: context.Context,string,string,string,string
		- To: context.Context,string,string,string,string,string
1. DataSourcesClient.ListByWorkspacePreparer
	- Params
		- From: context.Context,string,string,string,string
		- To: context.Context,string,string,string,string,string
1. LinkedServicesClient.CreateOrUpdate
	- Params
		- From: context.Context,string,string,string,LinkedService
		- To: context.Context,string,string,string,string,LinkedService
1. LinkedServicesClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context,string,string,string,LinkedService
		- To: context.Context,string,string,string,string,LinkedService
1. LinkedServicesClient.Delete
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. LinkedServicesClient.DeletePreparer
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. LinkedServicesClient.Get
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. LinkedServicesClient.GetPreparer
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. LinkedServicesClient.ListByWorkspace
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. LinkedServicesClient.ListByWorkspacePreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. New
	- Params
		- From: string
		- To: <none>
1. NewDataSourcesClient
	- Params
		- From: string
		- To: <none>
1. NewDataSourcesClientWithBaseURI
	- Params
		- From: string,string
		- To: string
1. NewLinkedServicesClient
	- Params
		- From: string
		- To: <none>
1. NewLinkedServicesClientWithBaseURI
	- Params
		- From: string,string
		- To: string
1. NewOperationsClient
	- Params
		- From: string
		- To: <none>
1. NewOperationsClientWithBaseURI
	- Params
		- From: string,string
		- To: string
1. NewWithBaseURI
	- Params
		- From: string,string
		- To: string
1. NewWorkspacesClient
	- Params
		- From: string
		- To: <none>
1. NewWorkspacesClientWithBaseURI
	- Params
		- From: string,string
		- To: string
1. WorkspacesClient.CreateOrUpdate
	- Params
		- From: context.Context,string,string,Workspace
		- To: context.Context,string,string,string,Workspace
1. WorkspacesClient.CreateOrUpdatePreparer
	- Params
		- From: context.Context,string,string,Workspace
		- To: context.Context,string,string,string,Workspace
1. WorkspacesClient.Delete
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.DeletePreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.DisableIntelligencePack
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. WorkspacesClient.DisableIntelligencePackPreparer
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. WorkspacesClient.EnableIntelligencePack
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. WorkspacesClient.EnableIntelligencePackPreparer
	- Params
		- From: context.Context,string,string,string
		- To: context.Context,string,string,string,string
1. WorkspacesClient.Get
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.GetPreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.GetSharedKeys
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.GetSharedKeysPreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.List
	- Params
		- From: context.Context
		- To: context.Context,string
1. WorkspacesClient.ListByResourceGroup
	- Params
		- From: context.Context,string
		- To: context.Context,string,string
1. WorkspacesClient.ListByResourceGroupPreparer
	- Params
		- From: context.Context,string
		- To: context.Context,string,string
1. WorkspacesClient.ListIntelligencePacks
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.ListIntelligencePacksPreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.ListManagementGroups
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.ListManagementGroupsPreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.ListPreparer
	- Params
		- From: context.Context
		- To: context.Context,string
1. WorkspacesClient.ListUsages
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.ListUsagesPreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,string
1. WorkspacesClient.Update
	- Params
		- From: context.Context,string,string,Workspace
		- To: context.Context,string,string,string,Workspace
1. WorkspacesClient.UpdatePreparer
	- Params
		- From: context.Context,string,string,Workspace
		- To: context.Context,string,string,string,Workspace

## New Content

### New Funcs

1. *Table.UnmarshalJSON([]byte) error
1. BaseClient.ListTables(context.Context,string,string,string) (TablesListResult,error)
1. BaseClient.ListTablesPreparer(context.Context,string,string,string) (*http.Request,error)
1. BaseClient.ListTablesResponder(*http.Response) (TablesListResult,error)
1. BaseClient.ListTablesSender(*http.Request) (*http.Response,error)
1. NewTablesClient() TablesClient
1. NewTablesClientWithBaseURI(string) TablesClient
1. Table.MarshalJSON() ([]byte,error)
1. TablesClient.CreateOrUpdate(context.Context,string,string,string,string,Table) (Table,error)
1. TablesClient.CreateOrUpdatePreparer(context.Context,string,string,string,string,Table) (*http.Request,error)
1. TablesClient.CreateOrUpdateResponder(*http.Response) (Table,error)
1. TablesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. TablesClient.Get(context.Context,string,string,string,string) (Table,error)
1. TablesClient.GetPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. TablesClient.GetResponder(*http.Response) (Table,error)
1. TablesClient.GetSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. Table
1. TableProperties
1. TablesClient
1. TablesListResult
