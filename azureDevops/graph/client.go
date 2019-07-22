// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package graph

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

var ResourceAreaId, _ = uuid.Parse("bb1e7ec9-e901-4b68-999a-de7012b920f8")

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
// subjectDescriptor (required)
func (client Client) DeleteAvatar(ctx context.Context, subjectDescriptor *string) error {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor

    locationId, _ := uuid.Parse("801eaf9c-0585-4be8-9cdb-b0efa074de91")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// subjectDescriptor (required)
// size (optional)
// format (optional)
func (client Client) GetAvatar(ctx context.Context, subjectDescriptor *string, size *AvatarSize, format *string) (*Avatar, error) {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor

    queryParams := url.Values{}
    if size != nil {
        queryParams.Add("size", string(*size))
    }
    if format != nil {
        queryParams.Add("format", *format)
    }
    locationId, _ := uuid.Parse("801eaf9c-0585-4be8-9cdb-b0efa074de91")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Avatar
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// avatar (required)
// subjectDescriptor (required)
func (client Client) SetAvatar(ctx context.Context, avatar *Avatar, subjectDescriptor *string) error {
    if avatar == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "avatar"}
    }
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor

    body, marshalErr := json.Marshal(*avatar)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("801eaf9c-0585-4be8-9cdb-b0efa074de91")
    _, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Resolve a storage key to a descriptor
// ctx
// storageKey (required): Storage key of the subject (user, group, scope, etc.) to resolve
func (client Client) GetDescriptor(ctx context.Context, storageKey *uuid.UUID) (*GraphDescriptorResult, error) {
    routeValues := make(map[string]string)
    if storageKey == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "storageKey"} 
    }
    routeValues["storageKey"] = (*storageKey).String()

    locationId, _ := uuid.Parse("048aee0a-7072-4cde-ab73-7af77b1e0b4e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphDescriptorResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a new Azure DevOps group or materialize an existing AAD group.
// ctx
// creationContext (required): The subset of the full graph group used to uniquely find the graph subject in an external provider.
// scopeDescriptor (optional): A descriptor referencing the scope (collection, project) in which the group should be created. If omitted, will be created in the scope of the enclosing account or organization. Valid only for VSTS groups.
// groupDescriptors (optional): A comma separated list of descriptors referencing groups you want the graph group to join
func (client Client) CreateGroup(ctx context.Context, creationContext *GraphGroupCreationContext, scopeDescriptor *string, groupDescriptors *[]string) (*GraphGroup, error) {
    if creationContext == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "creationContext"}
    }
    queryParams := url.Values{}
    if scopeDescriptor != nil {
        queryParams.Add("scopeDescriptor", *scopeDescriptor)
    }
    if groupDescriptors != nil {
        listAsString := strings.Join((*groupDescriptors)[:], ",")
        queryParams.Add("groupDescriptors", listAsString)
    }
    body, marshalErr := json.Marshal(*creationContext)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ebbe6af8-0b91-4c13-8cf1-777c14858188")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes an Azure DevOps group from all of its parent groups.
// ctx
// groupDescriptor (required): The descriptor of the group to delete.
func (client Client) DeleteGroup(ctx context.Context, groupDescriptor *string) error {
    routeValues := make(map[string]string)
    if groupDescriptor == nil || *groupDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupDescriptor"} 
    }
    routeValues["groupDescriptor"] = *groupDescriptor

    locationId, _ := uuid.Parse("ebbe6af8-0b91-4c13-8cf1-777c14858188")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a group by its descriptor.
// ctx
// groupDescriptor (required): The descriptor of the desired graph group.
func (client Client) GetGroup(ctx context.Context, groupDescriptor *string) (*GraphGroup, error) {
    routeValues := make(map[string]string)
    if groupDescriptor == nil || *groupDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupDescriptor"} 
    }
    routeValues["groupDescriptor"] = *groupDescriptor

    locationId, _ := uuid.Parse("ebbe6af8-0b91-4c13-8cf1-777c14858188")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a list of all groups in the current scope (usually organization or account).
// ctx
// scopeDescriptor (optional): Specify a non-default scope (collection, project) to search for groups.
// subjectTypes (optional): A comma separated list of user subject subtypes to reduce the retrieved results, e.g. Microsoft.IdentityModel.Claims.ClaimsIdentity
// continuationToken (optional): An opaque data blob that allows the next page of data to resume immediately after where the previous page ended. The only reliable way to know if there is more data left is the presence of a continuation token.
func (client Client) ListGroups(ctx context.Context, scopeDescriptor *string, subjectTypes *[]string, continuationToken *string) (*PagedGraphGroups, error) {
    queryParams := url.Values{}
    if scopeDescriptor != nil {
        queryParams.Add("scopeDescriptor", *scopeDescriptor)
    }
    if subjectTypes != nil {
        listAsString := strings.Join((*subjectTypes)[:], ",")
        queryParams.Add("subjectTypes", listAsString)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("ebbe6af8-0b91-4c13-8cf1-777c14858188")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue []GraphGroup
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *PagedGraphGroups
    if err == nil {
        responseValue = &PagedGraphGroups {
            GraphGroups: &responseBodyValue,
            ContinuationToken: &[]string{ resp.Header.Get("X-MS-ContinuationToken") },
        }
    }

    return responseValue, err
}

// [Preview API] Update the properties of an Azure DevOps group.
// ctx
// groupDescriptor (required): The descriptor of the group to modify.
// patchDocument (required): The JSON+Patch document containing the fields to alter.
func (client Client) UpdateGroup(ctx context.Context, groupDescriptor *string, patchDocument *[]JsonPatchOperation) (*GraphGroup, error) {
    if patchDocument == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "patchDocument"}
    }
    routeValues := make(map[string]string)
    if groupDescriptor == nil || *groupDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "groupDescriptor"} 
    }
    routeValues["groupDescriptor"] = *groupDescriptor

    body, marshalErr := json.Marshal(*patchDocument)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ebbe6af8-0b91-4c13-8cf1-777c14858188")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a new membership between a container and subject.
// ctx
// subjectDescriptor (required): A descriptor to a group or user that can be the child subject in the relationship.
// containerDescriptor (required): A descriptor to a group that can be the container in the relationship.
func (client Client) AddMembership(ctx context.Context, subjectDescriptor *string, containerDescriptor *string) (*GraphMembership, error) {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor
    if containerDescriptor == nil || *containerDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "containerDescriptor"} 
    }
    routeValues["containerDescriptor"] = *containerDescriptor

    locationId, _ := uuid.Parse("3fd2e6ca-fb30-443a-b579-95b19ed0934c")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphMembership
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Check to see if a membership relationship between a container and subject exists.
// ctx
// subjectDescriptor (required): The group or user that is a child subject of the relationship.
// containerDescriptor (required): The group that is the container in the relationship.
func (client Client) CheckMembershipExistence(ctx context.Context, subjectDescriptor *string, containerDescriptor *string) error {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor
    if containerDescriptor == nil || *containerDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "containerDescriptor"} 
    }
    routeValues["containerDescriptor"] = *containerDescriptor

    locationId, _ := uuid.Parse("3fd2e6ca-fb30-443a-b579-95b19ed0934c")
    _, err := client.Client.Send(ctx, http.MethodHead, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a membership relationship between a container and subject.
// ctx
// subjectDescriptor (required): A descriptor to the child subject in the relationship.
// containerDescriptor (required): A descriptor to the container in the relationship.
func (client Client) GetMembership(ctx context.Context, subjectDescriptor *string, containerDescriptor *string) (*GraphMembership, error) {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor
    if containerDescriptor == nil || *containerDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "containerDescriptor"} 
    }
    routeValues["containerDescriptor"] = *containerDescriptor

    locationId, _ := uuid.Parse("3fd2e6ca-fb30-443a-b579-95b19ed0934c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphMembership
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes a membership between a container and subject.
// ctx
// subjectDescriptor (required): A descriptor to a group or user that is the child subject in the relationship.
// containerDescriptor (required): A descriptor to a group that is the container in the relationship.
func (client Client) RemoveMembership(ctx context.Context, subjectDescriptor *string, containerDescriptor *string) error {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor
    if containerDescriptor == nil || *containerDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "containerDescriptor"} 
    }
    routeValues["containerDescriptor"] = *containerDescriptor

    locationId, _ := uuid.Parse("3fd2e6ca-fb30-443a-b579-95b19ed0934c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get all the memberships where this descriptor is a member in the relationship.
// ctx
// subjectDescriptor (required): Fetch all direct memberships of this descriptor.
// direction (optional): Defaults to Up.
// depth (optional): The maximum number of edges to traverse up or down the membership tree. Currently the only supported value is '1'.
func (client Client) ListMemberships(ctx context.Context, subjectDescriptor *string, direction *GraphTraversalDirection, depth *int) (*[]GraphMembership, error) {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor

    queryParams := url.Values{}
    if direction != nil {
        queryParams.Add("direction", string(*direction))
    }
    if depth != nil {
        queryParams.Add("depth", strconv.Itoa(*depth))
    }
    locationId, _ := uuid.Parse("e34b6394-6b30-4435-94a9-409a5eef3e31")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GraphMembership
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Check whether a subject is active or inactive.
// ctx
// subjectDescriptor (required): Descriptor of the subject (user, group, scope, etc.) to check state of
func (client Client) GetMembershipState(ctx context.Context, subjectDescriptor *string) (*GraphMembershipState, error) {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor

    locationId, _ := uuid.Parse("1ffe5c94-1144-4191-907b-d0211cad36a8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphMembershipState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// userDescriptor (required)
func (client Client) GetProviderInfo(ctx context.Context, userDescriptor *string) (*GraphProviderInfo, error) {
    routeValues := make(map[string]string)
    if userDescriptor == nil || *userDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userDescriptor"} 
    }
    routeValues["userDescriptor"] = *userDescriptor

    locationId, _ := uuid.Parse("1e377995-6fa2-4588-bd64-930186abdcfa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphProviderInfo
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// message (required)
func (client Client) RequestAccess(ctx context.Context, message *string) error {
    if message == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "message"}
    }
    body, marshalErr := json.Marshal(*message)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("8d54bf92-8c99-47f2-9972-b21341f1722e")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Resolve a descriptor to a storage key.
// ctx
// subjectDescriptor (required)
func (client Client) GetStorageKey(ctx context.Context, subjectDescriptor *string) (*GraphStorageKeyResult, error) {
    routeValues := make(map[string]string)
    if subjectDescriptor == nil || *subjectDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subjectDescriptor"} 
    }
    routeValues["subjectDescriptor"] = *subjectDescriptor

    locationId, _ := uuid.Parse("eb85f8cc-f0f6-4264-a5b1-ffe2e4d4801f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphStorageKeyResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Resolve descriptors to users, groups or scopes (Subjects) in a batch.
// ctx
// subjectLookup (required): A list of descriptors that specifies a subset of subjects to retrieve. Each descriptor uniquely identifies the subject across all instance scopes, but only at a single point in time.
func (client Client) LookupSubjects(ctx context.Context, subjectLookup *GraphSubjectLookup) (*map[string]GraphSubject, error) {
    if subjectLookup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subjectLookup"}
    }
    body, marshalErr := json.Marshal(*subjectLookup)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4dd4d168-11f2-48c4-83e8-756fa0de027c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]GraphSubject
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Materialize an existing AAD or MSA user into the VSTS account.
// ctx
// creationContext (required): The subset of the full graph user used to uniquely find the graph subject in an external provider.
// groupDescriptors (optional): A comma separated list of descriptors of groups you want the graph user to join
func (client Client) CreateUser(ctx context.Context, creationContext *GraphUserCreationContext, groupDescriptors *[]string) (*GraphUser, error) {
    if creationContext == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "creationContext"}
    }
    queryParams := url.Values{}
    if groupDescriptors != nil {
        listAsString := strings.Join((*groupDescriptors)[:], ",")
        queryParams.Add("groupDescriptors", listAsString)
    }
    body, marshalErr := json.Marshal(*creationContext)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("005e26ec-6b77-4e4f-a986-b3827bf241f5")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphUser
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Disables a user.
// ctx
// userDescriptor (required): The descriptor of the user to delete.
func (client Client) DeleteUser(ctx context.Context, userDescriptor *string) error {
    routeValues := make(map[string]string)
    if userDescriptor == nil || *userDescriptor == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userDescriptor"} 
    }
    routeValues["userDescriptor"] = *userDescriptor

    locationId, _ := uuid.Parse("005e26ec-6b77-4e4f-a986-b3827bf241f5")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a user by its descriptor.
// ctx
// userDescriptor (required): The descriptor of the desired user.
func (client Client) GetUser(ctx context.Context, userDescriptor *string) (*GraphUser, error) {
    routeValues := make(map[string]string)
    if userDescriptor == nil || *userDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userDescriptor"} 
    }
    routeValues["userDescriptor"] = *userDescriptor

    locationId, _ := uuid.Parse("005e26ec-6b77-4e4f-a986-b3827bf241f5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphUser
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of all users in a given scope.
// ctx
// subjectTypes (optional): A comma separated list of user subject subtypes to reduce the retrieved results, e.g. msa’, ‘aad’, ‘svc’ (service identity), ‘imp’ (imported identity), etc.
// continuationToken (optional): An opaque data blob that allows the next page of data to resume immediately after where the previous page ended. The only reliable way to know if there is more data left is the presence of a continuation token.
func (client Client) ListUsers(ctx context.Context, subjectTypes *[]string, continuationToken *string) (*PagedGraphUsers, error) {
    queryParams := url.Values{}
    if subjectTypes != nil {
        listAsString := strings.Join((*subjectTypes)[:], ",")
        queryParams.Add("subjectTypes", listAsString)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("005e26ec-6b77-4e4f-a986-b3827bf241f5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue []GraphUser
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *PagedGraphUsers
    if err == nil {
        responseValue = &PagedGraphUsers {
            GraphUsers: &responseBodyValue,
            ContinuationToken: &[]string{ resp.Header.Get("X-MS-ContinuationToken") },
        }
    }

    return responseValue, err
}

// [Preview API] Map an existing user to a different identity
// ctx
// updateContext (required): The subset of the full graph user used to uniquely find the graph subject in an external provider.
// userDescriptor (required): the descriptor of the user to update
func (client Client) UpdateUser(ctx context.Context, updateContext *GraphUserUpdateContext, userDescriptor *string) (*GraphUser, error) {
    if updateContext == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateContext"}
    }
    routeValues := make(map[string]string)
    if userDescriptor == nil || *userDescriptor == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userDescriptor"} 
    }
    routeValues["userDescriptor"] = *userDescriptor

    body, marshalErr := json.Marshal(*updateContext)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("005e26ec-6b77-4e4f-a986-b3827bf241f5")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GraphUser
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

