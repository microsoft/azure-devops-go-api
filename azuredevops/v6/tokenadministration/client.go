// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tokenadministration

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/delegatedauthorization"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/tokenadmin"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("95935461-9e54-44bd-b9fb-04f4dd05d640")

type Client interface {
	// [Preview API] Revokes the listed OAuth authorizations.
	ListIdentitiesWithGlobalAccessTokens(context.Context, ListIdentitiesWithGlobalAccessTokensArgs) (*[]uuid.UUID, error)
	// [Preview API] Lists of all the session token details of the personal access tokens (PATs) for a particular user.
	ListPersonalAccessTokens(context.Context, ListPersonalAccessTokensArgs) (*tokenadmin.TokenAdminPagedSessionTokens, error)
	// [Preview API] Revokes the listed OAuth authorizations.
	RevokeAuthorizations(context.Context, RevokeAuthorizationsArgs) (*[]delegatedauthorization.SessionToken, error)
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

// [Preview API] Revokes the listed OAuth authorizations.
func (client *ClientImpl) ListIdentitiesWithGlobalAccessTokens(ctx context.Context, args ListIdentitiesWithGlobalAccessTokensArgs) (*[]uuid.UUID, error) {
	if args.Revocations == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Revocations"}
	}
	queryParams := url.Values{}
	if args.IsPublic != nil {
		queryParams.Add("isPublic", strconv.FormatBool(*args.IsPublic))
	}
	body, marshalErr := json.Marshal(*args.Revocations)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("30d3a12b-66c3-4669-b016-ecb0706c8d0f")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []uuid.UUID
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListIdentitiesWithGlobalAccessTokens function
type ListIdentitiesWithGlobalAccessTokensArgs struct {
	// (required) The list of identities containing the authorization IDs of the OAuth authorizations, such as session tokens retrieved by listed a users PATs, that should be checked for global access tokens.
	Revocations *[]tokenadmin.TokenAdminRevocation
	// (optional) Set to false for PAT tokens and true for SSH tokens.
	IsPublic *bool
}

// [Preview API] Lists of all the session token details of the personal access tokens (PATs) for a particular user.
func (client *ClientImpl) ListPersonalAccessTokens(ctx context.Context, args ListPersonalAccessTokensArgs) (*tokenadmin.TokenAdminPagedSessionTokens, error) {
	if args.Audience == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Audience"}
	}
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
	body, marshalErr := json.Marshal(*args.Audience)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1bb7db14-87c5-4762-bf77-a70ad34a9ab3")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue tokenadmin.TokenAdminPagedSessionTokens
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListPersonalAccessTokens function
type ListPersonalAccessTokensArgs struct {
	// (required)
	Audience *[]string
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
func (client *ClientImpl) RevokeAuthorizations(ctx context.Context, args RevokeAuthorizationsArgs) (*[]delegatedauthorization.SessionToken, error) {
	if args.Revocations == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Revocations"}
	}
	queryParams := url.Values{}
	if args.HostId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "hostId"}
	}
	queryParams.Add("hostId", (*args.HostId).String())
	if args.IsPublic != nil {
		queryParams.Add("isPublic", strconv.FormatBool(*args.IsPublic))
	}
	body, marshalErr := json.Marshal(*args.Revocations)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("a2e4520b-1cc8-4526-871e-f3a8f865f221")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []delegatedauthorization.SessionToken
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the RevokeAuthorizations function
type RevokeAuthorizationsArgs struct {
	// (required) The list of objects containing the authorization IDs of the OAuth authorizations, such as session tokens retrieved by listed a users PATs, that should be revoked.
	Revocations *TokenAdministrationRevocation
	// (required) Host Id to display on the notification page to manage tokens.
	HostId *uuid.UUID
	// (optional) Set to false for PAT tokens and true for SSH tokens.
	IsPublic *bool
}
