// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package operations

import (
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// Gets an operation from the the operationId using the given pluginId.
// operationId (required): The ID for the operation.
// pluginId (optional): The ID for the plugin.
func (client Client) GetOperation(operationId *uuid.UUID, pluginId *uuid.UUID) (*Operation, error) {
    routeValues := make(map[string]string)
    if operationId == nil {
        return nil, errors.New("operationId is a required parameter")
    }
    routeValues["operationId"] = (*operationId).String()

    queryParams := url.Values{}
    if pluginId != nil {
        queryParams.Add("pluginId", (*pluginId).String())
    }
    locationId, _ := uuid.Parse("9a1b74b4-2ca8-4a9f-8470-c2f2e6fdc949")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Operation
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

