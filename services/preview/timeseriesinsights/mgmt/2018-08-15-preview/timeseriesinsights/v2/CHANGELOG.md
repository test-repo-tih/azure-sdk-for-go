## Breaking Changes

## Struct Changes

### Removed Struct Fields

1. LongTermEnvironmentMutableProperties.TimeSeriesIDProperties

## Signature Changes

### Funcs

1. EnvironmentsClient.Update
	- Params
		- From: context.Context,string,string,EnvironmentUpdateParameters
		- To: context.Context,string,string,StandardEnvironmentUpdateParameters
1. EnvironmentsClient.UpdatePreparer
	- Params
		- From: context.Context,string,string,EnvironmentUpdateParameters
		- To: context.Context,string,string,StandardEnvironmentUpdateParameters

## New Content

### New Constants

1. SkuName.P1
1. WarmStoragePropertiesState.WarmStoragePropertiesStateError
1. WarmStoragePropertiesState.WarmStoragePropertiesStateOk
1. WarmStoragePropertiesState.WarmStoragePropertiesStateUnknown

### New Funcs

1. *LongTermEnvironmentCreationProperties.UnmarshalJSON([]byte) error
1. *LongTermEnvironmentMutableProperties.UnmarshalJSON([]byte) error
1. *LongTermEnvironmentResourceProperties.UnmarshalJSON([]byte) error
1. *WarmStorageEnvironmentStatus.UnmarshalJSON([]byte) error
1. *WarmStoragePropertiesUsage.UnmarshalJSON([]byte) error
1. LongTermEnvironmentCreationProperties.MarshalJSON() ([]byte,error)
1. LongTermEnvironmentMutableProperties.MarshalJSON() ([]byte,error)
1. LongTermEnvironmentResourceProperties.MarshalJSON() ([]byte,error)
1. PossibleWarmStoragePropertiesStateValues() []WarmStoragePropertiesState
1. WarmStorageEnvironmentStatus.MarshalJSON() ([]byte,error)
1. WarmStoragePropertiesUsage.MarshalJSON() ([]byte,error)

## Struct Changes

### New Structs

1. WarmStorageEnvironmentStatus
1. WarmStoragePropertiesUsage
1. WarmStoragePropertiesUsageStateDetails
1. WarmStoreConfigurationProperties

### New Struct Fields

1. EnvironmentStatus.WarmStorage
1. LongTermEnvironmentCreationProperties.*WarmStoreConfigurationProperties
1. LongTermEnvironmentMutableProperties.*WarmStoreConfigurationProperties
1. LongTermEnvironmentResourceProperties.*WarmStoreConfigurationProperties
