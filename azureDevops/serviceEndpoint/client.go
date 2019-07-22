// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package serviceEndpoint

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
    "strings"
)

var ResourceAreaId, _ = uuid.Parse("1814ab31-2f4f-4a9f-8761-f4d77dc5a5d7")

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API] Proxy for a GET request defined by a service endpoint.
// ctx
// serviceEndpointRequest (required): Service endpoint request.
// project (required): Project ID or project name
// endpointId (required): Id of the service endpoint.
func (client Client) ExecuteServiceEndpointRequest(ctx context.Context, serviceEndpointRequest *ServiceEndpointRequest, project *string, endpointId *string) (*ServiceEndpointRequestResult, error) {
    if serviceEndpointRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "serviceEndpointRequest"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if endpointId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpointId"}
    }
    queryParams.Add("endpointId", *endpointId)
    body, marshalErr := json.Marshal(*serviceEndpointRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cc63bb57-2a5f-4a7a-b79c-c142d308657e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ServiceEndpointRequestResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a service endpoint.
// ctx
// endpoint (required): Service endpoint to create.
// project (required): Project ID or project name
func (client Client) CreateServiceEndpoint(ctx context.Context, endpoint *ServiceEndpoint, project *string) (*ServiceEndpoint, error) {
    if endpoint == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpoint"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*endpoint)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ServiceEndpoint
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a service endpoint.
// ctx
// project (required): Project ID or project name
// endpointId (required): Id of the service endpoint to delete.
// deep (optional): Specific to AzureRM endpoint created in Automatic flow. When set to true, this will also delete corresponding AAD application in Azure. Default value is true.
func (client Client) DeleteServiceEndpoint(ctx context.Context, project *string, endpointId *uuid.UUID, deep *bool) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if endpointId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "endpointId"} 
    }
    routeValues["endpointId"] = (*endpointId).String()

    queryParams := url.Values{}
    if deep != nil {
        queryParams.Add("deep", strconv.FormatBool(*deep))
    }
    locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the service endpoint details.
// ctx
// project (required): Project ID or project name
// endpointId (required): Id of the service endpoint.
func (client Client) GetServiceEndpointDetails(ctx context.Context, project *string, endpointId *uuid.UUID) (*ServiceEndpoint, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if endpointId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpointId"} 
    }
    routeValues["endpointId"] = (*endpointId).String()

    locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ServiceEndpoint
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the service endpoints.
// ctx
// project (required): Project ID or project name
// type_ (optional): Type of the service endpoints.
// authSchemes (optional): Authorization schemes used for service endpoints.
// endpointIds (optional): Ids of the service endpoints.
// owner (optional): Owner for service endpoints.
// includeFailed (optional): Failed flag for service endpoints.
// includeDetails (optional): Flag to include more details for service endpoints. This is for internal use only and the flag will be treated as false for all other requests
func (client Client) GetServiceEndpoints(ctx context.Context, project *string, type_ *string, authSchemes *[]string, endpointIds *[]uuid.UUID, owner *string, includeFailed *bool, includeDetails *bool) (*[]ServiceEndpoint, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    if authSchemes != nil {
        listAsString := strings.Join((*authSchemes)[:], ",")
        queryParams.Add("authSchemes", listAsString)
    }
    if endpointIds != nil {
        var stringList []string
        for _, item := range *endpointIds {
            stringList = append(stringList, item.String())
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if owner != nil {
        queryParams.Add("owner", *owner)
    }
    if includeFailed != nil {
        queryParams.Add("includeFailed", strconv.FormatBool(*includeFailed))
    }
    if includeDetails != nil {
        queryParams.Add("includeDetails", strconv.FormatBool(*includeDetails))
    }
    locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ServiceEndpoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the service endpoints by name.
// ctx
// project (required): Project ID or project name
// endpointNames (required): Names of the service endpoints.
// type_ (optional): Type of the service endpoints.
// authSchemes (optional): Authorization schemes used for service endpoints.
// owner (optional): Owner for service endpoints.
// includeFailed (optional): Failed flag for service endpoints.
// includeDetails (optional): Flag to include more details for service endpoints. This is for internal use only and the flag will be treated as false for all other requests
func (client Client) GetServiceEndpointsByNames(ctx context.Context, project *string, endpointNames *[]string, type_ *string, authSchemes *[]string, owner *string, includeFailed *bool, includeDetails *bool) (*[]ServiceEndpoint, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if endpointNames == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpointNames"}
    }
    listAsString := strings.Join((*endpointNames)[:], ",")
    queryParams.Add("endpointNames", listAsString)
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    if authSchemes != nil {
        listAsString := strings.Join((*authSchemes)[:], ",")
        queryParams.Add("authSchemes", listAsString)
    }
    if owner != nil {
        queryParams.Add("owner", *owner)
    }
    if includeFailed != nil {
        queryParams.Add("includeFailed", strconv.FormatBool(*includeFailed))
    }
    if includeDetails != nil {
        queryParams.Add("includeDetails", strconv.FormatBool(*includeDetails))
    }
    locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ServiceEndpoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a service endpoint.
// ctx
// endpoint (required): Service endpoint to update.
// project (required): Project ID or project name
// endpointId (required): Id of the service endpoint to update.
// operation (optional): Operation for the service endpoint.
func (client Client) UpdateServiceEndpoint(ctx context.Context, endpoint *ServiceEndpoint, project *string, endpointId *uuid.UUID, operation *string) (*ServiceEndpoint, error) {
    if endpoint == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpoint"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if endpointId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpointId"} 
    }
    routeValues["endpointId"] = (*endpointId).String()

    queryParams := url.Values{}
    if operation != nil {
        queryParams.Add("operation", *operation)
    }
    body, marshalErr := json.Marshal(*endpoint)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ServiceEndpoint
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update the service endpoints.
// ctx
// endpoints (required): Names of the service endpoints to update.
// project (required): Project ID or project name
func (client Client) UpdateServiceEndpoints(ctx context.Context, endpoints *[]ServiceEndpoint, project *string) (*[]ServiceEndpoint, error) {
    if endpoints == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpoints"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*endpoints)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ServiceEndpoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get service endpoint execution records.
// ctx
// project (required): Project ID or project name
// endpointId (required): Id of the service endpoint.
// top (optional): Number of service endpoint execution records to get.
// continuationToken (optional): A continuation token, returned by a previous call to this method, that can be used to return the next set of records
func (client Client) GetServiceEndpointExecutionRecords(ctx context.Context, project *string, endpointId *uuid.UUID, top *int, continuationToken *uint64) (*[]ServiceEndpointExecutionRecord, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if endpointId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "endpointId"} 
    }
    routeValues["endpointId"] = (*endpointId).String()

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", strconv.FormatUint(*continuationToken, 10))
    }
    locationId, _ := uuid.Parse("10a16738-9299-4cd1-9a81-fd23ad6200d0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ServiceEndpointExecutionRecord
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get service endpoint types.
// ctx
// type_ (optional): Type of service endpoint.
// scheme (optional): Scheme of service endpoint.
func (client Client) GetServiceEndpointTypes(ctx context.Context, type_ *string, scheme *string) (*[]ServiceEndpointType, error) {
    queryParams := url.Values{}
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    if scheme != nil {
        queryParams.Add("scheme", *scheme)
    }
    locationId, _ := uuid.Parse("5a7938a4-655e-486c-b562-b78c54a7e87b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ServiceEndpointType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

