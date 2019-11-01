## Breaking Changes

### Removed Constants

1. ResourceStatus.NotHealthy
1. ResourceStatus.OffByPolicy
1. Source.SourceAws

### Removed Funcs

1. AwsResourceDetails.AsAwsResourceDetails() (*AwsResourceDetails,bool)
1. AwsResourceDetails.AsAzureResourceDetails() (*AzureResourceDetails,bool)
1. AwsResourceDetails.AsBasicResourceDetails() (BasicResourceDetails,bool)
1. AwsResourceDetails.AsResourceDetails() (*ResourceDetails,bool)
1. AwsResourceDetails.MarshalJSON() ([]byte,error)
1. AzureResourceDetails.AsAwsResourceDetails() (*AwsResourceDetails,bool)
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
1. ResourceDetails.AsAwsResourceDetails() (*AwsResourceDetails,bool)

## Struct Changes

### Removed Structs

1. AwsResourceDetails
1. IoTSecuritySolutionsAnalyticsAggregatedAlertsClient
1. IoTSecuritySolutionsAnalyticsClient
1. IoTSecuritySolutionsAnalyticsRecommendationsClient
1. IoTSecuritySolutionsClient
1. IoTSecuritySolutionsResourceGroupClient

## Signature Changes

### Const Types

1. Healthy changed type from ResourceStatus to AssessmentStatusCode
1. NotApplicable changed type from ResourceStatus to AssessmentStatusCode

## New Content

### New Constants

1. AssessmentStatusCode.Unhealthy
1. ExpandEnum.Links
1. ExpandEnum.Metadata
1. ResourceStatus.ResourceStatusHealthy
1. ResourceStatus.ResourceStatusNotApplicable
1. ResourceStatus.ResourceStatusNotHealthy
1. ResourceStatus.ResourceStatusOffByPolicy
1. Source.SourceOnPremise

### New Funcs

1. *Assessment.UnmarshalJSON([]byte) error
1. *AssessmentListIterator.Next() error
1. *AssessmentListIterator.NextWithContext(context.Context) error
1. *AssessmentListPage.Next() error
1. *AssessmentListPage.NextWithContext(context.Context) error
1. *AssessmentProperties.UnmarshalJSON([]byte) error
1. Assessment.MarshalJSON() ([]byte,error)
1. AssessmentList.IsEmpty() bool
1. AssessmentListIterator.NotDone() bool
1. AssessmentListIterator.Response() AssessmentList
1. AssessmentListIterator.Value() Assessment
1. AssessmentListPage.NotDone() bool
1. AssessmentListPage.Response() AssessmentList
1. AssessmentListPage.Values() []Assessment
1. AssessmentProperties.MarshalJSON() ([]byte,error)
1. AssessmentsClient.Create(context.Context,string,string,Assessment) (Assessment,error)
1. AssessmentsClient.CreatePreparer(context.Context,string,string,Assessment) (*http.Request,error)
1. AssessmentsClient.CreateResponder(*http.Response) (Assessment,error)
1. AssessmentsClient.CreateSender(*http.Request) (*http.Response,error)
1. AssessmentsClient.Get(context.Context,string,string,ExpandEnum) (Assessment,error)
1. AssessmentsClient.GetPreparer(context.Context,string,string,ExpandEnum) (*http.Request,error)
1. AssessmentsClient.GetResponder(*http.Response) (Assessment,error)
1. AssessmentsClient.GetSender(*http.Request) (*http.Response,error)
1. AssessmentsClient.List(context.Context,string) (AssessmentListPage,error)
1. AssessmentsClient.ListComplete(context.Context,string) (AssessmentListIterator,error)
1. AssessmentsClient.ListPreparer(context.Context,string) (*http.Request,error)
1. AssessmentsClient.ListResponder(*http.Response) (AssessmentList,error)
1. AssessmentsClient.ListSender(*http.Request) (*http.Response,error)
1. AzureResourceDetails.AsOnPremiseResourceDetails() (*OnPremiseResourceDetails,bool)
1. NewAssessmentListIterator(AssessmentListPage) AssessmentListIterator
1. NewAssessmentListPage(func(context.Context, AssessmentList) (AssessmentList, error)) AssessmentListPage
1. NewAssessmentsClient(string,string) AssessmentsClient
1. NewAssessmentsClientWithBaseURI(string,string,string) AssessmentsClient
1. OnPremiseResourceDetails.AsAzureResourceDetails() (*AzureResourceDetails,bool)
1. OnPremiseResourceDetails.AsBasicResourceDetails() (BasicResourceDetails,bool)
1. OnPremiseResourceDetails.AsOnPremiseResourceDetails() (*OnPremiseResourceDetails,bool)
1. OnPremiseResourceDetails.AsResourceDetails() (*ResourceDetails,bool)
1. OnPremiseResourceDetails.MarshalJSON() ([]byte,error)
1. PossibleAssessmentStatusCodeValues() []AssessmentStatusCode
1. PossibleExpandEnumValues() []ExpandEnum
1. ResourceDetails.AsOnPremiseResourceDetails() (*OnPremiseResourceDetails,bool)

## Struct Changes

### New Structs

1. Assessment
1. AssessmentLinks
1. AssessmentList
1. AssessmentListIterator
1. AssessmentListPage
1. AssessmentProperties
1. AssessmentStatus
1. AssessmentsClient
1. IoTSecurityAggregatedAlertPropertiesTopDevicesListItem
1. OnPremiseResourceDetails

### New Struct Fields

1. IoTSecurityAggregatedAlertProperties.TopDevicesList
