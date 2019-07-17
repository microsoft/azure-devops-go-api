// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package contributions

import (
    "github.com/google/uuid"
    "time"
)

// Representation of a ContributionNode that can be used for serialized to clients.
type ClientContribution struct {
    // Description of the contribution/type
    Description *string `json:"description,omitempty"`
    // Fully qualified identifier of the contribution/type
    Id *string `json:"id,omitempty"`
    // Includes is a set of contributions that should have this contribution included in their targets list.
    Includes *[]string `json:"includes,omitempty"`
    // Properties/attributes of this contribution
    Properties *interface{} `json:"properties,omitempty"`
    // The ids of the contribution(s) that this contribution targets. (parent contributions)
    Targets *[]string `json:"targets,omitempty"`
    // Id of the Contribution Type
    Type *string `json:"type,omitempty"`
}

// Representation of a ContributionNode that can be used for serialized to clients.
type ClientContributionNode struct {
    // List of ids for contributions which are children to the current contribution.
    Children *[]string `json:"children,omitempty"`
    // Contribution associated with this node.
    Contribution *ClientContribution `json:"contribution,omitempty"`
    // List of ids for contributions which are parents to the current contribution.
    Parents *[]string `json:"parents,omitempty"`
}

type ClientContributionProviderDetails struct {
    // Friendly name for the provider.
    DisplayName *string `json:"displayName,omitempty"`
    // Unique identifier for this provider. The provider name can be used to cache the contribution data and refer back to it when looking for changes
    Name *string `json:"name,omitempty"`
    // Properties associated with the provider
    Properties *map[string]string `json:"properties,omitempty"`
    // Version of contributions associated with this contribution provider.
    Version *string `json:"version,omitempty"`
}

// A client data provider are the details needed to make the data provider request from the client.
type ClientDataProviderQuery struct {
    // Contextual information to pass to the data providers
    Context *DataProviderContext `json:"context,omitempty"`
    // The contribution ids of the data providers to resolve
    ContributionIds *[]string `json:"contributionIds,omitempty"`
    // The Id of the service instance type that should be communicated with in order to resolve the data providers from the client given the query values.
    QueryServiceInstanceType *uuid.UUID `json:"queryServiceInstanceType,omitempty"`
}

// An individual contribution made by an extension
type Contribution struct {
    // Description of the contribution/type
    Description *string `json:"description,omitempty"`
    // Fully qualified identifier of the contribution/type
    Id *string `json:"id,omitempty"`
    // VisibleTo can be used to restrict whom can reference a given contribution/type. This value should be a list of publishers or extensions access is restricted too.  Examples: "ms" - Means only the "ms" publisher can reference this. "ms.vss-web" - Means only the "vss-web" extension from the "ms" publisher can reference this.
    VisibleTo *[]string `json:"visibleTo,omitempty"`
    // List of constraints (filters) that should be applied to the availability of this contribution
    Constraints *[]ContributionConstraint `json:"constraints,omitempty"`
    // Includes is a set of contributions that should have this contribution included in their targets list.
    Includes *[]string `json:"includes,omitempty"`
    // Properties/attributes of this contribution
    Properties *interface{} `json:"properties,omitempty"`
    // List of demanded claims in order for the user to see this contribution (like anonymous, public, member...).
    RestrictedTo *[]string `json:"restrictedTo,omitempty"`
    // The ids of the contribution(s) that this contribution targets. (parent contributions)
    Targets *[]string `json:"targets,omitempty"`
    // Id of the Contribution Type
    Type *string `json:"type,omitempty"`
}

// Base class shared by contributions and contribution types
type ContributionBase struct {
    // Description of the contribution/type
    Description *string `json:"description,omitempty"`
    // Fully qualified identifier of the contribution/type
    Id *string `json:"id,omitempty"`
    // VisibleTo can be used to restrict whom can reference a given contribution/type. This value should be a list of publishers or extensions access is restricted too.  Examples: "ms" - Means only the "ms" publisher can reference this. "ms.vss-web" - Means only the "vss-web" extension from the "ms" publisher can reference this.
    VisibleTo *[]string `json:"visibleTo,omitempty"`
}

// Specifies a constraint that can be used to dynamically include/exclude a given contribution
type ContributionConstraint struct {
    // An optional property that can be specified to group constraints together. All constraints within a group are AND'd together (all must be evaluate to True in order for the contribution to be included). Different groups of constraints are OR'd (only one group needs to evaluate to True for the contribution to be included).
    Group *int `json:"group,omitempty"`
    // Fully qualified identifier of a shared constraint
    Id *string `json:"id,omitempty"`
    // If true, negate the result of the filter (include the contribution if the applied filter returns false instead of true)
    Inverse *bool `json:"inverse,omitempty"`
    // Name of the IContributionFilter plugin
    Name *string `json:"name,omitempty"`
    // Properties that are fed to the contribution filter class
    Properties *interface{} `json:"properties,omitempty"`
    // Constraints can be optionally be applied to one or more of the relationships defined in the contribution. If no relationships are defined then all relationships are associated with the constraint. This means the default behaviour will elimiate the contribution from the tree completely if the constraint is applied.
    Relationships *[]string `json:"relationships,omitempty"`
}

// Represents different ways of including contributions based on licensing
type ContributionLicensingBehaviorType string

// A query that can be issued for contribution nodes
type ContributionNodeQuery struct {
    // The contribution ids of the nodes to find.
    ContributionIds *[]string `json:"contributionIds,omitempty"`
    // Contextual information that can be leveraged by contribution constraints
    DataProviderContext *DataProviderContext `json:"dataProviderContext,omitempty"`
    // Indicator if contribution provider details should be included in the result.
    IncludeProviderDetails *bool `json:"includeProviderDetails,omitempty"`
    // Query options tpo be used when fetching ContributionNodes
    QueryOptions *ContributionQueryOptions `json:"queryOptions,omitempty"`
}

// Result of a contribution node query.  Wraps the resulting contribution nodes and provider details.
type ContributionNodeQueryResult struct {
    // Map of contribution ids to corresponding node.
    Nodes *map[string]ClientContributionNode `json:"nodes,omitempty"`
    // Map of provider ids to the corresponding provider details object.
    ProviderDetails *map[string]ClientContributionProviderDetails `json:"providerDetails,omitempty"`
}

// Description about a property of a contribution type
type ContributionPropertyDescription struct {
    // Description of the property
    Description *string `json:"description,omitempty"`
    // Name of the property
    Name *string `json:"name,omitempty"`
    // True if this property is required
    Required *bool `json:"required,omitempty"`
    // The type of value used for this property
    Type *ContributionPropertyType `json:"type,omitempty"`
}

// The type of value used for a property
type ContributionPropertyType string

// Options that control the contributions to include in a query
type ContributionQueryOptions string

// A contribution type, given by a json schema
type ContributionType struct {
    // Description of the contribution/type
    Description *string `json:"description,omitempty"`
    // Fully qualified identifier of the contribution/type
    Id *string `json:"id,omitempty"`
    // VisibleTo can be used to restrict whom can reference a given contribution/type. This value should be a list of publishers or extensions access is restricted too.  Examples: "ms" - Means only the "ms" publisher can reference this. "ms.vss-web" - Means only the "vss-web" extension from the "ms" publisher can reference this.
    VisibleTo *[]string `json:"visibleTo,omitempty"`
    // Controls whether or not contributions of this type have the type indexed for queries. This allows clients to find all extensions that have a contribution of this type.  NOTE: Only TrustedPartners are allowed to specify indexed contribution types.
    Indexed *bool `json:"indexed,omitempty"`
    // Friendly name of the contribution/type
    Name *string `json:"name,omitempty"`
    // Describes the allowed properties for this contribution type
    Properties *map[string]ContributionPropertyDescription `json:"properties,omitempty"`
}

// Contextual information that data providers can examine when populating their data
type DataProviderContext struct {
    // Generic property bag that contains context-specific properties that data providers can use when populating their data dictionary
    Properties *map[string]interface{} `json:"properties,omitempty"`
}

type DataProviderExceptionDetails struct {
    // The type of the exception that was thrown.
    ExceptionType *string `json:"exceptionType,omitempty"`
    // Message that is associated with the exception.
    Message *string `json:"message,omitempty"`
    // The StackTrace from the exception turned into a string.
    StackTrace *string `json:"stackTrace,omitempty"`
}

// A query that can be issued for data provider data
type DataProviderQuery struct {
    // Contextual information to pass to the data providers
    Context *DataProviderContext `json:"context,omitempty"`
    // The contribution ids of the data providers to resolve
    ContributionIds *[]string `json:"contributionIds,omitempty"`
}

// Result structure from calls to GetDataProviderData
type DataProviderResult struct {
    // This is the set of data providers that were requested, but either they were defined as client providers, or as remote providers that failed and may be retried by the client.
    ClientProviders *map[string]ClientDataProviderQuery `json:"clientProviders,omitempty"`
    // Property bag of data keyed off of the data provider contribution id
    Data *map[string]interface{} `json:"data,omitempty"`
    // Set of exceptions that occurred resolving the data providers.
    Exceptions *map[string]DataProviderExceptionDetails `json:"exceptions,omitempty"`
    // List of data providers resolved in the data-provider query
    ResolvedProviders *[]ResolvedDataProvider `json:"resolvedProviders,omitempty"`
    // Scope name applied to this data provider result.
    ScopeName *string `json:"scopeName,omitempty"`
    // Scope value applied to this data provider result.
    ScopeValue *string `json:"scopeValue,omitempty"`
    // Property bag of shared data that was contributed to by any of the individual data providers
    SharedData *map[string]interface{} `json:"sharedData,omitempty"`
}

// Base class for an event callback for an extension
type ExtensionEventCallback struct {
    // The uri of the endpoint that is hit when an event occurs
    Uri *string `json:"uri,omitempty"`
}

// Collection of event callbacks - endpoints called when particular extension events occur.
type ExtensionEventCallbackCollection struct {
    // Optional.  Defines an endpoint that gets called via a POST request to notify that an extension disable has occurred.
    PostDisable *ExtensionEventCallback `json:"postDisable,omitempty"`
    // Optional.  Defines an endpoint that gets called via a POST request to notify that an extension enable has occurred.
    PostEnable *ExtensionEventCallback `json:"postEnable,omitempty"`
    // Optional.  Defines an endpoint that gets called via a POST request to notify that an extension install has completed.
    PostInstall *ExtensionEventCallback `json:"postInstall,omitempty"`
    // Optional.  Defines an endpoint that gets called via a POST request to notify that an extension uninstall has occurred.
    PostUninstall *ExtensionEventCallback `json:"postUninstall,omitempty"`
    // Optional.  Defines an endpoint that gets called via a POST request to notify that an extension update has occurred.
    PostUpdate *ExtensionEventCallback `json:"postUpdate,omitempty"`
    // Optional.  Defines an endpoint that gets called via a POST request to notify that an extension install is about to occur.  Response indicates whether to proceed or abort.
    PreInstall *ExtensionEventCallback `json:"preInstall,omitempty"`
    // For multi-version extensions, defines an endpoint that gets called via an OPTIONS request to determine the particular version of the extension to be used
    VersionCheck *ExtensionEventCallback `json:"versionCheck,omitempty"`
}

type ExtensionFile struct {
    AssetType *string `json:"assetType,omitempty"`
    Language *string `json:"language,omitempty"`
    Source *string `json:"source,omitempty"`
}

// Set of flags applied to extensions that are relevant to contribution consumers
type ExtensionFlags string

// How an extension should handle including contributions based on licensing
type ExtensionLicensing struct {
    // A list of contributions which deviate from the default licensing behavior
    Overrides *[]LicensingOverride `json:"overrides,omitempty"`
}

// Base class for extension properties which are shared by the extension manifest and the extension model
type ExtensionManifest struct {
    // Uri used as base for other relative uri's defined in extension
    BaseUri *string `json:"baseUri,omitempty"`
    // List of shared constraints defined by this extension
    Constraints *[]ContributionConstraint `json:"constraints,omitempty"`
    // List of contributions made by this extension
    Contributions *[]Contribution `json:"contributions,omitempty"`
    // List of contribution types defined by this extension
    ContributionTypes *[]ContributionType `json:"contributionTypes,omitempty"`
    // List of explicit demands required by this extension
    Demands *[]string `json:"demands,omitempty"`
    // Collection of endpoints that get called when particular extension events occur
    EventCallbacks *ExtensionEventCallbackCollection `json:"eventCallbacks,omitempty"`
    // Secondary location that can be used as base for other relative uri's defined in extension
    FallbackBaseUri *string `json:"fallbackBaseUri,omitempty"`
    // Language Culture Name set by the Gallery
    Language *string `json:"language,omitempty"`
    // How this extension behaves with respect to licensing
    Licensing *ExtensionLicensing `json:"licensing,omitempty"`
    // Version of the extension manifest format/content
    ManifestVersion *float64 `json:"manifestVersion,omitempty"`
    // Default user claims applied to all contributions (except the ones which have been specified restrictedTo explicitly) to control the visibility of a contribution.
    RestrictedTo *[]string `json:"restrictedTo,omitempty"`
    // List of all oauth scopes required by this extension
    Scopes *[]string `json:"scopes,omitempty"`
    // The ServiceInstanceType(Guid) of the VSTS service that must be available to an account in order for the extension to be installed
    ServiceInstanceType *uuid.UUID `json:"serviceInstanceType,omitempty"`
}

// States of an extension Note:  If you add value to this enum, you need to do 2 other things.  First add the back compat enum in value src\Vssf\Sdk\Server\Contributions\InstalledExtensionMessage.cs.  Second, you can not send the new value on the message bus.  You need to remove it from the message bus event prior to being sent.
type ExtensionStateFlags string

// Represents a VSTS extension along with its installation state
type InstalledExtension struct {
    // Uri used as base for other relative uri's defined in extension
    BaseUri *string `json:"baseUri,omitempty"`
    // List of shared constraints defined by this extension
    Constraints *[]ContributionConstraint `json:"constraints,omitempty"`
    // List of contributions made by this extension
    Contributions *[]Contribution `json:"contributions,omitempty"`
    // List of contribution types defined by this extension
    ContributionTypes *[]ContributionType `json:"contributionTypes,omitempty"`
    // List of explicit demands required by this extension
    Demands *[]string `json:"demands,omitempty"`
    // Collection of endpoints that get called when particular extension events occur
    EventCallbacks *ExtensionEventCallbackCollection `json:"eventCallbacks,omitempty"`
    // Secondary location that can be used as base for other relative uri's defined in extension
    FallbackBaseUri *string `json:"fallbackBaseUri,omitempty"`
    // Language Culture Name set by the Gallery
    Language *string `json:"language,omitempty"`
    // How this extension behaves with respect to licensing
    Licensing *ExtensionLicensing `json:"licensing,omitempty"`
    // Version of the extension manifest format/content
    ManifestVersion *float64 `json:"manifestVersion,omitempty"`
    // Default user claims applied to all contributions (except the ones which have been specified restrictedTo explicitly) to control the visibility of a contribution.
    RestrictedTo *[]string `json:"restrictedTo,omitempty"`
    // List of all oauth scopes required by this extension
    Scopes *[]string `json:"scopes,omitempty"`
    // The ServiceInstanceType(Guid) of the VSTS service that must be available to an account in order for the extension to be installed
    ServiceInstanceType *uuid.UUID `json:"serviceInstanceType,omitempty"`
    // The friendly extension id for this extension - unique for a given publisher.
    ExtensionId *string `json:"extensionId,omitempty"`
    // The display name of the extension.
    ExtensionName *string `json:"extensionName,omitempty"`
    // This is the set of files available from the extension.
    Files *[]ExtensionFile `json:"files,omitempty"`
    // Extension flags relevant to contribution consumers
    Flags *ExtensionFlags `json:"flags,omitempty"`
    // Information about this particular installation of the extension
    InstallState *InstalledExtensionState `json:"installState,omitempty"`
    // This represents the date/time the extensions was last updated in the gallery. This doesnt mean this version was updated the value represents changes to any and all versions of the extension.
    LastPublished *time.Time `json:"lastPublished,omitempty"`
    // Unique id of the publisher of this extension
    PublisherId *string `json:"publisherId,omitempty"`
    // The display name of the publisher
    PublisherName *string `json:"publisherName,omitempty"`
    // Unique id for this extension (the same id is used for all versions of a single extension)
    RegistrationId *uuid.UUID `json:"registrationId,omitempty"`
    // Version of this extension
    Version *string `json:"version,omitempty"`
}

// The state of an installed extension
type InstalledExtensionState struct {
    // States of an installed extension
    Flags *ExtensionStateFlags `json:"flags,omitempty"`
    // List of installation issues
    InstallationIssues *[]InstalledExtensionStateIssue `json:"installationIssues,omitempty"`
    // The time at which this installation was last updated
    LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// Represents an installation issue
type InstalledExtensionStateIssue struct {
    // The error message
    Message *string `json:"message,omitempty"`
    // Source of the installation issue, for example  "Demands"
    Source *string `json:"source,omitempty"`
    // Installation issue type (Warning, Error)
    Type *InstalledExtensionStateIssueType `json:"type,omitempty"`
}

// Installation issue type (Warning, Error)
type InstalledExtensionStateIssueType string

// Maps a contribution to a licensing behavior
type LicensingOverride struct {
    // How the inclusion of this contribution should change based on licensing
    Behavior *ContributionLicensingBehaviorType `json:"behavior,omitempty"`
    // Fully qualified contribution id which we want to define licensing behavior for
    Id *string `json:"id,omitempty"`
}

// Entry for a specific data provider's resulting data
type ResolvedDataProvider struct {
    // The total time the data provider took to resolve its data (in milliseconds)
    Duration *int `json:"duration,omitempty"`
    Error *string `json:"error,omitempty"`
    Id *string `json:"id,omitempty"`
}
