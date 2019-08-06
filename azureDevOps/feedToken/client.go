// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package feedToken

import (
	"context"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azureDevOps"
	"net/http"
)

var ResourceAreaId, _ = uuid.Parse("cdeb6c7d-6b25-4d6f-b664-c2e3ede202e8")

type Client struct {
	Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) (*Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: *client,
	}, nil
}

// [Preview API] Get a time-limited session token representing the current user, with permissions scoped to the read/write of Artifacts.
func (client *Client) GetPersonalAccessToken(ctx context.Context, args GetPersonalAccessTokenArgs) (*FeedSessionToken, error) {
	routeValues := make(map[string]string)
	if args.FeedName != nil && *args.FeedName != "" {
		routeValues["feedName"] = *args.FeedName
	}

	locationId, _ := uuid.Parse("dfdb7ad7-3d8e-4907-911e-19b4a8330550")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue FeedSessionToken
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPersonalAccessToken function
type GetPersonalAccessTokenArgs struct {
	// (optional)
	FeedName *string
}
