// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testresults

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7.1"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7.1/test"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("c83eaf52-edf3-4034-ae11-17d38f25404c")

type Client interface {
	// [Preview API] Creates an attachment in the LogStore for the specified buildId.
	CreateBuildAttachmentInLogStore(context.Context, CreateBuildAttachmentInLogStoreArgs) error
	// [Preview API] Creates a new test failure type
	CreateFailureType(context.Context, CreateFailureTypeArgs) (*test.TestFailureType, error)
	// [Preview API] Creates an attachment in the LogStore for the specified runId.
	CreateTestRunLogStoreAttachment(context.Context, CreateTestRunLogStoreAttachmentArgs) (*test.TestLogStoreAttachmentReference, error)
	// [Preview API] Deletes a test failure type with specified failureTypeId
	DeleteFailureType(context.Context, DeleteFailureTypeArgs) error
	// [Preview API] Deletes the attachment with the specified filename for the specified runId from the LogStore.
	DeleteTestRunLogStoreAttachment(context.Context, DeleteTestRunLogStoreAttachmentArgs) error
	// [Preview API] Returns the list of test failure types.
	GetFailureTypes(context.Context, GetFailureTypesArgs) (*[]test.TestFailureType, error)
	// [Preview API] Get SAS Uri of a test results attachment
	GetTestLogStoreEndpointDetailsForResultLog(context.Context, GetTestLogStoreEndpointDetailsForResultLogArgs) (*test.TestLogStoreEndpointDetails, error)
	// [Preview API] Get SAS Uri of a test run attachment
	GetTestLogStoreEndpointDetailsForRunLog(context.Context, GetTestLogStoreEndpointDetailsForRunLogArgs) (*test.TestLogStoreEndpointDetails, error)
	// [Preview API] Get SAS Uri of a test subresults attachment
	GetTestLogStoreEndpointDetailsForSubResultLog(context.Context, GetTestLogStoreEndpointDetailsForSubResultLogArgs) (*test.TestLogStoreEndpointDetails, error)
	// [Preview API] Get list of test result attachments reference
	GetTestResultLogs(context.Context, GetTestResultLogsArgs) (*GetTestResultLogsResponseValue, error)
	// [Preview API] Get list of test run attachments reference
	GetTestRunLogs(context.Context, GetTestRunLogsArgs) (*GetTestRunLogsResponseValue, error)
	// [Preview API] Returns the attachment with the specified filename for the specified runId from the LogStore.
	GetTestRunLogStoreAttachmentContent(context.Context, GetTestRunLogStoreAttachmentContentArgs) (io.ReadCloser, error)
	// [Preview API] Returns a list of attachments for the specified runId from the LogStore.
	GetTestRunLogStoreAttachments(context.Context, GetTestRunLogStoreAttachmentsArgs) (*[]test.TestLogStoreAttachment, error)
	// [Preview API] Returns the attachment with the specified filename for the specified runId from the LogStore.
	GetTestRunLogStoreAttachmentZip(context.Context, GetTestRunLogStoreAttachmentZipArgs) (io.ReadCloser, error)
	// [Preview API] Get list of test subresult attachments reference
	GetTestSubResultLogs(context.Context, GetTestSubResultLogsArgs) (*GetTestSubResultLogsResponseValue, error)
	// [Preview API] Get list of test Result meta data details for corresponding testcasereferenceId
	QueryTestResultsMetaData(context.Context, QueryTestResultsMetaDataArgs) (*[]test.TestResultMetaData, error)
	// [Preview API] Create empty file for a result and Get Sas uri for the file
	TestLogStoreEndpointDetailsForResult(context.Context, TestLogStoreEndpointDetailsForResultArgs) (*test.TestLogStoreEndpointDetails, error)
	// [Preview API] Create empty file for a run and Get Sas uri for the file
	TestLogStoreEndpointDetailsForRun(context.Context, TestLogStoreEndpointDetailsForRunArgs) (*test.TestLogStoreEndpointDetails, error)
	// [Preview API] Update properties of test result meta data
	UpdateTestResultsMetaData(context.Context, UpdateTestResultsMetaDataArgs) (*test.TestResultMetaData, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// [Preview API] Creates an attachment in the LogStore for the specified buildId.
func (client *ClientImpl) CreateBuildAttachmentInLogStore(ctx context.Context, args CreateBuildAttachmentInLogStoreArgs) error {
	if args.AttachmentRequestModel == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.AttachmentRequestModel"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.BuildId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.BuildId"}
	}
	routeValues["buildId"] = strconv.Itoa(*args.BuildId)

	body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("6f747e16-18c2-435a-b4fb-fa05d6845fee")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the CreateBuildAttachmentInLogStore function
type CreateBuildAttachmentInLogStoreArgs struct {
	// (required) Contains attachment info like stream, filename, comment, attachmentType
	AttachmentRequestModel *test.TestAttachmentRequestModel
	// (required) Project ID or project name
	Project *string
	// (required) BuildId
	BuildId *int
}

// [Preview API] Creates a new test failure type
func (client *ClientImpl) CreateFailureType(ctx context.Context, args CreateFailureTypeArgs) (*test.TestFailureType, error) {
	if args.TestResultFailureType == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestResultFailureType"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.TestResultFailureType)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("c4ac0486-830c-4a2a-9ef9-e8a1791a70fd")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue test.TestFailureType
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateFailureType function
type CreateFailureTypeArgs struct {
	// (required)
	TestResultFailureType *test.TestFailureType
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Creates an attachment in the LogStore for the specified runId.
func (client *ClientImpl) CreateTestRunLogStoreAttachment(ctx context.Context, args CreateTestRunLogStoreAttachmentArgs) (*test.TestLogStoreAttachmentReference, error) {
	if args.AttachmentRequestModel == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.AttachmentRequestModel"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1026d5de-4b0b-46ae-a31f-7c59b6af51ef")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue test.TestLogStoreAttachmentReference
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateTestRunLogStoreAttachment function
type CreateTestRunLogStoreAttachmentArgs struct {
	// (required) Contains attachment info like stream, filename, comment, attachmentType
	AttachmentRequestModel *test.TestAttachmentRequestModel
	// (required) Project ID or project name
	Project *string
	// (required) Test RunId
	RunId *int
}

// [Preview API] Deletes a test failure type with specified failureTypeId
func (client *ClientImpl) DeleteFailureType(ctx context.Context, args DeleteFailureTypeArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.FailureTypeId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.FailureTypeId"}
	}
	routeValues["failureTypeId"] = strconv.Itoa(*args.FailureTypeId)

	locationId, _ := uuid.Parse("c4ac0486-830c-4a2a-9ef9-e8a1791a70fd")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteFailureType function
type DeleteFailureTypeArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required)
	FailureTypeId *int
}

// [Preview API] Deletes the attachment with the specified filename for the specified runId from the LogStore.
func (client *ClientImpl) DeleteTestRunLogStoreAttachment(ctx context.Context, args DeleteTestRunLogStoreAttachmentArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	queryParams := url.Values{}
	if args.Filename == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "filename"}
	}
	queryParams.Add("filename", *args.Filename)
	locationId, _ := uuid.Parse("1026d5de-4b0b-46ae-a31f-7c59b6af51ef")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTestRunLogStoreAttachment function
type DeleteTestRunLogStoreAttachmentArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Test RunId
	RunId *int
	// (required) Attachment FileName
	Filename *string
}

// [Preview API] Returns the list of test failure types.
func (client *ClientImpl) GetFailureTypes(ctx context.Context, args GetFailureTypesArgs) (*[]test.TestFailureType, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("c4ac0486-830c-4a2a-9ef9-e8a1791a70fd")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []test.TestFailureType
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetFailureTypes function
type GetFailureTypesArgs struct {
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Get SAS Uri of a test results attachment
func (client *ClientImpl) GetTestLogStoreEndpointDetailsForResultLog(ctx context.Context, args GetTestLogStoreEndpointDetailsForResultLogArgs) (*test.TestLogStoreEndpointDetails, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)
	if args.ResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ResultId"}
	}
	routeValues["resultId"] = strconv.Itoa(*args.ResultId)

	queryParams := url.Values{}
	if args.Type == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "type"}
	}
	queryParams.Add("type", string(*args.Type))
	if args.FilePath == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "filePath"}
	}
	queryParams.Add("filePath", *args.FilePath)
	locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	// (required) Id of the test run that contains result
	RunId *int
	// (required) Id of the test result whose files need to be downloaded
	ResultId *int
	// (required) type of the file
	Type *test.TestLogType
	// (required) filePath for which sas uri is needed
	FilePath *string
}

// [Preview API] Get SAS Uri of a test run attachment
func (client *ClientImpl) GetTestLogStoreEndpointDetailsForRunLog(ctx context.Context, args GetTestLogStoreEndpointDetailsForRunLogArgs) (*test.TestLogStoreEndpointDetails, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	queryParams := url.Values{}
	if args.Type == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "type"}
	}
	queryParams.Add("type", string(*args.Type))
	if args.FilePath == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "filePath"}
	}
	queryParams.Add("filePath", *args.FilePath)
	locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	// (required) Id of the test run whose file has to be downloaded
	RunId *int
	// (required) type of the file
	Type *test.TestLogType
	// (required) filePath for which sas uri is needed
	FilePath *string
}

// [Preview API] Get SAS Uri of a test subresults attachment
func (client *ClientImpl) GetTestLogStoreEndpointDetailsForSubResultLog(ctx context.Context, args GetTestLogStoreEndpointDetailsForSubResultLogArgs) (*test.TestLogStoreEndpointDetails, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)
	if args.ResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ResultId"}
	}
	routeValues["resultId"] = strconv.Itoa(*args.ResultId)

	queryParams := url.Values{}
	if args.SubResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "subResultId"}
	}
	queryParams.Add("subResultId", strconv.Itoa(*args.SubResultId))
	if args.Type == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "type"}
	}
	queryParams.Add("type", string(*args.Type))
	if args.FilePath == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "filePath"}
	}
	queryParams.Add("filePath", *args.FilePath)
	locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	// (required) Id of the test run that contains result
	RunId *int
	// (required) Id of the test result that contains subresult
	ResultId *int
	// (required) Id of the test subresult whose file sas uri is needed
	SubResultId *int
	// (required) type of the file
	Type *test.TestLogType
	// (required) filePath for which sas uri is needed
	FilePath *string
}

// [Preview API] Get list of test result attachments reference
func (client *ClientImpl) GetTestResultLogs(ctx context.Context, args GetTestResultLogsArgs) (*GetTestResultLogsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)
	if args.ResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ResultId"}
	}
	routeValues["resultId"] = strconv.Itoa(*args.ResultId)

	queryParams := url.Values{}
	if args.Type == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "type"}
	}
	queryParams.Add("type", string(*args.Type))
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestResultLogsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestResultLogs function
type GetTestResultLogsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the test run that contains the result
	RunId *int
	// (required) Id of the test result
	ResultId *int
	// (required) type of attachments to get
	Type *test.TestLogType
	// (optional) directory path of attachments to get
	DirectoryPath *string
	// (optional) file name prefix to filter the list of attachment
	FileNamePrefix *string
	// (optional) Default is false, set if metadata is needed
	FetchMetaData *bool
	// (optional) Numbe of attachments reference to return
	Top *int
	// (optional) Header to pass the continuationToken
	ContinuationToken *string
}

// Return type for the GetTestResultLogs function
type GetTestResultLogsResponseValue struct {
	Value             []test.TestLog
	ContinuationToken string
}

// [Preview API] Get list of test run attachments reference
func (client *ClientImpl) GetTestRunLogs(ctx context.Context, args GetTestRunLogsArgs) (*GetTestRunLogsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	queryParams := url.Values{}
	if args.Type == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "type"}
	}
	queryParams.Add("type", string(*args.Type))
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestRunLogsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestRunLogs function
type GetTestRunLogsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the test run
	RunId *int
	// (required) type of the attachments to get
	Type *test.TestLogType
	// (optional) directory path for which attachments are needed
	DirectoryPath *string
	// (optional) file name prefix to filter the list of attachment
	FileNamePrefix *string
	// (optional) Default is false, set if metadata is needed
	FetchMetaData *bool
	// (optional) Number of attachments reference to return
	Top *int
	// (optional) Header to pass the continuationToken
	ContinuationToken *string
}

// Return type for the GetTestRunLogs function
type GetTestRunLogsResponseValue struct {
	Value             []test.TestLog
	ContinuationToken string
}

// [Preview API] Returns the attachment with the specified filename for the specified runId from the LogStore.
func (client *ClientImpl) GetTestRunLogStoreAttachmentContent(ctx context.Context, args GetTestRunLogStoreAttachmentContentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	queryParams := url.Values{}
	if args.Filename == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "filename"}
	}
	queryParams.Add("filename", *args.Filename)
	locationId, _ := uuid.Parse("1026d5de-4b0b-46ae-a31f-7c59b6af51ef")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetTestRunLogStoreAttachmentContent function
type GetTestRunLogStoreAttachmentContentArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Test RunId
	RunId *int
	// (required) Attachment FileName
	Filename *string
}

// [Preview API] Returns a list of attachments for the specified runId from the LogStore.
func (client *ClientImpl) GetTestRunLogStoreAttachments(ctx context.Context, args GetTestRunLogStoreAttachmentsArgs) (*[]test.TestLogStoreAttachment, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	locationId, _ := uuid.Parse("1026d5de-4b0b-46ae-a31f-7c59b6af51ef")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []test.TestLogStoreAttachment
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTestRunLogStoreAttachments function
type GetTestRunLogStoreAttachmentsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Test RunId
	RunId *int
}

// [Preview API] Returns the attachment with the specified filename for the specified runId from the LogStore.
func (client *ClientImpl) GetTestRunLogStoreAttachmentZip(ctx context.Context, args GetTestRunLogStoreAttachmentZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	queryParams := url.Values{}
	if args.Filename == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "filename"}
	}
	queryParams.Add("filename", *args.Filename)
	locationId, _ := uuid.Parse("1026d5de-4b0b-46ae-a31f-7c59b6af51ef")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetTestRunLogStoreAttachmentZip function
type GetTestRunLogStoreAttachmentZipArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Test RunId
	RunId *int
	// (required) Attachment FileName
	Filename *string
}

// [Preview API] Get list of test subresult attachments reference
func (client *ClientImpl) GetTestSubResultLogs(ctx context.Context, args GetTestSubResultLogsArgs) (*GetTestSubResultLogsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)
	if args.ResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ResultId"}
	}
	routeValues["resultId"] = strconv.Itoa(*args.ResultId)

	queryParams := url.Values{}
	if args.SubResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "subResultId"}
	}
	queryParams.Add("subResultId", strconv.Itoa(*args.SubResultId))
	if args.Type == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "type"}
	}
	queryParams.Add("type", string(*args.Type))
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestSubResultLogsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestSubResultLogs function
type GetTestSubResultLogsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the test run that contains the results
	RunId *int
	// (required) Id of the test result that contains subresult
	ResultId *int
	// (required) Id of the test subresult
	SubResultId *int
	// (required) type of the attachments to get
	Type *test.TestLogType
	// (optional) directory path of the attachment to get
	DirectoryPath *string
	// (optional) file name prefix to filter the list of attachments
	FileNamePrefix *string
	// (optional) Default is false, set if metadata is needed
	FetchMetaData *bool
	// (optional) Number of attachments reference to return
	Top *int
	// (optional) Header to pass the continuationToken
	ContinuationToken *string
}

// Return type for the GetTestSubResultLogs function
type GetTestSubResultLogsResponseValue struct {
	Value             []test.TestLog
	ContinuationToken string
}

// [Preview API] Get list of test Result meta data details for corresponding testcasereferenceId
func (client *ClientImpl) QueryTestResultsMetaData(ctx context.Context, args QueryTestResultsMetaDataArgs) (*[]test.TestResultMetaData, error) {
	if args.TestCaseReferenceIds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestCaseReferenceIds"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.DetailsToInclude != nil {
		queryParams.Add("detailsToInclude", string(*args.DetailsToInclude))
	}
	body, marshalErr := json.Marshal(*args.TestCaseReferenceIds)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("b72ff4c0-4341-4213-ba27-f517cf341c95")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.4", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []test.TestResultMetaData
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the QueryTestResultsMetaData function
type QueryTestResultsMetaDataArgs struct {
	// (required) TestCaseReference Ids of the test Result to be queried, comma separated list of valid ids (limit no. of ids 200).
	TestCaseReferenceIds *[]string
	// (required) Project ID or project name
	Project *string
	// (optional) Details to include with test results metadata. Default is None. Other values are FlakyIdentifiers.
	DetailsToInclude *test.ResultMetaDataDetails
}

// [Preview API] Create empty file for a result and Get Sas uri for the file
func (client *ClientImpl) TestLogStoreEndpointDetailsForResult(ctx context.Context, args TestLogStoreEndpointDetailsForResultArgs) (*test.TestLogStoreEndpointDetails, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)
	if args.ResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ResultId"}
	}
	routeValues["resultId"] = strconv.Itoa(*args.ResultId)

	queryParams := url.Values{}
	if args.SubResultId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "subResultId"}
	}
	queryParams.Add("subResultId", strconv.Itoa(*args.SubResultId))
	if args.FilePath == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "filePath"}
	}
	queryParams.Add("filePath", *args.FilePath)
	if args.Type == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "type"}
	}
	queryParams.Add("type", string(*args.Type))
	locationId, _ := uuid.Parse("da630b37-1236-45b5-945e-1d7bdb673850")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	// (required) Id of the test run that contains the result
	RunId *int
	// (required) Id of the test results that contains sub result
	ResultId *int
	// (required) Id of the test sub result whose file sas uri is needed
	SubResultId *int
	// (required) file path inside the sub result for which sas uri is needed
	FilePath *string
	// (required) Type of the file for download
	Type *test.TestLogType
}

// [Preview API] Create empty file for a run and Get Sas uri for the file
func (client *ClientImpl) TestLogStoreEndpointDetailsForRun(ctx context.Context, args TestLogStoreEndpointDetailsForRunArgs) (*test.TestLogStoreEndpointDetails, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RunId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RunId"}
	}
	routeValues["runId"] = strconv.Itoa(*args.RunId)

	queryParams := url.Values{}
	if args.TestLogStoreOperationType == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "testLogStoreOperationType"}
	}
	queryParams.Add("testLogStoreOperationType", string(*args.TestLogStoreOperationType))
	if args.FilePath != nil {
		queryParams.Add("filePath", *args.FilePath)
	}
	if args.Type != nil {
		queryParams.Add("type", string(*args.Type))
	}
	locationId, _ := uuid.Parse("67eb3f92-6c97-4fd9-8b63-6cbdc7e526ea")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	// (required) Id of the run to get endpoint details
	RunId *int
	// (required) Type of operation to perform using sas uri
	TestLogStoreOperationType *test.TestLogStoreOperationType
	// (optional) file path to create an empty file
	FilePath *string
	// (optional) Default is GeneralAttachment, type of empty file to be created
	Type *test.TestLogType
}

// [Preview API] Update properties of test result meta data
func (client *ClientImpl) UpdateTestResultsMetaData(ctx context.Context, args UpdateTestResultsMetaDataArgs) (*test.TestResultMetaData, error) {
	if args.TestResultMetaDataUpdateInput == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestResultMetaDataUpdateInput"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TestCaseReferenceId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestCaseReferenceId"}
	}
	routeValues["testCaseReferenceId"] = strconv.Itoa(*args.TestCaseReferenceId)

	body, marshalErr := json.Marshal(*args.TestResultMetaDataUpdateInput)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("b72ff4c0-4341-4213-ba27-f517cf341c95")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.4", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue test.TestResultMetaData
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateTestResultsMetaData function
type UpdateTestResultsMetaDataArgs struct {
	// (required) TestResultMetaData update input TestResultMetaDataUpdateInput
	TestResultMetaDataUpdateInput *test.TestResultMetaDataUpdateInput
	// (required) Project ID or project name
	Project *string
	// (required) TestCaseReference Id of Test Result to be updated.
	TestCaseReferenceId *int
}
