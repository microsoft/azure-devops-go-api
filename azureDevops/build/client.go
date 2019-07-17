// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package build

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

var ResourceAreaId, _ = uuid.Parse("965220d5-5bb9-42cf-8d67-9b146df2a5a4")

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

// Associates an artifact with a build.
// artifact (required): The artifact.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) CreateArtifact(artifact *BuildArtifact, project *string, buildId *int) (*BuildArtifact, error) {
    if artifact == nil {
        return nil, errors.New("artifact is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    body, marshalErr := json.Marshal(*artifact)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildArtifact
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a specific artifact for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// artifactName (required): The name of the artifact.
func (client Client) GetArtifact(project *string, buildId *int, artifactName *string) (*BuildArtifact, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if artifactName == nil {
        return nil, errors.New("artifactName is a required parameter")
    }
    queryParams.Add("artifactName", *artifactName)
    locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildArtifact
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a specific artifact for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// artifactName (required): The name of the artifact.
func (client Client) GetArtifactContentZip(project *string, buildId *int, artifactName *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if artifactName == nil {
        return nil, errors.New("artifactName is a required parameter")
    }
    queryParams.Add("artifactName", *artifactName)
    locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets all artifacts for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) GetArtifacts(project *string, buildId *int) (*[]BuildArtifact, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildArtifact
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a file from the build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// artifactName (required): The name of the artifact.
// fileId (required): The primary key for the file.
// fileName (required): The name that the file will be set to.
func (client Client) GetFile(project *string, buildId *int, artifactName *string, fileId *string, fileName *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if artifactName == nil {
        return nil, errors.New("artifactName is a required parameter")
    }
    queryParams.Add("artifactName", *artifactName)
    if fileId == nil {
        return nil, errors.New("fileId is a required parameter")
    }
    queryParams.Add("fileId", *fileId)
    if fileName == nil {
        return nil, errors.New("fileName is a required parameter")
    }
    queryParams.Add("fileName", *fileName)
    locationId, _ := uuid.Parse("1db06c96-014e-44e1-ac91-90b2d4b3e984")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the list of attachments of a specific type that are associated with a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// type_ (required): The type of attachment.
func (client Client) GetAttachments(project *string, buildId *int, type_ *string) (*[]Attachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("f2192269-89fa-4f94-baf6-8fb128c55159")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Attachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a specific attachment.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// timelineId (required): The ID of the timeline.
// recordId (required): The ID of the timeline record.
// type_ (required): The type of the attachment.
// name (required): The name of the attachment.
func (client Client) GetAttachment(project *string, buildId *int, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, errors.New("recordId is a required parameter")
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("af5122d3-3438-485e-a25a-2dbbfde84ee6")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// resources (required)
// project (required): Project ID or project name
func (client Client) AuthorizeProjectResources(resources *[]DefinitionResourceReference, project *string) (*[]DefinitionResourceReference, error) {
    if resources == nil {
        return nil, errors.New("resources is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*resources)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("398c85bc-81aa-4822-947c-a194a05f0fef")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DefinitionResourceReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// type_ (optional)
// id (optional)
func (client Client) GetProjectResources(project *string, type_ *string, id *string) (*[]DefinitionResourceReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    if id != nil {
        queryParams.Add("id", *id)
    }
    locationId, _ := uuid.Parse("398c85bc-81aa-4822-947c-a194a05f0fef")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DefinitionResourceReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a list of branches for the given source code repository.
// project (required): Project ID or project name
// providerName (required): The name of the source provider.
// serviceEndpointId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
// repository (optional): The vendor-specific identifier or the name of the repository to get branches. Can only be omitted for providers that do not support multiple repositories.
// branchName (optional): If supplied, the name of the branch to check for specifically.
func (client Client) ListBranches(project *string, providerName *string, serviceEndpointId *uuid.UUID, repository *string, branchName *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if providerName == nil || *providerName == "" {
        return nil, errors.New("providerName is a required parameter")
    }
    routeValues["providerName"] = *providerName

    queryParams := url.Values{}
    if serviceEndpointId != nil {
        queryParams.Add("serviceEndpointId", (*serviceEndpointId).String())
    }
    if repository != nil {
        queryParams.Add("repository", *repository)
    }
    if branchName != nil {
        queryParams.Add("branchName", *branchName)
    }
    locationId, _ := uuid.Parse("e05d4403-9b81-4244-8763-20fde28d1976")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a badge that indicates the status of the most recent build for the specified branch.
// project (required): Project ID or project name
// repoType (required): The repository type.
// repoId (optional): The repository ID.
// branchName (optional): The branch name.
func (client Client) GetBuildBadge(project *string, repoType *string, repoId *string, branchName *string) (*BuildBadge, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if repoType == nil || *repoType == "" {
        return nil, errors.New("repoType is a required parameter")
    }
    routeValues["repoType"] = *repoType

    queryParams := url.Values{}
    if repoId != nil {
        queryParams.Add("repoId", *repoId)
    }
    if branchName != nil {
        queryParams.Add("branchName", *branchName)
    }
    locationId, _ := uuid.Parse("21b3b9ce-fad5-4567-9ad0-80679794e003")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildBadge
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a badge that indicates the status of the most recent build for the specified branch.
// project (required): Project ID or project name
// repoType (required): The repository type.
// repoId (optional): The repository ID.
// branchName (optional): The branch name.
func (client Client) GetBuildBadgeData(project *string, repoType *string, repoId *string, branchName *string) (*string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if repoType == nil || *repoType == "" {
        return nil, errors.New("repoType is a required parameter")
    }
    routeValues["repoType"] = *repoType

    queryParams := url.Values{}
    if repoId != nil {
        queryParams.Add("repoId", *repoId)
    }
    if branchName != nil {
        queryParams.Add("branchName", *branchName)
    }
    locationId, _ := uuid.Parse("21b3b9ce-fad5-4567-9ad0-80679794e003")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Deletes a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) DeleteBuild(project *string, buildId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Gets a build
// project (required): Project ID or project name
// buildId (required)
// propertyFilters (optional)
func (client Client) GetBuild(project *string, buildId *int, propertyFilters *string) (*Build, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if propertyFilters != nil {
        queryParams.Add("propertyFilters", *propertyFilters)
    }
    locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Build
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a list of builds.
// project (required): Project ID or project name
// definitions (optional): A comma-delimited list of definition IDs. If specified, filters to builds for these definitions.
// queues (optional): A comma-delimited list of queue IDs. If specified, filters to builds that ran against these queues.
// buildNumber (optional): If specified, filters to builds that match this build number. Append * to do a prefix search.
// minTime (optional): If specified, filters to builds that finished/started/queued after this date based on the queryOrder specified.
// maxTime (optional): If specified, filters to builds that finished/started/queued before this date based on the queryOrder specified.
// requestedFor (optional): If specified, filters to builds requested for the specified user.
// reasonFilter (optional): If specified, filters to builds that match this reason.
// statusFilter (optional): If specified, filters to builds that match this status.
// resultFilter (optional): If specified, filters to builds that match this result.
// tagFilters (optional): A comma-delimited list of tags. If specified, filters to builds that have the specified tags.
// properties (optional): A comma-delimited list of properties to retrieve.
// top (optional): The maximum number of builds to return.
// continuationToken (optional): A continuation token, returned by a previous call to this method, that can be used to return the next set of builds.
// maxBuildsPerDefinition (optional): The maximum number of builds to return per definition.
// deletedFilter (optional): Indicates whether to exclude, include, or only return deleted builds.
// queryOrder (optional): The order in which builds should be returned.
// branchName (optional): If specified, filters to builds that built branches that built this branch.
// buildIds (optional): A comma-delimited list that specifies the IDs of builds to retrieve.
// repositoryId (optional): If specified, filters to builds that built from this repository.
// repositoryType (optional): If specified, filters to builds that built from repositories of this type.
func (client Client) GetBuilds(project *string, definitions *[]int, queues *[]int, buildNumber *string, minTime *time.Time, maxTime *time.Time, requestedFor *string, reasonFilter *string, statusFilter *string, resultFilter *string, tagFilters *[]string, properties *[]string, top *int, continuationToken *string, maxBuildsPerDefinition *int, deletedFilter *string, queryOrder *string, branchName *string, buildIds *[]int, repositoryId *string, repositoryType *string) (*[]Build, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if definitions != nil {
        var stringList []string
        for _, item := range *definitions {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if queues != nil {
        var stringList []string
        for _, item := range *queues {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if buildNumber != nil {
        queryParams.Add("buildNumber", *buildNumber)
    }
    if minTime != nil {
        queryParams.Add("minTime", (*minTime).String())
    }
    if maxTime != nil {
        queryParams.Add("maxTime", (*maxTime).String())
    }
    if requestedFor != nil {
        queryParams.Add("requestedFor", *requestedFor)
    }
    if reasonFilter != nil {
        queryParams.Add("reasonFilter", *reasonFilter)
    }
    if statusFilter != nil {
        queryParams.Add("statusFilter", *statusFilter)
    }
    if resultFilter != nil {
        queryParams.Add("resultFilter", *resultFilter)
    }
    if tagFilters != nil {
        listAsString := strings.Join((*tagFilters)[:], ",")
        queryParams.Add("tagFilters", listAsString)
    }
    if properties != nil {
        listAsString := strings.Join((*properties)[:], ",")
        queryParams.Add("properties", listAsString)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if maxBuildsPerDefinition != nil {
        queryParams.Add("maxBuildsPerDefinition", strconv.Itoa(*maxBuildsPerDefinition))
    }
    if deletedFilter != nil {
        queryParams.Add("deletedFilter", *deletedFilter)
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", *queryOrder)
    }
    if branchName != nil {
        queryParams.Add("branchName", *branchName)
    }
    if buildIds != nil {
        var stringList []string
        for _, item := range *buildIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if repositoryId != nil {
        queryParams.Add("repositoryId", *repositoryId)
    }
    if repositoryType != nil {
        queryParams.Add("repositoryType", *repositoryType)
    }
    locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Build
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Queues a build
// build (required)
// project (required): Project ID or project name
// ignoreWarnings (optional)
// checkInTicket (optional)
// sourceBuildId (optional)
func (client Client) QueueBuild(build *Build, project *string, ignoreWarnings *bool, checkInTicket *string, sourceBuildId *int) (*Build, error) {
    if build == nil {
        return nil, errors.New("build is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if ignoreWarnings != nil {
        queryParams.Add("ignoreWarnings", strconv.FormatBool(*ignoreWarnings))
    }
    if checkInTicket != nil {
        queryParams.Add("checkInTicket", *checkInTicket)
    }
    if sourceBuildId != nil {
        queryParams.Add("sourceBuildId", strconv.Itoa(*sourceBuildId))
    }
    body, marshalErr := json.Marshal(*build)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Build
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Updates a build.
// build (required): The build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// retry (optional)
func (client Client) UpdateBuild(build *Build, project *string, buildId *int, retry *bool) (*Build, error) {
    if build == nil {
        return nil, errors.New("build is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if retry != nil {
        queryParams.Add("retry", strconv.FormatBool(*retry))
    }
    body, marshalErr := json.Marshal(*build)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Build
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Updates multiple builds.
// builds (required): The builds to update.
// project (required): Project ID or project name
func (client Client) UpdateBuilds(builds *[]Build, project *string) (*[]Build, error) {
    if builds == nil {
        return nil, errors.New("builds is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*builds)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0cd358e1-9217-4d94-8269-1c1ee6f93dcf")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Build
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the changes associated with a build
// project (required): Project ID or project name
// buildId (required)
// continuationToken (optional)
// top (optional): The maximum number of changes to return
// includeSourceChange (optional)
func (client Client) GetBuildChanges(project *string, buildId *int, continuationToken *string, top *int, includeSourceChange *bool) (*[]Change, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if includeSourceChange != nil {
        queryParams.Add("includeSourceChange", strconv.FormatBool(*includeSourceChange))
    }
    locationId, _ := uuid.Parse("54572c7b-bbd3-45d4-80dc-28be08941620")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Change
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the changes made to the repository between two given builds.
// project (required): Project ID or project name
// fromBuildId (optional): The ID of the first build.
// toBuildId (optional): The ID of the last build.
// top (optional): The maximum number of changes to return.
func (client Client) GetChangesBetweenBuilds(project *string, fromBuildId *int, toBuildId *int, top *int) (*[]Change, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if fromBuildId != nil {
        queryParams.Add("fromBuildId", strconv.Itoa(*fromBuildId))
    }
    if toBuildId != nil {
        queryParams.Add("toBuildId", strconv.Itoa(*toBuildId))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("f10f0ea5-18a1-43ec-a8fb-2042c7be9b43")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Change
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a controller
// controllerId (required)
func (client Client) GetBuildController(controllerId *int) (*BuildController, error) {
    routeValues := make(map[string]string)
    if controllerId == nil {
        return nil, errors.New("controllerId is a required parameter")
    }
    routeValues["controllerId"] = strconv.Itoa(*controllerId)

    locationId, _ := uuid.Parse("fcac1932-2ee1-437f-9b6f-7f696be858f6")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildController
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets controller, optionally filtered by name
// name (optional)
func (client Client) GetBuildControllers(name *string) (*[]BuildController, error) {
    queryParams := url.Values{}
    if name != nil {
        queryParams.Add("name", *name)
    }
    locationId, _ := uuid.Parse("fcac1932-2ee1-437f-9b6f-7f696be858f6")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildController
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Creates a new definition.
// definition (required): The definition.
// project (required): Project ID or project name
// definitionToCloneId (optional)
// definitionToCloneRevision (optional)
func (client Client) CreateDefinition(definition *BuildDefinition, project *string, definitionToCloneId *int, definitionToCloneRevision *int) (*BuildDefinition, error) {
    if definition == nil {
        return nil, errors.New("definition is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if definitionToCloneId != nil {
        queryParams.Add("definitionToCloneId", strconv.Itoa(*definitionToCloneId))
    }
    if definitionToCloneRevision != nil {
        queryParams.Add("definitionToCloneRevision", strconv.Itoa(*definitionToCloneRevision))
    }
    body, marshalErr := json.Marshal(*definition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Deletes a definition and all associated builds.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
func (client Client) DeleteDefinition(project *string, definitionId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Gets a definition, optionally at a specific revision.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
// revision (optional): The revision number to retrieve. If this is not specified, the latest version will be returned.
// minMetricsTime (optional): If specified, indicates the date from which metrics should be included.
// propertyFilters (optional): A comma-delimited list of properties to include in the results.
// includeLatestBuilds (optional)
func (client Client) GetDefinition(project *string, definitionId *int, revision *int, minMetricsTime *time.Time, propertyFilters *[]string, includeLatestBuilds *bool) (*BuildDefinition, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if revision != nil {
        queryParams.Add("revision", strconv.Itoa(*revision))
    }
    if minMetricsTime != nil {
        queryParams.Add("minMetricsTime", (*minMetricsTime).String())
    }
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if includeLatestBuilds != nil {
        queryParams.Add("includeLatestBuilds", strconv.FormatBool(*includeLatestBuilds))
    }
    locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a list of definitions.
// project (required): Project ID or project name
// name (optional): If specified, filters to definitions whose names match this pattern.
// repositoryId (optional): A repository ID. If specified, filters to definitions that use this repository.
// repositoryType (optional): If specified, filters to definitions that have a repository of this type.
// queryOrder (optional): Indicates the order in which definitions should be returned.
// top (optional): The maximum number of definitions to return.
// continuationToken (optional): A continuation token, returned by a previous call to this method, that can be used to return the next set of definitions.
// minMetricsTime (optional): If specified, indicates the date from which metrics should be included.
// definitionIds (optional): A comma-delimited list that specifies the IDs of definitions to retrieve.
// path (optional): If specified, filters to definitions under this folder.
// builtAfter (optional): If specified, filters to definitions that have builds after this date.
// notBuiltAfter (optional): If specified, filters to definitions that do not have builds after this date.
// includeAllProperties (optional): Indicates whether the full definitions should be returned. By default, shallow representations of the definitions are returned.
// includeLatestBuilds (optional): Indicates whether to return the latest and latest completed builds for this definition.
// taskIdFilter (optional): If specified, filters to definitions that use the specified task.
// processType (optional): If specified, filters to definitions with the given process type.
// yamlFilename (optional): If specified, filters to YAML definitions that match the given filename.
func (client Client) GetDefinitions(project *string, name *string, repositoryId *string, repositoryType *string, queryOrder *string, top *int, continuationToken *string, minMetricsTime *time.Time, definitionIds *[]int, path *string, builtAfter *time.Time, notBuiltAfter *time.Time, includeAllProperties *bool, includeLatestBuilds *bool, taskIdFilter *uuid.UUID, processType *int, yamlFilename *string) (*[]BuildDefinitionReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if name != nil {
        queryParams.Add("name", *name)
    }
    if repositoryId != nil {
        queryParams.Add("repositoryId", *repositoryId)
    }
    if repositoryType != nil {
        queryParams.Add("repositoryType", *repositoryType)
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", *queryOrder)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if minMetricsTime != nil {
        queryParams.Add("minMetricsTime", (*minMetricsTime).String())
    }
    if definitionIds != nil {
        var stringList []string
        for _, item := range *definitionIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if path != nil {
        queryParams.Add("path", *path)
    }
    if builtAfter != nil {
        queryParams.Add("builtAfter", (*builtAfter).String())
    }
    if notBuiltAfter != nil {
        queryParams.Add("notBuiltAfter", (*notBuiltAfter).String())
    }
    if includeAllProperties != nil {
        queryParams.Add("includeAllProperties", strconv.FormatBool(*includeAllProperties))
    }
    if includeLatestBuilds != nil {
        queryParams.Add("includeLatestBuilds", strconv.FormatBool(*includeLatestBuilds))
    }
    if taskIdFilter != nil {
        queryParams.Add("taskIdFilter", (*taskIdFilter).String())
    }
    if processType != nil {
        queryParams.Add("processType", strconv.Itoa(*processType))
    }
    if yamlFilename != nil {
        queryParams.Add("yamlFilename", *yamlFilename)
    }
    locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildDefinitionReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Restores a deleted definition
// project (required): Project ID or project name
// definitionId (required): The identifier of the definition to restore.
// deleted (required): When false, restores a deleted definition.
func (client Client) RestoreDefinition(project *string, definitionId *int, deleted *bool) (*BuildDefinition, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if deleted == nil {
        return nil, errors.New("deleted is a required parameter")
    }
    queryParams.Add("deleted", strconv.FormatBool(*deleted))
    locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Updates an existing definition.
// definition (required): The new version of the defintion.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
// secretsSourceDefinitionId (optional)
// secretsSourceDefinitionRevision (optional)
func (client Client) UpdateDefinition(definition *BuildDefinition, project *string, definitionId *int, secretsSourceDefinitionId *int, secretsSourceDefinitionRevision *int) (*BuildDefinition, error) {
    if definition == nil {
        return nil, errors.New("definition is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if secretsSourceDefinitionId != nil {
        queryParams.Add("secretsSourceDefinitionId", strconv.Itoa(*secretsSourceDefinitionId))
    }
    if secretsSourceDefinitionRevision != nil {
        queryParams.Add("secretsSourceDefinitionRevision", strconv.Itoa(*secretsSourceDefinitionRevision))
    }
    body, marshalErr := json.Marshal(*definition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("dbeaf647-6167-421a-bda9-c9327b25e2e6")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the contents of a file in the given source code repository.
// project (required): Project ID or project name
// providerName (required): The name of the source provider.
// serviceEndpointId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
// repository (optional): If specified, the vendor-specific identifier or the name of the repository to get branches. Can only be omitted for providers that do not support multiple repositories.
// commitOrBranch (optional): The identifier of the commit or branch from which a file's contents are retrieved.
// path (optional): The path to the file to retrieve, relative to the root of the repository.
func (client Client) GetFileContents(project *string, providerName *string, serviceEndpointId *uuid.UUID, repository *string, commitOrBranch *string, path *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if providerName == nil || *providerName == "" {
        return nil, errors.New("providerName is a required parameter")
    }
    routeValues["providerName"] = *providerName

    queryParams := url.Values{}
    if serviceEndpointId != nil {
        queryParams.Add("serviceEndpointId", (*serviceEndpointId).String())
    }
    if repository != nil {
        queryParams.Add("repository", *repository)
    }
    if commitOrBranch != nil {
        queryParams.Add("commitOrBranch", *commitOrBranch)
    }
    if path != nil {
        queryParams.Add("path", *path)
    }
    locationId, _ := uuid.Parse("29d12225-b1d9-425f-b668-6c594a981313")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Creates a new folder.
// folder (required): The folder.
// project (required): Project ID or project name
// path (required): The full path of the folder.
func (client Client) CreateFolder(folder *Folder, project *string, path *string) (*Folder, error) {
    if folder == nil {
        return nil, errors.New("folder is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if path == nil {
        return nil, errors.New("path is a required parameter")
    }
    queryParams.Add("path", *path)
    body, marshalErr := json.Marshal(*folder)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Folder
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes a definition folder. Definitions and their corresponding builds will also be deleted.
// project (required): Project ID or project name
// path (required): The full path to the folder.
func (client Client) DeleteFolder(project *string, path *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if path == nil {
        return errors.New("path is a required parameter")
    }
    queryParams.Add("path", *path)
    locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Gets a list of build definition folders.
// project (required): Project ID or project name
// path (optional): The path to start with.
// queryOrder (optional): The order in which folders should be returned.
func (client Client) GetFolders(project *string, path *string, queryOrder *string) (*[]Folder, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if path != nil && *path != "" {
        routeValues["path"] = *path
    }

    queryParams := url.Values{}
    if queryOrder != nil {
        queryParams.Add("queryOrder", *queryOrder)
    }
    locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Folder
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates an existing folder at given  existing path
// folder (required): The new version of the folder.
// project (required): Project ID or project name
// path (required): The full path to the folder.
func (client Client) UpdateFolder(folder *Folder, project *string, path *string) (*Folder, error) {
    if folder == nil {
        return nil, errors.New("folder is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if path == nil {
        return nil, errors.New("path is a required parameter")
    }
    queryParams.Add("path", *path)
    body, marshalErr := json.Marshal(*folder)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a906531b-d2da-4f55-bda7-f3e676cc50d9")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Folder
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the latest build for a definition, optionally scoped to a specific branch.
// project (required): Project ID or project name
// definition (required): definition name with optional leading folder path, or the definition id
// branchName (optional): optional parameter that indicates the specific branch to use
func (client Client) GetLatestBuild(project *string, definition *string, branchName *string) (*Build, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definition == nil || *definition == "" {
        return nil, errors.New("definition is a required parameter")
    }
    routeValues["definition"] = *definition

    queryParams := url.Values{}
    if branchName != nil {
        queryParams.Add("branchName", *branchName)
    }
    locationId, _ := uuid.Parse("54481611-01f4-47f3-998f-160da0f0c229")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Build
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets an individual log file for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// logId (required): The ID of the log file.
// startLine (optional): The start line.
// endLine (optional): The end line.
func (client Client) GetBuildLog(project *string, buildId *int, logId *int, startLine *uint64, endLine *uint64) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if logId == nil {
        return nil, errors.New("logId is a required parameter")
    }
    routeValues["logId"] = strconv.Itoa(*logId)

    queryParams := url.Values{}
    if startLine != nil {
        queryParams.Add("startLine", strconv.FormatUint(*startLine, 10))
    }
    if endLine != nil {
        queryParams.Add("endLine", strconv.FormatUint(*endLine, 10))
    }
    locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets an individual log file for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// logId (required): The ID of the log file.
// startLine (optional): The start line.
// endLine (optional): The end line.
func (client Client) GetBuildLogLines(project *string, buildId *int, logId *int, startLine *uint64, endLine *uint64) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if logId == nil {
        return nil, errors.New("logId is a required parameter")
    }
    routeValues["logId"] = strconv.Itoa(*logId)

    queryParams := url.Values{}
    if startLine != nil {
        queryParams.Add("startLine", strconv.FormatUint(*startLine, 10))
    }
    if endLine != nil {
        queryParams.Add("endLine", strconv.FormatUint(*endLine, 10))
    }
    locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the logs for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) GetBuildLogs(project *string, buildId *int) (*[]BuildLog, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the logs for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) GetBuildLogsZip(project *string, buildId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets an individual log file for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// logId (required): The ID of the log file.
// startLine (optional): The start line.
// endLine (optional): The end line.
func (client Client) GetBuildLogZip(project *string, buildId *int, logId *int, startLine *uint64, endLine *uint64) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if logId == nil {
        return nil, errors.New("logId is a required parameter")
    }
    routeValues["logId"] = strconv.Itoa(*logId)

    queryParams := url.Values{}
    if startLine != nil {
        queryParams.Add("startLine", strconv.FormatUint(*startLine, 10))
    }
    if endLine != nil {
        queryParams.Add("endLine", strconv.FormatUint(*endLine, 10))
    }
    locationId, _ := uuid.Parse("35a80daf-7f30-45fc-86e8-6b813d9c90df")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets build metrics for a project.
// project (required): Project ID or project name
// metricAggregationType (optional): The aggregation type to use (hourly, daily).
// minMetricsTime (optional): The date from which to calculate metrics.
func (client Client) GetProjectMetrics(project *string, metricAggregationType *string, minMetricsTime *time.Time) (*[]BuildMetric, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if metricAggregationType != nil && *metricAggregationType != "" {
        routeValues["metricAggregationType"] = *metricAggregationType
    }

    queryParams := url.Values{}
    if minMetricsTime != nil {
        queryParams.Add("minMetricsTime", (*minMetricsTime).String())
    }
    locationId, _ := uuid.Parse("7433fae7-a6bc-41dc-a6e2-eef9005ce41a")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildMetric
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets build metrics for a definition.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
// minMetricsTime (optional): The date from which to calculate metrics.
func (client Client) GetDefinitionMetrics(project *string, definitionId *int, minMetricsTime *time.Time) (*[]BuildMetric, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if minMetricsTime != nil {
        queryParams.Add("minMetricsTime", (*minMetricsTime).String())
    }
    locationId, _ := uuid.Parse("d973b939-0ce0-4fec-91d8-da3940fa1827")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildMetric
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets all build definition options supported by the system.
// project (optional): Project ID or project name
func (client Client) GetBuildOptionDefinitions(project *string) (*[]BuildOptionDefinition, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    locationId, _ := uuid.Parse("591cb5a4-2d46-4f3a-a697-5cd42b6bd332")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildOptionDefinition
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the contents of a directory in the given source code repository.
// project (required): Project ID or project name
// providerName (required): The name of the source provider.
// serviceEndpointId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
// repository (optional): If specified, the vendor-specific identifier or the name of the repository to get branches. Can only be omitted for providers that do not support multiple repositories.
// commitOrBranch (optional): The identifier of the commit or branch from which a file's contents are retrieved.
// path (optional): The path contents to list, relative to the root of the repository.
func (client Client) GetPathContents(project *string, providerName *string, serviceEndpointId *uuid.UUID, repository *string, commitOrBranch *string, path *string) (*[]SourceRepositoryItem, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if providerName == nil || *providerName == "" {
        return nil, errors.New("providerName is a required parameter")
    }
    routeValues["providerName"] = *providerName

    queryParams := url.Values{}
    if serviceEndpointId != nil {
        queryParams.Add("serviceEndpointId", (*serviceEndpointId).String())
    }
    if repository != nil {
        queryParams.Add("repository", *repository)
    }
    if commitOrBranch != nil {
        queryParams.Add("commitOrBranch", *commitOrBranch)
    }
    if path != nil {
        queryParams.Add("path", *path)
    }
    locationId, _ := uuid.Parse("7944d6fb-df01-4709-920a-7a189aa34037")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SourceRepositoryItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets properties for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// filter (optional): A comma-delimited list of properties. If specified, filters to these specific properties.
func (client Client) GetBuildProperties(project *string, buildId *int, filter *[]string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if filter != nil {
        listAsString := strings.Join((*filter)[:], ",")
        queryParams.Add("filter", listAsString)
    }
    locationId, _ := uuid.Parse("0a6312e9-0627-49b7-8083-7d74a64849c9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates properties for a build.
// document (required): A json-patch document describing the properties to update.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) UpdateBuildProperties(document *[]JsonPatchOperation, project *string, buildId *int) (*interface{}, error) {
    if document == nil {
        return nil, errors.New("document is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    body, marshalErr := json.Marshal(*document)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0a6312e9-0627-49b7-8083-7d74a64849c9")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets properties for a definition.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
// filter (optional): A comma-delimited list of properties. If specified, filters to these specific properties.
func (client Client) GetDefinitionProperties(project *string, definitionId *int, filter *[]string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if filter != nil {
        listAsString := strings.Join((*filter)[:], ",")
        queryParams.Add("filter", listAsString)
    }
    locationId, _ := uuid.Parse("d9826ad7-2a68-46a9-a6e9-677698777895")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates properties for a definition.
// document (required): A json-patch document describing the properties to update.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
func (client Client) UpdateDefinitionProperties(document *[]JsonPatchOperation, project *string, definitionId *int) (*interface{}, error) {
    if document == nil {
        return nil, errors.New("document is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    body, marshalErr := json.Marshal(*document)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d9826ad7-2a68-46a9-a6e9-677698777895")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a pull request object from source provider.
// project (required): Project ID or project name
// providerName (required): The name of the source provider.
// pullRequestId (required): Vendor-specific id of the pull request.
// repositoryId (optional): Vendor-specific identifier or the name of the repository that contains the pull request.
// serviceEndpointId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
func (client Client) GetPullRequest(project *string, providerName *string, pullRequestId *string, repositoryId *string, serviceEndpointId *uuid.UUID) (*PullRequest, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if providerName == nil || *providerName == "" {
        return nil, errors.New("providerName is a required parameter")
    }
    routeValues["providerName"] = *providerName
    if pullRequestId == nil || *pullRequestId == "" {
        return nil, errors.New("pullRequestId is a required parameter")
    }
    routeValues["pullRequestId"] = *pullRequestId

    queryParams := url.Values{}
    if repositoryId != nil {
        queryParams.Add("repositoryId", *repositoryId)
    }
    if serviceEndpointId != nil {
        queryParams.Add("serviceEndpointId", (*serviceEndpointId).String())
    }
    locationId, _ := uuid.Parse("d8763ec7-9ff0-4fb4-b2b2-9d757906ff14")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PullRequest
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a build report.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// type_ (optional)
func (client Client) GetBuildReport(project *string, buildId *int, type_ *string) (*BuildReportMetadata, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    locationId, _ := uuid.Parse("45bcaa88-67e1-4042-a035-56d3b4a7d44c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildReportMetadata
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a build report.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// type_ (optional)
func (client Client) GetBuildReportHtmlContent(project *string, buildId *int, type_ *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    locationId, _ := uuid.Parse("45bcaa88-67e1-4042-a035-56d3b4a7d44c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "text/html", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets a list of source code repositories.
// project (required): Project ID or project name
// providerName (required): The name of the source provider.
// serviceEndpointId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
// repository (optional): If specified, the vendor-specific identifier or the name of a single repository to get.
// resultSet (optional): 'top' for the repositories most relevant for the endpoint. If not set, all repositories are returned. Ignored if 'repository' is set.
// pageResults (optional): If set to true, this will limit the set of results and will return a continuation token to continue the query.
// continuationToken (optional): When paging results, this is a continuation token, returned by a previous call to this method, that can be used to return the next set of repositories.
func (client Client) ListRepositories(project *string, providerName *string, serviceEndpointId *uuid.UUID, repository *string, resultSet *string, pageResults *bool, continuationToken *string) (*SourceRepositories, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if providerName == nil || *providerName == "" {
        return nil, errors.New("providerName is a required parameter")
    }
    routeValues["providerName"] = *providerName

    queryParams := url.Values{}
    if serviceEndpointId != nil {
        queryParams.Add("serviceEndpointId", (*serviceEndpointId).String())
    }
    if repository != nil {
        queryParams.Add("repository", *repository)
    }
    if resultSet != nil {
        queryParams.Add("resultSet", *resultSet)
    }
    if pageResults != nil {
        queryParams.Add("pageResults", strconv.FormatBool(*pageResults))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("d44d1680-f978-4834-9b93-8c6e132329c9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SourceRepositories
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// resources (required)
// project (required): Project ID or project name
// definitionId (required)
func (client Client) AuthorizeDefinitionResources(resources *[]DefinitionResourceReference, project *string, definitionId *int) (*[]DefinitionResourceReference, error) {
    if resources == nil {
        return nil, errors.New("resources is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    body, marshalErr := json.Marshal(*resources)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ea623316-1967-45eb-89ab-e9e6110cf2d6")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DefinitionResourceReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// definitionId (required)
func (client Client) GetDefinitionResources(project *string, definitionId *int) (*[]DefinitionResourceReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    locationId, _ := uuid.Parse("ea623316-1967-45eb-89ab-e9e6110cf2d6")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []DefinitionResourceReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets information about build resources in the system.
func (client Client) GetResourceUsage() (*BuildResourceUsage, error) {
    locationId, _ := uuid.Parse("3813d06c-9e36-4ea1-aac3-61a485d60e3d")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildResourceUsage
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets all revisions of a definition.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
func (client Client) GetDefinitionRevisions(project *string, definitionId *int) (*[]BuildDefinitionRevision, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    locationId, _ := uuid.Parse("7c116775-52e5-453e-8c5d-914d9762d8c4")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildDefinitionRevision
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the build settings.
// project (optional): Project ID or project name
func (client Client) GetBuildSettings(project *string) (*BuildSettings, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    locationId, _ := uuid.Parse("aa8c1c9c-ef8b-474a-b8c4-785c7b191d0d")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Updates the build settings.
// settings (required): The new settings.
// project (optional): Project ID or project name
func (client Client) UpdateBuildSettings(settings *BuildSettings, project *string) (*BuildSettings, error) {
    if settings == nil {
        return nil, errors.New("settings is a required parameter")
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*settings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("aa8c1c9c-ef8b-474a-b8c4-785c7b191d0d")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of source providers and their capabilities.
// project (required): Project ID or project name
func (client Client) ListSourceProviders(project *string) (*[]SourceProviderAttributes, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("3ce81729-954f-423d-a581-9fea01d25186")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SourceProviderAttributes
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] <p>Gets the build status for a definition, optionally scoped to a specific branch, stage, job, and configuration.</p> <p>If there are more than one, then it is required to pass in a stageName value when specifying a jobName, and the same rule then applies for both if passing a configuration parameter.</p>
// project (required): Project ID or project name
// definition (required): Either the definition name with optional leading folder path, or the definition id.
// branchName (optional): Only consider the most recent build for this branch.
// stageName (optional): Use this stage within the pipeline to render the status.
// jobName (optional): Use this job within a stage of the pipeline to render the status.
// configuration (optional): Use this job configuration to render the status
// label (optional): Replaces the default text on the left side of the badge.
func (client Client) GetStatusBadge(project *string, definition *string, branchName *string, stageName *string, jobName *string, configuration *string, label *string) (*string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definition == nil || *definition == "" {
        return nil, errors.New("definition is a required parameter")
    }
    routeValues["definition"] = *definition

    queryParams := url.Values{}
    if branchName != nil {
        queryParams.Add("branchName", *branchName)
    }
    if stageName != nil {
        queryParams.Add("stageName", *stageName)
    }
    if jobName != nil {
        queryParams.Add("jobName", *jobName)
    }
    if configuration != nil {
        queryParams.Add("configuration", *configuration)
    }
    if label != nil {
        queryParams.Add("label", *label)
    }
    locationId, _ := uuid.Parse("07acfdce-4757-4439-b422-ddd13a2fcc10")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Adds a tag to a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// tag (required): The tag to add.
func (client Client) AddBuildTag(project *string, buildId *int, tag *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if tag == nil || *tag == "" {
        return nil, errors.New("tag is a required parameter")
    }
    routeValues["tag"] = *tag

    locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Adds tags to a build.
// tags (required): The tags to add.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) AddBuildTags(tags *[]string, project *string, buildId *int) (*[]string, error) {
    if tags == nil {
        return nil, errors.New("tags is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    body, marshalErr := json.Marshal(*tags)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Removes a tag from a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// tag (required): The tag to remove.
func (client Client) DeleteBuildTag(project *string, buildId *int, tag *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if tag == nil || *tag == "" {
        return nil, errors.New("tag is a required parameter")
    }
    routeValues["tag"] = *tag

    locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the tags for a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
func (client Client) GetBuildTags(project *string, buildId *int) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    locationId, _ := uuid.Parse("6e6114b2-8161-44c8-8f6c-c5505782427f")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets a list of all build and definition tags in the project.
// project (required): Project ID or project name
func (client Client) GetTags(project *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("d84ac5c6-edc7-43d5-adc9-1b34be5dea09")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds a tag to a definition
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
// tag (required): The tag to add.
func (client Client) AddDefinitionTag(project *string, definitionId *int, tag *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)
    if tag == nil || *tag == "" {
        return nil, errors.New("tag is a required parameter")
    }
    routeValues["tag"] = *tag

    locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Adds multiple tags to a definition.
// tags (required): The tags to add.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
func (client Client) AddDefinitionTags(tags *[]string, project *string, definitionId *int) (*[]string, error) {
    if tags == nil {
        return nil, errors.New("tags is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    body, marshalErr := json.Marshal(*tags)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes a tag from a definition.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
// tag (required): The tag to remove.
func (client Client) DeleteDefinitionTag(project *string, definitionId *int, tag *string) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)
    if tag == nil || *tag == "" {
        return nil, errors.New("tag is a required parameter")
    }
    routeValues["tag"] = *tag

    locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets the tags for a definition.
// project (required): Project ID or project name
// definitionId (required): The ID of the definition.
// revision (optional): The definition revision number. If not specified, uses the latest revision of the definition.
func (client Client) GetDefinitionTags(project *string, definitionId *int, revision *int) (*[]string, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, errors.New("definitionId is a required parameter")
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if revision != nil {
        queryParams.Add("revision", strconv.Itoa(*revision))
    }
    locationId, _ := uuid.Parse("cb894432-134a-4d31-a839-83beceaace4b")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Deletes a build definition template.
// project (required): Project ID or project name
// templateId (required): The ID of the template.
func (client Client) DeleteTemplate(project *string, templateId *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if templateId == nil || *templateId == "" {
        return errors.New("templateId is a required parameter")
    }
    routeValues["templateId"] = *templateId

    locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Gets a specific build definition template.
// project (required): Project ID or project name
// templateId (required): The ID of the requested template.
func (client Client) GetTemplate(project *string, templateId *string) (*BuildDefinitionTemplate, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if templateId == nil || *templateId == "" {
        return nil, errors.New("templateId is a required parameter")
    }
    routeValues["templateId"] = *templateId

    locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildDefinitionTemplate
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets all definition templates.
// project (required): Project ID or project name
func (client Client) GetTemplates(project *string) (*[]BuildDefinitionTemplate, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildDefinitionTemplate
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Updates an existing build definition template.
// template (required): The new version of the template.
// project (required): Project ID or project name
// templateId (required): The ID of the template.
func (client Client) SaveTemplate(template *BuildDefinitionTemplate, project *string, templateId *string) (*BuildDefinitionTemplate, error) {
    if template == nil {
        return nil, errors.New("template is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if templateId == nil || *templateId == "" {
        return nil, errors.New("templateId is a required parameter")
    }
    routeValues["templateId"] = *templateId

    body, marshalErr := json.Marshal(*template)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e884571e-7f92-4d6a-9274-3f5649900835")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue BuildDefinitionTemplate
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Gets details for a build
// project (required): Project ID or project name
// buildId (required)
// timelineId (optional)
// changeId (optional)
// planId (optional)
func (client Client) GetBuildTimeline(project *string, buildId *int, timelineId *uuid.UUID, changeId *int, planId *uuid.UUID) (*Timeline, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)
    if timelineId != nil {
        routeValues["timelineId"] = (*timelineId).String()
    }

    queryParams := url.Values{}
    if changeId != nil {
        queryParams.Add("changeId", strconv.Itoa(*changeId))
    }
    if planId != nil {
        queryParams.Add("planId", (*planId).String())
    }
    locationId, _ := uuid.Parse("8baac422-4c6e-4de5-8532-db96d92acffa")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Timeline
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Recreates the webhooks for the specified triggers in the given source code repository.
// triggerTypes (required): The types of triggers to restore webhooks for.
// project (required): Project ID or project name
// providerName (required): The name of the source provider.
// serviceEndpointId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
// repository (optional): If specified, the vendor-specific identifier or the name of the repository to get webhooks. Can only be omitted for providers that do not support multiple repositories.
func (client Client) RestoreWebhooks(triggerTypes *[]DefinitionTriggerType, project *string, providerName *string, serviceEndpointId *uuid.UUID, repository *string) error {
    if triggerTypes == nil {
        return errors.New("triggerTypes is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if providerName == nil || *providerName == "" {
        return errors.New("providerName is a required parameter")
    }
    routeValues["providerName"] = *providerName

    queryParams := url.Values{}
    if serviceEndpointId != nil {
        queryParams.Add("serviceEndpointId", (*serviceEndpointId).String())
    }
    if repository != nil {
        queryParams.Add("repository", *repository)
    }
    body, marshalErr := json.Marshal(*triggerTypes)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("793bceb8-9736-4030-bd2f-fb3ce6d6b478")
    _, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Gets a list of webhooks installed in the given source code repository.
// project (required): Project ID or project name
// providerName (required): The name of the source provider.
// serviceEndpointId (optional): If specified, the ID of the service endpoint to query. Can only be omitted for providers that do not use service endpoints, e.g. TFVC or TFGit.
// repository (optional): If specified, the vendor-specific identifier or the name of the repository to get webhooks. Can only be omitted for providers that do not support multiple repositories.
func (client Client) ListWebhooks(project *string, providerName *string, serviceEndpointId *uuid.UUID, repository *string) (*[]RepositoryWebhook, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if providerName == nil || *providerName == "" {
        return nil, errors.New("providerName is a required parameter")
    }
    routeValues["providerName"] = *providerName

    queryParams := url.Values{}
    if serviceEndpointId != nil {
        queryParams.Add("serviceEndpointId", (*serviceEndpointId).String())
    }
    if repository != nil {
        queryParams.Add("repository", *repository)
    }
    locationId, _ := uuid.Parse("8f20ff82-9498-4812-9f6e-9c01bdc50e99")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []RepositoryWebhook
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the work items associated with a build.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// top (optional): The maximum number of work items to return.
func (client Client) GetBuildWorkItemsRefs(project *string, buildId *int, top *int) (*[]ResourceRef, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("5a21f5d2-5642-47e4-a0bd-1356e6731bee")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ResourceRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Gets the work items associated with a build, filtered to specific commits.
// commitIds (required): A comma-delimited list of commit IDs.
// project (required): Project ID or project name
// buildId (required): The ID of the build.
// top (optional): The maximum number of work items to return, or the number of commits to consider if no commit IDs are specified.
func (client Client) GetBuildWorkItemsRefsFromCommits(commitIds *[]string, project *string, buildId *int, top *int) (*[]ResourceRef, error) {
    if commitIds == nil {
        return nil, errors.New("commitIds is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    routeValues["buildId"] = strconv.Itoa(*buildId)

    queryParams := url.Values{}
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    body, marshalErr := json.Marshal(*commitIds)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5a21f5d2-5642-47e4-a0bd-1356e6731bee")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ResourceRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets all the work items between two builds.
// project (required): Project ID or project name
// fromBuildId (required): The ID of the first build.
// toBuildId (required): The ID of the last build.
// top (optional): The maximum number of work items to return.
func (client Client) GetWorkItemsBetweenBuilds(project *string, fromBuildId *int, toBuildId *int, top *int) (*[]ResourceRef, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if fromBuildId == nil {
        return nil, errors.New("fromBuildId is a required parameter")
    }
    queryParams.Add("fromBuildId", strconv.Itoa(*fromBuildId))
    if toBuildId == nil {
        return nil, errors.New("toBuildId is a required parameter")
    }
    queryParams.Add("toBuildId", strconv.Itoa(*toBuildId))
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("52ba8915-5518-42e3-a4bb-b0182d159e2d")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ResourceRef
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

