// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package sbom

import (
    "context"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azuredevops/v7.1"
    "net/http"
)

var ResourceAreaId, _ = uuid.Parse("2e504d18-2c0c-46f8-af8f-322d2af0068a")

type Client interface {
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

