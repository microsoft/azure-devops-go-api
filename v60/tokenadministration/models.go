// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tokenadministration

import (
	"github.com/google/uuid"
)

type TokenAdministrationRevocation struct {
	// A list of audience (target accounts) to limit the revocations to
	Audience *[]string `json:"audience,omitempty"`
	// A list of authorization ID of the OAuth authorization to revoke.
	AuthorizationIds *[]uuid.UUID `json:"authorizationIds,omitempty"`
}
