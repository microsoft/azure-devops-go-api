// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package licensingrule

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/licensing"
	"github.com/microsoft/azure-devops-go-api/azuredevops/operations"
	"time"
)

type ApplicationStatus struct {
	Extensions  *[]ExtensionApplicationStatus `json:"extensions,omitempty"`
	IsTruncated *bool                         `json:"isTruncated,omitempty"`
	Licenses    *[]LicenseApplicationStatus   `json:"licenses,omitempty"`
	Option      *RuleOption                   `json:"option,omitempty"`
	Status      *operations.OperationStatus   `json:"status,omitempty"`
}

type ExtensionApplicationStatus struct {
	Assigned              *int    `json:"assigned,omitempty"`
	Failed                *int    `json:"failed,omitempty"`
	InsufficientResources *int    `json:"insufficientResources,omitempty"`
	ExtensionId           *string `json:"extensionId,omitempty"`
	Incompatible          *int    `json:"incompatible,omitempty"`
	Unassigned            *int    `json:"unassigned,omitempty"`
}

// Represents an Extension Rule
type ExtensionRule struct {
	// Extension Id
	ExtensionId *string `json:"extensionId,omitempty"`
	// Status of the group rule (applied, missing licenses, etc)
	Status *GroupLicensingRuleStatus `json:"status,omitempty"`
}

// Batching of subjects to lookup using the Graph API
type GraphSubjectLookup struct {
	LookupKeys *[]GraphSubjectLookupKey `json:"lookupKeys,omitempty"`
}

type GraphSubjectLookupKey struct {
	Descriptor *string `json:"descriptor,omitempty"`
}

// Do not attempt to use this type to create a new user. Use one of the subclasses instead. This type does not contain sufficient fields to create a new user.
type GraphUserCreationContext struct {
	// Optional: If provided, we will use this identifier for the storage key of the created user
	StorageKey *uuid.UUID `json:"storageKey,omitempty"`
}

// Use this type to create a new user using the mail address as a reference to an existing user from an external AD or AAD backed provider. This is the subset of GraphUser fields required for creation of a GraphUser for the AD and AAD use case when looking up the user by its mail address in the backing provider.
type GraphUserMailAddressCreationContext struct {
	// Optional: If provided, we will use this identifier for the storage key of the created user
	StorageKey  *uuid.UUID `json:"storageKey,omitempty"`
	MailAddress *string    `json:"mailAddress,omitempty"`
}

// Use this type to create a new user using the OriginID as a reference to an existing user from an external AD or AAD backed provider. This is the subset of GraphUser fields required for creation of a GraphUser for the AD and AAD use case when looking up the user by its unique ID in the backing provider.
type GraphUserOriginIdCreationContext struct {
	// Optional: If provided, we will use this identifier for the storage key of the created user
	StorageKey *uuid.UUID `json:"storageKey,omitempty"`
	// This should be the name of the origin provider. Example: github.com
	Origin *string `json:"origin,omitempty"`
	// This should be the object id or sid of the user from the source AD or AAD provider. Example: d47d025a-ce2f-4a79-8618-e8862ade30dd Team Services will communicate with the source provider to fill all other fields on creation.
	OriginId *string `json:"originId,omitempty"`
}

// Use this type to create a new user using the principal name as a reference to an existing user from an external AD or AAD backed provider. This is the subset of GraphUser fields required for creation of a GraphUser for the AD and AAD use case when looking up the user by its principal name in the backing provider.
type GraphUserPrincipalNameCreationContext struct {
	// Optional: If provided, we will use this identifier for the storage key of the created user
	StorageKey *uuid.UUID `json:"storageKey,omitempty"`
	// This should be the principal name or upn of the user in the source AD or AAD provider. Example: jamal@contoso.com Team Services will communicate with the source provider to fill all other fields on creation.
	PrincipalName *string `json:"principalName,omitempty"`
}

// Represents a GroupLicensingRule
type GroupLicensingRule struct {
	// Extension Rules
	ExtensionRules *[]ExtensionRule `json:"extensionRules,omitempty"`
	// License Rule
	LicenseRule *LicenseRule `json:"licenseRule,omitempty"`
	// SubjectDescriptor for the rule
	SubjectDescriptor *string `json:"subjectDescriptor,omitempty"`
}

type GroupLicensingRuleStatus string

type groupLicensingRuleStatusValuesType struct {
	ApplyPending  GroupLicensingRuleStatus
	Applied       GroupLicensingRuleStatus
	Incompatible  GroupLicensingRuleStatus
	UnableToApply GroupLicensingRuleStatus
}

var GroupLicensingRuleStatusValues = groupLicensingRuleStatusValuesType{
	// Rule is created or updated, but apply is pending
	ApplyPending: "applyPending",
	// Rule is applied
	Applied: "applied",
	// The group rule was incompatible
	Incompatible: "incompatible",
	// Rule failed to apply unexpectedly and should be retried
	UnableToApply: "unableToApply",
}

// Represents an GroupLicensingRuleUpdate Model
type GroupLicensingRuleUpdate struct {
	// Extensions to Add
	ExtensionsToAdd *[]string `json:"extensionsToAdd,omitempty"`
	// Extensions to Remove
	ExtensionsToRemove *[]string `json:"extensionsToRemove,omitempty"`
	// New License
	License *licensing.License `json:"license,omitempty"`
	// SubjectDescriptor for the rule
	SubjectDescriptor *string `json:"subjectDescriptor,omitempty"`
}

type LicenseApplicationStatus struct {
	Assigned              *int                          `json:"assigned,omitempty"`
	Failed                *int                          `json:"failed,omitempty"`
	InsufficientResources *int                          `json:"insufficientResources,omitempty"`
	AccountUserLicense    *licensing.AccountUserLicense `json:"accountUserLicense,omitempty"`
	License               *licensing.License            `json:"license,omitempty"`
}

// Represents a License Rule
type LicenseRule struct {
	// The last time the rule was executed (regardless of whether any changes were made)
	LastExecuted *time.Time `json:"lastExecuted,omitempty"`
	// License
	License *licensing.License `json:"license,omitempty"`
	// Status of the group rule (applied, missing licenses, etc)
	Status *GroupLicensingRuleStatus `json:"status,omitempty"`
}

type LicensingApplicationStatus struct {
	Assigned              *int `json:"assigned,omitempty"`
	Failed                *int `json:"failed,omitempty"`
	InsufficientResources *int `json:"insufficientResources,omitempty"`
}

type RuleOption string

type ruleOptionValuesType struct {
	ApplyGroupRule     RuleOption
	TestApplyGroupRule RuleOption
}

var RuleOptionValues = ruleOptionValuesType{
	ApplyGroupRule:     "applyGroupRule",
	TestApplyGroupRule: "testApplyGroupRule",
}
