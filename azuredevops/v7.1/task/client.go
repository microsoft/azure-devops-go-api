// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package task

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7.1"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	// [Preview API] Append a log to a task's log. The log should be sent in the body of the request as a TaskLog object stream.
	AppendLogContent(context.Context, AppendLogContentArgs) (*TaskLog, error)
	// [Preview API]
	AssociateLog(context.Context, AssociateLogArgs) (*TaskLog, error)
	// [Preview API]
	CreateAttachment(context.Context, CreateAttachmentArgs) (*TaskAttachment, error)
	// [Preview API]
	CreateAttachmentFromArtifact(context.Context, CreateAttachmentFromArtifactArgs) (*TaskAttachment, error)
	// [Preview API] Create a log and connect it to a pipeline run's execution plan.
	CreateLog(context.Context, CreateLogArgs) (*TaskLog, error)
	// [Preview API]
	CreateOidcToken(context.Context, CreateOidcTokenArgs) (*TaskHubOidcToken, error)
	// [Preview API]
	CreateTimeline(context.Context, CreateTimelineArgs) (*Timeline, error)
	// [Preview API]
	DeleteTimeline(context.Context, DeleteTimelineArgs) error
	// [Preview API]
	GetAttachment(context.Context, GetAttachmentArgs) (*TaskAttachment, error)
	// [Preview API]
	GetAttachmentContent(context.Context, GetAttachmentContentArgs) (io.ReadCloser, error)
	// [Preview API]
	GetAttachments(context.Context, GetAttachmentsArgs) (*[]TaskAttachment, error)
	// [Preview API]
	GetLog(context.Context, GetLogArgs) (*[]string, error)
	// [Preview API]
	GetLogs(context.Context, GetLogsArgs) (*[]TaskLog, error)
	// [Preview API]
	GetPlanAttachments(context.Context, GetPlanAttachmentsArgs) (*[]TaskAttachment, error)
	// [Preview API]
	GetRecords(context.Context, GetRecordsArgs) (*[]TimelineRecord, error)
	// [Preview API]
	GetTimeline(context.Context, GetTimelineArgs) (*Timeline, error)
	// [Preview API]
	GetTimelines(context.Context, GetTimelinesArgs) (*[]Timeline, error)
	// [Preview API] Update timeline records if they already exist, otherwise create new ones for the same timeline.
	UpdateRecords(context.Context, UpdateRecordsArgs) (*[]TimelineRecord, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) Client {
	client := connection.GetClientByUrl(connection.BaseUrl)
	return &ClientImpl{
		Client: *client,
	}
}

// [Preview API] Append a log to a task's log. The log should be sent in the body of the request as a TaskLog object stream.
func (client *ClientImpl) AppendLogContent(ctx context.Context, args AppendLogContentArgs) (*TaskLog, error) {
	if args.UploadStream == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.UploadStream"}
	}
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.LogId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.LogId"}
	}
	routeValues["logId"] = strconv.Itoa(*args.LogId)

	locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, args.UploadStream, "application/octet-stream", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskLog
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AppendLogContent function
type AppendLogContentArgs struct {
	// (required) Stream to upload
	UploadStream io.Reader
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required) The ID of the plan.
	PlanId *uuid.UUID
	// (required) The ID of the log.
	LogId *int
}

// [Preview API]
func (client *ClientImpl) AssociateLog(ctx context.Context, args AssociateLogArgs) (*TaskLog, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.LogId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.LogId"}
	}
	routeValues["logId"] = strconv.Itoa(*args.LogId)

	queryParams := url.Values{}
	if args.SerializedBlobId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "serializedBlobId"}
	}
	queryParams.Add("serializedBlobId", *args.SerializedBlobId)
	if args.LineCount == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "lineCount"}
	}
	queryParams.Add("lineCount", strconv.Itoa(*args.LineCount))
	locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskLog
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the AssociateLog function
type AssociateLogArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	LogId *int
	// (required)
	SerializedBlobId *string
	// (required)
	LineCount *int
}

// [Preview API]
func (client *ClientImpl) CreateAttachment(ctx context.Context, args CreateAttachmentArgs) (*TaskAttachment, error) {
	if args.UploadStream == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.UploadStream"}
	}
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()
	if args.RecordId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RecordId"}
	}
	routeValues["recordId"] = (*args.RecordId).String()
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "7.1-preview.1", routeValues, nil, args.UploadStream, "application/octet-stream", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskAttachment
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateAttachment function
type CreateAttachmentArgs struct {
	// (required) Stream to upload
	UploadStream io.Reader
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
	// (required)
	RecordId *uuid.UUID
	// (required)
	Type *string
	// (required)
	Name *string
}

// [Preview API]
func (client *ClientImpl) CreateAttachmentFromArtifact(ctx context.Context, args CreateAttachmentFromArtifactArgs) (*TaskAttachment, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()
	if args.RecordId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RecordId"}
	}
	routeValues["recordId"] = (*args.RecordId).String()
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	queryParams := url.Values{}
	if args.ArtifactHash == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "artifactHash"}
	}
	queryParams.Add("artifactHash", *args.ArtifactHash)
	if args.Length == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "length"}
	}
	queryParams.Add("length", strconv.FormatUint(*args.Length, 10))
	locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskAttachment
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateAttachmentFromArtifact function
type CreateAttachmentFromArtifactArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
	// (required)
	RecordId *uuid.UUID
	// (required)
	Type *string
	// (required)
	Name *string
	// (required)
	ArtifactHash *string
	// (required)
	Length *uint64
}

// [Preview API] Create a log and connect it to a pipeline run's execution plan.
func (client *ClientImpl) CreateLog(ctx context.Context, args CreateLogArgs) (*TaskLog, error) {
	if args.Log == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Log"}
	}
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()

	body, marshalErr := json.Marshal(*args.Log)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskLog
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateLog function
type CreateLogArgs struct {
	// (required) An object that contains information about log's path.
	Log *TaskLog
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required) The ID of the plan.
	PlanId *uuid.UUID
}

// [Preview API]
func (client *ClientImpl) CreateOidcToken(ctx context.Context, args CreateOidcTokenArgs) (*TaskHubOidcToken, error) {
	if args.Claims == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Claims"}
	}
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.JobId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.JobId"}
	}
	routeValues["jobId"] = (*args.JobId).String()

	queryParams := url.Values{}
	if args.ServiceConnectionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "serviceConnectionId"}
	}
	queryParams.Add("serviceConnectionId", (*args.ServiceConnectionId).String())
	body, marshalErr := json.Marshal(*args.Claims)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("69a319f4-28c1-4bfd-93e6-ea0ff5c6f1a2")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskHubOidcToken
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateOidcToken function
type CreateOidcTokenArgs struct {
	// (required)
	Claims *map[string]string
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	JobId *uuid.UUID
	// (required)
	ServiceConnectionId *uuid.UUID
}

// [Preview API]
func (client *ClientImpl) CreateTimeline(ctx context.Context, args CreateTimelineArgs) (*Timeline, error) {
	if args.Timeline == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Timeline"}
	}
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()

	body, marshalErr := json.Marshal(*args.Timeline)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Timeline
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateTimeline function
type CreateTimelineArgs struct {
	// (required)
	Timeline *Timeline
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
}

// [Preview API]
func (client *ClientImpl) DeleteTimeline(ctx context.Context, args DeleteTimelineArgs) error {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()

	locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteTimeline function
type DeleteTimelineArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
}

// [Preview API]
func (client *ClientImpl) GetAttachment(ctx context.Context, args GetAttachmentArgs) (*TaskAttachment, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()
	if args.RecordId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RecordId"}
	}
	routeValues["recordId"] = (*args.RecordId).String()
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue TaskAttachment
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAttachment function
type GetAttachmentArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
	// (required)
	RecordId *uuid.UUID
	// (required)
	Type *string
	// (required)
	Name *string
}

// [Preview API]
func (client *ClientImpl) GetAttachmentContent(ctx context.Context, args GetAttachmentContentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()
	if args.RecordId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RecordId"}
	}
	routeValues["recordId"] = (*args.RecordId).String()
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetAttachmentContent function
type GetAttachmentContentArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
	// (required)
	RecordId *uuid.UUID
	// (required)
	Type *string
	// (required)
	Name *string
}

// [Preview API]
func (client *ClientImpl) GetAttachments(ctx context.Context, args GetAttachmentsArgs) (*[]TaskAttachment, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()
	if args.RecordId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RecordId"}
	}
	routeValues["recordId"] = (*args.RecordId).String()
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type

	locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TaskAttachment
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAttachments function
type GetAttachmentsArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
	// (required)
	RecordId *uuid.UUID
	// (required)
	Type *string
}

// [Preview API]
func (client *ClientImpl) GetLog(ctx context.Context, args GetLogArgs) (*[]string, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.LogId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.LogId"}
	}
	routeValues["logId"] = strconv.Itoa(*args.LogId)

	queryParams := url.Values{}
	if args.StartLine != nil {
		queryParams.Add("startLine", strconv.FormatUint(*args.StartLine, 10))
	}
	if args.EndLine != nil {
		queryParams.Add("endLine", strconv.FormatUint(*args.EndLine, 10))
	}
	locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []string
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLog function
type GetLogArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	LogId *int
	// (optional)
	StartLine *uint64
	// (optional)
	EndLine *uint64
}

// [Preview API]
func (client *ClientImpl) GetLogs(ctx context.Context, args GetLogsArgs) (*[]TaskLog, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()

	locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TaskLog
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLogs function
type GetLogsArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
}

// [Preview API]
func (client *ClientImpl) GetPlanAttachments(ctx context.Context, args GetPlanAttachmentsArgs) (*[]TaskAttachment, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.Type == nil || *args.Type == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Type"}
	}
	routeValues["type"] = *args.Type

	locationId, _ := uuid.Parse("eb55e5d6-2f30-4295-b5ed-38da50b1fc52")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TaskAttachment
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPlanAttachments function
type GetPlanAttachmentsArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	Type *string
}

// [Preview API]
func (client *ClientImpl) GetRecords(ctx context.Context, args GetRecordsArgs) (*[]TimelineRecord, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()

	queryParams := url.Values{}
	if args.ChangeId != nil {
		queryParams.Add("changeId", strconv.Itoa(*args.ChangeId))
	}
	locationId, _ := uuid.Parse("8893bc5b-35b2-4be7-83cb-99e683551db4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TimelineRecord
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRecords function
type GetRecordsArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
	// (optional)
	ChangeId *int
}

// [Preview API]
func (client *ClientImpl) GetTimeline(ctx context.Context, args GetTimelineArgs) (*Timeline, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()

	queryParams := url.Values{}
	if args.ChangeId != nil {
		queryParams.Add("changeId", strconv.Itoa(*args.ChangeId))
	}
	if args.IncludeRecords != nil {
		queryParams.Add("includeRecords", strconv.FormatBool(*args.IncludeRecords))
	}
	locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Timeline
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTimeline function
type GetTimelineArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
	// (required)
	TimelineId *uuid.UUID
	// (optional)
	ChangeId *int
	// (optional)
	IncludeRecords *bool
}

// [Preview API]
func (client *ClientImpl) GetTimelines(ctx context.Context, args GetTimelinesArgs) (*[]Timeline, error) {
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()

	locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Timeline
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTimelines function
type GetTimelinesArgs struct {
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required)
	PlanId *uuid.UUID
}

// [Preview API] Update timeline records if they already exist, otherwise create new ones for the same timeline.
func (client *ClientImpl) UpdateRecords(ctx context.Context, args UpdateRecordsArgs) (*[]TimelineRecord, error) {
	if args.Records == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Records"}
	}
	routeValues := make(map[string]string)
	if args.ScopeIdentifier == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ScopeIdentifier"}
	}
	routeValues["scopeIdentifier"] = (*args.ScopeIdentifier).String()
	if args.HubName == nil || *args.HubName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.HubName"}
	}
	routeValues["hubName"] = *args.HubName
	if args.PlanId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PlanId"}
	}
	routeValues["planId"] = (*args.PlanId).String()
	if args.TimelineId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TimelineId"}
	}
	routeValues["timelineId"] = (*args.TimelineId).String()

	body, marshalErr := json.Marshal(*args.Records)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("8893bc5b-35b2-4be7-83cb-99e683551db4")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []TimelineRecord
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateRecords function
type UpdateRecordsArgs struct {
	// (required) The array of timeline records to be updated.
	Records *azuredevops.VssJsonCollectionWrapper
	// (required) The project GUID to scope the request
	ScopeIdentifier *uuid.UUID
	// (required) The name of the server hub. Common examples: "build", "rm", "checks"
	HubName *string
	// (required) The ID of the plan.
	PlanId *uuid.UUID
	// (required) The ID of the timeline.
	TimelineId *uuid.UUID
}
