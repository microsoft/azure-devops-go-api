// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package policy

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

var ResourceAreaId, _ = uuid.Parse("fb13a388-40dd-4a04-b530-013a739c72ef")

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

// Create a policy configuration of a given policy type.
// ctx
// configuration (required): The policy configuration to create.
// project (required): Project ID or project name
// configurationId (optional)
func (client Client) CreatePolicyConfiguration(ctx context.Context, configuration *PolicyConfiguration, project *string, configurationId *int) (*PolicyConfiguration, error) {
    if configuration == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "configuration"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if configurationId != nil {
        routeValues["configurationId"] = strconv.Itoa(*configurationId)
    }

    body, marshalErr := json.Marshal(*configuration)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PolicyConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a policy configuration by its ID.
// ctx
// project (required): Project ID or project name
// configurationId (required): ID of the policy configuration to delete.
func (client Client) DeletePolicyConfiguration(ctx context.Context, project *string, configurationId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if configurationId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "configurationId"} 
    }
    routeValues["configurationId"] = strconv.Itoa(*configurationId)

    locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a policy configuration by its ID.
// ctx
// project (required): Project ID or project name
// configurationId (required): ID of the policy configuration
func (client Client) GetPolicyConfiguration(ctx context.Context, project *string, configurationId *int) (*PolicyConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if configurationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "configurationId"} 
    }
    routeValues["configurationId"] = strconv.Itoa(*configurationId)

    locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PolicyConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of policy configurations in a project.
// ctx
// project (required): Project ID or project name
// scope (optional): [Provided for legacy reasons] The scope on which a subset of policies is defined.
// policyType (optional): Filter returned policies to only this type
func (client Client) GetPolicyConfigurations(ctx context.Context, project *string, scope *string, policyType *uuid.UUID) (*[]PolicyConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if scope != nil {
        queryParams.Add("scope", *scope)
    }
    if policyType != nil {
        queryParams.Add("policyType", (*policyType).String())
    }
    locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PolicyConfiguration
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a policy configuration by its ID.
// ctx
// configuration (required): The policy configuration to update.
// project (required): Project ID or project name
// configurationId (required): ID of the existing policy configuration to be updated.
func (client Client) UpdatePolicyConfiguration(ctx context.Context, configuration *PolicyConfiguration, project *string, configurationId *int) (*PolicyConfiguration, error) {
    if configuration == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "configuration"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if configurationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "configurationId"} 
    }
    routeValues["configurationId"] = strconv.Itoa(*configurationId)

    body, marshalErr := json.Marshal(*configuration)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("dad91cbe-d183-45f8-9c6e-9c1164472121")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PolicyConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the present evaluation state of a policy.
// ctx
// project (required): Project ID or project name
// evaluationId (required): ID of the policy evaluation to be retrieved.
func (client Client) GetPolicyEvaluation(ctx context.Context, project *string, evaluationId *uuid.UUID) (*PolicyEvaluationRecord, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if evaluationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "evaluationId"} 
    }
    routeValues["evaluationId"] = (*evaluationId).String()

    locationId, _ := uuid.Parse("46aecb7a-5d2c-4647-897b-0209505a9fe4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PolicyEvaluationRecord
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Requeue the policy evaluation.
// ctx
// project (required): Project ID or project name
// evaluationId (required): ID of the policy evaluation to be retrieved.
func (client Client) RequeuePolicyEvaluation(ctx context.Context, project *string, evaluationId *uuid.UUID) (*PolicyEvaluationRecord, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if evaluationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "evaluationId"} 
    }
    routeValues["evaluationId"] = (*evaluationId).String()

    locationId, _ := uuid.Parse("46aecb7a-5d2c-4647-897b-0209505a9fe4")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PolicyEvaluationRecord
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieves a list of all the policy evaluation statuses for a specific pull request.
// ctx
// project (required): Project ID or project name
// artifactId (required): A string which uniquely identifies the target of a policy evaluation.
// includeNotApplicable (optional): Some policies might determine that they do not apply to a specific pull request. Setting this parameter to true will return evaluation records even for policies which don't apply to this pull request.
// top (optional): The number of policy evaluation records to retrieve.
// skip (optional): The number of policy evaluation records to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
func (client Client) GetPolicyEvaluations(ctx context.Context, project *string, artifactId *string, includeNotApplicable *bool, top *int, skip *int) (*[]PolicyEvaluationRecord, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if artifactId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "artifactId"}
    }
    queryParams.Add("artifactId", *artifactId)
    if includeNotApplicable != nil {
        queryParams.Add("includeNotApplicable", strconv.FormatBool(*includeNotApplicable))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("c23ddff5-229c-4d04-a80b-0fdce9f360c8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PolicyEvaluationRecord
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a specific revision of a given policy by ID.
// ctx
// project (required): Project ID or project name
// configurationId (required): The policy configuration ID.
// revisionId (required): The revision ID.
func (client Client) GetPolicyConfigurationRevision(ctx context.Context, project *string, configurationId *int, revisionId *int) (*PolicyConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if configurationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "configurationId"} 
    }
    routeValues["configurationId"] = strconv.Itoa(*configurationId)
    if revisionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "revisionId"} 
    }
    routeValues["revisionId"] = strconv.Itoa(*revisionId)

    locationId, _ := uuid.Parse("fe1e68a2-60d3-43cb-855b-85e41ae97c95")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PolicyConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve all revisions for a given policy.
// ctx
// project (required): Project ID or project name
// configurationId (required): The policy configuration ID.
// top (optional): The number of revisions to retrieve.
// skip (optional): The number of revisions to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
func (client Client) GetPolicyConfigurationRevisions(ctx context.Context, project *string, configurationId *int, top *int, skip *int) (*[]PolicyConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if configurationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "configurationId"} 
    }
    routeValues["configurationId"] = strconv.Itoa(*configurationId)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("fe1e68a2-60d3-43cb-855b-85e41ae97c95")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PolicyConfiguration
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a specific policy type by ID.
// ctx
// project (required): Project ID or project name
// typeId (required): The policy ID.
func (client Client) GetPolicyType(ctx context.Context, project *string, typeId *uuid.UUID) (*PolicyType, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if typeId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "typeId"} 
    }
    routeValues["typeId"] = (*typeId).String()

    locationId, _ := uuid.Parse("44096322-2d3d-466a-bb30-d1b7de69f61f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PolicyType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve all available policy types.
// ctx
// project (required): Project ID or project name
func (client Client) GetPolicyTypes(ctx context.Context, project *string) (*[]PolicyType, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("44096322-2d3d-466a-bb30-d1b7de69f61f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PolicyType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

