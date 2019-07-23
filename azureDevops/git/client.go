// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package git

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "io"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("4e080c62-fa21-4fbc-8fef-2a10a2b38049")

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

// [Preview API] Create an annotated tag.
// ctx
// tagObject (required): Object containing details of tag to be created.
// project (required): Project ID or project name
// repositoryId (required): ID or name of the repository.
func (client Client) CreateAnnotatedTag(ctx context.Context, tagObject *GitAnnotatedTag, project *string, repositoryId *string) (*GitAnnotatedTag, error) {
    if tagObject == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "tagObject"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*tagObject)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5e8a8081-3851-4626-b677-9891cc04102e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitAnnotatedTag
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get an annotated tag.
// ctx
// project (required): Project ID or project name
// repositoryId (required): ID or name of the repository.
// objectId (required): ObjectId (Sha1Id) of tag to get.
func (client Client) GetAnnotatedTag(ctx context.Context, project *string, repositoryId *string, objectId *string) (*GitAnnotatedTag, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if objectId == nil || *objectId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "objectId"} 
    }
    routeValues["objectId"] = *objectId

    locationId, _ := uuid.Parse("5e8a8081-3851-4626-b677-9891cc04102e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitAnnotatedTag
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a single blob.
// ctx
// repositoryId (required): The name or ID of the repository.
// sha1 (required): SHA1 hash of the file. You can get the SHA1 of a file using the "Git/Items/Get Item" endpoint.
// project (optional): Project ID or project name
// download (optional): If true, prompt for a download rather than rendering in a browser. Note: this value defaults to true if $format is zip
// fileName (optional): Provide a fileName to use for a download.
// resolveLfs (optional): If true, try to resolve a blob to its LFS contents, if it's an LFS pointer file. Only compatible with octet-stream Accept headers or $format types
func (client Client) GetBlob(ctx context.Context, repositoryId *string, sha1 *string, project *string, download *bool, fileName *string, resolveLfs *bool) (*GitBlobRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if sha1 == nil || *sha1 == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sha1"} 
    }
    routeValues["sha1"] = *sha1

    queryParams := url.Values{}
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if resolveLfs != nil {
        queryParams.Add("resolveLfs", strconv.FormatBool(*resolveLfs))
    }
    locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitBlobRef
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a single blob.
// ctx
// repositoryId (required): The name or ID of the repository.
// sha1 (required): SHA1 hash of the file. You can get the SHA1 of a file using the "Git/Items/Get Item" endpoint.
// project (optional): Project ID or project name
// download (optional): If true, prompt for a download rather than rendering in a browser. Note: this value defaults to true if $format is zip
// fileName (optional): Provide a fileName to use for a download.
// resolveLfs (optional): If true, try to resolve a blob to its LFS contents, if it's an LFS pointer file. Only compatible with octet-stream Accept headers or $format types
func (client Client) GetBlobContent(ctx context.Context, repositoryId *string, sha1 *string, project *string, download *bool, fileName *string, resolveLfs *bool) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if sha1 == nil || *sha1 == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sha1"} 
    }
    routeValues["sha1"] = *sha1

    queryParams := url.Values{}
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if resolveLfs != nil {
        queryParams.Add("resolveLfs", strconv.FormatBool(*resolveLfs))
    }
    locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Gets one or more blobs in a zip file download.
// ctx
// blobIds (required): Blob IDs (SHA1 hashes) to be returned in the zip file.
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// filename (optional)
func (client Client) GetBlobsZip(ctx context.Context, blobIds *[]string, repositoryId *string, project *string, filename *string) (interface{}, error) {
    if blobIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "blobIds"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if filename != nil {
        queryParams.Add("filename", *filename)
    }
    body, marshalErr := json.Marshal(*blobIds)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Get a single blob.
// ctx
// repositoryId (required): The name or ID of the repository.
// sha1 (required): SHA1 hash of the file. You can get the SHA1 of a file using the "Git/Items/Get Item" endpoint.
// project (optional): Project ID or project name
// download (optional): If true, prompt for a download rather than rendering in a browser. Note: this value defaults to true if $format is zip
// fileName (optional): Provide a fileName to use for a download.
// resolveLfs (optional): If true, try to resolve a blob to its LFS contents, if it's an LFS pointer file. Only compatible with octet-stream Accept headers or $format types
func (client Client) GetBlobZip(ctx context.Context, repositoryId *string, sha1 *string, project *string, download *bool, fileName *string, resolveLfs *bool) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if sha1 == nil || *sha1 == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sha1"} 
    }
    routeValues["sha1"] = *sha1

    queryParams := url.Values{}
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    if resolveLfs != nil {
        queryParams.Add("resolveLfs", strconv.FormatBool(*resolveLfs))
    }
    locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Retrieve statistics about a single branch.
// ctx
// repositoryId (required): The name or ID of the repository.
// name (required): Name of the branch.
// project (optional): Project ID or project name
// baseVersionDescriptor (optional): Identifies the commit or branch to use as the base.
func (client Client) GetBranch(ctx context.Context, repositoryId *string, name *string, project *string, baseVersionDescriptor *GitVersionDescriptor) (*GitBranchStats, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if name == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "name"}
    }
    queryParams.Add("name", *name)
    if baseVersionDescriptor != nil {
        if baseVersionDescriptor.VersionType != nil {
            queryParams.Add("baseVersionDescriptor.versionType", string(*baseVersionDescriptor.VersionType))
        }
        if baseVersionDescriptor.Version != nil {
            queryParams.Add("baseVersionDescriptor.version", *baseVersionDescriptor.Version)
        }
        if baseVersionDescriptor.VersionOptions != nil {
            queryParams.Add("baseVersionDescriptor.versionOptions", string(*baseVersionDescriptor.VersionOptions))
        }
    }
    locationId, _ := uuid.Parse("d5b216de-d8d5-4d32-ae76-51df755b16d3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitBranchStats
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve statistics about all branches within a repository.
// ctx
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// baseVersionDescriptor (optional): Identifies the commit or branch to use as the base.
func (client Client) GetBranches(ctx context.Context, repositoryId *string, project *string, baseVersionDescriptor *GitVersionDescriptor) (*[]GitBranchStats, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if baseVersionDescriptor != nil {
        if baseVersionDescriptor.VersionType != nil {
            queryParams.Add("baseVersionDescriptor.versionType", string(*baseVersionDescriptor.VersionType))
        }
        if baseVersionDescriptor.Version != nil {
            queryParams.Add("baseVersionDescriptor.version", *baseVersionDescriptor.Version)
        }
        if baseVersionDescriptor.VersionOptions != nil {
            queryParams.Add("baseVersionDescriptor.versionOptions", string(*baseVersionDescriptor.VersionOptions))
        }
    }
    locationId, _ := uuid.Parse("d5b216de-d8d5-4d32-ae76-51df755b16d3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitBranchStats
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve changes for a particular commit.
// ctx
// commitId (required): The id of the commit.
// repositoryId (required): The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
// project (optional): Project ID or project name
// top (optional): The maximum number of changes to return.
// skip (optional): The number of changes to skip.
func (client Client) GetChanges(ctx context.Context, commitId *string, repositoryId *string, project *string, top *int, skip *int) (*GitCommitChanges, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if commitId == nil || *commitId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "commitId"} 
    }
    routeValues["commitId"] = *commitId
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("5bf884f5-3e07-42e9-afb8-1b872267bf16")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitCommitChanges
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Cherry pick a specific commit or commits that are associated to a pull request into a new branch.
// ctx
// cherryPickToCreate (required)
// project (required): Project ID or project name
// repositoryId (required): ID of the repository.
func (client Client) CreateCherryPick(ctx context.Context, cherryPickToCreate *GitAsyncRefOperationParameters, project *string, repositoryId *string) (*GitCherryPick, error) {
    if cherryPickToCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cherryPickToCreate"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*cherryPickToCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("033bad68-9a14-43d1-90e0-59cb8856fef6")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitCherryPick
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve information about a cherry pick by cherry pick Id.
// ctx
// project (required): Project ID or project name
// cherryPickId (required): ID of the cherry pick.
// repositoryId (required): ID of the repository.
func (client Client) GetCherryPick(ctx context.Context, project *string, cherryPickId *int, repositoryId *string) (*GitCherryPick, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if cherryPickId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cherryPickId"} 
    }
    routeValues["cherryPickId"] = strconv.Itoa(*cherryPickId)
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    locationId, _ := uuid.Parse("033bad68-9a14-43d1-90e0-59cb8856fef6")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitCherryPick
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve information about a cherry pick for a specific branch.
// ctx
// project (required): Project ID or project name
// repositoryId (required): ID of the repository.
// refName (required): The GitAsyncRefOperationParameters generatedRefName used for the cherry pick operation.
func (client Client) GetCherryPickForRefName(ctx context.Context, project *string, repositoryId *string, refName *string) (*GitCherryPick, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if refName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "refName"}
    }
    queryParams.Add("refName", *refName)
    locationId, _ := uuid.Parse("033bad68-9a14-43d1-90e0-59cb8856fef6")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitCherryPick
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Find the closest common commit (the merge base) between base and target commits, and get the diff between either the base and target commits or common and target commits.
// ctx
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// diffCommonCommit (optional): If true, diff between common and target commits. If false, diff between base and target commits.
// top (optional): Maximum number of changes to return. Defaults to 100.
// skip (optional): Number of changes to skip
// baseVersionDescriptor (optional): Descriptor for base commit.
// targetVersionDescriptor (optional): Descriptor for target commit.
func (client Client) GetCommitDiffs(ctx context.Context, repositoryId *string, project *string, diffCommonCommit *bool, top *int, skip *int, baseVersionDescriptor *GitBaseVersionDescriptor, targetVersionDescriptor *GitTargetVersionDescriptor) (*GitCommitDiffs, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if diffCommonCommit != nil {
        queryParams.Add("diffCommonCommit", strconv.FormatBool(*diffCommonCommit))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if baseVersionDescriptor != nil {
        if baseVersionDescriptor.BaseVersionType != nil {
            queryParams.Add("baseVersionType", string(*baseVersionDescriptor.BaseVersionType))
        }
        if baseVersionDescriptor.BaseVersion != nil {
            queryParams.Add("baseVersion", *baseVersionDescriptor.BaseVersion)
        }
        if baseVersionDescriptor.BaseVersionOptions != nil {
            queryParams.Add("baseVersionOptions", string(*baseVersionDescriptor.BaseVersionOptions))
        }
    }
    if targetVersionDescriptor != nil {
        if targetVersionDescriptor.TargetVersionType != nil {
            queryParams.Add("targetVersionType", string(*targetVersionDescriptor.TargetVersionType))
        }
        if targetVersionDescriptor.TargetVersion != nil {
            queryParams.Add("targetVersion", *targetVersionDescriptor.TargetVersion)
        }
        if targetVersionDescriptor.TargetVersionOptions != nil {
            queryParams.Add("targetVersionOptions", string(*targetVersionDescriptor.TargetVersionOptions))
        }
    }
    locationId, _ := uuid.Parse("615588d5-c0c7-4b88-88f8-e625306446e8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitCommitDiffs
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a particular commit.
// ctx
// commitId (required): The id of the commit.
// repositoryId (required): The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
// project (optional): Project ID or project name
// changeCount (optional): The number of changes to include in the result.
func (client Client) GetCommit(ctx context.Context, commitId *string, repositoryId *string, project *string, changeCount *int) (*GitCommit, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if commitId == nil || *commitId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "commitId"} 
    }
    routeValues["commitId"] = *commitId
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if changeCount != nil {
        queryParams.Add("changeCount", strconv.Itoa(*changeCount))
    }
    locationId, _ := uuid.Parse("c2570c3b-5b3f-41b8-98bf-5407bfde8d58")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitCommit
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve git commits for a project
// ctx
// repositoryId (required): The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
// searchCriteria (required)
// project (optional): Project ID or project name
func (client Client) GetCommits(ctx context.Context, repositoryId *string, searchCriteria *GitQueryCommitsCriteria, project *string) (*[]GitCommitRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if searchCriteria == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "searchCriteria"}
    }
    if searchCriteria.Ids != nil {
        for index, item := range *searchCriteria.Ids {
            queryParams.Add("searchCriteria.ids[" + strconv.Itoa(index) + "]", item)
        }
    }
    if searchCriteria.FromDate != nil {
        queryParams.Add("searchCriteria.fromDate", *searchCriteria.FromDate)
    }
    if searchCriteria.ToDate != nil {
        queryParams.Add("searchCriteria.toDate", *searchCriteria.ToDate)
    }
    if searchCriteria.ItemVersion != nil {
        if searchCriteria.ItemVersion.VersionType != nil {
            queryParams.Add("searchCriteria.itemVersion.versionType", string(*searchCriteria.ItemVersion.VersionType))
        }
        if searchCriteria.ItemVersion.Version != nil {
            queryParams.Add("searchCriteria.itemVersion.version", *searchCriteria.ItemVersion.Version)
        }
        if searchCriteria.ItemVersion.VersionOptions != nil {
            queryParams.Add("searchCriteria.itemVersion.versionOptions", string(*searchCriteria.ItemVersion.VersionOptions))
        }
    }
    if searchCriteria.CompareVersion != nil {
        if searchCriteria.CompareVersion.VersionType != nil {
            queryParams.Add("searchCriteria.compareVersion.versionType", string(*searchCriteria.CompareVersion.VersionType))
        }
        if searchCriteria.CompareVersion.Version != nil {
            queryParams.Add("searchCriteria.compareVersion.version", *searchCriteria.CompareVersion.Version)
        }
        if searchCriteria.CompareVersion.VersionOptions != nil {
            queryParams.Add("searchCriteria.compareVersion.versionOptions", string(*searchCriteria.CompareVersion.VersionOptions))
        }
    }
    if searchCriteria.FromCommitId != nil {
        queryParams.Add("searchCriteria.fromCommitId", *searchCriteria.FromCommitId)
    }
    if searchCriteria.ToCommitId != nil {
        queryParams.Add("searchCriteria.toCommitId", *searchCriteria.ToCommitId)
    }
    if searchCriteria.User != nil {
        queryParams.Add("searchCriteria.user", *searchCriteria.User)
    }
    if searchCriteria.Author != nil {
        queryParams.Add("searchCriteria.author", *searchCriteria.Author)
    }
    if searchCriteria.ItemPath != nil {
        queryParams.Add("searchCriteria.itemPath", *searchCriteria.ItemPath)
    }
    if searchCriteria.ExcludeDeletes != nil {
        queryParams.Add("searchCriteria.excludeDeletes", strconv.FormatBool(*searchCriteria.ExcludeDeletes))
    }
    if searchCriteria.Skip != nil {
        queryParams.Add("searchCriteria.$skip", strconv.Itoa(*searchCriteria.Skip))
    }
    if searchCriteria.Top != nil {
        queryParams.Add("searchCriteria.$top", strconv.Itoa(*searchCriteria.Top))
    }
    if searchCriteria.IncludeLinks != nil {
        queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*searchCriteria.IncludeLinks))
    }
    if searchCriteria.IncludeWorkItems != nil {
        queryParams.Add("searchCriteria.includeWorkItems", strconv.FormatBool(*searchCriteria.IncludeWorkItems))
    }
    if searchCriteria.IncludeUserImageUrl != nil {
        queryParams.Add("searchCriteria.includeUserImageUrl", strconv.FormatBool(*searchCriteria.IncludeUserImageUrl))
    }
    if searchCriteria.IncludePushData != nil {
        queryParams.Add("searchCriteria.includePushData", strconv.FormatBool(*searchCriteria.IncludePushData))
    }
    if searchCriteria.HistoryMode != nil {
        queryParams.Add("searchCriteria.historyMode", string(*searchCriteria.HistoryMode))
    }
    locationId, _ := uuid.Parse("c2570c3b-5b3f-41b8-98bf-5407bfde8d58")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitCommitRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a list of commits associated with a particular push.
// ctx
// repositoryId (required): The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
// pushId (required): The id of the push.
// project (optional): Project ID or project name
// top (optional): The maximum number of commits to return ("get the top x commits").
// skip (optional): The number of commits to skip.
// includeLinks (optional): Set to false to avoid including REST Url links for resources. Defaults to true.
func (client Client) GetPushCommits(ctx context.Context, repositoryId *string, pushId *int, project *string, top *int, skip *int, includeLinks *bool) (*[]GitCommitRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if pushId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pushId"}
    }
    queryParams.Add("pushId", strconv.Itoa(*pushId))
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("skip", strconv.Itoa(*skip))
    }
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    locationId, _ := uuid.Parse("c2570c3b-5b3f-41b8-98bf-5407bfde8d58")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitCommitRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve git commits for a project matching the search criteria
// ctx
// searchCriteria (required): Search options
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// skip (optional): Number of commits to skip.
// top (optional): Maximum number of commits to return.
// includeStatuses (optional): True to include additional commit status information.
func (client Client) GetCommitsBatch(ctx context.Context, searchCriteria *GitQueryCommitsCriteria, repositoryId *string, project *string, skip *int, top *int, includeStatuses *bool) (*[]GitCommitRef, error) {
    if searchCriteria == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "searchCriteria"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if includeStatuses != nil {
        queryParams.Add("includeStatuses", strconv.FormatBool(*includeStatuses))
    }
    body, marshalErr := json.Marshal(*searchCriteria)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6400dfb2-0bcb-462b-b992-5a57f8f1416c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitCommitRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve deleted git repositories.
// ctx
// project (required): Project ID or project name
func (client Client) GetDeletedRepositories(ctx context.Context, project *string) (*[]GitDeletedRepository, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("2b6869c4-cb25-42b5-b7a3-0d3e6be0a11a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitDeletedRepository
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve all forks of a repository in the collection.
// ctx
// repositoryNameOrId (required): The name or ID of the repository.
// collectionId (required): Team project collection ID.
// project (optional): Project ID or project name
// includeLinks (optional): True to include links.
func (client Client) GetForks(ctx context.Context, repositoryNameOrId *string, collectionId *uuid.UUID, project *string, includeLinks *bool) (*[]GitRepositoryRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryNameOrId == nil || *repositoryNameOrId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryNameOrId"} 
    }
    routeValues["repositoryNameOrId"] = *repositoryNameOrId
    if collectionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "collectionId"} 
    }
    routeValues["collectionId"] = (*collectionId).String()

    queryParams := url.Values{}
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    locationId, _ := uuid.Parse("158c0340-bf6f-489c-9625-d572a1480d57")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitRepositoryRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Request that another repository's refs be fetched into this one. It syncs two existing forks. To create a fork, please see the <a href="https://docs.microsoft.com/en-us/rest/api/vsts/git/repositories/create?view=azure-devops-rest-5.1"> repositories endpoint</a>
// ctx
// syncParams (required): Source repository and ref mapping.
// repositoryNameOrId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// includeLinks (optional): True to include links
func (client Client) CreateForkSyncRequest(ctx context.Context, syncParams *GitForkSyncRequestParameters, repositoryNameOrId *string, project *string, includeLinks *bool) (*GitForkSyncRequest, error) {
    if syncParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "syncParams"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryNameOrId == nil || *repositoryNameOrId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryNameOrId"} 
    }
    routeValues["repositoryNameOrId"] = *repositoryNameOrId

    queryParams := url.Values{}
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    body, marshalErr := json.Marshal(*syncParams)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1703f858-b9d1-46af-ab62-483e9e1055b5")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitForkSyncRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a specific fork sync operation's details.
// ctx
// repositoryNameOrId (required): The name or ID of the repository.
// forkSyncOperationId (required): OperationId of the sync request.
// project (optional): Project ID or project name
// includeLinks (optional): True to include links.
func (client Client) GetForkSyncRequest(ctx context.Context, repositoryNameOrId *string, forkSyncOperationId *int, project *string, includeLinks *bool) (*GitForkSyncRequest, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryNameOrId == nil || *repositoryNameOrId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryNameOrId"} 
    }
    routeValues["repositoryNameOrId"] = *repositoryNameOrId
    if forkSyncOperationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "forkSyncOperationId"} 
    }
    routeValues["forkSyncOperationId"] = strconv.Itoa(*forkSyncOperationId)

    queryParams := url.Values{}
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    locationId, _ := uuid.Parse("1703f858-b9d1-46af-ab62-483e9e1055b5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitForkSyncRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve all requested fork sync operations on this repository.
// ctx
// repositoryNameOrId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// includeAbandoned (optional): True to include abandoned requests.
// includeLinks (optional): True to include links.
func (client Client) GetForkSyncRequests(ctx context.Context, repositoryNameOrId *string, project *string, includeAbandoned *bool, includeLinks *bool) (*[]GitForkSyncRequest, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryNameOrId == nil || *repositoryNameOrId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryNameOrId"} 
    }
    routeValues["repositoryNameOrId"] = *repositoryNameOrId

    queryParams := url.Values{}
    if includeAbandoned != nil {
        queryParams.Add("includeAbandoned", strconv.FormatBool(*includeAbandoned))
    }
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    locationId, _ := uuid.Parse("1703f858-b9d1-46af-ab62-483e9e1055b5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitForkSyncRequest
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create an import request.
// ctx
// importRequest (required): The import request to create.
// project (required): Project ID or project name
// repositoryId (required): The name or ID of the repository.
func (client Client) CreateImportRequest(ctx context.Context, importRequest *GitImportRequest, project *string, repositoryId *string) (*GitImportRequest, error) {
    if importRequest == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "importRequest"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*importRequest)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitImportRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve a particular import request.
// ctx
// project (required): Project ID or project name
// repositoryId (required): The name or ID of the repository.
// importRequestId (required): The unique identifier for the import request.
func (client Client) GetImportRequest(ctx context.Context, project *string, repositoryId *string, importRequestId *int) (*GitImportRequest, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if importRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "importRequestId"} 
    }
    routeValues["importRequestId"] = strconv.Itoa(*importRequestId)

    locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitImportRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve import requests for a repository.
// ctx
// project (required): Project ID or project name
// repositoryId (required): The name or ID of the repository.
// includeAbandoned (optional): True to include abandoned import requests in the results.
func (client Client) QueryImportRequests(ctx context.Context, project *string, repositoryId *string, includeAbandoned *bool) (*[]GitImportRequest, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if includeAbandoned != nil {
        queryParams.Add("includeAbandoned", strconv.FormatBool(*includeAbandoned))
    }
    locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitImportRequest
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retry or abandon a failed import request.
// ctx
// importRequestToUpdate (required): The updated version of the import request. Currently, the only change allowed is setting the Status to Queued or Abandoned.
// project (required): Project ID or project name
// repositoryId (required): The name or ID of the repository.
// importRequestId (required): The unique identifier for the import request to update.
func (client Client) UpdateImportRequest(ctx context.Context, importRequestToUpdate *GitImportRequest, project *string, repositoryId *string, importRequestId *int) (*GitImportRequest, error) {
    if importRequestToUpdate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "importRequestToUpdate"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if importRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "importRequestId"} 
    }
    routeValues["importRequestId"] = strconv.Itoa(*importRequestId)

    body, marshalErr := json.Marshal(*importRequestToUpdate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitImportRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
// ctx
// repositoryId (required): The name or ID of the repository.
// path (required): The item path.
// project (optional): Project ID or project name
// scopePath (optional): The path scope.  The default is null.
// recursionLevel (optional): The recursion level of this request. The default is 'none', no recursion.
// includeContentMetadata (optional): Set to true to include content metadata.  Default is false.
// latestProcessedChange (optional): Set to true to include the latest changes.  Default is false.
// download (optional): Set to true to download the response as a file.  Default is false.
// versionDescriptor (optional): Version descriptor.  Default is the default branch for the repository.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
// resolveLfs (optional): Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
func (client Client) GetItem(ctx context.Context, repositoryId *string, path *string, project *string, scopePath *string, recursionLevel *VersionControlRecursionType, includeContentMetadata *bool, latestProcessedChange *bool, download *bool, versionDescriptor *GitVersionDescriptor, includeContent *bool, resolveLfs *bool) (*GitItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContentMetadata != nil {
        queryParams.Add("includeContentMetadata", strconv.FormatBool(*includeContentMetadata))
    }
    if latestProcessedChange != nil {
        queryParams.Add("latestProcessedChange", strconv.FormatBool(*latestProcessedChange))
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
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
    if resolveLfs != nil {
        queryParams.Add("resolveLfs", strconv.FormatBool(*resolveLfs))
    }
    locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitItem
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
// ctx
// repositoryId (required): The name or ID of the repository.
// path (required): The item path.
// project (optional): Project ID or project name
// scopePath (optional): The path scope.  The default is null.
// recursionLevel (optional): The recursion level of this request. The default is 'none', no recursion.
// includeContentMetadata (optional): Set to true to include content metadata.  Default is false.
// latestProcessedChange (optional): Set to true to include the latest changes.  Default is false.
// download (optional): Set to true to download the response as a file.  Default is false.
// versionDescriptor (optional): Version descriptor.  Default is the default branch for the repository.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
// resolveLfs (optional): Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
func (client Client) GetItemContent(ctx context.Context, repositoryId *string, path *string, project *string, scopePath *string, recursionLevel *VersionControlRecursionType, includeContentMetadata *bool, latestProcessedChange *bool, download *bool, versionDescriptor *GitVersionDescriptor, includeContent *bool, resolveLfs *bool) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContentMetadata != nil {
        queryParams.Add("includeContentMetadata", strconv.FormatBool(*includeContentMetadata))
    }
    if latestProcessedChange != nil {
        queryParams.Add("latestProcessedChange", strconv.FormatBool(*latestProcessedChange))
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
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
    if resolveLfs != nil {
        queryParams.Add("resolveLfs", strconv.FormatBool(*resolveLfs))
    }
    locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Get Item Metadata and/or Content for a collection of items. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
// ctx
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// scopePath (optional): The path scope.  The default is null.
// recursionLevel (optional): The recursion level of this request. The default is 'none', no recursion.
// includeContentMetadata (optional): Set to true to include content metadata.  Default is false.
// latestProcessedChange (optional): Set to true to include the latest changes.  Default is false.
// download (optional): Set to true to download the response as a file.  Default is false.
// includeLinks (optional): Set to true to include links to items.  Default is false.
// versionDescriptor (optional): Version descriptor.  Default is the default branch for the repository.
func (client Client) GetItems(ctx context.Context, repositoryId *string, project *string, scopePath *string, recursionLevel *VersionControlRecursionType, includeContentMetadata *bool, latestProcessedChange *bool, download *bool, includeLinks *bool, versionDescriptor *GitVersionDescriptor) (*[]GitItem, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContentMetadata != nil {
        queryParams.Add("includeContentMetadata", strconv.FormatBool(*includeContentMetadata))
    }
    if latestProcessedChange != nil {
        queryParams.Add("latestProcessedChange", strconv.FormatBool(*latestProcessedChange))
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
    }
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
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
    locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
// ctx
// repositoryId (required): The name or ID of the repository.
// path (required): The item path.
// project (optional): Project ID or project name
// scopePath (optional): The path scope.  The default is null.
// recursionLevel (optional): The recursion level of this request. The default is 'none', no recursion.
// includeContentMetadata (optional): Set to true to include content metadata.  Default is false.
// latestProcessedChange (optional): Set to true to include the latest changes.  Default is false.
// download (optional): Set to true to download the response as a file.  Default is false.
// versionDescriptor (optional): Version descriptor.  Default is the default branch for the repository.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
// resolveLfs (optional): Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
func (client Client) GetItemText(ctx context.Context, repositoryId *string, path *string, project *string, scopePath *string, recursionLevel *VersionControlRecursionType, includeContentMetadata *bool, latestProcessedChange *bool, download *bool, versionDescriptor *GitVersionDescriptor, includeContent *bool, resolveLfs *bool) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContentMetadata != nil {
        queryParams.Add("includeContentMetadata", strconv.FormatBool(*includeContentMetadata))
    }
    if latestProcessedChange != nil {
        queryParams.Add("latestProcessedChange", strconv.FormatBool(*latestProcessedChange))
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
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
    if resolveLfs != nil {
        queryParams.Add("resolveLfs", strconv.FormatBool(*resolveLfs))
    }
    locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
// ctx
// repositoryId (required): The name or ID of the repository.
// path (required): The item path.
// project (optional): Project ID or project name
// scopePath (optional): The path scope.  The default is null.
// recursionLevel (optional): The recursion level of this request. The default is 'none', no recursion.
// includeContentMetadata (optional): Set to true to include content metadata.  Default is false.
// latestProcessedChange (optional): Set to true to include the latest changes.  Default is false.
// download (optional): Set to true to download the response as a file.  Default is false.
// versionDescriptor (optional): Version descriptor.  Default is the default branch for the repository.
// includeContent (optional): Set to true to include item content when requesting json.  Default is false.
// resolveLfs (optional): Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
func (client Client) GetItemZip(ctx context.Context, repositoryId *string, path *string, project *string, scopePath *string, recursionLevel *VersionControlRecursionType, includeContentMetadata *bool, latestProcessedChange *bool, download *bool, versionDescriptor *GitVersionDescriptor, includeContent *bool, resolveLfs *bool) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if path == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "path"}
    }
    queryParams.Add("path", *path)
    if scopePath != nil {
        queryParams.Add("scopePath", *scopePath)
    }
    if recursionLevel != nil {
        queryParams.Add("recursionLevel", string(*recursionLevel))
    }
    if includeContentMetadata != nil {
        queryParams.Add("includeContentMetadata", strconv.FormatBool(*includeContentMetadata))
    }
    if latestProcessedChange != nil {
        queryParams.Add("latestProcessedChange", strconv.FormatBool(*latestProcessedChange))
    }
    if download != nil {
        queryParams.Add("download", strconv.FormatBool(*download))
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
    if resolveLfs != nil {
        queryParams.Add("resolveLfs", strconv.FormatBool(*resolveLfs))
    }
    locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Post for retrieving a creating a batch out of a set of items in a repo / project given a list of paths or a long path
// ctx
// requestData (required): Request data attributes: ItemDescriptors, IncludeContentMetadata, LatestProcessedChange, IncludeLinks. ItemDescriptors: Collection of items to fetch, including path, version, and recursion level. IncludeContentMetadata: Whether to include metadata for all items LatestProcessedChange: Whether to include shallow ref to commit that last changed each item. IncludeLinks: Whether to include the _links field on the shallow references.
// repositoryId (required): The name or ID of the repository
// project (optional): Project ID or project name
func (client Client) GetItemsBatch(ctx context.Context, requestData *GitItemRequestData, repositoryId *string, project *string) (*[][]GitItem, error) {
    if requestData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "requestData"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*requestData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("630fd2e4-fb88-4f85-ad21-13f3fd1fbca9")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue [][]GitItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Find the merge bases of two commits, optionally across forks. If otherRepositoryId is not specified, the merge bases will only be calculated within the context of the local repositoryNameOrId.
// ctx
// repositoryNameOrId (required): ID or name of the local repository.
// commitId (required): First commit, usually the tip of the target branch of the potential merge.
// otherCommitId (required): Other commit, usually the tip of the source branch of the potential merge.
// project (optional): Project ID or project name
// otherCollectionId (optional): The collection ID where otherCommitId lives.
// otherRepositoryId (optional): The repository ID where otherCommitId lives.
func (client Client) GetMergeBases(ctx context.Context, repositoryNameOrId *string, commitId *string, otherCommitId *string, project *string, otherCollectionId *uuid.UUID, otherRepositoryId *uuid.UUID) (*[]GitCommitRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryNameOrId == nil || *repositoryNameOrId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryNameOrId"} 
    }
    routeValues["repositoryNameOrId"] = *repositoryNameOrId
    if commitId == nil || *commitId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "commitId"} 
    }
    routeValues["commitId"] = *commitId

    queryParams := url.Values{}
    if otherCommitId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "otherCommitId"}
    }
    queryParams.Add("otherCommitId", *otherCommitId)
    if otherCollectionId != nil {
        queryParams.Add("otherCollectionId", (*otherCollectionId).String())
    }
    if otherRepositoryId != nil {
        queryParams.Add("otherRepositoryId", (*otherRepositoryId).String())
    }
    locationId, _ := uuid.Parse("7cf2abb6-c964-4f7e-9872-f78c66e72e9c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitCommitRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Request a git merge operation. Currently we support merging only 2 commits.
// ctx
// mergeParameters (required): Parents commitIds and merge commit messsage.
// project (required): Project ID or project name
// repositoryNameOrId (required): The name or ID of the repository.
// includeLinks (optional): True to include links
func (client Client) CreateMergeRequest(ctx context.Context, mergeParameters *GitMergeParameters, project *string, repositoryNameOrId *string, includeLinks *bool) (*GitMerge, error) {
    if mergeParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "mergeParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryNameOrId == nil || *repositoryNameOrId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryNameOrId"} 
    }
    routeValues["repositoryNameOrId"] = *repositoryNameOrId

    queryParams := url.Values{}
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    body, marshalErr := json.Marshal(*mergeParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("985f7ae9-844f-4906-9897-7ef41516c0e2")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitMerge
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a specific merge operation's details.
// ctx
// project (required): Project ID or project name
// repositoryNameOrId (required): The name or ID of the repository.
// mergeOperationId (required): OperationId of the merge request.
// includeLinks (optional): True to include links
func (client Client) GetMergeRequest(ctx context.Context, project *string, repositoryNameOrId *string, mergeOperationId *int, includeLinks *bool) (*GitMerge, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryNameOrId == nil || *repositoryNameOrId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryNameOrId"} 
    }
    routeValues["repositoryNameOrId"] = *repositoryNameOrId
    if mergeOperationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "mergeOperationId"} 
    }
    routeValues["mergeOperationId"] = strconv.Itoa(*mergeOperationId)

    queryParams := url.Values{}
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    locationId, _ := uuid.Parse("985f7ae9-844f-4906-9897-7ef41516c0e2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitMerge
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve a list of policy configurations by a given set of scope/filtering criteria.
// ctx
// project (required): Project ID or project name
// repositoryId (optional): The repository id.
// refName (optional): The fully-qualified Git ref name (e.g. refs/heads/master).
// policyType (optional): The policy type filter.
// top (optional): Maximum number of policies to return.
// continuationToken (optional): Pass a policy configuration ID to fetch the next page of results, up to top number of results, for this endpoint.
func (client Client) GetPolicyConfigurations(ctx context.Context, project *string, repositoryId *uuid.UUID, refName *string, policyType *uuid.UUID, top *int, continuationToken *string) (*GitPolicyConfigurationResponse, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if repositoryId != nil {
        queryParams.Add("repositoryId", (*repositoryId).String())
    }
    if refName != nil {
        queryParams.Add("refName", *refName)
    }
    if policyType != nil {
        queryParams.Add("policyType", (*policyType).String())
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("2c420070-a0a2-49cc-9639-c9f271c5ff07")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseBodyValue []PolicyConfiguration
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *GitPolicyConfigurationResponse
    xmsContinuationTokenHeader := resp.Header.Get("x-ms-continuationtoken")
    if err == nil {
        responseValue = &GitPolicyConfigurationResponse {
            PolicyConfigurations: &responseBodyValue,
            ContinuationToken: &xmsContinuationTokenHeader,
        }
    }

    return responseValue, err
}

// [Preview API] Attach a new file to a pull request.
// ctx
// uploadStream (required): Stream to upload
// fileName (required): The name of the file.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) CreateAttachment(ctx context.Context, uploadStream io.Reader, fileName *string, repositoryId *string, pullRequestId *int, project *string) (*Attachment, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if fileName == nil || *fileName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fileName"} 
    }
    routeValues["fileName"] = *fileName
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, uploadStream, "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Attachment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a pull request attachment.
// ctx
// fileName (required): The name of the attachment to delete.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) DeleteAttachment(ctx context.Context, fileName *string, repositoryId *string, pullRequestId *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if fileName == nil || *fileName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fileName"} 
    }
    routeValues["fileName"] = *fileName
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the file content of a pull request attachment.
// ctx
// fileName (required): The name of the attachment.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetAttachmentContent(ctx context.Context, fileName *string, repositoryId *string, pullRequestId *int, project *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if fileName == nil || *fileName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fileName"} 
    }
    routeValues["fileName"] = *fileName
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Get a list of files attached to a given pull request.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetAttachments(ctx context.Context, repositoryId *string, pullRequestId *int, project *string) (*[]Attachment, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Attachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the file content of a pull request attachment.
// ctx
// fileName (required): The name of the attachment.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetAttachmentZip(ctx context.Context, fileName *string, repositoryId *string, pullRequestId *int, project *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if fileName == nil || *fileName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "fileName"} 
    }
    routeValues["fileName"] = *fileName
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Add a like on a comment.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): The ID of the thread that contains the comment.
// commentId (required): The ID of the comment.
// project (optional): Project ID or project name
func (client Client) CreateLike(ctx context.Context, repositoryId *string, pullRequestId *int, threadId *int, commentId *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)
    if commentId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("5f2e2851-1389-425b-a00b-fb2adb3ef31b")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Delete a like on a comment.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): The ID of the thread that contains the comment.
// commentId (required): The ID of the comment.
// project (optional): Project ID or project name
func (client Client) DeleteLike(ctx context.Context, repositoryId *string, pullRequestId *int, threadId *int, commentId *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)
    if commentId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("5f2e2851-1389-425b-a00b-fb2adb3ef31b")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get likes for a comment.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): The ID of the thread that contains the comment.
// commentId (required): The ID of the comment.
// project (optional): Project ID or project name
func (client Client) GetLikes(ctx context.Context, repositoryId *string, pullRequestId *int, threadId *int, commentId *int, project *string) (*[]IdentityRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("5f2e2851-1389-425b-a00b-fb2adb3ef31b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []IdentityRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get the commits for the specified iteration of a pull request.
// ctx
// repositoryId (required): ID or name of the repository.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the iteration from which to get the commits.
// project (optional): Project ID or project name
// top (optional): Maximum number of commits to return. The maximum number of commits that can be returned per batch is 500.
// skip (optional): Number of commits to skip.
func (client Client) GetPullRequestIterationCommits(ctx context.Context, repositoryId *string, pullRequestId *int, iterationId *int, project *string, top *int, skip *int) (*[]GitCommitRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("skip", strconv.Itoa(*skip))
    }
    locationId, _ := uuid.Parse("e7ea0883-095f-4926-b5fb-f24691c26fb9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitCommitRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get the commits for the specified pull request.
// ctx
// repositoryId (required): ID or name of the repository.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetPullRequestCommits(ctx context.Context, repositoryId *string, pullRequestId *int, project *string) (*[]GitCommitRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("52823034-34a8-4576-922c-8d8b77e9e4c4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitCommitRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve the changes made in a pull request between two iterations.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the pull request iteration. <br /> Iteration IDs are zero-based with zero indicating the common commit between the source and target branches. Iteration one is the head of the source branch at the time the pull request is created and subsequent iterations are created when there are pushes to the source branch.
// project (optional): Project ID or project name
// top (optional): Optional. The number of changes to retrieve.  The default value is 100 and the maximum value is 2000.
// skip (optional): Optional. The number of changes to ignore.  For example, to retrieve changes 101-150, set top 50 and skip to 100.
// compareTo (optional): ID of the pull request iteration to compare against.  The default value is zero which indicates the comparison is made against the common commit between the source and target branches
func (client Client) GetPullRequestIterationChanges(ctx context.Context, repositoryId *string, pullRequestId *int, iterationId *int, project *string, top *int, skip *int, compareTo *int) (*GitPullRequestIterationChanges, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if compareTo != nil {
        queryParams.Add("$compareTo", strconv.Itoa(*compareTo))
    }
    locationId, _ := uuid.Parse("4216bdcf-b6b1-4d59-8b82-c34cc183fc8b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestIterationChanges
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get the specified iteration for a pull request.
// ctx
// repositoryId (required): ID or name of the repository.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the pull request iteration to return.
// project (optional): Project ID or project name
func (client Client) GetPullRequestIteration(ctx context.Context, repositoryId *string, pullRequestId *int, iterationId *int, project *string) (*GitPullRequestIteration, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    locationId, _ := uuid.Parse("d43911ee-6958-46b0-a42b-8445b8a0d004")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestIteration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get the list of iterations for the specified pull request.
// ctx
// repositoryId (required): ID or name of the repository.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
// includeCommits (optional): If true, include the commits associated with each iteration in the response.
func (client Client) GetPullRequestIterations(ctx context.Context, repositoryId *string, pullRequestId *int, project *string, includeCommits *bool) (*[]GitPullRequestIteration, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    queryParams := url.Values{}
    if includeCommits != nil {
        queryParams.Add("includeCommits", strconv.FormatBool(*includeCommits))
    }
    locationId, _ := uuid.Parse("d43911ee-6958-46b0-a42b-8445b8a0d004")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitPullRequestIteration
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a pull request status on the iteration. This operation will have the same result as Create status on pull request with specified iteration ID in the request body.
// ctx
// status (required): Pull request status to create.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the pull request iteration.
// project (optional): Project ID or project name
func (client Client) CreatePullRequestIterationStatus(ctx context.Context, status *GitPullRequestStatus, repositoryId *string, pullRequestId *int, iterationId *int, project *string) (*GitPullRequestStatus, error) {
    if status == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "status"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    body, marshalErr := json.Marshal(*status)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestStatus
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete pull request iteration status.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the pull request iteration.
// statusId (required): ID of the pull request status.
// project (optional): Project ID or project name
func (client Client) DeletePullRequestIterationStatus(ctx context.Context, repositoryId *string, pullRequestId *int, iterationId *int, statusId *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)
    if statusId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "statusId"} 
    }
    routeValues["statusId"] = strconv.Itoa(*statusId)

    locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the specific pull request iteration status by ID. The status ID is unique within the pull request across all iterations.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the pull request iteration.
// statusId (required): ID of the pull request status.
// project (optional): Project ID or project name
func (client Client) GetPullRequestIterationStatus(ctx context.Context, repositoryId *string, pullRequestId *int, iterationId *int, statusId *int, project *string) (*GitPullRequestStatus, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)
    if statusId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "statusId"} 
    }
    routeValues["statusId"] = strconv.Itoa(*statusId)

    locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestStatus
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all the statuses associated with a pull request iteration.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the pull request iteration.
// project (optional): Project ID or project name
func (client Client) GetPullRequestIterationStatuses(ctx context.Context, repositoryId *string, pullRequestId *int, iterationId *int, project *string) (*[]GitPullRequestStatus, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitPullRequestStatus
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update pull request iteration statuses collection. The only supported operation type is `remove`.
// ctx
// patchDocument (required): Operations to apply to the pull request statuses in JSON Patch format.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// iterationId (required): ID of the pull request iteration.
// project (optional): Project ID or project name
func (client Client) UpdatePullRequestIterationStatuses(ctx context.Context, patchDocument *[]JsonPatchOperation, repositoryId *string, pullRequestId *int, iterationId *int, project *string) error {
    if patchDocument == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "patchDocument"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if iterationId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    body, marshalErr := json.Marshal(*patchDocument)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Create a label for a specified pull request. The only required field is the name of the new label.
// ctx
// label (required): Label to assign to the pull request.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
// projectId (optional): Project ID or project name.
func (client Client) CreatePullRequestLabel(ctx context.Context, label *WebApiCreateTagRequestData, repositoryId *string, pullRequestId *int, project *string, projectId *string) (*WebApiTagDefinition, error) {
    if label == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "label"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    queryParams := url.Values{}
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    body, marshalErr := json.Marshal(*label)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WebApiTagDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a label from the set of those assigned to the pull request.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// labelIdOrName (required): The name or ID of the label requested.
// project (optional): Project ID or project name
// projectId (optional): Project ID or project name.
func (client Client) DeletePullRequestLabels(ctx context.Context, repositoryId *string, pullRequestId *int, labelIdOrName *string, project *string, projectId *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if labelIdOrName == nil || *labelIdOrName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "labelIdOrName"} 
    }
    routeValues["labelIdOrName"] = *labelIdOrName

    queryParams := url.Values{}
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Retrieves a single label that has been assigned to a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// labelIdOrName (required): The name or ID of the label requested.
// project (optional): Project ID or project name
// projectId (optional): Project ID or project name.
func (client Client) GetPullRequestLabel(ctx context.Context, repositoryId *string, pullRequestId *int, labelIdOrName *string, project *string, projectId *string) (*WebApiTagDefinition, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if labelIdOrName == nil || *labelIdOrName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "labelIdOrName"} 
    }
    routeValues["labelIdOrName"] = *labelIdOrName

    queryParams := url.Values{}
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WebApiTagDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all the labels assigned to a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
// projectId (optional): Project ID or project name.
func (client Client) GetPullRequestLabels(ctx context.Context, repositoryId *string, pullRequestId *int, project *string, projectId *string) (*[]WebApiTagDefinition, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    queryParams := url.Values{}
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WebApiTagDefinition
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get external properties of the pull request.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetPullRequestProperties(ctx context.Context, repositoryId *string, pullRequestId *int, project *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("48a52185-5b9e-4736-9dc1-bb1e2feac80b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Create or update pull request external properties. The patch operation can be `add`, `replace` or `remove`. For `add` operation, the path can be empty. If the path is empty, the value must be a list of key value pairs. For `replace` operation, the path cannot be empty. If the path does not exist, the property will be added to the collection. For `remove` operation, the path cannot be empty. If the path does not exist, no action will be performed.
// ctx
// patchDocument (required): Properties to add, replace or remove in JSON Patch format.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) UpdatePullRequestProperties(ctx context.Context, patchDocument *[]JsonPatchOperation, repositoryId *string, pullRequestId *int, project *string) (interface{}, error) {
    if patchDocument == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "patchDocument"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*patchDocument)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("48a52185-5b9e-4736-9dc1-bb1e2feac80b")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// This API is used to find what pull requests are related to a given commit.  It can be used to either find the pull request that created a particular merge commit or it can be used to find all pull requests that have ever merged a particular commit.  The input is a list of queries which each contain a list of commits. For each commit that you search against, you will get back a dictionary of commit -> pull requests.
// ctx
// queries (required): The list of queries to perform.
// repositoryId (required): ID of the repository.
// project (optional): Project ID or project name
func (client Client) GetPullRequestQuery(ctx context.Context, queries *GitPullRequestQuery, repositoryId *string, project *string) (*GitPullRequestQuery, error) {
    if queries == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queries"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*queries)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b3a6eebe-9cf0-49ea-b6cb-1a4c5f5007b0")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Add a reviewer to a pull request or cast a vote.
// ctx
// reviewer (required): Reviewer's vote.<br />If the reviewer's ID is included here, it must match the reviewerID parameter.<br />Reviewers can set their own vote with this method.  When adding other reviewers, vote must be set to zero.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// reviewerId (required): ID of the reviewer.
// project (optional): Project ID or project name
func (client Client) CreatePullRequestReviewer(ctx context.Context, reviewer *IdentityRefWithVote, repositoryId *string, pullRequestId *int, reviewerId *string, project *string) (*IdentityRefWithVote, error) {
    if reviewer == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "reviewer"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if reviewerId == nil || *reviewerId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "reviewerId"} 
    }
    routeValues["reviewerId"] = *reviewerId

    body, marshalErr := json.Marshal(*reviewer)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityRefWithVote
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Add reviewers to a pull request.
// ctx
// reviewers (required): Reviewers to add to the pull request.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) CreatePullRequestReviewers(ctx context.Context, reviewers *[]IdentityRef, repositoryId *string, pullRequestId *int, project *string) (*[]IdentityRefWithVote, error) {
    if reviewers == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "reviewers"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*reviewers)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []IdentityRefWithVote
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Remove a reviewer from a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// reviewerId (required): ID of the reviewer to remove.
// project (optional): Project ID or project name
func (client Client) DeletePullRequestReviewer(ctx context.Context, repositoryId *string, pullRequestId *int, reviewerId *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if reviewerId == nil || *reviewerId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "reviewerId"} 
    }
    routeValues["reviewerId"] = *reviewerId

    locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Retrieve information about a particular reviewer on a pull request
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// reviewerId (required): ID of the reviewer.
// project (optional): Project ID or project name
func (client Client) GetPullRequestReviewer(ctx context.Context, repositoryId *string, pullRequestId *int, reviewerId *string, project *string) (*IdentityRefWithVote, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if reviewerId == nil || *reviewerId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "reviewerId"} 
    }
    routeValues["reviewerId"] = *reviewerId

    locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue IdentityRefWithVote
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve the reviewers for a pull request
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetPullRequestReviewers(ctx context.Context, repositoryId *string, pullRequestId *int, project *string) (*[]IdentityRefWithVote, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []IdentityRefWithVote
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Reset the votes of multiple reviewers on a pull request.  NOTE: This endpoint only supports updating votes, but does not support updating required reviewers (use policy) or display names.
// ctx
// patchVotes (required): IDs of the reviewers whose votes will be reset to zero
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request
// project (optional): Project ID or project name
func (client Client) UpdatePullRequestReviewers(ctx context.Context, patchVotes *[]IdentityRefWithVote, repositoryId *string, pullRequestId *int, project *string) error {
    if patchVotes == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "patchVotes"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*patchVotes)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Retrieve a pull request.
// ctx
// pullRequestId (required): The ID of the pull request to retrieve.
// project (optional): Project ID or project name
func (client Client) GetPullRequestById(ctx context.Context, pullRequestId *int, project *string) (*GitPullRequest, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("01a46dea-7d46-4d40-bc84-319e7c260d99")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve all pull requests matching a specified criteria.
// ctx
// project (required): Project ID or project name
// searchCriteria (required): Pull requests will be returned that match this search criteria.
// maxCommentLength (optional): Not used.
// skip (optional): The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
// top (optional): The number of pull requests to retrieve.
func (client Client) GetPullRequestsByProject(ctx context.Context, project *string, searchCriteria *GitPullRequestSearchCriteria, maxCommentLength *int, skip *int, top *int) (*[]GitPullRequest, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if searchCriteria == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "searchCriteria"}
    }
    if searchCriteria.RepositoryId != nil {
        queryParams.Add("searchCriteria.repositoryId", (*searchCriteria.RepositoryId).String())
    }
    if searchCriteria.CreatorId != nil {
        queryParams.Add("searchCriteria.creatorId", (*searchCriteria.CreatorId).String())
    }
    if searchCriteria.ReviewerId != nil {
        queryParams.Add("searchCriteria.reviewerId", (*searchCriteria.ReviewerId).String())
    }
    if searchCriteria.Status != nil {
        queryParams.Add("searchCriteria.status", string(*searchCriteria.Status))
    }
    if searchCriteria.TargetRefName != nil {
        queryParams.Add("searchCriteria.targetRefName", *searchCriteria.TargetRefName)
    }
    if searchCriteria.SourceRepositoryId != nil {
        queryParams.Add("searchCriteria.sourceRepositoryId", (*searchCriteria.SourceRepositoryId).String())
    }
    if searchCriteria.SourceRefName != nil {
        queryParams.Add("searchCriteria.sourceRefName", *searchCriteria.SourceRefName)
    }
    if searchCriteria.IncludeLinks != nil {
        queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*searchCriteria.IncludeLinks))
    }
    if maxCommentLength != nil {
        queryParams.Add("maxCommentLength", strconv.Itoa(*maxCommentLength))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("a5d28130-9cd2-40fa-9f08-902e7daa9efb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitPullRequest
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a pull request.
// ctx
// gitPullRequestToCreate (required): The pull request to create.
// repositoryId (required): The repository ID of the pull request's target branch.
// project (optional): Project ID or project name
// supportsIterations (optional): If true, subsequent pushes to the pull request will be individually reviewable. Set this to false for large pull requests for performance reasons if this functionality is not needed.
func (client Client) CreatePullRequest(ctx context.Context, gitPullRequestToCreate *GitPullRequest, repositoryId *string, project *string, supportsIterations *bool) (*GitPullRequest, error) {
    if gitPullRequestToCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "gitPullRequestToCreate"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if supportsIterations != nil {
        queryParams.Add("supportsIterations", strconv.FormatBool(*supportsIterations))
    }
    body, marshalErr := json.Marshal(*gitPullRequestToCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): The ID of the pull request to retrieve.
// project (optional): Project ID or project name
// maxCommentLength (optional): Not used.
// skip (optional): Not used.
// top (optional): Not used.
// includeCommits (optional): If true, the pull request will be returned with the associated commits.
// includeWorkItemRefs (optional): If true, the pull request will be returned with the associated work item references.
func (client Client) GetPullRequest(ctx context.Context, repositoryId *string, pullRequestId *int, project *string, maxCommentLength *int, skip *int, top *int, includeCommits *bool, includeWorkItemRefs *bool) (*GitPullRequest, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

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
    if includeCommits != nil {
        queryParams.Add("includeCommits", strconv.FormatBool(*includeCommits))
    }
    if includeWorkItemRefs != nil {
        queryParams.Add("includeWorkItemRefs", strconv.FormatBool(*includeWorkItemRefs))
    }
    locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve all pull requests matching a specified criteria.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// searchCriteria (required): Pull requests will be returned that match this search criteria.
// project (optional): Project ID or project name
// maxCommentLength (optional): Not used.
// skip (optional): The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
// top (optional): The number of pull requests to retrieve.
func (client Client) GetPullRequests(ctx context.Context, repositoryId *string, searchCriteria *GitPullRequestSearchCriteria, project *string, maxCommentLength *int, skip *int, top *int) (*[]GitPullRequest, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if searchCriteria == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "searchCriteria"}
    }
    if searchCriteria.RepositoryId != nil {
        queryParams.Add("searchCriteria.repositoryId", (*searchCriteria.RepositoryId).String())
    }
    if searchCriteria.CreatorId != nil {
        queryParams.Add("searchCriteria.creatorId", (*searchCriteria.CreatorId).String())
    }
    if searchCriteria.ReviewerId != nil {
        queryParams.Add("searchCriteria.reviewerId", (*searchCriteria.ReviewerId).String())
    }
    if searchCriteria.Status != nil {
        queryParams.Add("searchCriteria.status", string(*searchCriteria.Status))
    }
    if searchCriteria.TargetRefName != nil {
        queryParams.Add("searchCriteria.targetRefName", *searchCriteria.TargetRefName)
    }
    if searchCriteria.SourceRepositoryId != nil {
        queryParams.Add("searchCriteria.sourceRepositoryId", (*searchCriteria.SourceRepositoryId).String())
    }
    if searchCriteria.SourceRefName != nil {
        queryParams.Add("searchCriteria.sourceRefName", *searchCriteria.SourceRefName)
    }
    if searchCriteria.IncludeLinks != nil {
        queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*searchCriteria.IncludeLinks))
    }
    if maxCommentLength != nil {
        queryParams.Add("maxCommentLength", strconv.Itoa(*maxCommentLength))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitPullRequest
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a pull request
// ctx
// gitPullRequestToUpdate (required): The pull request content that should be updated.
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request to update.
// project (optional): Project ID or project name
func (client Client) UpdatePullRequest(ctx context.Context, gitPullRequestToUpdate *GitPullRequest, repositoryId *string, pullRequestId *int, project *string) (*GitPullRequest, error) {
    if gitPullRequestToUpdate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "gitPullRequestToUpdate"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*gitPullRequestToUpdate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Sends an e-mail notification about a specific pull request to a set of recipients
// ctx
// userMessage (required)
// repositoryId (required): ID of the git repository.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) SharePullRequest(ctx context.Context, userMessage *ShareNotificationContext, repositoryId *string, pullRequestId *int, project *string) error {
    if userMessage == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "userMessage"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*userMessage)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("696f3a82-47c9-487f-9117-b9d00972ca84")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Create a pull request status.
// ctx
// status (required): Pull request status to create.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) CreatePullRequestStatus(ctx context.Context, status *GitPullRequestStatus, repositoryId *string, pullRequestId *int, project *string) (*GitPullRequestStatus, error) {
    if status == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "status"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*status)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestStatus
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete pull request status.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// statusId (required): ID of the pull request status.
// project (optional): Project ID or project name
func (client Client) DeletePullRequestStatus(ctx context.Context, repositoryId *string, pullRequestId *int, statusId *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if statusId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "statusId"} 
    }
    routeValues["statusId"] = strconv.Itoa(*statusId)

    locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the specific pull request status by ID. The status ID is unique within the pull request across all iterations.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// statusId (required): ID of the pull request status.
// project (optional): Project ID or project name
func (client Client) GetPullRequestStatus(ctx context.Context, repositoryId *string, pullRequestId *int, statusId *int, project *string) (*GitPullRequestStatus, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if statusId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "statusId"} 
    }
    routeValues["statusId"] = strconv.Itoa(*statusId)

    locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestStatus
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all the statuses associated with a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetPullRequestStatuses(ctx context.Context, repositoryId *string, pullRequestId *int, project *string) (*[]GitPullRequestStatus, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitPullRequestStatus
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update pull request statuses collection. The only supported operation type is `remove`.
// ctx
// patchDocument (required): Operations to apply to the pull request statuses in JSON Patch format.
// repositoryId (required): The repository ID of the pull request’s target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) UpdatePullRequestStatuses(ctx context.Context, patchDocument *[]JsonPatchOperation, repositoryId *string, pullRequestId *int, project *string) error {
    if patchDocument == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "patchDocument"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*patchDocument)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Create a comment on a specific thread in a pull request (up to 500 comments can be created per thread).
// ctx
// comment (required): The comment to create. Comments can be up to 150,000 characters.
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): ID of the thread that the desired comment is in.
// project (optional): Project ID or project name
func (client Client) CreateComment(ctx context.Context, comment *Comment, repositoryId *string, pullRequestId *int, threadId *int, project *string) (*Comment, error) {
    if comment == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "comment"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)

    body, marshalErr := json.Marshal(*comment)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Comment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a comment associated with a specific thread in a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): ID of the thread that the desired comment is in.
// commentId (required): ID of the comment.
// project (optional): Project ID or project name
func (client Client) DeleteComment(ctx context.Context, repositoryId *string, pullRequestId *int, threadId *int, commentId *int, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)
    if commentId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Retrieve a comment associated with a specific thread in a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): ID of the thread that the desired comment is in.
// commentId (required): ID of the comment.
// project (optional): Project ID or project name
func (client Client) GetComment(ctx context.Context, repositoryId *string, pullRequestId *int, threadId *int, commentId *int, project *string) (*Comment, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Comment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve all comments associated with a specific thread in a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): ID of the thread.
// project (optional): Project ID or project name
func (client Client) GetComments(ctx context.Context, repositoryId *string, pullRequestId *int, threadId *int, project *string) (*[]Comment, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)

    locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Comment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a comment associated with a specific thread in a pull request.
// ctx
// comment (required): The comment content that should be updated. Comments can be up to 150,000 characters.
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): ID of the thread that the desired comment is in.
// commentId (required): ID of the comment to update.
// project (optional): Project ID or project name
func (client Client) UpdateComment(ctx context.Context, comment *Comment, repositoryId *string, pullRequestId *int, threadId *int, commentId *int, project *string) (*Comment, error) {
    if comment == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "comment"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)
    if commentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentId"} 
    }
    routeValues["commentId"] = strconv.Itoa(*commentId)

    body, marshalErr := json.Marshal(*comment)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Comment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Create a thread in a pull request.
// ctx
// commentThread (required): The thread to create. Thread must contain at least one comment.
// repositoryId (required): Repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) CreateThread(ctx context.Context, commentThread *GitPullRequestCommentThread, repositoryId *string, pullRequestId *int, project *string) (*GitPullRequestCommentThread, error) {
    if commentThread == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentThread"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    body, marshalErr := json.Marshal(*commentThread)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestCommentThread
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a thread in a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): ID of the thread.
// project (optional): Project ID or project name
// iteration (optional): If specified, thread position will be tracked using this iteration as the right side of the diff.
// baseIteration (optional): If specified, thread position will be tracked using this iteration as the left side of the diff.
func (client Client) GetPullRequestThread(ctx context.Context, repositoryId *string, pullRequestId *int, threadId *int, project *string, iteration *int, baseIteration *int) (*GitPullRequestCommentThread, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)

    queryParams := url.Values{}
    if iteration != nil {
        queryParams.Add("$iteration", strconv.Itoa(*iteration))
    }
    if baseIteration != nil {
        queryParams.Add("$baseIteration", strconv.Itoa(*baseIteration))
    }
    locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestCommentThread
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve all threads in a pull request.
// ctx
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
// iteration (optional): If specified, thread positions will be tracked using this iteration as the right side of the diff.
// baseIteration (optional): If specified, thread positions will be tracked using this iteration as the left side of the diff.
func (client Client) GetThreads(ctx context.Context, repositoryId *string, pullRequestId *int, project *string, iteration *int, baseIteration *int) (*[]GitPullRequestCommentThread, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    queryParams := url.Values{}
    if iteration != nil {
        queryParams.Add("$iteration", strconv.Itoa(*iteration))
    }
    if baseIteration != nil {
        queryParams.Add("$baseIteration", strconv.Itoa(*baseIteration))
    }
    locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitPullRequestCommentThread
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a thread in a pull request.
// ctx
// commentThread (required): The thread content that should be updated.
// repositoryId (required): The repository ID of the pull request's target branch.
// pullRequestId (required): ID of the pull request.
// threadId (required): ID of the thread to update.
// project (optional): Project ID or project name
func (client Client) UpdateThread(ctx context.Context, commentThread *GitPullRequestCommentThread, repositoryId *string, pullRequestId *int, threadId *int, project *string) (*GitPullRequestCommentThread, error) {
    if commentThread == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "commentThread"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)
    if threadId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "threadId"} 
    }
    routeValues["threadId"] = strconv.Itoa(*threadId)

    body, marshalErr := json.Marshal(*commentThread)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPullRequestCommentThread
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a list of work items associated with a pull request.
// ctx
// repositoryId (required): ID or name of the repository.
// pullRequestId (required): ID of the pull request.
// project (optional): Project ID or project name
func (client Client) GetPullRequestWorkItemRefs(ctx context.Context, repositoryId *string, pullRequestId *int, project *string) (*[]ResourceRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pullRequestId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pullRequestId"} 
    }
    routeValues["pullRequestId"] = strconv.Itoa(*pullRequestId)

    locationId, _ := uuid.Parse("0a637fcc-5370-4ce8-b0e8-98091f5f9482")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ResourceRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Push changes to the repository.
// ctx
// push (required)
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
func (client Client) CreatePush(ctx context.Context, push *GitPush, repositoryId *string, project *string) (*GitPush, error) {
    if push == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "push"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*push)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ea98d07b-3c87-4971-8ede-a613694ffb55")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPush
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieves a particular push.
// ctx
// repositoryId (required): The name or ID of the repository.
// pushId (required): ID of the push.
// project (optional): Project ID or project name
// includeCommits (optional): The number of commits to include in the result.
// includeRefUpdates (optional): If true, include the list of refs that were updated by the push.
func (client Client) GetPush(ctx context.Context, repositoryId *string, pushId *int, project *string, includeCommits *int, includeRefUpdates *bool) (*GitPush, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if pushId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pushId"} 
    }
    routeValues["pushId"] = strconv.Itoa(*pushId)

    queryParams := url.Values{}
    if includeCommits != nil {
        queryParams.Add("includeCommits", strconv.Itoa(*includeCommits))
    }
    if includeRefUpdates != nil {
        queryParams.Add("includeRefUpdates", strconv.FormatBool(*includeRefUpdates))
    }
    locationId, _ := uuid.Parse("ea98d07b-3c87-4971-8ede-a613694ffb55")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitPush
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieves pushes associated with the specified repository.
// ctx
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// skip (optional): Number of pushes to skip.
// top (optional): Number of pushes to return.
// searchCriteria (optional): Search criteria attributes: fromDate, toDate, pusherId, refName, includeRefUpdates or includeLinks. fromDate: Start date to search from. toDate: End date to search to. pusherId: Identity of the person who submitted the push. refName: Branch name to consider. includeRefUpdates: If true, include the list of refs that were updated by the push. includeLinks: Whether to include the _links field on the shallow references.
func (client Client) GetPushes(ctx context.Context, repositoryId *string, project *string, skip *int, top *int, searchCriteria *GitPushSearchCriteria) (*[]GitPush, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if searchCriteria != nil {
        if searchCriteria.FromDate != nil {
            queryParams.Add("searchCriteria.fromDate", (*searchCriteria.FromDate).String())
        }
        if searchCriteria.ToDate != nil {
            queryParams.Add("searchCriteria.toDate", (*searchCriteria.ToDate).String())
        }
        if searchCriteria.PusherId != nil {
            queryParams.Add("searchCriteria.pusherId", (*searchCriteria.PusherId).String())
        }
        if searchCriteria.RefName != nil {
            queryParams.Add("searchCriteria.refName", *searchCriteria.RefName)
        }
        if searchCriteria.IncludeRefUpdates != nil {
            queryParams.Add("searchCriteria.includeRefUpdates", strconv.FormatBool(*searchCriteria.IncludeRefUpdates))
        }
        if searchCriteria.IncludeLinks != nil {
            queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*searchCriteria.IncludeLinks))
        }
    }
    locationId, _ := uuid.Parse("ea98d07b-3c87-4971-8ede-a613694ffb55")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitPush
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Destroy (hard delete) a soft-deleted Git repository.
// ctx
// project (required): Project ID or project name
// repositoryId (required): The ID of the repository.
func (client Client) DeleteRepositoryFromRecycleBin(ctx context.Context, project *string, repositoryId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = (*repositoryId).String()

    locationId, _ := uuid.Parse("a663da97-81db-4eb3-8b83-287670f63073")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Retrieve soft-deleted git repositories from the recycle bin.
// ctx
// project (required): Project ID or project name
func (client Client) GetRecycleBinRepositories(ctx context.Context, project *string) (*[]GitDeletedRepository, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("a663da97-81db-4eb3-8b83-287670f63073")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitDeletedRepository
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Recover a soft-deleted Git repository. Recently deleted repositories go into a soft-delete state for a period of time before they are hard deleted and become unrecoverable.
// ctx
// repositoryDetails (required)
// project (required): Project ID or project name
// repositoryId (required): The ID of the repository.
func (client Client) RestoreRepositoryFromRecycleBin(ctx context.Context, repositoryDetails *GitRecycleBinRepositoryDetails, project *string, repositoryId *uuid.UUID) (*GitRepository, error) {
    if repositoryDetails == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "repositoryDetails"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = (*repositoryId).String()

    body, marshalErr := json.Marshal(*repositoryDetails)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a663da97-81db-4eb3-8b83-287670f63073")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRepository
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Queries the provided repository for its refs and returns them.
// ctx
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// filter (optional): [optional] A filter to apply to the refs (starts with).
// includeLinks (optional): [optional] Specifies if referenceLinks should be included in the result. default is false.
// includeStatuses (optional): [optional] Includes up to the first 1000 commit statuses for each ref. The default value is false.
// includeMyBranches (optional): [optional] Includes only branches that the user owns, the branches the user favorites, and the default branch. The default value is false. Cannot be combined with the filter parameter.
// latestStatusesOnly (optional): [optional] True to include only the tip commit status for each ref. This option requires `includeStatuses` to be true. The default value is false.
// peelTags (optional): [optional] Annotated tags will populate the PeeledObjectId property. default is false.
// filterContains (optional): [optional] A filter to apply to the refs (contains).
func (client Client) GetRefs(ctx context.Context, repositoryId *string, project *string, filter *string, includeLinks *bool, includeStatuses *bool, includeMyBranches *bool, latestStatusesOnly *bool, peelTags *bool, filterContains *string) (*[]GitRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if filter != nil {
        queryParams.Add("filter", *filter)
    }
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    if includeStatuses != nil {
        queryParams.Add("includeStatuses", strconv.FormatBool(*includeStatuses))
    }
    if includeMyBranches != nil {
        queryParams.Add("includeMyBranches", strconv.FormatBool(*includeMyBranches))
    }
    if latestStatusesOnly != nil {
        queryParams.Add("latestStatusesOnly", strconv.FormatBool(*latestStatusesOnly))
    }
    if peelTags != nil {
        queryParams.Add("peelTags", strconv.FormatBool(*peelTags))
    }
    if filterContains != nil {
        queryParams.Add("filterContains", *filterContains)
    }
    locationId, _ := uuid.Parse("2d874a60-a811-4f62-9c9f-963a6ea0a55b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Lock or Unlock a branch.
// ctx
// newRefInfo (required): The ref update action (lock/unlock) to perform
// repositoryId (required): The name or ID of the repository.
// filter (required): The name of the branch to lock/unlock
// project (optional): Project ID or project name
// projectId (optional): ID or name of the team project. Optional if specifying an ID for repository.
func (client Client) UpdateRef(ctx context.Context, newRefInfo *GitRefUpdate, repositoryId *string, filter *string, project *string, projectId *string) (*GitRef, error) {
    if newRefInfo == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "newRefInfo"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if filter == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filter"}
    }
    queryParams.Add("filter", *filter)
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    body, marshalErr := json.Marshal(*newRefInfo)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2d874a60-a811-4f62-9c9f-963a6ea0a55b")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRef
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Creating, updating, or deleting refs(branches).
// ctx
// refUpdates (required): List of ref updates to attempt to perform
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
// projectId (optional): ID or name of the team project. Optional if specifying an ID for repository.
func (client Client) UpdateRefs(ctx context.Context, refUpdates *[]GitRefUpdate, repositoryId *string, project *string, projectId *string) (*[]GitRefUpdateResult, error) {
    if refUpdates == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "refUpdates"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    body, marshalErr := json.Marshal(*refUpdates)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2d874a60-a811-4f62-9c9f-963a6ea0a55b")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitRefUpdateResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a ref favorite
// ctx
// favorite (required): The ref favorite to create.
// project (required): Project ID or project name
func (client Client) CreateFavorite(ctx context.Context, favorite *GitRefFavorite, project *string) (*GitRefFavorite, error) {
    if favorite == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "favorite"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*favorite)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRefFavorite
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes the refs favorite specified
// ctx
// project (required): Project ID or project name
// favoriteId (required): The Id of the ref favorite to delete.
func (client Client) DeleteRefFavorite(ctx context.Context, project *string, favoriteId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if favoriteId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "favoriteId"} 
    }
    routeValues["favoriteId"] = strconv.Itoa(*favoriteId)

    locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Gets the refs favorite for a favorite Id.
// ctx
// project (required): Project ID or project name
// favoriteId (required): The Id of the requested ref favorite.
func (client Client) GetRefFavorite(ctx context.Context, project *string, favoriteId *int) (*GitRefFavorite, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if favoriteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "favoriteId"} 
    }
    routeValues["favoriteId"] = strconv.Itoa(*favoriteId)

    locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRefFavorite
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the refs favorites for a repo and an identity.
// ctx
// project (required): Project ID or project name
// repositoryId (optional): The id of the repository.
// identityId (optional): The id of the identity whose favorites are to be retrieved. If null, the requesting identity is used.
func (client Client) GetRefFavorites(ctx context.Context, project *string, repositoryId *string, identityId *string) (*[]GitRefFavorite, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if repositoryId != nil {
        queryParams.Add("repositoryId", *repositoryId)
    }
    if identityId != nil {
        queryParams.Add("identityId", *identityId)
    }
    locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitRefFavorite
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a git repository in a team project.
// ctx
// gitRepositoryToCreate (required): Specify the repo name, team project and/or parent repository. Team project information can be ommitted from gitRepositoryToCreate if the request is project-scoped (i.e., includes project Id).
// project (optional): Project ID or project name
// sourceRef (optional): [optional] Specify the source refs to use while creating a fork repo
func (client Client) CreateRepository(ctx context.Context, gitRepositoryToCreate *GitRepositoryCreateOptions, project *string, sourceRef *string) (*GitRepository, error) {
    if gitRepositoryToCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "gitRepositoryToCreate"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if sourceRef != nil {
        queryParams.Add("sourceRef", *sourceRef)
    }
    body, marshalErr := json.Marshal(*gitRepositoryToCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRepository
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a git repository
// ctx
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
func (client Client) DeleteRepository(ctx context.Context, repositoryId *uuid.UUID, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = (*repositoryId).String()

    locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Retrieve git repositories.
// ctx
// project (optional): Project ID or project name
// includeLinks (optional): [optional] True to include reference links. The default value is false.
// includeAllUrls (optional): [optional] True to include all remote URLs. The default value is false.
// includeHidden (optional): [optional] True to include hidden repositories. The default value is false.
func (client Client) GetRepositories(ctx context.Context, project *string, includeLinks *bool, includeAllUrls *bool, includeHidden *bool) (*[]GitRepository, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if includeLinks != nil {
        queryParams.Add("includeLinks", strconv.FormatBool(*includeLinks))
    }
    if includeAllUrls != nil {
        queryParams.Add("includeAllUrls", strconv.FormatBool(*includeAllUrls))
    }
    if includeHidden != nil {
        queryParams.Add("includeHidden", strconv.FormatBool(*includeHidden))
    }
    locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitRepository
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a git repository.
// ctx
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
func (client Client) GetRepository(ctx context.Context, repositoryId *string, project *string) (*GitRepository, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRepository
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Retrieve a git repository.
// ctx
// repositoryId (required): The name or ID of the repository.
// includeParent (required): True to include parent repository. Only available in authenticated calls.
// project (optional): Project ID or project name
func (client Client) GetRepositoryWithParent(ctx context.Context, repositoryId *string, includeParent *bool, project *string) (*GitRepository, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if includeParent == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "includeParent"}
    }
    queryParams.Add("includeParent", strconv.FormatBool(*includeParent))
    locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRepository
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Updates the Git repository with either a new repo name or a new default branch.
// ctx
// newRepositoryInfo (required): Specify a new repo name or a new default branch of the repository
// repositoryId (required): The name or ID of the repository.
// project (optional): Project ID or project name
func (client Client) UpdateRepository(ctx context.Context, newRepositoryInfo *GitRepository, repositoryId *uuid.UUID, project *string) (*GitRepository, error) {
    if newRepositoryInfo == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "newRepositoryInfo"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = (*repositoryId).String()

    body, marshalErr := json.Marshal(*newRepositoryInfo)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRepository
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Starts the operation to create a new branch which reverts changes introduced by either a specific commit or commits that are associated to a pull request.
// ctx
// revertToCreate (required)
// project (required): Project ID or project name
// repositoryId (required): ID of the repository.
func (client Client) CreateRevert(ctx context.Context, revertToCreate *GitAsyncRefOperationParameters, project *string, repositoryId *string) (*GitRevert, error) {
    if revertToCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "revertToCreate"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*revertToCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bc866058-5449-4715-9cf1-a510b6ff193c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRevert
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve information about a revert operation by revert Id.
// ctx
// project (required): Project ID or project name
// revertId (required): ID of the revert operation.
// repositoryId (required): ID of the repository.
func (client Client) GetRevert(ctx context.Context, project *string, revertId *int, repositoryId *string) (*GitRevert, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if revertId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "revertId"} 
    }
    routeValues["revertId"] = strconv.Itoa(*revertId)
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    locationId, _ := uuid.Parse("bc866058-5449-4715-9cf1-a510b6ff193c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRevert
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve information about a revert operation for a specific branch.
// ctx
// project (required): Project ID or project name
// repositoryId (required): ID of the repository.
// refName (required): The GitAsyncRefOperationParameters generatedRefName used for the revert operation.
func (client Client) GetRevertForRefName(ctx context.Context, project *string, repositoryId *string, refName *string) (*GitRevert, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if refName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "refName"}
    }
    queryParams.Add("refName", *refName)
    locationId, _ := uuid.Parse("bc866058-5449-4715-9cf1-a510b6ff193c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitRevert
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Create Git commit status.
// ctx
// gitCommitStatusToCreate (required): Git commit status object to create.
// commitId (required): ID of the Git commit.
// repositoryId (required): ID of the repository.
// project (optional): Project ID or project name
func (client Client) CreateCommitStatus(ctx context.Context, gitCommitStatusToCreate *GitStatus, commitId *string, repositoryId *string, project *string) (*GitStatus, error) {
    if gitCommitStatusToCreate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "gitCommitStatusToCreate"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if commitId == nil || *commitId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "commitId"} 
    }
    routeValues["commitId"] = *commitId
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    body, marshalErr := json.Marshal(*gitCommitStatusToCreate)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("428dd4fb-fda5-4722-af02-9313b80305da")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitStatus
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get statuses associated with the Git commit.
// ctx
// commitId (required): ID of the Git commit.
// repositoryId (required): ID of the repository.
// project (optional): Project ID or project name
// top (optional): Optional. The number of statuses to retrieve. Default is 1000.
// skip (optional): Optional. The number of statuses to ignore. Default is 0. For example, to retrieve results 101-150, set top to 50 and skip to 100.
// latestOnly (optional): The flag indicates whether to get only latest statuses grouped by `Context.Name` and `Context.Genre`.
func (client Client) GetStatuses(ctx context.Context, commitId *string, repositoryId *string, project *string, top *int, skip *int, latestOnly *bool) (*[]GitStatus, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if commitId == nil || *commitId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "commitId"} 
    }
    routeValues["commitId"] = *commitId
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("skip", strconv.Itoa(*skip))
    }
    if latestOnly != nil {
        queryParams.Add("latestOnly", strconv.FormatBool(*latestOnly))
    }
    locationId, _ := uuid.Parse("428dd4fb-fda5-4722-af02-9313b80305da")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitStatus
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve a pull request suggestion for a particular repository or team project.
// ctx
// repositoryId (required): ID of the git repository.
// project (optional): Project ID or project name
func (client Client) GetSuggestions(ctx context.Context, repositoryId *string, project *string) (*[]GitSuggestion, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId

    locationId, _ := uuid.Parse("9393b4fb-4445-4919-972b-9ad16f442d83")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GitSuggestion
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// The Tree endpoint returns the collection of objects underneath the specified tree. Trees are folders in a Git repository.
// ctx
// repositoryId (required): Repository Id.
// sha1 (required): SHA1 hash of the tree object.
// project (optional): Project ID or project name
// projectId (optional): Project Id.
// recursive (optional): Search recursively. Include trees underneath this tree. Default is false.
// fileName (optional): Name to use if a .zip file is returned. Default is the object ID.
func (client Client) GetTree(ctx context.Context, repositoryId *string, sha1 *string, project *string, projectId *string, recursive *bool, fileName *string) (*GitTreeRef, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if sha1 == nil || *sha1 == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sha1"} 
    }
    routeValues["sha1"] = *sha1

    queryParams := url.Values{}
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    if recursive != nil {
        queryParams.Add("recursive", strconv.FormatBool(*recursive))
    }
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    locationId, _ := uuid.Parse("729f6437-6f92-44ec-8bee-273a7111063c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue GitTreeRef
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// The Tree endpoint returns the collection of objects underneath the specified tree. Trees are folders in a Git repository.
// ctx
// repositoryId (required): Repository Id.
// sha1 (required): SHA1 hash of the tree object.
// project (optional): Project ID or project name
// projectId (optional): Project Id.
// recursive (optional): Search recursively. Include trees underneath this tree. Default is false.
// fileName (optional): Name to use if a .zip file is returned. Default is the object ID.
func (client Client) GetTreeZip(ctx context.Context, repositoryId *string, sha1 *string, project *string, projectId *string, recursive *bool, fileName *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if repositoryId == nil || *repositoryId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "repositoryId"} 
    }
    routeValues["repositoryId"] = *repositoryId
    if sha1 == nil || *sha1 == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "sha1"} 
    }
    routeValues["sha1"] = *sha1

    queryParams := url.Values{}
    if projectId != nil {
        queryParams.Add("projectId", *projectId)
    }
    if recursive != nil {
        queryParams.Add("recursive", strconv.FormatBool(*recursive))
    }
    if fileName != nil {
        queryParams.Add("fileName", *fileName)
    }
    locationId, _ := uuid.Parse("729f6437-6f92-44ec-8bee-273a7111063c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

