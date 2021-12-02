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
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
)

type Issue struct {
	Category *string            `json:"category,omitempty"`
	Data     *map[string]string `json:"data,omitempty"`
	Message  *string            `json:"message,omitempty"`
	Type     *IssueType         `json:"type,omitempty"`
}

type IssueType string

type issueTypeValuesType struct {
	Error   IssueType
	Warning IssueType
}

var IssueTypeValues = issueTypeValuesType{
	Error:   "error",
	Warning: "warning",
}

// Represents an option that may affect the way an agent runs the job.
type JobOption struct {
	Data *map[string]string `json:"data,omitempty"`
	// Gets the id of the option.
	Id *uuid.UUID `json:"id,omitempty"`
}

type MaskHint struct {
	Type  *MaskType `json:"type,omitempty"`
	Value *string   `json:"value,omitempty"`
}

type MaskType string

type maskTypeValuesType struct {
	Variable MaskType
	Regex    MaskType
}

var MaskTypeValues = maskTypeValuesType{
	Variable: "variable",
	Regex:    "regex",
}

type PlanEnvironment struct {
	Mask      *[]MaskHint              `json:"mask,omitempty"`
	Options   *map[uuid.UUID]JobOption `json:"options,omitempty"`
	Variables *map[string]string       `json:"variables,omitempty"`
}

// [Flags]
type PlanGroupStatus string

type planGroupStatusValuesType struct {
	Running PlanGroupStatus
	Queued  PlanGroupStatus
	All     PlanGroupStatus
}

var PlanGroupStatusValues = planGroupStatusValuesType{
	Running: "running",
	Queued:  "queued",
	All:     "all",
}

type ProjectReference struct {
	Id   *uuid.UUID `json:"id,omitempty"`
	Name *string    `json:"name,omitempty"`
}

type TaskAgentJob struct {
	Container         *string                 `json:"container,omitempty"`
	Id                *uuid.UUID              `json:"id,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	SidecarContainers *map[string]string      `json:"sidecarContainers,omitempty"`
	Steps             *[]TaskAgentJobStep     `json:"steps,omitempty"`
	Variables         *[]TaskAgentJobVariable `json:"variables,omitempty"`
}

type TaskAgentJobStep struct {
	Condition        *string               `json:"condition,omitempty"`
	ContinueOnError  *bool                 `json:"continueOnError,omitempty"`
	Enabled          *bool                 `json:"enabled,omitempty"`
	Env              *map[string]string    `json:"env,omitempty"`
	Id               *uuid.UUID            `json:"id,omitempty"`
	Inputs           *map[string]string    `json:"inputs,omitempty"`
	Name             *string               `json:"name,omitempty"`
	Task             *TaskAgentJobTask     `json:"task,omitempty"`
	TimeoutInMinutes *int                  `json:"timeoutInMinutes,omitempty"`
	Type             *TaskAgentJobStepType `json:"type,omitempty"`
}

type TaskAgentJobStepType string

type taskAgentJobStepTypeValuesType struct {
	Task   TaskAgentJobStepType
	Action TaskAgentJobStepType
}

var TaskAgentJobStepTypeValues = taskAgentJobStepTypeValuesType{
	Task:   "task",
	Action: "action",
}

type TaskAgentJobTask struct {
	Id      *uuid.UUID `json:"id,omitempty"`
	Name    *string    `json:"name,omitempty"`
	Version *string    `json:"version,omitempty"`
}

type TaskAgentJobVariable struct {
	Name   *string `json:"name,omitempty"`
	Secret *bool   `json:"secret,omitempty"`
	Value  *string `json:"value,omitempty"`
}

type TaskAttachment struct {
	Links         interface{}       `json:"_links,omitempty"`
	CreatedOn     *azuredevops.Time `json:"createdOn,omitempty"`
	LastChangedBy *uuid.UUID        `json:"lastChangedBy,omitempty"`
	LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
	Name          *string           `json:"name,omitempty"`
	RecordId      *uuid.UUID        `json:"recordId,omitempty"`
	TimelineId    *uuid.UUID        `json:"timelineId,omitempty"`
	Type          *string           `json:"type,omitempty"`
}

type TaskLog struct {
	Id            *int              `json:"id,omitempty"`
	Location      *string           `json:"location,omitempty"`
	CreatedOn     *azuredevops.Time `json:"createdOn,omitempty"`
	IndexLocation *string           `json:"indexLocation,omitempty"`
	LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
	LineCount     *uint64           `json:"lineCount,omitempty"`
	Path          *string           `json:"path,omitempty"`
}

type TaskLogReference struct {
	Id       *int    `json:"id,omitempty"`
	Location *string `json:"location,omitempty"`
}

type TaskOrchestrationContainer struct {
	ItemType        *TaskOrchestrationItemType  `json:"itemType,omitempty"`
	Children        *[]TaskOrchestrationItem    `json:"children,omitempty"`
	ContinueOnError *bool                       `json:"continueOnError,omitempty"`
	Data            *map[string]string          `json:"data,omitempty"`
	MaxConcurrency  *int                        `json:"maxConcurrency,omitempty"`
	Parallel        *bool                       `json:"parallel,omitempty"`
	Rollback        *TaskOrchestrationContainer `json:"rollback,omitempty"`
}

type TaskOrchestrationItem struct {
	ItemType *TaskOrchestrationItemType `json:"itemType,omitempty"`
}

type TaskOrchestrationItemType string

type taskOrchestrationItemTypeValuesType struct {
	Container TaskOrchestrationItemType
	Job       TaskOrchestrationItemType
}

var TaskOrchestrationItemTypeValues = taskOrchestrationItemTypeValuesType{
	Container: "container",
	Job:       "job",
}

type TaskOrchestrationOwner struct {
	Links interface{} `json:"_links,omitempty"`
	Id    *int        `json:"id,omitempty"`
	Name  *string     `json:"name,omitempty"`
}

type TaskOrchestrationPlan struct {
	ArtifactLocation  *string                     `json:"artifactLocation,omitempty"`
	ArtifactUri       *string                     `json:"artifactUri,omitempty"`
	Definition        *TaskOrchestrationOwner     `json:"definition,omitempty"`
	Owner             *TaskOrchestrationOwner     `json:"owner,omitempty"`
	PlanGroup         *string                     `json:"planGroup,omitempty"`
	PlanId            *uuid.UUID                  `json:"planId,omitempty"`
	PlanType          *string                     `json:"planType,omitempty"`
	ScopeIdentifier   *uuid.UUID                  `json:"scopeIdentifier,omitempty"`
	Version           *int                        `json:"version,omitempty"`
	Environment       *PlanEnvironment            `json:"environment,omitempty"`
	ExpandedYaml      *TaskLogReference           `json:"expandedYaml,omitempty"`
	FinishTime        *azuredevops.Time           `json:"finishTime,omitempty"`
	Implementation    *TaskOrchestrationContainer `json:"implementation,omitempty"`
	InitializationLog *TaskLogReference           `json:"initializationLog,omitempty"`
	RequestedById     *uuid.UUID                  `json:"requestedById,omitempty"`
	RequestedForId    *uuid.UUID                  `json:"requestedForId,omitempty"`
	Result            *TaskResult                 `json:"result,omitempty"`
	ResultCode        *string                     `json:"resultCode,omitempty"`
	StartTime         *azuredevops.Time           `json:"startTime,omitempty"`
	State             *TaskOrchestrationPlanState `json:"state,omitempty"`
	Timeline          *TimelineReference          `json:"timeline,omitempty"`
}

type TaskOrchestrationPlanGroupsQueueMetrics struct {
	Count  *int             `json:"count,omitempty"`
	Status *PlanGroupStatus `json:"status,omitempty"`
}

type TaskOrchestrationPlanReference struct {
	ArtifactLocation *string                 `json:"artifactLocation,omitempty"`
	ArtifactUri      *string                 `json:"artifactUri,omitempty"`
	Definition       *TaskOrchestrationOwner `json:"definition,omitempty"`
	Owner            *TaskOrchestrationOwner `json:"owner,omitempty"`
	PlanGroup        *string                 `json:"planGroup,omitempty"`
	PlanId           *uuid.UUID              `json:"planId,omitempty"`
	PlanType         *string                 `json:"planType,omitempty"`
	ScopeIdentifier  *uuid.UUID              `json:"scopeIdentifier,omitempty"`
	Version          *int                    `json:"version,omitempty"`
}

type TaskOrchestrationPlanState string

type taskOrchestrationPlanStateValuesType struct {
	InProgress TaskOrchestrationPlanState
	Queued     TaskOrchestrationPlanState
	Completed  TaskOrchestrationPlanState
	Throttled  TaskOrchestrationPlanState
}

var TaskOrchestrationPlanStateValues = taskOrchestrationPlanStateValuesType{
	InProgress: "inProgress",
	Queued:     "queued",
	Completed:  "completed",
	Throttled:  "throttled",
}

type TaskOrchestrationQueuedPlan struct {
	AssignTime      *azuredevops.Time       `json:"assignTime,omitempty"`
	Definition      *TaskOrchestrationOwner `json:"definition,omitempty"`
	Owner           *TaskOrchestrationOwner `json:"owner,omitempty"`
	PlanGroup       *string                 `json:"planGroup,omitempty"`
	PlanId          *uuid.UUID              `json:"planId,omitempty"`
	PoolId          *int                    `json:"poolId,omitempty"`
	QueuePosition   *int                    `json:"queuePosition,omitempty"`
	QueueTime       *azuredevops.Time       `json:"queueTime,omitempty"`
	ScopeIdentifier *uuid.UUID              `json:"scopeIdentifier,omitempty"`
}

type TaskOrchestrationQueuedPlanGroup struct {
	Definition    *TaskOrchestrationOwner        `json:"definition,omitempty"`
	Owner         *TaskOrchestrationOwner        `json:"owner,omitempty"`
	PlanGroup     *string                        `json:"planGroup,omitempty"`
	Plans         *[]TaskOrchestrationQueuedPlan `json:"plans,omitempty"`
	Project       *ProjectReference              `json:"project,omitempty"`
	QueuePosition *int                           `json:"queuePosition,omitempty"`
}

type TaskReference struct {
	Id      *uuid.UUID         `json:"id,omitempty"`
	Inputs  *map[string]string `json:"inputs,omitempty"`
	Name    *string            `json:"name,omitempty"`
	Version *string            `json:"version,omitempty"`
}

type TaskResult string

type taskResultValuesType struct {
	Succeeded           TaskResult
	SucceededWithIssues TaskResult
	Failed              TaskResult
	Canceled            TaskResult
	Skipped             TaskResult
	Abandoned           TaskResult
}

var TaskResultValues = taskResultValuesType{
	Succeeded:           "succeeded",
	SucceededWithIssues: "succeededWithIssues",
	Failed:              "failed",
	Canceled:            "canceled",
	Skipped:             "skipped",
	Abandoned:           "abandoned",
}

type Timeline struct {
	ChangeId      *int              `json:"changeId,omitempty"`
	Id            *uuid.UUID        `json:"id,omitempty"`
	Location      *string           `json:"location,omitempty"`
	LastChangedBy *uuid.UUID        `json:"lastChangedBy,omitempty"`
	LastChangedOn *azuredevops.Time `json:"lastChangedOn,omitempty"`
	Records       *[]TimelineRecord `json:"records,omitempty"`
}

type TimelineAttempt struct {
	// Gets or sets the attempt of the record.
	Attempt *int `json:"attempt,omitempty"`
	// Gets or sets the unique identifier for the record.
	Identifier *string `json:"identifier,omitempty"`
	// Gets or sets the record identifier located within the specified timeline.
	RecordId *uuid.UUID `json:"recordId,omitempty"`
	// Gets or sets the timeline identifier which owns the record representing this attempt.
	TimelineId *uuid.UUID `json:"timelineId,omitempty"`
}

type TimelineRecord struct {
	AgentSpecification interface{}             `json:"agentSpecification,omitempty"`
	Attempt            *int                    `json:"attempt,omitempty"`
	ChangeId           *int                    `json:"changeId,omitempty"`
	CurrentOperation   *string                 `json:"currentOperation,omitempty"`
	Details            *TimelineReference      `json:"details,omitempty"`
	ErrorCount         *int                    `json:"errorCount,omitempty"`
	FinishTime         *azuredevops.Time       `json:"finishTime,omitempty"`
	Id                 *uuid.UUID              `json:"id,omitempty"`
	Identifier         *string                 `json:"identifier,omitempty"`
	Issues             *[]Issue                `json:"issues,omitempty"`
	LastModified       *azuredevops.Time       `json:"lastModified,omitempty"`
	Location           *string                 `json:"location,omitempty"`
	Log                *TaskLogReference       `json:"log,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	Order              *int                    `json:"order,omitempty"`
	ParentId           *uuid.UUID              `json:"parentId,omitempty"`
	PercentComplete    *int                    `json:"percentComplete,omitempty"`
	PreviousAttempts   *[]TimelineAttempt      `json:"previousAttempts,omitempty"`
	QueueId            *int                    `json:"queueId,omitempty"`
	RefName            *string                 `json:"refName,omitempty"`
	Result             *TaskResult             `json:"result,omitempty"`
	ResultCode         *string                 `json:"resultCode,omitempty"`
	StartTime          *azuredevops.Time       `json:"startTime,omitempty"`
	State              *TimelineRecordState    `json:"state,omitempty"`
	Task               *TaskReference          `json:"task,omitempty"`
	Type               *string                 `json:"type,omitempty"`
	Variables          *map[string]interface{} `json:"variables,omitempty"`
	WarningCount       *int                    `json:"warningCount,omitempty"`
	WorkerName         *string                 `json:"workerName,omitempty"`
}

type TimelineRecordFeedLinesWrapper struct {
	Count     *int       `json:"count,omitempty"`
	EndLine   *uint64    `json:"endLine,omitempty"`
	StartLine *uint64    `json:"startLine,omitempty"`
	StepId    *uuid.UUID `json:"stepId,omitempty"`
	Value     *[]string  `json:"value,omitempty"`
}

type TimelineRecordState string

type timelineRecordStateValuesType struct {
	Pending    TimelineRecordState
	InProgress TimelineRecordState
	Completed  TimelineRecordState
}

var TimelineRecordStateValues = timelineRecordStateValuesType{
	Pending:    "pending",
	InProgress: "inProgress",
	Completed:  "completed",
}

type TimelineReference struct {
	ChangeId *int       `json:"changeId,omitempty"`
	Id       *uuid.UUID `json:"id,omitempty"`
	Location *string    `json:"location,omitempty"`
}

type VariableValue struct {
	IsReadOnly *bool   `json:"isReadOnly,omitempty"`
	IsSecret   *bool   `json:"isSecret,omitempty"`
	Value      *string `json:"value,omitempty"`
}
