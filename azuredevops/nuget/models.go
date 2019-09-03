// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package nuget

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/packagingshared"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

// Data required to unlist or relist multiple package versions. Pass this while performing NuGetBatchOperationTypes.List batch operation.
type BatchListData struct {
	// The desired listed status for the package versions.
	Listed *bool `json:"listed,omitempty"`
}

// Describes NuGet batch operation types.
type NuGetBatchOperationType string

type nuGetBatchOperationTypeValuesType struct {
	Promote         NuGetBatchOperationType
	List            NuGetBatchOperationType
	Delete          NuGetBatchOperationType
	PermanentDelete NuGetBatchOperationType
	RestoreToFeed   NuGetBatchOperationType
}

var NuGetBatchOperationTypeValues = nuGetBatchOperationTypeValuesType{
	// Promote package versions to a release view. If constructing a NuGetPackagesBatchRequest object with this type, use BatchPromoteData for its Data property. Not supported in the Recycle Bin.
	Promote: "promote",
	// Unlist or relist package versions. Not supported in the Recycle Bin.
	List: "list",
	// Move package versions to the feed's Recycle Bin. Not supported in the Recycle Bin.
	Delete: "delete",
	// Permanently delete package versions. Only supported in the Recycle Bin.
	PermanentDelete: "permanentDelete",
	// Restore deleted package versions to the feed. Only supported in the Recycle Bin.
	RestoreToFeed: "restoreToFeed",
}

// A batch of operations to apply to package versions.
type NuGetPackagesBatchRequest struct {
	// Data required to perform the operation. This is optional based on the type of the operation. Use BatchPromoteData if performing a promote operation.
	Data interface{} `json:"data,omitempty"`
	// Type of operation that needs to be performed on packages.
	Operation *NuGetBatchOperationType `json:"operation,omitempty"`
	// The packages onto which the operation will be performed.
	Packages *[]packagingshared.MinimalPackageDetails `json:"packages,omitempty"`
}

// Deletion state of a NuGet package.
type NuGetPackageVersionDeletionState struct {
	// Utc date the package was deleted.
	DeletedDate *azuredevops.Time `json:"deletedDate,omitempty"`
	// Name of the package.
	Name *string `json:"name,omitempty"`
	// Version of the package.
	Version *string `json:"version,omitempty"`
}

type NuGetRecycleBinPackageVersionDetails struct {
	// Setting to false will undo earlier deletion and restore the package to feed.
	Deleted *bool `json:"deleted,omitempty"`
}

// Package version metadata for a NuGet package
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

type PackageVersionDetails struct {
	// Indicates the listing state of a package
	Listed *bool `json:"listed,omitempty"`
	// The view to which the package version will be added
	Views *webapi.JsonPatchOperation `json:"views,omitempty"`
}
