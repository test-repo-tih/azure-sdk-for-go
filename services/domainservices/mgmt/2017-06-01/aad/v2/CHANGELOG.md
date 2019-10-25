## Breaking Changes

### Removed Funcs

1. *ReplicaSet.UnmarshalJSON([]byte) error
1. *ReplicaSetListResultIterator.Next() error
1. *ReplicaSetListResultIterator.NextWithContext(context.Context) error
1. *ReplicaSetListResultPage.Next() error
1. *ReplicaSetListResultPage.NextWithContext(context.Context) error
1. *ReplicaSetsCreateOrUpdateFuture.Result(ReplicaSetsClient) (ReplicaSet,error)
1. *ReplicaSetsDeleteFuture.Result(ReplicaSetsClient) (ReplicaSet,error)
1. *ReplicaSetsUpdateFuture.Result(ReplicaSetsClient) (ReplicaSet,error)
1. NewReplicaSetListResultIterator(ReplicaSetListResultPage) ReplicaSetListResultIterator
1. NewReplicaSetListResultPage(func(context.Context, ReplicaSetListResult) (ReplicaSetListResult, error)) ReplicaSetListResultPage
1. NewReplicaSetsClient(string) ReplicaSetsClient
1. NewReplicaSetsClientWithBaseURI(string,string) ReplicaSetsClient
1. ReplicaSet.MarshalJSON() ([]byte,error)
1. ReplicaSetListResult.IsEmpty() bool
1. ReplicaSetListResultIterator.NotDone() bool
1. ReplicaSetListResultIterator.Response() ReplicaSetListResult
1. ReplicaSetListResultIterator.Value() ReplicaSet
1. ReplicaSetListResultPage.NotDone() bool
1. ReplicaSetListResultPage.Response() ReplicaSetListResult
1. ReplicaSetListResultPage.Values() []ReplicaSet
1. ReplicaSetsClient.CreateOrUpdate(context.Context,string,string,string,ReplicaSet) (ReplicaSetsCreateOrUpdateFuture,error)
1. ReplicaSetsClient.CreateOrUpdatePreparer(context.Context,string,string,string,ReplicaSet) (*http.Request,error)
1. ReplicaSetsClient.CreateOrUpdateResponder(*http.Response) (ReplicaSet,error)
1. ReplicaSetsClient.CreateOrUpdateSender(*http.Request) (ReplicaSetsCreateOrUpdateFuture,error)
1. ReplicaSetsClient.Delete(context.Context,string,string,string) (ReplicaSetsDeleteFuture,error)
1. ReplicaSetsClient.DeletePreparer(context.Context,string,string,string) (*http.Request,error)
1. ReplicaSetsClient.DeleteResponder(*http.Response) (ReplicaSet,error)
1. ReplicaSetsClient.DeleteSender(*http.Request) (ReplicaSetsDeleteFuture,error)
1. ReplicaSetsClient.Get(context.Context,string,string,string) (ReplicaSet,error)
1. ReplicaSetsClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. ReplicaSetsClient.GetResponder(*http.Response) (ReplicaSet,error)
1. ReplicaSetsClient.GetSender(*http.Request) (*http.Response,error)
1. ReplicaSetsClient.ListByResourceGroup(context.Context,string,string) (ReplicaSetListResultPage,error)
1. ReplicaSetsClient.ListByResourceGroupComplete(context.Context,string,string) (ReplicaSetListResultIterator,error)
1. ReplicaSetsClient.ListByResourceGroupPreparer(context.Context,string,string) (*http.Request,error)
1. ReplicaSetsClient.ListByResourceGroupResponder(*http.Response) (ReplicaSetListResult,error)
1. ReplicaSetsClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. ReplicaSetsClient.Update(context.Context,string,string,string,ReplicaSet) (ReplicaSetsUpdateFuture,error)
1. ReplicaSetsClient.UpdatePreparer(context.Context,string,string,string,ReplicaSet) (*http.Request,error)
1. ReplicaSetsClient.UpdateResponder(*http.Response) (ReplicaSet,error)
1. ReplicaSetsClient.UpdateSender(*http.Request) (ReplicaSetsUpdateFuture,error)

## Struct Changes

### Removed Structs

1. ReplicaSet
1. ReplicaSetListResult
1. ReplicaSetListResultIterator
1. ReplicaSetListResultPage
1. ReplicaSetProperties
1. ReplicaSetsClient
1. ReplicaSetsCreateOrUpdateFuture
1. ReplicaSetsDeleteFuture
1. ReplicaSetsUpdateFuture

## Signature Changes

### Funcs

1. *DomainServicesDeleteFuture.Result
	- Returns
		- From: DomainService,error
		- To: autorest.Response,error
1. DomainServicesClient.DeleteResponder
	- Returns
		- From: DomainService,error
		- To: autorest.Response,error
