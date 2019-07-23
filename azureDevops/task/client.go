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
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API]
// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// type_ (required)
func (client Client) GetPlanAttachments(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, type_ *string) (*[]TaskAttachment, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("eb55e5d6-2f30-4295-b5ed-38da50b1fc52")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// uploadStream (required): Stream to upload
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
// name (required)
func (client Client) CreateAttachment(ctx context.Context, uploadStream interface{}, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (*TaskAttachment, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "recordId"} 
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *name

    body, marshalErr := json.Marshal(uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAttachment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
// name (required)
func (client Client) GetAttachment(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (*TaskAttachment, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "recordId"} 
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAttachment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
// name (required)
func (client Client) GetAttachmentContent(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "recordId"} 
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "name"} 
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API]
// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
func (client Client) GetAttachments(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string) (*[]TaskAttachment, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "recordId"} 
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "type_"} 
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// uploadStream (required): Stream to upload
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// logId (required)
func (client Client) AppendLogContent(ctx context.Context, uploadStream interface{}, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, logId *int) (*TaskLog, error) {
    if uploadStream == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "uploadStream"}
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if logId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "logId"} 
    }
    routeValues["logId"] = strconv.Itoa(*logId)

    body, marshalErr := json.Marshal(uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskLog
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// log (required)
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) CreateLog(ctx context.Context, log *TaskLog, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*TaskLog, error) {
    if log == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "log"}
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()

    body, marshalErr := json.Marshal(*log)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskLog
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// logId (required)
// startLine (optional)
// endLine (optional)
func (client Client) GetLog(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, logId *int, startLine *uint64, endLine *uint64) (*[]string, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if logId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "logId"} 
    }
    routeValues["logId"] = strconv.Itoa(*logId)

    queryParams := url.Values{}
    if startLine != nil {
        queryParams.Add("startLine", strconv.FormatUint(*startLine, 10))
    }
    if endLine != nil {
        queryParams.Add("endLine", strconv.FormatUint(*endLine, 10))
    }
    locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) GetLogs(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*[]TaskLog, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()

    locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// changeId (optional)
func (client Client) GetRecords(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, changeId *int) (*[]TimelineRecord, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()

    queryParams := url.Values{}
    if changeId != nil {
        queryParams.Add("changeId", strconv.Itoa(*changeId))
    }
    locationId, _ := uuid.Parse("8893bc5b-35b2-4be7-83cb-99e683551db4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TimelineRecord
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// records (required)
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
func (client Client) UpdateRecords(ctx context.Context, records *azureDevops.VssJsonCollectionWrapper, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID) (*[]TimelineRecord, error) {
    if records == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "records"}
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()

    body, marshalErr := json.Marshal(*records)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("8893bc5b-35b2-4be7-83cb-99e683551db4")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TimelineRecord
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// timeline (required)
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) CreateTimeline(ctx context.Context, timeline *Timeline, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*Timeline, error) {
    if timeline == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timeline"}
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()

    body, marshalErr := json.Marshal(*timeline)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Timeline
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
func (client Client) DeleteTimeline(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()

    locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// changeId (optional)
// includeRecords (optional)
func (client Client) GetTimeline(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, changeId *int, includeRecords *bool) (*Timeline, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "timelineId"} 
    }
    routeValues["timelineId"] = (*timelineId).String()

    queryParams := url.Values{}
    if changeId != nil {
        queryParams.Add("changeId", strconv.Itoa(*changeId))
    }
    if includeRecords != nil {
        queryParams.Add("includeRecords", strconv.FormatBool(*includeRecords))
    }
    locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Timeline
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) GetTimelines(ctx context.Context, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*[]Timeline, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scopeIdentifier"} 
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "hubName"} 
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = (*planId).String()

    locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Timeline
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

