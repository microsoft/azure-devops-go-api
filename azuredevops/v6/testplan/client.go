// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testplan

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	// [Preview API] Add test cases to a suite with specified configurations
	AddTestCasesToSuite(context.Context, AddTestCasesToSuiteArgs) (*[]TestCase, error)
	// [Preview API]
	CloneTestCase(context.Context, CloneTestCaseArgs) (*CloneTestCaseOperationInformation, error)
	// [Preview API] Clone test plan
	CloneTestPlan(context.Context, CloneTestPlanArgs) (*CloneTestPlanOperationInformation, error)
	// [Preview API] Clone test suite
	CloneTestSuite(context.Context, CloneTestSuiteArgs) (*CloneTestSuiteOperationInformation, error)
	// [Preview API] Create a test configuration.
	CreateTestConfiguration(context.Context, CreateTestConfigurationArgs) (*TestConfiguration, error)
	// [Preview API] Create a test plan.
	CreateTestPlan(context.Context, CreateTestPlanArgs) (*TestPlan, error)
	// [Preview API] Create test suite.
	CreateTestSuite(context.Context, CreateTestSuiteArgs) (*TestSuite, error)
	// [Preview API] Create a test variable.
	CreateTestVariable(context.Context, CreateTestVariableArgs) (*TestVariable, error)
	// [Preview API] Delete a test case.
	DeleteTestCase(context.Context, DeleteTestCaseArgs) error
	// [Preview API] Delete a test configuration by its ID.
	DeleteTestConfguration(context.Context, DeleteTestConfgurationArgs) error
	// [Preview API] Delete a test plan.
	DeleteTestPlan(context.Context, DeleteTestPlanArgs) error
	// [Preview API] Delete test suite.
	DeleteTestSuite(context.Context, DeleteTestSuiteArgs) error
	// [Preview API] Delete a test variable by its ID.
	DeleteTestVariable(context.Context, DeleteTestVariableArgs) error
	// [Preview API] Get clone information.
	GetCloneInformation(context.Context, GetCloneInformationArgs) (*CloneTestPlanOperationInformation, error)
	// [Preview API] Get a list of points based on point Ids provided.
	GetPoints(context.Context, GetPointsArgs) (*[]TestPoint, error)
	// [Preview API] Get all the points inside a suite based on some filters
	GetPointsList(context.Context, GetPointsListArgs) (*GetPointsListResponseValue, error)
	// [Preview API] Get clone information.
	GetSuiteCloneInformation(context.Context, GetSuiteCloneInformationArgs) (*CloneTestSuiteOperationInformation, error)
	// [Preview API] Get a list of test suite entries in the test suite.
	GetSuiteEntries(context.Context, GetSuiteEntriesArgs) (*[]SuiteEntry, error)
	// Find the list of all test suites in which a given test case is present. This is helpful if you need to find out which test suites are using a test case, when you need to make changes to a test case.
	GetSuitesByTestCaseId(context.Context, GetSuitesByTestCaseIdArgs) (*[]TestSuite, error)
	// [Preview API] Get Test Cases For a Suite.
	GetTestCase(context.Context, GetTestCaseArgs) (*[]TestCase, error)
	// [Preview API] Get clone information.
	GetTestCaseCloneInformation(context.Context, GetTestCaseCloneInformationArgs) (*CloneTestCaseOperationInformation, error)
	// [Preview API] Get Test Case List return those test cases which have all the configuration Ids as mentioned in the optional parameter. If configuration Ids is null, it return all the test cases
	GetTestCaseList(context.Context, GetTestCaseListArgs) (*GetTestCaseListResponseValue, error)
	// [Preview API] Get a test configuration
	GetTestConfigurationById(context.Context, GetTestConfigurationByIdArgs) (*TestConfiguration, error)
	// [Preview API] Get a list of test configurations.
	GetTestConfigurations(context.Context, GetTestConfigurationsArgs) (*GetTestConfigurationsResponseValue, error)
	// [Preview API] Get a test plan by Id.
	GetTestPlanById(context.Context, GetTestPlanByIdArgs) (*TestPlan, error)
	// [Preview API] Get a list of test plans
	GetTestPlans(context.Context, GetTestPlansArgs) (*GetTestPlansResponseValue, error)
	// [Preview API] Get test suite by suite id.
	GetTestSuiteById(context.Context, GetTestSuiteByIdArgs) (*TestSuite, error)
	// [Preview API] Get test suites for plan.
	GetTestSuitesForPlan(context.Context, GetTestSuitesForPlanArgs) (*GetTestSuitesForPlanResponseValue, error)
	// [Preview API] Get a test variable by its ID.
	GetTestVariableById(context.Context, GetTestVariableByIdArgs) (*TestVariable, error)
	// [Preview API] Get a list of test variables.
	GetTestVariables(context.Context, GetTestVariablesArgs) (*GetTestVariablesResponseValue, error)
	// [Preview API] Removes test cases from a suite based on the list of test case Ids provided.
	RemoveTestCasesFromSuite(context.Context, RemoveTestCasesFromSuiteArgs) error
	// [Preview API] Reorder test suite entries in the test suite.
	ReorderSuiteEntries(context.Context, ReorderSuiteEntriesArgs) (*[]SuiteEntry, error)
	// [Preview API] Update the configurations for test cases
	UpdateSuiteTestCases(context.Context, UpdateSuiteTestCasesArgs) (*[]TestCase, error)
	// [Preview API] Update a test configuration by its ID.
	UpdateTestConfiguration(context.Context, UpdateTestConfigurationArgs) (*TestConfiguration, error)
	// [Preview API] Update a test plan.
	UpdateTestPlan(context.Context, UpdateTestPlanArgs) (*TestPlan, error)
	// [Preview API] Update Test Points. This is used to Reset test point to active, update the outcome of a test point or update the tester of a test point
	UpdateTestPoints(context.Context, UpdateTestPointsArgs) (*[]TestPoint, error)
	// [Preview API] Update test suite.
	UpdateTestSuite(context.Context, UpdateTestSuiteArgs) (*TestSuite, error)
	// [Preview API] Update a test variable by its ID.
	UpdateTestVariable(context.Context, UpdateTestVariableArgs) (*TestVariable, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) Client {
	client := connection.GetClientByUrl(connection.BaseUrl)
	return &ClientImpl{
		Client: *client,
	}
}

// [Preview API] Add test cases to a suite with specified configurations
func (client *ClientImpl) AddTestCasesToSuite(ctx context.Context, args AddTestCasesToSuiteArgs) (*[]TestCase, error) {
	if args.SuiteTestCaseCreateUpdateParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteTestCaseCreateUpdateParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	body, marshalErr := json.Marshal(*args.SuiteTestCaseCreateUpdateParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TestCase
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AddTestCasesToSuite function
type AddTestCasesToSuiteArgs struct {
	// (required) SuiteTestCaseCreateUpdateParameters object.
	SuiteTestCaseCreateUpdateParameters *[]SuiteTestCaseCreateUpdateParameters
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan to which test cases are to be added.
	PlanId *int
	// (required) ID of the test suite to which test cases are to be added.
	SuiteId *int
}

// [Preview API]
func (client *ClientImpl) CloneTestCase(ctx context.Context, args CloneTestCaseArgs) (*CloneTestCaseOperationInformation, error) {
	if args.CloneRequestBody == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CloneRequestBody"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.CloneRequestBody)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("529b2b8d-82f4-4893-b1e4-1e74ea534673")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CloneTestCaseOperationInformation
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CloneTestCase function
type CloneTestCaseArgs struct {
	// (required)
	CloneRequestBody *CloneTestCaseParams
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Clone test plan
func (client *ClientImpl) CloneTestPlan(ctx context.Context, args CloneTestPlanArgs) (*CloneTestPlanOperationInformation, error) {
	if args.CloneRequestBody == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CloneRequestBody"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.DeepClone != nil {
		queryParams.Add("deepClone", strconv.FormatBool(*args.DeepClone))
	}
	body, marshalErr := json.Marshal(*args.CloneRequestBody)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e65df662-d8a3-46c7-ae1c-14e2d4df57e1")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CloneTestPlanOperationInformation
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CloneTestPlan function
type CloneTestPlanArgs struct {
	// (required) Plan Clone Request Body detail TestPlanCloneRequest
	CloneRequestBody *CloneTestPlanParams
	// (required) Project ID or project name
	Project *string
	// (optional) Clones all the associated test cases as well
	DeepClone *bool
}

// [Preview API] Clone test suite
func (client *ClientImpl) CloneTestSuite(ctx context.Context, args CloneTestSuiteArgs) (*CloneTestSuiteOperationInformation, error) {
	if args.CloneRequestBody == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CloneRequestBody"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.DeepClone != nil {
		queryParams.Add("deepClone", strconv.FormatBool(*args.DeepClone))
	}
	body, marshalErr := json.Marshal(*args.CloneRequestBody)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("181d4c97-0e98-4ee2-ad6a-4cada675e555")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CloneTestSuiteOperationInformation
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CloneTestSuite function
type CloneTestSuiteArgs struct {
	// (required) Suite Clone Request Body detail TestSuiteCloneRequest
	CloneRequestBody *CloneTestSuiteParams
	// (required) Project ID or project name
	Project *string
	// (optional) Clones all the associated test cases as well
	DeepClone *bool
}

// [Preview API] Create a test configuration.
func (client *ClientImpl) CreateTestConfiguration(ctx context.Context, args CreateTestConfigurationArgs) (*TestConfiguration, error) {
	if args.TestConfigurationCreateUpdateParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestConfigurationCreateUpdateParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.TestConfigurationCreateUpdateParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateTestConfiguration function
type CreateTestConfigurationArgs struct {
	// (required) TestConfigurationCreateUpdateParameters
	TestConfigurationCreateUpdateParameters *TestConfigurationCreateUpdateParameters
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Create a test plan.
func (client *ClientImpl) CreateTestPlan(ctx context.Context, args CreateTestPlanArgs) (*TestPlan, error) {
	if args.TestPlanCreateParams == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestPlanCreateParams"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.TestPlanCreateParams)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestPlan
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateTestPlan function
type CreateTestPlanArgs struct {
	// (required) A testPlanCreateParams object.TestPlanCreateParams
	TestPlanCreateParams *TestPlanCreateParams
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Create test suite.
func (client *ClientImpl) CreateTestSuite(ctx context.Context, args CreateTestSuiteArgs) (*TestSuite, error) {
	if args.TestSuiteCreateParams == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestSuiteCreateParams"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)

	body, marshalErr := json.Marshal(*args.TestSuiteCreateParams)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestSuite
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateTestSuite function
type CreateTestSuiteArgs struct {
	// (required) Parameters for suite creation
	TestSuiteCreateParams *TestSuiteCreateParams
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan that contains the suites.
	PlanId *int
}

// [Preview API] Create a test variable.
func (client *ClientImpl) CreateTestVariable(ctx context.Context, args CreateTestVariableArgs) (*TestVariable, error) {
	if args.TestVariableCreateUpdateParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestVariableCreateUpdateParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.TestVariableCreateUpdateParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestVariable
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateTestVariable function
type CreateTestVariableArgs struct {
	// (required) TestVariableCreateUpdateParameters
	TestVariableCreateUpdateParameters *TestVariableCreateUpdateParameters
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Delete a test case.
func (client *ClientImpl) DeleteTestCase(ctx context.Context, args DeleteTestCaseArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TestCaseId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.TestCaseId"}
	}
	routeValues["testCaseId"] = strconv.Itoa(*args.TestCaseId)

	locationId, _ := uuid.Parse("29006fb5-816b-4ff7-a329-599943569229")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTestCase function
type DeleteTestCaseArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of test case to be deleted.
	TestCaseId *int
}

// [Preview API] Delete a test configuration by its ID.
func (client *ClientImpl) DeleteTestConfguration(ctx context.Context, args DeleteTestConfgurationArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.TestConfiguartionId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "testConfiguartionId"}
	}
	queryParams.Add("testConfiguartionId", strconv.Itoa(*args.TestConfiguartionId))
	locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTestConfguration function
type DeleteTestConfgurationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test configuration to delete.
	TestConfiguartionId *int
}

// [Preview API] Delete a test plan.
func (client *ClientImpl) DeleteTestPlan(ctx context.Context, args DeleteTestPlanArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)

	locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTestPlan function
type DeleteTestPlanArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan to be deleted.
	PlanId *int
}

// [Preview API] Delete test suite.
func (client *ClientImpl) DeleteTestSuite(ctx context.Context, args DeleteTestSuiteArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTestSuite function
type DeleteTestSuiteArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan that contains the suite.
	PlanId *int
	// (required) ID of the test suite to delete.
	SuiteId *int
}

// [Preview API] Delete a test variable by its ID.
func (client *ClientImpl) DeleteTestVariable(ctx context.Context, args DeleteTestVariableArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TestVariableId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.TestVariableId"}
	}
	routeValues["testVariableId"] = strconv.Itoa(*args.TestVariableId)

	locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTestVariable function
type DeleteTestVariableArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test variable to delete.
	TestVariableId *int
}

// [Preview API] Get clone information.
func (client *ClientImpl) GetCloneInformation(ctx context.Context, args GetCloneInformationArgs) (*CloneTestPlanOperationInformation, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.CloneOperationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CloneOperationId"}
	}
	routeValues["cloneOperationId"] = strconv.Itoa(*args.CloneOperationId)

	locationId, _ := uuid.Parse("e65df662-d8a3-46c7-ae1c-14e2d4df57e1")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CloneTestPlanOperationInformation
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCloneInformation function
type GetCloneInformationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Operation ID returned when we queue a clone operation
	CloneOperationId *int
}

// [Preview API] Get a list of points based on point Ids provided.
func (client *ClientImpl) GetPoints(ctx context.Context, args GetPointsArgs) (*[]TestPoint, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
	if args.PointIds == nil || *args.PointIds == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PointIds"}
	}
	routeValues["pointIds"] = *args.PointIds

	queryParams := url.Values{}
	if args.ReturnIdentityRef != nil {
		queryParams.Add("returnIdentityRef", strconv.FormatBool(*args.ReturnIdentityRef))
	}
	if args.IncludePointDetails != nil {
		queryParams.Add("includePointDetails", strconv.FormatBool(*args.IncludePointDetails))
	}
	locationId, _ := uuid.Parse("52df686e-bae4-4334-b0ee-b6cf4e6f6b73")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TestPoint
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPoints function
type GetPointsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan for which test points are requested.
	PlanId *int
	// (required) ID of the test suite for which test points are requested.
	SuiteId *int
	// (required) ID of test points to be fetched.
	PointIds *string
	// (optional) If set to true, returns the AssignedTo field in TestCaseReference as IdentityRef object.
	ReturnIdentityRef *bool
	// (optional) If set to false, will get a smaller payload containing only basic details about the test point object
	IncludePointDetails *bool
}

// [Preview API] Get all the points inside a suite based on some filters
func (client *ClientImpl) GetPointsList(ctx context.Context, args GetPointsListArgs) (*GetPointsListResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	queryParams := url.Values{}
	if args.TestPointIds != nil {
		queryParams.Add("testPointIds", *args.TestPointIds)
	}
	if args.TestCaseId != nil {
		queryParams.Add("testCaseId", *args.TestCaseId)
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.ReturnIdentityRef != nil {
		queryParams.Add("returnIdentityRef", strconv.FormatBool(*args.ReturnIdentityRef))
	}
	if args.IncludePointDetails != nil {
		queryParams.Add("includePointDetails", strconv.FormatBool(*args.IncludePointDetails))
	}
	if args.IsRecursive != nil {
		queryParams.Add("isRecursive", strconv.FormatBool(*args.IsRecursive))
	}
	locationId, _ := uuid.Parse("52df686e-bae4-4334-b0ee-b6cf4e6f6b73")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetPointsListResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetPointsList function
type GetPointsListArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan for which test points are requested.
	PlanId *int
	// (required) ID of the test suite for which test points are requested
	SuiteId *int
	// (optional) ID of test points to fetch.
	TestPointIds *string
	// (optional) Get Test Points for specific test case Ids.
	TestCaseId *string
	// (optional) If the list of test point returned is not complete, a continuation token to query next batch of test points is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test points.
	ContinuationToken *string
	// (optional) If set to true, returns the AssignedTo field in TestCaseReference as IdentityRef object.
	ReturnIdentityRef *bool
	// (optional) If set to false, will get a smaller payload containing only basic details about the test point object
	IncludePointDetails *bool
	// (optional) If set to true, will also fetch test points belonging to child suites recursively.
	IsRecursive *bool
}

// Return type for the GetPointsList function
type GetPointsListResponseValue struct {
	Value             []TestPoint
	ContinuationToken string
}

// [Preview API] Get clone information.
func (client *ClientImpl) GetSuiteCloneInformation(ctx context.Context, args GetSuiteCloneInformationArgs) (*CloneTestSuiteOperationInformation, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.CloneOperationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CloneOperationId"}
	}
	routeValues["cloneOperationId"] = strconv.Itoa(*args.CloneOperationId)

	locationId, _ := uuid.Parse("181d4c97-0e98-4ee2-ad6a-4cada675e555")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CloneTestSuiteOperationInformation
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetSuiteCloneInformation function
type GetSuiteCloneInformationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Operation ID returned when we queue a clone operation
	CloneOperationId *int
}

// [Preview API] Get a list of test suite entries in the test suite.
func (client *ClientImpl) GetSuiteEntries(ctx context.Context, args GetSuiteEntriesArgs) (*[]SuiteEntry, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	queryParams := url.Values{}
	if args.SuiteEntryType != nil {
		queryParams.Add("suiteEntryType", string(*args.SuiteEntryType))
	}
	locationId, _ := uuid.Parse("d6733edf-72f1-4252-925b-c560dfe9b75a")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []SuiteEntry
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetSuiteEntries function
type GetSuiteEntriesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Id of the parent suite.
	SuiteId *int
	// (optional)
	SuiteEntryType *SuiteEntryTypes
}

// Find the list of all test suites in which a given test case is present. This is helpful if you need to find out which test suites are using a test case, when you need to make changes to a test case.
func (client *ClientImpl) GetSuitesByTestCaseId(ctx context.Context, args GetSuitesByTestCaseIdArgs) (*[]TestSuite, error) {
	queryParams := url.Values{}
	if args.TestCaseId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "testCaseId"}
	}
	queryParams.Add("testCaseId", strconv.Itoa(*args.TestCaseId))
	locationId, _ := uuid.Parse("a4080e84-f17b-4fad-84f1-7960b6525bf2")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TestSuite
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetSuitesByTestCaseId function
type GetSuitesByTestCaseIdArgs struct {
	// (required) ID of the test case for which suites need to be fetched.
	TestCaseId *int
}

// [Preview API] Get Test Cases For a Suite.
func (client *ClientImpl) GetTestCase(ctx context.Context, args GetTestCaseArgs) (*[]TestCase, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
	if args.TestCaseIds == nil || *args.TestCaseIds == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TestCaseIds"}
	}
	routeValues["testCaseIds"] = *args.TestCaseIds

	queryParams := url.Values{}
	if args.WitFields != nil {
		queryParams.Add("witFields", *args.WitFields)
	}
	if args.ReturnIdentityRef != nil {
		queryParams.Add("returnIdentityRef", strconv.FormatBool(*args.ReturnIdentityRef))
	}
	locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TestCase
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTestCase function
type GetTestCaseArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan for which test cases are requested.
	PlanId *int
	// (required) ID of the test suite for which test cases are requested.
	SuiteId *int
	// (required) Test Case Ids to be fetched.
	TestCaseIds *string
	// (optional) Get the list of witFields.
	WitFields *string
	// (optional) If set to true, returns all identity fields, like AssignedTo, ActivatedBy etc., as IdentityRef objects. If set to false, these fields are returned as unique names in string format. This is false by default.
	ReturnIdentityRef *bool
}

// [Preview API] Get clone information.
func (client *ClientImpl) GetTestCaseCloneInformation(ctx context.Context, args GetTestCaseCloneInformationArgs) (*CloneTestCaseOperationInformation, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.CloneOperationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CloneOperationId"}
	}
	routeValues["cloneOperationId"] = strconv.Itoa(*args.CloneOperationId)

	locationId, _ := uuid.Parse("529b2b8d-82f4-4893-b1e4-1e74ea534673")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CloneTestCaseOperationInformation
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTestCaseCloneInformation function
type GetTestCaseCloneInformationArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Operation ID returned when we queue a clone operation
	CloneOperationId *int
}

// [Preview API] Get Test Case List return those test cases which have all the configuration Ids as mentioned in the optional parameter. If configuration Ids is null, it return all the test cases
func (client *ClientImpl) GetTestCaseList(ctx context.Context, args GetTestCaseListArgs) (*GetTestCaseListResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	queryParams := url.Values{}
	if args.TestIds != nil {
		queryParams.Add("testIds", *args.TestIds)
	}
	if args.ConfigurationIds != nil {
		queryParams.Add("configurationIds", *args.ConfigurationIds)
	}
	if args.WitFields != nil {
		queryParams.Add("witFields", *args.WitFields)
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.ReturnIdentityRef != nil {
		queryParams.Add("returnIdentityRef", strconv.FormatBool(*args.ReturnIdentityRef))
	}
	if args.Expand != nil {
		queryParams.Add("expand", strconv.FormatBool(*args.Expand))
	}
	if args.ExcludeFlags != nil {
		queryParams.Add("excludeFlags", string(*args.ExcludeFlags))
	}
	locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestCaseListResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestCaseList function
type GetTestCaseListArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan for which test cases are requested.
	PlanId *int
	// (required) ID of the test suite for which test cases are requested.
	SuiteId *int
	// (optional) Test Case Ids to be fetched.
	TestIds *string
	// (optional) Fetch Test Cases which contains all the configuration Ids specified.
	ConfigurationIds *string
	// (optional) Get the list of witFields.
	WitFields *string
	// (optional) If the list of test cases returned is not complete, a continuation token to query next batch of test cases is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test cases.
	ContinuationToken *string
	// (optional) If set to true, returns all identity fields, like AssignedTo, ActivatedBy etc., as IdentityRef objects. If set to false, these fields are returned as unique names in string format. This is false by default.
	ReturnIdentityRef *bool
	// (optional) If set to false, will get a smaller payload containing only basic details about the suite test case object
	Expand *bool
	// (optional) Flag to exclude various values from payload. For example to remove point assignments pass exclude = 1. To remove extra information (links, test plan , test suite) pass exclude = 2. To remove both extra information and point assignments pass exclude = 3 (1 + 2).
	ExcludeFlags *ExcludeFlags
}

// Return type for the GetTestCaseList function
type GetTestCaseListResponseValue struct {
	Value             []TestCase
	ContinuationToken string
}

// [Preview API] Get a test configuration
func (client *ClientImpl) GetTestConfigurationById(ctx context.Context, args GetTestConfigurationByIdArgs) (*TestConfiguration, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TestConfigurationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestConfigurationId"}
	}
	routeValues["testConfigurationId"] = strconv.Itoa(*args.TestConfigurationId)

	locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTestConfigurationById function
type GetTestConfigurationByIdArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test configuration to get.
	TestConfigurationId *int
}

// [Preview API] Get a list of test configurations.
func (client *ClientImpl) GetTestConfigurations(ctx context.Context, args GetTestConfigurationsArgs) (*GetTestConfigurationsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestConfigurationsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestConfigurations function
type GetTestConfigurationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) If the list of configurations returned is not complete, a continuation token to query next batch of configurations is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test configurations.
	ContinuationToken *string
}

// Return type for the GetTestConfigurations function
type GetTestConfigurationsResponseValue struct {
	Value             []TestConfiguration
	ContinuationToken string
}

// [Preview API] Get a test plan by Id.
func (client *ClientImpl) GetTestPlanById(ctx context.Context, args GetTestPlanByIdArgs) (*TestPlan, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)

	locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestPlan
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTestPlanById function
type GetTestPlanByIdArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan to get.
	PlanId *int
}

// [Preview API] Get a list of test plans
func (client *ClientImpl) GetTestPlans(ctx context.Context, args GetTestPlansArgs) (*GetTestPlansResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.Owner != nil {
		queryParams.Add("owner", *args.Owner)
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.IncludePlanDetails != nil {
		queryParams.Add("includePlanDetails", strconv.FormatBool(*args.IncludePlanDetails))
	}
	if args.FilterActivePlans != nil {
		queryParams.Add("filterActivePlans", strconv.FormatBool(*args.FilterActivePlans))
	}
	locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestPlansResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestPlans function
type GetTestPlansArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) Filter for test plan by owner ID or name
	Owner *string
	// (optional) If the list of plans returned is not complete, a continuation token to query next batch of plans is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test plans.
	ContinuationToken *string
	// (optional) Get all properties of the test plan
	IncludePlanDetails *bool
	// (optional) Get just the active plans
	FilterActivePlans *bool
}

// Return type for the GetTestPlans function
type GetTestPlansResponseValue struct {
	Value             []TestPlan
	ContinuationToken string
}

// [Preview API] Get test suite by suite id.
func (client *ClientImpl) GetTestSuiteById(ctx context.Context, args GetTestSuiteByIdArgs) (*TestSuite, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("expand", string(*args.Expand))
	}
	locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestSuite
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTestSuiteById function
type GetTestSuiteByIdArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan that contains the suites.
	PlanId *int
	// (required) ID of the suite to get.
	SuiteId *int
	// (optional) Include the children suites and testers details
	Expand *SuiteExpand
}

// [Preview API] Get test suites for plan.
func (client *ClientImpl) GetTestSuitesForPlan(ctx context.Context, args GetTestSuitesForPlanArgs) (*GetTestSuitesForPlanResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)

	queryParams := url.Values{}
	if args.Expand != nil {
		queryParams.Add("expand", string(*args.Expand))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.AsTreeView != nil {
		queryParams.Add("asTreeView", strconv.FormatBool(*args.AsTreeView))
	}
	locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestSuitesForPlanResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestSuitesForPlan function
type GetTestSuitesForPlanArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan for which suites are requested.
	PlanId *int
	// (optional) Include the children suites and testers details.
	Expand *SuiteExpand
	// (optional) If the list of suites returned is not complete, a continuation token to query next batch of suites is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test suites.
	ContinuationToken *string
	// (optional) If the suites returned should be in a tree structure.
	AsTreeView *bool
}

// Return type for the GetTestSuitesForPlan function
type GetTestSuitesForPlanResponseValue struct {
	Value             []TestSuite
	ContinuationToken string
}

// [Preview API] Get a test variable by its ID.
func (client *ClientImpl) GetTestVariableById(ctx context.Context, args GetTestVariableByIdArgs) (*TestVariable, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TestVariableId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestVariableId"}
	}
	routeValues["testVariableId"] = strconv.Itoa(*args.TestVariableId)

	locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestVariable
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTestVariableById function
type GetTestVariableByIdArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test variable to get.
	TestVariableId *int
}

// [Preview API] Get a list of test variables.
func (client *ClientImpl) GetTestVariables(ctx context.Context, args GetTestVariablesArgs) (*GetTestVariablesResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetTestVariablesResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetTestVariables function
type GetTestVariablesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) If the list of variables returned is not complete, a continuation token to query next batch of variables is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test variables.
	ContinuationToken *string
}

// Return type for the GetTestVariables function
type GetTestVariablesResponseValue struct {
	Value             []TestVariable
	ContinuationToken string
}

// [Preview API] Removes test cases from a suite based on the list of test case Ids provided.
func (client *ClientImpl) RemoveTestCasesFromSuite(ctx context.Context, args RemoveTestCasesFromSuiteArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
	if args.TestCaseIds == nil || *args.TestCaseIds == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TestCaseIds"}
	}
	routeValues["testCaseIds"] = *args.TestCaseIds

	locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "6.0-preview.2", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RemoveTestCasesFromSuite function
type RemoveTestCasesFromSuiteArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan from which test cases are to be removed.
	PlanId *int
	// (required) ID of the test suite from which test cases are to be removed.
	SuiteId *int
	// (required) Test Case Ids to be removed.
	TestCaseIds *string
}

// [Preview API] Reorder test suite entries in the test suite.
func (client *ClientImpl) ReorderSuiteEntries(ctx context.Context, args ReorderSuiteEntriesArgs) (*[]SuiteEntry, error) {
	if args.SuiteEntries == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteEntries"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	body, marshalErr := json.Marshal(*args.SuiteEntries)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("d6733edf-72f1-4252-925b-c560dfe9b75a")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []SuiteEntry
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the ReorderSuiteEntries function
type ReorderSuiteEntriesArgs struct {
	// (required) List of SuiteEntry to reorder.
	SuiteEntries *[]SuiteEntryUpdateParams
	// (required) Project ID or project name
	Project *string
	// (required) Id of the parent test suite.
	SuiteId *int
}

// [Preview API] Update the configurations for test cases
func (client *ClientImpl) UpdateSuiteTestCases(ctx context.Context, args UpdateSuiteTestCasesArgs) (*[]TestCase, error) {
	if args.SuiteTestCaseCreateUpdateParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteTestCaseCreateUpdateParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	body, marshalErr := json.Marshal(*args.SuiteTestCaseCreateUpdateParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TestCase
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateSuiteTestCases function
type UpdateSuiteTestCasesArgs struct {
	// (required) A SuiteTestCaseCreateUpdateParameters object.
	SuiteTestCaseCreateUpdateParameters *[]SuiteTestCaseCreateUpdateParameters
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan to which test cases are to be updated.
	PlanId *int
	// (required) ID of the test suite to which test cases are to be updated.
	SuiteId *int
}

// [Preview API] Update a test configuration by its ID.
func (client *ClientImpl) UpdateTestConfiguration(ctx context.Context, args UpdateTestConfigurationArgs) (*TestConfiguration, error) {
	if args.TestConfigurationCreateUpdateParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestConfigurationCreateUpdateParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.TestConfiguartionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "testConfiguartionId"}
	}
	queryParams.Add("testConfiguartionId", strconv.Itoa(*args.TestConfiguartionId))
	body, marshalErr := json.Marshal(*args.TestConfigurationCreateUpdateParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestConfiguration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateTestConfiguration function
type UpdateTestConfigurationArgs struct {
	// (required) TestConfigurationCreateUpdateParameters
	TestConfigurationCreateUpdateParameters *TestConfigurationCreateUpdateParameters
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test configuration to update.
	TestConfiguartionId *int
}

// [Preview API] Update a test plan.
func (client *ClientImpl) UpdateTestPlan(ctx context.Context, args UpdateTestPlanArgs) (*TestPlan, error) {
	if args.TestPlanUpdateParams == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestPlanUpdateParams"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)

	body, marshalErr := json.Marshal(*args.TestPlanUpdateParams)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestPlan
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateTestPlan function
type UpdateTestPlanArgs struct {
	// (required) A testPlanUpdateParams object.TestPlanUpdateParams
	TestPlanUpdateParams *TestPlanUpdateParams
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan to be updated.
	PlanId *int
}

// [Preview API] Update Test Points. This is used to Reset test point to active, update the outcome of a test point or update the tester of a test point
func (client *ClientImpl) UpdateTestPoints(ctx context.Context, args UpdateTestPointsArgs) (*[]TestPoint, error) {
	if args.TestPointUpdateParams == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestPointUpdateParams"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	queryParams := url.Values{}
	if args.IncludePointDetails != nil {
		queryParams.Add("includePointDetails", strconv.FormatBool(*args.IncludePointDetails))
	}
	if args.ReturnIdentityRef != nil {
		queryParams.Add("returnIdentityRef", strconv.FormatBool(*args.ReturnIdentityRef))
	}
	body, marshalErr := json.Marshal(*args.TestPointUpdateParams)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("52df686e-bae4-4334-b0ee-b6cf4e6f6b73")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TestPoint
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateTestPoints function
type UpdateTestPointsArgs struct {
	// (required) A TestPointUpdateParams Object.
	TestPointUpdateParams *[]TestPointUpdateParams
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan for which test points are requested.
	PlanId *int
	// (required) ID of the test suite for which test points are requested.
	SuiteId *int
	// (optional) If set to false, will get a smaller payload containing only basic details about the test point object
	IncludePointDetails *bool
	// (optional) If set to true, returns the AssignedTo field in TestCaseReference as IdentityRef object.
	ReturnIdentityRef *bool
}

// [Preview API] Update test suite.
func (client *ClientImpl) UpdateTestSuite(ctx context.Context, args UpdateTestSuiteArgs) (*TestSuite, error) {
	if args.TestSuiteUpdateParams == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestSuiteUpdateParams"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = strconv.Itoa(*args.PlanId)
	if args.SuiteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SuiteId"}
	}
	routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

	body, marshalErr := json.Marshal(*args.TestSuiteUpdateParams)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestSuite
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateTestSuite function
type UpdateTestSuiteArgs struct {
	// (required) Parameters for suite updation
	TestSuiteUpdateParams *TestSuiteUpdateParams
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test plan that contains the suites.
	PlanId *int
	// (required) ID of the parent suite.
	SuiteId *int
}

// [Preview API] Update a test variable by its ID.
func (client *ClientImpl) UpdateTestVariable(ctx context.Context, args UpdateTestVariableArgs) (*TestVariable, error) {
	if args.TestVariableCreateUpdateParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestVariableCreateUpdateParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.TestVariableId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TestVariableId"}
	}
	routeValues["testVariableId"] = strconv.Itoa(*args.TestVariableId)

	body, marshalErr := json.Marshal(*args.TestVariableCreateUpdateParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TestVariable
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateTestVariable function
type UpdateTestVariableArgs struct {
	// (required) TestVariableCreateUpdateParameters
	TestVariableCreateUpdateParameters *TestVariableCreateUpdateParameters
	// (required) Project ID or project name
	Project *string
	// (required) ID of the test variable to update.
	TestVariableId *int
}
