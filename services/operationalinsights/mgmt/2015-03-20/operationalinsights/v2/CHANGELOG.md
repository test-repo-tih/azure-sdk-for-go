## Breaking Changes

### Removed Funcs

1. *WorkspacesGetSearchResultsFuture.Result(WorkspacesClient) (SearchResultsResponse,error)
1. SavedSearchesClient.GetResults(context.Context,string,string,string) (SearchResultsResponse,error)
1. SavedSearchesClient.GetResultsPreparer(context.Context,string,string,string) (*http.Request,error)
1. SavedSearchesClient.GetResultsResponder(*http.Response) (SearchResultsResponse,error)
1. SavedSearchesClient.GetResultsSender(*http.Request) (*http.Response,error)
1. WorkspacesClient.GetSearchResults(context.Context,string,string,SearchParameters) (WorkspacesGetSearchResultsFuture,error)
1. WorkspacesClient.GetSearchResultsPreparer(context.Context,string,string,SearchParameters) (*http.Request,error)
1. WorkspacesClient.GetSearchResultsResponder(*http.Response) (SearchResultsResponse,error)
1. WorkspacesClient.GetSearchResultsSender(*http.Request) (WorkspacesGetSearchResultsFuture,error)
1. WorkspacesClient.UpdateSearchResults(context.Context,string,string,string) (SearchResultsResponse,error)
1. WorkspacesClient.UpdateSearchResultsPreparer(context.Context,string,string,string) (*http.Request,error)
1. WorkspacesClient.UpdateSearchResultsResponder(*http.Response) (SearchResultsResponse,error)
1. WorkspacesClient.UpdateSearchResultsSender(*http.Request) (*http.Response,error)

## Struct Changes

### Removed Structs

1. SearchError
1. SearchHighlight
1. SearchParameters
1. SearchResultsResponse
1. WorkspacesGetSearchResultsFuture
