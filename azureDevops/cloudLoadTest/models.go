// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package cloudLoadTest

import (
    "time"
)

type AgentGroup struct {
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    CreationTime *time.Time `json:"creationTime,omitempty"`
    GroupId *string `json:"groupId,omitempty"`
    GroupName *string `json:"groupName,omitempty"`
    MachineAccessData *[]AgentGroupAccessData `json:"machineAccessData,omitempty"`
    MachineConfiguration *WebApiUserLoadTestMachineInput `json:"machineConfiguration,omitempty"`
    TenantId *string `json:"tenantId,omitempty"`
}

type AgentGroupAccessData struct {
    Details *string `json:"details,omitempty"`
    StorageConnectionString *string `json:"storageConnectionString,omitempty"`
    StorageEndPoint *string `json:"storageEndPoint,omitempty"`
    StorageName *string `json:"storageName,omitempty"`
    StorageType *string `json:"storageType,omitempty"`
}

type Application struct {
    ApplicationId *string `json:"applicationId,omitempty"`
    Description *string `json:"description,omitempty"`
    Name *string `json:"name,omitempty"`
    Path *string `json:"path,omitempty"`
    PathSeperator *string `json:"pathSeperator,omitempty"`
    Type *string `json:"type,omitempty"`
    Version *string `json:"version,omitempty"`
}

type ApplicationCounters struct {
    ApplicationId *string `json:"applicationId,omitempty"`
    Description *string `json:"description,omitempty"`
    Id *string `json:"id,omitempty"`
    IsDefault *bool `json:"isDefault,omitempty"`
    Name *string `json:"name,omitempty"`
    Path *string `json:"path,omitempty"`
}

type ApplicationType struct {
    ActionUriLink *string `json:"actionUriLink,omitempty"`
    AutPortalLink *string `json:"autPortalLink,omitempty"`
    IsEnabled *bool `json:"isEnabled,omitempty"`
    MaxComponentsAllowedForCollection *int `json:"maxComponentsAllowedForCollection,omitempty"`
    MaxCountersAllowed *int `json:"maxCountersAllowed,omitempty"`
    Type *string `json:"type,omitempty"`
}

type BrowserMix struct {
    BrowserName *string `json:"browserName,omitempty"`
    BrowserPercentage *int `json:"browserPercentage,omitempty"`
}

type CltCustomerIntelligenceData struct {
    Area *string `json:"area,omitempty"`
    Feature *string `json:"feature,omitempty"`
    Properties *map[string]interface{} `json:"properties,omitempty"`
}

type CounterGroup struct {
    GroupName *string `json:"groupName,omitempty"`
    Url *string `json:"url,omitempty"`
}

type CounterInstanceSamples struct {
    Count *int `json:"count,omitempty"`
    CounterInstanceId *string `json:"counterInstanceId,omitempty"`
    NextRefreshTime *time.Time `json:"nextRefreshTime,omitempty"`
    Values *[]CounterSample `json:"values,omitempty"`
}

type CounterSample struct {
    BaseValue *uint64 `json:"baseValue,omitempty"`
    ComputedValue *int `json:"computedValue,omitempty"`
    CounterFrequency *uint64 `json:"counterFrequency,omitempty"`
    CounterInstanceId *string `json:"counterInstanceId,omitempty"`
    CounterType *string `json:"counterType,omitempty"`
    IntervalEndDate *time.Time `json:"intervalEndDate,omitempty"`
    IntervalNumber *int `json:"intervalNumber,omitempty"`
    RawValue *uint64 `json:"rawValue,omitempty"`
    SystemFrequency *uint64 `json:"systemFrequency,omitempty"`
    TimeStamp *uint64 `json:"timeStamp,omitempty"`
}

type CounterSampleQueryDetails struct {
    CounterInstanceId *string `json:"counterInstanceId,omitempty"`
    FromInterval *int `json:"fromInterval,omitempty"`
    ToInterval *int `json:"toInterval,omitempty"`
}

type CounterSamplesResult struct {
    Count *int `json:"count,omitempty"`
    MaxBatchSize *int `json:"maxBatchSize,omitempty"`
    TotalSamplesCount *int `json:"totalSamplesCount,omitempty"`
    Values *[]CounterInstanceSamples `json:"values,omitempty"`
}

type Diagnostics struct {
    DiagnosticStoreConnectionString *string `json:"diagnosticStoreConnectionString,omitempty"`
    LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
    RelativePathToDiagnosticFiles *string `json:"relativePathToDiagnosticFiles,omitempty"`
}

type DropAccessData struct {
    DropContainerUrl *string `json:"dropContainerUrl,omitempty"`
    SasKey *string `json:"sasKey,omitempty"`
}

type ErrorDetails struct {
    LastErrorDate *time.Time `json:"lastErrorDate,omitempty"`
    MessageText *string `json:"messageText,omitempty"`
    Occurrences *int `json:"occurrences,omitempty"`
    Request *string `json:"request,omitempty"`
    ScenarioName *string `json:"scenarioName,omitempty"`
    StackTrace *string `json:"stackTrace,omitempty"`
    TestCaseName *string `json:"testCaseName,omitempty"`
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

type LoadGenerationGeoLocation struct {
    Location *string `json:"location,omitempty"`
    Percentage *int `json:"percentage,omitempty"`
}

type LoadTest struct {
}

type LoadTestDefinition struct {
    AgentCount *int `json:"agentCount,omitempty"`
    BrowserMixs *[]BrowserMix `json:"browserMixs,omitempty"`
    CoreCount *int `json:"coreCount,omitempty"`
    CoresPerAgent *int `json:"coresPerAgent,omitempty"`
    LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
    LoadPatternName *string `json:"loadPatternName,omitempty"`
    LoadTestName *string `json:"loadTestName,omitempty"`
    MaxVusers *int `json:"maxVusers,omitempty"`
    RunDuration *int `json:"runDuration,omitempty"`
    SamplingRate *int `json:"samplingRate,omitempty"`
    ThinkTime *int `json:"thinkTime,omitempty"`
    Urls *[]string `json:"urls,omitempty"`
}

type LoadTestErrors struct {
    Count *int `json:"count,omitempty"`
    Occurrences *int `json:"occurrences,omitempty"`
    Types *[]interface{} `json:"types,omitempty"`
    Url *string `json:"url,omitempty"`
}

type LoadTestMachineType string

type LoadTestRunDetails struct {
    AgentCount *int `json:"agentCount,omitempty"`
    CoreCount *int `json:"coreCount,omitempty"`
    CoresPerAgent *int `json:"coresPerAgent,omitempty"`
    Duration *int `json:"duration,omitempty"`
    LoadGeneratorMachinesType *LoadTestMachineType `json:"loadGeneratorMachinesType,omitempty"`
    SamplingInterval *int `json:"samplingInterval,omitempty"`
    WarmUpDuration *int `json:"warmUpDuration,omitempty"`
    VirtualUserCount *int `json:"virtualUserCount,omitempty"`
}

type LoadTestRunSettings struct {
    AgentCount *int `json:"agentCount,omitempty"`
    CoreCount *int `json:"coreCount,omitempty"`
    CoresPerAgent *int `json:"coresPerAgent,omitempty"`
    Duration *int `json:"duration,omitempty"`
    LoadGeneratorMachinesType *LoadTestMachineType `json:"loadGeneratorMachinesType,omitempty"`
    SamplingInterval *int `json:"samplingInterval,omitempty"`
    WarmUpDuration *int `json:"warmUpDuration,omitempty"`
}

type LoadTestTypes string

type MessageSource string

type MessageType string

type OverridableRunSettings struct {
    LoadGeneratorMachinesType *LoadTestMachineType `json:"loadGeneratorMachinesType,omitempty"`
    StaticAgentRunSettings *StaticAgentRunSetting `json:"staticAgentRunSettings,omitempty"`
}

type PageSummary struct {
    AveragePageTime *float64 `json:"averagePageTime,omitempty"`
    PageUrl *string `json:"pageUrl,omitempty"`
    PercentagePagesMeetingGoal *int `json:"percentagePagesMeetingGoal,omitempty"`
    PercentileData *[]SummaryPercentileData `json:"percentileData,omitempty"`
    ScenarioName *string `json:"scenarioName,omitempty"`
    TestName *string `json:"testName,omitempty"`
    TotalPages *int `json:"totalPages,omitempty"`
}

type ProcessorArchitecture string

type ReferenceLinks struct {
    Links *map[string]interface{} `json:"links,omitempty"`
}

type RequestSummary struct {
    AverageResponseTime *float64 `json:"averageResponseTime,omitempty"`
    FailedRequests *int `json:"failedRequests,omitempty"`
    PassedRequests *int `json:"passedRequests,omitempty"`
    PercentileData *[]SummaryPercentileData `json:"percentileData,omitempty"`
    RequestsPerSec *float64 `json:"requestsPerSec,omitempty"`
    RequestUrl *string `json:"requestUrl,omitempty"`
    ScenarioName *string `json:"scenarioName,omitempty"`
    TestName *string `json:"testName,omitempty"`
    TotalRequests *int `json:"totalRequests,omitempty"`
}

type ScenarioSummary struct {
    MaxUserLoad *int `json:"maxUserLoad,omitempty"`
    MinUserLoad *int `json:"minUserLoad,omitempty"`
    ScenarioName *string `json:"scenarioName,omitempty"`
}

type StaticAgentRunSetting struct {
    LoadGeneratorMachinesType *LoadTestMachineType `json:"loadGeneratorMachinesType,omitempty"`
    StaticAgentGroupName *string `json:"staticAgentGroupName,omitempty"`
}

type SubType struct {
    Count *int `json:"count,omitempty"`
    ErrorDetailList *[]ErrorDetails `json:"errorDetailList,omitempty"`
    Occurrences *int `json:"occurrences,omitempty"`
    SubTypeName *string `json:"subTypeName,omitempty"`
    Url *string `json:"url,omitempty"`
}

type SummaryPercentileData struct {
    Percentile *int `json:"percentile,omitempty"`
    PercentileValue *float64 `json:"percentileValue,omitempty"`
}

type TenantDetails struct {
    AccessDetails *[]AgentGroupAccessData `json:"accessDetails,omitempty"`
    Id *string `json:"id,omitempty"`
    StaticMachines *[]WebApiTestMachine `json:"staticMachines,omitempty"`
    UserLoadAgentInput *WebApiUserLoadTestMachineInput `json:"userLoadAgentInput,omitempty"`
    UserLoadAgentResourcesUri *string `json:"userLoadAgentResourcesUri,omitempty"`
    ValidGeoLocations *[]string `json:"validGeoLocations,omitempty"`
}

type TestDefinition struct {
    AccessData *DropAccessData `json:"accessData,omitempty"`
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    Id *string `json:"id,omitempty"`
    LastModifiedBy *IdentityRef `json:"lastModifiedBy,omitempty"`
    LastModifiedDate *time.Time `json:"lastModifiedDate,omitempty"`
    LoadTestType *LoadTestTypes `json:"loadTestType,omitempty"`
    Name *string `json:"name,omitempty"`
    Description *string `json:"description,omitempty"`
    LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
    LoadTestDefinitionSource *string `json:"loadTestDefinitionSource,omitempty"`
    RunSettings *LoadTestRunSettings `json:"runSettings,omitempty"`
    StaticAgentRunSettings *StaticAgentRunSetting `json:"staticAgentRunSettings,omitempty"`
    TestDetails *LoadTest `json:"testDetails,omitempty"`
}

type TestDefinitionBasic struct {
    AccessData *DropAccessData `json:"accessData,omitempty"`
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    Id *string `json:"id,omitempty"`
    LastModifiedBy *IdentityRef `json:"lastModifiedBy,omitempty"`
    LastModifiedDate *time.Time `json:"lastModifiedDate,omitempty"`
    LoadTestType *LoadTestTypes `json:"loadTestType,omitempty"`
    Name *string `json:"name,omitempty"`
}

type TestDrop struct {
    AccessData *DropAccessData `json:"accessData,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    DropType *string `json:"dropType,omitempty"`
    Id *string `json:"id,omitempty"`
    LoadTestDefinition *LoadTestDefinition `json:"loadTestDefinition,omitempty"`
    TestRunId *string `json:"testRunId,omitempty"`
}

type TestDropRef struct {
    Id *string `json:"id,omitempty"`
    Url *string `json:"url,omitempty"`
}

type TestResults struct {
    CloudLoadTestSolutionUrl *string `json:"cloudLoadTestSolutionUrl,omitempty"`
    CounterGroups *[]CounterGroup `json:"counterGroups,omitempty"`
    Diagnostics *Diagnostics `json:"diagnostics,omitempty"`
    ResultsUrl *string `json:"resultsUrl,omitempty"`
}

type TestResultsSummary struct {
    OverallPageSummary *PageSummary `json:"overallPageSummary,omitempty"`
    OverallRequestSummary *RequestSummary `json:"overallRequestSummary,omitempty"`
    OverallScenarioSummary *ScenarioSummary `json:"overallScenarioSummary,omitempty"`
    OverallTestSummary *TestSummary `json:"overallTestSummary,omitempty"`
    OverallTransactionSummary *TransactionSummary `json:"overallTransactionSummary,omitempty"`
    TopSlowPages *[]PageSummary `json:"topSlowPages,omitempty"`
    TopSlowRequests *[]RequestSummary `json:"topSlowRequests,omitempty"`
    TopSlowTests *[]TestSummary `json:"topSlowTests,omitempty"`
    TopSlowTransactions *[]TransactionSummary `json:"topSlowTransactions,omitempty"`
}

type TestRun struct {
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    DeletedBy *IdentityRef `json:"deletedBy,omitempty"`
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    FinishedDate *time.Time `json:"finishedDate,omitempty"`
    Id *string `json:"id,omitempty"`
    LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
    LoadTestFileName *string `json:"loadTestFileName,omitempty"`
    Name *string `json:"name,omitempty"`
    RunNumber *int `json:"runNumber,omitempty"`
    RunSource *string `json:"runSource,omitempty"`
    RunSpecificDetails *LoadTestRunDetails `json:"runSpecificDetails,omitempty"`
    RunType *TestRunType `json:"runType,omitempty"`
    State *TestRunState `json:"state,omitempty"`
    Url *string `json:"url,omitempty"`
    AbortMessage *TestRunAbortMessage `json:"abortMessage,omitempty"`
    AutInitializationError *bool `json:"autInitializationError,omitempty"`
    Chargeable *bool `json:"chargeable,omitempty"`
    ChargedVUserminutes *int `json:"chargedVUserminutes,omitempty"`
    Description *string `json:"description,omitempty"`
    ExecutionFinishedDate *time.Time `json:"executionFinishedDate,omitempty"`
    ExecutionStartedDate *time.Time `json:"executionStartedDate,omitempty"`
    QueuedDate *time.Time `json:"queuedDate,omitempty"`
    RetentionState *TestRunRetentionState `json:"retentionState,omitempty"`
    RunSourceIdentifier *string `json:"runSourceIdentifier,omitempty"`
    RunSourceUrl *string `json:"runSourceUrl,omitempty"`
    StartedBy *IdentityRef `json:"startedBy,omitempty"`
    StartedDate *time.Time `json:"startedDate,omitempty"`
    StoppedBy *IdentityRef `json:"stoppedBy,omitempty"`
    SubState *TestRunSubState `json:"subState,omitempty"`
    SupersedeRunSettings *OverridableRunSettings `json:"supersedeRunSettings,omitempty"`
    TestDrop *TestDropRef `json:"testDrop,omitempty"`
    TestSettings *TestSettings `json:"testSettings,omitempty"`
    WarmUpStartedDate *time.Time `json:"warmUpStartedDate,omitempty"`
    WebResultUrl *string `json:"webResultUrl,omitempty"`
}

type TestRunAbortMessage struct {
    Action *string `json:"action,omitempty"`
    Cause *string `json:"cause,omitempty"`
    Details *[]string `json:"details,omitempty"`
    LoggedDate *time.Time `json:"loggedDate,omitempty"`
    Source *string `json:"source,omitempty"`
}

type TestRunBasic struct {
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    DeletedBy *IdentityRef `json:"deletedBy,omitempty"`
    DeletedDate *time.Time `json:"deletedDate,omitempty"`
    FinishedDate *time.Time `json:"finishedDate,omitempty"`
    Id *string `json:"id,omitempty"`
    LoadGenerationGeoLocations *[]LoadGenerationGeoLocation `json:"loadGenerationGeoLocations,omitempty"`
    LoadTestFileName *string `json:"loadTestFileName,omitempty"`
    Name *string `json:"name,omitempty"`
    RunNumber *int `json:"runNumber,omitempty"`
    RunSource *string `json:"runSource,omitempty"`
    RunSpecificDetails *LoadTestRunDetails `json:"runSpecificDetails,omitempty"`
    RunType *TestRunType `json:"runType,omitempty"`
    State *TestRunState `json:"state,omitempty"`
    Url *string `json:"url,omitempty"`
}

type TestRunCounterInstance struct {
    CategoryName *string `json:"categoryName,omitempty"`
    CounterInstanceId *string `json:"counterInstanceId,omitempty"`
    CounterName *string `json:"counterName,omitempty"`
    CounterUnits *string `json:"counterUnits,omitempty"`
    InstanceName *string `json:"instanceName,omitempty"`
    IsPreselectedCounter *bool `json:"isPreselectedCounter,omitempty"`
    MachineName *string `json:"machineName,omitempty"`
    PartOfCounterGroups *[]string `json:"partOfCounterGroups,omitempty"`
    SummaryData *WebInstanceSummaryData `json:"summaryData,omitempty"`
    UniqueName *string `json:"uniqueName,omitempty"`
}

type TestRunMessage struct {
    AgentId *string `json:"agentId,omitempty"`
    ErrorCode *string `json:"errorCode,omitempty"`
    LoggedDate *time.Time `json:"loggedDate,omitempty"`
    Message *string `json:"message,omitempty"`
    MessageId *string `json:"messageId,omitempty"`
    MessageSource *MessageSource `json:"messageSource,omitempty"`
    MessageType *MessageType `json:"messageType,omitempty"`
    TestRunId *string `json:"testRunId,omitempty"`
    Url *string `json:"url,omitempty"`
}

type TestRunRetentionState string

type TestRunState string

type TestRunSubState string

type TestRunType string

type TestSettings struct {
    CleanupCommand *string `json:"cleanupCommand,omitempty"`
    HostProcessPlatform *ProcessorArchitecture `json:"hostProcessPlatform,omitempty"`
    SetupCommand *string `json:"setupCommand,omitempty"`
}

type TestSummary struct {
    AverageTestTime *float64 `json:"averageTestTime,omitempty"`
    FailedTests *int `json:"failedTests,omitempty"`
    PassedTests *int `json:"passedTests,omitempty"`
    PercentileData *[]SummaryPercentileData `json:"percentileData,omitempty"`
    ScenarioName *string `json:"scenarioName,omitempty"`
    TestName *string `json:"testName,omitempty"`
    TotalTests *int `json:"totalTests,omitempty"`
}

type TransactionSummary struct {
    AverageResponseTime *float64 `json:"averageResponseTime,omitempty"`
    AverageTransactionTime *float64 `json:"averageTransactionTime,omitempty"`
    PercentileData *[]SummaryPercentileData `json:"percentileData,omitempty"`
    ScenarioName *string `json:"scenarioName,omitempty"`
    TestName *string `json:"testName,omitempty"`
    TotalTransactions *int `json:"totalTransactions,omitempty"`
    TransactionName *string `json:"transactionName,omitempty"`
}

type WebApiLoadTestMachineInput struct {
    MachineGroupId *string `json:"machineGroupId,omitempty"`
    MachineType *LoadTestMachineType `json:"machineType,omitempty"`
    SetupConfiguration *WebApiSetupParamaters `json:"setupConfiguration,omitempty"`
    SupportedRunTypes *[]TestRunType `json:"supportedRunTypes,omitempty"`
}

type WebApiMachineConfiguration string

type WebApiSetupParamaters struct {
    Configurations *map[WebApiMachineConfiguration]string `json:"configurations,omitempty"`
}

type WebApiTestMachine struct {
    LastHeartBeat *time.Time `json:"lastHeartBeat,omitempty"`
    MachineName *string `json:"machineName,omitempty"`
    Status *string `json:"status,omitempty"`
}

type WebApiUserLoadTestMachineInput struct {
    MachineGroupId *string `json:"machineGroupId,omitempty"`
    MachineType *LoadTestMachineType `json:"machineType,omitempty"`
    SetupConfiguration *WebApiSetupParamaters `json:"setupConfiguration,omitempty"`
    SupportedRunTypes *[]TestRunType `json:"supportedRunTypes,omitempty"`
    AgentGroupName *string `json:"agentGroupName,omitempty"`
    TenantId *string `json:"tenantId,omitempty"`
    UserLoadAgentResourcesUri *string `json:"userLoadAgentResourcesUri,omitempty"`
    VSTSAccountUri *string `json:"vstsAccountUri,omitempty"`
}

type WebInstanceSummaryData struct {
    Average *float64 `json:"average,omitempty"`
    Max *float64 `json:"max,omitempty"`
    Min *float64 `json:"min,omitempty"`
}
