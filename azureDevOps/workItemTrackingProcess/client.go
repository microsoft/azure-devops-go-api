// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package workItemTrackingProcess

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azureDevOps"
	"net/http"
	"net/url"
)

var ResourceAreaId, _ = uuid.Parse("5264459e-e5e0-4bd8-b118-0985e68a4ec5")

type Client struct {
	Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) (*Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: *client,
	}, nil
}

// [Preview API] Creates a single behavior in the given process.
func (client *Client) CreateProcessBehavior(ctx context.Context, args CreateProcessBehaviorArgs) (*ProcessBehavior, error) {
	if args.Behavior == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "behavior"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()

	body, marshalErr := json.Marshal(*args.Behavior)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateProcessBehavior function
type CreateProcessBehaviorArgs struct {
	// (required)
	Behavior *ProcessBehaviorCreateRequest
	// (required) The ID of the process
	ProcessId *uuid.UUID
}

// [Preview API] Removes a behavior in the process.
func (client *Client) DeleteProcessBehavior(ctx context.Context, args DeleteProcessBehaviorArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.BehaviorRefName == nil || *args.BehaviorRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"}
	}
	routeValues["behaviorRefName"] = *args.BehaviorRefName

	locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteProcessBehavior function
type DeleteProcessBehaviorArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the behavior
	BehaviorRefName *string
}

// [Preview API] Returns a behavior of the process.
func (client *Client) GetProcessBehavior(ctx context.Context, args GetProcessBehaviorArgs) (*ProcessBehavior, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.BehaviorRefName == nil || *args.BehaviorRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"}
	}
	routeValues["behaviorRefName"] = *args.BehaviorRefName

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessBehavior function
type GetProcessBehaviorArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the behavior
	BehaviorRefName *string
	// (optional)
	Expand *GetBehaviorsExpand
}

// [Preview API] Returns a list of all behaviors in the process.
func (client *Client) GetProcessBehaviors(ctx context.Context, args GetProcessBehaviorsArgs) (*[]ProcessBehavior, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ProcessBehavior
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessBehaviors function
type GetProcessBehaviorsArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (optional)
	Expand *GetBehaviorsExpand
}

// [Preview API] Replaces a behavior in the process.
func (client *Client) UpdateProcessBehavior(ctx context.Context, args UpdateProcessBehaviorArgs) (*ProcessBehavior, error) {
	if args.BehaviorData == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "behaviorData"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.BehaviorRefName == nil || *args.BehaviorRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"}
	}
	routeValues["behaviorRefName"] = *args.BehaviorRefName

	body, marshalErr := json.Marshal(*args.BehaviorData)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateProcessBehavior function
type UpdateProcessBehaviorArgs struct {
	// (required)
	BehaviorData *ProcessBehaviorUpdateRequest
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the behavior
	BehaviorRefName *string
}

// [Preview API] Creates a control in a group.
func (client *Client) CreateControlInGroup(ctx context.Context, args CreateControlInGroupArgs) (*Control, error) {
	if args.Control == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "control"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId

	body, marshalErr := json.Marshal(*args.Control)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Control
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateControlInGroup function
type CreateControlInGroupArgs struct {
	// (required) The control.
	Control *Control
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the group to add the control to.
	GroupId *string
}

// [Preview API] Moves a control to a specified group.
func (client *Client) MoveControlToGroup(ctx context.Context, args MoveControlToGroupArgs) (*Control, error) {
	if args.Control == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "control"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ControlId == nil || *args.ControlId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "controlId"}
	}
	routeValues["controlId"] = *args.ControlId

	queryParams := url.Values{}
	if args.RemoveFromGroupId != nil {
		queryParams.Add("removeFromGroupId", *args.RemoveFromGroupId)
	}
	body, marshalErr := json.Marshal(*args.Control)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Control
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the MoveControlToGroup function
type MoveControlToGroupArgs struct {
	// (required) The control.
	Control *Control
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the group to move the control to.
	GroupId *string
	// (required) The ID of the control.
	ControlId *string
	// (optional) The group ID to remove the control from.
	RemoveFromGroupId *string
}

// [Preview API] Removes a control from the work item form.
func (client *Client) RemoveControlFromGroup(ctx context.Context, args RemoveControlFromGroupArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.GroupId == nil || *args.GroupId == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ControlId == nil || *args.ControlId == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "controlId"}
	}
	routeValues["controlId"] = *args.ControlId

	locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RemoveControlFromGroup function
type RemoveControlFromGroupArgs struct {
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the group.
	GroupId *string
	// (required) The ID of the control to remove.
	ControlId *string
}

// [Preview API] Updates a control on the work item form.
func (client *Client) UpdateControl(ctx context.Context, args UpdateControlArgs) (*Control, error) {
	if args.Control == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "control"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ControlId == nil || *args.ControlId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "controlId"}
	}
	routeValues["controlId"] = *args.ControlId

	body, marshalErr := json.Marshal(*args.Control)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Control
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateControl function
type UpdateControlArgs struct {
	// (required) The updated control.
	Control *Control
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the group.
	GroupId *string
	// (required) The ID of the control.
	ControlId *string
}

// [Preview API] Adds a field to a work item type.
func (client *Client) AddFieldToWorkItemType(ctx context.Context, args AddFieldToWorkItemTypeArgs) (*ProcessWorkItemTypeField, error) {
	if args.Field == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "field"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	body, marshalErr := json.Marshal(*args.Field)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessWorkItemTypeField
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddFieldToWorkItemType function
type AddFieldToWorkItemTypeArgs struct {
	// (required)
	Field *AddProcessWorkItemTypeFieldRequest
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
}

// [Preview API] Returns a list of all fields in a work item type.
func (client *Client) GetAllWorkItemTypeFields(ctx context.Context, args GetAllWorkItemTypeFieldsArgs) (*[]ProcessWorkItemTypeField, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ProcessWorkItemTypeField
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAllWorkItemTypeFields function
type GetAllWorkItemTypeFieldsArgs struct {
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
}

// [Preview API] Returns a field in a work item type.
func (client *Client) GetWorkItemTypeField(ctx context.Context, args GetWorkItemTypeFieldArgs) (*ProcessWorkItemTypeField, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.FieldRefName == nil || *args.FieldRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "fieldRefName"}
	}
	routeValues["fieldRefName"] = *args.FieldRefName

	locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessWorkItemTypeField
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetWorkItemTypeField function
type GetWorkItemTypeFieldArgs struct {
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The reference name of the field.
	FieldRefName *string
}

// [Preview API] Removes a field from a work item type. Does not permanently delete the field.
func (client *Client) RemoveWorkItemTypeField(ctx context.Context, args RemoveWorkItemTypeFieldArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.FieldRefName == nil || *args.FieldRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "fieldRefName"}
	}
	routeValues["fieldRefName"] = *args.FieldRefName

	locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RemoveWorkItemTypeField function
type RemoveWorkItemTypeFieldArgs struct {
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The reference name of the field.
	FieldRefName *string
}

// [Preview API] Updates a field in a work item type.
func (client *Client) UpdateWorkItemTypeField(ctx context.Context, args UpdateWorkItemTypeFieldArgs) (*ProcessWorkItemTypeField, error) {
	if args.Field == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "field"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.FieldRefName == nil || *args.FieldRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "fieldRefName"}
	}
	routeValues["fieldRefName"] = *args.FieldRefName

	body, marshalErr := json.Marshal(*args.Field)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessWorkItemTypeField
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateWorkItemTypeField function
type UpdateWorkItemTypeFieldArgs struct {
	// (required)
	Field *UpdateProcessWorkItemTypeFieldRequest
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The reference name of the field.
	FieldRefName *string
}

// [Preview API] Adds a group to the work item form.
func (client *Client) AddGroup(ctx context.Context, args AddGroupArgs) (*Group, error) {
	if args.Group == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "group"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.PageId == nil || *args.PageId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "pageId"}
	}
	routeValues["pageId"] = *args.PageId
	if args.SectionId == nil || *args.SectionId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "sectionId"}
	}
	routeValues["sectionId"] = *args.SectionId

	body, marshalErr := json.Marshal(*args.Group)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Group
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddGroup function
type AddGroupArgs struct {
	// (required) The group.
	Group *Group
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the page to add the group to.
	PageId *string
	// (required) The ID of the section to add the group to.
	SectionId *string
}

// [Preview API] Moves a group to a different page and section.
func (client *Client) MoveGroupToPage(ctx context.Context, args MoveGroupToPageArgs) (*Group, error) {
	if args.Group == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "group"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.PageId == nil || *args.PageId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "pageId"}
	}
	routeValues["pageId"] = *args.PageId
	if args.SectionId == nil || *args.SectionId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "sectionId"}
	}
	routeValues["sectionId"] = *args.SectionId
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId

	queryParams := url.Values{}
	if args.RemoveFromPageId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "removeFromPageId"}
	}
	queryParams.Add("removeFromPageId", *args.RemoveFromPageId)
	if args.RemoveFromSectionId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "removeFromSectionId"}
	}
	queryParams.Add("removeFromSectionId", *args.RemoveFromSectionId)
	body, marshalErr := json.Marshal(*args.Group)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Group
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the MoveGroupToPage function
type MoveGroupToPageArgs struct {
	// (required) The updated group.
	Group *Group
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the page the group is in.
	PageId *string
	// (required) The ID of the section the group is i.n
	SectionId *string
	// (required) The ID of the group.
	GroupId *string
	// (required) ID of the page to remove the group from.
	RemoveFromPageId *string
	// (required) ID of the section to remove the group from.
	RemoveFromSectionId *string
}

// [Preview API] Moves a group to a different section.
func (client *Client) MoveGroupToSection(ctx context.Context, args MoveGroupToSectionArgs) (*Group, error) {
	if args.Group == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "group"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.PageId == nil || *args.PageId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "pageId"}
	}
	routeValues["pageId"] = *args.PageId
	if args.SectionId == nil || *args.SectionId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "sectionId"}
	}
	routeValues["sectionId"] = *args.SectionId
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId

	queryParams := url.Values{}
	if args.RemoveFromSectionId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "removeFromSectionId"}
	}
	queryParams.Add("removeFromSectionId", *args.RemoveFromSectionId)
	body, marshalErr := json.Marshal(*args.Group)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Group
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the MoveGroupToSection function
type MoveGroupToSectionArgs struct {
	// (required) The updated group.
	Group *Group
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the page the group is in.
	PageId *string
	// (required) The ID of the section the group is in.
	SectionId *string
	// (required) The ID of the group.
	GroupId *string
	// (required) ID of the section to remove the group from.
	RemoveFromSectionId *string
}

// [Preview API] Removes a group from the work item form.
func (client *Client) RemoveGroup(ctx context.Context, args RemoveGroupArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.PageId == nil || *args.PageId == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "pageId"}
	}
	routeValues["pageId"] = *args.PageId
	if args.SectionId == nil || *args.SectionId == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "sectionId"}
	}
	routeValues["sectionId"] = *args.SectionId
	if args.GroupId == nil || *args.GroupId == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId

	locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RemoveGroup function
type RemoveGroupArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) The ID of the page the group is in
	PageId *string
	// (required) The ID of the section to the group is in
	SectionId *string
	// (required) The ID of the group
	GroupId *string
}

// [Preview API] Updates a group in the work item form.
func (client *Client) UpdateGroup(ctx context.Context, args UpdateGroupArgs) (*Group, error) {
	if args.Group == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "group"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.PageId == nil || *args.PageId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "pageId"}
	}
	routeValues["pageId"] = *args.PageId
	if args.SectionId == nil || *args.SectionId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "sectionId"}
	}
	routeValues["sectionId"] = *args.SectionId
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "groupId"}
	}
	routeValues["groupId"] = *args.GroupId

	body, marshalErr := json.Marshal(*args.Group)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Group
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateGroup function
type UpdateGroupArgs struct {
	// (required) The updated group.
	Group *Group
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
	// (required) The ID of the page the group is in.
	PageId *string
	// (required) The ID of the section the group is in.
	SectionId *string
	// (required) The ID of the group.
	GroupId *string
}

// [Preview API] Gets the form layout.
func (client *Client) GetFormLayout(ctx context.Context, args GetFormLayoutArgs) (*FormLayout, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	locationId, _ := uuid.Parse("fa8646eb-43cd-4b71-9564-40106fd63e40")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue FormLayout
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetFormLayout function
type GetFormLayoutArgs struct {
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
}

// [Preview API] Creates a picklist.
func (client *Client) CreateList(ctx context.Context, args CreateListArgs) (*PickList, error) {
	if args.Picklist == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "picklist"}
	}
	body, marshalErr := json.Marshal(*args.Picklist)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PickList
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateList function
type CreateListArgs struct {
	// (required) Picklist
	Picklist *PickList
}

// [Preview API] Removes a picklist.
func (client *Client) DeleteList(ctx context.Context, args DeleteListArgs) error {
	routeValues := make(map[string]string)
	if args.ListId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "listId"}
	}
	routeValues["listId"] = (*args.ListId).String()

	locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteList function
type DeleteListArgs struct {
	// (required) The ID of the list
	ListId *uuid.UUID
}

// [Preview API] Returns a picklist.
func (client *Client) GetList(ctx context.Context, args GetListArgs) (*PickList, error) {
	routeValues := make(map[string]string)
	if args.ListId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "listId"}
	}
	routeValues["listId"] = (*args.ListId).String()

	locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PickList
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetList function
type GetListArgs struct {
	// (required) The ID of the list
	ListId *uuid.UUID
}

// [Preview API] Returns meta data of the picklist.
func (client *Client) GetListsMetadata(ctx context.Context, args GetListsMetadataArgs) (*[]PickListMetadata, error) {
	locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []PickListMetadata
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetListsMetadata function
type GetListsMetadataArgs struct {
}

// [Preview API] Updates a list.
func (client *Client) UpdateList(ctx context.Context, args UpdateListArgs) (*PickList, error) {
	if args.Picklist == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "picklist"}
	}
	routeValues := make(map[string]string)
	if args.ListId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "listId"}
	}
	routeValues["listId"] = (*args.ListId).String()

	body, marshalErr := json.Marshal(*args.Picklist)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PickList
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateList function
type UpdateListArgs struct {
	// (required)
	Picklist *PickList
	// (required) The ID of the list
	ListId *uuid.UUID
}

// [Preview API] Adds a page to the work item form.
func (client *Client) AddPage(ctx context.Context, args AddPageArgs) (*Page, error) {
	if args.Page == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "page"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	body, marshalErr := json.Marshal(*args.Page)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1cc7b29f-6697-4d9d-b0a1-2650d3e1d584")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Page
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddPage function
type AddPageArgs struct {
	// (required) The page.
	Page *Page
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
}

// [Preview API] Removes a page from the work item form
func (client *Client) RemovePage(ctx context.Context, args RemovePageArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.PageId == nil || *args.PageId == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "pageId"}
	}
	routeValues["pageId"] = *args.PageId

	locationId, _ := uuid.Parse("1cc7b29f-6697-4d9d-b0a1-2650d3e1d584")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RemovePage function
type RemovePageArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) The ID of the page
	PageId *string
}

// [Preview API] Updates a page on the work item form
func (client *Client) UpdatePage(ctx context.Context, args UpdatePageArgs) (*Page, error) {
	if args.Page == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "page"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	body, marshalErr := json.Marshal(*args.Page)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1cc7b29f-6697-4d9d-b0a1-2650d3e1d584")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Page
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdatePage function
type UpdatePageArgs struct {
	// (required) The page
	Page *Page
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
}

// [Preview API] Creates a process.
func (client *Client) CreateNewProcess(ctx context.Context, args CreateNewProcessArgs) (*ProcessInfo, error) {
	if args.CreateRequest == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "createRequest"}
	}
	body, marshalErr := json.Marshal(*args.CreateRequest)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateNewProcess function
type CreateNewProcessArgs struct {
	// (required) CreateProcessModel.
	CreateRequest *CreateProcessModel
}

// [Preview API] Removes a process of a specific ID.
func (client *Client) DeleteProcessById(ctx context.Context, args DeleteProcessByIdArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessTypeId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processTypeId"}
	}
	routeValues["processTypeId"] = (*args.ProcessTypeId).String()

	locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteProcessById function
type DeleteProcessByIdArgs struct {
	// (required)
	ProcessTypeId *uuid.UUID
}

// [Preview API] Edit a process of a specific ID.
func (client *Client) EditProcess(ctx context.Context, args EditProcessArgs) (*ProcessInfo, error) {
	if args.UpdateRequest == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "updateRequest"}
	}
	routeValues := make(map[string]string)
	if args.ProcessTypeId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processTypeId"}
	}
	routeValues["processTypeId"] = (*args.ProcessTypeId).String()

	body, marshalErr := json.Marshal(*args.UpdateRequest)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the EditProcess function
type EditProcessArgs struct {
	// (required)
	UpdateRequest *UpdateProcessModel
	// (required)
	ProcessTypeId *uuid.UUID
}

// [Preview API] Get list of all processes including system and inherited.
func (client *Client) GetListOfProcesses(ctx context.Context, args GetListOfProcessesArgs) (*[]ProcessInfo, error) {
	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ProcessInfo
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetListOfProcesses function
type GetListOfProcessesArgs struct {
	// (optional)
	Expand *GetProcessExpandLevel
}

// [Preview API] Get a single process of a specified ID.
func (client *Client) GetProcessByItsId(ctx context.Context, args GetProcessByItsIdArgs) (*ProcessInfo, error) {
	routeValues := make(map[string]string)
	if args.ProcessTypeId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processTypeId"}
	}
	routeValues["processTypeId"] = (*args.ProcessTypeId).String()

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessInfo
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessByItsId function
type GetProcessByItsIdArgs struct {
	// (required)
	ProcessTypeId *uuid.UUID
	// (optional)
	Expand *GetProcessExpandLevel
}

// [Preview API] Adds a rule to work item type in the process.
func (client *Client) AddProcessWorkItemTypeRule(ctx context.Context, args AddProcessWorkItemTypeRuleArgs) (*ProcessRule, error) {
	if args.ProcessRuleCreate == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processRuleCreate"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	body, marshalErr := json.Marshal(*args.ProcessRuleCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessRule
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddProcessWorkItemTypeRule function
type AddProcessWorkItemTypeRuleArgs struct {
	// (required)
	ProcessRuleCreate *CreateProcessRuleRequest
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
}

// [Preview API] Removes a rule from the work item type in the process.
func (client *Client) DeleteProcessWorkItemTypeRule(ctx context.Context, args DeleteProcessWorkItemTypeRuleArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.RuleId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "ruleId"}
	}
	routeValues["ruleId"] = (*args.RuleId).String()

	locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteProcessWorkItemTypeRule function
type DeleteProcessWorkItemTypeRuleArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) The ID of the rule
	RuleId *uuid.UUID
}

// [Preview API] Returns a single rule in the work item type of the process.
func (client *Client) GetProcessWorkItemTypeRule(ctx context.Context, args GetProcessWorkItemTypeRuleArgs) (*ProcessRule, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.RuleId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "ruleId"}
	}
	routeValues["ruleId"] = (*args.RuleId).String()

	locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessRule
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessWorkItemTypeRule function
type GetProcessWorkItemTypeRuleArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) The ID of the rule
	RuleId *uuid.UUID
}

// [Preview API] Returns a list of all rules in the work item type of the process.
func (client *Client) GetProcessWorkItemTypeRules(ctx context.Context, args GetProcessWorkItemTypeRulesArgs) (*[]ProcessRule, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ProcessRule
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessWorkItemTypeRules function
type GetProcessWorkItemTypeRulesArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
}

// [Preview API] Updates a rule in the work item type of the process.
func (client *Client) UpdateProcessWorkItemTypeRule(ctx context.Context, args UpdateProcessWorkItemTypeRuleArgs) (*ProcessRule, error) {
	if args.ProcessRule == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processRule"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.RuleId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "ruleId"}
	}
	routeValues["ruleId"] = (*args.RuleId).String()

	body, marshalErr := json.Marshal(*args.ProcessRule)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessRule
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateProcessWorkItemTypeRule function
type UpdateProcessWorkItemTypeRuleArgs struct {
	// (required)
	ProcessRule *UpdateProcessRuleRequest
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) The ID of the rule
	RuleId *uuid.UUID
}

// [Preview API] Creates a state definition in the work item type of the process.
func (client *Client) CreateStateDefinition(ctx context.Context, args CreateStateDefinitionArgs) (*WorkItemStateResultModel, error) {
	if args.StateModel == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "stateModel"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	body, marshalErr := json.Marshal(*args.StateModel)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemStateResultModel
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateStateDefinition function
type CreateStateDefinitionArgs struct {
	// (required)
	StateModel *WorkItemStateInputModel
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
}

// [Preview API] Removes a state definition in the work item type of the process.
func (client *Client) DeleteStateDefinition(ctx context.Context, args DeleteStateDefinitionArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.StateId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "stateId"}
	}
	routeValues["stateId"] = (*args.StateId).String()

	locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteStateDefinition function
type DeleteStateDefinitionArgs struct {
	// (required) ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) ID of the state
	StateId *uuid.UUID
}

// [Preview API] Returns a single state definition in a work item type of the process.
func (client *Client) GetStateDefinition(ctx context.Context, args GetStateDefinitionArgs) (*WorkItemStateResultModel, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.StateId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "stateId"}
	}
	routeValues["stateId"] = (*args.StateId).String()

	locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemStateResultModel
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetStateDefinition function
type GetStateDefinitionArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) The ID of the state
	StateId *uuid.UUID
}

// [Preview API] Returns a list of all state definitions in a work item type of the process.
func (client *Client) GetStateDefinitions(ctx context.Context, args GetStateDefinitionsArgs) (*[]WorkItemStateResultModel, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []WorkItemStateResultModel
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetStateDefinitions function
type GetStateDefinitionsArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
}

// [Preview API] Hides a state definition in the work item type of the process.Only states with customizationType:System can be hidden.
func (client *Client) HideStateDefinition(ctx context.Context, args HideStateDefinitionArgs) (*WorkItemStateResultModel, error) {
	if args.HideStateModel == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "hideStateModel"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.StateId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "stateId"}
	}
	routeValues["stateId"] = (*args.StateId).String()

	body, marshalErr := json.Marshal(*args.HideStateModel)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemStateResultModel
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the HideStateDefinition function
type HideStateDefinitionArgs struct {
	// (required)
	HideStateModel *HideStateModel
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) The ID of the state
	StateId *uuid.UUID
}

// [Preview API] Updates a given state definition in the work item type of the process.
func (client *Client) UpdateStateDefinition(ctx context.Context, args UpdateStateDefinitionArgs) (*WorkItemStateResultModel, error) {
	if args.StateModel == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "stateModel"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName
	if args.StateId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "stateId"}
	}
	routeValues["stateId"] = (*args.StateId).String()

	body, marshalErr := json.Marshal(*args.StateModel)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemStateResultModel
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateStateDefinition function
type UpdateStateDefinitionArgs struct {
	// (required)
	StateModel *WorkItemStateInputModel
	// (required) ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (required) ID of the state
	StateId *uuid.UUID
}

// [Preview API] Creates a work item type in the process.
func (client *Client) CreateProcessWorkItemType(ctx context.Context, args CreateProcessWorkItemTypeArgs) (*ProcessWorkItemType, error) {
	if args.WorkItemType == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "workItemType"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()

	body, marshalErr := json.Marshal(*args.WorkItemType)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessWorkItemType
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateProcessWorkItemType function
type CreateProcessWorkItemTypeArgs struct {
	// (required)
	WorkItemType *CreateProcessWorkItemTypeRequest
	// (required) The ID of the process on which to create work item type.
	ProcessId *uuid.UUID
}

// [Preview API] Removes a work itewm type in the process.
func (client *Client) DeleteProcessWorkItemType(ctx context.Context, args DeleteProcessWorkItemTypeArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteProcessWorkItemType function
type DeleteProcessWorkItemTypeArgs struct {
	// (required) The ID of the process.
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type.
	WitRefName *string
}

// [Preview API] Returns a single work item type in a process.
func (client *Client) GetProcessWorkItemType(ctx context.Context, args GetProcessWorkItemTypeArgs) (*ProcessWorkItemType, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessWorkItemType
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessWorkItemType function
type GetProcessWorkItemTypeArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
	// (optional) Flag to determine what properties of work item type to return
	Expand *GetWorkItemTypeExpand
}

// [Preview API] Returns a list of all work item types in a process.
func (client *Client) GetProcessWorkItemTypes(ctx context.Context, args GetProcessWorkItemTypesArgs) (*[]ProcessWorkItemType, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("$expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ProcessWorkItemType
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetProcessWorkItemTypes function
type GetProcessWorkItemTypesArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (optional) Flag to determine what properties of work item type to return
	Expand *GetWorkItemTypeExpand
}

// [Preview API] Updates a work item type of the process.
func (client *Client) UpdateProcessWorkItemType(ctx context.Context, args UpdateProcessWorkItemTypeArgs) (*ProcessWorkItemType, error) {
	if args.WorkItemTypeUpdate == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "workItemTypeUpdate"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefName == nil || *args.WitRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefName"}
	}
	routeValues["witRefName"] = *args.WitRefName

	body, marshalErr := json.Marshal(*args.WorkItemTypeUpdate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ProcessWorkItemType
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateProcessWorkItemType function
type UpdateProcessWorkItemTypeArgs struct {
	// (required)
	WorkItemTypeUpdate *UpdateProcessWorkItemTypeRequest
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) The reference name of the work item type
	WitRefName *string
}

// [Preview API] Adds a behavior to the work item type of the process.
func (client *Client) AddBehaviorToWorkItemType(ctx context.Context, args AddBehaviorToWorkItemTypeArgs) (*WorkItemTypeBehavior, error) {
	if args.Behavior == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "behavior"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefNameForBehaviors == nil || *args.WitRefNameForBehaviors == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"}
	}
	routeValues["witRefNameForBehaviors"] = *args.WitRefNameForBehaviors

	body, marshalErr := json.Marshal(*args.Behavior)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemTypeBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddBehaviorToWorkItemType function
type AddBehaviorToWorkItemTypeArgs struct {
	// (required)
	Behavior *WorkItemTypeBehavior
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) Work item type reference name for the behavior
	WitRefNameForBehaviors *string
}

// [Preview API] Returns a behavior for the work item type of the process.
func (client *Client) GetBehaviorForWorkItemType(ctx context.Context, args GetBehaviorForWorkItemTypeArgs) (*WorkItemTypeBehavior, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefNameForBehaviors == nil || *args.WitRefNameForBehaviors == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"}
	}
	routeValues["witRefNameForBehaviors"] = *args.WitRefNameForBehaviors
	if args.BehaviorRefName == nil || *args.BehaviorRefName == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"}
	}
	routeValues["behaviorRefName"] = *args.BehaviorRefName

	locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemTypeBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBehaviorForWorkItemType function
type GetBehaviorForWorkItemTypeArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) Work item type reference name for the behavior
	WitRefNameForBehaviors *string
	// (required) The reference name of the behavior
	BehaviorRefName *string
}

// [Preview API] Returns a list of all behaviors for the work item type of the process.
func (client *Client) GetBehaviorsForWorkItemType(ctx context.Context, args GetBehaviorsForWorkItemTypeArgs) (*[]WorkItemTypeBehavior, error) {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefNameForBehaviors == nil || *args.WitRefNameForBehaviors == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"}
	}
	routeValues["witRefNameForBehaviors"] = *args.WitRefNameForBehaviors

	locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []WorkItemTypeBehavior
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBehaviorsForWorkItemType function
type GetBehaviorsForWorkItemTypeArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) Work item type reference name for the behavior
	WitRefNameForBehaviors *string
}

// [Preview API] Removes a behavior for the work item type of the process.
func (client *Client) RemoveBehaviorFromWorkItemType(ctx context.Context, args RemoveBehaviorFromWorkItemTypeArgs) error {
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefNameForBehaviors == nil || *args.WitRefNameForBehaviors == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"}
	}
	routeValues["witRefNameForBehaviors"] = *args.WitRefNameForBehaviors
	if args.BehaviorRefName == nil || *args.BehaviorRefName == "" {
		return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"}
	}
	routeValues["behaviorRefName"] = *args.BehaviorRefName

	locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RemoveBehaviorFromWorkItemType function
type RemoveBehaviorFromWorkItemTypeArgs struct {
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) Work item type reference name for the behavior
	WitRefNameForBehaviors *string
	// (required) The reference name of the behavior
	BehaviorRefName *string
}

// [Preview API] Updates a behavior for the work item type of the process.
func (client *Client) UpdateBehaviorToWorkItemType(ctx context.Context, args UpdateBehaviorToWorkItemTypeArgs) (*WorkItemTypeBehavior, error) {
	if args.Behavior == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "behavior"}
	}
	routeValues := make(map[string]string)
	if args.ProcessId == nil {
		return nil, &azureDevOps.ArgumentNilError{ArgumentName: "processId"}
	}
	routeValues["processId"] = (*args.ProcessId).String()
	if args.WitRefNameForBehaviors == nil || *args.WitRefNameForBehaviors == "" {
		return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"}
	}
	routeValues["witRefNameForBehaviors"] = *args.WitRefNameForBehaviors

	body, marshalErr := json.Marshal(*args.Behavior)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemTypeBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateBehaviorToWorkItemType function
type UpdateBehaviorToWorkItemTypeArgs struct {
	// (required)
	Behavior *WorkItemTypeBehavior
	// (required) The ID of the process
	ProcessId *uuid.UUID
	// (required) Work item type reference name for the behavior
	WitRefNameForBehaviors *string
}
