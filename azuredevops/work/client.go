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
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var ResourceAreaId, _ = uuid.Parse("1d4f49f9-02b9-4e26-b826-2cdb6195f2a9")

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

// Gets backlog configuration for a team
func (client *Client) GetBacklogConfigurations(ctx context.Context, args GetBacklogConfigurationsArgs) (*BacklogConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
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

// Arguments for the GetBacklogConfigurations function
type GetBacklogConfigurationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
}

// [Preview API] Get a list of work items within a backlog level
func (client *Client) GetBacklogLevelWorkItems(ctx context.Context, args GetBacklogLevelWorkItemsArgs) (*BacklogLevelWorkItems, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team == nil || *args.Team == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "team"}
	}
	routeValues["team"] = *args.Team
	if args.BacklogId == nil || *args.BacklogId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "backlogId"}
	}
	routeValues["backlogId"] = *args.BacklogId

	locationId, _ := uuid.Parse("7c468d96-ab1d-4294-a360-92f07e9ccd98")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BacklogLevelWorkItems
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBacklogLevelWorkItems function
type GetBacklogLevelWorkItemsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Team ID or team name
	Team *string
	// (required)
	BacklogId *string
}

// [Preview API] Get a backlog level
func (client *Client) GetBacklog(ctx context.Context, args GetBacklogArgs) (*BacklogLevelConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team == nil || *args.Team == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "team"}
	}
	routeValues["team"] = *args.Team
	if args.Id == nil || *args.Id == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "id"}
	}
	routeValues["id"] = *args.Id

	locationId, _ := uuid.Parse("a93726f9-7867-4e38-b4f2-0bfafc2f6a94")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BacklogLevelConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBacklog function
type GetBacklogArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Team ID or team name
	Team *string
	// (required) The id of the backlog level
	Id *string
}

// [Preview API] List all backlog levels
func (client *Client) GetBacklogs(ctx context.Context, args GetBacklogsArgs) (*[]BacklogLevelConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team == nil || *args.Team == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "team"}
	}
	routeValues["team"] = *args.Team

	locationId, _ := uuid.Parse("a93726f9-7867-4e38-b4f2-0bfafc2f6a94")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BacklogLevelConfiguration
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBacklogs function
type GetBacklogsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Team ID or team name
	Team *string
}

// Get available board columns in a project
func (client *Client) GetColumnSuggestedValues(ctx context.Context, args GetColumnSuggestedValuesArgs) (*[]BoardSuggestedValue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
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

// Arguments for the GetColumnSuggestedValues function
type GetColumnSuggestedValuesArgs struct {
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Returns the list of parent field filter model for the given list of workitem ids
func (client *Client) GetBoardMappingParentItems(ctx context.Context, args GetBoardMappingParentItemsArgs) (*[]ParentChildWIMap, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}

	queryParams := url.Values{}
	if args.ChildBacklogContextCategoryRefName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "childBacklogContextCategoryRefName"}
	}
	queryParams.Add("childBacklogContextCategoryRefName", *args.ChildBacklogContextCategoryRefName)
	if args.WorkitemIds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "workitemIds"}
	}
	var stringList []string
	for _, item := range *args.WorkitemIds {
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

// Arguments for the GetBoardMappingParentItems function
type GetBoardMappingParentItemsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	ChildBacklogContextCategoryRefName *string
	// (required)
	WorkitemIds *[]int
	// (optional) Team ID or team name
	Team *string
}

// Get available board rows in a project
func (client *Client) GetRowSuggestedValues(ctx context.Context, args GetRowSuggestedValuesArgs) (*[]BoardSuggestedValue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
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

// Arguments for the GetRowSuggestedValues function
type GetRowSuggestedValuesArgs struct {
	// (optional) Project ID or project name
	Project *string
}

// Get board
func (client *Client) GetBoard(ctx context.Context, args GetBoardArgs) (*Board, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Id == nil || *args.Id == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "id"}
	}
	routeValues["id"] = *args.Id

	locationId, _ := uuid.Parse("23ad19fc-3b8e-4877-8462-b3f92bc06b40")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Board
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoard function
type GetBoardArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) identifier for board, either board's backlog level name (Eg:"Stories") or Id
	Id *string
	// (optional) Team ID or team name
	Team *string
}

// Get boards
func (client *Client) GetBoards(ctx context.Context, args GetBoardsArgs) (*[]BoardReference, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
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

// Arguments for the GetBoards function
type GetBoardsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
}

// Update board options
func (client *Client) SetBoardOptions(ctx context.Context, args SetBoardOptionsArgs) (*map[string]string, error) {
	if args.Options == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "options"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Id == nil || *args.Id == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "id"}
	}
	routeValues["id"] = *args.Id

	body, marshalErr := json.Marshal(*args.Options)
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

// Arguments for the SetBoardOptions function
type SetBoardOptionsArgs struct {
	// (required) options to updated
	Options *map[string]string
	// (required) Project ID or project name
	Project *string
	// (required) identifier for board, either category plural name (Eg:"Stories") or guid
	Id *string
	// (optional) Team ID or team name
	Team *string
}

// [Preview API] Get board user settings for a board id
func (client *Client) GetBoardUserSettings(ctx context.Context, args GetBoardUserSettingsArgs) (*BoardUserSettings, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	locationId, _ := uuid.Parse("b30d9f58-1891-4b0a-b168-c46408f919b0")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BoardUserSettings
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoardUserSettings function
type GetBoardUserSettingsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Board ID or Name
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// [Preview API] Update board user settings for the board id
func (client *Client) UpdateBoardUserSettings(ctx context.Context, args UpdateBoardUserSettingsArgs) (*BoardUserSettings, error) {
	if args.BoardUserSettings == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "boardUserSettings"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	body, marshalErr := json.Marshal(*args.BoardUserSettings)
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

// Arguments for the UpdateBoardUserSettings function
type UpdateBoardUserSettingsArgs struct {
	// (required)
	BoardUserSettings *map[string]string
	// (required) Project ID or project name
	Project *string
	// (required)
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Get a team's capacity
func (client *Client) GetCapacitiesWithIdentityRef(ctx context.Context, args GetCapacitiesWithIdentityRefArgs) (*[]TeamMemberCapacityIdentityRef, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()

	locationId, _ := uuid.Parse("74412d15-8c1a-4352-a48d-ef1ed5587d57")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TeamMemberCapacityIdentityRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCapacitiesWithIdentityRef function
type GetCapacitiesWithIdentityRefArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	IterationId *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Get a team member's capacity
func (client *Client) GetCapacityWithIdentityRef(ctx context.Context, args GetCapacityWithIdentityRefArgs) (*TeamMemberCapacityIdentityRef, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()
	if args.TeamMemberId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "teamMemberId"}
	}
	routeValues["teamMemberId"] = (*args.TeamMemberId).String()

	locationId, _ := uuid.Parse("74412d15-8c1a-4352-a48d-ef1ed5587d57")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TeamMemberCapacityIdentityRef
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCapacityWithIdentityRef function
type GetCapacityWithIdentityRefArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	IterationId *uuid.UUID
	// (required) ID of the team member
	TeamMemberId *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Replace a team's capacity
func (client *Client) ReplaceCapacitiesWithIdentityRef(ctx context.Context, args ReplaceCapacitiesWithIdentityRefArgs) (*[]TeamMemberCapacityIdentityRef, error) {
	if args.Capacities == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "capacities"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()

	body, marshalErr := json.Marshal(*args.Capacities)
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

// Arguments for the ReplaceCapacitiesWithIdentityRef function
type ReplaceCapacitiesWithIdentityRefArgs struct {
	// (required) Team capacity to replace
	Capacities *[]TeamMemberCapacityIdentityRef
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	IterationId *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Update a team member's capacity
func (client *Client) UpdateCapacityWithIdentityRef(ctx context.Context, args UpdateCapacityWithIdentityRefArgs) (*TeamMemberCapacityIdentityRef, error) {
	if args.Patch == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "patch"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()
	if args.TeamMemberId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "teamMemberId"}
	}
	routeValues["teamMemberId"] = (*args.TeamMemberId).String()

	body, marshalErr := json.Marshal(*args.Patch)
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

// Arguments for the UpdateCapacityWithIdentityRef function
type UpdateCapacityWithIdentityRefArgs struct {
	// (required) Updated capacity
	Patch *CapacityPatch
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	IterationId *uuid.UUID
	// (required) ID of the team member
	TeamMemberId *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Get board card Rule settings for the board id or board by name
func (client *Client) GetBoardCardRuleSettings(ctx context.Context, args GetBoardCardRuleSettingsArgs) (*BoardCardRuleSettings, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	locationId, _ := uuid.Parse("b044a3d9-02ea-49c7-91a1-b730949cc896")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BoardCardRuleSettings
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoardCardRuleSettings function
type GetBoardCardRuleSettingsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Update board card Rule settings for the board id or board by name
func (client *Client) UpdateBoardCardRuleSettings(ctx context.Context, args UpdateBoardCardRuleSettingsArgs) (*BoardCardRuleSettings, error) {
	if args.BoardCardRuleSettings == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "boardCardRuleSettings"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	body, marshalErr := json.Marshal(*args.BoardCardRuleSettings)
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

// Arguments for the UpdateBoardCardRuleSettings function
type UpdateBoardCardRuleSettingsArgs struct {
	// (required)
	BoardCardRuleSettings *BoardCardRuleSettings
	// (required) Project ID or project name
	Project *string
	// (required)
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// [Preview API] Update taskboard card Rule settings
func (client *Client) UpdateTaskboardCardRuleSettings(ctx context.Context, args UpdateTaskboardCardRuleSettingsArgs) error {
	if args.BoardCardRuleSettings == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "boardCardRuleSettings"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team == nil || *args.Team == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "team"}
	}
	routeValues["team"] = *args.Team

	body, marshalErr := json.Marshal(*args.BoardCardRuleSettings)
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

// Arguments for the UpdateTaskboardCardRuleSettings function
type UpdateTaskboardCardRuleSettingsArgs struct {
	// (required)
	BoardCardRuleSettings *BoardCardRuleSettings
	// (required) Project ID or project name
	Project *string
	// (required) Team ID or team name
	Team *string
}

// Get board card settings for the board id or board by name
func (client *Client) GetBoardCardSettings(ctx context.Context, args GetBoardCardSettingsArgs) (*BoardCardSettings, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	locationId, _ := uuid.Parse("07c3b467-bc60-4f05-8e34-599ce288fafc")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BoardCardSettings
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoardCardSettings function
type GetBoardCardSettingsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Update board card settings for the board id or board by name
func (client *Client) UpdateBoardCardSettings(ctx context.Context, args UpdateBoardCardSettingsArgs) (*BoardCardSettings, error) {
	if args.BoardCardSettingsToSave == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "boardCardSettingsToSave"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	body, marshalErr := json.Marshal(*args.BoardCardSettingsToSave)
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

// Arguments for the UpdateBoardCardSettings function
type UpdateBoardCardSettingsArgs struct {
	// (required)
	BoardCardSettingsToSave *BoardCardSettings
	// (required) Project ID or project name
	Project *string
	// (required)
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// [Preview API] Update taskboard card settings
func (client *Client) UpdateTaskboardCardSettings(ctx context.Context, args UpdateTaskboardCardSettingsArgs) error {
	if args.BoardCardSettingsToSave == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "boardCardSettingsToSave"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team == nil || *args.Team == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "team"}
	}
	routeValues["team"] = *args.Team

	body, marshalErr := json.Marshal(*args.BoardCardSettingsToSave)
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

// Arguments for the UpdateTaskboardCardSettings function
type UpdateTaskboardCardSettingsArgs struct {
	// (required)
	BoardCardSettingsToSave *BoardCardSettings
	// (required) Project ID or project name
	Project *string
	// (required) Team ID or team name
	Team *string
}

// Get a board chart
func (client *Client) GetBoardChart(ctx context.Context, args GetBoardChartArgs) (*BoardChart, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "name"}
	}
	routeValues["name"] = *args.Name

	locationId, _ := uuid.Parse("45fe888c-239e-49fd-958c-df1a1ab21d97")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue BoardChart
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoardChart function
type GetBoardChartArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Identifier for board, either board's backlog level name (Eg:"Stories") or Id
	Board *string
	// (required) The chart name
	Name *string
	// (optional) Team ID or team name
	Team *string
}

// Get board charts
func (client *Client) GetBoardCharts(ctx context.Context, args GetBoardChartsArgs) (*[]BoardChartReference, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	locationId, _ := uuid.Parse("45fe888c-239e-49fd-958c-df1a1ab21d97")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BoardChartReference
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoardCharts function
type GetBoardChartsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Identifier for board, either board's backlog level name (Eg:"Stories") or Id
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Update a board chart
func (client *Client) UpdateBoardChart(ctx context.Context, args UpdateBoardChartArgs) (*BoardChart, error) {
	if args.Chart == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "chart"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "name"}
	}
	routeValues["name"] = *args.Name

	body, marshalErr := json.Marshal(*args.Chart)
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

// Arguments for the UpdateBoardChart function
type UpdateBoardChartArgs struct {
	// (required)
	Chart *BoardChart
	// (required) Project ID or project name
	Project *string
	// (required) Identifier for board, either board's backlog level name (Eg:"Stories") or Id
	Board *string
	// (required) The chart name
	Name *string
	// (optional) Team ID or team name
	Team *string
}

// Get columns on a board
func (client *Client) GetBoardColumns(ctx context.Context, args GetBoardColumnsArgs) (*[]BoardColumn, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	locationId, _ := uuid.Parse("c555d7ff-84e1-47df-9923-a3fe0cd8751b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BoardColumn
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoardColumns function
type GetBoardColumnsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Name or ID of the specific board
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Update columns on a board
func (client *Client) UpdateBoardColumns(ctx context.Context, args UpdateBoardColumnsArgs) (*[]BoardColumn, error) {
	if args.BoardColumns == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "boardColumns"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	body, marshalErr := json.Marshal(*args.BoardColumns)
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

// Arguments for the UpdateBoardColumns function
type UpdateBoardColumnsArgs struct {
	// (required) List of board columns to update
	BoardColumns *[]BoardColumn
	// (required) Project ID or project name
	Project *string
	// (required) Name or ID of the specific board
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Get Delivery View Data
func (client *Client) GetDeliveryTimelineData(ctx context.Context, args GetDeliveryTimelineDataArgs) (*DeliveryViewData, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Id == nil || *args.Id == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "id"}
	}
	routeValues["id"] = *args.Id

	queryParams := url.Values{}
	if args.Revision != nil {
		queryParams.Add("revision", strconv.Itoa(*args.Revision))
	}
	if args.StartDate != nil {
		queryParams.Add("startDate", (*args.StartDate).String())
	}
	if args.EndDate != nil {
		queryParams.Add("endDate", (*args.EndDate).String())
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

// Arguments for the GetDeliveryTimelineData function
type GetDeliveryTimelineDataArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Identifier for delivery view
	Id *string
	// (optional) Revision of the plan for which you want data. If the current plan is a different revision you will get an ViewRevisionMismatchException exception. If you do not supply a revision you will get data for the latest revision.
	Revision *int
	// (optional) The start date of timeline
	StartDate *time.Time
	// (optional) The end date of timeline
	EndDate *time.Time
}

// Delete a team's iteration by iterationId
func (client *Client) DeleteTeamIteration(ctx context.Context, args DeleteTeamIterationArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Id == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "id"}
	}
	routeValues["id"] = (*args.Id).String()

	locationId, _ := uuid.Parse("c9175577-28a1-4b06-9197-8636af9f64ad")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTeamIteration function
type DeleteTeamIterationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	Id *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Get team's iteration by iterationId
func (client *Client) GetTeamIteration(ctx context.Context, args GetTeamIterationArgs) (*TeamSettingsIteration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Id == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "id"}
	}
	routeValues["id"] = (*args.Id).String()

	locationId, _ := uuid.Parse("c9175577-28a1-4b06-9197-8636af9f64ad")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TeamSettingsIteration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTeamIteration function
type GetTeamIterationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	Id *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Get a team's iterations using timeframe filter
func (client *Client) GetTeamIterations(ctx context.Context, args GetTeamIterationsArgs) (*[]TeamSettingsIteration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}

	queryParams := url.Values{}
	if args.Timeframe != nil {
		queryParams.Add("$timeframe", *args.Timeframe)
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

// Arguments for the GetTeamIterations function
type GetTeamIterationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
	// (optional) A filter for which iterations are returned based on relative time. Only Current is supported currently.
	Timeframe *string
}

// Add an iteration to the team
func (client *Client) PostTeamIteration(ctx context.Context, args PostTeamIterationArgs) (*TeamSettingsIteration, error) {
	if args.Iteration == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iteration"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}

	body, marshalErr := json.Marshal(*args.Iteration)
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

// Arguments for the PostTeamIteration function
type PostTeamIterationArgs struct {
	// (required) Iteration to add
	Iteration *TeamSettingsIteration
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
}

// Add a new plan for the team
func (client *Client) CreatePlan(ctx context.Context, args CreatePlanArgs) (*Plan, error) {
	if args.PostedPlan == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "postedPlan"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.PostedPlan)
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

// Arguments for the CreatePlan function
type CreatePlanArgs struct {
	// (required) Plan definition
	PostedPlan *CreatePlan
	// (required) Project ID or project name
	Project *string
}

// Delete the specified plan
func (client *Client) DeletePlan(ctx context.Context, args DeletePlanArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Id == nil || *args.Id == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "id"}
	}
	routeValues["id"] = *args.Id

	locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePlan function
type DeletePlanArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Identifier of the plan
	Id *string
}

// Get the information for the specified plan
func (client *Client) GetPlan(ctx context.Context, args GetPlanArgs) (*Plan, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Id == nil || *args.Id == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "id"}
	}
	routeValues["id"] = *args.Id

	locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Plan
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPlan function
type GetPlanArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Identifier of the plan
	Id *string
}

// Get the information for all the plans configured for the given team
func (client *Client) GetPlans(ctx context.Context, args GetPlansArgs) (*[]Plan, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("0b42cb47-cd73-4810-ac90-19c9ba147453")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Plan
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPlans function
type GetPlansArgs struct {
	// (required) Project ID or project name
	Project *string
}

// Update the information for the specified plan
func (client *Client) UpdatePlan(ctx context.Context, args UpdatePlanArgs) (*Plan, error) {
	if args.UpdatedPlan == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "updatedPlan"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Id == nil || *args.Id == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "id"}
	}
	routeValues["id"] = *args.Id

	body, marshalErr := json.Marshal(*args.UpdatedPlan)
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

// Arguments for the UpdatePlan function
type UpdatePlanArgs struct {
	// (required) Plan definition to be updated
	UpdatedPlan *UpdatePlan
	// (required) Project ID or project name
	Project *string
	// (required) Identifier of the plan
	Id *string
}

// [Preview API] Get process configuration
func (client *Client) GetProcessConfiguration(ctx context.Context, args GetProcessConfigurationArgs) (*ProcessConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("f901ba42-86d2-4b0c-89c1-3f86d06daa84")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessConfiguration function
type GetProcessConfigurationArgs struct {
	// (required) Project ID or project name
	Project *string
}

// Get rows on a board
func (client *Client) GetBoardRows(ctx context.Context, args GetBoardRowsArgs) (*[]BoardRow, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	locationId, _ := uuid.Parse("0863355d-aefd-4d63-8669-984c9b7b0e78")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []BoardRow
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBoardRows function
type GetBoardRowsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Name or ID of the specific board
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Update rows on a board
func (client *Client) UpdateBoardRows(ctx context.Context, args UpdateBoardRowsArgs) (*[]BoardRow, error) {
	if args.BoardRows == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "boardRows"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.Board == nil || *args.Board == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "board"}
	}
	routeValues["board"] = *args.Board

	body, marshalErr := json.Marshal(*args.BoardRows)
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

// Arguments for the UpdateBoardRows function
type UpdateBoardRowsArgs struct {
	// (required) List of board rows to update
	BoardRows *[]BoardRow
	// (required) Project ID or project name
	Project *string
	// (required) Name or ID of the specific board
	Board *string
	// (optional) Team ID or team name
	Team *string
}

// Get team's days off for an iteration
func (client *Client) GetTeamDaysOff(ctx context.Context, args GetTeamDaysOffArgs) (*TeamSettingsDaysOff, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()

	locationId, _ := uuid.Parse("2d4faa2e-9150-4cbf-a47a-932b1b4a0773")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TeamSettingsDaysOff
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTeamDaysOff function
type GetTeamDaysOffArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	IterationId *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Set a team's days off for an iteration
func (client *Client) UpdateTeamDaysOff(ctx context.Context, args UpdateTeamDaysOffArgs) (*TeamSettingsDaysOff, error) {
	if args.DaysOffPatch == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "daysOffPatch"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()

	body, marshalErr := json.Marshal(*args.DaysOffPatch)
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

// Arguments for the UpdateTeamDaysOff function
type UpdateTeamDaysOffArgs struct {
	// (required) Team's days off patch containting a list of start and end dates
	DaysOffPatch *TeamSettingsDaysOffPatch
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	IterationId *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// Get a collection of team field values
func (client *Client) GetTeamFieldValues(ctx context.Context, args GetTeamFieldValuesArgs) (*TeamFieldValues, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
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

// Arguments for the GetTeamFieldValues function
type GetTeamFieldValuesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
}

// Update team field values
func (client *Client) UpdateTeamFieldValues(ctx context.Context, args UpdateTeamFieldValuesArgs) (*TeamFieldValues, error) {
	if args.Patch == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "patch"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}

	body, marshalErr := json.Marshal(*args.Patch)
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

// Arguments for the UpdateTeamFieldValues function
type UpdateTeamFieldValuesArgs struct {
	// (required)
	Patch *TeamFieldValuesPatch
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
}

// Get a team's settings
func (client *Client) GetTeamSettings(ctx context.Context, args GetTeamSettingsArgs) (*TeamSetting, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
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

// Arguments for the GetTeamSettings function
type GetTeamSettingsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
}

// Update a team's settings
func (client *Client) UpdateTeamSettings(ctx context.Context, args UpdateTeamSettingsArgs) (*TeamSetting, error) {
	if args.TeamSettingsPatch == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "teamSettingsPatch"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}

	body, marshalErr := json.Marshal(*args.TeamSettingsPatch)
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

// Arguments for the UpdateTeamSettings function
type UpdateTeamSettingsArgs struct {
	// (required) TeamSettings changes
	TeamSettingsPatch *TeamSettingsPatch
	// (required) Project ID or project name
	Project *string
	// (optional) Team ID or team name
	Team *string
}

// [Preview API] Get work items for iteration
func (client *Client) GetIterationWorkItems(ctx context.Context, args GetIterationWorkItemsArgs) (*IterationWorkItems, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team != nil && *args.Team != "" {
		routeValues["team"] = *args.Team
	}
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()

	locationId, _ := uuid.Parse("5b3ef1a6-d3ab-44cd-bafd-c7f45db850fa")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue IterationWorkItems
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetIterationWorkItems function
type GetIterationWorkItemsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the iteration
	IterationId *uuid.UUID
	// (optional) Team ID or team name
	Team *string
}

// [Preview API] Reorder Product Backlog/Boards Work Items
func (client *Client) ReorderBacklogWorkItems(ctx context.Context, args ReorderBacklogWorkItemsArgs) (*[]ReorderResult, error) {
	if args.Operation == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "operation"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team == nil || *args.Team == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "team"}
	}
	routeValues["team"] = *args.Team

	body, marshalErr := json.Marshal(*args.Operation)
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

// Arguments for the ReorderBacklogWorkItems function
type ReorderBacklogWorkItemsArgs struct {
	// (required)
	Operation *ReorderOperation
	// (required) Project ID or project name
	Project *string
	// (required) Team ID or team name
	Team *string
}

// [Preview API] Reorder Sprint Backlog/Taskboard Work Items
func (client *Client) ReorderIterationWorkItems(ctx context.Context, args ReorderIterationWorkItemsArgs) (*[]ReorderResult, error) {
	if args.Operation == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "operation"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "project"}
	}
	routeValues["project"] = *args.Project
	if args.Team == nil || *args.Team == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "team"}
	}
	routeValues["team"] = *args.Team
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "iterationId"}
	}
	routeValues["iterationId"] = (*args.IterationId).String()

	body, marshalErr := json.Marshal(*args.Operation)
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

// Arguments for the ReorderIterationWorkItems function
type ReorderIterationWorkItemsArgs struct {
	// (required)
	Operation *ReorderOperation
	// (required) Project ID or project name
	Project *string
	// (required) Team ID or team name
	Team *string
	// (required) The id of the iteration
	IterationId *uuid.UUID
}
