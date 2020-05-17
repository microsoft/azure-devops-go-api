// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package featureavailability

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	// [Preview API] Retrieve a listing of all feature flags and their current states for a user
	GetAllFeatureFlags(context.Context, GetAllFeatureFlagsArgs) (*[]FeatureFlag, error)
	// [Preview API] Retrieve information on a single feature flag and its current states
	GetFeatureFlagByName(context.Context, GetFeatureFlagByNameArgs) (*FeatureFlag, error)
	// [Preview API] Retrieve information on a single feature flag and its current states for a user
	GetFeatureFlagByNameAndUserEmail(context.Context, GetFeatureFlagByNameAndUserEmailArgs) (*FeatureFlag, error)
	// [Preview API] Retrieve information on a single feature flag and its current states for a user
	GetFeatureFlagByNameAndUserId(context.Context, GetFeatureFlagByNameAndUserIdArgs) (*FeatureFlag, error)
	// [Preview API] Change the state of an individual feature flag for a name
	UpdateFeatureFlag(context.Context, UpdateFeatureFlagArgs) (*FeatureFlag, error)
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

// [Preview API] Retrieve a listing of all feature flags and their current states for a user
func (client *ClientImpl) GetAllFeatureFlags(ctx context.Context, args GetAllFeatureFlagsArgs) (*[]FeatureFlag, error) {
	queryParams := url.Values{}
	if args.UserEmail != nil {
		queryParams.Add("userEmail", *args.UserEmail)
	}
	locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []FeatureFlag
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAllFeatureFlags function
type GetAllFeatureFlagsArgs struct {
	// (optional) The email of the user to check
	UserEmail *string
}

// [Preview API] Retrieve information on a single feature flag and its current states
func (client *ClientImpl) GetFeatureFlagByName(ctx context.Context, args GetFeatureFlagByNameArgs) (*FeatureFlag, error) {
	routeValues := make(map[string]string)
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	queryParams := url.Values{}
	if args.CheckFeatureExists != nil {
		queryParams.Add("checkFeatureExists", strconv.FormatBool(*args.CheckFeatureExists))
	}
	locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue FeatureFlag
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetFeatureFlagByName function
type GetFeatureFlagByNameArgs struct {
	// (required) The name of the feature to retrieve
	Name *string
	// (optional) Check if feature exists
	CheckFeatureExists *bool
}

// [Preview API] Retrieve information on a single feature flag and its current states for a user
func (client *ClientImpl) GetFeatureFlagByNameAndUserEmail(ctx context.Context, args GetFeatureFlagByNameAndUserEmailArgs) (*FeatureFlag, error) {
	routeValues := make(map[string]string)
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	queryParams := url.Values{}
	if args.UserEmail == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "userEmail"}
	}
	queryParams.Add("userEmail", *args.UserEmail)
	if args.CheckFeatureExists != nil {
		queryParams.Add("checkFeatureExists", strconv.FormatBool(*args.CheckFeatureExists))
	}
	locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue FeatureFlag
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetFeatureFlagByNameAndUserEmail function
type GetFeatureFlagByNameAndUserEmailArgs struct {
	// (required) The name of the feature to retrieve
	Name *string
	// (required) The email of the user to check
	UserEmail *string
	// (optional) Check if feature exists
	CheckFeatureExists *bool
}

// [Preview API] Retrieve information on a single feature flag and its current states for a user
func (client *ClientImpl) GetFeatureFlagByNameAndUserId(ctx context.Context, args GetFeatureFlagByNameAndUserIdArgs) (*FeatureFlag, error) {
	routeValues := make(map[string]string)
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	queryParams := url.Values{}
	if args.UserId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "userId"}
	}
	queryParams.Add("userId", (*args.UserId).String())
	if args.CheckFeatureExists != nil {
		queryParams.Add("checkFeatureExists", strconv.FormatBool(*args.CheckFeatureExists))
	}
	locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue FeatureFlag
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetFeatureFlagByNameAndUserId function
type GetFeatureFlagByNameAndUserIdArgs struct {
	// (required) The name of the feature to retrieve
	Name *string
	// (required) The id of the user to check
	UserId *uuid.UUID
	// (optional) Check if feature exists
	CheckFeatureExists *bool
}

// [Preview API] Change the state of an individual feature flag for a name
func (client *ClientImpl) UpdateFeatureFlag(ctx context.Context, args UpdateFeatureFlagArgs) (*FeatureFlag, error) {
	if args.State == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.State"}
	}
	routeValues := make(map[string]string)
	if args.Name == nil || *args.Name == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Name"}
	}
	routeValues["name"] = *args.Name

	queryParams := url.Values{}
	if args.UserEmail != nil {
		queryParams.Add("userEmail", *args.UserEmail)
	}
	if args.CheckFeatureExists != nil {
		queryParams.Add("checkFeatureExists", strconv.FormatBool(*args.CheckFeatureExists))
	}
	if args.SetAtApplicationLevelAlso != nil {
		queryParams.Add("setAtApplicationLevelAlso", strconv.FormatBool(*args.SetAtApplicationLevelAlso))
	}
	body, marshalErr := json.Marshal(*args.State)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("3e2b80f8-9e6f-441e-8393-005610692d9c")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue FeatureFlag
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateFeatureFlag function
type UpdateFeatureFlagArgs struct {
	// (required) State that should be set
	State *FeatureFlagPatch
	// (required) The name of the feature to change
	Name *string
	// (optional)
	UserEmail *string
	// (optional) Checks if the feature exists before setting the state
	CheckFeatureExists *bool
	// (optional)
	SetAtApplicationLevelAlso *bool
}
