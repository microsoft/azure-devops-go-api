// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package identity

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

var ResourceAreaId, _ = uuid.Parse("8a3d49b8-91f0-46ef-b33d-dda338c25db3")

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

// [Preview API]
// sourceIdentity (required)
func (client Client) CreateOrBindWithClaims(sourceIdentity *Identity) (*Identity, error) {
    if sourceIdentity == nil {
        return nil, errors.New("sourceIdentity is a required parameter")
    }
    body, marshalErr := json.Marshal(*sourceIdentity)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("90ddfe71-171c-446c-bf3b-b597cd562afd")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Identity
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// id (required)
// isMasterId (optional)
func (client Client) GetDescriptorById(id *uuid.UUID, isMasterId *bool) (*string, error) {
    routeValues := make(map[string]string)
    if id == nil {
        return nil, errors.New("id is a required parameter")
    }
    routeValues["id"] = (*id).String()

    queryParams := url.Values{}
    if isMasterId != nil {
        queryParams.Add("isMasterId", strconv.FormatBool(*isMasterId))
    }
    locationId, _ := uuid.Parse("a230389a-94f2-496c-839f-c929787496dd")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// container (required)
func (client Client) CreateGroups(container *interface{}) (*[]Identity, error) {
    if container == nil {
        return nil, errors.New("container is a required parameter")
    }
    body, marshalErr := json.Marshal(*container)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5966283b-4196-4d57-9211-1b68f41ec1c2")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// groupId (required)
func (client Client) DeleteGroup(groupId *string) error {
    routeValues := make(map[string]string)
    if groupId == nil || *groupId == "" {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId

    locationId, _ := uuid.Parse("5966283b-4196-4d57-9211-1b68f41ec1c2")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// scopeIds (optional)
// recurse (optional)
// deleted (optional)
// properties (optional)
func (client Client) ListGroups(scopeIds *string, recurse *bool, deleted *bool, properties *string) (*[]Identity, error) {
    queryParams := url.Values{}
    if scopeIds != nil {
        queryParams.Add("scopeIds", *scopeIds)
    }
    if recurse != nil {
        queryParams.Add("recurse", strconv.FormatBool(*recurse))
    }
    if deleted != nil {
        queryParams.Add("deleted", strconv.FormatBool(*deleted))
    }
    if properties != nil {
        queryParams.Add("properties", *properties)
    }
    locationId, _ := uuid.Parse("5966283b-4196-4d57-9211-1b68f41ec1c2")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// identitySequenceId (required)
// groupSequenceId (required)
// organizationIdentitySequenceId (optional)
// pageSize (optional)
// scopeId (optional)
func (client Client) GetIdentityChanges(identitySequenceId *int, groupSequenceId *int, organizationIdentitySequenceId *int, pageSize *int, scopeId *uuid.UUID) (*ChangedIdentities, error) {
    queryParams := url.Values{}
    if identitySequenceId == nil {
        return nil, errors.New("identitySequenceId is a required parameter")
    }
    queryParams.Add("identitySequenceId", strconv.Itoa(*identitySequenceId))
    if groupSequenceId == nil {
        return nil, errors.New("groupSequenceId is a required parameter")
    }
    queryParams.Add("groupSequenceId", strconv.Itoa(*groupSequenceId))
    if organizationIdentitySequenceId != nil {
        queryParams.Add("organizationIdentitySequenceId", strconv.Itoa(*organizationIdentitySequenceId))
    }
    if pageSize != nil {
        queryParams.Add("pageSize", strconv.Itoa(*pageSize))
    }
    if scopeId != nil {
        queryParams.Add("scopeId", (*scopeId).String())
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ChangedIdentities
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// domainId (required)
func (client Client) GetUserIdentityIdsByDomainId(domainId *uuid.UUID) (*[]uuid.UUID, error) {
    queryParams := url.Values{}
    if domainId == nil {
        return nil, errors.New("domainId is a required parameter")
    }
    queryParams.Add("domainId", (*domainId).String())
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []uuid.UUID
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// descriptors (optional)
// identityIds (optional)
// subjectDescriptors (optional)
// socialDescriptors (optional)
// searchFilter (optional)
// filterValue (optional)
// queryMembership (optional)
// properties (optional)
// includeRestrictedVisibility (optional)
// options (optional)
func (client Client) ReadIdentities(descriptors *string, identityIds *string, subjectDescriptors *string, socialDescriptors *string, searchFilter *string, filterValue *string, queryMembership *string, properties *string, includeRestrictedVisibility *bool, options *string) (*[]Identity, error) {
    queryParams := url.Values{}
    if descriptors != nil {
        queryParams.Add("descriptors", *descriptors)
    }
    if identityIds != nil {
        queryParams.Add("identityIds", *identityIds)
    }
    if subjectDescriptors != nil {
        queryParams.Add("subjectDescriptors", *subjectDescriptors)
    }
    if socialDescriptors != nil {
        queryParams.Add("socialDescriptors", *socialDescriptors)
    }
    if searchFilter != nil {
        queryParams.Add("searchFilter", *searchFilter)
    }
    if filterValue != nil {
        queryParams.Add("filterValue", *filterValue)
    }
    if queryMembership != nil {
        queryParams.Add("queryMembership", *queryMembership)
    }
    if properties != nil {
        queryParams.Add("properties", *properties)
    }
    if includeRestrictedVisibility != nil {
        queryParams.Add("includeRestrictedVisibility", strconv.FormatBool(*includeRestrictedVisibility))
    }
    if options != nil {
        queryParams.Add("options", *options)
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// scopeId (required)
// queryMembership (optional)
// properties (optional)
func (client Client) ReadIdentitiesByScope(scopeId *uuid.UUID, queryMembership *string, properties *string) (*[]Identity, error) {
    queryParams := url.Values{}
    if scopeId == nil {
        return nil, errors.New("scopeId is a required parameter")
    }
    queryParams.Add("scopeId", (*scopeId).String())
    if queryMembership != nil {
        queryParams.Add("queryMembership", *queryMembership)
    }
    if properties != nil {
        queryParams.Add("properties", *properties)
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// identityId (required)
// queryMembership (optional)
// properties (optional)
func (client Client) ReadIdentity(identityId *string, queryMembership *string, properties *string) (*Identity, error) {
    routeValues := make(map[string]string)
    if identityId == nil || *identityId == "" {
        return nil, errors.New("identityId is a required parameter")
    }
    routeValues["identityId"] = *identityId

    queryParams := url.Values{}
    if queryMembership != nil {
        queryParams.Add("queryMembership", *queryMembership)
    }
    if properties != nil {
        queryParams.Add("properties", *properties)
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Identity
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// identities (required)
func (client Client) UpdateIdentities(identities *azureDevops.VssJsonCollectionWrapper) (*[]IdentityUpdateData, error) {
    if identities == nil {
        return nil, errors.New("identities is a required parameter")
    }
    body, marshalErr := json.Marshal(*identities)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []IdentityUpdateData
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// identity (required)
// identityId (required)
func (client Client) UpdateIdentity(identity *Identity, identityId *uuid.UUID) error {
    if identity == nil {
        return errors.New("identity is a required parameter")
    }
    routeValues := make(map[string]string)
    if identityId == nil {
        return errors.New("identityId is a required parameter")
    }
    routeValues["identityId"] = (*identityId).String()

    body, marshalErr := json.Marshal(*identity)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("28010c54-d0c0-4c89-a5b0-1c9e188b9fb7")
    _, err := client.Client.Send(http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// frameworkIdentityInfo (required)
func (client Client) CreateIdentity(frameworkIdentityInfo *FrameworkIdentityInfo) (*Identity, error) {
    if frameworkIdentityInfo == nil {
        return nil, errors.New("frameworkIdentityInfo is a required parameter")
    }
    body, marshalErr := json.Marshal(*frameworkIdentityInfo)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("dd55f0eb-6ea2-4fe4-9ebe-919e7dd1dfb4")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Identity
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// batchInfo (required)
func (client Client) ReadIdentityBatch(batchInfo *IdentityBatchInfo) (*[]Identity, error) {
    if batchInfo == nil {
        return nil, errors.New("batchInfo is a required parameter")
    }
    body, marshalErr := json.Marshal(*batchInfo)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("299e50df-fe45-4d3a-8b5b-a5836fac74dc")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Identity
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// scopeId (required)
func (client Client) GetIdentitySnapshot(scopeId *string) (*IdentitySnapshot, error) {
    routeValues := make(map[string]string)
    if scopeId == nil || *scopeId == "" {
        return nil, errors.New("scopeId is a required parameter")
    }
    routeValues["scopeId"] = *scopeId

    locationId, _ := uuid.Parse("d56223df-8ccd-45c9-89b4-eddf692400d7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentitySnapshot
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Read the max sequence id of all the identities.
func (client Client) GetMaxSequenceId() (*uint64, error) {
    locationId, _ := uuid.Parse("e4a70778-cb2c-4e85-b7cc-3f3c7ae2d408")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue uint64
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Read identity of the home tenant request user.
func (client Client) GetSelf() (*IdentitySelf, error) {
    locationId, _ := uuid.Parse("4bb02b5b-c120-4be2-b68e-21f7c50a4b82")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentitySelf
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// containerId (required)
// memberId (required)
func (client Client) AddMember(containerId *string, memberId *string) (*bool, error) {
    routeValues := make(map[string]string)
    if containerId == nil || *containerId == "" {
        return nil, errors.New("containerId is a required parameter")
    }
    routeValues["containerId"] = *containerId
    if memberId == nil || *memberId == "" {
        return nil, errors.New("memberId is a required parameter")
    }
    routeValues["memberId"] = *memberId

    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// containerId (required)
// memberId (required)
// queryMembership (optional)
func (client Client) ReadMember(containerId *string, memberId *string, queryMembership *string) (*string, error) {
    routeValues := make(map[string]string)
    if containerId == nil || *containerId == "" {
        return nil, errors.New("containerId is a required parameter")
    }
    routeValues["containerId"] = *containerId
    if memberId == nil || *memberId == "" {
        return nil, errors.New("memberId is a required parameter")
    }
    routeValues["memberId"] = *memberId

    queryParams := url.Values{}
    if queryMembership != nil {
        queryParams.Add("queryMembership", *queryMembership)
    }
    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// containerId (required)
// queryMembership (optional)
func (client Client) ReadMembers(containerId *string, queryMembership *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if containerId == nil || *containerId == "" {
        return nil, errors.New("containerId is a required parameter")
    }
    routeValues["containerId"] = *containerId

    queryParams := url.Values{}
    if queryMembership != nil {
        queryParams.Add("queryMembership", *queryMembership)
    }
    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// containerId (required)
// memberId (required)
func (client Client) RemoveMember(containerId *string, memberId *string) (*bool, error) {
    routeValues := make(map[string]string)
    if containerId == nil || *containerId == "" {
        return nil, errors.New("containerId is a required parameter")
    }
    routeValues["containerId"] = *containerId
    if memberId == nil || *memberId == "" {
        return nil, errors.New("memberId is a required parameter")
    }
    routeValues["memberId"] = *memberId

    locationId, _ := uuid.Parse("8ba35978-138e-41f8-8963-7b1ea2c5f775")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// memberId (required)
// containerId (required)
// queryMembership (optional)
func (client Client) ReadMemberOf(memberId *string, containerId *string, queryMembership *string) (*string, error) {
    routeValues := make(map[string]string)
    if memberId == nil || *memberId == "" {
        return nil, errors.New("memberId is a required parameter")
    }
    routeValues["memberId"] = *memberId
    if containerId == nil || *containerId == "" {
        return nil, errors.New("containerId is a required parameter")
    }
    routeValues["containerId"] = *containerId

    queryParams := url.Values{}
    if queryMembership != nil {
        queryParams.Add("queryMembership", *queryMembership)
    }
    locationId, _ := uuid.Parse("22865b02-9e4a-479e-9e18-e35b8803b8a0")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// memberId (required)
// queryMembership (optional)
func (client Client) ReadMembersOf(memberId *string, queryMembership *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if memberId == nil || *memberId == "" {
        return nil, errors.New("memberId is a required parameter")
    }
    routeValues["memberId"] = *memberId

    queryParams := url.Values{}
    if queryMembership != nil {
        queryParams.Add("queryMembership", *queryMembership)
    }
    locationId, _ := uuid.Parse("22865b02-9e4a-479e-9e18-e35b8803b8a0")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// info (required)
// scopeId (required)
func (client Client) CreateScope(info *CreateScopeInfo, scopeId *uuid.UUID) (*IdentityScope, error) {
    if info == nil {
        return nil, errors.New("info is a required parameter")
    }
    routeValues := make(map[string]string)
    if scopeId == nil {
        return nil, errors.New("scopeId is a required parameter")
    }
    routeValues["scopeId"] = (*scopeId).String()

    body, marshalErr := json.Marshal(*info)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityScope
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// scopeId (required)
func (client Client) DeleteScope(scopeId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if scopeId == nil {
        return errors.New("scopeId is a required parameter")
    }
    routeValues["scopeId"] = (*scopeId).String()

    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// scopeId (required)
func (client Client) GetScopeById(scopeId *uuid.UUID) (*IdentityScope, error) {
    routeValues := make(map[string]string)
    if scopeId == nil {
        return nil, errors.New("scopeId is a required parameter")
    }
    routeValues["scopeId"] = (*scopeId).String()

    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityScope
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// scopeName (required)
func (client Client) GetScopeByName(scopeName *string) (*IdentityScope, error) {
    queryParams := url.Values{}
    if scopeName == nil {
        return nil, errors.New("scopeName is a required parameter")
    }
    queryParams.Add("scopeName", *scopeName)
    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityScope
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// patchDocument (required)
// scopeId (required)
func (client Client) UpdateScope(patchDocument *[]JsonPatchOperation, scopeId *uuid.UUID) error {
    if patchDocument == nil {
        return errors.New("patchDocument is a required parameter")
    }
    routeValues := make(map[string]string)
    if scopeId == nil {
        return errors.New("scopeId is a required parameter")
    }
    routeValues["scopeId"] = (*scopeId).String()

    body, marshalErr := json.Marshal(*patchDocument)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4e11e2bf-1e79-4eb5-8f34-a6337bd0de38")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
func (client Client) GetSignedInToken() (*AccessTokenResult, error) {
    locationId, _ := uuid.Parse("6074ff18-aaad-4abb-a41e-5c75f6178057")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AccessTokenResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
func (client Client) GetSignoutToken() (*AccessTokenResult, error) {
    locationId, _ := uuid.Parse("be39e83c-7529-45e9-9c67-0410885880da")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AccessTokenResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// tenantId (required)
func (client Client) GetTenant(tenantId *string) (*TenantInfo, error) {
    routeValues := make(map[string]string)
    if tenantId == nil || *tenantId == "" {
        return nil, errors.New("tenantId is a required parameter")
    }
    routeValues["tenantId"] = *tenantId

    locationId, _ := uuid.Parse("5f0a1723-2e2c-4c31-8cae-002d01bdd592")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TenantInfo
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

