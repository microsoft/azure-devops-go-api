// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package uPackApi

import (
    "time"
)

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

type Operation string

// Package version metadata for a Universal package
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
    // The version of the package.
    Version *string `json:"version,omitempty"`
}

type PackageVersionDetails struct {
    // The view to which the package version will be added
    Views *JsonPatchOperation `json:"views,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

// Describes UPack batch operation types.
type UPackBatchOperationType string

// Describes intent when calling the API GetPackageMetadata.
type UPackGetPackageMetadataIntent string

// A batch of operations to apply to package versions.
type UPackPackagesBatchRequest struct {
    // Data required to perform the operation. This is optional based on the type of the operation. Use BatchPromoteData if performing a promote operation.
    Data *BatchOperationData `json:"data,omitempty"`
    // Type of operation that needs to be performed on packages.
    Operation *UPackBatchOperationType `json:"operation,omitempty"`
    // The packages onto which the operation will be performed.
    Packages *[]MinimalPackageDetails `json:"packages,omitempty"`
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

type UPackRecycleBinPackageVersionDetails struct {
    // Setting to false will undo earlier deletion and restore the package to feed.
    Deleted *bool `json:"deleted,omitempty"`
}
