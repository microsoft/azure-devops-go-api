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
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("5264459e-e5e0-4bd8-b118-0985e68a4ec5")

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

// [Preview API] Creates a single behavior in the given process.
// ctx
// behavior (required)
// processId (required): The ID of the process
func (client Client) CreateProcessBehavior(ctx context.Context, behavior *ProcessBehaviorCreateRequest, processId *uuid.UUID) (*ProcessBehavior, error) {
    if behavior == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "behavior"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()

    body, marshalErr := json.Marshal(*behavior)
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

// [Preview API] Removes a behavior in the process.
// ctx
// processId (required): The ID of the process
// behaviorRefName (required): The reference name of the behavior
func (client Client) DeleteProcessBehavior(ctx context.Context, processId *uuid.UUID, behaviorRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if behaviorRefName == nil || *behaviorRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"} 
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a behavior of the process.
// ctx
// processId (required): The ID of the process
// behaviorRefName (required): The reference name of the behavior
// expand (optional)
func (client Client) GetProcessBehavior(ctx context.Context, processId *uuid.UUID, behaviorRefName *string, expand *GetBehaviorsExpand) (*ProcessBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if behaviorRefName == nil || *behaviorRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"} 
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
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

// [Preview API] Returns a list of all behaviors in the process.
// ctx
// processId (required): The ID of the process
// expand (optional)
func (client Client) GetProcessBehaviors(ctx context.Context, processId *uuid.UUID, expand *GetBehaviorsExpand) (*[]ProcessBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
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

// [Preview API] Replaces a behavior in the process.
// ctx
// behaviorData (required)
// processId (required): The ID of the process
// behaviorRefName (required): The reference name of the behavior
func (client Client) UpdateProcessBehavior(ctx context.Context, behaviorData *ProcessBehaviorUpdateRequest, processId *uuid.UUID, behaviorRefName *string) (*ProcessBehavior, error) {
    if behaviorData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "behaviorData"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if behaviorRefName == nil || *behaviorRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"} 
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    body, marshalErr := json.Marshal(*behaviorData)
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

// [Preview API] Creates a control in a group.
// ctx
// control (required): The control.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group to add the control to.
func (client Client) CreateControlInGroup(ctx context.Context, control *Control, processId *uuid.UUID, witRefName *string, groupId *string) (*Control, error) {
    if control == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "control"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId

    body, marshalErr := json.Marshal(*control)
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

// [Preview API] Moves a control to a specified group.
// ctx
// control (required): The control.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group to move the control to.
// controlId (required): The ID of the control.
// removeFromGroupId (optional): The group ID to remove the control from.
func (client Client) MoveControlToGroup(ctx context.Context, control *Control, processId *uuid.UUID, witRefName *string, groupId *string, controlId *string, removeFromGroupId *string) (*Control, error) {
    if control == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "control"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId
    if controlId == nil || *controlId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "controlId"} 
    }
    routeValues["controlId"] = *controlId

    queryParams := url.Values{}
    if removeFromGroupId != nil {
        queryParams.Add("removeFromGroupId", *removeFromGroupId)
    }
    body, marshalErr := json.Marshal(*control)
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

// [Preview API] Removes a control from the work item form.
// ctx
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group.
// controlId (required): The ID of the control to remove.
func (client Client) RemoveControlFromGroup(ctx context.Context, processId *uuid.UUID, witRefName *string, groupId *string, controlId *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId
    if controlId == nil || *controlId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "controlId"} 
    }
    routeValues["controlId"] = *controlId

    locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a control on the work item form.
// ctx
// control (required): The updated control.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group.
// controlId (required): The ID of the control.
func (client Client) UpdateControl(ctx context.Context, control *Control, processId *uuid.UUID, witRefName *string, groupId *string, controlId *string) (*Control, error) {
    if control == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "control"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId
    if controlId == nil || *controlId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "controlId"} 
    }
    routeValues["controlId"] = *controlId

    body, marshalErr := json.Marshal(*control)
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

// [Preview API] Adds a field to a work item type.
// ctx
// field (required)
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) AddFieldToWorkItemType(ctx context.Context, field *AddProcessWorkItemTypeFieldRequest, processId *uuid.UUID, witRefName *string) (*ProcessWorkItemTypeField, error) {
    if field == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "field"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*field)
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

// [Preview API] Returns a list of all fields in a work item type.
// ctx
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) GetAllWorkItemTypeFields(ctx context.Context, processId *uuid.UUID, witRefName *string) (*[]ProcessWorkItemTypeField, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProcessWorkItemTypeField
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a field in a work item type.
// ctx
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// fieldRefName (required): The reference name of the field.
func (client Client) GetWorkItemTypeField(ctx context.Context, processId *uuid.UUID, witRefName *string, fieldRefName *string) (*ProcessWorkItemTypeField, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if fieldRefName == nil || *fieldRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fieldRefName"} 
    }
    routeValues["fieldRefName"] = *fieldRefName

    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessWorkItemTypeField
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a field from a work item type. Does not permanently delete the field.
// ctx
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// fieldRefName (required): The reference name of the field.
func (client Client) RemoveWorkItemTypeField(ctx context.Context, processId *uuid.UUID, witRefName *string, fieldRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if fieldRefName == nil || *fieldRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fieldRefName"} 
    }
    routeValues["fieldRefName"] = *fieldRefName

    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a field in a work item type.
// ctx
// field (required)
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// fieldRefName (required): The reference name of the field.
func (client Client) UpdateWorkItemTypeField(ctx context.Context, field *UpdateProcessWorkItemTypeFieldRequest, processId *uuid.UUID, witRefName *string, fieldRefName *string) (*ProcessWorkItemTypeField, error) {
    if field == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "field"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if fieldRefName == nil || *fieldRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fieldRefName"} 
    }
    routeValues["fieldRefName"] = *fieldRefName

    body, marshalErr := json.Marshal(*field)
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

// [Preview API] Adds a group to the work item form.
// ctx
// group (required): The group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page to add the group to.
// sectionId (required): The ID of the section to add the group to.
func (client Client) AddGroup(ctx context.Context, group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string) (*Group, error) {
    if group == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "group"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pageId"} 
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sectionId"} 
    }
    routeValues["sectionId"] = *sectionId

    body, marshalErr := json.Marshal(*group)
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

// [Preview API] Moves a group to a different page and section.
// ctx
// group (required): The updated group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page the group is in.
// sectionId (required): The ID of the section the group is i.n
// groupId (required): The ID of the group.
// removeFromPageId (required): ID of the page to remove the group from.
// removeFromSectionId (required): ID of the section to remove the group from.
func (client Client) MoveGroupToPage(ctx context.Context, group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string, removeFromPageId *string, removeFromSectionId *string) (*Group, error) {
    if group == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "group"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pageId"} 
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sectionId"} 
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId

    queryParams := url.Values{}
    if removeFromPageId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "removeFromPageId"}
    }
    queryParams.Add("removeFromPageId", *removeFromPageId)
    if removeFromSectionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "removeFromSectionId"}
    }
    queryParams.Add("removeFromSectionId", *removeFromSectionId)
    body, marshalErr := json.Marshal(*group)
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

// [Preview API] Moves a group to a different section.
// ctx
// group (required): The updated group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page the group is in.
// sectionId (required): The ID of the section the group is in.
// groupId (required): The ID of the group.
// removeFromSectionId (required): ID of the section to remove the group from.
func (client Client) MoveGroupToSection(ctx context.Context, group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string, removeFromSectionId *string) (*Group, error) {
    if group == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "group"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pageId"} 
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sectionId"} 
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId

    queryParams := url.Values{}
    if removeFromSectionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "removeFromSectionId"}
    }
    queryParams.Add("removeFromSectionId", *removeFromSectionId)
    body, marshalErr := json.Marshal(*group)
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

// [Preview API] Removes a group from the work item form.
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// pageId (required): The ID of the page the group is in
// sectionId (required): The ID of the section to the group is in
// groupId (required): The ID of the group
func (client Client) RemoveGroup(ctx context.Context, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pageId"} 
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sectionId"} 
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId

    locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a group in the work item form.
// ctx
// group (required): The updated group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page the group is in.
// sectionId (required): The ID of the section the group is in.
// groupId (required): The ID of the group.
func (client Client) UpdateGroup(ctx context.Context, group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string) (*Group, error) {
    if group == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "group"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pageId"} 
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sectionId"} 
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupId"} 
    }
    routeValues["groupId"] = *groupId

    body, marshalErr := json.Marshal(*group)
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

// [Preview API] Gets the form layout.
// ctx
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) GetFormLayout(ctx context.Context, processId *uuid.UUID, witRefName *string) (*FormLayout, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("fa8646eb-43cd-4b71-9564-40106fd63e40")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FormLayout
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a picklist.
// ctx
// picklist (required): Picklist
func (client Client) CreateList(ctx context.Context, picklist *PickList) (*PickList, error) {
    if picklist == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "picklist"}
    }
    body, marshalErr := json.Marshal(*picklist)
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

// [Preview API] Removes a picklist.
// ctx
// listId (required): The ID of the list
func (client Client) DeleteList(ctx context.Context, listId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if listId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "listId"} 
    }
    routeValues["listId"] = (*listId).String()

    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a picklist.
// ctx
// listId (required): The ID of the list
func (client Client) GetList(ctx context.Context, listId *uuid.UUID) (*PickList, error) {
    routeValues := make(map[string]string)
    if listId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "listId"} 
    }
    routeValues["listId"] = (*listId).String()

    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PickList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns meta data of the picklist.
// ctx
func (client Client) GetListsMetadata(ctx context.Context, ) (*[]PickListMetadata, error) {
    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PickListMetadata
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates a list.
// ctx
// picklist (required)
// listId (required): The ID of the list
func (client Client) UpdateList(ctx context.Context, picklist *PickList, listId *uuid.UUID) (*PickList, error) {
    if picklist == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "picklist"}
    }
    routeValues := make(map[string]string)
    if listId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "listId"} 
    }
    routeValues["listId"] = (*listId).String()

    body, marshalErr := json.Marshal(*picklist)
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

// [Preview API] Adds a page to the work item form.
// ctx
// page (required): The page.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) AddPage(ctx context.Context, page *Page, processId *uuid.UUID, witRefName *string) (*Page, error) {
    if page == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "page"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*page)
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

// [Preview API] Removes a page from the work item form
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// pageId (required): The ID of the page
func (client Client) RemovePage(ctx context.Context, processId *uuid.UUID, witRefName *string, pageId *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pageId"} 
    }
    routeValues["pageId"] = *pageId

    locationId, _ := uuid.Parse("1cc7b29f-6697-4d9d-b0a1-2650d3e1d584")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a page on the work item form
// ctx
// page (required): The page
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) UpdatePage(ctx context.Context, page *Page, processId *uuid.UUID, witRefName *string) (*Page, error) {
    if page == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "page"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*page)
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

// [Preview API] Creates a process.
// ctx
// createRequest (required): CreateProcessModel.
func (client Client) CreateNewProcess(ctx context.Context, createRequest *CreateProcessModel) (*ProcessInfo, error) {
    if createRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "createRequest"}
    }
    body, marshalErr := json.Marshal(*createRequest)
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

// [Preview API] Removes a process of a specific ID.
// ctx
// processTypeId (required)
func (client Client) DeleteProcessById(ctx context.Context, processTypeId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if processTypeId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processTypeId"} 
    }
    routeValues["processTypeId"] = (*processTypeId).String()

    locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Edit a process of a specific ID.
// ctx
// updateRequest (required)
// processTypeId (required)
func (client Client) EditProcess(ctx context.Context, updateRequest *UpdateProcessModel, processTypeId *uuid.UUID) (*ProcessInfo, error) {
    if updateRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateRequest"}
    }
    routeValues := make(map[string]string)
    if processTypeId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processTypeId"} 
    }
    routeValues["processTypeId"] = (*processTypeId).String()

    body, marshalErr := json.Marshal(*updateRequest)
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

// [Preview API] Get list of all processes including system and inherited.
// ctx
// expand (optional)
func (client Client) GetListOfProcesses(ctx context.Context, expand *GetProcessExpandLevel) (*[]ProcessInfo, error) {
    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
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

// [Preview API] Get a single process of a specified ID.
// ctx
// processTypeId (required)
// expand (optional)
func (client Client) GetProcessByItsId(ctx context.Context, processTypeId *uuid.UUID, expand *GetProcessExpandLevel) (*ProcessInfo, error) {
    routeValues := make(map[string]string)
    if processTypeId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processTypeId"} 
    }
    routeValues["processTypeId"] = (*processTypeId).String()

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
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

// [Preview API] Adds a rule to work item type in the process.
// ctx
// processRuleCreate (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) AddProcessWorkItemTypeRule(ctx context.Context, processRuleCreate *CreateProcessRuleRequest, processId *uuid.UUID, witRefName *string) (*ProcessRule, error) {
    if processRuleCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processRuleCreate"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*processRuleCreate)
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

// [Preview API] Removes a rule from the work item type in the process.
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// ruleId (required): The ID of the rule
func (client Client) DeleteProcessWorkItemTypeRule(ctx context.Context, processId *uuid.UUID, witRefName *string, ruleId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if ruleId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "ruleId"} 
    }
    routeValues["ruleId"] = (*ruleId).String()

    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a single rule in the work item type of the process.
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// ruleId (required): The ID of the rule
func (client Client) GetProcessWorkItemTypeRule(ctx context.Context, processId *uuid.UUID, witRefName *string, ruleId *uuid.UUID) (*ProcessRule, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if ruleId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "ruleId"} 
    }
    routeValues["ruleId"] = (*ruleId).String()

    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessRule
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all rules in the work item type of the process.
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) GetProcessWorkItemTypeRules(ctx context.Context, processId *uuid.UUID, witRefName *string) (*[]ProcessRule, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProcessRule
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates a rule in the work item type of the process.
// ctx
// processRule (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// ruleId (required): The ID of the rule
func (client Client) UpdateProcessWorkItemTypeRule(ctx context.Context, processRule *UpdateProcessRuleRequest, processId *uuid.UUID, witRefName *string, ruleId *uuid.UUID) (*ProcessRule, error) {
    if processRule == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processRule"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if ruleId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "ruleId"} 
    }
    routeValues["ruleId"] = (*ruleId).String()

    body, marshalErr := json.Marshal(*processRule)
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

// [Preview API] Creates a state definition in the work item type of the process.
// ctx
// stateModel (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) CreateStateDefinition(ctx context.Context, stateModel *WorkItemStateInputModel, processId *uuid.UUID, witRefName *string) (*WorkItemStateResultModel, error) {
    if stateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "stateModel"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*stateModel)
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

// [Preview API] Removes a state definition in the work item type of the process.
// ctx
// processId (required): ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): ID of the state
func (client Client) DeleteStateDefinition(ctx context.Context, processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "stateId"} 
    }
    routeValues["stateId"] = (*stateId).String()

    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a single state definition in a work item type of the process.
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): The ID of the state
func (client Client) GetStateDefinition(ctx context.Context, processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) (*WorkItemStateResultModel, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "stateId"} 
    }
    routeValues["stateId"] = (*stateId).String()

    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemStateResultModel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all state definitions in a work item type of the process.
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) GetStateDefinitions(ctx context.Context, processId *uuid.UUID, witRefName *string) (*[]WorkItemStateResultModel, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemStateResultModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Hides a state definition in the work item type of the process.Only states with customizationType:System can be hidden.
// ctx
// hideStateModel (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): The ID of the state
func (client Client) HideStateDefinition(ctx context.Context, hideStateModel *HideStateModel, processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) (*WorkItemStateResultModel, error) {
    if hideStateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "hideStateModel"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "stateId"} 
    }
    routeValues["stateId"] = (*stateId).String()

    body, marshalErr := json.Marshal(*hideStateModel)
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

// [Preview API] Updates a given state definition in the work item type of the process.
// ctx
// stateModel (required)
// processId (required): ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): ID of the state
func (client Client) UpdateStateDefinition(ctx context.Context, stateModel *WorkItemStateInputModel, processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) (*WorkItemStateResultModel, error) {
    if stateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "stateModel"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "stateId"} 
    }
    routeValues["stateId"] = (*stateId).String()

    body, marshalErr := json.Marshal(*stateModel)
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

// [Preview API] Creates a work item type in the process.
// ctx
// workItemType (required)
// processId (required): The ID of the process on which to create work item type.
func (client Client) CreateProcessWorkItemType(ctx context.Context, workItemType *CreateProcessWorkItemTypeRequest, processId *uuid.UUID) (*ProcessWorkItemType, error) {
    if workItemType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemType"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()

    body, marshalErr := json.Marshal(*workItemType)
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

// [Preview API] Removes a work itewm type in the process.
// ctx
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) DeleteProcessWorkItemType(ctx context.Context, processId *uuid.UUID, witRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a single work item type in a process.
// ctx
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// expand (optional): Flag to determine what properties of work item type to return
func (client Client) GetProcessWorkItemType(ctx context.Context, processId *uuid.UUID, witRefName *string, expand *GetWorkItemTypeExpand) (*ProcessWorkItemType, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
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

// [Preview API] Returns a list of all work item types in a process.
// ctx
// processId (required): The ID of the process
// expand (optional): Flag to determine what properties of work item type to return
func (client Client) GetProcessWorkItemTypes(ctx context.Context, processId *uuid.UUID, expand *GetWorkItemTypeExpand) (*[]ProcessWorkItemType, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
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

// [Preview API] Updates a work item type of the process.
// ctx
// workItemTypeUpdate (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) UpdateProcessWorkItemType(ctx context.Context, workItemTypeUpdate *UpdateProcessWorkItemTypeRequest, processId *uuid.UUID, witRefName *string) (*ProcessWorkItemType, error) {
    if workItemTypeUpdate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemTypeUpdate"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefName"} 
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*workItemTypeUpdate)
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

// [Preview API] Adds a behavior to the work item type of the process.
// ctx
// behavior (required)
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
func (client Client) AddBehaviorToWorkItemType(ctx context.Context, behavior *WorkItemTypeBehavior, processId *uuid.UUID, witRefNameForBehaviors *string) (*WorkItemTypeBehavior, error) {
    if behavior == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "behavior"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"} 
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors

    body, marshalErr := json.Marshal(*behavior)
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

// [Preview API] Returns a behavior for the work item type of the process.
// ctx
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
// behaviorRefName (required): The reference name of the behavior
func (client Client) GetBehaviorForWorkItemType(ctx context.Context, processId *uuid.UUID, witRefNameForBehaviors *string, behaviorRefName *string) (*WorkItemTypeBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"} 
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors
    if behaviorRefName == nil || *behaviorRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"} 
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTypeBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all behaviors for the work item type of the process.
// ctx
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
func (client Client) GetBehaviorsForWorkItemType(ctx context.Context, processId *uuid.UUID, witRefNameForBehaviors *string) (*[]WorkItemTypeBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"} 
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors

    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemTypeBehavior
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a behavior for the work item type of the process.
// ctx
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
// behaviorRefName (required): The reference name of the behavior
func (client Client) RemoveBehaviorFromWorkItemType(ctx context.Context, processId *uuid.UUID, witRefNameForBehaviors *string, behaviorRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"} 
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors
    if behaviorRefName == nil || *behaviorRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "behaviorRefName"} 
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a behavior for the work item type of the process.
// ctx
// behavior (required)
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
func (client Client) UpdateBehaviorToWorkItemType(ctx context.Context, behavior *WorkItemTypeBehavior, processId *uuid.UUID, witRefNameForBehaviors *string) (*WorkItemTypeBehavior, error) {
    if behavior == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "behavior"}
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "processId"} 
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "witRefNameForBehaviors"} 
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors

    body, marshalErr := json.Marshal(*behavior)
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

