// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package customerIntelligence

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API]
// ctx
// events (required)
func (client Client) PublishEvents(ctx context.Context, events *[]CustomerIntelligenceEvent) error {
    if events == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "events"}
    }
    body, marshalErr := json.Marshal(*events)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("b5cc35c2-ff2b-491d-a085-24b6e9f396fd")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

