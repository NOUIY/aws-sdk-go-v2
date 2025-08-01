// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChannelRole string

// Enum values for ChannelRole
const (
	ChannelRoleAsker            ChannelRole = "ASKER"
	ChannelRoleExpert           ChannelRole = "EXPERT"
	ChannelRoleModerator        ChannelRole = "MODERATOR"
	ChannelRoleSupportrequestor ChannelRole = "SUPPORTREQUESTOR"
)

// Values returns all known values for ChannelRole. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChannelRole) Values() []ChannelRole {
	return []ChannelRole{
		"ASKER",
		"EXPERT",
		"MODERATOR",
		"SUPPORTREQUESTOR",
	}
}

type ChannelStatus string

// Enum values for ChannelStatus
const (
	ChannelStatusCreated      ChannelStatus = "CREATED"
	ChannelStatusCreating     ChannelStatus = "CREATING"
	ChannelStatusCreateFailed ChannelStatus = "CREATE_FAILED"
	ChannelStatusDeleted      ChannelStatus = "DELETED"
	ChannelStatusDeleting     ChannelStatus = "DELETING"
	ChannelStatusDeleteFailed ChannelStatus = "DELETE_FAILED"
)

// Values returns all known values for ChannelStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChannelStatus) Values() []ChannelStatus {
	return []ChannelStatus{
		"CREATED",
		"CREATING",
		"CREATE_FAILED",
		"DELETED",
		"DELETING",
		"DELETE_FAILED",
	}
}

type ConfigurationStatus string

// Enum values for ConfigurationStatus
const (
	ConfigurationStatusConfigured   ConfigurationStatus = "CONFIGURED"
	ConfigurationStatusUnconfigured ConfigurationStatus = "UNCONFIGURED"
)

// Values returns all known values for ConfigurationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationStatus) Values() []ConfigurationStatus {
	return []ConfigurationStatus{
		"CONFIGURED",
		"UNCONFIGURED",
	}
}

type FeatureEnableParameter string

// Enum values for FeatureEnableParameter
const (
	FeatureEnableParameterEnabled  FeatureEnableParameter = "ENABLED"
	FeatureEnableParameterDisabled FeatureEnableParameter = "DISABLED"
)

// Values returns all known values for FeatureEnableParameter. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FeatureEnableParameter) Values() []FeatureEnableParameter {
	return []FeatureEnableParameter{
		"ENABLED",
		"DISABLED",
	}
}

type FeatureEnableStatus string

// Enum values for FeatureEnableStatus
const (
	FeatureEnableStatusEnabled    FeatureEnableStatus = "ENABLED"
	FeatureEnableStatusDisabled   FeatureEnableStatus = "DISABLED"
	FeatureEnableStatusNotAllowed FeatureEnableStatus = "NOT_ALLOWED"
)

// Values returns all known values for FeatureEnableStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FeatureEnableStatus) Values() []FeatureEnableStatus {
	return []FeatureEnableStatus{
		"ENABLED",
		"DISABLED",
		"NOT_ALLOWED",
	}
}

type Role string

// Enum values for Role
const (
	RoleExpert           Role = "EXPERT"
	RoleModerator        Role = "MODERATOR"
	RoleAdministrator    Role = "ADMINISTRATOR"
	RoleSupportrequestor Role = "SUPPORTREQUESTOR"
)

// Values returns all known values for Role. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Role) Values() []Role {
	return []Role{
		"EXPERT",
		"MODERATOR",
		"ADMINISTRATOR",
		"SUPPORTREQUESTOR",
	}
}

type TierLevel string

// Enum values for TierLevel
const (
	TierLevelBasic    TierLevel = "BASIC"
	TierLevelStandard TierLevel = "STANDARD"
)

// Values returns all known values for TierLevel. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TierLevel) Values() []TierLevel {
	return []TierLevel{
		"BASIC",
		"STANDARD",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}

type VanityDomainStatus string

// Enum values for VanityDomainStatus
const (
	VanityDomainStatusPending    VanityDomainStatus = "PENDING"
	VanityDomainStatusApproved   VanityDomainStatus = "APPROVED"
	VanityDomainStatusUnapproved VanityDomainStatus = "UNAPPROVED"
)

// Values returns all known values for VanityDomainStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VanityDomainStatus) Values() []VanityDomainStatus {
	return []VanityDomainStatus{
		"PENDING",
		"APPROVED",
		"UNAPPROVED",
	}
}
