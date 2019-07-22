// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testResults

import (
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "time"
)

type AggregatedDataForResultTrend struct {
    // This is tests execution duration.
    Duration *interface{} `json:"duration,omitempty"`
    ResultsByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsByOutcome,omitempty"`
    RunSummaryByState *map[TestRunState]AggregatedRunsByState `json:"runSummaryByState,omitempty"`
    TestResultsContext *TestResultsContext `json:"testResultsContext,omitempty"`
    TotalTests *int `json:"totalTests,omitempty"`
}

type AggregatedResultsAnalysis struct {
    Duration *interface{} `json:"duration,omitempty"`
    NotReportedResultsByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"notReportedResultsByOutcome,omitempty"`
    PreviousContext *TestResultsContext `json:"previousContext,omitempty"`
    ResultsByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsByOutcome,omitempty"`
    ResultsDifference *AggregatedResultsDifference `json:"resultsDifference,omitempty"`
    RunSummaryByOutcome *map[TestRunOutcome]AggregatedRunsByOutcome `json:"runSummaryByOutcome,omitempty"`
    RunSummaryByState *map[TestRunState]AggregatedRunsByState `json:"runSummaryByState,omitempty"`
    TotalTests *int `json:"totalTests,omitempty"`
}

type AggregatedResultsByOutcome struct {
    Count *int `json:"count,omitempty"`
    Duration *interface{} `json:"duration,omitempty"`
    GroupByField *string `json:"groupByField,omitempty"`
    GroupByValue *interface{} `json:"groupByValue,omitempty"`
    Outcome *TestOutcome `json:"outcome,omitempty"`
    RerunResultCount *int `json:"rerunResultCount,omitempty"`
}

type AggregatedResultsDifference struct {
    IncreaseInDuration *interface{} `json:"increaseInDuration,omitempty"`
    IncreaseInFailures *int `json:"increaseInFailures,omitempty"`
    IncreaseInOtherTests *int `json:"increaseInOtherTests,omitempty"`
    IncreaseInPassedTests *int `json:"increaseInPassedTests,omitempty"`
    IncreaseInTotalTests *int `json:"increaseInTotalTests,omitempty"`
}

type AggregatedRunsByOutcome struct {
    Outcome *TestRunOutcome `json:"outcome,omitempty"`
    RunsCount *int `json:"runsCount,omitempty"`
}

type AggregatedRunsByState struct {
    ResultsByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsByOutcome,omitempty"`
    RunsCount *int `json:"runsCount,omitempty"`
    State *TestRunState `json:"state,omitempty"`
}

type Attachment struct {
    CompressionType *string `json:"compressionType,omitempty"`
    FileName *string `json:"fileName,omitempty"`
    Stream *interface{} `json:"stream,omitempty"`
}

// The types of test attachments.
type AttachmentType string

type attachmentTypeValuesType struct {
    GeneralAttachment AttachmentType
    CodeCoverage AttachmentType
    ConsoleLog AttachmentType
}

var AttachmentTypeValues = attachmentTypeValuesType{
    // Attachment type GeneralAttachment , use this as default type unless you have other type.
    GeneralAttachment: "generalAttachment",
    // Attachment type CodeCoverage.
    CodeCoverage: "codeCoverage",
    // Attachment type ConsoleLog.
    ConsoleLog: "consoleLog",
}

// BuildConfiguration Details.
type BuildConfiguration struct {
    // Branch name for which build is generated.
    BranchName *string `json:"branchName,omitempty"`
    // BuildDefinitionId for build.
    BuildDefinitionId *int `json:"buildDefinitionId,omitempty"`
    // Build system.
    BuildSystem *string `json:"buildSystem,omitempty"`
    // Build Creation Date.
    CreationDate *time.Time `json:"creationDate,omitempty"`
    // Build flavor (eg Build/Release).
    Flavor *string `json:"flavor,omitempty"`
    // BuildConfiguration Id.
    Id *int `json:"id,omitempty"`
    // Build Number.
    Number *string `json:"number,omitempty"`
    // BuildConfiguration Platform.
    Platform *string `json:"platform,omitempty"`
    // Project associated with this BuildConfiguration.
    Project *ShallowReference `json:"project,omitempty"`
    // Repository Guid for the Build.
    RepositoryGuid *string `json:"repositoryGuid,omitempty"`
    // Deprecated: Use RepositoryGuid instead
    RepositoryId *int `json:"repositoryId,omitempty"`
    // Repository Type (eg. TFSGit).
    RepositoryType *string `json:"repositoryType,omitempty"`
    // Source Version(/first commit) for the build was triggered.
    SourceVersion *string `json:"sourceVersion,omitempty"`
    // Target BranchName.
    TargetBranchName *string `json:"targetBranchName,omitempty"`
    // Build Uri.
    Uri *string `json:"uri,omitempty"`
}

// Build Coverage Detail
type BuildCoverage struct {
    // Code Coverage File Url
    CodeCoverageFileUrl *string `json:"codeCoverageFileUrl,omitempty"`
    // Build Configuration
    Configuration *BuildConfiguration `json:"configuration,omitempty"`
    // Last Error
    LastError *string `json:"lastError,omitempty"`
    // List of Modules
    Modules *[]ModuleCoverage `json:"modules,omitempty"`
    // State
    State *string `json:"state,omitempty"`
}

// Reference to a build.
type BuildReference struct {
    // Branch name.
    BranchName *string `json:"branchName,omitempty"`
    // Build system.
    BuildSystem *string `json:"buildSystem,omitempty"`
    // Build Definition ID.
    DefinitionId *int `json:"definitionId,omitempty"`
    // Build ID.
    Id *int `json:"id,omitempty"`
    // Build Number.
    Number *string `json:"number,omitempty"`
    // Repository ID.
    RepositoryId *string `json:"repositoryId,omitempty"`
    // Build URI.
    Uri *string `json:"uri,omitempty"`
}

// Represents the build configuration (platform, flavor) and coverage data for the build
type CodeCoverageData struct {
    // Flavor of build for which data is retrieved/published
    BuildFlavor *string `json:"buildFlavor,omitempty"`
    // Platform of build for which data is retrieved/published
    BuildPlatform *string `json:"buildPlatform,omitempty"`
    // List of coverage data for the build
    CoverageStats *[]CodeCoverageStatistics `json:"coverageStats,omitempty"`
}

// Represents the code coverage statistics for a particular coverage label (modules, statements, blocks, etc.)
type CodeCoverageStatistics struct {
    // Covered units
    Covered *int `json:"covered,omitempty"`
    // Delta of coverage
    Delta *float64 `json:"delta,omitempty"`
    // Is delta valid
    IsDeltaAvailable *bool `json:"isDeltaAvailable,omitempty"`
    // Label of coverage data ("Blocks", "Statements", "Modules", etc.)
    Label *string `json:"label,omitempty"`
    // Position of label
    Position *int `json:"position,omitempty"`
    // Total units
    Total *int `json:"total,omitempty"`
}

// Represents the code coverage summary results Used to publish or retrieve code coverage summary against a build
type CodeCoverageSummary struct {
    // Uri of build for which data is retrieved/published
    Build *ShallowReference `json:"build,omitempty"`
    // List of coverage data and details for the build
    CoverageData *[]CodeCoverageData `json:"coverageData,omitempty"`
    // Uri of build against which difference in coverage is computed
    DeltaBuild *ShallowReference `json:"deltaBuild,omitempty"`
    // Uri of build against which difference in coverage is computed
    Status *CoverageSummaryStatus `json:"status,omitempty"`
}

type CoverageStatistics struct {
    BlocksCovered *int `json:"blocksCovered,omitempty"`
    BlocksNotCovered *int `json:"blocksNotCovered,omitempty"`
    LinesCovered *int `json:"linesCovered,omitempty"`
    LinesNotCovered *int `json:"linesNotCovered,omitempty"`
    LinesPartiallyCovered *int `json:"linesPartiallyCovered,omitempty"`
}

// Represents status of code coverage summary for a build
type CoverageSummaryStatus string

type coverageSummaryStatusValuesType struct {
    None CoverageSummaryStatus
    InProgress CoverageSummaryStatus
    Completed CoverageSummaryStatus
    Finalized CoverageSummaryStatus
    Pending CoverageSummaryStatus
}

var CoverageSummaryStatusValues = coverageSummaryStatusValuesType{
    // No coverage status
    None: "none",
    // The summary evaluation is in progress
    InProgress: "inProgress",
    // The summary evaluation for the previous request is completed. Summary can change in future
    Completed: "completed",
    // The summary evaluation is finalized and won't change
    Finalized: "finalized",
    // The summary evaluation is pending
    Pending: "pending",
}

// A custom field information. Allowed Key : Value pairs - ( AttemptId: int value, IsTestResultFlaky: bool)
type CustomTestField struct {
    // Field Name.
    FieldName *string `json:"fieldName,omitempty"`
    // Field value.
    Value *interface{} `json:"value,omitempty"`
}

// This is a temporary class to provide the details for the test run environment.
type DtlEnvironmentDetails struct {
    CsmContent *string `json:"csmContent,omitempty"`
    CsmParameters *string `json:"csmParameters,omitempty"`
    SubscriptionName *string `json:"subscriptionName,omitempty"`
}

// Failing since information of a test result.
type FailingSince struct {
    // Build reference since failing.
    Build *BuildReference `json:"build,omitempty"`
    // Time since failing.
    Date *time.Time `json:"date,omitempty"`
    // Release reference since failing.
    Release *ReleaseReference `json:"release,omitempty"`
}

type FieldDetailsForTestResults struct {
    // Group by field name
    FieldName *string `json:"fieldName,omitempty"`
    // Group by field values
    GroupsForField *[]interface{} `json:"groupsForField,omitempty"`
}

type FileCoverageRequest struct {
    FilePath *string `json:"filePath,omitempty"`
    PullRequestBaseIterationId *int `json:"pullRequestBaseIterationId,omitempty"`
    PullRequestId *int `json:"pullRequestId,omitempty"`
    PullRequestIterationId *int `json:"pullRequestIterationId,omitempty"`
    RepoId *string `json:"repoId,omitempty"`
}

type FlakyDetection struct {
    // FlakyDetectionPipelines defines Pipelines for Detection.
    FlakyDetectionPipelines *FlakyDetectionPipelines `json:"flakyDetectionPipelines,omitempty"`
    // FlakyDetectionType defines Detection type i.e. 1. System or 2. Manual.
    FlakyDetectionType *FlakyDetectionType `json:"flakyDetectionType,omitempty"`
}

type FlakyDetectionPipelines struct {
    // AllowedPipelines - List All Pipelines allowed for detection.
    AllowedPipelines *[]int `json:"allowedPipelines,omitempty"`
    // IsAllPipelinesAllowed if users configure all system's pipelines.
    IsAllPipelinesAllowed *bool `json:"isAllPipelinesAllowed,omitempty"`
}

type FlakyDetectionType string

type flakyDetectionTypeValuesType struct {
    Custom FlakyDetectionType
    System FlakyDetectionType
}

var FlakyDetectionTypeValues = flakyDetectionTypeValuesType{
    // Custom defines manual detection type.
    Custom: "custom",
    // Defines System detection type.
    System: "system",
}

type FlakySettings struct {
    // FlakyDetection defines types of detection.
    FlakyDetection *FlakyDetection `json:"flakyDetection,omitempty"`
    // FlakyInSummaryReport defines flaky data should show in summary report or not.
    FlakyInSummaryReport *bool `json:"flakyInSummaryReport,omitempty"`
    // ManualMarkUnmarkFlaky defines manual marking unmarking of flaky testcase.
    ManualMarkUnmarkFlaky *bool `json:"manualMarkUnmarkFlaky,omitempty"`
}

type FunctionCoverage struct {
    Class *string `json:"class,omitempty"`
    Name *string `json:"name,omitempty"`
    Namespace *string `json:"namespace,omitempty"`
    SourceFile *string `json:"sourceFile,omitempty"`
    Statistics *CoverageStatistics `json:"statistics,omitempty"`
}

type GraphSubjectBase struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Descriptor *string `json:"descriptor,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    Url *string `json:"url,omitempty"`
}

type IdentityRef struct {
    Links *ReferenceLinks `json:"_links,omitempty"`
    Descriptor *string `json:"descriptor,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    Url *string `json:"url,omitempty"`
    DirectoryAlias *string `json:"directoryAlias,omitempty"`
    Id *string `json:"id,omitempty"`
    ImageUrl *string `json:"imageUrl,omitempty"`
    Inactive *bool `json:"inactive,omitempty"`
    IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
    IsContainer *bool `json:"isContainer,omitempty"`
    IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
    ProfileUrl *string `json:"profileUrl,omitempty"`
    UniqueName *string `json:"uniqueName,omitempty"`
}

// Job in pipeline. This is related to matrixing in YAML.
type JobReference struct {
    // Attempt number of the job
    Attempt *int `json:"attempt,omitempty"`
    // Matrixing in YAML generates copies of a job with different inputs in matrix. JobName is the name of those input. Maximum supported length for name is 256 character.
    JobName *string `json:"jobName,omitempty"`
}

type ModuleCoverage struct {
    BlockCount *int `json:"blockCount,omitempty"`
    BlockData *[]byte `json:"blockData,omitempty"`
    // Code Coverage File Url
    FileUrl *string `json:"fileUrl,omitempty"`
    Functions *[]FunctionCoverage `json:"functions,omitempty"`
    Name *string `json:"name,omitempty"`
    Signature *uuid.UUID `json:"signature,omitempty"`
    SignatureAge *int `json:"signatureAge,omitempty"`
    Statistics *CoverageStatistics `json:"statistics,omitempty"`
}

type OperationType string

type operationTypeValuesType struct {
    Add OperationType
    Delete OperationType
}

var OperationTypeValues = operationTypeValuesType{
    Add: "add",
    Delete: "delete",
}

// Phase in pipeline
type PhaseReference struct {
    // Attempt number of the phase
    Attempt *int `json:"attempt,omitempty"`
    // Name of the phase. Maximum supported length for name is 256 character.
    PhaseName *string `json:"phaseName,omitempty"`
}

// Pipeline reference
type PipelineReference struct {
    // Reference of the job
    JobReference *JobReference `json:"jobReference,omitempty"`
    // Reference of the phase.
    PhaseReference *PhaseReference `json:"phaseReference,omitempty"`
    // Reference of the pipeline with which this pipeline instance is related.
    PipelineId *int `json:"pipelineId,omitempty"`
    // Reference of the stage.
    StageReference *StageReference `json:"stageReference,omitempty"`
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
    Deleting: "deleting",
    New: "new",
    WellFormed: "wellFormed",
    CreatePending: "createPending",
    All: "all",
    Unchanged: "unchanged",
    Deleted: "deleted",
}

type ProjectVisibility string

type projectVisibilityValuesType struct {
    Private ProjectVisibility
    Public ProjectVisibility
}

var ProjectVisibilityValues = projectVisibilityValuesType{
    Private: "private",
    Public: "public",
}

type QueryModel struct {
    Query *string `json:"query,omitempty"`
}

type ReferenceLinks struct {
    Links *map[string]interface{} `json:"links,omitempty"`
}

// Reference to a release.
type ReleaseReference struct {
    // Number of Release Attempt.
    Attempt *int `json:"attempt,omitempty"`
    // Release Creation Date.
    CreationDate *time.Time `json:"creationDate,omitempty"`
    // Release definition ID.
    DefinitionId *int `json:"definitionId,omitempty"`
    // Environment creation Date.
    EnvironmentCreationDate *time.Time `json:"environmentCreationDate,omitempty"`
    // Release environment definition ID.
    EnvironmentDefinitionId *int `json:"environmentDefinitionId,omitempty"`
    // Release environment definition name.
    EnvironmentDefinitionName *string `json:"environmentDefinitionName,omitempty"`
    // Release environment ID.
    EnvironmentId *int `json:"environmentId,omitempty"`
    // Release environment name.
    EnvironmentName *string `json:"environmentName,omitempty"`
    // Release ID.
    Id *int `json:"id,omitempty"`
    // Release name.
    Name *string `json:"name,omitempty"`
}

// Additional details with test result
type ResultDetails string

type resultDetailsValuesType struct {
    None ResultDetails
    Iterations ResultDetails
    WorkItems ResultDetails
    SubResults ResultDetails
    Point ResultDetails
}

var ResultDetailsValues = resultDetailsValuesType{
    // Core fields of test result. Core fields includes State, Outcome, Priority, AutomatedTestName, AutomatedTestStorage, Comments, ErrorMessage etc.
    None: "none",
    // Test iteration details in a test result.
    Iterations: "iterations",
    // Workitems associated with a test result.
    WorkItems: "workItems",
    // Subresults in a test result.
    SubResults: "subResults",
    // Point and plan detail in a test result.
    Point: "point",
}

// Hierarchy type of the result/subresults.
type ResultGroupType string

type resultGroupTypeValuesType struct {
    None ResultGroupType
    Rerun ResultGroupType
    DataDriven ResultGroupType
    OrderedTest ResultGroupType
    Generic ResultGroupType
}

var ResultGroupTypeValues = resultGroupTypeValuesType{
    // Leaf node of test result.
    None: "none",
    // Hierarchy type of test result.
    Rerun: "rerun",
    // Hierarchy type of test result.
    DataDriven: "dataDriven",
    // Hierarchy type of test result.
    OrderedTest: "orderedTest",
    // Unknown hierarchy type.
    Generic: "generic",
}

// Additional details with test result metadata
type ResultMetaDataDetails string

type resultMetaDataDetailsValuesType struct {
    None ResultMetaDataDetails
    FlakyIdentifiers ResultMetaDataDetails
}

var ResultMetaDataDetailsValues = resultMetaDataDetailsValuesType{
    // Core fields of test result metadata.
    None: "none",
    // Test FlakyIdentifiers details in test result metadata.
    FlakyIdentifiers: "flakyIdentifiers",
}

type ResultsFilter struct {
    AutomatedTestName *string `json:"automatedTestName,omitempty"`
    Branch *string `json:"branch,omitempty"`
    ExecutedIn *Service `json:"executedIn,omitempty"`
    GroupBy *string `json:"groupBy,omitempty"`
    MaxCompleteDate *time.Time `json:"maxCompleteDate,omitempty"`
    ResultsCount *int `json:"resultsCount,omitempty"`
    TestCaseId *int `json:"testCaseId,omitempty"`
    TestCaseReferenceIds *[]int `json:"testCaseReferenceIds,omitempty"`
    TestPlanId *int `json:"testPlanId,omitempty"`
    TestPointIds *[]int `json:"testPointIds,omitempty"`
    TestResultsContext *TestResultsContext `json:"testResultsContext,omitempty"`
    TrendDays *int `json:"trendDays,omitempty"`
}

// Test run create details.
type RunCreateModel struct {
    // true if test run is automated, false otherwise. By default it will be false.
    Automated *bool `json:"automated,omitempty"`
    // An abstracted reference to the build that it belongs.
    Build *ShallowReference `json:"build,omitempty"`
    // Drop location of the build used for test run.
    BuildDropLocation *string `json:"buildDropLocation,omitempty"`
    // Flavor of the build used for test run. (E.g: Release, Debug)
    BuildFlavor *string `json:"buildFlavor,omitempty"`
    // Platform of the build used for test run. (E.g.: x86, amd64)
    BuildPlatform *string `json:"buildPlatform,omitempty"`
    // BuildReference of the test run.
    BuildReference *BuildConfiguration `json:"buildReference,omitempty"`
    // Comments entered by those analyzing the run.
    Comment *string `json:"comment,omitempty"`
    // Completed date time of the run.
    CompleteDate *string `json:"completeDate,omitempty"`
    // IDs of the test configurations associated with the run.
    ConfigurationIds *[]int `json:"configurationIds,omitempty"`
    // Name of the test controller used for automated run.
    Controller *string `json:"controller,omitempty"`
    // Additional properties of test Run.
    CustomTestFields *[]CustomTestField `json:"customTestFields,omitempty"`
    // An abstracted reference to DtlAutEnvironment.
    DtlAutEnvironment *ShallowReference `json:"dtlAutEnvironment,omitempty"`
    // An abstracted reference to DtlTestEnvironment.
    DtlTestEnvironment *ShallowReference `json:"dtlTestEnvironment,omitempty"`
    // Due date and time for test run.
    DueDate *string `json:"dueDate,omitempty"`
    EnvironmentDetails *DtlEnvironmentDetails `json:"environmentDetails,omitempty"`
    // Error message associated with the run.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // Filter used for discovering the Run.
    Filter *RunFilter `json:"filter,omitempty"`
    // The iteration in which to create the run. Root iteration of the team project will be default
    Iteration *string `json:"iteration,omitempty"`
    // Name of the test run.
    Name *string `json:"name,omitempty"`
    // Display name of the owner of the run.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Reference of the pipeline to which this test run belongs. PipelineReference.PipelineId should be equal to RunCreateModel.Build.Id
    PipelineReference *PipelineReference `json:"pipelineReference,omitempty"`
    // An abstracted reference to the plan that it belongs.
    Plan *ShallowReference `json:"plan,omitempty"`
    // IDs of the test points to use in the run.
    PointIds *[]int `json:"pointIds,omitempty"`
    // URI of release environment associated with the run.
    ReleaseEnvironmentUri *string `json:"releaseEnvironmentUri,omitempty"`
    // Reference to release associated with test run.
    ReleaseReference *ReleaseReference `json:"releaseReference,omitempty"`
    // URI of release associated with the run.
    ReleaseUri *string `json:"releaseUri,omitempty"`
    // Run summary for run Type = NoConfigRun.
    RunSummary *[]RunSummaryModel `json:"runSummary,omitempty"`
    // Timespan till the run times out.
    RunTimeout *interface{} `json:"runTimeout,omitempty"`
    // SourceWorkFlow(CI/CD) of the test run.
    SourceWorkflow *string `json:"sourceWorkflow,omitempty"`
    // Start date time of the run.
    StartDate *string `json:"startDate,omitempty"`
    // The state of the run. Type TestRunState Valid states - Unspecified ,NotStarted, InProgress, Completed, Waiting, Aborted, NeedsInvestigation
    State *string `json:"state,omitempty"`
    // Tags to attach with the test run, maximum of 5 tags can be added to run.
    Tags *[]TestTag `json:"tags,omitempty"`
    // TestConfigurationMapping of the test run.
    TestConfigurationsMapping *string `json:"testConfigurationsMapping,omitempty"`
    // ID of the test environment associated with the run.
    TestEnvironmentId *string `json:"testEnvironmentId,omitempty"`
    // An abstracted reference to the test settings resource.
    TestSettings *ShallowReference `json:"testSettings,omitempty"`
    // Type of the run(RunType) Valid Values : (Unspecified, Normal, Blocking, Web, MtrRunInitiatedFromWeb, RunWithDtlEnv, NoConfigRun)
    Type *string `json:"type,omitempty"`
}

// This class is used to provide the filters used for discovery
type RunFilter struct {
    // filter for the test case sources (test containers)
    SourceFilter *string `json:"sourceFilter,omitempty"`
    // filter for the test cases
    TestCaseFilter *string `json:"testCaseFilter,omitempty"`
}

// Test run statistics per outcome.
type RunStatistic struct {
    // Test result count fo the given outcome.
    Count *int `json:"count,omitempty"`
    // Test result outcome
    Outcome *string `json:"outcome,omitempty"`
    // Test run Resolution State.
    ResolutionState *TestResolutionState `json:"resolutionState,omitempty"`
    // State of the test run
    State *string `json:"state,omitempty"`
}

// Run summary for each output type of test.
type RunSummaryModel struct {
    // Total time taken in milliseconds.
    Duration *uint64 `json:"duration,omitempty"`
    // Number of results for Outcome TestOutcome
    ResultCount *int `json:"resultCount,omitempty"`
    // Summary is based on outcome
    TestOutcome *TestOutcome `json:"testOutcome,omitempty"`
}

type RunUpdateModel struct {
    // An abstracted reference to the build that it belongs.
    Build *ShallowReference `json:"build,omitempty"`
    // Drop location of the build used for test run.
    BuildDropLocation *string `json:"buildDropLocation,omitempty"`
    // Flavor of the build used for test run. (E.g: Release, Debug)
    BuildFlavor *string `json:"buildFlavor,omitempty"`
    // Platform of the build used for test run. (E.g.: x86, amd64)
    BuildPlatform *string `json:"buildPlatform,omitempty"`
    // Comments entered by those analyzing the run.
    Comment *string `json:"comment,omitempty"`
    // Completed date time of the run.
    CompletedDate *string `json:"completedDate,omitempty"`
    // Name of the test controller used for automated run.
    Controller *string `json:"controller,omitempty"`
    // true to delete inProgess Results , false otherwise.
    DeleteInProgressResults *bool `json:"deleteInProgressResults,omitempty"`
    // An abstracted reference to DtlAutEnvironment.
    DtlAutEnvironment *ShallowReference `json:"dtlAutEnvironment,omitempty"`
    // An abstracted reference to DtlEnvironment.
    DtlEnvironment *ShallowReference `json:"dtlEnvironment,omitempty"`
    DtlEnvironmentDetails *DtlEnvironmentDetails `json:"dtlEnvironmentDetails,omitempty"`
    // Due date and time for test run.
    DueDate *string `json:"dueDate,omitempty"`
    // Error message associated with the run.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // The iteration in which to create the run.
    Iteration *string `json:"iteration,omitempty"`
    // Log entries associated with the run. Use a comma-separated list of multiple log entry objects. { logEntry }, { logEntry }, ...
    LogEntries *[]TestMessageLogDetails `json:"logEntries,omitempty"`
    // Name of the test run.
    Name *string `json:"name,omitempty"`
    // URI of release environment associated with the run.
    ReleaseEnvironmentUri *string `json:"releaseEnvironmentUri,omitempty"`
    // URI of release associated with the run.
    ReleaseUri *string `json:"releaseUri,omitempty"`
    // Run summary for run Type = NoConfigRun.
    RunSummary *[]RunSummaryModel `json:"runSummary,omitempty"`
    // SourceWorkFlow(CI/CD) of the test run.
    SourceWorkflow *string `json:"sourceWorkflow,omitempty"`
    // Start date time of the run.
    StartedDate *string `json:"startedDate,omitempty"`
    // The state of the test run Below are the valid values - NotStarted, InProgress, Completed, Aborted, Waiting
    State *string `json:"state,omitempty"`
    // The types of sub states for test run.
    Substate *TestRunSubstate `json:"substate,omitempty"`
    // Tags to attach with the test run.
    Tags *[]TestTag `json:"tags,omitempty"`
    // ID of the test environment associated with the run.
    TestEnvironmentId *string `json:"testEnvironmentId,omitempty"`
    // An abstracted reference to test setting resource.
    TestSettings *ShallowReference `json:"testSettings,omitempty"`
}

type Service string

type serviceValuesType struct {
    Any Service
    Tcm Service
    Tfs Service
}

var ServiceValues = serviceValuesType{
    Any: "any",
    Tcm: "tcm",
    Tfs: "tfs",
}

// An abstracted reference to some other resource. This class is used to provide the build data contracts with a uniform way to reference other resources in a way that provides easy traversal through links.
type ShallowReference struct {
    // ID of the resource
    Id *string `json:"id,omitempty"`
    // Name of the linked resource (definition name, controller name, etc.)
    Name *string `json:"name,omitempty"`
    // Full http link to the resource
    Url *string `json:"url,omitempty"`
}

type ShallowTestCaseResult struct {
    AutomatedTestName *string `json:"automatedTestName,omitempty"`
    AutomatedTestStorage *string `json:"automatedTestStorage,omitempty"`
    DurationInMs *float64 `json:"durationInMs,omitempty"`
    Id *int `json:"id,omitempty"`
    IsReRun *bool `json:"isReRun,omitempty"`
    Outcome *string `json:"outcome,omitempty"`
    Owner *string `json:"owner,omitempty"`
    Priority *int `json:"priority,omitempty"`
    RefId *int `json:"refId,omitempty"`
    RunId *int `json:"runId,omitempty"`
    Tags *[]string `json:"tags,omitempty"`
    TestCaseTitle *string `json:"testCaseTitle,omitempty"`
}

// Reference to shared step workitem.
type SharedStepModel struct {
    // WorkItem shared step ID.
    Id *int `json:"id,omitempty"`
    // Shared step workitem revision.
    Revision *int `json:"revision,omitempty"`
}

// Stage in pipeline
type StageReference struct {
    // Attempt number of stage
    Attempt *int `json:"attempt,omitempty"`
    // Name of the stage. Maximum supported length for name is 256 character.
    StageName *string `json:"stageName,omitempty"`
}

type TeamProjectReference struct {
    Abbreviation *string `json:"abbreviation,omitempty"`
    DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
    Description *string `json:"description,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
    Name *string `json:"name,omitempty"`
    Revision *uint64 `json:"revision,omitempty"`
    State *ProjectState `json:"state,omitempty"`
    Url *string `json:"url,omitempty"`
    Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

// Represents a test step result.
type TestActionResultModel struct {
    // Comment in result.
    Comment *string `json:"comment,omitempty"`
    // Time when execution completed.
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Duration of execution.
    DurationInMs *float64 `json:"durationInMs,omitempty"`
    // Error message in result.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // Test outcome of result.
    Outcome *string `json:"outcome,omitempty"`
    // Time when execution started.
    StartedDate *time.Time `json:"startedDate,omitempty"`
    // Path identifier test step in test case workitem.
    ActionPath *string `json:"actionPath,omitempty"`
    // Iteration ID of test action result.
    IterationId *int `json:"iterationId,omitempty"`
    // Reference to shared step workitem.
    SharedStepModel *SharedStepModel `json:"sharedStepModel,omitempty"`
    // This is step Id of test case. For shared step, it is step Id of shared step in test case workitem; step Id in shared step. Example: TestCase workitem has two steps: 1) Normal step with Id = 1 2) Shared Step with Id = 2. Inside shared step: a) Normal Step with Id = 1 Value for StepIdentifier for First step: "1" Second step: "2;1"
    StepIdentifier *string `json:"stepIdentifier,omitempty"`
    // Url of test action result.
    Url *string `json:"url,omitempty"`
}

type TestAttachment struct {
    // Attachment type.
    AttachmentType *AttachmentType `json:"attachmentType,omitempty"`
    // Comment associated with attachment.
    Comment *string `json:"comment,omitempty"`
    // Attachment created date.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Attachment file name
    FileName *string `json:"fileName,omitempty"`
    // ID of the attachment.
    Id *int `json:"id,omitempty"`
    // Attachment size.
    Size *uint64 `json:"size,omitempty"`
    // Attachment Url.
    Url *string `json:"url,omitempty"`
}

// Reference to test attachment.
type TestAttachmentReference struct {
    // ID of the attachment.
    Id *int `json:"id,omitempty"`
    // Url to download the attachment.
    Url *string `json:"url,omitempty"`
}

// Test attachment request model
type TestAttachmentRequestModel struct {
    // Attachment type By Default it will be GeneralAttachment. It can be one of the following type. { GeneralAttachment, AfnStrip, BugFilingData, CodeCoverage, IntermediateCollectorData, RunConfig, TestImpactDetails, TmiTestRunDeploymentFiles, TmiTestRunReverseDeploymentFiles, TmiTestResultDetail, TmiTestRunSummary }
    AttachmentType *string `json:"attachmentType,omitempty"`
    // Comment associated with attachment
    Comment *string `json:"comment,omitempty"`
    // Attachment filename
    FileName *string `json:"fileName,omitempty"`
    // Base64 encoded file stream
    Stream *string `json:"stream,omitempty"`
}

// Represents a test result.
type TestCaseResult struct {
    // Test attachment ID of action recording.
    AfnStripId *int `json:"afnStripId,omitempty"`
    // Reference to area path of test.
    Area *ShallowReference `json:"area,omitempty"`
    // Reference to bugs linked to test result.
    AssociatedBugs *[]ShallowReference `json:"associatedBugs,omitempty"`
    // ID representing test method in a dll.
    AutomatedTestId *string `json:"automatedTestId,omitempty"`
    // Fully qualified name of test executed.
    AutomatedTestName *string `json:"automatedTestName,omitempty"`
    // Container to which test belongs.
    AutomatedTestStorage *string `json:"automatedTestStorage,omitempty"`
    // Type of automated test.
    AutomatedTestType *string `json:"automatedTestType,omitempty"`
    // TypeId of automated test.
    AutomatedTestTypeId *string `json:"automatedTestTypeId,omitempty"`
    // Shallow reference to build associated with test result.
    Build *ShallowReference `json:"build,omitempty"`
    // Reference to build associated with test result.
    BuildReference *BuildReference `json:"buildReference,omitempty"`
    // Comment in a test result with maxSize= 1000 chars.
    Comment *string `json:"comment,omitempty"`
    // Time when test execution completed. Completed date should be greater than StartedDate.
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Machine name where test executed.
    ComputerName *string `json:"computerName,omitempty"`
    // Reference to test configuration. Type ShallowReference.
    Configuration *ShallowReference `json:"configuration,omitempty"`
    // Timestamp when test result created.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Additional properties of test result.
    CustomFields *[]CustomTestField `json:"customFields,omitempty"`
    // Duration of test execution in milliseconds. If not provided value will be set as CompletedDate - StartedDate
    DurationInMs *float64 `json:"durationInMs,omitempty"`
    // Error message in test execution.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // Information when test results started failing.
    FailingSince *FailingSince `json:"failingSince,omitempty"`
    // Failure type of test result. Valid Value= (Known Issue, New Issue, Regression, Unknown, None)
    FailureType *string `json:"failureType,omitempty"`
    // ID of a test result.
    Id *int `json:"id,omitempty"`
    // Test result details of test iterations used only for Manual Testing.
    IterationDetails *[]TestIterationDetailsModel `json:"iterationDetails,omitempty"`
    // Reference to identity last updated test result.
    LastUpdatedBy *IdentityRef `json:"lastUpdatedBy,omitempty"`
    // Last updated datetime of test result.
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Test outcome of test result. Valid values = (Unspecified, None, Passed, Failed, Inconclusive, Timeout, Aborted, Blocked, NotExecuted, Warning, Error, NotApplicable, Paused, InProgress, NotImpacted)
    Outcome *string `json:"outcome,omitempty"`
    // Reference to test owner.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Priority of test executed.
    Priority *int `json:"priority,omitempty"`
    // Reference to team project.
    Project *ShallowReference `json:"project,omitempty"`
    // Shallow reference to release associated with test result.
    Release *ShallowReference `json:"release,omitempty"`
    // Reference to release associated with test result.
    ReleaseReference *ReleaseReference `json:"releaseReference,omitempty"`
    // ResetCount.
    ResetCount *int `json:"resetCount,omitempty"`
    // Resolution state of test result.
    ResolutionState *string `json:"resolutionState,omitempty"`
    // ID of resolution state.
    ResolutionStateId *int `json:"resolutionStateId,omitempty"`
    // Hierarchy type of the result, default value of None means its leaf node.
    ResultGroupType *ResultGroupType `json:"resultGroupType,omitempty"`
    // Revision number of test result.
    Revision *int `json:"revision,omitempty"`
    // Reference to identity executed the test.
    RunBy *IdentityRef `json:"runBy,omitempty"`
    // Stacktrace with maxSize= 1000 chars.
    StackTrace *string `json:"stackTrace,omitempty"`
    // Time when test execution started.
    StartedDate *time.Time `json:"startedDate,omitempty"`
    // State of test result. Type TestRunState.
    State *string `json:"state,omitempty"`
    // List of sub results inside a test result, if ResultGroupType is not None, it holds corresponding type sub results.
    SubResults *[]TestSubResult `json:"subResults,omitempty"`
    // Reference to the test executed.
    TestCase *ShallowReference `json:"testCase,omitempty"`
    // Reference ID of test used by test result. Type TestResultMetaData
    TestCaseReferenceId *int `json:"testCaseReferenceId,omitempty"`
    // TestCaseRevision Number.
    TestCaseRevision *int `json:"testCaseRevision,omitempty"`
    // Name of test.
    TestCaseTitle *string `json:"testCaseTitle,omitempty"`
    // Reference to test plan test case workitem is part of.
    TestPlan *ShallowReference `json:"testPlan,omitempty"`
    // Reference to the test point executed.
    TestPoint *ShallowReference `json:"testPoint,omitempty"`
    // Reference to test run.
    TestRun *ShallowReference `json:"testRun,omitempty"`
    // Reference to test suite test case workitem is part of.
    TestSuite *ShallowReference `json:"testSuite,omitempty"`
    // Url of test result.
    Url *string `json:"url,omitempty"`
}

// Test attachment information in a test iteration.
type TestCaseResultAttachmentModel struct {
    // Path identifier test step in test case workitem.
    ActionPath *string `json:"actionPath,omitempty"`
    // Attachment ID.
    Id *int `json:"id,omitempty"`
    // Iteration ID.
    IterationId *int `json:"iterationId,omitempty"`
    // Name of attachment.
    Name *string `json:"name,omitempty"`
    // Attachment size.
    Size *uint64 `json:"size,omitempty"`
    // Url to attachment.
    Url *string `json:"url,omitempty"`
}

// Reference to a test result.
type TestCaseResultIdentifier struct {
    // Test result ID.
    TestResultId *int `json:"testResultId,omitempty"`
    // Test run ID.
    TestRunId *int `json:"testRunId,omitempty"`
}

// Test environment Detail.
type TestEnvironment struct {
    // Test Environment Id.
    EnvironmentId *uuid.UUID `json:"environmentId,omitempty"`
    // Test Environment Name.
    EnvironmentName *string `json:"environmentName,omitempty"`
}

type TestFailureDetails struct {
    Count *int `json:"count,omitempty"`
    TestResults *[]TestCaseResultIdentifier `json:"testResults,omitempty"`
}

type TestFailuresAnalysis struct {
    ExistingFailures *TestFailureDetails `json:"existingFailures,omitempty"`
    FixedTests *TestFailureDetails `json:"fixedTests,omitempty"`
    NewFailures *TestFailureDetails `json:"newFailures,omitempty"`
    PreviousContext *TestResultsContext `json:"previousContext,omitempty"`
}

// Test Flaky Identifier
type TestFlakyIdentifier struct {
    // Branch Name where Flakiness has to be Marked/Unmarked
    BranchName *string `json:"branchName,omitempty"`
    // State for Flakiness
    IsFlaky *bool `json:"isFlaky,omitempty"`
}

// Filter to get TestCase result history.
type TestHistoryQuery struct {
    // Automated test name of the TestCase.
    AutomatedTestName *string `json:"automatedTestName,omitempty"`
    // Results to be get for a particular branches.
    Branch *string `json:"branch,omitempty"`
    // Get the results history only for this BuildDefinitionId. This to get used in query GroupBy should be Branch. If this is provided, Branch will have no use.
    BuildDefinitionId *int `json:"buildDefinitionId,omitempty"`
    // It will be filled by server. If not null means there are some results still to be get, and we need to call this REST API with this ContinuousToken. It is not supposed to be created (or altered, if received from server in last batch) by user.
    ContinuationToken *string `json:"continuationToken,omitempty"`
    // Group the result on the basis of TestResultGroupBy. This can be Branch, Environment or null(if results are fetched by BuildDefinitionId)
    GroupBy *TestResultGroupBy `json:"groupBy,omitempty"`
    // History to get between time interval MaxCompleteDate and  (MaxCompleteDate - TrendDays). Default is current date time.
    MaxCompleteDate *time.Time `json:"maxCompleteDate,omitempty"`
    // Get the results history only for this ReleaseEnvDefinitionId. This to get used in query GroupBy should be Environment.
    ReleaseEnvDefinitionId *int `json:"releaseEnvDefinitionId,omitempty"`
    // List of TestResultHistoryForGroup which are grouped by GroupBy
    ResultsForGroup *[]TestResultHistoryForGroup `json:"resultsForGroup,omitempty"`
    // Get the results history only for this testCaseId. This to get used in query to filter the result along with automatedtestname
    TestCaseId *int `json:"testCaseId,omitempty"`
    // Number of days for which history to collect. Maximum supported value is 7 days. Default is 7 days.
    TrendDays *int `json:"trendDays,omitempty"`
}

// Represents a test iteration result.
type TestIterationDetailsModel struct {
    // Test step results in an iteration.
    ActionResults *[]TestActionResultModel `json:"actionResults,omitempty"`
    // Reference to attachments in test iteration result.
    Attachments *[]TestCaseResultAttachmentModel `json:"attachments,omitempty"`
    // Comment in test iteration result.
    Comment *string `json:"comment,omitempty"`
    // Time when execution completed.
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Duration of execution.
    DurationInMs *float64 `json:"durationInMs,omitempty"`
    // Error message in test iteration result execution.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // ID of test iteration result.
    Id *int `json:"id,omitempty"`
    // Test outcome if test iteration result.
    Outcome *string `json:"outcome,omitempty"`
    // Test parameters in an iteration.
    Parameters *[]TestResultParameterModel `json:"parameters,omitempty"`
    // Time when execution started.
    StartedDate *time.Time `json:"startedDate,omitempty"`
    // Url to test iteration result.
    Url *string `json:"url,omitempty"`
}

// Represents Test Log Result object.
type TestLog struct {
    // Test Log Context run, build
    LogReference *TestLogReference `json:"logReference,omitempty"`
    MetaData *map[string]string `json:"metaData,omitempty"`
    // LastUpdatedDate for Log file
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Size in Bytes for Log file
    Size *uint64 `json:"size,omitempty"`
}

type TestLogReference struct {
    // BuildId for test log, if context is build
    BuildId *int `json:"buildId,omitempty"`
    // FileName for log file
    FilePath *string `json:"filePath,omitempty"`
    // ReleaseEnvId for test log, if context is Release
    ReleaseEnvId *int `json:"releaseEnvId,omitempty"`
    // ReleaseId for test log, if context is Release
    ReleaseId *int `json:"releaseId,omitempty"`
    // Resultid for test log, if context is run and log is related to result
    ResultId *int `json:"resultId,omitempty"`
    // runid for test log, if context is run
    RunId *int `json:"runId,omitempty"`
    // Test Log Reference object
    Scope *TestLogScope `json:"scope,omitempty"`
    // SubResultid for test log, if context is run and log is related to subresult
    SubResultId *int `json:"subResultId,omitempty"`
    // Log Type
    Type *TestLogType `json:"type,omitempty"`
}

// Test Log Context
type TestLogScope string

type testLogScopeValuesType struct {
    Run TestLogScope
    Build TestLogScope
    Release TestLogScope
}

var TestLogScopeValues = testLogScopeValuesType{
    // Log file is associated with Run, result, subresult
    Run: "run",
    // Log File associated with Build
    Build: "build",
    // Log File associated with Release
    Release: "release",
}

// Test Log Status codes.
type TestLogStatusCode string

type testLogStatusCodeValuesType struct {
    Success TestLogStatusCode
    Failed TestLogStatusCode
    FileAlreadyExists TestLogStatusCode
    InvalidInput TestLogStatusCode
    InvalidFileName TestLogStatusCode
    InvalidContainer TestLogStatusCode
    TransferFailed TestLogStatusCode
    FeatureDisabled TestLogStatusCode
    BuildDoesNotExist TestLogStatusCode
    RunDoesNotExist TestLogStatusCode
    ContainerNotCreated TestLogStatusCode
    ApiNotSupported TestLogStatusCode
    FileSizeExceeds TestLogStatusCode
    ContainerNotFound TestLogStatusCode
    FileNotFound TestLogStatusCode
    DirectoryNotFound TestLogStatusCode
}

var TestLogStatusCodeValues = testLogStatusCodeValuesType{
    Success: "success",
    Failed: "failed",
    FileAlreadyExists: "fileAlreadyExists",
    InvalidInput: "invalidInput",
    InvalidFileName: "invalidFileName",
    InvalidContainer: "invalidContainer",
    TransferFailed: "transferFailed",
    FeatureDisabled: "featureDisabled",
    BuildDoesNotExist: "buildDoesNotExist",
    RunDoesNotExist: "runDoesNotExist",
    ContainerNotCreated: "containerNotCreated",
    ApiNotSupported: "apiNotSupported",
    FileSizeExceeds: "fileSizeExceeds",
    ContainerNotFound: "containerNotFound",
    FileNotFound: "fileNotFound",
    DirectoryNotFound: "directoryNotFound",
}

// Represents Test Log store endpoint details.
type TestLogStoreEndpointDetails struct {
    // Test log store connection Uri.
    EndpointSASUri *string `json:"endpointSASUri,omitempty"`
    // Test log store endpoint type.
    EndpointType *TestLogStoreEndpointType `json:"endpointType,omitempty"`
    // Test log store status code
    Status *TestLogStatusCode `json:"status,omitempty"`
}

type TestLogStoreEndpointType string

type testLogStoreEndpointTypeValuesType struct {
    Root TestLogStoreEndpointType
    File TestLogStoreEndpointType
}

var TestLogStoreEndpointTypeValues = testLogStoreEndpointTypeValuesType{
    Root: "root",
    File: "file",
}

type TestLogStoreOperationType string

type testLogStoreOperationTypeValuesType struct {
    Read TestLogStoreOperationType
    Create TestLogStoreOperationType
    ReadAndCreate TestLogStoreOperationType
}

var TestLogStoreOperationTypeValues = testLogStoreOperationTypeValuesType{
    Read: "read",
    Create: "create",
    ReadAndCreate: "readAndCreate",
}

// Test Log Types
type TestLogType string

type testLogTypeValuesType struct {
    GeneralAttachment TestLogType
    CodeCoverage TestLogType
    TestImpact TestLogType
    Intermediate TestLogType
}

var TestLogTypeValues = testLogTypeValuesType{
    // Any gereric attachment.
    GeneralAttachment: "generalAttachment",
    // Code Coverage files
    CodeCoverage: "codeCoverage",
    // Test Impact details.
    TestImpact: "testImpact",
    // Temporary files
    Intermediate: "intermediate",
}

// An abstracted reference to some other resource. This class is used to provide the build data contracts with a uniform way to reference other resources in a way that provides easy traversal through links.
type TestMessageLogDetails struct {
    // Date when the resource is created
    DateCreated *time.Time `json:"dateCreated,omitempty"`
    // Id of the resource
    EntryId *int `json:"entryId,omitempty"`
    // Message of the resource
    Message *string `json:"message,omitempty"`
}

type TestMethod struct {
    Container *string `json:"container,omitempty"`
    Name *string `json:"name,omitempty"`
}

// Class representing a reference to an operation.
type TestOperationReference struct {
    Id *string `json:"id,omitempty"`
    Status *string `json:"status,omitempty"`
    Url *string `json:"url,omitempty"`
}

// Valid TestOutcome values.
type TestOutcome string

type testOutcomeValuesType struct {
    Unspecified TestOutcome
    None TestOutcome
    Passed TestOutcome
    Failed TestOutcome
    Inconclusive TestOutcome
    Timeout TestOutcome
    Aborted TestOutcome
    Blocked TestOutcome
    NotExecuted TestOutcome
    Warning TestOutcome
    Error TestOutcome
    NotApplicable TestOutcome
    Paused TestOutcome
    InProgress TestOutcome
    NotImpacted TestOutcome
}

var TestOutcomeValues = testOutcomeValuesType{
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
}

// Test Resolution State Details.
type TestResolutionState struct {
    // Test Resolution state Id.
    Id *int `json:"id,omitempty"`
    // Test Resolution State Name.
    Name *string `json:"name,omitempty"`
    Project *ShallowReference `json:"project,omitempty"`
}

type TestResultDocument struct {
    OperationReference *TestOperationReference `json:"operationReference,omitempty"`
    Payload *TestResultPayload `json:"payload,omitempty"`
}

// Group by for results
type TestResultGroupBy string

type testResultGroupByValuesType struct {
    Branch TestResultGroupBy
    Environment TestResultGroupBy
}

var TestResultGroupByValues = testResultGroupByValuesType{
    // Group the results by branches
    Branch: "branch",
    // Group the results by environment
    Environment: "environment",
}

type TestResultHistory struct {
    GroupByField *string `json:"groupByField,omitempty"`
    ResultsForGroup *[]TestResultHistoryDetailsForGroup `json:"resultsForGroup,omitempty"`
}

type TestResultHistoryDetailsForGroup struct {
    GroupByValue *interface{} `json:"groupByValue,omitempty"`
    LatestResult *TestCaseResult `json:"latestResult,omitempty"`
}

// List of test results filtered on the basis of GroupByValue
type TestResultHistoryForGroup struct {
    // Display name of the group.
    DisplayName *string `json:"displayName,omitempty"`
    // Name or Id of the group identifier by which results are grouped together.
    GroupByValue *string `json:"groupByValue,omitempty"`
    // List of results for GroupByValue
    Results *[]TestCaseResult `json:"results,omitempty"`
}

// Represents a Meta Data of a test result.
type TestResultMetaData struct {
    // AutomatedTestName of test result.
    AutomatedTestName *string `json:"automatedTestName,omitempty"`
    // AutomatedTestStorage of test result.
    AutomatedTestStorage *string `json:"automatedTestStorage,omitempty"`
    // List of Flaky Identifier for TestCaseReferenceId
    FlakyIdentifiers *[]TestFlakyIdentifier `json:"flakyIdentifiers,omitempty"`
    // Owner of test result.
    Owner *string `json:"owner,omitempty"`
    // Priority of test result.
    Priority *int `json:"priority,omitempty"`
    // ID of TestCaseReference.
    TestCaseReferenceId *int `json:"testCaseReferenceId,omitempty"`
    // TestCaseTitle of test result.
    TestCaseTitle *string `json:"testCaseTitle,omitempty"`
}

// Represents a TestResultMetaData Input
type TestResultMetaDataUpdateInput struct {
    // List of Flaky Identifiers
    FlakyIdentifiers *[]TestFlakyIdentifier `json:"flakyIdentifiers,omitempty"`
}

type TestResultModelBase struct {
    // Comment in result.
    Comment *string `json:"comment,omitempty"`
    // Time when execution completed.
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Duration of execution.
    DurationInMs *float64 `json:"durationInMs,omitempty"`
    // Error message in result.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // Test outcome of result.
    Outcome *string `json:"outcome,omitempty"`
    // Time when execution started.
    StartedDate *time.Time `json:"startedDate,omitempty"`
}

// Test parameter information in a test iteration.
type TestResultParameterModel struct {
    // Test step path where parameter is referenced.
    ActionPath *string `json:"actionPath,omitempty"`
    // Iteration ID.
    IterationId *int `json:"iterationId,omitempty"`
    // Name of parameter.
    ParameterName *string `json:"parameterName,omitempty"`
    // This is step Id of test case. For shared step, it is step Id of shared step in test case workitem; step Id in shared step. Example: TestCase workitem has two steps: 1) Normal step with Id = 1 2) Shared Step with Id = 2. Inside shared step: a) Normal Step with Id = 1 Value for StepIdentifier for First step: "1" Second step: "2;1"
    StepIdentifier *string `json:"stepIdentifier,omitempty"`
    // Url of test parameter.
    Url *string `json:"url,omitempty"`
    // Value of parameter.
    Value *string `json:"value,omitempty"`
}

type TestResultPayload struct {
    Comment *string `json:"comment,omitempty"`
    Name *string `json:"name,omitempty"`
    Stream *string `json:"stream,omitempty"`
}

type TestResultsContext struct {
    Build *BuildReference `json:"build,omitempty"`
    ContextType *TestResultsContextType `json:"contextType,omitempty"`
    Release *ReleaseReference `json:"release,omitempty"`
}

type TestResultsContextType string

type testResultsContextTypeValuesType struct {
    Build TestResultsContextType
    Release TestResultsContextType
}

var TestResultsContextTypeValues = testResultsContextTypeValuesType{
    Build: "build",
    Release: "release",
}

type TestResultsDetails struct {
    GroupByField *string `json:"groupByField,omitempty"`
    ResultsForGroup *[]TestResultsDetailsForGroup `json:"resultsForGroup,omitempty"`
}

type TestResultsDetailsForGroup struct {
    GroupByValue *interface{} `json:"groupByValue,omitempty"`
    Results *[]TestCaseResult `json:"results,omitempty"`
    ResultsCountByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsCountByOutcome,omitempty"`
    Tags *[]string `json:"tags,omitempty"`
}

type TestResultsQuery struct {
    Fields *[]string `json:"fields,omitempty"`
    Results *[]TestCaseResult `json:"results,omitempty"`
    ResultsFilter *ResultsFilter `json:"resultsFilter,omitempty"`
}

type TestResultsSettings struct {
    // IsRequired and EmitDefaultValue are passed as false as if users doesn't pass anything, should not come for serialisation and deserialisation.
    FlakySettings *FlakySettings `json:"flakySettings,omitempty"`
}

type TestResultsSettingsType string

type testResultsSettingsTypeValuesType struct {
    All TestResultsSettingsType
    Flaky TestResultsSettingsType
}

var TestResultsSettingsTypeValues = testResultsSettingsTypeValuesType{
    // Returns All Test Settings.
    All: "all",
    // Returns Flaky Test Settings.
    Flaky: "flaky",
}

type TestResultSummary struct {
    AggregatedResultsAnalysis *AggregatedResultsAnalysis `json:"aggregatedResultsAnalysis,omitempty"`
    NoConfigRunsCount *int `json:"noConfigRunsCount,omitempty"`
    TeamProject *TeamProjectReference `json:"teamProject,omitempty"`
    TestFailures *TestFailuresAnalysis `json:"testFailures,omitempty"`
    TestResultsContext *TestResultsContext `json:"testResultsContext,omitempty"`
    TotalRunsCount *int `json:"totalRunsCount,omitempty"`
}

type TestResultsUpdateSettings struct {
    // FlakySettings defines Flaky Settings Data.
    FlakySettings *FlakySettings `json:"flakySettings,omitempty"`
}

type TestResultTrendFilter struct {
    BranchNames *[]string `json:"branchNames,omitempty"`
    BuildCount *int `json:"buildCount,omitempty"`
    DefinitionIds *[]int `json:"definitionIds,omitempty"`
    EnvDefinitionIds *[]int `json:"envDefinitionIds,omitempty"`
    MaxCompleteDate *time.Time `json:"maxCompleteDate,omitempty"`
    PublishContext *string `json:"publishContext,omitempty"`
    TestRunTitles *[]string `json:"testRunTitles,omitempty"`
    TrendDays *int `json:"trendDays,omitempty"`
}

// Test run details.
type TestRun struct {
    // Build associated with this test run.
    Build *ShallowReference `json:"build,omitempty"`
    // Build configuration details associated with this test run.
    BuildConfiguration *BuildConfiguration `json:"buildConfiguration,omitempty"`
    // Comments entered by those analyzing the run.
    Comment *string `json:"comment,omitempty"`
    // Completed date time of the run.
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Test Run Controller.
    Controller *string `json:"controller,omitempty"`
    // Test Run CreatedDate.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // List of Custom Fields for TestRun.
    CustomFields *[]CustomTestField `json:"customFields,omitempty"`
    // Drop Location for the test Run.
    DropLocation *string `json:"dropLocation,omitempty"`
    DtlAutEnvironment *ShallowReference `json:"dtlAutEnvironment,omitempty"`
    DtlEnvironment *ShallowReference `json:"dtlEnvironment,omitempty"`
    DtlEnvironmentCreationDetails *DtlEnvironmentDetails `json:"dtlEnvironmentCreationDetails,omitempty"`
    // Due date and time for test run.
    DueDate *time.Time `json:"dueDate,omitempty"`
    // Error message associated with the run.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    Filter *RunFilter `json:"filter,omitempty"`
    // ID of the test run.
    Id *int `json:"id,omitempty"`
    // Number of Incomplete Tests.
    IncompleteTests *int `json:"incompleteTests,omitempty"`
    // true if test run is automated, false otherwise.
    IsAutomated *bool `json:"isAutomated,omitempty"`
    // The iteration to which the run belongs.
    Iteration *string `json:"iteration,omitempty"`
    // Team foundation ID of the last updated the test run.
    LastUpdatedBy *IdentityRef `json:"lastUpdatedBy,omitempty"`
    // Last updated date and time
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Name of the test run.
    Name *string `json:"name,omitempty"`
    // Number of Not Applicable Tests.
    NotApplicableTests *int `json:"notApplicableTests,omitempty"`
    // Team Foundation ID of the owner of the runs.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Number of passed tests in the run
    PassedTests *int `json:"passedTests,omitempty"`
    // Phase/State for the testRun.
    Phase *string `json:"phase,omitempty"`
    // Reference of the pipeline to which this test run belongs.
    PipelineReference *PipelineReference `json:"pipelineReference,omitempty"`
    // Test plan associated with this test run.
    Plan *ShallowReference `json:"plan,omitempty"`
    // Post Process State.
    PostProcessState *string `json:"postProcessState,omitempty"`
    // Project associated with this run.
    Project *ShallowReference `json:"project,omitempty"`
    // Release Reference for the Test Run.
    Release *ReleaseReference `json:"release,omitempty"`
    // Release Environment Uri for TestRun.
    ReleaseEnvironmentUri *string `json:"releaseEnvironmentUri,omitempty"`
    // Release Uri for TestRun.
    ReleaseUri *string `json:"releaseUri,omitempty"`
    Revision *int `json:"revision,omitempty"`
    // RunSummary by outcome.
    RunStatistics *[]RunStatistic `json:"runStatistics,omitempty"`
    // Start date time of the run.
    StartedDate *time.Time `json:"startedDate,omitempty"`
    // The state of the run. Type TestRunState Valid states - Unspecified ,NotStarted, InProgress, Completed, Waiting, Aborted, NeedsInvestigation
    State *string `json:"state,omitempty"`
    // TestRun Substate.
    Substate *TestRunSubstate `json:"substate,omitempty"`
    // Tags attached with this test run.
    Tags *[]TestTag `json:"tags,omitempty"`
    // Test environment associated with the run.
    TestEnvironment *TestEnvironment `json:"testEnvironment,omitempty"`
    TestMessageLogId *int `json:"testMessageLogId,omitempty"`
    TestSettings *ShallowReference `json:"testSettings,omitempty"`
    // Total tests in the run
    TotalTests *int `json:"totalTests,omitempty"`
    // Number of failed tests in the run.
    UnanalyzedTests *int `json:"unanalyzedTests,omitempty"`
    // Url of the test run
    Url *string `json:"url,omitempty"`
    // Web Access Url for TestRun.
    WebAccessUrl *string `json:"webAccessUrl,omitempty"`
}

// Test Run Code Coverage Details
type TestRunCoverage struct {
    // Last Error
    LastError *string `json:"lastError,omitempty"`
    // List of Modules Coverage
    Modules *[]ModuleCoverage `json:"modules,omitempty"`
    // State
    State *string `json:"state,omitempty"`
    // Reference of test Run.
    TestRun *ShallowReference `json:"testRun,omitempty"`
}

// The types of outcomes for test run.
type TestRunOutcome string

type testRunOutcomeValuesType struct {
    Passed TestRunOutcome
    Failed TestRunOutcome
    NotImpacted TestRunOutcome
    Others TestRunOutcome
}

var TestRunOutcomeValues = testRunOutcomeValuesType{
    // Run with zero failed tests and has at least one impacted test
    Passed: "passed",
    // Run with at-least one failed test.
    Failed: "failed",
    // Run with no impacted tests.
    NotImpacted: "notImpacted",
    // Runs with All tests in other category.
    Others: "others",
}

// The types of publish context for run.
type TestRunPublishContext string

type testRunPublishContextValuesType struct {
    Build TestRunPublishContext
    Release TestRunPublishContext
    All TestRunPublishContext
}

var TestRunPublishContextValues = testRunPublishContextValuesType{
    // Run is published for Build Context.
    Build: "build",
    // Run is published for Release Context.
    Release: "release",
    // Run is published for any Context.
    All: "all",
}

// The types of states for test run.
type TestRunState string

type testRunStateValuesType struct {
    Unspecified TestRunState
    NotStarted TestRunState
    InProgress TestRunState
    Completed TestRunState
    Aborted TestRunState
    Waiting TestRunState
    NeedsInvestigation TestRunState
}

var TestRunStateValues = testRunStateValuesType{
    // Only used during an update to preserve the existing value.
    Unspecified: "unspecified",
    // The run is still being created.  No tests have started yet.
    NotStarted: "notStarted",
    // Tests are running.
    InProgress: "inProgress",
    // All tests have completed or been skipped.
    Completed: "completed",
    // Run is stopped and remaining tests have been aborted
    Aborted: "aborted",
    // Run is currently initializing This is a legacy state and should not be used any more
    Waiting: "waiting",
    // Run requires investigation because of a test point failure This is a legacy state and should not be used any more
    NeedsInvestigation: "needsInvestigation",
}

// Test run statistics.
type TestRunStatistic struct {
    Run *ShallowReference `json:"run,omitempty"`
    RunStatistics *[]RunStatistic `json:"runStatistics,omitempty"`
}

// The types of sub states for test run. It gives the user more info about the test run beyond the high level test run state
type TestRunSubstate string

type testRunSubstateValuesType struct {
    None TestRunSubstate
    CreatingEnvironment TestRunSubstate
    RunningTests TestRunSubstate
    CanceledByUser TestRunSubstate
    AbortedBySystem TestRunSubstate
    TimedOut TestRunSubstate
    PendingAnalysis TestRunSubstate
    Analyzed TestRunSubstate
    CancellationInProgress TestRunSubstate
}

var TestRunSubstateValues = testRunSubstateValuesType{
    // Run with noState.
    None: "none",
    // Run state while Creating Environment.
    CreatingEnvironment: "creatingEnvironment",
    // Run state while Running Tests.
    RunningTests: "runningTests",
    // Run state while Creating Environment.
    CanceledByUser: "canceledByUser",
    // Run state when it is Aborted By the System.
    AbortedBySystem: "abortedBySystem",
    // Run state when run has timedOut.
    TimedOut: "timedOut",
    // Run state while Pending Analysis.
    PendingAnalysis: "pendingAnalysis",
    // Run state after being Analysed.
    Analyzed: "analyzed",
    // Run state when cancellation is in Progress.
    CancellationInProgress: "cancellationInProgress",
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

// Represents a sub result of a test result.
type TestSubResult struct {
    // Comment in sub result.
    Comment *string `json:"comment,omitempty"`
    // Time when test execution completed.
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Machine where test executed.
    ComputerName *string `json:"computerName,omitempty"`
    // Reference to test configuration.
    Configuration *ShallowReference `json:"configuration,omitempty"`
    // Additional properties of sub result.
    CustomFields *[]CustomTestField `json:"customFields,omitempty"`
    // Name of sub result.
    DisplayName *string `json:"displayName,omitempty"`
    // Duration of test execution.
    DurationInMs *uint64 `json:"durationInMs,omitempty"`
    // Error message in sub result.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // ID of sub result.
    Id *int `json:"id,omitempty"`
    // Time when result last updated.
    LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
    // Outcome of sub result.
    Outcome *string `json:"outcome,omitempty"`
    // Immediate parent ID of sub result.
    ParentId *int `json:"parentId,omitempty"`
    // Hierarchy type of the result, default value of None means its leaf node.
    ResultGroupType *ResultGroupType `json:"resultGroupType,omitempty"`
    // Index number of sub result.
    SequenceId *int `json:"sequenceId,omitempty"`
    // Stacktrace.
    StackTrace *string `json:"stackTrace,omitempty"`
    // Time when test execution started.
    StartedDate *time.Time `json:"startedDate,omitempty"`
    // List of sub results inside a sub result, if ResultGroupType is not None, it holds corresponding type sub results.
    SubResults *[]TestSubResult `json:"subResults,omitempty"`
    // Reference to test result.
    TestResult *TestCaseResultIdentifier `json:"testResult,omitempty"`
    // Url of sub result.
    Url *string `json:"url,omitempty"`
}

type TestSummaryForWorkItem struct {
    Summary *AggregatedDataForResultTrend `json:"summary,omitempty"`
    WorkItem *WorkItemReference `json:"workItem,omitempty"`
}

// Tag attached to a run or result.
type TestTag struct {
    // Name of the tag, alphanumeric value less than 30 chars
    Name *string `json:"name,omitempty"`
}

// Test tag summary for build or release grouped by test run.
type TestTagSummary struct {
    // Dictionary which contains tags associated with a test run.
    TagsGroupByTestArtifact *map[int][]TestTag `json:"tagsGroupByTestArtifact,omitempty"`
}

// Tags to update to a run or result.
type TestTagsUpdateModel struct {
    Tags *[]azureDevops.KeyValuePair `json:"tags,omitempty"`
}

type TestToWorkItemLinks struct {
    Test *TestMethod `json:"test,omitempty"`
    WorkItems *[]WorkItemReference `json:"workItems,omitempty"`
}

// WorkItem reference Details.
type WorkItemReference struct {
    // WorkItem Id.
    Id *string `json:"id,omitempty"`
    // WorkItem Name.
    Name *string `json:"name,omitempty"`
    // WorkItem Type.
    Type *string `json:"type,omitempty"`
    // WorkItem Url. Valid Values : (Bug, Task, User Story, Test Case)
    Url *string `json:"url,omitempty"`
    // WorkItem WebUrl.
    WebUrl *string `json:"webUrl,omitempty"`
}

type WorkItemToTestLinks struct {
    ExecutedIn *Service `json:"executedIn,omitempty"`
    Tests *[]TestMethod `json:"tests,omitempty"`
    WorkItem *WorkItemReference `json:"workItem,omitempty"`
}
