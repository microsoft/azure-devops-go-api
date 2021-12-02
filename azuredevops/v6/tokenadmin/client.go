// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tokenadmin

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/delegatedauthorization"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("af68438b-ed04-4407-9eb6-f1dbae3f922e")

type Client interface {
	// [Preview API] Creates a revocation rule to prevent the further usage of any OAuth authorizations that were created before the current point in time and which match the conditions in the rule.
	CreateRevocationRule(context.Context, CreateRevocationRuleArgs) error
	// [Preview API]
	GetPersonalAccessToken(context.Context, GetPersonalAccessTokenArgs) (*delegatedauthorization.SessionTokenResult, error)
	// [Preview API] Lists of all the session token details of the personal access tokens (PATs) for a particular user.
	ListPersonalAccessTokens(context.Context, ListPersonalAccessTokensArgs) (*TokenAdminPagedSessionTokens, error)
	// [Preview API] Revokes the listed OAuth authorizations.
	RevokeAuthorizations(context.Context, RevokeAuthorizationsArgs) error
	// [Preview API]
	RevokePersonalAccessToken(context.Context, RevokePersonalAccessTokenArgs) error
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// [Preview API] Creates a revocation rule to prevent the further usage of any OAuth authorizations that were created before the current point in time and which match the conditions in the rule.
func (client *ClientImpl) CreateRevocationRule(ctx context.Context, args CreateRevocationRuleArgs) error {
	if args.RevocationRule == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.RevocationRule"}
	}
	body, marshalErr := json.Marshal(*args.RevocationRule)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("ee4afb16-e7ab-4ed8-9d4b-4ef3e78f97e4")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the CreateRevocationRule function
type CreateRevocationRuleArgs struct {
	// (required) The revocation rule to create. The rule must specify a space-separated list of scopes, after which preexisting OAuth authorizations that match that any of the scopes will be rejected. For a list of all OAuth scopes supported by VSTS, see: https://docs.microsoft.com/en-us/vsts/integrate/get-started/authentication/oauth?view=vsts#scopes The rule may also specify the time before which to revoke tokens.
	RevocationRule *TokenAdminRevocationRule
}

// [Preview API]
func (client *ClientImpl) GetPersonalAccessToken(ctx context.Context, args GetPersonalAccessTokenArgs) (*delegatedauthorization.SessionTokenResult, error) {
	if args.AccessTokenKey == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AccessTokenKey"}
	}
	queryParams := url.Values{}
	if args.IsPublic == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "isPublic"}
	}
	queryParams.Add("isPublic", strconv.FormatBool(*args.IsPublic))
	body, marshalErr := json.Marshal(*args.AccessTokenKey)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("11e3d37f-fa7e-4721-ab2d-2d931bd944c4")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue delegatedauthorization.SessionTokenResult
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPersonalAccessToken function
type GetPersonalAccessTokenArgs struct {
	// (required)
	AccessTokenKey *string
	// (required)
	IsPublic *bool
}

// [Preview API] Lists of all the session token details of the personal access tokens (PATs) for a particular user.
func (client *ClientImpl) ListPersonalAccessTokens(ctx context.Context, args ListPersonalAccessTokensArgs) (*TokenAdminPagedSessionTokens, error) {
	routeValues := make(map[string]string)
	if args.SubjectDescriptor == nil || *args.SubjectDescriptor == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.SubjectDescriptor"}
	}
	routeValues["subjectDescriptor"] = *args.SubjectDescriptor

	queryParams := url.Values{}
	if args.PageSize != nil {
		queryParams.Add("pageSize", strconv.Itoa(*args.PageSize))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.IsPublic != nil {
		queryParams.Add("isPublic", strconv.FormatBool(*args.IsPublic))
	}
	locationId, _ := uuid.Parse("af68438b-ed04-4407-9eb6-f1dbae3f922e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TokenAdminPagedSessionTokens
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListPersonalAccessTokens function
type ListPersonalAccessTokensArgs struct {
	// (required) The descriptor of the target user.
	SubjectDescriptor *string
	// (optional) The maximum number of results to return on each page.
	PageSize *int
	// (optional) An opaque data blob that allows the next page of data to resume immediately after where the previous page ended. The only reliable way to know if there is more data left is the presence of a continuation token.
	ContinuationToken *string
	// (optional) Set to false for PAT tokens and true for SSH tokens.
	IsPublic *bool
}

// [Preview API] Revokes the listed OAuth authorizations.
func (client *ClientImpl) RevokeAuthorizations(ctx context.Context, args RevokeAuthorizationsArgs) error {
	if args.Revocations == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.Revocations"}
	}
	queryParams := url.Values{}
	if args.IsPublic != nil {
		queryParams.Add("isPublic", strconv.FormatBool(*args.IsPublic))
	}
	body, marshalErr := json.Marshal(*args.Revocations)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("a9c08b2c-5466-4e22-8626-1ff304ffdf0f")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RevokeAuthorizations function
type RevokeAuthorizationsArgs struct {
	// (required) The list of objects containing the authorization IDs of the OAuth authorizations, such as session tokens retrieved by listed a users PATs, that should be revoked.
	Revocations *[]TokenAdminRevocation
	// (optional) Set to false for PAT tokens and true for SSH tokens.
	IsPublic *bool
}

// [Preview API]
func (client *ClientImpl) RevokePersonalAccessToken(ctx context.Context, args RevokePersonalAccessTokenArgs) error {
	if args.AccessTokenKey == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.AccessTokenKey"}
	}
	queryParams := url.Values{}
	if args.IsPublic == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "isPublic"}
	}
	queryParams.Add("isPublic", strconv.FormatBool(*args.IsPublic))
	body, marshalErr := json.Marshal(*args.AccessTokenKey)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("55687c95-c811-41e7-889f-25afb03eda19")
	_, err := client.Client.Send(ctx, http.MethodPut, locationId, "6.0-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RevokePersonalAccessToken function
type RevokePersonalAccessTokenArgs struct {
	// (required)
	AccessTokenKey *string
	// (required)
	IsPublic *bool
}
