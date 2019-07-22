// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package taskAgent

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

var ResourceAreaId, _ = uuid.Parse("a85b8835-c1a1-4aac-ae97-1c3d0ba72dbd")

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
// agentCloud (required)
func (client Client) AddAgentCloud(ctx context.Context, agentCloud *TaskAgentCloud) (*TaskAgentCloud, error) {
    if agentCloud == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agentCloud"}
    }
    body, marshalErr := json.Marshal(*agentCloud)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bfa72b3d-0fc6-43fb-932b-a7f6559f93b9")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentCloud
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// agentCloudId (required)
func (client Client) DeleteAgentCloud(ctx context.Context, agentCloudId *int) (*TaskAgentCloud, error) {
    routeValues := make(map[string]string)
    if agentCloudId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agentCloudId"} 
    }
    routeValues["agentCloudId"] = strconv.Itoa(*agentCloudId)

    locationId, _ := uuid.Parse("bfa72b3d-0fc6-43fb-932b-a7f6559f93b9")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentCloud
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// agentCloudId (required)
func (client Client) GetAgentCloud(ctx context.Context, agentCloudId *int) (*TaskAgentCloud, error) {
    routeValues := make(map[string]string)
    if agentCloudId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agentCloudId"} 
    }
    routeValues["agentCloudId"] = strconv.Itoa(*agentCloudId)

    locationId, _ := uuid.Parse("bfa72b3d-0fc6-43fb-932b-a7f6559f93b9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentCloud
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
func (client Client) GetAgentClouds(ctx context.Context, ) (*[]TaskAgentCloud, error) {
    locationId, _ := uuid.Parse("bfa72b3d-0fc6-43fb-932b-a7f6559f93b9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentCloud
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get agent cloud types.
// ctx
func (client Client) GetAgentCloudTypes(ctx context.Context, ) (*[]TaskAgentCloudType, error) {
    locationId, _ := uuid.Parse("5932e193-f376-469d-9c3e-e5588ce12cb5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentCloudType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Adds an agent to a pool.  You probably don't want to call this endpoint directly. Instead, [configure an agent](https://docs.microsoft.com/azure/devops/pipelines/agents/agents) using the agent download package.
// ctx
// agent (required): Details about the agent being added
// poolId (required): The agent pool in which to add the agent
func (client Client) AddAgent(ctx context.Context, agent *TaskAgent, poolId *int) (*TaskAgent, error) {
    if agent == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agent"}
    }
    routeValues := make(map[string]string)
    if poolId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)

    body, marshalErr := json.Marshal(*agent)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e298ef32-5878-4cab-993c-043836571f42")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgent
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete an agent.  You probably don't want to call this endpoint directly. Instead, [use the agent configuration script](https://docs.microsoft.com/azure/devops/pipelines/agents/agents) to remove an agent from your organization.
// ctx
// poolId (required): The pool ID to remove the agent from
// agentId (required): The agent ID to remove
func (client Client) DeleteAgent(ctx context.Context, poolId *int, agentId *int) error {
    routeValues := make(map[string]string)
    if poolId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)
    if agentId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "agentId"} 
    }
    routeValues["agentId"] = strconv.Itoa(*agentId)

    locationId, _ := uuid.Parse("e298ef32-5878-4cab-993c-043836571f42")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get information about an agent.
// ctx
// poolId (required): The agent pool containing the agent
// agentId (required): The agent ID to get information about
// includeCapabilities (optional): Whether to include the agent's capabilities in the response
// includeAssignedRequest (optional): Whether to include details about the agent's current work
// includeLastCompletedRequest (optional): Whether to include details about the agents' most recent completed work
// propertyFilters (optional): Filter which custom properties will be returned
func (client Client) GetAgent(ctx context.Context, poolId *int, agentId *int, includeCapabilities *bool, includeAssignedRequest *bool, includeLastCompletedRequest *bool, propertyFilters *[]string) (*TaskAgent, error) {
    routeValues := make(map[string]string)
    if poolId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)
    if agentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agentId"} 
    }
    routeValues["agentId"] = strconv.Itoa(*agentId)

    queryParams := url.Values{}
    if includeCapabilities != nil {
        queryParams.Add("includeCapabilities", strconv.FormatBool(*includeCapabilities))
    }
    if includeAssignedRequest != nil {
        queryParams.Add("includeAssignedRequest", strconv.FormatBool(*includeAssignedRequest))
    }
    if includeLastCompletedRequest != nil {
        queryParams.Add("includeLastCompletedRequest", strconv.FormatBool(*includeLastCompletedRequest))
    }
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    locationId, _ := uuid.Parse("e298ef32-5878-4cab-993c-043836571f42")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgent
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of agents.
// ctx
// poolId (required): The agent pool containing the agents
// agentName (optional): Filter on agent name
// includeCapabilities (optional): Whether to include the agents' capabilities in the response
// includeAssignedRequest (optional): Whether to include details about the agents' current work
// includeLastCompletedRequest (optional): Whether to include details about the agents' most recent completed work
// propertyFilters (optional): Filter which custom properties will be returned
// demands (optional): Filter by demands the agents can satisfy
func (client Client) GetAgents(ctx context.Context, poolId *int, agentName *string, includeCapabilities *bool, includeAssignedRequest *bool, includeLastCompletedRequest *bool, propertyFilters *[]string, demands *[]string) (*[]TaskAgent, error) {
    routeValues := make(map[string]string)
    if poolId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)

    queryParams := url.Values{}
    if agentName != nil {
        queryParams.Add("agentName", *agentName)
    }
    if includeCapabilities != nil {
        queryParams.Add("includeCapabilities", strconv.FormatBool(*includeCapabilities))
    }
    if includeAssignedRequest != nil {
        queryParams.Add("includeAssignedRequest", strconv.FormatBool(*includeAssignedRequest))
    }
    if includeLastCompletedRequest != nil {
        queryParams.Add("includeLastCompletedRequest", strconv.FormatBool(*includeLastCompletedRequest))
    }
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if demands != nil {
        listAsString := strings.Join((*demands)[:], ",")
        queryParams.Add("demands", listAsString)
    }
    locationId, _ := uuid.Parse("e298ef32-5878-4cab-993c-043836571f42")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgent
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Replace an agent.  You probably don't want to call this endpoint directly. Instead, [use the agent configuration script](https://docs.microsoft.com/azure/devops/pipelines/agents/agents) to remove and reconfigure an agent from your organization.
// ctx
// agent (required): Updated details about the replacing agent
// poolId (required): The agent pool to use
// agentId (required): The agent to replace
func (client Client) ReplaceAgent(ctx context.Context, agent *TaskAgent, poolId *int, agentId *int) (*TaskAgent, error) {
    if agent == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agent"}
    }
    routeValues := make(map[string]string)
    if poolId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)
    if agentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agentId"} 
    }
    routeValues["agentId"] = strconv.Itoa(*agentId)

    body, marshalErr := json.Marshal(*agent)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e298ef32-5878-4cab-993c-043836571f42")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgent
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update agent details.
// ctx
// agent (required): Updated details about the agent
// poolId (required): The agent pool to use
// agentId (required): The agent to update
func (client Client) UpdateAgent(ctx context.Context, agent *TaskAgent, poolId *int, agentId *int) (*TaskAgent, error) {
    if agent == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agent"}
    }
    routeValues := make(map[string]string)
    if poolId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)
    if agentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agentId"} 
    }
    routeValues["agentId"] = strconv.Itoa(*agentId)

    body, marshalErr := json.Marshal(*agent)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e298ef32-5878-4cab-993c-043836571f42")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgent
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a deployment group.
// ctx
// deploymentGroup (required): Deployment group to create.
// project (required): Project ID or project name
func (client Client) AddDeploymentGroup(ctx context.Context, deploymentGroup *DeploymentGroupCreateParameter, project *string) (*DeploymentGroup, error) {
    if deploymentGroup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroup"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*deploymentGroup)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("083c4d89-ab35-45af-aa11-7cf66895c53e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DeploymentGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a deployment group.
// ctx
// project (required): Project ID or project name
// deploymentGroupId (required): ID of the deployment group to be deleted.
func (client Client) DeleteDeploymentGroup(ctx context.Context, project *string, deploymentGroupId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if deploymentGroupId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroupId"} 
    }
    routeValues["deploymentGroupId"] = strconv.Itoa(*deploymentGroupId)

    locationId, _ := uuid.Parse("083c4d89-ab35-45af-aa11-7cf66895c53e")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a deployment group by its ID.
// ctx
// project (required): Project ID or project name
// deploymentGroupId (required): ID of the deployment group.
// actionFilter (optional): Get the deployment group only if this action can be performed on it.
// expand (optional): Include these additional details in the returned object.
func (client Client) GetDeploymentGroup(ctx context.Context, project *string, deploymentGroupId *int, actionFilter *string, expand *string) (*DeploymentGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if deploymentGroupId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroupId"} 
    }
    routeValues["deploymentGroupId"] = strconv.Itoa(*deploymentGroupId)

    queryParams := url.Values{}
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("083c4d89-ab35-45af-aa11-7cf66895c53e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DeploymentGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of deployment groups by name or IDs.
// ctx
// project (required): Project ID or project name
// name (optional): Name of the deployment group.
// actionFilter (optional): Get only deployment groups on which this action can be performed.
// expand (optional): Include these additional details in the returned objects.
// continuationToken (optional): Get deployment groups with names greater than this continuationToken lexicographically.
// top (optional): Maximum number of deployment groups to return. Default is **1000**.
// ids (optional): Comma separated list of IDs of the deployment groups.
func (client Client) GetDeploymentGroups(ctx context.Context, project *string, name *string, actionFilter *string, expand *string, continuationToken *string, top *int, ids *[]int) (*[]DeploymentGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if name != nil {
        queryParams.Add("name", *name)
    }
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if ids != nil {
        var stringList []string
        for _, item := range *ids {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    locationId, _ := uuid.Parse("083c4d89-ab35-45af-aa11-7cf66895c53e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DeploymentGroup
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a deployment group.
// ctx
// deploymentGroup (required): Deployment group to update.
// project (required): Project ID or project name
// deploymentGroupId (required): ID of the deployment group.
func (client Client) UpdateDeploymentGroup(ctx context.Context, deploymentGroup *DeploymentGroupUpdateParameter, project *string, deploymentGroupId *int) (*DeploymentGroup, error) {
    if deploymentGroup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroup"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if deploymentGroupId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroupId"} 
    }
    routeValues["deploymentGroupId"] = strconv.Itoa(*deploymentGroupId)

    body, marshalErr := json.Marshal(*deploymentGroup)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("083c4d89-ab35-45af-aa11-7cf66895c53e")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DeploymentGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Create an agent pool.
// ctx
// pool (required): Details about the new agent pool
func (client Client) AddAgentPool(ctx context.Context, pool *TaskAgentPool) (*TaskAgentPool, error) {
    if pool == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pool"}
    }
    body, marshalErr := json.Marshal(*pool)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a8c47e17-4d56-4a56-92bb-de7ea7dc65be")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentPool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete an agent pool.
// ctx
// poolId (required): ID of the agent pool to delete
func (client Client) DeleteAgentPool(ctx context.Context, poolId *int) error {
    routeValues := make(map[string]string)
    if poolId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)

    locationId, _ := uuid.Parse("a8c47e17-4d56-4a56-92bb-de7ea7dc65be")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get information about an agent pool.
// ctx
// poolId (required): An agent pool ID
// properties (optional): Agent pool properties (comma-separated)
// actionFilter (optional): Filter by whether the calling user has use or manage permissions
func (client Client) GetAgentPool(ctx context.Context, poolId *int, properties *[]string, actionFilter *string) (*TaskAgentPool, error) {
    routeValues := make(map[string]string)
    if poolId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)

    queryParams := url.Values{}
    if properties != nil {
        listAsString := strings.Join((*properties)[:], ",")
        queryParams.Add("properties", listAsString)
    }
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    locationId, _ := uuid.Parse("a8c47e17-4d56-4a56-92bb-de7ea7dc65be")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentPool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of agent pools.
// ctx
// poolName (optional): Filter by name
// properties (optional): Filter by agent pool properties (comma-separated)
// poolType (optional): Filter by pool type
// actionFilter (optional): Filter by whether the calling user has use or manage permissions
func (client Client) GetAgentPools(ctx context.Context, poolName *string, properties *[]string, poolType *string, actionFilter *string) (*[]TaskAgentPool, error) {
    queryParams := url.Values{}
    if poolName != nil {
        queryParams.Add("poolName", *poolName)
    }
    if properties != nil {
        listAsString := strings.Join((*properties)[:], ",")
        queryParams.Add("properties", listAsString)
    }
    if poolType != nil {
        queryParams.Add("poolType", *poolType)
    }
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    locationId, _ := uuid.Parse("a8c47e17-4d56-4a56-92bb-de7ea7dc65be")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentPool
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of agent pools.
// ctx
// poolIds (required): pool Ids to fetch
// actionFilter (optional): Filter by whether the calling user has use or manage permissions
func (client Client) GetAgentPoolsByIds(ctx context.Context, poolIds *[]int, actionFilter *string) (*[]TaskAgentPool, error) {
    queryParams := url.Values{}
    if poolIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolIds"}
    }
    var stringList []string
    for _, item := range *poolIds {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    locationId, _ := uuid.Parse("a8c47e17-4d56-4a56-92bb-de7ea7dc65be")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentPool
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update properties on an agent pool
// ctx
// pool (required): Updated agent pool details
// poolId (required): The agent pool to update
func (client Client) UpdateAgentPool(ctx context.Context, pool *TaskAgentPool, poolId *int) (*TaskAgentPool, error) {
    if pool == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pool"}
    }
    routeValues := make(map[string]string)
    if poolId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "poolId"} 
    }
    routeValues["poolId"] = strconv.Itoa(*poolId)

    body, marshalErr := json.Marshal(*pool)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a8c47e17-4d56-4a56-92bb-de7ea7dc65be")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentPool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a new agent queue to connect a project to an agent pool.
// ctx
// queue (required): Details about the queue to create
// project (optional): Project ID or project name
// authorizePipelines (optional): Automatically authorize this queue when using YAML
func (client Client) AddAgentQueue(ctx context.Context, queue *TaskAgentQueue, project *string, authorizePipelines *bool) (*TaskAgentQueue, error) {
    if queue == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queue"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if authorizePipelines != nil {
        queryParams.Add("authorizePipelines", strconv.FormatBool(*authorizePipelines))
    }
    body, marshalErr := json.Marshal(*queue)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("900fa995-c559-4923-aae7-f8424fe4fbea")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentQueue
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes an agent queue from a project.
// ctx
// queueId (required): The agent queue to remove
// project (optional): Project ID or project name
func (client Client) DeleteAgentQueue(ctx context.Context, queueId *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if queueId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "queueId"} 
    }
    routeValues["queueId"] = strconv.Itoa(*queueId)

    locationId, _ := uuid.Parse("900fa995-c559-4923-aae7-f8424fe4fbea")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about an agent queue.
// ctx
// queueId (required): The agent queue to get information about
// project (optional): Project ID or project name
// actionFilter (optional): Filter by whether the calling user has use or manage permissions
func (client Client) GetAgentQueue(ctx context.Context, queueId *int, project *string, actionFilter *string) (*TaskAgentQueue, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if queueId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queueId"} 
    }
    routeValues["queueId"] = strconv.Itoa(*queueId)

    queryParams := url.Values{}
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    locationId, _ := uuid.Parse("900fa995-c559-4923-aae7-f8424fe4fbea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAgentQueue
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of agent queues.
// ctx
// project (optional): Project ID or project name
// queueName (optional): Filter on the agent queue name
// actionFilter (optional): Filter by whether the calling user has use or manage permissions
func (client Client) GetAgentQueues(ctx context.Context, project *string, queueName *string, actionFilter *string) (*[]TaskAgentQueue, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if queueName != nil {
        queryParams.Add("queueName", *queueName)
    }
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    locationId, _ := uuid.Parse("900fa995-c559-4923-aae7-f8424fe4fbea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentQueue
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of agent queues by their IDs
// ctx
// queueIds (required): A comma-separated list of agent queue IDs to retrieve
// project (optional): Project ID or project name
// actionFilter (optional): Filter by whether the calling user has use or manage permissions
func (client Client) GetAgentQueuesByIds(ctx context.Context, queueIds *[]int, project *string, actionFilter *string) (*[]TaskAgentQueue, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if queueIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queueIds"}
    }
    var stringList []string
    for _, item := range *queueIds {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    locationId, _ := uuid.Parse("900fa995-c559-4923-aae7-f8424fe4fbea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentQueue
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of agent queues by their names
// ctx
// queueNames (required): A comma-separated list of agent names to retrieve
// project (optional): Project ID or project name
// actionFilter (optional): Filter by whether the calling user has use or manage permissions
func (client Client) GetAgentQueuesByNames(ctx context.Context, queueNames *[]string, project *string, actionFilter *string) (*[]TaskAgentQueue, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if queueNames == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queueNames"}
    }
    listAsString := strings.Join((*queueNames)[:], ",")
    queryParams.Add("queueNames", listAsString)
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    locationId, _ := uuid.Parse("900fa995-c559-4923-aae7-f8424fe4fbea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentQueue
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// agentCloudId (required)
func (client Client) GetAgentCloudRequests(ctx context.Context, agentCloudId *int) (*[]TaskAgentCloudRequest, error) {
    routeValues := make(map[string]string)
    if agentCloudId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "agentCloudId"} 
    }
    routeValues["agentCloudId"] = strconv.Itoa(*agentCloudId)

    locationId, _ := uuid.Parse("20189bd7-5134-49c2-b8e9-f9e856eea2b2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAgentCloudRequest
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a deployment target in a deployment group. This deletes the agent from associated deployment pool too.
// ctx
// project (required): Project ID or project name
// deploymentGroupId (required): ID of the deployment group in which deployment target is deleted.
// targetId (required): ID of the deployment target to delete.
func (client Client) DeleteDeploymentTarget(ctx context.Context, project *string, deploymentGroupId *int, targetId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if deploymentGroupId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroupId"} 
    }
    routeValues["deploymentGroupId"] = strconv.Itoa(*deploymentGroupId)
    if targetId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "targetId"} 
    }
    routeValues["targetId"] = strconv.Itoa(*targetId)

    locationId, _ := uuid.Parse("2f0aa599-c121-4256-a5fd-ba370e0ae7b6")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a deployment target by its ID in a deployment group
// ctx
// project (required): Project ID or project name
// deploymentGroupId (required): ID of the deployment group to which deployment target belongs.
// targetId (required): ID of the deployment target to return.
// expand (optional): Include these additional details in the returned objects.
func (client Client) GetDeploymentTarget(ctx context.Context, project *string, deploymentGroupId *int, targetId *int, expand *string) (*DeploymentMachine, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if deploymentGroupId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroupId"} 
    }
    routeValues["deploymentGroupId"] = strconv.Itoa(*deploymentGroupId)
    if targetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "targetId"} 
    }
    routeValues["targetId"] = strconv.Itoa(*targetId)

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("2f0aa599-c121-4256-a5fd-ba370e0ae7b6")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DeploymentMachine
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of deployment targets in a deployment group.
// ctx
// project (required): Project ID or project name
// deploymentGroupId (required): ID of the deployment group.
// tags (optional): Get only the deployment targets that contain all these comma separted list of tags.
// name (optional): Name pattern of the deployment targets to return.
// partialNameMatch (optional): When set to true, treats **name** as pattern. Else treats it as absolute match. Default is **false**.
// expand (optional): Include these additional details in the returned objects.
// agentStatus (optional): Get only deployment targets that have this status.
// agentJobResult (optional): Get only deployment targets that have this last job result.
// continuationToken (optional): Get deployment targets with names greater than this continuationToken lexicographically.
// top (optional): Maximum number of deployment targets to return. Default is **1000**.
// enabled (optional): Get only deployment targets that are enabled or disabled. Default is 'null' which returns all the targets.
// propertyFilters (optional)
func (client Client) GetDeploymentTargets(ctx context.Context, project *string, deploymentGroupId *int, tags *[]string, name *string, partialNameMatch *bool, expand *string, agentStatus *string, agentJobResult *string, continuationToken *string, top *int, enabled *bool, propertyFilters *[]string) (*[]DeploymentMachine, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if deploymentGroupId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroupId"} 
    }
    routeValues["deploymentGroupId"] = strconv.Itoa(*deploymentGroupId)

    queryParams := url.Values{}
    if tags != nil {
        listAsString := strings.Join((*tags)[:], ",")
        queryParams.Add("tags", listAsString)
    }
    if name != nil {
        queryParams.Add("name", *name)
    }
    if partialNameMatch != nil {
        queryParams.Add("partialNameMatch", strconv.FormatBool(*partialNameMatch))
    }
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    if agentStatus != nil {
        queryParams.Add("agentStatus", *agentStatus)
    }
    if agentJobResult != nil {
        queryParams.Add("agentJobResult", *agentJobResult)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if enabled != nil {
        queryParams.Add("enabled", strconv.FormatBool(*enabled))
    }
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    locationId, _ := uuid.Parse("2f0aa599-c121-4256-a5fd-ba370e0ae7b6")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DeploymentMachine
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update tags of a list of deployment targets in a deployment group.
// ctx
// machines (required): Deployment targets with tags to udpdate.
// project (required): Project ID or project name
// deploymentGroupId (required): ID of the deployment group in which deployment targets are updated.
func (client Client) UpdateDeploymentTargets(ctx context.Context, machines *[]DeploymentTargetUpdateParameter, project *string, deploymentGroupId *int) (*[]DeploymentMachine, error) {
    if machines == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "machines"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if deploymentGroupId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "deploymentGroupId"} 
    }
    routeValues["deploymentGroupId"] = strconv.Itoa(*deploymentGroupId)

    body, marshalErr := json.Marshal(*machines)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2f0aa599-c121-4256-a5fd-ba370e0ae7b6")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DeploymentMachine
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a task group.
// ctx
// taskGroup (required): Task group object to create.
// project (required): Project ID or project name
func (client Client) AddTaskGroup(ctx context.Context, taskGroup *TaskGroupCreateParameter, project *string) (*TaskGroup, error) {
    if taskGroup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "taskGroup"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*taskGroup)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6c08ffbf-dbf1-4f9a-94e5-a1cbd47005e7")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a task group.
// ctx
// project (required): Project ID or project name
// taskGroupId (required): Id of the task group to be deleted.
// comment (optional): Comments to delete.
func (client Client) DeleteTaskGroup(ctx context.Context, project *string, taskGroupId *uuid.UUID, comment *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if taskGroupId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "taskGroupId"} 
    }
    routeValues["taskGroupId"] = (*taskGroupId).String()

    queryParams := url.Values{}
    if comment != nil {
        queryParams.Add("comment", *comment)
    }
    locationId, _ := uuid.Parse("6c08ffbf-dbf1-4f9a-94e5-a1cbd47005e7")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] List task groups.
// ctx
// project (required): Project ID or project name
// taskGroupId (optional): Id of the task group.
// expanded (optional): 'true' to recursively expand task groups. Default is 'false'.
// taskIdFilter (optional): Guid of the taskId to filter.
// deleted (optional): 'true'to include deleted task groups. Default is 'false'.
// top (optional): Number of task groups to get.
// continuationToken (optional): Gets the task groups after the continuation token provided.
// queryOrder (optional): Gets the results in the defined order. Default is 'CreatedOnDescending'.
func (client Client) GetTaskGroups(ctx context.Context, project *string, taskGroupId *uuid.UUID, expanded *bool, taskIdFilter *uuid.UUID, deleted *bool, top *int, continuationToken *time.Time, queryOrder *string) (*[]TaskGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if taskGroupId != nil {
        routeValues["taskGroupId"] = (*taskGroupId).String()
    }

    queryParams := url.Values{}
    if expanded != nil {
        queryParams.Add("expanded", strconv.FormatBool(*expanded))
    }
    if taskIdFilter != nil {
        queryParams.Add("taskIdFilter", (*taskIdFilter).String())
    }
    if deleted != nil {
        queryParams.Add("deleted", strconv.FormatBool(*deleted))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", (*continuationToken).String())
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", *queryOrder)
    }
    locationId, _ := uuid.Parse("6c08ffbf-dbf1-4f9a-94e5-a1cbd47005e7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskGroup
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a task group.
// ctx
// taskGroup (required): Task group to update.
// project (required): Project ID or project name
// taskGroupId (optional): Id of the task group to update.
func (client Client) UpdateTaskGroup(ctx context.Context, taskGroup *TaskGroupUpdateParameter, project *string, taskGroupId *uuid.UUID) (*TaskGroup, error) {
    if taskGroup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "taskGroup"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if taskGroupId != nil {
        routeValues["taskGroupId"] = (*taskGroupId).String()
    }

    body, marshalErr := json.Marshal(*taskGroup)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6c08ffbf-dbf1-4f9a-94e5-a1cbd47005e7")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Add a variable group.
// ctx
// group (required): Variable group to add.
// project (required): Project ID or project name
func (client Client) AddVariableGroup(ctx context.Context, group *VariableGroupParameters, project *string) (*VariableGroup, error) {
    if group == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "group"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue VariableGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a variable group
// ctx
// project (required): Project ID or project name
// groupId (required): Id of the variable group.
func (client Client) DeleteVariableGroup(ctx context.Context, project *string, groupId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if groupId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = strconv.Itoa(*groupId)

    locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a variable group.
// ctx
// project (required): Project ID or project name
// groupId (required): Id of the variable group.
func (client Client) GetVariableGroup(ctx context.Context, project *string, groupId *int) (*VariableGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if groupId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = strconv.Itoa(*groupId)

    locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue VariableGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get variable groups.
// ctx
// project (required): Project ID or project name
// groupName (optional): Name of variable group.
// actionFilter (optional): Action filter for the variable group. It specifies the action which can be performed on the variable groups.
// top (optional): Number of variable groups to get.
// continuationToken (optional): Gets the variable groups after the continuation token provided.
// queryOrder (optional): Gets the results in the defined order. Default is 'IdDescending'.
func (client Client) GetVariableGroups(ctx context.Context, project *string, groupName *string, actionFilter *string, top *int, continuationToken *int, queryOrder *string) (*[]VariableGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if groupName != nil {
        queryParams.Add("groupName", *groupName)
    }
    if actionFilter != nil {
        queryParams.Add("actionFilter", *actionFilter)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", strconv.Itoa(*continuationToken))
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", *queryOrder)
    }
    locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []VariableGroup
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get variable groups by ids.
// ctx
// project (required): Project ID or project name
// groupIds (required): Comma separated list of Ids of variable groups.
func (client Client) GetVariableGroupsById(ctx context.Context, project *string, groupIds *[]int) (*[]VariableGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if groupIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "groupIds"}
    }
    var stringList []string
    for _, item := range *groupIds {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []VariableGroup
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a variable group.
// ctx
// group (required): Variable group to update.
// project (required): Project ID or project name
// groupId (required): Id of the variable group to update.
func (client Client) UpdateVariableGroup(ctx context.Context, group *VariableGroupParameters, project *string, groupId *int) (*VariableGroup, error) {
    if group == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "group"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if groupId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = strconv.Itoa(*groupId)

    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue VariableGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
func (client Client) GetYamlSchema(ctx context.Context, ) (*interface{}, error) {
    locationId, _ := uuid.Parse("1f9990b9-1dba-441f-9c2e-6485888c42b6")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

