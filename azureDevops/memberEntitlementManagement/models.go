// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package memberEntitlementManagement

import (
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "time"
)

// License assigned to a user
type AccessLevel struct {
    // Type of Account License (e.g. Express, Stakeholder etc.)
    AccountLicenseType *AccountLicenseType `json:"accountLicenseType,omitempty"`
    // Assignment Source of the License (e.g. Group, Unknown etc.
    AssignmentSource *AssignmentSource `json:"assignmentSource,omitempty"`
    // Display name of the License
    LicenseDisplayName *string `json:"licenseDisplayName,omitempty"`
    // Licensing Source (e.g. Account. MSDN etc.)
    LicensingSource *LicensingSource `json:"licensingSource,omitempty"`
    // Type of MSDN License (e.g. Visual Studio Professional, Visual Studio Enterprise etc.)
    MsdnLicenseType *MsdnLicenseType `json:"msdnLicenseType,omitempty"`
    // User status in the account
    Status *AccountUserStatus `json:"status,omitempty"`
    // Status message.
    StatusMessage *string `json:"statusMessage,omitempty"`
}

type AccountLicenseType string

type accountLicenseTypeValuesType struct {
    None AccountLicenseType
    EarlyAdopter AccountLicenseType
    Express AccountLicenseType
    Professional AccountLicenseType
    Advanced AccountLicenseType
    Stakeholder AccountLicenseType
}

var AccountLicenseTypeValues = accountLicenseTypeValuesType{
    None: "none",
    EarlyAdopter: "earlyAdopter",
    Express: "express",
    Professional: "professional",
    Advanced: "advanced",
    Stakeholder: "stakeholder",
}

type AccountUserStatus string

type accountUserStatusValuesType struct {
    None AccountUserStatus
    Active AccountUserStatus
    Disabled AccountUserStatus
    Deleted AccountUserStatus
    Pending AccountUserStatus
    Expired AccountUserStatus
    PendingDisabled AccountUserStatus
}

var AccountUserStatusValues = accountUserStatusValuesType{
    None: "none",
    // User has signed in at least once to the VSTS account
    Active: "active",
    // User cannot sign in; primarily used by admin to temporarily remove a user due to absence or license reallocation
    Disabled: "disabled",
    // User is removed from the VSTS account by the VSTS account admin
    Deleted: "deleted",
    // User is invited to join the VSTS account by the VSTS account admin, but has not signed up/signed in yet
    Pending: "pending",
    // User can sign in; primarily used when license is in expired state and we give a grace period
    Expired: "expired",
    // User is disabled; if reenabled, they will still be in the Pending state
    PendingDisabled: "pendingDisabled",
}

type AssignmentSource string

type assignmentSourceValuesType struct {
    None AssignmentSource
    Unknown AssignmentSource
    GroupRule AssignmentSource
}

var AssignmentSourceValues = assignmentSourceValuesType{
    None: "none",
    Unknown: "unknown",
    GroupRule: "groupRule",
}

type BaseOperationResult struct {
    // List of error codes paired with their corresponding error messages
    Errors *[]azureDevops.KeyValuePair `json:"errors,omitempty"`
    // Success status of the operation
    IsSuccess *bool `json:"isSuccess,omitempty"`
}

// An extension assigned to a user
type Extension struct {
    // Assignment source for this extension. I.e. explicitly assigned or from a group rule.
    AssignmentSource *AssignmentSource `json:"assignmentSource,omitempty"`
    // Gallery Id of the Extension.
    Id *string `json:"id,omitempty"`
    // Friendly name of this extension.
    Name *string `json:"name,omitempty"`
    // Source of this extension assignment. Ex: msdn, account, none, etc.
    Source *LicensingSource `json:"source,omitempty"`
}

// Summary of Extensions in the organization.
type ExtensionSummaryData struct {
    // Count of Licenses already assigned.
    Assigned *int `json:"assigned,omitempty"`
    // Available Count.
    Available *int `json:"available,omitempty"`
    // Quantity
    IncludedQuantity *int `json:"includedQuantity,omitempty"`
    // Total Count.
    Total *int `json:"total,omitempty"`
    // Count of Extension Licenses assigned to users through msdn.
    AssignedThroughSubscription *int `json:"assignedThroughSubscription,omitempty"`
    // Gallery Id of the Extension
    ExtensionId *string `json:"extensionId,omitempty"`
    // Friendly name of this extension
    ExtensionName *string `json:"extensionName,omitempty"`
    // Whether its a Trial Version.
    IsTrialVersion *bool `json:"isTrialVersion,omitempty"`
    // Minimum License Required for the Extension.
    MinimumLicenseRequired *MinimumRequiredServiceLevel `json:"minimumLicenseRequired,omitempty"`
    // Days remaining for the Trial to expire.
    RemainingTrialDays *int `json:"remainingTrialDays,omitempty"`
    // Date on which the Trial expires.
    TrialExpiryDate *time.Time `json:"trialExpiryDate,omitempty"`
}

// Graph group entity
type GraphGroup struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // [Internal Use Only] The legacy descriptor is here in case you need to access old version IMS using identity descriptor.
    LegacyDescriptor *string `json:"legacyDescriptor,omitempty"`
    // The type of source provider for the origin identifier (ex:AD, AAD, MSA)
    Origin *string `json:"origin,omitempty"`
    // The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
    OriginId *string `json:"originId,omitempty"`
    // This field identifies the type of the graph subject (ex: Group, Scope, User).
    SubjectKind *string `json:"subjectKind,omitempty"`
    // This represents the name of the container of origin for a graph member. (For MSA this is "Windows Live ID", for AD the name of the domain, for AAD the tenantID of the directory, for VSTS groups the ScopeId, etc)
    Domain *string `json:"domain,omitempty"`
    // The email address of record for a given graph member. This may be different than the principal name.
    MailAddress *string `json:"mailAddress,omitempty"`
    // This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
    PrincipalName *string `json:"principalName,omitempty"`
    // A short phrase to help human readers disambiguate groups with similar names
    Description *string `json:"description,omitempty"`
}

type GraphMember struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // [Internal Use Only] The legacy descriptor is here in case you need to access old version IMS using identity descriptor.
    LegacyDescriptor *string `json:"legacyDescriptor,omitempty"`
    // The type of source provider for the origin identifier (ex:AD, AAD, MSA)
    Origin *string `json:"origin,omitempty"`
    // The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
    OriginId *string `json:"originId,omitempty"`
    // This field identifies the type of the graph subject (ex: Group, Scope, User).
    SubjectKind *string `json:"subjectKind,omitempty"`
    // This represents the name of the container of origin for a graph member. (For MSA this is "Windows Live ID", for AD the name of the domain, for AAD the tenantID of the directory, for VSTS groups the ScopeId, etc)
    Domain *string `json:"domain,omitempty"`
    // The email address of record for a given graph member. This may be different than the principal name.
    MailAddress *string `json:"mailAddress,omitempty"`
    // This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
    PrincipalName *string `json:"principalName,omitempty"`
}

// Top-level graph entity
type GraphSubject struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // [Internal Use Only] The legacy descriptor is here in case you need to access old version IMS using identity descriptor.
    LegacyDescriptor *string `json:"legacyDescriptor,omitempty"`
    // The type of source provider for the origin identifier (ex:AD, AAD, MSA)
    Origin *string `json:"origin,omitempty"`
    // The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
    OriginId *string `json:"originId,omitempty"`
    // This field identifies the type of the graph subject (ex: Group, Scope, User).
    SubjectKind *string `json:"subjectKind,omitempty"`
}

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

// Graph user entity
type GraphUser struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // [Internal Use Only] The legacy descriptor is here in case you need to access old version IMS using identity descriptor.
    LegacyDescriptor *string `json:"legacyDescriptor,omitempty"`
    // The type of source provider for the origin identifier (ex:AD, AAD, MSA)
    Origin *string `json:"origin,omitempty"`
    // The unique identifier from the system of origin. Typically a sid, object id or Guid. Linking and unlinking operations can cause this value to change for a user because the user is not backed by a different provider and has a different unique id in the new provider.
    OriginId *string `json:"originId,omitempty"`
    // This field identifies the type of the graph subject (ex: Group, Scope, User).
    SubjectKind *string `json:"subjectKind,omitempty"`
    // This represents the name of the container of origin for a graph member. (For MSA this is "Windows Live ID", for AD the name of the domain, for AAD the tenantID of the directory, for VSTS groups the ScopeId, etc)
    Domain *string `json:"domain,omitempty"`
    // The email address of record for a given graph member. This may be different than the principal name.
    MailAddress *string `json:"mailAddress,omitempty"`
    // This is the PrincipalName of this graph member from the source provider. The source provider may change this field over time and it is not guaranteed to be immutable for the life of the graph member by VSTS.
    PrincipalName *string `json:"principalName,omitempty"`
    // The short, generally unique name for the user in the backing directory. For AAD users, this corresponds to the mail nickname, which is often but not necessarily similar to the part of the user's mail address before the @ sign. For GitHub users, this corresponds to the GitHub user handle.
    DirectoryAlias *string `json:"directoryAlias,omitempty"`
    // When true, the group has been deleted in the identity provider
    IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
    // The meta type of the user in the origin, such as "member", "guest", etc. See UserMetaType for the set of possible values.
    MetaType *string `json:"metaType,omitempty"`
}

// Project Group (e.g. Contributor, Reader etc.)
type Group struct {
    // Display Name of the Group
    DisplayName *string `json:"displayName,omitempty"`
    // Group Type
    GroupType *GroupType `json:"groupType,omitempty"`
}

// A group entity with additional properties including its license, extensions, and project membership
type GroupEntitlement struct {
    // Extension Rules.
    ExtensionRules *[]Extension `json:"extensionRules,omitempty"`
    // Member reference.
    Group *GraphGroup `json:"group,omitempty"`
    // The unique identifier which matches the Id of the GraphMember.
    Id *uuid.UUID `json:"id,omitempty"`
    // [Readonly] The last time the group licensing rule was executed (regardless of whether any changes were made).
    LastExecuted *time.Time `json:"lastExecuted,omitempty"`
    // License Rule.
    LicenseRule *AccessLevel `json:"licenseRule,omitempty"`
    // Group members. Only used when creating a new group.
    Members *[]UserEntitlement `json:"members,omitempty"`
    // Relation between a project and the member's effective permissions in that project.
    ProjectEntitlements *[]ProjectEntitlement `json:"projectEntitlements,omitempty"`
    // The status of the group rule.
    Status *GroupLicensingRuleStatus `json:"status,omitempty"`
}

type GroupEntitlementOperationReference struct {
    // Unique identifier for the operation.
    Id *uuid.UUID `json:"id,omitempty"`
    // Unique identifier for the plugin.
    PluginId *uuid.UUID `json:"pluginId,omitempty"`
    // The current status of the operation.
    Status *OperationStatus `json:"status,omitempty"`
    // URL to get the full operation object.
    Url *string `json:"url,omitempty"`
    // Operation completed with success or failure.
    Completed *bool `json:"completed,omitempty"`
    // True if all operations were successful.
    HaveResultsSucceeded *bool `json:"haveResultsSucceeded,omitempty"`
    // List of results for each operation.
    Results *[]GroupOperationResult `json:"results,omitempty"`
}

type GroupLicensingRuleStatus string

type groupLicensingRuleStatusValuesType struct {
    ApplyPending GroupLicensingRuleStatus
    Applied GroupLicensingRuleStatus
    Incompatible GroupLicensingRuleStatus
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

type GroupOperationResult struct {
    // List of error codes paired with their corresponding error messages
    Errors *[]azureDevops.KeyValuePair `json:"errors,omitempty"`
    // Success status of the operation
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Identifier of the Group being acted upon
    GroupId *uuid.UUID `json:"groupId,omitempty"`
    // Result of the Groupentitlement after the operation
    Result *GroupEntitlement `json:"result,omitempty"`
}

// Group option to add a user to
type GroupOption struct {
    // Access Level
    AccessLevel *AccessLevel `json:"accessLevel,omitempty"`
    // Group
    Group *Group `json:"group,omitempty"`
}

// Used when adding users to a project. Each GroupType maps to a well-known group. The lowest GroupType should always be ProjectStakeholder
type GroupType string

type groupTypeValuesType struct {
    ProjectStakeholder GroupType
    ProjectReader GroupType
    ProjectContributor GroupType
    ProjectAdministrator GroupType
    Custom GroupType
}

var GroupTypeValues = groupTypeValuesType{
    ProjectStakeholder: "projectStakeholder",
    ProjectReader: "projectReader",
    ProjectContributor: "projectContributor",
    ProjectAdministrator: "projectAdministrator",
    Custom: "custom",
}

// The JSON model for a JSON Patch operation
type JsonPatchOperation struct {
    // The path to copy from for the Move/Copy operation.
    From *string `json:"from,omitempty"`
    // The patch operation
    Op *Operation `json:"op,omitempty"`
    // The path for the operation. In the case of an array, a zero based index can be used to specify the position in the array (e.g. /biscuits/0/name). The "-" character can be used instead of an index to insert at the end of the array (e.g. /biscuits/-).
    Path *string `json:"path,omitempty"`
    // The value for the operation. This is either a primitive or a JToken.
    Value *interface{} `json:"value,omitempty"`
}

// Summary of Licenses in the organization.
type LicenseSummaryData struct {
    // Count of Licenses already assigned.
    Assigned *int `json:"assigned,omitempty"`
    // Available Count.
    Available *int `json:"available,omitempty"`
    // Quantity
    IncludedQuantity *int `json:"includedQuantity,omitempty"`
    // Total Count.
    Total *int `json:"total,omitempty"`
    // Type of Account License.
    AccountLicenseType *AccountLicenseType `json:"accountLicenseType,omitempty"`
    // Count of Disabled Licenses.
    Disabled *int `json:"disabled,omitempty"`
    // Designates if this license quantity can be changed through purchase
    IsPurchasable *bool `json:"isPurchasable,omitempty"`
    // Name of the License.
    LicenseName *string `json:"licenseName,omitempty"`
    // Type of MSDN License.
    MsdnLicenseType *MsdnLicenseType `json:"msdnLicenseType,omitempty"`
    // Specifies the date when billing will charge for paid licenses
    NextBillingDate *time.Time `json:"nextBillingDate,omitempty"`
    // Source of the License.
    Source *LicensingSource `json:"source,omitempty"`
    // Total license count after next billing cycle
    TotalAfterNextBillingDate *int `json:"totalAfterNextBillingDate,omitempty"`
}

type LicensingSource string

type licensingSourceValuesType struct {
    None LicensingSource
    Account LicensingSource
    Msdn LicensingSource
    Profile LicensingSource
    Auto LicensingSource
    Trial LicensingSource
}

var LicensingSourceValues = licensingSourceValuesType{
    None: "none",
    Account: "account",
    Msdn: "msdn",
    Profile: "profile",
    Auto: "auto",
    Trial: "trial",
}

// Deprecated: Use UserEntitlement instead
type MemberEntitlement struct {
    // User's access level denoted by a license.
    AccessLevel *AccessLevel `json:"accessLevel,omitempty"`
    // [Readonly] Date the user was added to the collection.
    DateCreated *time.Time `json:"dateCreated,omitempty"`
    // User's extensions.
    Extensions *[]Extension `json:"extensions,omitempty"`
    // [Readonly] GroupEntitlements that this user belongs to.
    GroupAssignments *[]GroupEntitlement `json:"groupAssignments,omitempty"`
    // The unique identifier which matches the Id of the Identity associated with the GraphMember.
    Id *uuid.UUID `json:"id,omitempty"`
    // [Readonly] Date the user last accessed the collection.
    LastAccessedDate *time.Time `json:"lastAccessedDate,omitempty"`
    // Relation between a project and the user's effective permissions in that project.
    ProjectEntitlements *[]ProjectEntitlement `json:"projectEntitlements,omitempty"`
    // User reference.
    User *GraphUser `json:"user,omitempty"`
    // Member reference
    Member *GraphMember `json:"member,omitempty"`
}

type MemberEntitlementOperationReference struct {
    // Unique identifier for the operation.
    Id *uuid.UUID `json:"id,omitempty"`
    // Unique identifier for the plugin.
    PluginId *uuid.UUID `json:"pluginId,omitempty"`
    // The current status of the operation.
    Status *OperationStatus `json:"status,omitempty"`
    // URL to get the full operation object.
    Url *string `json:"url,omitempty"`
    // Operation completed with success or failure
    Completed *bool `json:"completed,omitempty"`
    // True if all operations were successful
    HaveResultsSucceeded *bool `json:"haveResultsSucceeded,omitempty"`
    // List of results for each operation
    Results *[]OperationResult `json:"results,omitempty"`
}

type MemberEntitlementsPatchResponse struct {
    // True if all operations were successful.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Result of the member entitlement after the operations. have been applied
    MemberEntitlement *MemberEntitlement `json:"memberEntitlement,omitempty"`
    // List of results for each operation
    OperationResults *[]OperationResult `json:"operationResults,omitempty"`
}

type MemberEntitlementsPostResponse struct {
    // True if all operations were successful.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Result of the member entitlement after the operations. have been applied
    MemberEntitlement *MemberEntitlement `json:"memberEntitlement,omitempty"`
    // Operation result
    OperationResult *OperationResult `json:"operationResult,omitempty"`
}

type MemberEntitlementsResponseBase struct {
    // True if all operations were successful.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Result of the member entitlement after the operations. have been applied
    MemberEntitlement *MemberEntitlement `json:"memberEntitlement,omitempty"`
}

type MinimumRequiredServiceLevel string

type minimumRequiredServiceLevelValuesType struct {
    None MinimumRequiredServiceLevel
    Express MinimumRequiredServiceLevel
    Advanced MinimumRequiredServiceLevel
    AdvancedPlus MinimumRequiredServiceLevel
    Stakeholder MinimumRequiredServiceLevel
}

var MinimumRequiredServiceLevelValues = minimumRequiredServiceLevelValuesType{
    // No service rights. The user cannot access the account
    None: "none",
    // Default or minimum service level
    Express: "express",
    // Premium service level - either by purchasing on the Azure portal or by purchasing the appropriate MSDN subscription
    Advanced: "advanced",
    // Only available to a specific set of MSDN Subscribers
    AdvancedPlus: "advancedPlus",
    // Stakeholder service level
    Stakeholder: "stakeholder",
}

type MsdnLicenseType string

type msdnLicenseTypeValuesType struct {
    None MsdnLicenseType
    Eligible MsdnLicenseType
    Professional MsdnLicenseType
    Platforms MsdnLicenseType
    TestProfessional MsdnLicenseType
    Premium MsdnLicenseType
    Ultimate MsdnLicenseType
    Enterprise MsdnLicenseType
}

var MsdnLicenseTypeValues = msdnLicenseTypeValuesType{
    None: "none",
    Eligible: "eligible",
    Professional: "professional",
    Platforms: "platforms",
    TestProfessional: "testProfessional",
    Premium: "premium",
    Ultimate: "ultimate",
    Enterprise: "enterprise",
}

type Operation string

type operationValuesType struct {
    Add Operation
    Remove Operation
    Replace Operation
    Move Operation
    Copy Operation
    Test Operation
}

var OperationValues = operationValuesType{
    Add: "add",
    Remove: "remove",
    Replace: "replace",
    Move: "move",
    Copy: "copy",
    Test: "test",
}

// Reference for an async operation.
type OperationReference struct {
    // Unique identifier for the operation.
    Id *uuid.UUID `json:"id,omitempty"`
    // Unique identifier for the plugin.
    PluginId *uuid.UUID `json:"pluginId,omitempty"`
    // The current status of the operation.
    Status *OperationStatus `json:"status,omitempty"`
    // URL to get the full operation object.
    Url *string `json:"url,omitempty"`
}

type OperationResult struct {
    // List of error codes paired with their corresponding error messages.
    Errors *[]azureDevops.KeyValuePair `json:"errors,omitempty"`
    // Success status of the operation.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Identifier of the Member being acted upon.
    MemberId *uuid.UUID `json:"memberId,omitempty"`
    // Result of the MemberEntitlement after the operation.
    Result *MemberEntitlement `json:"result,omitempty"`
}

// The status of an operation.
type OperationStatus string

type operationStatusValuesType struct {
    NotSet OperationStatus
    Queued OperationStatus
    InProgress OperationStatus
    Cancelled OperationStatus
    Succeeded OperationStatus
    Failed OperationStatus
}

var OperationStatusValues = operationStatusValuesType{
    // The operation does not have a status set.
    NotSet: "notSet",
    // The operation has been queued.
    Queued: "queued",
    // The operation is in progress.
    InProgress: "inProgress",
    // The operation was cancelled by the user.
    Cancelled: "cancelled",
    // The operation completed successfully.
    Succeeded: "succeeded",
    // The operation completed with a failure.
    Failed: "failed",
}

// A page of users
type PagedGraphMemberList struct {
    Members *[]UserEntitlement `json:"members,omitempty"`
}

type PagedList struct {
    ContinuationToken *string `json:"continuationToken,omitempty"`
    Items *[]interface{} `json:"items,omitempty"`
    TotalCount *int `json:"totalCount,omitempty"`
}

// Relation between a project and the user's effective permissions in that project.
type ProjectEntitlement struct {
    // Assignment Source (e.g. Group or Unknown).
    AssignmentSource *AssignmentSource `json:"assignmentSource,omitempty"`
    // Project Group (e.g. Contributor, Reader etc.)
    Group *Group `json:"group,omitempty"`
    // Deprecated: This property is deprecated. Please use ProjectPermissionInherited.
    IsProjectPermissionInherited *bool `json:"isProjectPermissionInherited,omitempty"`
    // Whether the user is inheriting permissions to a project through a Azure DevOps or AAD group membership.
    ProjectPermissionInherited *ProjectPermissionInherited `json:"projectPermissionInherited,omitempty"`
    // Project Ref
    ProjectRef *ProjectRef `json:"projectRef,omitempty"`
    // Team Ref.
    TeamRefs *[]TeamRef `json:"teamRefs,omitempty"`
}

type ProjectPermissionInherited string

type projectPermissionInheritedValuesType struct {
    NotSet ProjectPermissionInherited
    NotInherited ProjectPermissionInherited
    Inherited ProjectPermissionInherited
}

var ProjectPermissionInheritedValues = projectPermissionInheritedValuesType{
    NotSet: "notSet",
    NotInherited: "notInherited",
    Inherited: "inherited",
}

// A reference to a project
type ProjectRef struct {
    // Project ID.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project Name.
    Name *string `json:"name,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

type RuleOption string

type ruleOptionValuesType struct {
    ApplyGroupRule RuleOption
    TestApplyGroupRule RuleOption
}

var RuleOptionValues = ruleOptionValuesType{
    ApplyGroupRule: "applyGroupRule",
    TestApplyGroupRule: "testApplyGroupRule",
}

type SummaryData struct {
    // Count of Licenses already assigned.
    Assigned *int `json:"assigned,omitempty"`
    // Available Count.
    Available *int `json:"available,omitempty"`
    // Quantity
    IncludedQuantity *int `json:"includedQuantity,omitempty"`
    // Total Count.
    Total *int `json:"total,omitempty"`
}

type SummaryPropertyName string

type summaryPropertyNameValuesType struct {
    AccessLevels SummaryPropertyName
    Licenses SummaryPropertyName
    Extensions SummaryPropertyName
    Projects SummaryPropertyName
    Groups SummaryPropertyName
    All SummaryPropertyName
}

var SummaryPropertyNameValues = summaryPropertyNameValuesType{
    AccessLevels: "accessLevels",
    Licenses: "licenses",
    Extensions: "extensions",
    Projects: "projects",
    Groups: "groups",
    All: "all",
}

// A reference to a team
type TeamRef struct {
    // Team ID
    Id *uuid.UUID `json:"id,omitempty"`
    // Team Name
    Name *string `json:"name,omitempty"`
}

// A user entity with additional properties including thier license, extensions, and project membership
type UserEntitlement struct {
    // User's access level denoted by a license.
    AccessLevel *AccessLevel `json:"accessLevel,omitempty"`
    // [Readonly] Date the user was added to the collection.
    DateCreated *time.Time `json:"dateCreated,omitempty"`
    // User's extensions.
    Extensions *[]Extension `json:"extensions,omitempty"`
    // [Readonly] GroupEntitlements that this user belongs to.
    GroupAssignments *[]GroupEntitlement `json:"groupAssignments,omitempty"`
    // The unique identifier which matches the Id of the Identity associated with the GraphMember.
    Id *uuid.UUID `json:"id,omitempty"`
    // [Readonly] Date the user last accessed the collection.
    LastAccessedDate *time.Time `json:"lastAccessedDate,omitempty"`
    // Relation between a project and the user's effective permissions in that project.
    ProjectEntitlements *[]ProjectEntitlement `json:"projectEntitlements,omitempty"`
    // User reference.
    User *GraphUser `json:"user,omitempty"`
}

type UserEntitlementOperationReference struct {
    // Unique identifier for the operation.
    Id *uuid.UUID `json:"id,omitempty"`
    // Unique identifier for the plugin.
    PluginId *uuid.UUID `json:"pluginId,omitempty"`
    // The current status of the operation.
    Status *OperationStatus `json:"status,omitempty"`
    // URL to get the full operation object.
    Url *string `json:"url,omitempty"`
    // Operation completed with success or failure.
    Completed *bool `json:"completed,omitempty"`
    // True if all operations were successful.
    HaveResultsSucceeded *bool `json:"haveResultsSucceeded,omitempty"`
    // List of results for each operation.
    Results *[]UserEntitlementOperationResult `json:"results,omitempty"`
}

type UserEntitlementOperationResult struct {
    // List of error codes paired with their corresponding error messages.
    Errors *[]azureDevops.KeyValuePair `json:"errors,omitempty"`
    // Success status of the operation.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Result of the MemberEntitlement after the operation.
    Result *UserEntitlement `json:"result,omitempty"`
    // Identifier of the Member being acted upon.
    UserId *uuid.UUID `json:"userId,omitempty"`
}

type UserEntitlementProperty string

type userEntitlementPropertyValuesType struct {
    License UserEntitlementProperty
    Extensions UserEntitlementProperty
    Projects UserEntitlementProperty
    GroupRules UserEntitlementProperty
    All UserEntitlementProperty
}

var UserEntitlementPropertyValues = userEntitlementPropertyValuesType{
    License: "license",
    Extensions: "extensions",
    Projects: "projects",
    GroupRules: "groupRules",
    All: "all",
}

type UserEntitlementsPatchResponse struct {
    // True if all operations were successful.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Result of the user entitlement after the operations have been applied.
    UserEntitlement *UserEntitlement `json:"userEntitlement,omitempty"`
    // List of results for each operation.
    OperationResults *[]UserEntitlementOperationResult `json:"operationResults,omitempty"`
}

type UserEntitlementsPostResponse struct {
    // True if all operations were successful.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Result of the user entitlement after the operations have been applied.
    UserEntitlement *UserEntitlement `json:"userEntitlement,omitempty"`
    // Operation result.
    OperationResult *UserEntitlementOperationResult `json:"operationResult,omitempty"`
}

type UserEntitlementsResponseBase struct {
    // True if all operations were successful.
    IsSuccess *bool `json:"isSuccess,omitempty"`
    // Result of the user entitlement after the operations have been applied.
    UserEntitlement *UserEntitlement `json:"userEntitlement,omitempty"`
}

// Summary of licenses and extensions assigned to users in the organization
type UsersSummary struct {
    // Available Access Levels
    AvailableAccessLevels *[]AccessLevel `json:"availableAccessLevels,omitempty"`
    // Summary of Extensions in the organization
    Extensions *[]ExtensionSummaryData `json:"extensions,omitempty"`
    // Group Options
    GroupOptions *[]GroupOption `json:"groupOptions,omitempty"`
    // Summary of Licenses in the organization
    Licenses *[]LicenseSummaryData `json:"licenses,omitempty"`
    // Summary of Projects in the organization
    ProjectRefs *[]ProjectRef `json:"projectRefs,omitempty"`
}
