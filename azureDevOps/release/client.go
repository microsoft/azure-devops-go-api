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
    "github.com/microsoft/azure-devops-go-api/azureDevOps"
    "io"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

var ResourceAreaId, _ = uuid.Parse("efc2f575-36ef-48e9-b672-0c6fb4a48ac5")

type Client struct {
    Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// Get a list of approvals
func (client *Client) GetApprovals(ctx context.Context, args GetApprovalsArgs) (*[]ReleaseApproval, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.AssignedToFilter != nil {
        queryParams.Add("assignedToFilter", *args.AssignedToFilter)
    }
    if args.StatusFilter != nil {
        queryParams.Add("statusFilter", string(*args.StatusFilter))
    }
    if args.ReleaseIdsFilter != nil {
        var stringList []string
        for _, item := range *args.ReleaseIdsFilter {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.TypeFilter != nil {
        queryParams.Add("typeFilter", string(*args.TypeFilter))
    }
    if args.Top != nil {
        queryParams.Add("top", strconv.Itoa(*args.Top))
    }
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", strconv.Itoa(*args.ContinuationToken))
    }
    if args.QueryOrder != nil {
        queryParams.Add("queryOrder", string(*args.QueryOrder))
    }
    if args.IncludeMyGroupApprovals != nil {
        queryParams.Add("includeMyGroupApprovals", strconv.FormatBool(*args.IncludeMyGroupApprovals))
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

// Arguments for the GetApprovals function
type GetApprovalsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional) Approvals assigned to this user.
    AssignedToFilter *string
    // (optional) Approvals with this status. Default is 'pending'.
    StatusFilter *ApprovalStatus
    // (optional) Approvals for release id(s) mentioned in the filter. Multiple releases can be mentioned by separating them with ',' e.g. releaseIdsFilter=1,2,3,4.
    ReleaseIdsFilter *[]int
    // (optional) Approval with this type.
    TypeFilter *ApprovalType
    // (optional) Number of approvals to get. Default is 50.
    Top *int
    // (optional) Gets the approvals after the continuation token provided.
    ContinuationToken *int
    // (optional) Gets the results in the defined order of created approvals. Default is 'descending'.
    QueryOrder *ReleaseQueryOrder
    // (optional) 'true' to include my group approvals. Default is 'false'.
    IncludeMyGroupApprovals *bool
}

// Update status of an approval
func (client *Client) UpdateReleaseApproval(ctx context.Context, args UpdateReleaseApprovalArgs) (*ReleaseApproval, error) {
    if args.Approval == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "approval"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ApprovalId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "approvalId"} 
    }
    routeValues["approvalId"] = strconv.Itoa(*args.ApprovalId)

    body, marshalErr := json.Marshal(*args.Approval)
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

// Arguments for the UpdateReleaseApproval function
type UpdateReleaseApprovalArgs struct {
    // (required) ReleaseApproval object having status, approver and comments.
    Approval *ReleaseApproval
    // (required) Project ID or project name
    Project *string
    // (required) Id of the approval.
    ApprovalId *int
}

// [Preview API] Get a release task attachment.
func (client *Client) GetReleaseTaskAttachmentContent(ctx context.Context, args GetReleaseTaskAttachmentContentArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)
    if args.EnvironmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*args.EnvironmentId)
    if args.AttemptId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attemptId"} 
    }
    routeValues["attemptId"] = strconv.Itoa(*args.AttemptId)
    if args.PlanId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*args.PlanId).String()
    if args.TimelineId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*args.TimelineId).String()
    if args.RecordId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "recordId"} 
    }
    routeValues["recordId"] = (*args.RecordId).String()
    if args.Type_ == nil || *args.Type_ == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *args.Type_
    if args.Name == nil || *args.Name == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *args.Name

    locationId, _ := uuid.Parse("60b86efb-7b8c-4853-8f9f-aa142b77b479")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetReleaseTaskAttachmentContent function
type GetReleaseTaskAttachmentContentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Id of the release environment.
    EnvironmentId *int
    // (required) Attempt number of deployment.
    AttemptId *int
    // (required) Plan Id of the deploy phase.
    PlanId *uuid.UUID
    // (required) Timeline Id of the task.
    TimelineId *uuid.UUID
    // (required) Record Id of attachment.
    RecordId *uuid.UUID
    // (required) Type of the attachment.
    Type_ *string
    // (required) Name of the attachment.
    Name *string
}

// [Preview API] Get the release task attachments.
func (client *Client) GetReleaseTaskAttachments(ctx context.Context, args GetReleaseTaskAttachmentsArgs) (*[]ReleaseTaskAttachment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)
    if args.EnvironmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*args.EnvironmentId)
    if args.AttemptId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attemptId"} 
    }
    routeValues["attemptId"] = strconv.Itoa(*args.AttemptId)
    if args.PlanId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*args.PlanId).String()
    if args.Type_ == nil || *args.Type_ == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *args.Type_

    locationId, _ := uuid.Parse("a4d06688-0dfa-4895-82a5-f43ec9452306")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReleaseTaskAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetReleaseTaskAttachments function
type GetReleaseTaskAttachmentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Id of the release environment.
    EnvironmentId *int
    // (required) Attempt number of deployment.
    AttemptId *int
    // (required) Plan Id of the deploy phase.
    PlanId *uuid.UUID
    // (required) Type of the attachment.
    Type_ *string
}

// Create a release definition
func (client *Client) CreateReleaseDefinition(ctx context.Context, args CreateReleaseDefinitionArgs) (*ReleaseDefinition, error) {
    if args.ReleaseDefinition == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseDefinition"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.ReleaseDefinition)
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

// Arguments for the CreateReleaseDefinition function
type CreateReleaseDefinitionArgs struct {
    // (required) release definition object to create.
    ReleaseDefinition *ReleaseDefinition
    // (required) Project ID or project name
    Project *string
}

// Delete a release definition.
func (client *Client) DeleteReleaseDefinition(ctx context.Context, args DeleteReleaseDefinitionArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.DefinitionId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

    queryParams := url.Values{}
    if args.Comment != nil {
        queryParams.Add("comment", *args.Comment)
    }
    if args.ForceDelete != nil {
        queryParams.Add("forceDelete", strconv.FormatBool(*args.ForceDelete))
    }
    locationId, _ := uuid.Parse("d8f96f24-8ea7-4cb6-baab-2df8fc515665")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteReleaseDefinition function
type DeleteReleaseDefinitionArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release definition.
    DefinitionId *int
    // (optional) Comment for deleting a release definition.
    Comment *string
    // (optional) 'true' to automatically cancel any in-progress release deployments and proceed with release definition deletion . Default is 'false'.
    ForceDelete *bool
}

// Get a release definition.
func (client *Client) GetReleaseDefinition(ctx context.Context, args GetReleaseDefinitionArgs) (*ReleaseDefinition, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.DefinitionId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

    queryParams := url.Values{}
    if args.PropertyFilters != nil {
        listAsString := strings.Join((*args.PropertyFilters)[:], ",")
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

// Arguments for the GetReleaseDefinition function
type GetReleaseDefinitionArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release definition.
    DefinitionId *int
    // (optional) A comma-delimited list of extended properties to be retrieved. If set, the returned Release Definition will contain values for the specified property Ids (if they exist). If not set, properties will not be included.
    PropertyFilters *[]string
}

// Get a list of release definitions.
func (client *Client) GetReleaseDefinitions(ctx context.Context, args GetReleaseDefinitionsArgs) (*[]ReleaseDefinition, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.SearchText != nil {
        queryParams.Add("searchText", *args.SearchText)
    }
    if args.Expand != nil {
        queryParams.Add("$expand", string(*args.Expand))
    }
    if args.ArtifactType != nil {
        queryParams.Add("artifactType", *args.ArtifactType)
    }
    if args.ArtifactSourceId != nil {
        queryParams.Add("artifactSourceId", *args.ArtifactSourceId)
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", *args.ContinuationToken)
    }
    if args.QueryOrder != nil {
        queryParams.Add("queryOrder", string(*args.QueryOrder))
    }
    if args.Path != nil {
        queryParams.Add("path", *args.Path)
    }
    if args.IsExactNameMatch != nil {
        queryParams.Add("isExactNameMatch", strconv.FormatBool(*args.IsExactNameMatch))
    }
    if args.TagFilter != nil {
        listAsString := strings.Join((*args.TagFilter)[:], ",")
        queryParams.Add("tagFilter", listAsString)
    }
    if args.PropertyFilters != nil {
        listAsString := strings.Join((*args.PropertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if args.DefinitionIdFilter != nil {
        listAsString := strings.Join((*args.DefinitionIdFilter)[:], ",")
        queryParams.Add("definitionIdFilter", listAsString)
    }
    if args.IsDeleted != nil {
        queryParams.Add("isDeleted", strconv.FormatBool(*args.IsDeleted))
    }
    if args.SearchTextContainsFolderName != nil {
        queryParams.Add("searchTextContainsFolderName", strconv.FormatBool(*args.SearchTextContainsFolderName))
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

// Arguments for the GetReleaseDefinitions function
type GetReleaseDefinitionsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional) Get release definitions with names containing searchText.
    SearchText *string
    // (optional) The properties that should be expanded in the list of Release definitions.
    Expand *ReleaseDefinitionExpands
    // (optional) Release definitions with given artifactType will be returned. Values can be Build, Jenkins, GitHub, Nuget, Team Build (external), ExternalTFSBuild, Git, TFVC, ExternalTfsXamlBuild.
    ArtifactType *string
    // (optional) Release definitions with given artifactSourceId will be returned. e.g. For build it would be {projectGuid}:{BuildDefinitionId}, for Jenkins it would be {JenkinsConnectionId}:{JenkinsDefinitionId}, for TfsOnPrem it would be {TfsOnPremConnectionId}:{ProjectName}:{TfsOnPremDefinitionId}. For third-party artifacts e.g. TeamCity, BitBucket you may refer 'uniqueSourceIdentifier' inside vss-extension.json at https://github.com/Microsoft/vsts-rm-extensions/blob/master/Extensions.
    ArtifactSourceId *string
    // (optional) Number of release definitions to get.
    Top *int
    // (optional) Gets the release definitions after the continuation token provided.
    ContinuationToken *string
    // (optional) Gets the results in the defined order. Default is 'IdAscending'.
    QueryOrder *ReleaseDefinitionQueryOrder
    // (optional) Gets the release definitions under the specified path.
    Path *string
    // (optional) 'true'to gets the release definitions with exact match as specified in searchText. Default is 'false'.
    IsExactNameMatch *bool
    // (optional) A comma-delimited list of tags. Only release definitions with these tags will be returned.
    TagFilter *[]string
    // (optional) A comma-delimited list of extended properties to be retrieved. If set, the returned Release Definitions will contain values for the specified property Ids (if they exist). If not set, properties will not be included. Note that this will not filter out any Release Definition from results irrespective of whether it has property set or not.
    PropertyFilters *[]string
    // (optional) A comma-delimited list of release definitions to retrieve.
    DefinitionIdFilter *[]string
    // (optional) 'true' to get release definitions that has been deleted. Default is 'false'
    IsDeleted *bool
    // (optional) 'true' to get the release definitions under the folder with name as specified in searchText. Default is 'false'.
    SearchTextContainsFolderName *bool
}

// Update a release definition.
func (client *Client) UpdateReleaseDefinition(ctx context.Context, args UpdateReleaseDefinitionArgs) (*ReleaseDefinition, error) {
    if args.ReleaseDefinition == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseDefinition"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.ReleaseDefinition)
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

// Arguments for the UpdateReleaseDefinition function
type UpdateReleaseDefinitionArgs struct {
    // (required) Release definition object to update.
    ReleaseDefinition *ReleaseDefinition
    // (required) Project ID or project name
    Project *string
}

func (client *Client) GetDeployments(ctx context.Context, args GetDeploymentsArgs) (*[]Deployment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.DefinitionId != nil {
        queryParams.Add("definitionId", strconv.Itoa(*args.DefinitionId))
    }
    if args.DefinitionEnvironmentId != nil {
        queryParams.Add("definitionEnvironmentId", strconv.Itoa(*args.DefinitionEnvironmentId))
    }
    if args.CreatedBy != nil {
        queryParams.Add("createdBy", *args.CreatedBy)
    }
    if args.MinModifiedTime != nil {
        queryParams.Add("minModifiedTime", (*args.MinModifiedTime).String())
    }
    if args.MaxModifiedTime != nil {
        queryParams.Add("maxModifiedTime", (*args.MaxModifiedTime).String())
    }
    if args.DeploymentStatus != nil {
        queryParams.Add("deploymentStatus", string(*args.DeploymentStatus))
    }
    if args.OperationStatus != nil {
        queryParams.Add("operationStatus", string(*args.OperationStatus))
    }
    if args.LatestAttemptsOnly != nil {
        queryParams.Add("latestAttemptsOnly", strconv.FormatBool(*args.LatestAttemptsOnly))
    }
    if args.QueryOrder != nil {
        queryParams.Add("queryOrder", string(*args.QueryOrder))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", strconv.Itoa(*args.ContinuationToken))
    }
    if args.CreatedFor != nil {
        queryParams.Add("createdFor", *args.CreatedFor)
    }
    if args.MinStartedTime != nil {
        queryParams.Add("minStartedTime", (*args.MinStartedTime).String())
    }
    if args.MaxStartedTime != nil {
        queryParams.Add("maxStartedTime", (*args.MaxStartedTime).String())
    }
    if args.SourceBranch != nil {
        queryParams.Add("sourceBranch", *args.SourceBranch)
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

// Arguments for the GetDeployments function
type GetDeploymentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional)
    DefinitionId *int
    // (optional)
    DefinitionEnvironmentId *int
    // (optional)
    CreatedBy *string
    // (optional)
    MinModifiedTime *time.Time
    // (optional)
    MaxModifiedTime *time.Time
    // (optional)
    DeploymentStatus *DeploymentStatus
    // (optional)
    OperationStatus *DeploymentOperationStatus
    // (optional)
    LatestAttemptsOnly *bool
    // (optional)
    QueryOrder *ReleaseQueryOrder
    // (optional)
    Top *int
    // (optional)
    ContinuationToken *int
    // (optional)
    CreatedFor *string
    // (optional)
    MinStartedTime *time.Time
    // (optional)
    MaxStartedTime *time.Time
    // (optional)
    SourceBranch *string
}

// [Preview API] Get a release environment.
func (client *Client) GetReleaseEnvironment(ctx context.Context, args GetReleaseEnvironmentArgs) (*ReleaseEnvironment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)
    if args.EnvironmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*args.EnvironmentId)

    locationId, _ := uuid.Parse("a7e426b1-03dc-48af-9dfe-c98bac612dcb")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.6", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ReleaseEnvironment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetReleaseEnvironment function
type GetReleaseEnvironmentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Id of the release environment.
    EnvironmentId *int
}

// [Preview API] Update the status of a release environment
func (client *Client) UpdateReleaseEnvironment(ctx context.Context, args UpdateReleaseEnvironmentArgs) (*ReleaseEnvironment, error) {
    if args.EnvironmentUpdateData == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "environmentUpdateData"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)
    if args.EnvironmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*args.EnvironmentId)

    body, marshalErr := json.Marshal(*args.EnvironmentUpdateData)
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

// Arguments for the UpdateReleaseEnvironment function
type UpdateReleaseEnvironmentArgs struct {
    // (required) Environment update meta data.
    EnvironmentUpdateData *ReleaseEnvironmentUpdateMetadata
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Id of release environment.
    EnvironmentId *int
}

// [Preview API] Creates a new folder.
func (client *Client) CreateFolder(ctx context.Context, args CreateFolderArgs) (*Folder, error) {
    if args.Folder == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "folder"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Folder)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f7ddf76d-ce0c-4d68-94ff-becaec5d9dea")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Folder
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateFolder function
type CreateFolderArgs struct {
    // (required) Folder to create.
    Folder *Folder
    // (required) Project ID or project name
    Project *string
}

// [Preview API] Deletes a definition folder for given folder name and path and all it's existing definitions.
func (client *Client) DeleteFolder(ctx context.Context, args DeleteFolderArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.Path == nil || *args.Path == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "path"} 
    }
    routeValues["path"] = *args.Path

    locationId, _ := uuid.Parse("f7ddf76d-ce0c-4d68-94ff-becaec5d9dea")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteFolder function
type DeleteFolderArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Path of the folder to delete.
    Path *string
}

// [Preview API] Gets folders.
func (client *Client) GetFolders(ctx context.Context, args GetFoldersArgs) (*[]Folder, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.Path != nil && *args.Path != "" {
        routeValues["path"] = *args.Path
    }

    queryParams := url.Values{}
    if args.QueryOrder != nil {
        queryParams.Add("queryOrder", string(*args.QueryOrder))
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

// Arguments for the GetFolders function
type GetFoldersArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional) Path of the folder.
    Path *string
    // (optional) Gets the results in the defined order. Default is 'None'.
    QueryOrder *FolderPathQueryOrder
}

// [Preview API] Updates an existing folder at given existing path.
func (client *Client) UpdateFolder(ctx context.Context, args UpdateFolderArgs) (*Folder, error) {
    if args.Folder == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "folder"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.Path == nil || *args.Path == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "path"} 
    }
    routeValues["path"] = *args.Path

    body, marshalErr := json.Marshal(*args.Folder)
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

// Arguments for the UpdateFolder function
type UpdateFolderArgs struct {
    // (required) folder.
    Folder *Folder
    // (required) Project ID or project name
    Project *string
    // (required) Path of the folder to update.
    Path *string
}

// [Preview API] Updates the gate for a deployment.
func (client *Client) UpdateGates(ctx context.Context, args UpdateGatesArgs) (*ReleaseGates, error) {
    if args.GateUpdateMetadata == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "gateUpdateMetadata"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.GateStepId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "gateStepId"} 
    }
    routeValues["gateStepId"] = strconv.Itoa(*args.GateStepId)

    body, marshalErr := json.Marshal(*args.GateUpdateMetadata)
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

// Arguments for the UpdateGates function
type UpdateGatesArgs struct {
    // (required) Metadata to patch the Release Gates.
    GateUpdateMetadata *GateUpdateMetadata
    // (required) Project ID or project name
    Project *string
    // (required) Gate step Id.
    GateStepId *int
}

// [Preview API] Get logs for a release Id.
func (client *Client) GetLogs(ctx context.Context, args GetLogsArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)

    locationId, _ := uuid.Parse("c37fbab5-214b-48e4-a55b-cb6b4f6e4038")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetLogs function
type GetLogsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
}

// [Preview API] Gets the task log of a release as a plain text file.
func (client *Client) GetTaskLog(ctx context.Context, args GetTaskLogArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)
    if args.EnvironmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "environmentId"} 
    }
    routeValues["environmentId"] = strconv.Itoa(*args.EnvironmentId)
    if args.ReleaseDeployPhaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseDeployPhaseId"} 
    }
    routeValues["releaseDeployPhaseId"] = strconv.Itoa(*args.ReleaseDeployPhaseId)
    if args.TaskId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "taskId"} 
    }
    routeValues["taskId"] = strconv.Itoa(*args.TaskId)

    queryParams := url.Values{}
    if args.StartLine != nil {
        queryParams.Add("startLine", strconv.FormatUint(*args.StartLine, 10))
    }
    if args.EndLine != nil {
        queryParams.Add("endLine", strconv.FormatUint(*args.EndLine, 10))
    }
    locationId, _ := uuid.Parse("17c91af7-09fd-4256-bff1-c24ee4f73bc0")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTaskLog function
type GetTaskLogArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Id of release environment.
    EnvironmentId *int
    // (required) Release deploy phase Id.
    ReleaseDeployPhaseId *int
    // (required) ReleaseTask Id for the log.
    TaskId *int
    // (optional) Starting line number for logs
    StartLine *uint64
    // (optional) Ending line number for logs
    EndLine *uint64
}

// Get manual intervention for a given release and manual intervention id.
func (client *Client) GetManualIntervention(ctx context.Context, args GetManualInterventionArgs) (*ManualIntervention, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)
    if args.ManualInterventionId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "manualInterventionId"} 
    }
    routeValues["manualInterventionId"] = strconv.Itoa(*args.ManualInterventionId)

    locationId, _ := uuid.Parse("616c46e4-f370-4456-adaa-fbaf79c7b79e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ManualIntervention
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetManualIntervention function
type GetManualInterventionArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Id of the manual intervention.
    ManualInterventionId *int
}

// List all manual interventions for a given release.
func (client *Client) GetManualInterventions(ctx context.Context, args GetManualInterventionsArgs) (*[]ManualIntervention, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)

    locationId, _ := uuid.Parse("616c46e4-f370-4456-adaa-fbaf79c7b79e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ManualIntervention
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetManualInterventions function
type GetManualInterventionsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
}

// Update manual intervention.
func (client *Client) UpdateManualIntervention(ctx context.Context, args UpdateManualInterventionArgs) (*ManualIntervention, error) {
    if args.ManualInterventionUpdateMetadata == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "manualInterventionUpdateMetadata"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)
    if args.ManualInterventionId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "manualInterventionId"} 
    }
    routeValues["manualInterventionId"] = strconv.Itoa(*args.ManualInterventionId)

    body, marshalErr := json.Marshal(*args.ManualInterventionUpdateMetadata)
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

// Arguments for the UpdateManualIntervention function
type UpdateManualInterventionArgs struct {
    // (required) Meta data to update manual intervention.
    ManualInterventionUpdateMetadata *ManualInterventionUpdateMetadata
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Id of the manual intervention.
    ManualInterventionId *int
}

// Get a list of releases
func (client *Client) GetReleases(ctx context.Context, args GetReleasesArgs) (*[]Release, error) {
    routeValues := make(map[string]string)
    if args.Project != nil && *args.Project != "" {
        routeValues["project"] = *args.Project
    }

    queryParams := url.Values{}
    if args.DefinitionId != nil {
        queryParams.Add("definitionId", strconv.Itoa(*args.DefinitionId))
    }
    if args.DefinitionEnvironmentId != nil {
        queryParams.Add("definitionEnvironmentId", strconv.Itoa(*args.DefinitionEnvironmentId))
    }
    if args.SearchText != nil {
        queryParams.Add("searchText", *args.SearchText)
    }
    if args.CreatedBy != nil {
        queryParams.Add("createdBy", *args.CreatedBy)
    }
    if args.StatusFilter != nil {
        queryParams.Add("statusFilter", string(*args.StatusFilter))
    }
    if args.EnvironmentStatusFilter != nil {
        queryParams.Add("environmentStatusFilter", strconv.Itoa(*args.EnvironmentStatusFilter))
    }
    if args.MinCreatedTime != nil {
        queryParams.Add("minCreatedTime", (*args.MinCreatedTime).String())
    }
    if args.MaxCreatedTime != nil {
        queryParams.Add("maxCreatedTime", (*args.MaxCreatedTime).String())
    }
    if args.QueryOrder != nil {
        queryParams.Add("queryOrder", string(*args.QueryOrder))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", strconv.Itoa(*args.ContinuationToken))
    }
    if args.Expand != nil {
        queryParams.Add("$expand", string(*args.Expand))
    }
    if args.ArtifactTypeId != nil {
        queryParams.Add("artifactTypeId", *args.ArtifactTypeId)
    }
    if args.SourceId != nil {
        queryParams.Add("sourceId", *args.SourceId)
    }
    if args.ArtifactVersionId != nil {
        queryParams.Add("artifactVersionId", *args.ArtifactVersionId)
    }
    if args.SourceBranchFilter != nil {
        queryParams.Add("sourceBranchFilter", *args.SourceBranchFilter)
    }
    if args.IsDeleted != nil {
        queryParams.Add("isDeleted", strconv.FormatBool(*args.IsDeleted))
    }
    if args.TagFilter != nil {
        listAsString := strings.Join((*args.TagFilter)[:], ",")
        queryParams.Add("tagFilter", listAsString)
    }
    if args.PropertyFilters != nil {
        listAsString := strings.Join((*args.PropertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if args.ReleaseIdFilter != nil {
        var stringList []string
        for _, item := range *args.ReleaseIdFilter {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.Path != nil {
        queryParams.Add("path", *args.Path)
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

// Arguments for the GetReleases function
type GetReleasesArgs struct {
    // (optional) Project ID or project name
    Project *string
    // (optional) Releases from this release definition Id.
    DefinitionId *int
    // (optional)
    DefinitionEnvironmentId *int
    // (optional) Releases with names containing searchText.
    SearchText *string
    // (optional) Releases created by this user.
    CreatedBy *string
    // (optional) Releases that have this status.
    StatusFilter *ReleaseStatus
    // (optional)
    EnvironmentStatusFilter *int
    // (optional) Releases that were created after this time.
    MinCreatedTime *time.Time
    // (optional) Releases that were created before this time.
    MaxCreatedTime *time.Time
    // (optional) Gets the results in the defined order of created date for releases. Default is descending.
    QueryOrder *ReleaseQueryOrder
    // (optional) Number of releases to get. Default is 50.
    Top *int
    // (optional) Gets the releases after the continuation token provided.
    ContinuationToken *int
    // (optional) The property that should be expanded in the list of releases.
    Expand *ReleaseExpands
    // (optional) Releases with given artifactTypeId will be returned. Values can be Build, Jenkins, GitHub, Nuget, Team Build (external), ExternalTFSBuild, Git, TFVC, ExternalTfsXamlBuild.
    ArtifactTypeId *string
    // (optional) Unique identifier of the artifact used. e.g. For build it would be {projectGuid}:{BuildDefinitionId}, for Jenkins it would be {JenkinsConnectionId}:{JenkinsDefinitionId}, for TfsOnPrem it would be {TfsOnPremConnectionId}:{ProjectName}:{TfsOnPremDefinitionId}. For third-party artifacts e.g. TeamCity, BitBucket you may refer 'uniqueSourceIdentifier' inside vss-extension.json https://github.com/Microsoft/vsts-rm-extensions/blob/master/Extensions.
    SourceId *string
    // (optional) Releases with given artifactVersionId will be returned. E.g. in case of Build artifactType, it is buildId.
    ArtifactVersionId *string
    // (optional) Releases with given sourceBranchFilter will be returned.
    SourceBranchFilter *string
    // (optional) Gets the soft deleted releases, if true.
    IsDeleted *bool
    // (optional) A comma-delimited list of tags. Only releases with these tags will be returned.
    TagFilter *[]string
    // (optional) A comma-delimited list of extended properties to be retrieved. If set, the returned Releases will contain values for the specified property Ids (if they exist). If not set, properties will not be included. Note that this will not filter out any Release from results irrespective of whether it has property set or not.
    PropertyFilters *[]string
    // (optional) A comma-delimited list of releases Ids. Only releases with these Ids will be returned.
    ReleaseIdFilter *[]int
    // (optional) Releases under this folder path will be returned
    Path *string
}

// Create a release.
func (client *Client) CreateRelease(ctx context.Context, args CreateReleaseArgs) (*Release, error) {
    if args.ReleaseStartMetadata == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseStartMetadata"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.ReleaseStartMetadata)
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

// Arguments for the CreateRelease function
type CreateReleaseArgs struct {
    // (required) Metadata to create a release.
    ReleaseStartMetadata *ReleaseStartMetadata
    // (required) Project ID or project name
    Project *string
}

// Get a Release
func (client *Client) GetRelease(ctx context.Context, args GetReleaseArgs) (*Release, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)

    queryParams := url.Values{}
    if args.ApprovalFilters != nil {
        queryParams.Add("approvalFilters", string(*args.ApprovalFilters))
    }
    if args.PropertyFilters != nil {
        listAsString := strings.Join((*args.PropertyFilters)[:], ",")
        queryParams.Add("propertyFilters", listAsString)
    }
    if args.Expand != nil {
        queryParams.Add("$expand", string(*args.Expand))
    }
    if args.TopGateRecords != nil {
        queryParams.Add("$topGateRecords", strconv.Itoa(*args.TopGateRecords))
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

// Arguments for the GetRelease function
type GetReleaseArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (optional) A filter which would allow fetching approval steps selectively based on whether it is automated, or manual. This would also decide whether we should fetch pre and post approval snapshots. Assumes All by default
    ApprovalFilters *ApprovalFilters
    // (optional) A comma-delimited list of extended properties to be retrieved. If set, the returned Release will contain values for the specified property Ids (if they exist). If not set, properties will not be included.
    PropertyFilters *[]string
    // (optional) A property that should be expanded in the release.
    Expand *SingleReleaseExpands
    // (optional) Number of release gate records to get. Default is 5.
    TopGateRecords *int
}

// Get release for a given revision number.
func (client *Client) GetReleaseRevision(ctx context.Context, args GetReleaseRevisionArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)

    queryParams := url.Values{}
    if args.DefinitionSnapshotRevision == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "definitionSnapshotRevision"}
    }
    queryParams.Add("definitionSnapshotRevision", strconv.Itoa(*args.DefinitionSnapshotRevision))
    locationId, _ := uuid.Parse("a166fde7-27ad-408e-ba75-703c2cc9d500")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetReleaseRevision function
type GetReleaseRevisionArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release.
    ReleaseId *int
    // (required) Definition snapshot revision number.
    DefinitionSnapshotRevision *int
}

// Update a complete release object.
func (client *Client) UpdateRelease(ctx context.Context, args UpdateReleaseArgs) (*Release, error) {
    if args.Release == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "release"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)

    body, marshalErr := json.Marshal(*args.Release)
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

// Arguments for the UpdateRelease function
type UpdateReleaseArgs struct {
    // (required) Release object for update.
    Release *Release
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release to update.
    ReleaseId *int
}

// Update few properties of a release.
func (client *Client) UpdateReleaseResource(ctx context.Context, args UpdateReleaseResourceArgs) (*Release, error) {
    if args.ReleaseUpdateMetadata == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseUpdateMetadata"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"} 
    }
    routeValues["releaseId"] = strconv.Itoa(*args.ReleaseId)

    body, marshalErr := json.Marshal(*args.ReleaseUpdateMetadata)
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

// Arguments for the UpdateReleaseResource function
type UpdateReleaseResourceArgs struct {
    // (required) Properties of release to update.
    ReleaseUpdateMetadata *ReleaseUpdateMetadata
    // (required) Project ID or project name
    Project *string
    // (required) Id of the release to update.
    ReleaseId *int
}

// [Preview API] Get release definition for a given definitionId and revision
func (client *Client) GetDefinitionRevision(ctx context.Context, args GetDefinitionRevisionArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.DefinitionId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)
    if args.Revision == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "revision"} 
    }
    routeValues["revision"] = strconv.Itoa(*args.Revision)

    locationId, _ := uuid.Parse("258b82e0-9d41-43f3-86d6-fef14ddd44bc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetDefinitionRevision function
type GetDefinitionRevisionArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the definition.
    DefinitionId *int
    // (required) Id of the revision.
    Revision *int
}

// [Preview API] Get revision history for a release definition
func (client *Client) GetReleaseDefinitionHistory(ctx context.Context, args GetReleaseDefinitionHistoryArgs) (*[]ReleaseDefinitionRevision, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.DefinitionId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "definitionId"} 
    }
    routeValues["definitionId"] = strconv.Itoa(*args.DefinitionId)

    locationId, _ := uuid.Parse("258b82e0-9d41-43f3-86d6-fef14ddd44bc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ReleaseDefinitionRevision
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetReleaseDefinitionHistory function
type GetReleaseDefinitionHistoryArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the definition.
    DefinitionId *int
}

