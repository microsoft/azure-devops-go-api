// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package location

import (
    "github.com/google/uuid"
    "time"
)

type AccessMapping struct {
    AccessPoint *string `json:"accessPoint,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    Moniker *string `json:"moniker,omitempty"`
    // The service which owns this access mapping e.g. TFS, ELS, etc.
    ServiceOwner *uuid.UUID `json:"serviceOwner,omitempty"`
    // Part of the access mapping which applies context after the access point of the server.
    VirtualDirectory *string `json:"virtualDirectory,omitempty"`
}

// Data transfer class that holds information needed to set up a connection with a VSS server.
type ConnectionData struct {
    // The Id of the authenticated user who made this request. More information about the user can be obtained by passing this Id to the Identity service
    AuthenticatedUser *Identity `json:"authenticatedUser,omitempty"`
    // The Id of the authorized user who made this request. More information about the user can be obtained by passing this Id to the Identity service
    AuthorizedUser *Identity `json:"authorizedUser,omitempty"`
    // The id for the server.
    DeploymentId *uuid.UUID `json:"deploymentId,omitempty"`
    // The type for the server Hosted/OnPremises.
    DeploymentType *DeploymentFlags `json:"deploymentType,omitempty"`
    // The instance id for this host.
    InstanceId *uuid.UUID `json:"instanceId,omitempty"`
    // The last user access for this instance.  Null if not requested specifically.
    LastUserAccess *time.Time `json:"lastUserAccess,omitempty"`
    // Data that the location service holds.
    LocationServiceData *LocationServiceData `json:"locationServiceData,omitempty"`
    // The virtual directory of the host we are talking to.
    WebApplicationRelativeDirectory *string `json:"webApplicationRelativeDirectory,omitempty"`
}

// Enumeration of the options that can be passed in on Connect.
type ConnectOptions string

type connectOptionsValuesType struct {
    None ConnectOptions
    IncludeServices ConnectOptions
    IncludeLastUserAccess ConnectOptions
    IncludeInheritedDefinitionsOnly ConnectOptions
    IncludeNonInheritedDefinitionsOnly ConnectOptions
}

var ConnectOptionsValues = connectOptionsValuesType{
    // Retrieve no optional data.
    None: "none",
    // Includes information about AccessMappings and ServiceDefinitions.
    IncludeServices: "includeServices",
    // Includes the last user access for this host.
    IncludeLastUserAccess: "includeLastUserAccess",
    // This is only valid on the deployment host and when true. Will only return inherited definitions.
    IncludeInheritedDefinitionsOnly: "includeInheritedDefinitionsOnly",
    // When true will only return non inherited definitions. Only valid at non-deployment host.
    IncludeNonInheritedDefinitionsOnly: "includeNonInheritedDefinitionsOnly",
}

type DeploymentFlags string

type deploymentFlagsValuesType struct {
    None DeploymentFlags
    Hosted DeploymentFlags
    OnPremises DeploymentFlags
}

var DeploymentFlagsValues = deploymentFlagsValuesType{
    None: "none",
    Hosted: "hosted",
    OnPremises: "onPremises",
}

type Identity struct {
    // The custom display name for the identity (if any). Setting this property to an empty string will clear the existing custom display name. Setting this property to null will not affect the existing persisted value (since null values do not get sent over the wire or to the database)
    CustomDisplayName *string `json:"customDisplayName,omitempty"`
    Descriptor *string `json:"descriptor,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    IsActive *bool `json:"isActive,omitempty"`
    IsContainer *bool `json:"isContainer,omitempty"`
    MasterId *uuid.UUID `json:"masterId,omitempty"`
    MemberIds *[]uuid.UUID `json:"memberIds,omitempty"`
    MemberOf *[]string `json:"memberOf,omitempty"`
    Members *[]string `json:"members,omitempty"`
    MetaTypeId *int `json:"metaTypeId,omitempty"`
    Properties *interface{} `json:"properties,omitempty"`
    // The display name for the identity as specified by the source identity provider.
    ProviderDisplayName *string `json:"providerDisplayName,omitempty"`
    ResourceVersion *int `json:"resourceVersion,omitempty"`
    SocialDescriptor *string `json:"socialDescriptor,omitempty"`
    SubjectDescriptor *string `json:"subjectDescriptor,omitempty"`
    UniqueUserId *int `json:"uniqueUserId,omitempty"`
}

// Base Identity class to allow "trimmed" identity class in the GetConnectionData API Makes sure that on-the-wire representations of the derived classes are compatible with each other (e.g. Server responds with PublicIdentity object while client deserializes it as Identity object) Derived classes should not have additional [DataMember] properties
type IdentityBase struct {
    // The custom display name for the identity (if any). Setting this property to an empty string will clear the existing custom display name. Setting this property to null will not affect the existing persisted value (since null values do not get sent over the wire or to the database)
    CustomDisplayName *string `json:"customDisplayName,omitempty"`
    Descriptor *string `json:"descriptor,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    IsActive *bool `json:"isActive,omitempty"`
    IsContainer *bool `json:"isContainer,omitempty"`
    MasterId *uuid.UUID `json:"masterId,omitempty"`
    MemberIds *[]uuid.UUID `json:"memberIds,omitempty"`
    MemberOf *[]string `json:"memberOf,omitempty"`
    Members *[]string `json:"members,omitempty"`
    MetaTypeId *int `json:"metaTypeId,omitempty"`
    Properties *interface{} `json:"properties,omitempty"`
    // The display name for the identity as specified by the source identity provider.
    ProviderDisplayName *string `json:"providerDisplayName,omitempty"`
    ResourceVersion *int `json:"resourceVersion,omitempty"`
    SocialDescriptor *string `json:"socialDescriptor,omitempty"`
    SubjectDescriptor *string `json:"subjectDescriptor,omitempty"`
    UniqueUserId *int `json:"uniqueUserId,omitempty"`
}

type InheritLevel string

type inheritLevelValuesType struct {
    None InheritLevel
    Deployment InheritLevel
    Account InheritLevel
    Collection InheritLevel
    All InheritLevel
}

var InheritLevelValues = inheritLevelValuesType{
    None: "none",
    Deployment: "deployment",
    Account: "account",
    Collection: "collection",
    All: "all",
}

type LocationMapping struct {
    AccessMappingMoniker *string `json:"accessMappingMoniker,omitempty"`
    Location *string `json:"location,omitempty"`
}

// Data transfer class used to transfer data about the location service data over the web service.
type LocationServiceData struct {
    // Data about the access mappings contained by this location service.
    AccessMappings *[]AccessMapping `json:"accessMappings,omitempty"`
    // Data that the location service holds.
    ClientCacheFresh *bool `json:"clientCacheFresh,omitempty"`
    // The time to live on the location service cache.
    ClientCacheTimeToLive *int `json:"clientCacheTimeToLive,omitempty"`
    // The default access mapping moniker for the server.
    DefaultAccessMappingMoniker *string `json:"defaultAccessMappingMoniker,omitempty"`
    // The obsolete id for the last change that took place on the server (use LastChangeId64).
    LastChangeId *int `json:"lastChangeId,omitempty"`
    // The non-truncated 64-bit id for the last change that took place on the server.
    LastChangeId64 *uint64 `json:"lastChangeId64,omitempty"`
    // Data about the service definitions contained by this location service.
    ServiceDefinitions *[]ServiceDefinition `json:"serviceDefinitions,omitempty"`
    // The identifier of the deployment which is hosting this location data (e.g. SPS, TFS, ELS, Napa, etc.)
    ServiceOwner *uuid.UUID `json:"serviceOwner,omitempty"`
}

type RelativeToSetting string

type relativeToSettingValuesType struct {
    Context RelativeToSetting
    WebApplication RelativeToSetting
    FullyQualified RelativeToSetting
}

var RelativeToSettingValues = relativeToSettingValuesType{
    Context: "context",
    WebApplication: "webApplication",
    FullyQualified: "fullyQualified",
}

type ResourceAreaInfo struct {
    Id *uuid.UUID `json:"id,omitempty"`
    LocationUrl *string `json:"locationUrl,omitempty"`
    Name *string `json:"name,omitempty"`
}

type ServiceDefinition struct {
    Description *string `json:"description,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    Identifier *uuid.UUID `json:"identifier,omitempty"`
    InheritLevel *InheritLevel `json:"inheritLevel,omitempty"`
    LocationMappings *[]LocationMapping `json:"locationMappings,omitempty"`
    // Maximum api version that this resource supports (current server version for this resource). Copied from <c>ApiResourceLocation</c>.
    MaxVersion *string `json:"maxVersion,omitempty"`
    // Minimum api version that this resource supports. Copied from <c>ApiResourceLocation</c>.
    MinVersion *string `json:"minVersion,omitempty"`
    ParentIdentifier *uuid.UUID `json:"parentIdentifier,omitempty"`
    ParentServiceType *string `json:"parentServiceType,omitempty"`
    Properties *interface{} `json:"properties,omitempty"`
    RelativePath *string `json:"relativePath,omitempty"`
    RelativeToSetting *RelativeToSetting `json:"relativeToSetting,omitempty"`
    // The latest version of this resource location that is in "Release" (non-preview) mode. Copied from <c>ApiResourceLocation</c>.
    ReleasedVersion *string `json:"releasedVersion,omitempty"`
    // The current resource version supported by this resource location. Copied from <c>ApiResourceLocation</c>.
    ResourceVersion *int `json:"resourceVersion,omitempty"`
    // The service which owns this definition e.g. TFS, ELS, etc.
    ServiceOwner *uuid.UUID `json:"serviceOwner,omitempty"`
    ServiceType *string `json:"serviceType,omitempty"`
    Status *ServiceStatus `json:"status,omitempty"`
    ToolId *string `json:"toolId,omitempty"`
}

type ServiceStatus string

type serviceStatusValuesType struct {
    Assigned ServiceStatus
    Active ServiceStatus
    Moving ServiceStatus
}

var ServiceStatusValues = serviceStatusValuesType{
    Assigned: "assigned",
    Active: "active",
    Moving: "moving",
}
