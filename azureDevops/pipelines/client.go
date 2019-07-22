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

// [Preview API]
// ctx
// project (required): Project ID or project name
// pipelineId (required)
// runId (required)
// logId (required)
// expand (optional)
func (client Client) GetLog(ctx context.Context, project *string, pipelineId *int, runId *int, logId *int, expand *string) (*Log, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if pipelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pipelineId"} 
    }
    routeValues["pipelineId"] = strconv.Itoa(*pipelineId)
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if logId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "logId"} 
    }
    routeValues["logId"] = strconv.Itoa(*logId)

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
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

// [Preview API]
// ctx
// project (required): Project ID or project name
// pipelineId (required)
// runId (required)
// expand (optional)
func (client Client) ListLogs(ctx context.Context, project *string, pipelineId *int, runId *int, expand *string) (*LogCollection, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if pipelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pipelineId"} 
    }
    routeValues["pipelineId"] = strconv.Itoa(*pipelineId)
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
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

// [Preview API]
// ctx
// inputParameters (required)
// project (required): Project ID or project name
func (client Client) CreatePipeline(ctx context.Context, inputParameters *CreatePipelineParameters, project *string) (*Pipeline, error) {
    if inputParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "inputParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*inputParameters)
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

// [Preview API] Gets a pipeline, optionally at the specified version
// ctx
// project (required): Project ID or project name
// pipelineId (required): The pipeline id
// pipelineVersion (optional): The pipeline version
func (client Client) GetPipeline(ctx context.Context, project *string, pipelineId *int, pipelineVersion *int) (*Pipeline, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if pipelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pipelineId"} 
    }
    routeValues["pipelineId"] = strconv.Itoa(*pipelineId)

    queryParams := url.Values{}
    if pipelineVersion != nil {
        queryParams.Add("pipelineVersion", strconv.Itoa(*pipelineVersion))
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

// [Preview API] Gets a list of pipelines.
// ctx
// project (required): Project ID or project name
// orderBy (optional): A sort expression. Defaults to "name asc"
// top (optional): The maximum number of pipelines to return
// continuationToken (optional): A continuation token from a previous request, to retrieve the next page of results
func (client Client) ListPipelines(ctx context.Context, project *string, orderBy *string, top *int, continuationToken *string) (*[]Pipeline, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if orderBy != nil {
        queryParams.Add("orderBy", *orderBy)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("28e1305e-2afe-47bf-abaf-cbb0e6a91988")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Pipeline
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a run for a particular pipeline.
// ctx
// project (required): Project ID or project name
// pipelineId (required): The pipeline id
// runId (required): The run id
func (client Client) GetRun(ctx context.Context, project *string, pipelineId *int, runId *int) (*Run, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if pipelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pipelineId"} 
    }
    routeValues["pipelineId"] = strconv.Itoa(*pipelineId)
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    locationId, _ := uuid.Parse("7859261e-d2e9-4a68-b820-a5d84cc5bb3d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Run
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets top 10000 runs for a particular pipeline.
// ctx
// project (required): Project ID or project name
// pipelineId (required): The pipeline id
func (client Client) ListRuns(ctx context.Context, project *string, pipelineId *int) (*[]Run, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if pipelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pipelineId"} 
    }
    routeValues["pipelineId"] = strconv.Itoa(*pipelineId)

    locationId, _ := uuid.Parse("7859261e-d2e9-4a68-b820-a5d84cc5bb3d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Run
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Runs a pipeline.
// ctx
// runParameters (required): Optional.
// project (required): Project ID or project name
// pipelineId (required): The pipeline id
// pipelineVersion (optional): The pipeline version
func (client Client) RunPipeline(ctx context.Context, runParameters *RunPipelineParameters, project *string, pipelineId *int, pipelineVersion *int) (*Run, error) {
    if runParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if pipelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pipelineId"} 
    }
    routeValues["pipelineId"] = strconv.Itoa(*pipelineId)

    queryParams := url.Values{}
    if pipelineVersion != nil {
        queryParams.Add("pipelineVersion", strconv.Itoa(*pipelineVersion))
    }
    body, marshalErr := json.Marshal(*runParameters)
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

