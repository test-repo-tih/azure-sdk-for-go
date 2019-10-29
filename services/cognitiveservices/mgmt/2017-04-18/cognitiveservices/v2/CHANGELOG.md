## Breaking Changes

### Removed Constants

1. NetworkRuleBypassOptions.AzureServices
1. NetworkRuleBypassOptions.None

### Removed Funcs

1. *Account.UnmarshalJSON([]byte) error
1. AccountCreateParameters.MarshalJSON() ([]byte,error)
1. AccountUpdateParameters.MarshalJSON() ([]byte,error)
1. CheckSkuAvailabilityClient.List(context.Context,string,CheckSkuAvailabilityParameter) (CheckSkuAvailabilityResultList,error)
1. CheckSkuAvailabilityClient.ListPreparer(context.Context,string,CheckSkuAvailabilityParameter) (*http.Request,error)
1. CheckSkuAvailabilityClient.ListResponder(*http.Response) (CheckSkuAvailabilityResultList,error)
1. CheckSkuAvailabilityClient.ListSender(*http.Request) (*http.Response,error)
1. NewCheckSkuAvailabilityClient(string) CheckSkuAvailabilityClient
1. NewCheckSkuAvailabilityClientWithBaseURI(string,string) CheckSkuAvailabilityClient
1. PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions

## Struct Changes

### Removed Structs

1. AccountCreateParameters
1. AccountUpdateParameters
1. CheckSkuAvailabilityClient

### Removed Struct Fields

1. Account.*AccountProperties
1. NetworkRuleSet.Bypass

## Signature Changes

### Funcs

1. AccountsClient.Create
	- Params
		- From: context.Context,string,string,AccountCreateParameters
		- To: context.Context,string,string,Account
1. AccountsClient.CreatePreparer
	- Params
		- From: context.Context,string,string,AccountCreateParameters
		- To: context.Context,string,string,Account
1. AccountsClient.Update
	- Params
		- From: context.Context,string,string,AccountUpdateParameters
		- To: context.Context,string,string,Account
1. AccountsClient.UpdatePreparer
	- Params
		- From: context.Context,string,string,AccountUpdateParameters
		- To: context.Context,string,string,Account

## New Content

### New Funcs

1. BaseClient.CheckSkuAvailability(context.Context,string,CheckSkuAvailabilityParameter) (CheckSkuAvailabilityResultList,error)
1. BaseClient.CheckSkuAvailabilityPreparer(context.Context,string,CheckSkuAvailabilityParameter) (*http.Request,error)
1. BaseClient.CheckSkuAvailabilityResponder(*http.Response) (CheckSkuAvailabilityResultList,error)
1. BaseClient.CheckSkuAvailabilitySender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. AccountAPIProperties

### New Struct Fields

1. Account.Properties
1. AccountProperties.APIProperties
