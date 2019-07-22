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
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("d397749b-f115-4027-b6dd-77a65dd10d21")

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API]
// ctx
// metadata (required)
// feedId (required)
// packageName (required)
// packageVersion (required)
func (client Client) AddPackage(ctx context.Context, metadata *UPackPackagePushMetadata, feedId *string, packageName *string, packageVersion *string) error {
    if metadata == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "metadata"}
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*metadata)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    _, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API]
// ctx
// feedId (required)
// packageName (required)
// packageVersion (required)
// intent (optional)
func (client Client) GetPackageMetadata(ctx context.Context, feedId *string, packageName *string, packageVersion *string, intent *UPackGetPackageMetadataIntent) (*UPackPackageMetadata, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    queryParams := url.Values{}
    if intent != nil {
        queryParams.Add("intent", string(*intent))
    }
    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackPackageMetadata
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// feedId (required)
// packageName (required)
func (client Client) GetPackageVersionsMetadata(ctx context.Context, feedId *string, packageName *string) (*UPackLimitedPackageMetadataListResponse, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName

    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackLimitedPackageMetadataListResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

