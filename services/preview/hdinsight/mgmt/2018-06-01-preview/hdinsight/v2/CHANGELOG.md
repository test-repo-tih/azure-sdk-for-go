## Breaking Changes

### Removed Constants

1. ApplicationHTTPSEndpointAccessMode.WebPage
1. ApplicationType.CustomApplication
1. ApplicationType.RServer

### Removed Funcs

1. PossibleApplicationHTTPSEndpointAccessModeValues() []ApplicationHTTPSEndpointAccessMode
1. PossibleApplicationTypeValues() []ApplicationType

## Signature Changes

### Struct Fields

1. ApplicationGetHTTPSEndpoint.AccessModes changed type from *[]ApplicationHTTPSEndpointAccessMode to *[]string
1. ApplicationProperties.ApplicationType changed type from ApplicationType to *string
