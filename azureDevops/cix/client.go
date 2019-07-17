// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package cix

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API] Gets a list of existing configuration files for the given repository.
// project (required): Project ID or project name
// repositoryType (optional): The type of the repository such as GitHub, TfsGit (i.e. Azure Repos), Bitbucket, etc.
// repositoryId (optional): The vendor-specific identifier or the name of the repository, e.g. Microsoft/vscode (GitHub) or e9d82045-ddba-4e01-a63d-2ab9f040af62 (Azure Repos)
// branch (optional): The repository branch where to look for the configuration file.
// serviceConnectionId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TfsGit (i.e. Azure Repos).
func (client Client) GetConfigurations(project *string, repositoryType *string, repositoryId *string, branch *string, serviceConnectionId *uuid.UUID) (*[]ConfigurationFile, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if repositoryType != nil {
        queryParams.Add("repositoryType", *repositoryType)
    }
    if repositoryId != nil {
        queryParams.Add("repositoryId", *repositoryId)
    }
    if branch != nil {
        queryParams.Add("branch", *branch)
    }
    if serviceConnectionId != nil {
        queryParams.Add("serviceConnectionId", (*serviceConnectionId).String())
    }
    locationId, _ := uuid.Parse("8fc87684-9ebc-4c37-ab92-f4ac4a58cb3a")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ConfigurationFile
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a new Pipeline connection between the provider installation and the specified project. Returns the PipelineConnection object created.
// createConnectionInputs (required)
// project (required)
func (client Client) CreateProjectConnection(createConnectionInputs *CreatePipelineConnectionInputs, project *string) (*PipelineConnection, error) {
    if createConnectionInputs == nil {
        return nil, errors.New("createConnectionInputs is a required parameter")
    }
    queryParams := url.Values{}
    if project == nil {
        return nil, errors.New("project is a required parameter")
    }
    queryParams.Add("project", *project)
    body, marshalErr := json.Marshal(*createConnectionInputs)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("00df4879-9216-45d5-b38d-4a487b626b2c")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PipelineConnection
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of build frameworks that best match the given repository based on its contents.
// project (required): Project ID or project name
// repositoryType (optional): The type of the repository such as GitHub, TfsGit (i.e. Azure Repos), Bitbucket, etc.
// repositoryId (optional): The vendor-specific identifier or the name of the repository, e.g. Microsoft/vscode (GitHub) or e9d82045-ddba-4e01-a63d-2ab9f040af62 (Azure Repos)
// branch (optional): The repository branch to detect build frameworks for.
// detectionType (optional)
// serviceConnectionId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TfsGit (i.e. Azure Repos).
func (client Client) GetDetectedBuildFrameworks(project *string, repositoryType *string, repositoryId *string, branch *string, detectionType *string, serviceConnectionId *uuid.UUID) (*[]DetectedBuildFramework, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if repositoryType != nil {
        queryParams.Add("repositoryType", *repositoryType)
    }
    if repositoryId != nil {
        queryParams.Add("repositoryId", *repositoryId)
    }
    if branch != nil {
        queryParams.Add("branch", *branch)
    }
    if detectionType != nil {
        queryParams.Add("detectionType", *detectionType)
    }
    if serviceConnectionId != nil {
        queryParams.Add("serviceConnectionId", (*serviceConnectionId).String())
    }
    locationId, _ := uuid.Parse("29a30bab-9efb-4652-bf1b-9269baca0980")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DetectedBuildFramework
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all YAML templates with weighting based on which would best fit the given repository.
// project (required): Project ID or project name
// repositoryType (optional): The type of the repository such as GitHub, TfsGit (i.e. Azure Repos), Bitbucket, etc.
// repositoryId (optional): The vendor-specific identifier or the name of the repository, e.g. Microsoft/vscode (GitHub) or e9d82045-ddba-4e01-a63d-2ab9f040af62 (Azure Repos)
// branch (optional): The repository branch which to find matching templates for.
// serviceConnectionId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TfsGit (i.e. Azure Repos).
func (client Client) GetTemplateRecommendations(project *string, repositoryType *string, repositoryId *string, branch *string, serviceConnectionId *uuid.UUID) (*[]Template, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if repositoryType != nil {
        queryParams.Add("repositoryType", *repositoryType)
    }
    if repositoryId != nil {
        queryParams.Add("repositoryId", *repositoryId)
    }
    if branch != nil {
        queryParams.Add("branch", *branch)
    }
    if serviceConnectionId != nil {
        queryParams.Add("serviceConnectionId", (*serviceConnectionId).String())
    }
    locationId, _ := uuid.Parse("63ea8f13-b563-4be7-bc31-3a96eda27220")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Template
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// creationParameters (required)
// project (required): Project ID or project name
func (client Client) CreateResources(creationParameters *map[string]ResourceCreationParameter, project *string) (*CreatedResources, error) {
    if creationParameters == nil {
        return nil, errors.New("creationParameters is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*creationParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("43201899-7690-4870-9c79-ab69605f21ed")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CreatedResources
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// templateParameters (required)
// templateId (required)
func (client Client) RenderTemplate(templateParameters *TemplateParameters, templateId *string) (*Template, error) {
    if templateParameters == nil {
        return nil, errors.New("templateParameters is a required parameter")
    }
    routeValues := make(map[string]string)
    if templateId == nil || *templateId == "" {
        return nil, errors.New("templateId is a required parameter")
    }
    routeValues["templateId"] = *templateId

    body, marshalErr := json.Marshal(*templateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("eb5d6d1d-98a2-4bbd-9028-f9a6b2d66515")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Template
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

