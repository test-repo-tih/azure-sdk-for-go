## Breaking Changes

### Removed Funcs

1. ServicesClient.Get(context.Context,string,string) (ServiceResource,error)
1. ServicesClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. ServicesClient.GetResponder(*http.Response) (ServiceResource,error)
1. ServicesClient.GetSender(*http.Request) (*http.Response,error)

## New Content

### New Funcs

1. NewServicesTestClient(string) ServicesTestClient
1. NewServicesTestClientWithBaseURI(string,string) ServicesTestClient
1. ServicesTestClient.Get(context.Context,string,string) (ServiceResource,error)
1. ServicesTestClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. ServicesTestClient.GetResponder(*http.Response) (ServiceResource,error)
1. ServicesTestClient.GetSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. ServicesTestClient
