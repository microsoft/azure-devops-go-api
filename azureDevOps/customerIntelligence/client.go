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
	"github.com/microsoft/azure-devops-go-api/azureDevOps"
	"net/http"
)

type Client struct {
	Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) *Client {
	client := connection.GetClientByUrl(connection.BaseUrl)
	return &Client{
		Client: *client,
	}
}

// [Preview API]
func (client *Client) PublishEvents(ctx context.Context, args PublishEventsArgs) error {
	if args.Events == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "events"}
	}
	body, marshalErr := json.Marshal(*args.Events)
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

// Arguments for the PublishEvents function
type PublishEventsArgs struct {
	// (required)
	Events *[]CustomerIntelligenceEvent
}
