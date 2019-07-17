// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package fileContainer

import (
    "github.com/google/uuid"
    "time"
)

// Status of a container item.
type ContainerItemStatus string

// Type of a container item.
type ContainerItemType string

// Options a container can have.
type ContainerOptions string

// Represents a container that encapsulates a hierarchical file system.
type FileContainer struct {
    // Uri of the artifact associated with the container.
    ArtifactUri *string `json:"artifactUri,omitempty"`
    // Download Url for the content of this item.
    ContentLocation *string `json:"contentLocation,omitempty"`
    // Owner.
    CreatedBy *uuid.UUID `json:"createdBy,omitempty"`
    // Creation date.
    DateCreated *time.Time `json:"dateCreated,omitempty"`
    // Description.
    Description *string `json:"description,omitempty"`
    // Id.
    Id *uint64 `json:"id,omitempty"`
    // Location of the item resource.
    ItemLocation *string `json:"itemLocation,omitempty"`
    // ItemStore Locator for this container.
    LocatorPath *string `json:"locatorPath,omitempty"`
    // Name.
    Name *string `json:"name,omitempty"`
    // Options the container can have.
    Options *ContainerOptions `json:"options,omitempty"`
    // Project Id.
    ScopeIdentifier *uuid.UUID `json:"scopeIdentifier,omitempty"`
    // Security token of the artifact associated with the container.
    SecurityToken *string `json:"securityToken,omitempty"`
    // Identifier of the optional encryption key.
    SigningKeyId *uuid.UUID `json:"signingKeyId,omitempty"`
    // Total size of the files in bytes.
    Size *uint64 `json:"size,omitempty"`
}

// Represents an item in a container.
type FileContainerItem struct {
    // Container Id.
    ContainerId *uint64 `json:"containerId,omitempty"`
    ContentId *[]byte `json:"contentId,omitempty"`
    // Download Url for the content of this item.
    ContentLocation *string `json:"contentLocation,omitempty"`
    // Creator.
    CreatedBy *uuid.UUID `json:"createdBy,omitempty"`
    // Creation date.
    DateCreated *time.Time `json:"dateCreated,omitempty"`
    // Last modified date.
    DateLastModified *time.Time `json:"dateLastModified,omitempty"`
    // Encoding of the file. Zero if not a file.
    FileEncoding *int `json:"fileEncoding,omitempty"`
    // Hash value of the file. Null if not a file.
    FileHash *[]byte `json:"fileHash,omitempty"`
    // Id of the file content.
    FileId *int `json:"fileId,omitempty"`
    // Length of the file. Zero if not of a file.
    FileLength *uint64 `json:"fileLength,omitempty"`
    // Type of the file. Zero if not a file.
    FileType *int `json:"fileType,omitempty"`
    // Location of the item resource.
    ItemLocation *string `json:"itemLocation,omitempty"`
    // Type of the item: Folder, File or String.
    ItemType *ContainerItemType `json:"itemType,omitempty"`
    // Modifier.
    LastModifiedBy *uuid.UUID `json:"lastModifiedBy,omitempty"`
    // Unique path that identifies the item.
    Path *string `json:"path,omitempty"`
    // Project Id.
    ScopeIdentifier *uuid.UUID `json:"scopeIdentifier,omitempty"`
    // Status of the item: Created or Pending Upload.
    Status *ContainerItemStatus `json:"status,omitempty"`
    Ticket *string `json:"ticket,omitempty"`
}
