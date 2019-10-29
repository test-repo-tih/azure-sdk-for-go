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

## Struct Changes

### New Structs

1. AutomaticRepairsPolicy

### New Struct Fields

1. OSProfile.RequireGuestProvisionSignal
1. VirtualMachineScaleSetProperties.AutomaticRepairsPolicy
1. VirtualMachineScaleSetUpdateProperties.AutomaticRepairsPolicy
1. VirtualMachineScaleSetUpdateProperties.DoNotRunExtensionsOnOverprovisionedVMs
