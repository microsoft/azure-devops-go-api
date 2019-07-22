// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package work

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

var ResourceAreaId, _ = uuid.Parse("1d4f49f9-02b9-4e26-b826-2cdb6195f2a9")

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

// Gets backlog configuration for a team
// ctx
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) GetBacklogConfigurations(ctx context.Context, project *string, team *string) (*BacklogConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    locationId, _ := uuid.Parse("7799f497-3cb5-4f16-ad4f-5cd06012db64")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BacklogConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of work items within a backlog level
// ctx
// project (required): Project ID or project name
// team (required): Team ID or team name
// backlogId (required)
func (client Client) GetBacklogLevelWorkItems(ctx context.Context, project *string, team *string, backlogId *string) (*BacklogLevelWorkItems, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team
    if backlogId == nil || *backlogId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "backlogId"} 
    }
    routeValues["backlogId"] = *backlogId

    locationId, _ := uuid.Parse("7c468d96-ab1d-4294-a360-92f07e9ccd98")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BacklogLevelWorkItems
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a backlog level
// ctx
// project (required): Project ID or project name
// team (required): Team ID or team name
// id (required): The id of the backlog level
func (client Client) GetBacklog(ctx context.Context, project *string, team *string, id *string) (*BacklogLevelConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team
    if id == nil || *id == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    locationId, _ := uuid.Parse("a93726f9-7867-4e38-b4f2-0bfafc2f6a94")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BacklogLevelConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] List all backlog levels
// ctx
// project (required): Project ID or project name
// team (required): Team ID or team name
func (client Client) GetBacklogs(ctx context.Context, project *string, team *string) (*[]BacklogLevelConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team

    locationId, _ := uuid.Parse("a93726f9-7867-4e38-b4f2-0bfafc2f6a94")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BacklogLevelConfiguration
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get available board columns in a project
// ctx
// project (optional): Project ID or project name
func (client Client) GetColumnSuggestedValues(ctx context.Context, project *string) (*[]BoardSuggestedValue, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    locationId, _ := uuid.Parse("eb7ec5a3-1ba3-4fd1-b834-49a5a387e57d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardSuggestedValue
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns the list of parent field filter model for the given list of workitem ids
// ctx
// project (required): Project ID or project name
// childBacklogContextCategoryRefName (required)
// workitemIds (required)
// team (optional): Team ID or team name
func (client Client) GetBoardMappingParentItems(ctx context.Context, project *string, childBacklogContextCategoryRefName *string, workitemIds *[]int, team *string) (*[]ParentChildWIMap, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    queryParams := url.Values{}
    if childBacklogContextCategoryRefName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "childBacklogContextCategoryRefName"}
    }
    queryParams.Add("childBacklogContextCategoryRefName", *childBacklogContextCategoryRefName)
    if workitemIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workitemIds"}
    }
    var stringList []string
    for _, item := range *workitemIds {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    locationId, _ := uuid.Parse("186abea3-5c35-432f-9e28-7a15b4312a0e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ParentChildWIMap
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get available board rows in a project
// ctx
// project (optional): Project ID or project name
func (client Client) GetRowSuggestedValues(ctx context.Context, project *string) (*[]BoardSuggestedValue, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    locationId, _ := uuid.Parse("bb494cc6-a0f5-4c6c-8dca-ea6912e79eb9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardSuggestedValue
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get board
// ctx
// project (required): Project ID or project name
// id (required): identifier for board, either board's backlog level name (Eg:"Stories") or Id
// team (optional): Team ID or team name
func (client Client) GetBoard(ctx context.Context, project *string, id *string, team *string) (*Board, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if id == nil || *id == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    locationId, _ := uuid.Parse("23ad19fc-3b8e-4877-8462-b3f92bc06b40")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Board
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get boards
// ctx
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) GetBoards(ctx context.Context, project *string, team *string) (*[]BoardReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    locationId, _ := uuid.Parse("23ad19fc-3b8e-4877-8462-b3f92bc06b40")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update board options
// ctx
// options (required): options to updated
// project (required): Project ID or project name
// id (required): identifier for board, either category plural name (Eg:"Stories") or guid
// team (optional): Team ID or team name
func (client Client) SetBoardOptions(ctx context.Context, options *map[string]string, project *string, id *string, team *string) (*map[string]string, error) {
    if options == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "options"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if id == nil || *id == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    body, marshalErr := json.Marshal(*options)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("23ad19fc-3b8e-4877-8462-b3f92bc06b40")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get board user settings for a board id
// ctx
// project (required): Project ID or project name
// board (required): Board ID or Name
// team (optional): Team ID or team name
func (client Client) GetBoardUserSettings(ctx context.Context, project *string, board *string, team *string) (*BoardUserSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    locationId, _ := uuid.Parse("b30d9f58-1891-4b0a-b168-c46408f919b0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardUserSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update board user settings for the board id
// ctx
// boardUserSettings (required)
// project (required): Project ID or project name
// board (required)
// team (optional): Team ID or team name
func (client Client) UpdateBoardUserSettings(ctx context.Context, boardUserSettings *map[string]string, project *string, board *string, team *string) (*BoardUserSettings, error) {
    if boardUserSettings == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "boardUserSettings"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    body, marshalErr := json.Marshal(*boardUserSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b30d9f58-1891-4b0a-b168-c46408f919b0")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardUserSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a team's capacity
// ctx
// project (required): Project ID or project name
// iterationId (required): ID of the iteration
// team (optional): Team ID or team name
func (client Client) GetCapacitiesWithIdentityRef(ctx context.Context, project *string, iterationId *uuid.UUID, team *string) (*[]TeamMemberCapacityIdentityRef, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()

    locationId, _ := uuid.Parse("74412d15-8c1a-4352-a48d-ef1ed5587d57")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TeamMemberCapacityIdentityRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a team member's capacity
// ctx
// project (required): Project ID or project name
// iterationId (required): ID of the iteration
// teamMemberId (required): ID of the team member
// team (optional): Team ID or team name
func (client Client) GetCapacityWithIdentityRef(ctx context.Context, project *string, iterationId *uuid.UUID, teamMemberId *uuid.UUID, team *string) (*TeamMemberCapacityIdentityRef, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()
    if teamMemberId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "teamMemberId"} 
    }
    routeValues["teamMemberId"] = (*teamMemberId).String()

    locationId, _ := uuid.Parse("74412d15-8c1a-4352-a48d-ef1ed5587d57")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamMemberCapacityIdentityRef
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Replace a team's capacity
// ctx
// capacities (required): Team capacity to replace
// project (required): Project ID or project name
// iterationId (required): ID of the iteration
// team (optional): Team ID or team name
func (client Client) ReplaceCapacitiesWithIdentityRef(ctx context.Context, capacities *[]TeamMemberCapacityIdentityRef, project *string, iterationId *uuid.UUID, team *string) (*[]TeamMemberCapacityIdentityRef, error) {
    if capacities == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "capacities"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()

    body, marshalErr := json.Marshal(*capacities)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("74412d15-8c1a-4352-a48d-ef1ed5587d57")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TeamMemberCapacityIdentityRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a team member's capacity
// ctx
// patch (required): Updated capacity
// project (required): Project ID or project name
// iterationId (required): ID of the iteration
// teamMemberId (required): ID of the team member
// team (optional): Team ID or team name
func (client Client) UpdateCapacityWithIdentityRef(ctx context.Context, patch *CapacityPatch, project *string, iterationId *uuid.UUID, teamMemberId *uuid.UUID, team *string) (*TeamMemberCapacityIdentityRef, error) {
    if patch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "patch"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()
    if teamMemberId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "teamMemberId"} 
    }
    routeValues["teamMemberId"] = (*teamMemberId).String()

    body, marshalErr := json.Marshal(*patch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("74412d15-8c1a-4352-a48d-ef1ed5587d57")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamMemberCapacityIdentityRef
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get board card Rule settings for the board id or board by name
// ctx
// project (required): Project ID or project name
// board (required)
// team (optional): Team ID or team name
func (client Client) GetBoardCardRuleSettings(ctx context.Context, project *string, board *string, team *string) (*BoardCardRuleSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    locationId, _ := uuid.Parse("b044a3d9-02ea-49c7-91a1-b730949cc896")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardCardRuleSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update board card Rule settings for the board id or board by name
// ctx
// boardCardRuleSettings (required)
// project (required): Project ID or project name
// board (required)
// team (optional): Team ID or team name
func (client Client) UpdateBoardCardRuleSettings(ctx context.Context, boardCardRuleSettings *BoardCardRuleSettings, project *string, board *string, team *string) (*BoardCardRuleSettings, error) {
    if boardCardRuleSettings == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "boardCardRuleSettings"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    body, marshalErr := json.Marshal(*boardCardRuleSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b044a3d9-02ea-49c7-91a1-b730949cc896")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardCardRuleSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update taskboard card Rule settings
// ctx
// boardCardRuleSettings (required)
// project (required): Project ID or project name
// team (required): Team ID or team name
func (client Client) UpdateTaskboardCardRuleSettings(ctx context.Context, boardCardRuleSettings *BoardCardRuleSettings, project *string, team *string) error {
    if boardCardRuleSettings == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "boardCardRuleSettings"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team

    body, marshalErr := json.Marshal(*boardCardRuleSettings)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("3f84a8d1-1aab-423e-a94b-6dcbdcca511f")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get board card settings for the board id or board by name
// ctx
// project (required): Project ID or project name
// board (required)
// team (optional): Team ID or team name
func (client Client) GetBoardCardSettings(ctx context.Context, project *string, board *string, team *string) (*BoardCardSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    locationId, _ := uuid.Parse("07c3b467-bc60-4f05-8e34-599ce288fafc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardCardSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update board card settings for the board id or board by name
// ctx
// boardCardSettingsToSave (required)
// project (required): Project ID or project name
// board (required)
// team (optional): Team ID or team name
func (client Client) UpdateBoardCardSettings(ctx context.Context, boardCardSettingsToSave *BoardCardSettings, project *string, board *string, team *string) (*BoardCardSettings, error) {
    if boardCardSettingsToSave == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "boardCardSettingsToSave"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    body, marshalErr := json.Marshal(*boardCardSettingsToSave)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("07c3b467-bc60-4f05-8e34-599ce288fafc")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardCardSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update taskboard card settings
// ctx
// boardCardSettingsToSave (required)
// project (required): Project ID or project name
// team (required): Team ID or team name
func (client Client) UpdateTaskboardCardSettings(ctx context.Context, boardCardSettingsToSave *BoardCardSettings, project *string, team *string) error {
    if boardCardSettingsToSave == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "boardCardSettingsToSave"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team

    body, marshalErr := json.Marshal(*boardCardSettingsToSave)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("0d63745f-31f3-4cf3-9056-2a064e567637")
    _, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a board chart
// ctx
// project (required): Project ID or project name
// board (required): Identifier for board, either board's backlog level name (Eg:"Stories") or Id
// name (required): The chart name
// team (optional): Team ID or team name
func (client Client) GetBoardChart(ctx context.Context, project *string, board *string, name *string, team *string) (*BoardChart, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board
    if name == nil || *name == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("45fe888c-239e-49fd-958c-df1a1ab21d97")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardChart
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get board charts
// ctx
// project (required): Project ID or project name
// board (required): Identifier for board, either board's backlog level name (Eg:"Stories") or Id
// team (optional): Team ID or team name
func (client Client) GetBoardCharts(ctx context.Context, project *string, board *string, team *string) (*[]BoardChartReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    locationId, _ := uuid.Parse("45fe888c-239e-49fd-958c-df1a1ab21d97")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardChartReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a board chart
// ctx
// chart (required)
// project (required): Project ID or project name
// board (required): Identifier for board, either board's backlog level name (Eg:"Stories") or Id
// name (required): The chart name
// team (optional): Team ID or team name
func (client Client) UpdateBoardChart(ctx context.Context, chart *BoardChart, project *string, board *string, name *string, team *string) (*BoardChart, error) {
    if chart == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "chart"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board
    if name == nil || *name == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *name

    body, marshalErr := json.Marshal(*chart)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("45fe888c-239e-49fd-958c-df1a1ab21d97")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BoardChart
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get columns on a board
// ctx
// project (required): Project ID or project name
// board (required): Name or ID of the specific board
// team (optional): Team ID or team name
func (client Client) GetBoardColumns(ctx context.Context, project *string, board *string, team *string) (*[]BoardColumn, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    locationId, _ := uuid.Parse("c555d7ff-84e1-47df-9923-a3fe0cd8751b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardColumn
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update columns on a board
// ctx
// boardColumns (required): List of board columns to update
// project (required): Project ID or project name
// board (required): Name or ID of the specific board
// team (optional): Team ID or team name
func (client Client) UpdateBoardColumns(ctx context.Context, boardColumns *[]BoardColumn, project *string, board *string, team *string) (*[]BoardColumn, error) {
    if boardColumns == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "boardColumns"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    body, marshalErr := json.Marshal(*boardColumns)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("c555d7ff-84e1-47df-9923-a3fe0cd8751b")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardColumn
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get Delivery View Data
// ctx
// project (required): Project ID or project name
// id (required): Identifier for delivery view
// revision (optional): Revision of the plan for which you want data. If the current plan is a different revision you will get an ViewRevisionMismatchException exception. If you do not supply a revision you will get data for the latest revision.
// startDate (optional): The start date of timeline
// endDate (optional): The end date of timeline
func (client Client) GetDeliveryTimelineData(ctx context.Context, project *string, id *string, revision *int, startDate *time.Time, endDate *time.Time) (*DeliveryViewData, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if id == nil || *id == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    queryParams := url.Values{}
    if revision != nil {
        queryParams.Add("revision", strconv.Itoa(*revision))
    }
    if startDate != nil {
        queryParams.Add("startDate", (*startDate).String())
    }
    if endDate != nil {
        queryParams.Add("endDate", (*endDate).String())
    }
    locationId, _ := uuid.Parse("bdd0834e-101f-49f0-a6ae-509f384a12b4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DeliveryViewData
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a team's iteration by iterationId
// ctx
// project (required): Project ID or project name
// id (required): ID of the iteration
// team (optional): Team ID or team name
func (client Client) DeleteTeamIteration(ctx context.Context, project *string, id *uuid.UUID, team *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if id == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = (*id).String()

    locationId, _ := uuid.Parse("c9175577-28a1-4b06-9197-8636af9f64ad")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get team's iteration by iterationId
// ctx
// project (required): Project ID or project name
// id (required): ID of the iteration
// team (optional): Team ID or team name
func (client Client) GetTeamIteration(ctx context.Context, project *string, id *uuid.UUID, team *string) (*TeamSettingsIteration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = (*id).String()

    locationId, _ := uuid.Parse("c9175577-28a1-4b06-9197-8636af9f64ad")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamSettingsIteration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a team's iterations using timeframe filter
// ctx
// project (required): Project ID or project name
// team (optional): Team ID or team name
// timeframe (optional): A filter for which iterations are returned based on relative time. Only Current is supported currently.
func (client Client) GetTeamIterations(ctx context.Context, project *string, team *string, timeframe *string) (*[]TeamSettingsIteration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    queryParams := url.Values{}
    if timeframe != nil {
        queryParams.Add("$timeframe", *timeframe)
    }
    locationId, _ := uuid.Parse("c9175577-28a1-4b06-9197-8636af9f64ad")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TeamSettingsIteration
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Add an iteration to the team
// ctx
// iteration (required): Iteration to add
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) PostTeamIteration(ctx context.Context, iteration *TeamSettingsIteration, project *string, team *string) (*TeamSettingsIteration, error) {
    if iteration == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iteration"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    body, marshalErr := json.Marshal(*iteration)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("c9175577-28a1-4b06-9197-8636af9f64ad")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamSettingsIteration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Add a new plan for the team
// ctx
// postedPlan (required): Plan definition
// project (required): Project ID or project name
func (client Client) CreatePlan(ctx context.Context, postedPlan *CreatePlan, project *string) (*Plan, error) {
    if postedPlan == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "postedPlan"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*postedPlan)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Plan
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete the specified plan
// ctx
// project (required): Project ID or project name
// id (required): Identifier of the plan
func (client Client) DeletePlan(ctx context.Context, project *string, id *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if id == nil || *id == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get the information for the specified plan
// ctx
// project (required): Project ID or project name
// id (required): Identifier of the plan
func (client Client) GetPlan(ctx context.Context, project *string, id *string) (*Plan, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if id == nil || *id == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Plan
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get the information for all the plans configured for the given team
// ctx
// project (required): Project ID or project name
func (client Client) GetPlans(ctx context.Context, project *string) (*[]Plan, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Plan
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update the information for the specified plan
// ctx
// updatedPlan (required): Plan definition to be updated
// project (required): Project ID or project name
// id (required): Identifier of the plan
func (client Client) UpdatePlan(ctx context.Context, updatedPlan *UpdatePlan, project *string, id *string) (*Plan, error) {
    if updatedPlan == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updatedPlan"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if id == nil || *id == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "id"} 
    }
    routeValues["id"] = *id

    body, marshalErr := json.Marshal(*updatedPlan)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Plan
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get process configuration
// ctx
// project (required): Project ID or project name
func (client Client) GetProcessConfiguration(ctx context.Context, project *string) (*ProcessConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("f901ba42-86d2-4b0c-89c1-3f86d06daa84")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get rows on a board
// ctx
// project (required): Project ID or project name
// board (required): Name or ID of the specific board
// team (optional): Team ID or team name
func (client Client) GetBoardRows(ctx context.Context, project *string, board *string, team *string) (*[]BoardRow, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    locationId, _ := uuid.Parse("0863355d-aefd-4d63-8669-984c9b7b0e78")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardRow
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update rows on a board
// ctx
// boardRows (required): List of board rows to update
// project (required): Project ID or project name
// board (required): Name or ID of the specific board
// team (optional): Team ID or team name
func (client Client) UpdateBoardRows(ctx context.Context, boardRows *[]BoardRow, project *string, board *string, team *string) (*[]BoardRow, error) {
    if boardRows == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "boardRows"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if board == nil || *board == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "board"} 
    }
    routeValues["board"] = *board

    body, marshalErr := json.Marshal(*boardRows)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0863355d-aefd-4d63-8669-984c9b7b0e78")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BoardRow
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get team's days off for an iteration
// ctx
// project (required): Project ID or project name
// iterationId (required): ID of the iteration
// team (optional): Team ID or team name
func (client Client) GetTeamDaysOff(ctx context.Context, project *string, iterationId *uuid.UUID, team *string) (*TeamSettingsDaysOff, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()

    locationId, _ := uuid.Parse("2d4faa2e-9150-4cbf-a47a-932b1b4a0773")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamSettingsDaysOff
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Set a team's days off for an iteration
// ctx
// daysOffPatch (required): Team's days off patch containting a list of start and end dates
// project (required): Project ID or project name
// iterationId (required): ID of the iteration
// team (optional): Team ID or team name
func (client Client) UpdateTeamDaysOff(ctx context.Context, daysOffPatch *TeamSettingsDaysOffPatch, project *string, iterationId *uuid.UUID, team *string) (*TeamSettingsDaysOff, error) {
    if daysOffPatch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "daysOffPatch"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()

    body, marshalErr := json.Marshal(*daysOffPatch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2d4faa2e-9150-4cbf-a47a-932b1b4a0773")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamSettingsDaysOff
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a collection of team field values
// ctx
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) GetTeamFieldValues(ctx context.Context, project *string, team *string) (*TeamFieldValues, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    locationId, _ := uuid.Parse("07ced576-58ed-49e6-9c1e-5cb53ab8bf2a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamFieldValues
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update team field values
// ctx
// patch (required)
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) UpdateTeamFieldValues(ctx context.Context, patch *TeamFieldValuesPatch, project *string, team *string) (*TeamFieldValues, error) {
    if patch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "patch"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    body, marshalErr := json.Marshal(*patch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("07ced576-58ed-49e6-9c1e-5cb53ab8bf2a")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamFieldValues
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a team's settings
// ctx
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) GetTeamSettings(ctx context.Context, project *string, team *string) (*TeamSetting, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    locationId, _ := uuid.Parse("c3c1012b-bea7-49d7-b45e-1664e566f84c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamSetting
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update a team's settings
// ctx
// teamSettingsPatch (required): TeamSettings changes
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) UpdateTeamSettings(ctx context.Context, teamSettingsPatch *TeamSettingsPatch, project *string, team *string) (*TeamSetting, error) {
    if teamSettingsPatch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "teamSettingsPatch"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    body, marshalErr := json.Marshal(*teamSettingsPatch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("c3c1012b-bea7-49d7-b45e-1664e566f84c")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TeamSetting
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get work items for iteration
// ctx
// project (required): Project ID or project name
// iterationId (required): ID of the iteration
// team (optional): Team ID or team name
func (client Client) GetIterationWorkItems(ctx context.Context, project *string, iterationId *uuid.UUID, team *string) (*IterationWorkItems, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()

    locationId, _ := uuid.Parse("5b3ef1a6-d3ab-44cd-bafd-c7f45db850fa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IterationWorkItems
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Reorder Product Backlog/Boards Work Items
// ctx
// operation (required)
// project (required): Project ID or project name
// team (required): Team ID or team name
func (client Client) ReorderBacklogWorkItems(ctx context.Context, operation *ReorderOperation, project *string, team *string) (*[]ReorderResult, error) {
    if operation == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "operation"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team

    body, marshalErr := json.Marshal(*operation)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1c22b714-e7e4-41b9-85e0-56ee13ef55ed")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReorderResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Reorder Sprint Backlog/Taskboard Work Items
// ctx
// operation (required)
// project (required): Project ID or project name
// team (required): Team ID or team name
// iterationId (required): The id of the iteration
func (client Client) ReorderIterationWorkItems(ctx context.Context, operation *ReorderOperation, project *string, team *string, iterationId *uuid.UUID) (*[]ReorderResult, error) {
    if operation == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "operation"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = (*iterationId).String()

    body, marshalErr := json.Marshal(*operation)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("47755db2-d7eb-405a-8c25-675401525fc9")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReorderResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

