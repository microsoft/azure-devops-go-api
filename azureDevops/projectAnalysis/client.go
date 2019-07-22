// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package projectAnalysis

import (
    "context"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
    "time"
)

var ResourceAreaId, _ = uuid.Parse("7658fa33-b1bf-4580-990f-fac5896773d3")

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

// [Preview API]
// ctx
// project (required): Project ID or project name
func (client Client) GetProjectLanguageAnalytics(ctx context.Context, project *string) (*ProjectLanguageAnalytics, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("5b02a779-1867-433f-90b7-d23ed5e33e57")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProjectLanguageAnalytics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// fromDate (required)
// aggregationType (required)
func (client Client) GetProjectActivityMetrics(ctx context.Context, project *string, fromDate *time.Time, aggregationType *AggregationType) (*ProjectActivityMetrics, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if fromDate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "fromDate"}
    }
    queryParams.Add("fromDate", (*fromDate).String())
    if aggregationType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "aggregationType"}
    }
    queryParams.Add("aggregationType", string(*aggregationType))
    locationId, _ := uuid.Parse("e40ae584-9ea6-4f06-a7c7-6284651b466b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProjectActivityMetrics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieves git activity metrics for repositories matching a specified criteria.
// ctx
// project (required): Project ID or project name
// fromDate (required): Date from which, the trends are to be fetched.
// aggregationType (required): Bucket size on which, trends are to be aggregated.
// skip (required): The number of repositories to ignore.
// top (required): The number of repositories for which activity metrics are to be retrieved.
func (client Client) GetGitRepositoriesActivityMetrics(ctx context.Context, project *string, fromDate *time.Time, aggregationType *AggregationType, skip *int, top *int) (*[]RepositoryActivityMetrics, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if fromDate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "fromDate"}
    }
    queryParams.Add("fromDate", (*fromDate).String())
    if aggregationType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "aggregationType"}
    }
    queryParams.Add("aggregationType", string(*aggregationType))
    if skip == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "skip"}
    }
    queryParams.Add("$skip", strconv.Itoa(*skip))
    if top == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "top"}
    }
    queryParams.Add("$top", strconv.Itoa(*top))
    locationId, _ := uuid.Parse("df7fbbca-630a-40e3-8aa3-7a3faf66947e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []RepositoryActivityMetrics
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// repositoryId (required)
// fromDate (required)
// aggregationType (required)
func (client Client) GetRepositoryActivityMetrics(ctx context.Context, project *string, repositoryId *uuid.UUID, fromDate *time.Time, aggregationType *AggregationType) (*RepositoryActivityMetrics, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = (*repositoryId).String()

    queryParams := url.Values{}
    if fromDate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "fromDate"}
    }
    queryParams.Add("fromDate", (*fromDate).String())
    if aggregationType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "aggregationType"}
    }
    queryParams.Add("aggregationType", string(*aggregationType))
    locationId, _ := uuid.Parse("df7fbbca-630a-40e3-8aa3-7a3faf66947e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue RepositoryActivityMetrics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

