// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tokenadmin

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/delegatedauthorization"
)

// A paginated list of session tokens. Session tokens correspond to OAuth credentials such as personal access tokens (PATs) and other OAuth authorizations.
type TokenAdminPagedSessionTokens struct {
	// The continuation token that can be used to retrieve the next page of session tokens, or <code>null</code> if there is no next page.
	ContinuationToken *uuid.UUID `json:"continuationToken,omitempty"`
	// The list of all session tokens in the current page.
	Value *[]delegatedauthorization.SessionToken `json:"value,omitempty"`
}

// A request to revoke a particular delegated authorization.
type TokenAdminRevocation struct {
	// The authorization ID of the OAuth authorization to revoke.
	AuthorizationId *uuid.UUID `json:"authorizationId,omitempty"`
}

// A rule which is applied to disable any incoming delegated authorization which matches the given properties.
type TokenAdminRevocationRule struct {
	// A datetime cutoff. Tokens created before this time will be rejected. This is an optional parameter. If omitted, defaults to the time at which the rule was created.
	CreatedBefore *azuredevops.Time `json:"createdBefore,omitempty"`
	// A string containing a space-delimited list of OAuth scopes. A token matching any one of the scopes will be rejected. For a list of all OAuth scopes supported by Azure DevOps, see: https://docs.microsoft.com/en-us/azure/devops/integrate/get-started/authentication/oauth?view=azure-devops#scopes This is a mandatory parameter.
	Scopes *string `json:"scopes,omitempty"`
}
