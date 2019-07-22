// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package contributions

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
    "strings"
)

var ResourceAreaId, _ = uuid.Parse("8477aec9-a4c7-4bd4-a456-ba4c53c989cb")

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

// [Preview API] Query for contribution nodes and provider details according the parameters in the passed in query object.
// ctx
// query (required)
func (client Client) QueryContributionNodes(ctx context.Context, query *ContributionNodeQuery) (*ContributionNodeQueryResult, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("db7f2146-2309-4cee-b39c-c767777a1c55")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributionNodeQueryResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// query (required)
// scopeName (optional)
// scopeValue (optional)
func (client Client) QueryDataProviders(ctx context.Context, query *DataProviderQuery, scopeName *string, scopeValue *string) (*DataProviderResult, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    routeValues := make(map[string]string)
    if scopeName != nil && *scopeName != "" {
        routeValues["scopeName"] = *scopeName
    }
    if scopeValue != nil && *scopeValue != "" {
        routeValues["scopeValue"] = *scopeValue
    }

    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("738368db-35ee-4b85-9f94-77ed34af2b0d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DataProviderResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// contributionIds (optional)
// includeDisabledApps (optional)
// assetTypes (optional)
func (client Client) GetInstalledExtensions(ctx context.Context, contributionIds *[]string, includeDisabledApps *bool, assetTypes *[]string) (*[]InstalledExtension, error) {
    queryParams := url.Values{}
    if contributionIds != nil {
        listAsString := strings.Join((*contributionIds)[:], ";")
        queryParams.Add("contributionIds", listAsString)
    }
    if includeDisabledApps != nil {
        queryParams.Add("includeDisabledApps", strconv.FormatBool(*includeDisabledApps))
    }
    if assetTypes != nil {
        listAsString := strings.Join((*assetTypes)[:], ":")
        queryParams.Add("assetTypes", listAsString)
    }
    locationId, _ := uuid.Parse("2648442b-fd63-4b9a-902f-0c913510f139")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []InstalledExtension
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// publisherName (required)
// extensionName (required)
// assetTypes (optional)
func (client Client) GetInstalledExtensionByName(ctx context.Context, publisherName *string, extensionName *string, assetTypes *[]string) (*InstalledExtension, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherName"} 
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "extensionName"} 
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if assetTypes != nil {
        listAsString := strings.Join((*assetTypes)[:], ":")
        queryParams.Add("assetTypes", listAsString)
    }
    locationId, _ := uuid.Parse("3e2f6668-0798-4dcb-b592-bfe2fa57fde2")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue InstalledExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

