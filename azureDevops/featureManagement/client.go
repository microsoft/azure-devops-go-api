// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package featureManagement

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
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

// [Preview API] Get a specific feature by its id
// ctx
// featureId (required): The contribution id of the feature
func (client Client) GetFeature(ctx context.Context, featureId *string) (*ContributedFeature, error) {
    routeValues := make(map[string]string)
    if featureId == nil || *featureId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "featureId"} 
    }
    routeValues["featureId"] = *featureId

    locationId, _ := uuid.Parse("c4209f25-7a27-41dd-9f04-06080c7b6afd")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeature
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of all defined features
// ctx
// targetContributionId (optional): Optional target contribution. If null/empty, return all features. If specified include the features that target the specified contribution.
func (client Client) GetFeatures(ctx context.Context, targetContributionId *string) (*[]ContributedFeature, error) {
    queryParams := url.Values{}
    if targetContributionId != nil {
        queryParams.Add("targetContributionId", *targetContributionId)
    }
    locationId, _ := uuid.Parse("c4209f25-7a27-41dd-9f04-06080c7b6afd")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ContributedFeature
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the state of the specified feature for the given user/all-users scope
// ctx
// featureId (required): Contribution id of the feature
// userScope (required): User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
func (client Client) GetFeatureState(ctx context.Context, featureId *string, userScope *string) (*ContributedFeatureState, error) {
    routeValues := make(map[string]string)
    if featureId == nil || *featureId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "featureId"} 
    }
    routeValues["featureId"] = *featureId
    if userScope == nil || *userScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope

    locationId, _ := uuid.Parse("98911314-3f9b-4eaf-80e8-83900d8e85d9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeatureState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Set the state of a feature
// ctx
// feature (required): Posted feature state object. Should specify the effective value.
// featureId (required): Contribution id of the feature
// userScope (required): User-Scope at which to set the value. Should be "me" for the current user or "host" for all users.
// reason (optional): Reason for changing the state
// reasonCode (optional): Short reason code
func (client Client) SetFeatureState(ctx context.Context, feature *ContributedFeatureState, featureId *string, userScope *string, reason *string, reasonCode *string) (*ContributedFeatureState, error) {
    if feature == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "feature"}
    }
    routeValues := make(map[string]string)
    if featureId == nil || *featureId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "featureId"} 
    }
    routeValues["featureId"] = *featureId
    if userScope == nil || *userScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope

    queryParams := url.Values{}
    if reason != nil {
        queryParams.Add("reason", *reason)
    }
    if reasonCode != nil {
        queryParams.Add("reasonCode", *reasonCode)
    }
    body, marshalErr := json.Marshal(*feature)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("98911314-3f9b-4eaf-80e8-83900d8e85d9")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeatureState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the state of the specified feature for the given named scope
// ctx
// featureId (required): Contribution id of the feature
// userScope (required): User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
// scopeName (required): Scope at which to get the feature setting for (e.g. "project" or "team")
// scopeValue (required): Value of the scope (e.g. the project or team id)
func (client Client) GetFeatureStateForScope(ctx context.Context, featureId *string, userScope *string, scopeName *string, scopeValue *string) (*ContributedFeatureState, error) {
    routeValues := make(map[string]string)
    if featureId == nil || *featureId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "featureId"} 
    }
    routeValues["featureId"] = *featureId
    if userScope == nil || *userScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope
    if scopeName == nil || *scopeName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "scopeName"} 
    }
    routeValues["scopeName"] = *scopeName
    if scopeValue == nil || *scopeValue == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "scopeValue"} 
    }
    routeValues["scopeValue"] = *scopeValue

    locationId, _ := uuid.Parse("dd291e43-aa9f-4cee-8465-a93c78e414a4")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeatureState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Set the state of a feature at a specific scope
// ctx
// feature (required): Posted feature state object. Should specify the effective value.
// featureId (required): Contribution id of the feature
// userScope (required): User-Scope at which to set the value. Should be "me" for the current user or "host" for all users.
// scopeName (required): Scope at which to get the feature setting for (e.g. "project" or "team")
// scopeValue (required): Value of the scope (e.g. the project or team id)
// reason (optional): Reason for changing the state
// reasonCode (optional): Short reason code
func (client Client) SetFeatureStateForScope(ctx context.Context, feature *ContributedFeatureState, featureId *string, userScope *string, scopeName *string, scopeValue *string, reason *string, reasonCode *string) (*ContributedFeatureState, error) {
    if feature == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "feature"}
    }
    routeValues := make(map[string]string)
    if featureId == nil || *featureId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "featureId"} 
    }
    routeValues["featureId"] = *featureId
    if userScope == nil || *userScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope
    if scopeName == nil || *scopeName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "scopeName"} 
    }
    routeValues["scopeName"] = *scopeName
    if scopeValue == nil || *scopeValue == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "scopeValue"} 
    }
    routeValues["scopeValue"] = *scopeValue

    queryParams := url.Values{}
    if reason != nil {
        queryParams.Add("reason", *reason)
    }
    if reasonCode != nil {
        queryParams.Add("reasonCode", *reasonCode)
    }
    body, marshalErr := json.Marshal(*feature)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("dd291e43-aa9f-4cee-8465-a93c78e414a4")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeatureState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the effective state for a list of feature ids
// ctx
// query (required): Features to query along with current scope values
func (client Client) QueryFeatureStates(ctx context.Context, query *ContributedFeatureStateQuery) (*ContributedFeatureStateQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2b4486ad-122b-400c-ae65-17b6672c1f9d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeatureStateQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the states of the specified features for the default scope
// ctx
// query (required): Query describing the features to query.
// userScope (required)
func (client Client) QueryFeatureStatesForDefaultScope(ctx context.Context, query *ContributedFeatureStateQuery, userScope *string) (*ContributedFeatureStateQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope

    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("3f810f28-03e2-4239-b0bc-788add3005e5")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeatureStateQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the states of the specified features for the specific named scope
// ctx
// query (required): Query describing the features to query.
// userScope (required)
// scopeName (required)
// scopeValue (required)
func (client Client) QueryFeatureStatesForNamedScope(ctx context.Context, query *ContributedFeatureStateQuery, userScope *string, scopeName *string, scopeValue *string) (*ContributedFeatureStateQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "userScope"} 
    }
    routeValues["userScope"] = *userScope
    if scopeName == nil || *scopeName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "scopeName"} 
    }
    routeValues["scopeName"] = *scopeName
    if scopeValue == nil || *scopeValue == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "scopeValue"} 
    }
    routeValues["scopeValue"] = *scopeValue

    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("f29e997b-c2da-4d15-8380-765788a1a74c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ContributedFeatureStateQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

