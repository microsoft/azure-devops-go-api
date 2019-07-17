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
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API]
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// type_ (required)
func (client Client) GetPlanAttachments(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, type_ *string) (*[]TaskAttachment, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("eb55e5d6-2f30-4295-b5ed-38da50b1fc52")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// uploadStream (required): Stream to upload
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
// name (required)
func (client Client) CreateAttachment(uploadStream *interface{}, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (*TaskAttachment, error) {
    if uploadStream == nil {
        return nil, errors.New("uploadStream is a required parameter")
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, errors.New("recordId is a required parameter")
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAttachment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
// name (required)
func (client Client) GetAttachment(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (*TaskAttachment, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, errors.New("recordId is a required parameter")
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskAttachment
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
// name (required)
func (client Client) GetAttachmentContent(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string, name *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, errors.New("recordId is a required parameter")
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// recordId (required)
// type_ (required)
func (client Client) GetAttachments(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, recordId *uuid.UUID, type_ *string) (*[]TaskAttachment, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()
    if recordId == nil {
        return nil, errors.New("recordId is a required parameter")
    }
    routeValues["recordId"] = (*recordId).String()
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("7898f959-9cdf-4096-b29e-7f293031629e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// uploadStream (required): Stream to upload
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// logId (required)
func (client Client) AppendLogContent(uploadStream *interface{}, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, logId *int) (*TaskLog, error) {
    if uploadStream == nil {
        return nil, errors.New("uploadStream is a required parameter")
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if logId == nil {
        return nil, errors.New("logId is a required parameter")
    }
    routeValues["logId"] = strconv.Itoa(*logId)

    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskLog
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// log (required)
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) CreateLog(log *TaskLog, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*TaskLog, error) {
    if log == nil {
        return nil, errors.New("log is a required parameter")
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()

    body, marshalErr := json.Marshal(*log)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TaskLog
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// logId (required)
// startLine (optional)
// endLine (optional)
func (client Client) GetLog(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, logId *int, startLine *uint64, endLine *uint64) (*[]string, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if logId == nil {
        return nil, errors.New("logId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []string
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) GetLogs(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*[]TaskLog, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()

    locationId, _ := uuid.Parse("46f5667d-263a-4684-91b1-dff7fdcf64e2")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TaskLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// changeId (optional)
func (client Client) GetRecords(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, changeId *int) (*[]TimelineRecord, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()

    queryParams := url.Values{}
    if changeId != nil {
        queryParams.Add("changeId", strconv.Itoa(*changeId))
    }
    locationId, _ := uuid.Parse("8893bc5b-35b2-4be7-83cb-99e683551db4")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TimelineRecord
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// records (required)
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
func (client Client) UpdateRecords(records *azureDevops.VssJsonCollectionWrapper, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID) (*[]TimelineRecord, error) {
    if records == nil {
        return nil, errors.New("records is a required parameter")
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()

    body, marshalErr := json.Marshal(*records)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("8893bc5b-35b2-4be7-83cb-99e683551db4")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TimelineRecord
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// timeline (required)
// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) CreateTimeline(timeline *Timeline, scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*Timeline, error) {
    if timeline == nil {
        return nil, errors.New("timeline is a required parameter")
    }
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()

    body, marshalErr := json.Marshal(*timeline)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Timeline
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
func (client Client) DeleteTimeline(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return errors.New("timelineId is a required parameter")
    }
    routeValues["timelineId"] = (*timelineId).String()

    locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
// timelineId (required)
// changeId (optional)
// includeRecords (optional)
func (client Client) GetTimeline(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID, timelineId *uuid.UUID, changeId *int, includeRecords *bool) (*Timeline, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()
    if timelineId == nil {
        return nil, errors.New("timelineId is a required parameter")
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
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Timeline
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// scopeIdentifier (required): The project GUID to scope the request
// hubName (required): The name of the server hub: "build" for the Build server or "rm" for the Release Management server
// planId (required)
func (client Client) GetTimelines(scopeIdentifier *uuid.UUID, hubName *string, planId *uuid.UUID) (*[]Timeline, error) {
    routeValues := make(map[string]string)
    if scopeIdentifier == nil {
        return nil, errors.New("scopeIdentifier is a required parameter")
    }
    routeValues["scopeIdentifier"] = (*scopeIdentifier).String()
    if hubName == nil || *hubName == "" {
        return nil, errors.New("hubName is a required parameter")
    }
    routeValues["hubName"] = *hubName
    if planId == nil {
        return nil, errors.New("planId is a required parameter")
    }
    routeValues["planId"] = (*planId).String()

    locationId, _ := uuid.Parse("83597576-cc2c-453c-bea6-2882ae6a1653")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Timeline
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

