// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package task

import (
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azuredevops/v7.1"
)

// An issue (error, warning) associated with a pipeline run.
type Issue struct {
    // The category of the issue. <br />Example: Code - refers to compilation errors <br />Example: General - refers to generic errors
    Category *string `json:"category,omitempty"`
    // A dictionary containing details about the issue.
    Data *map[string]string `json:"data,omitempty"`
    // A description of issue.
    Message *string `json:"message,omitempty"`
    // The type (error, warning) of the issue.
    Type *IssueType `json:"type,omitempty"`
}

// The type of issue based on severity.
type IssueType string

type issueTypeValuesType struct {
    Error IssueType
    Warning IssueType
}

var IssueTypeValues = issueTypeValuesType{
    Error: "error",
    Warning: "warning",
}

// A pipeline job event to be processed by the execution plan.
type JobEvent struct {
    // The ID of the pipeline job affected by the event.
    JobId *uuid.UUID `json:"jobId,omitempty"`
    // The name of the pipeline job event.
    Name *string `json:"name,omitempty"`
}

// Represents an option that may affect the way an agent runs the job.
type JobOption struct {
    Data *map[string]string `json:"data,omitempty"`
    // Gets the id of the option.
    Id *uuid.UUID `json:"id,omitempty"`
}

type MaskHint struct {
    Type *MaskType `json:"type,omitempty"`
    Value *string `json:"value,omitempty"`
}

type MaskType string

type maskTypeValuesType struct {
    Variable MaskType
    Regex MaskType
}

var MaskTypeValues = maskTypeValuesType{
    Variable: "variable",
    Regex: "regex",
}

type PlanEnvironment struct {
    Mask *[]MaskHint `json:"mask,omitempty"`
    Options *map[uuid.UUID]JobOption `json:"options,omitempty"`
    Variables *map[string]string `json:"variables,omitempty"`
}

// [Flags]
type PlanGroupStatus string

type planGroupStatusValuesType struct {
    Running PlanGroupStatus
    Queued PlanGroupStatus
    All PlanGroupStatus
}

var PlanGroupStatusValues = planGroupStatusValuesType{
    Running: "running",
    Queued: "queued",
    All: "all",
}

type ProjectReference struct {
    Id *uuid.UUID `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}

type TaskAgentJob struct {
    Container *string `json:"container,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    SidecarContainers *map[string]string `json:"sidecarContainers,omitempty"`
    Steps *[]TaskAgentJobStep `json:"steps,omitempty"`
    Variables *[]TaskAgentJobVariable `json:"variables,omitempty"`
}

type TaskAgentJobStep struct {
    Condition *string `json:"condition,omitempty"`
    ContinueOnError *bool `json:"continueOnError,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    Env *map[string]string `json:"env,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    Inputs *map[string]string `json:"inputs,omitempty"`
    Name *string `json:"name,omitempty"`
    RetryCountOnTaskFailure *int `json:"retryCountOnTaskFailure,omitempty"`
    Task *TaskAgentJobTask `json:"task,omitempty"`
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
    Type *TaskAgentJobStepType `json:"type,omitempty"`
}

type TaskAgentJobStepType string

type taskAgentJobStepTypeValuesType struct {
    Task TaskAgentJobStepType
    Action TaskAgentJobStepType
}

var TaskAgentJobStepTypeValues = taskAgentJobStepTypeValuesType{
    Task: "task",
    Action: "action",
}

type TaskAgentJobTask struct {
    Id *uuid.UUID `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
    Version *string `json:"version,omitempty"`
}

type TaskAgentJobVariable struct {
    Name *string `json:"name,omitempty"`
    Secret *bool `json:"secret,omitempty"`
    Value *string `json:"value,omitempty"`
}

type TaskAttachment struct {
    Links interface{} `json:"_links,omitempty"`
    CreatedOn *azuredevops.Time `json:"createdOn,omitempty"`
    LastChangedBy *uuid.UUID `json:"lastChangedBy,omitempty"`
    LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
    Name *string `json:"name,omitempty"`
    RecordId *uuid.UUID `json:"recordId,omitempty"`
    TimelineId *uuid.UUID `json:"timelineId,omitempty"`
    Type *string `json:"type,omitempty"`
}

type TaskHubOidcToken struct {
    OidcToken *string `json:"oidcToken,omitempty"`
}

// A task log connected to a timeline record.
type TaskLog struct {
    // The ID of the task log.
    Id *int `json:"id,omitempty"`
    // The REST URL of the task log.
    Location *string `json:"location,omitempty"`
    // The time of the task log creation.
    CreatedOn *azuredevops.Time `json:"createdOn,omitempty"`
    // The REST URL of the task log when indexed.
    IndexLocation *string `json:"indexLocation,omitempty"`
    // The time of the last modification of the task log.
    LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
    // The number of the task log lines.
    LineCount *uint64 `json:"lineCount,omitempty"`
    // The path of the task log.
    Path *string `json:"path,omitempty"`
}

// A reference to a task log. This class contains information about the output printed to the timeline record's logs console during pipeline run.
type TaskLogReference struct {
    // The ID of the task log.
    Id *int `json:"id,omitempty"`
    // The REST URL of the task log.
    Location *string `json:"location,omitempty"`
}

type TaskOrchestrationContainer struct {
    ItemType *TaskOrchestrationItemType `json:"itemType,omitempty"`
    ContinueOnError *bool `json:"continueOnError,omitempty"`
    Data *map[string]string `json:"data,omitempty"`
    Children *[]TaskOrchestrationItem `json:"children,omitempty"`
    MaxConcurrency *int `json:"maxConcurrency,omitempty"`
    Parallel *bool `json:"parallel,omitempty"`
    Rollback *TaskOrchestrationContainer `json:"rollback,omitempty"`
}

type TaskOrchestrationItem struct {
    ItemType *TaskOrchestrationItemType `json:"itemType,omitempty"`
}

type TaskOrchestrationItemType string

type taskOrchestrationItemTypeValuesType struct {
    Container TaskOrchestrationItemType
    Job TaskOrchestrationItemType
}

var TaskOrchestrationItemTypeValues = taskOrchestrationItemTypeValuesType{
    Container: "container",
    Job: "job",
}

type TaskOrchestrationOwner struct {
    Links interface{} `json:"_links,omitempty"`
    Id *int `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}

type TaskOrchestrationPlan struct {
    ArtifactLocation *string `json:"artifactLocation,omitempty"`
    ArtifactUri *string `json:"artifactUri,omitempty"`
    Definition *TaskOrchestrationOwner `json:"definition,omitempty"`
    Owner *TaskOrchestrationOwner `json:"owner,omitempty"`
    PlanGroup *string `json:"planGroup,omitempty"`
    PlanId *uuid.UUID `json:"planId,omitempty"`
    PlanType *string `json:"planType,omitempty"`
    ScopeIdentifier *uuid.UUID `json:"scopeIdentifier,omitempty"`
    Version *int `json:"version,omitempty"`
    Environment *PlanEnvironment `json:"environment,omitempty"`
    ExpandedYaml *TaskLogReference `json:"expandedYaml,omitempty"`
    FinishTime *azuredevops.Time `json:"finishTime,omitempty"`
    Implementation *TaskOrchestrationContainer `json:"implementation,omitempty"`
    InitializationLog *TaskLogReference `json:"initializationLog,omitempty"`
    RequestedById *uuid.UUID `json:"requestedById,omitempty"`
    RequestedForId *uuid.UUID `json:"requestedForId,omitempty"`
    Result *TaskResult `json:"result,omitempty"`
    ResultCode *string `json:"resultCode,omitempty"`
    StartTime *azuredevops.Time `json:"startTime,omitempty"`
    State *TaskOrchestrationPlanState `json:"state,omitempty"`
    Timeline *TimelineReference `json:"timeline,omitempty"`
}

type TaskOrchestrationPlanGroupsQueueMetrics struct {
    Count *int `json:"count,omitempty"`
    Status *PlanGroupStatus `json:"status,omitempty"`
}

type TaskOrchestrationPlanReference struct {
    ArtifactLocation *string `json:"artifactLocation,omitempty"`
    ArtifactUri *string `json:"artifactUri,omitempty"`
    Definition *TaskOrchestrationOwner `json:"definition,omitempty"`
    Owner *TaskOrchestrationOwner `json:"owner,omitempty"`
    PlanGroup *string `json:"planGroup,omitempty"`
    PlanId *uuid.UUID `json:"planId,omitempty"`
    PlanType *string `json:"planType,omitempty"`
    ScopeIdentifier *uuid.UUID `json:"scopeIdentifier,omitempty"`
    Version *int `json:"version,omitempty"`
}

type TaskOrchestrationPlanState string

type taskOrchestrationPlanStateValuesType struct {
    InProgress TaskOrchestrationPlanState
    Queued TaskOrchestrationPlanState
    Completed TaskOrchestrationPlanState
    Throttled TaskOrchestrationPlanState
}

var TaskOrchestrationPlanStateValues = taskOrchestrationPlanStateValuesType{
    InProgress: "inProgress",
    Queued: "queued",
    Completed: "completed",
    Throttled: "throttled",
}

type TaskOrchestrationQueuedPlan struct {
    AssignTime *azuredevops.Time `json:"assignTime,omitempty"`
    Definition *TaskOrchestrationOwner `json:"definition,omitempty"`
    Owner *TaskOrchestrationOwner `json:"owner,omitempty"`
    PlanGroup *string `json:"planGroup,omitempty"`
    PlanId *uuid.UUID `json:"planId,omitempty"`
    PoolId *int `json:"poolId,omitempty"`
    QueuePosition *int `json:"queuePosition,omitempty"`
    QueueTime *azuredevops.Time `json:"queueTime,omitempty"`
    ScopeIdentifier *uuid.UUID `json:"scopeIdentifier,omitempty"`
}

type TaskOrchestrationQueuedPlanGroup struct {
    Definition *TaskOrchestrationOwner `json:"definition,omitempty"`
    Owner *TaskOrchestrationOwner `json:"owner,omitempty"`
    PlanGroup *string `json:"planGroup,omitempty"`
    Plans *[]TaskOrchestrationQueuedPlan `json:"plans,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
    QueuePosition *int `json:"queuePosition,omitempty"`
}

// A reference to a task.
type TaskReference struct {
    // The ID of the task definition. Corresponds to the id value of task.json file. <br />Example: CmdLineV2 { "id": "D9BAFED4-0B18-4F58-968D-86655B4D2CE9" }
    Id *uuid.UUID `json:"id,omitempty"`
    // A dictionary of inputs specific to a task definition. Corresponds to inputs value of task.json file.
    Inputs *map[string]string `json:"inputs,omitempty"`
    // The name of the task definition. Corresponds to the name value of task.json file. <br />Example: CmdLineV2 { "name": "CmdLine" }
    Name *string `json:"name,omitempty"`
    // The version of the task definition. Corresponds to the version value of task.json file. <br />Example: CmdLineV2 { "version": { "Major": 2, "Minor": 212, "Patch": 0 } }
    Version *string `json:"version,omitempty"`
}

// The result of an operation tracked by a timeline record.
type TaskResult string

type taskResultValuesType struct {
    Succeeded TaskResult
    SucceededWithIssues TaskResult
    Failed TaskResult
    Canceled TaskResult
    Skipped TaskResult
    Abandoned TaskResult
}

var TaskResultValues = taskResultValuesType{
    Succeeded: "succeeded",
    SucceededWithIssues: "succeededWithIssues",
    Failed: "failed",
    Canceled: "canceled",
    Skipped: "skipped",
    Abandoned: "abandoned",
}

type Timeline struct {
    // The change ID.
    ChangeId *int `json:"changeId,omitempty"`
    // The ID of the timeline.
    Id *uuid.UUID `json:"id,omitempty"`
    // The REST URL of the timeline.
    Location *string `json:"location,omitempty"`
    LastChangedBy *uuid.UUID `json:"lastChangedBy,omitempty"`
    LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
    Records *[]TimelineRecord `json:"records,omitempty"`
}

// An attempt to update a TimelineRecord.
type TimelineAttempt struct {
    // The attempt of the record.
    Attempt *int `json:"attempt,omitempty"`
    // The unique identifier for the record.
    Identifier *string `json:"identifier,omitempty"`
    // The record identifier located within the specified timeline.
    RecordId *uuid.UUID `json:"recordId,omitempty"`
    // The timeline identifier which owns the record representing this attempt.
    TimelineId *uuid.UUID `json:"timelineId,omitempty"`
}

// Detailed information about the execution of different operations during pipeline run.
type TimelineRecord struct {
    // The specification of an agent running a pipeline job, in binary format. Applicable when record is of type Job. <br />Example: { "VMImage" : "windows-2019" }
    AgentSpecification interface{} `json:"agentSpecification,omitempty"`
    // The number of record attempts.
    Attempt *int `json:"attempt,omitempty"`
    // A string that indicates the current operation.
    CurrentOperation *string `json:"currentOperation,omitempty"`
    // A reference to a sub-timeline.
    Details *TimelineReference `json:"details,omitempty"`
    // The number of errors produced by this operation.
    ErrorCount *int `json:"errorCount,omitempty"`
    // The finish time of the record.
    FinishTime *azuredevops.Time `json:"finishTime,omitempty"`
    // The ID connecting all records updated at the same time. This value is taken from timeline's ChangeId.
    ChangeId *int `json:"changeId,omitempty"`
    // The ID of the record.
    Id *uuid.UUID `json:"id,omitempty"`
    // String identifier that is consistent across attempts.
    Identifier *string `json:"identifier,omitempty"`
    // The list of issues produced by this operation.
    Issues *[]Issue `json:"issues,omitempty"`
    // The time the record was last modified.
    LastModified *azuredevops.Time `json:"lastModified,omitempty"`
    // The REST URL of the record.
    Location *string `json:"location,omitempty"`
    // A reference to the log produced by this operation.
    Log *TaskLogReference `json:"log,omitempty"`
    // The name of the record.
    Name *string `json:"name,omitempty"`
    // An ordinal value relative to other records within the timeline.
    Order *int `json:"order,omitempty"`
    // The ID of the record's parent. <br />Example: Stage is a parent of a Phase, Phase is a parent of a Job, Job is a parent of a Task.
    ParentId *uuid.UUID `json:"parentId,omitempty"`
    // The percentage of record completion.
    PercentComplete *int `json:"percentComplete,omitempty"`
    // The previous record attempts.
    PreviousAttempts *[]TimelineAttempt `json:"previousAttempts,omitempty"`
    // The ID of the queue which connects projects to agent pools on which the operation ran on. Applicable when record is of type Job.
    QueueId *int `json:"queueId,omitempty"`
    // Name of the referenced record.
    RefName *string `json:"refName,omitempty"`
    // The result of the record.
    Result *TaskResult `json:"result,omitempty"`
    // Evaluation of predefined conditions upon completion of record's operation. <br />Example: Evaluating `succeeded()`, Result = True <br />Example: Evaluating `and(succeeded(), eq(variables['system.debug'], False))`, Result = False
    ResultCode *string `json:"resultCode,omitempty"`
    // The start time of the record.
    StartTime *azuredevops.Time `json:"startTime,omitempty"`
    // The state of the record.
    State *TimelineRecordState `json:"state,omitempty"`
    // A reference to the task. Applicable when record is of type Task.
    Task *TaskReference `json:"task,omitempty"`
    // The type of operation being tracked by the record. <br />Example: Stage, Phase, Job, Task...
    Type *string `json:"type,omitempty"`
    // The variables of the record.
    Variables *map[string]interface{} `json:"variables,omitempty"`
    // The number of warnings produced by this operation.
    WarningCount *int `json:"warningCount,omitempty"`
    // The name of the agent running the operation. Applicable when record is of type Job.
    WorkerName *string `json:"workerName,omitempty"`
}

type TimelineRecordFeedLinesWrapper struct {
    Count *int `json:"count,omitempty"`
    EndLine *uint64 `json:"endLine,omitempty"`
    StartLine *uint64 `json:"startLine,omitempty"`
    StepId *uuid.UUID `json:"stepId,omitempty"`
    Value *[]string `json:"value,omitempty"`
}

// The state of the timeline record.
type TimelineRecordState string

type timelineRecordStateValuesType struct {
    Pending TimelineRecordState
    InProgress TimelineRecordState
    Completed TimelineRecordState
}

var TimelineRecordStateValues = timelineRecordStateValuesType{
    Pending: "pending",
    InProgress: "inProgress",
    Completed: "completed",
}

// A reference to a timeline.
type TimelineReference struct {
    // The change ID.
    ChangeId *int `json:"changeId,omitempty"`
    // The ID of the timeline.
    Id *uuid.UUID `json:"id,omitempty"`
    // The REST URL of the timeline.
    Location *string `json:"location,omitempty"`
}

// A wrapper class for a generic variable.
type VariableValue struct {
    // Indicates whether the variable can be changed during script's execution runtime.
    IsReadOnly *bool `json:"isReadOnly,omitempty"`
    // Indicates whether the variable should be encrypted at rest.
    IsSecret *bool `json:"isSecret,omitempty"`
    // The value of the variable.
    Value *string `json:"value,omitempty"`
}
