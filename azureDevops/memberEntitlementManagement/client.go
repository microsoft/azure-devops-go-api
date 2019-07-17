// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package memberEntitlementManagement

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("68ddce18-2501-45f1-a17b-7931a9922690")

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

// [Preview API] Create a group entitlement with license rule, extension rule.
// groupEntitlement (required): GroupEntitlement object specifying License Rule, Extensions Rule for the group. Based on the rules the members of the group will be given licenses and extensions. The Group Entitlement can be used to add the group to another project level groups
// ruleOption (optional): RuleOption [ApplyGroupRule/TestApplyGroupRule] - specifies if the rules defined in group entitlement should be created and applied to it’s members (default option) or just be tested
func (client Client) AddGroupEntitlement(groupEntitlement *GroupEntitlement, ruleOption *string) (*GroupEntitlementOperationReference, error) {
    if groupEntitlement == nil {
        return nil, errors.New("groupEntitlement is a required parameter")
    }
    queryParams := url.Values{}
    if ruleOption != nil {
        queryParams.Add("ruleOption", *ruleOption)
    }
    body, marshalErr := json.Marshal(*groupEntitlement)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2280bffa-58a2-49da-822e-0764a1bb44f7")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GroupEntitlementOperationReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a group entitlement.
// groupId (required): ID of the group to delete.
// ruleOption (optional): RuleOption [ApplyGroupRule/TestApplyGroupRule] - specifies if the rules defined in group entitlement should be deleted and the changes are applied to it’s members (default option) or just be tested
// removeGroupMembership (optional): Optional parameter that specifies whether the group with the given ID should be removed from all other groups
func (client Client) DeleteGroupEntitlement(groupId *uuid.UUID, ruleOption *string, removeGroupMembership *bool) (*GroupEntitlementOperationReference, error) {
    routeValues := make(map[string]string)
    if groupId == nil {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = (*groupId).String()

    queryParams := url.Values{}
    if ruleOption != nil {
        queryParams.Add("ruleOption", *ruleOption)
    }
    if removeGroupMembership != nil {
        queryParams.Add("removeGroupMembership", strconv.FormatBool(*removeGroupMembership))
    }
    locationId, _ := uuid.Parse("2280bffa-58a2-49da-822e-0764a1bb44f7")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GroupEntitlementOperationReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a group entitlement.
// groupId (required): ID of the group.
func (client Client) GetGroupEntitlement(groupId *uuid.UUID) (*GroupEntitlement, error) {
    routeValues := make(map[string]string)
    if groupId == nil {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = (*groupId).String()

    locationId, _ := uuid.Parse("2280bffa-58a2-49da-822e-0764a1bb44f7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GroupEntitlement
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the group entitlements for an account.
func (client Client) GetGroupEntitlements() (*[]GroupEntitlement, error) {
    locationId, _ := uuid.Parse("2280bffa-58a2-49da-822e-0764a1bb44f7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GroupEntitlement
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update entitlements (License Rule, Extensions Rule, Project memberships etc.) for a group.
// document (required): JsonPatchDocument containing the operations to perform on the group.
// groupId (required): ID of the group.
// ruleOption (optional): RuleOption [ApplyGroupRule/TestApplyGroupRule] - specifies if the rules defined in group entitlement should be updated and the changes are applied to it’s members (default option) or just be tested
func (client Client) UpdateGroupEntitlement(document *[]JsonPatchOperation, groupId *uuid.UUID, ruleOption *string) (*GroupEntitlementOperationReference, error) {
    if document == nil {
        return nil, errors.New("document is a required parameter")
    }
    routeValues := make(map[string]string)
    if groupId == nil {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = (*groupId).String()

    queryParams := url.Values{}
    if ruleOption != nil {
        queryParams.Add("ruleOption", *ruleOption)
    }
    body, marshalErr := json.Marshal(*document)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2280bffa-58a2-49da-822e-0764a1bb44f7")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GroupEntitlementOperationReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Add a member to a Group.
// groupId (required): Id of the Group.
// memberId (required): Id of the member to add.
func (client Client) AddMemberToGroup(groupId *uuid.UUID, memberId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if groupId == nil {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = (*groupId).String()
    if memberId == nil {
        return errors.New("memberId is a required parameter")
    }
    routeValues["memberId"] = (*memberId).String()

    locationId, _ := uuid.Parse("45a36e53-5286-4518-aa72-2d29f7acc5d8")
    _, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get direct members of a Group.
// groupId (required): Id of the Group.
// maxResults (optional): Maximum number of results to retrieve.
// pagingToken (optional): Paging Token from the previous page fetched. If the 'pagingToken' is null, the results would be fetched from the begining of the Members List.
func (client Client) GetGroupMembers(groupId *uuid.UUID, maxResults *int, pagingToken *string) (*PagedGraphMemberList, error) {
    routeValues := make(map[string]string)
    if groupId == nil {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = (*groupId).String()

    queryParams := url.Values{}
    if maxResults != nil {
        queryParams.Add("maxResults", strconv.Itoa(*maxResults))
    }
    if pagingToken != nil {
        queryParams.Add("pagingToken", *pagingToken)
    }
    locationId, _ := uuid.Parse("45a36e53-5286-4518-aa72-2d29f7acc5d8")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PagedGraphMemberList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Remove a member from a Group.
// groupId (required): Id of the group.
// memberId (required): Id of the member to remove.
func (client Client) RemoveMemberFromGroup(groupId *uuid.UUID, memberId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if groupId == nil {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = (*groupId).String()
    if memberId == nil {
        return errors.New("memberId is a required parameter")
    }
    routeValues["memberId"] = (*memberId).String()

    locationId, _ := uuid.Parse("45a36e53-5286-4518-aa72-2d29f7acc5d8")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Add a user, assign license and extensions and make them a member of a project group in an account.
// userEntitlement (required): UserEntitlement object specifying License, Extensions and Project/Team groups the user should be added to.
func (client Client) AddUserEntitlement(userEntitlement *UserEntitlement) (*UserEntitlementsPostResponse, error) {
    if userEntitlement == nil {
        return nil, errors.New("userEntitlement is a required parameter")
    }
    body, marshalErr := json.Marshal(*userEntitlement)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("387f832c-dbf2-4643-88e9-c1aa94dbb737")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UserEntitlementsPostResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a paged set of user entitlements matching the filter criteria. If no filter is is passed, a page from all the account users is returned.
// top (optional): Maximum number of the user entitlements to return. Max value is 10000. Default value is 100
// skip (optional): Offset: Number of records to skip. Default value is 0
// filter (optional): Comma (",") separated list of properties and their values to filter on. Currently, the API only supports filtering by ExtensionId. An example parameter would be filter=extensionId eq search.
// sortOption (optional): PropertyName and Order (separated by a space ( )) to sort on (e.g. LastAccessDate Desc)
func (client Client) GetUserEntitlements(top *int, skip *int, filter *string, sortOption *string) (*PagedGraphMemberList, error) {
    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("skip", strconv.Itoa(*skip))
    }
    if filter != nil {
        queryParams.Add("filter", *filter)
    }
    if sortOption != nil {
        queryParams.Add("sortOption", *sortOption)
    }
    locationId, _ := uuid.Parse("387f832c-dbf2-4643-88e9-c1aa94dbb737")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PagedGraphMemberList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Edit the entitlements (License, Extensions, Projects, Teams etc) for one or more users.
// document (required): JsonPatchDocument containing the operations to perform.
// doNotSendInviteForNewUsers (optional): Whether to send email invites to new users or not
func (client Client) UpdateUserEntitlements(document *[]JsonPatchOperation, doNotSendInviteForNewUsers *bool) (*UserEntitlementOperationReference, error) {
    if document == nil {
        return nil, errors.New("document is a required parameter")
    }
    queryParams := url.Values{}
    if doNotSendInviteForNewUsers != nil {
        queryParams.Add("doNotSendInviteForNewUsers", strconv.FormatBool(*doNotSendInviteForNewUsers))
    }
    body, marshalErr := json.Marshal(*document)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("387f832c-dbf2-4643-88e9-c1aa94dbb737")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", nil, queryParams, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UserEntitlementOperationReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a user from the account.
// userId (required): ID of the user.
func (client Client) DeleteUserEntitlement(userId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if userId == nil {
        return errors.New("userId is a required parameter")
    }
    routeValues["userId"] = (*userId).String()

    locationId, _ := uuid.Parse("8480c6eb-ce60-47e9-88df-eca3c801638b")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get User Entitlement for a user.
// userId (required): ID of the user.
func (client Client) GetUserEntitlement(userId *uuid.UUID) (*UserEntitlement, error) {
    routeValues := make(map[string]string)
    if userId == nil {
        return nil, errors.New("userId is a required parameter")
    }
    routeValues["userId"] = (*userId).String()

    locationId, _ := uuid.Parse("8480c6eb-ce60-47e9-88df-eca3c801638b")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UserEntitlement
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Edit the entitlements (License, Extensions, Projects, Teams etc) for a user.
// document (required): JsonPatchDocument containing the operations to perform on the user.
// userId (required): ID of the user.
func (client Client) UpdateUserEntitlement(document *[]JsonPatchOperation, userId *uuid.UUID) (*UserEntitlementsPatchResponse, error) {
    if document == nil {
        return nil, errors.New("document is a required parameter")
    }
    routeValues := make(map[string]string)
    if userId == nil {
        return nil, errors.New("userId is a required parameter")
    }
    routeValues["userId"] = (*userId).String()

    body, marshalErr := json.Marshal(*document)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("8480c6eb-ce60-47e9-88df-eca3c801638b")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UserEntitlementsPatchResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get summary of Licenses, Extension, Projects, Groups and their assignments in the collection.
// select_ (optional): Comma (",") separated list of properties to select. Supported property names are {AccessLevels, Licenses, Extensions, Projects, Groups}.
func (client Client) GetUsersSummary(select_ *string) (*UsersSummary, error) {
    queryParams := url.Values{}
    if select_ != nil {
        queryParams.Add("select_", *select_)
    }
    locationId, _ := uuid.Parse("5ae55b13-c9dd-49d1-957e-6e76c152e3d9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UsersSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

