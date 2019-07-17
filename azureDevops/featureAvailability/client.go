// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package featureAvailability

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

// [Preview API] Retrieve a listing of all feature flags and their current states for a user
// userEmail (optional): The email of the user to check
func (client Client) GetAllFeatureFlags(userEmail *string) (*[]FeatureFlag, error) {
    queryParams := url.Values{}
    if userEmail != nil {
        queryParams.Add("userEmail", *userEmail)
    }
    locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []FeatureFlag
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve information on a single feature flag and its current states
// name (required): The name of the feature to retrieve
// checkFeatureExists (optional): Check if feature exists
func (client Client) GetFeatureFlagByName(name *string, checkFeatureExists *bool) (*FeatureFlag, error) {
    routeValues := make(map[string]string)
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    queryParams := url.Values{}
    if checkFeatureExists != nil {
        queryParams.Add("checkFeatureExists", strconv.FormatBool(*checkFeatureExists))
    }
    locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeatureFlag
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve information on a single feature flag and its current states for a user
// name (required): The name of the feature to retrieve
// userEmail (required): The email of the user to check
// checkFeatureExists (optional): Check if feature exists
func (client Client) GetFeatureFlagByNameAndUserEmail(name *string, userEmail *string, checkFeatureExists *bool) (*FeatureFlag, error) {
    routeValues := make(map[string]string)
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    queryParams := url.Values{}
    if userEmail == nil {
        return nil, errors.New("userEmail is a required parameter")
    }
    queryParams.Add("userEmail", *userEmail)
    if checkFeatureExists != nil {
        queryParams.Add("checkFeatureExists", strconv.FormatBool(*checkFeatureExists))
    }
    locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeatureFlag
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Retrieve information on a single feature flag and its current states for a user
// name (required): The name of the feature to retrieve
// userId (required): The id of the user to check
// checkFeatureExists (optional): Check if feature exists
func (client Client) GetFeatureFlagByNameAndUserId(name *string, userId *uuid.UUID, checkFeatureExists *bool) (*FeatureFlag, error) {
    routeValues := make(map[string]string)
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    queryParams := url.Values{}
    if userId == nil {
        return nil, errors.New("userId is a required parameter")
    }
    queryParams.Add("userId", (*userId).String())
    if checkFeatureExists != nil {
        queryParams.Add("checkFeatureExists", strconv.FormatBool(*checkFeatureExists))
    }
    locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeatureFlag
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Change the state of an individual feature flag for a name
// state (required): State that should be set
// name (required): The name of the feature to change
// userEmail (optional)
// checkFeatureExists (optional): Checks if the feature exists before setting the state
// setAtApplicationLevelAlso (optional)
func (client Client) UpdateFeatureFlag(state *FeatureFlagPatch, name *string, userEmail *string, checkFeatureExists *bool, setAtApplicationLevelAlso *bool) (*FeatureFlag, error) {
    if state == nil {
        return nil, errors.New("state is a required parameter")
    }
    routeValues := make(map[string]string)
    if name == nil || *name == "" {
        return nil, errors.New("name is a required parameter")
    }
    routeValues["name"] = *name

    queryParams := url.Values{}
    if userEmail != nil {
        queryParams.Add("userEmail", *userEmail)
    }
    if checkFeatureExists != nil {
        queryParams.Add("checkFeatureExists", strconv.FormatBool(*checkFeatureExists))
    }
    if setAtApplicationLevelAlso != nil {
        queryParams.Add("setAtApplicationLevelAlso", strconv.FormatBool(*setAtApplicationLevelAlso))
    }
    body, marshalErr := json.Marshal(*state)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeatureFlag
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

