// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package serviceendpoint

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var ResourceAreaId, _ = uuid.Parse("1814ab31-2f4f-4a9f-8761-f4d77dc5a5d7")

type Client interface {
	// [Preview API] Create a service endpoint.
	CreateServiceEndpoint(context.Context, CreateServiceEndpointArgs) (*ServiceEndpoint, error)
	// [Preview API] Delete a service endpoint.
	DeleteServiceEndpoint(context.Context, DeleteServiceEndpointArgs) error
	// [Preview API] Proxy for a GET request defined by a service endpoint.
	ExecuteServiceEndpointRequest(context.Context, ExecuteServiceEndpointRequestArgs) (*ServiceEndpointRequestResult, error)
	// [Preview API] Get the service endpoint details.
	GetServiceEndpointDetails(context.Context, GetServiceEndpointDetailsArgs) (*ServiceEndpoint, error)
	// [Preview API] Get service endpoint execution records.
	GetServiceEndpointExecutionRecords(context.Context, GetServiceEndpointExecutionRecordsArgs) (*GetServiceEndpointExecutionRecordsResponseValue, error)
	// [Preview API] Get the service endpoints.
	GetServiceEndpoints(context.Context, GetServiceEndpointsArgs) (*[]ServiceEndpoint, error)
	// [Preview API] Get the service endpoints by name.
	GetServiceEndpointsByNames(context.Context, GetServiceEndpointsByNamesArgs) (*[]ServiceEndpoint, error)
	// [Preview API] Get service endpoint types.
	GetServiceEndpointTypes(context.Context, GetServiceEndpointTypesArgs) (*[]ServiceEndpointType, error)
	// [Preview API] Update a service endpoint.
	UpdateServiceEndpoint(context.Context, UpdateServiceEndpointArgs) (*ServiceEndpoint, error)
	// [Preview API] Update the service endpoints.
	UpdateServiceEndpoints(context.Context, UpdateServiceEndpointsArgs) (*[]ServiceEndpoint, error)
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

// [Preview API] Create a service endpoint.
func (client *ClientImpl) CreateServiceEndpoint(ctx context.Context, args CreateServiceEndpointArgs) (*ServiceEndpoint, error) {
	if args.Endpoint == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Endpoint"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.Endpoint)
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

// Arguments for the CreateServiceEndpoint function
type CreateServiceEndpointArgs struct {
	// (required) Service endpoint to create.
	Endpoint *ServiceEndpoint
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Delete a service endpoint.
func (client *ClientImpl) DeleteServiceEndpoint(ctx context.Context, args DeleteServiceEndpointArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.EndpointId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.EndpointId"}
	}
	routeValues["endpointId"] = (*args.EndpointId).String()

	queryParams := url.Values{}
	if args.Deep != nil {
		queryParams.Add("deep", strconv.FormatBool(*args.Deep))
	}
	locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteServiceEndpoint function
type DeleteServiceEndpointArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the service endpoint to delete.
	EndpointId *uuid.UUID
	// (optional) Specific to AzureRM endpoint created in Automatic flow. When set to true, this will also delete corresponding AAD application in Azure. Default value is true.
	Deep *bool
}

// [Preview API] Proxy for a GET request defined by a service endpoint.
func (client *ClientImpl) ExecuteServiceEndpointRequest(ctx context.Context, args ExecuteServiceEndpointRequestArgs) (*ServiceEndpointRequestResult, error) {
	if args.ServiceEndpointRequest == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ServiceEndpointRequest"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.EndpointId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "endpointId"}
	}
	queryParams.Add("endpointId", *args.EndpointId)
	body, marshalErr := json.Marshal(*args.ServiceEndpointRequest)
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

// Arguments for the ExecuteServiceEndpointRequest function
type ExecuteServiceEndpointRequestArgs struct {
	// (required) Service endpoint request.
	ServiceEndpointRequest *ServiceEndpointRequest
	// (required) Project ID or project name
	Project *string
	// (required) Id of the service endpoint.
	EndpointId *string
}

// [Preview API] Get the service endpoint details.
func (client *ClientImpl) GetServiceEndpointDetails(ctx context.Context, args GetServiceEndpointDetailsArgs) (*ServiceEndpoint, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.EndpointId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.EndpointId"}
	}
	routeValues["endpointId"] = (*args.EndpointId).String()

	locationId, _ := uuid.Parse("e85f1c62-adfc-4b74-b618-11a150fb195e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ServiceEndpoint
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetServiceEndpointDetails function
type GetServiceEndpointDetailsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the service endpoint.
	EndpointId *uuid.UUID
}

// [Preview API] Get service endpoint execution records.
func (client *ClientImpl) GetServiceEndpointExecutionRecords(ctx context.Context, args GetServiceEndpointExecutionRecordsArgs) (*GetServiceEndpointExecutionRecordsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.EndpointId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.EndpointId"}
	}
	routeValues["endpointId"] = (*args.EndpointId).String()

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", strconv.FormatUint(*args.ContinuationToken, 10))
	}
	locationId, _ := uuid.Parse("10a16738-9299-4cd1-9a81-fd23ad6200d0")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetServiceEndpointExecutionRecordsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetServiceEndpointExecutionRecords function
type GetServiceEndpointExecutionRecordsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the service endpoint.
	EndpointId *uuid.UUID
	// (optional) Number of service endpoint execution records to get.
	Top *int
	// (optional) A continuation token, returned by a previous call to this method, that can be used to return the next set of records
	ContinuationToken *uint64
}

// Return type for the GetServiceEndpointExecutionRecords function
type GetServiceEndpointExecutionRecordsResponseValue struct {
	Value []ServiceEndpointExecutionRecord
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// [Preview API] Get the service endpoints.
func (client *ClientImpl) GetServiceEndpoints(ctx context.Context, args GetServiceEndpointsArgs) (*[]ServiceEndpoint, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Type != nil {
		queryParams.Add("type", *args.Type)
	}
	if args.AuthSchemes != nil {
		listAsString := strings.Join((*args.AuthSchemes)[:], ",")
		queryParams.Add("authSchemes", listAsString)
	}
	if args.EndpointIds != nil {
		var stringList []string
		for _, item := range *args.EndpointIds {
			stringList = append(stringList, item.String())
		}
		listAsString := strings.Join((stringList)[:], ",")
		queryParams.Add("endpointIds", listAsString)
	}
	if args.Owner != nil {
		queryParams.Add("owner", *args.Owner)
	}
	if args.IncludeFailed != nil {
		queryParams.Add("includeFailed", strconv.FormatBool(*args.IncludeFailed))
	}
	if args.IncludeDetails != nil {
		queryParams.Add("includeDetails", strconv.FormatBool(*args.IncludeDetails))
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

// Arguments for the GetServiceEndpoints function
type GetServiceEndpointsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Type of the service endpoints.
	Type *string
	// (optional) Authorization schemes used for service endpoints.
	AuthSchemes *[]string
	// (optional) Ids of the service endpoints.
	EndpointIds *[]uuid.UUID
	// (optional) Owner for service endpoints.
	Owner *string
	// (optional) Failed flag for service endpoints.
	IncludeFailed *bool
	// (optional) Flag to include more details for service endpoints. This is for internal use only and the flag will be treated as false for all other requests
	IncludeDetails *bool
}

// [Preview API] Get the service endpoints by name.
func (client *ClientImpl) GetServiceEndpointsByNames(ctx context.Context, args GetServiceEndpointsByNamesArgs) (*[]ServiceEndpoint, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.EndpointNames == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "endpointNames"}
	}
	listAsString := strings.Join((*args.EndpointNames)[:], ",")
	queryParams.Add("endpointNames", listAsString)
	if args.Type != nil {
		queryParams.Add("type", *args.Type)
	}
	if args.AuthSchemes != nil {
		listAsString := strings.Join((*args.AuthSchemes)[:], ",")
		queryParams.Add("authSchemes", listAsString)
	}
	if args.Owner != nil {
		queryParams.Add("owner", *args.Owner)
	}
	if args.IncludeFailed != nil {
		queryParams.Add("includeFailed", strconv.FormatBool(*args.IncludeFailed))
	}
	if args.IncludeDetails != nil {
		queryParams.Add("includeDetails", strconv.FormatBool(*args.IncludeDetails))
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

// Arguments for the GetServiceEndpointsByNames function
type GetServiceEndpointsByNamesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Names of the service endpoints.
	EndpointNames *[]string
	// (optional) Type of the service endpoints.
	Type *string
	// (optional) Authorization schemes used for service endpoints.
	AuthSchemes *[]string
	// (optional) Owner for service endpoints.
	Owner *string
	// (optional) Failed flag for service endpoints.
	IncludeFailed *bool
	// (optional) Flag to include more details for service endpoints. This is for internal use only and the flag will be treated as false for all other requests
	IncludeDetails *bool
}

// [Preview API] Get service endpoint types.
func (client *ClientImpl) GetServiceEndpointTypes(ctx context.Context, args GetServiceEndpointTypesArgs) (*[]ServiceEndpointType, error) {
	queryParams := url.Values{}
	if args.Type != nil {
		queryParams.Add("type", *args.Type)
	}
	if args.Scheme != nil {
		queryParams.Add("scheme", *args.Scheme)
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

// Arguments for the GetServiceEndpointTypes function
type GetServiceEndpointTypesArgs struct {
	// (optional) Type of service endpoint.
	Type *string
	// (optional) Scheme of service endpoint.
	Scheme *string
}

// [Preview API] Update a service endpoint.
func (client *ClientImpl) UpdateServiceEndpoint(ctx context.Context, args UpdateServiceEndpointArgs) (*ServiceEndpoint, error) {
	if args.Endpoint == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Endpoint"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.EndpointId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.EndpointId"}
	}
	routeValues["endpointId"] = (*args.EndpointId).String()

	queryParams := url.Values{}
	if args.Operation != nil {
		queryParams.Add("operation", *args.Operation)
	}
	body, marshalErr := json.Marshal(*args.Endpoint)
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

// Arguments for the UpdateServiceEndpoint function
type UpdateServiceEndpointArgs struct {
	// (required) Service endpoint to update.
	Endpoint *ServiceEndpoint
	// (required) Project ID or project name
	Project *string
	// (required) Id of the service endpoint to update.
	EndpointId *uuid.UUID
	// (optional) Operation for the service endpoint.
	Operation *string
}

// [Preview API] Update the service endpoints.
func (client *ClientImpl) UpdateServiceEndpoints(ctx context.Context, args UpdateServiceEndpointsArgs) (*[]ServiceEndpoint, error) {
	if args.Endpoints == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Endpoints"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.Endpoints)
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

// Arguments for the UpdateServiceEndpoints function
type UpdateServiceEndpointsArgs struct {
	// (required) Names of the service endpoints to update.
	Endpoints *[]ServiceEndpoint
	// (required) Project ID or project name
	Project *string
}
