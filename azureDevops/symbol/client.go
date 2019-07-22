// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package symbol

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

var ResourceAreaId, _ = uuid.Parse("af607f94-69ba-4821-8159-f04e37b66350")

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

// [Preview API] Check the availability of symbol service. This includes checking for feature flag, and possibly license in future. Note this is NOT an anonymous endpoint, and the caller will be redirected to authentication before hitting it.
// ctx
func (client Client) CheckAvailability(ctx context.Context, ) error {
    locationId, _ := uuid.Parse("97c893cc-e861-4ef4-8c43-9bad4a963dee")
    _, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the client package.
// ctx
// clientType (required): Either "EXE" for a zip file containing a Windows symbol client (a.k.a. symbol.exe) along with dependencies, or "TASK" for a VSTS task that can be run on a VSTS build agent. All the other values are invalid. The parameter is case-insensitive.
func (client Client) GetClient(ctx context.Context, clientType *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if clientType == nil || *clientType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "clientType"} 
    }
    routeValues["clientType"] = *clientType

    locationId, _ := uuid.Parse("79c83865-4de3-460c-8a16-01be238e0818")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get client version information.
// ctx
func (client Client) HeadClient(ctx context.Context, ) error {
    locationId, _ := uuid.Parse("79c83865-4de3-460c-8a16-01be238e0818")
    _, err := client.Client.Send(ctx, http.MethodHead, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Create a new symbol request.
// ctx
// requestToCreate (required): The symbol request to create.
func (client Client) CreateRequests(ctx context.Context, requestToCreate *Request) (*Request, error) {
    if requestToCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "requestToCreate"}
    }
    body, marshalErr := json.Marshal(*requestToCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Request
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create debug entries for a symbol request as specified by its identifier.
// ctx
// batch (required): A batch that contains debug entries to create.
// requestId (required): The symbol request identifier.
// collection (required): A valid debug entry collection name. Must be "debugentries".
func (client Client) CreateRequestsRequestIdDebugEntries(ctx context.Context, batch *DebugEntryCreateBatch, requestId *string, collection *string) (*[]DebugEntry, error) {
    if batch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "batch"}
    }
    routeValues := make(map[string]string)
    if requestId == nil || *requestId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "requestId"} 
    }
    routeValues["requestId"] = *requestId

    queryParams := url.Values{}
    if collection == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "collection"}
    }
    queryParams.Add("collection", *collection)
    body, marshalErr := json.Marshal(*batch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DebugEntry
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create debug entries for a symbol request as specified by its name.
// ctx
// batch (required): A batch that contains debug entries to create.
// requestName (required): The symbol request name.
// collection (required): A valid debug entry collection name. Must be "debugentries".
func (client Client) CreateRequestsRequestNameDebugEntries(ctx context.Context, batch *DebugEntryCreateBatch, requestName *string, collection *string) (*[]DebugEntry, error) {
    if batch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "batch"}
    }
    queryParams := url.Values{}
    if requestName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "requestName"}
    }
    queryParams.Add("requestName", *requestName)
    if collection == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "collection"}
    }
    queryParams.Add("collection", *collection)
    body, marshalErr := json.Marshal(*batch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DebugEntry
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a symbol request by request identifier.
// ctx
// requestId (required): The symbol request identifier.
// synchronous (optional): If true, delete all the debug entries under this request synchronously in the current session. If false, the deletion will be postponed to a later point and be executed automatically by the system.
func (client Client) DeleteRequestsRequestId(ctx context.Context, requestId *string, synchronous *bool) error {
    routeValues := make(map[string]string)
    if requestId == nil || *requestId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "requestId"} 
    }
    routeValues["requestId"] = *requestId

    queryParams := url.Values{}
    if synchronous != nil {
        queryParams.Add("synchronous", strconv.FormatBool(*synchronous))
    }
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Delete a symbol request by request name.
// ctx
// requestName (required): The symbol request name.
// synchronous (optional): If true, delete all the debug entries under this request synchronously in the current session. If false, the deletion will be postponed to a later point and be executed automatically by the system.
func (client Client) DeleteRequestsRequestName(ctx context.Context, requestName *string, synchronous *bool) error {
    queryParams := url.Values{}
    if requestName == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "requestName"}
    }
    queryParams.Add("requestName", *requestName)
    if synchronous != nil {
        queryParams.Add("synchronous", strconv.FormatBool(*synchronous))
    }
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a symbol request by request identifier.
// ctx
// requestId (required): The symbol request identifier.
func (client Client) GetRequestsRequestId(ctx context.Context, requestId *string) (*Request, error) {
    routeValues := make(map[string]string)
    if requestId == nil || *requestId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "requestId"} 
    }
    routeValues["requestId"] = *requestId

    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Request
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a symbol request by request name.
// ctx
// requestName (required): The symbol request name.
func (client Client) GetRequestsRequestName(ctx context.Context, requestName *string) (*Request, error) {
    queryParams := url.Values{}
    if requestName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "requestName"}
    }
    queryParams.Add("requestName", *requestName)
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Request
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a symbol request by request identifier.
// ctx
// updateRequest (required): The symbol request.
// requestId (required): The symbol request identifier.
func (client Client) UpdateRequestsRequestId(ctx context.Context, updateRequest *Request, requestId *string) (*Request, error) {
    if updateRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateRequest"}
    }
    routeValues := make(map[string]string)
    if requestId == nil || *requestId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "requestId"} 
    }
    routeValues["requestId"] = *requestId

    body, marshalErr := json.Marshal(*updateRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Request
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a symbol request by request name.
// ctx
// updateRequest (required): The symbol request.
// requestName (required): The symbol request name.
func (client Client) UpdateRequestsRequestName(ctx context.Context, updateRequest *Request, requestName *string) (*Request, error) {
    if updateRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateRequest"}
    }
    queryParams := url.Values{}
    if requestName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "requestName"}
    }
    queryParams.Add("requestName", *requestName)
    body, marshalErr := json.Marshal(*updateRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Request
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Given a client key, returns the best matched debug entry.
// ctx
// debugEntryClientKey (required): A "client key" used by both ends of Microsoft's symbol protocol to identify a debug entry. The semantics of client key is governed by symsrv and is beyond the scope of this documentation.
func (client Client) GetSymSrvDebugEntryClientKey(ctx context.Context, debugEntryClientKey *string) error {
    routeValues := make(map[string]string)
    if debugEntryClientKey == nil || *debugEntryClientKey == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "debugEntryClientKey"} 
    }
    routeValues["debugEntryClientKey"] = *debugEntryClientKey

    locationId, _ := uuid.Parse("9648e256-c9f9-4f16-8a27-630b06396942")
    _, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

