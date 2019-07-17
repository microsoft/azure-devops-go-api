// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package uPackPackaging

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("d397749b-f115-4027-b6dd-77a65dd10d21")

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

// [Preview API]
// metadata (required)
// feedId (required)
// packageName (required)
// packageVersion (required)
func (client Client) AddPackage(metadata *UPackPackagePushMetadata, feedId *string, packageName *string, packageVersion *string) error {
    if metadata == nil {
        return errors.New("metadata is a required parameter")
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*metadata)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    _, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// feedId (required)
// packageName (required)
// packageVersion (required)
// intent (optional)
func (client Client) GetPackageMetadata(feedId *string, packageName *string, packageVersion *string, intent *string) (*UPackPackageMetadata, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    queryParams := url.Values{}
    if intent != nil {
        queryParams.Add("intent", *intent)
    }
    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackPackageMetadata
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// feedId (required)
// packageName (required)
func (client Client) GetPackageVersionsMetadata(feedId *string, packageName *string) (*UPackLimitedPackageMetadataListResponse, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName

    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackLimitedPackageMetadataListResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

