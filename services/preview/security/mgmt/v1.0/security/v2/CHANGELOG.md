## Breaking Changes

### Removed Constants

1. ValueType.IPCidr

## Signature Changes

### Const Types

1. String changed type from ValueType to PropertyType

## New Content

### New Constants

1. ActionType.ActionTypeAutomationAction
1. ActionType.ActionTypeEventHub
1. ActionType.ActionTypeLogicApp
1. ActionType.ActionTypeSecurityEmail
1. ActionType.ActionTypeWorkspace
1. EventSource.Alerts
1. EventSource.Assessments
1. Operator.Contains
1. Operator.EndsWith
1. Operator.Equals
1. Operator.GreaterThan
1. Operator.GreaterThanOrEqualTo
1. Operator.LesserThan
1. Operator.LesserThanOrEqualTo
1. Operator.NotEquals
1. Operator.StartsWith
1. PropertyType.Boolean
1. PropertyType.Integer
1. PropertyType.Number
1. SubscriptionRbacRoles.AccountAdmin
1. SubscriptionRbacRoles.Contributor
1. SubscriptionRbacRoles.Owner
1. SubscriptionRbacRoles.ServiceAdmin
1. ValueType.ValueTypeIPCidr
1. ValueType.ValueTypeString

### New Funcs

1. *Automation.UnmarshalJSON([]byte) error
1. *AutomationListIterator.Next() error
1. *AutomationListIterator.NextWithContext(context.Context) error
1. *AutomationListPage.Next() error
1. *AutomationListPage.NextWithContext(context.Context) error
1. *AutomationProperties.UnmarshalJSON([]byte) error
1. Automation.MarshalJSON() ([]byte,error)
1. AutomationAction.AsAutomationAction() (*AutomationAction,bool)
1. AutomationAction.AsAutomationActionEventHub() (*AutomationActionEventHub,bool)
1. AutomationAction.AsAutomationActionLogicApp() (*AutomationActionLogicApp,bool)
1. AutomationAction.AsAutomationActionSecurityEmail() (*AutomationActionSecurityEmail,bool)
1. AutomationAction.AsAutomationActionWorkspace() (*AutomationActionWorkspace,bool)
1. AutomationAction.AsBasicAutomationAction() (BasicAutomationAction,bool)
1. AutomationAction.MarshalJSON() ([]byte,error)
1. AutomationActionEventHub.AsAutomationAction() (*AutomationAction,bool)
1. AutomationActionEventHub.AsAutomationActionEventHub() (*AutomationActionEventHub,bool)
1. AutomationActionEventHub.AsAutomationActionLogicApp() (*AutomationActionLogicApp,bool)
1. AutomationActionEventHub.AsAutomationActionSecurityEmail() (*AutomationActionSecurityEmail,bool)
1. AutomationActionEventHub.AsAutomationActionWorkspace() (*AutomationActionWorkspace,bool)
1. AutomationActionEventHub.AsBasicAutomationAction() (BasicAutomationAction,bool)
1. AutomationActionEventHub.MarshalJSON() ([]byte,error)
1. AutomationActionLogicApp.AsAutomationAction() (*AutomationAction,bool)
1. AutomationActionLogicApp.AsAutomationActionEventHub() (*AutomationActionEventHub,bool)
1. AutomationActionLogicApp.AsAutomationActionLogicApp() (*AutomationActionLogicApp,bool)
1. AutomationActionLogicApp.AsAutomationActionSecurityEmail() (*AutomationActionSecurityEmail,bool)
1. AutomationActionLogicApp.AsAutomationActionWorkspace() (*AutomationActionWorkspace,bool)
1. AutomationActionLogicApp.AsBasicAutomationAction() (BasicAutomationAction,bool)
1. AutomationActionLogicApp.MarshalJSON() ([]byte,error)
1. AutomationActionSecurityEmail.AsAutomationAction() (*AutomationAction,bool)
1. AutomationActionSecurityEmail.AsAutomationActionEventHub() (*AutomationActionEventHub,bool)
1. AutomationActionSecurityEmail.AsAutomationActionLogicApp() (*AutomationActionLogicApp,bool)
1. AutomationActionSecurityEmail.AsAutomationActionSecurityEmail() (*AutomationActionSecurityEmail,bool)
1. AutomationActionSecurityEmail.AsAutomationActionWorkspace() (*AutomationActionWorkspace,bool)
1. AutomationActionSecurityEmail.AsBasicAutomationAction() (BasicAutomationAction,bool)
1. AutomationActionSecurityEmail.MarshalJSON() ([]byte,error)
1. AutomationActionWorkspace.AsAutomationAction() (*AutomationAction,bool)
1. AutomationActionWorkspace.AsAutomationActionEventHub() (*AutomationActionEventHub,bool)
1. AutomationActionWorkspace.AsAutomationActionLogicApp() (*AutomationActionLogicApp,bool)
1. AutomationActionWorkspace.AsAutomationActionSecurityEmail() (*AutomationActionSecurityEmail,bool)
1. AutomationActionWorkspace.AsAutomationActionWorkspace() (*AutomationActionWorkspace,bool)
1. AutomationActionWorkspace.AsBasicAutomationAction() (BasicAutomationAction,bool)
1. AutomationActionWorkspace.MarshalJSON() ([]byte,error)
1. AutomationList.IsEmpty() bool
1. AutomationListIterator.NotDone() bool
1. AutomationListIterator.Response() AutomationList
1. AutomationListIterator.Value() Automation
1. AutomationListPage.NotDone() bool
1. AutomationListPage.Response() AutomationList
1. AutomationListPage.Values() []Automation
1. AutomationsClient.CreateOrUpdate(context.Context,string,string,Automation) (Automation,error)
1. AutomationsClient.CreateOrUpdatePreparer(context.Context,string,string,Automation) (*http.Request,error)
1. AutomationsClient.CreateOrUpdateResponder(*http.Response) (Automation,error)
1. AutomationsClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. AutomationsClient.Delete(context.Context,string,string) (autorest.Response,error)
1. AutomationsClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. AutomationsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. AutomationsClient.DeleteSender(*http.Request) (*http.Response,error)
1. AutomationsClient.Get(context.Context,string,string) (Automation,error)
1. AutomationsClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. AutomationsClient.GetResponder(*http.Response) (Automation,error)
1. AutomationsClient.GetSender(*http.Request) (*http.Response,error)
1. AutomationsClient.List(context.Context) (AutomationListPage,error)
1. AutomationsClient.ListByResourceGroup(context.Context,string) (AutomationListPage,error)
1. AutomationsClient.ListByResourceGroupComplete(context.Context,string) (AutomationListIterator,error)
1. AutomationsClient.ListByResourceGroupPreparer(context.Context,string) (*http.Request,error)
1. AutomationsClient.ListByResourceGroupResponder(*http.Response) (AutomationList,error)
1. AutomationsClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. AutomationsClient.ListComplete(context.Context) (AutomationListIterator,error)
1. AutomationsClient.ListPreparer(context.Context) (*http.Request,error)
1. AutomationsClient.ListResponder(*http.Response) (AutomationList,error)
1. AutomationsClient.ListSender(*http.Request) (*http.Response,error)
1. AutomationsClient.Validate(context.Context,string,string,Automation) (AutomationValidationStatus,error)
1. AutomationsClient.ValidatePreparer(context.Context,string,string,Automation) (*http.Request,error)
1. AutomationsClient.ValidateResponder(*http.Response) (AutomationValidationStatus,error)
1. AutomationsClient.ValidateSender(*http.Request) (*http.Response,error)
1. NewAutomationListIterator(AutomationListPage) AutomationListIterator
1. NewAutomationListPage(func(context.Context, AutomationList) (AutomationList, error)) AutomationListPage
1. NewAutomationsClient(string,string) AutomationsClient
1. NewAutomationsClientWithBaseURI(string,string,string) AutomationsClient
1. PossibleActionTypeValues() []ActionType
1. PossibleEventSourceValues() []EventSource
1. PossibleOperatorValues() []Operator
1. PossiblePropertyTypeValues() []PropertyType
1. PossibleSubscriptionRbacRolesValues() []SubscriptionRbacRoles
1. Tags.MarshalJSON() ([]byte,error)
1. TrackedResource.MarshalJSON() ([]byte,error)

## Struct Changes

### New Structs

1. Automation
1. AutomationAction
1. AutomationActionEventHub
1. AutomationActionLogicApp
1. AutomationActionSecurityEmail
1. AutomationActionWorkspace
1. AutomationList
1. AutomationListIterator
1. AutomationListPage
1. AutomationMetadata
1. AutomationProperties
1. AutomationRuleSet
1. AutomationScope
1. AutomationSource
1. AutomationTriggeringRule
1. AutomationValidationStatus
1. AutomationsClient
1. ETag
1. Tags
1. TrackedResource
