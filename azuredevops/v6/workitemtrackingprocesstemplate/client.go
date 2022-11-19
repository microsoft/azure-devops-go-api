// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package workitemtrackingprocesstemplate

import (
	"context"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("5264459e-e5e0-4bd8-b118-0985e68a4ec5")

type Client interface {
	// [Preview API] Returns requested process template.
	ExportProcessTemplate(context.Context, ExportProcessTemplateArgs) (io.ReadCloser, error)
	// [Preview API] Returns a behavior for the process.
	GetBehavior(context.Context, GetBehaviorArgs) (*AdminBehavior, error)
	// [Preview API] Returns a list of behaviors for the process.
	GetBehaviors(context.Context, GetBehaviorsArgs) (*[]AdminBehavior, error)
	// [Preview API] Imports a process from zip file.
	ImportProcessTemplate(context.Context, ImportProcessTemplateArgs) (*ProcessImportResult, error)
	// [Preview API] Tells whether promote has completed for the specified promote job ID.
	ImportProcessTemplateStatus(context.Context, ImportProcessTemplateStatusArgs) (*ProcessPromoteStatus, error)
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

// [Preview API] Returns requested process template.
func (client *ClientImpl) ExportProcessTemplate(ctx context.Context, args ExportProcessTemplateArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Id == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Id"}
	}
	routeValues["id"] = (*args.Id).String()
	routeValues["action"] = "Export"

	locationId, _ := uuid.Parse("29e1f38d-9e9c-4358-86a5-cdf9896a5759")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the ExportProcessTemplate function
type ExportProcessTemplateArgs struct {
	// (required) The ID of the process
	Id *uuid.UUID
}

// [Preview API] Returns a behavior for the process.
func (client *ClientImpl) GetBehavior(ctx context.Context, args GetBehaviorArgs) (*AdminBehavior, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProcessId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()

	queryParams := url.Values{}
	if args.BehaviorRefName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "behaviorRefName"}
	}
	queryParams.Add("behaviorRefName", *args.BehaviorRefName)
	locationId, _ := uuid.Parse("90bf9317-3571-487b-bc8c-a523ba0e05d7")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AdminBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBehavior function
type GetBehaviorArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the behavior
	BehaviorRefName *string
}

// [Preview API] Returns a list of behaviors for the process.
func (client *ClientImpl) GetBehaviors(ctx context.Context, args GetBehaviorsArgs) (*[]AdminBehavior, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProcessId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()

	locationId, _ := uuid.Parse("90bf9317-3571-487b-bc8c-a523ba0e05d7")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []AdminBehavior
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBehaviors function
type GetBehaviorsArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
}

// [Preview API] Imports a process from zip file.
func (client *ClientImpl) ImportProcessTemplate(ctx context.Context, args ImportProcessTemplateArgs) (*ProcessImportResult, error) {
	if args.UploadStream == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.UploadStream"}
	}
	routeValues := make(map[string]string)
	routeValues["action"] = "Import"

	queryParams := url.Values{}
	if args.IgnoreWarnings != nil {
		queryParams.Add("ignoreWarnings", strconv.FormatBool(*args.IgnoreWarnings))
	}
	if args.ReplaceExistingTemplate != nil {
		queryParams.Add("replaceExistingTemplate", strconv.FormatBool(*args.ReplaceExistingTemplate))
	}
	locationId, _ := uuid.Parse("29e1f38d-9e9c-4358-86a5-cdf9896a5759")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, queryParams, args.UploadStream, "application/octet-stream", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessImportResult
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ImportProcessTemplate function
type ImportProcessTemplateArgs struct {
	// (required) Stream to upload
	UploadStream io.Reader
	// (optional) Ignores validation warnings. Default value is false.
	IgnoreWarnings *bool
	// (optional) Replaces the existing template. Default value is true.
	ReplaceExistingTemplate *bool
}

// [Preview API] Tells whether promote has completed for the specified promote job ID.
func (client *ClientImpl) ImportProcessTemplateStatus(ctx context.Context, args ImportProcessTemplateStatusArgs) (*ProcessPromoteStatus, error) {
	routeValues := make(map[string]string)
	if args.Id == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Id"}
	}
	routeValues["id"] = (*args.Id).String()
	routeValues["action"] = "Status"

	locationId, _ := uuid.Parse("29e1f38d-9e9c-4358-86a5-cdf9896a5759")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessPromoteStatus
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ImportProcessTemplateStatus function
type ImportProcessTemplateStatusArgs struct {
	// (required) The ID of the promote job operation
	Id *uuid.UUID
}
