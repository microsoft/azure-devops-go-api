// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tokenAdministration

import (
    "github.com/google/uuid"
    "time"
)

type SessionToken struct {
    AccessId *uuid.UUID `json:"accessId,omitempty"`
    // This is populated when user requests a compact token. The alternate token value is self describing token.
    AlternateToken *string `json:"alternateToken,omitempty"`
    AuthorizationId *uuid.UUID `json:"authorizationId,omitempty"`
    Claims *map[string]string `json:"claims,omitempty"`
    ClientId *uuid.UUID `json:"clientId,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    HostAuthorizationId *uuid.UUID `json:"hostAuthorizationId,omitempty"`
    IsPublic *bool `json:"isPublic,omitempty"`
    IsValid *bool `json:"isValid,omitempty"`
    PublicData *string `json:"publicData,omitempty"`
    Scope *string `json:"scope,omitempty"`
    Source *string `json:"source,omitempty"`
    TargetAccounts *[]uuid.UUID `json:"targetAccounts,omitempty"`
    // This is computed and not returned in Get queries
    Token *string `json:"token,omitempty"`
    UserId *uuid.UUID `json:"userId,omitempty"`
    ValidFrom *time.Time `json:"validFrom,omitempty"`
    ValidTo *time.Time `json:"validTo,omitempty"`
}

type TokenAdministrationRevocation struct {
    // A list of audience (target accounts) to limit the revocations to
    Audience *[]string `json:"audience,omitempty"`
    // A list of authorization ID of the OAuth authorization to revoke.
    AuthorizationIds *[]uuid.UUID `json:"authorizationIds,omitempty"`
}

// A paginated list of session tokens. Session tokens correspond to OAuth credentials such as personal access tokens (PATs) and other OAuth authorizations.
type TokenAdminPagedSessionTokens struct {
    // The continuation token that can be used to retrieve the next page of session tokens, or <code>null</code> if there is no next page.
    ContinuationToken *uuid.UUID `json:"continuationToken,omitempty"`
    // The list of all session tokens in the current page.
    Value *[]SessionToken `json:"value,omitempty"`
}

// A request to revoke a particular delegated authorization.
type TokenAdminRevocation struct {
    // The authorization ID of the OAuth authorization to revoke.
    AuthorizationId *uuid.UUID `json:"authorizationId,omitempty"`
}
