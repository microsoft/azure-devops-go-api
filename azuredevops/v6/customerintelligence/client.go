// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package customerintelligence

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
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
	locationId, _ := uuid.Parse("b5cc35c2-ff2b-491d-a085-24b6e9f396fd")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the PublishEvents function
type PublishEventsArgs struct {
	// (required)
	Events *[]CustomerIntelligenceEvent
}
