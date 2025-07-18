// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armstorageactions

import "time"

// ElseCondition - The else block of storage task operation
type ElseCondition struct {
	// REQUIRED; List of operations to execute in the else block
	Operations []*StorageTaskOperation
}

// IfCondition - The if block of storage task operation
type IfCondition struct {
	// REQUIRED; Condition predicate to evaluate each object. See https://aka.ms/storagetaskconditions for valid properties and
	// operators.
	Condition *string

	// REQUIRED; List of operations to execute when the condition predicate satisfies.
	Operations []*StorageTaskOperation
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedServiceIdentityType

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// Operation - REST API Operation
//
// Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure
	// Resource Manager/control-plane operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; The Operation items on this page
	Value []*Operation

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// StorageTask - Represents Storage Task.
type StorageTask struct {
	// REQUIRED; The managed service identity of the resource.
	Identity *ManagedServiceIdentity

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Properties of the storage task.
	Properties *StorageTaskProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The name of the storage task within the specified resource group. Storage task names must be between 3 and 18
	// characters in length and use numbers and lower-case letters only.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// StorageTaskAction - The storage task action represents conditional statements and operations to be performed on target
// objects.
type StorageTaskAction struct {
	// REQUIRED; The if block of storage task operation
	If *IfCondition

	// The else block of storage task operation
	Else *ElseCondition
}

// StorageTaskAssignment - Storage Task Assignment associated with this Storage Task.
type StorageTaskAssignment struct {
	// READ-ONLY; Resource ID of the Storage Task Assignment.
	ID *string
}

// StorageTaskAssignmentsListResult - The response from the List Storage Tasks operation.
type StorageTaskAssignmentsListResult struct {
	// READ-ONLY; List of Storage Task Assignment Resource IDs associated with this Storage Task.
	Value []*StorageTaskAssignment

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// StorageTaskOperation - Represents an operation to be performed on the object
type StorageTaskOperation struct {
	// REQUIRED; The operation to be performed on the object.
	Name *StorageTaskOperationName

	// Action to be taken when the operation fails for a object.
	OnFailure *OnFailure

	// Action to be taken when the operation is successful for a object.
	OnSuccess *OnSuccess

	// Key-value parameters for the operation.
	Parameters map[string]*string
}

// StorageTaskPreviewAction - Storage Task Preview Action.
type StorageTaskPreviewAction struct {
	// REQUIRED; Properties of the storage task preview.
	Properties *StorageTaskPreviewActionProperties
}

// StorageTaskPreviewActionCondition - Represents the storage task conditions to be tested for a match with container and
// blob properties.
type StorageTaskPreviewActionCondition struct {
	// REQUIRED; Specify whether the else block is present in the condition.
	ElseBlockExists *bool

	// REQUIRED; The condition to be tested for a match with container and blob properties.
	If *StorageTaskPreviewActionIfCondition
}

// StorageTaskPreviewActionIfCondition - Represents storage task preview action condition.
type StorageTaskPreviewActionIfCondition struct {
	// Storage task condition to bes tested for a match.
	Condition *string
}

// StorageTaskPreviewActionProperties - Storage task preview action properties.
type StorageTaskPreviewActionProperties struct {
	// REQUIRED; Preview action to test
	Action *StorageTaskPreviewActionCondition

	// REQUIRED; Properties of some sample blobs in the container to test for matches with the preview action.
	Blobs []*StorageTaskPreviewBlobProperties

	// REQUIRED; Properties of a sample container to test for a match with the preview action.
	Container *StorageTaskPreviewContainerProperties
}

// StorageTaskPreviewBlobProperties - Storage task preview container properties
type StorageTaskPreviewBlobProperties struct {
	// metadata key value pairs to be tested for a match against the provided condition.
	Metadata []*StorageTaskPreviewKeyValueProperties

	// Name of test blob
	Name *string

	// properties key value pairs to be tested for a match against the provided condition.
	Properties []*StorageTaskPreviewKeyValueProperties

	// tags key value pairs to be tested for a match against the provided condition.
	Tags []*StorageTaskPreviewKeyValueProperties

	// READ-ONLY; Represents the condition block name that matched blob properties.
	MatchedBlock *MatchedBlockName
}

// StorageTaskPreviewContainerProperties - Storage task preview container properties
type StorageTaskPreviewContainerProperties struct {
	// metadata key value pairs to be tested for a match against the provided condition.
	Metadata []*StorageTaskPreviewKeyValueProperties

	// Name of test container
	Name *string
}

// StorageTaskPreviewKeyValueProperties - Storage task preview object key value pair properties.
type StorageTaskPreviewKeyValueProperties struct {
	// Represents the key property of the pair.
	Key *string

	// Represents the value property of the pair.
	Value *string
}

// StorageTaskProperties - Properties of the storage task.
type StorageTaskProperties struct {
	// REQUIRED; The storage task action that is executed
	Action *StorageTaskAction

	// REQUIRED; Text that describes the purpose of the storage task
	Description *string

	// REQUIRED; Storage Task is enabled when set to true and disabled when set to false
	Enabled *bool

	// READ-ONLY; The creation date and time of the storage task in UTC.
	CreationTimeInUTC *time.Time

	// READ-ONLY; Represents the provisioning state of the storage task.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Storage task version.
	TaskVersion *int64
}

// StorageTaskReportInstance - Storage Tasks run report instance
type StorageTaskReportInstance struct {
	// Storage task execution report for a run instance.
	Properties *StorageTaskReportProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// StorageTaskReportProperties - Storage task execution report for a run instance.
type StorageTaskReportProperties struct {
	// READ-ONLY; End time of the run instance. Filter options such as startTime gt '2023-06-26T20:51:24.4494016Z' and other comparison
	// operators can be used as described for DateTime properties in https://learn.microsoft.com/en-us/rest/api/storageservices/querying-tables-and-entities#supported-comparison-operators
	FinishTime *string

	// READ-ONLY; Total number of objects where task operation failed when was attempted. Filter options such as objectFailedCount
	// eq 0 and other comparison operators can be used as described for Numerical properties in https://learn.microsoft.com/en-us/rest/api/storageservices/querying-tables-and-entities#supported-comparison-operators
	ObjectFailedCount *string

	// READ-ONLY; Total number of objects that meet the storage tasks condition and were operated upon. Filter options such as
	// objectsOperatedOnCount ge 100 and other comparison operators can be used as described for Numerical properties in https://learn.microsoft.com/en-us/rest/api/storageservices/querying-tables-and-entities#supported-comparison-operators
	ObjectsOperatedOnCount *string

	// READ-ONLY; Total number of objects where task operation succeeded when was attempted.Filter options such as objectsSucceededCount
	// gt 150 and other comparison operators can be used as described for Numerical properties in https://learn.microsoft.com/en-us/rest/api/storageservices/querying-tables-and-entities#supported-comparison-operators
	ObjectsSucceededCount *string

	// READ-ONLY; Total number of objects that meet the condition as defined in the storage task assignment execution context.
	// Filter options such as objectsTargetedCount gt 50 and other comparison operators can be used as described for Numerical
	// properties in https://learn.microsoft.com/en-us/rest/api/storageservices/querying-tables-and-entities#supported-comparison-operators
	ObjectsTargetedCount *string

	// READ-ONLY; Represents the overall result of the execution for the run instance
	RunResult *RunResult

	// READ-ONLY; Represents the status of the execution.
	RunStatusEnum *RunStatusEnum

	// READ-ONLY; Well known Azure Storage error code that represents the error encountered during execution of the run instance.
	RunStatusError *string

	// READ-ONLY; Start time of the run instance. Filter options such as startTime gt '2023-06-26T20:51:24.4494016Z' and other
	// comparison operators can be used as described for DateTime properties in https://learn.microsoft.com/en-us/rest/api/storageservices/querying-tables-and-entities#supported-comparison-operators
	StartTime *string

	// READ-ONLY; Resource ID of the Storage Account where this reported run executed.
	StorageAccountID *string

	// READ-ONLY; Full path to the verbose report stored in the reporting container as specified in the assignment execution context
	// for the storage account.
	SummaryReportPath *string

	// READ-ONLY; Resource ID of the Storage Task Assignment associated with this reported run.
	TaskAssignmentID *string

	// READ-ONLY; Resource ID of the Storage Task applied during this run.
	TaskID *string

	// READ-ONLY; Storage Task Version
	TaskVersion *string
}

// StorageTaskReportSummary - Fetch Storage Tasks Run Summary.
type StorageTaskReportSummary struct {
	// READ-ONLY; Gets storage tasks run result summary.
	Value []*StorageTaskReportInstance

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// StorageTaskUpdateParameters - Parameters of the storage task update request
type StorageTaskUpdateParameters struct {
	// The identity of the resource.
	Identity *ManagedServiceIdentity

	// Properties of the storage task.
	Properties *StorageTaskUpdateProperties

	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater
	// in length than 128 characters and a value no greater in length than 256 characters.
	Tags map[string]*string
}

// StorageTaskUpdateProperties - Properties of the storage task.
type StorageTaskUpdateProperties struct {
	// The storage task action that is executed
	Action *StorageTaskAction

	// Text that describes the purpose of the storage task
	Description *string

	// Storage Task is enabled when set to true and disabled when set to false
	Enabled *bool

	// READ-ONLY; The creation date and time of the storage task in UTC.
	CreationTimeInUTC *time.Time

	// READ-ONLY; Represents the provisioning state of the storage task.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Storage task version.
	TaskVersion *int64
}

// StorageTasksListResult - The response from the List Storage Task operation.
type StorageTasksListResult struct {
	// READ-ONLY; Gets the list of storage tasks and their properties.
	Value []*StorageTask

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
