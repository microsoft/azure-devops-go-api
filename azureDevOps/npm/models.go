// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package npm

import (
    "github.com/google/uuid"
    "time"
)

// Data required to deprecate multiple package versions. Pass this while performing NpmBatchOperationTypes.Deprecate batch operation.
type BatchDeprecateData struct {
    // Deprecate message that will be added to packages
    Message *string `json:"message,omitempty"`
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
    Value interface{} `json:"value,omitempty"`
}

// Minimal package details required to identify a package within a protocol.
type MinimalPackageDetails struct {
    // Package name.
    Id *string `json:"id,omitempty"`
    // Package version.
    Version *string `json:"version,omitempty"`
}

// Describes Npm batch operation types.
type NpmBatchOperationType string

type npmBatchOperationTypeValuesType struct {
    Promote NpmBatchOperationType
    Deprecate NpmBatchOperationType
    Unpublish NpmBatchOperationType
    PermanentDelete NpmBatchOperationType
    RestoreToFeed NpmBatchOperationType
    Delete NpmBatchOperationType
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
    Data *BatchOperationData `json:"data,omitempty"`
    // Type of operation that needs to be performed on packages.
    Operation *NpmBatchOperationType `json:"operation,omitempty"`
    // The packages onto which the operation will be performed.
    Packages *[]MinimalPackageDetails `json:"packages,omitempty"`
}

// Deletion state of an npm package.
type NpmPackageVersionDeletionState struct {
    // Name of the package.
    Name *string `json:"name,omitempty"`
    // UTC date the package was unpublished.
    UnpublishedDate *time.Time `json:"unpublishedDate,omitempty"`
    // Version of the package.
    Version *string `json:"version,omitempty"`
}

type NpmRecycleBinPackageVersionDetails struct {
    // Setting to false will undo earlier deletion and restore the package to feed.
    Deleted *bool `json:"deleted,omitempty"`
}

type Operation string

type operationValuesType struct {
    Add Operation
    Remove Operation
    Replace Operation
    Move Operation
    Copy Operation
    Test Operation
}

var OperationValues = operationValuesType{
    Add: "add",
    Remove: "remove",
    Replace: "replace",
    Move: "move",
    Copy: "copy",
    Test: "test",
}

// Package version metadata for an npm package
type Package struct {
    // Related REST links.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Deprecated message, if any, for the package.
    DeprecateMessage *string `json:"deprecateMessage,omitempty"`
    // Package Id.
    Id *string `json:"id,omitempty"`
    // The display name of the package.
    Name *string `json:"name,omitempty"`
    // If and when the package was permanently deleted.
    PermanentlyDeletedDate *time.Time `json:"permanentlyDeletedDate,omitempty"`
    // The history of upstream sources for this package. The first source in the list is the immediate source from which this package was saved.
    SourceChain *[]UpstreamSourceInfo `json:"sourceChain,omitempty"`
    // If and when the package was deleted.
    UnpublishedDate *time.Time `json:"unpublishedDate,omitempty"`
    // The version of the package.
    Version *string `json:"version,omitempty"`
}

type PackageVersionDetails struct {
    // Indicates the deprecate message of a package version
    DeprecateMessage *string `json:"deprecateMessage,omitempty"`
    // The view to which the package version will be added
    Views *JsonPatchOperation `json:"views,omitempty"`
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
