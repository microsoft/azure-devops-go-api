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
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
)

var ResourceAreaId, _ = uuid.Parse("b40c1171-807a-493a-8f3f-5c26d5e2f5aa")

type Client interface {
	// [Preview API] Creates a session, a wrapper around a feed that can store additional metadata on the packages published to it.
	CreateSession(context.Context, CreateSessionArgs) (*SessionResponse, error)
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

// [Preview API] Creates a session, a wrapper around a feed that can store additional metadata on the packages published to it.
func (client *ClientImpl) CreateSession(ctx context.Context, args CreateSessionArgs) (*SessionResponse, error) {
	if args.SessionRequest == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SessionRequest"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.Protocol == nil || *args.Protocol == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Protocol"}
	}
	routeValues["protocol"] = *args.Protocol

	body, marshalErr := json.Marshal(*args.SessionRequest)
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

// Arguments for the CreateSession function
type CreateSessionArgs struct {
	// (required) The feed and metadata for the session
	SessionRequest *SessionRequest
	// (required) The protocol that the session will target
	Protocol *string
	// (optional) Project ID or project name
	Project *string
}
