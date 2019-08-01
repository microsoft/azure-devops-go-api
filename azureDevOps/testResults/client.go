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
    "github.com/microsoft/azure-devops-go-api/azureDevOps"
    "github.com/microsoft/azure-devops-go-api/azureDevOps/test"
    "io"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

type Client struct {
    Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API]
func (client *Client) CreateTestIterationResultAttachment(ctx context.Context, args CreateTestIterationResultAttachmentArgs) (*test.TestAttachmentReference, error) {
    if args.AttachmentRequestModel == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    queryParams := url.Values{}
    if args.IterationId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "iterationId"}
    }
    queryParams.Add("iterationId", strconv.Itoa(*args.IterationId))
    if args.ActionPath != nil {
        queryParams.Add("actionPath", *args.ActionPath)
    }
    body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestIterationResultAttachment function
type CreateTestIterationResultAttachmentArgs struct {
    // (required)
    AttachmentRequestModel *test.TestAttachmentRequestModel
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    IterationId *int
    // (optional)
    ActionPath *string
}

// [Preview API]
func (client *Client) CreateTestResultAttachment(ctx context.Context, args CreateTestResultAttachmentArgs) (*test.TestAttachmentReference, error) {
    if args.AttachmentRequestModel == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestResultAttachment function
type CreateTestResultAttachmentArgs struct {
    // (required)
    AttachmentRequestModel *test.TestAttachmentRequestModel
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
}

// [Preview API]
func (client *Client) CreateTestSubResultAttachment(ctx context.Context, args CreateTestSubResultAttachmentArgs) (*test.TestAttachmentReference, error) {
    if args.AttachmentRequestModel == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestSubResultAttachment function
type CreateTestSubResultAttachmentArgs struct {
    // (required)
    AttachmentRequestModel *test.TestAttachmentRequestModel
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    TestSubResultId *int
}

// [Preview API]
func (client *Client) DeleteTestResultAttachment(ctx context.Context, args DeleteTestResultAttachmentArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteTestResultAttachment function
type DeleteTestResultAttachmentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    AttachmentId *int
}

// [Preview API] Returns a test result attachment
func (client *Client) GetTestResultAttachmentContent(ctx context.Context, args GetTestResultAttachmentContentArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestResultAttachmentContent function
type GetTestResultAttachmentContentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    AttachmentId *int
}

// [Preview API]
func (client *Client) GetTestResultAttachments(ctx context.Context, args GetTestResultAttachmentsArgs) (*[]test.TestAttachment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultAttachments function
type GetTestResultAttachmentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
}

// [Preview API] Returns a test result attachment
func (client *Client) GetTestResultAttachmentZip(ctx context.Context, args GetTestResultAttachmentZipArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestResultAttachmentZip function
type GetTestResultAttachmentZipArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    AttachmentId *int
}

// [Preview API] Returns a test sub result attachment
func (client *Client) GetTestSubResultAttachmentContent(ctx context.Context, args GetTestSubResultAttachmentContentArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestSubResultAttachmentContent function
type GetTestSubResultAttachmentContentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    AttachmentId *int
    // (required)
    TestSubResultId *int
}

// [Preview API] Returns attachment references for test sub result.
func (client *Client) GetTestSubResultAttachments(ctx context.Context, args GetTestSubResultAttachmentsArgs) (*[]test.TestAttachment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestSubResultAttachments function
type GetTestSubResultAttachmentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    TestSubResultId *int
}

// [Preview API] Returns a test sub result attachment
func (client *Client) GetTestSubResultAttachmentZip(ctx context.Context, args GetTestSubResultAttachmentZipArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    locationId, _ := uuid.Parse("2a632e97-e014-4275-978f-8e5c4906d4b3")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestSubResultAttachmentZip function
type GetTestSubResultAttachmentZipArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
    // (required)
    AttachmentId *int
    // (required)
    TestSubResultId *int
}

// [Preview API]
func (client *Client) CreateTestRunAttachment(ctx context.Context, args CreateTestRunAttachmentArgs) (*test.TestAttachmentReference, error) {
    if args.AttachmentRequestModel == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestRunAttachment function
type CreateTestRunAttachmentArgs struct {
    // (required)
    AttachmentRequestModel *test.TestAttachmentRequestModel
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
}

// [Preview API]
func (client *Client) DeleteTestRunAttachment(ctx context.Context, args DeleteTestRunAttachmentArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.AttachmentId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteTestRunAttachment function
type DeleteTestRunAttachmentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    AttachmentId *int
}

// [Preview API] Returns a test run attachment
func (client *Client) GetTestRunAttachmentContent(ctx context.Context, args GetTestRunAttachmentContentArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.AttachmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestRunAttachmentContent function
type GetTestRunAttachmentContentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    AttachmentId *int
}

// [Preview API]
func (client *Client) GetTestRunAttachments(ctx context.Context, args GetTestRunAttachmentsArgs) (*[]test.TestAttachment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunAttachments function
type GetTestRunAttachmentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
}

// [Preview API] Returns a test run attachment
func (client *Client) GetTestRunAttachmentZip(ctx context.Context, args GetTestRunAttachmentZipArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.AttachmentId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("b5731898-8206-477a-a51d-3fdf116fc6bf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestRunAttachmentZip function
type GetTestRunAttachmentZipArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    AttachmentId *int
}

// [Preview API]
func (client *Client) GetBugsLinkedToTestResult(ctx context.Context, args GetBugsLinkedToTestResultArgs) (*[]test.WorkItemReference, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    locationId, _ := uuid.Parse("d8dbf98f-eb34-4f8d-8365-47972af34f29")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.WorkItemReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetBugsLinkedToTestResult function
type GetBugsLinkedToTestResultArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestCaseResultId *int
}

// [Preview API]
func (client *Client) GetBuildCodeCoverage(ctx context.Context, args GetBuildCodeCoverageArgs) (*[]test.BuildCoverage, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    if args.Flags == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "flags"}
    }
    queryParams.Add("flags", strconv.Itoa(*args.Flags))
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.BuildCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetBuildCodeCoverage function
type GetBuildCodeCoverageArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    BuildId *int
    // (required)
    Flags *int
}

// [Preview API]
func (client *Client) GetCodeCoverageSummary(ctx context.Context, args GetCodeCoverageSummaryArgs) (*test.CodeCoverageSummary, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    if args.DeltaBuildId != nil {
        queryParams.Add("deltaBuildId", strconv.Itoa(*args.DeltaBuildId))
    }
    locationId, _ := uuid.Parse("9b3e1ece-c6ab-4fbb-8167-8a32a0c92216")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.CodeCoverageSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetCodeCoverageSummary function
type GetCodeCoverageSummaryArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    BuildId *int
    // (optional)
    DeltaBuildId *int
}

// [Preview API] http://(tfsserver):8080/tfs/DefaultCollection/_apis/test/CodeCoverage?buildId=10 Request: Json of code coverage summary
func (client *Client) UpdateCodeCoverageSummary(ctx context.Context, args UpdateCodeCoverageSummaryArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    body, marshalErr := json.Marshal(*args.CoverageData)
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

// Arguments for the UpdateCodeCoverageSummary function
type UpdateCodeCoverageSummaryArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    BuildId *int
    // (optional)
    CoverageData *test.CodeCoverageData
}

// [Preview API]
func (client *Client) GetTestRunCodeCoverage(ctx context.Context, args GetTestRunCodeCoverageArgs) (*[]test.TestRunCoverage, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.Flags == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "flags"}
    }
    queryParams.Add("flags", strconv.Itoa(*args.Flags))
    locationId, _ := uuid.Parse("5641efbc-6f9b-401a-baeb-d3da22489e5e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestRunCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunCodeCoverage function
type GetTestRunCodeCoverageArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    Flags *int
}

// [Preview API]
func (client *Client) QueryTestResultHistory(ctx context.Context, args QueryTestResultHistoryArgs) (*test.TestResultHistory, error) {
    if args.Filter == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filter"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdf7a97b-0395-4da8-9d5d-f957619327d1")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultHistory
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestResultHistory function
type QueryTestResultHistoryArgs struct {
    // (required)
    Filter *test.ResultsFilter
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) GetTestResultDetailsForBuild(ctx context.Context, args GetTestResultDetailsForBuildArgs) (*test.TestResultsDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    if args.PublishContext != nil {
        queryParams.Add("publishContext", *args.PublishContext)
    }
    if args.GroupBy != nil {
        queryParams.Add("groupBy", *args.GroupBy)
    }
    if args.Filter != nil {
        queryParams.Add("$filter", *args.Filter)
    }
    if args.Orderby != nil {
        queryParams.Add("$orderby", *args.Orderby)
    }
    if args.ShouldIncludeResults != nil {
        queryParams.Add("shouldIncludeResults", strconv.FormatBool(*args.ShouldIncludeResults))
    }
    if args.QueryRunSummaryForInProgress != nil {
        queryParams.Add("queryRunSummaryForInProgress", strconv.FormatBool(*args.QueryRunSummaryForInProgress))
    }
    locationId, _ := uuid.Parse("a518c749-4524-45b2-a7ef-1ac009b312cd")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultsDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultDetailsForBuild function
type GetTestResultDetailsForBuildArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    BuildId *int
    // (optional)
    PublishContext *string
    // (optional)
    GroupBy *string
    // (optional)
    Filter *string
    // (optional)
    Orderby *string
    // (optional)
    ShouldIncludeResults *bool
    // (optional)
    QueryRunSummaryForInProgress *bool
}

// [Preview API]
func (client *Client) GetTestResultDetailsForRelease(ctx context.Context, args GetTestResultDetailsForReleaseArgs) (*test.TestResultsDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"}
    }
    queryParams.Add("releaseId", strconv.Itoa(*args.ReleaseId))
    if args.ReleaseEnvId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseEnvId"}
    }
    queryParams.Add("releaseEnvId", strconv.Itoa(*args.ReleaseEnvId))
    if args.PublishContext != nil {
        queryParams.Add("publishContext", *args.PublishContext)
    }
    if args.GroupBy != nil {
        queryParams.Add("groupBy", *args.GroupBy)
    }
    if args.Filter != nil {
        queryParams.Add("$filter", *args.Filter)
    }
    if args.Orderby != nil {
        queryParams.Add("$orderby", *args.Orderby)
    }
    if args.ShouldIncludeResults != nil {
        queryParams.Add("shouldIncludeResults", strconv.FormatBool(*args.ShouldIncludeResults))
    }
    if args.QueryRunSummaryForInProgress != nil {
        queryParams.Add("queryRunSummaryForInProgress", strconv.FormatBool(*args.QueryRunSummaryForInProgress))
    }
    locationId, _ := uuid.Parse("19a8183a-69fb-47d7-bfbf-1b6b0d921294")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultsDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultDetailsForRelease function
type GetTestResultDetailsForReleaseArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    ReleaseId *int
    // (required)
    ReleaseEnvId *int
    // (optional)
    PublishContext *string
    // (optional)
    GroupBy *string
    // (optional)
    Filter *string
    // (optional)
    Orderby *string
    // (optional)
    ShouldIncludeResults *bool
    // (optional)
    QueryRunSummaryForInProgress *bool
}

// [Preview API]
func (client *Client) GetTestResultsByQuery(ctx context.Context, args GetTestResultsByQueryArgs) (*test.TestResultsQuery, error) {
    if args.Query == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "query"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("14033a2c-af25-4af1-9e39-8ef6900482e3")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultsQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultsByQuery function
type GetTestResultsByQueryArgs struct {
    // (required)
    Query *test.TestResultsQuery
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) GetTestResultsByQueryWiql(ctx context.Context, args GetTestResultsByQueryWiqlArgs) (*[]test.TestCaseResult, error) {
    if args.QueryModel == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "queryModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.IncludeResultDetails != nil {
        queryParams.Add("includeResultDetails", strconv.FormatBool(*args.IncludeResultDetails))
    }
    if args.IncludeIterationDetails != nil {
        queryParams.Add("includeIterationDetails", strconv.FormatBool(*args.IncludeIterationDetails))
    }
    if args.Skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    body, marshalErr := json.Marshal(*args.QueryModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("5ea78be3-2f5a-4110-8034-c27f24c62db1")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultsByQueryWiql function
type GetTestResultsByQueryWiqlArgs struct {
    // (required)
    QueryModel *test.QueryModel
    // (required) Project ID or project name
    Project *string
    // (optional)
    IncludeResultDetails *bool
    // (optional)
    IncludeIterationDetails *bool
    // (optional)
    Skip *int
    // (optional)
    Top *int
}

// [Preview API]
func (client *Client) AddTestResultsToTestRun(ctx context.Context, args AddTestResultsToTestRunArgs) (*[]test.TestCaseResult, error) {
    if args.Results == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "results"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.Results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the AddTestResultsToTestRun function
type AddTestResultsToTestRunArgs struct {
    // (required)
    Results *[]test.TestCaseResult
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
}

// [Preview API]
func (client *Client) GetTestResultById(ctx context.Context, args GetTestResultByIdArgs) (*test.TestCaseResult, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testResultId"} 
    }
    routeValues["testResultId"] = strconv.Itoa(*args.TestResultId)

    queryParams := url.Values{}
    if args.DetailsToInclude != nil {
        queryParams.Add("detailsToInclude", string(*args.DetailsToInclude))
    }
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestCaseResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultById function
type GetTestResultByIdArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestResultId *int
    // (optional)
    DetailsToInclude *test.ResultDetails
}

// [Preview API]
func (client *Client) GetTestResults(ctx context.Context, args GetTestResultsArgs) (*[]test.TestCaseResult, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.DetailsToInclude != nil {
        queryParams.Add("detailsToInclude", string(*args.DetailsToInclude))
    }
    if args.Skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    if args.Outcomes != nil {
        var stringList []string
        for _, item := range *args.Outcomes {
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

    var responseValue []test.TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResults function
type GetTestResultsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (optional)
    DetailsToInclude *test.ResultDetails
    // (optional)
    Skip *int
    // (optional)
    Top *int
    // (optional)
    Outcomes *[]test.TestOutcome
}

// [Preview API]
func (client *Client) UpdateTestResults(ctx context.Context, args UpdateTestResultsArgs) (*[]test.TestCaseResult, error) {
    if args.Results == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "results"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.Results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("02afa165-e79a-4d70-8f0c-2af0f35b4e07")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateTestResults function
type UpdateTestResultsArgs struct {
    // (required)
    Results *[]test.TestCaseResult
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
}

// [Preview API]
func (client *Client) QueryTestResultsReportForBuild(ctx context.Context, args QueryTestResultsReportForBuildArgs) (*test.TestResultSummary, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    if args.PublishContext != nil {
        queryParams.Add("publishContext", *args.PublishContext)
    }
    if args.IncludeFailureDetails != nil {
        queryParams.Add("includeFailureDetails", strconv.FormatBool(*args.IncludeFailureDetails))
    }
    if args.BuildToCompare != nil {
        if args.BuildToCompare.Id != nil {
            queryParams.Add("buildToCompare.id", strconv.Itoa(*args.BuildToCompare.Id))
        }
        if args.BuildToCompare.DefinitionId != nil {
            queryParams.Add("buildToCompare.definitionId", strconv.Itoa(*args.BuildToCompare.DefinitionId))
        }
        if args.BuildToCompare.Number != nil {
            queryParams.Add("buildToCompare.number", *args.BuildToCompare.Number)
        }
        if args.BuildToCompare.Uri != nil {
            queryParams.Add("buildToCompare.uri", *args.BuildToCompare.Uri)
        }
        if args.BuildToCompare.BuildSystem != nil {
            queryParams.Add("buildToCompare.buildSystem", *args.BuildToCompare.BuildSystem)
        }
        if args.BuildToCompare.BranchName != nil {
            queryParams.Add("buildToCompare.branchName", *args.BuildToCompare.BranchName)
        }
        if args.BuildToCompare.RepositoryId != nil {
            queryParams.Add("buildToCompare.repositoryId", *args.BuildToCompare.RepositoryId)
        }
    }
    locationId, _ := uuid.Parse("e009fa95-95a5-4ad4-9681-590043ce2423")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestResultsReportForBuild function
type QueryTestResultsReportForBuildArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    BuildId *int
    // (optional)
    PublishContext *string
    // (optional)
    IncludeFailureDetails *bool
    // (optional)
    BuildToCompare *test.BuildReference
}

// [Preview API]
func (client *Client) QueryTestResultsReportForRelease(ctx context.Context, args QueryTestResultsReportForReleaseArgs) (*test.TestResultSummary, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.ReleaseId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseId"}
    }
    queryParams.Add("releaseId", strconv.Itoa(*args.ReleaseId))
    if args.ReleaseEnvId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releaseEnvId"}
    }
    queryParams.Add("releaseEnvId", strconv.Itoa(*args.ReleaseEnvId))
    if args.PublishContext != nil {
        queryParams.Add("publishContext", *args.PublishContext)
    }
    if args.IncludeFailureDetails != nil {
        queryParams.Add("includeFailureDetails", strconv.FormatBool(*args.IncludeFailureDetails))
    }
    if args.ReleaseToCompare != nil {
        if args.ReleaseToCompare.Id != nil {
            queryParams.Add("releaseToCompare.id", strconv.Itoa(*args.ReleaseToCompare.Id))
        }
        if args.ReleaseToCompare.Name != nil {
            queryParams.Add("releaseToCompare.name", *args.ReleaseToCompare.Name)
        }
        if args.ReleaseToCompare.EnvironmentId != nil {
            queryParams.Add("releaseToCompare.environmentId", strconv.Itoa(*args.ReleaseToCompare.EnvironmentId))
        }
        if args.ReleaseToCompare.EnvironmentName != nil {
            queryParams.Add("releaseToCompare.environmentName", *args.ReleaseToCompare.EnvironmentName)
        }
        if args.ReleaseToCompare.DefinitionId != nil {
            queryParams.Add("releaseToCompare.definitionId", strconv.Itoa(*args.ReleaseToCompare.DefinitionId))
        }
        if args.ReleaseToCompare.EnvironmentDefinitionId != nil {
            queryParams.Add("releaseToCompare.environmentDefinitionId", strconv.Itoa(*args.ReleaseToCompare.EnvironmentDefinitionId))
        }
        if args.ReleaseToCompare.EnvironmentDefinitionName != nil {
            queryParams.Add("releaseToCompare.environmentDefinitionName", *args.ReleaseToCompare.EnvironmentDefinitionName)
        }
        if args.ReleaseToCompare.CreationDate != nil {
            queryParams.Add("releaseToCompare.creationDate", (*args.ReleaseToCompare.CreationDate).String())
        }
        if args.ReleaseToCompare.EnvironmentCreationDate != nil {
            queryParams.Add("releaseToCompare.environmentCreationDate", (*args.ReleaseToCompare.EnvironmentCreationDate).String())
        }
        if args.ReleaseToCompare.Attempt != nil {
            queryParams.Add("releaseToCompare.attempt", strconv.Itoa(*args.ReleaseToCompare.Attempt))
        }
    }
    locationId, _ := uuid.Parse("f10f9577-2c04-45ab-8c99-b26567a7cd55")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultSummary
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestResultsReportForRelease function
type QueryTestResultsReportForReleaseArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    ReleaseId *int
    // (required)
    ReleaseEnvId *int
    // (optional)
    PublishContext *string
    // (optional)
    IncludeFailureDetails *bool
    // (optional)
    ReleaseToCompare *test.ReleaseReference
}

// [Preview API]
func (client *Client) QueryTestResultsSummaryForReleases(ctx context.Context, args QueryTestResultsSummaryForReleasesArgs) (*[]test.TestResultSummary, error) {
    if args.Releases == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "releases"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Releases)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f10f9577-2c04-45ab-8c99-b26567a7cd55")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestResultSummary
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestResultsSummaryForReleases function
type QueryTestResultsSummaryForReleasesArgs struct {
    // (required)
    Releases *[]test.ReleaseReference
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) QueryTestSummaryByRequirement(ctx context.Context, args QueryTestSummaryByRequirementArgs) (*[]test.TestSummaryForWorkItem, error) {
    if args.ResultsContext == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "resultsContext"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.WorkItemIds != nil {
        var stringList []string
        for _, item := range *args.WorkItemIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    body, marshalErr := json.Marshal(*args.ResultsContext)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("3b7fd26f-c335-4e55-afc1-a588f5e2af3c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestSummaryForWorkItem
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestSummaryByRequirement function
type QueryTestSummaryByRequirementArgs struct {
    // (required)
    ResultsContext *test.TestResultsContext
    // (required) Project ID or project name
    Project *string
    // (optional)
    WorkItemIds *[]int
}

// [Preview API]
func (client *Client) QueryResultTrendForBuild(ctx context.Context, args QueryResultTrendForBuildArgs) (*[]test.AggregatedDataForResultTrend, error) {
    if args.Filter == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filter"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0886a7ae-315a-4dba-9122-bcce93301f3a")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.AggregatedDataForResultTrend
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryResultTrendForBuild function
type QueryResultTrendForBuildArgs struct {
    // (required)
    Filter *test.TestResultTrendFilter
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) QueryResultTrendForRelease(ctx context.Context, args QueryResultTrendForReleaseArgs) (*[]test.AggregatedDataForResultTrend, error) {
    if args.Filter == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filter"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("107f23c3-359a-460a-a70c-63ee739f9f9a")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.AggregatedDataForResultTrend
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryResultTrendForRelease function
type QueryResultTrendForReleaseArgs struct {
    // (required)
    Filter *test.TestResultTrendFilter
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) CreateTestRun(ctx context.Context, args CreateTestRunArgs) (*test.TestRun, error) {
    if args.TestRun == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testRun"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.TestRun)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestRun function
type CreateTestRunArgs struct {
    // (required)
    TestRun *test.RunCreateModel
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) DeleteTestRun(ctx context.Context, args DeleteTestRunArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteTestRun function
type DeleteTestRunArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
}

// [Preview API]
func (client *Client) GetTestRunById(ctx context.Context, args GetTestRunByIdArgs) (*test.TestRun, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.IncludeDetails != nil {
        queryParams.Add("includeDetails", strconv.FormatBool(*args.IncludeDetails))
    }
    if args.IncludeTags != nil {
        queryParams.Add("includeTags", strconv.FormatBool(*args.IncludeTags))
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunById function
type GetTestRunByIdArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (optional)
    IncludeDetails *bool
    // (optional)
    IncludeTags *bool
}

// [Preview API]
func (client *Client) GetTestRuns(ctx context.Context, args GetTestRunsArgs) (*[]test.TestRun, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildUri != nil {
        queryParams.Add("buildUri", *args.BuildUri)
    }
    if args.Owner != nil {
        queryParams.Add("owner", *args.Owner)
    }
    if args.TmiRunId != nil {
        queryParams.Add("tmiRunId", *args.TmiRunId)
    }
    if args.PlanId != nil {
        queryParams.Add("planId", strconv.Itoa(*args.PlanId))
    }
    if args.IncludeRunDetails != nil {
        queryParams.Add("includeRunDetails", strconv.FormatBool(*args.IncludeRunDetails))
    }
    if args.Automated != nil {
        queryParams.Add("automated", strconv.FormatBool(*args.Automated))
    }
    if args.Skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRuns function
type GetTestRunsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional)
    BuildUri *string
    // (optional)
    Owner *string
    // (optional)
    TmiRunId *string
    // (optional)
    PlanId *int
    // (optional)
    IncludeRunDetails *bool
    // (optional)
    Automated *bool
    // (optional)
    Skip *int
    // (optional)
    Top *int
}

// [Preview API] Query Test Runs based on filters. Mandatory fields are minLastUpdatedDate and maxLastUpdatedDate.
func (client *Client) QueryTestRuns(ctx context.Context, args QueryTestRunsArgs) (*[]test.TestRun, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.MinLastUpdatedDate == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "minLastUpdatedDate"}
    }
    queryParams.Add("minLastUpdatedDate", (*args.MinLastUpdatedDate).String())
    if args.MaxLastUpdatedDate == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "maxLastUpdatedDate"}
    }
    queryParams.Add("maxLastUpdatedDate", (*args.MaxLastUpdatedDate).String())
    if args.State != nil {
        queryParams.Add("state", string(*args.State))
    }
    if args.PlanIds != nil {
        var stringList []string
        for _, item := range *args.PlanIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.IsAutomated != nil {
        queryParams.Add("isAutomated", strconv.FormatBool(*args.IsAutomated))
    }
    if args.PublishContext != nil {
        queryParams.Add("publishContext", string(*args.PublishContext))
    }
    if args.BuildIds != nil {
        var stringList []string
        for _, item := range *args.BuildIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.BuildDefIds != nil {
        var stringList []string
        for _, item := range *args.BuildDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.BranchName != nil {
        queryParams.Add("branchName", *args.BranchName)
    }
    if args.ReleaseIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.ReleaseDefIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.ReleaseEnvIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseEnvIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.ReleaseEnvDefIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseEnvDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.RunTitle != nil {
        queryParams.Add("runTitle", *args.RunTitle)
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", *args.ContinuationToken)
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestRuns function
type QueryTestRunsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Minimum Last Modified Date of run to be queried (Mandatory).
    MinLastUpdatedDate *time.Time
    // (required) Maximum Last Modified Date of run to be queried (Mandatory, difference between min and max date can be atmost 7 days).
    MaxLastUpdatedDate *time.Time
    // (optional) Current state of the Runs to be queried.
    State *test.TestRunState
    // (optional) Plan Ids of the Runs to be queried, comma seperated list of valid ids.
    PlanIds *[]int
    // (optional) Automation type of the Runs to be queried.
    IsAutomated *bool
    // (optional) PublishContext of the Runs to be queried.
    PublishContext *test.TestRunPublishContext
    // (optional) Build Ids of the Runs to be queried, comma seperated list of valid ids.
    BuildIds *[]int
    // (optional) Build Definition Ids of the Runs to be queried, comma seperated list of valid ids.
    BuildDefIds *[]int
    // (optional) Source Branch name of the Runs to be queried.
    BranchName *string
    // (optional) Release Ids of the Runs to be queried, comma seperated list of valid ids.
    ReleaseIds *[]int
    // (optional) Release Definition Ids of the Runs to be queried, comma seperated list of valid ids.
    ReleaseDefIds *[]int
    // (optional) Release Environment Ids of the Runs to be queried, comma seperated list of valid ids.
    ReleaseEnvIds *[]int
    // (optional) Release Environment Definition Ids of the Runs to be queried, comma seperated list of valid ids.
    ReleaseEnvDefIds *[]int
    // (optional) Run Title of the Runs to be queried.
    RunTitle *string
    // (optional) Number of runs to be queried. Limit is 100
    Top *int
    // (optional) continuationToken received from previous batch or null for first batch. It is not supposed to be created (or altered, if received from last batch) by user.
    ContinuationToken *string
}

// [Preview API]
func (client *Client) UpdateTestRun(ctx context.Context, args UpdateTestRunArgs) (*test.TestRun, error) {
    if args.RunUpdateModel == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runUpdateModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.RunUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("364538f9-8062-4ce0-b024-75a0fb463f0d")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateTestRun function
type UpdateTestRunArgs struct {
    // (required)
    RunUpdateModel *test.RunUpdateModel
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
}

// Get test run statistics , used when we want to get summary of a run by outcome.
func (client *Client) GetTestRunStatistics(ctx context.Context, args GetTestRunStatisticsArgs) (*test.TestRunStatistic, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    locationId, _ := uuid.Parse("82b986e8-ca9e-4a89-b39e-f65c69bc104a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestRunStatistic
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunStatistics function
type GetTestRunStatisticsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the run to get.
    RunId *int
}

// [Preview API] Get TestResultsSettings data
func (client *Client) GetTestResultsSettings(ctx context.Context, args GetTestResultsSettingsArgs) (*test.TestResultsSettings, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.SettingsType != nil {
        queryParams.Add("settingsType", string(*args.SettingsType))
    }
    locationId, _ := uuid.Parse("7319952e-e5a9-4e19-a006-84f3be8b7c68")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.3", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultsSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultsSettings function
type GetTestResultsSettingsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional)
    SettingsType *test.TestResultsSettingsType
}

// [Preview API] Update project settings of test results
func (client *Client) UpdatePipelinesTestSettings(ctx context.Context, args UpdatePipelinesTestSettingsArgs) (*test.TestResultsSettings, error) {
    if args.TestResultsUpdateSettings == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testResultsUpdateSettings"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.TestResultsUpdateSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7319952e-e5a9-4e19-a006-84f3be8b7c68")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.3", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestResultsSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdatePipelinesTestSettings function
type UpdatePipelinesTestSettingsArgs struct {
    // (required)
    TestResultsUpdateSettings *test.TestResultsUpdateSettings
    // (required) Project ID or project name
    Project *string
}

// [Preview API] Get history of a test method using TestHistoryQuery
func (client *Client) QueryTestHistory(ctx context.Context, args QueryTestHistoryArgs) (*test.TestHistoryQuery, error) {
    if args.Filter == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filter"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2a41bd6a-8118-4403-b74e-5ba7492aed9d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestHistoryQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestHistory function
type QueryTestHistoryArgs struct {
    // (required) TestHistoryQuery to get history
    Filter *test.TestHistoryQuery
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) GetTestLogsForBuild(ctx context.Context, args GetTestLogsForBuildArgs) (*[]test.TestLog, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.DirectoryPath != nil {
        queryParams.Add("directoryPath", *args.DirectoryPath)
    }
    if args.FileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *args.FileNamePrefix)
    }
    if args.FetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*args.FetchMetaData))
    }
    if args.Top != nil {
        queryParams.Add("top", strconv.Itoa(*args.Top))
    }
    additionalHeaders := make(map[string]string)
    if args.ContinuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *args.ContinuationToken
    }
    locationId, _ := uuid.Parse("dff8ce3a-e539-4817-a405-d968491a88f1")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestLogsForBuild function
type GetTestLogsForBuildArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    BuildId *int
    // (required)
    Type_ *test.TestLogType
    // (optional)
    DirectoryPath *string
    // (optional)
    FileNamePrefix *string
    // (optional)
    FetchMetaData *bool
    // (optional)
    Top *int
    // (optional) Header to pass the continuationToken
    ContinuationToken *string
}

// [Preview API]
func (client *Client) GetTestResultLogs(ctx context.Context, args GetTestResultLogsArgs) (*[]test.TestLog, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.ResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*args.ResultId)

    queryParams := url.Values{}
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.DirectoryPath != nil {
        queryParams.Add("directoryPath", *args.DirectoryPath)
    }
    if args.FileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *args.FileNamePrefix)
    }
    if args.FetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*args.FetchMetaData))
    }
    if args.Top != nil {
        queryParams.Add("top", strconv.Itoa(*args.Top))
    }
    additionalHeaders := make(map[string]string)
    if args.ContinuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *args.ContinuationToken
    }
    locationId, _ := uuid.Parse("714caaac-ae1e-4869-8323-9bc0f5120dbf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultLogs function
type GetTestResultLogsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    ResultId *int
    // (required)
    Type_ *test.TestLogType
    // (optional)
    DirectoryPath *string
    // (optional)
    FileNamePrefix *string
    // (optional)
    FetchMetaData *bool
    // (optional)
    Top *int
    // (optional) Header to pass the continuationToken
    ContinuationToken *string
}

// [Preview API]
func (client *Client) GetTestSubResultLogs(ctx context.Context, args GetTestSubResultLogsArgs) (*[]test.TestLog, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.ResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*args.ResultId)

    queryParams := url.Values{}
    if args.SubResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "subResultId"}
    }
    queryParams.Add("subResultId", strconv.Itoa(*args.SubResultId))
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.DirectoryPath != nil {
        queryParams.Add("directoryPath", *args.DirectoryPath)
    }
    if args.FileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *args.FileNamePrefix)
    }
    if args.FetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*args.FetchMetaData))
    }
    if args.Top != nil {
        queryParams.Add("top", strconv.Itoa(*args.Top))
    }
    additionalHeaders := make(map[string]string)
    if args.ContinuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *args.ContinuationToken
    }
    locationId, _ := uuid.Parse("714caaac-ae1e-4869-8323-9bc0f5120dbf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestSubResultLogs function
type GetTestSubResultLogsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    ResultId *int
    // (required)
    SubResultId *int
    // (required)
    Type_ *test.TestLogType
    // (optional)
    DirectoryPath *string
    // (optional)
    FileNamePrefix *string
    // (optional)
    FetchMetaData *bool
    // (optional)
    Top *int
    // (optional) Header to pass the continuationToken
    ContinuationToken *string
}

// [Preview API]
func (client *Client) GetTestRunLogs(ctx context.Context, args GetTestRunLogsArgs) (*[]test.TestLog, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.DirectoryPath != nil {
        queryParams.Add("directoryPath", *args.DirectoryPath)
    }
    if args.FileNamePrefix != nil {
        queryParams.Add("fileNamePrefix", *args.FileNamePrefix)
    }
    if args.FetchMetaData != nil {
        queryParams.Add("fetchMetaData", strconv.FormatBool(*args.FetchMetaData))
    }
    if args.Top != nil {
        queryParams.Add("top", strconv.Itoa(*args.Top))
    }
    additionalHeaders := make(map[string]string)
    if args.ContinuationToken != nil {
        additionalHeaders["x-ms-continuationtoken"] = *args.ContinuationToken
    }
    locationId, _ := uuid.Parse("5b47b946-e875-4c9a-acdc-2a20996caebe")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseValue []test.TestLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunLogs function
type GetTestRunLogsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    Type_ *test.TestLogType
    // (optional)
    DirectoryPath *string
    // (optional)
    FileNamePrefix *string
    // (optional)
    FetchMetaData *bool
    // (optional)
    Top *int
    // (optional) Header to pass the continuationToken
    ContinuationToken *string
}

// [Preview API]
func (client *Client) GetTestLogStoreEndpointDetailsForBuildLog(ctx context.Context, args GetTestLogStoreEndpointDetailsForBuildLogArgs) (*test.TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.Build == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "build"}
    }
    queryParams.Add("build", strconv.Itoa(*args.Build))
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.FilePath == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *args.FilePath)
    locationId, _ := uuid.Parse("39b09be7-f0c9-4a83-a513-9ae31b45c56f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestLogStoreEndpointDetailsForBuildLog function
type GetTestLogStoreEndpointDetailsForBuildLogArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    Build *int
    // (required)
    Type_ *test.TestLogType
    // (required)
    FilePath *string
}

// [Preview API]
func (client *Client) TestLogStoreEndpointDetailsForBuild(ctx context.Context, args TestLogStoreEndpointDetailsForBuildArgs) (*test.TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    if args.TestLogStoreOperationType == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testLogStoreOperationType"}
    }
    queryParams.Add("testLogStoreOperationType", string(*args.TestLogStoreOperationType))
    locationId, _ := uuid.Parse("39b09be7-f0c9-4a83-a513-9ae31b45c56f")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the TestLogStoreEndpointDetailsForBuild function
type TestLogStoreEndpointDetailsForBuildArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    BuildId *int
    // (required)
    TestLogStoreOperationType *test.TestLogStoreOperationType
}

// [Preview API]
func (client *Client) GetTestLogStoreEndpointDetailsForResultLog(ctx context.Context, args GetTestLogStoreEndpointDetailsForResultLogArgs) (*test.TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.ResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*args.ResultId)

    queryParams := url.Values{}
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.FilePath == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *args.FilePath)
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestLogStoreEndpointDetailsForResultLog function
type GetTestLogStoreEndpointDetailsForResultLogArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    ResultId *int
    // (required)
    Type_ *test.TestLogType
    // (required)
    FilePath *string
}

// [Preview API]
func (client *Client) GetTestLogStoreEndpointDetailsForSubResultLog(ctx context.Context, args GetTestLogStoreEndpointDetailsForSubResultLogArgs) (*test.TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.ResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*args.ResultId)

    queryParams := url.Values{}
    if args.SubResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "subResultId"}
    }
    queryParams.Add("subResultId", strconv.Itoa(*args.SubResultId))
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.FilePath == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *args.FilePath)
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestLogStoreEndpointDetailsForSubResultLog function
type GetTestLogStoreEndpointDetailsForSubResultLogArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    ResultId *int
    // (required)
    SubResultId *int
    // (required)
    Type_ *test.TestLogType
    // (required)
    FilePath *string
}

// [Preview API]
func (client *Client) TestLogStoreEndpointDetailsForResult(ctx context.Context, args TestLogStoreEndpointDetailsForResultArgs) (*test.TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.ResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "resultId"} 
    }
    routeValues["resultId"] = strconv.Itoa(*args.ResultId)

    queryParams := url.Values{}
    if args.SubResultId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "subResultId"}
    }
    queryParams.Add("subResultId", strconv.Itoa(*args.SubResultId))
    if args.FilePath == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *args.FilePath)
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the TestLogStoreEndpointDetailsForResult function
type TestLogStoreEndpointDetailsForResultArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    ResultId *int
    // (required)
    SubResultId *int
    // (required)
    FilePath *string
    // (required)
    Type_ *test.TestLogType
}

// [Preview API]
func (client *Client) GetTestLogStoreEndpointDetailsForRunLog(ctx context.Context, args GetTestLogStoreEndpointDetailsForRunLogArgs) (*test.TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.Type_ == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "type_"}
    }
    queryParams.Add("type_", string(*args.Type_))
    if args.FilePath == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "filePath"}
    }
    queryParams.Add("filePath", *args.FilePath)
    locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestLogStoreEndpointDetailsForRunLog function
type GetTestLogStoreEndpointDetailsForRunLogArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    Type_ *test.TestLogType
    // (required)
    FilePath *string
}

// [Preview API]
func (client *Client) TestLogStoreEndpointDetailsForRun(ctx context.Context, args TestLogStoreEndpointDetailsForRunArgs) (*test.TestLogStoreEndpointDetails, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.TestLogStoreOperationType == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testLogStoreOperationType"}
    }
    queryParams.Add("testLogStoreOperationType", string(*args.TestLogStoreOperationType))
    if args.FilePath != nil {
        queryParams.Add("filePath", *args.FilePath)
    }
    if args.Type_ != nil {
        queryParams.Add("type_", string(*args.Type_))
    }
    locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestLogStoreEndpointDetails
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the TestLogStoreEndpointDetailsForRun function
type TestLogStoreEndpointDetailsForRunArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    RunId *int
    // (required)
    TestLogStoreOperationType *test.TestLogStoreOperationType
    // (optional)
    FilePath *string
    // (optional)
    Type_ *test.TestLogType
}

// [Preview API]
func (client *Client) CreateTestSettings(ctx context.Context, args CreateTestSettingsArgs) (*int, error) {
    if args.TestSettings == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testSettings"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.TestSettings)
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

// Arguments for the CreateTestSettings function
type CreateTestSettingsArgs struct {
    // (required)
    TestSettings *test.TestSettings
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) DeleteTestSettings(ctx context.Context, args DeleteTestSettingsArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.TestSettingsId == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "testSettingsId"}
    }
    queryParams.Add("testSettingsId", strconv.Itoa(*args.TestSettingsId))
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteTestSettings function
type DeleteTestSettingsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    TestSettingsId *int
}

// [Preview API]
func (client *Client) GetTestSettingsById(ctx context.Context, args GetTestSettingsByIdArgs) (*test.TestSettings, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.TestSettingsId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testSettingsId"}
    }
    queryParams.Add("testSettingsId", strconv.Itoa(*args.TestSettingsId))
    locationId, _ := uuid.Parse("930bad47-f826-4099-9597-f44d0a9c735c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestSettingsById function
type GetTestSettingsByIdArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    TestSettingsId *int
}

// [Preview API]
func (client *Client) AddWorkItemToTestLinks(ctx context.Context, args AddWorkItemToTestLinksArgs) (*test.WorkItemToTestLinks, error) {
    if args.WorkItemToTestLinks == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "workItemToTestLinks"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.WorkItemToTestLinks)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4e3abe63-ca46-4fe0-98b2-363f7ec7aa5f")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.WorkItemToTestLinks
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the AddWorkItemToTestLinks function
type AddWorkItemToTestLinksArgs struct {
    // (required)
    WorkItemToTestLinks *test.WorkItemToTestLinks
    // (required) Project ID or project name
    Project *string
}

// [Preview API]
func (client *Client) DeleteTestMethodToWorkItemLink(ctx context.Context, args DeleteTestMethodToWorkItemLinkArgs) (*bool, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.TestName == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testName"}
    }
    queryParams.Add("testName", *args.TestName)
    if args.WorkItemId == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "workItemId"}
    }
    queryParams.Add("workItemId", strconv.Itoa(*args.WorkItemId))
    locationId, _ := uuid.Parse("cbd50bd7-f7ed-4e35-b127-4408ae6bfa2c")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue bool
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the DeleteTestMethodToWorkItemLink function
type DeleteTestMethodToWorkItemLinkArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    TestName *string
    // (required)
    WorkItemId *int
}

// [Preview API]
func (client *Client) QueryTestMethodLinkedWorkItems(ctx context.Context, args QueryTestMethodLinkedWorkItemsArgs) (*test.TestToWorkItemLinks, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.TestName == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testName"}
    }
    queryParams.Add("testName", *args.TestName)
    locationId, _ := uuid.Parse("cbd50bd7-f7ed-4e35-b127-4408ae6bfa2c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue test.TestToWorkItemLinks
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestMethodLinkedWorkItems function
type QueryTestMethodLinkedWorkItemsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    TestName *string
}

// [Preview API]
func (client *Client) QueryTestResultWorkItems(ctx context.Context, args QueryTestResultWorkItemsArgs) (*[]test.WorkItemReference, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.WorkItemCategory == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "workItemCategory"}
    }
    queryParams.Add("workItemCategory", *args.WorkItemCategory)
    if args.AutomatedTestName != nil {
        queryParams.Add("automatedTestName", *args.AutomatedTestName)
    }
    if args.TestCaseId != nil {
        queryParams.Add("testCaseId", strconv.Itoa(*args.TestCaseId))
    }
    if args.MaxCompleteDate != nil {
        queryParams.Add("maxCompleteDate", (*args.MaxCompleteDate).String())
    }
    if args.Days != nil {
        queryParams.Add("days", strconv.Itoa(*args.Days))
    }
    if args.WorkItemCount != nil {
        queryParams.Add("$workItemCount", strconv.Itoa(*args.WorkItemCount))
    }
    locationId, _ := uuid.Parse("f7401a26-331b-44fe-a470-f7ed35138e4a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []test.WorkItemReference
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestResultWorkItems function
type QueryTestResultWorkItemsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required)
    WorkItemCategory *string
    // (optional)
    AutomatedTestName *string
    // (optional)
    TestCaseId *int
    // (optional)
    MaxCompleteDate *time.Time
    // (optional)
    Days *int
    // (optional)
    WorkItemCount *int
}

