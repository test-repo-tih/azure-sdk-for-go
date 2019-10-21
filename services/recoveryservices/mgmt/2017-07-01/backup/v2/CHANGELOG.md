## Breaking Changes

### Removed Constants

1. ProtectableItemType.ProtectableItemTypeSAPAseDatabase

### Removed Funcs

1. AzureFileShareProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureIaaSClassicComputeVMProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureIaaSComputeVMProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureFileShareProtectableItem() (*AzureFileShareProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureIaaSClassicComputeVMProtectableItem() (*AzureIaaSClassicComputeVMProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureIaaSComputeVMProtectableItem() (*AzureIaaSComputeVMProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadProtectableItem() (*AzureVMWorkloadProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadSAPAseSystemProtectableItem() (*AzureVMWorkloadSAPAseSystemProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadSAPHanaDatabaseProtectableItem() (*AzureVMWorkloadSAPHanaDatabaseProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadSAPHanaSystemProtectableItem() (*AzureVMWorkloadSAPHanaSystemProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadSQLAvailabilityGroupProtectableItem() (*AzureVMWorkloadSQLAvailabilityGroupProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadSQLDatabaseProtectableItem() (*AzureVMWorkloadSQLDatabaseProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsAzureVMWorkloadSQLInstanceProtectableItem() (*AzureVMWorkloadSQLInstanceProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsBasicAzureVMWorkloadProtectableItem() (BasicAzureVMWorkloadProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsBasicIaaSVMProtectableItem() (BasicIaaSVMProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsBasicWorkloadProtectableItem() (BasicWorkloadProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsIaaSVMProtectableItem() (*IaaSVMProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.AsWorkloadProtectableItem() (*WorkloadProtectableItem,bool)
1. AzureVMWorkloadSAPAseDatabaseProtectableItem.MarshalJSON() ([]byte,error)
1. AzureVMWorkloadSAPAseSystemProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadSAPHanaDatabaseProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadSAPHanaSystemProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadSQLAvailabilityGroupProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadSQLDatabaseProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. AzureVMWorkloadSQLInstanceProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. IaaSVMProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)
1. ProtectedItemsClient.List(context.Context,string,string,string,string) (ProtectedItemResourceListPage,error)
1. ProtectedItemsClient.ListComplete(context.Context,string,string,string,string) (ProtectedItemResourceListIterator,error)
1. ProtectedItemsClient.ListPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. ProtectedItemsClient.ListResponder(*http.Response) (ProtectedItemResourceList,error)
1. ProtectedItemsClient.ListSender(*http.Request) (*http.Response,error)
1. ProtectedItemsGroupClient.CreateOrUpdate(context.Context,string,string,string,string,string,ProtectedItemResource) (ProtectedItemResource,error)
1. ProtectedItemsGroupClient.CreateOrUpdatePreparer(context.Context,string,string,string,string,string,ProtectedItemResource) (*http.Request,error)
1. ProtectedItemsGroupClient.CreateOrUpdateResponder(*http.Response) (ProtectedItemResource,error)
1. ProtectedItemsGroupClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. ProtectedItemsGroupClient.Delete(context.Context,string,string,string,string,string) (autorest.Response,error)
1. ProtectedItemsGroupClient.DeletePreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. ProtectedItemsGroupClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. ProtectedItemsGroupClient.DeleteSender(*http.Request) (*http.Response,error)
1. ProtectedItemsGroupClient.Get(context.Context,string,string,string,string,string,string) (ProtectedItemResource,error)
1. ProtectedItemsGroupClient.GetPreparer(context.Context,string,string,string,string,string,string) (*http.Request,error)
1. ProtectedItemsGroupClient.GetResponder(*http.Response) (ProtectedItemResource,error)
1. ProtectedItemsGroupClient.GetSender(*http.Request) (*http.Response,error)
1. WorkloadProtectableItem.AsAzureVMWorkloadSAPAseDatabaseProtectableItem() (*AzureVMWorkloadSAPAseDatabaseProtectableItem,bool)

## Struct Changes

### Removed Structs

1. AzureVMWorkloadSAPAseDatabaseProtectableItem

## New Content

### New Constants

1. JobOperationType.JobOperationTypeCrossRegionRestore
1. JobOperationType.JobOperationTypeUndelete
1. ObjectTypeBasicILRRequest.ObjectTypeAzureFileShareProvisionILRRequest
1. ObjectTypeBasicOperationStatusExtendedInfo.ObjectTypeOperationStatusRecoveryPointExtendedInfo
1. RecoveryMode.RecoveryModeFileRecovery
1. RecoveryMode.RecoveryModeInvalid
1. RecoveryMode.RecoveryModeWorkloadRecovery
1. RecoveryType.RecoveryTypeOffline
1. SoftDeleteFeatureState.SoftDeleteFeatureStateDisabled
1. SoftDeleteFeatureState.SoftDeleteFeatureStateEnabled
1. SoftDeleteFeatureState.SoftDeleteFeatureStateInvalid

### New Funcs

1. *CrossRegionRestoreRequest.UnmarshalJSON([]byte) error
1. *CrrOperationResultsGetFuture.Result(CrrOperationResultsClient) (autorest.Response,error)
1. *OperationStatusRecoveryPointExtendedInfo.UnmarshalJSON([]byte) error
1. AADPropertiesResource.MarshalJSON() ([]byte,error)
1. AadPropertiesClient.Get(context.Context,string) (AADPropertiesResource,error)
1. AadPropertiesClient.GetPreparer(context.Context,string) (*http.Request,error)
1. AadPropertiesClient.GetResponder(*http.Response) (AADPropertiesResource,error)
1. AadPropertiesClient.GetSender(*http.Request) (*http.Response,error)
1. AzureFileShareProvisionILRRequest.AsAzureFileShareProvisionILRRequest() (*AzureFileShareProvisionILRRequest,bool)
1. AzureFileShareProvisionILRRequest.AsBasicILRRequest() (BasicILRRequest,bool)
1. AzureFileShareProvisionILRRequest.AsILRRequest() (*ILRRequest,bool)
1. AzureFileShareProvisionILRRequest.AsIaasVMILRRegistrationRequest() (*IaasVMILRRegistrationRequest,bool)
1. AzureFileShareProvisionILRRequest.MarshalJSON() ([]byte,error)
1. CrossRegionRestoreClient.Trigger(context.Context,string,CrossRegionRestoreRequestResource) (autorest.Response,error)
1. CrossRegionRestoreClient.TriggerPreparer(context.Context,string,CrossRegionRestoreRequestResource) (*http.Request,error)
1. CrossRegionRestoreClient.TriggerResponder(*http.Response) (autorest.Response,error)
1. CrossRegionRestoreClient.TriggerSender(*http.Request) (*http.Response,error)
1. CrossRegionRestoreRequestResource.MarshalJSON() ([]byte,error)
1. CrrAccessTokenResource.MarshalJSON() ([]byte,error)
1. CrrJobDetailsClient.Get(context.Context,string) (JobResource,error)
1. CrrJobDetailsClient.GetPreparer(context.Context,string) (*http.Request,error)
1. CrrJobDetailsClient.GetResponder(*http.Response) (JobResource,error)
1. CrrJobDetailsClient.GetSender(*http.Request) (*http.Response,error)
1. CrrJobsClient.List(context.Context,string) (JobResourceListPage,error)
1. CrrJobsClient.ListComplete(context.Context,string) (JobResourceListIterator,error)
1. CrrJobsClient.ListPreparer(context.Context,string) (*http.Request,error)
1. CrrJobsClient.ListResponder(*http.Response) (JobResourceList,error)
1. CrrJobsClient.ListSender(*http.Request) (*http.Response,error)
1. CrrOperationResultsClient.Get(context.Context,string,string) (CrrOperationResultsGetFuture,error)
1. CrrOperationResultsClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. CrrOperationResultsClient.GetResponder(*http.Response) (autorest.Response,error)
1. CrrOperationResultsClient.GetSender(*http.Request) (CrrOperationResultsGetFuture,error)
1. CrrOperationStatusClient.Get(context.Context,string,string) (OperationStatus,error)
1. CrrOperationStatusClient.GetPreparer(context.Context,string,string) (*http.Request,error)
1. CrrOperationStatusClient.GetResponder(*http.Response) (OperationStatus,error)
1. CrrOperationStatusClient.GetSender(*http.Request) (*http.Response,error)
1. ILRRequest.AsAzureFileShareProvisionILRRequest() (*AzureFileShareProvisionILRRequest,bool)
1. IaasVMILRRegistrationRequest.AsAzureFileShareProvisionILRRequest() (*AzureFileShareProvisionILRRequest,bool)
1. NewAadPropertiesClient(string) AadPropertiesClient
1. NewAadPropertiesClientWithBaseURI(string,string) AadPropertiesClient
1. NewCrossRegionRestoreClient(string) CrossRegionRestoreClient
1. NewCrossRegionRestoreClientWithBaseURI(string,string) CrossRegionRestoreClient
1. NewCrrJobDetailsClient(string) CrrJobDetailsClient
1. NewCrrJobDetailsClientWithBaseURI(string,string) CrrJobDetailsClient
1. NewCrrJobsClient(string) CrrJobsClient
1. NewCrrJobsClientWithBaseURI(string,string) CrrJobsClient
1. NewCrrOperationResultsClient(string) CrrOperationResultsClient
1. NewCrrOperationResultsClientWithBaseURI(string,string) CrrOperationResultsClient
1. NewCrrOperationStatusClient(string) CrrOperationStatusClient
1. NewCrrOperationStatusClientWithBaseURI(string,string) CrrOperationStatusClient
1. NewProtectedItemsCrrClient(string) ProtectedItemsCrrClient
1. NewProtectedItemsCrrClientWithBaseURI(string,string) ProtectedItemsCrrClient
1. NewRecoveryPointsCrrClient(string) RecoveryPointsCrrClient
1. NewRecoveryPointsCrrClientWithBaseURI(string,string) RecoveryPointsCrrClient
1. OperationStatusExtendedInfo.AsOperationStatusRecoveryPointExtendedInfo() (*OperationStatusRecoveryPointExtendedInfo,bool)
1. OperationStatusJobExtendedInfo.AsOperationStatusRecoveryPointExtendedInfo() (*OperationStatusRecoveryPointExtendedInfo,bool)
1. OperationStatusJobsExtendedInfo.AsOperationStatusRecoveryPointExtendedInfo() (*OperationStatusRecoveryPointExtendedInfo,bool)
1. OperationStatusProvisionILRExtendedInfo.AsOperationStatusRecoveryPointExtendedInfo() (*OperationStatusRecoveryPointExtendedInfo,bool)
1. OperationStatusRecoveryPointExtendedInfo.AsBasicOperationStatusExtendedInfo() (BasicOperationStatusExtendedInfo,bool)
1. OperationStatusRecoveryPointExtendedInfo.AsOperationStatusExtendedInfo() (*OperationStatusExtendedInfo,bool)
1. OperationStatusRecoveryPointExtendedInfo.AsOperationStatusJobExtendedInfo() (*OperationStatusJobExtendedInfo,bool)
1. OperationStatusRecoveryPointExtendedInfo.AsOperationStatusJobsExtendedInfo() (*OperationStatusJobsExtendedInfo,bool)
1. OperationStatusRecoveryPointExtendedInfo.AsOperationStatusProvisionILRExtendedInfo() (*OperationStatusProvisionILRExtendedInfo,bool)
1. OperationStatusRecoveryPointExtendedInfo.AsOperationStatusRecoveryPointExtendedInfo() (*OperationStatusRecoveryPointExtendedInfo,bool)
1. OperationStatusRecoveryPointExtendedInfo.MarshalJSON() ([]byte,error)
1. PossibleRecoveryModeValues() []RecoveryMode
1. PossibleSoftDeleteFeatureStateValues() []SoftDeleteFeatureState
1. ProtectedItemsClient.CreateOrUpdate(context.Context,string,string,string,string,string,ProtectedItemResource) (ProtectedItemResource,error)
1. ProtectedItemsClient.CreateOrUpdatePreparer(context.Context,string,string,string,string,string,ProtectedItemResource) (*http.Request,error)
1. ProtectedItemsClient.CreateOrUpdateResponder(*http.Response) (ProtectedItemResource,error)
1. ProtectedItemsClient.CreateOrUpdateSender(*http.Request) (*http.Response,error)
1. ProtectedItemsClient.Delete(context.Context,string,string,string,string,string) (autorest.Response,error)
1. ProtectedItemsClient.DeletePreparer(context.Context,string,string,string,string,string) (*http.Request,error)
1. ProtectedItemsClient.DeleteResponder(*http.Response) (autorest.Response,error)
1. ProtectedItemsClient.DeleteSender(*http.Request) (*http.Response,error)
1. ProtectedItemsClient.Get(context.Context,string,string,string,string,string,string) (ProtectedItemResource,error)
1. ProtectedItemsClient.GetPreparer(context.Context,string,string,string,string,string,string) (*http.Request,error)
1. ProtectedItemsClient.GetResponder(*http.Response) (ProtectedItemResource,error)
1. ProtectedItemsClient.GetSender(*http.Request) (*http.Response,error)
1. ProtectedItemsCrrClient.List(context.Context,string,string,string,string) (ProtectedItemResourceListPage,error)
1. ProtectedItemsCrrClient.ListComplete(context.Context,string,string,string,string) (ProtectedItemResourceListIterator,error)
1. ProtectedItemsCrrClient.ListPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. ProtectedItemsCrrClient.ListResponder(*http.Response) (ProtectedItemResourceList,error)
1. ProtectedItemsCrrClient.ListSender(*http.Request) (*http.Response,error)
1. ProtectedItemsGroupClient.List(context.Context,string,string,string,string) (ProtectedItemResourceListPage,error)
1. ProtectedItemsGroupClient.ListComplete(context.Context,string,string,string,string) (ProtectedItemResourceListIterator,error)
1. ProtectedItemsGroupClient.ListPreparer(context.Context,string,string,string,string) (*http.Request,error)
1. ProtectedItemsGroupClient.ListResponder(*http.Response) (ProtectedItemResourceList,error)
1. ProtectedItemsGroupClient.ListSender(*http.Request) (*http.Response,error)
1. RecoveryPointsClient.GetAccessToken(context.Context,string,string,string,string,string,string) (CrrAccessTokenResource,error)
1. RecoveryPointsClient.GetAccessTokenPreparer(context.Context,string,string,string,string,string,string) (*http.Request,error)
1. RecoveryPointsClient.GetAccessTokenResponder(*http.Response) (CrrAccessTokenResource,error)
1. RecoveryPointsClient.GetAccessTokenSender(*http.Request) (*http.Response,error)
1. RecoveryPointsCrrClient.List(context.Context,string,string,string,string,string,string) (RecoveryPointResourceListPage,error)
1. RecoveryPointsCrrClient.ListComplete(context.Context,string,string,string,string,string,string) (RecoveryPointResourceListIterator,error)
1. RecoveryPointsCrrClient.ListPreparer(context.Context,string,string,string,string,string,string) (*http.Request,error)
1. RecoveryPointsCrrClient.ListResponder(*http.Response) (RecoveryPointResourceList,error)
1. RecoveryPointsCrrClient.ListSender(*http.Request) (*http.Response,error)
1. TargetRestoreInfo.MarshalJSON() ([]byte,error)

## Struct Changes

### New Structs

1. AADProperties
1. AADPropertiesResource
1. AadPropertiesClient
1. AzureFileShareProvisionILRRequest
1. CrossRegionRestoreClient
1. CrossRegionRestoreRequest
1. CrossRegionRestoreRequestResource
1. CrrAccessToken
1. CrrAccessTokenResource
1. CrrJobDetailsClient
1. CrrJobsClient
1. CrrOperationResultsClient
1. CrrOperationResultsGetFuture
1. CrrOperationStatusClient
1. DiskExclusionProperties
1. DiskInformation
1. ExtendedProperties
1. InstantRPAdditionalDetails
1. OperationStatusRecoveryPointExtendedInfo
1. ProtectedItemsCrrClient
1. RecoveryPointDiskConfiguration
1. RecoveryPointsCrrClient

### New Struct Fields

1. AzureFileShareRecoveryPoint.RecoveryPointSizeInGB
1. AzureFileshareProtectedItemExtendedInfo.ResourceState
1. AzureFileshareProtectedItemExtendedInfo.ResourceStateSyncTime
1. AzureIaaSClassicComputeVMProtectedItem.ExtendedProperties
1. AzureIaaSComputeVMProtectedItem.ExtendedProperties
1. AzureIaaSVMProtectedItem.ExtendedProperties
1. AzureVMWorkloadProtectionPolicy.MakePolicyConsistent
1. AzureWorkloadPointInTimeRestoreRequest.RecoveryMode
1. AzureWorkloadRestoreRequest.RecoveryMode
1. AzureWorkloadSAPHanaPointInTimeRestoreRequest.RecoveryMode
1. AzureWorkloadSAPHanaRestoreRequest.RecoveryMode
1. AzureWorkloadSQLPointInTimeRestoreRequest.RecoveryMode
1. AzureWorkloadSQLRestoreRequest.RecoveryMode
1. ExportJobsOperationResultInfo.ExcelFileBlobSasKey
1. ExportJobsOperationResultInfo.ExcelFileBlobURL
1. IaasVMRecoveryPoint.RecoveryPointDiskConfiguration
1. IaasVMRestoreRequest.RestoreDiskLunList
1. InquiryValidation.AdditionalDetail
1. MabFileFolderProtectedItem.LastBackupTime
1. ResourceVaultConfig.SoftDeleteFeatureState
1. TargetRestoreInfo.TargetDirectoryMapping
