// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package projectanalysis

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

type AggregationType string

type aggregationTypeValuesType struct {
	Hourly AggregationType
	Daily  AggregationType
}

var AggregationTypeValues = aggregationTypeValuesType{
	Hourly: "hourly",
	Daily:  "daily",
}

type AnalyzerDescriptor struct {
	Description  *string    `json:"description,omitempty"`
	Id           *uuid.UUID `json:"id,omitempty"`
	MajorVersion *int       `json:"majorVersion,omitempty"`
	MinorVersion *int       `json:"minorVersion,omitempty"`
	Name         *string    `json:"name,omitempty"`
	PatchVersion *int       `json:"patchVersion,omitempty"`
}

type CodeChangeTrendItem struct {
	Time  *azuredevops.Time `json:"time,omitempty"`
	Value *int              `json:"value,omitempty"`
}

type LanguageMetricsSecuredObject struct {
	NamespaceId         *uuid.UUID `json:"namespaceId,omitempty"`
	ProjectId           *uuid.UUID `json:"projectId,omitempty"`
	RequiredPermissions *int       `json:"requiredPermissions,omitempty"`
}

type LanguageStatistics struct {
	NamespaceId         *uuid.UUID `json:"namespaceId,omitempty"`
	ProjectId           *uuid.UUID `json:"projectId,omitempty"`
	RequiredPermissions *int       `json:"requiredPermissions,omitempty"`
	Bytes               *uint64    `json:"bytes,omitempty"`
	Files               *int       `json:"files,omitempty"`
	FilesPercentage     *float64   `json:"filesPercentage,omitempty"`
	LanguagePercentage  *float64   `json:"languagePercentage,omitempty"`
	Name                *string    `json:"name,omitempty"`
}

type ProjectActivityMetrics struct {
	AuthorsCount               *int                   `json:"authorsCount,omitempty"`
	CodeChangesCount           *int                   `json:"codeChangesCount,omitempty"`
	CodeChangesTrend           *[]CodeChangeTrendItem `json:"codeChangesTrend,omitempty"`
	ProjectId                  *uuid.UUID             `json:"projectId,omitempty"`
	PullRequestsCompletedCount *int                   `json:"pullRequestsCompletedCount,omitempty"`
	PullRequestsCreatedCount   *int                   `json:"pullRequestsCreatedCount,omitempty"`
}

type ProjectLanguageAnalytics struct {
	NamespaceId                 *uuid.UUID                     `json:"namespaceId,omitempty"`
	ProjectId                   *uuid.UUID                     `json:"projectId,omitempty"`
	RequiredPermissions         *int                           `json:"requiredPermissions,omitempty"`
	Id                          *uuid.UUID                     `json:"id,omitempty"`
	LanguageBreakdown           *[]LanguageStatistics          `json:"languageBreakdown,omitempty"`
	RepositoryLanguageAnalytics *[]RepositoryLanguageAnalytics `json:"repositoryLanguageAnalytics,omitempty"`
	ResultPhase                 *ResultPhase                   `json:"resultPhase,omitempty"`
	Url                         *string                        `json:"url,omitempty"`
}

type RepositoryActivityMetrics struct {
	CodeChangesCount *int                   `json:"codeChangesCount,omitempty"`
	CodeChangesTrend *[]CodeChangeTrendItem `json:"codeChangesTrend,omitempty"`
	RepositoryId     *uuid.UUID             `json:"repositoryId,omitempty"`
}

type RepositoryLanguageAnalytics struct {
	NamespaceId         *uuid.UUID            `json:"namespaceId,omitempty"`
	ProjectId           *uuid.UUID            `json:"projectId,omitempty"`
	RequiredPermissions *int                  `json:"requiredPermissions,omitempty"`
	Id                  *uuid.UUID            `json:"id,omitempty"`
	LanguageBreakdown   *[]LanguageStatistics `json:"languageBreakdown,omitempty"`
	Name                *string               `json:"name,omitempty"`
	ResultPhase         *ResultPhase          `json:"resultPhase,omitempty"`
	UpdatedTime         *azuredevops.Time     `json:"updatedTime,omitempty"`
}

type ResultPhase string

type resultPhaseValuesType struct {
	Preliminary ResultPhase
	Full        ResultPhase
}

var ResultPhaseValues = resultPhaseValuesType{
	Preliminary: "preliminary",
	Full:        "full",
}
