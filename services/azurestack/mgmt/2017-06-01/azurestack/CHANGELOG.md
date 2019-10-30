## New Content

### New Constants

1. Category.ADFS
1. Category.AzureAD
1. CompatibilityIssue.ADFSIdentitySystemRequired
1. CompatibilityIssue.AzureADIdentitySystemRequired
1. CompatibilityIssue.CapacityBillingModelRequired
1. CompatibilityIssue.ConnectionToAzureRequired
1. CompatibilityIssue.ConnectionToInternetRequired
1. CompatibilityIssue.DevelopmentBillingModelRequired
1. CompatibilityIssue.DisconnectedEnvironmentRequired
1. CompatibilityIssue.HigherDeviceVersionRequired
1. CompatibilityIssue.LowerDeviceVersionRequired
1. CompatibilityIssue.PayAsYouGoBillingModelRequired

### New Funcs

1. PossibleCategoryValues() []Category
1. PossibleCompatibilityIssueValues() []CompatibilityIssue
1. ProductsClient.GetProduct(context.Context,string,string,string,*DeviceConfiguration) (Product,error)
1. ProductsClient.GetProductPreparer(context.Context,string,string,string,*DeviceConfiguration) (*http.Request,error)
1. ProductsClient.GetProductResponder(*http.Response) (Product,error)
1. ProductsClient.GetProductSender(*http.Request) (*http.Response,error)
1. ProductsClient.GetProducts(context.Context,string,string,*DeviceConfiguration) (ProductList,error)
1. ProductsClient.GetProductsPreparer(context.Context,string,string,*DeviceConfiguration) (*http.Request,error)
1. ProductsClient.GetProductsResponder(*http.Response) (ProductList,error)
1. ProductsClient.GetProductsSender(*http.Request) (*http.Response,error)
1. ProductsClient.UploadLog(context.Context,string,string,string,*MarketplaceProductLogUpdate) (ProductLog,error)
1. ProductsClient.UploadLogPreparer(context.Context,string,string,string,*MarketplaceProductLogUpdate) (*http.Request,error)
1. ProductsClient.UploadLogResponder(*http.Response) (ProductLog,error)
1. ProductsClient.UploadLogSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. Compatibility
1. DeviceConfiguration
1. MarketplaceProductLogUpdate
1. ProductLog

### New Struct Fields

1. ProductNestedProperties.Compatibility
