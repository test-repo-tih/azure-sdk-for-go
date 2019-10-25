## Breaking Changes

### Removed Constants

1. TaskType.TaskTypeMigratePostgreSQLAzureDbForPostgreSQLSync

### Removed Funcs

1. CheckOCIDriverTaskOutput.MarshalJSON() ([]byte,error)

## Signature Changes

### Struct Fields

1. CheckOCIDriverTaskOutput.InstalledDriver changed type from map[string][]OracleOCIDriverInfo to *OracleOCIDriverInfo

## New Content

### New Constants

1. ReplicateMigrationState.ReplicateMigrationStateACTIONREQUIRED
1. ReplicateMigrationState.ReplicateMigrationStateCOMPLETE
1. ReplicateMigrationState.ReplicateMigrationStateFAILED
1. ReplicateMigrationState.ReplicateMigrationStatePENDING
1. ReplicateMigrationState.ReplicateMigrationStateUNDEFINED
1. ReplicateMigrationState.ReplicateMigrationStateVALIDATING
1. ScenarioSource.Access
1. ScenarioSource.DB2
1. ScenarioSource.MongoDB
1. ScenarioSource.MySQL
1. ScenarioSource.MySQLRDS
1. ScenarioSource.Oracle
1. ScenarioSource.PostgreSQL
1. ScenarioSource.PostgreSQLRDS
1. ScenarioSource.SQL
1. ScenarioSource.SQLRDS
1. ScenarioSource.Sybase
1. ScenarioTarget.ScenarioTargetAzureDBForMySQL
1. ScenarioTarget.ScenarioTargetAzureDBForPostgreeSQL
1. ScenarioTarget.ScenarioTargetMongoDB
1. ScenarioTarget.ScenarioTargetSQLDB
1. ScenarioTarget.ScenarioTargetSQLDW
1. ScenarioTarget.ScenarioTargetSQLMI
1. ScenarioTarget.ScenarioTargetSQLServer
1. SyncDatabaseMigrationReportingState.SyncDatabaseMigrationReportingStateBACKUPCOMPLETED
1. SyncDatabaseMigrationReportingState.SyncDatabaseMigrationReportingStateBACKUPINPROGRESS
1. SyncDatabaseMigrationReportingState.SyncDatabaseMigrationReportingStateRESTORECOMPLETED
1. SyncDatabaseMigrationReportingState.SyncDatabaseMigrationReportingStateRESTOREINPROGRESS
1. SyncDatabaseMigrationReportingState.SyncDatabaseMigrationReportingStateVALIDATING
1. SyncDatabaseMigrationReportingState.SyncDatabaseMigrationReportingStateVALIDATIONCOMPLETE
1. SyncDatabaseMigrationReportingState.SyncDatabaseMigrationReportingStateVALIDATIONFAILED
1. TaskType.TaskTypeMigratePostgreSQLAzureDbForPostgreSQLSyncV2

### New Funcs

1. PossibleReplicateMigrationStateValues() []ReplicateMigrationState
1. PossibleScenarioSourceValues() []ScenarioSource
1. PossibleScenarioTargetValues() []ScenarioTarget

## Struct Changes

### New Struct Fields

1. MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputMigrationLevel.SourceServerType
1. MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputMigrationLevel.State
1. MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputMigrationLevel.TargetServerType
