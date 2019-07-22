// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tfvc

import (
    "github.com/google/uuid"
    "time"
)

type AssociatedWorkItem struct {
    AssignedTo *string `json:"assignedTo,omitempty"`
    // Id of associated the work item.
    Id *int `json:"id,omitempty"`
    State *string `json:"state,omitempty"`
    Title *string `json:"title,omitempty"`
    // REST Url of the work item.
    Url *string `json:"url,omitempty"`
    WebUrl *string `json:"webUrl,omitempty"`
    WorkItemType *string `json:"workItemType,omitempty"`
}

type Change struct {
    // The type of change that was made to the item.
    ChangeType *VersionControlChangeType `json:"changeType,omitempty"`
    // Current version.
    Item *interface{} `json:"item,omitempty"`
    // Content of the item after the change.
    NewContent *ItemContent `json:"newContent,omitempty"`
    // Path of the item on the server.
    SourceServerItem *string `json:"sourceServerItem,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
}

type CheckinNote struct {
    Name *string `json:"name,omitempty"`
    Value *string `json:"value,omitempty"`
}

type FileContentMetadata struct {
    ContentType *string `json:"contentType,omitempty"`
    Encoding *int `json:"encoding,omitempty"`
    Extension *string `json:"extension,omitempty"`
    FileName *string `json:"fileName,omitempty"`
    IsBinary *bool `json:"isBinary,omitempty"`
    IsImage *bool `json:"isImage,omitempty"`
    VsLink *string `json:"vsLink,omitempty"`
}

type GitRepository struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    DefaultBranch *string `json:"defaultBranch,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    // True if the repository was created as a fork
    IsFork *bool `json:"isFork,omitempty"`
    Name *string `json:"name,omitempty"`
    ParentRepository *GitRepositoryRef `json:"parentRepository,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
    RemoteUrl *string `json:"remoteUrl,omitempty"`
    // Compressed size (bytes) of the repository.
    Size *uint64 `json:"size,omitempty"`
    SshUrl *string `json:"sshUrl,omitempty"`
    Url *string `json:"url,omitempty"`
    ValidRemoteUrls *[]string `json:"validRemoteUrls,omitempty"`
    WebUrl *string `json:"webUrl,omitempty"`
}

type GitRepositoryRef struct {
    // Team Project Collection where this Fork resides
    Collection *TeamProjectCollectionReference `json:"collection,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    // True if the repository was created as a fork
    IsFork *bool `json:"isFork,omitempty"`
    Name *string `json:"name,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
    RemoteUrl *string `json:"remoteUrl,omitempty"`
    SshUrl *string `json:"sshUrl,omitempty"`
    Url *string `json:"url,omitempty"`
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

type ItemContent struct {
    Content *string `json:"content,omitempty"`
    ContentType *ItemContentType `json:"contentType,omitempty"`
}

type ItemContentType string

type itemContentTypeValuesType struct {
    RawText ItemContentType
    Base64Encoded ItemContentType
}

var ItemContentTypeValues = itemContentTypeValuesType{
    RawText: "rawText",
    Base64Encoded: "base64Encoded",
}

type ItemModel struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Content *string `json:"content,omitempty"`
    ContentMetadata *FileContentMetadata `json:"contentMetadata,omitempty"`
    IsFolder *bool `json:"isFolder,omitempty"`
    IsSymLink *bool `json:"isSymLink,omitempty"`
    Path *string `json:"path,omitempty"`
    Url *string `json:"url,omitempty"`
}

type ProjectState string

type projectStateValuesType struct {
    Deleting ProjectState
    New ProjectState
    WellFormed ProjectState
    CreatePending ProjectState
    All ProjectState
    Unchanged ProjectState
    Deleted ProjectState
}

var ProjectStateValues = projectStateValuesType{
    // Project is in the process of being deleted.
    Deleting: "deleting",
    // Project is in the process of being created.
    New: "new",
    // Project is completely created and ready to use.
    WellFormed: "wellFormed",
    // Project has been queued for creation, but the process has not yet started.
    CreatePending: "createPending",
    // All projects regardless of state.
    All: "all",
    // Project has not been changed.
    Unchanged: "unchanged",
    // Project has been deleted.
    Deleted: "deleted",
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
    State *ProjectState `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

type TfvcBranch struct {
    // Path for the branch.
    Path *string `json:"path,omitempty"`
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Creation date of the branch.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Description of the branch.
    Description *string `json:"description,omitempty"`
    // Is the branch deleted?
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Alias or display name of user
    Owner *IdentityRef `json:"owner,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
    // List of children for the branch.
    Children *[]TfvcBranch `json:"children,omitempty"`
    // List of branch mappings.
    Mappings *[]TfvcBranchMapping `json:"mappings,omitempty"`
    // Path of the branch's parent.
    Parent *TfvcShallowBranchRef `json:"parent,omitempty"`
    // List of paths of the related branches.
    RelatedBranches *[]TfvcShallowBranchRef `json:"relatedBranches,omitempty"`
}

type TfvcBranchMapping struct {
    // Depth of the branch.
    Depth *string `json:"depth,omitempty"`
    // Server item for the branch.
    ServerItem *string `json:"serverItem,omitempty"`
    // Type of the branch.
    Type *string `json:"type,omitempty"`
}

type TfvcBranchRef struct {
    // Path for the branch.
    Path *string `json:"path,omitempty"`
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Creation date of the branch.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Description of the branch.
    Description *string `json:"description,omitempty"`
    // Is the branch deleted?
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Alias or display name of user
    Owner *IdentityRef `json:"owner,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
}

type TfvcChange struct {
    // List of merge sources in case of rename or branch creation.
    MergeSources *[]TfvcMergeSource `json:"mergeSources,omitempty"`
    // Version at which a (shelved) change was pended against
    PendingVersion *int `json:"pendingVersion,omitempty"`
}

type TfvcChangeset struct {
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Alias or display name of user
    Author *IdentityRef `json:"author,omitempty"`
    // Id of the changeset.
    ChangesetId *int `json:"changesetId,omitempty"`
    // Alias or display name of user
    CheckedInBy *IdentityRef `json:"checkedInBy,omitempty"`
    // Comment for the changeset.
    Comment *string `json:"comment,omitempty"`
    // Was the Comment result truncated?
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    // Creation date of the changeset.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
    // Account Id of the changeset.
    AccountId *uuid.UUID `json:"accountId,omitempty"`
    // List of associated changes.
    Changes *[]TfvcChange `json:"changes,omitempty"`
    // Checkin Notes for the changeset.
    CheckinNotes *[]CheckinNote `json:"checkinNotes,omitempty"`
    // Collection Id of the changeset.
    CollectionId *uuid.UUID `json:"collectionId,omitempty"`
    // Are more changes available.
    HasMoreChanges *bool `json:"hasMoreChanges,omitempty"`
    // Policy Override for the changeset.
    PolicyOverride *TfvcPolicyOverrideInfo `json:"policyOverride,omitempty"`
    // Team Project Ids for the changeset.
    TeamProjectIds *[]uuid.UUID `json:"teamProjectIds,omitempty"`
    // List of work items associated with the changeset.
    WorkItems *[]AssociatedWorkItem `json:"workItems,omitempty"`
}

type TfvcChangesetRef struct {
    // A collection of REST reference links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Alias or display name of user
    Author *IdentityRef `json:"author,omitempty"`
    // Id of the changeset.
    ChangesetId *int `json:"changesetId,omitempty"`
    // Alias or display name of user
    CheckedInBy *IdentityRef `json:"checkedInBy,omitempty"`
    // Comment for the changeset.
    Comment *string `json:"comment,omitempty"`
    // Was the Comment result truncated?
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    // Creation date of the changeset.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // URL to retrieve the item.
    Url *string `json:"url,omitempty"`
}

// Criteria used in a search for change lists
type TfvcChangesetSearchCriteria struct {
    // Alias or display name of user who made the changes
    Author *string `json:"author,omitempty"`
    // Whether or not to follow renames for the given item being queried
    FollowRenames *bool `json:"followRenames,omitempty"`
    // If provided, only include changesets created after this date (string) Think of a better name for this.
    FromDate *string `json:"fromDate,omitempty"`
    // If provided, only include changesets after this changesetID
    FromId *int `json:"fromId,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    // Path of item to search under
    ItemPath *string `json:"itemPath,omitempty"`
    Mappings *[]TfvcMappingFilter `json:"mappings,omitempty"`
    // If provided, only include changesets created before this date (string) Think of a better name for this.
    ToDate *string `json:"toDate,omitempty"`
    // If provided, a version descriptor for the latest change list to include
    ToId *int `json:"toId,omitempty"`
}

type TfvcChangesetsRequestData struct {
    // List of changeset Ids.
    ChangesetIds *[]int `json:"changesetIds,omitempty"`
    // Length of the comment.
    CommentLength *int `json:"commentLength,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
}

type TfvcItem struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Content *string `json:"content,omitempty"`
    ContentMetadata *FileContentMetadata `json:"contentMetadata,omitempty"`
    IsFolder *bool `json:"isFolder,omitempty"`
    IsSymLink *bool `json:"isSymLink,omitempty"`
    Path *string `json:"path,omitempty"`
    Url *string `json:"url,omitempty"`
    ChangeDate *time.Time `json:"changeDate,omitempty"`
    DeletionId *int `json:"deletionId,omitempty"`
    // File encoding from database, -1 represents binary.
    Encoding *int `json:"encoding,omitempty"`
    // MD5 hash as a base 64 string, applies to files only.
    HashValue *string `json:"hashValue,omitempty"`
    IsBranch *bool `json:"isBranch,omitempty"`
    IsPendingChange *bool `json:"isPendingChange,omitempty"`
    // The size of the file, if applicable.
    Size *uint64 `json:"size,omitempty"`
    Version *int `json:"version,omitempty"`
}

// Item path and Version descriptor properties
type TfvcItemDescriptor struct {
    Path *string `json:"path,omitempty"`
    RecursionLevel *VersionControlRecursionType `json:"recursionLevel,omitempty"`
    Version *string `json:"version,omitempty"`
    VersionOption *TfvcVersionOption `json:"versionOption,omitempty"`
    VersionType *TfvcVersionType `json:"versionType,omitempty"`
}

type TfvcItemRequestData struct {
    // If true, include metadata about the file type
    IncludeContentMetadata *bool `json:"includeContentMetadata,omitempty"`
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    ItemDescriptors *[]TfvcItemDescriptor `json:"itemDescriptors,omitempty"`
}

type TfvcLabel struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Description *string `json:"description,omitempty"`
    Id *int `json:"id,omitempty"`
    LabelScope *string `json:"labelScope,omitempty"`
    ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
    Items *[]TfvcItem `json:"items,omitempty"`
}

type TfvcLabelRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Description *string `json:"description,omitempty"`
    Id *int `json:"id,omitempty"`
    LabelScope *string `json:"labelScope,omitempty"`
    ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
}

type TfvcLabelRequestData struct {
    // Whether to include the _links field on the shallow references
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    ItemLabelFilter *string `json:"itemLabelFilter,omitempty"`
    LabelScope *string `json:"labelScope,omitempty"`
    MaxItemCount *int `json:"maxItemCount,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *string `json:"owner,omitempty"`
}

type TfvcMappingFilter struct {
    Exclude *bool `json:"exclude,omitempty"`
    ServerPath *string `json:"serverPath,omitempty"`
}

type TfvcMergeSource struct {
    // Indicates if this a rename source. If false, it is a merge source.
    IsRename *bool `json:"isRename,omitempty"`
    // The server item of the merge source
    ServerItem *string `json:"serverItem,omitempty"`
    // Start of the version range
    VersionFrom *int `json:"versionFrom,omitempty"`
    // End of the version range
    VersionTo *int `json:"versionTo,omitempty"`
}

type TfvcPolicyFailureInfo struct {
    Message *string `json:"message,omitempty"`
    PolicyName *string `json:"policyName,omitempty"`
}

type TfvcPolicyOverrideInfo struct {
    Comment *string `json:"comment,omitempty"`
    PolicyFailures *[]TfvcPolicyFailureInfo `json:"policyFailures,omitempty"`
}

type TfvcShallowBranchRef struct {
    // Path for the branch.
    Path *string `json:"path,omitempty"`
}

// This is the deep shelveset class
type TfvcShelveset struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Comment *string `json:"comment,omitempty"`
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    Id *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
    Changes *[]TfvcChange `json:"changes,omitempty"`
    Notes *[]CheckinNote `json:"notes,omitempty"`
    PolicyOverride *TfvcPolicyOverrideInfo `json:"policyOverride,omitempty"`
    WorkItems *[]AssociatedWorkItem `json:"workItems,omitempty"`
}

// This is the shallow shelveset class
type TfvcShelvesetRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Comment *string `json:"comment,omitempty"`
    CommentTruncated *bool `json:"commentTruncated,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    Id *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Owner *IdentityRef `json:"owner,omitempty"`
    Url *string `json:"url,omitempty"`
}

type TfvcShelvesetRequestData struct {
    // Whether to include policyOverride and notes Only applies when requesting a single deep shelveset
    IncludeDetails *bool `json:"includeDetails,omitempty"`
    // Whether to include the _links field on the shallow references. Does not apply when requesting a single deep shelveset object. Links will always be included in the deep shelveset.
    IncludeLinks *bool `json:"includeLinks,omitempty"`
    // Whether to include workItems
    IncludeWorkItems *bool `json:"includeWorkItems,omitempty"`
    // Max number of changes to include
    MaxChangeCount *int `json:"maxChangeCount,omitempty"`
    // Max length of comment
    MaxCommentLength *int `json:"maxCommentLength,omitempty"`
    // Shelveset's name
    Name *string `json:"name,omitempty"`
    // Owner's ID. Could be a name or a guid.
    Owner *string `json:"owner,omitempty"`
}

type TfvcStatistics struct {
    // Id of the last changeset the stats are based on.
    ChangesetId *int `json:"changesetId,omitempty"`
    // Count of files at the requested scope.
    FileCountTotal *uint64 `json:"fileCountTotal,omitempty"`
}

type TfvcVersionDescriptor struct {
    Version *string `json:"version,omitempty"`
    VersionOption *TfvcVersionOption `json:"versionOption,omitempty"`
    VersionType *TfvcVersionType `json:"versionType,omitempty"`
}

type TfvcVersionOption string

type tfvcVersionOptionValuesType struct {
    None TfvcVersionOption
    Previous TfvcVersionOption
    UseRename TfvcVersionOption
}

var TfvcVersionOptionValues = tfvcVersionOptionValuesType{
    None: "none",
    Previous: "previous",
    UseRename: "useRename",
}

type TfvcVersionType string

type tfvcVersionTypeValuesType struct {
    None TfvcVersionType
    Changeset TfvcVersionType
    Shelveset TfvcVersionType
    Change TfvcVersionType
    Date TfvcVersionType
    Latest TfvcVersionType
    Tip TfvcVersionType
    MergeSource TfvcVersionType
}

var TfvcVersionTypeValues = tfvcVersionTypeValuesType{
    None: "none",
    Changeset: "changeset",
    Shelveset: "shelveset",
    Change: "change",
    Date: "date",
    Latest: "latest",
    Tip: "tip",
    MergeSource: "mergeSource",
}

type VersionControlChangeType string

type versionControlChangeTypeValuesType struct {
    None VersionControlChangeType
    Add VersionControlChangeType
    Edit VersionControlChangeType
    Encoding VersionControlChangeType
    Rename VersionControlChangeType
    Delete VersionControlChangeType
    Undelete VersionControlChangeType
    Branch VersionControlChangeType
    Merge VersionControlChangeType
    Lock VersionControlChangeType
    Rollback VersionControlChangeType
    SourceRename VersionControlChangeType
    TargetRename VersionControlChangeType
    Property VersionControlChangeType
    All VersionControlChangeType
}

var VersionControlChangeTypeValues = versionControlChangeTypeValuesType{
    None: "none",
    Add: "add",
    Edit: "edit",
    Encoding: "encoding",
    Rename: "rename",
    Delete: "delete",
    Undelete: "undelete",
    Branch: "branch",
    Merge: "merge",
    Lock: "lock",
    Rollback: "rollback",
    SourceRename: "sourceRename",
    TargetRename: "targetRename",
    Property: "property",
    All: "all",
}

type VersionControlProjectInfo struct {
    DefaultSourceControlType *SourceControlTypes `json:"defaultSourceControlType,omitempty"`
    Project *TeamProjectReference `json:"project,omitempty"`
    SupportsGit *bool `json:"supportsGit,omitempty"`
    SupportsTFVC *bool `json:"supportsTFVC,omitempty"`
}

type VersionControlRecursionType string

type versionControlRecursionTypeValuesType struct {
    None VersionControlRecursionType
    OneLevel VersionControlRecursionType
    OneLevelPlusNestedEmptyFolders VersionControlRecursionType
    Full VersionControlRecursionType
}

var VersionControlRecursionTypeValues = versionControlRecursionTypeValuesType{
    // Only return the specified item.
    None: "none",
    // Return the specified item and its direct children.
    OneLevel: "oneLevel",
    // Return the specified item and its direct children, as well as recursive chains of nested child folders that only contain a single folder.
    OneLevelPlusNestedEmptyFolders: "oneLevelPlusNestedEmptyFolders",
    // Return specified item and all descendants
    Full: "full",
}
