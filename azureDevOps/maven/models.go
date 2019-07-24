// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package maven

import (
    "github.com/google/uuid"
    "time"
)

// Do not attempt to use this type to create a new BatchOperationData. This type does not contain sufficient fields to create a new batch operation data.
type BatchOperationData struct {
}

type MavenBatchOperationType string

type mavenBatchOperationTypeValuesType struct {
    Promote MavenBatchOperationType
    Delete MavenBatchOperationType
    PermanentDelete MavenBatchOperationType
    RestoreToFeed MavenBatchOperationType
}

var MavenBatchOperationTypeValues = mavenBatchOperationTypeValuesType{
    // Promote package versions to a release view. If constructing a MavenPackagesBatchRequest object with this type, use BatchPromoteData for its Data property. Not supported in the Recycle Bin.
    Promote: "promote",
    // Delete package versions. Not supported in the Recycle Bin.
    Delete: "delete",
    // Permanently delete package versions. Only supported in the Recycle Bin.
    PermanentDelete: "permanentDelete",
    // Restore unpublished package versions to the feed. Only supported in the Recycle Bin.
    RestoreToFeed: "restoreToFeed",
}

type MavenDistributionManagement struct {
    Repository *MavenRepository `json:"repository,omitempty"`
    SnapshotRepository *MavenSnapshotRepository `json:"snapshotRepository,omitempty"`
}

// Identifies a particular Maven package version
type MavenMinimalPackageDetails struct {
    // Package artifact ID
    Artifact *string `json:"artifact,omitempty"`
    // Package group ID
    Group *string `json:"group,omitempty"`
    // Package version
    Version *string `json:"version,omitempty"`
}

type MavenPackage struct {
    ArtifactId *string `json:"artifactId,omitempty"`
    ArtifactIndex *ReferenceLink `json:"artifactIndex,omitempty"`
    ArtifactMetadata *ReferenceLink `json:"artifactMetadata,omitempty"`
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    Files *ReferenceLinks `json:"files,omitempty"`
    GroupId *string `json:"groupId,omitempty"`
    Pom *MavenPomMetadata `json:"pom,omitempty"`
    RequestedFile *ReferenceLink `json:"requestedFile,omitempty"`
    SnapshotMetadata *ReferenceLink `json:"snapshotMetadata,omitempty"`
    Version *string `json:"version,omitempty"`
    Versions *ReferenceLinks `json:"versions,omitempty"`
    VersionsIndex *ReferenceLink `json:"versionsIndex,omitempty"`
}

// A batch of operations to apply to package versions.
type MavenPackagesBatchRequest struct {
    // Data required to perform the operation. This is optional based on type of operation. Use BatchPromoteData if performing a promote operation.
    Data *BatchOperationData `json:"data,omitempty"`
    // Type of operation that needs to be performed on packages.
    Operation *MavenBatchOperationType `json:"operation,omitempty"`
    // The packages onto which the operation will be performed.
    Packages *[]MavenMinimalPackageDetails `json:"packages,omitempty"`
}

// Deletion state of a maven package.
type MavenPackageVersionDeletionState struct {
    // Artifact Id of the package.
    ArtifactId *string `json:"artifactId,omitempty"`
    // UTC date the package was deleted.
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    // Group Id of the package.
    GroupId *string `json:"groupId,omitempty"`
    // Version of the package.
    Version *string `json:"version,omitempty"`
}

type MavenPomBuild struct {
    Plugins *[]Plugin `json:"plugins,omitempty"`
}

type MavenPomCi struct {
    Notifiers *[]MavenPomCiNotifier `json:"notifiers,omitempty"`
    System *string `json:"system,omitempty"`
    Url *string `json:"url,omitempty"`
}

type MavenPomCiNotifier struct {
    Configuration *[]string `json:"configuration,omitempty"`
    SendOnError *string `json:"sendOnError,omitempty"`
    SendOnFailure *string `json:"sendOnFailure,omitempty"`
    SendOnSuccess *string `json:"sendOnSuccess,omitempty"`
    SendOnWarning *string `json:"sendOnWarning,omitempty"`
    Type *string `json:"type,omitempty"`
}

type MavenPomDependency struct {
    ArtifactId *string `json:"artifactId,omitempty"`
    GroupId *string `json:"groupId,omitempty"`
    Version *string `json:"version,omitempty"`
    Optional *bool `json:"optional,omitempty"`
    Scope *string `json:"scope,omitempty"`
    Type *string `json:"type,omitempty"`
}

type MavenPomDependencyManagement struct {
    Dependencies *[]MavenPomDependency `json:"dependencies,omitempty"`
}

type MavenPomGav struct {
    ArtifactId *string `json:"artifactId,omitempty"`
    GroupId *string `json:"groupId,omitempty"`
    Version *string `json:"version,omitempty"`
}

type MavenPomIssueManagement struct {
    System *string `json:"system,omitempty"`
    Url *string `json:"url,omitempty"`
}

type MavenPomLicense struct {
    Name *string `json:"name,omitempty"`
    Url *string `json:"url,omitempty"`
    Distribution *string `json:"distribution,omitempty"`
}

type MavenPomMailingList struct {
    Archive *string `json:"archive,omitempty"`
    Name *string `json:"name,omitempty"`
    OtherArchives *[]string `json:"otherArchives,omitempty"`
    Post *string `json:"post,omitempty"`
    Subscribe *string `json:"subscribe,omitempty"`
    Unsubscribe *string `json:"unsubscribe,omitempty"`
}

type MavenPomMetadata struct {
    ArtifactId *string `json:"artifactId,omitempty"`
    GroupId *string `json:"groupId,omitempty"`
    Version *string `json:"version,omitempty"`
    Build *MavenPomBuild `json:"build,omitempty"`
    CiManagement *MavenPomCi `json:"ciManagement,omitempty"`
    Contributors *[]MavenPomPerson `json:"contributors,omitempty"`
    Dependencies *[]MavenPomDependency `json:"dependencies,omitempty"`
    DependencyManagement *MavenPomDependencyManagement `json:"dependencyManagement,omitempty"`
    Description *string `json:"description,omitempty"`
    Developers *[]MavenPomPerson `json:"developers,omitempty"`
    DistributionManagement *MavenDistributionManagement `json:"distributionManagement,omitempty"`
    InceptionYear *string `json:"inceptionYear,omitempty"`
    IssueManagement *MavenPomIssueManagement `json:"issueManagement,omitempty"`
    Licenses *[]MavenPomLicense `json:"licenses,omitempty"`
    MailingLists *[]MavenPomMailingList `json:"mailingLists,omitempty"`
    ModelVersion *string `json:"modelVersion,omitempty"`
    Modules *[]string `json:"modules,omitempty"`
    Name *string `json:"name,omitempty"`
    Organization *MavenPomOrganization `json:"organization,omitempty"`
    Packaging *string `json:"packaging,omitempty"`
    Parent *MavenPomParent `json:"parent,omitempty"`
    Prerequisites *map[string]string `json:"prerequisites,omitempty"`
    Properties *map[string]string `json:"properties,omitempty"`
    Scm *MavenPomScm `json:"scm,omitempty"`
    Url *string `json:"url,omitempty"`
}

type MavenPomOrganization struct {
    Name *string `json:"name,omitempty"`
    Url *string `json:"url,omitempty"`
}

type MavenPomParent struct {
    ArtifactId *string `json:"artifactId,omitempty"`
    GroupId *string `json:"groupId,omitempty"`
    Version *string `json:"version,omitempty"`
    RelativePath *string `json:"relativePath,omitempty"`
}

type MavenPomPerson struct {
    Email *string `json:"email,omitempty"`
    Id *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Organization *string `json:"organization,omitempty"`
    OrganizationUrl *string `json:"organizationUrl,omitempty"`
    Roles *[]string `json:"roles,omitempty"`
    Timezone *string `json:"timezone,omitempty"`
    Url *string `json:"url,omitempty"`
}

type MavenPomScm struct {
    Connection *string `json:"connection,omitempty"`
    DeveloperConnection *string `json:"developerConnection,omitempty"`
    Tag *string `json:"tag,omitempty"`
    Url *string `json:"url,omitempty"`
}

type MavenRecycleBinPackageVersionDetails struct {
    // Setting to false will undo earlier deletion and restore the package to feed.
    Deleted *bool `json:"deleted,omitempty"`
}

type MavenRepository struct {
    UniqueVersion *bool `json:"uniqueVersion,omitempty"`
}

type MavenSnapshotRepository struct {
    UniqueVersion *bool `json:"uniqueVersion,omitempty"`
}

// Package version metadata for a Maven package
type Package struct {
    // Related REST links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // If and when the package was deleted.
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    // Package Id.
    Id *string `json:"id,omitempty"`
    // The display name of the package.
    Name *string `json:"name,omitempty"`
    // If and when the package was permanently deleted.
    PermanentlyDeletedDate *time.Time `json:"permanentlyDeletedDate,omitempty"`
    // The history of upstream sources for this package. The first source in the list is the immediate source from which this package was saved.
    SourceChain *[]UpstreamSourceInfo `json:"sourceChain,omitempty"`
    // The version of the package.
    Version *string `json:"version,omitempty"`
}

// Type of an upstream source, such as Public or Internal.
type PackagingSourceType string

type packagingSourceTypeValuesType struct {
    Public PackagingSourceType
    Internal PackagingSourceType
}

var PackagingSourceTypeValues = packagingSourceTypeValuesType{
    // Publicly available source.
    Public: "public",
    // Azure DevOps upstream source.
    Internal: "internal",
}

type Plugin struct {
    ArtifactId *string `json:"artifactId,omitempty"`
    GroupId *string `json:"groupId,omitempty"`
    Version *string `json:"version,omitempty"`
    Configuration *PluginConfiguration `json:"configuration,omitempty"`
}

type PluginConfiguration struct {
    GoalPrefix *string `json:"goalPrefix,omitempty"`
}

// The class to represent a REST reference link.  RFC: http://tools.ietf.org/html/draft-kelly-json-hal-06  The RFC is not fully implemented, additional properties are allowed on the reference link but as of yet we don't have a need for them.
type ReferenceLink struct {
    Href *string `json:"href,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

// Upstream source definition, including its Identity, package type, and other associated information.
type UpstreamSourceInfo struct {
    // Locator for connecting to the upstream source in a user friendly format, that may potentially change over time
    DisplayLocation *string `json:"displayLocation,omitempty"`
    // Identity of the upstream source.
    Id *uuid.UUID `json:"id,omitempty"`
    // Locator for connecting to the upstream source
    Location *string `json:"location,omitempty"`
    // Display name.
    Name *string `json:"name,omitempty"`
    // Source type, such as Public or Internal.
    SourceType *PackagingSourceType `json:"sourceType,omitempty"`
}
