## Breaking Changes

### Removed Funcs

1. IoTSecuritySolutionsAnalyticsAggregatedAlertsClient.List(context.Context,string,string,*int32) (IoTSecurityAggregatedAlertListPage,error)
1. IoTSecuritySolutionsAnalyticsAggregatedAlertsClient.ListComplete(context.Context,string,string,*int32) (IoTSecurityAggregatedAlertListIterator,error)
1. IoTSecuritySolutionsAnalyticsAggregatedAlertsClient.ListPreparer(context.Context,string,string,*int32) (*http.Request,error)
1. IoTSecuritySolutionsAnalyticsAggregatedAlertsClient.ListResponder(*http.Response) (IoTSecurityAggregatedAlertList,error)
1. IoTSecuritySolutionsAnalyticsAggregatedAlertsClient.ListSender(*http.Request) (*http.Response,error)
1. IoTSecuritySolutionsAnalyticsClient.GetAll(context.Context,string,string) (IoTSecuritySolutionAnalyticsModelList,error)
1. IoTSecuritySolutionsAnalyticsClient.GetAllPreparer(context.Context,string,string) (*http.Request,error)
1. IoTSecuritySolutionsAnalyticsClient.GetAllResponder(*http.Response) (IoTSecuritySolutionAnalyticsModelList,error)
1. IoTSecuritySolutionsAnalyticsClient.GetAllSender(*http.Request) (*http.Response,error)
1. IoTSecuritySolutionsAnalyticsClient.GetDefault(context.Context,string,string) (IoTSecuritySolutionAnalyticsModel,error)
1. IoTSecuritySolutionsAnalyticsClient.GetDefaultPreparer(context.Context,string,string) (*http.Request,error)
1. IoTSecuritySolutionsAnalyticsClient.GetDefaultResponder(*http.Response) (IoTSecuritySolutionAnalyticsModel,error)
1. IoTSecuritySolutionsAnalyticsClient.GetDefaultSender(*http.Request) (*http.Response,error)
1. IoTSecuritySolutionsAnalyticsRecommendationsClient.List(context.Context,string,string,*int32) (IoTSecurityAggregatedRecommendationListPage,error)
1. IoTSecuritySolutionsAnalyticsRecommendationsClient.ListComplete(context.Context,string,string,*int32) (IoTSecurityAggregatedRecommendationListIterator,error)
1. IoTSecuritySolutionsAnalyticsRecommendationsClient.ListPreparer(context.Context,string,string,*int32) (*http.Request,error)
1. IoTSecuritySolutionsAnalyticsRecommendationsClient.ListResponder(*http.Response) (IoTSecurityAggregatedRecommendationList,error)
1. IoTSecuritySolutionsAnalyticsRecommendationsClient.ListSender(*http.Request) (*http.Response,error)
1. IoTSecuritySolutionsClient.List(context.Context,string) (IoTSecuritySolutionsListPage,error)
1. IoTSecuritySolutionsClient.ListComplete(context.Context,string) (IoTSecuritySolutionsListIterator,error)
1. IoTSecuritySolutionsClient.ListPreparer(context.Context,string) (*http.Request,error)
1. IoTSecuritySolutionsClient.ListResponder(*http.Response) (IoTSecuritySolutionsList,error)
1. IoTSecuritySolutionsClient.ListSender(*http.Request) (*http.Response,error)
1. IoTSecuritySolutionsResourceGroupClient.List(context.Context,string,string) (IoTSecuritySolutionsListPage,error)
1. IoTSecuritySolutionsResourceGroupClient.ListComplete(context.Context,string,string) (IoTSecuritySolutionsListIterator,error)
1. IoTSecuritySolutionsResourceGroupClient.ListPreparer(context.Context,string,string) (*http.Request,error)
1. IoTSecuritySolutionsResourceGroupClient.ListResponder(*http.Response) (IoTSecuritySolutionsList,error)
1. IoTSecuritySolutionsResourceGroupClient.ListSender(*http.Request) (*http.Response,error)
1. NewIoTSecuritySolutionsAnalyticsAggregatedAlertsClient(string,string) IoTSecuritySolutionsAnalyticsAggregatedAlertsClient
1. NewIoTSecuritySolutionsAnalyticsAggregatedAlertsClientWithBaseURI(string,string,string) IoTSecuritySolutionsAnalyticsAggregatedAlertsClient
1. NewIoTSecuritySolutionsAnalyticsClient(string,string) IoTSecuritySolutionsAnalyticsClient
1. NewIoTSecuritySolutionsAnalyticsClientWithBaseURI(string,string,string) IoTSecuritySolutionsAnalyticsClient
1. NewIoTSecuritySolutionsAnalyticsRecommendationsClient(string,string) IoTSecuritySolutionsAnalyticsRecommendationsClient
1. NewIoTSecuritySolutionsAnalyticsRecommendationsClientWithBaseURI(string,string,string) IoTSecuritySolutionsAnalyticsRecommendationsClient
1. NewIoTSecuritySolutionsClient(string,string) IoTSecuritySolutionsClient
1. NewIoTSecuritySolutionsClientWithBaseURI(string,string,string) IoTSecuritySolutionsClient
1. NewIoTSecuritySolutionsResourceGroupClient(string,string) IoTSecuritySolutionsResourceGroupClient
1. NewIoTSecuritySolutionsResourceGroupClientWithBaseURI(string,string,string) IoTSecuritySolutionsResourceGroupClient

## Struct Changes

### Removed Structs

1. IoTSecuritySolutionsAnalyticsAggregatedAlertsClient
1. IoTSecuritySolutionsAnalyticsClient
1. IoTSecuritySolutionsAnalyticsRecommendationsClient
1. IoTSecuritySolutionsClient
1. IoTSecuritySolutionsResourceGroupClient

## New Content

### New Constants

1. AssessmentType.BuiltIn
1. AssessmentType.CustomPolicy
1. AssessmentType.CustomerManaged
1. Category.Compute
1. Category.Data
1. Category.IdentityAndAccess
1. Category.IoT
1. Category.Networking

### New Funcs

1. *AssessmentMetadata.UnmarshalJSON([]byte) error
1. *AssessmentMetadataListIterator.Next() error
1. *AssessmentMetadataListIterator.NextWithContext(context.Context) error
1. *AssessmentMetadataListPage.Next() error
1. *AssessmentMetadataListPage.NextWithContext(context.Context) error
1. AssessmentMetadata.MarshalJSON() ([]byte,error)
1. AssessmentMetadataList.IsEmpty() bool
1. AssessmentMetadataListIterator.NotDone() bool
1. AssessmentMetadataListIterator.Response() AssessmentMetadataList
1. AssessmentMetadataListIterator.Value() AssessmentMetadata
1. AssessmentMetadataListPage.NotDone() bool
1. AssessmentMetadataListPage.Response() AssessmentMetadataList
1. AssessmentMetadataListPage.Values() []AssessmentMetadata
1. AssessmentsMetadataClient.Get(context.Context,string) (AssessmentMetadata,error)
1. AssessmentsMetadataClient.GetPreparer(context.Context,string) (*http.Request,error)
1. AssessmentsMetadataClient.GetResponder(*http.Response) (AssessmentMetadata,error)
1. AssessmentsMetadataClient.GetSender(*http.Request) (*http.Response,error)
1. AssessmentsMetadataClient.List(context.Context) (AssessmentMetadataListPage,error)
1. AssessmentsMetadataClient.ListComplete(context.Context) (AssessmentMetadataListIterator,error)
1. AssessmentsMetadataClient.ListPreparer(context.Context) (*http.Request,error)
1. AssessmentsMetadataClient.ListResponder(*http.Response) (AssessmentMetadataList,error)
1. AssessmentsMetadataClient.ListSender(*http.Request) (*http.Response,error)
1. AssessmentsMetadataSubscriptionClient.Create(context.Context,string,AssessmentMetadata) (AssessmentMetadata,error)
1. AssessmentsMetadataSubscriptionClient.CreatePreparer(context.Context,string,AssessmentMetadata) (*http.Request,error)
1. AssessmentsMetadataSubscriptionClient.CreateResponder(*http.Response) (AssessmentMetadata,error)
1. AssessmentsMetadataSubscriptionClient.CreateSender(*http.Request) (*http.Response,error)
1. AssessmentsMetadataSubscriptionClient.Get(context.Context,string) (AssessmentMetadata,error)
1. AssessmentsMetadataSubscriptionClient.GetPreparer(context.Context,string) (*http.Request,error)
1. AssessmentsMetadataSubscriptionClient.GetResponder(*http.Response) (AssessmentMetadata,error)
1. AssessmentsMetadataSubscriptionClient.GetSender(*http.Request) (*http.Response,error)
1. AssessmentsMetadataSubscriptionClient.List(context.Context) (AssessmentMetadataListPage,error)
1. AssessmentsMetadataSubscriptionClient.ListComplete(context.Context) (AssessmentMetadataListIterator,error)
1. AssessmentsMetadataSubscriptionClient.ListPreparer(context.Context) (*http.Request,error)
1. AssessmentsMetadataSubscriptionClient.ListResponder(*http.Response) (AssessmentMetadataList,error)
1. AssessmentsMetadataSubscriptionClient.ListSender(*http.Request) (*http.Response,error)
1. NewAssessmentMetadataListIterator(AssessmentMetadataListPage) AssessmentMetadataListIterator
1. NewAssessmentMetadataListPage(func(context.Context, AssessmentMetadataList) (AssessmentMetadataList, error)) AssessmentMetadataListPage
1. NewAssessmentsMetadataClient(string,string) AssessmentsMetadataClient
1. NewAssessmentsMetadataClientWithBaseURI(string,string,string) AssessmentsMetadataClient
1. NewAssessmentsMetadataSubscriptionClient(string,string) AssessmentsMetadataSubscriptionClient
1. NewAssessmentsMetadataSubscriptionClientWithBaseURI(string,string,string) AssessmentsMetadataSubscriptionClient
1. PossibleAssessmentTypeValues() []AssessmentType
1. PossibleCategoryValues() []Category

## Struct Changes

### New Structs

1. AssessmentMetadata
1. AssessmentMetadataList
1. AssessmentMetadataListIterator
1. AssessmentMetadataListPage
1. AssessmentMetadataProperties
1. AssessmentsMetadataClient
1. AssessmentsMetadataSubscriptionClient
1. IoTSecurityAggregatedAlertPropertiesTopDevicesListItem

### New Struct Fields

1. IoTSecurityAggregatedAlertProperties.TopDevicesList
