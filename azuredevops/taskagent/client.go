// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package taskagent

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var ResourceAreaId, _ = uuid.Parse("a85b8835-c1a1-4aac-ae97-1c3d0ba72dbd")

type Client struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (*Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: *client,
	}, nil
}

// [Preview API]
func (client *Client) AddAgentCloud(ctx context.Context, args AddAgentCloudArgs) (*TaskAgentCloud, error) {
	if args.AgentCloud == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AgentCloud"}
	}
	body, marshalErr := json.Marshal(*args.AgentCloud)
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

// Arguments for the AddAgentCloud function
type AddAgentCloudArgs struct {
	// (required)
	AgentCloud *TaskAgentCloud
}

// [Preview API]
func (client *Client) DeleteAgentCloud(ctx context.Context, args DeleteAgentCloudArgs) (*TaskAgentCloud, error) {
	routeValues := make(map[string]string)
	if args.AgentCloudId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AgentCloudId"}
	}
	routeValues["agentCloudId"] = strconv.Itoa(*args.AgentCloudId)

	locationId, _ := uuid.Parse("bfa72b3d-0fc6-43fb-932b-a7f6559f93b9")
	resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskAgentCloud
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the DeleteAgentCloud function
type DeleteAgentCloudArgs struct {
	// (required)
	AgentCloudId *int
}

// [Preview API]
func (client *Client) GetAgentCloud(ctx context.Context, args GetAgentCloudArgs) (*TaskAgentCloud, error) {
	routeValues := make(map[string]string)
	if args.AgentCloudId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AgentCloudId"}
	}
	routeValues["agentCloudId"] = strconv.Itoa(*args.AgentCloudId)

	locationId, _ := uuid.Parse("bfa72b3d-0fc6-43fb-932b-a7f6559f93b9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskAgentCloud
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAgentCloud function
type GetAgentCloudArgs struct {
	// (required)
	AgentCloudId *int
}

// [Preview API]
func (client *Client) GetAgentClouds(ctx context.Context, args GetAgentCloudsArgs) (*[]TaskAgentCloud, error) {
	locationId, _ := uuid.Parse("bfa72b3d-0fc6-43fb-932b-a7f6559f93b9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TaskAgentCloud
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAgentClouds function
type GetAgentCloudsArgs struct {
}

// [Preview API] Get agent cloud types.
func (client *Client) GetAgentCloudTypes(ctx context.Context, args GetAgentCloudTypesArgs) (*[]TaskAgentCloudType, error) {
	locationId, _ := uuid.Parse("5932e193-f376-469d-9c3e-e5588ce12cb5")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TaskAgentCloudType
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAgentCloudTypes function
type GetAgentCloudTypesArgs struct {
}

// Adds an agent to a pool.  You probably don't want to call this endpoint directly. Instead, [configure an agent](https://docs.microsoft.com/azure/devops/pipelines/agents/agents) using the agent download package.
func (client *Client) AddAgent(ctx context.Context, args AddAgentArgs) (*TaskAgent, error) {
	if args.Agent == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Agent"}
	}
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)

	body, marshalErr := json.Marshal(*args.Agent)
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

// Arguments for the AddAgent function
type AddAgentArgs struct {
	// (required) Details about the agent being added
	Agent *TaskAgent
	// (required) The agent pool in which to add the agent
	PoolId *int
}

// Delete an agent.  You probably don't want to call this endpoint directly. Instead, [use the agent configuration script](https://docs.microsoft.com/azure/devops/pipelines/agents/agents) to remove an agent from your organization.
func (client *Client) DeleteAgent(ctx context.Context, args DeleteAgentArgs) error {
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)
	if args.AgentId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.AgentId"}
	}
	routeValues["agentId"] = strconv.Itoa(*args.AgentId)

	locationId, _ := uuid.Parse("e298ef32-5878-4cab-993c-043836571f42")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteAgent function
type DeleteAgentArgs struct {
	// (required) The pool ID to remove the agent from
	PoolId *int
	// (required) The agent ID to remove
	AgentId *int
}

// Get information about an agent.
func (client *Client) GetAgent(ctx context.Context, args GetAgentArgs) (*TaskAgent, error) {
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)
	if args.AgentId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AgentId"}
	}
	routeValues["agentId"] = strconv.Itoa(*args.AgentId)

	queryParams := url.Values{}
	if args.IncludeCapabilities != nil {
		queryParams.Add("includeCapabilities", strconv.FormatBool(*args.IncludeCapabilities))
	}
	if args.IncludeAssignedRequest != nil {
		queryParams.Add("includeAssignedRequest", strconv.FormatBool(*args.IncludeAssignedRequest))
	}
	if args.IncludeLastCompletedRequest != nil {
		queryParams.Add("includeLastCompletedRequest", strconv.FormatBool(*args.IncludeLastCompletedRequest))
	}
	if args.PropertyFilters != nil {
		listAsString := strings.Join((*args.PropertyFilters)[:], ",")
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

// Arguments for the GetAgent function
type GetAgentArgs struct {
	// (required) The agent pool containing the agent
	PoolId *int
	// (required) The agent ID to get information about
	AgentId *int
	// (optional) Whether to include the agent's capabilities in the response
	IncludeCapabilities *bool
	// (optional) Whether to include details about the agent's current work
	IncludeAssignedRequest *bool
	// (optional) Whether to include details about the agents' most recent completed work
	IncludeLastCompletedRequest *bool
	// (optional) Filter which custom properties will be returned
	PropertyFilters *[]string
}

// Get a list of agents.
func (client *Client) GetAgents(ctx context.Context, args GetAgentsArgs) (*[]TaskAgent, error) {
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)

	queryParams := url.Values{}
	if args.AgentName != nil {
		queryParams.Add("agentName", *args.AgentName)
	}
	if args.IncludeCapabilities != nil {
		queryParams.Add("includeCapabilities", strconv.FormatBool(*args.IncludeCapabilities))
	}
	if args.IncludeAssignedRequest != nil {
		queryParams.Add("includeAssignedRequest", strconv.FormatBool(*args.IncludeAssignedRequest))
	}
	if args.IncludeLastCompletedRequest != nil {
		queryParams.Add("includeLastCompletedRequest", strconv.FormatBool(*args.IncludeLastCompletedRequest))
	}
	if args.PropertyFilters != nil {
		listAsString := strings.Join((*args.PropertyFilters)[:], ",")
		queryParams.Add("propertyFilters", listAsString)
	}
	if args.Demands != nil {
		listAsString := strings.Join((*args.Demands)[:], ",")
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

// Arguments for the GetAgents function
type GetAgentsArgs struct {
	// (required) The agent pool containing the agents
	PoolId *int
	// (optional) Filter on agent name
	AgentName *string
	// (optional) Whether to include the agents' capabilities in the response
	IncludeCapabilities *bool
	// (optional) Whether to include details about the agents' current work
	IncludeAssignedRequest *bool
	// (optional) Whether to include details about the agents' most recent completed work
	IncludeLastCompletedRequest *bool
	// (optional) Filter which custom properties will be returned
	PropertyFilters *[]string
	// (optional) Filter by demands the agents can satisfy
	Demands *[]string
}

// Replace an agent.  You probably don't want to call this endpoint directly. Instead, [use the agent configuration script](https://docs.microsoft.com/azure/devops/pipelines/agents/agents) to remove and reconfigure an agent from your organization.
func (client *Client) ReplaceAgent(ctx context.Context, args ReplaceAgentArgs) (*TaskAgent, error) {
	if args.Agent == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Agent"}
	}
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)
	if args.AgentId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AgentId"}
	}
	routeValues["agentId"] = strconv.Itoa(*args.AgentId)

	body, marshalErr := json.Marshal(*args.Agent)
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

// Arguments for the ReplaceAgent function
type ReplaceAgentArgs struct {
	// (required) Updated details about the replacing agent
	Agent *TaskAgent
	// (required) The agent pool to use
	PoolId *int
	// (required) The agent to replace
	AgentId *int
}

// Update agent details.
func (client *Client) UpdateAgent(ctx context.Context, args UpdateAgentArgs) (*TaskAgent, error) {
	if args.Agent == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Agent"}
	}
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)
	if args.AgentId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AgentId"}
	}
	routeValues["agentId"] = strconv.Itoa(*args.AgentId)

	body, marshalErr := json.Marshal(*args.Agent)
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

// Arguments for the UpdateAgent function
type UpdateAgentArgs struct {
	// (required) Updated details about the agent
	Agent *TaskAgent
	// (required) The agent pool to use
	PoolId *int
	// (required) The agent to update
	AgentId *int
}

// [Preview API] Create a deployment group.
func (client *Client) AddDeploymentGroup(ctx context.Context, args AddDeploymentGroupArgs) (*DeploymentGroup, error) {
	if args.DeploymentGroup == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroup"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.DeploymentGroup)
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

// Arguments for the AddDeploymentGroup function
type AddDeploymentGroupArgs struct {
	// (required) Deployment group to create.
	DeploymentGroup *DeploymentGroupCreateParameter
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Delete a deployment group.
func (client *Client) DeleteDeploymentGroup(ctx context.Context, args DeleteDeploymentGroupArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DeploymentGroupId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroupId"}
	}
	routeValues["deploymentGroupId"] = strconv.Itoa(*args.DeploymentGroupId)

	locationId, _ := uuid.Parse("083c4d89-ab35-45af-aa11-7cf66895c53e")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteDeploymentGroup function
type DeleteDeploymentGroupArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the deployment group to be deleted.
	DeploymentGroupId *int
}

// [Preview API] Get a deployment group by its ID.
func (client *Client) GetDeploymentGroup(ctx context.Context, args GetDeploymentGroupArgs) (*DeploymentGroup, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DeploymentGroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroupId"}
	}
	routeValues["deploymentGroupId"] = strconv.Itoa(*args.DeploymentGroupId)

	queryParams := url.Values{}
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
	}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
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

// Arguments for the GetDeploymentGroup function
type GetDeploymentGroupArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the deployment group.
	DeploymentGroupId *int
	// (optional) Get the deployment group only if this action can be performed on it.
	ActionFilter *DeploymentGroupActionFilter
	// (optional) Include these additional details in the returned object.
	Expand *DeploymentGroupExpands
}

// [Preview API] Get a list of deployment groups by name or IDs.
func (client *Client) GetDeploymentGroups(ctx context.Context, args GetDeploymentGroupsArgs) (*GetDeploymentGroupsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Name != nil {
		queryParams.Add("name", *args.Name)
	}
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
	}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Ids != nil {
		var stringList []string
		for _, item := range *args.Ids {
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

	var responseValue GetDeploymentGroupsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetDeploymentGroups function
type GetDeploymentGroupsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Name of the deployment group.
	Name *string
	// (optional) Get only deployment groups on which this action can be performed.
	ActionFilter *DeploymentGroupActionFilter
	// (optional) Include these additional details in the returned objects.
	Expand *DeploymentGroupExpands
	// (optional) Get deployment groups with names greater than this continuationToken lexicographically.
	ContinuationToken *string
	// (optional) Maximum number of deployment groups to return. Default is **1000**.
	Top *int
	// (optional) Comma separated list of IDs of the deployment groups.
	Ids *[]int
}

// Return type for the GetDeploymentGroups function
type GetDeploymentGroupsResponseValue struct {
	Value []DeploymentGroup
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// [Preview API] Update a deployment group.
func (client *Client) UpdateDeploymentGroup(ctx context.Context, args UpdateDeploymentGroupArgs) (*DeploymentGroup, error) {
	if args.DeploymentGroup == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroup"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DeploymentGroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroupId"}
	}
	routeValues["deploymentGroupId"] = strconv.Itoa(*args.DeploymentGroupId)

	body, marshalErr := json.Marshal(*args.DeploymentGroup)
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

// Arguments for the UpdateDeploymentGroup function
type UpdateDeploymentGroupArgs struct {
	// (required) Deployment group to update.
	DeploymentGroup *DeploymentGroupUpdateParameter
	// (required) Project ID or project name
	Project *string
	// (required) ID of the deployment group.
	DeploymentGroupId *int
}

// Create an agent pool.
func (client *Client) AddAgentPool(ctx context.Context, args AddAgentPoolArgs) (*TaskAgentPool, error) {
	if args.Pool == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Pool"}
	}
	body, marshalErr := json.Marshal(*args.Pool)
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

// Arguments for the AddAgentPool function
type AddAgentPoolArgs struct {
	// (required) Details about the new agent pool
	Pool *TaskAgentPool
}

// Delete an agent pool.
func (client *Client) DeleteAgentPool(ctx context.Context, args DeleteAgentPoolArgs) error {
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)

	locationId, _ := uuid.Parse("a8c47e17-4d56-4a56-92bb-de7ea7dc65be")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteAgentPool function
type DeleteAgentPoolArgs struct {
	// (required) ID of the agent pool to delete
	PoolId *int
}

// Get information about an agent pool.
func (client *Client) GetAgentPool(ctx context.Context, args GetAgentPoolArgs) (*TaskAgentPool, error) {
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)

	queryParams := url.Values{}
	if args.Properties != nil {
		listAsString := strings.Join((*args.Properties)[:], ",")
		queryParams.Add("properties", listAsString)
	}
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
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

// Arguments for the GetAgentPool function
type GetAgentPoolArgs struct {
	// (required) An agent pool ID
	PoolId *int
	// (optional) Agent pool properties (comma-separated)
	Properties *[]string
	// (optional) Filter by whether the calling user has use or manage permissions
	ActionFilter *TaskAgentPoolActionFilter
}

// Get a list of agent pools.
func (client *Client) GetAgentPools(ctx context.Context, args GetAgentPoolsArgs) (*[]TaskAgentPool, error) {
	queryParams := url.Values{}
	if args.PoolName != nil {
		queryParams.Add("poolName", *args.PoolName)
	}
	if args.Properties != nil {
		listAsString := strings.Join((*args.Properties)[:], ",")
		queryParams.Add("properties", listAsString)
	}
	if args.PoolType != nil {
		queryParams.Add("poolType", string(*args.PoolType))
	}
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
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

// Arguments for the GetAgentPools function
type GetAgentPoolsArgs struct {
	// (optional) Filter by name
	PoolName *string
	// (optional) Filter by agent pool properties (comma-separated)
	Properties *[]string
	// (optional) Filter by pool type
	PoolType *TaskAgentPoolType
	// (optional) Filter by whether the calling user has use or manage permissions
	ActionFilter *TaskAgentPoolActionFilter
}

// Get a list of agent pools.
func (client *Client) GetAgentPoolsByIds(ctx context.Context, args GetAgentPoolsByIdsArgs) (*[]TaskAgentPool, error) {
	queryParams := url.Values{}
	if args.PoolIds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "poolIds"}
	}
	var stringList []string
	for _, item := range *args.PoolIds {
		stringList = append(stringList, strconv.Itoa(item))
	}
	listAsString := strings.Join((stringList)[:], ",")
	queryParams.Add("definitions", listAsString)
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
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

// Arguments for the GetAgentPoolsByIds function
type GetAgentPoolsByIdsArgs struct {
	// (required) pool Ids to fetch
	PoolIds *[]int
	// (optional) Filter by whether the calling user has use or manage permissions
	ActionFilter *TaskAgentPoolActionFilter
}

// Update properties on an agent pool
func (client *Client) UpdateAgentPool(ctx context.Context, args UpdateAgentPoolArgs) (*TaskAgentPool, error) {
	if args.Pool == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Pool"}
	}
	routeValues := make(map[string]string)
	if args.PoolId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PoolId"}
	}
	routeValues["poolId"] = strconv.Itoa(*args.PoolId)

	body, marshalErr := json.Marshal(*args.Pool)
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

// Arguments for the UpdateAgentPool function
type UpdateAgentPoolArgs struct {
	// (required) Updated agent pool details
	Pool *TaskAgentPool
	// (required) The agent pool to update
	PoolId *int
}

// [Preview API] Create a new agent queue to connect a project to an agent pool.
func (client *Client) AddAgentQueue(ctx context.Context, args AddAgentQueueArgs) (*TaskAgentQueue, error) {
	if args.Queue == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Queue"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.AuthorizePipelines != nil {
		queryParams.Add("authorizePipelines", strconv.FormatBool(*args.AuthorizePipelines))
	}
	body, marshalErr := json.Marshal(*args.Queue)
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

// Arguments for the AddAgentQueue function
type AddAgentQueueArgs struct {
	// (required) Details about the queue to create
	Queue *TaskAgentQueue
	// (optional) Project ID or project name
	Project *string
	// (optional) Automatically authorize this queue when using YAML
	AuthorizePipelines *bool
}

// [Preview API] Removes an agent queue from a project.
func (client *Client) DeleteAgentQueue(ctx context.Context, args DeleteAgentQueueArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.QueueId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.QueueId"}
	}
	routeValues["queueId"] = strconv.Itoa(*args.QueueId)

	locationId, _ := uuid.Parse("900fa995-c559-4923-aae7-f8424fe4fbea")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteAgentQueue function
type DeleteAgentQueueArgs struct {
	// (required) The agent queue to remove
	QueueId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get information about an agent queue.
func (client *Client) GetAgentQueue(ctx context.Context, args GetAgentQueueArgs) (*TaskAgentQueue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.QueueId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.QueueId"}
	}
	routeValues["queueId"] = strconv.Itoa(*args.QueueId)

	queryParams := url.Values{}
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
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

// Arguments for the GetAgentQueue function
type GetAgentQueueArgs struct {
	// (required) The agent queue to get information about
	QueueId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Filter by whether the calling user has use or manage permissions
	ActionFilter *TaskAgentQueueActionFilter
}

// [Preview API] Get a list of agent queues.
func (client *Client) GetAgentQueues(ctx context.Context, args GetAgentQueuesArgs) (*[]TaskAgentQueue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.QueueName != nil {
		queryParams.Add("queueName", *args.QueueName)
	}
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
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

// Arguments for the GetAgentQueues function
type GetAgentQueuesArgs struct {
	// (optional) Project ID or project name
	Project *string
	// (optional) Filter on the agent queue name
	QueueName *string
	// (optional) Filter by whether the calling user has use or manage permissions
	ActionFilter *TaskAgentQueueActionFilter
}

// [Preview API] Get a list of agent queues by their IDs
func (client *Client) GetAgentQueuesByIds(ctx context.Context, args GetAgentQueuesByIdsArgs) (*[]TaskAgentQueue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.QueueIds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "queueIds"}
	}
	var stringList []string
	for _, item := range *args.QueueIds {
		stringList = append(stringList, strconv.Itoa(item))
	}
	listAsString := strings.Join((stringList)[:], ",")
	queryParams.Add("definitions", listAsString)
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
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

// Arguments for the GetAgentQueuesByIds function
type GetAgentQueuesByIdsArgs struct {
	// (required) A comma-separated list of agent queue IDs to retrieve
	QueueIds *[]int
	// (optional) Project ID or project name
	Project *string
	// (optional) Filter by whether the calling user has use or manage permissions
	ActionFilter *TaskAgentQueueActionFilter
}

// [Preview API] Get a list of agent queues by their names
func (client *Client) GetAgentQueuesByNames(ctx context.Context, args GetAgentQueuesByNamesArgs) (*[]TaskAgentQueue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.QueueNames == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "queueNames"}
	}
	listAsString := strings.Join((*args.QueueNames)[:], ",")
	queryParams.Add("queueNames", listAsString)
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
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

// Arguments for the GetAgentQueuesByNames function
type GetAgentQueuesByNamesArgs struct {
	// (required) A comma-separated list of agent names to retrieve
	QueueNames *[]string
	// (optional) Project ID or project name
	Project *string
	// (optional) Filter by whether the calling user has use or manage permissions
	ActionFilter *TaskAgentQueueActionFilter
}

// [Preview API]
func (client *Client) GetAgentCloudRequests(ctx context.Context, args GetAgentCloudRequestsArgs) (*[]TaskAgentCloudRequest, error) {
	routeValues := make(map[string]string)
	if args.AgentCloudId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AgentCloudId"}
	}
	routeValues["agentCloudId"] = strconv.Itoa(*args.AgentCloudId)

	locationId, _ := uuid.Parse("20189bd7-5134-49c2-b8e9-f9e856eea2b2")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TaskAgentCloudRequest
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAgentCloudRequests function
type GetAgentCloudRequestsArgs struct {
	// (required)
	AgentCloudId *int
}

// [Preview API] Delete a deployment target in a deployment group. This deletes the agent from associated deployment pool too.
func (client *Client) DeleteDeploymentTarget(ctx context.Context, args DeleteDeploymentTargetArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DeploymentGroupId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroupId"}
	}
	routeValues["deploymentGroupId"] = strconv.Itoa(*args.DeploymentGroupId)
	if args.TargetId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.TargetId"}
	}
	routeValues["targetId"] = strconv.Itoa(*args.TargetId)

	locationId, _ := uuid.Parse("2f0aa599-c121-4256-a5fd-ba370e0ae7b6")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteDeploymentTarget function
type DeleteDeploymentTargetArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the deployment group in which deployment target is deleted.
	DeploymentGroupId *int
	// (required) ID of the deployment target to delete.
	TargetId *int
}

// [Preview API] Get a deployment target by its ID in a deployment group
func (client *Client) GetDeploymentTarget(ctx context.Context, args GetDeploymentTargetArgs) (*DeploymentMachine, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DeploymentGroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroupId"}
	}
	routeValues["deploymentGroupId"] = strconv.Itoa(*args.DeploymentGroupId)
	if args.TargetId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TargetId"}
	}
	routeValues["targetId"] = strconv.Itoa(*args.TargetId)

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
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

// Arguments for the GetDeploymentTarget function
type GetDeploymentTargetArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the deployment group to which deployment target belongs.
	DeploymentGroupId *int
	// (required) ID of the deployment target to return.
	TargetId *int
	// (optional) Include these additional details in the returned objects.
	Expand *DeploymentTargetExpands
}

// [Preview API] Get a list of deployment targets in a deployment group.
func (client *Client) GetDeploymentTargets(ctx context.Context, args GetDeploymentTargetsArgs) (*GetDeploymentTargetsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DeploymentGroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroupId"}
	}
	routeValues["deploymentGroupId"] = strconv.Itoa(*args.DeploymentGroupId)

	queryParams := url.Values{}
	if args.Tags != nil {
		listAsString := strings.Join((*args.Tags)[:], ",")
		queryParams.Add("tags", listAsString)
	}
	if args.Name != nil {
		queryParams.Add("name", *args.Name)
	}
	if args.PartialNameMatch != nil {
		queryParams.Add("partialNameMatch", strconv.FormatBool(*args.PartialNameMatch))
	}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	if args.AgentStatus != nil {
		queryParams.Add("agentStatus", string(*args.AgentStatus))
	}
	if args.AgentJobResult != nil {
		queryParams.Add("agentJobResult", string(*args.AgentJobResult))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Enabled != nil {
		queryParams.Add("enabled", strconv.FormatBool(*args.Enabled))
	}
	if args.PropertyFilters != nil {
		listAsString := strings.Join((*args.PropertyFilters)[:], ",")
		queryParams.Add("propertyFilters", listAsString)
	}
	locationId, _ := uuid.Parse("2f0aa599-c121-4256-a5fd-ba370e0ae7b6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetDeploymentTargetsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetDeploymentTargets function
type GetDeploymentTargetsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the deployment group.
	DeploymentGroupId *int
	// (optional) Get only the deployment targets that contain all these comma separted list of tags.
	Tags *[]string
	// (optional) Name pattern of the deployment targets to return.
	Name *string
	// (optional) When set to true, treats **name** as pattern. Else treats it as absolute match. Default is **false**.
	PartialNameMatch *bool
	// (optional) Include these additional details in the returned objects.
	Expand *DeploymentTargetExpands
	// (optional) Get only deployment targets that have this status.
	AgentStatus *TaskAgentStatusFilter
	// (optional) Get only deployment targets that have this last job result.
	AgentJobResult *TaskAgentJobResultFilter
	// (optional) Get deployment targets with names greater than this continuationToken lexicographically.
	ContinuationToken *string
	// (optional) Maximum number of deployment targets to return. Default is **1000**.
	Top *int
	// (optional) Get only deployment targets that are enabled or disabled. Default is 'null' which returns all the targets.
	Enabled *bool
	// (optional)
	PropertyFilters *[]string
}

// Return type for the GetDeploymentTargets function
type GetDeploymentTargetsResponseValue struct {
	Value []DeploymentMachine
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// [Preview API] Update tags of a list of deployment targets in a deployment group.
func (client *Client) UpdateDeploymentTargets(ctx context.Context, args UpdateDeploymentTargetsArgs) (*[]DeploymentMachine, error) {
	if args.Machines == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Machines"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.DeploymentGroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.DeploymentGroupId"}
	}
	routeValues["deploymentGroupId"] = strconv.Itoa(*args.DeploymentGroupId)

	body, marshalErr := json.Marshal(*args.Machines)
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

// Arguments for the UpdateDeploymentTargets function
type UpdateDeploymentTargetsArgs struct {
	// (required) Deployment targets with tags to udpdate.
	Machines *[]DeploymentTargetUpdateParameter
	// (required) Project ID or project name
	Project *string
	// (required) ID of the deployment group in which deployment targets are updated.
	DeploymentGroupId *int
}

// [Preview API] Create a task group.
func (client *Client) AddTaskGroup(ctx context.Context, args AddTaskGroupArgs) (*TaskGroup, error) {
	if args.TaskGroup == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TaskGroup"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.TaskGroup)
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

// Arguments for the AddTaskGroup function
type AddTaskGroupArgs struct {
	// (required) Task group object to create.
	TaskGroup *TaskGroupCreateParameter
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Delete a task group.
func (client *Client) DeleteTaskGroup(ctx context.Context, args DeleteTaskGroupArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TaskGroupId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.TaskGroupId"}
	}
	routeValues["taskGroupId"] = (*args.TaskGroupId).String()

	queryParams := url.Values{}
	if args.Comment != nil {
		queryParams.Add("comment", *args.Comment)
	}
	locationId, _ := uuid.Parse("6c08ffbf-dbf1-4f9a-94e5-a1cbd47005e7")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTaskGroup function
type DeleteTaskGroupArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the task group to be deleted.
	TaskGroupId *uuid.UUID
	// (optional) Comments to delete.
	Comment *string
}

// [Preview API] List task groups.
func (client *Client) GetTaskGroups(ctx context.Context, args GetTaskGroupsArgs) (*[]TaskGroup, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TaskGroupId != nil {
		routeValues["taskGroupId"] = (*args.TaskGroupId).String()
	}

	queryParams := url.Values{}
	if args.Expanded != nil {
		queryParams.Add("expanded", strconv.FormatBool(*args.Expanded))
	}
	if args.TaskIdFilter != nil {
		queryParams.Add("taskIdFilter", (*args.TaskIdFilter).String())
	}
	if args.Deleted != nil {
		queryParams.Add("deleted", strconv.FormatBool(*args.Deleted))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", (*args.ContinuationToken).String())
	}
	if args.QueryOrder != nil {
		queryParams.Add("queryOrder", string(*args.QueryOrder))
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

// Arguments for the GetTaskGroups function
type GetTaskGroupsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Id of the task group.
	TaskGroupId *uuid.UUID
	// (optional) 'true' to recursively expand task groups. Default is 'false'.
	Expanded *bool
	// (optional) Guid of the taskId to filter.
	TaskIdFilter *uuid.UUID
	// (optional) 'true'to include deleted task groups. Default is 'false'.
	Deleted *bool
	// (optional) Number of task groups to get.
	Top *int
	// (optional) Gets the task groups after the continuation token provided.
	ContinuationToken *time.Time
	// (optional) Gets the results in the defined order. Default is 'CreatedOnDescending'.
	QueryOrder *TaskGroupQueryOrder
}

// [Preview API] Update a task group.
func (client *Client) UpdateTaskGroup(ctx context.Context, args UpdateTaskGroupArgs) (*TaskGroup, error) {
	if args.TaskGroup == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TaskGroup"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TaskGroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TaskGroupId"}
	}
	routeValues["taskGroupId"] = (*args.TaskGroupId).String()

	body, marshalErr := json.Marshal(*args.TaskGroup)
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

// Arguments for the UpdateTaskGroup function
type UpdateTaskGroupArgs struct {
	// (required) Task group to update.
	TaskGroup *TaskGroupUpdateParameter
	// (required) Project ID or project name
	Project *string
	// (required) Id of the task group to update.
	TaskGroupId *uuid.UUID
}

// [Preview API] Add a variable group.
func (client *Client) AddVariableGroup(ctx context.Context, args AddVariableGroupArgs) (*VariableGroup, error) {
	if args.Group == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Group"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.Group)
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

// Arguments for the AddVariableGroup function
type AddVariableGroupArgs struct {
	// (required) Variable group to add.
	Group *VariableGroupParameters
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Delete a variable group
func (client *Client) DeleteVariableGroup(ctx context.Context, args DeleteVariableGroupArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.GroupId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = strconv.Itoa(*args.GroupId)

	locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteVariableGroup function
type DeleteVariableGroupArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the variable group.
	GroupId *int
}

// [Preview API] Get a variable group.
func (client *Client) GetVariableGroup(ctx context.Context, args GetVariableGroupArgs) (*VariableGroup, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.GroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = strconv.Itoa(*args.GroupId)

	locationId, _ := uuid.Parse("f5b09dd5-9d54-45a1-8b5a-1c8287d634cc")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue VariableGroup
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetVariableGroup function
type GetVariableGroupArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the variable group.
	GroupId *int
}

// [Preview API] Get variable groups.
func (client *Client) GetVariableGroups(ctx context.Context, args GetVariableGroupsArgs) (*[]VariableGroup, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.GroupName != nil {
		queryParams.Add("groupName", *args.GroupName)
	}
	if args.ActionFilter != nil {
		queryParams.Add("actionFilter", string(*args.ActionFilter))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", strconv.Itoa(*args.ContinuationToken))
	}
	if args.QueryOrder != nil {
		queryParams.Add("queryOrder", string(*args.QueryOrder))
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

// Arguments for the GetVariableGroups function
type GetVariableGroupsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Name of variable group.
	GroupName *string
	// (optional) Action filter for the variable group. It specifies the action which can be performed on the variable groups.
	ActionFilter *VariableGroupActionFilter
	// (optional) Number of variable groups to get.
	Top *int
	// (optional) Gets the variable groups after the continuation token provided.
	ContinuationToken *int
	// (optional) Gets the results in the defined order. Default is 'IdDescending'.
	QueryOrder *VariableGroupQueryOrder
}

// [Preview API] Get variable groups by ids.
func (client *Client) GetVariableGroupsById(ctx context.Context, args GetVariableGroupsByIdArgs) (*[]VariableGroup, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.GroupIds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "groupIds"}
	}
	var stringList []string
	for _, item := range *args.GroupIds {
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

// Arguments for the GetVariableGroupsById function
type GetVariableGroupsByIdArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Comma separated list of Ids of variable groups.
	GroupIds *[]int
}

// [Preview API] Update a variable group.
func (client *Client) UpdateVariableGroup(ctx context.Context, args UpdateVariableGroupArgs) (*VariableGroup, error) {
	if args.Group == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Group"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.GroupId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = strconv.Itoa(*args.GroupId)

	body, marshalErr := json.Marshal(*args.Group)
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

// Arguments for the UpdateVariableGroup function
type UpdateVariableGroupArgs struct {
	// (required) Variable group to update.
	Group *VariableGroupParameters
	// (required) Project ID or project name
	Project *string
	// (required) Id of the variable group to update.
	GroupId *int
}

func (client *Client) GetYamlSchema(ctx context.Context, args GetYamlSchemaArgs) (interface{}, error) {
	locationId, _ := uuid.Parse("1f9990b9-1dba-441f-9c2e-6485888c42b6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the GetYamlSchema function
type GetYamlSchemaArgs struct {
}
