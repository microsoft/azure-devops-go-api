// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package packagingshared

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7.1"
)

// Data required to delist or relist multiple package versions. Pass this while performing {protocol}BatchOperationTypes.List batch operation. There is another BatchListData in nuget assembly to maintain serialization compatibility.
type BatchListData struct {
	// The desired listed status for the package versions.
	Listed *bool `json:"listed,omitempty"`
}

// Data required for promoting multiple package versions. Pass this while performing {protocol}BatchOperationTypes.Promote batch operation.
type BatchPromoteData struct {
	// Id or Name of the view, packages need to be promoted to.
	ViewId *string `json:"viewId,omitempty"`
}

type MinimalPackage struct {
	// The display name of the package.
	Name *string `json:"name,omitempty"`
	// Type of the package.
	ProtocolType *string `json:"protocolType,omitempty"`
	// All versions for this package with size information within its feed.
	Versions *[]PackageVersionWithSize `json:"versions,omitempty"`
}

// Minimal package details required to identify a package within a protocol.
type MinimalPackageDetails struct {
	// Package name.
	Id *string `json:"id,omitempty"`
	// Package version.
	Version *string `json:"version,omitempty"`
}

// Describes how a package version instance *originally* entered the Azure Artifacts system, regardless of any intermediate upstreams
type PackageOrigin string

type packageOriginValuesType struct {
	Unknown  PackageOrigin
	External PackageOrigin
	Internal PackageOrigin
}

var PackageOriginValues = packageOriginValuesType{
	// Can't tell where the package came from because the remote scale unit doesn't provide source chains yet.
	Unknown: "unknown",
	// The package was ingested from an external upstream
	External: "external",
	// The package was directly pushed to an Azure Artifacts feed
	Internal: "internal",
}

type PackageVersionWithSize struct {
	// Indicates whether or not the version was soft-deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Display version.
	Version *string `json:"version,omitempty"`
	// Size in bytes for the version.
	VersionSizeInBytes *float64 `json:"versionSizeInBytes,omitempty"`
}

// Type of an upstream source, such as Public or Internal.
type PackagingSourceType string

type packagingSourceTypeValuesType struct {
	Public   PackagingSourceType
	Internal PackagingSourceType
}

var PackagingSourceTypeValues = packagingSourceTypeValuesType{
	// Publicly available source.
	Public: "public",
	// Azure DevOps upstream source.
	Internal: "internal",
}

type ProblemPackage struct {
	// Display name of the problem package.
	PackageName *string `json:"packageName,omitempty"`
	// Version that was blocked.
	PackageVersion *string `json:"packageVersion,omitempty"`
	// Package's protocol.
	Protocol *string `json:"protocol,omitempty"`
	// Reasons the package version was blocked.
	Reasons *[]ProblemPackageReason `json:"reasons,omitempty"`
	// Timestamp when the package was blocked.
	Timestamp *azuredevops.Time `json:"timestamp,omitempty"`
	// From where the package was blocked.
	UpstreamSource *UpstreamSourceInfo `json:"upstreamSource,omitempty"`
}

type ProblemPackageReason struct {
	// Code from Terrapin.
	Code *string `json:"code,omitempty"`
	// Message from Terrapin.
	Message *string `json:"message,omitempty"`
}

// Describes upstreaming behavior for a given feed/protocol/package
type UpstreamingBehavior struct {
	// Indicates whether external upstream versions should be considered for this package
	VersionsFromExternalUpstreams *UpstreamVersionVisibility `json:"versionsFromExternalUpstreams,omitempty"`
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

type UpstreamVersionsData struct {
	// The current <em>actual</em> decision as to whether external versions will be available. If the upstreaming behavior is set to AllowExternalVersions, this will be <c>true</c>. If the behavior is set to Auto, this will be computed based on the current state of the feed and its upstreams.
	ExternalVersionsFromUpstreamAvailable *bool   `json:"externalVersionsFromUpstreamAvailable,omitempty"`
	PackageDisplayName                    *string `json:"packageDisplayName,omitempty"`
	PackageNormalizedName                 *string `json:"packageNormalizedName,omitempty"`
	// The current upstreaming behavior settings for the package in the feed
	UpstreamingBehavior *UpstreamingBehavior `json:"upstreamingBehavior,omitempty"`
	// True if the user has permission to ingest new versions from upstream
	UserHasPermissionToIngestFromUpstream *bool `json:"userHasPermissionToIngestFromUpstream,omitempty"`
	// All versions which are available locally in the feed or in any of its upstreams
	Versions *[]UpstreamVersionsDataVersion `json:"versions,omitempty"`
}

type UpstreamVersionsDataVersion struct {
	DisplayVersion *string `json:"displayVersion,omitempty"`
	// True if the version is actually present in the feed
	IsLocal *bool `json:"isLocal,omitempty"`
	// Data about the local instance actually present in the feed. <c>null</c> if notIsLocal
	LocalInstance     *UpstreamVersionsDataVersionLocalInstance `json:"localInstance,omitempty"`
	NormalizedVersion *string                                   `json:"normalizedVersion,omitempty"`
	// The upstream that has been <em>selected</em> to provide the version, or <c>null</c> if no upstream has been selected (e.g. the package was directly pushed to the feed, or all upstreams that provide the version are hidden. If IsLocal is true, this is the upstream that actually provided the version. If IsLocal is false, this is the upstream that will provide the version on ingestion. This value may change over time as this feed's upstream configuration changes, and as the contents of this feed's upstreams change.
	SelectedUpstreamId *uuid.UUID `json:"selectedUpstreamId,omitempty"`
	// All upstreams this version appears in
	Upstreams *[]UpstreamVersionsDataVersionUpstream `json:"upstreams,omitempty"`
}

type UpstreamVersionsDataVersionLocalInstance struct {
	// Indicates whether the local package version instance has been deleted. Deleted versions cannot be downloaded, but also still cannot be overwritten by saving an instance from upstream.
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// Origin type of the local package version instance, for the purposes of the dependency-confusion vulnerability mitigation. For example, a package that originally came through an upstream to nuget.org will be External even if it has passed through several Azure Artifacts upstreams.
	Origin *PackageOrigin `json:"origin,omitempty"`
	// Source chain of the local version, from the perspective of <em>this feed</em>. Includes, as the first entry, the upstream the local version came from. If the version was directly pushed to the feed, this will be an empty list.
	SourceChain *[]UpstreamSourceInfo `json:"sourceChain,omitempty"`
}

type UpstreamVersionsDataVersionUpstream struct {
	// True if this is the upstream that provided the actually-ingested local version
	IsUpstreamForLocalVersion *bool `json:"isUpstreamForLocalVersion,omitempty"`
	// Origin type of this upstream's instance of the package version, for the purposes of the dependency- confusion vulnerability mitigation. For example, a package that originally came through an upstream to nuget.org will be External even if it has passed through several Azure Artifacts upstreams.
	Origin *PackageOrigin `json:"origin,omitempty"`
	// Source chain of the version in this upstream, from the perspective of the <em>upstream</em>. Does not include this upstream itself.
	SourceChain  *[]UpstreamSourceInfo `json:"sourceChain,omitempty"`
	UpstreamId   *uuid.UUID            `json:"upstreamId,omitempty"`
	UpstreamName *string               `json:"upstreamName,omitempty"`
}

type UpstreamVersionVisibility string

type upstreamVersionVisibilityValuesType struct {
	Auto                  UpstreamVersionVisibility
	AllowExternalVersions UpstreamVersionVisibility
}

var UpstreamVersionVisibilityValues = upstreamVersionVisibilityValuesType{
	Auto:                  "auto",
	AllowExternalVersions: "allowExternalVersions",
}
