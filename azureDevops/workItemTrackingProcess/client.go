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
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("5264459e-e5e0-4bd8-b118-0985e68a4ec5")

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API] Creates a single behavior in the given process.
// behavior (required)
// processId (required): The ID of the process
func (client Client) CreateProcessBehavior(behavior *ProcessBehaviorCreateRequest, processId *uuid.UUID) (*ProcessBehavior, error) {
    if behavior == nil {
        return nil, errors.New("behavior is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()

    body, marshalErr := json.Marshal(*behavior)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a behavior in the process.
// processId (required): The ID of the process
// behaviorRefName (required): The reference name of the behavior
func (client Client) DeleteProcessBehavior(processId *uuid.UUID, behaviorRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if behaviorRefName == nil || *behaviorRefName == "" {
        return errors.New("behaviorRefName is a required parameter")
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a behavior of the process.
// processId (required): The ID of the process
// behaviorRefName (required): The reference name of the behavior
// expand (optional)
func (client Client) GetProcessBehavior(processId *uuid.UUID, behaviorRefName *string, expand *string) (*ProcessBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if behaviorRefName == nil || *behaviorRefName == "" {
        return nil, errors.New("behaviorRefName is a required parameter")
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all behaviors in the process.
// processId (required): The ID of the process
// expand (optional)
func (client Client) GetProcessBehaviors(processId *uuid.UUID, expand *string) (*[]ProcessBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProcessBehavior
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Replaces a behavior in the process.
// behaviorData (required)
// processId (required): The ID of the process
// behaviorRefName (required): The reference name of the behavior
func (client Client) UpdateProcessBehavior(behaviorData *ProcessBehaviorUpdateRequest, processId *uuid.UUID, behaviorRefName *string) (*ProcessBehavior, error) {
    if behaviorData == nil {
        return nil, errors.New("behaviorData is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if behaviorRefName == nil || *behaviorRefName == "" {
        return nil, errors.New("behaviorRefName is a required parameter")
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    body, marshalErr := json.Marshal(*behaviorData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d1800200-f184-4e75-a5f2-ad0b04b4373e")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a control in a group.
// control (required): The control.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group to add the control to.
func (client Client) CreateControlInGroup(control *Control, processId *uuid.UUID, witRefName *string, groupId *string) (*Control, error) {
    if control == nil {
        return nil, errors.New("control is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId

    body, marshalErr := json.Marshal(*control)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Control
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Moves a control to a specified group.
// control (required): The control.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group to move the control to.
// controlId (required): The ID of the control.
// removeFromGroupId (optional): The group ID to remove the control from.
func (client Client) MoveControlToGroup(control *Control, processId *uuid.UUID, witRefName *string, groupId *string, controlId *string, removeFromGroupId *string) (*Control, error) {
    if control == nil {
        return nil, errors.New("control is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if controlId == nil || *controlId == "" {
        return nil, errors.New("controlId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Control
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a control from the work item form.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group.
// controlId (required): The ID of the control to remove.
func (client Client) RemoveControlFromGroup(processId *uuid.UUID, witRefName *string, groupId *string, controlId *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if controlId == nil || *controlId == "" {
        return errors.New("controlId is a required parameter")
    }
    routeValues["controlId"] = *controlId

    locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a control on the work item form.
// control (required): The updated control.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// groupId (required): The ID of the group.
// controlId (required): The ID of the control.
func (client Client) UpdateControl(control *Control, processId *uuid.UUID, witRefName *string, groupId *string, controlId *string) (*Control, error) {
    if control == nil {
        return nil, errors.New("control is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if controlId == nil || *controlId == "" {
        return nil, errors.New("controlId is a required parameter")
    }
    routeValues["controlId"] = *controlId

    body, marshalErr := json.Marshal(*control)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1f59b363-a2d0-4b7e-9bc6-eb9f5f3f0e58")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Control
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds a field to a work item type.
// field (required)
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) AddFieldToWorkItemType(field *AddProcessWorkItemTypeFieldRequest, processId *uuid.UUID, witRefName *string) (*ProcessWorkItemTypeField, error) {
    if field == nil {
        return nil, errors.New("field is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*field)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessWorkItemTypeField
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all fields in a work item type.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) GetAllWorkItemTypeFields(processId *uuid.UUID, witRefName *string) (*[]ProcessWorkItemTypeField, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProcessWorkItemTypeField
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a field in a work item type.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// fieldRefName (required): The reference name of the field.
func (client Client) GetWorkItemTypeField(processId *uuid.UUID, witRefName *string, fieldRefName *string) (*ProcessWorkItemTypeField, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if fieldRefName == nil || *fieldRefName == "" {
        return nil, errors.New("fieldRefName is a required parameter")
    }
    routeValues["fieldRefName"] = *fieldRefName

    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessWorkItemTypeField
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a field from a work item type. Does not permanently delete the field.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// fieldRefName (required): The reference name of the field.
func (client Client) RemoveWorkItemTypeField(processId *uuid.UUID, witRefName *string, fieldRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if fieldRefName == nil || *fieldRefName == "" {
        return errors.New("fieldRefName is a required parameter")
    }
    routeValues["fieldRefName"] = *fieldRefName

    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a field in a work item type.
// field (required)
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// fieldRefName (required): The reference name of the field.
func (client Client) UpdateWorkItemTypeField(field *UpdateProcessWorkItemTypeFieldRequest, processId *uuid.UUID, witRefName *string, fieldRefName *string) (*ProcessWorkItemTypeField, error) {
    if field == nil {
        return nil, errors.New("field is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if fieldRefName == nil || *fieldRefName == "" {
        return nil, errors.New("fieldRefName is a required parameter")
    }
    routeValues["fieldRefName"] = *fieldRefName

    body, marshalErr := json.Marshal(*field)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bc0ad8dc-e3f3-46b0-b06c-5bf861793196")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessWorkItemTypeField
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds a group to the work item form.
// group (required): The group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page to add the group to.
// sectionId (required): The ID of the section to add the group to.
func (client Client) AddGroup(group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string) (*Group, error) {
    if group == nil {
        return nil, errors.New("group is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, errors.New("pageId is a required parameter")
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, errors.New("sectionId is a required parameter")
    }
    routeValues["sectionId"] = *sectionId

    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Group
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Moves a group to a different page and section.
// group (required): The updated group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page the group is in.
// sectionId (required): The ID of the section the group is i.n
// groupId (required): The ID of the group.
// removeFromPageId (required): ID of the page to remove the group from.
// removeFromSectionId (required): ID of the section to remove the group from.
func (client Client) MoveGroupToPage(group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string, removeFromPageId *string, removeFromSectionId *string) (*Group, error) {
    if group == nil {
        return nil, errors.New("group is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, errors.New("pageId is a required parameter")
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, errors.New("sectionId is a required parameter")
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId

    queryParams := url.Values{}
    if removeFromPageId == nil {
        return nil, errors.New("removeFromPageId is a required parameter")
    }
    queryParams.Add("removeFromPageId", *removeFromPageId)
    if removeFromSectionId == nil {
        return nil, errors.New("removeFromSectionId is a required parameter")
    }
    queryParams.Add("removeFromSectionId", *removeFromSectionId)
    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Group
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Moves a group to a different section.
// group (required): The updated group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page the group is in.
// sectionId (required): The ID of the section the group is in.
// groupId (required): The ID of the group.
// removeFromSectionId (required): ID of the section to remove the group from.
func (client Client) MoveGroupToSection(group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string, removeFromSectionId *string) (*Group, error) {
    if group == nil {
        return nil, errors.New("group is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, errors.New("pageId is a required parameter")
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, errors.New("sectionId is a required parameter")
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId

    queryParams := url.Values{}
    if removeFromSectionId == nil {
        return nil, errors.New("removeFromSectionId is a required parameter")
    }
    queryParams.Add("removeFromSectionId", *removeFromSectionId)
    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Group
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a group from the work item form.
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// pageId (required): The ID of the page the group is in
// sectionId (required): The ID of the section to the group is in
// groupId (required): The ID of the group
func (client Client) RemoveGroup(processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return errors.New("pageId is a required parameter")
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return errors.New("sectionId is a required parameter")
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId

    locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a group in the work item form.
// group (required): The updated group.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
// pageId (required): The ID of the page the group is in.
// sectionId (required): The ID of the section the group is in.
// groupId (required): The ID of the group.
func (client Client) UpdateGroup(group *Group, processId *uuid.UUID, witRefName *string, pageId *string, sectionId *string, groupId *string) (*Group, error) {
    if group == nil {
        return nil, errors.New("group is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return nil, errors.New("pageId is a required parameter")
    }
    routeValues["pageId"] = *pageId
    if sectionId == nil || *sectionId == "" {
        return nil, errors.New("sectionId is a required parameter")
    }
    routeValues["sectionId"] = *sectionId
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId

    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("766e44e1-36a8-41d7-9050-c343ff02f7a5")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Group
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the form layout.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) GetFormLayout(processId *uuid.UUID, witRefName *string) (*FormLayout, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("fa8646eb-43cd-4b71-9564-40106fd63e40")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FormLayout
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a picklist.
// picklist (required): Picklist
func (client Client) CreateList(picklist *PickList) (*PickList, error) {
    if picklist == nil {
        return nil, errors.New("picklist is a required parameter")
    }
    body, marshalErr := json.Marshal(*picklist)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PickList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a picklist.
// listId (required): The ID of the list
func (client Client) DeleteList(listId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if listId == nil {
        return errors.New("listId is a required parameter")
    }
    routeValues["listId"] = (*listId).String()

    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a picklist.
// listId (required): The ID of the list
func (client Client) GetList(listId *uuid.UUID) (*PickList, error) {
    routeValues := make(map[string]string)
    if listId == nil {
        return nil, errors.New("listId is a required parameter")
    }
    routeValues["listId"] = (*listId).String()

    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PickList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns meta data of the picklist.
func (client Client) GetListsMetadata() (*[]PickListMetadata, error) {
    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PickListMetadata
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates a list.
// picklist (required)
// listId (required): The ID of the list
func (client Client) UpdateList(picklist *PickList, listId *uuid.UUID) (*PickList, error) {
    if picklist == nil {
        return nil, errors.New("picklist is a required parameter")
    }
    routeValues := make(map[string]string)
    if listId == nil {
        return nil, errors.New("listId is a required parameter")
    }
    routeValues["listId"] = (*listId).String()

    body, marshalErr := json.Marshal(*picklist)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("01e15468-e27c-4e20-a974-bd957dcccebc")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PickList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds a page to the work item form.
// page (required): The page.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) AddPage(page *Page, processId *uuid.UUID, witRefName *string) (*Page, error) {
    if page == nil {
        return nil, errors.New("page is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*page)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1cc7b29f-6697-4d9d-b0a1-2650d3e1d584")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Page
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a page from the work item form
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// pageId (required): The ID of the page
func (client Client) RemovePage(processId *uuid.UUID, witRefName *string, pageId *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if pageId == nil || *pageId == "" {
        return errors.New("pageId is a required parameter")
    }
    routeValues["pageId"] = *pageId

    locationId, _ := uuid.Parse("1cc7b29f-6697-4d9d-b0a1-2650d3e1d584")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a page on the work item form
// page (required): The page
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) UpdatePage(page *Page, processId *uuid.UUID, witRefName *string) (*Page, error) {
    if page == nil {
        return nil, errors.New("page is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*page)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1cc7b29f-6697-4d9d-b0a1-2650d3e1d584")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Page
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a process.
// createRequest (required): CreateProcessModel.
func (client Client) CreateNewProcess(createRequest *CreateProcessModel) (*ProcessInfo, error) {
    if createRequest == nil {
        return nil, errors.New("createRequest is a required parameter")
    }
    body, marshalErr := json.Marshal(*createRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessInfo
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a process of a specific ID.
// processTypeId (required)
func (client Client) DeleteProcessById(processTypeId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if processTypeId == nil {
        return errors.New("processTypeId is a required parameter")
    }
    routeValues["processTypeId"] = (*processTypeId).String()

    locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Edit a process of a specific ID.
// updateRequest (required)
// processTypeId (required)
func (client Client) EditProcess(updateRequest *UpdateProcessModel, processTypeId *uuid.UUID) (*ProcessInfo, error) {
    if updateRequest == nil {
        return nil, errors.New("updateRequest is a required parameter")
    }
    routeValues := make(map[string]string)
    if processTypeId == nil {
        return nil, errors.New("processTypeId is a required parameter")
    }
    routeValues["processTypeId"] = (*processTypeId).String()

    body, marshalErr := json.Marshal(*updateRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessInfo
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get list of all processes including system and inherited.
// expand (optional)
func (client Client) GetListOfProcesses(expand *string) (*[]ProcessInfo, error) {
    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProcessInfo
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a single process of a specified ID.
// processTypeId (required)
// expand (optional)
func (client Client) GetProcessByItsId(processTypeId *uuid.UUID, expand *string) (*ProcessInfo, error) {
    routeValues := make(map[string]string)
    if processTypeId == nil {
        return nil, errors.New("processTypeId is a required parameter")
    }
    routeValues["processTypeId"] = (*processTypeId).String()

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("02cc6a73-5cfb-427d-8c8e-b49fb086e8af")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessInfo
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds a rule to work item type in the process.
// processRuleCreate (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) AddProcessWorkItemTypeRule(processRuleCreate *CreateProcessRuleRequest, processId *uuid.UUID, witRefName *string) (*ProcessRule, error) {
    if processRuleCreate == nil {
        return nil, errors.New("processRuleCreate is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*processRuleCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessRule
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a rule from the work item type in the process.
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// ruleId (required): The ID of the rule
func (client Client) DeleteProcessWorkItemTypeRule(processId *uuid.UUID, witRefName *string, ruleId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if ruleId == nil {
        return errors.New("ruleId is a required parameter")
    }
    routeValues["ruleId"] = (*ruleId).String()

    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a single rule in the work item type of the process.
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// ruleId (required): The ID of the rule
func (client Client) GetProcessWorkItemTypeRule(processId *uuid.UUID, witRefName *string, ruleId *uuid.UUID) (*ProcessRule, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if ruleId == nil {
        return nil, errors.New("ruleId is a required parameter")
    }
    routeValues["ruleId"] = (*ruleId).String()

    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessRule
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all rules in the work item type of the process.
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) GetProcessWorkItemTypeRules(processId *uuid.UUID, witRefName *string) (*[]ProcessRule, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProcessRule
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates a rule in the work item type of the process.
// processRule (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// ruleId (required): The ID of the rule
func (client Client) UpdateProcessWorkItemTypeRule(processRule *UpdateProcessRuleRequest, processId *uuid.UUID, witRefName *string, ruleId *uuid.UUID) (*ProcessRule, error) {
    if processRule == nil {
        return nil, errors.New("processRule is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if ruleId == nil {
        return nil, errors.New("ruleId is a required parameter")
    }
    routeValues["ruleId"] = (*ruleId).String()

    body, marshalErr := json.Marshal(*processRule)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("76fe3432-d825-479d-a5f6-983bbb78b4f3")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessRule
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a state definition in the work item type of the process.
// stateModel (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) CreateStateDefinition(stateModel *WorkItemStateInputModel, processId *uuid.UUID, witRefName *string) (*WorkItemStateResultModel, error) {
    if stateModel == nil {
        return nil, errors.New("stateModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*stateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemStateResultModel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a state definition in the work item type of the process.
// processId (required): ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): ID of the state
func (client Client) DeleteStateDefinition(processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return errors.New("stateId is a required parameter")
    }
    routeValues["stateId"] = (*stateId).String()

    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a single state definition in a work item type of the process.
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): The ID of the state
func (client Client) GetStateDefinition(processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) (*WorkItemStateResultModel, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return nil, errors.New("stateId is a required parameter")
    }
    routeValues["stateId"] = (*stateId).String()

    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemStateResultModel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all state definitions in a work item type of the process.
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) GetStateDefinitions(processId *uuid.UUID, witRefName *string) (*[]WorkItemStateResultModel, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemStateResultModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Hides a state definition in the work item type of the process.Only states with customizationType:System can be hidden.
// hideStateModel (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): The ID of the state
func (client Client) HideStateDefinition(hideStateModel *HideStateModel, processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) (*WorkItemStateResultModel, error) {
    if hideStateModel == nil {
        return nil, errors.New("hideStateModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return nil, errors.New("stateId is a required parameter")
    }
    routeValues["stateId"] = (*stateId).String()

    body, marshalErr := json.Marshal(*hideStateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemStateResultModel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates a given state definition in the work item type of the process.
// stateModel (required)
// processId (required): ID of the process
// witRefName (required): The reference name of the work item type
// stateId (required): ID of the state
func (client Client) UpdateStateDefinition(stateModel *WorkItemStateInputModel, processId *uuid.UUID, witRefName *string, stateId *uuid.UUID) (*WorkItemStateResultModel, error) {
    if stateModel == nil {
        return nil, errors.New("stateModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName
    if stateId == nil {
        return nil, errors.New("stateId is a required parameter")
    }
    routeValues["stateId"] = (*stateId).String()

    body, marshalErr := json.Marshal(*stateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("31015d57-2dff-4a46-adb3-2fb4ee3dcec9")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemStateResultModel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a work item type in the process.
// workItemType (required)
// processId (required): The ID of the process on which to create work item type.
func (client Client) CreateProcessWorkItemType(workItemType *CreateProcessWorkItemTypeRequest, processId *uuid.UUID) (*ProcessWorkItemType, error) {
    if workItemType == nil {
        return nil, errors.New("workItemType is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()

    body, marshalErr := json.Marshal(*workItemType)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessWorkItemType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a work itewm type in the process.
// processId (required): The ID of the process.
// witRefName (required): The reference name of the work item type.
func (client Client) DeleteProcessWorkItemType(processId *uuid.UUID, witRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a single work item type in a process.
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
// expand (optional): Flag to determine what properties of work item type to return
func (client Client) GetProcessWorkItemType(processId *uuid.UUID, witRefName *string, expand *string) (*ProcessWorkItemType, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessWorkItemType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all work item types in a process.
// processId (required): The ID of the process
// expand (optional): Flag to determine what properties of work item type to return
func (client Client) GetProcessWorkItemTypes(processId *uuid.UUID, expand *string) (*[]ProcessWorkItemType, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", *expand)
    }
    locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ProcessWorkItemType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates a work item type of the process.
// workItemTypeUpdate (required)
// processId (required): The ID of the process
// witRefName (required): The reference name of the work item type
func (client Client) UpdateProcessWorkItemType(workItemTypeUpdate *UpdateProcessWorkItemTypeRequest, processId *uuid.UUID, witRefName *string) (*ProcessWorkItemType, error) {
    if workItemTypeUpdate == nil {
        return nil, errors.New("workItemTypeUpdate is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefName == nil || *witRefName == "" {
        return nil, errors.New("witRefName is a required parameter")
    }
    routeValues["witRefName"] = *witRefName

    body, marshalErr := json.Marshal(*workItemTypeUpdate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e2e9d1a6-432d-4062-8870-bfcb8c324ad7")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessWorkItemType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds a behavior to the work item type of the process.
// behavior (required)
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
func (client Client) AddBehaviorToWorkItemType(behavior *WorkItemTypeBehavior, processId *uuid.UUID, witRefNameForBehaviors *string) (*WorkItemTypeBehavior, error) {
    if behavior == nil {
        return nil, errors.New("behavior is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, errors.New("witRefNameForBehaviors is a required parameter")
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors

    body, marshalErr := json.Marshal(*behavior)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTypeBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a behavior for the work item type of the process.
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
// behaviorRefName (required): The reference name of the behavior
func (client Client) GetBehaviorForWorkItemType(processId *uuid.UUID, witRefNameForBehaviors *string, behaviorRefName *string) (*WorkItemTypeBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, errors.New("witRefNameForBehaviors is a required parameter")
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors
    if behaviorRefName == nil || *behaviorRefName == "" {
        return nil, errors.New("behaviorRefName is a required parameter")
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTypeBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of all behaviors for the work item type of the process.
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
func (client Client) GetBehaviorsForWorkItemType(processId *uuid.UUID, witRefNameForBehaviors *string) (*[]WorkItemTypeBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, errors.New("witRefNameForBehaviors is a required parameter")
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors

    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemTypeBehavior
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a behavior for the work item type of the process.
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
// behaviorRefName (required): The reference name of the behavior
func (client Client) RemoveBehaviorFromWorkItemType(processId *uuid.UUID, witRefNameForBehaviors *string, behaviorRefName *string) error {
    routeValues := make(map[string]string)
    if processId == nil {
        return errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return errors.New("witRefNameForBehaviors is a required parameter")
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors
    if behaviorRefName == nil || *behaviorRefName == "" {
        return errors.New("behaviorRefName is a required parameter")
    }
    routeValues["behaviorRefName"] = *behaviorRefName

    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Updates a behavior for the work item type of the process.
// behavior (required)
// processId (required): The ID of the process
// witRefNameForBehaviors (required): Work item type reference name for the behavior
func (client Client) UpdateBehaviorToWorkItemType(behavior *WorkItemTypeBehavior, processId *uuid.UUID, witRefNameForBehaviors *string) (*WorkItemTypeBehavior, error) {
    if behavior == nil {
        return nil, errors.New("behavior is a required parameter")
    }
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()
    if witRefNameForBehaviors == nil || *witRefNameForBehaviors == "" {
        return nil, errors.New("witRefNameForBehaviors is a required parameter")
    }
    routeValues["witRefNameForBehaviors"] = *witRefNameForBehaviors

    body, marshalErr := json.Marshal(*behavior)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6d765a2e-4e1b-4b11-be93-f953be676024")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTypeBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

