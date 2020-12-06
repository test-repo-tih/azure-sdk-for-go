## Breaking Changes

## Signature Changes

### Funcs

1. AccountsClient.ListKeys
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,ListKeyExpand
1. AccountsClient.ListKeysPreparer
	- Params
		- From: context.Context,string,string
		- To: context.Context,string,string,ListKeyExpand

## New Content

### New Constants

1. DirectoryServiceOptions.DirectoryServiceOptionsAD
1. ListKeyExpand.Kerb

### New Funcs

1. PossibleListKeyExpandValues() []ListKeyExpand

## Struct Changes

### New Structs

1. ActiveDirectoryProperties

### New Struct Fields

1. AzureFilesIdentityBasedAuthentication.ActiveDirectoryProperties
