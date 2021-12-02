// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package analytics

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
)

type AnalyticsState string

type analyticsStateValuesType struct {
	Disabled  AnalyticsState
	Enabled   AnalyticsState
	Paused    AnalyticsState
	Deleting  AnalyticsState
	Preparing AnalyticsState
}

var AnalyticsStateValues = analyticsStateValuesType{
	// The Analytics feature is disabled, and no data exists.
	Disabled: "disabled",
	// The Analytics feature is enabled, including all features.
	Enabled: "enabled",
	// The Analytics feature is paused. No data is flowing, and experiences may be unavailable.
	Paused: "paused",
	// The Analytics feature is undergoing deletion before becoming disabled.
	Deleting: "deleting",
	// The Analytics feature is enabled and undergoing first time data load, and is not yet reached model ready. Some experiences may not be available yet.
	Preparing: "preparing",
}

type AnalyticsStateDetails struct {
	// The Analytics feature state.
	AnalyticsState *AnalyticsState `json:"analyticsState,omitempty"`
	// When this state was set.
	ChangedDate *azuredevops.Time `json:"changedDate,omitempty"`
}

// Definition Contract for Analytics Views.
type AnalyticsViewWorkItemsDefinition struct {
	// List of area paths to filter by.
	AreaPathFilters *[]AreaPathFilter `json:"areaPathFilters,omitempty"`
	// List of backlog names to filter by.
	Backlogs *[]string `json:"backlogs,omitempty"`
	// List of field criteria to filter by.
	FieldFilters *[]FieldFilter `json:"fieldFilters,omitempty"`
	// Selection of fields for an analytics view. Field set may include common fields, or custom fields selected by the user.
	FieldSet *FieldSet `json:"fieldSet,omitempty"`
	// History/time-period constraints to filter an analytics view by.
	HistoryConfiguration *HistoryConfiguration `json:"historyConfiguration,omitempty"`
	// Defines the user-selected parameter by which an analytics view is filtered. True - filter by team. False -filter by area path.
	IsTeamFilterBySelectionMode *bool `json:"isTeamFilterBySelectionMode,omitempty"`
	// List of projects and teams to filter by.
	ProjectTeamFilters *[]ProjectTeamFilter `json:"projectTeamFilters,omitempty"`
	// List of work item types to filter by.
	WorkItemTypes interface{} `json:"workItemTypes,omitempty"`
}

// Project and Area path filter metadata for analytics view definitions.
type AreaPathFilter struct {
	// Area Id to filter by.
	AreaId *string `json:"areaId,omitempty"`
	// Enumeration representing operations while filtering an analytics view by area path.
	Operation *FieldOperation `json:"operation,omitempty"`
	// Unique identifier (Guid) for a project.
	ProjectId *uuid.UUID `json:"projectId,omitempty"`
}

// Date range to filter analytics view rows.
type DateRange struct {
	// End date for date range interval. If null, data spans until current date.
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// Start date for date range interval.
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
}

// Field filter criteria metadata for analytics view definitions.
type FieldFilter struct {
	// Name of work item field.
	FieldName *string `json:"fieldName,omitempty"`
	// Query comparison operator for selected field.
	FieldOperation *FieldOperation `json:"fieldOperation,omitempty"`
	// Deprecated: Use Values instead
	FieldValue *[]string `json:"fieldValue,omitempty"`
	// Deprecated: Use Values instead
	FieldValuesOperator *FieldValuesOperator `json:"fieldValuesOperator,omitempty"`
	// Value(s) of selected field.
	Value interface{} `json:"value,omitempty"`
}

// Enumeration representing operations while filtering field values an analytics view
type FieldOperation string

type fieldOperationValuesType struct {
	Eq          FieldOperation
	Ne          FieldOperation
	Ge          FieldOperation
	Le          FieldOperation
	Gt          FieldOperation
	Lt          FieldOperation
	Contains    FieldOperation
	Notcontains FieldOperation
	Under       FieldOperation
	Notunder    FieldOperation
}

var FieldOperationValues = fieldOperationValuesType{
	Eq:          "eq",
	Ne:          "ne",
	Ge:          "ge",
	Le:          "le",
	Gt:          "gt",
	Lt:          "lt",
	Contains:    "contains",
	Notcontains: "notcontains",
	Under:       "under",
	Notunder:    "notunder",
}

// Fields metadata for analytics view definitions.
type FieldSet struct {
	// List of fields in the analytics view field set.
	Fields *[]string `json:"fields,omitempty"`
	// Enumeration representing type of work item field set.
	FieldType *FieldType `json:"fieldType,omitempty"`
}

// Enumeration representing type of work item field set.
type FieldType string

type fieldTypeValuesType struct {
	Common FieldType
	Custom FieldType
}

var FieldTypeValues = fieldTypeValuesType{
	// DEPRECATED: When this option is provided, the analytics view uses the common field set.
	Common: "common",
	// When this option is provided, the analytics view uses custom fields specified by the user.
	Custom: "custom",
}

// Enumeration represanting multi-value operation for a single filter  i.e. what is the condition that should be applied _between_ values
type FieldValuesOperator string

type fieldValuesOperatorValuesType struct {
	And FieldValuesOperator
	Or  FieldValuesOperator
}

var FieldValuesOperatorValues = fieldValuesOperatorValuesType{
	And: "and",
	Or:  "or",
}

// Values of a generic type used for filtering information in an analytics view. Part of metadata for an analytics view definition.
type FilterValues struct {
	// Enumeration to specify whether filtering is applied on an entity.
	Mode *ValueMode `json:"mode,omitempty"`
	// Value(s) of selected field.
	Operator *FieldValuesOperator `json:"operator,omitempty"`
	// List of entity values of a generic type.
	Values *[]interface{} `json:"values,omitempty"`
}

// History metadata for analytics view definitions.
type HistoryConfiguration struct {
	// Date range to filter analytics view rows.
	DateRange *DateRange `json:"dateRange,omitempty"`
	// If true, then a clause will be added to filter out work items with a CompletedDateSK before OldCompletedItemsCutoffDate or the start of the the rolling period when HistoryType is Rolling.  When not present, the value is assumed to be false. This also achieves backwards compatibility.
	ExcludeOldCompletedWorkItems *bool `json:"excludeOldCompletedWorkItems,omitempty"`
	// Enumeration representing interval type for Analytics View History.
	HistoryType *HistoryType `json:"historyType,omitempty"`
	// When HistoryType is anything EXCEPT Rolling, and ExcludeOldCompletedWorkItems is true, this is the date to use to cutoff old completed work items. Work items closed on the cutoff date will not be filtered out.  For HistoryType Rolling, it is always calculated to be the beginning of the rolling period instead of using this.
	OldCompletedItemsCutoffDate *azuredevops.Time `json:"oldCompletedItemsCutoffDate,omitempty"`
	// Rolling period interval by number of days. Null when history type is not set to rolling period.
	RollingDays *int `json:"rollingDays,omitempty"`
	// Trend Granularity for filtering an analytics view by history.
	TrendGranularity *TrendGranularity `json:"trendGranularity,omitempty"`
}

// Enumeration representing history interval type for Analytics View History.
type HistoryType string

type historyTypeValuesType struct {
	None    HistoryType
	Rolling HistoryType
	Range   HistoryType
	All     HistoryType
}

var HistoryTypeValues = historyTypeValuesType{
	// When this option is provided, the view does not take history into account.
	None: "none",
	// When this option is provided, the view history comprises of a rolling period in number of days.
	Rolling: "rolling",
	// When this option is provided, the view history spans a range of dates.
	Range: "range",
	// When this option is provided, the view takes all available history into account.
	All: "all",
}

type IngestResult struct {
	Count     *int  `json:"count,omitempty"`
	Terminate *bool `json:"terminate,omitempty"`
}

// Teams filter metadata for analytics view definitions.
type ProjectTeamFilter struct {
	// Unique identifier (Guid) for a project.
	ProjectId *uuid.UUID `json:"projectId,omitempty"`
	// List of some or all teams under a project.
	Teams interface{} `json:"teams,omitempty"`
}

type StageProviderShardInfo struct {
	ProviderShardId *int               `json:"providerShardId,omitempty"`
	Streams         *[]StageStreamInfo `json:"streams,omitempty"`
}

type StageShardInvalid struct {
	DisableCurrentStream *bool     `json:"disableCurrentStream,omitempty"`
	InvalidFields        *[]string `json:"invalidFields,omitempty"`
	KeysOnly             *bool     `json:"keysOnly,omitempty"`
}

type StageStreamInfo struct {
	CreatedTime           *azuredevops.Time `json:"createdTime,omitempty"`
	Current               *bool             `json:"current,omitempty"`
	InitialContentVersion *int              `json:"initialContentVersion,omitempty"`
	KeysOnly              *bool             `json:"keysOnly,omitempty"`
	LatestContentVersion  *int              `json:"latestContentVersion,omitempty"`
	Priority              *int              `json:"priority,omitempty"`
	StreamId              *int              `json:"streamId,omitempty"`
	Watermark             *string           `json:"watermark,omitempty"`
}

type StageTableInfo struct {
	Shards       *[]StageProviderShardInfo `json:"shards,omitempty"`
	StageVersion *int                      `json:"stageVersion,omitempty"`
	TableName    *string                   `json:"tableName,omitempty"`
}

// Trend Granularity for filtering an analytics view by history.
type TrendGranularity struct {
	// Enumeration representing granularity of data being represented in an analytics view. Value is null when history type is 'None'.
	GranularityType *TrendGranularityType `json:"granularityType,omitempty"`
	// End day of week for weekly trend granularity. Value is null when trend granularity is daily or monthly.
	WeeklyEndDay *string `json:"weeklyEndDay,omitempty"`
}

// Enumeration representing granularity of data being represented in an analytics view.
type TrendGranularityType string

type trendGranularityTypeValuesType struct {
	Daily   TrendGranularityType
	Weekly  TrendGranularityType
	Monthly TrendGranularityType
}

var TrendGranularityTypeValues = trendGranularityTypeValuesType{
	// When this option is provided, analytics view rows represent data using a day-by-day granularity.
	Daily: "daily",
	// When this option is provided, analytics view rows represent data using weekly granularity.
	Weekly: "weekly",
	// When this option is provided, analytics view rows represent data using monthly granularity.
	Monthly: "monthly",
}

// Enumeration to specify whether filtering is applied on an entity.
type ValueMode string

type valueModeValuesType struct {
	Filter ValueMode
	All    ValueMode
}

var ValueModeValues = valueModeValuesType{
	// When this option is provided, filtering is applied on the entity.
	Filter: "filter",
	// When this option is provided, all values of the entity are taken into consideration.
	All: "all",
}
