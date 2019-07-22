// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testPlan

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
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

// [Preview API] Create a test configuration.
// ctx
// testConfigurationCreateUpdateParameters (required): TestConfigurationCreateUpdateParameters
// project (required): Project ID or project name
func (client Client) CreateTestConfiguration(ctx context.Context, testConfigurationCreateUpdateParameters *TestConfigurationCreateUpdateParameters, project *string) (*TestConfiguration, error) {
    if testConfigurationCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfigurationCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testConfigurationCreateUpdateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a test configuration by its ID.
// ctx
// project (required): Project ID or project name
// testConfiguartionId (required): ID of the test configuration to delete.
func (client Client) DeleteTestConfguration(ctx context.Context, project *string, testConfiguartionId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testConfiguartionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testConfiguartionId"}
    }
    queryParams.Add("testConfiguartionId", strconv.Itoa(*testConfiguartionId))
    locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a test configuration
// ctx
// project (required): Project ID or project name
// testConfigurationId (required): ID of the test configuration to get.
func (client Client) GetTestConfigurationById(ctx context.Context, project *string, testConfigurationId *int) (*TestConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if testConfigurationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfigurationId"} 
    }
    routeValues["testConfigurationId"] = strconv.Itoa(*testConfigurationId)

    locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of test configurations.
// ctx
// project (required): Project ID or project name
// continuationToken (optional): If the list of configurations returned is not complete, a continuation token to query next batch of configurations is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test configurations.
func (client Client) GetTestConfigurations(ctx context.Context, project *string, continuationToken *string) (*[]TestConfiguration, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestConfiguration
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a test configuration by its ID.
// ctx
// testConfigurationCreateUpdateParameters (required): TestConfigurationCreateUpdateParameters
// project (required): Project ID or project name
// testConfiguartionId (required): ID of the test configuration to update.
func (client Client) UpdateTestConfiguration(ctx context.Context, testConfigurationCreateUpdateParameters *TestConfigurationCreateUpdateParameters, project *string, testConfiguartionId *int) (*TestConfiguration, error) {
    if testConfigurationCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfigurationCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if testConfiguartionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfiguartionId"}
    }
    queryParams.Add("testConfiguartionId", strconv.Itoa(*testConfiguartionId))
    body, marshalErr := json.Marshal(*testConfigurationCreateUpdateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestConfiguration
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a test plan.
// ctx
// testPlanCreateParams (required): A testPlanCreateParams object.TestPlanCreateParams
// project (required): Project ID or project name
func (client Client) CreateTestPlan(ctx context.Context, testPlanCreateParams *TestPlanCreateParams, project *string) (*TestPlan, error) {
    if testPlanCreateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testPlanCreateParams"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testPlanCreateParams)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestPlan
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a test plan.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan to be deleted.
func (client Client) DeleteTestPlan(ctx context.Context, project *string, planId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)

    locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a test plan by Id.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan to get.
func (client Client) GetTestPlanById(ctx context.Context, project *string, planId *int) (*TestPlan, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)

    locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestPlan
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of test plans
// ctx
// project (required): Project ID or project name
// owner (optional): Filter for test plan by owner ID or name
// continuationToken (optional): If the list of plans returned is not complete, a continuation token to query next batch of plans is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test plans.
// includePlanDetails (optional): Get all properties of the test plan
// filterActivePlans (optional): Get just the active plans
func (client Client) GetTestPlans(ctx context.Context, project *string, owner *string, continuationToken *string, includePlanDetails *bool, filterActivePlans *bool) (*[]TestPlan, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if owner != nil {
        queryParams.Add("owner", *owner)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if includePlanDetails != nil {
        queryParams.Add("includePlanDetails", strconv.FormatBool(*includePlanDetails))
    }
    if filterActivePlans != nil {
        queryParams.Add("filterActivePlans", strconv.FormatBool(*filterActivePlans))
    }
    locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPlan
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a test plan.
// ctx
// testPlanUpdateParams (required): A testPlanUpdateParams object.TestPlanUpdateParams
// project (required): Project ID or project name
// planId (required): ID of the test plan to be updated.
func (client Client) UpdateTestPlan(ctx context.Context, testPlanUpdateParams *TestPlanUpdateParams, project *string, planId *int) (*TestPlan, error) {
    if testPlanUpdateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testPlanUpdateParams"}
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

    body, marshalErr := json.Marshal(*testPlanUpdateParams)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestPlan
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of test suite entries in the test suite.
// ctx
// project (required): Project ID or project name
// suiteId (required): Id of the parent suite.
// suiteEntryType (optional)
func (client Client) GetSuiteEntries(ctx context.Context, project *string, suiteId *int, suiteEntryType *SuiteEntryTypes) (*[]SuiteEntry, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)

    queryParams := url.Values{}
    if suiteEntryType != nil {
        queryParams.Add("suiteEntryType", *suiteEntryType)
    }
    locationId, _ := uuid.Parse("d6733edf-72f1-4252-925b-c560dfe9b75a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteEntry
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Reorder test suite entries in the test suite.
// ctx
// suiteEntries (required): List of SuiteEntry to reorder.
// project (required): Project ID or project name
// suiteId (required): Id of the parent test suite.
func (client Client) ReorderSuiteEntries(ctx context.Context, suiteEntries *[]SuiteEntryUpdateParams, project *string, suiteId *int) (*[]SuiteEntry, error) {
    if suiteEntries == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteEntries"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if suiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*suiteId)

    body, marshalErr := json.Marshal(*suiteEntries)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d6733edf-72f1-4252-925b-c560dfe9b75a")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteEntry
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create test suite.
// ctx
// testSuiteCreateParams (required): Parameters for suite creation
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suites.
func (client Client) CreateTestSuite(ctx context.Context, testSuiteCreateParams *TestSuiteCreateParams, project *string, planId *int) (*TestSuite, error) {
    if testSuiteCreateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSuiteCreateParams"}
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

    body, marshalErr := json.Marshal(*testSuiteCreateParams)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSuite
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete test suite.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suite.
// suiteId (required): ID of the test suite to delete.
func (client Client) DeleteTestSuite(ctx context.Context, project *string, planId *int, suiteId *int) error {
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

    locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get test suite by suite id.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suites.
// suiteId (required): ID of the suite to get.
// expand (optional): Include the children suites and testers details
func (client Client) GetTestSuiteById(ctx context.Context, project *string, planId *int, suiteId *int, expand *SuiteExpand) (*TestSuite, error) {
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
    if expand != nil {
        queryParams.Add("expand", *expand)
    }
    locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSuite
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get test suites for plan.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan for which suites are requested.
// expand (optional): Include the children suites and testers details.
// continuationToken (optional): If the list of suites returned is not complete, a continuation token to query next batch of suites is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test suites.
// asTreeView (optional): If the suites returned should be in a tree structure.
func (client Client) GetTestSuitesForPlan(ctx context.Context, project *string, planId *int, expand *SuiteExpand, continuationToken *string, asTreeView *bool) (*[]TestSuite, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*planId)

    queryParams := url.Values{}
    if expand != nil {
        queryParams.Add("expand", *expand)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if asTreeView != nil {
        queryParams.Add("asTreeView", strconv.FormatBool(*asTreeView))
    }
    locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestSuite
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update test suite.
// ctx
// testSuiteUpdateParams (required): Parameters for suite updation
// project (required): Project ID or project name
// planId (required): ID of the test plan that contains the suites.
// suiteId (required): ID of the parent suite.
func (client Client) UpdateTestSuite(ctx context.Context, testSuiteUpdateParams *TestSuiteUpdateParams, project *string, planId *int, suiteId *int) (*TestSuite, error) {
    if testSuiteUpdateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSuiteUpdateParams"}
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

    body, marshalErr := json.Marshal(*testSuiteUpdateParams)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSuite
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Find the list of all test suites in which a given test case is present. This is helpful if you need to find out which test suites are using a test case, when you need to make changes to a test case.
// ctx
// testCaseId (required): ID of the test case for which suites need to be fetched.
func (client Client) GetSuitesByTestCaseId(ctx context.Context, testCaseId *int) (*[]TestSuite, error) {
    queryParams := url.Values{}
    if testCaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseId"}
    }
    queryParams.Add("testCaseId", strconv.Itoa(*testCaseId))
    locationId, _ := uuid.Parse("a4080e84-f17b-4fad-84f1-7960b6525bf2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestSuite
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Add test cases to a suite with specified configurations
// ctx
// suiteTestCaseCreateUpdateParameters (required): SuiteTestCaseCreateUpdateParameters object.
// project (required): Project ID or project name
// planId (required): ID of the test plan to which test cases are to be added.
// suiteId (required): ID of the test suite to which test cases are to be added.
func (client Client) AddTestCasesToSuite(ctx context.Context, suiteTestCaseCreateUpdateParameters *[]SuiteTestCaseCreateUpdateParameters, project *string, planId *int, suiteId *int) (*[]TestCase, error) {
    if suiteTestCaseCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteTestCaseCreateUpdateParameters"}
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

    body, marshalErr := json.Marshal(*suiteTestCaseCreateUpdateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get Test Cases For a Suite.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan for which test cases are requested.
// suiteId (required): ID of the test suite for which test cases are requested.
// testCaseIds (required): Test Case Ids to be fetched.
// witFields (optional): Get the list of witFields.
// returnIdentityRef (optional): If set to true, returns all identity fields, like AssignedTo, ActivatedBy etc., as IdentityRef objects. If set to false, these fields are returned as unique names in string format. This is false by default.
func (client Client) GetTestCase(ctx context.Context, project *string, planId *int, suiteId *int, testCaseIds *string, witFields *string, returnIdentityRef *bool) (*[]TestCase, error) {
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

    queryParams := url.Values{}
    if witFields != nil {
        queryParams.Add("witFields", *witFields)
    }
    if returnIdentityRef != nil {
        queryParams.Add("returnIdentityRef", strconv.FormatBool(*returnIdentityRef))
    }
    locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get Test Case List return those test cases which have all the configuration Ids as mentioned in the optional paramter. If configuration Ids is null, it return all the test cases
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan for which test cases are requested.
// suiteId (required): ID of the test suite for which test cases are requested.
// testIds (optional): Test Case Ids to be fetched.
// configurationIds (optional): Fetch Test Cases which contains all the configuration Ids specified.
// witFields (optional): Get the list of witFields.
// continuationToken (optional): If the list of test cases returned is not complete, a continuation token to query next batch of test cases is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test cases.
// returnIdentityRef (optional): If set to true, returns all identity fields, like AssignedTo, ActivatedBy etc., as IdentityRef objects. If set to false, these fields are returned as unique names in string format. This is false by default.
// expand (optional): If set to false, will get a smaller payload containing only basic details about the suite test case object
func (client Client) GetTestCaseList(ctx context.Context, project *string, planId *int, suiteId *int, testIds *string, configurationIds *string, witFields *string, continuationToken *string, returnIdentityRef *bool, expand *bool) (*[]TestCase, error) {
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
    if testIds != nil {
        queryParams.Add("testIds", *testIds)
    }
    if configurationIds != nil {
        queryParams.Add("configurationIds", *configurationIds)
    }
    if witFields != nil {
        queryParams.Add("witFields", *witFields)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if returnIdentityRef != nil {
        queryParams.Add("returnIdentityRef", strconv.FormatBool(*returnIdentityRef))
    }
    if expand != nil {
        queryParams.Add("expand", strconv.FormatBool(*expand))
    }
    locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Removes test cases from a suite based on the list of test case Ids provided.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan from which test cases are to be removed.
// suiteId (required): ID of the test suite from which test cases are to be removed.
// testCaseIds (required): Test Case Ids to be removed.
func (client Client) RemoveTestCasesFromSuite(ctx context.Context, project *string, planId *int, suiteId *int, testCaseIds *string) error {
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

    locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Update the configurations for test cases
// ctx
// suiteTestCaseCreateUpdateParameters (required): A SuiteTestCaseCreateUpdateParameters object.
// project (required): Project ID or project name
// planId (required): ID of the test plan to which test cases are to be updated.
// suiteId (required): ID of the test suite to which test cases are to be updated.
func (client Client) UpdateSuiteTestCases(ctx context.Context, suiteTestCaseCreateUpdateParameters *[]SuiteTestCaseCreateUpdateParameters, project *string, planId *int, suiteId *int) (*[]TestCase, error) {
    if suiteTestCaseCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteTestCaseCreateUpdateParameters"}
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

    body, marshalErr := json.Marshal(*suiteTestCaseCreateUpdateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a test case.
// ctx
// project (required): Project ID or project name
// testCaseId (required): Id of test case to be deleted.
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

    locationId, _ := uuid.Parse("29006fb5-816b-4ff7-a329-599943569229")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Clone test plan
// ctx
// cloneRequestBody (required): Plan Clone Request Body detail TestPlanCloneRequest
// project (required): Project ID or project name
// deepClone (optional): Clones all the associated test cases as well
func (client Client) CloneTestPlan(ctx context.Context, cloneRequestBody *CloneTestPlanParams, project *string, deepClone *bool) (*CloneTestPlanOperationInformation, error) {
    if cloneRequestBody == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneRequestBody"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if deepClone != nil {
        queryParams.Add("deepClone", strconv.FormatBool(*deepClone))
    }
    body, marshalErr := json.Marshal(*cloneRequestBody)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e65df662-d8a3-46c7-ae1c-14e2d4df57e1")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CloneTestPlanOperationInformation
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get clone information.
// ctx
// project (required): Project ID or project name
// cloneOperationId (required): Operation ID returned when we queue a clone operation
func (client Client) GetCloneInformation(ctx context.Context, project *string, cloneOperationId *int) (*CloneTestPlanOperationInformation, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if cloneOperationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneOperationId"} 
    }
    routeValues["cloneOperationId"] = strconv.Itoa(*cloneOperationId)

    locationId, _ := uuid.Parse("e65df662-d8a3-46c7-ae1c-14e2d4df57e1")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CloneTestPlanOperationInformation
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of points based on point Ids provided.
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan for which test points are requested.
// suiteId (required): ID of the test suite for which test points are requested.
// pointIds (required): ID of test points to be fetched.
// returnIdentityRef (optional): If set to true, returns the AssignedTo field in TestCaseReference as IdentityRef object.
func (client Client) GetPoints(ctx context.Context, project *string, planId *int, suiteId *int, pointIds *string, returnIdentityRef *bool) (*[]TestPoint, error) {
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

    queryParams := url.Values{}
    if returnIdentityRef != nil {
        queryParams.Add("returnIdentityRef", strconv.FormatBool(*returnIdentityRef))
    }
    locationId, _ := uuid.Parse("52df686e-bae4-4334-b0ee-b6cf4e6f6b73")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all the points inside a suite based on some filters
// ctx
// project (required): Project ID or project name
// planId (required): ID of the test plan for which test points are requested.
// suiteId (required): ID of the test suite for which test points are requested
// testPointIds (optional): ID of test points to fetch.
// testCaseId (optional): Get Test Points for specific test case Ids.
// continuationToken (optional): If the list of test point returned is not complete, a continuation token to query next batch of test points is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test points.
// returnIdentityRef (optional): If set to true, returns the AssignedTo field in TestCaseReference as IdentityRef object.
// includePointDetails (optional): If set to false, returns only necessary information
func (client Client) GetPointsList(ctx context.Context, project *string, planId *int, suiteId *int, testPointIds *string, testCaseId *string, continuationToken *string, returnIdentityRef *bool, includePointDetails *bool) (*[]TestPoint, error) {
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
    if testPointIds != nil {
        queryParams.Add("testPointIds", *testPointIds)
    }
    if testCaseId != nil {
        queryParams.Add("testCaseId", *testCaseId)
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    if returnIdentityRef != nil {
        queryParams.Add("returnIdentityRef", strconv.FormatBool(*returnIdentityRef))
    }
    if includePointDetails != nil {
        queryParams.Add("includePointDetails", strconv.FormatBool(*includePointDetails))
    }
    locationId, _ := uuid.Parse("52df686e-bae4-4334-b0ee-b6cf4e6f6b73")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update Test Points. This is used to Reset test point to active, update the outcome of a test point or update the tester of a test point
// ctx
// testPointUpdateParams (required): A TestPointUpdateParams Object.
// project (required): Project ID or project name
// planId (required): ID of the test plan for which test points are requested.
// suiteId (required): ID of the test suite for which test points are requested.
func (client Client) UpdateTestPoints(ctx context.Context, testPointUpdateParams *[]TestPointUpdateParams, project *string, planId *int, suiteId *int) (*[]TestPoint, error) {
    if testPointUpdateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testPointUpdateParams"}
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

    body, marshalErr := json.Marshal(*testPointUpdateParams)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("52df686e-bae4-4334-b0ee-b6cf4e6f6b73")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Clone test suite
// ctx
// cloneRequestBody (required): Suite Clone Request Body detail TestSuiteCloneRequest
// project (required): Project ID or project name
// deepClone (optional): Clones all the associated test cases as well
func (client Client) CloneTestSuite(ctx context.Context, cloneRequestBody *CloneTestSuiteParams, project *string, deepClone *bool) (*CloneTestSuiteOperationInformation, error) {
    if cloneRequestBody == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneRequestBody"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if deepClone != nil {
        queryParams.Add("deepClone", strconv.FormatBool(*deepClone))
    }
    body, marshalErr := json.Marshal(*cloneRequestBody)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("181d4c97-0e98-4ee2-ad6a-4cada675e555")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CloneTestSuiteOperationInformation
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get clone information.
// ctx
// project (required): Project ID or project name
// cloneOperationId (required): Operation ID returned when we queue a clone operation
func (client Client) GetSuiteCloneInformation(ctx context.Context, project *string, cloneOperationId *int) (*CloneTestSuiteOperationInformation, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if cloneOperationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneOperationId"} 
    }
    routeValues["cloneOperationId"] = strconv.Itoa(*cloneOperationId)

    locationId, _ := uuid.Parse("181d4c97-0e98-4ee2-ad6a-4cada675e555")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CloneTestSuiteOperationInformation
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a test variable.
// ctx
// testVariableCreateUpdateParameters (required): TestVariableCreateUpdateParameters
// project (required): Project ID or project name
func (client Client) CreateTestVariable(ctx context.Context, testVariableCreateUpdateParameters *TestVariableCreateUpdateParameters, project *string) (*TestVariable, error) {
    if testVariableCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    body, marshalErr := json.Marshal(*testVariableCreateUpdateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestVariable
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a test variable by its ID.
// ctx
// project (required): Project ID or project name
// testVariableId (required): ID of the test variable to delete.
func (client Client) DeleteTestVariable(ctx context.Context, project *string, testVariableId *int) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if testVariableId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testVariableId"} 
    }
    routeValues["testVariableId"] = strconv.Itoa(*testVariableId)

    locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a test variable by its ID.
// ctx
// project (required): Project ID or project name
// testVariableId (required): ID of the test variable to get.
func (client Client) GetTestVariableById(ctx context.Context, project *string, testVariableId *int) (*TestVariable, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if testVariableId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableId"} 
    }
    routeValues["testVariableId"] = strconv.Itoa(*testVariableId)

    locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestVariable
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of test variables.
// ctx
// project (required): Project ID or project name
// continuationToken (optional): If the list of variables returned is not complete, a continuation token to query next batch of variables is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test variables.
func (client Client) GetTestVariables(ctx context.Context, project *string, continuationToken *string) (*[]TestVariable, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project

    queryParams := url.Values{}
    if continuationToken != nil {
        queryParams.Add("continuationToken", *continuationToken)
    }
    locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestVariable
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a test variable by its ID.
// ctx
// testVariableCreateUpdateParameters (required): TestVariableCreateUpdateParameters
// project (required): Project ID or project name
// testVariableId (required): ID of the test variable to update.
func (client Client) UpdateTestVariable(ctx context.Context, testVariableCreateUpdateParameters *TestVariableCreateUpdateParameters, project *string, testVariableId *int) (*TestVariable, error) {
    if testVariableCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if testVariableId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableId"} 
    }
    routeValues["testVariableId"] = strconv.Itoa(*testVariableId)

    body, marshalErr := json.Marshal(*testVariableCreateUpdateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestVariable
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

