// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package nuGet

import (
    "github.com/google/uuid"
    "time"
)

// Data required to unlist or relist multiple package versions. Pass this while performing NuGetBatchOperationTypes.List batch operation.
type BatchListData struct {
    // The desired listed status for the package versions.
    Listed *bool `json:"listed,omitempty"`
}

// Do not attempt to use this type to create a new BatchOperationData. This type does not contain sufficient fields to create a new batch operation data.
type BatchOperationData struct {
}

// The JSON model for a JSON Patch operation
type JsonPatchOperation struct {
    // The path to copy from for the Move/Copy operation.
    From *string `json:"from,omitempty"`
    // The patch operation
    Op *Operation `json:"op,omitempty"`
    // The path for the operation. In the case of an array, a zero based index can be used to specify the position in the array (e.g. /biscuits/0/name). The "-" character can be used instead of an index to insert at the end of the array (e.g. /biscuits/-).
    Path *string `json:"path,omitempty"`
    // The value for the operation. This is either a primitive or a JToken.
    Value *interface{} `json:"value,omitempty"`
}

// Minimal package details required to identify a package within a protocol.
type MinimalPackageDetails struct {
    // Package name.
    Id *string `json:"id,omitempty"`
    // Package version.
    Version *string `json:"version,omitempty"`
}

// Describes NuGet batch operation types.
type NuGetBatchOperationType string

// A batch of operations to apply to package versions.
type NuGetPackagesBatchRequest struct {
    // Data required to perform the operation. This is optional based on the type of the operation. Use BatchPromoteData if performing a promote operation.
    Data *BatchOperationData `json:"data,omitempty"`
    // Type of operation that needs to be performed on packages.
    Operation *NuGetBatchOperationType `json:"operation,omitempty"`
    // The packages onto which the operation will be performed.
    Packages *[]MinimalPackageDetails `json:"packages,omitempty"`
}

// Deletion state of a NuGet package.
type NuGetPackageVersionDeletionState struct {
    // Utc date the package was deleted.
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    // Name of the package.
    Name *string `json:"name,omitempty"`
    // Version of the package.
    Version *string `json:"version,omitempty"`
}

type NuGetRecycleBinPackageVersionDetails struct {
    // Setting to false will undo earlier deletion and restore the package to feed.
    Deleted *bool `json:"deleted,omitempty"`
}

type Operation string

// Package version metadata for a NuGet package
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

type PackageVersionDetails struct {
    // Indicates the listing state of a package
    Listed *bool `json:"listed,omitempty"`
    // The view to which the package version will be added
    Views *JsonPatchOperation `json:"views,omitempty"`
}

// Type of an upstream source, such as Public or Internal.
type PackagingSourceType string

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
