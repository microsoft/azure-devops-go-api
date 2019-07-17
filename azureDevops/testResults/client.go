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

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API]
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// iterationId (required)
// actionPath (optional)
func (client Client) CreateTestIterationResultAttachment(attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int, testCaseResultId *int, iterationId *int, actionPath *string) (*TestAttachmentReference, error) {
    if attachmentRequestModel == nil {
        return nil, errors.New("attachmentRequestModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    queryParams := url.Values{}
    if iterationId == nil {
        return nil, errors.New("iterationId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
func (client Client) CreateTestResultAttachment(attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int, testCaseResultId *int) (*TestAttachmentReference, error) {
    if attachmentRequestModel == nil {
        return nil, errors.New("attachmentRequestModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    body, marshalErr := json.Marshal(*attachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// testSubResultId (required)
func (client Client) CreateTestSubResultAttachment(attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int, testCaseResultId *int, testSubResultId *int) (*TestAttachmentReference, error) {
    if attachmentRequestModel == nil {
        return nil, errors.New("attachmentRequestModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, errors.New("testSubResultId is a required parameter")
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    body, marshalErr := json.Marshal(*attachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
func (client Client) DeleteTestResultAttachment(project *string, runId *int, testCaseResultId *int, attachmentId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a test result attachment
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
func (client Client) GetTestResultAttachmentContent(project *string, runId *int, testCaseResultId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
func (client Client) GetTestResultAttachments(project *string, runId *int, testCaseResultId *int) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a test result attachment
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
func (client Client) GetTestResultAttachmentZip(project *string, runId *int, testCaseResultId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a test sub result attachment
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
// testSubResultId (required)
func (client Client) GetTestSubResultAttachmentContent(project *string, runId *int, testCaseResultId *int, attachmentId *int, testSubResultId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, errors.New("testSubResultId is a required parameter")
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns attachment references for test sub result.
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// testSubResultId (required)
func (client Client) GetTestSubResultAttachments(project *string, runId *int, testCaseResultId *int, testSubResultId *int) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, errors.New("testSubResultId is a required parameter")
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a test sub result attachment
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
// attachmentId (required)
// testSubResultId (required)
func (client Client) GetTestSubResultAttachmentZip(project *string, runId *int, testCaseResultId *int, attachmentId *int, testSubResultId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)
    if attachmentId == nil {
        return nil, errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    queryParams := url.Values{}
    if testSubResultId == nil {
        return nil, errors.New("testSubResultId is a required parameter")
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*testSubResultId))
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// attachmentRequestModel (required)
// project (required): Project ID or project name
// runId (required)
func (client Client) CreateTestRunAttachment(attachmentRequestModel *TestAttachmentRequestModel, project *string, runId *int) (*TestAttachmentReference, error) {
    if attachmentRequestModel == nil {
        return nil, errors.New("attachmentRequestModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*attachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// attachmentId (required)
func (client Client) DeleteTestRunAttachment(project *string, runId *int, attachmentId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if attachmentId == nil {
        return errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Returns a test run attachment
// project (required): Project ID or project name
// runId (required)
// attachmentId (required)
func (client Client) GetTestRunAttachmentContent(project *string, runId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if attachmentId == nil {
        return nil, errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
func (client Client) GetTestRunAttachments(project *string, runId *int) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a test run attachment
// project (required): Project ID or project name
// runId (required)
// attachmentId (required)
func (client Client) GetTestRunAttachmentZip(project *string, runId *int, attachmentId *int) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if attachmentId == nil {
        return nil, errors.New("attachmentId is a required parameter")
    }
    routeValues["attachmentId"] = strconv.Itoa(*attachmentId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// testCaseResultId (required)
func (client Client) GetBugsLinkedToTestResult(project *string, runId *int, testCaseResultId *int) (*[]WorkItemReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testCaseResultId == nil {
        return nil, errors.New("testCaseResultId is a required parameter")
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*testCaseResultId)

    locationId, _ := uuid.Parse("d8dbf98f-eb34-4f8d-8365-47972af34f29")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// buildId (required)
// flags (required)
func (client Client) GetBuildCodeCoverage(project *string, buildId *int, flags *int) (*[]BuildCoverage, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    queryParams.Add("buildId", strconv.Itoa(*buildId))
    if flags == nil {
        return nil, errors.New("flags is a required parameter")
    }
    queryParams.Add("flags", strconv.Itoa(*flags))
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// buildId (required)
// deltaBuildId (optional)
func (client Client) GetCodeCoverageSummary(project *string, buildId *int, deltaBuildId *int) (*CodeCoverageSummary, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    queryParams.Add("buildId", strconv.Itoa(*buildId))
    if deltaBuildId != nil {
        queryParams.Add("deltaBuildId", strconv.Itoa(*deltaBuildId))
    }
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CodeCoverageSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] http://(tfsserver):8080/tfs/DefaultCollection/_apis/test/CodeCoverage?buildId=10 Request: Json of code coverage summary
// project (required): Project ID or project name
// buildId (required)
// coverageData (optional)
func (client Client) UpdateCodeCoverageSummary(project *string, buildId *int, coverageData *CodeCoverageData) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return errors.New("buildId is a required parameter")
    }
    queryParams.Add("buildId", strconv.Itoa(*buildId))
    body, marshalErr := json.Marshal(*coverageData)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    _, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// flags (required)
func (client Client) GetTestRunCodeCoverage(project *string, runId *int, flags *int) (*[]TestRunCoverage, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if flags == nil {
        return nil, errors.New("flags is a required parameter")
    }
    queryParams.Add("flags", strconv.Itoa(*flags))
    locationId, _ := uuid.Parse("5641efbc-6f9b-401a-baeb-d3da22489e5e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRunCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// filter (required)
// project (required): Project ID or project name
func (client Client) QueryTestResultHistory(filter *ResultsFilter, project *string) (*TestResultHistory, error) {
    if filter == nil {
        return nil, errors.New("filter is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdf7a97b-0395-4da8-9d5d-f957619327d1")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultHistory
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// buildId (required)
// publishContext (optional)
// groupBy (optional)
// filter (optional)
// orderby (optional)
// shouldIncludeResults (optional)
// queryRunSummaryForInProgress (optional)
func (client Client) GetTestResultDetailsForBuild(project *string, buildId *int, publishContext *string, groupBy *string, filter *string, orderby *string, shouldIncludeResults *bool, queryRunSummaryForInProgress *bool) (*TestResultsDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// releaseId (required)
// releaseEnvId (required)
// publishContext (optional)
// groupBy (optional)
// filter (optional)
// orderby (optional)
// shouldIncludeResults (optional)
// queryRunSummaryForInProgress (optional)
func (client Client) GetTestResultDetailsForRelease(project *string, releaseId *int, releaseEnvId *int, publishContext *string, groupBy *string, filter *string, orderby *string, shouldIncludeResults *bool, queryRunSummaryForInProgress *bool) (*TestResultsDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if releaseId == nil {
        return nil, errors.New("releaseId is a required parameter")
    }
    queryParams.Add("releaseId", strconv.Itoa(*releaseId))
    if releaseEnvId == nil {
        return nil, errors.New("releaseEnvId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// query (required)
// project (required): Project ID or project name
func (client Client) GetTestResultsByQuery(query *TestResultsQuery, project *string) (*TestResultsQuery, error) {
    if query == nil {
        return nil, errors.New("query is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("14033a2c-af25-4af1-9e39-8ef6900482e3")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// queryModel (required)
// project (required): Project ID or project name
// includeResultDetails (optional)
// includeIterationDetails (optional)
// skip (optional)
// top (optional)
func (client Client) GetTestResultsByQueryWiql(queryModel *QueryModel, project *string, includeResultDetails *bool, includeIterationDetails *bool, skip *int, top *int) (*[]TestCaseResult, error) {
    if queryModel == nil {
        return nil, errors.New("queryModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
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
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// results (required)
// project (required): Project ID or project name
// runId (required)
func (client Client) AddTestResultsToTestRun(results *[]TestCaseResult, project *string, runId *int) (*[]TestCaseResult, error) {
    if results == nil {
        return nil, errors.New("results is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// testResultId (required)
// detailsToInclude (optional)
func (client Client) GetTestResultById(project *string, runId *int, testResultId *int, detailsToInclude *string) (*TestCaseResult, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if testResultId == nil {
        return nil, errors.New("testResultId is a required parameter")
    }
    routeValues["testResultId"] = strconv.Itoa(*testResultId)

    queryParams := url.Values{}
    if detailsToInclude != nil {
        queryParams.Add("detailsToInclude", *detailsToInclude)
    }
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestCaseResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// detailsToInclude (optional)
// skip (optional)
// top (optional)
// outcomes (optional)
func (client Client) GetTestResults(project *string, runId *int, detailsToInclude *string, skip *int, top *int, outcomes *[]TestOutcome) (*[]TestCaseResult, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if detailsToInclude != nil {
        queryParams.Add("detailsToInclude", *detailsToInclude)
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// results (required)
// project (required): Project ID or project name
// runId (required)
func (client Client) UpdateTestResults(results *[]TestCaseResult, project *string, runId *int) (*[]TestCaseResult, error) {
    if results == nil {
        return nil, errors.New("results is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// buildId (required)
// publishContext (optional)
// includeFailureDetails (optional)
// buildToCompare (optional)
func (client Client) QueryTestResultsReportForBuild(project *string, buildId *int, publishContext *string, includeFailureDetails *bool, buildToCompare *BuildReference) (*TestResultSummary, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// releaseId (required)
// releaseEnvId (required)
// publishContext (optional)
// includeFailureDetails (optional)
// releaseToCompare (optional)
func (client Client) QueryTestResultsReportForRelease(project *string, releaseId *int, releaseEnvId *int, publishContext *string, includeFailureDetails *bool, releaseToCompare *ReleaseReference) (*TestResultSummary, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if releaseId == nil {
        return nil, errors.New("releaseId is a required parameter")
    }
    queryParams.Add("releaseId", strconv.Itoa(*releaseId))
    if releaseEnvId == nil {
        return nil, errors.New("releaseEnvId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// releases (required)
// project (required): Project ID or project name
func (client Client) QueryTestResultsSummaryForReleases(releases *[]ReleaseReference, project *string) (*[]TestResultSummary, error) {
    if releases == nil {
        return nil, errors.New("releases is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*releases)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f10f9577-2c04-45ab-8c99-b26567a7cd55")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestResultSummary
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// resultsContext (required)
// project (required): Project ID or project name
// workItemIds (optional)
func (client Client) QueryTestSummaryByRequirement(resultsContext *TestResultsContext, project *string, workItemIds *[]int) (*[]TestSummaryForWorkItem, error) {
    if resultsContext == nil {
        return nil, errors.New("resultsContext is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
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
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestSummaryForWorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// filter (required)
// project (required): Project ID or project name
func (client Client) QueryResultTrendForBuild(filter *TestResultTrendFilter, project *string) (*[]AggregatedDataForResultTrend, error) {
    if filter == nil {
        return nil, errors.New("filter is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0886a7ae-315a-4dba-9122-bcce93301f3a")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AggregatedDataForResultTrend
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// filter (required)
// project (required): Project ID or project name
func (client Client) QueryResultTrendForRelease(filter *TestResultTrendFilter, project *string) (*[]AggregatedDataForResultTrend, error) {
    if filter == nil {
        return nil, errors.New("filter is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("107f23c3-359a-460a-a70c-63ee739f9f9a")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AggregatedDataForResultTrend
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// testRun (required)
// project (required): Project ID or project name
func (client Client) CreateTestRun(testRun *RunCreateModel, project *string) (*TestRun, error) {
    if testRun == nil {
        return nil, errors.New("testRun is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testRun)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
func (client Client) DeleteTestRun(project *string, runId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// includeDetails (optional)
// includeTags (optional)
func (client Client) GetTestRunById(project *string, runId *int, includeDetails *bool, includeTags *bool) (*TestRun, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// buildUri (optional)
// owner (optional)
// tmiRunId (optional)
// planId (optional)
// includeRunDetails (optional)
// automated (optional)
// skip (optional)
// top (optional)
func (client Client) GetTestRuns(project *string, buildUri *string, owner *string, tmiRunId *string, planId *int, includeRunDetails *bool, automated *bool, skip *int, top *int) (*[]TestRun, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Query Test Runs based on filters. Mandatory fields are minLastUpdatedDate and maxLastUpdatedDate.
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
func (client Client) QueryTestRuns(project *string, minLastUpdatedDate *time.Time, maxLastUpdatedDate *time.Time, state *string, planIds *[]int, isAutomated *bool, publishContext *string, buildIds *[]int, buildDefIds *[]int, branchName *string, releaseIds *[]int, releaseDefIds *[]int, releaseEnvIds *[]int, releaseEnvDefIds *[]int, runTitle *string, top *int, continuationToken *string) (*[]TestRun, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if minLastUpdatedDate == nil {
        return nil, errors.New("minLastUpdatedDate is a required parameter")
    }
    queryParams.Add("minLastUpdatedDate", (*minLastUpdatedDate).String())
    if maxLastUpdatedDate == nil {
        return nil, errors.New("maxLastUpdatedDate is a required parameter")
    }
    queryParams.Add("maxLastUpdatedDate", (*maxLastUpdatedDate).String())
    if state != nil {
        queryParams.Add("state", *state)
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
        queryParams.Add("publishContext", *publishContext)
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// runUpdateModel (required)
// project (required): Project ID or project name
// runId (required)
func (client Client) UpdateTestRun(runUpdateModel *RunUpdateModel, project *string, runId *int) (*TestRun, error) {
    if runUpdateModel == nil {
        return nil, errors.New("runUpdateModel is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    body, marshalErr := json.Marshal(*runUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get test run statistics , used when we want to get summary of a run by outcome.
// project (required): Project ID or project name
// runId (required): ID of the run to get.
func (client Client) GetTestRunStatistics(project *string, runId *int) (*TestRunStatistic, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    locationId, _ := uuid.Parse("82b986e8-ca9e-4a89-b39e-f65c69bc104a")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRunStatistic
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get TestResultsSettings data
// project (required): Project ID or project name
// settingsType (optional)
func (client Client) GetTestResultsSettings(project *string, settingsType *string) (*TestResultsSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if settingsType != nil {
        queryParams.Add("settingsType", *settingsType)
    }
    locationId, _ := uuid.Parse("7319952e-e5a9-4e19-a006-84f3be8b7c68")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.3", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update project settings of test results
// testResultsUpdateSettings (required)
// project (required): Project ID or project name
func (client Client) UpdatePipelinesTestSettings(testResultsUpdateSettings *TestResultsUpdateSettings, project *string) (*TestResultsSettings, error) {
    if testResultsUpdateSettings == nil {
        return nil, errors.New("testResultsUpdateSettings is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testResultsUpdateSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7319952e-e5a9-4e19-a006-84f3be8b7c68")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.3", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResultsSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get history of a test method using TestHistoryQuery
// filter (required): TestHistoryQuery to get history
// project (required): Project ID or project name
func (client Client) QueryTestHistory(filter *TestHistoryQuery, project *string) (*TestHistoryQuery, error) {
    if filter == nil {
        return nil, errors.New("filter is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a41bd6a-8118-4403-b74e-5ba7492aed9d")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestHistoryQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// buildId (required)
// type_ (required)
// directoryPath (optional)
// fileNamePrefix (optional)
// fetchMetaData (optional)
// top (optional)
// continuationToken (optional): Header to pass the continuationToken
func (client Client) GetTestLogsForBuild(project *string, buildId *int, type_ *string, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    queryParams.Add("buildId", strconv.Itoa(*buildId))
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// type_ (required)
// directoryPath (optional)
// fileNamePrefix (optional)
// fetchMetaData (optional)
// top (optional)
// continuationToken (optional): Header to pass the continuationToken
func (client Client) GetTestResultLogs(project *string, runId *int, resultId *int, type_ *string, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, errors.New("resultId is a required parameter")
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
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
func (client Client) GetTestSubResultLogs(project *string, runId *int, resultId *int, subResultId *int, type_ *string, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, errors.New("resultId is a required parameter")
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if subResultId == nil {
        return nil, errors.New("subResultId is a required parameter")
    }
    queryParams.Add("subResultId", strconv.Itoa(*subResultId))
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// type_ (required)
// directoryPath (optional)
// fileNamePrefix (optional)
// fetchMetaData (optional)
// top (optional)
// continuationToken (optional): Header to pass the continuationToken
func (client Client) GetTestRunLogs(project *string, runId *int, type_ *string, directoryPath *string, fileNamePrefix *string, fetchMetaData *bool, top *int, continuationToken *string) (*[]TestLog, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// build (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForBuildLog(project *string, build *int, type_ *string, filePath *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if build == nil {
        return nil, errors.New("build is a required parameter")
    }
    queryParams.Add("build", strconv.Itoa(*build))
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
    if filePath == nil {
        return nil, errors.New("filePath is a required parameter")
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("39b09be7-f0c9-4a83-a513-9ae31b45c56f")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// buildId (required)
// testLogStoreOperationType (required)
func (client Client) TestLogStoreEndpointDetailsForBuild(project *string, buildId *int, testLogStoreOperationType *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if buildId == nil {
        return nil, errors.New("buildId is a required parameter")
    }
    queryParams.Add("buildId", strconv.Itoa(*buildId))
    if testLogStoreOperationType == nil {
        return nil, errors.New("testLogStoreOperationType is a required parameter")
    }
    queryParams.Add("testLogStoreOperationType", *testLogStoreOperationType)
    locationId, _ := uuid.Parse("39b09be7-f0c9-4a83-a513-9ae31b45c56f")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForResultLog(project *string, runId *int, resultId *int, type_ *string, filePath *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, errors.New("resultId is a required parameter")
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
    if filePath == nil {
        return nil, errors.New("filePath is a required parameter")
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// subResultId (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForSubResultLog(project *string, runId *int, resultId *int, subResultId *int, type_ *string, filePath *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, errors.New("resultId is a required parameter")
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if subResultId == nil {
        return nil, errors.New("subResultId is a required parameter")
    }
    queryParams.Add("subResultId", strconv.Itoa(*subResultId))
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
    if filePath == nil {
        return nil, errors.New("filePath is a required parameter")
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// resultId (required)
// subResultId (required)
// filePath (required)
// type_ (required)
func (client Client) TestLogStoreEndpointDetailsForResult(project *string, runId *int, resultId *int, subResultId *int, filePath *string, type_ *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)
    if resultId == nil {
        return nil, errors.New("resultId is a required parameter")
    }
    routeValues["resultId"] = strconv.Itoa(*resultId)

    queryParams := url.Values{}
    if subResultId == nil {
        return nil, errors.New("subResultId is a required parameter")
    }
    queryParams.Add("subResultId", strconv.Itoa(*subResultId))
    if filePath == nil {
        return nil, errors.New("filePath is a required parameter")
    }
    queryParams.Add("filePath", *filePath)
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// type_ (required)
// filePath (required)
func (client Client) GetTestLogStoreEndpointDetailsForRunLog(project *string, runId *int, type_ *string, filePath *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if type_ == nil {
        return nil, errors.New("type_ is a required parameter")
    }
    queryParams.Add("type_", *type_)
    if filePath == nil {
        return nil, errors.New("filePath is a required parameter")
    }
    queryParams.Add("filePath", *filePath)
    locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// runId (required)
// testLogStoreOperationType (required)
// filePath (optional)
// type_ (optional)
func (client Client) TestLogStoreEndpointDetailsForRun(project *string, runId *int, testLogStoreOperationType *string, filePath *string, type_ *string) (*TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if runId == nil {
        return nil, errors.New("runId is a required parameter")
    }
    routeValues["runId"] = strconv.Itoa(*runId)

    queryParams := url.Values{}
    if testLogStoreOperationType == nil {
        return nil, errors.New("testLogStoreOperationType is a required parameter")
    }
    queryParams.Add("testLogStoreOperationType", *testLogStoreOperationType)
    if filePath != nil {
        queryParams.Add("filePath", *filePath)
    }
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// testSettings (required)
// project (required): Project ID or project name
func (client Client) CreateTestSettings(testSettings *TestSettings, project *string) (*int, error) {
    if testSettings == nil {
        return nil, errors.New("testSettings is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue int
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// testSettingsId (required)
func (client Client) DeleteTestSettings(project *string, testSettingsId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testSettingsId == nil {
        return errors.New("testSettingsId is a required parameter")
    }
    queryParams.Add("testSettingsId", strconv.Itoa(*testSettingsId))
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// project (required): Project ID or project name
// testSettingsId (required)
func (client Client) GetTestSettingsById(project *string, testSettingsId *int) (*TestSettings, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testSettingsId == nil {
        return nil, errors.New("testSettingsId is a required parameter")
    }
    queryParams.Add("testSettingsId", strconv.Itoa(*testSettingsId))
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// workItemToTestLinks (required)
// project (required): Project ID or project name
func (client Client) AddWorkItemToTestLinks(workItemToTestLinks *WorkItemToTestLinks, project *string) (*WorkItemToTestLinks, error) {
    if workItemToTestLinks == nil {
        return nil, errors.New("workItemToTestLinks is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*workItemToTestLinks)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4e3abe63-ca46-4fe0-98b2-363f7ec7aa5f")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WorkItemToTestLinks
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// testName (required)
// workItemId (required)
func (client Client) DeleteTestMethodToWorkItemLink(project *string, testName *string, workItemId *int) (*bool, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testName == nil {
        return nil, errors.New("testName is a required parameter")
    }
    queryParams.Add("testName", *testName)
    if workItemId == nil {
        return nil, errors.New("workItemId is a required parameter")
    }
    queryParams.Add("workItemId", strconv.Itoa(*workItemId))
    locationId, _ := uuid.Parse("cbd50bd7-f7ed-4e35-b127-4408ae6bfa2c")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// testName (required)
func (client Client) QueryTestMethodLinkedWorkItems(project *string, testName *string) (*TestToWorkItemLinks, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testName == nil {
        return nil, errors.New("testName is a required parameter")
    }
    queryParams.Add("testName", *testName)
    locationId, _ := uuid.Parse("cbd50bd7-f7ed-4e35-b127-4408ae6bfa2c")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestToWorkItemLinks
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// project (required): Project ID or project name
// workItemCategory (required)
// automatedTestName (optional)
// testCaseId (optional)
// maxCompleteDate (optional)
// days (optional)
// workItemCount (optional)
func (client Client) QueryTestResultWorkItems(project *string, workItemCategory *string, automatedTestName *string, testCaseId *int, maxCompleteDate *time.Time, days *int, workItemCount *int) (*[]WorkItemReference, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if workItemCategory == nil {
        return nil, errors.New("workItemCategory is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []WorkItemReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

