## New Content

### New Funcs

1. AppsClient.CreateOrUpdateFunctionSecret(context.Context,string,string,string,string,KeyInfo) (KeyInfo,error)
1. AppsClient.CreateOrUpdateFunctionSecretPreparer(context.Context,string,string,string,string,KeyInfo) (*http.Request,error)
1. AppsClient.CreateOrUpdateFunctionSecretResponder(*http.Response) (KeyInfo,error)
1. AppsClient.CreateOrUpdateFunctionSecretSender(*http.Request) (*http.Response,error)
1. AppsClient.CreateOrUpdateFunctionSecretSlot(context.Context,string,string,string,string,string,KeyInfo) (KeyInfo,error)
1. AppsClient.CreateOrUpdateFunctionSecretSlotPreparer(context.Context,string,string,string,string,string,KeyInfo) (*http.Request,error)
1. AppsClient.CreateOrUpdateFunctionSecretSlotResponder(*http.Response) (KeyInfo,error)
1. AppsClient.CreateOrUpdateFunctionSecretSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.CreateOrUpdateHostSecret(context.Context,string,string,string,string,KeyInfo) (KeyInfo,error)
1. AppsClient.CreateOrUpdateHostSecretPreparer(context.Context,string,string,string,string,KeyInfo) (*http.Request,error)
1. AppsClient.CreateOrUpdateHostSecretResponder(*http.Response) (KeyInfo,error)
1. AppsClient.CreateOrUpdateHostSecretSender(*http.Request) (*http.Response,error)
1. AppsClient.CreateOrUpdateHostSecretSlot(context.Context,string,string,string,string,string,KeyInfo) (KeyInfo,error)
1. AppsClient.CreateOrUpdateHostSecretSlotPreparer(context.Context,string,string,string,string,string,KeyInfo) (*http.Request,error)
1. AppsClient.CreateOrUpdateHostSecretSlotResponder(*http.Response) (KeyInfo,error)
1. AppsClient.CreateOrUpdateHostSecretSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.DeleteFunctionSecret(context.Context,string,string,string,string) (autorest.Response,error)
1. AppsClient.DeleteFunctionSecretPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. AppsClient.DeleteFunctionSecretResponder(*http.Response) (autorest.Response,error)
1. AppsClient.DeleteFunctionSecretSender(*http.Request) (*http.Response,error)
1. AppsClient.DeleteFunctionSecretSlot(context.Context,string,string,string,string,string) (autorest.Response,error)
1. AppsClient.DeleteFunctionSecretSlotPreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. AppsClient.DeleteFunctionSecretSlotResponder(*http.Response) (autorest.Response,error)
1. AppsClient.DeleteFunctionSecretSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.DeleteHostSecret(context.Context,string,string,string,string) (autorest.Response,error)
1. AppsClient.DeleteHostSecretPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. AppsClient.DeleteHostSecretResponder(*http.Response) (autorest.Response,error)
1. AppsClient.DeleteHostSecretSender(*http.Request) (*http.Response,error)
1. AppsClient.DeleteHostSecretSlot(context.Context,string,string,string,string,string) (autorest.Response,error)
1. AppsClient.DeleteHostSecretSlotPreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. AppsClient.DeleteHostSecretSlotResponder(*http.Response) (autorest.Response,error)
1. AppsClient.DeleteHostSecretSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.ListFunctionKeys(context.Context,string,string,string) (StringDictionary,error)
1. AppsClient.ListFunctionKeysPreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.ListFunctionKeysResponder(*http.Response) (StringDictionary,error)
1. AppsClient.ListFunctionKeysSender(*http.Request) (*http.Response,error)
1. AppsClient.ListFunctionKeysSlot(context.Context,string,string,string,string) (StringDictionary,error)
1. AppsClient.ListFunctionKeysSlotPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. AppsClient.ListFunctionKeysSlotResponder(*http.Response) (StringDictionary,error)
1. AppsClient.ListFunctionKeysSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.ListHostKeys(context.Context,string,string) (HostKeys,error)
1. AppsClient.ListHostKeysPreparer(context.Context,string,string) (*http.Request,error)
1. AppsClient.ListHostKeysResponder(*http.Response) (HostKeys,error)
1. AppsClient.ListHostKeysSender(*http.Request) (*http.Response,error)
1. AppsClient.ListHostKeysSlot(context.Context,string,string,string) (HostKeys,error)
1. AppsClient.ListHostKeysSlotPreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.ListHostKeysSlotResponder(*http.Response) (HostKeys,error)
1. AppsClient.ListHostKeysSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.ListSyncStatus(context.Context,string,string) (autorest.Response,error)
1. AppsClient.ListSyncStatusPreparer(context.Context,string,string) (*http.Request,error)
1. AppsClient.ListSyncStatusResponder(*http.Response) (autorest.Response,error)
1. AppsClient.ListSyncStatusSender(*http.Request) (*http.Response,error)
1. AppsClient.ListSyncStatusSlot(context.Context,string,string,string) (autorest.Response,error)
1. AppsClient.ListSyncStatusSlotPreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.ListSyncStatusSlotResponder(*http.Response) (autorest.Response,error)
1. AppsClient.ListSyncStatusSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.SyncFunctions(context.Context,string,string) (autorest.Response,error)
1. AppsClient.SyncFunctionsPreparer(context.Context,string,string) (*http.Request,error)
1. AppsClient.SyncFunctionsResponder(*http.Response) (autorest.Response,error)
1. AppsClient.SyncFunctionsSender(*http.Request) (*http.Response,error)
1. AppsClient.SyncFunctionsSlot(context.Context,string,string,string) (autorest.Response,error)
1. AppsClient.SyncFunctionsSlotPreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.SyncFunctionsSlotResponder(*http.Response) (autorest.Response,error)
1. AppsClient.SyncFunctionsSlotSender(*http.Request) (*http.Response,error)
1. HostKeys.MarshalJSON() ([]byte,error)

## Struct Changes

### New Structs

1. HostKeys
1. KeyInfo

### New Struct Fields

1. FunctionEnvelopeProperties.InvokeURLTemplate
1. FunctionEnvelopeProperties.IsDisabled
1. FunctionEnvelopeProperties.Language
1. FunctionEnvelopeProperties.TestDataHref
