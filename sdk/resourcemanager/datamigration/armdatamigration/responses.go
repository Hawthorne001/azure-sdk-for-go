// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

// DatabaseMigrationsMongoToCosmosDbRUMongoClientCreateResponse contains the response from method DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginCreate.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientCreateResponse struct {
	// Database Migration Resource for Mongo to CosmosDb.
	DatabaseMigrationCosmosDbMongo
}

// DatabaseMigrationsMongoToCosmosDbRUMongoClientDeleteResponse contains the response from method DatabaseMigrationsMongoToCosmosDbRUMongoClient.BeginDelete.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeResponse contains the response from method DatabaseMigrationsMongoToCosmosDbRUMongoClient.NewGetForScopePager.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientGetForScopeResponse struct {
	// A list of Database Migrations.
	DatabaseMigrationCosmosDbMongoListResult
}

// DatabaseMigrationsMongoToCosmosDbRUMongoClientGetResponse contains the response from method DatabaseMigrationsMongoToCosmosDbRUMongoClient.Get.
type DatabaseMigrationsMongoToCosmosDbRUMongoClientGetResponse struct {
	// Database Migration Resource for Mongo to CosmosDb.
	DatabaseMigrationCosmosDbMongo
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientCreateResponse contains the response from method DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginCreate.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientCreateResponse struct {
	// Database Migration Resource for Mongo to CosmosDb.
	DatabaseMigrationCosmosDbMongo
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientDeleteResponse contains the response from method DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.BeginDelete.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse contains the response from method DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.NewGetForScopePager.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetForScopeResponse struct {
	// A list of Database Migrations.
	DatabaseMigrationCosmosDbMongoListResult
}

// DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse contains the response from method DatabaseMigrationsMongoToCosmosDbvCoreMongoClient.Get.
type DatabaseMigrationsMongoToCosmosDbvCoreMongoClientGetResponse struct {
	// Database Migration Resource for Mongo to CosmosDb.
	DatabaseMigrationCosmosDbMongo
}

// DatabaseMigrationsSQLDbClientCancelResponse contains the response from method DatabaseMigrationsSQLDbClient.BeginCancel.
type DatabaseMigrationsSQLDbClientCancelResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsSQLDbClientCreateOrUpdateResponse contains the response from method DatabaseMigrationsSQLDbClient.BeginCreateOrUpdate.
type DatabaseMigrationsSQLDbClientCreateOrUpdateResponse struct {
	// Database Migration Resource for SQL Database.
	DatabaseMigrationSQLDb
}

// DatabaseMigrationsSQLDbClientDeleteResponse contains the response from method DatabaseMigrationsSQLDbClient.BeginDelete.
type DatabaseMigrationsSQLDbClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsSQLDbClientGetResponse contains the response from method DatabaseMigrationsSQLDbClient.Get.
type DatabaseMigrationsSQLDbClientGetResponse struct {
	// Database Migration Resource for SQL Database.
	DatabaseMigrationSQLDb
}

// DatabaseMigrationsSQLMiClientCancelResponse contains the response from method DatabaseMigrationsSQLMiClient.BeginCancel.
type DatabaseMigrationsSQLMiClientCancelResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsSQLMiClientCreateOrUpdateResponse contains the response from method DatabaseMigrationsSQLMiClient.BeginCreateOrUpdate.
type DatabaseMigrationsSQLMiClientCreateOrUpdateResponse struct {
	// Database Migration Resource for SQL Managed Instance.
	DatabaseMigrationSQLMi
}

// DatabaseMigrationsSQLMiClientCutoverResponse contains the response from method DatabaseMigrationsSQLMiClient.BeginCutover.
type DatabaseMigrationsSQLMiClientCutoverResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsSQLMiClientGetResponse contains the response from method DatabaseMigrationsSQLMiClient.Get.
type DatabaseMigrationsSQLMiClientGetResponse struct {
	// Database Migration Resource for SQL Managed Instance.
	DatabaseMigrationSQLMi
}

// DatabaseMigrationsSQLVMClientCancelResponse contains the response from method DatabaseMigrationsSQLVMClient.BeginCancel.
type DatabaseMigrationsSQLVMClientCancelResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsSQLVMClientCreateOrUpdateResponse contains the response from method DatabaseMigrationsSQLVMClient.BeginCreateOrUpdate.
type DatabaseMigrationsSQLVMClientCreateOrUpdateResponse struct {
	// Database Migration Resource for SQL Virtual Machine.
	DatabaseMigrationSQLVM
}

// DatabaseMigrationsSQLVMClientCutoverResponse contains the response from method DatabaseMigrationsSQLVMClient.BeginCutover.
type DatabaseMigrationsSQLVMClientCutoverResponse struct {
	// placeholder for future response values
}

// DatabaseMigrationsSQLVMClientGetResponse contains the response from method DatabaseMigrationsSQLVMClient.Get.
type DatabaseMigrationsSQLVMClientGetResponse struct {
	// Database Migration Resource for SQL Virtual Machine.
	DatabaseMigrationSQLVM
}

// FilesClientCreateOrUpdateResponse contains the response from method FilesClient.CreateOrUpdate.
type FilesClientCreateOrUpdateResponse struct {
	// A file resource
	ProjectFile
}

// FilesClientDeleteResponse contains the response from method FilesClient.Delete.
type FilesClientDeleteResponse struct {
	// placeholder for future response values
}

// FilesClientGetResponse contains the response from method FilesClient.Get.
type FilesClientGetResponse struct {
	// A file resource
	ProjectFile
}

// FilesClientListResponse contains the response from method FilesClient.NewListPager.
type FilesClientListResponse struct {
	// OData page of files
	FileList
}

// FilesClientReadResponse contains the response from method FilesClient.Read.
type FilesClientReadResponse struct {
	// File storage information.
	FileStorageInfo
}

// FilesClientReadWriteResponse contains the response from method FilesClient.ReadWrite.
type FilesClientReadWriteResponse struct {
	// File storage information.
	FileStorageInfo
}

// FilesClientUpdateResponse contains the response from method FilesClient.Update.
type FilesClientUpdateResponse struct {
	// A file resource
	ProjectFile
}

// MigrationServicesClientCreateOrUpdateResponse contains the response from method MigrationServicesClient.BeginCreateOrUpdate.
type MigrationServicesClientCreateOrUpdateResponse struct {
	// A Migration Service.
	MigrationService
}

// MigrationServicesClientDeleteResponse contains the response from method MigrationServicesClient.BeginDelete.
type MigrationServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// MigrationServicesClientGetResponse contains the response from method MigrationServicesClient.Get.
type MigrationServicesClientGetResponse struct {
	// A Migration Service.
	MigrationService
}

// MigrationServicesClientListByResourceGroupResponse contains the response from method MigrationServicesClient.NewListByResourceGroupPager.
type MigrationServicesClientListByResourceGroupResponse struct {
	// A list of Migration Service.
	MigrationServiceListResult
}

// MigrationServicesClientListBySubscriptionResponse contains the response from method MigrationServicesClient.NewListBySubscriptionPager.
type MigrationServicesClientListBySubscriptionResponse struct {
	// A list of Migration Service.
	MigrationServiceListResult
}

// MigrationServicesClientListMigrationsResponse contains the response from method MigrationServicesClient.NewListMigrationsPager.
type MigrationServicesClientListMigrationsResponse struct {
	// A list of Database Migrations.
	DatabaseMigrationBaseListResult
}

// MigrationServicesClientUpdateResponse contains the response from method MigrationServicesClient.BeginUpdate.
type MigrationServicesClientUpdateResponse struct {
	// A Migration Service.
	MigrationService
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list SQL operations.
	OperationListResult
}

// ProjectsClientCreateOrUpdateResponse contains the response from method ProjectsClient.CreateOrUpdate.
type ProjectsClientCreateOrUpdateResponse struct {
	// A project resource
	Project
}

// ProjectsClientDeleteResponse contains the response from method ProjectsClient.Delete.
type ProjectsClientDeleteResponse struct {
	// placeholder for future response values
}

// ProjectsClientGetResponse contains the response from method ProjectsClient.Get.
type ProjectsClientGetResponse struct {
	// A project resource
	Project
}

// ProjectsClientListResponse contains the response from method ProjectsClient.NewListPager.
type ProjectsClientListResponse struct {
	// OData page of project resources
	ProjectList
}

// ProjectsClientUpdateResponse contains the response from method ProjectsClient.Update.
type ProjectsClientUpdateResponse struct {
	// A project resource
	Project
}

// ResourceSKUsClientListSKUsResponse contains the response from method ResourceSKUsClient.NewListSKUsPager.
type ResourceSKUsClientListSKUsResponse struct {
	// The DMS (classic) List SKUs operation response.
	ResourceSKUsResult
}

// SQLMigrationServicesClientCreateOrUpdateResponse contains the response from method SQLMigrationServicesClient.BeginCreateOrUpdate.
type SQLMigrationServicesClientCreateOrUpdateResponse struct {
	// A SQL Migration Service.
	SQLMigrationService
}

// SQLMigrationServicesClientDeleteNodeResponse contains the response from method SQLMigrationServicesClient.DeleteNode.
type SQLMigrationServicesClientDeleteNodeResponse struct {
	// Details of node to be deleted.
	DeleteNode
}

// SQLMigrationServicesClientDeleteResponse contains the response from method SQLMigrationServicesClient.BeginDelete.
type SQLMigrationServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// SQLMigrationServicesClientGetResponse contains the response from method SQLMigrationServicesClient.Get.
type SQLMigrationServicesClientGetResponse struct {
	// A SQL Migration Service.
	SQLMigrationService
}

// SQLMigrationServicesClientListAuthKeysResponse contains the response from method SQLMigrationServicesClient.ListAuthKeys.
type SQLMigrationServicesClientListAuthKeysResponse struct {
	// An authentication key.
	AuthenticationKeys
}

// SQLMigrationServicesClientListByResourceGroupResponse contains the response from method SQLMigrationServicesClient.NewListByResourceGroupPager.
type SQLMigrationServicesClientListByResourceGroupResponse struct {
	// A list of SQL Migration Service.
	SQLMigrationListResult
}

// SQLMigrationServicesClientListBySubscriptionResponse contains the response from method SQLMigrationServicesClient.NewListBySubscriptionPager.
type SQLMigrationServicesClientListBySubscriptionResponse struct {
	// A list of SQL Migration Service.
	SQLMigrationListResult
}

// SQLMigrationServicesClientListMigrationsResponse contains the response from method SQLMigrationServicesClient.NewListMigrationsPager.
type SQLMigrationServicesClientListMigrationsResponse struct {
	// A list of Database Migrations.
	DatabaseMigrationListResult
}

// SQLMigrationServicesClientListMonitoringDataResponse contains the response from method SQLMigrationServicesClient.ListMonitoringData.
type SQLMigrationServicesClientListMonitoringDataResponse struct {
	// Integration Runtime Monitoring Data.
	IntegrationRuntimeMonitoringData
}

// SQLMigrationServicesClientRegenerateAuthKeysResponse contains the response from method SQLMigrationServicesClient.RegenerateAuthKeys.
type SQLMigrationServicesClientRegenerateAuthKeysResponse struct {
	// An authentication key to regenerate.
	RegenAuthKeys
}

// SQLMigrationServicesClientUpdateResponse contains the response from method SQLMigrationServicesClient.BeginUpdate.
type SQLMigrationServicesClientUpdateResponse struct {
	// A SQL Migration Service.
	SQLMigrationService
}

// ServiceTasksClientCancelResponse contains the response from method ServiceTasksClient.Cancel.
type ServiceTasksClientCancelResponse struct {
	// A task resource
	ProjectTask
}

// ServiceTasksClientCreateOrUpdateResponse contains the response from method ServiceTasksClient.CreateOrUpdate.
type ServiceTasksClientCreateOrUpdateResponse struct {
	// A task resource
	ProjectTask
}

// ServiceTasksClientDeleteResponse contains the response from method ServiceTasksClient.Delete.
type ServiceTasksClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceTasksClientGetResponse contains the response from method ServiceTasksClient.Get.
type ServiceTasksClientGetResponse struct {
	// A task resource
	ProjectTask
}

// ServiceTasksClientListResponse contains the response from method ServiceTasksClient.NewListPager.
type ServiceTasksClientListResponse struct {
	// OData page of tasks
	TaskList
}

// ServiceTasksClientUpdateResponse contains the response from method ServiceTasksClient.Update.
type ServiceTasksClientUpdateResponse struct {
	// A task resource
	ProjectTask
}

// ServicesClientCheckChildrenNameAvailabilityResponse contains the response from method ServicesClient.CheckChildrenNameAvailability.
type ServicesClientCheckChildrenNameAvailabilityResponse struct {
	// Indicates whether a proposed resource name is available
	NameAvailabilityResponse
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	// Indicates whether a proposed resource name is available
	NameAvailabilityResponse
}

// ServicesClientCheckStatusResponse contains the response from method ServicesClient.CheckStatus.
type ServicesClientCheckStatusResponse struct {
	// Service health status
	ServiceStatusResponse
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// An Azure Database Migration Service (classic) resource
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// An Azure Database Migration Service (classic) resource
	Service
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// OData page of service objects
	ServiceList
}

// ServicesClientListResponse contains the response from method ServicesClient.NewListPager.
type ServicesClientListResponse struct {
	// OData page of service objects
	ServiceList
}

// ServicesClientListSKUsResponse contains the response from method ServicesClient.NewListSKUsPager.
type ServicesClientListSKUsResponse struct {
	// OData page of available SKUs
	ServiceSKUList
}

// ServicesClientStartResponse contains the response from method ServicesClient.BeginStart.
type ServicesClientStartResponse struct {
	// placeholder for future response values
}

// ServicesClientStopResponse contains the response from method ServicesClient.BeginStop.
type ServicesClientStopResponse struct {
	// placeholder for future response values
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.BeginUpdate.
type ServicesClientUpdateResponse struct {
	// An Azure Database Migration Service (classic) resource
	Service
}

// TasksClientCancelResponse contains the response from method TasksClient.Cancel.
type TasksClientCancelResponse struct {
	// A task resource
	ProjectTask
}

// TasksClientCommandResponse contains the response from method TasksClient.Command.
type TasksClientCommandResponse struct {
	// Base class for all types of DMS (classic) command properties. If command is not supported by current client, this object
	// is returned.
	CommandPropertiesClassification
}

// TasksClientCreateOrUpdateResponse contains the response from method TasksClient.CreateOrUpdate.
type TasksClientCreateOrUpdateResponse struct {
	// A task resource
	ProjectTask
}

// TasksClientDeleteResponse contains the response from method TasksClient.Delete.
type TasksClientDeleteResponse struct {
	// placeholder for future response values
}

// TasksClientGetResponse contains the response from method TasksClient.Get.
type TasksClientGetResponse struct {
	// A task resource
	ProjectTask
}

// TasksClientListResponse contains the response from method TasksClient.NewListPager.
type TasksClientListResponse struct {
	// OData page of tasks
	TaskList
}

// TasksClientUpdateResponse contains the response from method TasksClient.Update.
type TasksClientUpdateResponse struct {
	// A task resource
	ProjectTask
}

// UsagesClientListResponse contains the response from method UsagesClient.NewListPager.
type UsagesClientListResponse struct {
	// OData page of quota objects
	QuotaList
}
