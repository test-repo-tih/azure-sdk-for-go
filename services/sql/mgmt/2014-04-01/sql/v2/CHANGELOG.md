## Breaking Changes

### Removed Constants

1. CreateMode.Copy
1. CreateMode.NonReadableSecondary
1. CreateMode.OnlineSecondary
1. CreateMode.PointInTimeRestore
1. CreateMode.Recovery
1. CreateMode.Restore
1. CreateMode.RestoreLongTermRetentionBackup

## Signature Changes

### Const Types

1. Creating changed type from ElasticPoolState to DisasterRecoveryConfigurationStatus
1. Default changed type from CreateMode to AutoExecuteStatus
1. Disabled changed type from ElasticPoolState to AutoExecuteStatus
1. Ready changed type from ElasticPoolState to DisasterRecoveryConfigurationStatus

## New Content

### New Constants

1. AdvisorStatus.GA
1. AdvisorStatus.LimitedPublicPreview
1. AdvisorStatus.PrivatePreview
1. AdvisorStatus.PublicPreview
1. AutoExecuteStatus.Enabled
1. BackupLongTermRetentionPolicyState.BackupLongTermRetentionPolicyStateDisabled
1. BackupLongTermRetentionPolicyState.BackupLongTermRetentionPolicyStateEnabled
1. CapabilityStatus.CapabilityStatusAvailable
1. CapabilityStatus.CapabilityStatusDefault
1. CapabilityStatus.CapabilityStatusDisabled
1. CapabilityStatus.CapabilityStatusVisible
1. CreateMode.CreateModeCopy
1. CreateMode.CreateModeDefault
1. CreateMode.CreateModeNonReadableSecondary
1. CreateMode.CreateModeOnlineSecondary
1. CreateMode.CreateModePointInTimeRestore
1. CreateMode.CreateModeRecovery
1. CreateMode.CreateModeRestore
1. CreateMode.CreateModeRestoreLongTermRetentionBackup
1. DataMaskingFunction.DataMaskingFunctionCCN
1. DataMaskingFunction.DataMaskingFunctionDefault
1. DataMaskingFunction.DataMaskingFunctionEmail
1. DataMaskingFunction.DataMaskingFunctionNumber
1. DataMaskingFunction.DataMaskingFunctionSSN
1. DataMaskingFunction.DataMaskingFunctionText
1. DataMaskingRuleState.DataMaskingRuleStateDisabled
1. DataMaskingRuleState.DataMaskingRuleStateEnabled
1. DataMaskingState.DataMaskingStateDisabled
1. DataMaskingState.DataMaskingStateEnabled
1. DisasterRecoveryConfigurationAutoFailover.Off
1. DisasterRecoveryConfigurationAutoFailover.On
1. DisasterRecoveryConfigurationFailoverPolicy.Automatic
1. DisasterRecoveryConfigurationFailoverPolicy.Manual
1. DisasterRecoveryConfigurationRole.None
1. DisasterRecoveryConfigurationRole.Primary
1. DisasterRecoveryConfigurationRole.Secondary
1. DisasterRecoveryConfigurationStatus.Dropping
1. DisasterRecoveryConfigurationStatus.FailingOver
1. ElasticPoolState.ElasticPoolStateCreating
1. ElasticPoolState.ElasticPoolStateDisabled
1. ElasticPoolState.ElasticPoolStateReady
1. GeoBackupPolicyState.GeoBackupPolicyStateDisabled
1. GeoBackupPolicyState.GeoBackupPolicyStateEnabled
1. MaxSizeUnits.Gigabytes
1. MaxSizeUnits.Megabytes
1. MaxSizeUnits.Petabytes
1. MaxSizeUnits.Terabytes
1. PerformanceLevelUnit.DTU
1. PrimaryAggregationType.PrimaryAggregationTypeAverage
1. PrimaryAggregationType.PrimaryAggregationTypeCount
1. PrimaryAggregationType.PrimaryAggregationTypeMaximum
1. PrimaryAggregationType.PrimaryAggregationTypeMinimum
1. PrimaryAggregationType.PrimaryAggregationTypeNone
1. PrimaryAggregationType.PrimaryAggregationTypeTotal
1. QueryAggregationFunction.Avg
1. QueryAggregationFunction.Max
1. QueryAggregationFunction.Min
1. QueryAggregationFunction.Sum
1. QueryExecutionType.Aborted
1. QueryExecutionType.Any
1. QueryExecutionType.Exception
1. QueryExecutionType.Irregular
1. QueryExecutionType.Regular
1. QueryMetricUnit.KB
1. QueryMetricUnit.Microseconds
1. QueryMetricUnit.Percentage
1. QueryObservedMetricType.CPU
1. QueryObservedMetricType.Duration
1. QueryObservedMetricType.ExecutionCount
1. QueryObservedMetricType.Io
1. QueryObservedMetricType.Logio
1. RestorePointType.CONTINUOUS
1. RestorePointType.DISCRETE
1. ServerConnectionType.ServerConnectionTypeDefault
1. ServerConnectionType.ServerConnectionTypeProxy
1. ServerConnectionType.ServerConnectionTypeRedirect
1. ServerState.ServerStateDisabled
1. ServerState.ServerStateReady
1. ServerVersion.OneTwoFullStopZero
1. ServerVersion.TwoFullStopZero
1. UnitDefinitionType.Bytes
1. UnitDefinitionType.BytesPerSecond
1. UnitDefinitionType.Count
1. UnitDefinitionType.CountPerSecond
1. UnitDefinitionType.Percent
1. UnitDefinitionType.Seconds
1. UnitType.UnitTypeBytes
1. UnitType.UnitTypeBytesPerSecond
1. UnitType.UnitTypeCount
1. UnitType.UnitTypeCountPerSecond
1. UnitType.UnitTypePercent
1. UnitType.UnitTypeSeconds

### New Funcs

1. *Advisor.UnmarshalJSON([]byte) error
1. *BackupLongTermRetentionPoliciesCreateOrUpdateFuture.Result(BackupLongTermRetentionPoliciesClient) (BackupLongTermRetentionPolicy,error)
1. *BackupLongTermRetentionPolicy.UnmarshalJSON([]byte) error
1. *BackupLongTermRetentionVault.UnmarshalJSON([]byte) error
1. *BackupLongTermRetentionVaultsCreateOrUpdateFuture.Result(BackupLongTermRetentionVaultsClient) (BackupLongTermRetentionVault,error)
1. *DataMaskingPolicy.UnmarshalJSON([]byte) error
1. *DataMaskingRule.UnmarshalJSON([]byte) error
1. *DatabaseConnectionPolicy.UnmarshalJSON([]byte) error
1. *DatabaseTableAuditingPolicy.UnmarshalJSON([]byte) error
1. *DisasterRecoveryConfiguration.UnmarshalJSON([]byte) error
1. *DisasterRecoveryConfigurationsCreateOrUpdateFuture.Result(DisasterRecoveryConfigurationsClient) (DisasterRecoveryConfiguration,error)
1. *DisasterRecoveryConfigurationsDeleteFuture.Result(DisasterRecoveryConfigurationsClient) (autorest.Response,error)
1. *DisasterRecoveryConfigurationsFailoverAllowDataLossFuture.Result(DisasterRecoveryConfigurationsClient) (autorest.Response,error)
1. *DisasterRecoveryConfigurationsFailoverFuture.Result(DisasterRecoveryConfigurationsClient) (autorest.Response,error)
1. *GeoBackupPolicy.UnmarshalJSON([]byte) error
1. *RecoverableDatabase.UnmarshalJSON([]byte) error
1. *RestorableDroppedDatabase.UnmarshalJSON([]byte) error
1. *RestorePoint.UnmarshalJSON([]byte) error
1. *Server.UnmarshalJSON([]byte) error
1. *ServerAzureADAdministrator.UnmarshalJSON([]byte) error
1. *ServerAzureADAdministratorsCreateOrUpdateFuture.Result(ServerAzureADAdministratorsClient) (ServerAzureADAdministrator,error)
1. *ServerAzureADAdministratorsDeleteFuture.Result(ServerAzureADAdministratorsClient) (ServerAzureADAdministrator,error)
1. *ServerCommunicationLink.UnmarshalJSON([]byte) error
1. *ServerCommunicationLinksCreateOrUpdateFuture.Result(ServerCommunicationLinksClient) (ServerCommunicationLink,error)
1. *ServerConnectionPolicy.UnmarshalJSON([]byte) error
1. *ServerTableAuditingPolicy.UnmarshalJSON([]byte) error
1. *ServerUpdate.UnmarshalJSON([]byte) error
1. *ServiceObjective.UnmarshalJSON([]byte) error
1. *ServiceObjectiveCapability.UnmarshalJSON([]byte) error
1. Advisor.MarshalJSON() ([]byte,error)
1. BackupLongTermRetentionPoliciesClient.CreateOrUpdate(context.Context,string,string,string,BackupLongTermRetentionPolicy) (BackupLongTermRetentionPoliciesCreateOrUpdateFuture,error)
1. BackupLongTermRetentionPoliciesClient.CreateOrUpdatePreparer(context.Context,string,string,string,BackupLongTermRetentionPolicy) (*http.Request,error)
1. BackupLongTermRetentionPoliciesClient.CreateOrUpdateResponder(*http.Response) (BackupLongTermRetentionPolicy,error)
1. BackupLongTermRetentionPoliciesClient.CreateOrUpdateSender(*http.Request) (BackupLongTermRetentionPoliciesCreateOrUpdateFuture,error)
1. BackupLongTermRetentionPoliciesClient.Get(context.Context,string,string,string) (BackupLongTermRetentionPolicy,error)
1. BackupLongTermRetentionPoliciesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. BackupLongTermRetentionPoliciesClient.GetResponder(*http.Response) (BackupLongTermRetentionPolicy,error)
1. BackupLongTermRetentionPoliciesClient.GetSender(*http.Request) (*http.Response,error)
1. BackupLongTermRetentionPoliciesClient.ListByDatabase(context.Context,string,string,string) (BackupLongTermRetentionPolicyListResult,error)
1. BackupLongTermRetentionPoliciesClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. BackupLongTermRetentionPoliciesClient.ListByDatabaseResponder(*http.Response) (BackupLongTermRetentionPolicyListResult,error)
1. BackupLongTermRetentionPoliciesClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. BackupLongTermRetentionPolicy.MarshalJSON() ([]byte,error)
1. BackupLongTermRetentionVault.MarshalJSON() ([]byte,error)
1. BackupLongTermRetentionVaultsClient.CreateOrUpdate(context.Context,string,string,BackupLongTermRetentionVault) (BackupLongTermRetentionVaultsCreateOrUpdateFuture,error)
1. BackupLongTermRetentionVaultsClient.CreateOrUpdatePreparer(context.Context,string,string,BackupLongTermRetentionVault) (*http.Request,error)
1. BackupLongTermRetentionVaultsClient.CreateOrUpdateResponder(*http.Response) (BackupLongTermRetentionVault,error)
1. BackupLongTermRetentionVaultsClient.CreateOrUpdateSender(*http.Request) (BackupLongTermRetentionVaultsCreateOrUpdateFuture,error)
1. BackupLongTermRetentionVaultsClient.Get(context.Context,string,string) (BackupLongTermRetentionVault,error)
1. BackupLongTermRetentionVaultsClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. BackupLongTermRetentionVaultsClient.GetResponder(*http.Response) (BackupLongTermRetentionVault,error)
1. BackupLongTermRetentionVaultsClient.GetSender(*http.Request) (*http.Response,error)
1. BackupLongTermRetentionVaultsClient.ListByServer(context.Context,string,string) (BackupLongTermRetentionVaultListResult,error)
1. BackupLongTermRetentionVaultsClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. BackupLongTermRetentionVaultsClient.ListByServerResponder(*http.Response) (BackupLongTermRetentionVaultListResult,error)
1. BackupLongTermRetentionVaultsClient.ListByServerSender(*http.Request) (*http.Response,error)
1. CapabilitiesClient.ListByLocation(context.Context,string) (LocationCapabilities,error)
1. CapabilitiesClient.ListByLocationPreparer(context.Context,string) (*http.Request,error)
1. CapabilitiesClient.ListByLocationResponder(*http.Response) (LocationCapabilities,error)
1. CapabilitiesClient.ListByLocationSender(*http.Request) (*http.Response,error)
1. DataMaskingPoliciesClient.CreateOrUpdate(context.Context,string,string,string,DataMaskingPolicy) (DataMaskingPolicy,error)
1. DataMaskingPoliciesClient.CreateOrUpdatePreparer(context.Context,string,string,string,DataMaskingPolicy) (*http.Request,error)
1. DataMaskingPoliciesClient.CreateOrUpdateResponder(*http.Response) (DataMaskingPolicy,error)
1. DataMaskingPoliciesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. DataMaskingPoliciesClient.Get(context.Context,string,string,string) (DataMaskingPolicy,error)
1. DataMaskingPoliciesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. DataMaskingPoliciesClient.GetResponder(*http.Response) (DataMaskingPolicy,error)
1. DataMaskingPoliciesClient.GetSender(*http.Request) (*http.Response,error)
1. DataMaskingPolicy.MarshalJSON() ([]byte,error)
1. DataMaskingRule.MarshalJSON() ([]byte,error)
1. DataMaskingRulesClient.CreateOrUpdate(context.Context,string,string,string,string,DataMaskingRule) (DataMaskingRule,error)
1. DataMaskingRulesClient.CreateOrUpdatePreparer(context.Context,string,string,string,string,DataMaskingRule) (*http.Request,error)
1. DataMaskingRulesClient.CreateOrUpdateResponder(*http.Response) (DataMaskingRule,error)
1. DataMaskingRulesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. DataMaskingRulesClient.ListByDatabase(context.Context,string,string,string) (DataMaskingRuleListResult,error)
1. DataMaskingRulesClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. DataMaskingRulesClient.ListByDatabaseResponder(*http.Response) (DataMaskingRuleListResult,error)
1. DataMaskingRulesClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. DatabaseAdvisorsClient.CreateOrUpdate(context.Context,string,string,string,string,Advisor) (Advisor,error)
1. DatabaseAdvisorsClient.CreateOrUpdatePreparer(context.Context,string,string,string,string,Advisor) (*http.Request,error)
1. DatabaseAdvisorsClient.CreateOrUpdateResponder(*http.Response) (Advisor,error)
1. DatabaseAdvisorsClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. DatabaseAdvisorsClient.Get(context.Context,string,string,string,string) (Advisor,error)
1. DatabaseAdvisorsClient.GetPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. DatabaseAdvisorsClient.GetResponder(*http.Response) (Advisor,error)
1. DatabaseAdvisorsClient.GetSender(*http.Request) (*http.Response,error)
1. DatabaseAdvisorsClient.ListByDatabase(context.Context,string,string,string) (AdvisorListResult,error)
1. DatabaseAdvisorsClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. DatabaseAdvisorsClient.ListByDatabaseResponder(*http.Response) (AdvisorListResult,error)
1. DatabaseAdvisorsClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. DatabaseConnectionPoliciesClient.CreateOrUpdate(context.Context,string,string,string,DatabaseConnectionPolicy) (DatabaseConnectionPolicy,error)
1. DatabaseConnectionPoliciesClient.CreateOrUpdatePreparer(context.Context,string,string,string,DatabaseConnectionPolicy) (*http.Request,error)
1. DatabaseConnectionPoliciesClient.CreateOrUpdateResponder(*http.Response) (DatabaseConnectionPolicy,error)
1. DatabaseConnectionPoliciesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. DatabaseConnectionPoliciesClient.Get(context.Context,string,string,string) (DatabaseConnectionPolicy,error)
1. DatabaseConnectionPoliciesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. DatabaseConnectionPoliciesClient.GetResponder(*http.Response) (DatabaseConnectionPolicy,error)
1. DatabaseConnectionPoliciesClient.GetSender(*http.Request) (*http.Response,error)
1. DatabaseConnectionPolicy.MarshalJSON() ([]byte,error)
1. DatabaseTableAuditingPoliciesClient.CreateOrUpdate(context.Context,string,string,string,DatabaseTableAuditingPolicy) (DatabaseTableAuditingPolicy,error)
1. DatabaseTableAuditingPoliciesClient.CreateOrUpdatePreparer(context.Context,string,string,string,DatabaseTableAuditingPolicy) (*http.Request,error)
1. DatabaseTableAuditingPoliciesClient.CreateOrUpdateResponder(*http.Response) (DatabaseTableAuditingPolicy,error)
1. DatabaseTableAuditingPoliciesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. DatabaseTableAuditingPoliciesClient.Get(context.Context,string,string,string) (DatabaseTableAuditingPolicy,error)
1. DatabaseTableAuditingPoliciesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. DatabaseTableAuditingPoliciesClient.GetResponder(*http.Response) (DatabaseTableAuditingPolicy,error)
1. DatabaseTableAuditingPoliciesClient.GetSender(*http.Request) (*http.Response,error)
1. DatabaseTableAuditingPoliciesClient.ListByDatabase(context.Context,string,string,string) (DatabaseTableAuditingPolicyListResult,error)
1. DatabaseTableAuditingPoliciesClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. DatabaseTableAuditingPoliciesClient.ListByDatabaseResponder(*http.Response) (DatabaseTableAuditingPolicyListResult,error)
1. DatabaseTableAuditingPoliciesClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. DatabaseTableAuditingPolicy.MarshalJSON() ([]byte,error)
1. DatabaseUsagesClient.ListByDatabase(context.Context,string,string,string) (DatabaseUsageListResult,error)
1. DatabaseUsagesClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. DatabaseUsagesClient.ListByDatabaseResponder(*http.Response) (DatabaseUsageListResult,error)
1. DatabaseUsagesClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. DatabasesClient.ListMetricDefinitions(context.Context,string,string,string) (MetricDefinitionListResult,error)
1. DatabasesClient.ListMetricDefinitionsPreparer(context.Context,string,string,string) (*http.Request,error)
1. DatabasesClient.ListMetricDefinitionsResponder(*http.Response) (MetricDefinitionListResult,error)
1. DatabasesClient.ListMetricDefinitionsSender(*http.Request) (*http.Response,error)
1. DatabasesClient.ListMetrics(context.Context,string,string,string,string) (MetricListResult,error)
1. DatabasesClient.ListMetricsPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. DatabasesClient.ListMetricsResponder(*http.Response) (MetricListResult,error)
1. DatabasesClient.ListMetricsSender(*http.Request) (*http.Response,error)
1. DisasterRecoveryConfiguration.MarshalJSON() ([]byte,error)
1. DisasterRecoveryConfigurationsClient.CreateOrUpdate(context.Context,string,string,string) (DisasterRecoveryConfigurationsCreateOrUpdateFuture,error)
1. DisasterRecoveryConfigurationsClient.CreateOrUpdatePreparer(context.Context,string,string,string) (*http.Request,error)
1. DisasterRecoveryConfigurationsClient.CreateOrUpdateResponder(*http.Response) (DisasterRecoveryConfiguration,error)
1. DisasterRecoveryConfigurationsClient.CreateOrUpdateSender(*http.Request) (DisasterRecoveryConfigurationsCreateOrUpdateFuture,error)
1. DisasterRecoveryConfigurationsClient.Delete(context.Context,string,string,string) (DisasterRecoveryConfigurationsDeleteFuture,error)
1. DisasterRecoveryConfigurationsClient.DeletePreparer(context.Context,string,string,string) (*http.Request,error)
1. DisasterRecoveryConfigurationsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. DisasterRecoveryConfigurationsClient.DeleteSender(*http.Request) (DisasterRecoveryConfigurationsDeleteFuture,error)
1. DisasterRecoveryConfigurationsClient.Failover(context.Context,string,string,string) (DisasterRecoveryConfigurationsFailoverFuture,error)
1. DisasterRecoveryConfigurationsClient.FailoverAllowDataLoss(context.Context,string,string,string) (DisasterRecoveryConfigurationsFailoverAllowDataLossFuture,error)
1. DisasterRecoveryConfigurationsClient.FailoverAllowDataLossPreparer(context.Context,string,string,string) (*http.Request,error)
1. DisasterRecoveryConfigurationsClient.FailoverAllowDataLossResponder(*http.Response) (autorest.Response,error)
1. DisasterRecoveryConfigurationsClient.FailoverAllowDataLossSender(*http.Request) (DisasterRecoveryConfigurationsFailoverAllowDataLossFuture,error)
1. DisasterRecoveryConfigurationsClient.FailoverPreparer(context.Context,string,string,string) (*http.Request,error)
1. DisasterRecoveryConfigurationsClient.FailoverResponder(*http.Response) (autorest.Response,error)
1. DisasterRecoveryConfigurationsClient.FailoverSender(*http.Request) (DisasterRecoveryConfigurationsFailoverFuture,error)
1. DisasterRecoveryConfigurationsClient.Get(context.Context,string,string,string) (DisasterRecoveryConfiguration,error)
1. DisasterRecoveryConfigurationsClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. DisasterRecoveryConfigurationsClient.GetResponder(*http.Response) (DisasterRecoveryConfiguration,error)
1. DisasterRecoveryConfigurationsClient.GetSender(*http.Request) (*http.Response,error)
1. DisasterRecoveryConfigurationsClient.List(context.Context,string,string) (DisasterRecoveryConfigurationListResult,error)
1. DisasterRecoveryConfigurationsClient.ListPreparer(context.Context,string,string) (*http.Request,error)
1. DisasterRecoveryConfigurationsClient.ListResponder(*http.Response) (DisasterRecoveryConfigurationListResult,error)
1. DisasterRecoveryConfigurationsClient.ListSender(*http.Request) (*http.Response,error)
1. ElasticPoolsClient.ListMetricDefinitions(context.Context,string,string,string) (MetricDefinitionListResult,error)
1. ElasticPoolsClient.ListMetricDefinitionsPreparer(context.Context,string,string,string) (*http.Request,error)
1. ElasticPoolsClient.ListMetricDefinitionsResponder(*http.Response) (MetricDefinitionListResult,error)
1. ElasticPoolsClient.ListMetricDefinitionsSender(*http.Request) (*http.Response,error)
1. ElasticPoolsClient.ListMetrics(context.Context,string,string,string,string) (MetricListResult,error)
1. ElasticPoolsClient.ListMetricsPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. ElasticPoolsClient.ListMetricsResponder(*http.Response) (MetricListResult,error)
1. ElasticPoolsClient.ListMetricsSender(*http.Request) (*http.Response,error)
1. ExtensionsClient.Get(context.Context,string,string,string) (autorest.Response,error)
1. ExtensionsClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. ExtensionsClient.GetResponder(*http.Response) (autorest.Response,error)
1. ExtensionsClient.GetSender(*http.Request) (*http.Response,error)
1. ExtensionsClient.ListByDatabase(context.Context,string,string,string) (ExtensionListResult,error)
1. ExtensionsClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. ExtensionsClient.ListByDatabaseResponder(*http.Response) (ExtensionListResult,error)
1. ExtensionsClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. GeoBackupPoliciesClient.CreateOrUpdate(context.Context,string,string,string,GeoBackupPolicy) (GeoBackupPolicy,error)
1. GeoBackupPoliciesClient.CreateOrUpdatePreparer(context.Context,string,string,string,GeoBackupPolicy) (*http.Request,error)
1. GeoBackupPoliciesClient.CreateOrUpdateResponder(*http.Response) (GeoBackupPolicy,error)
1. GeoBackupPoliciesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. GeoBackupPoliciesClient.Get(context.Context,string,string,string) (GeoBackupPolicy,error)
1. GeoBackupPoliciesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. GeoBackupPoliciesClient.GetResponder(*http.Response) (GeoBackupPolicy,error)
1. GeoBackupPoliciesClient.GetSender(*http.Request) (*http.Response,error)
1. GeoBackupPoliciesClient.ListByDatabase(context.Context,string,string,string) (GeoBackupPolicyListResult,error)
1. GeoBackupPoliciesClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. GeoBackupPoliciesClient.ListByDatabaseResponder(*http.Response) (GeoBackupPolicyListResult,error)
1. GeoBackupPoliciesClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. GeoBackupPolicy.MarshalJSON() ([]byte,error)
1. NewBackupLongTermRetentionPoliciesClient(string) BackupLongTermRetentionPoliciesClient
1. NewBackupLongTermRetentionPoliciesClientWithBaseURI(string,string) BackupLongTermRetentionPoliciesClient
1. NewBackupLongTermRetentionVaultsClient(string) BackupLongTermRetentionVaultsClient
1. NewBackupLongTermRetentionVaultsClientWithBaseURI(string,string) BackupLongTermRetentionVaultsClient
1. NewCapabilitiesClient(string) CapabilitiesClient
1. NewCapabilitiesClientWithBaseURI(string,string) CapabilitiesClient
1. NewDataMaskingPoliciesClient(string) DataMaskingPoliciesClient
1. NewDataMaskingPoliciesClientWithBaseURI(string,string) DataMaskingPoliciesClient
1. NewDataMaskingRulesClient(string) DataMaskingRulesClient
1. NewDataMaskingRulesClientWithBaseURI(string,string) DataMaskingRulesClient
1. NewDatabaseAdvisorsClient(string) DatabaseAdvisorsClient
1. NewDatabaseAdvisorsClientWithBaseURI(string,string) DatabaseAdvisorsClient
1. NewDatabaseConnectionPoliciesClient(string) DatabaseConnectionPoliciesClient
1. NewDatabaseConnectionPoliciesClientWithBaseURI(string,string) DatabaseConnectionPoliciesClient
1. NewDatabaseTableAuditingPoliciesClient(string) DatabaseTableAuditingPoliciesClient
1. NewDatabaseTableAuditingPoliciesClientWithBaseURI(string,string) DatabaseTableAuditingPoliciesClient
1. NewDatabaseUsagesClient(string) DatabaseUsagesClient
1. NewDatabaseUsagesClientWithBaseURI(string,string) DatabaseUsagesClient
1. NewDisasterRecoveryConfigurationsClient(string) DisasterRecoveryConfigurationsClient
1. NewDisasterRecoveryConfigurationsClientWithBaseURI(string,string) DisasterRecoveryConfigurationsClient
1. NewExtensionsClient(string) ExtensionsClient
1. NewExtensionsClientWithBaseURI(string,string) ExtensionsClient
1. NewGeoBackupPoliciesClient(string) GeoBackupPoliciesClient
1. NewGeoBackupPoliciesClientWithBaseURI(string,string) GeoBackupPoliciesClient
1. NewOperationsClient(string) OperationsClient
1. NewOperationsClientWithBaseURI(string,string) OperationsClient
1. NewQueriesClient(string) QueriesClient
1. NewQueriesClientWithBaseURI(string,string) QueriesClient
1. NewQueryStatisticsClient(string) QueryStatisticsClient
1. NewQueryStatisticsClientWithBaseURI(string,string) QueryStatisticsClient
1. NewQueryTextsClient(string) QueryTextsClient
1. NewQueryTextsClientWithBaseURI(string,string) QueryTextsClient
1. NewRecoverableDatabasesClient(string) RecoverableDatabasesClient
1. NewRecoverableDatabasesClientWithBaseURI(string,string) RecoverableDatabasesClient
1. NewRestorableDroppedDatabasesClient(string) RestorableDroppedDatabasesClient
1. NewRestorableDroppedDatabasesClientWithBaseURI(string,string) RestorableDroppedDatabasesClient
1. NewRestorePointsClient(string) RestorePointsClient
1. NewRestorePointsClientWithBaseURI(string,string) RestorePointsClient
1. NewServerAdvisorsClient(string) ServerAdvisorsClient
1. NewServerAdvisorsClientWithBaseURI(string,string) ServerAdvisorsClient
1. NewServerAzureADAdministratorsClient(string) ServerAzureADAdministratorsClient
1. NewServerAzureADAdministratorsClientWithBaseURI(string,string) ServerAzureADAdministratorsClient
1. NewServerCommunicationLinksClient(string) ServerCommunicationLinksClient
1. NewServerCommunicationLinksClientWithBaseURI(string,string) ServerCommunicationLinksClient
1. NewServerConnectionPoliciesClient(string) ServerConnectionPoliciesClient
1. NewServerConnectionPoliciesClientWithBaseURI(string,string) ServerConnectionPoliciesClient
1. NewServerTableAuditingPoliciesClient(string) ServerTableAuditingPoliciesClient
1. NewServerTableAuditingPoliciesClientWithBaseURI(string,string) ServerTableAuditingPoliciesClient
1. NewServerUsagesClient(string) ServerUsagesClient
1. NewServerUsagesClientWithBaseURI(string,string) ServerUsagesClient
1. NewServiceObjectivesClient(string) ServiceObjectivesClient
1. NewServiceObjectivesClientWithBaseURI(string,string) ServiceObjectivesClient
1. NewTransparentDataEncryptionConfigurationsClient(string) TransparentDataEncryptionConfigurationsClient
1. NewTransparentDataEncryptionConfigurationsClientWithBaseURI(string,string) TransparentDataEncryptionConfigurationsClient
1. OperationsClient.List(context.Context) (OperationListResult,error)
1. OperationsClient.ListPreparer(context.Context) (*http.Request,error)
1. OperationsClient.ListResponder(*http.Response) (OperationListResult,error)
1. OperationsClient.ListSender(*http.Request) (*http.Response,error)
1. PossibleAdvisorStatusValues() []AdvisorStatus
1. PossibleAutoExecuteStatusValues() []AutoExecuteStatus
1. PossibleBackupLongTermRetentionPolicyStateValues() []BackupLongTermRetentionPolicyState
1. PossibleCapabilityStatusValues() []CapabilityStatus
1. PossibleDataMaskingFunctionValues() []DataMaskingFunction
1. PossibleDataMaskingRuleStateValues() []DataMaskingRuleState
1. PossibleDataMaskingStateValues() []DataMaskingState
1. PossibleDisasterRecoveryConfigurationAutoFailoverValues() []DisasterRecoveryConfigurationAutoFailover
1. PossibleDisasterRecoveryConfigurationFailoverPolicyValues() []DisasterRecoveryConfigurationFailoverPolicy
1. PossibleDisasterRecoveryConfigurationRoleValues() []DisasterRecoveryConfigurationRole
1. PossibleDisasterRecoveryConfigurationStatusValues() []DisasterRecoveryConfigurationStatus
1. PossibleGeoBackupPolicyStateValues() []GeoBackupPolicyState
1. PossibleMaxSizeUnitsValues() []MaxSizeUnits
1. PossiblePerformanceLevelUnitValues() []PerformanceLevelUnit
1. PossiblePrimaryAggregationTypeValues() []PrimaryAggregationType
1. PossibleQueryAggregationFunctionValues() []QueryAggregationFunction
1. PossibleQueryExecutionTypeValues() []QueryExecutionType
1. PossibleQueryMetricUnitValues() []QueryMetricUnit
1. PossibleQueryObservedMetricTypeValues() []QueryObservedMetricType
1. PossibleRestorePointTypeValues() []RestorePointType
1. PossibleServerConnectionTypeValues() []ServerConnectionType
1. PossibleServerStateValues() []ServerState
1. PossibleServerVersionValues() []ServerVersion
1. PossibleUnitDefinitionTypeValues() []UnitDefinitionType
1. PossibleUnitTypeValues() []UnitType
1. QueriesClient.ListByDatabase(context.Context,string,string,string) (TopQueriesListResult,error)
1. QueriesClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. QueriesClient.ListByDatabaseResponder(*http.Response) (TopQueriesListResult,error)
1. QueriesClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. QueryStatisticsClient.ListByQuery(context.Context,string,string,string,string) (QueryStatisticListResult,error)
1. QueryStatisticsClient.ListByQueryPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. QueryStatisticsClient.ListByQueryResponder(*http.Response) (QueryStatisticListResult,error)
1. QueryStatisticsClient.ListByQuerySender(*http.Request) (*http.Response,error)
1. QueryTextsClient.ListByQuery(context.Context,string,string,string,string) (QueryTextListResult,error)
1. QueryTextsClient.ListByQueryPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. QueryTextsClient.ListByQueryResponder(*http.Response) (QueryTextListResult,error)
1. QueryTextsClient.ListByQuerySender(*http.Request) (*http.Response,error)
1. RecoverableDatabase.MarshalJSON() ([]byte,error)
1. RecoverableDatabasesClient.Get(context.Context,string,string,string) (RecoverableDatabase,error)
1. RecoverableDatabasesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. RecoverableDatabasesClient.GetResponder(*http.Response) (RecoverableDatabase,error)
1. RecoverableDatabasesClient.GetSender(*http.Request) (*http.Response,error)
1. RecoverableDatabasesClient.ListByServer(context.Context,string,string) (RecoverableDatabaseListResult,error)
1. RecoverableDatabasesClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. RecoverableDatabasesClient.ListByServerResponder(*http.Response) (RecoverableDatabaseListResult,error)
1. RecoverableDatabasesClient.ListByServerSender(*http.Request) (*http.Response,error)
1. RestorableDroppedDatabase.MarshalJSON() ([]byte,error)
1. RestorableDroppedDatabasesClient.Get(context.Context,string,string,string) (RestorableDroppedDatabase,error)
1. RestorableDroppedDatabasesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. RestorableDroppedDatabasesClient.GetResponder(*http.Response) (RestorableDroppedDatabase,error)
1. RestorableDroppedDatabasesClient.GetSender(*http.Request) (*http.Response,error)
1. RestorableDroppedDatabasesClient.ListByServer(context.Context,string,string) (RestorableDroppedDatabaseListResult,error)
1. RestorableDroppedDatabasesClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. RestorableDroppedDatabasesClient.ListByServerResponder(*http.Response) (RestorableDroppedDatabaseListResult,error)
1. RestorableDroppedDatabasesClient.ListByServerSender(*http.Request) (*http.Response,error)
1. RestorePoint.MarshalJSON() ([]byte,error)
1. RestorePointsClient.ListByDatabase(context.Context,string,string,string) (RestorePointListResult,error)
1. RestorePointsClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. RestorePointsClient.ListByDatabaseResponder(*http.Response) (RestorePointListResult,error)
1. RestorePointsClient.ListByDatabaseSender(*http.Request) (*http.Response,error)
1. Server.MarshalJSON() ([]byte,error)
1. ServerAdvisorsClient.CreateOrUpdate(context.Context,string,string,string,Advisor) (Advisor,error)
1. ServerAdvisorsClient.CreateOrUpdatePreparer(context.Context,string,string,string,Advisor) (*http.Request,error)
1. ServerAdvisorsClient.CreateOrUpdateResponder(*http.Response) (Advisor,error)
1. ServerAdvisorsClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. ServerAdvisorsClient.Get(context.Context,string,string,string) (Advisor,error)
1. ServerAdvisorsClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. ServerAdvisorsClient.GetResponder(*http.Response) (Advisor,error)
1. ServerAdvisorsClient.GetSender(*http.Request) (*http.Response,error)
1. ServerAdvisorsClient.ListByServer(context.Context,string,string) (AdvisorListResult,error)
1. ServerAdvisorsClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. ServerAdvisorsClient.ListByServerResponder(*http.Response) (AdvisorListResult,error)
1. ServerAdvisorsClient.ListByServerSender(*http.Request) (*http.Response,error)
1. ServerAdvisorsClient.Update(context.Context,string,string,string,Advisor) (Advisor,error)
1. ServerAdvisorsClient.UpdatePreparer(context.Context,string,string,string,Advisor) (*http.Request,error)
1. ServerAdvisorsClient.UpdateResponder(*http.Response) (Advisor,error)
1. ServerAdvisorsClient.UpdateSender(*http.Request) (*http.Response,error)
1. ServerAzureADAdministrator.MarshalJSON() ([]byte,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdate(context.Context,string,string,ServerAzureADAdministrator) (ServerAzureADAdministratorsCreateOrUpdateFuture,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdatePreparer(context.Context,string,string,ServerAzureADAdministrator) (*http.Request,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdateResponder(*http.Response) (ServerAzureADAdministrator,error)
1. ServerAzureADAdministratorsClient.CreateOrUpdateSender(*http.Request) (ServerAzureADAdministratorsCreateOrUpdateFuture,error)
1. ServerAzureADAdministratorsClient.Delete(context.Context,string,string) (ServerAzureADAdministratorsDeleteFuture,error)
1. ServerAzureADAdministratorsClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. ServerAzureADAdministratorsClient.DeleteResponder(*http.Response) (ServerAzureADAdministrator,error)
1. ServerAzureADAdministratorsClient.DeleteSender(*http.Request) (ServerAzureADAdministratorsDeleteFuture,error)
1. ServerAzureADAdministratorsClient.Get(context.Context,string,string) (ServerAzureADAdministrator,error)
1. ServerAzureADAdministratorsClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. ServerAzureADAdministratorsClient.GetResponder(*http.Response) (ServerAzureADAdministrator,error)
1. ServerAzureADAdministratorsClient.GetSender(*http.Request) (*http.Response,error)
1. ServerAzureADAdministratorsClient.ListByServer(context.Context,string,string) (ServerAdministratorListResult,error)
1. ServerAzureADAdministratorsClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. ServerAzureADAdministratorsClient.ListByServerResponder(*http.Response) (ServerAdministratorListResult,error)
1. ServerAzureADAdministratorsClient.ListByServerSender(*http.Request) (*http.Response,error)
1. ServerCommunicationLink.MarshalJSON() ([]byte,error)
1. ServerCommunicationLinksClient.CreateOrUpdate(context.Context,string,string,string,ServerCommunicationLink) (ServerCommunicationLinksCreateOrUpdateFuture,error)
1. ServerCommunicationLinksClient.CreateOrUpdatePreparer(context.Context,string,string,string,ServerCommunicationLink) (*http.Request,error)
1. ServerCommunicationLinksClient.CreateOrUpdateResponder(*http.Response) (ServerCommunicationLink,error)
1. ServerCommunicationLinksClient.CreateOrUpdateSender(*http.Request) (ServerCommunicationLinksCreateOrUpdateFuture,error)
1. ServerCommunicationLinksClient.Delete(context.Context,string,string,string) (autorest.Response,error)
1. ServerCommunicationLinksClient.DeletePreparer(context.Context,string,string,string) (*http.Request,error)
1. ServerCommunicationLinksClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. ServerCommunicationLinksClient.DeleteSender(*http.Request) (*http.Response,error)
1. ServerCommunicationLinksClient.Get(context.Context,string,string,string) (ServerCommunicationLink,error)
1. ServerCommunicationLinksClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. ServerCommunicationLinksClient.GetResponder(*http.Response) (ServerCommunicationLink,error)
1. ServerCommunicationLinksClient.GetSender(*http.Request) (*http.Response,error)
1. ServerCommunicationLinksClient.ListByServer(context.Context,string,string) (ServerCommunicationLinkListResult,error)
1. ServerCommunicationLinksClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. ServerCommunicationLinksClient.ListByServerResponder(*http.Response) (ServerCommunicationLinkListResult,error)
1. ServerCommunicationLinksClient.ListByServerSender(*http.Request) (*http.Response,error)
1. ServerConnectionPoliciesClient.CreateOrUpdate(context.Context,string,string,ServerConnectionPolicy) (ServerConnectionPolicy,error)
1. ServerConnectionPoliciesClient.CreateOrUpdatePreparer(context.Context,string,string,ServerConnectionPolicy) (*http.Request,error)
1. ServerConnectionPoliciesClient.CreateOrUpdateResponder(*http.Response) (ServerConnectionPolicy,error)
1. ServerConnectionPoliciesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. ServerConnectionPoliciesClient.Get(context.Context,string,string) (ServerConnectionPolicy,error)
1. ServerConnectionPoliciesClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. ServerConnectionPoliciesClient.GetResponder(*http.Response) (ServerConnectionPolicy,error)
1. ServerConnectionPoliciesClient.GetSender(*http.Request) (*http.Response,error)
1. ServerConnectionPolicy.MarshalJSON() ([]byte,error)
1. ServerTableAuditingPoliciesClient.CreateOrUpdate(context.Context,string,string,ServerTableAuditingPolicy) (ServerTableAuditingPolicy,error)
1. ServerTableAuditingPoliciesClient.CreateOrUpdatePreparer(context.Context,string,string,ServerTableAuditingPolicy) (*http.Request,error)
1. ServerTableAuditingPoliciesClient.CreateOrUpdateResponder(*http.Response) (ServerTableAuditingPolicy,error)
1. ServerTableAuditingPoliciesClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. ServerTableAuditingPoliciesClient.Get(context.Context,string,string) (ServerTableAuditingPolicy,error)
1. ServerTableAuditingPoliciesClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. ServerTableAuditingPoliciesClient.GetResponder(*http.Response) (ServerTableAuditingPolicy,error)
1. ServerTableAuditingPoliciesClient.GetSender(*http.Request) (*http.Response,error)
1. ServerTableAuditingPoliciesClient.ListByServer(context.Context,string,string) (ServerTableAuditingPolicyListResult,error)
1. ServerTableAuditingPoliciesClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. ServerTableAuditingPoliciesClient.ListByServerResponder(*http.Response) (ServerTableAuditingPolicyListResult,error)
1. ServerTableAuditingPoliciesClient.ListByServerSender(*http.Request) (*http.Response,error)
1. ServerTableAuditingPolicy.MarshalJSON() ([]byte,error)
1. ServerUpdate.MarshalJSON() ([]byte,error)
1. ServerUsagesClient.ListByServer(context.Context,string,string) (ServerUsageListResult,error)
1. ServerUsagesClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. ServerUsagesClient.ListByServerResponder(*http.Response) (ServerUsageListResult,error)
1. ServerUsagesClient.ListByServerSender(*http.Request) (*http.Response,error)
1. ServersClient.CreateOrUpdate(context.Context,string,string,Server) (Server,error)
1. ServersClient.CreateOrUpdatePreparer(context.Context,string,string,Server) (*http.Request,error)
1. ServersClient.CreateOrUpdateResponder(*http.Response) (Server,error)
1. ServersClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. ServersClient.Delete(context.Context,string,string) (autorest.Response,error)
1. ServersClient.DeletePreparer(context.Context,string,string) (*http.Request,error)
1. ServersClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. ServersClient.DeleteSender(*http.Request) (*http.Response,error)
1. ServersClient.Get(context.Context,string,string) (Server,error)
1. ServersClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. ServersClient.GetResponder(*http.Response) (Server,error)
1. ServersClient.GetSender(*http.Request) (*http.Response,error)
1. ServersClient.List(context.Context) (ServerListResult,error)
1. ServersClient.ListByResourceGroup(context.Context,string) (ServerListResult,error)
1. ServersClient.ListByResourceGroupPreparer(context.Context,string) (*http.Request,error)
1. ServersClient.ListByResourceGroupResponder(*http.Response) (ServerListResult,error)
1. ServersClient.ListByResourceGroupSender(*http.Request) (*http.Response,error)
1. ServersClient.ListPreparer(context.Context) (*http.Request,error)
1. ServersClient.ListResponder(*http.Response) (ServerListResult,error)
1. ServersClient.ListSender(*http.Request) (*http.Response,error)
1. ServersClient.Update(context.Context,string,string,ServerUpdate) (Server,error)
1. ServersClient.UpdatePreparer(context.Context,string,string,ServerUpdate) (*http.Request,error)
1. ServersClient.UpdateResponder(*http.Response) (Server,error)
1. ServersClient.UpdateSender(*http.Request) (*http.Response,error)
1. ServiceObjective.MarshalJSON() ([]byte,error)
1. ServiceObjectiveCapability.MarshalJSON() ([]byte,error)
1. ServiceObjectivesClient.Get(context.Context,string,string,string) (ServiceObjective,error)
1. ServiceObjectivesClient.GetPreparer(context.Context,string,string,string) (*http.Request,error)
1. ServiceObjectivesClient.GetResponder(*http.Response) (ServiceObjective,error)
1. ServiceObjectivesClient.GetSender(*http.Request) (*http.Response,error)
1. ServiceObjectivesClient.ListByServer(context.Context,string,string) (ServiceObjectiveListResult,error)
1. ServiceObjectivesClient.ListByServerPreparer(context.Context,string,string) (*http.Request,error)
1. ServiceObjectivesClient.ListByServerResponder(*http.Response) (ServiceObjectiveListResult,error)
1. ServiceObjectivesClient.ListByServerSender(*http.Request) (*http.Response,error)
1. TransparentDataEncryptionConfigurationsClient.ListByDatabase(context.Context,string,string,string) (TransparentDataEncryptionListResult,error)
1. TransparentDataEncryptionConfigurationsClient.ListByDatabasePreparer(context.Context,string,string,string) (*http.Request,error)
1. TransparentDataEncryptionConfigurationsClient.ListByDatabaseResponder(*http.Response) (TransparentDataEncryptionListResult,error)
1. TransparentDataEncryptionConfigurationsClient.ListByDatabaseSender(*http.Request) (*http.Response,error)

## Struct Changes

### New Structs

1. Advisor
1. AdvisorListResult
1. AdvisorProperties
1. BackupLongTermRetentionPoliciesClient
1. BackupLongTermRetentionPoliciesCreateOrUpdateFuture
1. BackupLongTermRetentionPolicy
1. BackupLongTermRetentionPolicyListResult
1. BackupLongTermRetentionPolicyProperties
1. BackupLongTermRetentionVault
1. BackupLongTermRetentionVaultListResult
1. BackupLongTermRetentionVaultProperties
1. BackupLongTermRetentionVaultsClient
1. BackupLongTermRetentionVaultsCreateOrUpdateFuture
1. CapabilitiesClient
1. DataMaskingPoliciesClient
1. DataMaskingPolicy
1. DataMaskingPolicyProperties
1. DataMaskingRule
1. DataMaskingRuleListResult
1. DataMaskingRuleProperties
1. DataMaskingRulesClient
1. DatabaseAdvisorsClient
1. DatabaseConnectionPoliciesClient
1. DatabaseConnectionPolicy
1. DatabaseConnectionPolicyProperties
1. DatabaseTableAuditingPoliciesClient
1. DatabaseTableAuditingPolicy
1. DatabaseTableAuditingPolicyListResult
1. DatabaseTableAuditingPolicyProperties
1. DatabaseUsage
1. DatabaseUsageListResult
1. DatabaseUsagesClient
1. DisasterRecoveryConfiguration
1. DisasterRecoveryConfigurationListResult
1. DisasterRecoveryConfigurationProperties
1. DisasterRecoveryConfigurationsClient
1. DisasterRecoveryConfigurationsCreateOrUpdateFuture
1. DisasterRecoveryConfigurationsDeleteFuture
1. DisasterRecoveryConfigurationsFailoverAllowDataLossFuture
1. DisasterRecoveryConfigurationsFailoverFuture
1. EditionCapability
1. ElasticPoolDtuCapability
1. ElasticPoolEditionCapability
1. ElasticPoolPerDatabaseMaxDtuCapability
1. ElasticPoolPerDatabaseMinDtuCapability
1. ExtensionListResult
1. ExtensionsClient
1. GeoBackupPoliciesClient
1. GeoBackupPolicy
1. GeoBackupPolicyListResult
1. GeoBackupPolicyProperties
1. LocationCapabilities
1. MaxSizeCapability
1. Metric
1. MetricAvailability
1. MetricDefinition
1. MetricDefinitionListResult
1. MetricListResult
1. MetricName
1. MetricValue
1. Operation
1. OperationDisplay
1. OperationListResult
1. OperationsClient
1. PerformanceLevel
1. QueriesClient
1. QueryInterval
1. QueryMetric
1. QueryStatistic
1. QueryStatisticListResult
1. QueryStatisticsClient
1. QueryText
1. QueryTextListResult
1. QueryTextsClient
1. RecoverableDatabase
1. RecoverableDatabaseListResult
1. RecoverableDatabaseProperties
1. RecoverableDatabasesClient
1. RestorableDroppedDatabase
1. RestorableDroppedDatabaseListResult
1. RestorableDroppedDatabaseProperties
1. RestorableDroppedDatabasesClient
1. RestorePoint
1. RestorePointListResult
1. RestorePointProperties
1. RestorePointsClient
1. Server
1. ServerAdministratorListResult
1. ServerAdministratorProperties
1. ServerAdvisorsClient
1. ServerAzureADAdministrator
1. ServerAzureADAdministratorsClient
1. ServerAzureADAdministratorsCreateOrUpdateFuture
1. ServerAzureADAdministratorsDeleteFuture
1. ServerCommunicationLink
1. ServerCommunicationLinkListResult
1. ServerCommunicationLinkProperties
1. ServerCommunicationLinksClient
1. ServerCommunicationLinksCreateOrUpdateFuture
1. ServerConnectionPoliciesClient
1. ServerConnectionPolicy
1. ServerConnectionPolicyProperties
1. ServerListResult
1. ServerProperties
1. ServerTableAuditingPoliciesClient
1. ServerTableAuditingPolicy
1. ServerTableAuditingPolicyListResult
1. ServerTableAuditingPolicyProperties
1. ServerUpdate
1. ServerUsage
1. ServerUsageListResult
1. ServerUsagesClient
1. ServerVersionCapability
1. ServiceObjective
1. ServiceObjectiveCapability
1. ServiceObjectiveListResult
1. ServiceObjectiveProperties
1. ServiceObjectivesClient
1. TableAuditingPolicyProperties
1. TopQueries
1. TopQueriesListResult
1. TransparentDataEncryptionConfigurationsClient
1. TransparentDataEncryptionListResult
