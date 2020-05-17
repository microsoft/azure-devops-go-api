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
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
)

type Client interface {
	// [Preview API] Creates a new Pipeline connection between the provider installation and the specified project. Returns the PipelineConnection object created.
	CreateProjectConnection(context.Context, CreateProjectConnectionArgs) (*PipelineConnection, error)
	// [Preview API]
	CreateResources(context.Context, CreateResourcesArgs) (*CreatedResources, error)
	// [Preview API] Gets a list of existing configuration files for the given repository.
	GetConfigurations(context.Context, GetConfigurationsArgs) (*[]ConfigurationFile, error)
	// [Preview API] Returns a list of build frameworks that best match the given repository based on its contents.
	GetDetectedBuildFrameworks(context.Context, GetDetectedBuildFrameworksArgs) (*[]DetectedBuildFramework, error)
	// [Preview API] Returns a list of all YAML templates with weighting based on which would best fit the given repository.
	GetTemplateRecommendations(context.Context, GetTemplateRecommendationsArgs) (*[]Template, error)
	// [Preview API]
	RenderTemplate(context.Context, RenderTemplateArgs) (*Template, error)
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

// [Preview API] Creates a new Pipeline connection between the provider installation and the specified project. Returns the PipelineConnection object created.
func (client *ClientImpl) CreateProjectConnection(ctx context.Context, args CreateProjectConnectionArgs) (*PipelineConnection, error) {
	if args.CreateConnectionInputs == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CreateConnectionInputs"}
	}
	queryParams := url.Values{}
	if args.Project == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "project"}
	}
	queryParams.Add("project", *args.Project)
	body, marshalErr := json.Marshal(*args.CreateConnectionInputs)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("00df4879-9216-45d5-b38d-4a487b626b2c")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PipelineConnection
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateProjectConnection function
type CreateProjectConnectionArgs struct {
	// (required)
	CreateConnectionInputs *CreatePipelineConnectionInputs
	// (required)
	Project *string
}

// [Preview API]
func (client *ClientImpl) CreateResources(ctx context.Context, args CreateResourcesArgs) (*CreatedResources, error) {
	if args.CreationParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CreationParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.CreationParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("43201899-7690-4870-9c79-ab69605f21ed")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CreatedResources
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateResources function
type CreateResourcesArgs struct {
	// (required)
	CreationParameters *map[string]ResourceCreationParameter
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Gets a list of existing configuration files for the given repository.
func (client *ClientImpl) GetConfigurations(ctx context.Context, args GetConfigurationsArgs) (*[]ConfigurationFile, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.RepositoryType != nil {
		queryParams.Add("repositoryType", *args.RepositoryType)
	}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", *args.RepositoryId)
	}
	if args.Branch != nil {
		queryParams.Add("branch", *args.Branch)
	}
	if args.ServiceConnectionId != nil {
		queryParams.Add("serviceConnectionId", (*args.ServiceConnectionId).String())
	}
	locationId, _ := uuid.Parse("8fc87684-9ebc-4c37-ab92-f4ac4a58cb3a")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ConfigurationFile
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetConfigurations function
type GetConfigurationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The type of the repository such as GitHub, TfsGit (i.e. Azure Repos), Bitbucket, etc.
	RepositoryType *string
	// (optional) The vendor-specific identifier or the name of the repository, e.g. Microsoft/vscode (GitHub) or e9d82045-ddba-4e01-a63d-2ab9f040af62 (Azure Repos)
	RepositoryId *string
	// (optional) The repository branch where to look for the configuration file.
	Branch *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TfsGit (i.e. Azure Repos).
	ServiceConnectionId *uuid.UUID
}

// [Preview API] Returns a list of build frameworks that best match the given repository based on its contents.
func (client *ClientImpl) GetDetectedBuildFrameworks(ctx context.Context, args GetDetectedBuildFrameworksArgs) (*[]DetectedBuildFramework, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.RepositoryType != nil {
		queryParams.Add("repositoryType", *args.RepositoryType)
	}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", *args.RepositoryId)
	}
	if args.Branch != nil {
		queryParams.Add("branch", *args.Branch)
	}
	if args.DetectionType != nil {
		queryParams.Add("detectionType", string(*args.DetectionType))
	}
	if args.ServiceConnectionId != nil {
		queryParams.Add("serviceConnectionId", (*args.ServiceConnectionId).String())
	}
	locationId, _ := uuid.Parse("29a30bab-9efb-4652-bf1b-9269baca0980")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []DetectedBuildFramework
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetDetectedBuildFrameworks function
type GetDetectedBuildFrameworksArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The type of the repository such as GitHub, TfsGit (i.e. Azure Repos), Bitbucket, etc.
	RepositoryType *string
	// (optional) The vendor-specific identifier or the name of the repository, e.g. Microsoft/vscode (GitHub) or e9d82045-ddba-4e01-a63d-2ab9f040af62 (Azure Repos)
	RepositoryId *string
	// (optional) The repository branch to detect build frameworks for.
	Branch *string
	// (optional)
	DetectionType *BuildFrameworkDetectionType
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TfsGit (i.e. Azure Repos).
	ServiceConnectionId *uuid.UUID
}

// [Preview API] Returns a list of all YAML templates with weighting based on which would best fit the given repository.
func (client *ClientImpl) GetTemplateRecommendations(ctx context.Context, args GetTemplateRecommendationsArgs) (*[]Template, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.RepositoryType != nil {
		queryParams.Add("repositoryType", *args.RepositoryType)
	}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", *args.RepositoryId)
	}
	if args.Branch != nil {
		queryParams.Add("branch", *args.Branch)
	}
	if args.ServiceConnectionId != nil {
		queryParams.Add("serviceConnectionId", (*args.ServiceConnectionId).String())
	}
	locationId, _ := uuid.Parse("63ea8f13-b563-4be7-bc31-3a96eda27220")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Template
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTemplateRecommendations function
type GetTemplateRecommendationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The type of the repository such as GitHub, TfsGit (i.e. Azure Repos), Bitbucket, etc.
	RepositoryType *string
	// (optional) The vendor-specific identifier or the name of the repository, e.g. Microsoft/vscode (GitHub) or e9d82045-ddba-4e01-a63d-2ab9f040af62 (Azure Repos)
	RepositoryId *string
	// (optional) The repository branch which to find matching templates for.
	Branch *string
	// (optional) If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TfsGit (i.e. Azure Repos).
	ServiceConnectionId *uuid.UUID
}

// [Preview API]
func (client *ClientImpl) RenderTemplate(ctx context.Context, args RenderTemplateArgs) (*Template, error) {
	if args.TemplateParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TemplateParameters"}
	}
	routeValues := make(map[string]string)
	if args.TemplateId == nil || *args.TemplateId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TemplateId"}
	}
	routeValues["templateId"] = *args.TemplateId

	body, marshalErr := json.Marshal(*args.TemplateParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("eb5d6d1d-98a2-4bbd-9028-f9a6b2d66515")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Template
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the RenderTemplate function
type RenderTemplateArgs struct {
	// (required)
	TemplateParameters *TemplateParameters
	// (required)
	TemplateId *string
}
