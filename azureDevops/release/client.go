// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package release

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

var ResourceAreaId, _ = uuid.Parse("efc2f575-36ef-48e9-b672-0c6fb4a48ac5")

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

// Get a list of approvals
// ctx
// project (required): Project ID or project name
// assignedToFilter (optional): Approvals assigned to this user.
// statusFilter (optional): Approvals with this status. Default is 'pending'.
// releaseIdsFilter (optional): Approvals for release id(s) mentioned in the filter. Multiple releases can be mentioned by separating them with ',' e.g. releaseIdsFilter=1,2,3,4.
// typeFilter (optional): Approval with this type.
// top (optional): Number of approvals to get. Default is 50.
// continuationToken (optional): Gets the approvals after the continuation token provided.
// queryOrder (optional): Gets the results in the defined order of created approvals. Default is 'descending'.
// includeMyGroupApprovals (optional): 'true' to include my group approvals. Default is 'false'.
func (client Client) GetApprovals(ctx context.Context, project *string, assignedToFilter *string, statusFilter *ApprovalStatus, releaseIdsFilter *[]int, typeFilter *ApprovalType, top *int, continuationToken *int, queryOrder *ReleaseQueryOrder, includeMyGroupApprovals *bool) (*[]ReleaseApproval, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if assignedToFilter != nil {
        queryParams.Add("assignedToFilter", *assignedToFilter)
    }
    if statusFilter != nil {
        queryParams.Add("statusFilter", string(*statusFilter))
    }
    if releaseIdsFilter != nil {
        var stringList []string
        for _, item := range *releaseIdsFilter {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if typeFilter != nil {
        queryParams.Add("typeFilter", string(*typeFilter))
    }
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", strconv.Itoa(*continuationToken))
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", string(*queryOrder))
    }
    if includeMyGroupApprovals != nil {
        queryParams.Add("includeMyGroupApprovals", strconv.FormatBool(*includeMyGroupApprovals))
    }
    locationId, _ := uuid.Parse("b47c6458-e73b-47cb-a770-4df1e8813a91")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReleaseApproval
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update status of an approval
// ctx
// approval (required): ReleaseApproval object having status, approver and comments.
// project (required): Project ID or project name
// approvalId (required): Id of the approval.
func (client Client) UpdateReleaseApproval(ctx context.Context, approval *ReleaseApproval, project *string, approvalId *int) (*ReleaseApproval, error) {
    if approval == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "approval"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if approvalId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "approvalId"} 
    }
    routeValues["approvalId"] = strconv.Itoa(*approvalId)

    body, marshalErr := json.Marshal(*approval)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("9328e074-59fb-465a-89d9-b09c82ee5109")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseApproval
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a release task attachment.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// environmentId (required): Id of the release environment.
// attemptId (required): Attempt number of deployment.
// planId (required): Plan Id of the deploy phase.
// timelineId (required): Timeline Id of the task.
// recordId (required): Record Id of attachment.
// type_ (required): Type of the attachment.
// name (required): Name of the attachment.
func (client Client) GetReleaseTaskAttachmentContent(ctx context.Context, project *string, releaseId *int, environmentId *int, attemptId *int, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)
    if environmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*environmentId)
    if attemptId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attemptId"} 
    }
    routeValues["attemptId"] = strconv.Itoa(*attemptId)
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "recordId"} 
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("60b86efb-7b8c-4853-8f9f-aa142b77b479")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Get the release task attachments.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// environmentId (required): Id of the release environment.
// attemptId (required): Attempt number of deployment.
// planId (required): Plan Id of the deploy phase.
// type_ (required): Type of the attachment.
func (client Client) GetReleaseTaskAttachments(ctx context.Context, project *string, releaseId *int, environmentId *int, attemptId *int, planId *uuid.UUID, type_ *string) (*[]ReleaseTaskAttachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)
    if environmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*environmentId)
    if attemptId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attemptId"} 
    }
    routeValues["attemptId"] = strconv.Itoa(*attemptId)
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("a4d06688-0dfa-4895-82a5-f43ec9452306")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReleaseTaskAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a release definition
// ctx
// releaseDefinition (required): release definition object to create.
// project (required): Project ID or project name
func (client Client) CreateReleaseDefinition(ctx context.Context, releaseDefinition *ReleaseDefinition, project *string) (*ReleaseDefinition, error) {
    if releaseDefinition == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseDefinition"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*releaseDefinition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d8f96f24-8ea7-4cb6-baab-2df8fc515665")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a release definition.
// ctx
// project (required): Project ID or project name
// definitionId (required): Id of the release definition.
// comment (optional): Comment for deleting a release definition.
// forceDelete (optional): 'true' to automatically cancel any in-progress release deployments and proceed with release definition deletion . Default is 'false'.
func (client Client) DeleteReleaseDefinition(ctx context.Context, project *string, definitionId *int, comment *string, forceDelete *bool) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if comment != nil {
        queryParams.Add("comment", *comment)
    }
    if forceDelete != nil {
        queryParams.Add("forceDelete", strconv.FormatBool(*forceDelete))
    }
    locationId, _ := uuid.Parse("d8f96f24-8ea7-4cb6-baab-2df8fc515665")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a release definition.
// ctx
// project (required): Project ID or project name
// definitionId (required): Id of the release definition.
// propertyFilters (optional): A comma-delimited list of extended properties to be retrieved. If set, the returned Release Definition will contain values for the specified property Ids (if they exist). If not set, properties will not be included.
func (client Client) GetReleaseDefinition(ctx context.Context, project *string, definitionId *int, propertyFilters *[]string) (*ReleaseDefinition, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    queryParams := url.Values{}
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    locationId, _ := uuid.Parse("d8f96f24-8ea7-4cb6-baab-2df8fc515665")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of release definitions.
// ctx
// project (required): Project ID or project name
// searchText (optional): Get release definitions with names containing searchText.
// expand (optional): The properties that should be expanded in the list of Release definitions.
// artifactType (optional): Release definitions with given artifactType will be returned. Values can be Build, Jenkins, GitHub, Nuget, Team Build (external), ExternalTFSBuild, Git, TFVC, ExternalTfsXamlBuild.
// artifactSourceId (optional): Release definitions with given artifactSourceId will be returned. e.g. For build it would be {projectGuid}:{BuildDefinitionId}, for Jenkins it would be {JenkinsConnectionId}:{JenkinsDefinitionId}, for TfsOnPrem it would be {TfsOnPremConnectionId}:{ProjectName}:{TfsOnPremDefinitionId}. For third-party artifacts e.g. TeamCity, BitBucket you may refer 'uniqueSourceIdentifier' inside vss-extension.json at https://github.com/Microsoft/vsts-rm-extensions/blob/master/Extensions.
// top (optional): Number of release definitions to get.
// continuationToken (optional): Gets the release definitions after the continuation token provided.
// queryOrder (optional): Gets the results in the defined order. Default is 'IdAscending'.
// path (optional): Gets the release definitions under the specified path.
// isExactNameMatch (optional): 'true'to gets the release definitions with exact match as specified in searchText. Default is 'false'.
// tagFilter (optional): A comma-delimited list of tags. Only release definitions with these tags will be returned.
// propertyFilters (optional): A comma-delimited list of extended properties to be retrieved. If set, the returned Release Definitions will contain values for the specified property Ids (if they exist). If not set, properties will not be included. Note that this will not filter out any Release Definition from results irrespective of whether it has property set or not.
// definitionIdFilter (optional): A comma-delimited list of release definitions to retrieve.
// isDeleted (optional): 'true' to get release definitions that has been deleted. Default is 'false'
// searchTextContainsFolderName (optional): 'true' to get the release definitions under the folder with name as specified in searchText. Default is 'false'.
func (client Client) GetReleaseDefinitions(ctx context.Context, project *string, searchText *string, expand *ReleaseDefinitionExpands, artifactType *string, artifactSourceId *string, top *int, continuationToken *string, queryOrder *ReleaseDefinitionQueryOrder, path *string, isExactNameMatch *bool, tagFilter *[]string, propertyFilters *[]string, definitionIdFilter *[]string, isDeleted *bool, searchTextContainsFolderName *bool) (*[]ReleaseDefinition, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if searchText != nil {
        queryParams.Add("searchText", *searchText)
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if artifactType != nil {
        queryParams.Add("artifactType", *artifactType)
    }
    if artifactSourceId != nil {
        queryParams.Add("artifactSourceId", *artifactSourceId)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", string(*queryOrder))
    }
    if path != nil {
        queryParams.Add("path", *path)
    }
    if isExactNameMatch != nil {
        queryParams.Add("isExactNameMatch", strconv.FormatBool(*isExactNameMatch))
    }
    if tagFilter != nil {
        listAsString := strings.Join((*tagFilter)[:], ",")
        queryParams.Add("tagFilter", listAsString)
    }
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if definitionIdFilter != nil {
        listAsString := strings.Join((*definitionIdFilter)[:], ",")
        queryParams.Add("definitionIdFilter", listAsString)
    }
    if isDeleted != nil {
        queryParams.Add("isDeleted", strconv.FormatBool(*isDeleted))
    }
    if searchTextContainsFolderName != nil {
        queryParams.Add("searchTextContainsFolderName", strconv.FormatBool(*searchTextContainsFolderName))
    }
    locationId, _ := uuid.Parse("d8f96f24-8ea7-4cb6-baab-2df8fc515665")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReleaseDefinition
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a release definition.
// ctx
// releaseDefinition (required): Release definition object to update.
// project (required): Project ID or project name
func (client Client) UpdateReleaseDefinition(ctx context.Context, releaseDefinition *ReleaseDefinition, project *string) (*ReleaseDefinition, error) {
    if releaseDefinition == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseDefinition"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*releaseDefinition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d8f96f24-8ea7-4cb6-baab-2df8fc515665")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// project (required): Project ID or project name
// definitionId (optional)
// definitionEnvironmentId (optional)
// createdBy (optional)
// minModifiedTime (optional)
// maxModifiedTime (optional)
// deploymentStatus (optional)
// operationStatus (optional)
// latestAttemptsOnly (optional)
// queryOrder (optional)
// top (optional)
// continuationToken (optional)
// createdFor (optional)
// minStartedTime (optional)
// maxStartedTime (optional)
// sourceBranch (optional)
func (client Client) GetDeployments(ctx context.Context, project *string, definitionId *int, definitionEnvironmentId *int, createdBy *string, minModifiedTime *time.Time, maxModifiedTime *time.Time, deploymentStatus *DeploymentStatus, operationStatus *DeploymentOperationStatus, latestAttemptsOnly *bool, queryOrder *ReleaseQueryOrder, top *int, continuationToken *int, createdFor *string, minStartedTime *time.Time, maxStartedTime *time.Time, sourceBranch *string) (*[]Deployment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if definitionId != nil {
        queryParams.Add("definitionId", strconv.Itoa(*definitionId))
    }
    if definitionEnvironmentId != nil {
        queryParams.Add("definitionEnvironmentId", strconv.Itoa(*definitionEnvironmentId))
    }
    if createdBy != nil {
        queryParams.Add("createdBy", *createdBy)
    }
    if minModifiedTime != nil {
        queryParams.Add("minModifiedTime", (*minModifiedTime).String())
    }
    if maxModifiedTime != nil {
        queryParams.Add("maxModifiedTime", (*maxModifiedTime).String())
    }
    if deploymentStatus != nil {
        queryParams.Add("deploymentStatus", string(*deploymentStatus))
    }
    if operationStatus != nil {
        queryParams.Add("operationStatus", string(*operationStatus))
    }
    if latestAttemptsOnly != nil {
        queryParams.Add("latestAttemptsOnly", strconv.FormatBool(*latestAttemptsOnly))
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", string(*queryOrder))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", strconv.Itoa(*continuationToken))
    }
    if createdFor != nil {
        queryParams.Add("createdFor", *createdFor)
    }
    if minStartedTime != nil {
        queryParams.Add("minStartedTime", (*minStartedTime).String())
    }
    if maxStartedTime != nil {
        queryParams.Add("maxStartedTime", (*maxStartedTime).String())
    }
    if sourceBranch != nil {
        queryParams.Add("sourceBranch", *sourceBranch)
    }
    locationId, _ := uuid.Parse("b005ef73-cddc-448e-9ba2-5193bf36b19f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Deployment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a release environment.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// environmentId (required): Id of the release environment.
func (client Client) GetReleaseEnvironment(ctx context.Context, project *string, releaseId *int, environmentId *int) (*ReleaseEnvironment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)
    if environmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*environmentId)

    locationId, _ := uuid.Parse("a7e426b1-03dc-48af-9dfe-c98bac612dcb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.6", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseEnvironment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update the status of a release environment
// ctx
// environmentUpdateData (required): Environment update meta data.
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// environmentId (required): Id of release environment.
func (client Client) UpdateReleaseEnvironment(ctx context.Context, environmentUpdateData *ReleaseEnvironmentUpdateMetadata, project *string, releaseId *int, environmentId *int) (*ReleaseEnvironment, error) {
    if environmentUpdateData == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "environmentUpdateData"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)
    if environmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*environmentId)

    body, marshalErr := json.Marshal(*environmentUpdateData)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a7e426b1-03dc-48af-9dfe-c98bac612dcb")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.6", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseEnvironment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Deletes a definition folder for given folder name and path and all it's existing definitions.
// ctx
// project (required): Project ID or project name
// path (required): Path of the folder to delete.
func (client Client) DeleteFolder(ctx context.Context, project *string, path *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if path == nil || *path == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "path"} 
    }
    routeValues["path"] = *path

    locationId, _ := uuid.Parse("f7ddf76d-ce0c-4d68-94ff-becaec5d9dea")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Gets folders.
// ctx
// project (required): Project ID or project name
// path (optional): Path of the folder.
// queryOrder (optional): Gets the results in the defined order. Default is 'None'.
func (client Client) GetFolders(ctx context.Context, project *string, path *string, queryOrder *FolderPathQueryOrder) (*[]Folder, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if path != nil && *path != "" {
        routeValues["path"] = *path
    }

    queryParams := url.Values{}
    if queryOrder != nil {
        queryParams.Add("queryOrder", string(*queryOrder))
    }
    locationId, _ := uuid.Parse("f7ddf76d-ce0c-4d68-94ff-becaec5d9dea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Folder
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates an existing folder at given existing path.
// ctx
// folder (required): folder.
// project (required): Project ID or project name
// path (required): Path of the folder to update.
func (client Client) UpdateFolder(ctx context.Context, folder *Folder, project *string, path *string) (*Folder, error) {
    if folder == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "folder"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if path == nil || *path == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "path"} 
    }
    routeValues["path"] = *path

    body, marshalErr := json.Marshal(*folder)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f7ddf76d-ce0c-4d68-94ff-becaec5d9dea")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Folder
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Updates the gate for a deployment.
// ctx
// gateUpdateMetadata (required): Metadata to patch the Release Gates.
// project (required): Project ID or project name
// gateStepId (required): Gate step Id.
func (client Client) UpdateGates(ctx context.Context, gateUpdateMetadata *GateUpdateMetadata, project *string, gateStepId *int) (*ReleaseGates, error) {
    if gateUpdateMetadata == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "gateUpdateMetadata"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if gateStepId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "gateStepId"} 
    }
    routeValues["gateStepId"] = strconv.Itoa(*gateStepId)

    body, marshalErr := json.Marshal(*gateUpdateMetadata)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2666a539-2001-4f80-bcc7-0379956749d4")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseGates
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get logs for a release Id.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
func (client Client) GetLogs(ctx context.Context, project *string, releaseId *int) (interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)

    locationId, _ := uuid.Parse("c37fbab5-214b-48e4-a55b-cb6b4f6e4038")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Gets the task log of a release as a plain text file.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// environmentId (required): Id of release environment.
// releaseDeployPhaseId (required): Release deploy phase Id.
// taskId (required): ReleaseTask Id for the log.
// startLine (optional): Starting line number for logs
// endLine (optional): Ending line number for logs
func (client Client) GetTaskLog(ctx context.Context, project *string, releaseId *int, environmentId *int, releaseDeployPhaseId *int, taskId *int, startLine *uint64, endLine *uint64) (interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)
    if environmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*environmentId)
    if releaseDeployPhaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseDeployPhaseId"} 
    }
    routeValues["releaseDeployPhaseId"] = strconv.Itoa(*releaseDeployPhaseId)
    if taskId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "taskId"} 
    }
    routeValues["taskId"] = strconv.Itoa(*taskId)

    queryParams := url.Values{}
    if startLine != nil {
        queryParams.Add("startLine", strconv.FormatUint(*startLine, 10))
    }
    if endLine != nil {
        queryParams.Add("endLine", strconv.FormatUint(*endLine, 10))
    }
    locationId, _ := uuid.Parse("17c91af7-09fd-4256-bff1-c24ee4f73bc0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Get manual intervention for a given release and manual intervention id.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// manualInterventionId (required): Id of the manual intervention.
func (client Client) GetManualIntervention(ctx context.Context, project *string, releaseId *int, manualInterventionId *int) (*ManualIntervention, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)
    if manualInterventionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "manualInterventionId"} 
    }
    routeValues["manualInterventionId"] = strconv.Itoa(*manualInterventionId)

    locationId, _ := uuid.Parse("616c46e4-f370-4456-adaa-fbaf79c7b79e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ManualIntervention
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// List all manual interventions for a given release.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
func (client Client) GetManualInterventions(ctx context.Context, project *string, releaseId *int) (*[]ManualIntervention, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)

    locationId, _ := uuid.Parse("616c46e4-f370-4456-adaa-fbaf79c7b79e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ManualIntervention
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update manual intervention.
// ctx
// manualInterventionUpdateMetadata (required): Meta data to update manual intervention.
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// manualInterventionId (required): Id of the manual intervention.
func (client Client) UpdateManualIntervention(ctx context.Context, manualInterventionUpdateMetadata *ManualInterventionUpdateMetadata, project *string, releaseId *int, manualInterventionId *int) (*ManualIntervention, error) {
    if manualInterventionUpdateMetadata == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "manualInterventionUpdateMetadata"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)
    if manualInterventionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "manualInterventionId"} 
    }
    routeValues["manualInterventionId"] = strconv.Itoa(*manualInterventionId)

    body, marshalErr := json.Marshal(*manualInterventionUpdateMetadata)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("616c46e4-f370-4456-adaa-fbaf79c7b79e")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ManualIntervention
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of releases
// ctx
// project (optional): Project ID or project name
// definitionId (optional): Releases from this release definition Id.
// definitionEnvironmentId (optional)
// searchText (optional): Releases with names containing searchText.
// createdBy (optional): Releases created by this user.
// statusFilter (optional): Releases that have this status.
// environmentStatusFilter (optional)
// minCreatedTime (optional): Releases that were created after this time.
// maxCreatedTime (optional): Releases that were created before this time.
// queryOrder (optional): Gets the results in the defined order of created date for releases. Default is descending.
// top (optional): Number of releases to get. Default is 50.
// continuationToken (optional): Gets the releases after the continuation token provided.
// expand (optional): The property that should be expanded in the list of releases.
// artifactTypeId (optional): Releases with given artifactTypeId will be returned. Values can be Build, Jenkins, GitHub, Nuget, Team Build (external), ExternalTFSBuild, Git, TFVC, ExternalTfsXamlBuild.
// sourceId (optional): Unique identifier of the artifact used. e.g. For build it would be {projectGuid}:{BuildDefinitionId}, for Jenkins it would be {JenkinsConnectionId}:{JenkinsDefinitionId}, for TfsOnPrem it would be {TfsOnPremConnectionId}:{ProjectName}:{TfsOnPremDefinitionId}. For third-party artifacts e.g. TeamCity, BitBucket you may refer 'uniqueSourceIdentifier' inside vss-extension.json https://github.com/Microsoft/vsts-rm-extensions/blob/master/Extensions.
// artifactVersionId (optional): Releases with given artifactVersionId will be returned. E.g. in case of Build artifactType, it is buildId.
// sourceBranchFilter (optional): Releases with given sourceBranchFilter will be returned.
// isDeleted (optional): Gets the soft deleted releases, if true.
// tagFilter (optional): A comma-delimited list of tags. Only releases with these tags will be returned.
// propertyFilters (optional): A comma-delimited list of extended properties to be retrieved. If set, the returned Releases will contain values for the specified property Ids (if they exist). If not set, properties will not be included. Note that this will not filter out any Release from results irrespective of whether it has property set or not.
// releaseIdFilter (optional): A comma-delimited list of releases Ids. Only releases with these Ids will be returned.
// path (optional): Releases under this folder path will be returned
func (client Client) GetReleases(ctx context.Context, project *string, definitionId *int, definitionEnvironmentId *int, searchText *string, createdBy *string, statusFilter *ReleaseStatus, environmentStatusFilter *int, minCreatedTime *time.Time, maxCreatedTime *time.Time, queryOrder *ReleaseQueryOrder, top *int, continuationToken *int, expand *ReleaseExpands, artifactTypeId *string, sourceId *string, artifactVersionId *string, sourceBranchFilter *string, isDeleted *bool, tagFilter *[]string, propertyFilters *[]string, releaseIdFilter *[]int, path *string) (*[]Release, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if definitionId != nil {
        queryParams.Add("definitionId", strconv.Itoa(*definitionId))
    }
    if definitionEnvironmentId != nil {
        queryParams.Add("definitionEnvironmentId", strconv.Itoa(*definitionEnvironmentId))
    }
    if searchText != nil {
        queryParams.Add("searchText", *searchText)
    }
    if createdBy != nil {
        queryParams.Add("createdBy", *createdBy)
    }
    if statusFilter != nil {
        queryParams.Add("statusFilter", string(*statusFilter))
    }
    if environmentStatusFilter != nil {
        queryParams.Add("environmentStatusFilter", strconv.Itoa(*environmentStatusFilter))
    }
    if minCreatedTime != nil {
        queryParams.Add("minCreatedTime", (*minCreatedTime).String())
    }
    if maxCreatedTime != nil {
        queryParams.Add("maxCreatedTime", (*maxCreatedTime).String())
    }
    if queryOrder != nil {
        queryParams.Add("queryOrder", string(*queryOrder))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", strconv.Itoa(*continuationToken))
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if artifactTypeId != nil {
        queryParams.Add("artifactTypeId", *artifactTypeId)
    }
    if sourceId != nil {
        queryParams.Add("sourceId", *sourceId)
    }
    if artifactVersionId != nil {
        queryParams.Add("artifactVersionId", *artifactVersionId)
    }
    if sourceBranchFilter != nil {
        queryParams.Add("sourceBranchFilter", *sourceBranchFilter)
    }
    if isDeleted != nil {
        queryParams.Add("isDeleted", strconv.FormatBool(*isDeleted))
    }
    if tagFilter != nil {
        listAsString := strings.Join((*tagFilter)[:], ",")
        queryParams.Add("tagFilter", listAsString)
    }
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if releaseIdFilter != nil {
        var stringList []string
        for _, item := range *releaseIdFilter {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if path != nil {
        queryParams.Add("path", *path)
    }
    locationId, _ := uuid.Parse("a166fde7-27ad-408e-ba75-703c2cc9d500")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Release
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a release.
// ctx
// releaseStartMetadata (required): Metadata to create a release.
// project (required): Project ID or project name
func (client Client) CreateRelease(ctx context.Context, releaseStartMetadata *ReleaseStartMetadata, project *string) (*Release, error) {
    if releaseStartMetadata == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseStartMetadata"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*releaseStartMetadata)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a166fde7-27ad-408e-ba75-703c2cc9d500")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Release
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a Release
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// approvalFilters (optional): A filter which would allow fetching approval steps selectively based on whether it is automated, or manual. This would also decide whether we should fetch pre and post approval snapshots. Assumes All by default
// propertyFilters (optional): A comma-delimited list of extended properties to be retrieved. If set, the returned Release will contain values for the specified property Ids (if they exist). If not set, properties will not be included.
// expand (optional): A property that should be expanded in the release.
// topGateRecords (optional): Number of release gate records to get. Default is 5.
func (client Client) GetRelease(ctx context.Context, project *string, releaseId *int, approvalFilters *ApprovalFilters, propertyFilters *[]string, expand *SingleReleaseExpands, topGateRecords *int) (*Release, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)

    queryParams := url.Values{}
    if approvalFilters != nil {
        queryParams.Add("approvalFilters", string(*approvalFilters))
    }
    if propertyFilters != nil {
        listAsString := strings.Join((*propertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if expand != nil {
        queryParams.Add("$expand", string(*expand))
    }
    if topGateRecords != nil {
        queryParams.Add("$topGateRecords", strconv.Itoa(*topGateRecords))
    }
    locationId, _ := uuid.Parse("a166fde7-27ad-408e-ba75-703c2cc9d500")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Release
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get release for a given revision number.
// ctx
// project (required): Project ID or project name
// releaseId (required): Id of the release.
// definitionSnapshotRevision (required): Definition snapshot revision number.
func (client Client) GetReleaseRevision(ctx context.Context, project *string, releaseId *int, definitionSnapshotRevision *int) (interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)

    queryParams := url.Values{}
    if definitionSnapshotRevision == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "definitionSnapshotRevision"}
    }
    queryParams.Add("definitionSnapshotRevision", strconv.Itoa(*definitionSnapshotRevision))
    locationId, _ := uuid.Parse("a166fde7-27ad-408e-ba75-703c2cc9d500")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Update a complete release object.
// ctx
// release (required): Release object for update.
// project (required): Project ID or project name
// releaseId (required): Id of the release to update.
func (client Client) UpdateRelease(ctx context.Context, release *Release, project *string, releaseId *int) (*Release, error) {
    if release == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "release"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)

    body, marshalErr := json.Marshal(*release)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a166fde7-27ad-408e-ba75-703c2cc9d500")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Release
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Update few properties of a release.
// ctx
// releaseUpdateMetadata (required): Properties of release to update.
// project (required): Project ID or project name
// releaseId (required): Id of the release to update.
func (client Client) UpdateReleaseResource(ctx context.Context, releaseUpdateMetadata *ReleaseUpdateMetadata, project *string, releaseId *int) (*Release, error) {
    if releaseUpdateMetadata == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseUpdateMetadata"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*releaseId)

    body, marshalErr := json.Marshal(*releaseUpdateMetadata)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a166fde7-27ad-408e-ba75-703c2cc9d500")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Release
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get release definition for a given definitionId and revision
// ctx
// project (required): Project ID or project name
// definitionId (required): Id of the definition.
// revision (required): Id of the revision.
func (client Client) GetDefinitionRevision(ctx context.Context, project *string, definitionId *int, revision *int) (interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)
    if revision == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "revision"} 
    }
    routeValues["revision"] = strconv.Itoa(*revision)

    locationId, _ := uuid.Parse("258b82e0-9d41-43f3-86d6-fef14ddd44bc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Get revision history for a release definition
// ctx
// project (required): Project ID or project name
// definitionId (required): Id of the definition.
func (client Client) GetReleaseDefinitionHistory(ctx context.Context, project *string, definitionId *int) (*[]ReleaseDefinitionRevision, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if definitionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*definitionId)

    locationId, _ := uuid.Parse("258b82e0-9d41-43f3-86d6-fef14ddd44bc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReleaseDefinitionRevision
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

