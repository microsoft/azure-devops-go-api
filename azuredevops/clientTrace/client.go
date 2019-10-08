// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package clienttrace

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
)

type Client interface {
	// [Preview API]
	PublishEvents(context.Context, PublishEventsArgs) error
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) Client {
	client := connection.GetClientByUrl(connection.BaseUrl)
	return &ClientImpl{
		Client: *client,
	}
}

// [Preview API]
func (client *ClientImpl) PublishEvents(ctx context.Context, args PublishEventsArgs) error {
	if args.Events == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.Events"}
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
