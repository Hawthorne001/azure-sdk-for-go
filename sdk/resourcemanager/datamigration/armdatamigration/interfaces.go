// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

// CommandPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetCommandProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CommandProperties, *MigrateMISyncCompleteCommandProperties, *MigrateSyncCompleteCommandProperties, *MongoDbCancelCommand,
// - *MongoDbFinishCommand, *MongoDbRestartCommand
type CommandPropertiesClassification interface {
	// GetCommandProperties returns the CommandProperties content of the underlying type.
	GetCommandProperties() *CommandProperties
}

// ConnectToSourceSQLServerTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetConnectToSourceSQLServerTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ConnectToSourceSQLServerTaskOutput, *ConnectToSourceSQLServerTaskOutputAgentJobLevel, *ConnectToSourceSQLServerTaskOutputDatabaseLevel,
// - *ConnectToSourceSQLServerTaskOutputLoginLevel, *ConnectToSourceSQLServerTaskOutputTaskLevel
type ConnectToSourceSQLServerTaskOutputClassification interface {
	// GetConnectToSourceSQLServerTaskOutput returns the ConnectToSourceSQLServerTaskOutput content of the underlying type.
	GetConnectToSourceSQLServerTaskOutput() *ConnectToSourceSQLServerTaskOutput
}

// ConnectionInfoClassification provides polymorphic access to related types.
// Call the interface's GetConnectionInfo() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *ConnectionInfo, *MiSQLConnectionInfo, *MongoDbConnectionInfo, *MySQLConnectionInfo, *OracleConnectionInfo, *PostgreSQLConnectionInfo,
// - *SQLConnectionInfo
type ConnectionInfoClassification interface {
	// GetConnectionInfo returns the ConnectionInfo content of the underlying type.
	GetConnectionInfo() *ConnectionInfo
}

// DatabaseMigrationBasePropertiesClassification provides polymorphic access to related types.
// Call the interface's GetDatabaseMigrationBaseProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DatabaseMigrationBaseProperties, *DatabaseMigrationProperties, *DatabaseMigrationPropertiesCosmosDbMongo, *DatabaseMigrationPropertiesSQLDb,
// - *DatabaseMigrationPropertiesSQLMi, *DatabaseMigrationPropertiesSQLVM
type DatabaseMigrationBasePropertiesClassification interface {
	// GetDatabaseMigrationBaseProperties returns the DatabaseMigrationBaseProperties content of the underlying type.
	GetDatabaseMigrationBaseProperties() *DatabaseMigrationBaseProperties
}

// DatabaseMigrationPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetDatabaseMigrationProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DatabaseMigrationProperties, *DatabaseMigrationPropertiesSQLDb, *DatabaseMigrationPropertiesSQLMi, *DatabaseMigrationPropertiesSQLVM
type DatabaseMigrationPropertiesClassification interface {
	DatabaseMigrationBasePropertiesClassification
	// GetDatabaseMigrationProperties returns the DatabaseMigrationProperties content of the underlying type.
	GetDatabaseMigrationProperties() *DatabaseMigrationProperties
}

// MigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateMySQLAzureDbForMySQLOfflineTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateMySQLAzureDbForMySQLOfflineTaskOutput, *MigrateMySQLAzureDbForMySQLOfflineTaskOutputDatabaseLevel, *MigrateMySQLAzureDbForMySQLOfflineTaskOutputError,
// - *MigrateMySQLAzureDbForMySQLOfflineTaskOutputMigrationLevel, *MigrateMySQLAzureDbForMySQLOfflineTaskOutputTableLevel
type MigrateMySQLAzureDbForMySQLOfflineTaskOutputClassification interface {
	// GetMigrateMySQLAzureDbForMySQLOfflineTaskOutput returns the MigrateMySQLAzureDbForMySQLOfflineTaskOutput content of the underlying type.
	GetMigrateMySQLAzureDbForMySQLOfflineTaskOutput() *MigrateMySQLAzureDbForMySQLOfflineTaskOutput
}

// MigrateMySQLAzureDbForMySQLSyncTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateMySQLAzureDbForMySQLSyncTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateMySQLAzureDbForMySQLSyncTaskOutput, *MigrateMySQLAzureDbForMySQLSyncTaskOutputDatabaseError, *MigrateMySQLAzureDbForMySQLSyncTaskOutputDatabaseLevel,
// - *MigrateMySQLAzureDbForMySQLSyncTaskOutputError, *MigrateMySQLAzureDbForMySQLSyncTaskOutputMigrationLevel, *MigrateMySQLAzureDbForMySQLSyncTaskOutputTableLevel
type MigrateMySQLAzureDbForMySQLSyncTaskOutputClassification interface {
	// GetMigrateMySQLAzureDbForMySQLSyncTaskOutput returns the MigrateMySQLAzureDbForMySQLSyncTaskOutput content of the underlying type.
	GetMigrateMySQLAzureDbForMySQLSyncTaskOutput() *MigrateMySQLAzureDbForMySQLSyncTaskOutput
}

// MigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateOracleAzureDbPostgreSQLSyncTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateOracleAzureDbPostgreSQLSyncTaskOutput, *MigrateOracleAzureDbPostgreSQLSyncTaskOutputDatabaseError, *MigrateOracleAzureDbPostgreSQLSyncTaskOutputDatabaseLevel,
// - *MigrateOracleAzureDbPostgreSQLSyncTaskOutputError, *MigrateOracleAzureDbPostgreSQLSyncTaskOutputMigrationLevel, *MigrateOracleAzureDbPostgreSQLSyncTaskOutputTableLevel
type MigrateOracleAzureDbPostgreSQLSyncTaskOutputClassification interface {
	// GetMigrateOracleAzureDbPostgreSQLSyncTaskOutput returns the MigrateOracleAzureDbPostgreSQLSyncTaskOutput content of the underlying type.
	GetMigrateOracleAzureDbPostgreSQLSyncTaskOutput() *MigrateOracleAzureDbPostgreSQLSyncTaskOutput
}

// MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutput, *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputDatabaseError,
// - *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputDatabaseLevel, *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputError,
// - *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputMigrationLevel, *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputTableLevel
type MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutputClassification interface {
	// GetMigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutput returns the MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutput content of the underlying type.
	GetMigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutput() *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskOutput
}

// MigrateSQLServerSQLDbSyncTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateSQLServerSQLDbSyncTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateSQLServerSQLDbSyncTaskOutput, *MigrateSQLServerSQLDbSyncTaskOutputDatabaseError, *MigrateSQLServerSQLDbSyncTaskOutputDatabaseLevel,
// - *MigrateSQLServerSQLDbSyncTaskOutputError, *MigrateSQLServerSQLDbSyncTaskOutputMigrationLevel, *MigrateSQLServerSQLDbSyncTaskOutputTableLevel
type MigrateSQLServerSQLDbSyncTaskOutputClassification interface {
	// GetMigrateSQLServerSQLDbSyncTaskOutput returns the MigrateSQLServerSQLDbSyncTaskOutput content of the underlying type.
	GetMigrateSQLServerSQLDbSyncTaskOutput() *MigrateSQLServerSQLDbSyncTaskOutput
}

// MigrateSQLServerSQLDbTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateSQLServerSQLDbTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateSQLServerSQLDbTaskOutput, *MigrateSQLServerSQLDbTaskOutputDatabaseLevel, *MigrateSQLServerSQLDbTaskOutputDatabaseLevelValidationResult,
// - *MigrateSQLServerSQLDbTaskOutputError, *MigrateSQLServerSQLDbTaskOutputMigrationLevel, *MigrateSQLServerSQLDbTaskOutputTableLevel,
// - *MigrateSQLServerSQLDbTaskOutputValidationResult
type MigrateSQLServerSQLDbTaskOutputClassification interface {
	// GetMigrateSQLServerSQLDbTaskOutput returns the MigrateSQLServerSQLDbTaskOutput content of the underlying type.
	GetMigrateSQLServerSQLDbTaskOutput() *MigrateSQLServerSQLDbTaskOutput
}

// MigrateSQLServerSQLMISyncTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateSQLServerSQLMISyncTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateSQLServerSQLMISyncTaskOutput, *MigrateSQLServerSQLMISyncTaskOutputDatabaseLevel, *MigrateSQLServerSQLMISyncTaskOutputError,
// - *MigrateSQLServerSQLMISyncTaskOutputMigrationLevel
type MigrateSQLServerSQLMISyncTaskOutputClassification interface {
	// GetMigrateSQLServerSQLMISyncTaskOutput returns the MigrateSQLServerSQLMISyncTaskOutput content of the underlying type.
	GetMigrateSQLServerSQLMISyncTaskOutput() *MigrateSQLServerSQLMISyncTaskOutput
}

// MigrateSQLServerSQLMITaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateSQLServerSQLMITaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateSQLServerSQLMITaskOutput, *MigrateSQLServerSQLMITaskOutputAgentJobLevel, *MigrateSQLServerSQLMITaskOutputDatabaseLevel,
// - *MigrateSQLServerSQLMITaskOutputError, *MigrateSQLServerSQLMITaskOutputLoginLevel, *MigrateSQLServerSQLMITaskOutputMigrationLevel
type MigrateSQLServerSQLMITaskOutputClassification interface {
	// GetMigrateSQLServerSQLMITaskOutput returns the MigrateSQLServerSQLMITaskOutput content of the underlying type.
	GetMigrateSQLServerSQLMITaskOutput() *MigrateSQLServerSQLMITaskOutput
}

// MigrateSchemaSQLServerSQLDbTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateSchemaSQLServerSQLDbTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateSchemaSQLServerSQLDbTaskOutput, *MigrateSchemaSQLServerSQLDbTaskOutputDatabaseLevel, *MigrateSchemaSQLServerSQLDbTaskOutputError,
// - *MigrateSchemaSQLServerSQLDbTaskOutputMigrationLevel, *MigrateSchemaSQLTaskOutputError
type MigrateSchemaSQLServerSQLDbTaskOutputClassification interface {
	// GetMigrateSchemaSQLServerSQLDbTaskOutput returns the MigrateSchemaSQLServerSQLDbTaskOutput content of the underlying type.
	GetMigrateSchemaSQLServerSQLDbTaskOutput() *MigrateSchemaSQLServerSQLDbTaskOutput
}

// MigrateSsisTaskOutputClassification provides polymorphic access to related types.
// Call the interface's GetMigrateSsisTaskOutput() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MigrateSsisTaskOutput, *MigrateSsisTaskOutputMigrationLevel, *MigrateSsisTaskOutputProjectLevel
type MigrateSsisTaskOutputClassification interface {
	// GetMigrateSsisTaskOutput returns the MigrateSsisTaskOutput content of the underlying type.
	GetMigrateSsisTaskOutput() *MigrateSsisTaskOutput
}

// MongoDbProgressClassification provides polymorphic access to related types.
// Call the interface's GetMongoDbProgress() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MongoDbCollectionProgress, *MongoDbDatabaseProgress, *MongoDbMigrationProgress, *MongoDbProgress
type MongoDbProgressClassification interface {
	// GetMongoDbProgress returns the MongoDbProgress content of the underlying type.
	GetMongoDbProgress() *MongoDbProgress
}

// ProjectTaskPropertiesClassification provides polymorphic access to related types.
// Call the interface's GetProjectTaskProperties() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *CheckOCIDriverTaskProperties, *ConnectToMongoDbTaskProperties, *ConnectToSourceMySQLTaskProperties, *ConnectToSourceOracleSyncTaskProperties,
// - *ConnectToSourcePostgreSQLSyncTaskProperties, *ConnectToSourceSQLServerSyncTaskProperties, *ConnectToSourceSQLServerTaskProperties,
// - *ConnectToTargetAzureDbForMySQLTaskProperties, *ConnectToTargetAzureDbForPostgreSQLSyncTaskProperties, *ConnectToTargetOracleAzureDbForPostgreSQLSyncTaskProperties,
// - *ConnectToTargetSQLDbSyncTaskProperties, *ConnectToTargetSQLDbTaskProperties, *ConnectToTargetSQLMISyncTaskProperties,
// - *ConnectToTargetSQLMITaskProperties, *GetTdeCertificatesSQLTaskProperties, *GetUserTablesMySQLTaskProperties, *GetUserTablesOracleTaskProperties,
// - *GetUserTablesPostgreSQLTaskProperties, *GetUserTablesSQLSyncTaskProperties, *GetUserTablesSQLTaskProperties, *InstallOCIDriverTaskProperties,
// - *MigrateMongoDbTaskProperties, *MigrateMySQLAzureDbForMySQLOfflineTaskProperties, *MigrateMySQLAzureDbForMySQLSyncTaskProperties,
// - *MigrateOracleAzureDbForPostgreSQLSyncTaskProperties, *MigratePostgreSQLAzureDbForPostgreSQLSyncTaskProperties, *MigrateSQLServerSQLDbSyncTaskProperties,
// - *MigrateSQLServerSQLDbTaskProperties, *MigrateSQLServerSQLMISyncTaskProperties, *MigrateSQLServerSQLMITaskProperties,
// - *MigrateSchemaSQLServerSQLDbTaskProperties, *MigrateSsisTaskProperties, *ProjectTaskProperties, *UploadOCIDriverTaskProperties,
// - *ValidateMigrationInputSQLServerSQLDbSyncTaskProperties, *ValidateMigrationInputSQLServerSQLMISyncTaskProperties, *ValidateMigrationInputSQLServerSQLMITaskProperties,
// - *ValidateMongoDbTaskProperties, *ValidateOracleAzureDbForPostgreSQLSyncTaskProperties
type ProjectTaskPropertiesClassification interface {
	// GetProjectTaskProperties returns the ProjectTaskProperties content of the underlying type.
	GetProjectTaskProperties() *ProjectTaskProperties
}
