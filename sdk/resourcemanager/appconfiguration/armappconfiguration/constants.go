//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
	moduleVersion = "v2.2.0"
)

// ActionsRequired - Any action that is required beyond basic workflow (approve/ reject/ disconnect)
type ActionsRequired string

const (
	ActionsRequiredNone     ActionsRequired = "None"
	ActionsRequiredRecreate ActionsRequired = "Recreate"
)

// PossibleActionsRequiredValues returns the possible values for the ActionsRequired const type.
func PossibleActionsRequiredValues() []ActionsRequired {
	return []ActionsRequired{
		ActionsRequiredNone,
		ActionsRequiredRecreate,
	}
}

// AuthenticationMode - The data plane proxy authentication mode. This property manages the authentication mode of request
// to the data plane resources.
type AuthenticationMode string

const (
	// AuthenticationModeLocal - The local authentication mode. Users are not required to have data plane permissions if local
	// authentication is not disabled.
	AuthenticationModeLocal AuthenticationMode = "Local"
	// AuthenticationModePassThrough - The pass-through authentication mode. User identity will be passed through from Azure Resource
	// Manager (ARM), requiring user to have data plane action permissions (Available via App Configuration Data Owner/ App Configuration
	// Data Reader).
	AuthenticationModePassThrough AuthenticationMode = "Pass-through"
)

// PossibleAuthenticationModeValues returns the possible values for the AuthenticationMode const type.
func PossibleAuthenticationModeValues() []AuthenticationMode {
	return []AuthenticationMode{
		AuthenticationModeLocal,
		AuthenticationModePassThrough,
	}
}

// CompositionType - The composition type describes how the key-values within the snapshot are composed. The 'key' composition
// type ensures there are no two key-values containing the same key. The 'key_label' composition
// type ensures there are no two key-values containing the same key and label.
type CompositionType string

const (
	CompositionTypeKey      CompositionType = "Key"
	CompositionTypeKeyLabel CompositionType = "Key_Label"
)

// PossibleCompositionTypeValues returns the possible values for the CompositionType const type.
func PossibleCompositionTypeValues() []CompositionType {
	return []CompositionType{
		CompositionTypeKey,
		CompositionTypeKeyLabel,
	}
}

// ConfigurationResourceType - The resource type to check for name availability.
type ConfigurationResourceType string

const (
	ConfigurationResourceTypeMicrosoftAppConfigurationConfigurationStores ConfigurationResourceType = "Microsoft.AppConfiguration/configurationStores"
)

// PossibleConfigurationResourceTypeValues returns the possible values for the ConfigurationResourceType const type.
func PossibleConfigurationResourceTypeValues() []ConfigurationResourceType {
	return []ConfigurationResourceType{
		ConfigurationResourceTypeMicrosoftAppConfigurationConfigurationStores,
	}
}

// ConnectionStatus - The private link service connection status.
type ConnectionStatus string

const (
	ConnectionStatusApproved     ConnectionStatus = "Approved"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusPending      ConnectionStatus = "Pending"
	ConnectionStatusRejected     ConnectionStatus = "Rejected"
)

// PossibleConnectionStatusValues returns the possible values for the ConnectionStatus const type.
func PossibleConnectionStatusValues() []ConnectionStatus {
	return []ConnectionStatus{
		ConnectionStatusApproved,
		ConnectionStatusDisconnected,
		ConnectionStatusPending,
		ConnectionStatusRejected,
	}
}

// CreateMode - Indicates whether the configuration store need to be recovered.
type CreateMode string

const (
	CreateModeDefault CreateMode = "Default"
	CreateModeRecover CreateMode = "Recover"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeRecover,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// IdentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created
// identity and a set of user-assigned identities. The type 'None' will remove any
// identities.
type IdentityType string

const (
	IdentityTypeNone                       IdentityType = "None"
	IdentityTypeSystemAssigned             IdentityType = "SystemAssigned"
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned, UserAssigned"
	IdentityTypeUserAssigned               IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeSystemAssignedUserAssigned,
		IdentityTypeUserAssigned,
	}
}

// PrivateLinkDelegation - The data plane proxy private link delegation. This property manages if a request from delegated
// Azure Resource Manager (ARM) private link is allowed when the data plane resource requires private link.
type PrivateLinkDelegation string

const (
	// PrivateLinkDelegationDisabled - Request is denied if the resource requires private link.
	PrivateLinkDelegationDisabled PrivateLinkDelegation = "Disabled"
	// PrivateLinkDelegationEnabled - Azure Resource Manager (ARM) private endpoint is required if the resource requires private
	// link.
	PrivateLinkDelegationEnabled PrivateLinkDelegation = "Enabled"
)

// PossiblePrivateLinkDelegationValues returns the possible values for the PrivateLinkDelegation const type.
func PossiblePrivateLinkDelegationValues() []PrivateLinkDelegation {
	return []PrivateLinkDelegation{
		PrivateLinkDelegationDisabled,
		PrivateLinkDelegationEnabled,
	}
}

// ProvisioningState - The provisioning state of the configuration store.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Control permission for data plane traffic coming from public networks while private endpoint is enabled.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ReplicaProvisioningState - The provisioning state of the replica.
type ReplicaProvisioningState string

const (
	ReplicaProvisioningStateCanceled  ReplicaProvisioningState = "Canceled"
	ReplicaProvisioningStateCreating  ReplicaProvisioningState = "Creating"
	ReplicaProvisioningStateDeleting  ReplicaProvisioningState = "Deleting"
	ReplicaProvisioningStateFailed    ReplicaProvisioningState = "Failed"
	ReplicaProvisioningStateSucceeded ReplicaProvisioningState = "Succeeded"
)

// PossibleReplicaProvisioningStateValues returns the possible values for the ReplicaProvisioningState const type.
func PossibleReplicaProvisioningStateValues() []ReplicaProvisioningState {
	return []ReplicaProvisioningState{
		ReplicaProvisioningStateCanceled,
		ReplicaProvisioningStateCreating,
		ReplicaProvisioningStateDeleting,
		ReplicaProvisioningStateFailed,
		ReplicaProvisioningStateSucceeded,
	}
}

// SnapshotStatus - The current status of the snapshot.
type SnapshotStatus string

const (
	SnapshotStatusArchived     SnapshotStatus = "Archived"
	SnapshotStatusFailed       SnapshotStatus = "Failed"
	SnapshotStatusProvisioning SnapshotStatus = "Provisioning"
	SnapshotStatusReady        SnapshotStatus = "Ready"
)

// PossibleSnapshotStatusValues returns the possible values for the SnapshotStatus const type.
func PossibleSnapshotStatusValues() []SnapshotStatus {
	return []SnapshotStatus{
		SnapshotStatusArchived,
		SnapshotStatusFailed,
		SnapshotStatusProvisioning,
		SnapshotStatusReady,
	}
}
