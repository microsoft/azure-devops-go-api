// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package upackpackaging

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
)

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

type UPackLimitedPackageMetadata struct {
	Version *string `json:"version,omitempty"`
}

type UPackLimitedPackageMetadataListResponse struct {
	Count *int                           `json:"count,omitempty"`
	Value *[]UPackLimitedPackageMetadata `json:"value,omitempty"`
}

type UPackPackageMetadata struct {
	Description *string `json:"description,omitempty"`
	ManifestId  *string `json:"manifestId,omitempty"`
	SuperRootId *string `json:"superRootId,omitempty"`
	Version     *string `json:"version,omitempty"`
}

type UPackPackagePushMetadata struct {
	Description *string   `json:"description,omitempty"`
	ManifestId  *string   `json:"manifestId,omitempty"`
	SuperRootId *string   `json:"superRootId,omitempty"`
	Version     *string   `json:"version,omitempty"`
	ProofNodes  *[]string `json:"proofNodes,omitempty"`
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
