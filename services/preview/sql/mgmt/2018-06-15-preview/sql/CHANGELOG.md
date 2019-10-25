## New Content

### New Funcs

1. *WorkloadClassifier.UnmarshalJSON([]byte) error
1. *WorkloadClassifierListResultIterator.Next() error
1. *WorkloadClassifierListResultIterator.NextWithContext(context.Context) error
1. *WorkloadClassifierListResultPage.Next() error
1. *WorkloadClassifierListResultPage.NextWithContext(context.Context) error
1. *WorkloadClassifiersCreateOrUpdateFuture.Result(WorkloadClassifiersClient) (WorkloadClassifier,error)
1. *WorkloadClassifiersDeleteFuture.Result(WorkloadClassifiersClient) (autorest.Response,error)
1. *WorkloadGroup.UnmarshalJSON([]byte) error
1. *WorkloadGroupListResultIterator.Next() error
1. *WorkloadGroupListResultIterator.NextWithContext(context.Context) error
1. *WorkloadGroupListResultPage.Next() error
1. *WorkloadGroupListResultPage.NextWithContext(context.Context) error
1. *WorkloadGroupsCreateOrUpdateFuture.Result(WorkloadGroupsClient) (WorkloadGroup,error)
1. *WorkloadGroupsDeleteFuture.Result(WorkloadGroupsClient) (autorest.Response,error)
1. NewWorkloadClassifierListResultIterator(WorkloadClassifierListResultPage) WorkloadClassifierListResultIterator
1. NewWorkloadClassifierListResultPage(func(context.Context, WorkloadClassifierListResult) (WorkloadClassifierListResult, error)) WorkloadClassifierListResultPage
1. NewWorkloadClassifiersClient(string) WorkloadClassifiersClient
1. NewWorkloadClassifiersClientWithBaseURI(string,string) WorkloadClassifiersClient
1. NewWorkloadGroupListResultIterator(WorkloadGroupListResultPage) WorkloadGroupListResultIterator
1. NewWorkloadGroupListResultPage(func(context.Context, WorkloadGroupListResult) (WorkloadGroupListResult, error)) WorkloadGroupListResultPage
1. NewWorkloadGroupsClient(string) WorkloadGroupsClient
1. NewWorkloadGroupsClientWithBaseURI(string,string) WorkloadGroupsClient
1. WorkloadClassifier.MarshalJSON() ([]byte,error)
1. WorkloadClassifierListResult.IsEmpty() bool
1. WorkloadClassifierListResultIterator.NotDone() bool
1. WorkloadClassifierListResultIterator.Response() WorkloadClassifierListResult
1. WorkloadClassifierListResultIterator.Value() WorkloadClassifier
1. WorkloadClassifierListResultPage.NotDone() bool
1. WorkloadClassifierListResultPage.Response() WorkloadClassifierListResult
1. WorkloadClassifierListResultPage.Values() []WorkloadClassifier
1. WorkloadClassifiersClient.CreateOrUpdate(context.Context,string,string,string,string,string,WorkloadClassifier) (WorkloadClassifiersCreateOrUpdateFuture,error)
1. WorkloadClassifiersClient.CreateOrUpdatePreparer(context.Context,string,string,string,string,string,WorkloadClassifier) (*http.Request,error)
1. WorkloadClassifiersClient.CreateOrUpdateResponder(*http.Response) (WorkloadClassifier,error)
1. WorkloadClassifiersClient.CreateOrUpdateSender(*http.Request) (WorkloadClassifiersCreateOrUpdateFuture,error)
1. WorkloadClassifiersClient.Delete(context.Context,string,string,string,string,string) (WorkloadClassifiersDeleteFuture,error)
1. WorkloadClassifiersClient.DeletePreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. WorkloadClassifiersClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. WorkloadClassifiersClient.DeleteSender(*http.Request) (WorkloadClassifiersDeleteFuture,error)
1. WorkloadClassifiersClient.Get(context.Context,string,string,string,string,string) (WorkloadClassifier,error)
1. WorkloadClassifiersClient.GetPreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. WorkloadClassifiersClient.GetResponder(*http.Response) (WorkloadClassifier,error)
1. WorkloadClassifiersClient.GetSender(*http.Request) (*http.Response,error)
1. WorkloadClassifiersClient.ListByWorkloadGroup(context.Context,string,string,string,string) (WorkloadClassifierListResultPage,error)
1. WorkloadClassifiersClient.ListByWorkloadGroupComplete(context.Context,string,string,string,string) (WorkloadClassifierListResultIterator,error)
1. WorkloadClassifiersClient.ListByWorkloadGroupPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. WorkloadClassifiersClient.ListByWorkloadGroupResponder(*http.Response) (WorkloadClassifierListResult,error)
1. WorkloadClassifiersClient.ListByWorkloadGroupSender(*http.Request) (*http.Response,error)
1. WorkloadGroup.MarshalJSON() ([]byte,error)
1. WorkloadGroupListResult.IsEmpty() bool
1. WorkloadGroupListResultIterator.NotDone() bool
1. WorkloadGroupListResultIterator.Response() WorkloadGroupListResult
1. WorkloadGroupListResultIterator.Value() WorkloadGroup
1. WorkloadGroupListResultPage.NotDone() bool
1. WorkloadGroupListResultPage.Response() WorkloadGroupListResult
1. WorkloadGroupListResultPage.Values() []WorkloadGroup
1. WorkloadGroupsClient.CreateOrUpdate(context.Context,string,string,string,string,WorkloadGroup) (WorkloadGroupsCreateOrUpdateFuture,error)
1. WorkloadGroupsClient.CreateOrUpdatePreparer(context.Context,string,string,string,string,WorkloadGroup) (*http.Request,error)
1. WorkloadGroupsClient.CreateOrUpdateResponder(*http.Response) (WorkloadGroup,error)
1. WorkloadGroupsClient.CreateOrUpdateSender(*http.Request) (WorkloadGroupsCreateOrUpdateFuture,error)
1. WorkloadGroupsClient.Delete(context.Context,string,string,string,string) (WorkloadGroupsDeleteFuture,error)
1. WorkloadGroupsClient.DeletePreparer(context.Context,string,string,string,string) (*http.Request,error)
1. WorkloadGroupsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. WorkloadGroupsClient.DeleteSender(*http.Request) (WorkloadGroupsDeleteFuture,error)
1. WorkloadGroupsClient.Get(context.Context,string,string,string,string) (WorkloadGroup,error)
1. WorkloadGroupsClient.GetPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. WorkloadGroupsClient.GetResponder(*http.Response) (WorkloadGroup,error)
1. WorkloadGroupsClient.GetSender(*http.Request) (*http.Response,error)
1. WorkloadGroupsClient.ListByDatabase(context.Context,string,string,string) (WorkloadGroupListResultPage,error)
1. WorkloadGroupsClient.ListByDatabaseComplete(context.Context,string,string,string) (WorkloadGroupListResultIterator,error)
1. WorkloadGroupsClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. WorkloadGroupsClient.ListByDatabaseResponder(*http.Response) (WorkloadGroupListResult,error)
1. WorkloadGroupsClient.ListByDatabaseSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. WorkloadClassifier
1. WorkloadClassifierListResult
1. WorkloadClassifierListResultIterator
1. WorkloadClassifierListResultPage
1. WorkloadClassifierProperties
1. WorkloadClassifiersClient
1. WorkloadClassifiersCreateOrUpdateFuture
1. WorkloadClassifiersDeleteFuture
1. WorkloadGroup
1. WorkloadGroupListResult
1. WorkloadGroupListResultIterator
1. WorkloadGroupListResultPage
1. WorkloadGroupProperties
1. WorkloadGroupsClient
1. WorkloadGroupsCreateOrUpdateFuture
1. WorkloadGroupsDeleteFuture
