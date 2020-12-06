## Breaking Changes

## Signature Changes

### Funcs

1. ResourceSkusClient.List
	- Params
		- From: context.Context
		- To: context.Context,string
1. ResourceSkusClient.ListComplete
	- Params
		- From: context.Context
		- To: context.Context,string
1. ResourceSkusClient.ListPreparer
	- Params
		- From: context.Context
		- To: context.Context,string

## New Content

### New Constants

1. DiskEncryptionSetIdentityType.SystemAssigned
1. EncryptionType.EncryptionAtRestWithCustomerKey
1. EncryptionType.EncryptionAtRestWithPlatformKey

### New Funcs

1. *DiskEncryptionSet.UnmarshalJSON([]byte) error
1. *DiskEncryptionSetListIterator.Next() error
1. *DiskEncryptionSetListIterator.NextWithContext(context.Context) error
1. *DiskEncryptionSetListPage.Next() error
1. *DiskEncryptionSetListPage.NextWithContext(context.Context) error
1. *DiskEncryptionSetUpdate.UnmarshalJSON([]byte) error
1. *DiskEncryptionSetsCreateOrUpdateFuture.Result(DiskEncryptionSetsClient) (DiskEncryptionSet,error)
1. *DiskEncryptionSetsDeleteFuture.Result(DiskEncryptionSetsClient) (autorest.Response,error)
1. *DiskEncryptionSetsUpdateFuture.Result(DiskEncryptionSetsClient) (DiskEncryptionSet,error)
1. DiskEncryptionSet.MarshalJSON() ([]byte,error)
1. DiskEncryptionSetList.IsEmpty() bool
1. DiskEncryptionSetListIterator.NotDone() bool
1. DiskEncryptionSetListIterator.Response() DiskEncryptionSetList
1. DiskEncryptionSetListIterator.Value() DiskEncryptionSet
1. DiskEncryptionSetListPage.NotDone() bool
1. DiskEncryptionSetListPage.Response() DiskEncryptionSetList
1. DiskEncryptionSetListPage.Values() []DiskEncryptionSet
1. DiskEncryptionSetUpdate.MarshalJSON() ([]byte,error)
1. DiskEncryptionSetsClient.CreateOrUpdate(context.Context,string,string,DiskEncryptionSet) (DiskEncryptionSetsCreateOrUpdateFuture,error)
1. DiskEncryptionSetsClient.CreateOrUpdatePreparer(context.Context,string,string,DiskEncryptionSet) (*http.Request,error)
1. DiskEncryptionSetsClient.CreateOrUpdateResponder(*http.Response) (DiskEncryptionSet,error)
1. DiskEncryptionSetsClient.CreateOrUpdateSender(*http.Request) (DiskEncryptionSetsCreateOrUpdateFuture,error)
1. DiskEncryptionSetsClient.Delete(context.Context,string,string) (DiskEncryptionSetsDeleteFuture,error)
1. DiskEncryptionSetsClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. DiskEncryptionSetsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. DiskEncryptionSetsClient.DeleteSender(*http.Request) (DiskEncryptionSetsDeleteFuture,error)
1. DiskEncryptionSetsClient.Get(context.Context,string,string) (DiskEncryptionSet,error)
1. DiskEncryptionSetsClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. DiskEncryptionSetsClient.GetResponder(*http.Response) (DiskEncryptionSet,error)
1. DiskEncryptionSetsClient.GetSender(*http.Request) (*http.Response,error)
1. DiskEncryptionSetsClient.List(context.Context) (DiskEncryptionSetListPage,error)
1. DiskEncryptionSetsClient.ListByResourceGroup(context.Context,string) (DiskEncryptionSetListPage,error)
1. DiskEncryptionSetsClient.ListByResourceGroupComplete(context.Context,string) (DiskEncryptionSetListIterator,error)
1. DiskEncryptionSetsClient.ListByResourceGroupPreparer(context.Context,string) (*http.Request,error)
1. DiskEncryptionSetsClient.ListByResourceGroupResponder(*http.Response) (DiskEncryptionSetList,error)
1. DiskEncryptionSetsClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. DiskEncryptionSetsClient.ListComplete(context.Context) (DiskEncryptionSetListIterator,error)
1. DiskEncryptionSetsClient.ListPreparer(context.Context) (*http.Request,error)
1. DiskEncryptionSetsClient.ListResponder(*http.Response) (DiskEncryptionSetList,error)
1. DiskEncryptionSetsClient.ListSender(*http.Request) (*http.Response,error)
1. DiskEncryptionSetsClient.Update(context.Context,string,string,DiskEncryptionSetUpdate) (DiskEncryptionSetsUpdateFuture,error)
1. DiskEncryptionSetsClient.UpdatePreparer(context.Context,string,string,DiskEncryptionSetUpdate) (*http.Request,error)
1. DiskEncryptionSetsClient.UpdateResponder(*http.Response) (DiskEncryptionSet,error)
1. DiskEncryptionSetsClient.UpdateSender(*http.Request) (DiskEncryptionSetsUpdateFuture,error)
1. NewDiskEncryptionSetListIterator(DiskEncryptionSetListPage) DiskEncryptionSetListIterator
1. NewDiskEncryptionSetListPage(func(context.Context, DiskEncryptionSetList) (DiskEncryptionSetList, error)) DiskEncryptionSetListPage
1. NewDiskEncryptionSetsClient(string) DiskEncryptionSetsClient
1. NewDiskEncryptionSetsClientWithBaseURI(string,string) DiskEncryptionSetsClient
1. PossibleDiskEncryptionSetIdentityTypeValues() []DiskEncryptionSetIdentityType
1. PossibleEncryptionTypeValues() []EncryptionType

## Struct Changes

### New Structs

1. AutomaticRepairsPolicy
1. DiskEncryptionSet
1. DiskEncryptionSetList
1. DiskEncryptionSetListIterator
1. DiskEncryptionSetListPage
1. DiskEncryptionSetUpdate
1. DiskEncryptionSetUpdateProperties
1. DiskEncryptionSetsClient
1. DiskEncryptionSetsCreateOrUpdateFuture
1. DiskEncryptionSetsDeleteFuture
1. DiskEncryptionSetsUpdateFuture
1. Encryption
1. EncryptionSetProperties
1. ResourceIdentity

### New Struct Fields

1. DiskProperties.Encryption
1. OSProfile.RequireGuestProvisionSignal
1. SnapshotProperties.Encryption
1. VirtualMachineScaleSetProperties.AutomaticRepairsPolicy
1. VirtualMachineScaleSetUpdateProperties.AutomaticRepairsPolicy
1. VirtualMachineScaleSetUpdateProperties.DoNotRunExtensionsOnOverprovisionedVMs
