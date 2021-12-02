// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testplan

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/test"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/webapi"
)

// The build definition reference resource
type BuildDefinitionReference struct {
	// ID of the build definition
	Id *int `json:"id,omitempty"`
	// Name of the build definition
	Name *string `json:"name,omitempty"`
}

// Common Response for clone operation
type CloneOperationCommonResponse struct {
	// Various statistics related to the clone operation
	CloneStatistics *test.CloneStatistics `json:"cloneStatistics,omitempty"`
	// Completion data of the operation
	CompletionDate *azuredevops.Time `json:"completionDate,omitempty"`
	// Creation data of the operation
	CreationDate *azuredevops.Time `json:"creationDate,omitempty"`
	// Reference links
	Links interface{} `json:"links,omitempty"`
	// Message related to the job
	Message *string `json:"message,omitempty"`
	// Clone operation Id
	OpId *int `json:"opId,omitempty"`
	// Clone operation state
	State *test.CloneOperationState `json:"state,omitempty"`
}

type CloneTestCaseOperationInformation struct {
	// Various information related to the clone
	CloneOperationResponse *CloneOperationCommonResponse `json:"cloneOperationResponse,omitempty"`
	// Test Plan Clone create parameters
	CloneOptions *test.CloneTestCaseOptions `json:"cloneOptions,omitempty"`
	// Information of destination Test Suite
	DestinationTestSuite *TestSuiteReferenceWithProject `json:"destinationTestSuite,omitempty"`
	// Information of source Test Suite
	SourceTestSuite *SourceTestSuiteResponse `json:"sourceTestSuite,omitempty"`
}

// Parameters for Test Suite clone operation
type CloneTestCaseParams struct {
	// Test Case Clone create parameters
	CloneOptions *test.CloneTestCaseOptions `json:"cloneOptions,omitempty"`
	// Information about destination Test Plan
	DestinationTestPlan *TestPlanReference `json:"destinationTestPlan,omitempty"`
	// Information about destination Test Suite
	DestinationTestSuite *DestinationTestSuiteInfo `json:"destinationTestSuite,omitempty"`
	// Information about source Test Plan
	SourceTestPlan *TestPlanReference `json:"sourceTestPlan,omitempty"`
	// Information about source Test Suite
	SourceTestSuite *SourceTestSuiteInfo `json:"sourceTestSuite,omitempty"`
	// Test Case IDs
	TestCaseIds *[]int `json:"testCaseIds,omitempty"`
}

// Response for Test Plan clone operation
type CloneTestPlanOperationInformation struct {
	// Various information related to the clone
	CloneOperationResponse *CloneOperationCommonResponse `json:"cloneOperationResponse,omitempty"`
	// Test Plan Clone create parameters
	CloneOptions *test.CloneOptions `json:"cloneOptions,omitempty"`
	// Information of destination Test Plan
	DestinationTestPlan *TestPlan `json:"destinationTestPlan,omitempty"`
	// Information of source Test Plan
	SourceTestPlan *SourceTestplanResponse `json:"sourceTestPlan,omitempty"`
}

// Parameters for Test Plan clone operation
type CloneTestPlanParams struct {
	// Test Plan Clone create parameters
	CloneOptions *test.CloneOptions `json:"cloneOptions,omitempty"`
	// Information about destination Test Plan
	DestinationTestPlan *DestinationTestPlanCloneParams `json:"destinationTestPlan,omitempty"`
	// Information about source Test Plan
	SourceTestPlan *SourceTestPlanInfo `json:"sourceTestPlan,omitempty"`
}

// Response for Test Suite clone operation
type CloneTestSuiteOperationInformation struct {
	// Information of newly cloned Test Suite
	ClonedTestSuite *TestSuiteReferenceWithProject `json:"clonedTestSuite,omitempty"`
	// Various information related to the clone
	CloneOperationResponse *CloneOperationCommonResponse `json:"cloneOperationResponse,omitempty"`
	// Test Plan Clone create parameters
	CloneOptions *test.CloneOptions `json:"cloneOptions,omitempty"`
	// Information of destination Test Suite
	DestinationTestSuite *TestSuiteReferenceWithProject `json:"destinationTestSuite,omitempty"`
	// Information of source Test Suite
	SourceTestSuite *TestSuiteReferenceWithProject `json:"sourceTestSuite,omitempty"`
}

// Parameters for Test Suite clone operation
type CloneTestSuiteParams struct {
	// Test Plan Clone create parameters
	CloneOptions *test.CloneOptions `json:"cloneOptions,omitempty"`
	// Information about destination Test Suite
	DestinationTestSuite *DestinationTestSuiteInfo `json:"destinationTestSuite,omitempty"`
	// Information about source Test Suite
	SourceTestSuite *SourceTestSuiteInfo `json:"sourceTestSuite,omitempty"`
}

// Configuration of the Test Point
type Configuration struct {
	// Id of the Configuration Assigned to the Test Point
	ConfigurationId *int `json:"configurationId,omitempty"`
}

// Destination Test Plan create parameters
type DestinationTestPlanCloneParams struct {
	// Area of the test plan.
	AreaPath *string `json:"areaPath,omitempty"`
	// The Build Definition that generates a build associated with this test plan.
	BuildDefinition *BuildDefinitionReference `json:"buildDefinition,omitempty"`
	// Build to be tested.
	BuildId *int `json:"buildId,omitempty"`
	// Description of the test plan.
	Description *string `json:"description,omitempty"`
	// End date for the test plan.
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// Iteration path of the test plan.
	Iteration *string `json:"iteration,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// Owner of the test plan.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Release Environment to be used to deploy the build and run automated tests from this test plan.
	ReleaseEnvironmentDefinition *test.ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
	// Start date for the test plan.
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
	// State of the test plan.
	State *string `json:"state,omitempty"`
	// Value to configure how same tests across test suites under a test plan need to behave
	TestOutcomeSettings *test.TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
	// Destination Project Name
	Project *string `json:"project,omitempty"`
}

// Destination Test Suite information for Test Suite clone operation
type DestinationTestSuiteInfo struct {
	// Destination Suite Id
	Id *int `json:"id,omitempty"`
	// Destination Project Name
	Project *string `json:"project,omitempty"`
}

// [Flags] Exclude Flags for suite test case object. Exclude Flags exclude various objects from payload depending on the value passed
type ExcludeFlags string

type excludeFlagsValuesType struct {
	None             ExcludeFlags
	PointAssignments ExcludeFlags
	ExtraInformation ExcludeFlags
}

var ExcludeFlagsValues = excludeFlagsValuesType{
	// To exclude nothing
	None: "none",
	// To exclude point assignments, pass exclude = 1
	PointAssignments: "pointAssignments",
	// To exclude extra information (links, test plan, test suite), pass exclude = 2
	ExtraInformation: "extraInformation",
}

type FailureType string

type failureTypeValuesType struct {
	None        FailureType
	Regression  FailureType
	New_Issue   FailureType
	Known_Issue FailureType
	Unknown     FailureType
	Null_Value  FailureType
	MaxValue    FailureType
}

var FailureTypeValues = failureTypeValuesType{
	None:        "none",
	Regression:  "regression",
	New_Issue:   "new_Issue",
	Known_Issue: "known_Issue",
	Unknown:     "unknown",
	Null_Value:  "null_Value",
	MaxValue:    "maxValue",
}

type LastResolutionState string

type lastResolutionStateValuesType struct {
	None               LastResolutionState
	NeedsInvestigation LastResolutionState
	TestIssue          LastResolutionState
	ProductIssue       LastResolutionState
	ConfigurationIssue LastResolutionState
	NullValue          LastResolutionState
	MaxValue           LastResolutionState
}

var LastResolutionStateValues = lastResolutionStateValuesType{
	None:               "none",
	NeedsInvestigation: "needsInvestigation",
	TestIssue:          "testIssue",
	ProductIssue:       "productIssue",
	ConfigurationIssue: "configurationIssue",
	NullValue:          "nullValue",
	MaxValue:           "maxValue",
}

// Enum representing the return code of data provider.
type LibraryTestCasesDataReturnCode string

type libraryTestCasesDataReturnCodeValuesType struct {
	Success LibraryTestCasesDataReturnCode
	Error   LibraryTestCasesDataReturnCode
}

var LibraryTestCasesDataReturnCodeValues = libraryTestCasesDataReturnCodeValuesType{
	Success: "success",
	Error:   "error",
}

// This data model is used in Work item-based tabs of Test Plans Library.
type LibraryWorkItemsData struct {
	// Specifies the column option field names
	ColumnOptions *[]string `json:"columnOptions,omitempty"`
	// Continuation token to fetch next set of elements. Present only when HasMoreElements is true.
	ContinuationToken *string `json:"continuationToken,omitempty"`
	// Boolean indicating if the WIQL query has exceeded the limit of items returned.
	ExceededWorkItemQueryLimit *bool `json:"exceededWorkItemQueryLimit,omitempty"`
	// Boolean indicating if there are more elements present than what are being sent.
	HasMoreElements *bool `json:"hasMoreElements,omitempty"`
	// Specifies if there was an error while execution of data provider.
	ReturnCode *LibraryTestCasesDataReturnCode `json:"returnCode,omitempty"`
	// List of work items returned when OrderByField is sent something other than Id.
	WorkItemIds *[]int `json:"workItemIds,omitempty"`
	// List of work items to be returned.
	WorkItems *[]WorkItemDetails `json:"workItems,omitempty"`
}

// This is the request data contract for LibraryTestCaseDataProvider.
type LibraryWorkItemsDataProviderRequest struct {
	// Specifies the list of column options to show in test cases table.
	ColumnOptions *[]string `json:"columnOptions,omitempty"`
	// The continuation token required for paging of work items. This is required when getting subsequent sets of work items when OrderByField is Id.
	ContinuationToken *string `json:"continuationToken,omitempty"`
	// List of filter values to be supplied. Currently supported filters are Title, State, AssignedTo, Priority, AreaPath.
	FilterValues *[]TestPlansLibraryWorkItemFilter `json:"filterValues,omitempty"`
	// Whether the data is to be sorted in ascending or descending order. When not supplied, defaults to descending.
	IsAscending *bool `json:"isAscending,omitempty"`
	// The type of query to run.
	LibraryQueryType *TestPlansLibraryQuery `json:"libraryQueryType,omitempty"`
	// Work item field on which to order the results. When not supplied, defaults to work item IDs.
	OrderByField *string `json:"orderByField,omitempty"`
	// List of work items to query for field details. This is required when getting subsequent sets of work item fields when OrderByField is other than Id.
	WorkItemIds *[]int `json:"workItemIds,omitempty"`
}

type Outcome string

type outcomeValuesType struct {
	Unspecified   Outcome
	None          Outcome
	Passed        Outcome
	Failed        Outcome
	Inconclusive  Outcome
	Timeout       Outcome
	Aborted       Outcome
	Blocked       Outcome
	NotExecuted   Outcome
	Warning       Outcome
	Error         Outcome
	NotApplicable Outcome
	Paused        Outcome
	InProgress    Outcome
	NotImpacted   Outcome
	MaxValue      Outcome
}

var OutcomeValues = outcomeValuesType{
	// Only used during an update to preserve the existing value.
	Unspecified: "unspecified",
	// Test has not been completed, or the test type does not report pass/failure.
	None: "none",
	// Test was executed w/o any issues.
	Passed: "passed",
	// Test was executed, but there were issues. Issues may involve exceptions or failed assertions.
	Failed: "failed",
	// Test has completed, but we can't say if it passed or failed. May be used for aborted tests...
	Inconclusive: "inconclusive",
	// The test timed out
	Timeout: "timeout",
	// Test was aborted. This was not caused by a user gesture, but rather by a framework decision.
	Aborted: "aborted",
	// Test had it chance for been executed but was not, as ITestElement.IsRunnable == false.
	Blocked: "blocked",
	// Test was not executed. This was caused by a user gesture - e.g. user hit stop button.
	NotExecuted: "notExecuted",
	// To be used by Run level results. This is not a failure.
	Warning: "warning",
	// There was a system error while we were trying to execute a test.
	Error: "error",
	// Test is Not Applicable for execution.
	NotApplicable: "notApplicable",
	// Test is paused.
	Paused: "paused",
	// Test is currently executing. Added this for TCM charts
	InProgress: "inProgress",
	// Test is not impacted. Added fot TIA.
	NotImpacted: "notImpacted",
	MaxValue:    "maxValue",
}

// Assignments for the Test Point
type PointAssignment struct {
	// Id of the Configuration Assigned to the Test Point
	ConfigurationId *int `json:"configurationId,omitempty"`
	// Name of the Configuration Assigned to the Test Point
	ConfigurationName *string `json:"configurationName,omitempty"`
	// Id of the Test Point
	Id *int `json:"id,omitempty"`
	// Tester Assigned to the Test Point
	Tester *webapi.IdentityRef `json:"tester,omitempty"`
}

type PointState string

type pointStateValuesType struct {
	None       PointState
	Ready      PointState
	Completed  PointState
	NotReady   PointState
	InProgress PointState
	MaxValue   PointState
}

var PointStateValues = pointStateValuesType{
	// Default
	None: "none",
	// The test point needs to be executed in order for the test pass to be considered complete.  Either the test has not been run before or the previous run failed.
	Ready: "ready",
	// The test has passed successfully and does not need to be re-run for the test pass to be considered complete.
	Completed: "completed",
	// The test point needs to be executed but is not able to.
	NotReady: "notReady",
	// The test is being executed.
	InProgress: "inProgress",
	MaxValue:   "maxValue",
}

// Results class for Test Point
type Results struct {
	// Outcome of the Test Point
	Outcome *Outcome `json:"outcome,omitempty"`
}

type ResultState string

type resultStateValuesType struct {
	Unspecified ResultState
	Pending     ResultState
	Queued      ResultState
	InProgress  ResultState
	Paused      ResultState
	Completed   ResultState
	MaxValue    ResultState
}

var ResultStateValues = resultStateValuesType{
	// Only used during an update to preserve the existing value.
	Unspecified: "unspecified",
	// Test is in the execution queue, was not started yet.
	Pending: "pending",
	// Test has been queued. This is applicable when a test case is queued for execution
	Queued: "queued",
	// Test is currently executing.
	InProgress: "inProgress",
	// Test has been paused. This is applicable when a test case is paused by the user (For e.g. Manual Tester can pause the execution of the manual test case)
	Paused: "paused",
	// Test has completed, but there is no quantitative measure of completeness. This may apply to load tests.
	Completed: "completed",
	MaxValue:  "maxValue",
}

// Source Test Plan information for Test Plan clone operation
type SourceTestPlanInfo struct {
	// ID of the source Test Plan
	Id *int `json:"id,omitempty"`
	// Id of suites to be cloned inside source Test Plan
	SuiteIds *[]int `json:"suiteIds,omitempty"`
}

// Source Test Plan Response for Test Plan clone operation
type SourceTestplanResponse struct {
	// ID of the test plan.
	Id *int `json:"id,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// project reference
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// Id of suites to be cloned inside source Test Plan
	SuiteIds *[]int `json:"suiteIds,omitempty"`
}

// Source Test Suite information for Test Suite clone operation
type SourceTestSuiteInfo struct {
	// Id of the Source Test Suite
	Id *int `json:"id,omitempty"`
}

// Source Test Suite Response for Test Case clone operation
type SourceTestSuiteResponse struct {
	// ID of the test suite.
	Id *int `json:"id,omitempty"`
	// Name of the test suite.
	Name *string `json:"name,omitempty"`
	// project reference
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// Id of suites to be cloned inside source Test Plan
	TestCaseIds *[]int `json:"testCaseIds,omitempty"`
}

// A suite entry defines properties for a test suite.
type SuiteEntry struct {
	// Id of the suite entry in the test suite: either a test case id or child suite id.
	Id *int `json:"id,omitempty"`
	// Sequence number for the suite entry object in the test suite.
	SequenceNumber *int `json:"sequenceNumber,omitempty"`
	// Defines whether the entry is of type test case or suite.
	SuiteEntryType *SuiteEntryTypes `json:"suiteEntryType,omitempty"`
	// Id for the test suite.
	SuiteId *int `json:"suiteId,omitempty"`
}

type SuiteEntryTypes string

type suiteEntryTypesValuesType struct {
	TestCase SuiteEntryTypes
	Suite    SuiteEntryTypes
}

var SuiteEntryTypesValues = suiteEntryTypesValuesType{
	// Test Case
	TestCase: "testCase",
	// Child Suite
	Suite: "suite",
}

// A suite entry defines properties for a test suite.
type SuiteEntryUpdateParams struct {
	// Id of the suite entry in the test suite: either a test case id or child suite id.
	Id *int `json:"id,omitempty"`
	// Sequence number for the suite entry object in the test suite.
	SequenceNumber *int `json:"sequenceNumber,omitempty"`
	// Defines whether the entry is of type test case or suite.
	SuiteEntryType *SuiteEntryTypes `json:"suiteEntryType,omitempty"`
}

// [Flags] Option to get details in response
type SuiteExpand string

type suiteExpandValuesType struct {
	None           SuiteExpand
	Children       SuiteExpand
	DefaultTesters SuiteExpand
}

var SuiteExpandValues = suiteExpandValuesType{
	// Dont include any of the expansions in output.
	None: "none",
	// Include children in response.
	Children: "children",
	// Include default testers in response.
	DefaultTesters: "defaultTesters",
}

// Create and Update Suite Test Case Parameters
type SuiteTestCaseCreateUpdateParameters struct {
	// Configurations Ids
	PointAssignments *[]Configuration `json:"pointAssignments,omitempty"`
	// Id of Test Case to be updated or created
	WorkItem *WorkItem `json:"workItem,omitempty"`
}

// Test Case Class
type TestCase struct {
	// Reference links
	Links interface{} `json:"links,omitempty"`
	// Order of the TestCase in the Suite
	Order *int `json:"order,omitempty"`
	// List of Points associated with the Test Case
	PointAssignments *[]PointAssignment `json:"pointAssignments,omitempty"`
	// Project under which the Test Case is
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// Test Plan under which the Test Case is
	TestPlan *TestPlanReference `json:"testPlan,omitempty"`
	// Test Suite under which the Test Case is
	TestSuite *TestSuiteReference `json:"testSuite,omitempty"`
	// Work Item details of the TestCase
	WorkItem *WorkItemDetails `json:"workItem,omitempty"`
}

type TestCaseAssociatedResult struct {
	CompletedDate *azuredevops.Time           `json:"completedDate,omitempty"`
	Configuration *TestConfigurationReference `json:"configuration,omitempty"`
	Outcome       *UserFriendlyTestOutcome    `json:"outcome,omitempty"`
	Plan          *TestPlanReference          `json:"plan,omitempty"`
	PointId       *int                        `json:"pointId,omitempty"`
	ResultId      *int                        `json:"resultId,omitempty"`
	RunBy         *webapi.IdentityRef         `json:"runBy,omitempty"`
	RunId         *int                        `json:"runId,omitempty"`
	Suite         *TestSuiteReference         `json:"suite,omitempty"`
	Tester        *webapi.IdentityRef         `json:"tester,omitempty"`
}

// Test Case Reference
type TestCaseReference struct {
	// Identity to whom the test case is assigned
	AssignedTo *webapi.IdentityRef `json:"assignedTo,omitempty"`
	// Test Case Id
	Id *int `json:"id,omitempty"`
	// Test Case Name
	Name *string `json:"name,omitempty"`
	// State of the test case work item
	State *string `json:"state,omitempty"`
}

// This data model is used in TestCaseResultsDataProvider and populates the data required for initial page load
type TestCaseResultsData struct {
	// Point information from where the execution history was viewed. Used to set initial filters.
	ContextPoint *TestPointDetailedReference `json:"contextPoint,omitempty"`
	// Use to store the results displayed in the table
	Results *[]TestCaseAssociatedResult `json:"results,omitempty"`
	// Test Case Name to be displayed in the table header
	TestCaseName *string `json:"testCaseName,omitempty"`
}

// Test configuration
type TestConfiguration struct {
	// Description of the configuration
	Description *string `json:"description,omitempty"`
	// Is the configuration a default for the test plans
	IsDefault *bool `json:"isDefault,omitempty"`
	// Name of the configuration
	Name *string `json:"name,omitempty"`
	// State of the configuration
	State *test.TestConfigurationState `json:"state,omitempty"`
	// Dictionary of Test Variable, Selected Value
	Values *[]test.NameValuePair `json:"values,omitempty"`
	// Id of the configuration
	Id *int `json:"id,omitempty"`
	// Id of the test configuration variable
	Project *core.TeamProjectReference `json:"project,omitempty"`
}

// Test Configuration Create or Update Parameters
type TestConfigurationCreateUpdateParameters struct {
	// Description of the configuration
	Description *string `json:"description,omitempty"`
	// Is the configuration a default for the test plans
	IsDefault *bool `json:"isDefault,omitempty"`
	// Name of the configuration
	Name *string `json:"name,omitempty"`
	// State of the configuration
	State *test.TestConfigurationState `json:"state,omitempty"`
	// Dictionary of Test Variable, Selected Value
	Values *[]test.NameValuePair `json:"values,omitempty"`
}

// Test Configuration Reference
type TestConfigurationReference struct {
	// Id of the configuration
	Id *int `json:"id,omitempty"`
	// Name of the configuration
	Name *string `json:"name,omitempty"`
}

// Test Entity Count Used to store test cases count (define tab) and test point count (execute tab) Used to store test cases count (define tab) and test point count (execute tab)
type TestEntityCount struct {
	// Test Entity Count
	Count *int `json:"count,omitempty"`
	// Test Plan under which the Test Entities are
	TestPlanId *int `json:"testPlanId,omitempty"`
	// Test Suite under which the Test Entities are
	TestSuiteId *int `json:"testSuiteId,omitempty"`
	// Total test entities in the suite without the applied filters
	TotalCount *int `json:"totalCount,omitempty"`
}

type TestEntityTypes string

type testEntityTypesValuesType struct {
	TestCase  TestEntityTypes
	TestPoint TestEntityTypes
}

var TestEntityTypesValues = testEntityTypesValuesType{
	TestCase:  "testCase",
	TestPoint: "testPoint",
}

// The test plan resource.
type TestPlan struct {
	// Area of the test plan.
	AreaPath *string `json:"areaPath,omitempty"`
	// The Build Definition that generates a build associated with this test plan.
	BuildDefinition *BuildDefinitionReference `json:"buildDefinition,omitempty"`
	// Build to be tested.
	BuildId *int `json:"buildId,omitempty"`
	// Description of the test plan.
	Description *string `json:"description,omitempty"`
	// End date for the test plan.
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// Iteration path of the test plan.
	Iteration *string `json:"iteration,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// Owner of the test plan.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Release Environment to be used to deploy the build and run automated tests from this test plan.
	ReleaseEnvironmentDefinition *test.ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
	// Start date for the test plan.
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
	// State of the test plan.
	State *string `json:"state,omitempty"`
	// Value to configure how same tests across test suites under a test plan need to behave
	TestOutcomeSettings *test.TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
	// Revision of the test plan.
	Revision *int `json:"revision,omitempty"`
	// Relevant links
	Links interface{} `json:"_links,omitempty"`
	// ID of the test plan.
	Id *int `json:"id,omitempty"`
	// Previous build Id associated with the test plan
	PreviousBuildId *int `json:"previousBuildId,omitempty"`
	// Project which contains the test plan.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// Root test suite of the test plan.
	RootSuite *TestSuiteReference `json:"rootSuite,omitempty"`
	// Identity Reference for the last update of the test plan
	UpdatedBy *webapi.IdentityRef `json:"updatedBy,omitempty"`
	// Updated date of the test plan
	UpdatedDate *azuredevops.Time `json:"updatedDate,omitempty"`
}

// The test plan create parameters.
type TestPlanCreateParams struct {
	// Area of the test plan.
	AreaPath *string `json:"areaPath,omitempty"`
	// The Build Definition that generates a build associated with this test plan.
	BuildDefinition *BuildDefinitionReference `json:"buildDefinition,omitempty"`
	// Build to be tested.
	BuildId *int `json:"buildId,omitempty"`
	// Description of the test plan.
	Description *string `json:"description,omitempty"`
	// End date for the test plan.
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// Iteration path of the test plan.
	Iteration *string `json:"iteration,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// Owner of the test plan.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Release Environment to be used to deploy the build and run automated tests from this test plan.
	ReleaseEnvironmentDefinition *test.ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
	// Start date for the test plan.
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
	// State of the test plan.
	State *string `json:"state,omitempty"`
	// Value to configure how same tests across test suites under a test plan need to behave
	TestOutcomeSettings *test.TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
}

// The test plan detailed reference resource. Contains additional workitem realted information
type TestPlanDetailedReference struct {
	// ID of the test plan.
	Id *int `json:"id,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// Area of the test plan.
	AreaPath *string `json:"areaPath,omitempty"`
	// End date for the test plan.
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// Iteration path of the test plan.
	Iteration *string `json:"iteration,omitempty"`
	// Root Suite Id
	RootSuiteId *int `json:"rootSuiteId,omitempty"`
	// Start date for the test plan.
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
}

// The test plan reference resource.
type TestPlanReference struct {
	// ID of the test plan.
	Id *int `json:"id,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
}

// This data model is used in TestPlansHubRefreshDataProvider and populates the data required for initial page load
type TestPlansHubRefreshData struct {
	DefineColumnOptionFields    *[]string                  `json:"defineColumnOptionFields,omitempty"`
	ErrorMessage                *string                    `json:"errorMessage,omitempty"`
	ExecuteColumnOptionFields   *[]string                  `json:"executeColumnOptionFields,omitempty"`
	IsAdvancedExtensionEnabled  *bool                      `json:"isAdvancedExtensionEnabled,omitempty"`
	SelectedPivotId             *string                    `json:"selectedPivotId,omitempty"`
	SelectedSuiteId             *int                       `json:"selectedSuiteId,omitempty"`
	TestCasePageSize            *int                       `json:"testCasePageSize,omitempty"`
	TestCases                   *[]TestCase                `json:"testCases,omitempty"`
	TestCasesContinuationToken  *string                    `json:"testCasesContinuationToken,omitempty"`
	TestPlan                    *TestPlanDetailedReference `json:"testPlan,omitempty"`
	TestPointPageSize           *int                       `json:"testPointPageSize,omitempty"`
	TestPoints                  *[]TestPoint               `json:"testPoints,omitempty"`
	TestPointsContinuationToken *string                    `json:"testPointsContinuationToken,omitempty"`
	TestSuites                  *[]TestSuite               `json:"testSuites,omitempty"`
	TestSuitesContinuationToken *string                    `json:"testSuitesContinuationToken,omitempty"`
}

// Enum used to define the queries used in Test Plans Library.
type TestPlansLibraryQuery string

type testPlansLibraryQueryValuesType struct {
	None                              TestPlansLibraryQuery
	AllTestCases                      TestPlansLibraryQuery
	TestCasesWithActiveBugs           TestPlansLibraryQuery
	TestCasesNotLinkedToRequirements  TestPlansLibraryQuery
	TestCasesLinkedToRequirements     TestPlansLibraryQuery
	AllSharedSteps                    TestPlansLibraryQuery
	SharedStepsNotLinkedToRequirement TestPlansLibraryQuery
}

var TestPlansLibraryQueryValues = testPlansLibraryQueryValuesType{
	None:                              "none",
	AllTestCases:                      "allTestCases",
	TestCasesWithActiveBugs:           "testCasesWithActiveBugs",
	TestCasesNotLinkedToRequirements:  "testCasesNotLinkedToRequirements",
	TestCasesLinkedToRequirements:     "testCasesLinkedToRequirements",
	AllSharedSteps:                    "allSharedSteps",
	SharedStepsNotLinkedToRequirement: "sharedStepsNotLinkedToRequirement",
}

// Container to hold information about a filter being applied in Test Plans Library.
type TestPlansLibraryWorkItemFilter struct {
	// Work item field name on which the items are to be filtered.
	FieldName *string `json:"fieldName,omitempty"`
	// Work item field values corresponding to the field name.
	FieldValues *[]string `json:"fieldValues,omitempty"`
	// Mode of the filter.
	FilterMode *TestPlansLibraryWorkItemFilterMode `json:"filterMode,omitempty"`
}

type TestPlansLibraryWorkItemFilterMode string

type testPlansLibraryWorkItemFilterModeValuesType struct {
	Or  TestPlansLibraryWorkItemFilterMode
	And TestPlansLibraryWorkItemFilterMode
}

var TestPlansLibraryWorkItemFilterModeValues = testPlansLibraryWorkItemFilterModeValuesType{
	// Default. Have the field values separated by an OR clause.
	Or: "or",
	// Have the field values separated by an AND clause.
	And: "and",
}

// The test plan update parameters.
type TestPlanUpdateParams struct {
	// Area of the test plan.
	AreaPath *string `json:"areaPath,omitempty"`
	// The Build Definition that generates a build associated with this test plan.
	BuildDefinition *BuildDefinitionReference `json:"buildDefinition,omitempty"`
	// Build to be tested.
	BuildId *int `json:"buildId,omitempty"`
	// Description of the test plan.
	Description *string `json:"description,omitempty"`
	// End date for the test plan.
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// Iteration path of the test plan.
	Iteration *string `json:"iteration,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// Owner of the test plan.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Release Environment to be used to deploy the build and run automated tests from this test plan.
	ReleaseEnvironmentDefinition *test.ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
	// Start date for the test plan.
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
	// State of the test plan.
	State *string `json:"state,omitempty"`
	// Value to configure how same tests across test suites under a test plan need to behave
	TestOutcomeSettings *test.TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
	// Revision of the test plan.
	Revision *int `json:"revision,omitempty"`
}

// Test Point Class
type TestPoint struct {
	// Comment associated to the Test Point
	Comment *string `json:"comment,omitempty"`
	// Configuration associated with the Test Point
	Configuration *TestConfigurationReference `json:"configuration,omitempty"`
	// Id of the Test Point
	Id *int `json:"id,omitempty"`
	// Variable to decide whether the test case is Active or not
	IsActive *bool `json:"isActive,omitempty"`
	// Is the Test Point for Automated Test Case or Manual
	IsAutomated *bool `json:"isAutomated,omitempty"`
	// Last Reset to Active Time Stamp for the Test Point
	LastResetToActive *azuredevops.Time `json:"lastResetToActive,omitempty"`
	// Last Updated details for the Test Point
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last Update Time Stamp for the Test Point
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Reference links
	Links interface{} `json:"links,omitempty"`
	// Project under which the Test Point is
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// Results associated to the Test Point
	Results *TestPointResults `json:"results,omitempty"`
	// Test Case Reference
	TestCaseReference *TestCaseReference `json:"testCaseReference,omitempty"`
	// Tester associated with the Test Point
	Tester *webapi.IdentityRef `json:"tester,omitempty"`
	// Test Plan under which the Test Point is
	TestPlan *TestPlanReference `json:"testPlan,omitempty"`
	// Test Suite under which the Test Point is
	TestSuite *TestSuiteReference `json:"testSuite,omitempty"`
}

type TestPointDetailedReference struct {
	Configuration *TestConfigurationReference `json:"configuration,omitempty"`
	Plan          *TestPlanReference          `json:"plan,omitempty"`
	PointId       *int                        `json:"pointId,omitempty"`
	Suite         *TestSuiteReference         `json:"suite,omitempty"`
	Tester        *webapi.IdentityRef         `json:"tester,omitempty"`
}

// Test Point Results
type TestPointResults struct {
	// Failure Type for the Test Point
	FailureType *FailureType `json:"failureType,omitempty"`
	// Last Resolution State Id for the Test Point
	LastResolutionState *LastResolutionState `json:"lastResolutionState,omitempty"`
	// Last Result Details for the Test Point
	LastResultDetails *test.LastResultDetails `json:"lastResultDetails,omitempty"`
	// Last Result Id
	LastResultId *int `json:"lastResultId,omitempty"`
	// Last Result State of the Test Point
	LastResultState *ResultState `json:"lastResultState,omitempty"`
	// Last RUn Build Number for the Test Point
	LastRunBuildNumber *string `json:"lastRunBuildNumber,omitempty"`
	// Last Test Run Id for the Test Point
	LastTestRunId *int `json:"lastTestRunId,omitempty"`
	// Outcome of the Test Point
	Outcome *Outcome `json:"outcome,omitempty"`
	// State of the Test Point
	State *PointState `json:"state,omitempty"`
}

// Test Point Update Parameters
type TestPointUpdateParams struct {
	// Id of Test Point to be updated
	Id *int `json:"id,omitempty"`
	// Reset the Test Point to Active
	IsActive *bool `json:"isActive,omitempty"`
	// Results of the test point
	Results *Results `json:"results,omitempty"`
	// Tester of the Test Point
	Tester *webapi.IdentityRef `json:"tester,omitempty"`
}

// Test suite
type TestSuite struct {
	// Test suite default configurations.
	DefaultConfigurations *[]TestConfigurationReference `json:"defaultConfigurations,omitempty"`
	// Test suite default testers.
	DefaultTesters *[]webapi.IdentityRef `json:"defaultTesters,omitempty"`
	// Default configuration was inherited or not.
	InheritDefaultConfigurations *bool `json:"inheritDefaultConfigurations,omitempty"`
	// Name of test suite.
	Name *string `json:"name,omitempty"`
	// Test suite parent shallow reference.
	ParentSuite *TestSuiteReference `json:"parentSuite,omitempty"`
	// Test suite query string, for dynamic suites.
	QueryString *string `json:"queryString,omitempty"`
	// Test suite requirement id.
	RequirementId *int `json:"requirementId,omitempty"`
	// Test suite type.
	SuiteType *TestSuiteType `json:"suiteType,omitempty"`
	// Links: self, testPoints, testCases, parent
	Links interface{} `json:"_links,omitempty"`
	// Child test suites of current test suite.
	Children *[]TestSuite `json:"children,omitempty"`
	// Boolean value dictating if Child test suites are present
	HasChildren *bool `json:"hasChildren,omitempty"`
	// Id of test suite.
	Id *int `json:"id,omitempty"`
	// Last error for test suite.
	LastError *string `json:"lastError,omitempty"`
	// Last populated date.
	LastPopulatedDate *azuredevops.Time `json:"lastPopulatedDate,omitempty"`
	// IdentityRef of user who has updated test suite recently.
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last update date.
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Test plan to which the test suite belongs.
	Plan *TestPlanReference `json:"plan,omitempty"`
	// Test suite project shallow reference.
	Project *core.TeamProjectReference `json:"project,omitempty"`
	// Test suite revision.
	Revision *int `json:"revision,omitempty"`
}

// Test suite Create Parameters
type TestSuiteCreateParams struct {
	// Test suite default configurations.
	DefaultConfigurations *[]TestConfigurationReference `json:"defaultConfigurations,omitempty"`
	// Test suite default testers.
	DefaultTesters *[]webapi.IdentityRef `json:"defaultTesters,omitempty"`
	// Default configuration was inherited or not.
	InheritDefaultConfigurations *bool `json:"inheritDefaultConfigurations,omitempty"`
	// Name of test suite.
	Name *string `json:"name,omitempty"`
	// Test suite parent shallow reference.
	ParentSuite *TestSuiteReference `json:"parentSuite,omitempty"`
	// Test suite query string, for dynamic suites.
	QueryString *string `json:"queryString,omitempty"`
	// Test suite requirement id.
	RequirementId *int `json:"requirementId,omitempty"`
	// Test suite type.
	SuiteType *TestSuiteType `json:"suiteType,omitempty"`
}

// Test Suite Create/Update Common Parameters
type TestSuiteCreateUpdateCommonParams struct {
	// Test suite default configurations.
	DefaultConfigurations *[]TestConfigurationReference `json:"defaultConfigurations,omitempty"`
	// Test suite default testers.
	DefaultTesters *[]webapi.IdentityRef `json:"defaultTesters,omitempty"`
	// Default configuration was inherited or not.
	InheritDefaultConfigurations *bool `json:"inheritDefaultConfigurations,omitempty"`
	// Name of test suite.
	Name *string `json:"name,omitempty"`
	// Test suite parent shallow reference.
	ParentSuite *TestSuiteReference `json:"parentSuite,omitempty"`
	// Test suite query string, for dynamic suites.
	QueryString *string `json:"queryString,omitempty"`
}

// The test suite reference resource.
type TestSuiteReference struct {
	// ID of the test suite.
	Id *int `json:"id,omitempty"`
	// Name of the test suite.
	Name *string `json:"name,omitempty"`
}

// Test Suite Reference with Project
type TestSuiteReferenceWithProject struct {
	// ID of the test suite.
	Id *int `json:"id,omitempty"`
	// Name of the test suite.
	Name *string `json:"name,omitempty"`
	// Reference of destination Project
	Project *core.TeamProjectReference `json:"project,omitempty"`
}

// Type of TestSuite
type TestSuiteType string

type testSuiteTypeValuesType struct {
	None                 TestSuiteType
	DynamicTestSuite     TestSuiteType
	StaticTestSuite      TestSuiteType
	RequirementTestSuite TestSuiteType
}

var TestSuiteTypeValues = testSuiteTypeValuesType{
	// Default suite type
	None: "none",
	// Query Based test Suite
	DynamicTestSuite: "dynamicTestSuite",
	// Static Test Suite
	StaticTestSuite: "staticTestSuite",
	// Requirement based Test Suite
	RequirementTestSuite: "requirementTestSuite",
}

// Test Suite Update Parameters
type TestSuiteUpdateParams struct {
	// Test suite default configurations.
	DefaultConfigurations *[]TestConfigurationReference `json:"defaultConfigurations,omitempty"`
	// Test suite default testers.
	DefaultTesters *[]webapi.IdentityRef `json:"defaultTesters,omitempty"`
	// Default configuration was inherited or not.
	InheritDefaultConfigurations *bool `json:"inheritDefaultConfigurations,omitempty"`
	// Name of test suite.
	Name *string `json:"name,omitempty"`
	// Test suite parent shallow reference.
	ParentSuite *TestSuiteReference `json:"parentSuite,omitempty"`
	// Test suite query string, for dynamic suites.
	QueryString *string `json:"queryString,omitempty"`
	// Test suite revision.
	Revision *int `json:"revision,omitempty"`
}

// Test Variable
type TestVariable struct {
	// Description of the test variable
	Description *string `json:"description,omitempty"`
	// Name of the test variable
	Name *string `json:"name,omitempty"`
	// List of allowed values
	Values *[]string `json:"values,omitempty"`
	// Id of the test variable
	Id *int `json:"id,omitempty"`
	// Id of the test variable
	Project *core.TeamProjectReference `json:"project,omitempty"`
}

// Test Variable Create or Update Parameters
type TestVariableCreateUpdateParameters struct {
	// Description of the test variable
	Description *string `json:"description,omitempty"`
	// Name of the test variable
	Name *string `json:"name,omitempty"`
	// List of allowed values
	Values *[]string `json:"values,omitempty"`
}

type UserFriendlyTestOutcome string

type userFriendlyTestOutcomeValuesType struct {
	InProgress    UserFriendlyTestOutcome
	Blocked       UserFriendlyTestOutcome
	Failed        UserFriendlyTestOutcome
	Passed        UserFriendlyTestOutcome
	Ready         UserFriendlyTestOutcome
	NotApplicable UserFriendlyTestOutcome
	Paused        UserFriendlyTestOutcome
	Timeout       UserFriendlyTestOutcome
	Warning       UserFriendlyTestOutcome
	Error         UserFriendlyTestOutcome
	NotExecuted   UserFriendlyTestOutcome
	Inconclusive  UserFriendlyTestOutcome
	Aborted       UserFriendlyTestOutcome
	None          UserFriendlyTestOutcome
	NotImpacted   UserFriendlyTestOutcome
	Unspecified   UserFriendlyTestOutcome
	MaxValue      UserFriendlyTestOutcome
}

var UserFriendlyTestOutcomeValues = userFriendlyTestOutcomeValuesType{
	InProgress:    "inProgress",
	Blocked:       "blocked",
	Failed:        "failed",
	Passed:        "passed",
	Ready:         "ready",
	NotApplicable: "notApplicable",
	Paused:        "paused",
	Timeout:       "timeout",
	Warning:       "warning",
	Error:         "error",
	NotExecuted:   "notExecuted",
	Inconclusive:  "inconclusive",
	Aborted:       "aborted",
	None:          "none",
	NotImpacted:   "notImpacted",
	Unspecified:   "unspecified",
	MaxValue:      "maxValue",
}

// Work Item
type WorkItem struct {
	// Id of the Work Item
	Id *int `json:"id,omitempty"`
}

// Work Item Class
type WorkItemDetails struct {
	// Work Item Id
	Id *int `json:"id,omitempty"`
	// Work Item Name
	Name *string `json:"name,omitempty"`
	// Work Item Fields
	WorkItemFields *[]interface{} `json:"workItemFields,omitempty"`
}
