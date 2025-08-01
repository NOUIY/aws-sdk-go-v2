// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessPolicyType string

// Enum values for AccessPolicyType
const (
	// data policy type
	AccessPolicyTypeData AccessPolicyType = "data"
)

// Values returns all known values for AccessPolicyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessPolicyType) Values() []AccessPolicyType {
	return []AccessPolicyType{
		"data",
	}
}

type CollectionStatus string

// Enum values for CollectionStatus
const (
	// Creating collection resource
	CollectionStatusCreating CollectionStatus = "CREATING"
	// Deleting collection resource
	CollectionStatusDeleting CollectionStatus = "DELETING"
	// Collection resource is ready to use
	CollectionStatusActive CollectionStatus = "ACTIVE"
	// Collection resource create or delete failed
	CollectionStatusFailed CollectionStatus = "FAILED"
)

// Values returns all known values for CollectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CollectionStatus) Values() []CollectionStatus {
	return []CollectionStatus{
		"CREATING",
		"DELETING",
		"ACTIVE",
		"FAILED",
	}
}

type CollectionType string

// Enum values for CollectionType
const (
	// Search collection type
	CollectionTypeSearch CollectionType = "SEARCH"
	// Timeseries collection type
	CollectionTypeTimeseries CollectionType = "TIMESERIES"
	// Vectorsearch collection type
	CollectionTypeVectorsearch CollectionType = "VECTORSEARCH"
)

// Values returns all known values for CollectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CollectionType) Values() []CollectionType {
	return []CollectionType{
		"SEARCH",
		"TIMESERIES",
		"VECTORSEARCH",
	}
}

type IamIdentityCenterGroupAttribute string

// Enum values for IamIdentityCenterGroupAttribute
const (
	// Group ID
	IamIdentityCenterGroupAttributeGroupId IamIdentityCenterGroupAttribute = "GroupId"
	// Group Name
	IamIdentityCenterGroupAttributeGroupName IamIdentityCenterGroupAttribute = "GroupName"
)

// Values returns all known values for IamIdentityCenterGroupAttribute. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IamIdentityCenterGroupAttribute) Values() []IamIdentityCenterGroupAttribute {
	return []IamIdentityCenterGroupAttribute{
		"GroupId",
		"GroupName",
	}
}

type IamIdentityCenterUserAttribute string

// Enum values for IamIdentityCenterUserAttribute
const (
	// User ID
	IamIdentityCenterUserAttributeUserId IamIdentityCenterUserAttribute = "UserId"
	// User Name
	IamIdentityCenterUserAttributeUserName IamIdentityCenterUserAttribute = "UserName"
	// Email
	IamIdentityCenterUserAttributeEmail IamIdentityCenterUserAttribute = "Email"
)

// Values returns all known values for IamIdentityCenterUserAttribute. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IamIdentityCenterUserAttribute) Values() []IamIdentityCenterUserAttribute {
	return []IamIdentityCenterUserAttribute{
		"UserId",
		"UserName",
		"Email",
	}
}

type LifecyclePolicyType string

// Enum values for LifecyclePolicyType
const (
	// retention policy type
	LifecyclePolicyTypeRetention LifecyclePolicyType = "retention"
)

// Values returns all known values for LifecyclePolicyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (LifecyclePolicyType) Values() []LifecyclePolicyType {
	return []LifecyclePolicyType{
		"retention",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	// index resource type
	ResourceTypeIndex ResourceType = "index"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"index",
	}
}

type SecurityConfigType string

// Enum values for SecurityConfigType
const (
	// saml provider
	SecurityConfigTypeSaml SecurityConfigType = "saml"
	// iam identity center
	SecurityConfigTypeIamidentitycenter SecurityConfigType = "iamidentitycenter"
	// iam federation
	SecurityConfigTypeIamfederation SecurityConfigType = "iamfederation"
)

// Values returns all known values for SecurityConfigType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SecurityConfigType) Values() []SecurityConfigType {
	return []SecurityConfigType{
		"saml",
		"iamidentitycenter",
		"iamfederation",
	}
}

type SecurityPolicyType string

// Enum values for SecurityPolicyType
const (
	// encryption policy type
	SecurityPolicyTypeEncryption SecurityPolicyType = "encryption"
	// network policy type
	SecurityPolicyTypeNetwork SecurityPolicyType = "network"
)

// Values returns all known values for SecurityPolicyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SecurityPolicyType) Values() []SecurityPolicyType {
	return []SecurityPolicyType{
		"encryption",
		"network",
	}
}

type StandbyReplicas string

// Enum values for StandbyReplicas
const (
	// Standby replicas enabled
	StandbyReplicasEnabled StandbyReplicas = "ENABLED"
	// Standby replicas disabled
	StandbyReplicasDisabled StandbyReplicas = "DISABLED"
)

// Values returns all known values for StandbyReplicas. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StandbyReplicas) Values() []StandbyReplicas {
	return []StandbyReplicas{
		"ENABLED",
		"DISABLED",
	}
}

type VpcEndpointStatus string

// Enum values for VpcEndpointStatus
const (
	// Pending VPCEndpoint resource
	VpcEndpointStatusPending VpcEndpointStatus = "PENDING"
	// Deleting VPCEndpoint resource
	VpcEndpointStatusDeleting VpcEndpointStatus = "DELETING"
	// VPCEndpoint resource is ready to use
	VpcEndpointStatusActive VpcEndpointStatus = "ACTIVE"
	// VPCEndpoint resource create or delete failed
	VpcEndpointStatusFailed VpcEndpointStatus = "FAILED"
)

// Values returns all known values for VpcEndpointStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (VpcEndpointStatus) Values() []VpcEndpointStatus {
	return []VpcEndpointStatus{
		"PENDING",
		"DELETING",
		"ACTIVE",
		"FAILED",
	}
}
