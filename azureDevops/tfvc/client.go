// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tfvc

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

var ResourceAreaId, _ = uuid.Parse("8aa40520-446d-40e6-89f6-9c9f9ce44c48")

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

// Get a single branch hierarchy at the given path with parents or children as specified.
// ctx
// path (required): Full path to the branch.  Default: $/ Examples: $/, $/MyProject, $/MyProject/SomeFolder.
// project (optional): Project ID or project name
// includeParent (optional): Return the parent branch, if there is one. Default: False
// includeChildren (optional): Return child branches, if there are any. Default: False
func (client Client) GetBranch(ctx context.Context, path *string, project *string, includeParent *bool, includeChildren *bool) (*TfvcBranch, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if includeParent != nil {
        queryParams.Add("includeParent", strconv.FormatBool(*includeParent))
    }
    if includeChildren != nil {
        queryParams.Add("includeChildren", strconv.FormatBool(*includeChildren))
    }
    locationId, _ := uuid.Parse("bc1f417e-239d-42e7-85e1-76e80cb2d6eb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TfvcBranch
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a collection of branch roots -- first-level children, branches with no parents.
// ctx
// project (optional): Project ID or project name
// includeParent (optional): Return the parent branch, if there is one. Default: False
// includeChildren (optional): Return the child branches for each root branch. Default: False
// includeDeleted (optional): Return deleted branches. Default: False
// includeLinks (optional): Return links. Default: False
func (client Client) GetBranches(ctx context.Context, project *string, includeParent *bool, includeChildren *bool, includeDeleted *bool, includeLinks *bool) (*[]TfvcBranch, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if includeParent != nil {
        queryParams.Add("includeParent", strconv.FormatBool(*includeParent))
    }
    if includeChildren != nil {
        queryParams.Add("includeChildren", strconv.FormatBool(*includeChildren))
    }
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    locationId, _ := uuid.Parse("bc1f417e-239d-42e7-85e1-76e80cb2d6eb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcBranch
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get branch hierarchies below the specified scopePath
// ctx
// scopePath (required): Full path to the branch.  Default: $/ Examples: $/, $/MyProject, $/MyProject/SomeFolder.
// project (optional): Project ID or project name
// includeDeleted (optional): Return deleted branches. Default: False
// includeLinks (optional): Return links. Default: False
func (client Client) GetBranchRefs(ctx context.Context, scopePath *string, project *string, includeDeleted *bool, includeLinks *bool) (*[]TfvcBranchRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if scopePath == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopePath"}
    }
    queryParams.Add("scopePath", *scopePath)
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    locationId, _ := uuid.Parse("bc1f417e-239d-42e7-85e1-76e80cb2d6eb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcBranchRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve Tfvc changes for a given changeset.
// ctx
// id (optional): ID of the changeset. Default: null
// skip (optional): Number of results to skip. Default: null
// top (optional): The maximum number of results to return. Default: null
func (client Client) GetChangesetChanges(ctx context.Context, id *int, skip *int, top *int) (*[]TfvcChange, error) {
    routeValues := make(map[string]string)
    if id != nil {
        routeValues["id"] = strconv.Itoa(*id)
    }

    queryParams := url.Values{}
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("f32b86f2-15b9-4fe6-81b1-6f8938617ee5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcChange
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a new changeset.
// ctx
// changeset (required)
// project (optional): Project ID or project name
func (client Client) CreateChangeset(ctx context.Context, changeset *TfvcChangeset, project *string) (*TfvcChangesetRef, error) {
    if changeset == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "changeset"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*changeset)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0bc8f0a4-6bfb-42a9-ba84-139da7b99c49")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TfvcChangesetRef
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a Tfvc Changeset
// ctx
// id (required): Changeset Id to retrieve.
// project (optional): Project ID or project name
// maxChangeCount (optional): Number of changes to return (maximum 100 changes) Default: 0
// includeDetails (optional): Include policy details and check-in notes in the response. Default: false
// includeWorkItems (optional): Include workitems. Default: false
// maxCommentLength (optional): Include details about associated work items in the response. Default: null
// includeSourceRename (optional): Include renames.  Default: false
// skip (optional): Number of results to skip. Default: null
// top (optional): The maximum number of results to return. Default: null
// orderby (optional): Results are sorted by ID in descending order by default. Use id asc to sort by ID in ascending order.
// searchCriteria (optional): Following criteria available (.itemPath, .version, .versionType, .versionOption, .author, .fromId, .toId, .fromDate, .toDate) Default: null
func (client Client) GetChangeset(ctx context.Context, id *int, project *string, maxChangeCount *int, includeDetails *bool, includeWorkItems *bool, maxCommentLength *int, includeSourceRename *bool, skip *int, top *int, orderby *string, searchCriteria *TfvcChangesetSearchCriteria) (*TfvcChangeset, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if id == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "id"} 
    }
    routeValues["id"] = strconv.Itoa(*id)

    queryParams := url.Values{}
    if maxChangeCount != nil {
        queryParams.Add("maxChangeCount", strconv.Itoa(*maxChangeCount))
    }
    if includeDetails != nil {
        queryParams.Add("includeDetails", strconv.FormatBool(*includeDetails))
    }
    if includeWorkItems != nil {
        queryParams.Add("includeWorkItems", strconv.FormatBool(*includeWorkItems))
    }
    if maxCommentLength != nil {
        queryParams.Add("maxCommentLength", strconv.Itoa(*maxCommentLength))
    }
    if includeSourceRename != nil {
        queryParams.Add("includeSourceRename", strconv.FormatBool(*includeSourceRename))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if orderby != nil {
        queryParams.Add("$orderby", *orderby)
    }
    if searchCriteria != nil {
        if searchCriteria.ItemPath != nil {
            queryParams.Add("searchCriteria.itemPath", *searchCriteria.ItemPath)
        }
        if searchCriteria.Author != nil {
            queryParams.Add("searchCriteria.author", *searchCriteria.Author)
        }
        if searchCriteria.FromDate != nil {
            queryParams.Add("searchCriteria.fromDate", *searchCriteria.FromDate)
        }
        if searchCriteria.ToDate != nil {
            queryParams.Add("searchCriteria.toDate", *searchCriteria.ToDate)
        }
        if searchCriteria.FromId != nil {
            queryParams.Add("searchCriteria.fromId", strconv.Itoa(*searchCriteria.FromId))
        }
        if searchCriteria.ToId != nil {
            queryParams.Add("searchCriteria.toId", strconv.Itoa(*searchCriteria.ToId))
        }
        if searchCriteria.FollowRenames != nil {
            queryParams.Add("searchCriteria.followRenames", strconv.FormatBool(*searchCriteria.FollowRenames))
        }
        if searchCriteria.IncludeLinks != nil {
            queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*searchCriteria.IncludeLinks))
        }
    }
    locationId, _ := uuid.Parse("0bc8f0a4-6bfb-42a9-ba84-139da7b99c49")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TfvcChangeset
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve Tfvc Changesets
// ctx
// project (optional): Project ID or project name
// maxCommentLength (optional): Include details about associated work items in the response. Default: null
// skip (optional): Number of results to skip. Default: null
// top (optional): The maximum number of results to return. Default: null
// orderby (optional): Results are sorted by ID in descending order by default. Use id asc to sort by ID in ascending order.
// searchCriteria (optional): Following criteria available (.itemPath, .version, .versionType, .versionOption, .author, .fromId, .toId, .fromDate, .toDate) Default: null
func (client Client) GetChangesets(ctx context.Context, project *string, maxCommentLength *int, skip *int, top *int, orderby *string, searchCriteria *TfvcChangesetSearchCriteria) (*[]TfvcChangesetRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if maxCommentLength != nil {
        queryParams.Add("maxCommentLength", strconv.Itoa(*maxCommentLength))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if orderby != nil {
        queryParams.Add("$orderby", *orderby)
    }
    if searchCriteria != nil {
        if searchCriteria.ItemPath != nil {
            queryParams.Add("searchCriteria.itemPath", *searchCriteria.ItemPath)
        }
        if searchCriteria.Author != nil {
            queryParams.Add("searchCriteria.author", *searchCriteria.Author)
        }
        if searchCriteria.FromDate != nil {
            queryParams.Add("searchCriteria.fromDate", *searchCriteria.FromDate)
        }
        if searchCriteria.ToDate != nil {
            queryParams.Add("searchCriteria.toDate", *searchCriteria.ToDate)
        }
        if searchCriteria.FromId != nil {
            queryParams.Add("searchCriteria.fromId", strconv.Itoa(*searchCriteria.FromId))
        }
        if searchCriteria.ToId != nil {
            queryParams.Add("searchCriteria.toId", strconv.Itoa(*searchCriteria.ToId))
        }
        if searchCriteria.FollowRenames != nil {
            queryParams.Add("searchCriteria.followRenames", strconv.FormatBool(*searchCriteria.FollowRenames))
        }
        if searchCriteria.IncludeLinks != nil {
            queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*searchCriteria.IncludeLinks))
        }
    }
    locationId, _ := uuid.Parse("0bc8f0a4-6bfb-42a9-ba84-139da7b99c49")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcChangesetRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Returns changesets for a given list of changeset Ids.
// ctx
// changesetsRequestData (required): List of changeset IDs.
func (client Client) GetBatchedChangesets(ctx context.Context, changesetsRequestData *TfvcChangesetsRequestData) (*[]TfvcChangesetRef, error) {
    if changesetsRequestData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "changesetsRequestData"}
    }
    body, marshalErr := json.Marshal(*changesetsRequestData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b7e7c173-803c-4fea-9ec8-31ee35c5502a")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcChangesetRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieves the work items associated with a particular changeset.
// ctx
// id (optional): ID of the changeset. Default: null
func (client Client) GetChangesetWorkItems(ctx context.Context, id *int) (*[]AssociatedWorkItem, error) {
    routeValues := make(map[string]string)
    if id != nil {
        routeValues["id"] = strconv.Itoa(*id)
    }

    locationId, _ := uuid.Parse("64ae0bea-1d71-47c9-a9e5-fe73f5ea0ff4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AssociatedWorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Post for retrieving a set of items given a list of paths or a long path. Allows for specifying the recursionLevel and version descriptors for each path.
// ctx
// itemRequestData (required)
// project (optional): Project ID or project name
func (client Client) GetItemsBatch(ctx context.Context, itemRequestData *TfvcItemRequestData, project *string) (*[][]TfvcItem, error) {
    if itemRequestData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "itemRequestData"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*itemRequestData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("fe6f827b-5f64-480f-b8af-1eca3b80e833")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue [][]TfvcItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Post for retrieving a set of items given a list of paths or a long path. Allows for specifying the recursionLevel and version descriptors for each path.
// ctx
// itemRequestData (required)
// project (optional): Project ID or project name
func (client Client) GetItemsBatchZip(ctx context.Context, itemRequestData *TfvcItemRequestData, project *string) (*interface{}, error) {
    if itemRequestData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "itemRequestData"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*itemRequestData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("fe6f827b-5f64-480f-b8af-1eca3b80e833")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
// ctx
// path (required): Version control path of an individual item to return.
// project (optional): Project ID or project name
// fileName (optional): file name of item returned.
// download (optional): If true, create a downloadable attachment.
// scopePath (optional): Version control path of a folder to return multiple items.
// recursionLevel (optional): None (just the item), or OneLevel (contents of a folder).
// versionDescriptor (optional): Version descriptor.  Default is null.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
func (client Client) GetItem(ctx context.Context, path *string, project *string, fileName *string, download *bool, scopePath *string, recursionLevel *VersionControlRecursionType, versionDescriptor *TfvcVersionDescriptor, includeContent *bool) (*TfvcItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", *recursionLevel)
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionOption != nil {
            queryParams.Add("versionDescriptor.versionOption", string(*versionDescriptor.VersionOption))
        }
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TfvcItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
// ctx
// path (required): Version control path of an individual item to return.
// project (optional): Project ID or project name
// fileName (optional): file name of item returned.
// download (optional): If true, create a downloadable attachment.
// scopePath (optional): Version control path of a folder to return multiple items.
// recursionLevel (optional): None (just the item), or OneLevel (contents of a folder).
// versionDescriptor (optional): Version descriptor.  Default is null.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
func (client Client) GetItemContent(ctx context.Context, path *string, project *string, fileName *string, download *bool, scopePath *string, recursionLevel *VersionControlRecursionType, versionDescriptor *TfvcVersionDescriptor, includeContent *bool) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", *recursionLevel)
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionOption != nil {
            queryParams.Add("versionDescriptor.versionOption", string(*versionDescriptor.VersionOption))
        }
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of Tfvc items
// ctx
// project (optional): Project ID or project name
// scopePath (optional): Version control path of a folder to return multiple items.
// recursionLevel (optional): None (just the item), or OneLevel (contents of a folder).
// includeLinks (optional): True to include links.
// versionDescriptor (optional)
func (client Client) GetItems(ctx context.Context, project *string, scopePath *string, recursionLevel *VersionControlRecursionType, includeLinks *bool, versionDescriptor *TfvcVersionDescriptor) (*[]TfvcItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", *recursionLevel)
    }
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionOption != nil {
            queryParams.Add("versionDescriptor.versionOption", string(*versionDescriptor.VersionOption))
        }
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
    }
    locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
// ctx
// path (required): Version control path of an individual item to return.
// project (optional): Project ID or project name
// fileName (optional): file name of item returned.
// download (optional): If true, create a downloadable attachment.
// scopePath (optional): Version control path of a folder to return multiple items.
// recursionLevel (optional): None (just the item), or OneLevel (contents of a folder).
// versionDescriptor (optional): Version descriptor.  Default is null.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
func (client Client) GetItemText(ctx context.Context, path *string, project *string, fileName *string, download *bool, scopePath *string, recursionLevel *VersionControlRecursionType, versionDescriptor *TfvcVersionDescriptor, includeContent *bool) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", *recursionLevel)
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionOption != nil {
            queryParams.Add("versionDescriptor.versionOption", string(*versionDescriptor.VersionOption))
        }
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
// ctx
// path (required): Version control path of an individual item to return.
// project (optional): Project ID or project name
// fileName (optional): file name of item returned.
// download (optional): If true, create a downloadable attachment.
// scopePath (optional): Version control path of a folder to return multiple items.
// recursionLevel (optional): None (just the item), or OneLevel (contents of a folder).
// versionDescriptor (optional): Version descriptor.  Default is null.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
func (client Client) GetItemZip(ctx context.Context, path *string, project *string, fileName *string, download *bool, scopePath *string, recursionLevel *VersionControlRecursionType, versionDescriptor *TfvcVersionDescriptor, includeContent *bool) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", *recursionLevel)
    }
    if versionDescriptor != nil {
        if versionDescriptor.VersionOption != nil {
            queryParams.Add("versionDescriptor.versionOption", string(*versionDescriptor.VersionOption))
        }
        if versionDescriptor.VersionType != nil {
            queryParams.Add("versionDescriptor.versionType", string(*versionDescriptor.VersionType))
        }
        if versionDescriptor.Version != nil {
            queryParams.Add("versionDescriptor.version", *versionDescriptor.Version)
        }
    }
    if includeContent != nil {
        queryParams.Add("includeContent", strconv.FormatBool(*includeContent))
    }
    locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get items under a label.
// ctx
// labelId (required): Unique identifier of label
// top (optional): Max number of items to return
// skip (optional): Number of items to skip
func (client Client) GetLabelItems(ctx context.Context, labelId *string, top *int, skip *int) (*[]TfvcItem, error) {
    routeValues := make(map[string]string)
    if labelId == nil || *labelId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "labelId"} 
    }
    routeValues["labelId"] = *labelId

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("06166e34-de17-4b60-8cd1-23182a346fda")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a single deep label.
// ctx
// labelId (required): Unique identifier of label
// requestData (required): maxItemCount
// project (optional): Project ID or project name
func (client Client) GetLabel(ctx context.Context, labelId *string, requestData *TfvcLabelRequestData, project *string) (*TfvcLabel, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if labelId == nil || *labelId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "labelId"} 
    }
    routeValues["labelId"] = *labelId

    queryParams := url.Values{}
    if requestData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "requestData"}
    }
    if requestData.LabelScope != nil {
        queryParams.Add("requestData.labelScope", *requestData.LabelScope)
    }
    if requestData.Name != nil {
        queryParams.Add("requestData.name", *requestData.Name)
    }
    if requestData.Owner != nil {
        queryParams.Add("requestData.owner", *requestData.Owner)
    }
    if requestData.ItemLabelFilter != nil {
        queryParams.Add("requestData.itemLabelFilter", *requestData.ItemLabelFilter)
    }
    if requestData.MaxItemCount != nil {
        queryParams.Add("requestData.maxItemCount", strconv.Itoa(*requestData.MaxItemCount))
    }
    if requestData.IncludeLinks != nil {
        queryParams.Add("requestData.includeLinks", strconv.FormatBool(*requestData.IncludeLinks))
    }
    locationId, _ := uuid.Parse("a5d9bd7f-b661-4d0e-b9be-d9c16affae54")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TfvcLabel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a collection of shallow label references.
// ctx
// requestData (required): labelScope, name, owner, and itemLabelFilter
// project (optional): Project ID or project name
// top (optional): Max number of labels to return
// skip (optional): Number of labels to skip
func (client Client) GetLabels(ctx context.Context, requestData *TfvcLabelRequestData, project *string, top *int, skip *int) (*[]TfvcLabelRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if requestData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "requestData"}
    }
    if requestData.LabelScope != nil {
        queryParams.Add("requestData.labelScope", *requestData.LabelScope)
    }
    if requestData.Name != nil {
        queryParams.Add("requestData.name", *requestData.Name)
    }
    if requestData.Owner != nil {
        queryParams.Add("requestData.owner", *requestData.Owner)
    }
    if requestData.ItemLabelFilter != nil {
        queryParams.Add("requestData.itemLabelFilter", *requestData.ItemLabelFilter)
    }
    if requestData.MaxItemCount != nil {
        queryParams.Add("requestData.maxItemCount", strconv.Itoa(*requestData.MaxItemCount))
    }
    if requestData.IncludeLinks != nil {
        queryParams.Add("requestData.includeLinks", strconv.FormatBool(*requestData.IncludeLinks))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("a5d9bd7f-b661-4d0e-b9be-d9c16affae54")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcLabelRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get changes included in a shelveset.
// ctx
// shelvesetId (required): Shelveset's unique ID
// top (optional): Max number of changes to return
// skip (optional): Number of changes to skip
func (client Client) GetShelvesetChanges(ctx context.Context, shelvesetId *string, top *int, skip *int) (*[]TfvcChange, error) {
    queryParams := url.Values{}
    if shelvesetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "shelvesetId"}
    }
    queryParams.Add("shelvesetId", *shelvesetId)
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("dbaf075b-0445-4c34-9e5b-82292f856522")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcChange
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a single deep shelveset.
// ctx
// shelvesetId (required): Shelveset's unique ID
// requestData (optional): includeDetails, includeWorkItems, maxChangeCount, and maxCommentLength
func (client Client) GetShelveset(ctx context.Context, shelvesetId *string, requestData *TfvcShelvesetRequestData) (*TfvcShelveset, error) {
    queryParams := url.Values{}
    if shelvesetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "shelvesetId"}
    }
    queryParams.Add("shelvesetId", *shelvesetId)
    if requestData != nil {
        if requestData.Name != nil {
            queryParams.Add("requestData.name", *requestData.Name)
        }
        if requestData.Owner != nil {
            queryParams.Add("requestData.owner", *requestData.Owner)
        }
        if requestData.MaxCommentLength != nil {
            queryParams.Add("requestData.maxCommentLength", strconv.Itoa(*requestData.MaxCommentLength))
        }
        if requestData.MaxChangeCount != nil {
            queryParams.Add("requestData.maxChangeCount", strconv.Itoa(*requestData.MaxChangeCount))
        }
        if requestData.IncludeDetails != nil {
            queryParams.Add("requestData.includeDetails", strconv.FormatBool(*requestData.IncludeDetails))
        }
        if requestData.IncludeWorkItems != nil {
            queryParams.Add("requestData.includeWorkItems", strconv.FormatBool(*requestData.IncludeWorkItems))
        }
        if requestData.IncludeLinks != nil {
            queryParams.Add("requestData.includeLinks", strconv.FormatBool(*requestData.IncludeLinks))
        }
    }
    locationId, _ := uuid.Parse("e36d44fb-e907-4b0a-b194-f83f1ed32ad3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TfvcShelveset
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Return a collection of shallow shelveset references.
// ctx
// requestData (optional): name, owner, and maxCommentLength
// top (optional): Max number of shelvesets to return
// skip (optional): Number of shelvesets to skip
func (client Client) GetShelvesets(ctx context.Context, requestData *TfvcShelvesetRequestData, top *int, skip *int) (*[]TfvcShelvesetRef, error) {
    queryParams := url.Values{}
    if requestData != nil {
        if requestData.Name != nil {
            queryParams.Add("requestData.name", *requestData.Name)
        }
        if requestData.Owner != nil {
            queryParams.Add("requestData.owner", *requestData.Owner)
        }
        if requestData.MaxCommentLength != nil {
            queryParams.Add("requestData.maxCommentLength", strconv.Itoa(*requestData.MaxCommentLength))
        }
        if requestData.MaxChangeCount != nil {
            queryParams.Add("requestData.maxChangeCount", strconv.Itoa(*requestData.MaxChangeCount))
        }
        if requestData.IncludeDetails != nil {
            queryParams.Add("requestData.includeDetails", strconv.FormatBool(*requestData.IncludeDetails))
        }
        if requestData.IncludeWorkItems != nil {
            queryParams.Add("requestData.includeWorkItems", strconv.FormatBool(*requestData.IncludeWorkItems))
        }
        if requestData.IncludeLinks != nil {
            queryParams.Add("requestData.includeLinks", strconv.FormatBool(*requestData.IncludeLinks))
        }
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("e36d44fb-e907-4b0a-b194-f83f1ed32ad3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TfvcShelvesetRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get work items associated with a shelveset.
// ctx
// shelvesetId (required): Shelveset's unique ID
func (client Client) GetShelvesetWorkItems(ctx context.Context, shelvesetId *string) (*[]AssociatedWorkItem, error) {
    queryParams := url.Values{}
    if shelvesetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "shelvesetId"}
    }
    queryParams.Add("shelvesetId", *shelvesetId)
    locationId, _ := uuid.Parse("a7a0c1c1-373e-425a-b031-a519474d743d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AssociatedWorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

