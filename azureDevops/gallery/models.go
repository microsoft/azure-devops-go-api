// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package gallery

import (
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "time"
)

// How the acquisition is assigned
type AcquisitionAssignmentType string

type AcquisitionOperation struct {
    // State of the the AcquisitionOperation for the current user
    OperationState *AcquisitionOperationState `json:"operationState,omitempty"`
    // AcquisitionOperationType: install, request, buy, etc...
    OperationType *AcquisitionOperationType `json:"operationType,omitempty"`
    // Optional reason to justify current state. Typically used with Disallow state.
    Reason *string `json:"reason,omitempty"`
}

type AcquisitionOperationState string

// Set of different types of operations that can be requested.
type AcquisitionOperationType string

// Market item acquisition options (install, buy, etc) for an installation target.
type AcquisitionOptions struct {
    // Default Operation for the ItemId in this target
    DefaultOperation *AcquisitionOperation `json:"defaultOperation,omitempty"`
    // The item id that this options refer to
    ItemId *string `json:"itemId,omitempty"`
    // Operations allowed for the ItemId in this target
    Operations *[]AcquisitionOperation `json:"operations,omitempty"`
    // The target that this options refer to
    Target *string `json:"target,omitempty"`
}

type Answers struct {
    // Gets or sets the vs marketplace extension name
    VSMarketplaceExtensionName *string `json:"vsMarketplaceExtensionName,omitempty"`
    // Gets or sets the vs marketplace publsiher name
    VSMarketplacePublisherName *string `json:"vsMarketplacePublisherName,omitempty"`
}

type AssetDetails struct {
    // Gets or sets the Answers, which contains vs marketplace extension name and publisher name
    Answers *Answers `json:"answers,omitempty"`
    // Gets or sets the VS publisher Id
    PublisherNaturalIdentifier *string `json:"publisherNaturalIdentifier,omitempty"`
}

type AzurePublisher struct {
    AzurePublisherId *string `json:"azurePublisherId,omitempty"`
    PublisherName *string `json:"publisherName,omitempty"`
}

type AzureRestApiRequestModel struct {
    // Gets or sets the Asset details
    AssetDetails *AssetDetails `json:"assetDetails,omitempty"`
    // Gets or sets the asset id
    AssetId *string `json:"assetId,omitempty"`
    // Gets or sets the asset version
    AssetVersion *uint64 `json:"assetVersion,omitempty"`
    // Gets or sets the customer support email
    CustomerSupportEmail *string `json:"customerSupportEmail,omitempty"`
    // Gets or sets the integration contact email
    IntegrationContactEmail *string `json:"integrationContactEmail,omitempty"`
    // Gets or sets the asset version
    Operation *string `json:"operation,omitempty"`
    // Gets or sets the plan identifier if any.
    PlanId *string `json:"planId,omitempty"`
    // Gets or sets the publisher id
    PublisherId *string `json:"publisherId,omitempty"`
    // Gets or sets the resource type
    Type *string `json:"type,omitempty"`
}

type AzureRestApiResponseModel struct {
    // Gets or sets the Asset details
    AssetDetails *AssetDetails `json:"assetDetails,omitempty"`
    // Gets or sets the asset id
    AssetId *string `json:"assetId,omitempty"`
    // Gets or sets the asset version
    AssetVersion *uint64 `json:"assetVersion,omitempty"`
    // Gets or sets the customer support email
    CustomerSupportEmail *string `json:"customerSupportEmail,omitempty"`
    // Gets or sets the integration contact email
    IntegrationContactEmail *string `json:"integrationContactEmail,omitempty"`
    // Gets or sets the asset version
    Operation *string `json:"operation,omitempty"`
    // Gets or sets the plan identifier if any.
    PlanId *string `json:"planId,omitempty"`
    // Gets or sets the publisher id
    PublisherId *string `json:"publisherId,omitempty"`
    // Gets or sets the resource type
    Type *string `json:"type,omitempty"`
    // Gets or sets the Asset operation status
    OperationStatus *RestApiResponseStatusModel `json:"operationStatus,omitempty"`
}

// This is the set of categories in response to the get category query
type CategoriesResult struct {
    Categories *[]ExtensionCategory `json:"categories,omitempty"`
}

// Definition of one title of a category
type CategoryLanguageTitle struct {
    // The language for which the title is applicable
    Lang *string `json:"lang,omitempty"`
    // The language culture id of the lang parameter
    Lcid *int `json:"lcid,omitempty"`
    // Actual title to be shown on the UI
    Title *string `json:"title,omitempty"`
}

// The structure of a Concern Rather than defining a separate data structure having same fields as QnAItem, we are inheriting from the QnAItem.
type Concern struct {
    // Time when the review was first created
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Unique identifier of a QnA item
    Id *uint64 `json:"id,omitempty"`
    // Get status of item
    Status *QnAItemStatus `json:"status,omitempty"`
    // Text description of the QnA item
    Text *string `json:"text,omitempty"`
    // Time when the review was edited/updated
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
    // User details for the item.
    User *UserIdentityRef `json:"user,omitempty"`
    // Category of the concern
    Category *ConcernCategory `json:"category,omitempty"`
}

type ConcernCategory string

// Stores Last Contact Date
type CustomerLastContact struct {
    // account for which customer was last contacted
    Account *string `json:"account,omitempty"`
    // Date on which the custoemr was last contacted
    LastContactDate *time.Time `json:"lastContactDate,omitempty"`
}

type DraftPatchOperation string

type DraftStateType string

type EventCounts struct {
    // Average rating on the day for extension
    AverageRating *int `json:"averageRating,omitempty"`
    // Number of times the extension was bought in hosted scenario (applies only to VSTS extensions)
    BuyCount *int `json:"buyCount,omitempty"`
    // Number of times the extension was bought in connected scenario (applies only to VSTS extensions)
    ConnectedBuyCount *int `json:"connectedBuyCount,omitempty"`
    // Number of times the extension was installed in connected scenario (applies only to VSTS extensions)
    ConnectedInstallCount *int `json:"connectedInstallCount,omitempty"`
    // Number of times the extension was installed
    InstallCount *uint64 `json:"installCount,omitempty"`
    // Number of times the extension was installed as a trial (applies only to VSTS extensions)
    TryCount *int `json:"tryCount,omitempty"`
    // Number of times the extension was uninstalled (applies only to VSTS extensions)
    UninstallCount *int `json:"uninstallCount,omitempty"`
    // Number of times the extension was downloaded (applies to VSTS extensions and VSCode marketplace click installs)
    WebDownloadCount *uint64 `json:"webDownloadCount,omitempty"`
    // Number of detail page views
    WebPageViews *uint64 `json:"webPageViews,omitempty"`
}

// Contract for handling the extension acquisition process
type ExtensionAcquisitionRequest struct {
    // How the item is being assigned
    AssignmentType *AcquisitionAssignmentType `json:"assignmentType,omitempty"`
    // The id of the subscription used for purchase
    BillingId *string `json:"billingId,omitempty"`
    // The marketplace id (publisherName.extensionName) for the item
    ItemId *string `json:"itemId,omitempty"`
    // The type of operation, such as install, request, purchase
    OperationType *AcquisitionOperationType `json:"operationType,omitempty"`
    // Additional properties which can be added to the request.
    Properties *interface{} `json:"properties,omitempty"`
    // How many licenses should be purchased
    Quantity *int `json:"quantity,omitempty"`
    // A list of target guids where the item should be acquired (installed, requested, etc.), such as account id
    Targets *[]string `json:"targets,omitempty"`
}

type ExtensionBadge struct {
    Description *string `json:"description,omitempty"`
    ImgUri *string `json:"imgUri,omitempty"`
    Link *string `json:"link,omitempty"`
}

type ExtensionCategory struct {
    // The name of the products with which this category is associated to.
    AssociatedProducts *[]string `json:"associatedProducts,omitempty"`
    CategoryId *int `json:"categoryId,omitempty"`
    // This is the internal name for a category
    CategoryName *string `json:"categoryName,omitempty"`
    // This parameter is obsolete. Refer to LanguageTitles for langauge specific titles
    Language *string `json:"language,omitempty"`
    // The list of all the titles of this category in various languages
    LanguageTitles *[]CategoryLanguageTitle `json:"languageTitles,omitempty"`
    // This is the internal name of the parent if this is associated with a parent
    ParentCategoryName *string `json:"parentCategoryName,omitempty"`
}

type ExtensionDailyStat struct {
    // Stores the event counts
    Counts *EventCounts `json:"counts,omitempty"`
    // Generic key/value pair to store extended statistics. Used for sending paid extension stats like Upgrade, Downgrade, Cancel trend etc.
    ExtendedStats *map[string]interface{} `json:"extendedStats,omitempty"`
    // Timestamp of this data point
    StatisticDate *time.Time `json:"statisticDate,omitempty"`
    // Version of the extension
    Version *string `json:"version,omitempty"`
}

type ExtensionDailyStats struct {
    // List of extension statistics data points
    DailyStats *[]ExtensionDailyStat `json:"dailyStats,omitempty"`
    // Id of the extension, this will never be sent back to the client. For internal use only.
    ExtensionId *uuid.UUID `json:"extensionId,omitempty"`
    // Name of the extension
    ExtensionName *string `json:"extensionName,omitempty"`
    // Name of the publisher
    PublisherName *string `json:"publisherName,omitempty"`
    // Count of stats
    StatCount *int `json:"statCount,omitempty"`
}

type ExtensionDeploymentTechnology string

type ExtensionDraft struct {
    Assets *[]ExtensionDraftAsset `json:"assets,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    DraftState *DraftStateType `json:"draftState,omitempty"`
    ExtensionName *string `json:"extensionName,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    LastUpdated *time.Time `json:"lastUpdated,omitempty"`
    Payload *ExtensionPayload `json:"payload,omitempty"`
    Product *string `json:"product,omitempty"`
    PublisherName *string `json:"publisherName,omitempty"`
    ValidationErrors *[]azureDevops.KeyValuePair `json:"validationErrors,omitempty"`
    ValidationWarnings *[]azureDevops.KeyValuePair `json:"validationWarnings,omitempty"`
}

type ExtensionDraftAsset struct {
    AssetType *string `json:"assetType,omitempty"`
    Language *string `json:"language,omitempty"`
    Source *string `json:"source,omitempty"`
}

type ExtensionDraftPatch struct {
    ExtensionData *UnpackagedExtensionData `json:"extensionData,omitempty"`
    Operation *DraftPatchOperation `json:"operation,omitempty"`
}

// Stores details of each event
type ExtensionEvent struct {
    // Id which identifies each data point uniquely
    Id *uint64 `json:"id,omitempty"`
    Properties *interface{} `json:"properties,omitempty"`
    // Timestamp of when the event occurred
    StatisticDate *time.Time `json:"statisticDate,omitempty"`
    // Version of the extension
    Version *string `json:"version,omitempty"`
}

// Container object for all extension events. Stores all install and uninstall events related to an extension. The events container is generic so can store data of any type of event. New event types can be added without altering the contract.
type ExtensionEvents struct {
    // Generic container for events data. The dictionary key denotes the type of event and the list contains properties related to that event
    Events *map[string][]ExtensionEvent `json:"events,omitempty"`
    // Id of the extension, this will never be sent back to the client. This field will mainly be used when EMS calls into Gallery REST API to update install/uninstall events for various extensions in one go.
    ExtensionId *uuid.UUID `json:"extensionId,omitempty"`
    // Name of the extension
    ExtensionName *string `json:"extensionName,omitempty"`
    // Name of the publisher
    PublisherName *string `json:"publisherName,omitempty"`
}

type ExtensionFile struct {
    AssetType *string `json:"assetType,omitempty"`
    Language *string `json:"language,omitempty"`
    Source *string `json:"source,omitempty"`
}

// The FilterResult is the set of extensions that matched a particular query filter.
type ExtensionFilterResult struct {
    // This is the set of applications that matched the query filter supplied.
    Extensions *[]PublishedExtension `json:"extensions,omitempty"`
    // The PagingToken is returned from a request when more records exist that match the result than were requested or could be returned. A follow-up query with this paging token can be used to retrieve more results.
    PagingToken *string `json:"pagingToken,omitempty"`
    // This is the additional optional metadata for the given result. E.g. Total count of results which is useful in case of paged results
    ResultMetadata *[]ExtensionFilterResultMetadata `json:"resultMetadata,omitempty"`
}

// ExtensionFilterResultMetadata is one set of metadata for the result e.g. Total count. There can be multiple metadata items for one metadata.
type ExtensionFilterResultMetadata struct {
    // The metadata items for the category
    MetadataItems *[]MetadataItem `json:"metadataItems,omitempty"`
    // Defines the category of metadata items
    MetadataType *string `json:"metadataType,omitempty"`
}

// Represents the component pieces of an extensions fully qualified name, along with the fully qualified name.
type ExtensionIdentifier struct {
    // The ExtensionName component part of the fully qualified ExtensionIdentifier
    ExtensionName *string `json:"extensionName,omitempty"`
    // The PublisherName component part of the fully qualified ExtensionIdentifier
    PublisherName *string `json:"publisherName,omitempty"`
}

// Type of event
type ExtensionLifecycleEventType string

// Package that will be used to create or update a published extension
type ExtensionPackage struct {
    // Base 64 encoded extension package
    ExtensionManifest *string `json:"extensionManifest,omitempty"`
}

type ExtensionPayload struct {
    Description *string `json:"description,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    FileName *string `json:"fileName,omitempty"`
    InstallationTargets *[]InstallationTarget `json:"installationTargets,omitempty"`
    IsPreview *bool `json:"isPreview,omitempty"`
    IsSignedByMicrosoft *bool `json:"isSignedByMicrosoft,omitempty"`
    IsValid *bool `json:"isValid,omitempty"`
    Metadata *[]azureDevops.KeyValuePair `json:"metadata,omitempty"`
    Type *ExtensionDeploymentTechnology `json:"type,omitempty"`
}

// Policy with a set of permissions on extension operations
type ExtensionPolicy struct {
    // Permissions on 'Install' operation
    Install *ExtensionPolicyFlags `json:"install,omitempty"`
    // Permission on 'Request' operation
    Request *ExtensionPolicyFlags `json:"request,omitempty"`
}

// Set of flags that can be associated with a given permission over an extension
type ExtensionPolicyFlags string

// An ExtensionQuery is used to search the gallery for a set of extensions that match one of many filter values.
type ExtensionQuery struct {
    // When retrieving extensions with a query; frequently the caller only needs a small subset of the assets. The caller may specify a list of asset types that should be returned if the extension contains it. All other assets will not be returned.
    AssetTypes *[]string `json:"assetTypes,omitempty"`
    // Each filter is a unique query and will have matching set of extensions returned from the request. Each result will have the same index in the resulting array that the filter had in the incoming query.
    Filters *[]QueryFilter `json:"filters,omitempty"`
    // The Flags are used to determine which set of information the caller would like returned for the matched extensions.
    Flags *ExtensionQueryFlags `json:"flags,omitempty"`
}

// Type of extension filters that are supported in the queries.
type ExtensionQueryFilterType string

// Set of flags used to determine which set of information is retrieved when reading published extensions
type ExtensionQueryFlags string

// This is the set of extensions that matched a supplied query through the filters given.
type ExtensionQueryResult struct {
    // For each filter supplied in the query, a filter result will be returned in the query result.
    Results *[]ExtensionFilterResult `json:"results,omitempty"`
}

type ExtensionShare struct {
    Id *string `json:"id,omitempty"`
    IsOrg *bool `json:"isOrg,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *string `json:"type,omitempty"`
}

type ExtensionStatistic struct {
    StatisticName *string `json:"statisticName,omitempty"`
    Value *float64 `json:"value,omitempty"`
}

type ExtensionStatisticOperation string

type ExtensionStatisticUpdate struct {
    ExtensionName *string `json:"extensionName,omitempty"`
    Operation *ExtensionStatisticOperation `json:"operation,omitempty"`
    PublisherName *string `json:"publisherName,omitempty"`
    Statistic *ExtensionStatistic `json:"statistic,omitempty"`
}

// Stats aggregation type
type ExtensionStatsAggregateType string

type ExtensionVersion struct {
    AssetUri *string `json:"assetUri,omitempty"`
    Badges *[]ExtensionBadge `json:"badges,omitempty"`
    FallbackAssetUri *string `json:"fallbackAssetUri,omitempty"`
    Files *[]ExtensionFile `json:"files,omitempty"`
    Flags *ExtensionVersionFlags `json:"flags,omitempty"`
    LastUpdated *time.Time `json:"lastUpdated,omitempty"`
    Properties *[]azureDevops.KeyValuePair `json:"properties,omitempty"`
    ValidationResultMessage *string `json:"validationResultMessage,omitempty"`
    Version *string `json:"version,omitempty"`
    VersionDescription *string `json:"versionDescription,omitempty"`
}

// Set of flags that can be associated with a given extension version. These flags apply to a specific version of the extension.
type ExtensionVersionFlags string

// One condition in a QueryFilter.
type FilterCriteria struct {
    FilterType *int `json:"filterType,omitempty"`
    // The value used in the match based on the filter type.
    Value *string `json:"value,omitempty"`
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

type InstallationTarget struct {
    Target *string `json:"target,omitempty"`
    TargetVersion *string `json:"targetVersion,omitempty"`
}

// MetadataItem is one value of metadata under a given category of metadata
type MetadataItem struct {
    // The count of the metadata item
    Count *int `json:"count,omitempty"`
    // The name of the metadata item
    Name *string `json:"name,omitempty"`
}

// Information needed for sending mail notification
type NotificationsData struct {
    // Notification data needed
    Data *map[string]interface{} `json:"data,omitempty"`
    // List of users who should get the notification
    Identities *map[string]interface{} `json:"identities,omitempty"`
    // Type of Mail Notification.Can be Qna , review or CustomerContact
    Type *NotificationTemplateType `json:"type,omitempty"`
}

// Type of event
type NotificationTemplateType string

// PagingDirection is used to define which set direction to move the returned result set based on a previous query.
type PagingDirection string

// This is the set of categories in response to the get category query
type ProductCategoriesResult struct {
    Categories *[]ProductCategory `json:"categories,omitempty"`
}

// This is the interface object to be used by Root Categories and Category Tree APIs for Visual Studio Ide.
type ProductCategory struct {
    Children *[]ProductCategory `json:"children,omitempty"`
    // Indicator whether this is a leaf or there are children under this category
    HasChildren *bool `json:"hasChildren,omitempty"`
    // Individual Guid of the Category
    Id *uuid.UUID `json:"id,omitempty"`
    // Category Title in the requested language
    Title *string `json:"title,omitempty"`
}

type PublishedExtension struct {
    Categories *[]string `json:"categories,omitempty"`
    DeploymentType *ExtensionDeploymentTechnology `json:"deploymentType,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    ExtensionId *uuid.UUID `json:"extensionId,omitempty"`
    ExtensionName *string `json:"extensionName,omitempty"`
    Flags *PublishedExtensionFlags `json:"flags,omitempty"`
    InstallationTargets *[]InstallationTarget `json:"installationTargets,omitempty"`
    LastUpdated *time.Time `json:"lastUpdated,omitempty"`
    LongDescription *string `json:"longDescription,omitempty"`
    // Date on which the extension was first uploaded.
    PublishedDate *time.Time `json:"publishedDate,omitempty"`
    Publisher *PublisherFacts `json:"publisher,omitempty"`
    // Date on which the extension first went public.
    ReleaseDate *time.Time `json:"releaseDate,omitempty"`
    SharedWith *[]ExtensionShare `json:"sharedWith,omitempty"`
    ShortDescription *string `json:"shortDescription,omitempty"`
    Statistics *[]ExtensionStatistic `json:"statistics,omitempty"`
    Tags *[]string `json:"tags,omitempty"`
    Versions *[]ExtensionVersion `json:"versions,omitempty"`
}

// Set of flags that can be associated with a given extension. These flags apply to all versions of the extension and not to a specific version.
type PublishedExtensionFlags string

type Publisher struct {
    DisplayName *string `json:"displayName,omitempty"`
    EmailAddress *[]string `json:"emailAddress,omitempty"`
    Extensions *[]PublishedExtension `json:"extensions,omitempty"`
    Flags *PublisherFlags `json:"flags,omitempty"`
    LastUpdated *time.Time `json:"lastUpdated,omitempty"`
    LongDescription *string `json:"longDescription,omitempty"`
    PublisherId *uuid.UUID `json:"publisherId,omitempty"`
    PublisherName *string `json:"publisherName,omitempty"`
    ShortDescription *string `json:"shortDescription,omitempty"`
    State *PublisherState `json:"state,omitempty"`
    Links *ReferenceLinks `json:"_links,omitempty"`
}

// Keeping base class separate since publisher DB model class and publisher contract class share these common properties
type PublisherBase struct {
    DisplayName *string `json:"displayName,omitempty"`
    EmailAddress *[]string `json:"emailAddress,omitempty"`
    Extensions *[]PublishedExtension `json:"extensions,omitempty"`
    Flags *PublisherFlags `json:"flags,omitempty"`
    LastUpdated *time.Time `json:"lastUpdated,omitempty"`
    LongDescription *string `json:"longDescription,omitempty"`
    PublisherId *uuid.UUID `json:"publisherId,omitempty"`
    PublisherName *string `json:"publisherName,omitempty"`
    ShortDescription *string `json:"shortDescription,omitempty"`
    State *PublisherState `json:"state,omitempty"`
}

// High-level information about the publisher, like id's and names
type PublisherFacts struct {
    DisplayName *string `json:"displayName,omitempty"`
    Flags *PublisherFlags `json:"flags,omitempty"`
    PublisherId *uuid.UUID `json:"publisherId,omitempty"`
    PublisherName *string `json:"publisherName,omitempty"`
}

// The FilterResult is the set of publishers that matched a particular query filter.
type PublisherFilterResult struct {
    // This is the set of applications that matched the query filter supplied.
    Publishers *[]Publisher `json:"publishers,omitempty"`
}

type PublisherFlags string

type PublisherPermissions string

// An PublisherQuery is used to search the gallery for a set of publishers that match one of many filter values.
type PublisherQuery struct {
    // Each filter is a unique query and will have matching set of publishers returned from the request. Each result will have the same index in the resulting array that the filter had in the incoming query.
    Filters *[]QueryFilter `json:"filters,omitempty"`
    // The Flags are used to determine which set of information the caller would like returned for the matched publishers.
    Flags *PublisherQueryFlags `json:"flags,omitempty"`
}

// Set of flags used to define the attributes requested when a publisher is returned. Some API's allow the caller to specify the level of detail needed.
type PublisherQueryFlags string

// This is the set of publishers that matched a supplied query through the filters given.
type PublisherQueryResult struct {
    // For each filter supplied in the query, a filter result will be returned in the query result.
    Results *[]PublisherFilterResult `json:"results,omitempty"`
}

// Access definition for a RoleAssignment.
type PublisherRoleAccess string

type PublisherRoleAssignment struct {
    // Designates the role as explicitly assigned or inherited.
    Access *PublisherRoleAccess `json:"access,omitempty"`
    // User friendly description of access assignment.
    AccessDisplayName *string `json:"accessDisplayName,omitempty"`
    // The user to whom the role is assigned.
    Identity *IdentityRef `json:"identity,omitempty"`
    // The role assigned to the user.
    Role *PublisherSecurityRole `json:"role,omitempty"`
}

type PublisherSecurityRole struct {
    // Permissions the role is allowed.
    AllowPermissions *int `json:"allowPermissions,omitempty"`
    // Permissions the role is denied.
    DenyPermissions *int `json:"denyPermissions,omitempty"`
    // Description of user access defined by the role
    Description *string `json:"description,omitempty"`
    // User friendly name of the role.
    DisplayName *string `json:"displayName,omitempty"`
    // Globally unique identifier for the role.
    Identifier *string `json:"identifier,omitempty"`
    // Unique name of the role in the scope.
    Name *string `json:"name,omitempty"`
    // Returns the id of the ParentScope.
    Scope *string `json:"scope,omitempty"`
}

type PublisherState string

type PublisherUserRoleAssignmentRef struct {
    // The name of the role assigned.
    RoleName *string `json:"roleName,omitempty"`
    // Identifier of the user given the role assignment.
    UniqueName *string `json:"uniqueName,omitempty"`
    // Unique id of the user given the role assignment.
    UserId *uuid.UUID `json:"userId,omitempty"`
}

// The core structure of a QnA item
type QnAItem struct {
    // Time when the review was first created
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Unique identifier of a QnA item
    Id *uint64 `json:"id,omitempty"`
    // Get status of item
    Status *QnAItemStatus `json:"status,omitempty"`
    // Text description of the QnA item
    Text *string `json:"text,omitempty"`
    // Time when the review was edited/updated
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
    // User details for the item.
    User *UserIdentityRef `json:"user,omitempty"`
}

// Denotes the status of the QnA Item
type QnAItemStatus string

// A filter used to define a set of extensions to return during a query.
type QueryFilter struct {
    // The filter values define the set of values in this query. They are applied based on the QueryFilterType.
    Criteria *[]FilterCriteria `json:"criteria,omitempty"`
    // The PagingDirection is applied to a paging token if one exists. If not the direction is ignored, and Forward from the start of the resultset is used. Direction should be left out of the request unless a paging token is used to help prevent future issues.
    Direction *PagingDirection `json:"direction,omitempty"`
    // The page number requested by the user. If not provided 1 is assumed by default.
    PageNumber *int `json:"pageNumber,omitempty"`
    // The page size defines the number of results the caller wants for this filter. The count can't exceed the overall query size limits.
    PageSize *int `json:"pageSize,omitempty"`
    // The paging token is a distinct type of filter and the other filter fields are ignored. The paging token represents the continuation of a previously executed query. The information about where in the result and what fields are being filtered are embeded in the token.
    PagingToken *string `json:"pagingToken,omitempty"`
    // Defines the type of sorting to be applied on the results. The page slice is cut of the sorted results only.
    SortBy *int `json:"sortBy,omitempty"`
    // Defines the order of sorting, 1 for Ascending, 2 for Descending, else default ordering based on the SortBy value
    SortOrder *int `json:"sortOrder,omitempty"`
}

// The structure of the question / thread
type Question struct {
    // Time when the review was first created
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Unique identifier of a QnA item
    Id *uint64 `json:"id,omitempty"`
    // Get status of item
    Status *QnAItemStatus `json:"status,omitempty"`
    // Text description of the QnA item
    Text *string `json:"text,omitempty"`
    // Time when the review was edited/updated
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
    // User details for the item.
    User *UserIdentityRef `json:"user,omitempty"`
    // List of answers in for the question / thread
    Responses *[]Response `json:"responses,omitempty"`
}

type QuestionsResult struct {
    // Flag indicating if there are more QnA threads to be shown (for paging)
    HasMoreQuestions *bool `json:"hasMoreQuestions,omitempty"`
    // List of the QnA threads
    Questions *[]Question `json:"questions,omitempty"`
}

type RatingCountPerRating struct {
    // Rating value
    Rating *byte `json:"rating,omitempty"`
    // Count of total ratings
    RatingCount *uint64 `json:"ratingCount,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

// The structure of a response
type Response struct {
    // Time when the review was first created
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Unique identifier of a QnA item
    Id *uint64 `json:"id,omitempty"`
    // Get status of item
    Status *QnAItemStatus `json:"status,omitempty"`
    // Text description of the QnA item
    Text *string `json:"text,omitempty"`
    // Time when the review was edited/updated
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
    // User details for the item.
    User *UserIdentityRef `json:"user,omitempty"`
}

// The status of a REST Api response status.
type RestApiResponseStatus string

// REST Api Response
type RestApiResponseStatusModel struct {
    // Gets or sets the operation details
    OperationDetails *interface{} `json:"operationDetails,omitempty"`
    // Gets or sets the operation id
    OperationId *string `json:"operationId,omitempty"`
    // Gets or sets the completed status percentage
    PercentageCompleted *int `json:"percentageCompleted,omitempty"`
    // Gets or sets the status
    Status *RestApiResponseStatus `json:"status,omitempty"`
    // Gets or sets the status message
    StatusMessage *string `json:"statusMessage,omitempty"`
}

type Review struct {
    // Admin Reply, if any, for this review
    AdminReply *ReviewReply `json:"adminReply,omitempty"`
    // Unique identifier of a review item
    Id *uint64 `json:"id,omitempty"`
    // Flag for soft deletion
    IsDeleted *bool `json:"isDeleted,omitempty"`
    IsIgnored *bool `json:"isIgnored,omitempty"`
    // Version of the product for which review was submitted
    ProductVersion *string `json:"productVersion,omitempty"`
    // Rating procided by the user
    Rating *byte `json:"rating,omitempty"`
    // Reply, if any, for this review
    Reply *ReviewReply `json:"reply,omitempty"`
    // Text description of the review
    Text *string `json:"text,omitempty"`
    // Title of the review
    Title *string `json:"title,omitempty"`
    // Time when the review was edited/updated
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
    // Name of the user
    UserDisplayName *string `json:"userDisplayName,omitempty"`
    // Id of the user who submitted the review
    UserId *uuid.UUID `json:"userId,omitempty"`
}

// Type of operation
type ReviewEventOperation string

// Properties associated with Review event
type ReviewEventProperties struct {
    // Operation performed on Event - Create\Update
    EventOperation *ReviewEventOperation `json:"eventOperation,omitempty"`
    // Flag to see if reply is admin reply
    IsAdminReply *bool `json:"isAdminReply,omitempty"`
    // Flag to record if the review is ignored
    IsIgnored *bool `json:"isIgnored,omitempty"`
    // Rating at the time of event
    Rating *int `json:"rating,omitempty"`
    // Reply update date
    ReplyDate *time.Time `json:"replyDate,omitempty"`
    // Publisher reply text or admin reply text
    ReplyText *string `json:"replyText,omitempty"`
    // User who responded to the review
    ReplyUserId *uuid.UUID `json:"replyUserId,omitempty"`
    // Review Event Type - Review
    ResourceType *ReviewResourceType `json:"resourceType,omitempty"`
    // Review update date
    ReviewDate *time.Time `json:"reviewDate,omitempty"`
    // ReviewId of the review  on which the operation is performed
    ReviewId *uint64 `json:"reviewId,omitempty"`
    // Text in Review Text
    ReviewText *string `json:"reviewText,omitempty"`
    // User display name at the time of review
    UserDisplayName *string `json:"userDisplayName,omitempty"`
    // User who gave review
    UserId *uuid.UUID `json:"userId,omitempty"`
}

// Options to GetReviews query
type ReviewFilterOptions string

type ReviewPatch struct {
    // Denotes the patch operation type
    Operation *ReviewPatchOperation `json:"operation,omitempty"`
    // Use when patch operation is FlagReview
    ReportedConcern *UserReportedConcern `json:"reportedConcern,omitempty"`
    // Use when patch operation is EditReview
    ReviewItem *Review `json:"reviewItem,omitempty"`
}

// Denotes the patch operation type
type ReviewPatchOperation string

type ReviewReply struct {
    // Id of the reply
    Id *uint64 `json:"id,omitempty"`
    // Flag for soft deletion
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Version of the product when the reply was submitted or updated
    ProductVersion *string `json:"productVersion,omitempty"`
    // Content of the reply
    ReplyText *string `json:"replyText,omitempty"`
    // Id of the review, to which this reply belongs
    ReviewId *uint64 `json:"reviewId,omitempty"`
    // Title of the reply
    Title *string `json:"title,omitempty"`
    // Date the reply was submitted or updated
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
    // Id of the user who left the reply
    UserId *uuid.UUID `json:"userId,omitempty"`
}

// Type of event
type ReviewResourceType string

type ReviewsResult struct {
    // Flag indicating if there are more reviews to be shown (for paging)
    HasMoreReviews *bool `json:"hasMoreReviews,omitempty"`
    // List of reviews
    Reviews *[]Review `json:"reviews,omitempty"`
    // Count of total review items
    TotalReviewCount *uint64 `json:"totalReviewCount,omitempty"`
}

type ReviewSummary struct {
    // Average Rating
    AverageRating *int `json:"averageRating,omitempty"`
    // Count of total ratings
    RatingCount *uint64 `json:"ratingCount,omitempty"`
    // Split of count across rating
    RatingSplit *[]RatingCountPerRating `json:"ratingSplit,omitempty"`
}

// Defines the sort order that can be defined for Extensions query
type SortByType string

// Defines the sort order that can be defined for Extensions query
type SortOrderType string

type UnpackagedExtensionData struct {
    Categories *[]string `json:"categories,omitempty"`
    Description *string `json:"description,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    DraftId *uuid.UUID `json:"draftId,omitempty"`
    ExtensionName *string `json:"extensionName,omitempty"`
    InstallationTargets *[]InstallationTarget `json:"installationTargets,omitempty"`
    IsConvertedToMarkdown *bool `json:"isConvertedToMarkdown,omitempty"`
    IsPreview *bool `json:"isPreview,omitempty"`
    PricingCategory *string `json:"pricingCategory,omitempty"`
    Product *string `json:"product,omitempty"`
    PublisherName *string `json:"publisherName,omitempty"`
    QnAEnabled *bool `json:"qnAEnabled,omitempty"`
    ReferralUrl *string `json:"referralUrl,omitempty"`
    RepositoryUrl *string `json:"repositoryUrl,omitempty"`
    Tags *[]string `json:"tags,omitempty"`
    Version *string `json:"version,omitempty"`
    VsixId *string `json:"vsixId,omitempty"`
}

// Represents the extension policy applied to a given user
type UserExtensionPolicy struct {
    // User display name that this policy refers to
    DisplayName *string `json:"displayName,omitempty"`
    // The extension policy applied to the user
    Permissions *ExtensionPolicy `json:"permissions,omitempty"`
    // User id that this policy refers to
    UserId *string `json:"userId,omitempty"`
}

// Identity reference with name and guid
type UserIdentityRef struct {
    // User display name
    DisplayName *string `json:"displayName,omitempty"`
    // User VSID
    Id *uuid.UUID `json:"id,omitempty"`
}

type UserReportedConcern struct {
    // Category of the concern
    Category *ConcernCategory `json:"category,omitempty"`
    // User comment associated with the report
    ConcernText *string `json:"concernText,omitempty"`
    // Id of the review which was reported
    ReviewId *uint64 `json:"reviewId,omitempty"`
    // Date the report was submitted
    SubmittedDate *time.Time `json:"submittedDate,omitempty"`
    // Id of the user who reported a review
    UserId *uuid.UUID `json:"userId,omitempty"`
}
