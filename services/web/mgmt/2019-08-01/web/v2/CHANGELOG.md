## Breaking Changes

### Removed Constants

1. ManagedServiceIdentityType.ManagedServiceIdentityTypeSystemAssignedUserAssigned

### Removed Funcs

1. *MetricDefinition.UnmarshalJSON([]byte) error
1. *ResourceMetricCollectionIterator.Next() error
1. *ResourceMetricCollectionIterator.NextWithContext(context.Context) error
1. *ResourceMetricCollectionPage.Next() error
1. *ResourceMetricCollectionPage.NextWithContext(context.Context) error
1. AppServiceEnvironmentsClient.ListMetricDefinitions(context.Context,string,string) (MetricDefinition,error)
1. AppServiceEnvironmentsClient.ListMetricDefinitionsPreparer(context.Context,string,string) (*http.Request,error)
1. AppServiceEnvironmentsClient.ListMetricDefinitionsResponder(*http.Response) (MetricDefinition,error)
1. AppServiceEnvironmentsClient.ListMetricDefinitionsSender(*http.Request) (*http.Response,error)
1. AppServiceEnvironmentsClient.ListMetrics(context.Context,string,string,*bool,string) (ResourceMetricCollectionPage,error)
1. AppServiceEnvironmentsClient.ListMetricsComplete(context.Context,string,string,*bool,string) (ResourceMetricCollectionIterator,error)
1. AppServiceEnvironmentsClient.ListMetricsPreparer(context.Context,string,string,*bool,string) (*http.Request,error)
1. AppServiceEnvironmentsClient.ListMetricsResponder(*http.Response) (ResourceMetricCollection,error)
1. AppServiceEnvironmentsClient.ListMetricsSender(*http.Request) (*http.Response,error)
1. AppServiceEnvironmentsClient.ListMultiRoleMetrics(context.Context,string,string,string,string,string,*bool,string) (ResourceMetricCollectionPage,error)
1. AppServiceEnvironmentsClient.ListMultiRoleMetricsComplete(context.Context,string,string,string,string,string,*bool,string) (ResourceMetricCollectionIterator,error)
1. AppServiceEnvironmentsClient.ListMultiRoleMetricsPreparer(context.Context,string,string,string,string,string,*bool,string) (*http.Request,error)
1. AppServiceEnvironmentsClient.ListMultiRoleMetricsResponder(*http.Response) (ResourceMetricCollection,error)
1. AppServiceEnvironmentsClient.ListMultiRoleMetricsSender(*http.Request) (*http.Response,error)
1. AppServiceEnvironmentsClient.ListMultiRolePoolInstanceMetrics(context.Context,string,string,string,*bool) (ResourceMetricCollectionPage,error)
1. AppServiceEnvironmentsClient.ListMultiRolePoolInstanceMetricsComplete(context.Context,string,string,string,*bool) (ResourceMetricCollectionIterator,error)
1. AppServiceEnvironmentsClient.ListMultiRolePoolInstanceMetricsPreparer(context.Context,string,string,string,*bool) (*http.Request,error)
1. AppServiceEnvironmentsClient.ListMultiRolePoolInstanceMetricsResponder(*http.Response) (ResourceMetricCollection,error)
1. AppServiceEnvironmentsClient.ListMultiRolePoolInstanceMetricsSender(*http.Request) (*http.Response,error)
1. AppServiceEnvironmentsClient.ListVips(context.Context,string,string) (AddressResponse,error)
1. AppServiceEnvironmentsClient.ListVipsPreparer(context.Context,string,string) (*http.Request,error)
1. AppServiceEnvironmentsClient.ListVipsResponder(*http.Response) (AddressResponse,error)
1. AppServiceEnvironmentsClient.ListVipsSender(*http.Request) (*http.Response,error)
1. AppServiceEnvironmentsClient.ListWebWorkerMetrics(context.Context,string,string,string,*bool,string) (ResourceMetricCollectionPage,error)
1. AppServiceEnvironmentsClient.ListWebWorkerMetricsComplete(context.Context,string,string,string,*bool,string) (ResourceMetricCollectionIterator,error)
1. AppServiceEnvironmentsClient.ListWebWorkerMetricsPreparer(context.Context,string,string,string,*bool,string) (*http.Request,error)
1. AppServiceEnvironmentsClient.ListWebWorkerMetricsResponder(*http.Response) (ResourceMetricCollection,error)
1. AppServiceEnvironmentsClient.ListWebWorkerMetricsSender(*http.Request) (*http.Response,error)
1. AppServiceEnvironmentsClient.ListWorkerPoolInstanceMetrics(context.Context,string,string,string,string,*bool,string) (ResourceMetricCollectionPage,error)
1. AppServiceEnvironmentsClient.ListWorkerPoolInstanceMetricsComplete(context.Context,string,string,string,string,*bool,string) (ResourceMetricCollectionIterator,error)
1. AppServiceEnvironmentsClient.ListWorkerPoolInstanceMetricsPreparer(context.Context,string,string,string,string,*bool,string) (*http.Request,error)
1. AppServiceEnvironmentsClient.ListWorkerPoolInstanceMetricsResponder(*http.Response) (ResourceMetricCollection,error)
1. AppServiceEnvironmentsClient.ListWorkerPoolInstanceMetricsSender(*http.Request) (*http.Response,error)
1. AppServicePlansClient.ListMetricDefintions(context.Context,string,string) (ResourceMetricDefinitionCollectionPage,error)
1. AppServicePlansClient.ListMetricDefintionsComplete(context.Context,string,string) (ResourceMetricDefinitionCollectionIterator,error)
1. AppServicePlansClient.ListMetricDefintionsPreparer(context.Context,string,string) (*http.Request,error)
1. AppServicePlansClient.ListMetricDefintionsResponder(*http.Response) (ResourceMetricDefinitionCollection,error)
1. AppServicePlansClient.ListMetricDefintionsSender(*http.Request) (*http.Response,error)
1. AppServicePlansClient.ListMetrics(context.Context,string,string,*bool,string) (ResourceMetricCollectionPage,error)
1. AppServicePlansClient.ListMetricsComplete(context.Context,string,string,*bool,string) (ResourceMetricCollectionIterator,error)
1. AppServicePlansClient.ListMetricsPreparer(context.Context,string,string,*bool,string) (*http.Request,error)
1. AppServicePlansClient.ListMetricsResponder(*http.Response) (ResourceMetricCollection,error)
1. AppServicePlansClient.ListMetricsSender(*http.Request) (*http.Response,error)
1. AppsClient.GetInstanceProcessThread(context.Context,string,string,string,string,string) (ProcessThreadInfo,error)
1. AppsClient.GetInstanceProcessThreadPreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. AppsClient.GetInstanceProcessThreadResponder(*http.Response) (ProcessThreadInfo,error)
1. AppsClient.GetInstanceProcessThreadSender(*http.Request) (*http.Response,error)
1. AppsClient.GetInstanceProcessThreadSlot(context.Context,string,string,string,string,string,string) (ProcessThreadInfo,error)
1. AppsClient.GetInstanceProcessThreadSlotPreparer(context.Context,string,string,string,string,string,string) (*http.Request,error)
1. AppsClient.GetInstanceProcessThreadSlotResponder(*http.Response) (ProcessThreadInfo,error)
1. AppsClient.GetInstanceProcessThreadSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.GetProcessThread(context.Context,string,string,string,string) (ProcessThreadInfo,error)
1. AppsClient.GetProcessThreadPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. AppsClient.GetProcessThreadResponder(*http.Response) (ProcessThreadInfo,error)
1. AppsClient.GetProcessThreadSender(*http.Request) (*http.Response,error)
1. AppsClient.GetProcessThreadSlot(context.Context,string,string,string,string,string) (ProcessThreadInfo,error)
1. AppsClient.GetProcessThreadSlotPreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. AppsClient.GetProcessThreadSlotResponder(*http.Response) (ProcessThreadInfo,error)
1. AppsClient.GetProcessThreadSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.ListHybridConnectionKeys(context.Context,string,string,string,string) (HybridConnectionKey,error)
1. AppsClient.ListHybridConnectionKeysPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. AppsClient.ListHybridConnectionKeysResponder(*http.Response) (HybridConnectionKey,error)
1. AppsClient.ListHybridConnectionKeysSender(*http.Request) (*http.Response,error)
1. AppsClient.ListHybridConnectionKeysSlot(context.Context,string,string,string,string,string) (HybridConnectionKey,error)
1. AppsClient.ListHybridConnectionKeysSlotPreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. AppsClient.ListHybridConnectionKeysSlotResponder(*http.Response) (HybridConnectionKey,error)
1. AppsClient.ListHybridConnectionKeysSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.ListMetricDefinitions(context.Context,string,string) (ResourceMetricDefinitionCollectionPage,error)
1. AppsClient.ListMetricDefinitionsComplete(context.Context,string,string) (ResourceMetricDefinitionCollectionIterator,error)
1. AppsClient.ListMetricDefinitionsPreparer(context.Context,string,string) (*http.Request,error)
1. AppsClient.ListMetricDefinitionsResponder(*http.Response) (ResourceMetricDefinitionCollection,error)
1. AppsClient.ListMetricDefinitionsSender(*http.Request) (*http.Response,error)
1. AppsClient.ListMetricDefinitionsSlot(context.Context,string,string,string) (ResourceMetricDefinitionCollectionPage,error)
1. AppsClient.ListMetricDefinitionsSlotComplete(context.Context,string,string,string) (ResourceMetricDefinitionCollectionIterator,error)
1. AppsClient.ListMetricDefinitionsSlotPreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.ListMetricDefinitionsSlotResponder(*http.Response) (ResourceMetricDefinitionCollection,error)
1. AppsClient.ListMetricDefinitionsSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.ListMetrics(context.Context,string,string,*bool,string) (ResourceMetricCollectionPage,error)
1. AppsClient.ListMetricsComplete(context.Context,string,string,*bool,string) (ResourceMetricCollectionIterator,error)
1. AppsClient.ListMetricsPreparer(context.Context,string,string,*bool,string) (*http.Request,error)
1. AppsClient.ListMetricsResponder(*http.Response) (ResourceMetricCollection,error)
1. AppsClient.ListMetricsSender(*http.Request) (*http.Response,error)
1. AppsClient.ListMetricsSlot(context.Context,string,string,string,*bool,string) (ResourceMetricCollectionPage,error)
1. AppsClient.ListMetricsSlotComplete(context.Context,string,string,string,*bool,string) (ResourceMetricCollectionIterator,error)
1. AppsClient.ListMetricsSlotPreparer(context.Context,string,string,string,*bool,string) (*http.Request,error)
1. AppsClient.ListMetricsSlotResponder(*http.Response) (ResourceMetricCollection,error)
1. AppsClient.ListMetricsSlotSender(*http.Request) (*http.Response,error)
1. BaseClient.ValidateContainerSettings(context.Context,ValidateContainerSettingsRequest,string) (SetObject,error)
1. BaseClient.ValidateContainerSettingsPreparer(context.Context,ValidateContainerSettingsRequest,string) (*http.Request,error)
1. BaseClient.ValidateContainerSettingsResponder(*http.Response) (SetObject,error)
1. BaseClient.ValidateContainerSettingsSender(*http.Request) (*http.Response,error)
1. DiagnosticsClient.GetSiteDetectorComplete(context.Context,string,string,string,string) (DiagnosticDetectorCollectionIterator,error)
1. DiagnosticsClient.GetSiteDetectorSlotComplete(context.Context,string,string,string,string,string) (DiagnosticDetectorCollectionIterator,error)
1. MetricDefinition.MarshalJSON() ([]byte,error)
1. NewResourceMetricCollectionIterator(ResourceMetricCollectionPage) ResourceMetricCollectionIterator
1. NewResourceMetricCollectionPage(func(context.Context, ResourceMetricCollection) (ResourceMetricCollection, error)) ResourceMetricCollectionPage
1. ResourceMetricCollection.IsEmpty() bool
1. ResourceMetricCollectionIterator.NotDone() bool
1. ResourceMetricCollectionIterator.Response() ResourceMetricCollection
1. ResourceMetricCollectionIterator.Value() ResourceMetric
1. ResourceMetricCollectionPage.NotDone() bool
1. ResourceMetricCollectionPage.Response() ResourceMetricCollection
1. ResourceMetricCollectionPage.Values() []ResourceMetric
1. SiteConfig.MarshalJSON() ([]byte,error)

## Struct Changes

### Removed Structs

1. DomainAvailablilityCheckResult
1. GeoDistribution
1. MetricAvailabilily
1. MetricDefinition
1. MetricDefinitionProperties
1. ResourceMetric
1. ResourceMetricCollection
1. ResourceMetricCollectionIterator
1. ResourceMetricCollectionPage
1. ResourceMetricName
1. ResourceMetricProperty
1. ResourceMetricValue
1. ValidateContainerSettingsRequest

### Removed Struct Fields

1. AddressResponse.InternalIPAddress
1. AddressResponse.OutboundIPAddresses
1. AddressResponse.ServiceIPAddress
1. AddressResponse.VipMappings
1. EndpointDetail.IsAccessable
1. HostingEnvironmentDiagnostics.DiagnosicsOutput
1. IdentifierProperties.ID
1. ProcessThreadInfo.autorest.Response
1. ProcessThreadInfoProperties.PriviledgedProcessorTime
1. SiteConfig.AzureStorageAccounts
1. SiteConfig.ReservedInstanceCount
1. SitePatchResourceProperties.GeoDistributions
1. SiteProperties.GeoDistributions

## Signature Changes

### Funcs

1. ApplicationStackCollectionIterator.Value
	- Returns
		- From: ApplicationStack
		- To: ApplicationStackResource
1. ApplicationStackCollectionPage.Values
	- Returns
		- From: []ApplicationStack
		- To: []ApplicationStackResource
1. DiagnosticsClient.GetSiteAnalysis
	- Returns
		- From: DiagnosticAnalysis,error
		- To: AnalysisDefinition,error
1. DiagnosticsClient.GetSiteAnalysisResponder
	- Returns
		- From: DiagnosticAnalysis,error
		- To: AnalysisDefinition,error
1. DiagnosticsClient.GetSiteAnalysisSlot
	- Returns
		- From: DiagnosticAnalysis,error
		- To: AnalysisDefinition,error
1. DiagnosticsClient.GetSiteAnalysisSlotResponder
	- Returns
		- From: DiagnosticAnalysis,error
		- To: AnalysisDefinition,error
1. DiagnosticsClient.GetSiteDetector
	- Returns
		- From: DiagnosticDetectorCollectionPage,error
		- To: DetectorDefinition,error
1. DiagnosticsClient.GetSiteDetectorResponder
	- Returns
		- From: DiagnosticDetectorCollection,error
		- To: DetectorDefinition,error
1. DiagnosticsClient.GetSiteDetectorSlot
	- Returns
		- From: DiagnosticDetectorCollectionPage,error
		- To: DetectorDefinition,error
1. DiagnosticsClient.GetSiteDetectorSlotResponder
	- Returns
		- From: DiagnosticDetectorCollection,error
		- To: DetectorDefinition,error
1. DomainsClient.CheckAvailability
	- Returns
		- From: DomainAvailablilityCheckResult,error
		- To: DomainAvailabilityCheckResult,error
1. DomainsClient.CheckAvailabilityResponder
	- Returns
		- From: DomainAvailablilityCheckResult,error
		- To: DomainAvailabilityCheckResult,error

### Struct Fields

1. ApplicationStackCollection.Value changed type from *[]ApplicationStack to *[]ApplicationStackResource

## New Content

### New Constants

1. ConfigReferenceLocation.ApplicationSetting
1. ConfigReferenceSource.KeyVault
1. ResolveStatus.AccessToKeyVaultDenied
1. ResolveStatus.Initialized
1. ResolveStatus.InvalidSyntax
1. ResolveStatus.MSINotEnabled
1. ResolveStatus.OtherReasons
1. ResolveStatus.Resolved
1. ResolveStatus.SecretNotFound
1. ResolveStatus.SecretVersionNotFound
1. ResolveStatus.VaultNotFound
1. ScmType.ScmTypeVSTSRM
1. SiteRuntimeState.READY
1. SiteRuntimeState.STOPPED
1. SiteRuntimeState.UNKNOWN
1. WorkerSizeOptions.WorkerSizeOptionsNestedSmall

### New Funcs

1. *AddressResponse.UnmarshalJSON([]byte) error
1. *ApplicationStackResource.UnmarshalJSON([]byte) error
1. *AppsCopyProductionSlotFuture.Result(AppsClient) (autorest.Response,error)
1. *AppsCopySlotSlotFuture.Result(AppsClient) (autorest.Response,error)
1. *KeyVaultReferenceCollection.UnmarshalJSON([]byte) error
1. *KeyVaultReferenceResource.UnmarshalJSON([]byte) error
1. *SiteInstanceStatus.UnmarshalJSON([]byte) error
1. AddressResponse.MarshalJSON() ([]byte,error)
1. AppServiceEnvironmentsClient.GetVipInfo(context.Context,string,string) (AddressResponse,error)
1. AppServiceEnvironmentsClient.GetVipInfoPreparer(context.Context,string,string) (*http.Request,error)
1. AppServiceEnvironmentsClient.GetVipInfoResponder(*http.Response) (AddressResponse,error)
1. AppServiceEnvironmentsClient.GetVipInfoSender(*http.Request) (*http.Response,error)
1. ApplicationStackResource.MarshalJSON() ([]byte,error)
1. AppsClient.CopyProductionSlot(context.Context,string,string,CsmCopySlotEntity) (AppsCopyProductionSlotFuture,error)
1. AppsClient.CopyProductionSlotPreparer(context.Context,string,string,CsmCopySlotEntity) (*http.Request,error)
1. AppsClient.CopyProductionSlotResponder(*http.Response) (autorest.Response,error)
1. AppsClient.CopyProductionSlotSender(*http.Request) (AppsCopyProductionSlotFuture,error)
1. AppsClient.CopySlotSlot(context.Context,string,string,CsmCopySlotEntity,string) (AppsCopySlotSlotFuture,error)
1. AppsClient.CopySlotSlotPreparer(context.Context,string,string,CsmCopySlotEntity,string) (*http.Request,error)
1. AppsClient.CopySlotSlotResponder(*http.Response) (autorest.Response,error)
1. AppsClient.CopySlotSlotSender(*http.Request) (AppsCopySlotSlotFuture,error)
1. AppsClient.GetAppSettingKeyVaultReference(context.Context,string,string,string) (KeyVaultReferenceResource,error)
1. AppsClient.GetAppSettingKeyVaultReferencePreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.GetAppSettingKeyVaultReferenceResponder(*http.Response) (KeyVaultReferenceResource,error)
1. AppsClient.GetAppSettingKeyVaultReferenceSender(*http.Request) (*http.Response,error)
1. AppsClient.GetAppSettingsKeyVaultReferences(context.Context,string,string) (KeyVaultReferenceCollection,error)
1. AppsClient.GetAppSettingsKeyVaultReferencesPreparer(context.Context,string,string) (*http.Request,error)
1. AppsClient.GetAppSettingsKeyVaultReferencesResponder(*http.Response) (KeyVaultReferenceCollection,error)
1. AppsClient.GetAppSettingsKeyVaultReferencesSender(*http.Request) (*http.Response,error)
1. AppsClient.GetInstanceInfo(context.Context,string,string,string) (SiteInstanceStatus,error)
1. AppsClient.GetInstanceInfoPreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.GetInstanceInfoResponder(*http.Response) (SiteInstanceStatus,error)
1. AppsClient.GetInstanceInfoSender(*http.Request) (*http.Response,error)
1. AppsClient.GetInstanceInfoSlot(context.Context,string,string,string,string) (SiteInstanceStatus,error)
1. AppsClient.GetInstanceInfoSlotPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. AppsClient.GetInstanceInfoSlotResponder(*http.Response) (SiteInstanceStatus,error)
1. AppsClient.GetInstanceInfoSlotSender(*http.Request) (*http.Response,error)
1. AppsClient.ListSiteBackups(context.Context,string,string) (BackupItemCollectionPage,error)
1. AppsClient.ListSiteBackupsComplete(context.Context,string,string) (BackupItemCollectionIterator,error)
1. AppsClient.ListSiteBackupsPreparer(context.Context,string,string) (*http.Request,error)
1. AppsClient.ListSiteBackupsResponder(*http.Response) (BackupItemCollection,error)
1. AppsClient.ListSiteBackupsSender(*http.Request) (*http.Response,error)
1. AppsClient.ListSiteBackupsSlot(context.Context,string,string,string) (BackupItemCollectionPage,error)
1. AppsClient.ListSiteBackupsSlotComplete(context.Context,string,string,string) (BackupItemCollectionIterator,error)
1. AppsClient.ListSiteBackupsSlotPreparer(context.Context,string,string,string) (*http.Request,error)
1. AppsClient.ListSiteBackupsSlotResponder(*http.Response) (BackupItemCollection,error)
1. AppsClient.ListSiteBackupsSlotSender(*http.Request) (*http.Response,error)
1. KeyVaultReferenceCollection.MarshalJSON() ([]byte,error)
1. KeyVaultReferenceCollectionProperties.MarshalJSON() ([]byte,error)
1. KeyVaultReferenceResource.MarshalJSON() ([]byte,error)
1. PossibleConfigReferenceLocationValues() []ConfigReferenceLocation
1. PossibleConfigReferenceSourceValues() []ConfigReferenceSource
1. PossibleResolveStatusValues() []ResolveStatus
1. PossibleSiteRuntimeStateValues() []SiteRuntimeState
1. SiteInstanceStatus.MarshalJSON() ([]byte,error)
1. SiteInstanceStatusProperties.MarshalJSON() ([]byte,error)

## Struct Changes

### New Structs

1. APIKVReference
1. APIManagementConfig
1. AddressResponseProperties
1. ApplicationStackResource
1. AppsCopyProductionSlotFuture
1. AppsCopySlotSlotFuture
1. ContainerCPUStatistics
1. ContainerCPUUsage
1. ContainerInfo
1. ContainerMemoryStatistics
1. ContainerNetworkInterfaceStatistics
1. ContainerThrottlingData
1. CsmCopySlotEntity
1. DomainAvailabilityCheckResult
1. KeyVaultReferenceCollection
1. KeyVaultReferenceCollectionProperties
1. KeyVaultReferenceResource
1. SiteInstanceStatus
1. SiteInstanceStatusProperties

### New Struct Fields

1. AddressResponse.*AddressResponseProperties
1. AddressResponse.ID
1. AddressResponse.Kind
1. AddressResponse.Name
1. AddressResponse.Type
1. AnalysisDefinition.autorest.Response
1. CertificatePatchResourceProperties.CanonicalName
1. CertificateProperties.CanonicalName
1. DetectorDefinition.autorest.Response
1. EndpointDetail.IsAccessible
1. FunctionEnvelopeProperties.InvokeURLTemplate
1. FunctionEnvelopeProperties.IsDisabled
1. FunctionEnvelopeProperties.Language
1. FunctionEnvelopeProperties.TestDataHref
1. GeoRegionProperties.OrgDomain
1. HostingEnvironmentDiagnostics.DiagnosticsOutput
1. IdentifierProperties.Value
1. MetricSpecification.SupportedTimeGrainTypes
1. SiteConfig.APIManagementConfig
1. SiteConfig.HealthCheckPath
1. SiteConfig.PreWarmedInstanceCount
1. StackMajorVersion.IsDeprecated
1. StackMajorVersion.IsHidden
1. StackMajorVersion.IsPreview
1. StackMinorVersion.TestOnly
1. ValidateProperties.ContainerImagePlatform
1. ValidateProperties.ContainerImageRepository
1. ValidateProperties.ContainerImageTag
1. ValidateProperties.ContainerRegistryBaseURL
1. ValidateProperties.ContainerRegistryPassword
1. ValidateProperties.ContainerRegistryUsername
1. VirtualIPMapping.ServiceName
