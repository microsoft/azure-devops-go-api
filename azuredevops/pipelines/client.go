// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package pipelines

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	// [Preview API]
	CreatePipeline(context.Context, CreatePipelineArgs) (*Pipeline, error)
	// [Preview API]
	GetLog(context.Context, GetLogArgs) (*Log, error)
	// [Preview API] Gets a pipeline, optionally at the specified version
	GetPipeline(context.Context, GetPipelineArgs) (*Pipeline, error)
	// [Preview API] Gets a run for a particular pipeline.
	GetRun(context.Context, GetRunArgs) (*Run, error)
	// [Preview API]
	ListLogs(context.Context, ListLogsArgs) (*LogCollection, error)
	// [Preview API] Gets a list of pipelines.
	ListPipelines(context.Context, ListPipelinesArgs) (*ListPipelinesResponseValue, error)
	// [Preview API] Gets top 10000 runs for a particular pipeline.
	ListRuns(context.Context, ListRunsArgs) (*[]Run, error)
	// [Preview API] Runs a pipeline.
	RunPipeline(context.Context, RunPipelineArgs) (*Run, error)
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
func (client *ClientImpl) CreatePipeline(ctx context.Context, args CreatePipelineArgs) (*Pipeline, error) {
	if args.InputParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.InputParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.InputParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("28e1305e-2afe-47bf-abaf-cbb0e6a91988")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Pipeline
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePipeline function
type CreatePipelineArgs struct {
	// (required)
	InputParameters *CreatePipelineParameters
	// (required) Project ID or project name
	Project *string
}

// [Preview API]
func (client *ClientImpl) GetLog(ctx context.Context, args GetLogArgs) (*Log, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PipelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PipelineId"}
	}
	routeValues["pipelineId"] = strconv.Itoa(*args.PipelineId)
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)
	if args.LogId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.LogId"}
	}
	routeValues["logId"] = strconv.Itoa(*args.LogId)

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("fb1b6d27-3957-43d5-a14b-a2d70403e545")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Log
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLog function
type GetLogArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	PipelineId *int
	// (required)
	RunId *int
	// (required)
	LogId *int
	// (optional)
	Expand *GetLogExpandOptions
}

// [Preview API] Gets a pipeline, optionally at the specified version
func (client *ClientImpl) GetPipeline(ctx context.Context, args GetPipelineArgs) (*Pipeline, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PipelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PipelineId"}
	}
	routeValues["pipelineId"] = strconv.Itoa(*args.PipelineId)

	queryParams := url.Values{}
	if args.PipelineVersion != nil {
		queryParams.Add("pipelineVersion", strconv.Itoa(*args.PipelineVersion))
	}
	locationId, _ := uuid.Parse("28e1305e-2afe-47bf-abaf-cbb0e6a91988")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Pipeline
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPipeline function
type GetPipelineArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The pipeline id
	PipelineId *int
	// (optional) The pipeline version
	PipelineVersion *int
}

// [Preview API] Gets a run for a particular pipeline.
func (client *ClientImpl) GetRun(ctx context.Context, args GetRunArgs) (*Run, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PipelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PipelineId"}
	}
	routeValues["pipelineId"] = strconv.Itoa(*args.PipelineId)
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	locationId, _ := uuid.Parse("7859261e-d2e9-4a68-b820-a5d84cc5bb3d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Run
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRun function
type GetRunArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The pipeline id
	PipelineId *int
	// (required) The run id
	RunId *int
}

// [Preview API]
func (client *ClientImpl) ListLogs(ctx context.Context, args ListLogsArgs) (*LogCollection, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PipelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PipelineId"}
	}
	routeValues["pipelineId"] = strconv.Itoa(*args.PipelineId)
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("fb1b6d27-3957-43d5-a14b-a2d70403e545")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue LogCollection
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListLogs function
type ListLogsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	PipelineId *int
	// (required)
	RunId *int
	// (optional)
	Expand *GetLogExpandOptions
}

// [Preview API] Gets a list of pipelines.
func (client *ClientImpl) ListPipelines(ctx context.Context, args ListPipelinesArgs) (*ListPipelinesResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.OrderBy != nil {
		queryParams.Add("orderBy", *args.OrderBy)
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	locationId, _ := uuid.Parse("28e1305e-2afe-47bf-abaf-cbb0e6a91988")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ListPipelinesResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the ListPipelines function
type ListPipelinesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) A sort expression. Defaults to "name asc"
	OrderBy *string
	// (optional) The maximum number of pipelines to return
	Top *int
	// (optional) A continuation token from a previous request, to retrieve the next page of results
	ContinuationToken *string
}

// Return type for the ListPipelines function
type ListPipelinesResponseValue struct {
	Value []Pipeline
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// [Preview API] Gets top 10000 runs for a particular pipeline.
func (client *ClientImpl) ListRuns(ctx context.Context, args ListRunsArgs) (*[]Run, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PipelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PipelineId"}
	}
	routeValues["pipelineId"] = strconv.Itoa(*args.PipelineId)

	locationId, _ := uuid.Parse("7859261e-d2e9-4a68-b820-a5d84cc5bb3d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Run
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ListRuns function
type ListRunsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The pipeline id
	PipelineId *int
}

// [Preview API] Runs a pipeline.
func (client *ClientImpl) RunPipeline(ctx context.Context, args RunPipelineArgs) (*Run, error) {
	if args.RunParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PipelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PipelineId"}
	}
	routeValues["pipelineId"] = strconv.Itoa(*args.PipelineId)

	queryParams := url.Values{}
	if args.PipelineVersion != nil {
		queryParams.Add("pipelineVersion", strconv.Itoa(*args.PipelineVersion))
	}
	body, marshalErr := json.Marshal(*args.RunParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("7859261e-d2e9-4a68-b820-a5d84cc5bb3d")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Run
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the RunPipeline function
type RunPipelineArgs struct {
	// (required) Optional.
	RunParameters *RunPipelineParameters
	// (required) Project ID or project name
	Project *string
	// (required) The pipeline id
	PipelineId *int
	// (optional) The pipeline version
	PipelineVersion *int
}
