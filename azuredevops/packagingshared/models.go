// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package packagingshared

import (
	"github.com/google/uuid"
)

// Data required for promoting multiple package versions. Pass this while performing {protocol}BatchOperationTypes.Promote batch operation.
type BatchPromoteData struct {
	// Id or Name of the view, packages need to be promoted to.
	ViewId *string `json:"viewId,omitempty"`
}

// Minimal package details required to identify a package within a protocol.
type MinimalPackageDetails struct {
	// Package name.
	Id *string `json:"id,omitempty"`
	// Package version.
	Version *string `json:"version,omitempty"`
}

// Type of an upstream source, such as Public or Internal.
type PackagingSourceType string

type packagingSourceTypeValuesType struct {
	Public   PackagingSourceType
	Internal PackagingSourceType
}

var PackagingSourceTypeValues = packagingSourceTypeValuesType{
	// Publicly available source.
	Public: "public",
	// Azure DevOps upstream source.
	Internal: "internal",
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
