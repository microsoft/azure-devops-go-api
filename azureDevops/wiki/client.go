// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package wiki

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

var ResourceAreaId, _ = uuid.Parse("bf7d82a0-8aa5-4613-94ef-6172a5ea01f3")

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

// Creates an attachment in the wiki.
// ctx
// uploadStream (required): Stream to upload
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// name (required): Wiki attachment name.
// versionDescriptor (optional): GitVersionDescriptor for the page. (Optional in case of ProjectWiki).
func (client Client) CreateAttachment(ctx context.Context, uploadStream *interface{}, project *string, wikiIdentifier *string, name *string, versionDescriptor *GitVersionDescriptor) (*WikiAttachmentResponse, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    queryParams := url.Values{}
    if name == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "name"}
    }
    queryParams.Add("name", *name)
    if versionDescriptor != nil {
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
        if versionDescriptor.VersionOptions != nil {
            queryParams.Add("versionDescriptor.versionOptions", string(*versionDescriptor.VersionOptions))
        }
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("c4382d8d-fefc-40e0-92c5-49852e9e17c0")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiAttachment
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiAttachmentResponse
    if err == nil {
        responseValue = &WikiAttachmentResponse {
            Attachment: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// Creates a page move operation that updates the path and order of the page as provided in the parameters.
// ctx
// pageMoveParameters (required): Page more operation parameters.
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// comment (optional): Comment that is to be associated with this page move.
// versionDescriptor (optional): GitVersionDescriptor for the page. (Optional in case of ProjectWiki).
func (client Client) CreatePageMove(ctx context.Context, pageMoveParameters *WikiPageMoveParameters, project *string, wikiIdentifier *string, comment *string, versionDescriptor *GitVersionDescriptor) (*WikiPageMoveResponse, error) {
    if pageMoveParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pageMoveParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    queryParams := url.Values{}
    if comment != nil {
        queryParams.Add("comment", *comment)
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
        if versionDescriptor.VersionOptions != nil {
            queryParams.Add("versionDescriptor.versionOptions", string(*versionDescriptor.VersionOptions))
        }
    }
    body, marshalErr := json.Marshal(*pageMoveParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e37bbe71-cbae-49e5-9a4e-949143b9d910")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiPageMove
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiPageMoveResponse
    if err == nil {
        responseValue = &WikiPageMoveResponse {
            PageMove: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// [Preview API] Deletes a wiki page.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// id (required): Wiki page id.
// comment (optional): Comment to be associated with this page delete.
func (client Client) DeletePageById(ctx context.Context, project *string, wikiIdentifier *string, id *int, comment *string) (*WikiPageResponse, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if comment != nil {
        queryParams.Add("comment", *comment)
    }
    locationId, _ := uuid.Parse("ceddcf75-1068-452d-8b13-2d4d76e1f970")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiPage
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiPageResponse
    if err == nil {
        responseValue = &WikiPageResponse {
            Page: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// [Preview API] Gets metadata or content of the wiki page for the provided page id. Content negotiation is done based on the `Accept` header sent in the request.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// id (required): Wiki page id.
// recursionLevel (optional): Recursion level for subpages retrieval. Defaults to `None` (Optional).
// includeContent (optional): True to include the content of the page in the response for Json content type. Defaults to false (Optional)
func (client Client) GetPageById(ctx context.Context, project *string, wikiIdentifier *string, id *int, recursionLevel *VersionControlRecursionType, includeContent *bool) (*WikiPageResponse, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("ceddcf75-1068-452d-8b13-2d4d76e1f970")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiPage
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiPageResponse
    if err == nil {
        responseValue = &WikiPageResponse {
            Page: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// [Preview API] Gets metadata or content of the wiki page for the provided page id. Content negotiation is done based on the `Accept` header sent in the request.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// id (required): Wiki page id.
// recursionLevel (optional): Recursion level for subpages retrieval. Defaults to `None` (Optional).
// includeContent (optional): True to include the content of the page in the response for Json content type. Defaults to false (Optional)
func (client Client) GetPageByIdText(ctx context.Context, project *string, wikiIdentifier *string, id *int, recursionLevel *VersionControlRecursionType, includeContent *bool) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("ceddcf75-1068-452d-8b13-2d4d76e1f970")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets metadata or content of the wiki page for the provided page id. Content negotiation is done based on the `Accept` header sent in the request.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// id (required): Wiki page id.
// recursionLevel (optional): Recursion level for subpages retrieval. Defaults to `None` (Optional).
// includeContent (optional): True to include the content of the page in the response for Json content type. Defaults to false (Optional)
func (client Client) GetPageByIdZip(ctx context.Context, project *string, wikiIdentifier *string, id *int, recursionLevel *VersionControlRecursionType, includeContent *bool) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("ceddcf75-1068-452d-8b13-2d4d76e1f970")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Edits a wiki page.
// ctx
// parameters (required): Wiki update operation parameters.
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// id (required): Wiki page id.
// version (required): Version of the page on which the change is to be made. Mandatory for `Edit` scenario. To be populated in the If-Match header of the request.
// comment (optional): Comment to be associated with the page operation.
func (client Client) UpdatePageById(ctx context.Context, parameters *WikiPageCreateOrUpdateParameters, project *string, wikiIdentifier *string, id *int, version *string, comment *string) (*WikiPageResponse, error) {
    if parameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "parameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if comment != nil {
        queryParams.Add("comment", *comment)
    }
    additionalHeaders := make(map[string]string)
    if version != nil {
        additionalHeaders["If-Match"] = *version
    }
    body, marshalErr := json.Marshal(*parameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ceddcf75-1068-452d-8b13-2d4d76e1f970")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiPage
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiPageResponse
    if err == nil {
        responseValue = &WikiPageResponse {
            Page: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// Creates or edits a wiki page.
// ctx
// parameters (required): Wiki create or update operation parameters.
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// path (required): Wiki page path.
// version (required): Version of the page on which the change is to be made. Mandatory for `Edit` scenario. To be populated in the If-Match header of the request.
// comment (optional): Comment to be associated with the page operation.
// versionDescriptor (optional): GitVersionDescriptor for the page. (Optional in case of ProjectWiki).
func (client Client) CreateOrUpdatePage(ctx context.Context, parameters *WikiPageCreateOrUpdateParameters, project *string, wikiIdentifier *string, path *string, version *string, comment *string, versionDescriptor *GitVersionDescriptor) (*WikiPageResponse, error) {
    if parameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "parameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if comment != nil {
        queryParams.Add("comment", *comment)
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
        if versionDescriptor.VersionOptions != nil {
            queryParams.Add("versionDescriptor.versionOptions", string(*versionDescriptor.VersionOptions))
        }
    }
    additionalHeaders := make(map[string]string)
    if version != nil {
        additionalHeaders["If-Match"] = *version
    }
    body, marshalErr := json.Marshal(*parameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("25d3fbc7-fe3d-46cb-b5a5-0b6f79caf27b")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiPage
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiPageResponse
    if err == nil {
        responseValue = &WikiPageResponse {
            Page: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// Deletes a wiki page.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// path (required): Wiki page path.
// comment (optional): Comment to be associated with this page delete.
// versionDescriptor (optional): GitVersionDescriptor for the page. (Optional in case of ProjectWiki).
func (client Client) DeletePage(ctx context.Context, project *string, wikiIdentifier *string, path *string, comment *string, versionDescriptor *GitVersionDescriptor) (*WikiPageResponse, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if comment != nil {
        queryParams.Add("comment", *comment)
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
        if versionDescriptor.VersionOptions != nil {
            queryParams.Add("versionDescriptor.versionOptions", string(*versionDescriptor.VersionOptions))
        }
    }
    locationId, _ := uuid.Parse("25d3fbc7-fe3d-46cb-b5a5-0b6f79caf27b")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiPage
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiPageResponse
    if err == nil {
        responseValue = &WikiPageResponse {
            Page: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// Gets metadata or content of the wiki page for the provided path. Content negotiation is done based on the `Accept` header sent in the request.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// path (optional): Wiki page path.
// recursionLevel (optional): Recursion level for subpages retrieval. Defaults to `None` (Optional).
// versionDescriptor (optional): GitVersionDescriptor for the page. Defaults to the default branch (Optional).
// includeContent (optional): True to include the content of the page in the response for Json content type. Defaults to false (Optional)
func (client Client) GetPage(ctx context.Context, project *string, wikiIdentifier *string, path *string, recursionLevel *VersionControlRecursionType, versionDescriptor *GitVersionDescriptor, includeContent *bool) (*WikiPageResponse, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    queryParams := url.Values{}
    if path != nil {
        queryParams.Add("path", *path)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
        if versionDescriptor.VersionOptions != nil {
            queryParams.Add("versionDescriptor.versionOptions", string(*versionDescriptor.VersionOptions))
        }
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("25d3fbc7-fe3d-46cb-b5a5-0b6f79caf27b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue WikiPage
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WikiPageResponse
    if err == nil {
        responseValue = &WikiPageResponse {
            Page: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// Gets metadata or content of the wiki page for the provided path. Content negotiation is done based on the `Accept` header sent in the request.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// path (optional): Wiki page path.
// recursionLevel (optional): Recursion level for subpages retrieval. Defaults to `None` (Optional).
// versionDescriptor (optional): GitVersionDescriptor for the page. Defaults to the default branch (Optional).
// includeContent (optional): True to include the content of the page in the response for Json content type. Defaults to false (Optional)
func (client Client) GetPageText(ctx context.Context, project *string, wikiIdentifier *string, path *string, recursionLevel *VersionControlRecursionType, versionDescriptor *GitVersionDescriptor, includeContent *bool) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    queryParams := url.Values{}
    if path != nil {
        queryParams.Add("path", *path)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
        if versionDescriptor.VersionOptions != nil {
            queryParams.Add("versionDescriptor.versionOptions", string(*versionDescriptor.VersionOptions))
        }
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("25d3fbc7-fe3d-46cb-b5a5-0b6f79caf27b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets metadata or content of the wiki page for the provided path. Content negotiation is done based on the `Accept` header sent in the request.
// ctx
// project (required): Project ID or project name
// wikiIdentifier (required): Wiki Id or name.
// path (optional): Wiki page path.
// recursionLevel (optional): Recursion level for subpages retrieval. Defaults to `None` (Optional).
// versionDescriptor (optional): GitVersionDescriptor for the page. Defaults to the default branch (Optional).
// includeContent (optional): True to include the content of the page in the response for Json content type. Defaults to false (Optional)
func (client Client) GetPageZip(ctx context.Context, project *string, wikiIdentifier *string, path *string, recursionLevel *VersionControlRecursionType, versionDescriptor *GitVersionDescriptor, includeContent *bool) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    queryParams := url.Values{}
    if path != nil {
        queryParams.Add("path", *path)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
        if versionDescriptor.VersionOptions != nil {
            queryParams.Add("versionDescriptor.versionOptions", string(*versionDescriptor.VersionOptions))
        }
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("25d3fbc7-fe3d-46cb-b5a5-0b6f79caf27b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Creates the wiki resource.
// ctx
// wikiCreateParams (required): Parameters for the wiki creation.
// project (optional): Project ID or project name
func (client Client) CreateWiki(ctx context.Context, wikiCreateParams *WikiCreateParametersV2, project *string) (*WikiV2, error) {
    if wikiCreateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "wikiCreateParams"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*wikiCreateParams)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("288d122c-dbd4-451d-aa5f-7dbbba070728")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WikiV2
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Deletes the wiki corresponding to the wiki name or Id provided.
// ctx
// wikiIdentifier (required): Wiki name or Id.
// project (optional): Project ID or project name
func (client Client) DeleteWiki(ctx context.Context, wikiIdentifier *string, project *string) (*WikiV2, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    locationId, _ := uuid.Parse("288d122c-dbd4-451d-aa5f-7dbbba070728")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WikiV2
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets all wikis in a project or collection.
// ctx
// project (optional): Project ID or project name
func (client Client) GetAllWikis(ctx context.Context, project *string) (*[]WikiV2, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    locationId, _ := uuid.Parse("288d122c-dbd4-451d-aa5f-7dbbba070728")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WikiV2
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the wiki corresponding to the wiki name or Id provided.
// ctx
// wikiIdentifier (required): Wiki name or id.
// project (optional): Project ID or project name
func (client Client) GetWiki(ctx context.Context, wikiIdentifier *string, project *string) (*WikiV2, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    locationId, _ := uuid.Parse("288d122c-dbd4-451d-aa5f-7dbbba070728")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WikiV2
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Updates the wiki corresponding to the wiki Id or name provided using the update parameters.
// ctx
// updateParameters (required): Update parameters.
// wikiIdentifier (required): Wiki name or Id.
// project (optional): Project ID or project name
func (client Client) UpdateWiki(ctx context.Context, updateParameters *WikiUpdateParameters, wikiIdentifier *string, project *string) (*WikiV2, error) {
    if updateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if wikiIdentifier == nil || *wikiIdentifier == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "wikiIdentifier"} 
    }
    routeValues["wikiIdentifier"] = *wikiIdentifier

    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("288d122c-dbd4-451d-aa5f-7dbbba070728")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WikiV2
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

