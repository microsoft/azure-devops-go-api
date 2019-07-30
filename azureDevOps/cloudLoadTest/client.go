// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package cloudLoadTest

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevOps"
    "github.com/microsoft/azure-devops-go-api/azureDevOps/testService"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("7ae6d0a6-cda5-44cf-a261-28c392bed25c")

type Client struct {
    Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection azureDevOps.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

func (client Client) CreateAgentGroup(ctx context.Context, args CreateAgentGroupArgs) (*testService.AgentGroup, error) {
    if args.Group == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "group"}
    }
    body, marshalErr := json.Marshal(*args.Group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ab8d91c1-12d9-4ec5-874d-1ddb23e17720")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.AgentGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateAgentGroup function
type CreateAgentGroupArgs struct {
    // (required) Agent group to be created
    Group *testService.AgentGroup
}

func (client Client) GetAgentGroups(ctx context.Context, args GetAgentGroupsArgs) (interface{}, error) {
    routeValues := make(map[string]string)
    if args.AgentGroupId != nil && *args.AgentGroupId != "" {
        routeValues["agentGroupId"] = *args.AgentGroupId
    }

    queryParams := url.Values{}
    if args.MachineSetupInput != nil {
        queryParams.Add("machineSetupInput", strconv.FormatBool(*args.MachineSetupInput))
    }
    if args.MachineAccessData != nil {
        queryParams.Add("machineAccessData", strconv.FormatBool(*args.MachineAccessData))
    }
    if args.OutgoingRequestUrls != nil {
        queryParams.Add("outgoingRequestUrls", strconv.FormatBool(*args.OutgoingRequestUrls))
    }
    if args.AgentGroupName != nil {
        queryParams.Add("agentGroupName", *args.AgentGroupName)
    }
    locationId, _ := uuid.Parse("ab8d91c1-12d9-4ec5-874d-1ddb23e17720")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Arguments for the GetAgentGroups function
type GetAgentGroupsArgs struct {
    // (optional) The agent group identifier
    AgentGroupId *string
    // (optional)
    MachineSetupInput *bool
    // (optional)
    MachineAccessData *bool
    // (optional)
    OutgoingRequestUrls *bool
    // (optional) Name of the agent group
    AgentGroupName *string
}

func (client Client) DeleteStaticAgent(ctx context.Context, args DeleteStaticAgentArgs) (*string, error) {
    routeValues := make(map[string]string)
    if args.AgentGroupId == nil || *args.AgentGroupId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "agentGroupId"} 
    }
    routeValues["agentGroupId"] = *args.AgentGroupId

    queryParams := url.Values{}
    if args.AgentName == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "agentName"}
    }
    queryParams.Add("agentName", *args.AgentName)
    locationId, _ := uuid.Parse("87e4b63d-7142-4b50-801e-72ba9ff8ee9b")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the DeleteStaticAgent function
type DeleteStaticAgentArgs struct {
    // (required) The agent group identifier
    AgentGroupId *string
    // (required) Name of the static agent
    AgentName *string
}

func (client Client) GetStaticAgents(ctx context.Context, args GetStaticAgentsArgs) (interface{}, error) {
    routeValues := make(map[string]string)
    if args.AgentGroupId == nil || *args.AgentGroupId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "agentGroupId"} 
    }
    routeValues["agentGroupId"] = *args.AgentGroupId

    queryParams := url.Values{}
    if args.AgentName != nil {
        queryParams.Add("agentName", *args.AgentName)
    }
    locationId, _ := uuid.Parse("87e4b63d-7142-4b50-801e-72ba9ff8ee9b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Arguments for the GetStaticAgents function
type GetStaticAgentsArgs struct {
    // (required) The agent group identifier
    AgentGroupId *string
    // (optional) Name of the static agent
    AgentName *string
}

func (client Client) GetApplication(ctx context.Context, args GetApplicationArgs) (*testService.Application, error) {
    routeValues := make(map[string]string)
    if args.ApplicationId == nil || *args.ApplicationId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "applicationId"} 
    }
    routeValues["applicationId"] = *args.ApplicationId

    locationId, _ := uuid.Parse("2c986dce-8e8d-4142-b541-d016d5aff764")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.Application
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetApplication function
type GetApplicationArgs struct {
    // (required) Filter by APM application identifier.
    ApplicationId *string
}

func (client Client) GetApplications(ctx context.Context, args GetApplicationsArgs) (*[]testService.Application, error) {
    queryParams := url.Values{}
    if args.Type_ != nil {
        queryParams.Add("type_", *args.Type_)
    }
    locationId, _ := uuid.Parse("2c986dce-8e8d-4142-b541-d016d5aff764")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []testService.Application
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetApplications function
type GetApplicationsArgs struct {
    // (optional) Filters the results based on the plugin type.
    Type_ *string
}

func (client Client) GetCounters(ctx context.Context, args GetCountersArgs) (*[]testService.TestRunCounterInstance, error) {
    routeValues := make(map[string]string)
    if args.TestRunId == nil || *args.TestRunId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testRunId"} 
    }
    routeValues["testRunId"] = *args.TestRunId

    queryParams := url.Values{}
    if args.GroupNames == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "groupNames"}
    }
    queryParams.Add("groupNames", *args.GroupNames)
    if args.IncludeSummary != nil {
        queryParams.Add("includeSummary", strconv.FormatBool(*args.IncludeSummary))
    }
    locationId, _ := uuid.Parse("29265ea4-b5a5-4b2e-b054-47f5f6f00183")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []testService.TestRunCounterInstance
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetCounters function
type GetCountersArgs struct {
    // (required) The test run identifier
    TestRunId *string
    // (required) Comma separated names of counter groups, such as 'Application', 'Performance' and 'Throughput'
    GroupNames *string
    // (optional)
    IncludeSummary *bool
}

func (client Client) GetApplicationCounters(ctx context.Context, args GetApplicationCountersArgs) (*[]testService.ApplicationCounters, error) {
    queryParams := url.Values{}
    if args.ApplicationId != nil {
        queryParams.Add("applicationId", *args.ApplicationId)
    }
    if args.Plugintype != nil {
        queryParams.Add("plugintype", *args.Plugintype)
    }
    locationId, _ := uuid.Parse("c1275ce9-6d26-4bc6-926b-b846502e812d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []testService.ApplicationCounters
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetApplicationCounters function
type GetApplicationCountersArgs struct {
    // (optional) Filter by APM application identifier.
    ApplicationId *string
    // (optional) Currently ApplicationInsights is the only available plugin type.
    Plugintype *string
}

func (client Client) GetCounterSamples(ctx context.Context, args GetCounterSamplesArgs) (*testService.CounterSamplesResult, error) {
    if args.CounterSampleQueryDetails == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "counterSampleQueryDetails"}
    }
    routeValues := make(map[string]string)
    if args.TestRunId == nil || *args.TestRunId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testRunId"} 
    }
    routeValues["testRunId"] = *args.TestRunId

    body, marshalErr := json.Marshal(*args.CounterSampleQueryDetails)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bad18480-7193-4518-992a-37289c5bb92d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.CounterSamplesResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetCounterSamples function
type GetCounterSamplesArgs struct {
    // (required)
    CounterSampleQueryDetails *azureDevOps.VssJsonCollectionWrapper
    // (required) The test run identifier
    TestRunId *string
}

func (client Client) GetLoadTestRunErrors(ctx context.Context, args GetLoadTestRunErrorsArgs) (*testService.LoadTestErrors, error) {
    routeValues := make(map[string]string)
    if args.TestRunId == nil || *args.TestRunId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testRunId"} 
    }
    routeValues["testRunId"] = *args.TestRunId

    queryParams := url.Values{}
    if args.Type_ != nil {
        queryParams.Add("type_", *args.Type_)
    }
    if args.SubType != nil {
        queryParams.Add("subType", *args.SubType)
    }
    if args.Detailed != nil {
        queryParams.Add("detailed", strconv.FormatBool(*args.Detailed))
    }
    locationId, _ := uuid.Parse("b52025a7-3fb4-4283-8825-7079e75bd402")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.LoadTestErrors
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetLoadTestRunErrors function
type GetLoadTestRunErrorsArgs struct {
    // (required) The test run identifier
    TestRunId *string
    // (optional) Filter for the particular type of errors.
    Type_ *string
    // (optional) Filter for a particular subtype of errors. You should not provide error subtype without error type.
    SubType *string
    // (optional) To include the details of test errors such as messagetext, request, stacktrace, testcasename, scenarioname, and lasterrordate.
    Detailed *bool
}

func (client Client) GetTestRunMessages(ctx context.Context, args GetTestRunMessagesArgs) (*[]testService.TestRunMessage, error) {
    routeValues := make(map[string]string)
    if args.TestRunId == nil || *args.TestRunId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testRunId"} 
    }
    routeValues["testRunId"] = *args.TestRunId

    locationId, _ := uuid.Parse("2e7ba122-f522-4205-845b-2d270e59850a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []testService.TestRunMessage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunMessages function
type GetTestRunMessagesArgs struct {
    // (required) Id of the test run
    TestRunId *string
}

func (client Client) GetPlugin(ctx context.Context, args GetPluginArgs) (*testService.ApplicationType, error) {
    routeValues := make(map[string]string)
    if args.Type_ == nil || *args.Type_ == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *args.Type_

    locationId, _ := uuid.Parse("7dcb0bb2-42d5-4729-9958-c0401d5e7693")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.ApplicationType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPlugin function
type GetPluginArgs struct {
    // (required) Currently ApplicationInsights is the only available plugin type.
    Type_ *string
}

func (client Client) GetPlugins(ctx context.Context, args GetPluginsArgs) (*[]testService.ApplicationType, error) {
    locationId, _ := uuid.Parse("7dcb0bb2-42d5-4729-9958-c0401d5e7693")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []testService.ApplicationType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPlugins function
type GetPluginsArgs struct {
}

func (client Client) GetLoadTestResult(ctx context.Context, args GetLoadTestResultArgs) (*testService.TestResults, error) {
    routeValues := make(map[string]string)
    if args.TestRunId == nil || *args.TestRunId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testRunId"} 
    }
    routeValues["testRunId"] = *args.TestRunId

    locationId, _ := uuid.Parse("5ed69bd8-4557-4cec-9b75-1ad67d0c257b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestResults
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetLoadTestResult function
type GetLoadTestResultArgs struct {
    // (required) The test run identifier
    TestRunId *string
}

func (client Client) CreateTestDefinition(ctx context.Context, args CreateTestDefinitionArgs) (*testService.TestDefinition, error) {
    if args.TestDefinition == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testDefinition"}
    }
    body, marshalErr := json.Marshal(*args.TestDefinition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestDefinition function
type CreateTestDefinitionArgs struct {
    // (required) Test definition to be created
    TestDefinition *testService.TestDefinition
}

func (client Client) GetTestDefinition(ctx context.Context, args GetTestDefinitionArgs) (*testService.TestDefinition, error) {
    routeValues := make(map[string]string)
    if args.TestDefinitionId == nil || *args.TestDefinitionId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testDefinitionId"} 
    }
    routeValues["testDefinitionId"] = *args.TestDefinitionId

    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestDefinition function
type GetTestDefinitionArgs struct {
    // (required) The test definition identifier
    TestDefinitionId *string
}

func (client Client) GetTestDefinitions(ctx context.Context, args GetTestDefinitionsArgs) (*[]testService.TestDefinitionBasic, error) {
    queryParams := url.Values{}
    if args.FromDate != nil {
        queryParams.Add("fromDate", *args.FromDate)
    }
    if args.ToDate != nil {
        queryParams.Add("toDate", *args.ToDate)
    }
    if args.Top != nil {
        queryParams.Add("top", strconv.Itoa(*args.Top))
    }
    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []testService.TestDefinitionBasic
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestDefinitions function
type GetTestDefinitionsArgs struct {
    // (optional) Date after which test definitions were created
    FromDate *string
    // (optional) Date before which test definitions were crated
    ToDate *string
    // (optional)
    Top *int
}

func (client Client) UpdateTestDefinition(ctx context.Context, args UpdateTestDefinitionArgs) (*testService.TestDefinition, error) {
    if args.TestDefinition == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "testDefinition"}
    }
    body, marshalErr := json.Marshal(*args.TestDefinition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateTestDefinition function
type UpdateTestDefinitionArgs struct {
    // (required)
    TestDefinition *testService.TestDefinition
}

func (client Client) CreateTestDrop(ctx context.Context, args CreateTestDropArgs) (*testService.TestDrop, error) {
    if args.WebTestDrop == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "webTestDrop"}
    }
    body, marshalErr := json.Marshal(*args.WebTestDrop)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d89d0e08-505c-4357-96f6-9729311ce8ad")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestDrop
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestDrop function
type CreateTestDropArgs struct {
    // (required) Test drop to be created
    WebTestDrop *testService.TestDrop
}

func (client Client) GetTestDrop(ctx context.Context, args GetTestDropArgs) (*testService.TestDrop, error) {
    routeValues := make(map[string]string)
    if args.TestDropId == nil || *args.TestDropId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testDropId"} 
    }
    routeValues["testDropId"] = *args.TestDropId

    locationId, _ := uuid.Parse("d89d0e08-505c-4357-96f6-9729311ce8ad")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestDrop
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestDrop function
type GetTestDropArgs struct {
    // (required) The test drop identifier
    TestDropId *string
}

func (client Client) CreateTestRun(ctx context.Context, args CreateTestRunArgs) (*testService.TestRun, error) {
    if args.WebTestRun == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "webTestRun"}
    }
    body, marshalErr := json.Marshal(*args.WebTestRun)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestRun function
type CreateTestRunArgs struct {
    // (required)
    WebTestRun *testService.TestRun
}

func (client Client) GetTestRun(ctx context.Context, args GetTestRunArgs) (*testService.TestRun, error) {
    routeValues := make(map[string]string)
    if args.TestRunId == nil || *args.TestRunId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testRunId"} 
    }
    routeValues["testRunId"] = *args.TestRunId

    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue testService.TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRun function
type GetTestRunArgs struct {
    // (required) Unique ID of the test run
    TestRunId *string
}

// Returns test runs based on the filter specified. Returns all runs of the tenant if there is no filter.
func (client Client) GetTestRuns(ctx context.Context, args GetTestRunsArgs) (interface{}, error) {
    queryParams := url.Values{}
    if args.Name != nil {
        queryParams.Add("name", *args.Name)
    }
    if args.RequestedBy != nil {
        queryParams.Add("requestedBy", *args.RequestedBy)
    }
    if args.Status != nil {
        queryParams.Add("status", *args.Status)
    }
    if args.RunType != nil {
        queryParams.Add("runType", *args.RunType)
    }
    if args.FromDate != nil {
        queryParams.Add("fromDate", *args.FromDate)
    }
    if args.ToDate != nil {
        queryParams.Add("toDate", *args.ToDate)
    }
    if args.Detailed != nil {
        queryParams.Add("detailed", strconv.FormatBool(*args.Detailed))
    }
    if args.Top != nil {
        queryParams.Add("top", strconv.Itoa(*args.Top))
    }
    if args.Runsourceidentifier != nil {
        queryParams.Add("runsourceidentifier", *args.Runsourceidentifier)
    }
    if args.RetentionState != nil {
        queryParams.Add("retentionState", *args.RetentionState)
    }
    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// Arguments for the GetTestRuns function
type GetTestRunsArgs struct {
    // (optional) Name for the test run. Names are not unique. Test runs with same name are assigned sequential rolling numbers.
    Name *string
    // (optional) Filter by the user who requested the test run. Here requestedBy should be the display name of the user.
    RequestedBy *string
    // (optional) Filter by the test run status.
    Status *string
    // (optional) Valid values include: null, one of TestRunType, or "*"
    RunType *string
    // (optional) Filter by the test runs that have been modified after the fromDate timestamp.
    FromDate *string
    // (optional) Filter by the test runs that have been modified before the toDate timestamp.
    ToDate *string
    // (optional) Include the detailed test run attributes.
    Detailed *bool
    // (optional) The maximum number of test runs to return.
    Top *int
    // (optional)
    Runsourceidentifier *string
    // (optional)
    RetentionState *string
}

func (client Client) UpdateTestRun(ctx context.Context, args UpdateTestRunArgs) error {
    if args.WebTestRun == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "webTestRun"}
    }
    routeValues := make(map[string]string)
    if args.TestRunId == nil || *args.TestRunId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "testRunId"} 
    }
    routeValues["testRunId"] = *args.TestRunId

    body, marshalErr := json.Marshal(*args.WebTestRun)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the UpdateTestRun function
type UpdateTestRunArgs struct {
    // (required)
    WebTestRun *testService.TestRun
    // (required)
    TestRunId *string
}

