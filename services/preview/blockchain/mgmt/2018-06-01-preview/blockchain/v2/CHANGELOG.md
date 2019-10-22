## Breaking Changes

### Removed Funcs

1. MembersClient.ListRegenerateAPIKeys(context.Context,string,string,*APIKey) (APIKeyCollection,error)
1. MembersClient.ListRegenerateAPIKeysPreparer(context.Context,string,string,*APIKey) (*http.Request,error)
1. MembersClient.ListRegenerateAPIKeysResponder(*http.Response) (APIKeyCollection,error)
1. MembersClient.ListRegenerateAPIKeysSender(*http.Request) (*http.Response,error)
1. TransactionNodesClient.ListRegenerateAPIKeys(context.Context,string,string,string,*APIKey) (APIKeyCollection,error)
1. TransactionNodesClient.ListRegenerateAPIKeysPreparer(context.Context,string,string,string,*APIKey) (*http.Request,error)
1. TransactionNodesClient.ListRegenerateAPIKeysResponder(*http.Response) (APIKeyCollection,error)
1. TransactionNodesClient.ListRegenerateAPIKeysSender(*http.Request) (*http.Response,error)

## New Content

### New Funcs

1. MembersClient.RegenerateAPIKeys(context.Context,string,string,*APIKey) (APIKeyCollection,error)
1. MembersClient.RegenerateAPIKeysPreparer(context.Context,string,string,*APIKey) (*http.Request,error)
1. MembersClient.RegenerateAPIKeysResponder(*http.Response) (APIKeyCollection,error)
1. MembersClient.RegenerateAPIKeysSender(*http.Request) (*http.Response,error)
1. TransactionNodesClient.RegenerateAPIKeys(context.Context,string,string,string,*APIKey) (APIKeyCollection,error)
1. TransactionNodesClient.RegenerateAPIKeysPreparer(context.Context,string,string,string,*APIKey) (*http.Request,error)
1. TransactionNodesClient.RegenerateAPIKeysResponder(*http.Response) (APIKeyCollection,error)
1. TransactionNodesClient.RegenerateAPIKeysSender(*http.Request) (*http.Response,error)
