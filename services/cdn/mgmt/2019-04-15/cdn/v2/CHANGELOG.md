## Breaking Changes

## Signature Changes

### Const Types

1. Any changed type from PostArgsOperator to CookiesOperator
1. BeginsWith changed type from PostArgsOperator to CookiesOperator
1. Contains changed type from PostArgsOperator to CookiesOperator
1. EndsWith changed type from PostArgsOperator to CookiesOperator
1. Equal changed type from PostArgsOperator to CookiesOperator
1. GreaterThan changed type from PostArgsOperator to CookiesOperator
1. GreaterThanOrEqual changed type from PostArgsOperator to CookiesOperator
1. LessThan changed type from PostArgsOperator to CookiesOperator
1. LessThanOrEqual changed type from PostArgsOperator to CookiesOperator

## New Content

### New Constants

1. MinimumTLSVersion.None
1. MinimumTLSVersion.TLS10
1. MinimumTLSVersion.TLS12
1. Name.NameCookies
1. Name.NameHTTPVersion
1. PostArgsOperator.PostArgsOperatorAny
1. PostArgsOperator.PostArgsOperatorBeginsWith
1. PostArgsOperator.PostArgsOperatorContains
1. PostArgsOperator.PostArgsOperatorEndsWith
1. PostArgsOperator.PostArgsOperatorEqual
1. PostArgsOperator.PostArgsOperatorGreaterThan
1. PostArgsOperator.PostArgsOperatorGreaterThanOrEqual
1. PostArgsOperator.PostArgsOperatorLessThan
1. PostArgsOperator.PostArgsOperatorLessThanOrEqual

### New Funcs

1. *CustomDomainProperties.UnmarshalJSON([]byte) error
1. DeliveryRuleCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleCookiesCondition.AsBasicDeliveryRuleCondition() (BasicDeliveryRuleCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleCondition() (*DeliveryRuleCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleIsDeviceCondition() (*DeliveryRuleIsDeviceCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRulePostArgsCondition() (*DeliveryRulePostArgsCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleQueryStringCondition() (*DeliveryRuleQueryStringCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleRemoteAddressCondition() (*DeliveryRuleRemoteAddressCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleRequestBodyCondition() (*DeliveryRuleRequestBodyCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleRequestHeaderCondition() (*DeliveryRuleRequestHeaderCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleRequestMethodCondition() (*DeliveryRuleRequestMethodCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleRequestSchemeCondition() (*DeliveryRuleRequestSchemeCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleRequestURICondition() (*DeliveryRuleRequestURICondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleURLFileExtensionCondition() (*DeliveryRuleURLFileExtensionCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleURLFileNameCondition() (*DeliveryRuleURLFileNameCondition,bool)
1. DeliveryRuleCookiesCondition.AsDeliveryRuleURLPathCondition() (*DeliveryRuleURLPathCondition,bool)
1. DeliveryRuleCookiesCondition.MarshalJSON() ([]byte,error)
1. DeliveryRuleHTTPVersionCondition.AsBasicDeliveryRuleCondition() (BasicDeliveryRuleCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleCondition() (*DeliveryRuleCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleIsDeviceCondition() (*DeliveryRuleIsDeviceCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRulePostArgsCondition() (*DeliveryRulePostArgsCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleQueryStringCondition() (*DeliveryRuleQueryStringCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleRemoteAddressCondition() (*DeliveryRuleRemoteAddressCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleRequestBodyCondition() (*DeliveryRuleRequestBodyCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleRequestHeaderCondition() (*DeliveryRuleRequestHeaderCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleRequestMethodCondition() (*DeliveryRuleRequestMethodCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleRequestSchemeCondition() (*DeliveryRuleRequestSchemeCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleRequestURICondition() (*DeliveryRuleRequestURICondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleURLFileExtensionCondition() (*DeliveryRuleURLFileExtensionCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleURLFileNameCondition() (*DeliveryRuleURLFileNameCondition,bool)
1. DeliveryRuleHTTPVersionCondition.AsDeliveryRuleURLPathCondition() (*DeliveryRuleURLPathCondition,bool)
1. DeliveryRuleHTTPVersionCondition.MarshalJSON() ([]byte,error)
1. DeliveryRuleIsDeviceCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleIsDeviceCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRulePostArgsCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRulePostArgsCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleQueryStringCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleQueryStringCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleRemoteAddressCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleRemoteAddressCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleRequestBodyCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleRequestBodyCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleRequestHeaderCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleRequestHeaderCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleRequestMethodCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleRequestMethodCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleRequestSchemeCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleRequestSchemeCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleRequestURICondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleRequestURICondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleURLFileExtensionCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleURLFileExtensionCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleURLFileNameCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleURLFileNameCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. DeliveryRuleURLPathCondition.AsDeliveryRuleCookiesCondition() (*DeliveryRuleCookiesCondition,bool)
1. DeliveryRuleURLPathCondition.AsDeliveryRuleHTTPVersionCondition() (*DeliveryRuleHTTPVersionCondition,bool)
1. PossibleCookiesOperatorValues() []CookiesOperator
1. PossibleMinimumTLSVersionValues() []MinimumTLSVersion

## Struct Changes

### New Structs

1. CookiesMatchConditionParameters
1. DeliveryRuleCookiesCondition
1. DeliveryRuleHTTPVersionCondition
1. HTTPVersionMatchConditionParameters

### New Struct Fields

1. CustomDomainHTTPSParameters.MinimumTLSVersion
1. CustomDomainProperties.CustomHTTPSParameters
1. ManagedHTTPSParameters.MinimumTLSVersion
1. UserManagedHTTPSParameters.MinimumTLSVersion
