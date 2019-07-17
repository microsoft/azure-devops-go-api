// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package accounts

import (
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("0d55247a-1c47-4462-9b1f-5e2125590ee6")

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

// Get a list of accounts for a specific owner or a specific member.
// ownerId (optional): ID for the owner of the accounts.
// memberId (optional): ID for a member of the accounts.
// properties (optional)
func (client Client) GetAccounts(ownerId *uuid.UUID, memberId *uuid.UUID, properties *string) (*[]Account, error) {
    queryParams := url.Values{}
    if ownerId != nil {
        queryParams.Add("ownerId", (*ownerId).String())
    }
    if memberId != nil {
        queryParams.Add("memberId", (*memberId).String())
    }
    if properties != nil {
        queryParams.Add("properties", *properties)
    }
    locationId, _ := uuid.Parse("229a6a53-b428-4ffb-a835-e8f36b5b4b1e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Account
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

