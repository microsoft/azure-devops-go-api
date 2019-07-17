// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package workItemTrackingProcessTemplate

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

var ResourceAreaId, _ = uuid.Parse("5264459e-e5e0-4bd8-b118-0985e68a4ec5")

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API] Returns a behavior for the process.
// processId (required): The ID of the process
// behaviorRefName (required): The reference name of the behavior
func (client Client) GetBehavior(processId *uuid.UUID, behaviorRefName *string) (*AdminBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()

    queryParams := url.Values{}
    if behaviorRefName == nil {
        return nil, errors.New("behaviorRefName is a required parameter")
    }
    queryParams.Add("behaviorRefName", *behaviorRefName)
    locationId, _ := uuid.Parse("90bf9317-3571-487b-bc8c-a523ba0e05d7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AdminBehavior
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns a list of behaviors for the process.
// processId (required): The ID of the process
func (client Client) GetBehaviors(processId *uuid.UUID) (*[]AdminBehavior, error) {
    routeValues := make(map[string]string)
    if processId == nil {
        return nil, errors.New("processId is a required parameter")
    }
    routeValues["processId"] = (*processId).String()

    locationId, _ := uuid.Parse("90bf9317-3571-487b-bc8c-a523ba0e05d7")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []AdminBehavior
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Returns requested process template.
// id (required): The ID of the process
func (client Client) ExportProcessTemplate(id *uuid.UUID) (*interface{}, error) {
    routeValues := make(map[string]string)
    if id == nil {
        return nil, errors.New("id is a required parameter")
    }
    routeValues["id"] = (*id).String()
    routeValues["action"] = "Export"

    locationId, _ := uuid.Parse("29e1f38d-9e9c-4358-86a5-cdf9896a5759")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Imports a process from zip file.
// uploadStream (required): Stream to upload
// ignoreWarnings (optional): Ignores validation warnings. Default value is false.
// replaceExistingTemplate (optional): Replaces the existing template. Default value is true.
func (client Client) ImportProcessTemplate(uploadStream *interface{}, ignoreWarnings *bool, replaceExistingTemplate *bool) (*ProcessImportResult, error) {
    if uploadStream == nil {
        return nil, errors.New("uploadStream is a required parameter")
    }
    routeValues := make(map[string]string)
    routeValues["action"] = "Import"

    queryParams := url.Values{}
    if ignoreWarnings != nil {
        queryParams.Add("ignoreWarnings", strconv.FormatBool(*ignoreWarnings))
    }
    if replaceExistingTemplate != nil {
        queryParams.Add("replaceExistingTemplate", strconv.FormatBool(*replaceExistingTemplate))
    }
    body, marshalErr := json.Marshal(*uploadStream)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("29e1f38d-9e9c-4358-86a5-cdf9896a5759")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/octet-stream", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessImportResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Tells whether promote has completed for the specified promote job ID.
// id (required): The ID of the promote job operation
func (client Client) ImportProcessTemplateStatus(id *uuid.UUID) (*ProcessPromoteStatus, error) {
    routeValues := make(map[string]string)
    if id == nil {
        return nil, errors.New("id is a required parameter")
    }
    routeValues["id"] = (*id).String()
    routeValues["action"] = "Status"

    locationId, _ := uuid.Parse("29e1f38d-9e9c-4358-86a5-cdf9896a5759")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ProcessPromoteStatus
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

