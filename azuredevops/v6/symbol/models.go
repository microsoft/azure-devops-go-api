// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package symbol

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/symbolcommon"
)

// A dual-purpose data object, the debug entry is used by the client to publish the symbol file (with file's blob identifier, which can be calculated from VSTS hashing algorithm) or query the file (with a client key). Since the symbol server tries to return a matched symbol file with the richest information level, it may not always point to the same symbol file for different queries with same client key.
type DebugEntry struct {
	// The ID of user who created this item. Optional.
	CreatedBy *uuid.UUID `json:"createdBy,omitempty"`
	// The date time when this item is created. Optional.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// An identifier for this item. Optional.
	Id *string `json:"id,omitempty"`
	// An opaque ETag used to synchronize with the version stored at server end. Optional.
	StorageETag *string `json:"storageETag,omitempty"`
	// A URI which can be used to retrieve this item in its raw format. Optional. Note this is distinguished from other URIs that are present in a derived resource.
	Url *string `json:"url,omitempty"`
	// Details of the blob formatted to be deserialized for symbol service.
	BlobDetails *JsonBlobIdentifierWithBlocks `json:"blobDetails,omitempty"`
	// A blob identifier of the symbol file to upload to this debug entry. This property is mostly used during creation of debug entry (a.k.a. symbol publishing) to allow the server to query the existence of the blob.
	BlobIdentifier *JsonBlobIdentifier `json:"blobIdentifier,omitempty"`
	// The URI to get the symbol file. Provided by the server, the URI contains authentication information and is readily accessible by plain HTTP GET request. The client is recommended to retrieve the file as soon as it can since the URI will expire in a short period.
	BlobUri *string `json:"blobUri,omitempty"`
	// A key the client (debugger, for example) uses to find the debug entry. Note it is not unique for each different symbol file as it does not distinguish between those which only differ by information level.
	ClientKey *string `json:"clientKey,omitempty"`
	// The information level this debug entry contains.
	InformationLevel *symbolcommon.DebugInformationLevel `json:"informationLevel,omitempty"`
	// The identifier of symbol request to which this debug entry belongs.
	RequestId *string `json:"requestId,omitempty"`
	// The size for the debug entry.
	Size *uint64 `json:"size,omitempty"`
	// The status of debug entry.
	Status *DebugEntryStatus `json:"status,omitempty"`
}

// A batch of debug entry to create.
type DebugEntryCreateBatch struct {
	// Defines what to do when a debug entry in the batch already exists.
	CreateBehavior *DebugEntryCreateBehavior `json:"createBehavior,omitempty"`
	// The debug entries.
	DebugEntries *[]DebugEntry `json:"debugEntries,omitempty"`
}

// The expected behavior when a debug entry to add already exists.
type DebugEntryCreateBehavior string

type debugEntryCreateBehaviorValuesType struct {
	ThrowIfExists     DebugEntryCreateBehavior
	SkipIfExists      DebugEntryCreateBehavior
	OverwriteIfExists DebugEntryCreateBehavior
}

var DebugEntryCreateBehaviorValues = debugEntryCreateBehaviorValuesType{
	// Throw exceptions at server end. This will translate to 409 (Conflict) HTTP status code.
	ThrowIfExists: "throwIfExists",
	// Do not add this debug entry. The rest of the batch, if any, is not affected.
	SkipIfExists: "skipIfExists",
	// Overwrite the existing debug entry.
	OverwriteIfExists: "overwriteIfExists",
}

// The status of debug entry.
type DebugEntryStatus string

type debugEntryStatusValuesType struct {
	None        DebugEntryStatus
	Created     DebugEntryStatus
	BlobMissing DebugEntryStatus
}

var DebugEntryStatusValues = debugEntryStatusValuesType{
	// The status of this debug entry is undefined or irrelevant in the current context.
	None: "none",
	// The debug entry is created and read to use.
	Created: "created",
	// The symbol file for the requested debug entry is missing.
	BlobMissing: "blobMissing",
}

// BlobBlock hash formatted to be deserialized for symbol service.
type JsonBlobBlockHash struct {
	// Array of hash bytes.
	HashBytes *[]byte `json:"hashBytes,omitempty"`
}

type JsonBlobIdentifier struct {
	IdentifierValue *[]byte `json:"identifierValue,omitempty"`
}

// BlobIdentifier with block hashes formatted to be deserialzied for symbol service.
type JsonBlobIdentifierWithBlocks struct {
	// List of blob block hashes.
	BlockHashes *[]JsonBlobBlockHash `json:"blockHashes,omitempty"`
	// Array of blobId bytes.
	IdentifierValue *[]byte `json:"identifierValue,omitempty"`
}

// Symbol request.
type Request struct {
	// The ID of user who created this item. Optional.
	CreatedBy *uuid.UUID `json:"createdBy,omitempty"`
	// The date time when this item is created. Optional.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// An identifier for this item. Optional.
	Id *string `json:"id,omitempty"`
	// An opaque ETag used to synchronize with the version stored at server end. Optional.
	StorageETag *string `json:"storageETag,omitempty"`
	// A URI which can be used to retrieve this item in its raw format. Optional. Note this is distinguished from other URIs that are present in a derived resource.
	Url *string `json:"url,omitempty"`
	// An optional human-facing description.
	Description *string `json:"description,omitempty"`
	// An optional expiration date for the request. The request will become inaccessible and get deleted after the date, regardless of its status.  On an HTTP POST, if expiration date is null/missing, the server will assign a default expiration data (30 days unless overwridden in the registry at the account level). On PATCH, if expiration date is null/missing, the behavior is to not change whatever the request's current expiration date is.
	ExpirationDate *azuredevops.Time `json:"expirationDate,omitempty"`
	// A human-facing name for the request. Required on POST, ignored on PATCH.
	Name *string `json:"name,omitempty"`
	// The total Size for this request.
	Size *uint64 `json:"size,omitempty"`
	// The status for this request.
	Status *RequestStatus `json:"status,omitempty"`
}

// The status of request.
type RequestStatus string

type requestStatusValuesType struct {
	None        RequestStatus
	Created     RequestStatus
	Sealed      RequestStatus
	Unavailable RequestStatus
}

var RequestStatusValues = requestStatusValuesType{
	// The status of this request is undefined or irrelevant in the current context.
	None: "none",
	// The request is created, and is open to accepting debug entries.
	Created: "created",
	// The request is sealed. No more debug entries can be added to it.
	Sealed: "sealed",
	// The request is no longer available, possibly deleted.
	Unavailable: "unavailable",
}

type ResourceBase struct {
	// The ID of user who created this item. Optional.
	CreatedBy *uuid.UUID `json:"createdBy,omitempty"`
	// The date time when this item is created. Optional.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// An identifier for this item. Optional.
	Id *string `json:"id,omitempty"`
	// An opaque ETag used to synchronize with the version stored at server end. Optional.
	StorageETag *string `json:"storageETag,omitempty"`
	// A URI which can be used to retrieve this item in its raw format. Optional. Note this is distinguished from other URIs that are present in a derived resource.
	Url *string `json:"url,omitempty"`
}
