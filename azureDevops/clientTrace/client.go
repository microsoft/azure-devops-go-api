// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package clientTrace

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
func (client Client) PublishEvents(ctx context.Context, args PublishEventsArgs) error {
    if args.Events == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "events"}
    }
    body, marshalErr := json.Marshal(*args.Events)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("06bcc74a-1491-4eb8-a0eb-704778f9d041")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the PublishEvents function
type PublishEventsArgs struct {
    // (required)
    Events *[]ClientTraceEvent
}

