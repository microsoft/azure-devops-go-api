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
func (client Client) CreateTestConfiguration(ctx context.Context, args CreateTestConfigurationArgs) (*TestConfiguration, error) {
    if args.TestConfigurationCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfigurationCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.TestConfigurationCreateUpdateParameters)
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

// Arguments for the CreateTestConfiguration function
type CreateTestConfigurationArgs struct {
    // (required) TestConfigurationCreateUpdateParameters
    TestConfigurationCreateUpdateParameters *TestConfigurationCreateUpdateParameters
    // (required) Project ID or project name
    Project *string
}

// [Preview API] Delete a test configuration by its ID.
func (client Client) DeleteTestConfguration(ctx context.Context, args DeleteTestConfgurationArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.TestConfiguartionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testConfiguartionId"}
    }
    queryParams.Add("testConfiguartionId", strconv.Itoa(*args.TestConfiguartionId))
    locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
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

// [Preview API] Get a test configuration
func (client Client) GetTestConfigurationById(ctx context.Context, args GetTestConfigurationByIdArgs) (*TestConfiguration, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.TestConfigurationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfigurationId"} 
    }
    routeValues["testConfigurationId"] = strconv.Itoa(*args.TestConfigurationId)

    locationId, _ := uuid.Parse("8369318e-38fa-4e84-9043-4b2a75d2c256")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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
func (client Client) GetTestConfigurations(ctx context.Context, args GetTestConfigurationsArgs) (*[]TestConfiguration, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", *args.ContinuationToken)
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

// Arguments for the GetTestConfigurations function
type GetTestConfigurationsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional) If the list of configurations returned is not complete, a continuation token to query next batch of configurations is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test configurations.
    ContinuationToken *string
}

// [Preview API] Update a test configuration by its ID.
func (client Client) UpdateTestConfiguration(ctx context.Context, args UpdateTestConfigurationArgs) (*TestConfiguration, error) {
    if args.TestConfigurationCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfigurationCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.TestConfiguartionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testConfiguartionId"}
    }
    queryParams.Add("testConfiguartionId", strconv.Itoa(*args.TestConfiguartionId))
    body, marshalErr := json.Marshal(*args.TestConfigurationCreateUpdateParameters)
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

// Arguments for the UpdateTestConfiguration function
type UpdateTestConfigurationArgs struct {
    // (required) TestConfigurationCreateUpdateParameters
    TestConfigurationCreateUpdateParameters *TestConfigurationCreateUpdateParameters
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test configuration to update.
    TestConfiguartionId *int
}

// [Preview API] Create a test plan.
func (client Client) CreateTestPlan(ctx context.Context, args CreateTestPlanArgs) (*TestPlan, error) {
    if args.TestPlanCreateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testPlanCreateParams"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.TestPlanCreateParams)
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

// Arguments for the CreateTestPlan function
type CreateTestPlanArgs struct {
    // (required) A testPlanCreateParams object.TestPlanCreateParams
    TestPlanCreateParams *TestPlanCreateParams
    // (required) Project ID or project name
    Project *string
}

// [Preview API] Delete a test plan.
func (client Client) DeleteTestPlan(ctx context.Context, args DeleteTestPlanArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)

    locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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

// [Preview API] Get a test plan by Id.
func (client Client) GetTestPlanById(ctx context.Context, args GetTestPlanByIdArgs) (*TestPlan, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)

    locationId, _ := uuid.Parse("0e292477-a0c2-47f3-a9b6-34f153d627f4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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
func (client Client) GetTestPlans(ctx context.Context, args GetTestPlansArgs) (*[]TestPlan, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
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
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPlan
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
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

// [Preview API] Update a test plan.
func (client Client) UpdateTestPlan(ctx context.Context, args UpdateTestPlanArgs) (*TestPlan, error) {
    if args.TestPlanUpdateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testPlanUpdateParams"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)

    body, marshalErr := json.Marshal(*args.TestPlanUpdateParams)
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

// Arguments for the UpdateTestPlan function
type UpdateTestPlanArgs struct {
    // (required) A testPlanUpdateParams object.TestPlanUpdateParams
    TestPlanUpdateParams *TestPlanUpdateParams
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan to be updated.
    PlanId *int
}

// [Preview API] Get a list of test suite entries in the test suite.
func (client Client) GetSuiteEntries(ctx context.Context, args GetSuiteEntriesArgs) (*[]SuiteEntry, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    queryParams := url.Values{}
    if args.SuiteEntryType != nil {
        queryParams.Add("suiteEntryType", string(*args.SuiteEntryType))
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

// Arguments for the GetSuiteEntries function
type GetSuiteEntriesArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of the parent suite.
    SuiteId *int
    // (optional)
    SuiteEntryType *SuiteEntryTypes
}

// [Preview API] Reorder test suite entries in the test suite.
func (client Client) ReorderSuiteEntries(ctx context.Context, args ReorderSuiteEntriesArgs) (*[]SuiteEntry, error) {
    if args.SuiteEntries == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteEntries"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    body, marshalErr := json.Marshal(*args.SuiteEntries)
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

// Arguments for the ReorderSuiteEntries function
type ReorderSuiteEntriesArgs struct {
    // (required) List of SuiteEntry to reorder.
    SuiteEntries *[]SuiteEntryUpdateParams
    // (required) Project ID or project name
    Project *string
    // (required) Id of the parent test suite.
    SuiteId *int
}

// [Preview API] Create test suite.
func (client Client) CreateTestSuite(ctx context.Context, args CreateTestSuiteArgs) (*TestSuite, error) {
    if args.TestSuiteCreateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSuiteCreateParams"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)

    body, marshalErr := json.Marshal(*args.TestSuiteCreateParams)
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

// Arguments for the CreateTestSuite function
type CreateTestSuiteArgs struct {
    // (required) Parameters for suite creation
    TestSuiteCreateParams *TestSuiteCreateParams
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan that contains the suites.
    PlanId *int
}

// [Preview API] Delete test suite.
func (client Client) DeleteTestSuite(ctx context.Context, args DeleteTestSuiteArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    locationId, _ := uuid.Parse("1046d5d3-ab61-4ca7-a65a-36118a978256")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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

// [Preview API] Get test suite by suite id.
func (client Client) GetTestSuiteById(ctx context.Context, args GetTestSuiteByIdArgs) (*TestSuite, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    queryParams := url.Values{}
    if args.Expand != nil {
        queryParams.Add("expand", string(*args.Expand))
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
func (client Client) GetTestSuitesForPlan(ctx context.Context, args GetTestSuitesForPlanArgs) (*[]TestSuite, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
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
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestSuite
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
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

// [Preview API] Update test suite.
func (client Client) UpdateTestSuite(ctx context.Context, args UpdateTestSuiteArgs) (*TestSuite, error) {
    if args.TestSuiteUpdateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSuiteUpdateParams"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    body, marshalErr := json.Marshal(*args.TestSuiteUpdateParams)
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

// Find the list of all test suites in which a given test case is present. This is helpful if you need to find out which test suites are using a test case, when you need to make changes to a test case.
func (client Client) GetSuitesByTestCaseId(ctx context.Context, args GetSuitesByTestCaseIdArgs) (*[]TestSuite, error) {
    queryParams := url.Values{}
    if args.TestCaseId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseId"}
    }
    queryParams.Add("testCaseId", strconv.Itoa(*args.TestCaseId))
    locationId, _ := uuid.Parse("a4080e84-f17b-4fad-84f1-7960b6525bf2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
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

// [Preview API] Add test cases to a suite with specified configurations
func (client Client) AddTestCasesToSuite(ctx context.Context, args AddTestCasesToSuiteArgs) (*[]TestCase, error) {
    if args.SuiteTestCaseCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteTestCaseCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    body, marshalErr := json.Marshal(*args.SuiteTestCaseCreateUpdateParameters)
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

// [Preview API] Get Test Cases For a Suite.
func (client Client) GetTestCase(ctx context.Context, args GetTestCaseArgs) (*[]TestCase, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.TestCaseIds == nil || *args.TestCaseIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
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
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
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

// [Preview API] Get Test Case List return those test cases which have all the configuration Ids as mentioned in the optional paramter. If configuration Ids is null, it return all the test cases
func (client Client) GetTestCaseList(ctx context.Context, args GetTestCaseListArgs) (*[]TestCase, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
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
    locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
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
}

// [Preview API] Removes test cases from a suite based on the list of test case Ids provided.
func (client Client) RemoveTestCasesFromSuite(ctx context.Context, args RemoveTestCasesFromSuiteArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.TestCaseIds == nil || *args.TestCaseIds == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = *args.TestCaseIds

    locationId, _ := uuid.Parse("a9bd61ac-45cf-4d13-9441-43dcd01edf8d")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
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

// [Preview API] Update the configurations for test cases
func (client Client) UpdateSuiteTestCases(ctx context.Context, args UpdateSuiteTestCasesArgs) (*[]TestCase, error) {
    if args.SuiteTestCaseCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteTestCaseCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    body, marshalErr := json.Marshal(*args.SuiteTestCaseCreateUpdateParameters)
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

// [Preview API] Delete a test case.
func (client Client) DeleteTestCase(ctx context.Context, args DeleteTestCaseArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.TestCaseId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testCaseId"} 
    }
    routeValues["testCaseId"] = strconv.Itoa(*args.TestCaseId)

    locationId, _ := uuid.Parse("29006fb5-816b-4ff7-a329-599943569229")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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

// [Preview API] Clone test plan
func (client Client) CloneTestPlan(ctx context.Context, args CloneTestPlanArgs) (*CloneTestPlanOperationInformation, error) {
    if args.CloneRequestBody == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneRequestBody"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
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
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
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

// [Preview API] Get clone information.
func (client Client) GetCloneInformation(ctx context.Context, args GetCloneInformationArgs) (*CloneTestPlanOperationInformation, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.CloneOperationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneOperationId"} 
    }
    routeValues["cloneOperationId"] = strconv.Itoa(*args.CloneOperationId)

    locationId, _ := uuid.Parse("e65df662-d8a3-46c7-ae1c-14e2d4df57e1")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
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
func (client Client) GetPoints(ctx context.Context, args GetPointsArgs) (*[]TestPoint, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.PointIds == nil || *args.PointIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pointIds"} 
    }
    routeValues["pointIds"] = *args.PointIds

    queryParams := url.Values{}
    if args.ReturnIdentityRef != nil {
        queryParams.Add("returnIdentityRef", strconv.FormatBool(*args.ReturnIdentityRef))
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
}

// [Preview API] Get all the points inside a suite based on some filters
func (client Client) GetPointsList(ctx context.Context, args GetPointsListArgs) (*[]TestPoint, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
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
    locationId, _ := uuid.Parse("52df686e-bae4-4334-b0ee-b6cf4e6f6b73")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
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
    // (optional) If set to false, returns only necessary information
    IncludePointDetails *bool
}

// [Preview API] Update Test Points. This is used to Reset test point to active, update the outcome of a test point or update the tester of a test point
func (client Client) UpdateTestPoints(ctx context.Context, args UpdateTestPointsArgs) (*[]TestPoint, error) {
    if args.TestPointUpdateParams == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testPointUpdateParams"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    body, marshalErr := json.Marshal(*args.TestPointUpdateParams)
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
}

// [Preview API] Clone test suite
func (client Client) CloneTestSuite(ctx context.Context, args CloneTestSuiteArgs) (*CloneTestSuiteOperationInformation, error) {
    if args.CloneRequestBody == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneRequestBody"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
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
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
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

// [Preview API] Get clone information.
func (client Client) GetSuiteCloneInformation(ctx context.Context, args GetSuiteCloneInformationArgs) (*CloneTestSuiteOperationInformation, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.CloneOperationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "cloneOperationId"} 
    }
    routeValues["cloneOperationId"] = strconv.Itoa(*args.CloneOperationId)

    locationId, _ := uuid.Parse("181d4c97-0e98-4ee2-ad6a-4cada675e555")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
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

// [Preview API] Create a test variable.
func (client Client) CreateTestVariable(ctx context.Context, args CreateTestVariableArgs) (*TestVariable, error) {
    if args.TestVariableCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.TestVariableCreateUpdateParameters)
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

// Arguments for the CreateTestVariable function
type CreateTestVariableArgs struct {
    // (required) TestVariableCreateUpdateParameters
    TestVariableCreateUpdateParameters *TestVariableCreateUpdateParameters
    // (required) Project ID or project name
    Project *string
}

// [Preview API] Delete a test variable by its ID.
func (client Client) DeleteTestVariable(ctx context.Context, args DeleteTestVariableArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.TestVariableId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testVariableId"} 
    }
    routeValues["testVariableId"] = strconv.Itoa(*args.TestVariableId)

    locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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

// [Preview API] Get a test variable by its ID.
func (client Client) GetTestVariableById(ctx context.Context, args GetTestVariableByIdArgs) (*TestVariable, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.TestVariableId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableId"} 
    }
    routeValues["testVariableId"] = strconv.Itoa(*args.TestVariableId)

    locationId, _ := uuid.Parse("2c61fac6-ac4e-45a5-8c38-1c2b8fd8ea6c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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
func (client Client) GetTestVariables(ctx context.Context, args GetTestVariablesArgs) (*[]TestVariable, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", *args.ContinuationToken)
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

// Arguments for the GetTestVariables function
type GetTestVariablesArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional) If the list of variables returned is not complete, a continuation token to query next batch of variables is included in the response header as "x-ms-continuationtoken". Omit this parameter to get the first batch of test variables.
    ContinuationToken *string
}

// [Preview API] Update a test variable by its ID.
func (client Client) UpdateTestVariable(ctx context.Context, args UpdateTestVariableArgs) (*TestVariable, error) {
    if args.TestVariableCreateUpdateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableCreateUpdateParameters"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.TestVariableId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testVariableId"} 
    }
    routeValues["testVariableId"] = strconv.Itoa(*args.TestVariableId)

    body, marshalErr := json.Marshal(*args.TestVariableCreateUpdateParameters)
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

// Arguments for the UpdateTestVariable function
type UpdateTestVariableArgs struct {
    // (required) TestVariableCreateUpdateParameters
    TestVariableCreateUpdateParameters *TestVariableCreateUpdateParameters
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test variable to update.
    TestVariableId *int
}

