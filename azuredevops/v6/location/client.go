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
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/webapi"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	// [Preview API]
	DeleteServiceDefinition(context.Context, DeleteServiceDefinitionArgs) error
	// [Preview API] This was copied and adapted from TeamFoundationConnectionService.Connect()
	GetConnectionData(context.Context, GetConnectionDataArgs) (*ConnectionData, error)
	// [Preview API]
	GetResourceArea(context.Context, GetResourceAreaArgs) (*ResourceAreaInfo, error)
	// [Preview API]
	GetResourceAreaByHost(context.Context, GetResourceAreaByHostArgs) (*ResourceAreaInfo, error)
	// [Preview API]
	GetResourceAreas(context.Context, GetResourceAreasArgs) (*[]ResourceAreaInfo, error)
	// [Preview API]
	GetResourceAreasByHost(context.Context, GetResourceAreasByHostArgs) (*[]ResourceAreaInfo, error)
	// [Preview API] Finds a given service definition.
	GetServiceDefinition(context.Context, GetServiceDefinitionArgs) (*ServiceDefinition, error)
	// [Preview API]
	GetServiceDefinitions(context.Context, GetServiceDefinitionsArgs) (*[]ServiceDefinition, error)
	// [Preview API]
	UpdateServiceDefinitions(context.Context, UpdateServiceDefinitionsArgs) error
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
func (client *ClientImpl) DeleteServiceDefinition(ctx context.Context, args DeleteServiceDefinitionArgs) error {
	routeValues := make(map[string]string)
	if args.ServiceType == nil || *args.ServiceType == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ServiceType"}
	}
	routeValues["serviceType"] = *args.ServiceType
	if args.Identifier == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.Identifier"}
	}
	routeValues["identifier"] = (*args.Identifier).String()

	locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
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

// [Preview API] This was copied and adapted from TeamFoundationConnectionService.Connect()
func (client *ClientImpl) GetConnectionData(ctx context.Context, args GetConnectionDataArgs) (*ConnectionData, error) {
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, queryParams, nil, "", "application/json", nil)
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
	ConnectOptions *webapi.ConnectOptions
	// (optional) Obsolete 32-bit LastChangeId
	LastChangeId *int
	// (optional) Non-truncated 64-bit LastChangeId
	LastChangeId64 *uint64
}

// [Preview API]
func (client *ClientImpl) GetResourceArea(ctx context.Context, args GetResourceAreaArgs) (*ResourceAreaInfo, error) {
	routeValues := make(map[string]string)
	if args.AreaId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AreaId"}
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
func (client *ClientImpl) GetResourceAreaByHost(ctx context.Context, args GetResourceAreaByHostArgs) (*ResourceAreaInfo, error) {
	routeValues := make(map[string]string)
	if args.AreaId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AreaId"}
	}
	routeValues["areaId"] = (*args.AreaId).String()

	queryParams := url.Values{}
	if args.HostId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "hostId"}
	}
	queryParams.Add("hostId", (*args.HostId).String())
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
func (client *ClientImpl) GetResourceAreas(ctx context.Context, args GetResourceAreasArgs) (*[]ResourceAreaInfo, error) {
	queryParams := url.Values{}
	if args.EnterpriseName != nil {
		queryParams.Add("enterpriseName", *args.EnterpriseName)
	}
	if args.OrganizationName != nil {
		queryParams.Add("organizationName", *args.OrganizationName)
	}
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, queryParams, nil, "", "application/json", nil)
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
func (client *ClientImpl) GetResourceAreasByHost(ctx context.Context, args GetResourceAreasByHostArgs) (*[]ResourceAreaInfo, error) {
	queryParams := url.Values{}
	if args.HostId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "hostId"}
	}
	queryParams.Add("hostId", (*args.HostId).String())
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, queryParams, nil, "", "application/json", nil)
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

// [Preview API] Finds a given service definition.
func (client *ClientImpl) GetServiceDefinition(ctx context.Context, args GetServiceDefinitionArgs) (*ServiceDefinition, error) {
	routeValues := make(map[string]string)
	if args.ServiceType == nil || *args.ServiceType == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ServiceType"}
	}
	routeValues["serviceType"] = *args.ServiceType
	if args.Identifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Identifier"}
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
func (client *ClientImpl) GetServiceDefinitions(ctx context.Context, args GetServiceDefinitionsArgs) (*[]ServiceDefinition, error) {
	routeValues := make(map[string]string)
	if args.ServiceType != nil && *args.ServiceType != "" {
		routeValues["serviceType"] = *args.ServiceType
	}

	locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
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
func (client *ClientImpl) UpdateServiceDefinitions(ctx context.Context, args UpdateServiceDefinitionsArgs) error {
	if args.ServiceDefinitions == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ServiceDefinitions"}
	}
	body, marshalErr := json.Marshal(*args.ServiceDefinitions)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdateServiceDefinitions function
type UpdateServiceDefinitionsArgs struct {
	// (required)
	ServiceDefinitions *azuredevops.VssJsonCollectionWrapper
}
