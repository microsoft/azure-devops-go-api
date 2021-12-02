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
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("af607f94-69ba-4821-8159-f04e37b66350")

type Client interface {
	// [Preview API] Check the availability of symbol service. This includes checking for feature flag, and possibly license in future. Note this is NOT an anonymous endpoint, and the caller will be redirected to authentication before hitting it.
	CheckAvailability(context.Context, CheckAvailabilityArgs) error
	// [Preview API] Create a new symbol request.
	CreateRequests(context.Context, CreateRequestsArgs) (*Request, error)
	// [Preview API] Create debug entries for a symbol request as specified by its identifier.
	CreateRequestsRequestIdDebugEntries(context.Context, CreateRequestsRequestIdDebugEntriesArgs) (*[]DebugEntry, error)
	// [Preview API] Create debug entries for a symbol request as specified by its name.
	CreateRequestsRequestNameDebugEntries(context.Context, CreateRequestsRequestNameDebugEntriesArgs) (*[]DebugEntry, error)
	// [Preview API] Delete a symbol request by request identifier.
	DeleteRequestsRequestId(context.Context, DeleteRequestsRequestIdArgs) error
	// [Preview API] Delete a symbol request by request name.
	DeleteRequestsRequestName(context.Context, DeleteRequestsRequestNameArgs) error
	// [Preview API] Get the client package.
	GetClient(context.Context, GetClientArgs) (interface{}, error)
	// [Preview API] Get a symbol request by request identifier.
	GetRequestsRequestId(context.Context, GetRequestsRequestIdArgs) (*Request, error)
	// [Preview API] Get a symbol request by request name.
	GetRequestsRequestName(context.Context, GetRequestsRequestNameArgs) (*Request, error)
	// [Preview API] Given a client key, returns the best matched debug entry.
	GetSymSrvDebugEntryClientKey(context.Context, GetSymSrvDebugEntryClientKeyArgs) error
	// [Preview API] Get client version information.
	HeadClient(context.Context, HeadClientArgs) error
	// [Preview API] Update a symbol request by request identifier.
	UpdateRequestsRequestId(context.Context, UpdateRequestsRequestIdArgs) (*Request, error)
	// [Preview API] Update a symbol request by request name.
	UpdateRequestsRequestName(context.Context, UpdateRequestsRequestNameArgs) (*Request, error)
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

// [Preview API] Check the availability of symbol service. This includes checking for feature flag, and possibly license in future. Note this is NOT an anonymous endpoint, and the caller will be redirected to authentication before hitting it.
func (client *ClientImpl) CheckAvailability(ctx context.Context, args CheckAvailabilityArgs) error {
	locationId, _ := uuid.Parse("97c893cc-e861-4ef4-8c43-9bad4a963dee")
	_, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the CheckAvailability function
type CheckAvailabilityArgs struct {
}

// [Preview API] Create a new symbol request.
func (client *ClientImpl) CreateRequests(ctx context.Context, args CreateRequestsArgs) (*Request, error) {
	if args.RequestToCreate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RequestToCreate"}
	}
	body, marshalErr := json.Marshal(*args.RequestToCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Request
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateRequests function
type CreateRequestsArgs struct {
	// (required) The symbol request to create.
	RequestToCreate *Request
}

// [Preview API] Create debug entries for a symbol request as specified by its identifier.
func (client *ClientImpl) CreateRequestsRequestIdDebugEntries(ctx context.Context, args CreateRequestsRequestIdDebugEntriesArgs) (*[]DebugEntry, error) {
	if args.Batch == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Batch"}
	}
	routeValues := make(map[string]string)
	if args.RequestId == nil || *args.RequestId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RequestId"}
	}
	routeValues["requestId"] = *args.RequestId

	queryParams := url.Values{}
	if args.Collection == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "collection"}
	}
	queryParams.Add("collection", *args.Collection)
	body, marshalErr := json.Marshal(*args.Batch)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []DebugEntry
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateRequestsRequestIdDebugEntries function
type CreateRequestsRequestIdDebugEntriesArgs struct {
	// (required) A batch that contains debug entries to create.
	Batch *DebugEntryCreateBatch
	// (required) The symbol request identifier.
	RequestId *string
	// (required) A valid debug entry collection name. Must be "debugentries".
	Collection *string
}

// [Preview API] Create debug entries for a symbol request as specified by its name.
func (client *ClientImpl) CreateRequestsRequestNameDebugEntries(ctx context.Context, args CreateRequestsRequestNameDebugEntriesArgs) (*[]DebugEntry, error) {
	if args.Batch == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Batch"}
	}
	queryParams := url.Values{}
	if args.RequestName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "requestName"}
	}
	queryParams.Add("requestName", *args.RequestName)
	if args.Collection == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "collection"}
	}
	queryParams.Add("collection", *args.Collection)
	body, marshalErr := json.Marshal(*args.Batch)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []DebugEntry
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateRequestsRequestNameDebugEntries function
type CreateRequestsRequestNameDebugEntriesArgs struct {
	// (required) A batch that contains debug entries to create.
	Batch *DebugEntryCreateBatch
	// (required) The symbol request name.
	RequestName *string
	// (required) A valid debug entry collection name. Must be "debugentries".
	Collection *string
}

// [Preview API] Delete a symbol request by request identifier.
func (client *ClientImpl) DeleteRequestsRequestId(ctx context.Context, args DeleteRequestsRequestIdArgs) error {
	routeValues := make(map[string]string)
	if args.RequestId == nil || *args.RequestId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RequestId"}
	}
	routeValues["requestId"] = *args.RequestId

	queryParams := url.Values{}
	if args.Synchronous != nil {
		queryParams.Add("synchronous", strconv.FormatBool(*args.Synchronous))
	}
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteRequestsRequestId function
type DeleteRequestsRequestIdArgs struct {
	// (required) The symbol request identifier.
	RequestId *string
	// (optional) If true, delete all the debug entries under this request synchronously in the current session. If false, the deletion will be postponed to a later point and be executed automatically by the system.
	Synchronous *bool
}

// [Preview API] Delete a symbol request by request name.
func (client *ClientImpl) DeleteRequestsRequestName(ctx context.Context, args DeleteRequestsRequestNameArgs) error {
	queryParams := url.Values{}
	if args.RequestName == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "requestName"}
	}
	queryParams.Add("requestName", *args.RequestName)
	if args.Synchronous != nil {
		queryParams.Add("synchronous", strconv.FormatBool(*args.Synchronous))
	}
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteRequestsRequestName function
type DeleteRequestsRequestNameArgs struct {
	// (required) The symbol request name.
	RequestName *string
	// (optional) If true, delete all the debug entries under this request synchronously in the current session. If false, the deletion will be postponed to a later point and be executed automatically by the system.
	Synchronous *bool
}

// [Preview API] Get the client package.
func (client *ClientImpl) GetClient(ctx context.Context, args GetClientArgs) (interface{}, error) {
	routeValues := make(map[string]string)
	if args.ClientType == nil || *args.ClientType == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ClientType"}
	}
	routeValues["clientType"] = *args.ClientType

	locationId, _ := uuid.Parse("79c83865-4de3-460c-8a16-01be238e0818")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the GetClient function
type GetClientArgs struct {
	// (required) Either "EXE" for a zip file containing a Windows symbol client (a.k.a. symbol.exe) along with dependencies, or "TASK" for a VSTS task that can be run on a VSTS build agent. All the other values are invalid. The parameter is case-insensitive.
	ClientType *string
}

// [Preview API] Get a symbol request by request identifier.
func (client *ClientImpl) GetRequestsRequestId(ctx context.Context, args GetRequestsRequestIdArgs) (*Request, error) {
	routeValues := make(map[string]string)
	if args.RequestId == nil || *args.RequestId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RequestId"}
	}
	routeValues["requestId"] = *args.RequestId

	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Request
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRequestsRequestId function
type GetRequestsRequestIdArgs struct {
	// (required) The symbol request identifier.
	RequestId *string
}

// [Preview API] Get a symbol request by request name.
func (client *ClientImpl) GetRequestsRequestName(ctx context.Context, args GetRequestsRequestNameArgs) (*Request, error) {
	queryParams := url.Values{}
	if args.RequestName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "requestName"}
	}
	queryParams.Add("requestName", *args.RequestName)
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Request
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRequestsRequestName function
type GetRequestsRequestNameArgs struct {
	// (required) The symbol request name.
	RequestName *string
}

// [Preview API] Given a client key, returns the best matched debug entry.
func (client *ClientImpl) GetSymSrvDebugEntryClientKey(ctx context.Context, args GetSymSrvDebugEntryClientKeyArgs) error {
	routeValues := make(map[string]string)
	if args.DebugEntryClientKey == nil || *args.DebugEntryClientKey == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.DebugEntryClientKey"}
	}
	routeValues["debugEntryClientKey"] = *args.DebugEntryClientKey

	locationId, _ := uuid.Parse("9648e256-c9f9-4f16-8a27-630b06396942")
	_, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the GetSymSrvDebugEntryClientKey function
type GetSymSrvDebugEntryClientKeyArgs struct {
	// (required) A "client key" used by both ends of Microsoft's symbol protocol to identify a debug entry. The semantics of client key is governed by symsrv and is beyond the scope of this documentation.
	DebugEntryClientKey *string
}

// [Preview API] Get client version information.
func (client *ClientImpl) HeadClient(ctx context.Context, args HeadClientArgs) error {
	locationId, _ := uuid.Parse("79c83865-4de3-460c-8a16-01be238e0818")
	_, err := client.Client.Send(ctx, http.MethodHead, locationId, "6.0-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the HeadClient function
type HeadClientArgs struct {
}

// [Preview API] Update a symbol request by request identifier.
func (client *ClientImpl) UpdateRequestsRequestId(ctx context.Context, args UpdateRequestsRequestIdArgs) (*Request, error) {
	if args.UpdateRequest == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.UpdateRequest"}
	}
	routeValues := make(map[string]string)
	if args.RequestId == nil || *args.RequestId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RequestId"}
	}
	routeValues["requestId"] = *args.RequestId

	body, marshalErr := json.Marshal(*args.UpdateRequest)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Request
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateRequestsRequestId function
type UpdateRequestsRequestIdArgs struct {
	// (required) The symbol request.
	UpdateRequest *Request
	// (required) The symbol request identifier.
	RequestId *string
}

// [Preview API] Update a symbol request by request name.
func (client *ClientImpl) UpdateRequestsRequestName(ctx context.Context, args UpdateRequestsRequestNameArgs) (*Request, error) {
	if args.UpdateRequest == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.UpdateRequest"}
	}
	queryParams := url.Values{}
	if args.RequestName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "requestName"}
	}
	queryParams.Add("requestName", *args.RequestName)
	body, marshalErr := json.Marshal(*args.UpdateRequest)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ebc09fe3-1b20-4667-abc5-f2b60fe8de52")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Request
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateRequestsRequestName function
type UpdateRequestsRequestNameArgs struct {
	// (required) The symbol request.
	UpdateRequest *Request
	// (required) The symbol request name.
	RequestName *string
}
