// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package profile

import (
    "context"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("8ccfef3d-2b87-4e99-8ccb-66e343d2daa8")

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// Gets a user profile.
// ctx
// id (required): The ID of the target user profile within the same organization, or 'me' to get the profile of the current authenticated user.
// details (optional): Return public profile information such as display name, email address, country, etc. If false, the withAttributes parameter is ignored.
// withAttributes (optional): If true, gets the attributes (named key-value pairs of arbitrary data) associated with the profile. The partition parameter must also have a value.
// partition (optional): The partition (named group) of attributes to return.
// coreAttributes (optional): A comma-delimited list of core profile attributes to return. Valid values are Email, Avatar, DisplayName, and ContactWithOffers.
// forceRefresh (optional): Not used in this version of the API.
func (client Client) GetProfile(ctx context.Context, id *string, details *bool, withAttributes *bool, partition *string, coreAttributes *string, forceRefresh *bool) (*Profile, error) {
    routeValues := make(map[string]string)
    if id == nil || *id == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    queryParams := url.Values{}
    if details != nil {
        queryParams.Add("details", strconv.FormatBool(*details))
    }
    if withAttributes != nil {
        queryParams.Add("withAttributes", strconv.FormatBool(*withAttributes))
    }
    if partition != nil {
        queryParams.Add("partition", *partition)
    }
    if coreAttributes != nil {
        queryParams.Add("coreAttributes", *coreAttributes)
    }
    if forceRefresh != nil {
        queryParams.Add("forceRefresh", strconv.FormatBool(*forceRefresh))
    }
    locationId, _ := uuid.Parse("f83735dc-483f-4238-a291-d45f6080a9af")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Profile
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

