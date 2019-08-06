// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package fileContainer

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azureDevOps"
	"net/http"
	"net/url"
	"strconv"
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

// [Preview API] Creates the specified items in in the referenced container.
func (client *Client) CreateItems(ctx context.Context, args CreateItemsArgs) (*[]FileContainerItem, error) {
	if args.Items == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "items"}
	}
	routeValues := make(map[string]string)
	if args.ContainerId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "containerId"}
	}
	routeValues["containerId"] = strconv.Itoa(*args.ContainerId)

	queryParams := url.Values{}
	if args.Scope != nil {
		queryParams.Add("scope", (*args.Scope).String())
	}
	body, marshalErr := json.Marshal(*args.Items)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.4", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []FileContainerItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateItems function
type CreateItemsArgs struct {
	// (required)
	Items *azureDevOps.VssJsonCollectionWrapper
	// (required)
	ContainerId *int
	// (optional) A guid representing the scope of the container. This is often the project id.
	Scope *uuid.UUID
}

// [Preview API] Deletes the specified items in a container.
func (client *Client) DeleteItem(ctx context.Context, args DeleteItemArgs) error {
	routeValues := make(map[string]string)
	if args.ContainerId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "containerId"}
	}
	routeValues["containerId"] = strconv.FormatUint(*args.ContainerId, 10)

	queryParams := url.Values{}
	if args.ItemPath == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "itemPath"}
	}
	queryParams.Add("itemPath", *args.ItemPath)
	if args.Scope != nil {
		queryParams.Add("scope", (*args.Scope).String())
	}
	locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.4", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteItem function
type DeleteItemArgs struct {
	// (required) Container Id.
	ContainerId *uint64
	// (required) Path to delete.
	ItemPath *string
	// (optional) A guid representing the scope of the container. This is often the project id.
	Scope *uuid.UUID
}

// [Preview API] Gets containers filtered by a comma separated list of artifact uris within the same scope, if not specified returns all containers
func (client *Client) GetContainers(ctx context.Context, args GetContainersArgs) (*[]FileContainer, error) {
	queryParams := url.Values{}
	if args.Scope != nil {
		queryParams.Add("scope", (*args.Scope).String())
	}
	if args.ArtifactUris != nil {
		queryParams.Add("artifactUris", *args.ArtifactUris)
	}
	locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.4", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []FileContainer
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetContainers function
type GetContainersArgs struct {
	// (optional) A guid representing the scope of the container. This is often the project id.
	Scope *uuid.UUID
	// (optional)
	ArtifactUris *string
}

// [Preview API]
func (client *Client) GetItems(ctx context.Context, args GetItemsArgs) (*[]FileContainerItem, error) {
	routeValues := make(map[string]string)
	if args.ContainerId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "containerId"}
	}
	routeValues["containerId"] = strconv.FormatUint(*args.ContainerId, 10)

	queryParams := url.Values{}
	if args.Scope != nil {
		queryParams.Add("scope", (*args.Scope).String())
	}
	if args.ItemPath != nil {
		queryParams.Add("itemPath", *args.ItemPath)
	}
	if args.Metadata != nil {
		queryParams.Add("metadata", strconv.FormatBool(*args.Metadata))
	}
	if args.Format != nil {
		queryParams.Add("$format", *args.Format)
	}
	if args.DownloadFileName != nil {
		queryParams.Add("downloadFileName", *args.DownloadFileName)
	}
	if args.IncludeDownloadTickets != nil {
		queryParams.Add("includeDownloadTickets", strconv.FormatBool(*args.IncludeDownloadTickets))
	}
	if args.IsShallow != nil {
		queryParams.Add("isShallow", strconv.FormatBool(*args.IsShallow))
	}
	locationId, _ := uuid.Parse("e4f5c81e-e250-447b-9fef-bd48471bea5e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.4", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []FileContainerItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetItems function
type GetItemsArgs struct {
	// (required)
	ContainerId *uint64
	// (optional)
	Scope *uuid.UUID
	// (optional)
	ItemPath *string
	// (optional)
	Metadata *bool
	// (optional)
	Format *string
	// (optional)
	DownloadFileName *string
	// (optional)
	IncludeDownloadTickets *bool
	// (optional)
	IsShallow *bool
}
