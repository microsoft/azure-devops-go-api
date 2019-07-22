// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package security

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// Remove the specified ACEs from the ACL belonging to the specified token.
// ctx
// securityNamespaceId (required): Security namespace identifier.
// token (optional): The token whose ACL should be modified.
// descriptors (optional): String containing a list of identity descriptors separated by ',' whose entries should be removed.
func (client Client) RemoveAccessControlEntries(ctx context.Context, securityNamespaceId *uuid.UUID, token *string, descriptors *string) (*bool, error) {
    routeValues := make(map[string]string)
    if securityNamespaceId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "securityNamespaceId"} 
    }
    routeValues["securityNamespaceId"] = (*securityNamespaceId).String()

    queryParams := url.Values{}
    if token != nil {
        queryParams.Add("token", *token)
    }
    if descriptors != nil {
        queryParams.Add("descriptors", *descriptors)
    }
    locationId, _ := uuid.Parse("ac08c8ff-4323-4b08-af90-bcd018d380ce")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Add or update ACEs in the ACL for the provided token. The request body contains the target token, a list of [ACEs](https://docs.microsoft.com/en-us/rest/api/azure/devops/security/access%20control%20entries/set%20access%20control%20entries?#accesscontrolentry) and a optional merge parameter. In the case of a collision (by identity descriptor) with an existing ACE in the ACL, the "merge" parameter determines the behavior. If set, the existing ACE has its allow and deny merged with the incoming ACE's allow and deny. If unset, the existing ACE is displaced.
// ctx
// container (required)
// securityNamespaceId (required): Security namespace identifier.
func (client Client) SetAccessControlEntries(ctx context.Context, container *interface{}, securityNamespaceId *uuid.UUID) (*[]AccessControlEntry, error) {
    if container == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "container"}
    }
    routeValues := make(map[string]string)
    if securityNamespaceId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "securityNamespaceId"} 
    }
    routeValues["securityNamespaceId"] = (*securityNamespaceId).String()

    body, marshalErr := json.Marshal(*container)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ac08c8ff-4323-4b08-af90-bcd018d380ce")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AccessControlEntry
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Return a list of access control lists for the specified security namespace and token. All ACLs in the security namespace will be retrieved if no optional parameters are provided.
// ctx
// securityNamespaceId (required): Security namespace identifier.
// token (optional): Security token
// descriptors (optional): An optional filter string containing a list of identity descriptors separated by ',' whose ACEs should be retrieved. If this is left null, entire ACLs will be returned.
// includeExtendedInfo (optional): If true, populate the extended information properties for the access control entries contained in the returned lists.
// recurse (optional): If true and this is a hierarchical namespace, return child ACLs of the specified token.
func (client Client) QueryAccessControlLists(ctx context.Context, securityNamespaceId *uuid.UUID, token *string, descriptors *string, includeExtendedInfo *bool, recurse *bool) (*[]AccessControlList, error) {
    routeValues := make(map[string]string)
    if securityNamespaceId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "securityNamespaceId"} 
    }
    routeValues["securityNamespaceId"] = (*securityNamespaceId).String()

    queryParams := url.Values{}
    if token != nil {
        queryParams.Add("token", *token)
    }
    if descriptors != nil {
        queryParams.Add("descriptors", *descriptors)
    }
    if includeExtendedInfo != nil {
        queryParams.Add("includeExtendedInfo", strconv.FormatBool(*includeExtendedInfo))
    }
    if recurse != nil {
        queryParams.Add("recurse", strconv.FormatBool(*recurse))
    }
    locationId, _ := uuid.Parse("18a2ad18-7571-46ae-bec7-0c7da1495885")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AccessControlList
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Remove access control lists under the specfied security namespace.
// ctx
// securityNamespaceId (required): Security namespace identifier.
// tokens (optional): One or more comma-separated security tokens
// recurse (optional): If true and this is a hierarchical namespace, also remove child ACLs of the specified tokens.
func (client Client) RemoveAccessControlLists(ctx context.Context, securityNamespaceId *uuid.UUID, tokens *string, recurse *bool) (*bool, error) {
    routeValues := make(map[string]string)
    if securityNamespaceId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "securityNamespaceId"} 
    }
    routeValues["securityNamespaceId"] = (*securityNamespaceId).String()

    queryParams := url.Values{}
    if tokens != nil {
        queryParams.Add("tokens", *tokens)
    }
    if recurse != nil {
        queryParams.Add("recurse", strconv.FormatBool(*recurse))
    }
    locationId, _ := uuid.Parse("18a2ad18-7571-46ae-bec7-0c7da1495885")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Create or update one or more access control lists. All data that currently exists for the ACLs supplied will be overwritten.
// ctx
// accessControlLists (required): A list of ACLs to create or update.
// securityNamespaceId (required): Security namespace identifier.
func (client Client) SetAccessControlLists(ctx context.Context, accessControlLists *azureDevops.VssJsonCollectionWrapper, securityNamespaceId *uuid.UUID) error {
    if accessControlLists == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "accessControlLists"}
    }
    routeValues := make(map[string]string)
    if securityNamespaceId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "securityNamespaceId"} 
    }
    routeValues["securityNamespaceId"] = (*securityNamespaceId).String()

    body, marshalErr := json.Marshal(*accessControlLists)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("18a2ad18-7571-46ae-bec7-0c7da1495885")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Evaluates multiple permissions for the calling user.  Note: This method does not aggregate the results, nor does it short-circuit if one of the permissions evaluates to false.
// ctx
// evalBatch (required): The set of evaluation requests.
func (client Client) HasPermissionsBatch(ctx context.Context, evalBatch *PermissionEvaluationBatch) (*PermissionEvaluationBatch, error) {
    if evalBatch == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "evalBatch"}
    }
    body, marshalErr := json.Marshal(*evalBatch)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cf1faa59-1b63-4448-bf04-13d981a46f5d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PermissionEvaluationBatch
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Evaluates whether the caller has the specified permissions on the specified set of security tokens.
// ctx
// securityNamespaceId (required): Security namespace identifier.
// permissions (optional): Permissions to evaluate.
// tokens (optional): One or more security tokens to evaluate.
// alwaysAllowAdministrators (optional): If true and if the caller is an administrator, always return true.
// delimiter (optional): Optional security token separator. Defaults to ",".
func (client Client) HasPermissions(ctx context.Context, securityNamespaceId *uuid.UUID, permissions *int, tokens *string, alwaysAllowAdministrators *bool, delimiter *string) (*[]bool, error) {
    routeValues := make(map[string]string)
    if securityNamespaceId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "securityNamespaceId"} 
    }
    routeValues["securityNamespaceId"] = (*securityNamespaceId).String()
    if permissions != nil {
        routeValues["permissions"] = strconv.Itoa(*permissions)
    }

    queryParams := url.Values{}
    if tokens != nil {
        queryParams.Add("tokens", *tokens)
    }
    if alwaysAllowAdministrators != nil {
        queryParams.Add("alwaysAllowAdministrators", strconv.FormatBool(*alwaysAllowAdministrators))
    }
    if delimiter != nil {
        queryParams.Add("delimiter", *delimiter)
    }
    locationId, _ := uuid.Parse("dd3b8bd6-c7fc-4cbd-929a-933d9c011c9d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []bool
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Removes the specified permissions on a security token for a user or group.
// ctx
// securityNamespaceId (required): Security namespace identifier.
// descriptor (required): Identity descriptor of the user to remove permissions for.
// permissions (optional): Permissions to remove.
// token (optional): Security token to remove permissions for.
func (client Client) RemovePermission(ctx context.Context, securityNamespaceId *uuid.UUID, descriptor *string, permissions *int, token *string) (*AccessControlEntry, error) {
    routeValues := make(map[string]string)
    if securityNamespaceId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "securityNamespaceId"} 
    }
    routeValues["securityNamespaceId"] = (*securityNamespaceId).String()
    if permissions != nil {
        routeValues["permissions"] = strconv.Itoa(*permissions)
    }

    queryParams := url.Values{}
    if descriptor == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "descriptor"}
    }
    queryParams.Add("descriptor", *descriptor)
    if token != nil {
        queryParams.Add("token", *token)
    }
    locationId, _ := uuid.Parse("dd3b8bd6-c7fc-4cbd-929a-933d9c011c9d")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AccessControlEntry
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// List all security namespaces or just the specified namespace.
// ctx
// securityNamespaceId (optional): Security namespace identifier.
// localOnly (optional): If true, retrieve only local security namespaces.
func (client Client) QuerySecurityNamespaces(ctx context.Context, securityNamespaceId *uuid.UUID, localOnly *bool) (*[]SecurityNamespaceDescription, error) {
    routeValues := make(map[string]string)
    if securityNamespaceId != nil {
        routeValues["securityNamespaceId"] = (*securityNamespaceId).String()
    }

    queryParams := url.Values{}
    if localOnly != nil {
        queryParams.Add("localOnly", strconv.FormatBool(*localOnly))
    }
    locationId, _ := uuid.Parse("ce7b9f95-fde9-4be8-a86d-83b366f0b87a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SecurityNamespaceDescription
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

