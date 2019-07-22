// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testPlan

import (
    "github.com/google/uuid"
    "time"
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
    CloneStatistics *CloneStatistics `json:"cloneStatistics,omitempty"`
    // Completion data of the operation
    CompletionDate *time.Time `json:"completionDate,omitempty"`
    // Creation data of the operation
    CreationDate *time.Time `json:"creationDate,omitempty"`
    // Reference links
    Links *ReferenceLinks `json:"links,omitempty"`
    // Message related to the job
    Message *string `json:"message,omitempty"`
    // Clone operation Id
    OpId *int `json:"opId,omitempty"`
    // Clone operation state
    State *CloneOperationState `json:"state,omitempty"`
}

// Enum of type Clone Operation Type.
type CloneOperationState string

type cloneOperationStateValuesType struct {
    Failed CloneOperationState
    InProgress CloneOperationState
    Queued CloneOperationState
    Succeeded CloneOperationState
}

var CloneOperationStateValues = cloneOperationStateValuesType{
    // value for Failed State
    Failed: "failed",
    // value for Inprogress state
    InProgress: "inProgress",
    // Value for Queued State
    Queued: "queued",
    // value for Success state
    Succeeded: "succeeded",
}

// Clone options for cloning the test suite.
type CloneOptions struct {
    // If set to true requirements will be cloned
    CloneRequirements *bool `json:"cloneRequirements,omitempty"`
    // copy all suites from a source plan
    CopyAllSuites *bool `json:"copyAllSuites,omitempty"`
    // copy ancestor hierarchy
    CopyAncestorHierarchy *bool `json:"copyAncestorHierarchy,omitempty"`
    // Name of the workitem type of the clone
    DestinationWorkItemType *string `json:"destinationWorkItemType,omitempty"`
    // Key value pairs where the key value is overridden by the value.
    OverrideParameters *map[string]string `json:"overrideParameters,omitempty"`
    // Comment on the link that will link the new clone  test case to the original Set null for no comment
    RelatedLinkComment *string `json:"relatedLinkComment,omitempty"`
}

// Clone Statistics Details.
type CloneStatistics struct {
    // Number of requirements cloned so far.
    ClonedRequirementsCount *int `json:"clonedRequirementsCount,omitempty"`
    // Number of shared steps cloned so far.
    ClonedSharedStepsCount *int `json:"clonedSharedStepsCount,omitempty"`
    // Number of test cases cloned so far
    ClonedTestCasesCount *int `json:"clonedTestCasesCount,omitempty"`
    // Total number of requirements to be cloned
    TotalRequirementsCount *int `json:"totalRequirementsCount,omitempty"`
    // Total number of test cases to be cloned
    TotalTestCasesCount *int `json:"totalTestCasesCount,omitempty"`
}

// Response for Test Plan clone operation
type CloneTestPlanOperationInformation struct {
    // Various information related to the clone
    CloneOperationResponse *CloneOperationCommonResponse `json:"cloneOperationResponse,omitempty"`
    // Test Plan Clone create parameters
    CloneOptions *CloneOptions `json:"cloneOptions,omitempty"`
    // Information of destination Test Plan
    DestinationTestPlan *TestPlan `json:"destinationTestPlan,omitempty"`
    // Information of source Test Plan
    SourceTestPlan *SourceTestplanResponse `json:"sourceTestPlan,omitempty"`
}

// Parameters for Test Plan clone operation
type CloneTestPlanParams struct {
    // Test Plan Clone create parameters
    CloneOptions *CloneOptions `json:"cloneOptions,omitempty"`
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
    CloneOptions *CloneOptions `json:"cloneOptions,omitempty"`
    // Information of destination Test Suite
    DestinationTestSuite *TestSuiteReferenceWithProject `json:"destinationTestSuite,omitempty"`
    // Information of source Test Suite
    SourceTestSuite *TestSuiteReferenceWithProject `json:"sourceTestSuite,omitempty"`
}

// Parameters for Test Suite clone operation
type CloneTestSuiteParams struct {
    // Test Plan Clone create parameters
    CloneOptions *CloneOptions `json:"cloneOptions,omitempty"`
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
    EndDate *time.Time `json:"endDate,omitempty"`
    // Iteration path of the test plan.
    Iteration *string `json:"iteration,omitempty"`
    // Name of the test plan.
    Name *string `json:"name,omitempty"`
    // Owner of the test plan.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Release Environment to be used to deploy the build and run automated tests from this test plan.
    ReleaseEnvironmentDefinition *ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
    // Start date for the test plan.
    StartDate *time.Time `json:"startDate,omitempty"`
    // State of the test plan.
    State *string `json:"state,omitempty"`
    // Value to configure how same tests across test suites under a test plan need to behave
    TestOutcomeSettings *TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
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

type FailureType string

type failureTypeValuesType struct {
    None FailureType
    Regression FailureType
    New_Issue FailureType
    Known_Issue FailureType
    Unknown FailureType
    Null_Value FailureType
    MaxValue FailureType
}

var FailureTypeValues = failureTypeValuesType{
    None: "none",
    Regression: "regression",
    New_Issue: "new_Issue",
    Known_Issue: "known_Issue",
    Unknown: "unknown",
    Null_Value: "null_Value",
    MaxValue: "maxValue",
}

type GraphSubjectBase struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
}

type IdentityRef struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph user referenced in the "self" entry of the IdentityRef "_links" dictionary
    DirectoryAlias *string `json:"directoryAlias,omitempty"`
    Id *string `json:"id,omitempty"`
    // Deprecated - Available in the "avatar" entry of the IdentityRef "_links" dictionary
    ImageUrl *string `json:"imageUrl,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph membership state referenced in the "membershipState" entry of the GraphUser "_links" dictionary
    Inactive *bool `json:"inactive,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsAadUserType/Descriptor.IsAadGroupType)
    IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsGroupType)
    IsContainer *bool `json:"isContainer,omitempty"`
    IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
    // Deprecated - not in use in most preexisting implementations of ToIdentityRef
    ProfileUrl *string `json:"profileUrl,omitempty"`
    // Deprecated - use Domain+PrincipalName instead
    UniqueName *string `json:"uniqueName,omitempty"`
}

type LastResolutionState string

type lastResolutionStateValuesType struct {
    None LastResolutionState
    NeedsInvestigation LastResolutionState
    TestIssue LastResolutionState
    ProductIssue LastResolutionState
    ConfigurationIssue LastResolutionState
    NullValue LastResolutionState
    MaxValue LastResolutionState
}

var LastResolutionStateValues = lastResolutionStateValuesType{
    None: "none",
    NeedsInvestigation: "needsInvestigation",
    TestIssue: "testIssue",
    ProductIssue: "productIssue",
    ConfigurationIssue: "configurationIssue",
    NullValue: "nullValue",
    MaxValue: "maxValue",
}

// Last result details of test point.
type LastResultDetails struct {
    // CompletedDate of LastResult.
    DateCompleted *time.Time `json:"dateCompleted,omitempty"`
    // Duration of LastResult.
    Duration *uint64 `json:"duration,omitempty"`
    // RunBy.
    RunBy *IdentityRef `json:"runBy,omitempty"`
}

// Name value pair
type NameValuePair struct {
    // Name
    Name *string `json:"name,omitempty"`
    // Value
    Value *string `json:"value,omitempty"`
}

type Outcome string

type outcomeValuesType struct {
    Unspecified Outcome
    None Outcome
    Passed Outcome
    Failed Outcome
    Inconclusive Outcome
    Timeout Outcome
    Aborted Outcome
    Blocked Outcome
    NotExecuted Outcome
    Warning Outcome
    Error Outcome
    NotApplicable Outcome
    Paused Outcome
    InProgress Outcome
    NotImpacted Outcome
    MaxValue Outcome
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
    MaxValue: "maxValue",
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
    Tester *IdentityRef `json:"tester,omitempty"`
}

type PointState string

type pointStateValuesType struct {
    None PointState
    Ready PointState
    Completed PointState
    NotReady PointState
    InProgress PointState
    MaxValue PointState
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
    MaxValue: "maxValue",
}

type ProjectState string

type projectStateValuesType struct {
    Deleting ProjectState
    New ProjectState
    WellFormed ProjectState
    CreatePending ProjectState
    All ProjectState
    Unchanged ProjectState
    Deleted ProjectState
}

var ProjectStateValues = projectStateValuesType{
    // Project is in the process of being deleted.
    Deleting: "deleting",
    // Project is in the process of being created.
    New: "new",
    // Project is completely created and ready to use.
    WellFormed: "wellFormed",
    // Project has been queued for creation, but the process has not yet started.
    CreatePending: "createPending",
    // All projects regardless of state.
    All: "all",
    // Project has not been changed.
    Unchanged: "unchanged",
    // Project has been deleted.
    Deleted: "deleted",
}

type ProjectVisibility string

type projectVisibilityValuesType struct {
    Private ProjectVisibility
    Public ProjectVisibility
}

var ProjectVisibilityValues = projectVisibilityValuesType{
    // The project is only visible to users with explicit access.
    Private: "private",
    // The project is visible to all.
    Public: "public",
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

// Reference to release environment resource.
type ReleaseEnvironmentDefinitionReference struct {
    // ID of the release definition that contains the release environment definition.
    DefinitionId *int `json:"definitionId,omitempty"`
    // ID of the release environment definition.
    EnvironmentDefinitionId *int `json:"environmentDefinitionId,omitempty"`
}

// Results class for Test Point
type Results struct {
    // Outcome of the Test Point
    Outcome *Outcome `json:"outcome,omitempty"`
}

type ResultState string

type resultStateValuesType struct {
    Unspecified ResultState
    Pending ResultState
    Queued ResultState
    InProgress ResultState
    Paused ResultState
    Completed ResultState
    MaxValue ResultState
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
    MaxValue: "maxValue",
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
    Project *TeamProjectReference `json:"project,omitempty"`
    // Id of suites to be cloned inside source Test Plan
    SuiteIds *[]int `json:"suiteIds,omitempty"`
}

// Source Test Suite information for Test Suite clone operation
type SourceTestSuiteInfo struct {
    // Id of the Source Test Suite
    Id *int `json:"id,omitempty"`
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
    Suite SuiteEntryTypes
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

// Option to get details in response
type SuiteExpand string

type suiteExpandValuesType struct {
    None SuiteExpand
    Children SuiteExpand
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

// Represents a shallow reference to a TeamProject.
type TeamProjectReference struct {
    // Project abbreviation.
    Abbreviation *string `json:"abbreviation,omitempty"`
    // Url to default team identity image.
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    // The project's description (if any).
    Description *string `json:"description,omitempty"`
    // Project identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Project last update time.
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    // Project name.
    Name *string `json:"name,omitempty"`
    // Project revision.
    Revision *uint64 `json:"revision,omitempty"`
    // Project state.
    State *ProjectState `json:"state,omitempty"`
    // Url to the full version of the object.
    Url *string `json:"url,omitempty"`
    // Project visibility.
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

// Test Case Class
type TestCase struct {
    // Reference links
    Links *ReferenceLinks `json:"links,omitempty"`
    // Order of the TestCase in the Suite
    Order *int `json:"order,omitempty"`
    // List of Points associated with the Test Case
    PointAssignments *[]PointAssignment `json:"pointAssignments,omitempty"`
    // Project under which the Test Case is
    Project *TeamProjectReference `json:"project,omitempty"`
    // Test Plan under which the Test Case is
    TestPlan *TestPlanReference `json:"testPlan,omitempty"`
    // Test Suite under which the Test Case is
    TestSuite *TestSuiteReference `json:"testSuite,omitempty"`
    // Work Item details of the TestCase
    WorkItem *WorkItemDetails `json:"workItem,omitempty"`
}

// Test Case Reference
type TestCaseReference struct {
    // Identity to whom the test case is assigned
    AssignedTo *IdentityRef `json:"assignedTo,omitempty"`
    // Test Case Id
    Id *int `json:"id,omitempty"`
    // Test Case Name
    Name *string `json:"name,omitempty"`
    // State of the test case work item
    State *string `json:"state,omitempty"`
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
    State *TestConfigurationState `json:"state,omitempty"`
    // Dictionary of Test Variable, Selected Value
    Values *[]NameValuePair `json:"values,omitempty"`
    // Id of the configuration
    Id *int `json:"id,omitempty"`
    // Id of the test configuration variable
    Project *TeamProjectReference `json:"project,omitempty"`
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
    State *TestConfigurationState `json:"state,omitempty"`
    // Dictionary of Test Variable, Selected Value
    Values *[]NameValuePair `json:"values,omitempty"`
}

// Test Configuration Reference
type TestConfigurationReference struct {
    // Id of the configuration
    Id *int `json:"id,omitempty"`
    // Name of the configuration
    Name *string `json:"name,omitempty"`
}

// Represents the state of an ITestConfiguration object.
type TestConfigurationState string

type testConfigurationStateValuesType struct {
    Active TestConfigurationState
    Inactive TestConfigurationState
}

var TestConfigurationStateValues = testConfigurationStateValuesType{
    // The configuration can be used for new test runs.
    Active: "active",
    // The configuration has been retired and should not be used for new test runs.
    Inactive: "inactive",
}

// Test environment Detail.
type TestEnvironment struct {
    // Test Environment Id.
    EnvironmentId *uuid.UUID `json:"environmentId,omitempty"`
    // Test Environment Name.
    EnvironmentName *string `json:"environmentName,omitempty"`
}

// Test outcome settings
type TestOutcomeSettings struct {
    // Value to configure how test outcomes for the same tests across suites are shown
    SyncOutcomeAcrossSuites *bool `json:"syncOutcomeAcrossSuites,omitempty"`
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
    EndDate *time.Time `json:"endDate,omitempty"`
    // Iteration path of the test plan.
    Iteration *string `json:"iteration,omitempty"`
    // Name of the test plan.
    Name *string `json:"name,omitempty"`
    // Owner of the test plan.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Release Environment to be used to deploy the build and run automated tests from this test plan.
    ReleaseEnvironmentDefinition *ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
    // Start date for the test plan.
    StartDate *time.Time `json:"startDate,omitempty"`
    // State of the test plan.
    State *string `json:"state,omitempty"`
    // Value to configure how same tests across test suites under a test plan need to behave
    TestOutcomeSettings *TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
    // Revision of the test plan.
    Revision *int `json:"revision,omitempty"`
    // Relevant links
    Links *ReferenceLinks `json:"_links,omitempty"`
    // ID of the test plan.
    Id *int `json:"id,omitempty"`
    // Previous build Id associated with the test plan
    PreviousBuildId *int `json:"previousBuildId,omitempty"`
    // Project which contains the test plan.
    Project *TeamProjectReference `json:"project,omitempty"`
    // Root test suite of the test plan.
    RootSuite *TestSuiteReference `json:"rootSuite,omitempty"`
    // Identity Reference for the last update of the test plan
    UpdatedBy *IdentityRef `json:"updatedBy,omitempty"`
    // Updated date of the test plan
    UpdatedDate *time.Time `json:"updatedDate,omitempty"`
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
    EndDate *time.Time `json:"endDate,omitempty"`
    // Iteration path of the test plan.
    Iteration *string `json:"iteration,omitempty"`
    // Name of the test plan.
    Name *string `json:"name,omitempty"`
    // Owner of the test plan.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Release Environment to be used to deploy the build and run automated tests from this test plan.
    ReleaseEnvironmentDefinition *ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
    // Start date for the test plan.
    StartDate *time.Time `json:"startDate,omitempty"`
    // State of the test plan.
    State *string `json:"state,omitempty"`
    // Value to configure how same tests across test suites under a test plan need to behave
    TestOutcomeSettings *TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
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
    EndDate *time.Time `json:"endDate,omitempty"`
    // Iteration path of the test plan.
    Iteration *string `json:"iteration,omitempty"`
    // Root Suite Id
    RootSuiteId *int `json:"rootSuiteId,omitempty"`
    // Start date for the test plan.
    StartDate *time.Time `json:"startDate,omitempty"`
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
    IsAdvancedExtensionEnabled *bool `json:"isAdvancedExtensionEnabled,omitempty"`
    SelectedSuiteId *int `json:"selectedSuiteId,omitempty"`
    SelectedSuiteName *string `json:"selectedSuiteName,omitempty"`
    TestCasePageSize *int `json:"testCasePageSize,omitempty"`
    TestCases *[]TestCase `json:"testCases,omitempty"`
    TestCasesContinuationToken *string `json:"testCasesContinuationToken,omitempty"`
    TestPlan *TestPlanDetailedReference `json:"testPlan,omitempty"`
    TestPointPageSize *int `json:"testPointPageSize,omitempty"`
    TestPoints *[]TestPoint `json:"testPoints,omitempty"`
    TestPointsContinuationToken *string `json:"testPointsContinuationToken,omitempty"`
    TestSuites *[]TestSuite `json:"testSuites,omitempty"`
    TestSuitesContinuationToken *string `json:"testSuitesContinuationToken,omitempty"`
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
    EndDate *time.Time `json:"endDate,omitempty"`
    // Iteration path of the test plan.
    Iteration *string `json:"iteration,omitempty"`
    // Name of the test plan.
    Name *string `json:"name,omitempty"`
    // Owner of the test plan.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Release Environment to be used to deploy the build and run automated tests from this test plan.
    ReleaseEnvironmentDefinition *ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
    // Start date for the test plan.
    StartDate *time.Time `json:"startDate,omitempty"`
    // State of the test plan.
    State *string `json:"state,omitempty"`
    // Value to configure how same tests across test suites under a test plan need to behave
    TestOutcomeSettings *TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
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
    LastResetToActive *time.Time `json:"lastResetToActive,omitempty"`
    // Last Updated details for the Test Point
    LastUpdatedBy *IdentityRef `json:"lastUpdatedBy,omitempty"`
    // Last Update Time Stamp for the Test Point
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Reference links
    Links *ReferenceLinks `json:"links,omitempty"`
    // Project under which the Test Point is
    Project *TeamProjectReference `json:"project,omitempty"`
    // Results associated to the Test Point
    Results *TestPointResults `json:"results,omitempty"`
    // Test Case Reference
    TestCaseReference *TestCaseReference `json:"testCaseReference,omitempty"`
    // Tester associated with the Test Point
    Tester *IdentityRef `json:"tester,omitempty"`
    // Test Plan under which the Test Point is
    TestPlan *TestPlanReference `json:"testPlan,omitempty"`
    // Test Suite under which the Test Point is
    TestSuite *TestSuiteReference `json:"testSuite,omitempty"`
}

// Test Point Count
type TestPointCount struct {
    // Test Point Count
    Count *int `json:"count,omitempty"`
    // Test Plan under which the Test Points are
    TestPlanId *int `json:"testPlanId,omitempty"`
    // Test Suite under which the Test Points are
    TestSuiteId *int `json:"testSuiteId,omitempty"`
}

// Test Point Results
type TestPointResults struct {
    // Failure Type for the Test Point
    FailureType *FailureType `json:"failureType,omitempty"`
    // Last Resolution State Id for the TEst Point
    LastResolutionState *LastResolutionState `json:"lastResolutionState,omitempty"`
    // Last Result Details for the Test Point
    LastResultDetails *LastResultDetails `json:"lastResultDetails,omitempty"`
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
    Tester *IdentityRef `json:"tester,omitempty"`
}

// Represents the test settings of the run. Used to create test settings and fetch test settings
type TestSettings struct {
    // Area path required to create test settings
    AreaPath *string `json:"areaPath,omitempty"`
    // Description of the test settings. Used in create test settings.
    Description *string `json:"description,omitempty"`
    // Indicates if the tests settings is public or private.Used in create test settings.
    IsPublic *bool `json:"isPublic,omitempty"`
    // Xml string of machine roles. Used in create test settings.
    MachineRoles *string `json:"machineRoles,omitempty"`
    // Test settings content.
    TestSettingsContent *string `json:"testSettingsContent,omitempty"`
    // Test settings id.
    TestSettingsId *int `json:"testSettingsId,omitempty"`
    // Test settings name.
    TestSettingsName *string `json:"testSettingsName,omitempty"`
}

// Test suite
type TestSuite struct {
    // Test suite default configurations.
    DefaultConfigurations *[]TestConfigurationReference `json:"defaultConfigurations,omitempty"`
    // Test suite default testers.
    DefaultTesters *[]IdentityRef `json:"defaultTesters,omitempty"`
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
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Child test suites of current test suite.
    Children *[]TestSuite `json:"children,omitempty"`
    // Boolean value dictating if Child test suites are present
    HasChildren *bool `json:"hasChildren,omitempty"`
    // Id of test suite.
    Id *int `json:"id,omitempty"`
    // Last error for test suite.
    LastError *string `json:"lastError,omitempty"`
    // Last populated date.
    LastPopulatedDate *time.Time `json:"lastPopulatedDate,omitempty"`
    // IdentityRef of user who has updated test suite recently.
    LastUpdatedBy *IdentityRef `json:"lastUpdatedBy,omitempty"`
    // Last update date.
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Test plan to which the test suite belongs.
    Plan *TestPlanReference `json:"plan,omitempty"`
    // Test suite project shallow reference.
    Project *TeamProjectReference `json:"project,omitempty"`
    // Test suite revision.
    Revision *int `json:"revision,omitempty"`
}

// Test suite Create Parameters
type TestSuiteCreateParams struct {
    // Test suite default configurations.
    DefaultConfigurations *[]TestConfigurationReference `json:"defaultConfigurations,omitempty"`
    // Test suite default testers.
    DefaultTesters *[]IdentityRef `json:"defaultTesters,omitempty"`
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
    DefaultTesters *[]IdentityRef `json:"defaultTesters,omitempty"`
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
    Project *TeamProjectReference `json:"project,omitempty"`
}

// Type of TestSuite
type TestSuiteType string

type testSuiteTypeValuesType struct {
    None TestSuiteType
    DynamicTestSuite TestSuiteType
    StaticTestSuite TestSuiteType
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
    DefaultTesters *[]IdentityRef `json:"defaultTesters,omitempty"`
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
    Project *TeamProjectReference `json:"project,omitempty"`
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
    InProgress UserFriendlyTestOutcome
    Blocked UserFriendlyTestOutcome
    Failed UserFriendlyTestOutcome
    Passed UserFriendlyTestOutcome
    Ready UserFriendlyTestOutcome
    NotApplicable UserFriendlyTestOutcome
    Paused UserFriendlyTestOutcome
    MaxValue UserFriendlyTestOutcome
}

var UserFriendlyTestOutcomeValues = userFriendlyTestOutcomeValuesType{
    InProgress: "inProgress",
    Blocked: "blocked",
    Failed: "failed",
    Passed: "passed",
    Ready: "ready",
    NotApplicable: "notApplicable",
    Paused: "paused",
    MaxValue: "maxValue",
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
