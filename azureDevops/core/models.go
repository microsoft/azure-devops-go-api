// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package core

import (
    "github.com/google/uuid"
    "time"
)

type ConnectedServiceKind string

type connectedServiceKindValuesType struct {
    Custom ConnectedServiceKind
    AzureSubscription ConnectedServiceKind
    Chef ConnectedServiceKind
    Generic ConnectedServiceKind
}

var ConnectedServiceKindValues = connectedServiceKindValuesType{
    // Custom or unknown service
    Custom: "custom",
    // Azure Subscription
    AzureSubscription: "azureSubscription",
    // Chef Connection
    Chef: "chef",
    // Generic Connection
    Generic: "generic",
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

type IdentityData struct {
    IdentityIds *[]uuid.UUID `json:"identityIds,omitempty"`
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

type Process struct {
    Name *string `json:"name,omitempty"`
    Url *string `json:"url,omitempty"`
    Links *ReferenceLinks `json:"_links,omitempty"`
    Description *string `json:"description,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    IsDefault *bool `json:"isDefault,omitempty"`
    Type *ProcessType `json:"type,omitempty"`
}

// Type of process customization on a collection.
type ProcessCustomizationType string

type processCustomizationTypeValuesType struct {
    Xml ProcessCustomizationType
    Inherited ProcessCustomizationType
}

var ProcessCustomizationTypeValues = processCustomizationTypeValuesType{
    // Customization based on project-scoped xml customization
    Xml: "xml",
    // Customization based on process inheritance
    Inherited: "inherited",
}

type ProcessReference struct {
    Name *string `json:"name,omitempty"`
    Url *string `json:"url,omitempty"`
}

type ProcessType string

type processTypeValuesType struct {
    System ProcessType
    Custom ProcessType
    Inherited ProcessType
}

var ProcessTypeValues = processTypeValuesType{
    System: "system",
    Custom: "custom",
    Inherited: "inherited",
}

// Contains the image data for project avatar.
type ProjectAvatar struct {
    // The avatar image represented as a byte array.
    Image *[]byte `json:"image,omitempty"`
}

type ProjectChangeType string

type projectChangeTypeValuesType struct {
    Modified ProjectChangeType
    Deleted ProjectChangeType
    Added ProjectChangeType
}

var ProjectChangeTypeValues = projectChangeTypeValuesType{
    Modified: "modified",
    Deleted: "deleted",
    Added: "added",
}

// Contains information describing a project.
type ProjectInfo struct {
    // The abbreviated name of the project.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // The description of the project.
    Description *string `json:"description,omitempty"`
    // The id of the project.
    Id *uuid.UUID `json:"id,omitempty"`
    // The time that this project was last updated.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // The name of the project.
    Name *string `json:"name,omitempty"`
    // A set of name-value pairs storing additional property data related to the project.
    Properties *[]ProjectProperty `json:"properties,omitempty"`
    // The current revision of the project.
    Revision *uint64 `json:"revision,omitempty"`
    // The current state of the project.
    State *interface{} `json:"state,omitempty"`
    // A Uri that can be used to refer to this project.
    Uri *string `json:"uri,omitempty"`
    // The version number of the project.
    Version *uint64 `json:"version,omitempty"`
    // Indicates whom the project is visible to.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

type ProjectMessage struct {
    Project *ProjectInfo `json:"project,omitempty"`
    ProjectChangeType *ProjectChangeType `json:"projectChangeType,omitempty"`
    ShouldInvalidateSystemStore *bool `json:"shouldInvalidateSystemStore,omitempty"`
}

type ProjectProperties struct {
    // The team project Id
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    // The collection of team project properties
    Properties *[]ProjectProperty `json:"properties,omitempty"`
}

// A named value associated with a project.
type ProjectProperty struct {
    // The name of the property.
    Name *string `json:"name,omitempty"`
    // The value of the property.
    Value *interface{} `json:"value,omitempty"`
}

type ProjectVisibility string

type projectVisibilityValuesType struct {
    Private ProjectVisibility
    Public ProjectVisibility
}

var ProjectVisibilityValues = projectVisibilityValuesType{
    // The project is only visible to users with explicit access.
    Private: "private",
    // The project is visible to all.
    Public: "public",
}

type Proxy struct {
    Authorization *ProxyAuthorization `json:"authorization,omitempty"`
    // This is a description string
    Description *string `json:"description,omitempty"`
    // The friendly name of the server
    FriendlyName *string `json:"friendlyName,omitempty"`
    GlobalDefault *bool `json:"globalDefault,omitempty"`
    // This is a string representation of the site that the proxy server is located in (e.g. "NA-WA-RED")
    Site *string `json:"site,omitempty"`
    SiteDefault *bool `json:"siteDefault,omitempty"`
    // The URL of the proxy server
    Url *string `json:"url,omitempty"`
}

type ProxyAuthorization struct {
    // Gets or sets the endpoint used to obtain access tokens from the configured token service.
    AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
    // Gets or sets the client identifier for this proxy.
    ClientId *uuid.UUID `json:"clientId,omitempty"`
    // Gets or sets the user identity to authorize for on-prem.
    Identity *string `json:"identity,omitempty"`
    // Gets or sets the public key used to verify the identity of this proxy. Only specify on hosted.
    PublicKey *PublicKey `json:"publicKey,omitempty"`
}

// Represents the public key portion of an RSA asymmetric key.
type PublicKey struct {
    // Gets or sets the exponent for the public key.
    Exponent *[]byte `json:"exponent,omitempty"`
    // Gets or sets the modulus for the public key.
    Modulus *[]byte `json:"modulus,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

type SourceControlTypes string

type sourceControlTypesValuesType struct {
    Tfvc SourceControlTypes
    Git SourceControlTypes
}

var SourceControlTypesValues = sourceControlTypesValuesType{
    Tfvc: "tfvc",
    Git: "git",
}

// The Team Context for an operation.
type TeamContext struct {
    // The team project Id or name.  Ignored if ProjectId is set.
    Project *string `json:"project,omitempty"`
    // The Team Project ID.  Required if Project is not set.
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    // The Team Id or name.  Ignored if TeamId is set.
    Team *string `json:"team,omitempty"`
    // The Team Id
    TeamId *uuid.UUID `json:"teamId,omitempty"`
}

type TeamMember struct {
    Identity *IdentityRef `json:"identity,omitempty"`
    IsTeamAdmin *bool `json:"isTeamAdmin,omitempty"`
}

// Represents a Team Project object.
type TeamProject struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *interface{} `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
    // The links to other objects related to this object.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Set of capabilities this project has (such as process template & version control).
    Capabilities *map[string]map[string]string `json:"capabilities,omitempty"`
    // The shallow ref to the default team.
    DefaultTeam *WebApiTeamRef `json:"defaultTeam,omitempty"`
}

// Data contract for a TeamProjectCollection.
type TeamProjectCollection struct {
    // Collection Id.
    Id *uuid.UUID `json:"id,omitempty"`
    // Collection Name.
    Name *string `json:"name,omitempty"`
    // Collection REST Url.
    Url *string `json:"url,omitempty"`
    // The links to other objects related to this object.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Project collection description.
    Description *string `json:"description,omitempty"`
    // Process customization type on this collection. It can be Xml or Inherited.
    ProcessCustomizationType *ProcessCustomizationType `json:"processCustomizationType,omitempty"`
    // Project collection state.
    State *string `json:"state,omitempty"`
}

// Reference object for a TeamProjectCollection.
type TeamProjectCollectionReference struct {
    // Collection Id.
    Id *uuid.UUID `json:"id,omitempty"`
    // Collection Name.
    Name *string `json:"name,omitempty"`
    // Collection REST Url.
    Url *string `json:"url,omitempty"`
}

// Represents a shallow reference to a TeamProject.
type TeamProjectReference struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *interface{} `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

// A data transfer object that stores the metadata associated with the creation of temporary data.
type TemporaryDataCreatedDTO struct {
    ExpirationSeconds *int `json:"expirationSeconds,omitempty"`
    Origin *string `json:"origin,omitempty"`
    Value *interface{} `json:"value,omitempty"`
    ExpirationDate *time.Time `json:"expirationDate,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    Url *string `json:"url,omitempty"`
}

// A data transfer object that stores the metadata associated with the temporary data.
type TemporaryDataDTO struct {
    ExpirationSeconds *int `json:"expirationSeconds,omitempty"`
    Origin *string `json:"origin,omitempty"`
    Value *interface{} `json:"value,omitempty"`
}

// Updateable properties for a WebApiTeam.
type UpdateTeam struct {
    // New description for the team.
    Description *string `json:"description,omitempty"`
    // New name for the team.
    Name *string `json:"name,omitempty"`
}

type WebApiConnectedService struct {
    Url *string `json:"url,omitempty"`
    // The user who did the OAuth authentication to created this service
    AuthenticatedBy *IdentityRef `json:"authenticatedBy,omitempty"`
    // Extra description on the service.
    Description *string `json:"description,omitempty"`
    // Friendly Name of service connection
    FriendlyName *string `json:"friendlyName,omitempty"`
    // Id/Name of the connection service. For Ex: Subscription Id for Azure Connection
    Id *string `json:"id,omitempty"`
    // The kind of service.
    Kind *string `json:"kind,omitempty"`
    // The project associated with this service
    Project *TeamProjectReference `json:"project,omitempty"`
    // Optional uri to connect directly to the service such as https://windows.azure.com
    ServiceUri *string `json:"serviceUri,omitempty"`
}

type WebApiConnectedServiceDetails struct {
    Id *string `json:"id,omitempty"`
    Url *string `json:"url,omitempty"`
    // Meta data for service connection
    ConnectedServiceMetaData *WebApiConnectedService `json:"connectedServiceMetaData,omitempty"`
    // Credential info
    CredentialsXml *string `json:"credentialsXml,omitempty"`
    // Optional uri to connect directly to the service such as https://windows.azure.com
    EndPoint *string `json:"endPoint,omitempty"`
}

type WebApiConnectedServiceRef struct {
    Id *string `json:"id,omitempty"`
    Url *string `json:"url,omitempty"`
}

// The representation of data needed to create a tag definition which is sent across the wire.
type WebApiCreateTagRequestData struct {
    // Name of the tag definition that will be created.
    Name *string `json:"name,omitempty"`
}

type WebApiProject struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *interface{} `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
    // Set of capabilities this project has
    Capabilities *map[string]map[string]string `json:"capabilities,omitempty"`
    // Reference to collection which contains this project
    Collection *WebApiProjectCollectionRef `json:"collection,omitempty"`
    // Default team for this project
    DefaultTeam *WebApiTeamRef `json:"defaultTeam,omitempty"`
}

type WebApiProjectCollection struct {
    // Collection Tfs Url (Host Url)
    CollectionUrl *string `json:"collectionUrl,omitempty"`
    // Collection Guid
    Id *uuid.UUID `json:"id,omitempty"`
    // Collection Name
    Name *string `json:"name,omitempty"`
    // Collection REST Url
    Url *string `json:"url,omitempty"`
    // Project collection description
    Description *string `json:"description,omitempty"`
    // Project collection state
    State *string `json:"state,omitempty"`
}

type WebApiProjectCollectionRef struct {
    // Collection Tfs Url (Host Url)
    CollectionUrl *string `json:"collectionUrl,omitempty"`
    // Collection Guid
    Id *uuid.UUID `json:"id,omitempty"`
    // Collection Name
    Name *string `json:"name,omitempty"`
    // Collection REST Url
    Url *string `json:"url,omitempty"`
}

// The representation of a tag definition which is sent across the wire.
type WebApiTagDefinition struct {
    // Whether or not the tag definition is active.
    Active *bool `json:"active,omitempty"`
    // ID of the tag definition.
    Id *uuid.UUID `json:"id,omitempty"`
    // The name of the tag definition.
    Name *string `json:"name,omitempty"`
    // Resource URL for the Tag Definition.
    Url *string `json:"url,omitempty"`
}

type WebApiTeam struct {
    // Team (Identity) Guid. A Team Foundation ID.
    Id *uuid.UUID `json:"id,omitempty"`
    // Team name
    Name *string `json:"name,omitempty"`
    // Team REST API Url
    Url *string `json:"url,omitempty"`
    // Team description
    Description *string `json:"description,omitempty"`
    // Team identity.
    Identity *Identity `json:"identity,omitempty"`
    // Identity REST API Url to this team
    IdentityUrl *string `json:"identityUrl,omitempty"`
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    ProjectName *string `json:"projectName,omitempty"`
}

type WebApiTeamRef struct {
    // Team (Identity) Guid. A Team Foundation ID.
    Id *uuid.UUID `json:"id,omitempty"`
    // Team name
    Name *string `json:"name,omitempty"`
    // Team REST API Url
    Url *string `json:"url,omitempty"`
}
