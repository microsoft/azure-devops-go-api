// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package location

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azureDevOps"
	"github.com/microsoft/azure-devops-go-api/azureDevOps/webApi"
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

// [Preview API] This was copied and adapted from TeamFoundationConnectionService.Connect()
func (client *Client) GetConnectionData(ctx context.Context, args GetConnectionDataArgs) (*ConnectionData, error) {
	queryParams := url.Values{}
	if args.ConnectOptions != nil {
		queryParams.Add("connectOptions", string(*args.ConnectOptions))
	}
	if args.LastChangeId != nil {
		queryParams.Add("lastChangeId", strconv.Itoa(*args.LastChangeId))
	}
	if args.LastChangeId64 != nil {
		queryParams.Add("lastChangeId64", strconv.FormatUint(*args.LastChangeId64, 10))
	}
	locationId, _ := uuid.Parse("00d9565f-ed9c-4a06-9a50-00e7896ccab4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ConnectionData
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetConnectionData function
type GetConnectionDataArgs struct {
	// (optional)
	ConnectOptions *webApi.ConnectOptions
	// (optional) Obsolete 32-bit LastChangeId
	LastChangeId *int
	// (optional) Non-truncated 64-bit LastChangeId
	LastChangeId64 *uint64
}

// [Preview API]
func (client *Client) GetResourceArea(ctx context.Context, args GetResourceAreaArgs) (*ResourceAreaInfo, error) {
	routeValues := make(map[string]string)
	if args.AreaId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "areaId"}
	}
	routeValues["areaId"] = (*args.AreaId).String()

	queryParams := url.Values{}
	if args.EnterpriseName != nil {
		queryParams.Add("enterpriseName", *args.EnterpriseName)
	}
	if args.OrganizationName != nil {
		queryParams.Add("organizationName", *args.OrganizationName)
	}
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ResourceAreaInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetResourceArea function
type GetResourceAreaArgs struct {
	// (required)
	AreaId *uuid.UUID
	// (optional)
	EnterpriseName *string
	// (optional)
	OrganizationName *string
}

// [Preview API]
func (client *Client) GetResourceAreaByHost(ctx context.Context, args GetResourceAreaByHostArgs) (*ResourceAreaInfo, error) {
	routeValues := make(map[string]string)
	if args.AreaId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "areaId"}
	}
	routeValues["areaId"] = (*args.AreaId).String()

	queryParams := url.Values{}
	if args.HostId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "hostId"}
	}
	queryParams.Add("hostId", (*args.HostId).String())
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ResourceAreaInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetResourceAreaByHost function
type GetResourceAreaByHostArgs struct {
	// (required)
	AreaId *uuid.UUID
	// (required)
	HostId *uuid.UUID
}

// [Preview API]
func (client *Client) GetResourceAreas(ctx context.Context, args GetResourceAreasArgs) (*[]ResourceAreaInfo, error) {
	queryParams := url.Values{}
	if args.EnterpriseName != nil {
		queryParams.Add("enterpriseName", *args.EnterpriseName)
	}
	if args.OrganizationName != nil {
		queryParams.Add("organizationName", *args.OrganizationName)
	}
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ResourceAreaInfo
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetResourceAreas function
type GetResourceAreasArgs struct {
	// (optional)
	EnterpriseName *string
	// (optional)
	OrganizationName *string
}

// [Preview API]
func (client *Client) GetResourceAreasByHost(ctx context.Context, args GetResourceAreasByHostArgs) (*[]ResourceAreaInfo, error) {
	queryParams := url.Values{}
	if args.HostId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "hostId"}
	}
	queryParams.Add("hostId", (*args.HostId).String())
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ResourceAreaInfo
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetResourceAreasByHost function
type GetResourceAreasByHostArgs struct {
	// (required)
	HostId *uuid.UUID
}

// [Preview API]
func (client *Client) DeleteServiceDefinition(ctx context.Context, args DeleteServiceDefinitionArgs) error {
	routeValues := make(map[string]string)
	if args.ServiceType == nil || *args.ServiceType == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "serviceType"}
	}
	routeValues["serviceType"] = *args.ServiceType
	if args.Identifier == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "identifier"}
	}
	routeValues["identifier"] = (*args.Identifier).String()

	locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteServiceDefinition function
type DeleteServiceDefinitionArgs struct {
	// (required)
	ServiceType *string
	// (required)
	Identifier *uuid.UUID
}

// [Preview API] Finds a given service definition.
func (client *Client) GetServiceDefinition(ctx context.Context, args GetServiceDefinitionArgs) (*ServiceDefinition, error) {
	routeValues := make(map[string]string)
	if args.ServiceType == nil || *args.ServiceType == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "serviceType"}
	}
	routeValues["serviceType"] = *args.ServiceType
	if args.Identifier == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "identifier"}
	}
	routeValues["identifier"] = (*args.Identifier).String()

	queryParams := url.Values{}
	if args.AllowFaultIn != nil {
		queryParams.Add("allowFaultIn", strconv.FormatBool(*args.AllowFaultIn))
	}
	if args.PreviewFaultIn != nil {
		queryParams.Add("previewFaultIn", strconv.FormatBool(*args.PreviewFaultIn))
	}
	locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ServiceDefinition
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetServiceDefinition function
type GetServiceDefinitionArgs struct {
	// (required)
	ServiceType *string
	// (required)
	Identifier *uuid.UUID
	// (optional) If true, we will attempt to fault in a host instance mapping if in SPS.
	AllowFaultIn *bool
	// (optional) If true, we will calculate and return a host instance mapping, but not persist it.
	PreviewFaultIn *bool
}

// [Preview API]
func (client *Client) GetServiceDefinitions(ctx context.Context, args GetServiceDefinitionsArgs) (*[]ServiceDefinition, error) {
	routeValues := make(map[string]string)
	if args.ServiceType != nil && *args.ServiceType != "" {
		routeValues["serviceType"] = *args.ServiceType
	}

	locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ServiceDefinition
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetServiceDefinitions function
type GetServiceDefinitionsArgs struct {
	// (optional)
	ServiceType *string
}

// [Preview API]
func (client *Client) UpdateServiceDefinitions(ctx context.Context, args UpdateServiceDefinitionsArgs) error {
	if args.ServiceDefinitions == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "serviceDefinitions"}
	}
	body, marshalErr := json.Marshal(*args.ServiceDefinitions)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdateServiceDefinitions function
type UpdateServiceDefinitionsArgs struct {
	// (required)
	ServiceDefinitions *azureDevOps.VssJsonCollectionWrapper
}
