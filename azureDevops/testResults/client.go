// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testResults

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
    "strings"
    "time"
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

// [Preview API]
// ctx
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// iterationId (required)
// actionPath (optional)
func (client Client) CreateTestIterationResultAttachment(ctx context.Context, attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int, testCaseResultId *int, iterationId *int, actionPath *string) (*TestAttachmentReference, error) {
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
    if iterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"}
    }
    queryParams.Add("iterationId", strconv.Itoa(*iterationId))
    if actionPath != nil {
        queryParams.Add("actionPath", *actionPath)
    }
    body, marshalErr := json.Marshal(*attachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
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
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// testSubResultId (required)
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
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
func (client Client) DeleteTestResultAttachment(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a test result attachment
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
func (client Client) GetTestResultAttachmentContent(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int) (io.ReadCloser, error) {
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

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
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

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a test result attachment
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
func (client Client) GetTestResultAttachmentZip(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int) (io.ReadCloser, error) {
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

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// [Preview API] Returns a test sub result attachment
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
// testSubResultId (required)
func (client Client) GetTestSubResultAttachmentContent(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int, testSubResultId *int) (io.ReadCloser, error) {
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
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// [Preview API] Returns attachment references for test sub result.
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// testSubResultId (required)
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
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a test sub result attachment
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
// testSubResultId (required)
func (client Client) GetTestSubResultAttachmentZip(ctx context.Context, project *string, runId *int, testCaseResultId *int, attachmentId *int, testSubResultId *int) (io.ReadCloser, error) {
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
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// [Preview API]
// ctx
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
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
    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// attachmentId (required)
func (client Client) DeleteTestRunAttachment(ctx context.Context, project *string, runId *int, attachmentId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if attachmentId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a test run attachment
// ctx
// project (required): Project ID or project name
// runId (required)
// attachmentId (required)
func (client Client) GetTestRunAttachmentContent(ctx context.Context, project *string, runId *int, attachmentId *int) (io.ReadCloser, error) {
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

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
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

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a test run attachment
// ctx
// project (required): Project ID or project name
// runId (required)
// attachmentId (required)
func (client Client) GetTestRunAttachmentZip(ctx context.Context, project *string, runId *int, attachmentId *int) (io.ReadCloser, error) {
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

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
func (client Client) GetBugsLinkedToTestResult(ctx context.Context, project *string, runId *int, testCaseResultId *int) (*[]WorkItemReference, error) {
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

    locationId, _ := uuid.Parse("d8dbf98f-eb34-4f8d-8365-47972af34f29")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// buildId (required)
// flags (required)
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
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// buildId (required)
// deltaBuildId (optional)
func (client Client) GetCodeCoverageSummary(ctx context.Context, project *string, buildId *int, deltaBuildId *int) (*CodeCoverageSummary, error) {
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
    if deltaBuildId != nil {
        queryParams.Add("deltaBuildId", strconv.Itoa(*deltaBuildId))
    }
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CodeCoverageSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] http://(tfsserver):8080/tfs/DefaultCollection/_apis/test/CodeCoverage?buildId=10 Request: Json of code coverage summary
// ctx
// project (required): Project ID or project name
// buildId (required)
// coverageData (optional)
func (client Client) UpdateCodeCoverageSummary(ctx context.Context, project *string, buildId *int, coverageData *CodeCoverageData) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*buildId))
    body, marshalErr := json.Marshal(*coverageData)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    _, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// flags (required)
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
    locationId, _ := uuid.Parse("5641efbc-6f9b-401a-baeb-d3da22489e5e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRunCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// filter (required)
// project (required): Project ID or project name
func (client Client) QueryTestResultHistory(ctx context.Context, filter *ResultsFilter, project *string) (*TestResultHistory, error) {
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
    locationId, _ := uuid.Parse("bdf7a97b-0395-4da8-9d5d-f957619327d1")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultHistory
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// buildId (required)
// publishContext (optional)
// groupBy (optional)
// filter (optional)
// orderby (optional)
// shouldIncludeResults (optional)
// queryRunSummaryForInProgress (optional)
func (client Client) GetTestResultDetailsForBuild(ctx context.Context, project *string, buildId *int, publishContext *string, groupBy *string, filter *string, orderby *string, shouldIncludeResults *bool, queryRunSummaryForInProgress *bool) (*TestResultsDetails, error) {
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
    if publishContext != nil {
        queryParams.Add("publishContext", *publishContext)
    }
    if groupBy != nil {
        queryParams.Add("groupBy", *groupBy)
    }
    if filter != nil {
        queryParams.Add("$filter", *filter)
    }
    if orderby != nil {
        queryParams.Add("$orderby", *orderby)
    }
    if shouldIncludeResults != nil {
        queryParams.Add("shouldIncludeResults", strconv.FormatBool(*shouldIncludeResults))
    }
    if queryRunSummaryForInProgress != nil {
        queryParams.Add("queryRunSummaryForInProgress", strconv.FormatBool(*queryRunSummaryForInProgress))
    }
    locationId, _ := uuid.Parse("a518c749-4524-45b2-a7ef-1ac009b312cd")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// releaseId (required)
// releaseEnvId (required)
// publishContext (optional)
// groupBy (optional)
// filter (optional)
// orderby (optional)
// shouldIncludeResults (optional)
// queryRunSummaryForInProgress (optional)
func (client Client) GetTestResultDetailsForRelease(ctx context.Context, project *string, releaseId *int, releaseEnvId *int, publishContext *string, groupBy *string, filter *string, orderby *string, shouldIncludeResults *bool, queryRunSummaryForInProgress *bool) (*TestResultsDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"}
    }
    queryParams.Add("releaseId", strconv.Itoa(*releaseId))
    if releaseEnvId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseEnvId"}
    }
    queryParams.Add("releaseEnvId", strconv.Itoa(*releaseEnvId))
    if publishContext != nil {
        queryParams.Add("publishContext", *publishContext)
    }
    if groupBy != nil {
        queryParams.Add("groupBy", *groupBy)
    }
    if filter != nil {
        queryParams.Add("$filter", *filter)
    }
    if orderby != nil {
        queryParams.Add("$orderby", *orderby)
    }
    if shouldIncludeResults != nil {
        queryParams.Add("shouldIncludeResults", strconv.FormatBool(*shouldIncludeResults))
    }
    if queryRunSummaryForInProgress != nil {
        queryParams.Add("queryRunSummaryForInProgress", strconv.FormatBool(*queryRunSummaryForInProgress))
    }
    locationId, _ := uuid.Parse("19a8183a-69fb-47d7-bfbf-1b6b0d921294")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// query (required)
// project (required): Project ID or project name
func (client Client) GetTestResultsByQuery(ctx context.Context, query *TestResultsQuery, project *string) (*TestResultsQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("14033a2c-af25-4af1-9e39-8ef6900482e3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// queryModel (required)
// project (required): Project ID or project name
// includeResultDetails (optional)
// includeIterationDetails (optional)
// skip (optional)
// top (optional)
func (client Client) GetTestResultsByQueryWiql(ctx context.Context, queryModel *QueryModel, project *string, includeResultDetails *bool, includeIterationDetails *bool, skip *int, top *int) (*[]TestCaseResult, error) {
    if queryModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "queryModel"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if includeResultDetails != nil {
        queryParams.Add("includeResultDetails", strconv.FormatBool(*includeResultDetails))
    }
    if includeIterationDetails != nil {
        queryParams.Add("includeIterationDetails", strconv.FormatBool(*includeIterationDetails))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    body, marshalErr := json.Marshal(*queryModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5ea78be3-2f5a-4110-8034-c27f24c62db1")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// results (required)
// project (required): Project ID or project name
// runId (required)
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
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// testResultId (required)
// detailsToInclude (optional)
func (client Client) GetTestResultById(ctx context.Context, project *string, runId *int, testResultId *int, detailsToInclude *ResultDetails) (*TestCaseResult, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testResultId"} 
    }
    routeValues["testResultId"] = strconv.Itoa(*testResultId)

    queryParams := url.Values{}
    if detailsToInclude != nil {
        queryParams.Add("detailsToInclude", string(*detailsToInclude))
    }
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestCaseResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// detailsToInclude (optional)
// skip (optional)
// top (optional)
// outcomes (optional)
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
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// results (required)
// project (required): Project ID or project name
// runId (required)
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
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// buildId (required)
// publishContext (optional)
// includeFailureDetails (optional)
// buildToCompare (optional)
func (client Client) QueryTestResultsReportForBuild(ctx context.Context, project *string, buildId *int, publishContext *string, includeFailureDetails *bool, buildToCompare *BuildReference) (*TestResultSummary, error) {
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
    if publishContext != nil {
        queryParams.Add("publishContext", *publishContext)
    }
    if includeFailureDetails != nil {
        queryParams.Add("includeFailureDetails", strconv.FormatBool(*includeFailureDetails))
    }
    if buildToCompare != nil {
        if buildToCompare.Id != nil {
            queryParams.Add("buildToCompare.id", strconv.Itoa(*buildToCompare.Id))
        }
        if buildToCompare.DefinitionId != nil {
            queryParams.Add("buildToCompare.definitionId", strconv.Itoa(*buildToCompare.DefinitionId))
        }
        if buildToCompare.Number != nil {
            queryParams.Add("buildToCompare.number", *buildToCompare.Number)
        }
        if buildToCompare.Uri != nil {
            queryParams.Add("buildToCompare.uri", *buildToCompare.Uri)
        }
        if buildToCompare.BuildSystem != nil {
            queryParams.Add("buildToCompare.buildSystem", *buildToCompare.BuildSystem)
        }
        if buildToCompare.BranchName != nil {
            queryParams.Add("buildToCompare.branchName", *buildToCompare.BranchName)
        }
        if buildToCompare.RepositoryId != nil {
            queryParams.Add("buildToCompare.repositoryId", *buildToCompare.RepositoryId)
        }
    }
    locationId, _ := uuid.Parse("e009fa95-95a5-4ad4-9681-590043ce2423")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// releaseId (required)
// releaseEnvId (required)
// publishContext (optional)
// includeFailureDetails (optional)
// releaseToCompare (optional)
func (client Client) QueryTestResultsReportForRelease(ctx context.Context, project *string, releaseId *int, releaseEnvId *int, publishContext *string, includeFailureDetails *bool, releaseToCompare *ReleaseReference) (*TestResultSummary, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if releaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseId"}
    }
    queryParams.Add("releaseId", strconv.Itoa(*releaseId))
    if releaseEnvId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releaseEnvId"}
    }
    queryParams.Add("releaseEnvId", strconv.Itoa(*releaseEnvId))
    if publishContext != nil {
        queryParams.Add("publishContext", *publishContext)
    }
    if includeFailureDetails != nil {
        queryParams.Add("includeFailureDetails", strconv.FormatBool(*includeFailureDetails))
    }
    if releaseToCompare != nil {
        if releaseToCompare.Id != nil {
            queryParams.Add("releaseToCompare.id", strconv.Itoa(*releaseToCompare.Id))
        }
        if releaseToCompare.Name != nil {
            queryParams.Add("releaseToCompare.name", *releaseToCompare.Name)
        }
        if releaseToCompare.EnvironmentId != nil {
            queryParams.Add("releaseToCompare.environmentId", strconv.Itoa(*releaseToCompare.EnvironmentId))
        }
        if releaseToCompare.EnvironmentName != nil {
            queryParams.Add("releaseToCompare.environmentName", *releaseToCompare.EnvironmentName)
        }
        if releaseToCompare.DefinitionId != nil {
            queryParams.Add("releaseToCompare.definitionId", strconv.Itoa(*releaseToCompare.DefinitionId))
        }
        if releaseToCompare.EnvironmentDefinitionId != nil {
            queryParams.Add("releaseToCompare.environmentDefinitionId", strconv.Itoa(*releaseToCompare.EnvironmentDefinitionId))
        }
        if releaseToCompare.EnvironmentDefinitionName != nil {
            queryParams.Add("releaseToCompare.environmentDefinitionName", *releaseToCompare.EnvironmentDefinitionName)
        }
        if releaseToCompare.CreationDate != nil {
            queryParams.Add("releaseToCompare.creationDate", (*releaseToCompare.CreationDate).String())
        }
        if releaseToCompare.EnvironmentCreationDate != nil {
            queryParams.Add("releaseToCompare.environmentCreationDate", (*releaseToCompare.EnvironmentCreationDate).String())
        }
        if releaseToCompare.Attempt != nil {
            queryParams.Add("releaseToCompare.attempt", strconv.Itoa(*releaseToCompare.Attempt))
        }
    }
    locationId, _ := uuid.Parse("f10f9577-2c04-45ab-8c99-b26567a7cd55")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// releases (required)
// project (required): Project ID or project name
func (client Client) QueryTestResultsSummaryForReleases(ctx context.Context, releases *[]ReleaseReference, project *string) (*[]TestResultSummary, error) {
    if releases == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "releases"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*releases)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f10f9577-2c04-45ab-8c99-b26567a7cd55")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestResultSummary
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// resultsContext (required)
// project (required): Project ID or project name
// workItemIds (optional)
func (client Client) QueryTestSummaryByRequirement(ctx context.Context, resultsContext *TestResultsContext, project *string, workItemIds *[]int) (*[]TestSummaryForWorkItem, error) {
    if resultsContext == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "resultsContext"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if workItemIds != nil {
        var stringList []string
        for _, item := range *workItemIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    body, marshalErr := json.Marshal(*resultsContext)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("3b7fd26f-c335-4e55-afc1-a588f5e2af3c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestSummaryForWorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// filter (required)
// project (required): Project ID or project name
func (client Client) QueryResultTrendForBuild(ctx context.Context, filter *TestResultTrendFilter, project *string) (*[]AggregatedDataForResultTrend, error) {
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
    locationId, _ := uuid.Parse("0886a7ae-315a-4dba-9122-bcce93301f3a")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AggregatedDataForResultTrend
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// filter (required)
// project (required): Project ID or project name
func (client Client) QueryResultTrendForRelease(ctx context.Context, filter *TestResultTrendFilter, project *string) (*[]AggregatedDataForResultTrend, error) {
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
    locationId, _ := uuid.Parse("107f23c3-359a-460a-a70c-63ee739f9f9a")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AggregatedDataForResultTrend
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// testRun (required)
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
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
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

    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// includeDetails (optional)
// includeTags (optional)
func (client Client) GetTestRunById(ctx context.Context, project *string, runId *int, includeDetails *bool, includeTags *bool) (*TestRun, error) {
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
    if includeTags != nil {
        queryParams.Add("includeTags", strconv.FormatBool(*includeTags))
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// buildUri (optional)
// owner (optional)
// tmiRunId (optional)
// planId (optional)
// includeRunDetails (optional)
// automated (optional)
// skip (optional)
// top (optional)
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
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Query Test Runs based on filters. Mandatory fields are minLastUpdatedDate and maxLastUpdatedDate.
// ctx
// project (required): Project ID or project name
// minLastUpdatedDate (required): Minimum Last Modified Date of run to be queried (Mandatory).
// maxLastUpdatedDate (required): Maximum Last Modified Date of run to be queried (Mandatory, difference between min and max date can be atmost 7 days).
// state (optional): Current state of the Runs to be queried.
// planIds (optional): Plan Ids of the Runs to be queried, comma seperated list of valid ids.
// isAutomated (optional): Automation type of the Runs to be queried.
// publishContext (optional): PublishContext of the Runs to be queried.
// buildIds (optional): Build Ids of the Runs to be queried, comma seperated list of valid ids.
// buildDefIds (optional): Build Definition Ids of the Runs to be queried, comma seperated list of valid ids.
// branchName (optional): Source Branch name of the Runs to be queried.
// releaseIds (optional): Release Ids of the Runs to be queried, comma seperated list of valid ids.
// releaseDefIds (optional): Release Definition Ids of the Runs to be queried, comma seperated list of valid ids.
// releaseEnvIds (optional): Release Environment Ids of the Runs to be queried, comma seperated list of valid ids.
// releaseEnvDefIds (optional): Release Environment Definition Ids of the Runs to be queried, comma seperated list of valid ids.
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
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// runUpdateModel (required)
// project (required): Project ID or project name
// runId (required)
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
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
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

    locationId, _ := uuid.Parse("82b986e8-ca9e-4a89-b39e-f65c69bc104a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRunStatistic
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get TestResultsSettings data
// ctx
// project (required): Project ID or project name
// settingsType (optional)
func (client Client) GetTestResultsSettings(ctx context.Context, project *string, settingsType *TestResultsSettingsType) (*TestResultsSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if settingsType != nil {
        queryParams.Add("settingsType", string(*settingsType))
    }
    locationId, _ := uuid.Parse("7319952e-e5a9-4e19-a006-84f3be8b7c68")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.3", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update project settings of test results
// ctx
// testResultsUpdateSettings (required)
// project (required): Project ID or project name
func (client Client) UpdatePipelinesTestSettings(ctx context.Context, testResultsUpdateSettings *TestResultsUpdateSettings, project *string) (*TestResultsSettings, error) {
    if testResultsUpdateSettings == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testResultsUpdateSettings"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testResultsUpdateSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7319952e-e5a9-4e19-a006-84f3be8b7c68")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.3", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
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
    locationId, _ := uuid.Parse("2a41bd6a-8118-4403-b74e-5ba7492aed9d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestHistoryQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// buildId (required)
// type_ (required)
// directoryPath (optional)
// fileNamePrefix (optional)
// fetchMetaData (optional)
// top (optional)
// continuationToken (optional): Header to pass the continuationToken
func (client Client) GetTestLogsForBuild(ctx context.Context, project *string, buildId *int, type_ *TestLogType, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
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
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if directoryPath != nil {
        queryParams.Add("directoryPath", *directoryPath)
    }
    if fileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *fileNamePrefix)
    }
    if fetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*fetchMetaData))
    }
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    additionalHeaders := make(map[string]string)
    if continuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *continuationToken
    }
    locationId, _ := uuid.Parse("dff8ce3a-e539-4817-a405-d968491a88f1")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// type_ (required)
// directoryPath (optional)
// fileNamePrefix (optional)
// fetchMetaData (optional)
// top (optional)
// continuationToken (optional): Header to pass the continuationToken
func (client Client) GetTestResultLogs(ctx context.Context, project *string, runId *int, resultId *int, type_ *TestLogType, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if directoryPath != nil {
        queryParams.Add("directoryPath", *directoryPath)
    }
    if fileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *fileNamePrefix)
    }
    if fetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*fetchMetaData))
    }
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    additionalHeaders := make(map[string]string)
    if continuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *continuationToken
    }
    locationId, _ := uuid.Parse("714caaac-ae1e-4869-8323-9bc0f5120dbf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// subResultId (required)
// type_ (required)
// directoryPath (optional)
// fileNamePrefix (optional)
// fetchMetaData (optional)
// top (optional)
// continuationToken (optional): Header to pass the continuationToken
func (client Client) GetTestSubResultLogs(ctx context.Context, project *string, runId *int, resultId *int, subResultId *int, type_ *TestLogType, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if subResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subResultId"}
    }
    queryParams.Add("subResultId", strconv.Itoa(*subResultId))
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if directoryPath != nil {
        queryParams.Add("directoryPath", *directoryPath)
    }
    if fileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *fileNamePrefix)
    }
    if fetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*fetchMetaData))
    }
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    additionalHeaders := make(map[string]string)
    if continuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *continuationToken
    }
    locationId, _ := uuid.Parse("714caaac-ae1e-4869-8323-9bc0f5120dbf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// type_ (required)
// directoryPath (optional)
// fileNamePrefix (optional)
// fetchMetaData (optional)
// top (optional)
// continuationToken (optional): Header to pass the continuationToken
func (client Client) GetTestRunLogs(ctx context.Context, project *string, runId *int, type_ *TestLogType, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
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
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if directoryPath != nil {
        queryParams.Add("directoryPath", *directoryPath)
    }
    if fileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *fileNamePrefix)
    }
    if fetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*fetchMetaData))
    }
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    additionalHeaders := make(map[string]string)
    if continuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *continuationToken
    }
    locationId, _ := uuid.Parse("5b47b946-e875-4c9a-acdc-2a20996caebe")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// build (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForBuildLog(ctx context.Context, project *string, build *int, type_ *TestLogType, filePath *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if build == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "build"}
    }
    queryParams.Add("build", strconv.Itoa(*build))
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if filePath == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("39b09be7-f0c9-4a83-a513-9ae31b45c56f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// buildId (required)
// testLogStoreOperationType (required)
func (client Client) TestLogStoreEndpointDetailsForBuild(ctx context.Context, project *string, buildId *int, testLogStoreOperationType *TestLogStoreOperationType) (*TestLogStoreEndpointDetails, error) {
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
    if testLogStoreOperationType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testLogStoreOperationType"}
    }
    queryParams.Add("testLogStoreOperationType", string(*testLogStoreOperationType))
    locationId, _ := uuid.Parse("39b09be7-f0c9-4a83-a513-9ae31b45c56f")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForResultLog(ctx context.Context, project *string, runId *int, resultId *int, type_ *TestLogType, filePath *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if filePath == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// subResultId (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForSubResultLog(ctx context.Context, project *string, runId *int, resultId *int, subResultId *int, type_ *TestLogType, filePath *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if subResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subResultId"}
    }
    queryParams.Add("subResultId", strconv.Itoa(*subResultId))
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if filePath == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// subResultId (required)
// filePath (required)
// type_ (required)
func (client Client) TestLogStoreEndpointDetailsForResult(ctx context.Context, project *string, runId *int, resultId *int, subResultId *int, filePath *string, type_ *TestLogType) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if subResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subResultId"}
    }
    queryParams.Add("subResultId", strconv.Itoa(*subResultId))
    if filePath == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *filePath)
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForRunLog(ctx context.Context, project *string, runId *int, type_ *TestLogType, filePath *string) (*TestLogStoreEndpointDetails, error) {
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
    if type_ == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*type_))
    if filePath == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// runId (required)
// testLogStoreOperationType (required)
// filePath (optional)
// type_ (optional)
func (client Client) TestLogStoreEndpointDetailsForRun(ctx context.Context, project *string, runId *int, testLogStoreOperationType *TestLogStoreOperationType, filePath *string, type_ *TestLogType) (*TestLogStoreEndpointDetails, error) {
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
    if testLogStoreOperationType == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testLogStoreOperationType"}
    }
    queryParams.Add("testLogStoreOperationType", string(*testLogStoreOperationType))
    if filePath != nil {
        queryParams.Add("filePath", *filePath)
    }
    if type_ != nil {
        queryParams.Add("type_", string(*type_))
    }
    locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// testSettings (required)
// project (required): Project ID or project name
func (client Client) CreateTestSettings(ctx context.Context, testSettings *TestSettings, project *string) (*int, error) {
    if testSettings == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSettings"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue int
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// testSettingsId (required)
func (client Client) DeleteTestSettings(ctx context.Context, project *string, testSettingsId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testSettingsId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testSettingsId"}
    }
    queryParams.Add("testSettingsId", strconv.Itoa(*testSettingsId))
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// testSettingsId (required)
func (client Client) GetTestSettingsById(ctx context.Context, project *string, testSettingsId *int) (*TestSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testSettingsId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSettingsId"}
    }
    queryParams.Add("testSettingsId", strconv.Itoa(*testSettingsId))
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// workItemToTestLinks (required)
// project (required): Project ID or project name
func (client Client) AddWorkItemToTestLinks(ctx context.Context, workItemToTestLinks *WorkItemToTestLinks, project *string) (*WorkItemToTestLinks, error) {
    if workItemToTestLinks == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemToTestLinks"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*workItemToTestLinks)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4e3abe63-ca46-4fe0-98b2-363f7ec7aa5f")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemToTestLinks
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// testName (required)
// workItemId (required)
func (client Client) DeleteTestMethodToWorkItemLink(ctx context.Context, project *string, testName *string, workItemId *int) (*bool, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testName"}
    }
    queryParams.Add("testName", *testName)
    if workItemId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemId"}
    }
    queryParams.Add("workItemId", strconv.Itoa(*workItemId))
    locationId, _ := uuid.Parse("cbd50bd7-f7ed-4e35-b127-4408ae6bfa2c")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// testName (required)
func (client Client) QueryTestMethodLinkedWorkItems(ctx context.Context, project *string, testName *string) (*TestToWorkItemLinks, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testName == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testName"}
    }
    queryParams.Add("testName", *testName)
    locationId, _ := uuid.Parse("cbd50bd7-f7ed-4e35-b127-4408ae6bfa2c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestToWorkItemLinks
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// project (required): Project ID or project name
// workItemCategory (required)
// automatedTestName (optional)
// testCaseId (optional)
// maxCompleteDate (optional)
// days (optional)
// workItemCount (optional)
func (client Client) QueryTestResultWorkItems(ctx context.Context, project *string, workItemCategory *string, automatedTestName *string, testCaseId *int, maxCompleteDate *time.Time, days *int, workItemCount *int) (*[]WorkItemReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if workItemCategory == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "workItemCategory"}
    }
    queryParams.Add("workItemCategory", *workItemCategory)
    if automatedTestName != nil {
        queryParams.Add("automatedTestName", *automatedTestName)
    }
    if testCaseId != nil {
        queryParams.Add("testCaseId", strconv.Itoa(*testCaseId))
    }
    if maxCompleteDate != nil {
        queryParams.Add("maxCompleteDate", (*maxCompleteDate).String())
    }
    if days != nil {
        queryParams.Add("days", strconv.Itoa(*days))
    }
    if workItemCount != nil {
        queryParams.Add("$workItemCount", strconv.Itoa(*workItemCount))
    }
    locationId, _ := uuid.Parse("f7401a26-331b-44fe-a470-f7ed35138e4a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

