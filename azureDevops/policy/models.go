// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package policy

import (
    "github.com/google/uuid"
    "time"
)

type GraphSubjectBase struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
}

type IdentityRef struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph user referenced in the "self" entry of the IdentityRef "_links" dictionary
    DirectoryAlias *string `json:"directoryAlias,omitempty"`
    Id *string `json:"id,omitempty"`
    // Deprecated - Available in the "avatar" entry of the IdentityRef "_links" dictionary
    ImageUrl *string `json:"imageUrl,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph membership state referenced in the "membershipState" entry of the GraphUser "_links" dictionary
    Inactive *bool `json:"inactive,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsAadUserType/Descriptor.IsAadGroupType)
    IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsGroupType)
    IsContainer *bool `json:"isContainer,omitempty"`
    IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
    // Deprecated - not in use in most preexisting implementations of ToIdentityRef
    ProfileUrl *string `json:"profileUrl,omitempty"`
    // Deprecated - use Domain+PrincipalName instead
    UniqueName *string `json:"uniqueName,omitempty"`
}

// The full policy configuration with settings.
type PolicyConfiguration struct {
    // The policy configuration ID.
    Id *int `json:"id,omitempty"`
    // The policy configuration type.
    Type *PolicyTypeRef `json:"type,omitempty"`
    // The URL where the policy configuration can be retrieved.
    Url *string `json:"url,omitempty"`
    // The policy configuration revision ID.
    Revision *int `json:"revision,omitempty"`
    // The links to other objects related to this object.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // A reference to the identity that created the policy.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // The date and time when the policy was created.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Indicates whether the policy is blocking.
    IsBlocking *bool `json:"isBlocking,omitempty"`
    // Indicates whether the policy has been (soft) deleted.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Indicates whether the policy is enabled.
    IsEnabled *bool `json:"isEnabled,omitempty"`
    // The policy configuration settings.
    Settings *interface{} `json:"settings,omitempty"`
}

// Policy configuration reference.
type PolicyConfigurationRef struct {
    // The policy configuration ID.
    Id *int `json:"id,omitempty"`
    // The policy configuration type.
    Type *PolicyTypeRef `json:"type,omitempty"`
    // The URL where the policy configuration can be retrieved.
    Url *string `json:"url,omitempty"`
}

// This record encapsulates the current state of a policy as it applies to one specific pull request. Each pull request has a unique PolicyEvaluationRecord for each pull request which the policy applies to.
type PolicyEvaluationRecord struct {
    // Links to other related objects
    Links *ReferenceLinks `json:"_links,omitempty"`
    // A string which uniquely identifies the target of a policy evaluation.
    ArtifactId *string `json:"artifactId,omitempty"`
    // Time when this policy finished evaluating on this pull request.
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Contains all configuration data for the policy which is being evaluated.
    Configuration *PolicyConfiguration `json:"configuration,omitempty"`
    // Internal context data of this policy evaluation.
    Context *interface{} `json:"context,omitempty"`
    // Guid which uniquely identifies this evaluation record (one policy running on one pull request).
    EvaluationId *uuid.UUID `json:"evaluationId,omitempty"`
    // Time when this policy was first evaluated on this pull request.
    StartedDate *time.Time `json:"startedDate,omitempty"`
    // Status of the policy (Running, Approved, Failed, etc.)
    Status *PolicyEvaluationStatus `json:"status,omitempty"`
}

// Status of a policy which is running against a specific pull request.
type PolicyEvaluationStatus string

// User-friendly policy type with description (used for querying policy types).
type PolicyType struct {
    // Display name of the policy type.
    DisplayName *string `json:"displayName,omitempty"`
    // The policy type ID.
    Id *uuid.UUID `json:"id,omitempty"`
    // The URL where the policy type can be retrieved.
    Url *string `json:"url,omitempty"`
    // The links to other objects related to this object.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Detailed description of the policy type.
    Description *string `json:"description,omitempty"`
}

// Policy type reference.
type PolicyTypeRef struct {
    // Display name of the policy type.
    DisplayName *string `json:"displayName,omitempty"`
    // The policy type ID.
    Id *uuid.UUID `json:"id,omitempty"`
    // The URL where the policy type can be retrieved.
    Url *string `json:"url,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

// A particular revision for a policy configuration.
type VersionedPolicyConfigurationRef struct {
    // The policy configuration ID.
    Id *int `json:"id,omitempty"`
    // The policy configuration type.
    Type *PolicyTypeRef `json:"type,omitempty"`
    // The URL where the policy configuration can be retrieved.
    Url *string `json:"url,omitempty"`
    // The policy configuration revision ID.
    Revision *int `json:"revision,omitempty"`
}
