// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package audit

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
)

// Defines all the categories an AuditAction can be
type AuditActionCategory string

type auditActionCategoryValuesType struct {
	Unknown AuditActionCategory
	Modify  AuditActionCategory
	Remove  AuditActionCategory
	Create  AuditActionCategory
	Access  AuditActionCategory
}

var AuditActionCategoryValues = auditActionCategoryValuesType{
	// The category is not known
	Unknown: "unknown",
	// An artifact has been Modified
	Modify: "modify",
	// An artifact has been Removed
	Remove: "remove",
	// An artifact has been Created
	Create: "create",
	// An artifact has been Accessed
	Access: "access",
}

// The object returned when the audit log is queried. It contains the log and the information needed to query more audit entries.
type AuditLogQueryResult struct {
	// The continuation token to pass to get the next set of results
	ContinuationToken *string `json:"continuationToken,omitempty"`
	// The list of audit log entries
	DecoratedAuditLogEntries *[]DecoratedAuditLogEntry `json:"decoratedAuditLogEntries,omitempty"`
	// True when there are more matching results to be fetched, false otherwise.
	HasMore *bool `json:"hasMore,omitempty"`
}

type DecoratedAuditLogEntry struct {
	// The action if for the event, i.e Git.CreateRepo, Project.RenameProject
	ActionId *string `json:"actionId,omitempty"`
	// ActivityId
	ActivityId *uuid.UUID `json:"activityId,omitempty"`
	// The Actor's CUID
	ActorCUID *uuid.UUID `json:"actorCUID,omitempty"`
	// DisplayName of the user who initiated the action
	ActorDisplayName *string `json:"actorDisplayName,omitempty"`
	// URL of Actor's Profile image
	ActorImageUrl *string `json:"actorImageUrl,omitempty"`
	// The Actor's User Id
	ActorUserId *uuid.UUID `json:"actorUserId,omitempty"`
	// Area of Azure DevOps the action occurred
	Area *string `json:"area,omitempty"`
	// Type of authentication used by the actor
	AuthenticationMechanism *string `json:"authenticationMechanism,omitempty"`
	// Type of action executed
	Category *AuditActionCategory `json:"category,omitempty"`
	// DisplayName of the category
	CategoryDisplayName *string `json:"categoryDisplayName,omitempty"`
	// This allows related audit entries to be grouped together. Generally this occurs when a single action causes a cascade of audit entries. For example, project creation.
	CorrelationId *uuid.UUID `json:"correlationId,omitempty"`
	// External data such as CUIDs, item names, etc.
	Data *map[string]interface{} `json:"data,omitempty"`
	// Decorated details
	Details *string `json:"details,omitempty"`
	// EventId - Needs to be unique per service
	Id *string `json:"id,omitempty"`
	// IP Address where the event was originated
	IpAddress *string `json:"ipAddress,omitempty"`
	// DisplayName of the scope
	ScopeDisplayName *string `json:"scopeDisplayName,omitempty"`
	// The organization or project Id
	ScopeId *uuid.UUID `json:"scopeId,omitempty"`
	// The type of the scope, Organization or Project
	ScopeType *string `json:"scopeType,omitempty"`
	// The time when the event occurred in UTC
	Timestamp *azuredevops.Time `json:"timestamp,omitempty"`
	// The user agent from the request
	UserAgent *string `json:"userAgent,omitempty"`
}
