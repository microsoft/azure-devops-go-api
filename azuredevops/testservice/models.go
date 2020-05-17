// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package testservice

import (
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

type AgentGroup struct {
	// User that created the agent group
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// Time agent group was created
	CreationTime *azuredevops.Time `json:"creationTime,omitempty"`
	// Id of the agent group
	GroupId *string `json:"groupId,omitempty"`
	// The name of the agent group
	GroupName         *string                 `json:"groupName,omitempty"`
	MachineAccessData *[]AgentGroupAccessData `json:"machineAccessData,omitempty"`
	// Machine configuration
	MachineConfiguration *WebApiUserLoadTestMachineInput `json:"machineConfiguration,omitempty"`
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`
}

type AgentGroupAccessData struct {
	// Type Specific details
	Details *string `json:"details,omitempty"`
	// Access string
	StorageConnectionString *string `json:"storageConnectionString,omitempty"`
	// Endpoint for the service
	StorageEndPoint *string `json:"storageEndPoint,omitempty"`
	// Identifier for the storage (eg. table name)
	StorageName *string `json:"storageName,omitempty"`
	// Type of the store (table, queue, blob)
	StorageType *string `json:"storageType,omitempty"`
}

type Application struct {
	// Unique Id of the Application Component
	ApplicationId *string `json:"applicationId,omitempty"`
	// Description of the Application component
	Description *string `json:"description,omitempty"`
	// The Name of the Application component
	Name *string `json:"name,omitempty"`
	// Path identifier of the Application component
	Path *string `json:"path,omitempty"`
	// Character used to separate paths for counters
	PathSeperator *string `json:"pathSeperator,omitempty"`
	// Type identifier of the Application component under test
	Type *string `json:"type,omitempty"`
	// Version of the Application Component
	Version *string `json:"version,omitempty"`
}

type ApplicationCounters struct {
	// The unique Id of the Application that the counter belongs
	ApplicationId *string `json:"applicationId,omitempty"`
	// Description of autCounter
	Description *string `json:"description,omitempty"`
	// The unique Id for the AutCounter
	Id *string `json:"id,omitempty"`
	// Whether the autCounter is a default counter or not
	IsDefault *bool `json:"isDefault,omitempty"`
	// Name of the AutCounter
	Name *string `json:"name,omitempty"`
	// The Path of the the autcounter wrt to hierarchy
	Path *string `json:"path,omitempty"`
}

type ApplicationType struct {
	// Helper link url
	ActionUriLink *string `json:"actionUriLink,omitempty"`
	// The link that points to aut results site
	AutPortalLink *string `json:"autPortalLink,omitempty"`
	// true if application results collection is enabled for this tenant
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// the max no. of application components allowed for collection per run
	MaxComponentsAllowedForCollection *int `json:"maxComponentsAllowedForCollection,omitempty"`
	// The max no. of counters that can be collected per aut
	MaxCountersAllowed *int `json:"maxCountersAllowed,omitempty"`
	// Application Type
	Type *string `json:"type,omitempty"`
}

type BrowserMix struct {
	BrowserName       *string  `json:"browserName,omitempty"`
	BrowserPercentage *float32 `json:"browserPercentage,omitempty"`
}

type CltCustomerIntelligenceData struct {
	Area       *string                 `json:"area,omitempty"`
	Feature    *string                 `json:"feature,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

type CounterGroup struct {
	GroupName *string `json:"groupName,omitempty"`
	Url       *string `json:"url,omitempty"`
}

type CounterInstanceSamples struct {
	Count             *int    `json:"count,omitempty"`
	CounterInstanceId *string `json:"counterInstanceId,omitempty"`
	// The time of next refresh
	NextRefreshTime *azuredevops.Time `json:"nextRefreshTime,omitempty"`
	Values          *[]CounterSample  `json:"values,omitempty"`
}

type CounterSample struct {
	BaseValue         *uint64           `json:"baseValue,omitempty"`
	ComputedValue     *float32          `json:"computedValue,omitempty"`
	CounterFrequency  *uint64           `json:"counterFrequency,omitempty"`
	CounterInstanceId *string           `json:"counterInstanceId,omitempty"`
	CounterType       *string           `json:"counterType,omitempty"`
	IntervalEndDate   *azuredevops.Time `json:"intervalEndDate,omitempty"`
	IntervalNumber    *int              `json:"intervalNumber,omitempty"`
	RawValue          *uint64           `json:"rawValue,omitempty"`
	SystemFrequency   *uint64           `json:"systemFrequency,omitempty"`
	TimeStamp         *uint64           `json:"timeStamp,omitempty"`
}

type CounterSampleQueryDetails struct {
	// The instanceId for which samples are required
	CounterInstanceId *string `json:"counterInstanceId,omitempty"`
	FromInterval      *int    `json:"fromInterval,omitempty"`
	ToInterval        *int    `json:"toInterval,omitempty"`
}

type CounterSamplesResult struct {
	// Count of the samples
	Count *int `json:"count,omitempty"`
	// Maximum number of samples returned in this object
	MaxBatchSize *int `json:"maxBatchSize,omitempty"`
	// Count of the samples
	TotalSamplesCount *int `json:"totalSamplesCount,omitempty"`
	// The result samples
	Values *[]CounterInstanceSamples `json:"values,omitempty"`
}

type Diagnostics struct {
	DiagnosticStoreConnectionString *string           `json:"diagnosticStoreConnectionString,omitempty"`
	LastModifiedTime                *azuredevops.Time `json:"lastModifiedTime,omitempty"`
	RelativePathToDiagnosticFiles   *string           `json:"relativePathToDiagnosticFiles,omitempty"`
}

type DropAccessData struct {
	DropContainerUrl *string `json:"dropContainerUrl,omitempty"`
	// The SaSkey to use for the drop.
	SasKey *string `json:"sasKey,omitempty"`
}

type ErrorDetails struct {
	LastErrorDate *azuredevops.Time `json:"lastErrorDate,omitempty"`
	MessageText   *string           `json:"messageText,omitempty"`
	Occurrences   *int              `json:"occurrences,omitempty"`
	Request       *string           `json:"request,omitempty"`
	ScenarioName  *string           `json:"scenarioName,omitempty"`
	StackTrace    *string           `json:"stackTrace,omitempty"`
	TestCaseName  *string           `json:"testCaseName,omitempty"`
}

type LoadGenerationGeoLocation struct {
	Location   *string `json:"location,omitempty"`
	Percentage *int    `json:"percentage,omitempty"`
}

type LoadTest struct {
}

type LoadTestDefinition struct {
	AgentCount                 *int                         `json:"agentCount,omitempty"`
	BrowserMixs                *[]BrowserMix                `json:"browserMixs,omitempty"`
	CoreCount                  *int                         `json:"coreCount,omitempty"`
	CoresPerAgent              *int                         `json:"coresPerAgent,omitempty"`
	LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
	LoadPatternName            *string                      `json:"loadPatternName,omitempty"`
	LoadTestName               *string                      `json:"loadTestName,omitempty"`
	MaxVusers                  *int                         `json:"maxVusers,omitempty"`
	RunDuration                *int                         `json:"runDuration,omitempty"`
	SamplingRate               *int                         `json:"samplingRate,omitempty"`
	ThinkTime                  *int                         `json:"thinkTime,omitempty"`
	Urls                       *[]string                    `json:"urls,omitempty"`
}

type LoadTestErrorDetails struct {
	LastErrorDate *azuredevops.Time `json:"lastErrorDate,omitempty"`
	MessageText   *string           `json:"messageText,omitempty"`
	Occurrences   *int              `json:"occurrences,omitempty"`
	Request       *string           `json:"request,omitempty"`
	ScenarioName  *string           `json:"scenarioName,omitempty"`
	StackTrace    *string           `json:"stackTrace,omitempty"`
	SubType       *string           `json:"subType,omitempty"`
	TestCaseName  *string           `json:"testCaseName,omitempty"`
	Type          *string           `json:"type,omitempty"`
}

type LoadTestErrors struct {
	Count       *int           `json:"count,omitempty"`
	Occurrences *int           `json:"occurrences,omitempty"`
	Types       *[]interface{} `json:"types,omitempty"`
	Url         *string        `json:"url,omitempty"`
}

type LoadTestMachineType string

type loadTestMachineTypeValuesType struct {
	Default       LoadTestMachineType
	CltLoadAgent  LoadTestMachineType
	UserLoadAgent LoadTestMachineType
}

var LoadTestMachineTypeValues = loadTestMachineTypeValuesType{
	Default:       "default",
	CltLoadAgent:  "cltLoadAgent",
	UserLoadAgent: "userLoadAgent",
}

type LoadTestRunDetails struct {
	AgentCount                *int                 `json:"agentCount,omitempty"`
	CoreCount                 *int                 `json:"coreCount,omitempty"`
	CoresPerAgent             *int                 `json:"coresPerAgent,omitempty"`
	Duration                  *int                 `json:"duration,omitempty"`
	LoadGeneratorMachinesType *LoadTestMachineType `json:"loadGeneratorMachinesType,omitempty"`
	SamplingInterval          *int                 `json:"samplingInterval,omitempty"`
	WarmUpDuration            *int                 `json:"warmUpDuration,omitempty"`
	VirtualUserCount          *int                 `json:"virtualUserCount,omitempty"`
}

type LoadTestRunSettings struct {
	AgentCount                *int                 `json:"agentCount,omitempty"`
	CoreCount                 *int                 `json:"coreCount,omitempty"`
	CoresPerAgent             *int                 `json:"coresPerAgent,omitempty"`
	Duration                  *int                 `json:"duration,omitempty"`
	LoadGeneratorMachinesType *LoadTestMachineType `json:"loadGeneratorMachinesType,omitempty"`
	SamplingInterval          *int                 `json:"samplingInterval,omitempty"`
	WarmUpDuration            *int                 `json:"warmUpDuration,omitempty"`
}

type LoadTestTypes string

type loadTestTypesValuesType struct {
	VisualStudioLoadTest LoadTestTypes
	JMeter               LoadTestTypes
	OldLoadTestFile      LoadTestTypes
}

var LoadTestTypesValues = loadTestTypesValuesType{
	VisualStudioLoadTest: "visualStudioLoadTest",
	JMeter:               "jMeter",
	OldLoadTestFile:      "oldLoadTestFile",
}

type MessageSource string

type messageSourceValuesType struct {
	SetupScript          MessageSource
	CleanupScript        MessageSource
	Validation           MessageSource
	Other                MessageSource
	AutCounterCollection MessageSource
}

var MessageSourceValues = messageSourceValuesType{
	SetupScript:          "setupScript",
	CleanupScript:        "cleanupScript",
	Validation:           "validation",
	Other:                "other",
	AutCounterCollection: "autCounterCollection",
}

type MessageType string

type messageTypeValuesType struct {
	Info     MessageType
	Output   MessageType
	Error    MessageType
	Warning  MessageType
	Critical MessageType
}

var MessageTypeValues = messageTypeValuesType{
	Info:     "info",
	Output:   "output",
	Error:    "error",
	Warning:  "warning",
	Critical: "critical",
}

type OverridableRunSettings struct {
	LoadGeneratorMachinesType *LoadTestMachineType   `json:"loadGeneratorMachinesType,omitempty"`
	StaticAgentRunSettings    *StaticAgentRunSetting `json:"staticAgentRunSettings,omitempty"`
}

type PageSummary struct {
	AveragePageTime            *float64                 `json:"averagePageTime,omitempty"`
	PageUrl                    *string                  `json:"pageUrl,omitempty"`
	PercentagePagesMeetingGoal *int                     `json:"percentagePagesMeetingGoal,omitempty"`
	PercentileData             *[]SummaryPercentileData `json:"percentileData,omitempty"`
	ScenarioName               *string                  `json:"scenarioName,omitempty"`
	TestName                   *string                  `json:"testName,omitempty"`
	TotalPages                 *int                     `json:"totalPages,omitempty"`
}

type ProcessorArchitecture string

type processorArchitectureValuesType struct {
	None  ProcessorArchitecture
	Msil  ProcessorArchitecture
	X86   ProcessorArchitecture
	Ia64  ProcessorArchitecture
	Amd64 ProcessorArchitecture
	Arm   ProcessorArchitecture
}

var ProcessorArchitectureValues = processorArchitectureValuesType{
	None:  "none",
	Msil:  "msil",
	X86:   "x86",
	Ia64:  "ia64",
	Amd64: "amd64",
	Arm:   "arm",
}

type RequestSummary struct {
	AverageResponseTime *float64                 `json:"averageResponseTime,omitempty"`
	FailedRequests      *int                     `json:"failedRequests,omitempty"`
	PassedRequests      *int                     `json:"passedRequests,omitempty"`
	PercentileData      *[]SummaryPercentileData `json:"percentileData,omitempty"`
	RequestsPerSec      *float64                 `json:"requestsPerSec,omitempty"`
	RequestUrl          *string                  `json:"requestUrl,omitempty"`
	ScenarioName        *string                  `json:"scenarioName,omitempty"`
	TestName            *string                  `json:"testName,omitempty"`
	TotalRequests       *int                     `json:"totalRequests,omitempty"`
}

type ScenarioSummary struct {
	MaxUserLoad  *int    `json:"maxUserLoad,omitempty"`
	MinUserLoad  *int    `json:"minUserLoad,omitempty"`
	ScenarioName *string `json:"scenarioName,omitempty"`
}

type StaticAgent struct {
	AgentGroupId   *string           `json:"agentGroupId,omitempty"`
	AgentGroupName *string           `json:"agentGroupName,omitempty"`
	LastHeartBeat  *azuredevops.Time `json:"lastHeartBeat,omitempty"`
	Name           *string           `json:"name,omitempty"`
	State          *string           `json:"state,omitempty"`
}

type StaticAgentRunSetting struct {
	LoadGeneratorMachinesType *LoadTestMachineType `json:"loadGeneratorMachinesType,omitempty"`
	StaticAgentGroupName      *string              `json:"staticAgentGroupName,omitempty"`
}

type SubType struct {
	Count           *int            `json:"count,omitempty"`
	ErrorDetailList *[]ErrorDetails `json:"errorDetailList,omitempty"`
	Occurrences     *int            `json:"occurrences,omitempty"`
	SubTypeName     *string         `json:"subTypeName,omitempty"`
	Url             *string         `json:"url,omitempty"`
}

type SummaryPercentileData struct {
	Percentile      *int     `json:"percentile,omitempty"`
	PercentileValue *float64 `json:"percentileValue,omitempty"`
}

type TestDefinition struct {
	// Data for accessing the drop and not persisted in storage
	AccessData       *DropAccessData     `json:"accessData,omitempty"`
	CreatedBy        *webapi.IdentityRef `json:"createdBy,omitempty"`
	CreatedDate      *azuredevops.Time   `json:"createdDate,omitempty"`
	Id               *string             `json:"id,omitempty"`
	LastModifiedBy   *webapi.IdentityRef `json:"lastModifiedBy,omitempty"`
	LastModifiedDate *azuredevops.Time   `json:"lastModifiedDate,omitempty"`
	LoadTestType     *LoadTestTypes      `json:"loadTestType,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Description      *string             `json:"description,omitempty"`
	// Geo location from where load is generated
	LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
	LoadTestDefinitionSource   *string                      `json:"loadTestDefinitionSource,omitempty"`
	RunSettings                *LoadTestRunSettings         `json:"runSettings,omitempty"`
	StaticAgentRunSettings     *StaticAgentRunSetting       `json:"staticAgentRunSettings,omitempty"`
	TestDetails                *LoadTest                    `json:"testDetails,omitempty"`
}

type TestDefinitionBasic struct {
	// Data for accessing the drop and not persisted in storage
	AccessData       *DropAccessData     `json:"accessData,omitempty"`
	CreatedBy        *webapi.IdentityRef `json:"createdBy,omitempty"`
	CreatedDate      *azuredevops.Time   `json:"createdDate,omitempty"`
	Id               *string             `json:"id,omitempty"`
	LastModifiedBy   *webapi.IdentityRef `json:"lastModifiedBy,omitempty"`
	LastModifiedDate *azuredevops.Time   `json:"lastModifiedDate,omitempty"`
	LoadTestType     *LoadTestTypes      `json:"loadTestType,omitempty"`
	Name             *string             `json:"name,omitempty"`
}

type TestDrop struct {
	// Data for accessing the drop and not persisted in storage
	AccessData *DropAccessData `json:"accessData,omitempty"`
	// Time at which the drop is created
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// Identifies the type of drop
	DropType *string `json:"dropType,omitempty"`
	// Drop Id
	Id *string `json:"id,omitempty"`
	// LoadTest definition of the run for which testdrop is created
	LoadTestDefinition *LoadTestDefinition `json:"loadTestDefinition,omitempty"`
	// Test Run Id
	TestRunId *string `json:"testRunId,omitempty"`
}

// An abstracted reference to some other resource. This class is used to provide the load test data contracts with a uniform way to reference other resources in a way that provides easy traversal through links.
type TestDropRef struct {
	// Id of the resource
	Id *string `json:"id,omitempty"`
	// Full http link to the resource
	Url *string `json:"url,omitempty"`
}

type TestResults struct {
	// The uri to the test run results file.
	CloudLoadTestSolutionUrl *string         `json:"cloudLoadTestSolutionUrl,omitempty"`
	CounterGroups            *[]CounterGroup `json:"counterGroups,omitempty"`
	// The object contains diagnostic details
	Diagnostics *Diagnostics `json:"diagnostics,omitempty"`
	// The uri to the test run results file.
	ResultsUrl *string `json:"resultsUrl,omitempty"`
}

type TestResultsSummary struct {
	OverallPageSummary        *PageSummary          `json:"overallPageSummary,omitempty"`
	OverallRequestSummary     *RequestSummary       `json:"overallRequestSummary,omitempty"`
	OverallScenarioSummary    *ScenarioSummary      `json:"overallScenarioSummary,omitempty"`
	OverallTestSummary        *TestSummary          `json:"overallTestSummary,omitempty"`
	OverallTransactionSummary *TransactionSummary   `json:"overallTransactionSummary,omitempty"`
	TopSlowPages              *[]PageSummary        `json:"topSlowPages,omitempty"`
	TopSlowRequests           *[]RequestSummary     `json:"topSlowRequests,omitempty"`
	TopSlowTests              *[]TestSummary        `json:"topSlowTests,omitempty"`
	TopSlowTransactions       *[]TransactionSummary `json:"topSlowTransactions,omitempty"`
}

type TestRun struct {
	// Vss User identity who created the test run.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// Gets the creation time of the test run
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// Vss User identity who deleted the test run.
	DeletedBy *webapi.IdentityRef `json:"deletedBy,omitempty"`
	// Gets the deleted time of the test run
	DeletedDate *azuredevops.Time `json:"deletedDate,omitempty"`
	// Gets the finish time of the test run
	FinishedDate *azuredevops.Time `json:"finishedDate,omitempty"`
	// Gets the unique identifier for the test run definition.
	Id                         *string                      `json:"id,omitempty"`
	LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
	// Gets the load test file of the test run definition.
	LoadTestFileName *string `json:"loadTestFileName,omitempty"`
	// Gets the name of the test run definition.
	Name *string `json:"name,omitempty"`
	// Gets the number of the test run (unique within a tenant)
	RunNumber *int `json:"runNumber,omitempty"`
	// Test run source like Ibiza,VSO,BuildVNext, etc.
	RunSource *string `json:"runSource,omitempty"`
	// Run specific details.
	RunSpecificDetails *LoadTestRunDetails `json:"runSpecificDetails,omitempty"`
	// Run type like VisualStudioLoadTest or JMeterLoadTest
	RunType *TestRunType `json:"runType,omitempty"`
	// State of the test run.
	State *TestRunState `json:"state,omitempty"`
	Url   *string       `json:"url,omitempty"`
	// Message associated to state change, contains details of infrastructure error.
	AbortMessage *TestRunAbortMessage `json:"abortMessage,omitempty"`
	// true if aut counter collection could not start due to some critical error for this run.
	AutInitializationError *bool `json:"autInitializationError,omitempty"`
	// Whether run is chargeable or not Its chargeable once we configured agent and sent start signal
	Chargeable *bool `json:"chargeable,omitempty"`
	// Whether run is chargeable or not The Charged VUser Minutes for the RUN
	ChargedVUserminutes *int `json:"chargedVUserminutes,omitempty"`
	// Test run description.
	Description *string `json:"description,omitempty"`
	// Gets the time when the test run execution finished
	ExecutionFinishedDate *azuredevops.Time `json:"executionFinishedDate,omitempty"`
	// Gets the time when the test run warmup finished(if warmup was specified) and load test started
	ExecutionStartedDate *azuredevops.Time `json:"executionStartedDate,omitempty"`
	// Gets the time when the test run was queued
	QueuedDate *azuredevops.Time `json:"queuedDate,omitempty"`
	// Retention state of the run
	RetentionState      *TestRunRetentionState `json:"retentionState,omitempty"`
	RunSourceIdentifier *string                `json:"runSourceIdentifier,omitempty"`
	// The uri to the run source.
	RunSourceUrl *string `json:"runSourceUrl,omitempty"`
	// Vss User identity who created the test run.
	StartedBy *webapi.IdentityRef `json:"startedBy,omitempty"`
	// When the test run started execution.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
	// Vss User identity who created the test run.
	StoppedBy *webapi.IdentityRef `json:"stoppedBy,omitempty"`
	// SubState is more granular description of the state
	SubState             *TestRunSubState        `json:"subState,omitempty"`
	SupersedeRunSettings *OverridableRunSettings `json:"supersedeRunSettings,omitempty"`
	// Drop associated with this test run
	TestDrop *TestDropRef `json:"testDrop,omitempty"`
	// The Test settings for the test run
	TestSettings *TestSettings `json:"testSettings,omitempty"`
	// Gets the time when the test run warmup started
	WarmUpStartedDate *azuredevops.Time `json:"warmUpStartedDate,omitempty"`
	// The uri to the vso detailed result.
	WebResultUrl *string `json:"webResultUrl,omitempty"`
}

type TestRunAbortMessage struct {
	Action     *string           `json:"action,omitempty"`
	Cause      *string           `json:"cause,omitempty"`
	Details    *[]string         `json:"details,omitempty"`
	LoggedDate *azuredevops.Time `json:"loggedDate,omitempty"`
	Source     *string           `json:"source,omitempty"`
}

type TestRunBasic struct {
	// Vss User identity who created the test run.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// Gets the creation time of the test run
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// Vss User identity who deleted the test run.
	DeletedBy *webapi.IdentityRef `json:"deletedBy,omitempty"`
	// Gets the deleted time of the test run
	DeletedDate *azuredevops.Time `json:"deletedDate,omitempty"`
	// Gets the finish time of the test run
	FinishedDate *azuredevops.Time `json:"finishedDate,omitempty"`
	// Gets the unique identifier for the test run definition.
	Id                         *string                      `json:"id,omitempty"`
	LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
	// Gets the load test file of the test run definition.
	LoadTestFileName *string `json:"loadTestFileName,omitempty"`
	// Gets the name of the test run definition.
	Name *string `json:"name,omitempty"`
	// Gets the number of the test run (unique within a tenant)
	RunNumber *int `json:"runNumber,omitempty"`
	// Test run source like Ibiza,VSO,BuildVNext, etc.
	RunSource *string `json:"runSource,omitempty"`
	// Run specific details.
	RunSpecificDetails *LoadTestRunDetails `json:"runSpecificDetails,omitempty"`
	// Run type like VisualStudioLoadTest or JMeterLoadTest
	RunType *TestRunType `json:"runType,omitempty"`
	// State of the test run.
	State *TestRunState `json:"state,omitempty"`
	Url   *string       `json:"url,omitempty"`
}

type TestRunCounterInstance struct {
	// CategoryName for this counter
	CategoryName *string `json:"categoryName,omitempty"`
	// Combination of source and SourceInstanceId
	CounterInstanceId *string `json:"counterInstanceId,omitempty"`
	// Name of the counter Eg: Errors/Sec
	CounterName *string `json:"counterName,omitempty"`
	// Units for this counter. Empty string for mere numbers
	CounterUnits *string `json:"counterUnits,omitempty"`
	// Instance Name Eg: _Avg,_Total etc
	InstanceName *string `json:"instanceName,omitempty"`
	// true if this counter instance is a default counter
	IsPreselectedCounter *bool `json:"isPreselectedCounter,omitempty"`
	// Machine from where this counter was collected Used in case of machine specific counters like - Agent CPU and memory etc.
	MachineName *string `json:"machineName,omitempty"`
	// Counter Groups to which this counter instance is part of
	PartOfCounterGroups *[]string `json:"partOfCounterGroups,omitempty"`
	// Summary result for this counter instance
	SummaryData *WebInstanceSummaryData `json:"summaryData,omitempty"`
	// A unique name for this counter instance
	UniqueName *string `json:"uniqueName,omitempty"`
}

type TestRunMessage struct {
	// Agent Id
	AgentId    *string           `json:"agentId,omitempty"`
	ErrorCode  *string           `json:"errorCode,omitempty"`
	LoggedDate *azuredevops.Time `json:"loggedDate,omitempty"`
	Message    *string           `json:"message,omitempty"`
	// Message Id
	MessageId     *string        `json:"messageId,omitempty"`
	MessageSource *MessageSource `json:"messageSource,omitempty"`
	MessageType   *MessageType   `json:"messageType,omitempty"`
	// Id of the test run
	TestRunId *string `json:"testRunId,omitempty"`
	Url       *string `json:"url,omitempty"`
}

type TestRunRetentionState string

type testRunRetentionStateValuesType struct {
	None              TestRunRetentionState
	MarkedForDeletion TestRunRetentionState
	Deleted           TestRunRetentionState
	Retain            TestRunRetentionState
}

var TestRunRetentionStateValues = testRunRetentionStateValuesType{
	None:              "none",
	MarkedForDeletion: "markedForDeletion",
	Deleted:           "deleted",
	Retain:            "retain",
}

type TestRunState string

type testRunStateValuesType struct {
	Pending    TestRunState
	Queued     TestRunState
	InProgress TestRunState
	Stopping   TestRunState
	Completed  TestRunState
	Aborted    TestRunState
	Error      TestRunState
}

var TestRunStateValues = testRunStateValuesType{
	Pending:    "pending",
	Queued:     "queued",
	InProgress: "inProgress",
	Stopping:   "stopping",
	Completed:  "completed",
	Aborted:    "aborted",
	Error:      "error",
}

type TestRunSubState string

type testRunSubStateValuesType struct {
	None                   TestRunSubState
	ValidatingTestRun      TestRunSubState
	AcquiringResources     TestRunSubState
	ConfiguringAgents      TestRunSubState
	ExecutingSetupScript   TestRunSubState
	WarmingUp              TestRunSubState
	RunningTest            TestRunSubState
	ExecutingCleanupScript TestRunSubState
	CollectingResults      TestRunSubState
	Success                TestRunSubState
	PartialSuccess         TestRunSubState
}

var TestRunSubStateValues = testRunSubStateValuesType{
	None:                   "none",
	ValidatingTestRun:      "validatingTestRun",
	AcquiringResources:     "acquiringResources",
	ConfiguringAgents:      "configuringAgents",
	ExecutingSetupScript:   "executingSetupScript",
	WarmingUp:              "warmingUp",
	RunningTest:            "runningTest",
	ExecutingCleanupScript: "executingCleanupScript",
	CollectingResults:      "collectingResults",
	Success:                "success",
	PartialSuccess:         "partialSuccess",
}

type TestRunType string

type testRunTypeValuesType struct {
	VisualStudioLoadTest TestRunType
	JMeterLoadTest       TestRunType
}

var TestRunTypeValues = testRunTypeValuesType{
	VisualStudioLoadTest: "visualStudioLoadTest",
	JMeterLoadTest:       "jMeterLoadTest",
}

type TestSettings struct {
	// Cleanup command
	CleanupCommand *string `json:"cleanupCommand,omitempty"`
	// Processor Architecture chosen
	HostProcessPlatform *ProcessorArchitecture `json:"hostProcessPlatform,omitempty"`
	// Setup command
	SetupCommand *string `json:"setupCommand,omitempty"`
}

type TestSummary struct {
	AverageTestTime *float64                 `json:"averageTestTime,omitempty"`
	FailedTests     *int                     `json:"failedTests,omitempty"`
	PassedTests     *int                     `json:"passedTests,omitempty"`
	PercentileData  *[]SummaryPercentileData `json:"percentileData,omitempty"`
	ScenarioName    *string                  `json:"scenarioName,omitempty"`
	TestName        *string                  `json:"testName,omitempty"`
	TotalTests      *int                     `json:"totalTests,omitempty"`
}

type TransactionSummary struct {
	AverageResponseTime    *float64                 `json:"averageResponseTime,omitempty"`
	AverageTransactionTime *float64                 `json:"averageTransactionTime,omitempty"`
	PercentileData         *[]SummaryPercentileData `json:"percentileData,omitempty"`
	ScenarioName           *string                  `json:"scenarioName,omitempty"`
	TestName               *string                  `json:"testName,omitempty"`
	TotalTransactions      *int                     `json:"totalTransactions,omitempty"`
	TransactionName        *string                  `json:"transactionName,omitempty"`
}

type Type struct {
	Count       *int       `json:"count,omitempty"`
	Occurrences *int       `json:"occurrences,omitempty"`
	SubTypes    *[]SubType `json:"subTypes,omitempty"`
	TypeName    *string    `json:"typeName,omitempty"`
	Url         *string    `json:"url,omitempty"`
}

type WebApiLoadTestMachineInput struct {
	MachineGroupId     *string                `json:"machineGroupId,omitempty"`
	MachineType        *LoadTestMachineType   `json:"machineType,omitempty"`
	SetupConfiguration *WebApiSetupParamaters `json:"setupConfiguration,omitempty"`
	SupportedRunTypes  *[]TestRunType         `json:"supportedRunTypes,omitempty"`
}

type WebApiMachineConfiguration string

type webApiMachineConfigurationValuesType struct {
	UseXcopyTmiAgent                WebApiMachineConfiguration
	DisablingStrongNameVerification WebApiMachineConfiguration
	TempFolderPath                  WebApiMachineConfiguration
	ConfigureTcpParameters          WebApiMachineConfiguration
}

var WebApiMachineConfigurationValues = webApiMachineConfigurationValuesType{
	UseXcopyTmiAgent:                "useXcopyTmiAgent",
	DisablingStrongNameVerification: "disablingStrongNameVerification",
	TempFolderPath:                  "tempFolderPath",
	ConfigureTcpParameters:          "configureTcpParameters",
}

type WebApiSetupParamaters struct {
	Configurations *map[WebApiMachineConfiguration]string `json:"configurations,omitempty"`
}

// This can eventually evolve as the ultimate JSON file that user can use to configure their machine(s) against CLT
type WebApiUserLoadTestMachineInput struct {
	MachineGroupId            *string                `json:"machineGroupId,omitempty"`
	MachineType               *LoadTestMachineType   `json:"machineType,omitempty"`
	SetupConfiguration        *WebApiSetupParamaters `json:"setupConfiguration,omitempty"`
	SupportedRunTypes         *[]TestRunType         `json:"supportedRunTypes,omitempty"`
	AgentGroupName            *string                `json:"agentGroupName,omitempty"`
	TenantId                  *string                `json:"tenantId,omitempty"`
	UserLoadAgentResourcesUri *string                `json:"userLoadAgentResourcesUri,omitempty"`
	VstsAccountUri            *string                `json:"vstsAccountUri,omitempty"`
}

type WebInstanceSummaryData struct {
	Average *float64 `json:"average,omitempty"`
	Max     *float64 `json:"max,omitempty"`
	Min     *float64 `json:"min,omitempty"`
}
