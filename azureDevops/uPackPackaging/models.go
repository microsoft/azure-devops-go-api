// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package uPackPackaging

import (
    "time"
)

// Describes UPack batch operation types.
type UPackBatchOperationType string

// Describes intent when calling the API GetPackageMetadata.
type UPackGetPackageMetadataIntent string

type UPackLimitedPackageMetadata struct {
    Version *string `json:"version,omitempty"`
}

type UPackLimitedPackageMetadataListResponse struct {
    Count *int `json:"count,omitempty"`
    Value *[]UPackLimitedPackageMetadata `json:"value,omitempty"`
}

type UPackPackageMetadata struct {
    Description *string `json:"description,omitempty"`
    ManifestId *string `json:"manifestId,omitempty"`
    SuperRootId *string `json:"superRootId,omitempty"`
    Version *string `json:"version,omitempty"`
}

type UPackPackagePushMetadata struct {
    Description *string `json:"description,omitempty"`
    ManifestId *string `json:"manifestId,omitempty"`
    SuperRootId *string `json:"superRootId,omitempty"`
    Version *string `json:"version,omitempty"`
    ProofNodes *[]string `json:"proofNodes,omitempty"`
}

// Deletion state of a Universal package.
type UPackPackageVersionDeletionState struct {
    // UTC date the package was deleted.
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    // Name of the package.
    Name *string `json:"name,omitempty"`
    // Version of the package.
    Version *string `json:"version,omitempty"`
}
