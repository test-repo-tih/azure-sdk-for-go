## New Content

### New Constants

1. EvaluationJobStatus.Completed
1. EvaluationJobStatus.Failed
1. EvaluationJobStatus.NotSubmitted
1. EvaluationJobStatus.Pending

### New Funcs

1. EvaluationsClient.Create(context.Context,EvaluationContract) (Evaluation,error)
1. EvaluationsClient.CreatePreparer(context.Context,EvaluationContract) (*http.Request,error)
1. EvaluationsClient.CreateResponder(*http.Response) (Evaluation,error)
1. EvaluationsClient.CreateSender(*http.Request) (*http.Response,error)
1. EvaluationsClient.Delete(context.Context,string) (autorest.Response,error)
1. EvaluationsClient.DeletePreparer(context.Context,string) (*http.Request,error)
1. EvaluationsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. EvaluationsClient.DeleteSender(*http.Request) (*http.Response,error)
1. EvaluationsClient.Get(context.Context,string) (Evaluation,error)
1. EvaluationsClient.GetPreparer(context.Context,string) (*http.Request,error)
1. EvaluationsClient.GetResponder(*http.Response) (Evaluation,error)
1. EvaluationsClient.GetSender(*http.Request) (*http.Response,error)
1. EvaluationsClient.List(context.Context) (ListEvaluation,error)
1. EvaluationsClient.ListPreparer(context.Context) (*http.Request,error)
1. EvaluationsClient.ListResponder(*http.Response) (ListEvaluation,error)
1. EvaluationsClient.ListSender(*http.Request) (*http.Response,error)
1. LogClient.Delete(context.Context) (autorest.Response,error)
1. LogClient.DeletePreparer(context.Context) (*http.Request,error)
1. LogClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. LogClient.DeleteSender(*http.Request) (*http.Response,error)
1. LogClient.GetProperties(context.Context) (LogsProperties,error)
1. LogClient.GetPropertiesPreparer(context.Context) (*http.Request,error)
1. LogClient.GetPropertiesResponder(*http.Response) (LogsProperties,error)
1. LogClient.GetPropertiesSender(*http.Request) (*http.Response,error)
1. ModelClient.Delete(context.Context) (autorest.Response,error)
1. ModelClient.DeletePreparer(context.Context) (*http.Request,error)
1. ModelClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. ModelClient.DeleteSender(*http.Request) (*http.Response,error)
1. ModelClient.Get(context.Context) (ReadCloser,error)
1. ModelClient.GetPreparer(context.Context) (*http.Request,error)
1. ModelClient.GetProperties(context.Context) (ModelProperties,error)
1. ModelClient.GetPropertiesPreparer(context.Context) (*http.Request,error)
1. ModelClient.GetPropertiesResponder(*http.Response) (ModelProperties,error)
1. ModelClient.GetPropertiesSender(*http.Request) (*http.Response,error)
1. ModelClient.GetResponder(*http.Response) (ReadCloser,error)
1. ModelClient.GetSender(*http.Request) (*http.Response,error)
1. NewEvaluationsClient(string) EvaluationsClient
1. NewLogClient(string) LogClient
1. NewModelClient(string) ModelClient
1. NewPolicyClient(string) PolicyClient
1. NewServiceConfigurationClient(string) ServiceConfigurationClient
1. PolicyClient.Delete(context.Context) (PolicyContract,error)
1. PolicyClient.DeletePreparer(context.Context) (*http.Request,error)
1. PolicyClient.DeleteResponder(*http.Response) (PolicyContract,error)
1. PolicyClient.DeleteSender(*http.Request) (*http.Response,error)
1. PolicyClient.Get(context.Context) (PolicyContract,error)
1. PolicyClient.GetPreparer(context.Context) (*http.Request,error)
1. PolicyClient.GetResponder(*http.Response) (PolicyContract,error)
1. PolicyClient.GetSender(*http.Request) (*http.Response,error)
1. PolicyClient.Update(context.Context,PolicyContract) (PolicyContract,error)
1. PolicyClient.UpdatePreparer(context.Context,PolicyContract) (*http.Request,error)
1. PolicyClient.UpdateResponder(*http.Response) (PolicyContract,error)
1. PolicyClient.UpdateSender(*http.Request) (*http.Response,error)
1. PossibleEvaluationJobStatusValues() []EvaluationJobStatus
1. ServiceConfigurationClient.Get(context.Context) (ServiceConfiguration,error)
1. ServiceConfigurationClient.GetPreparer(context.Context) (*http.Request,error)
1. ServiceConfigurationClient.GetResponder(*http.Response) (ServiceConfiguration,error)
1. ServiceConfigurationClient.GetSender(*http.Request) (*http.Response,error)
1. ServiceConfigurationClient.Update(context.Context,ServiceConfiguration) (ServiceConfiguration,error)
1. ServiceConfigurationClient.UpdatePreparer(context.Context,ServiceConfiguration) (*http.Request,error)
1. ServiceConfigurationClient.UpdateResponder(*http.Response) (ServiceConfiguration,error)
1. ServiceConfigurationClient.UpdateSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. ContainerStatus
1. DateRange
1. Evaluation
1. EvaluationContract
1. EvaluationsClient
1. ListEvaluation
1. LogClient
1. LogsProperties
1. LogsPropertiesDateRange
1. ModelClient
1. ModelProperties
1. PolicyClient
1. PolicyContract
1. PolicyResult
1. PolicyResultSummary
1. PolicyResultTotalSummary
1. ReadCloser
1. ServiceConfiguration
1. ServiceConfigurationClient
