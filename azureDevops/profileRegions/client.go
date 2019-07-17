// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package profileRegions

import (
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("8ccfef3d-2b87-4e99-8ccb-66e343d2daa8")

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API] Lookup up country/region based on provided IPv4, null if using the remote IPv4 address.
// ip (required)
func (client Client) GetGeoRegion(ip *string) (*GeoRegion, error) {
    queryParams := url.Values{}
    if ip == nil {
        return nil, errors.New("ip is a required parameter")
    }
    queryParams.Add("ip", *ip)
    locationId, _ := uuid.Parse("35b3ff1d-ab4c-4d1c-98bb-f6ea21d86bd9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GeoRegion
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
func (client Client) GetRegions() (*ProfileRegions, error) {
    locationId, _ := uuid.Parse("b129ca90-999d-47bb-ab37-0dcf784ee633")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProfileRegions
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

