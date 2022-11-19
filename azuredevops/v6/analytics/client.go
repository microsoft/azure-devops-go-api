// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package analytics

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"io"
	"net/http"
	"strconv"
)

type Client interface {
	// [Preview API]
	CreateShard(context.Context, CreateShardArgs) (*StageProviderShardInfo, error)
	// [Preview API]
	DeleteShard(context.Context, DeleteShardArgs) error
	// [Preview API]
	GetShard(context.Context, GetShardArgs) (*StageProviderShardInfo, error)
	// [Preview API] Gets the current state of the Analytics feature.
	GetState(context.Context, GetStateArgs) (*AnalyticsStateDetails, error)
	// [Preview API]
	GetTable(context.Context, GetTableArgs) (*StageTableInfo, error)
	// [Preview API]
	InvalidateShard(context.Context, InvalidateShardArgs) error
	// [Preview API]
	StageRecords(context.Context, StageRecordsArgs) (*IngestResult, error)
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
func (client *ClientImpl) CreateShard(ctx context.Context, args CreateShardArgs) (*StageProviderShardInfo, error) {
	routeValues := make(map[string]string)
	if args.Table == nil || *args.Table == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Table"}
	}
	routeValues["table"] = *args.Table
	if args.ProviderShard == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProviderShard"}
	}
	routeValues["providerShard"] = strconv.Itoa(*args.ProviderShard)

	locationId, _ := uuid.Parse("9bd3f7d0-e20d-4e7b-95ba-854704939f9e")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue StageProviderShardInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateShard function
type CreateShardArgs struct {
	// (required)
	Table *string
	// (required)
	ProviderShard *int
}

// [Preview API]
func (client *ClientImpl) DeleteShard(ctx context.Context, args DeleteShardArgs) error {
	routeValues := make(map[string]string)
	if args.Table == nil || *args.Table == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Table"}
	}
	routeValues["table"] = *args.Table
	if args.ProviderShard == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ProviderShard"}
	}
	routeValues["providerShard"] = strconv.Itoa(*args.ProviderShard)

	locationId, _ := uuid.Parse("9bd3f7d0-e20d-4e7b-95ba-854704939f9e")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteShard function
type DeleteShardArgs struct {
	// (required)
	Table *string
	// (required)
	ProviderShard *int
}

// [Preview API]
func (client *ClientImpl) GetShard(ctx context.Context, args GetShardArgs) (*StageProviderShardInfo, error) {
	routeValues := make(map[string]string)
	if args.Table == nil || *args.Table == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Table"}
	}
	routeValues["table"] = *args.Table
	if args.ProviderShard == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProviderShard"}
	}
	routeValues["providerShard"] = strconv.Itoa(*args.ProviderShard)

	locationId, _ := uuid.Parse("9bd3f7d0-e20d-4e7b-95ba-854704939f9e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue StageProviderShardInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetShard function
type GetShardArgs struct {
	// (required)
	Table *string
	// (required)
	ProviderShard *int
}

// [Preview API] Gets the current state of the Analytics feature.
func (client *ClientImpl) GetState(ctx context.Context, args GetStateArgs) (*AnalyticsStateDetails, error) {
	locationId, _ := uuid.Parse("0b79c382-d776-40b9-87b4-407fb8f7df24")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AnalyticsStateDetails
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetState function
type GetStateArgs struct {
}

// [Preview API]
func (client *ClientImpl) GetTable(ctx context.Context, args GetTableArgs) (*StageTableInfo, error) {
	routeValues := make(map[string]string)
	if args.Table == nil || *args.Table == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Table"}
	}
	routeValues["table"] = *args.Table

	locationId, _ := uuid.Parse("9bd3f7d0-e20d-4e7b-95ba-854704939f9e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue StageTableInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTable function
type GetTableArgs struct {
	// (required)
	Table *string
}

// [Preview API]
func (client *ClientImpl) InvalidateShard(ctx context.Context, args InvalidateShardArgs) error {
	if args.Shard == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.Shard"}
	}
	routeValues := make(map[string]string)
	if args.Table == nil || *args.Table == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Table"}
	}
	routeValues["table"] = *args.Table
	if args.ProviderShard == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ProviderShard"}
	}
	routeValues["providerShard"] = strconv.Itoa(*args.ProviderShard)

	body, marshalErr := json.Marshal(*args.Shard)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("328a8d58-1727-4715-9a3d-e236feebd247")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the InvalidateShard function
type InvalidateShardArgs struct {
	// (required)
	Shard *StageShardInvalid
	// (required)
	Table *string
	// (required)
	ProviderShard *int
}

// [Preview API]
func (client *ClientImpl) StageRecords(ctx context.Context, args StageRecordsArgs) (*IngestResult, error) {
	if args.UploadStream == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.UploadStream"}
	}
	routeValues := make(map[string]string)
	if args.Table == nil || *args.Table == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Table"}
	}
	routeValues["table"] = *args.Table
	if args.ProviderShard == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProviderShard"}
	}
	routeValues["providerShard"] = strconv.Itoa(*args.ProviderShard)
	if args.Stream == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Stream"}
	}
	routeValues["stream"] = strconv.Itoa(*args.Stream)

	locationId, _ := uuid.Parse("9bd3f7d0-e20d-4e7b-95ba-854704939f9e")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, nil, args.UploadStream, "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue IngestResult
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the StageRecords function
type StageRecordsArgs struct {
	// (required) Stream to upload
	UploadStream io.Reader
	// (required)
	Table *string
	// (required)
	ProviderShard *int
	// (required)
	Stream *int
}
