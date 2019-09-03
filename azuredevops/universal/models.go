// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package universal

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/packagingshared"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

// Package version metadata for a Universal package
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
	// The version of the package.
	Version *string `json:"version,omitempty"`
}

type PackageVersionDetails struct {
	// The view to which the package version will be added
	Views *webapi.JsonPatchOperation `json:"views,omitempty"`
}

// Describes UPack batch operation types.
type UPackBatchOperationType string

type uPackBatchOperationTypeValuesType struct {
	Promote         UPackBatchOperationType
	Delete          UPackBatchOperationType
	PermanentDelete UPackBatchOperationType
	RestoreToFeed   UPackBatchOperationType
}

var UPackBatchOperationTypeValues = uPackBatchOperationTypeValuesType{
	// Promote package versions to a release view. If constructing a UPackPackagesBatchRequest object with this type, use BatchPromoteData for its Data property. Not supported in the Recycle Bin.
	Promote: "promote",
	// Move package versions to the feed's Recycle Bin. Not supported in the Recycle Bin.
	Delete: "delete",
	// Permanently delete package versions. Only supported in the Recycle Bin.
	PermanentDelete: "permanentDelete",
	// Restore deleted package versions to the feed. Only supported in the Recycle Bin.
	RestoreToFeed: "restoreToFeed",
}

// Describes intent when calling the API GetPackageMetadata.
type UPackGetPackageMetadataIntent string

type uPackGetPackageMetadataIntentValuesType struct {
	FetchMetadataOnly UPackGetPackageMetadataIntent
	Download          UPackGetPackageMetadataIntent
}

var UPackGetPackageMetadataIntentValues = uPackGetPackageMetadataIntentValuesType{
	// Default. The call intends only to retrieve the package metadata.
	FetchMetadataOnly: "fetchMetadataOnly",
	// The call is part of the download flow.
	Download: "download",
}

// A batch of operations to apply to package versions.
type UPackPackagesBatchRequest struct {
	// Data required to perform the operation. This is optional based on the type of the operation. Use BatchPromoteData if performing a promote operation.
	Data interface{} `json:"data,omitempty"`
	// Type of operation that needs to be performed on packages.
	Operation *UPackBatchOperationType `json:"operation,omitempty"`
	// The packages onto which the operation will be performed.
	Packages *[]packagingshared.MinimalPackageDetails `json:"packages,omitempty"`
}

// Deletion state of a Universal package.
type UPackPackageVersionDeletionState struct {
	// UTC date the package was deleted.
	DeletedDate *azuredevops.Time `json:"deletedDate,omitempty"`
	// Name of the package.
	Name *string `json:"name,omitempty"`
	// Version of the package.
	Version *string `json:"version,omitempty"`
}

type UPackRecycleBinPackageVersionDetails struct {
	// Setting to false will undo earlier deletion and restore the package to feed.
	Deleted *bool `json:"deleted,omitempty"`
}
