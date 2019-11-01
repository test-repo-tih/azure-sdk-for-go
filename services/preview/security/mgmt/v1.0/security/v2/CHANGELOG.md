## Breaking Changes

### Removed Constants

1. Source.SourceAws

### Removed Funcs

1. AwsResourceDetails.AsAwsResourceDetails() (*AwsResourceDetails,bool)
1. AwsResourceDetails.AsAzureResourceDetails() (*AzureResourceDetails,bool)
1. AwsResourceDetails.AsBasicResourceDetails() (BasicResourceDetails,bool)
1. AwsResourceDetails.AsResourceDetails() (*ResourceDetails,bool)
1. AwsResourceDetails.MarshalJSON() ([]byte,error)
1. AzureResourceDetails.AsAwsResourceDetails() (*AwsResourceDetails,bool)
1. ResourceDetails.AsAwsResourceDetails() (*AwsResourceDetails,bool)

## Struct Changes

### Removed Structs

1. AwsResourceDetails

## New Content

### New Constants

1. Source.SourceOnPremise

### New Funcs

1. AzureResourceDetails.AsOnPremiseResourceDetails() (*OnPremiseResourceDetails,bool)
1. OnPremiseResourceDetails.AsAzureResourceDetails() (*AzureResourceDetails,bool)
1. OnPremiseResourceDetails.AsBasicResourceDetails() (BasicResourceDetails,bool)
1. OnPremiseResourceDetails.AsOnPremiseResourceDetails() (*OnPremiseResourceDetails,bool)
1. OnPremiseResourceDetails.AsResourceDetails() (*ResourceDetails,bool)
1. OnPremiseResourceDetails.MarshalJSON() ([]byte,error)
1. ResourceDetails.AsOnPremiseResourceDetails() (*OnPremiseResourceDetails,bool)

## Struct Changes

### New Structs

1. OnPremiseResourceDetails
