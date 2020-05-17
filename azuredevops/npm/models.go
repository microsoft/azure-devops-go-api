// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package npm

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/packagingshared"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

// Data required to deprecate multiple package versions. Pass this while performing NpmBatchOperationTypes.Deprecate batch operation.
type BatchDeprecateData struct {
	// Deprecate message that will be added to packages
	Message *string `json:"message,omitempty"`
}

// Describes Npm batch operation types.
type NpmBatchOperationType string

type npmBatchOperationTypeValuesType struct {
	Promote         NpmBatchOperationType
	Deprecate       NpmBatchOperationType
	Unpublish       NpmBatchOperationType
	PermanentDelete NpmBatchOperationType
	RestoreToFeed   NpmBatchOperationType
	Delete          NpmBatchOperationType
}

var NpmBatchOperationTypeValues = npmBatchOperationTypeValuesType{
	// Promote package versions to a release view. If constructing a NpmPackagesBatchRequest object with this type, use BatchPromoteData for its Data property. Not supported in the Recycle Bin.
	Promote: "promote",
	// Deprecate or undeprecate package versions. Not supported in the Recycle Bin.
	Deprecate: "deprecate",
	// Unpublish package versions. Npm-specific alias for the Delete operation. Not supported in the Recycle Bin.
	Unpublish: "unpublish",
	// Permanently delete package versions. Only supported in the Recycle Bin.
	PermanentDelete: "permanentDelete",
	// Restore unpublished package versions to the feed. Only supported in the Recycle Bin.
	RestoreToFeed: "restoreToFeed",
	// Delete package versions (equivalent to Unpublish). Not supported in the Recycle Bin.
	Delete: "delete",
}

// A batch of operations to apply to package versions.
type NpmPackagesBatchRequest struct {
	// Data required to perform the operation. This is optional based on type of operation. Use BatchPromoteData if performing a promote operation.
	Data interface{} `json:"data,omitempty"`
	// Type of operation that needs to be performed on packages.
	Operation *NpmBatchOperationType `json:"operation,omitempty"`
	// The packages onto which the operation will be performed.
	Packages *[]packagingshared.MinimalPackageDetails `json:"packages,omitempty"`
}

// Deletion state of an npm package.
type NpmPackageVersionDeletionState struct {
	// Name of the package.
	Name *string `json:"name,omitempty"`
	// UTC date the package was unpublished.
	UnpublishedDate *azuredevops.Time `json:"unpublishedDate,omitempty"`
	// Version of the package.
	Version *string `json:"version,omitempty"`
}

type NpmRecycleBinPackageVersionDetails struct {
	// Setting to false will undo earlier deletion and restore the package to feed.
	Deleted *bool `json:"deleted,omitempty"`
}

// Package version metadata for an npm package
type Package struct {
	// Related REST links.
	Links interface{} `json:"_links,omitempty"`
	// Deprecated message, if any, for the package.
	DeprecateMessage *string `json:"deprecateMessage,omitempty"`
	// Package Id.
	Id *string `json:"id,omitempty"`
	// The display name of the package.
	Name *string `json:"name,omitempty"`
	// If and when the package was permanently deleted.
	PermanentlyDeletedDate *azuredevops.Time `json:"permanentlyDeletedDate,omitempty"`
	// The history of upstream sources for this package. The first source in the list is the immediate source from which this package was saved.
	SourceChain *[]packagingshared.UpstreamSourceInfo `json:"sourceChain,omitempty"`
	// If and when the package was deleted.
	UnpublishedDate *azuredevops.Time `json:"unpublishedDate,omitempty"`
	// The version of the package.
	Version *string `json:"version,omitempty"`
}

type PackageVersionDetails struct {
	// Indicates the deprecate message of a package version
	DeprecateMessage *string `json:"deprecateMessage,omitempty"`
	// The view to which the package version will be added
	Views *webapi.JsonPatchOperation `json:"views,omitempty"`
}
