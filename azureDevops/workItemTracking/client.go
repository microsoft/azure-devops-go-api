// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package workItemTracking

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

// [Preview API] Gets recent work item activities
// ctx
func (client Client) GetRecentActivityData(ctx context.Context, ) (*[]AccountRecentActivityWorkItemModel2, error) {
    locationId, _ := uuid.Parse("1bc988f4-c15f-4072-ad35-497c87e3a909")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AccountRecentActivityWorkItemModel2
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the list of work item tracking outbound artifact link types.
// ctx
func (client Client) GetWorkArtifactLinkTypes(ctx context.Context, ) (*[]WorkArtifactLink, error) {
    locationId, _ := uuid.Parse("1a31de40-e318-41cd-a6c6-881077df52e3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkArtifactLink
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Queries work items linked to a given list of artifact URI.
// ctx
// artifactUriQuery (required): Defines a list of artifact URI for querying work items.
// project (optional): Project ID or project name
func (client Client) QueryWorkItemsForArtifactUris(ctx context.Context, artifactUriQuery *ArtifactUriQuery, project *string) (*ArtifactUriQueryResult, error) {
    if artifactUriQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "artifactUriQuery"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*artifactUriQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a9a9aa7a-8c09-44d3-ad1b-46e855c1e3d3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ArtifactUriQueryResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Uploads an attachment.
// ctx
// uploadStream (required): Stream to upload
// project (optional): Project ID or project name
// fileName (optional): The name of the file
// uploadType (optional): Attachment upload type: Simple or Chunked
// areaPath (optional): Target project Area Path
func (client Client) CreateAttachment(ctx context.Context, uploadStream interface{}, project *string, fileName *string, uploadType *string, areaPath *string) (*AttachmentReference, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if uploadType != nil {
        queryParams.Add("uploadType", *uploadType)
    }
    if areaPath != nil {
        queryParams.Add("areaPath", *areaPath)
    }
    body, marshalErr := json.Marshal(uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e07b5fa4-1499-494d-a496-64b860fd64ff")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Downloads an attachment.
// ctx
// id (required): Attachment ID
// project (optional): Project ID or project name
// fileName (optional): Name of the file
// download (optional): If set to <c>true</c> always download attachment
func (client Client) GetAttachmentContent(ctx context.Context, id *uuid.UUID, project *string, fileName *string, download *bool) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = (*id).String()

    queryParams := url.Values{}
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    locationId, _ := uuid.Parse("e07b5fa4-1499-494d-a496-64b860fd64ff")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Downloads an attachment.
// ctx
// id (required): Attachment ID
// project (optional): Project ID or project name
// fileName (optional): Name of the file
// download (optional): If set to <c>true</c> always download attachment
func (client Client) GetAttachmentZip(ctx context.Context, id *uuid.UUID, project *string, fileName *string, download *bool) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = (*id).String()

    queryParams := url.Values{}
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    locationId, _ := uuid.Parse("e07b5fa4-1499-494d-a496-64b860fd64ff")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Gets root classification nodes or list of classification nodes for a given list of nodes ids, for a given project. In case ids parameter is supplied you will  get list of classification nodes for those ids. Otherwise you will get root classification nodes for this project.
// ctx
// project (required): Project ID or project name
// ids (required): Comma separated integer classification nodes ids. It's not required, if you want root nodes.
// depth (optional): Depth of children to fetch.
// errorPolicy (optional): Flag to handle errors in getting some nodes. Possible options are Fail and Omit.
func (client Client) GetClassificationNodes(ctx context.Context, project *string, ids *[]int, depth *int, errorPolicy *ClassificationNodesErrorPolicy) (*[]WorkItemClassificationNode, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if ids == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "ids"}
    }
    var stringList []string
    for _, item := range *ids {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    if depth != nil {
        queryParams.Add("$depth", strconv.Itoa(*depth))
    }
    if errorPolicy != nil {
        queryParams.Add("errorPolicy", string(*errorPolicy))
    }
    locationId, _ := uuid.Parse("a70579d1-f53a-48ee-a5be-7be8659023b9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemClassificationNode
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets root classification nodes under the project.
// ctx
// project (required): Project ID or project name
// depth (optional): Depth of children to fetch.
func (client Client) GetRootNodes(ctx context.Context, project *string, depth *int) (*[]WorkItemClassificationNode, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if depth != nil {
        queryParams.Add("$depth", strconv.Itoa(*depth))
    }
    locationId, _ := uuid.Parse("a70579d1-f53a-48ee-a5be-7be8659023b9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemClassificationNode
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create new or update an existing classification node.
// ctx
// postedNode (required): Node to create or update.
// project (required): Project ID or project name
// structureGroup (required): Structure group of the classification node, area or iteration.
// path (optional): Path of the classification node.
func (client Client) CreateOrUpdateClassificationNode(ctx context.Context, postedNode *WorkItemClassificationNode, project *string, structureGroup *TreeStructureGroup, path *string) (*WorkItemClassificationNode, error) {
    if postedNode == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "postedNode"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if structureGroup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "structureGroup"} 
    }
    routeValues["structureGroup"] = string(*structureGroup)
    if path != nil && *path != "" {
        routeValues["path"] = *path
    }

    body, marshalErr := json.Marshal(*postedNode)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5a172953-1b41-49d3-840a-33f79c3ce89f")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemClassificationNode
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete an existing classification node.
// ctx
// project (required): Project ID or project name
// structureGroup (required): Structure group of the classification node, area or iteration.
// path (optional): Path of the classification node.
// reclassifyId (optional): Id of the target classification node for reclassification.
func (client Client) DeleteClassificationNode(ctx context.Context, project *string, structureGroup *TreeStructureGroup, path *string, reclassifyId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if structureGroup == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "structureGroup"} 
    }
    routeValues["structureGroup"] = string(*structureGroup)
    if path != nil && *path != "" {
        routeValues["path"] = *path
    }

    queryParams := url.Values{}
    if reclassifyId != nil {
        queryParams.Add("$reclassifyId", strconv.Itoa(*reclassifyId))
    }
    locationId, _ := uuid.Parse("5a172953-1b41-49d3-840a-33f79c3ce89f")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Gets the classification node for a given node path.
// ctx
// project (required): Project ID or project name
// structureGroup (required): Structure group of the classification node, area or iteration.
// path (optional): Path of the classification node.
// depth (optional): Depth of children to fetch.
func (client Client) GetClassificationNode(ctx context.Context, project *string, structureGroup *TreeStructureGroup, path *string, depth *int) (*WorkItemClassificationNode, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if structureGroup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "structureGroup"} 
    }
    routeValues["structureGroup"] = string(*structureGroup)
    if path != nil && *path != "" {
        routeValues["path"] = *path
    }

    queryParams := url.Values{}
    if depth != nil {
        queryParams.Add("$depth", strconv.Itoa(*depth))
    }
    locationId, _ := uuid.Parse("5a172953-1b41-49d3-840a-33f79c3ce89f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemClassificationNode
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update an existing classification node.
// ctx
// postedNode (required): Node to create or update.
// project (required): Project ID or project name
// structureGroup (required): Structure group of the classification node, area or iteration.
// path (optional): Path of the classification node.
func (client Client) UpdateClassificationNode(ctx context.Context, postedNode *WorkItemClassificationNode, project *string, structureGroup *TreeStructureGroup, path *string) (*WorkItemClassificationNode, error) {
    if postedNode == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "postedNode"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if structureGroup == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "structureGroup"} 
    }
    routeValues["structureGroup"] = string(*structureGroup)
    if path != nil && *path != "" {
        routeValues["path"] = *path
    }

    body, marshalErr := json.Marshal(*postedNode)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5a172953-1b41-49d3-840a-33f79c3ce89f")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemClassificationNode
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get users who reacted on the comment.
// ctx
// project (required): Project ID or project name
// workItemId (required): WorkItem ID.
// commentId (required): Comment ID.
// reactionType (required): Type of the reaction.
// top (optional)
// skip (optional)
func (client Client) GetEngagedUsers(ctx context.Context, project *string, workItemId *int, commentId *int, reactionType *CommentReactionType, top *int, skip *int) (*[]IdentityRef, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)
    if reactionType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "reactionType"} 
    }
    routeValues["reactionType"] = string(*reactionType)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("e33ca5e0-2349-4285-af3d-d72d86781c35")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []IdentityRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Add a comment on a work item.
// ctx
// request (required): Comment create request.
// project (required): Project ID or project name
// workItemId (required): Id of a work item.
func (client Client) AddComment(ctx context.Context, request *CommentCreate, project *string, workItemId *int) (*Comment, error) {
    if request == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "request"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)

    body, marshalErr := json.Marshal(*request)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("608aac0a-32e1-4493-a863-b9cf4566d257")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.3", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Comment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a comment on a work item.
// ctx
// project (required): Project ID or project name
// workItemId (required): Id of a work item.
// commentId (required)
func (client Client) DeleteComment(ctx context.Context, project *string, workItemId *int, commentId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("608aac0a-32e1-4493-a863-b9cf4566d257")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.3", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a work item comment.
// ctx
// project (required): Project ID or project name
// workItemId (required): Id of a work item to get the comment.
// commentId (required): Id of the comment to return.
// includeDeleted (optional): Specify if the deleted comment should be retrieved.
// expand (optional): Specifies the additional data retrieval options for work item comments.
func (client Client) GetComment(ctx context.Context, project *string, workItemId *int, commentId *int, includeDeleted *bool, expand *CommentExpandOptions) (*Comment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    queryParams := url.Values{}
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("608aac0a-32e1-4493-a863-b9cf4566d257")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.3", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Comment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of work item comments, pageable.
// ctx
// project (required): Project ID or project name
// workItemId (required): Id of a work item to get comments for.
// top (optional): Max number of comments to return.
// continuationToken (optional): Used to query for the next page of comments.
// includeDeleted (optional): Specify if the deleted comments should be retrieved.
// expand (optional): Specifies the additional data retrieval options for work item comments.
// order (optional): Order in which the comments should be returned.
func (client Client) GetComments(ctx context.Context, project *string, workItemId *int, top *int, continuationToken *string, includeDeleted *bool, expand *CommentExpandOptions, order *CommentSortOrder) (*CommentList, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if order != nil {
        queryParams.Add("order", string(*order))
    }
    locationId, _ := uuid.Parse("608aac0a-32e1-4493-a863-b9cf4566d257")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.3", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CommentList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of work item comments by ids.
// ctx
// project (required): Project ID or project name
// workItemId (required): Id of a work item to get comments for.
// ids (required): Comma-separated list of comment ids to return.
// includeDeleted (optional): Specify if the deleted comments should be retrieved.
// expand (optional): Specifies the additional data retrieval options for work item comments.
func (client Client) GetCommentsBatch(ctx context.Context, project *string, workItemId *int, ids *[]int, includeDeleted *bool, expand *CommentExpandOptions) (*CommentList, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)

    queryParams := url.Values{}
    if ids == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "ids"}
    }
    var stringList []string
    for _, item := range *ids {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("608aac0a-32e1-4493-a863-b9cf4566d257")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.3", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CommentList
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a comment on a work item.
// ctx
// request (required): Comment update request.
// project (required): Project ID or project name
// workItemId (required): Id of a work item.
// commentId (required)
func (client Client) UpdateComment(ctx context.Context, request *CommentUpdate, project *string, workItemId *int, commentId *int) (*Comment, error) {
    if request == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "request"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    body, marshalErr := json.Marshal(*request)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("608aac0a-32e1-4493-a863-b9cf4566d257")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.3", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Comment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds a new reaction to a comment.
// ctx
// project (required): Project ID or project name
// workItemId (required): WorkItem ID
// commentId (required): Comment ID
// reactionType (required): Type of the reaction
func (client Client) CreateCommentReaction(ctx context.Context, project *string, workItemId *int, commentId *int, reactionType *CommentReactionType) (*CommentReaction, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)
    if reactionType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "reactionType"} 
    }
    routeValues["reactionType"] = string(*reactionType)

    locationId, _ := uuid.Parse("f6cb3f27-1028-4851-af96-887e570dc21f")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CommentReaction
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes an existing reaction on a comment.
// ctx
// project (required): Project ID or project name
// workItemId (required): WorkItem ID
// commentId (required): Comment ID
// reactionType (required): Type of the reaction
func (client Client) DeleteCommentReaction(ctx context.Context, project *string, workItemId *int, commentId *int, reactionType *CommentReactionType) (*CommentReaction, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)
    if reactionType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "reactionType"} 
    }
    routeValues["reactionType"] = string(*reactionType)

    locationId, _ := uuid.Parse("f6cb3f27-1028-4851-af96-887e570dc21f")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CommentReaction
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets reactions of a comment.
// ctx
// project (required): Project ID or project name
// workItemId (required): WorkItem ID
// commentId (required): Comment ID
func (client Client) GetCommentReactions(ctx context.Context, project *string, workItemId *int, commentId *int) (*[]CommentReaction, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("f6cb3f27-1028-4851-af96-887e570dc21f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []CommentReaction
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// workItemId (required)
// commentId (required)
// version (required)
func (client Client) GetCommentVersion(ctx context.Context, project *string, workItemId *int, commentId *int, version *int) (*CommentVersion, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)
    if version == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "version"} 
    }
    routeValues["version"] = strconv.Itoa(*version)

    locationId, _ := uuid.Parse("49e03b34-3be0-42e3-8a5d-e8dfb88ac954")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CommentVersion
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// workItemId (required)
// commentId (required)
func (client Client) GetCommentVersions(ctx context.Context, project *string, workItemId *int, commentId *int) (*[]CommentVersion, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"} 
    }
    routeValues["workItemId"] = strconv.Itoa(*workItemId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("49e03b34-3be0-42e3-8a5d-e8dfb88ac954")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []CommentVersion
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a new field.
// ctx
// workItemField (required): New field definition
// project (optional): Project ID or project name
func (client Client) CreateField(ctx context.Context, workItemField *WorkItemField, project *string) (*WorkItemField, error) {
    if workItemField == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemField"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*workItemField)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b51fd764-e5c2-4b9b-aaf7-3395cf4bdd94")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemField
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Deletes the field.
// ctx
// fieldNameOrRefName (required): Field simple name or reference name
// project (optional): Project ID or project name
func (client Client) DeleteField(ctx context.Context, fieldNameOrRefName *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if fieldNameOrRefName == nil || *fieldNameOrRefName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fieldNameOrRefName"} 
    }
    routeValues["fieldNameOrRefName"] = *fieldNameOrRefName

    locationId, _ := uuid.Parse("b51fd764-e5c2-4b9b-aaf7-3395cf4bdd94")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Gets information on a specific field.
// ctx
// fieldNameOrRefName (required): Field simple name or reference name
// project (optional): Project ID or project name
func (client Client) GetField(ctx context.Context, fieldNameOrRefName *string, project *string) (*WorkItemField, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if fieldNameOrRefName == nil || *fieldNameOrRefName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fieldNameOrRefName"} 
    }
    routeValues["fieldNameOrRefName"] = *fieldNameOrRefName

    locationId, _ := uuid.Parse("b51fd764-e5c2-4b9b-aaf7-3395cf4bdd94")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemField
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns information for all fields.
// ctx
// project (optional): Project ID or project name
// expand (optional): Use ExtensionFields to include extension fields, otherwise exclude them. Unless the feature flag for this parameter is enabled, extension fields are always included.
func (client Client) GetFields(ctx context.Context, project *string, expand *GetFieldsExpand) (*[]WorkItemField, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("b51fd764-e5c2-4b9b-aaf7-3395cf4bdd94")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemField
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Creates a query, or moves a query.
// ctx
// postedQuery (required): The query to create.
// project (required): Project ID or project name
// query (required): The parent id or path under which the query is to be created.
// validateWiqlOnly (optional): If you only want to validate your WIQL query without actually creating one, set it to true. Default is false.
func (client Client) CreateQuery(ctx context.Context, postedQuery *QueryHierarchyItem, project *string, query *string, validateWiqlOnly *bool) (*QueryHierarchyItem, error) {
    if postedQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "postedQuery"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if query == nil || *query == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "query"} 
    }
    routeValues["query"] = *query

    queryParams := url.Values{}
    if validateWiqlOnly != nil {
        queryParams.Add("validateWiqlOnly", strconv.FormatBool(*validateWiqlOnly))
    }
    body, marshalErr := json.Marshal(*postedQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a67d190c-c41f-424b-814d-0e906f659301")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue QueryHierarchyItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a query or a folder. This deletes any permission change on the deleted query or folder and any of its descendants if it is a folder. It is important to note that the deleted permission changes cannot be recovered upon undeleting the query or folder.
// ctx
// project (required): Project ID or project name
// query (required): ID or path of the query or folder to delete.
func (client Client) DeleteQuery(ctx context.Context, project *string, query *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if query == nil || *query == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "query"} 
    }
    routeValues["query"] = *query

    locationId, _ := uuid.Parse("a67d190c-c41f-424b-814d-0e906f659301")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Gets the root queries and their children
// ctx
// project (required): Project ID or project name
// expand (optional): Include the query string (wiql), clauses, query result columns, and sort options in the results.
// depth (optional): In the folder of queries, return child queries and folders to this depth.
// includeDeleted (optional): Include deleted queries and folders
func (client Client) GetQueries(ctx context.Context, project *string, expand *QueryExpand, depth *int, includeDeleted *bool) (*[]QueryHierarchyItem, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if depth != nil {
        queryParams.Add("$depth", strconv.Itoa(*depth))
    }
    if includeDeleted != nil {
        queryParams.Add("$includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    locationId, _ := uuid.Parse("a67d190c-c41f-424b-814d-0e906f659301")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []QueryHierarchyItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieves an individual query and its children
// ctx
// project (required): Project ID or project name
// query (required): ID or path of the query.
// expand (optional): Include the query string (wiql), clauses, query result columns, and sort options in the results.
// depth (optional): In the folder of queries, return child queries and folders to this depth.
// includeDeleted (optional): Include deleted queries and folders
func (client Client) GetQuery(ctx context.Context, project *string, query *string, expand *QueryExpand, depth *int, includeDeleted *bool) (*QueryHierarchyItem, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if query == nil || *query == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "query"} 
    }
    routeValues["query"] = *query

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if depth != nil {
        queryParams.Add("$depth", strconv.Itoa(*depth))
    }
    if includeDeleted != nil {
        queryParams.Add("$includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    locationId, _ := uuid.Parse("a67d190c-c41f-424b-814d-0e906f659301")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue QueryHierarchyItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Searches all queries the user has access to in the current project
// ctx
// project (required): Project ID or project name
// filter (required): The text to filter the queries with.
// top (optional): The number of queries to return (Default is 50 and maximum is 200).
// expand (optional)
// includeDeleted (optional): Include deleted queries and folders
func (client Client) SearchQueries(ctx context.Context, project *string, filter *string, top *int, expand *QueryExpand, includeDeleted *bool) (*QueryHierarchyItemsResult, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if filter == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filter"}
    }
    queryParams.Add("$filter", *filter)
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if includeDeleted != nil {
        queryParams.Add("$includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    locationId, _ := uuid.Parse("a67d190c-c41f-424b-814d-0e906f659301")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue QueryHierarchyItemsResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update a query or a folder. This allows you to update, rename and move queries and folders.
// ctx
// queryUpdate (required): The query to update.
// project (required): Project ID or project name
// query (required): The ID or path for the query to update.
// undeleteDescendants (optional): Undelete the children of this folder. It is important to note that this will not bring back the permission changes that were previously applied to the descendants.
func (client Client) UpdateQuery(ctx context.Context, queryUpdate *QueryHierarchyItem, project *string, query *string, undeleteDescendants *bool) (*QueryHierarchyItem, error) {
    if queryUpdate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queryUpdate"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if query == nil || *query == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "query"} 
    }
    routeValues["query"] = *query

    queryParams := url.Values{}
    if undeleteDescendants != nil {
        queryParams.Add("$undeleteDescendants", strconv.FormatBool(*undeleteDescendants))
    }
    body, marshalErr := json.Marshal(*queryUpdate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a67d190c-c41f-424b-814d-0e906f659301")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue QueryHierarchyItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a list of queries by ids (Maximum 1000)
// ctx
// queryGetRequest (required)
// project (required): Project ID or project name
func (client Client) GetQueriesBatch(ctx context.Context, queryGetRequest *QueryBatchGetRequest, project *string) (*[]QueryHierarchyItem, error) {
    if queryGetRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queryGetRequest"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*queryGetRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("549816f9-09b0-4e75-9e81-01fbfcd07426")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []QueryHierarchyItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Destroys the specified work item permanently from the Recycle Bin. This action can not be undone.
// ctx
// id (required): ID of the work item to be destroyed permanently
// project (optional): Project ID or project name
func (client Client) DestroyWorkItem(ctx context.Context, id *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    locationId, _ := uuid.Parse("b70d8d39-926c-465e-b927-b1bf0e5ca0e0")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Gets a deleted work item from Recycle Bin.
// ctx
// id (required): ID of the work item to be returned
// project (optional): Project ID or project name
func (client Client) GetDeletedWorkItem(ctx context.Context, id *int, project *string) (*WorkItemDelete, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    locationId, _ := uuid.Parse("b70d8d39-926c-465e-b927-b1bf0e5ca0e0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemDelete
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the work items from the recycle bin, whose IDs have been specified in the parameters
// ctx
// ids (required): Comma separated list of IDs of the deleted work items to be returned
// project (optional): Project ID or project name
func (client Client) GetDeletedWorkItems(ctx context.Context, ids *[]int, project *string) (*[]WorkItemDeleteReference, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if ids == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "ids"}
    }
    var stringList []string
    for _, item := range *ids {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    locationId, _ := uuid.Parse("b70d8d39-926c-465e-b927-b1bf0e5ca0e0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemDeleteReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a list of the IDs and the URLs of the deleted the work items in the Recycle Bin.
// ctx
// project (optional): Project ID or project name
func (client Client) GetDeletedWorkItemShallowReferences(ctx context.Context, project *string) (*[]WorkItemDeleteShallowReference, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    locationId, _ := uuid.Parse("b70d8d39-926c-465e-b927-b1bf0e5ca0e0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemDeleteShallowReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Restores the deleted work item from Recycle Bin.
// ctx
// payload (required): Paylod with instructions to update the IsDeleted flag to false
// id (required): ID of the work item to be restored
// project (optional): Project ID or project name
func (client Client) RestoreWorkItem(ctx context.Context, payload *WorkItemDeleteUpdate, id *int, project *string) (*WorkItemDelete, error) {
    if payload == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "payload"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    body, marshalErr := json.Marshal(*payload)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b70d8d39-926c-465e-b927-b1bf0e5ca0e0")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemDelete
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns a fully hydrated work item for the requested revision
// ctx
// id (required)
// revisionNumber (required)
// project (optional): Project ID or project name
// expand (optional)
func (client Client) GetRevision(ctx context.Context, id *int, revisionNumber *int, project *string, expand *WorkItemExpand) (*WorkItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)
    if revisionNumber == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "revisionNumber"} 
    }
    routeValues["revisionNumber"] = strconv.Itoa(*revisionNumber)

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("a00c85a5-80fa-4565-99c3-bcd2181434bb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns the list of fully hydrated work item revisions, paged.
// ctx
// id (required)
// project (optional): Project ID or project name
// top (optional)
// skip (optional)
// expand (optional)
func (client Client) GetRevisions(ctx context.Context, id *int, project *string, top *int, skip *int, expand *WorkItemExpand) (*[]WorkItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("a00c85a5-80fa-4565-99c3-bcd2181434bb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a template
// ctx
// template (required): Template contents
// project (required): Project ID or project name
// team (required): Team ID or team name
func (client Client) CreateTemplate(ctx context.Context, template *WorkItemTemplate, project *string, team *string) (*WorkItemTemplate, error) {
    if template == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "template"}
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

    body, marshalErr := json.Marshal(*template)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6a90345f-a676-4969-afce-8e163e1d5642")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTemplate
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets template
// ctx
// project (required): Project ID or project name
// team (required): Team ID or team name
// workitemtypename (optional): Optional, When specified returns templates for given Work item type.
func (client Client) GetTemplates(ctx context.Context, project *string, team *string, workitemtypename *string) (*[]WorkItemTemplateReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team

    queryParams := url.Values{}
    if workitemtypename != nil {
        queryParams.Add("workitemtypename", *workitemtypename)
    }
    locationId, _ := uuid.Parse("6a90345f-a676-4969-afce-8e163e1d5642")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemTemplateReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes the template with given id
// ctx
// project (required): Project ID or project name
// team (required): Team ID or team name
// templateId (required): Template id
func (client Client) DeleteTemplate(ctx context.Context, project *string, team *string, templateId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team
    if templateId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "templateId"} 
    }
    routeValues["templateId"] = (*templateId).String()

    locationId, _ := uuid.Parse("fb10264a-8836-48a0-8033-1b0ccd2748d5")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Gets the template with specified id
// ctx
// project (required): Project ID or project name
// team (required): Team ID or team name
// templateId (required): Template Id
func (client Client) GetTemplate(ctx context.Context, project *string, team *string, templateId *uuid.UUID) (*WorkItemTemplate, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team == nil || *team == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "team"} 
    }
    routeValues["team"] = *team
    if templateId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "templateId"} 
    }
    routeValues["templateId"] = (*templateId).String()

    locationId, _ := uuid.Parse("fb10264a-8836-48a0-8033-1b0ccd2748d5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTemplate
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Replace template contents
// ctx
// templateContent (required): Template contents to replace with
// project (required): Project ID or project name
// team (required): Team ID or team name
// templateId (required): Template id
func (client Client) ReplaceTemplate(ctx context.Context, templateContent *WorkItemTemplate, project *string, team *string, templateId *uuid.UUID) (*WorkItemTemplate, error) {
    if templateContent == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "templateContent"}
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
    if templateId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "templateId"} 
    }
    routeValues["templateId"] = (*templateId).String()

    body, marshalErr := json.Marshal(*templateContent)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("fb10264a-8836-48a0-8033-1b0ccd2748d5")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTemplate
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns a single update for a work item
// ctx
// id (required)
// updateNumber (required)
// project (optional): Project ID or project name
func (client Client) GetUpdate(ctx context.Context, id *int, updateNumber *int, project *string) (*WorkItemUpdate, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)
    if updateNumber == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateNumber"} 
    }
    routeValues["updateNumber"] = strconv.Itoa(*updateNumber)

    locationId, _ := uuid.Parse("6570bf97-d02c-4a91-8d93-3abe9895b1a9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemUpdate
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns a the deltas between work item revisions
// ctx
// id (required)
// project (optional): Project ID or project name
// top (optional)
// skip (optional)
func (client Client) GetUpdates(ctx context.Context, id *int, project *string, top *int, skip *int) (*[]WorkItemUpdate, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("6570bf97-d02c-4a91-8d93-3abe9895b1a9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemUpdate
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the results of the query given its WIQL.
// ctx
// wiql (required): The query containing the WIQL.
// project (optional): Project ID or project name
// team (optional): Team ID or team name
// timePrecision (optional): Whether or not to use time precision.
// top (optional): The max number of results to return.
func (client Client) QueryByWiql(ctx context.Context, wiql *Wiql, project *string, team *string, timePrecision *bool, top *int) (*WorkItemQueryResult, error) {
    if wiql == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "wiql"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    queryParams := url.Values{}
    if timePrecision != nil {
        queryParams.Add("timePrecision", strconv.FormatBool(*timePrecision))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    body, marshalErr := json.Marshal(*wiql)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1a9c53f7-f243-4447-b110-35ef023636e4")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemQueryResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the results of the query given the query ID.
// ctx
// id (required): The query ID.
// project (optional): Project ID or project name
// team (optional): Team ID or team name
// timePrecision (optional): Whether or not to use time precision.
// top (optional): The max number of results to return.
func (client Client) GetQueryResultCount(ctx context.Context, id *uuid.UUID, project *string, team *string, timePrecision *bool, top *int) (*int, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = (*id).String()

    queryParams := url.Values{}
    if timePrecision != nil {
        queryParams.Add("timePrecision", strconv.FormatBool(*timePrecision))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("a02355f5-5f8a-4671-8e32-369d23aac83d")
    resp, err := client.Client.Send(ctx, http.MethodHead, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue int
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the results of the query given the query ID.
// ctx
// id (required): The query ID.
// project (optional): Project ID or project name
// team (optional): Team ID or team name
// timePrecision (optional): Whether or not to use time precision.
// top (optional): The max number of results to return.
func (client Client) QueryById(ctx context.Context, id *uuid.UUID, project *string, team *string, timePrecision *bool, top *int) (*WorkItemQueryResult, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = (*id).String()

    queryParams := url.Values{}
    if timePrecision != nil {
        queryParams.Add("timePrecision", strconv.FormatBool(*timePrecision))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("a02355f5-5f8a-4671-8e32-369d23aac83d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemQueryResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a work item icon given the friendly name and icon color.
// ctx
// icon (required): The name of the icon
// color (optional): The 6-digit hex color for the icon
// v (optional): The version of the icon (used only for cache invalidation)
func (client Client) GetWorkItemIconJson(ctx context.Context, icon *string, color *string, v *int) (*WorkItemIcon, error) {
    routeValues := make(map[string]string)
    if icon == nil || *icon == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "icon"} 
    }
    routeValues["icon"] = *icon

    queryParams := url.Values{}
    if color != nil {
        queryParams.Add("color", *color)
    }
    if v != nil {
        queryParams.Add("v", strconv.Itoa(*v))
    }
    locationId, _ := uuid.Parse("4e1eb4a5-1970-4228-a682-ec48eb2dca30")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemIcon
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of all work item icons.
// ctx
func (client Client) GetWorkItemIcons(ctx context.Context, ) (*[]WorkItemIcon, error) {
    locationId, _ := uuid.Parse("4e1eb4a5-1970-4228-a682-ec48eb2dca30")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemIcon
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a work item icon given the friendly name and icon color.
// ctx
// icon (required): The name of the icon
// color (optional): The 6-digit hex color for the icon
// v (optional): The version of the icon (used only for cache invalidation)
func (client Client) GetWorkItemIconSvg(ctx context.Context, icon *string, color *string, v *int) (interface{}, error) {
    routeValues := make(map[string]string)
    if icon == nil || *icon == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "icon"} 
    }
    routeValues["icon"] = *icon

    queryParams := url.Values{}
    if color != nil {
        queryParams.Add("color", *color)
    }
    if v != nil {
        queryParams.Add("v", strconv.Itoa(*v))
    }
    locationId, _ := uuid.Parse("4e1eb4a5-1970-4228-a682-ec48eb2dca30")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "image/svg+xml", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Get a work item icon given the friendly name and icon color.
// ctx
// icon (required): The name of the icon
// color (optional): The 6-digit hex color for the icon
// v (optional): The version of the icon (used only for cache invalidation)
func (client Client) GetWorkItemIconXaml(ctx context.Context, icon *string, color *string, v *int) (interface{}, error) {
    routeValues := make(map[string]string)
    if icon == nil || *icon == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "icon"} 
    }
    routeValues["icon"] = *icon

    queryParams := url.Values{}
    if color != nil {
        queryParams.Add("color", *color)
    }
    if v != nil {
        queryParams.Add("v", strconv.Itoa(*v))
    }
    locationId, _ := uuid.Parse("4e1eb4a5-1970-4228-a682-ec48eb2dca30")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "image/xaml+xml", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Get a batch of work item links
// ctx
// project (optional): Project ID or project name
// linkTypes (optional): A list of types to filter the results to specific link types. Omit this parameter to get work item links of all link types.
// types (optional): A list of types to filter the results to specific work item types. Omit this parameter to get work item links of all work item types.
// continuationToken (optional): Specifies the continuationToken to start the batch from. Omit this parameter to get the first batch of links.
// startDateTime (optional): Date/time to use as a starting point for link changes. Only link changes that occurred after that date/time will be returned. Cannot be used in conjunction with 'watermark' parameter.
func (client Client) GetReportingLinksByLinkType(ctx context.Context, project *string, linkTypes *[]string, types *[]string, continuationToken *string, startDateTime *time.Time) (*ReportingWorkItemLinksBatch, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if linkTypes != nil {
        listAsString := strings.Join((*linkTypes)[:], ",")
        queryParams.Add("linkTypes", listAsString)
    }
    if types != nil {
        listAsString := strings.Join((*types)[:], ",")
        queryParams.Add("types", listAsString)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if startDateTime != nil {
        queryParams.Add("startDateTime", (*startDateTime).String())
    }
    locationId, _ := uuid.Parse("b5b5b6d0-0308-40a1-b3f4-b9bb3c66878f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReportingWorkItemLinksBatch
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the work item relation type definition.
// ctx
// relation (required): The relation name
func (client Client) GetRelationType(ctx context.Context, relation *string) (*WorkItemRelationType, error) {
    routeValues := make(map[string]string)
    if relation == nil || *relation == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "relation"} 
    }
    routeValues["relation"] = *relation

    locationId, _ := uuid.Parse("f5d33bc9-5b49-4a3c-a9bd-f3cd46dd2165")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemRelationType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the work item relation types.
// ctx
func (client Client) GetRelationTypes(ctx context.Context, ) (*[]WorkItemRelationType, error) {
    locationId, _ := uuid.Parse("f5d33bc9-5b49-4a3c-a9bd-f3cd46dd2165")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemRelationType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a batch of work item revisions with the option of including deleted items
// ctx
// project (optional): Project ID or project name
// fields (optional): A list of fields to return in work item revisions. Omit this parameter to get all reportable fields.
// types (optional): A list of types to filter the results to specific work item types. Omit this parameter to get work item revisions of all work item types.
// continuationToken (optional): Specifies the watermark to start the batch from. Omit this parameter to get the first batch of revisions.
// startDateTime (optional): Date/time to use as a starting point for revisions, all revisions will occur after this date/time. Cannot be used in conjunction with 'watermark' parameter.
// includeIdentityRef (optional): Return an identity reference instead of a string value for identity fields.
// includeDeleted (optional): Specify if the deleted item should be returned.
// includeTagRef (optional): Specify if the tag objects should be returned for System.Tags field.
// includeLatestOnly (optional): Return only the latest revisions of work items, skipping all historical revisions
// expand (optional): Return all the fields in work item revisions, including long text fields which are not returned by default
// includeDiscussionChangesOnly (optional): Return only the those revisions of work items, where only history field was changed
// maxPageSize (optional): The maximum number of results to return in this batch
func (client Client) ReadReportingRevisionsGet(ctx context.Context, project *string, fields *[]string, types *[]string, continuationToken *string, startDateTime *time.Time, includeIdentityRef *bool, includeDeleted *bool, includeTagRef *bool, includeLatestOnly *bool, expand *ReportingRevisionsExpand, includeDiscussionChangesOnly *bool, maxPageSize *int) (*ReportingWorkItemRevisionsBatch, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if fields != nil {
        listAsString := strings.Join((*fields)[:], ",")
        queryParams.Add("fields", listAsString)
    }
    if types != nil {
        listAsString := strings.Join((*types)[:], ",")
        queryParams.Add("types", listAsString)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if startDateTime != nil {
        queryParams.Add("startDateTime", (*startDateTime).String())
    }
    if includeIdentityRef != nil {
        queryParams.Add("includeIdentityRef", strconv.FormatBool(*includeIdentityRef))
    }
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if includeTagRef != nil {
        queryParams.Add("includeTagRef", strconv.FormatBool(*includeTagRef))
    }
    if includeLatestOnly != nil {
        queryParams.Add("includeLatestOnly", strconv.FormatBool(*includeLatestOnly))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if includeDiscussionChangesOnly != nil {
        queryParams.Add("includeDiscussionChangesOnly", strconv.FormatBool(*includeDiscussionChangesOnly))
    }
    if maxPageSize != nil {
        queryParams.Add("$maxPageSize", strconv.Itoa(*maxPageSize))
    }
    locationId, _ := uuid.Parse("f828fe59-dd87-495d-a17c-7a8d6211ca6c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReportingWorkItemRevisionsBatch
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a batch of work item revisions. This request may be used if your list of fields is large enough that it may run the URL over the length limit.
// ctx
// filter (required): An object that contains request settings: field filter, type filter, identity format
// project (optional): Project ID or project name
// continuationToken (optional): Specifies the watermark to start the batch from. Omit this parameter to get the first batch of revisions.
// startDateTime (optional): Date/time to use as a starting point for revisions, all revisions will occur after this date/time. Cannot be used in conjunction with 'watermark' parameter.
// expand (optional)
func (client Client) ReadReportingRevisionsPost(ctx context.Context, filter *ReportingWorkItemRevisionsFilter, project *string, continuationToken *string, startDateTime *time.Time, expand *ReportingRevisionsExpand) (*ReportingWorkItemRevisionsBatch, error) {
    if filter == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filter"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if startDateTime != nil {
        queryParams.Add("startDateTime", (*startDateTime).String())
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    body, marshalErr := json.Marshal(*filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f828fe59-dd87-495d-a17c-7a8d6211ca6c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReportingWorkItemRevisionsBatch
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (optional): Project ID or project name
// continuationToken (optional)
// maxPageSize (optional)
func (client Client) ReadReportingDiscussions(ctx context.Context, project *string, continuationToken *string, maxPageSize *int) (*ReportingWorkItemRevisionsBatch, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if maxPageSize != nil {
        queryParams.Add("$maxPageSize", strconv.Itoa(*maxPageSize))
    }
    locationId, _ := uuid.Parse("4a644469-90c5-4fcc-9a9f-be0827d369ec")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReportingWorkItemRevisionsBatch
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Creates a single work item.
// ctx
// document (required): The JSON Patch document representing the work item
// project (required): Project ID or project name
// type_ (required): The work item type of the work item to create
// validateOnly (optional): Indicate if you only want to validate the changes without saving the work item
// bypassRules (optional): Do not enforce the work item type rules on this update
// suppressNotifications (optional): Do not fire any notifications for this change
// expand (optional): The expand parameters for work item attributes. Possible options are { None, Relations, Fields, Links, All }.
func (client Client) CreateWorkItem(ctx context.Context, document *[]JsonPatchOperation, project *string, type_ *string, validateOnly *bool, bypassRules *bool, suppressNotifications *bool, expand *WorkItemExpand) (*WorkItem, error) {
    if document == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "document"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    queryParams := url.Values{}
    if validateOnly != nil {
        queryParams.Add("validateOnly", strconv.FormatBool(*validateOnly))
    }
    if bypassRules != nil {
        queryParams.Add("bypassRules", strconv.FormatBool(*bypassRules))
    }
    if suppressNotifications != nil {
        queryParams.Add("suppressNotifications", strconv.FormatBool(*suppressNotifications))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    body, marshalErr := json.Marshal(*document)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("62d3d110-0047-428c-ad3c-4fe872c91c74")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns a single work item from a template.
// ctx
// project (required): Project ID or project name
// type_ (required): The work item type name
// fields (optional): Comma-separated list of requested fields
// asOf (optional): AsOf UTC date time string
// expand (optional): The expand parameters for work item attributes. Possible options are { None, Relations, Fields, Links, All }.
func (client Client) GetWorkItemTemplate(ctx context.Context, project *string, type_ *string, fields *string, asOf *time.Time, expand *WorkItemExpand) (*WorkItem, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    queryParams := url.Values{}
    if fields != nil {
        queryParams.Add("fields", *fields)
    }
    if asOf != nil {
        queryParams.Add("asOf", (*asOf).String())
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("62d3d110-0047-428c-ad3c-4fe872c91c74")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Deletes the specified work item and sends it to the Recycle Bin, so that it can be restored back, if required. Optionally, if the destroy parameter has been set to true, it destroys the work item permanently. WARNING: If the destroy parameter is set to true, work items deleted by this command will NOT go to recycle-bin and there is no way to restore/recover them after deletion. It is recommended NOT to use this parameter. If you do, please use this parameter with extreme caution.
// ctx
// id (required): ID of the work item to be deleted
// project (optional): Project ID or project name
// destroy (optional): Optional parameter, if set to true, the work item is deleted permanently. Please note: the destroy action is PERMANENT and cannot be undone.
func (client Client) DeleteWorkItem(ctx context.Context, id *int, project *string, destroy *bool) (*WorkItemDelete, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if destroy != nil {
        queryParams.Add("destroy", strconv.FormatBool(*destroy))
    }
    locationId, _ := uuid.Parse("72c7ddf8-2cdc-4f60-90cd-ab71c14a399b")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemDelete
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns a single work item.
// ctx
// id (required): The work item id
// project (optional): Project ID or project name
// fields (optional): Comma-separated list of requested fields
// asOf (optional): AsOf UTC date time string
// expand (optional): The expand parameters for work item attributes. Possible options are { None, Relations, Fields, Links, All }.
func (client Client) GetWorkItem(ctx context.Context, id *int, project *string, fields *[]string, asOf *time.Time, expand *WorkItemExpand) (*WorkItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if fields != nil {
        listAsString := strings.Join((*fields)[:], ",")
        queryParams.Add("fields", listAsString)
    }
    if asOf != nil {
        queryParams.Add("asOf", (*asOf).String())
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("72c7ddf8-2cdc-4f60-90cd-ab71c14a399b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns a list of work items (Maximum 200)
// ctx
// ids (required): The comma-separated list of requested work item ids. (Maximum 200 ids allowed).
// project (optional): Project ID or project name
// fields (optional): Comma-separated list of requested fields
// asOf (optional): AsOf UTC date time string
// expand (optional): The expand parameters for work item attributes. Possible options are { None, Relations, Fields, Links, All }.
// errorPolicy (optional): The flag to control error policy in a bulk get work items request. Possible options are {Fail, Omit}.
func (client Client) GetWorkItems(ctx context.Context, ids *[]int, project *string, fields *[]string, asOf *time.Time, expand *WorkItemExpand, errorPolicy *WorkItemErrorPolicy) (*[]WorkItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if ids == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "ids"}
    }
    var stringList []string
    for _, item := range *ids {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    if fields != nil {
        listAsString := strings.Join((*fields)[:], ",")
        queryParams.Add("fields", listAsString)
    }
    if asOf != nil {
        queryParams.Add("asOf", (*asOf).String())
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if errorPolicy != nil {
        queryParams.Add("errorPolicy", string(*errorPolicy))
    }
    locationId, _ := uuid.Parse("72c7ddf8-2cdc-4f60-90cd-ab71c14a399b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Updates a single work item.
// ctx
// document (required): The JSON Patch document representing the update
// id (required): The id of the work item to update
// project (optional): Project ID or project name
// validateOnly (optional): Indicate if you only want to validate the changes without saving the work item
// bypassRules (optional): Do not enforce the work item type rules on this update
// suppressNotifications (optional): Do not fire any notifications for this change
// expand (optional): The expand parameters for work item attributes. Possible options are { None, Relations, Fields, Links, All }.
func (client Client) UpdateWorkItem(ctx context.Context, document *[]JsonPatchOperation, id *int, project *string, validateOnly *bool, bypassRules *bool, suppressNotifications *bool, expand *WorkItemExpand) (*WorkItem, error) {
    if document == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "document"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if validateOnly != nil {
        queryParams.Add("validateOnly", strconv.FormatBool(*validateOnly))
    }
    if bypassRules != nil {
        queryParams.Add("bypassRules", strconv.FormatBool(*bypassRules))
    }
    if suppressNotifications != nil {
        queryParams.Add("suppressNotifications", strconv.FormatBool(*suppressNotifications))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    body, marshalErr := json.Marshal(*document)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("72c7ddf8-2cdc-4f60-90cd-ab71c14a399b")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets work items for a list of work item ids (Maximum 200)
// ctx
// workItemGetRequest (required)
// project (optional): Project ID or project name
func (client Client) GetWorkItemsBatch(ctx context.Context, workItemGetRequest *WorkItemBatchGetRequest, project *string) (*[]WorkItem, error) {
    if workItemGetRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemGetRequest"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*workItemGetRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("908509b6-4248-4475-a1cd-829139ba419f")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns the next state on the given work item IDs.
// ctx
// ids (required): list of work item ids
// action (optional): possible actions. Currently only supports checkin
func (client Client) GetWorkItemNextStatesOnCheckinAction(ctx context.Context, ids *[]int, action *string) (*[]WorkItemNextStateOnTransition, error) {
    queryParams := url.Values{}
    if ids == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "ids"}
    }
    var stringList []string
    for _, item := range *ids {
        stringList = append(stringList, strconv.Itoa(item))
    }
    listAsString := strings.Join((stringList)[:], ",")
    queryParams.Add("definitions", listAsString)
    if action != nil {
        queryParams.Add("action", *action)
    }
    locationId, _ := uuid.Parse("afae844b-e2f6-44c2-8053-17b3bb936a40")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemNextStateOnTransition
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get all work item type categories.
// ctx
// project (required): Project ID or project name
func (client Client) GetWorkItemTypeCategories(ctx context.Context, project *string) (*[]WorkItemTypeCategory, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("9b9f5734-36c8-415e-ba67-f83b45c31408")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemTypeCategory
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get specific work item type category by name.
// ctx
// project (required): Project ID or project name
// category (required): The category name
func (client Client) GetWorkItemTypeCategory(ctx context.Context, project *string, category *string) (*WorkItemTypeCategory, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if category == nil || *category == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "category"} 
    }
    routeValues["category"] = *category

    locationId, _ := uuid.Parse("9b9f5734-36c8-415e-ba67-f83b45c31408")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTypeCategory
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns a work item type definition.
// ctx
// project (required): Project ID or project name
// type_ (required): Work item type name
func (client Client) GetWorkItemType(ctx context.Context, project *string, type_ *string) (*WorkItemType, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("7c8d7a76-4a09-43e8-b5df-bd792f4ac6aa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns the list of work item types
// ctx
// project (required): Project ID or project name
func (client Client) GetWorkItemTypes(ctx context.Context, project *string) (*[]WorkItemType, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("7c8d7a76-4a09-43e8-b5df-bd792f4ac6aa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of fields for a work item type with detailed references.
// ctx
// project (required): Project ID or project name
// type_ (required): Work item type.
// expand (optional): Expand level for the API response. Properties: to include allowedvalues, default value, isRequired etc. as a part of response; None: to skip these properties.
func (client Client) GetWorkItemTypeFieldsWithReferences(ctx context.Context, project *string, type_ *string, expand *WorkItemTypeFieldsExpandLevel) (*[]WorkItemTypeFieldWithReferences, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("bd293ce5-3d25-4192-8e67-e8092e879efb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemTypeFieldWithReferences
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a field for a work item type with detailed references.
// ctx
// project (required): Project ID or project name
// type_ (required): Work item type.
// field (required)
// expand (optional): Expand level for the API response. Properties: to include allowedvalues, default value, isRequired etc. as a part of response; None: to skip these properties.
func (client Client) GetWorkItemTypeFieldWithReferences(ctx context.Context, project *string, type_ *string, field *string, expand *WorkItemTypeFieldsExpandLevel) (*WorkItemTypeFieldWithReferences, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_
    if field == nil || *field == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "field"} 
    }
    routeValues["field"] = *field

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    locationId, _ := uuid.Parse("bd293ce5-3d25-4192-8e67-e8092e879efb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemTypeFieldWithReferences
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns the state names and colors for a work item type.
// ctx
// project (required): Project ID or project name
// type_ (required): The state name
func (client Client) GetWorkItemTypeStates(ctx context.Context, project *string, type_ *string) (*[]WorkItemStateColor, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("7c9d7a76-4a09-43e8-b5df-bd792f4ac6aa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemStateColor
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

