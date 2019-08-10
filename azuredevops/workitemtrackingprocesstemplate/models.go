// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package workitemtrackingprocesstemplate

import (
	"github.com/google/uuid"
)

// Describes an admin behavior for a process.
type AdminBehavior struct {
	// Is the behavior abstract (i.e. can not be associated with any work item type).
	Abstract *bool `json:"abstract,omitempty"`
	// The color associated with the behavior.
	Color *string `json:"color,omitempty"`
	// Indicates if the behavior is custom.
	Custom *bool `json:"custom,omitempty"`
	// The description of the behavior.
	Description *string `json:"description,omitempty"`
	// List of behavior fields.
	Fields *[]AdminBehaviorField `json:"fields,omitempty"`
	// Behavior ID.
	Id *string `json:"id,omitempty"`
	// Parent behavior reference.
	Inherits *string `json:"inherits,omitempty"`
	// The behavior name.
	Name *string `json:"name,omitempty"`
	// Is the behavior overrides a behavior from system process.
	Overriden *bool `json:"overriden,omitempty"`
	// The rank.
	Rank *int `json:"rank,omitempty"`
}

// Describes an admin behavior field.
type AdminBehaviorField struct {
	// The behavior field identifier.
	BehaviorFieldId *string `json:"behaviorFieldId,omitempty"`
	// The behavior ID.
	Id *string `json:"id,omitempty"`
	// The behavior name.
	Name *string `json:"name,omitempty"`
}

// Describes result of a check template existence request.
type CheckTemplateExistenceResult struct {
	// Indicates whether a template exists.
	DoesTemplateExist *bool `json:"doesTemplateExist,omitempty"`
	// The name of the existing template.
	ExistingTemplateName *string `json:"existingTemplateName,omitempty"`
	// The existing template type identifier.
	ExistingTemplateTypeId *uuid.UUID `json:"existingTemplateTypeId,omitempty"`
	// The name of the requested template.
	RequestedTemplateName *string `json:"requestedTemplateName,omitempty"`
}

// Describes the result of a Process Import request.
type ProcessImportResult struct {
	// Check template existence result.
	CheckExistenceResult *CheckTemplateExistenceResult `json:"checkExistenceResult,omitempty"`
	// Help URL.
	HelpUrl *string `json:"helpUrl,omitempty"`
	// ID of the import operation.
	Id *uuid.UUID `json:"id,omitempty"`
	// Whether this imported process is new.
	IsNew *bool `json:"isNew,omitempty"`
	// The promote job identifier.
	PromoteJobId *uuid.UUID `json:"promoteJobId,omitempty"`
	// The list of validation results.
	ValidationResults *[]ValidationIssue `json:"validationResults,omitempty"`
}

// Describes result of process operation promote.
type ProcessPromoteStatus struct {
	// Number of projects for which promote is complete.
	Complete *int `json:"complete,omitempty"`
	// ID of the promote operation.
	Id *uuid.UUID `json:"id,omitempty"`
	// The error message associated with the promote operation. The string will be empty if there are no errors.
	Message *string `json:"message,omitempty"`
	// Number of projects for which promote is pending.
	Pending *int `json:"pending,omitempty"`
	// The remaining retries.
	RemainingRetries *int `json:"remainingRetries,omitempty"`
	// True if promote finished all the projects successfully. False if still in progress or any project promote failed.
	Successful *bool `json:"successful,omitempty"`
}

type ValidationIssue struct {
	Description *string              `json:"description,omitempty"`
	File        *string              `json:"file,omitempty"`
	HelpLink    *string              `json:"helpLink,omitempty"`
	IssueType   *ValidationIssueType `json:"issueType,omitempty"`
	Line        *int                 `json:"line,omitempty"`
}

type ValidationIssueType string

type validationIssueTypeValuesType struct {
	Warning ValidationIssueType
	Error   ValidationIssueType
}

var ValidationIssueTypeValues = validationIssueTypeValuesType{
	Warning: "warning",
	Error:   "error",
}
