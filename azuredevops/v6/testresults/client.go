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
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/test"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("c83eaf52-edf3-4034-ae11-17d38f25404c")

type Client interface {
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
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
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", additionalHeaders)
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
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.4", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
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
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.4", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
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
