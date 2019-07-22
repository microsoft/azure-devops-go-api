// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package provenance

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
)

var ResourceAreaId, _ = uuid.Parse("b40c1171-807a-493a-8f3f-5c26d5e2f5aa")

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

// [Preview API] Creates a session, a wrapper around a feed that can store additional metadata on the packages published to it.
// ctx
// sessionRequest (required): The feed and metadata for the session
// protocol (required): The protocol that the session will target
// project (optional): Project ID or project name
func (client Client) CreateSession(ctx context.Context, sessionRequest *SessionRequest, protocol *string, project *string) (*SessionResponse, error) {
    if sessionRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "sessionRequest"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if protocol == nil || *protocol == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "protocol"} 
    }
    routeValues["protocol"] = *protocol

    body, marshalErr := json.Marshal(*sessionRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("503b4e54-ebf4-4d04-8eee-21c00823c2ac")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SessionResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

