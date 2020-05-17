// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package profileregions

import (
	"context"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
)

var ResourceAreaId, _ = uuid.Parse("8ccfef3d-2b87-4e99-8ccb-66e343d2daa8")

type Client interface {
	// [Preview API] Lookup up country/region based on provided IPv4, null if using the remote IPv4 address.
	GetGeoRegion(context.Context, GetGeoRegionArgs) (*GeoRegion, error)
	// [Preview API]
	GetRegions(context.Context, GetRegionsArgs) (*ProfileRegions, error)
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

// [Preview API] Lookup up country/region based on provided IPv4, null if using the remote IPv4 address.
func (client *ClientImpl) GetGeoRegion(ctx context.Context, args GetGeoRegionArgs) (*GeoRegion, error) {
	queryParams := url.Values{}
	if args.Ip == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "ip"}
	}
	queryParams.Add("ip", *args.Ip)
	locationId, _ := uuid.Parse("35b3ff1d-ab4c-4d1c-98bb-f6ea21d86bd9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GeoRegion
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetGeoRegion function
type GetGeoRegionArgs struct {
	// (required)
	Ip *string
}

// [Preview API]
func (client *ClientImpl) GetRegions(ctx context.Context, args GetRegionsArgs) (*ProfileRegions, error) {
	locationId, _ := uuid.Parse("b129ca90-999d-47bb-ab37-0dcf784ee633")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProfileRegions
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRegions function
type GetRegionsArgs struct {
}
