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
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API] This was copied and adapted from TeamFoundationConnectionService.Connect()
// ctx
// connectOptions (optional)
// lastChangeId (optional): Obsolete 32-bit LastChangeId
// lastChangeId64 (optional): Non-truncated 64-bit LastChangeId
func (client Client) GetConnectionData(ctx context.Context, connectOptions *string, lastChangeId *int, lastChangeId64 *uint64) (*ConnectionData, error) {
    queryParams := url.Values{}
    if connectOptions != nil {
        queryParams.Add("connectOptions", *connectOptions)
    }
    if lastChangeId != nil {
        queryParams.Add("lastChangeId", strconv.Itoa(*lastChangeId))
    }
    if lastChangeId64 != nil {
        queryParams.Add("lastChangeId64", strconv.FormatUint(*lastChangeId64, 10))
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

// [Preview API]
// ctx
// areaId (required)
// enterpriseName (optional)
// organizationName (optional)
func (client Client) GetResourceArea(ctx context.Context, areaId *uuid.UUID, enterpriseName *string, organizationName *string) (*ResourceAreaInfo, error) {
    routeValues := make(map[string]string)
    if areaId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "areaId"} 
    }
    routeValues["areaId"] = (*areaId).String()

    queryParams := url.Values{}
    if enterpriseName != nil {
        queryParams.Add("enterpriseName", *enterpriseName)
    }
    if organizationName != nil {
        queryParams.Add("organizationName", *organizationName)
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

// [Preview API]
// ctx
// areaId (required)
// hostId (required)
func (client Client) GetResourceAreaByHost(ctx context.Context, areaId *uuid.UUID, hostId *uuid.UUID) (*ResourceAreaInfo, error) {
    routeValues := make(map[string]string)
    if areaId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "areaId"} 
    }
    routeValues["areaId"] = (*areaId).String()

    queryParams := url.Values{}
    if hostId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "hostId"}
    }
    queryParams.Add("hostId", (*hostId).String())
    locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ResourceAreaInfo
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// enterpriseName (optional)
// organizationName (optional)
func (client Client) GetResourceAreas(ctx context.Context, enterpriseName *string, organizationName *string) (*[]ResourceAreaInfo, error) {
    queryParams := url.Values{}
    if enterpriseName != nil {
        queryParams.Add("enterpriseName", *enterpriseName)
    }
    if organizationName != nil {
        queryParams.Add("organizationName", *organizationName)
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

// [Preview API]
// ctx
// hostId (required)
func (client Client) GetResourceAreasByHost(ctx context.Context, hostId *uuid.UUID) (*[]ResourceAreaInfo, error) {
    queryParams := url.Values{}
    if hostId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "hostId"}
    }
    queryParams.Add("hostId", (*hostId).String())
    locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ResourceAreaInfo
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// serviceType (required)
// identifier (required)
func (client Client) DeleteServiceDefinition(ctx context.Context, serviceType *string, identifier *uuid.UUID) error {
    routeValues := make(map[string]string)
    if serviceType == nil || *serviceType == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "serviceType"} 
    }
    routeValues["serviceType"] = *serviceType
    if identifier == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "identifier"} 
    }
    routeValues["identifier"] = (*identifier).String()

    locationId, _ := uuid.Parse("d810a47d-f4f4-4a62-a03f-fa1860585c4c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Finds a given service definition.
// ctx
// serviceType (required)
// identifier (required)
// allowFaultIn (optional): If true, we will attempt to fault in a host instance mapping if in SPS.
// previewFaultIn (optional): If true, we will calculate and return a host instance mapping, but not persist it.
func (client Client) GetServiceDefinition(ctx context.Context, serviceType *string, identifier *uuid.UUID, allowFaultIn *bool, previewFaultIn *bool) (*ServiceDefinition, error) {
    routeValues := make(map[string]string)
    if serviceType == nil || *serviceType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "serviceType"} 
    }
    routeValues["serviceType"] = *serviceType
    if identifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "identifier"} 
    }
    routeValues["identifier"] = (*identifier).String()

    queryParams := url.Values{}
    if allowFaultIn != nil {
        queryParams.Add("allowFaultIn", strconv.FormatBool(*allowFaultIn))
    }
    if previewFaultIn != nil {
        queryParams.Add("previewFaultIn", strconv.FormatBool(*previewFaultIn))
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

// [Preview API]
// ctx
// serviceType (optional)
func (client Client) GetServiceDefinitions(ctx context.Context, serviceType *string) (*[]ServiceDefinition, error) {
    routeValues := make(map[string]string)
    if serviceType != nil && *serviceType != "" {
        routeValues["serviceType"] = *serviceType
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

// [Preview API]
// ctx
// serviceDefinitions (required)
func (client Client) UpdateServiceDefinitions(ctx context.Context, serviceDefinitions *azureDevops.VssJsonCollectionWrapper) error {
    if serviceDefinitions == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "serviceDefinitions"}
    }
    body, marshalErr := json.Marshal(*serviceDefinitions)
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

