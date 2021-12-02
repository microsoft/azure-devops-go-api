// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package maven

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/packagingshared"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/webapi"
)

type MavenBatchOperationType string

type mavenBatchOperationTypeValuesType struct {
	Promote         MavenBatchOperationType
	Delete          MavenBatchOperationType
	PermanentDelete MavenBatchOperationType
	RestoreToFeed   MavenBatchOperationType
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
	Repository         *MavenRepository         `json:"repository,omitempty"`
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
	ArtifactId       *string               `json:"artifactId,omitempty"`
	ArtifactIndex    *webapi.ReferenceLink `json:"artifactIndex,omitempty"`
	ArtifactMetadata *webapi.ReferenceLink `json:"artifactMetadata,omitempty"`
	DeletedDate      *azuredevops.Time     `json:"deletedDate,omitempty"`
	Files            interface{}           `json:"files,omitempty"`
	GroupId          *string               `json:"groupId,omitempty"`
	Pom              *MavenPomMetadata     `json:"pom,omitempty"`
	RequestedFile    *webapi.ReferenceLink `json:"requestedFile,omitempty"`
	SnapshotMetadata *webapi.ReferenceLink `json:"snapshotMetadata,omitempty"`
	Version          *string               `json:"version,omitempty"`
	Versions         interface{}           `json:"versions,omitempty"`
	VersionsIndex    *webapi.ReferenceLink `json:"versionsIndex,omitempty"`
}

// A batch of operations to apply to package versions.
type MavenPackagesBatchRequest struct {
	// Data required to perform the operation. This is optional based on type of operation. Use BatchPromoteData if performing a promote operation.
	Data interface{} `json:"data,omitempty"`
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
	DeletedDate *azuredevops.Time `json:"deletedDate,omitempty"`
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
	System    *string               `json:"system,omitempty"`
	Url       *string               `json:"url,omitempty"`
}

type MavenPomCiNotifier struct {
	Configuration *[]string `json:"configuration,omitempty"`
	SendOnError   *string   `json:"sendOnError,omitempty"`
	SendOnFailure *string   `json:"sendOnFailure,omitempty"`
	SendOnSuccess *string   `json:"sendOnSuccess,omitempty"`
	SendOnWarning *string   `json:"sendOnWarning,omitempty"`
	Type          *string   `json:"type,omitempty"`
}

type MavenPomDependency struct {
	ArtifactId *string `json:"artifactId,omitempty"`
	GroupId    *string `json:"groupId,omitempty"`
	Version    *string `json:"version,omitempty"`
	Optional   *bool   `json:"optional,omitempty"`
	Scope      *string `json:"scope,omitempty"`
	Type       *string `json:"type,omitempty"`
}

type MavenPomDependencyManagement struct {
	Dependencies *[]MavenPomDependency `json:"dependencies,omitempty"`
}

type MavenPomGav struct {
	ArtifactId *string `json:"artifactId,omitempty"`
	GroupId    *string `json:"groupId,omitempty"`
	Version    *string `json:"version,omitempty"`
}

type MavenPomIssueManagement struct {
	System *string `json:"system,omitempty"`
	Url    *string `json:"url,omitempty"`
}

type MavenPomLicense struct {
	Name         *string `json:"name,omitempty"`
	Url          *string `json:"url,omitempty"`
	Distribution *string `json:"distribution,omitempty"`
}

type MavenPomMailingList struct {
	Archive       *string   `json:"archive,omitempty"`
	Name          *string   `json:"name,omitempty"`
	OtherArchives *[]string `json:"otherArchives,omitempty"`
	Post          *string   `json:"post,omitempty"`
	Subscribe     *string   `json:"subscribe,omitempty"`
	Unsubscribe   *string   `json:"unsubscribe,omitempty"`
}

type MavenPomMetadata struct {
	ArtifactId             *string                       `json:"artifactId,omitempty"`
	GroupId                *string                       `json:"groupId,omitempty"`
	Version                *string                       `json:"version,omitempty"`
	Build                  *MavenPomBuild                `json:"build,omitempty"`
	CiManagement           *MavenPomCi                   `json:"ciManagement,omitempty"`
	Contributors           *[]MavenPomPerson             `json:"contributors,omitempty"`
	Dependencies           *[]MavenPomDependency         `json:"dependencies,omitempty"`
	DependencyManagement   *MavenPomDependencyManagement `json:"dependencyManagement,omitempty"`
	Description            *string                       `json:"description,omitempty"`
	Developers             *[]MavenPomPerson             `json:"developers,omitempty"`
	DistributionManagement *MavenDistributionManagement  `json:"distributionManagement,omitempty"`
	InceptionYear          *string                       `json:"inceptionYear,omitempty"`
	IssueManagement        *MavenPomIssueManagement      `json:"issueManagement,omitempty"`
	Licenses               *[]MavenPomLicense            `json:"licenses,omitempty"`
	MailingLists           *[]MavenPomMailingList        `json:"mailingLists,omitempty"`
	ModelVersion           *string                       `json:"modelVersion,omitempty"`
	Modules                *[]string                     `json:"modules,omitempty"`
	Name                   *string                       `json:"name,omitempty"`
	Organization           *MavenPomOrganization         `json:"organization,omitempty"`
	Packaging              *string                       `json:"packaging,omitempty"`
	Parent                 *MavenPomParent               `json:"parent,omitempty"`
	Prerequisites          *map[string]string            `json:"prerequisites,omitempty"`
	Properties             *map[string]string            `json:"properties,omitempty"`
	Scm                    *MavenPomScm                  `json:"scm,omitempty"`
	Url                    *string                       `json:"url,omitempty"`
}

type MavenPomOrganization struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}

type MavenPomParent struct {
	ArtifactId   *string `json:"artifactId,omitempty"`
	GroupId      *string `json:"groupId,omitempty"`
	Version      *string `json:"version,omitempty"`
	RelativePath *string `json:"relativePath,omitempty"`
}

type MavenPomPerson struct {
	Email           *string   `json:"email,omitempty"`
	Id              *string   `json:"id,omitempty"`
	Name            *string   `json:"name,omitempty"`
	Organization    *string   `json:"organization,omitempty"`
	OrganizationUrl *string   `json:"organizationUrl,omitempty"`
	Roles           *[]string `json:"roles,omitempty"`
	Timezone        *string   `json:"timezone,omitempty"`
	Url             *string   `json:"url,omitempty"`
}

type MavenPomScm struct {
	Connection          *string `json:"connection,omitempty"`
	DeveloperConnection *string `json:"developerConnection,omitempty"`
	Tag                 *string `json:"tag,omitempty"`
	Url                 *string `json:"url,omitempty"`
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
	Links interface{} `json:"_links,omitempty"`
	// If and when the package was deleted.
	DeletedDate *azuredevops.Time `json:"deletedDate,omitempty"`
	// Package Id.
	Id *string `json:"id,omitempty"`
	// The display name of the package.
	Name *string `json:"name,omitempty"`
	// If and when the package was permanently deleted.
	PermanentlyDeletedDate *azuredevops.Time `json:"permanentlyDeletedDate,omitempty"`
	// The history of upstream sources for this package. The first source in the list is the immediate source from which this package was saved.
	SourceChain *[]packagingshared.UpstreamSourceInfo `json:"sourceChain,omitempty"`
	// The version of the package.
	Version *string `json:"version,omitempty"`
}

type Plugin struct {
	ArtifactId    *string              `json:"artifactId,omitempty"`
	GroupId       *string              `json:"groupId,omitempty"`
	Version       *string              `json:"version,omitempty"`
	Configuration *PluginConfiguration `json:"configuration,omitempty"`
}

type PluginConfiguration struct {
	GoalPrefix *string `json:"goalPrefix,omitempty"`
}
