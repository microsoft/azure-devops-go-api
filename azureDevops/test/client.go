// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package test

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

var ResourceAreaId, _ = uuid.Parse("c2aa639c-3ccc-4740-b3b6-ce2a1e1d984e")

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

// Gets the action results for an iteration in a test result.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test result that contains the iterations.
// iterationId (required): ID of the iteration that contains the actions.
// actionPath (optional): Path of a specific action, used to get just that action.
func (client Client) GetActionResults(ctx context.Context, project *string, runId *int, testCaseResultId *int, iterationId *int, actionPath *string) (*[]TestActionResultModel, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)
    if actionPath != nil && *actionPath != "" {
        routeValues["actionPath"] = *actionPath
    }

    locationId, _ := uuid.Parse("eaf40c31-ff84-4062-aafd-d5664be11a37")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestActionResultModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Attach a file to a test result.
// ctx
// attachmentRequestModel (required): Attachment details TestAttachmentRequestModel
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test result against which attachment has to be uploaded.
func (client Client) CreateTestResultAttachment(ctx context.Context, attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int, testCaseResultId *int) (*TestAttachmentReference, error) {
    if attachmentRequestModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    body, marshalErr := json.Marshal(*attachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Attach a file to a test result
// ctx
// attachmentRequestModel (required): Attachment Request Model.
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test results that contains sub result.
// testSubResultId (required): ID of the test sub results against which attachment has to be uploaded.
func (client Client) CreateTestSubResultAttachment(ctx context.Context, attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int, testCaseResultId *int, testSubResultId *int) (*TestAttachmentReference, error) {
    if attachmentRequestModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    body, marshalErr := json.Marshal(*attachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Download a test result attachment by its ID.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the testCaseResultId.
// testCaseResultId (required): ID of the test result whose attachment has to be downloaded.
// attachmentId (required): ID of the test result attachment to be downloaded.
func (client Client) GetTestResultAttachmentContent(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get list of test result attachments reference.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test result.
func (client Client) GetTestResultAttachments(ctx context.Context, project *string, runId *int, testCaseResultId *int) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Download a test result attachment by its ID.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the testCaseResultId.
// testCaseResultId (required): ID of the test result whose attachment has to be downloaded.
// attachmentId (required): ID of the test result attachment to be downloaded.
func (client Client) GetTestResultAttachmentZip(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Download a test sub result attachment
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test results that contains sub result.
// attachmentId (required): ID of the test result attachment to be downloaded
// testSubResultId (required): ID of the test sub result whose attachment has to be downloaded
func (client Client) GetTestSubResultAttachmentContent(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int, testSubResultId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get list of test sub result attachments
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test results that contains sub result.
// testSubResultId (required): ID of the test sub result whose attachment has to be downloaded
func (client Client) GetTestSubResultAttachments(ctx context.Context, project *string, runId *int, testCaseResultId *int, testSubResultId *int) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Download a test sub result attachment
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test results that contains sub result.
// attachmentId (required): ID of the test result attachment to be downloaded
// testSubResultId (required): ID of the test sub result whose attachment has to be downloaded
func (client Client) GetTestSubResultAttachmentZip(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int, testSubResultId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Attach a file to a test run.
// ctx
// attachmentRequestModel (required): Attachment details TestAttachmentRequestModel
// project (required): Project ID or project name
// runId (required): ID of the test run against which attachment has to be uploaded.
func (client Client) CreateTestRunAttachment(ctx context.Context, attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int) (*TestAttachmentReference, error) {
    if attachmentRequestModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*attachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Download a test run attachment by its ID.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run whose attachment has to be downloaded.
// attachmentId (required): ID of the test run attachment to be downloaded.
func (client Client) GetTestRunAttachmentContent(ctx context.Context, project *string, runId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if attachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get list of test run attachments reference.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run.
func (client Client) GetTestRunAttachments(ctx context.Context, project *string, runId *int) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Download a test run attachment by its ID.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run whose attachment has to be downloaded.
// attachmentId (required): ID of the test run attachment to be downloaded.
func (client Client) GetTestRunAttachmentZip(ctx context.Context, project *string, runId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if attachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get code coverage data for a build.
// ctx
// project (required): Project ID or project name
// buildId (required): ID of the build for which code coverage data needs to be fetched.
// flags (required): Value of flags determine the level of code coverage details to be fetched. Flags are additive. Expected Values are 1 for Modules, 2 for Functions, 4 for BlockData.
func (client Client) GetBuildCodeCoverage(ctx context.Context, project *string, buildId *int, flags *int) (*[]BuildCoverage, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*buildId))
    if flags == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "flags"}
    }
    queryParams.Add("flags", strconv.Itoa(*flags))
    locationId, _ := uuid.Parse("77560e8a-4e8c-4d59-894e-a5f264c24444")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get code coverage data for a test run
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run for which code coverage data needs to be fetched.
// flags (required): Value of flags determine the level of code coverage details to be fetched. Flags are additive. Expected Values are 1 for Modules, 2 for Functions, 4 for BlockData.
func (client Client) GetTestRunCodeCoverage(ctx context.Context, project *string, runId *int, flags *int) (*[]TestRunCoverage, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if flags == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "flags"}
    }
    queryParams.Add("flags", strconv.Itoa(*flags))
    locationId, _ := uuid.Parse("9629116f-3b89-4ed8-b358-d4694efda160")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRunCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get iteration for a result
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test result that contains the iterations.
// iterationId (required): Id of the test results Iteration.
// includeActionResults (optional): Include result details for each action performed in the test iteration. ActionResults refer to outcome (pass/fail) of test steps that are executed as part of a running a manual test. Including the ActionResults flag gets the outcome of test steps in the actionResults section and test parameters in the parameters section for each test iteration.
func (client Client) GetTestIteration(ctx context.Context, project *string, runId *int, testCaseResultId *int, iterationId *int, includeActionResults *bool) (*TestIterationDetailsModel, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    queryParams := url.Values{}
    if includeActionResults != nil {
        queryParams.Add("includeActionResults", strconv.FormatBool(*includeActionResults))
    }
    locationId, _ := uuid.Parse("73eb9074-3446-4c44-8296-2f811950ff8d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestIterationDetailsModel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get iterations for a result
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test result that contains the iterations.
// includeActionResults (optional): Include result details for each action performed in the test iteration. ActionResults refer to outcome (pass/fail) of test steps that are executed as part of a running a manual test. Including the ActionResults flag gets the outcome of test steps in the actionResults section and test parameters in the parameters section for each test iteration.
func (client Client) GetTestIterations(ctx context.Context, project *string, runId *int, testCaseResultId *int, includeActionResults *bool) (*[]TestIterationDetailsModel, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    queryParams := url.Values{}
    if includeActionResults != nil {
        queryParams.Add("includeActionResults", strconv.FormatBool(*includeActionResults))
    }
    locationId, _ := uuid.Parse("73eb9074-3446-4c44-8296-2f811950ff8d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestIterationDetailsModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of parameterized results
// ctx
// project (required): Project ID or project name
// runId (required): ID of the test run that contains the result.
// testCaseResultId (required): ID of the test result that contains the iterations.
// iterationId (required): ID of the iteration that contains the parameterized results.
// paramName (optional): Name of the parameter.
func (client Client) GetResultParameters(ctx context.Context, project *string, runId *int, testCaseResultId *int, iterationId *int, paramName *string) (*[]TestResultParameterModel, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*iterationId)

    queryParams := url.Values{}
    if paramName != nil {
        queryParams.Add("paramName", *paramName)
    }
    locationId, _ := uuid.Parse("7c69810d-3354-4af3-844a-180bd25db08a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestResultParameterModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a test point.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan.
// suiteId (required): ID of the suite that contains the point.
// pointIds (required): ID of the test point to get.
// witFields (optional): Comma-separated list of work item field names.
func (client Client) GetPoint(ctx context.Context, project *string, planId *int, suiteId *int, pointIds *int, witFields *string) (*TestPoint, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)
    if pointIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pointIds"} 
    }
    routeValues["pointIds"] = strconv.Itoa(*pointIds)

    queryParams := url.Values{}
    if witFields != nil {
        queryParams.Add("witFields", *witFields)
    }
    locationId, _ := uuid.Parse("3bcfd5c8-be62-488e-b1da-b8289ce9299c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestPoint
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of test points.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan.
// suiteId (required): ID of the suite that contains the points.
// witFields (optional): Comma-separated list of work item field names.
// configurationId (optional): Get test points for specific configuration.
// testCaseId (optional): Get test points for a specific test case, valid when configurationId is not set.
// testPointIds (optional): Get test points for comma-separated list of test point IDs, valid only when configurationId and testCaseId are not set.
// includePointDetails (optional): Include all properties for the test point.
// skip (optional): Number of test points to skip..
// top (optional): Number of test points to return.
func (client Client) GetPoints(ctx context.Context, project *string, planId *int, suiteId *int, witFields *string, configurationId *string, testCaseId *string, testPointIds *string, includePointDetails *bool, skip *int, top *int) (*[]TestPoint, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)

    queryParams := url.Values{}
    if witFields != nil {
        queryParams.Add("witFields", *witFields)
    }
    if configurationId != nil {
        queryParams.Add("configurationId", *configurationId)
    }
    if testCaseId != nil {
        queryParams.Add("testCaseId", *testCaseId)
    }
    if testPointIds != nil {
        queryParams.Add("testPointIds", *testPointIds)
    }
    if includePointDetails != nil {
        queryParams.Add("includePointDetails", strconv.FormatBool(*includePointDetails))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("3bcfd5c8-be62-488e-b1da-b8289ce9299c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update test points.
// ctx
// pointUpdateModel (required): Data to update.
// project (required): Project ID or project name
// planId (required): ID of the test plan.
// suiteId (required): ID of the suite that contains the points.
// pointIds (required): ID of the test point to get. Use a comma-separated list of IDs to update multiple test points.
func (client Client) UpdateTestPoints(ctx context.Context, pointUpdateModel *PointUpdateModel, project *string, planId *int, suiteId *int, pointIds *string) (*[]TestPoint, error) {
    if pointUpdateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pointUpdateModel"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)
    if pointIds == nil || *pointIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pointIds"} 
    }
    routeValues["pointIds"] = *pointIds

    body, marshalErr := json.Marshal(*pointUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("3bcfd5c8-be62-488e-b1da-b8289ce9299c")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get test points using query.
// ctx
// query (required): TestPointsQuery to get test points.
// project (required): Project ID or project name
// skip (optional): Number of test points to skip..
// top (optional): Number of test points to return.
func (client Client) GetPointsByQuery(ctx context.Context, query *TestPointsQuery, project *string, skip *int, top *int) (*TestPointsQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b4264fd0-a5d1-43e2-82a5-b9c46b7da9ce")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestPointsQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get test result retention settings
// ctx
// project (required): Project ID or project name
func (client Client) GetResultRetentionSettings(ctx context.Context, project *string) (*ResultRetentionSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    locationId, _ := uuid.Parse("a3206d9e-fa8d-42d3-88cb-f75c51e69cde")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ResultRetentionSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update test result retention settings
// ctx
// retentionSettings (required): Test result retention settings details to be updated
// project (required): Project ID or project name
func (client Client) UpdateResultRetentionSettings(ctx context.Context, retentionSettings *ResultRetentionSettings, project *string) (*ResultRetentionSettings, error) {
    if retentionSettings == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "retentionSettings"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*retentionSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a3206d9e-fa8d-42d3-88cb-f75c51e69cde")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ResultRetentionSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Add test results to a test run.
// ctx
// results (required): List of test results to add.
// project (required): Project ID or project name
// runId (required): Test run ID into which test results to add.
func (client Client) AddTestResultsToTestRun(ctx context.Context, results *[]TestCaseResult, project *string, runId *int) (*[]TestCaseResult, error) {
    if results == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "results"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a test result for a test run.
// ctx
// project (required): Project ID or project name
// runId (required): Test run ID of a test result to fetch.
// testCaseResultId (required): Test result ID.
// detailsToInclude (optional): Details to include with test results. Default is None. Other values are Iterations, WorkItems and SubResults.
func (client Client) GetTestResultById(ctx context.Context, project *string, runId *int, testCaseResultId *int, detailsToInclude *ResultDetails) (*TestCaseResult, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    queryParams := url.Values{}
    if detailsToInclude != nil {
        queryParams.Add("detailsToInclude", string(*detailsToInclude))
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestCaseResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get test results for a test run.
// ctx
// project (required): Project ID or project name
// runId (required): Test run ID of test results to fetch.
// detailsToInclude (optional): Details to include with test results. Default is None. Other values are Iterations and WorkItems.
// skip (optional): Number of test results to skip from beginning.
// top (optional): Number of test results to return. Maximum is 1000 when detailsToInclude is None and 200 otherwise.
// outcomes (optional): Comma separated list of test outcomes to filter test results.
func (client Client) GetTestResults(ctx context.Context, project *string, runId *int, detailsToInclude *ResultDetails, skip *int, top *int, outcomes *[]TestOutcome) (*[]TestCaseResult, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if detailsToInclude != nil {
        queryParams.Add("detailsToInclude", string(*detailsToInclude))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if outcomes != nil {
        var stringList []string
        for _, item := range *outcomes {
            stringList = append(stringList, string(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update test results in a test run.
// ctx
// results (required): List of test results to update.
// project (required): Project ID or project name
// runId (required): Test run ID whose test results to update.
func (client Client) UpdateTestResults(ctx context.Context, results *[]TestCaseResult, project *string, runId *int) (*[]TestCaseResult, error) {
    if results == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "results"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get test run statistics , used when we want to get summary of a run by outcome.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the run to get.
func (client Client) GetTestRunStatistics(ctx context.Context, project *string, runId *int) (*TestRunStatistic, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    locationId, _ := uuid.Parse("0a42c424-d764-4a16-a2d5-5c85f87d0ae8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRunStatistic
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Create new test run.
// ctx
// testRun (required): Run details RunCreateModel
// project (required): Project ID or project name
func (client Client) CreateTestRun(ctx context.Context, testRun *RunCreateModel, project *string) (*TestRun, error) {
    if testRun == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testRun"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testRun)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a test run by its ID.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the run to delete.
func (client Client) DeleteTestRun(ctx context.Context, project *string, runId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a test run by its ID.
// ctx
// project (required): Project ID or project name
// runId (required): ID of the run to get.
// includeDetails (optional): Defualt value is true. It includes details like run statistics,release,build,Test enviornment,Post process state and more
func (client Client) GetTestRunById(ctx context.Context, project *string, runId *int, includeDetails *bool) (*TestRun, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if includeDetails != nil {
        queryParams.Add("includeDetails", strconv.FormatBool(*includeDetails))
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of test runs.
// ctx
// project (required): Project ID or project name
// buildUri (optional): URI of the build that the runs used.
// owner (optional): Team foundation ID of the owner of the runs.
// tmiRunId (optional)
// planId (optional): ID of the test plan that the runs are a part of.
// includeRunDetails (optional): If true, include all the properties of the runs.
// automated (optional): If true, only returns automated runs.
// skip (optional): Number of test runs to skip.
// top (optional): Number of test runs to return.
func (client Client) GetTestRuns(ctx context.Context, project *string, buildUri *string, owner *string, tmiRunId *string, planId *int, includeRunDetails *bool, automated *bool, skip *int, top *int) (*[]TestRun, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildUri != nil {
        queryParams.Add("buildUri", *buildUri)
    }
    if owner != nil {
        queryParams.Add("owner", *owner)
    }
    if tmiRunId != nil {
        queryParams.Add("tmiRunId", *tmiRunId)
    }
    if planId != nil {
        queryParams.Add("planId", strconv.Itoa(*planId))
    }
    if includeRunDetails != nil {
        queryParams.Add("includeRunDetails", strconv.FormatBool(*includeRunDetails))
    }
    if automated != nil {
        queryParams.Add("automated", strconv.FormatBool(*automated))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Query Test Runs based on filters. Mandatory fields are minLastUpdatedDate and maxLastUpdatedDate.
// ctx
// project (required): Project ID or project name
// minLastUpdatedDate (required): Minimum Last Modified Date of run to be queried (Mandatory).
// maxLastUpdatedDate (required): Maximum Last Modified Date of run to be queried (Mandatory, difference between min and max date can be atmost 7 days).
// state (optional): Current state of the Runs to be queried.
// planIds (optional): Plan Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
// isAutomated (optional): Automation type of the Runs to be queried.
// publishContext (optional): PublishContext of the Runs to be queried.
// buildIds (optional): Build Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
// buildDefIds (optional): Build Definition Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
// branchName (optional): Source Branch name of the Runs to be queried.
// releaseIds (optional): Release Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
// releaseDefIds (optional): Release Definition Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
// releaseEnvIds (optional): Release Environment Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
// releaseEnvDefIds (optional): Release Environment Definition Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
// runTitle (optional): Run Title of the Runs to be queried.
// top (optional): Number of runs to be queried. Limit is 100
// continuationToken (optional): continuationToken received from previous batch or null for first batch. It is not supposed to be created (or altered, if received from last batch) by user.
func (client Client) QueryTestRuns(ctx context.Context, project *string, minLastUpdatedDate *time.Time, maxLastUpdatedDate *time.Time, state *TestRunState, planIds *[]int, isAutomated *bool, publishContext *TestRunPublishContext, buildIds *[]int, buildDefIds *[]int, branchName *string, releaseIds *[]int, releaseDefIds *[]int, releaseEnvIds *[]int, releaseEnvDefIds *[]int, runTitle *string, top *int, continuationToken *string) (*[]TestRun, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if minLastUpdatedDate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "minLastUpdatedDate"}
    }
    queryParams.Add("minLastUpdatedDate", (*minLastUpdatedDate).String())
    if maxLastUpdatedDate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "maxLastUpdatedDate"}
    }
    queryParams.Add("maxLastUpdatedDate", (*maxLastUpdatedDate).String())
    if state != nil {
        queryParams.Add("state", string(*state))
    }
    if planIds != nil {
        var stringList []string
        for _, item := range *planIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if isAutomated != nil {
        queryParams.Add("isAutomated", strconv.FormatBool(*isAutomated))
    }
    if publishContext != nil {
        queryParams.Add("publishContext", string(*publishContext))
    }
    if buildIds != nil {
        var stringList []string
        for _, item := range *buildIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if buildDefIds != nil {
        var stringList []string
        for _, item := range *buildDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if branchName != nil {
        queryParams.Add("branchName", *branchName)
    }
    if releaseIds != nil {
        var stringList []string
        for _, item := range *releaseIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if releaseDefIds != nil {
        var stringList []string
        for _, item := range *releaseDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if releaseEnvIds != nil {
        var stringList []string
        for _, item := range *releaseEnvIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if releaseEnvDefIds != nil {
        var stringList []string
        for _, item := range *releaseEnvDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if runTitle != nil {
        queryParams.Add("runTitle", *runTitle)
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update test run by its ID.
// ctx
// runUpdateModel (required): Run details RunUpdateModel
// project (required): Project ID or project name
// runId (required): ID of the run to update.
func (client Client) UpdateTestRun(ctx context.Context, runUpdateModel *RunUpdateModel, project *string, runId *int) (*TestRun, error) {
    if runUpdateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runUpdateModel"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*runUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a test session
// ctx
// testSession (required): Test session details for creation
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) CreateTestSession(ctx context.Context, testSession *TestSession, project *string, team *string) (*TestSession, error) {
    if testSession == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSession"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    body, marshalErr := json.Marshal(*testSession)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1500b4b4-6c69-4ca6-9b18-35e9e97fe2ac")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSession
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of test sessions
// ctx
// project (required): Project ID or project name
// team (optional): Team ID or team name
// period (optional): Period in days from now, for which test sessions are fetched.
// allSessions (optional): If false, returns test sessions for current user. Otherwise, it returns test sessions for all users
// includeAllProperties (optional): If true, it returns all properties of the test sessions. Otherwise, it returns the skinny version.
// source (optional): Source of the test session.
// includeOnlyCompletedSessions (optional): If true, it returns test sessions in completed state. Otherwise, it returns test sessions for all states
func (client Client) GetTestSessions(ctx context.Context, project *string, team *string, period *int, allSessions *bool, includeAllProperties *bool, source *TestSessionSource, includeOnlyCompletedSessions *bool) (*[]TestSession, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    queryParams := url.Values{}
    if period != nil {
        queryParams.Add("period", strconv.Itoa(*period))
    }
    if allSessions != nil {
        queryParams.Add("allSessions", strconv.FormatBool(*allSessions))
    }
    if includeAllProperties != nil {
        queryParams.Add("includeAllProperties", strconv.FormatBool(*includeAllProperties))
    }
    if source != nil {
        queryParams.Add("source", string(*source))
    }
    if includeOnlyCompletedSessions != nil {
        queryParams.Add("includeOnlyCompletedSessions", strconv.FormatBool(*includeOnlyCompletedSessions))
    }
    locationId, _ := uuid.Parse("1500b4b4-6c69-4ca6-9b18-35e9e97fe2ac")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestSession
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a test session
// ctx
// testSession (required): Test session details for update
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) UpdateTestSession(ctx context.Context, testSession *TestSession, project *string, team *string) (*TestSession, error) {
    if testSession == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSession"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    body, marshalErr := json.Marshal(*testSession)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1500b4b4-6c69-4ca6-9b18-35e9e97fe2ac")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSession
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Add test cases to suite.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suite.
// suiteId (required): ID of the test suite to which the test cases must be added.
// testCaseIds (required): IDs of the test cases to add to the suite. Ids are specified in comma separated format.
func (client Client) AddTestCasesToSuite(ctx context.Context, project *string, planId *int, suiteId *int, testCaseIds *string) (*[]SuiteTestCase, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)
    if testCaseIds == nil || *testCaseIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = *testCaseIds
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteTestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a specific test case in a test suite with test case id.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suites.
// suiteId (required): ID of the suite that contains the test case.
// testCaseIds (required): ID of the test case to get.
func (client Client) GetTestCaseById(ctx context.Context, project *string, planId *int, suiteId *int, testCaseIds *int) (*SuiteTestCase, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)
    if testCaseIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = strconv.Itoa(*testCaseIds)
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SuiteTestCase
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get all test cases in a suite.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suites.
// suiteId (required): ID of the suite to get.
func (client Client) GetTestCases(ctx context.Context, project *string, planId *int, suiteId *int) (*[]SuiteTestCase, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteTestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// The test points associated with the test cases are removed from the test suite. The test case work item is not deleted from the system. See test cases resource to delete a test case permanently.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suite.
// suiteId (required): ID of the suite to get.
// testCaseIds (required): IDs of the test cases to remove from the suite.
func (client Client) RemoveTestCasesFromSuiteUrl(ctx context.Context, project *string, planId *int, suiteId *int, testCaseIds *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)
    if testCaseIds == nil || *testCaseIds == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = *testCaseIds
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Updates the properties of the test case association in a suite.
// ctx
// suiteTestCaseUpdateModel (required): Model for updation of the properties of test case suite association.
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suite.
// suiteId (required): ID of the test suite to which the test cases must be added.
// testCaseIds (required): IDs of the test cases to add to the suite. Ids are specified in comma separated format.
func (client Client) UpdateSuiteTestCases(ctx context.Context, suiteTestCaseUpdateModel *SuiteTestCaseUpdateModel, project *string, planId *int, suiteId *int, testCaseIds *string) (*[]SuiteTestCase, error) {
    if suiteTestCaseUpdateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteTestCaseUpdateModel"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)
    if testCaseIds == nil || *testCaseIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = *testCaseIds
    routeValues["action"] = "TestCases"

    body, marshalErr := json.Marshal(*suiteTestCaseUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteTestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a test case.
// ctx
// project (required): Project ID or project name
// testCaseId (required): Id of test case to delete.
func (client Client) DeleteTestCase(ctx context.Context, project *string, testCaseId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if testCaseId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testCaseId"} 
    }
    routeValues["testCaseId"] = strconv.Itoa(*testCaseId)

    locationId, _ := uuid.Parse("4d472e0f-e32c-4ef8-adf4-a4078772889c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get history of a test method using TestHistoryQuery
// ctx
// filter (required): TestHistoryQuery to get history
// project (required): Project ID or project name
func (client Client) QueryTestHistory(ctx context.Context, filter *TestHistoryQuery, project *string) (*TestHistoryQuery, error) {
    if filter == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filter"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("929fd86c-3e38-4d8c-b4b6-90df256e5971")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestHistoryQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

