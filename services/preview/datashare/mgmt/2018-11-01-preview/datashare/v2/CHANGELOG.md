## Breaking Changes

## Struct Changes

### Removed Struct Fields

1. AccountProperties.CreatedBy
1. ConsumerInvitationProperties.Sender
1. ConsumerInvitationProperties.SenderCompanyName
1. InvitationProperties.Sender
1. ProviderShareSubscriptionProperties.Company
1. ProviderShareSubscriptionProperties.CreatedBy
1. ProviderShareSubscriptionProperties.SharedBy
1. ScheduledSynchronizationSettingProperties.CreatedBy
1. ScheduledTriggerProperties.CreatedBy
1. ShareProperties.CreatedBy
1. ShareSubscriptionProperties.CreatedBy
1. ShareSubscriptionProperties.ShareSender
1. ShareSubscriptionProperties.ShareSenderCompanyName
1. ShareSynchronization.Company
1. ShareSynchronization.Recipient

## Signature Changes

### Funcs

1. DataSetsClient.Delete
	- Returns
		- From: autorest.Response,error
		- To: DataSetsDeleteFuture,error
1. DataSetsClient.DeleteSender
	- Returns
		- From: *http.Response,error
		- To: DataSetsDeleteFuture,error

## New Content

### New Constants

1. DataSetType.KustoCluster
1. DataSetType.KustoDatabase
1. Kind.KindKustoCluster
1. Kind.KindKustoDatabase
1. KindBasicDataSetMapping.KindBasicDataSetMappingKindKustoCluster
1. KindBasicDataSetMapping.KindBasicDataSetMappingKindKustoDatabase
1. ShareKind.InPlace

### New Funcs

1. *DataSetsDeleteFuture.Result(DataSetsClient) (autorest.Response,error)
1. *KustoClusterDataSet.UnmarshalJSON([]byte) error
1. *KustoClusterDataSetMapping.UnmarshalJSON([]byte) error
1. *KustoDatabaseDataSet.UnmarshalJSON([]byte) error
1. *KustoDatabaseDataSetMapping.UnmarshalJSON([]byte) error
1. ADLSGen1FileDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. ADLSGen1FileDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. ADLSGen1FolderDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. ADLSGen1FolderDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. ADLSGen2FileDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. ADLSGen2FileDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. ADLSGen2FileDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. ADLSGen2FileDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. ADLSGen2FileSystemDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. ADLSGen2FileSystemDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. ADLSGen2FileSystemDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. ADLSGen2FileSystemDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. ADLSGen2FolderDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. ADLSGen2FolderDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. ADLSGen2FolderDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. ADLSGen2FolderDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. BlobContainerDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. BlobContainerDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. BlobContainerDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. BlobContainerDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. BlobDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. BlobDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. BlobDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. BlobDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. BlobFolderDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. BlobFolderDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. BlobFolderDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. BlobFolderDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. DataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. DataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. DataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. DataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. KustoClusterDataSet.AsADLSGen1FileDataSet() (*ADLSGen1FileDataSet,bool)
1. KustoClusterDataSet.AsADLSGen1FolderDataSet() (*ADLSGen1FolderDataSet,bool)
1. KustoClusterDataSet.AsADLSGen2FileDataSet() (*ADLSGen2FileDataSet,bool)
1. KustoClusterDataSet.AsADLSGen2FileSystemDataSet() (*ADLSGen2FileSystemDataSet,bool)
1. KustoClusterDataSet.AsADLSGen2FolderDataSet() (*ADLSGen2FolderDataSet,bool)
1. KustoClusterDataSet.AsBasicDataSet() (BasicDataSet,bool)
1. KustoClusterDataSet.AsBlobContainerDataSet() (*BlobContainerDataSet,bool)
1. KustoClusterDataSet.AsBlobDataSet() (*BlobDataSet,bool)
1. KustoClusterDataSet.AsBlobFolderDataSet() (*BlobFolderDataSet,bool)
1. KustoClusterDataSet.AsDataSet() (*DataSet,bool)
1. KustoClusterDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. KustoClusterDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. KustoClusterDataSet.AsSQLDBTableDataSet() (*SQLDBTableDataSet,bool)
1. KustoClusterDataSet.AsSQLDWTableDataSet() (*SQLDWTableDataSet,bool)
1. KustoClusterDataSet.MarshalJSON() ([]byte,error)
1. KustoClusterDataSetMapping.AsADLSGen2FileDataSetMapping() (*ADLSGen2FileDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsADLSGen2FileSystemDataSetMapping() (*ADLSGen2FileSystemDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsADLSGen2FolderDataSetMapping() (*ADLSGen2FolderDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsBasicDataSetMapping() (BasicDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsBlobContainerDataSetMapping() (*BlobContainerDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsBlobDataSetMapping() (*BlobDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsBlobFolderDataSetMapping() (*BlobFolderDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsDataSetMapping() (*DataSetMapping,bool)
1. KustoClusterDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsSQLDBTableDataSetMapping() (*SQLDBTableDataSetMapping,bool)
1. KustoClusterDataSetMapping.AsSQLDWTableDataSetMapping() (*SQLDWTableDataSetMapping,bool)
1. KustoClusterDataSetMapping.MarshalJSON() ([]byte,error)
1. KustoDatabaseDataSet.AsADLSGen1FileDataSet() (*ADLSGen1FileDataSet,bool)
1. KustoDatabaseDataSet.AsADLSGen1FolderDataSet() (*ADLSGen1FolderDataSet,bool)
1. KustoDatabaseDataSet.AsADLSGen2FileDataSet() (*ADLSGen2FileDataSet,bool)
1. KustoDatabaseDataSet.AsADLSGen2FileSystemDataSet() (*ADLSGen2FileSystemDataSet,bool)
1. KustoDatabaseDataSet.AsADLSGen2FolderDataSet() (*ADLSGen2FolderDataSet,bool)
1. KustoDatabaseDataSet.AsBasicDataSet() (BasicDataSet,bool)
1. KustoDatabaseDataSet.AsBlobContainerDataSet() (*BlobContainerDataSet,bool)
1. KustoDatabaseDataSet.AsBlobDataSet() (*BlobDataSet,bool)
1. KustoDatabaseDataSet.AsBlobFolderDataSet() (*BlobFolderDataSet,bool)
1. KustoDatabaseDataSet.AsDataSet() (*DataSet,bool)
1. KustoDatabaseDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. KustoDatabaseDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. KustoDatabaseDataSet.AsSQLDBTableDataSet() (*SQLDBTableDataSet,bool)
1. KustoDatabaseDataSet.AsSQLDWTableDataSet() (*SQLDWTableDataSet,bool)
1. KustoDatabaseDataSet.MarshalJSON() ([]byte,error)
1. KustoDatabaseDataSetMapping.AsADLSGen2FileDataSetMapping() (*ADLSGen2FileDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsADLSGen2FileSystemDataSetMapping() (*ADLSGen2FileSystemDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsADLSGen2FolderDataSetMapping() (*ADLSGen2FolderDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsBasicDataSetMapping() (BasicDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsBlobContainerDataSetMapping() (*BlobContainerDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsBlobDataSetMapping() (*BlobDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsBlobFolderDataSetMapping() (*BlobFolderDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsDataSetMapping() (*DataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsSQLDBTableDataSetMapping() (*SQLDBTableDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.AsSQLDWTableDataSetMapping() (*SQLDWTableDataSetMapping,bool)
1. KustoDatabaseDataSetMapping.MarshalJSON() ([]byte,error)
1. SQLDBTableDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. SQLDBTableDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. SQLDBTableDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. SQLDBTableDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)
1. SQLDWTableDataSet.AsKustoClusterDataSet() (*KustoClusterDataSet,bool)
1. SQLDWTableDataSet.AsKustoDatabaseDataSet() (*KustoDatabaseDataSet,bool)
1. SQLDWTableDataSetMapping.AsKustoClusterDataSetMapping() (*KustoClusterDataSetMapping,bool)
1. SQLDWTableDataSetMapping.AsKustoDatabaseDataSetMapping() (*KustoDatabaseDataSetMapping,bool)

## Struct Changes

### New Structs

1. DataSetsDeleteFuture
1. KustoClusterDataSet
1. KustoClusterDataSetMapping
1. KustoClusterDataSetMappingProperties
1. KustoClusterDataSetProperties
1. KustoDatabaseDataSet
1. KustoDatabaseDataSetMapping
1. KustoDatabaseDataSetMappingProperties
1. KustoDatabaseDataSetProperties

### New Struct Fields

1. ADLSGen2FileDataSetMappingProperties.ProvisioningState
1. ADLSGen2FileSystemDataSetMappingProperties.ProvisioningState
1. ADLSGen2FolderDataSetMappingProperties.ProvisioningState
1. AccountProperties.UserEmail
1. AccountProperties.UserName
1. BlobContainerMappingProperties.ProvisioningState
1. BlobFolderMappingProperties.ProvisioningState
1. BlobMappingProperties.ProvisioningState
1. ConsumerInvitationProperties.ProviderEmail
1. ConsumerInvitationProperties.ProviderName
1. ConsumerInvitationProperties.ProviderTenantName
1. ConsumerInvitationProperties.UserEmail
1. ConsumerInvitationProperties.UserName
1. ConsumerSourceDataSetProperties.DataSetLocation
1. ConsumerSourceDataSetProperties.DataSetPath
1. InvitationProperties.UserEmail
1. InvitationProperties.UserName
1. ProviderShareSubscriptionProperties.ConsumerEmail
1. ProviderShareSubscriptionProperties.ConsumerName
1. ProviderShareSubscriptionProperties.ConsumerTenantName
1. ProviderShareSubscriptionProperties.ProviderName
1. SQLDBTableDataSetMappingProperties.ProvisioningState
1. SQLDBTableDataSetMappingProperties.SchemaName
1. SQLDBTableProperties.SchemaName
1. SQLDWTableDataSetMappingProperties.ProvisioningState
1. SQLDWTableDataSetMappingProperties.SchemaName
1. SQLDWTableProperties.SchemaName
1. ScheduledSynchronizationSettingProperties.UserName
1. ScheduledTriggerProperties.UserName
1. ShareProperties.UserEmail
1. ShareProperties.UserName
1. ShareSubscriptionProperties.ProviderEmail
1. ShareSubscriptionProperties.ProviderName
1. ShareSubscriptionProperties.ProviderTenantName
1. ShareSubscriptionProperties.UserEmail
1. ShareSubscriptionProperties.UserName
1. ShareSubscriptionSynchronization.SynchronizationMode
1. ShareSynchronization.ConsumerEmail
1. ShareSynchronization.ConsumerName
1. ShareSynchronization.ConsumerTenantName
1. ShareSynchronization.SynchronizationMode
