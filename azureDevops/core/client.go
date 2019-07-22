// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package core

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
)

var ResourceAreaId, _ = uuid.Parse("79134c72-4a58-4b42-976c-04e7115f32bf")

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

// [Preview API] Removes the avatar for the project.
// ctx
// projectId (required): The ID or name of the project.
func (client Client) RemoveProjectAvatar(ctx context.Context, projectId *string) error {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId

    locationId, _ := uuid.Parse("54b2a2a0-859b-4d05-827c-ec4c862f641a")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Sets the avatar for the project.
// ctx
// avatarBlob (required): The avatar blob data object to upload.
// projectId (required): The ID or name of the project.
func (client Client) SetProjectAvatar(ctx context.Context, avatarBlob *ProjectAvatar, projectId *string) error {
    if avatarBlob == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "avatarBlob"}
    }
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId

    body, marshalErr := json.Marshal(*avatarBlob)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("54b2a2a0-859b-4d05-827c-ec4c862f641a")
    _, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// connectedServiceCreationData (required)
// projectId (required)
func (client Client) CreateConnectedService(ctx context.Context, connectedServiceCreationData *WebApiConnectedServiceDetails, projectId *string) (*WebApiConnectedService, error) {
    if connectedServiceCreationData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "connectedServiceCreationData"}
    }
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId

    body, marshalErr := json.Marshal(*connectedServiceCreationData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b4f70219-e18b-42c5-abe3-98b07d35525e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WebApiConnectedService
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// projectId (required)
// name (required)
func (client Client) GetConnectedServiceDetails(ctx context.Context, projectId *string, name *string) (*WebApiConnectedServiceDetails, error) {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId
    if name == nil || *name == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("b4f70219-e18b-42c5-abe3-98b07d35525e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WebApiConnectedServiceDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// projectId (required)
// kind (optional)
func (client Client) GetConnectedServices(ctx context.Context, projectId *string, kind *string) (*[]WebApiConnectedService, error) {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId

    queryParams := url.Values{}
    if kind != nil {
        queryParams.Add("kind", *kind)
    }
    locationId, _ := uuid.Parse("b4f70219-e18b-42c5-abe3-98b07d35525e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WebApiConnectedService
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of members for a specific team.
// ctx
// projectId (required): The name or ID (GUID) of the team project the team belongs to.
// teamId (required): The name or ID (GUID) of the team .
// top (optional)
// skip (optional)
func (client Client) GetTeamMembersWithExtendedProperties(ctx context.Context, projectId *string, teamId *string, top *int, skip *int) (*[]TeamMember, error) {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId
    if teamId == nil || *teamId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "teamId"} 
    }
    routeValues["teamId"] = *teamId

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("294c494c-2600-4d7e-b76c-3dd50c3c95be")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TeamMember
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a process by ID.
// ctx
// processId (required): ID for a process.
func (client Client) GetProcessById(ctx context.Context, processId *uuid.UUID) (*Process, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()

    locationId, _ := uuid.Parse("93878975-88c5-4e6a-8abb-7ddd77a8a7d8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Process
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of processes.
// ctx
func (client Client) GetProcesses(ctx context.Context, ) (*[]Process, error) {
    locationId, _ := uuid.Parse("93878975-88c5-4e6a-8abb-7ddd77a8a7d8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Process
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get project collection with the specified id or name.
// ctx
// collectionId (required)
func (client Client) GetProjectCollection(ctx context.Context, collectionId *string) (*TeamProjectCollection, error) {
    routeValues := make(map[string]string)
    if collectionId == nil || *collectionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "collectionId"} 
    }
    routeValues["collectionId"] = *collectionId

    locationId, _ := uuid.Parse("8031090f-ef1d-4af6-85fc-698cd75d42bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamProjectCollection
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get project collection references for this application.
// ctx
// top (optional)
// skip (optional)
func (client Client) GetProjectCollections(ctx context.Context, top *int, skip *int) (*[]TeamProjectCollectionReference, error) {
    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("8031090f-ef1d-4af6-85fc-698cd75d42bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TeamProjectCollectionReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get project with the specified id or name, optionally including capabilities.
// ctx
// projectId (required)
// includeCapabilities (optional): Include capabilities (such as source control) in the team project result (default: false).
// includeHistory (optional): Search within renamed projects (that had such name in the past).
func (client Client) GetProject(ctx context.Context, projectId *string, includeCapabilities *bool, includeHistory *bool) (*TeamProject, error) {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId

    queryParams := url.Values{}
    if includeCapabilities != nil {
        queryParams.Add("includeCapabilities", strconv.FormatBool(*includeCapabilities))
    }
    if includeHistory != nil {
        queryParams.Add("includeHistory", strconv.FormatBool(*includeHistory))
    }
    locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamProject
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get all projects in the organization that the authenticated user has access to.
// ctx
// stateFilter (optional): Filter on team projects in a specific team project state (default: WellFormed).
// top (optional)
// skip (optional)
// continuationToken (optional)
// getDefaultTeamImageUrl (optional)
func (client Client) GetProjects(ctx context.Context, stateFilter *string, top *int, skip *int, continuationToken *string, getDefaultTeamImageUrl *bool) (*[]TeamProjectReference, error) {
    queryParams := url.Values{}
    if stateFilter != nil {
        queryParams.Add("stateFilter", *stateFilter)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if getDefaultTeamImageUrl != nil {
        queryParams.Add("getDefaultTeamImageUrl", strconv.FormatBool(*getDefaultTeamImageUrl))
    }
    locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TeamProjectReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Queues a project to be created. Use the [GetOperation](../../operations/operations/get) to periodically check for create project status.
// ctx
// projectToCreate (required): The project to create.
func (client Client) QueueCreateProject(ctx context.Context, projectToCreate *TeamProject) (*OperationReference, error) {
    if projectToCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "projectToCreate"}
    }
    body, marshalErr := json.Marshal(*projectToCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue OperationReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Queues a project to be deleted. Use the [GetOperation](../../operations/operations/get) to periodically check for delete project status.
// ctx
// projectId (required): The project id of the project to delete.
func (client Client) QueueDeleteProject(ctx context.Context, projectId *uuid.UUID) (*OperationReference, error) {
    routeValues := make(map[string]string)
    if projectId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = (*projectId).String()

    locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue OperationReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update an existing project's name, abbreviation, description, or restore a project.
// ctx
// projectUpdate (required): The updates for the project. The state must be set to wellFormed to restore the project.
// projectId (required): The project id of the project to update.
func (client Client) UpdateProject(ctx context.Context, projectUpdate *TeamProject, projectId *uuid.UUID) (*OperationReference, error) {
    if projectUpdate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "projectUpdate"}
    }
    routeValues := make(map[string]string)
    if projectId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = (*projectId).String()

    body, marshalErr := json.Marshal(*projectUpdate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue OperationReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a collection of team project properties.
// ctx
// projectId (required): The team project ID.
// keys (optional): A comma-delimited string of team project property names. Wildcard characters ("?" and "*") are supported. If no key is specified, all properties will be returned.
func (client Client) GetProjectProperties(ctx context.Context, projectId *uuid.UUID, keys *[]string) (*[]ProjectProperty, error) {
    routeValues := make(map[string]string)
    if projectId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = (*projectId).String()

    queryParams := url.Values{}
    if keys != nil {
        listAsString := strings.Join((*keys)[:], ",")
        queryParams.Add("keys", listAsString)
    }
    locationId, _ := uuid.Parse("4976a71a-4487-49aa-8aab-a1eda469037a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProjectProperty
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create, update, and delete team project properties.
// ctx
// projectId (required): The team project ID.
// patchDocument (required): A JSON Patch document that represents an array of property operations. See RFC 6902 for more details on JSON Patch. The accepted operation verbs are Add and Remove, where Add is used for both creating and updating properties. The path consists of a forward slash and a property name.
func (client Client) SetProjectProperties(ctx context.Context, projectId *uuid.UUID, patchDocument *[]JsonPatchOperation) error {
    if patchDocument == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "patchDocument"}
    }
    routeValues := make(map[string]string)
    if projectId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = (*projectId).String()

    body, marshalErr := json.Marshal(*patchDocument)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4976a71a-4487-49aa-8aab-a1eda469037a")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// proxy (required)
func (client Client) CreateOrUpdateProxy(ctx context.Context, proxy *Proxy) (*Proxy, error) {
    if proxy == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "proxy"}
    }
    body, marshalErr := json.Marshal(*proxy)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ec1f4311-f2b4-4c15-b2b8-8990b80d2908")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Proxy
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// proxyUrl (required)
// site (optional)
func (client Client) DeleteProxy(ctx context.Context, proxyUrl *string, site *string) error {
    queryParams := url.Values{}
    if proxyUrl == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "proxyUrl"}
    }
    queryParams.Add("proxyUrl", *proxyUrl)
    if site != nil {
        queryParams.Add("site", *site)
    }
    locationId, _ := uuid.Parse("ec1f4311-f2b4-4c15-b2b8-8990b80d2908")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// proxyUrl (optional)
func (client Client) GetProxies(ctx context.Context, proxyUrl *string) (*[]Proxy, error) {
    queryParams := url.Values{}
    if proxyUrl != nil {
        queryParams.Add("proxyUrl", *proxyUrl)
    }
    locationId, _ := uuid.Parse("ec1f4311-f2b4-4c15-b2b8-8990b80d2908")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Proxy
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a team in a team project.
// ctx
// team (required): The team data used to create the team.
// projectId (required): The name or ID (GUID) of the team project in which to create the team.
func (client Client) CreateTeam(ctx context.Context, team *WebApiTeam, projectId *string) (*WebApiTeam, error) {
    if team == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "team"}
    }
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId

    body, marshalErr := json.Marshal(*team)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WebApiTeam
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a team.
// ctx
// projectId (required): The name or ID (GUID) of the team project containing the team to delete.
// teamId (required): The name or ID of the team to delete.
func (client Client) DeleteTeam(ctx context.Context, projectId *string, teamId *string) error {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId
    if teamId == nil || *teamId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "teamId"} 
    }
    routeValues["teamId"] = *teamId

    locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a specific team.
// ctx
// projectId (required): The name or ID (GUID) of the team project containing the team.
// teamId (required): The name or ID (GUID) of the team.
// expandIdentity (optional): A value indicating whether or not to expand Identity information in the result WebApiTeam object.
func (client Client) GetTeam(ctx context.Context, projectId *string, teamId *string, expandIdentity *bool) (*WebApiTeam, error) {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId
    if teamId == nil || *teamId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "teamId"} 
    }
    routeValues["teamId"] = *teamId

    queryParams := url.Values{}
    if expandIdentity != nil {
        queryParams.Add("$expandIdentity", strconv.FormatBool(*expandIdentity))
    }
    locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WebApiTeam
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of teams.
// ctx
// projectId (required)
// mine (optional): If true return all the teams requesting user is member, otherwise return all the teams user has read access.
// top (optional): Maximum number of teams to return.
// skip (optional): Number of teams to skip.
// expandIdentity (optional): A value indicating whether or not to expand Identity information in the result WebApiTeam object.
func (client Client) GetTeams(ctx context.Context, projectId *string, mine *bool, top *int, skip *int, expandIdentity *bool) (*[]WebApiTeam, error) {
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId

    queryParams := url.Values{}
    if mine != nil {
        queryParams.Add("$mine", strconv.FormatBool(*mine))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if expandIdentity != nil {
        queryParams.Add("$expandIdentity", strconv.FormatBool(*expandIdentity))
    }
    locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WebApiTeam
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a team's name and/or description.
// ctx
// teamData (required)
// projectId (required): The name or ID (GUID) of the team project containing the team to update.
// teamId (required): The name of ID of the team to update.
func (client Client) UpdateTeam(ctx context.Context, teamData *WebApiTeam, projectId *string, teamId *string) (*WebApiTeam, error) {
    if teamData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "teamData"}
    }
    routeValues := make(map[string]string)
    if projectId == nil || *projectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "projectId"} 
    }
    routeValues["projectId"] = *projectId
    if teamId == nil || *teamId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "teamId"} 
    }
    routeValues["teamId"] = *teamId

    body, marshalErr := json.Marshal(*teamData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WebApiTeam
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of all teams.
// ctx
// mine (optional): If true, then return all teams requesting user is member. Otherwise return all teams user has read access.
// top (optional): Maximum number of teams to return.
// skip (optional): Number of teams to skip.
// expandIdentity (optional): A value indicating whether or not to expand Identity information in the result WebApiTeam object.
func (client Client) GetAllTeams(ctx context.Context, mine *bool, top *int, skip *int, expandIdentity *bool) (*[]WebApiTeam, error) {
    queryParams := url.Values{}
    if mine != nil {
        queryParams.Add("$mine", strconv.FormatBool(*mine))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if expandIdentity != nil {
        queryParams.Add("$expandIdentity", strconv.FormatBool(*expandIdentity))
    }
    locationId, _ := uuid.Parse("7a4d9ee9-3433-4347-b47a-7a80f1cf307e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.3", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WebApiTeam
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

