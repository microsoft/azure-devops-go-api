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
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var ResourceAreaId, _ = uuid.Parse("8477aec9-a4c7-4bd4-a456-ba4c53c989cb")

type Client interface {
	// [Preview API]
	GetInstalledExtensionByName(context.Context, GetInstalledExtensionByNameArgs) (*InstalledExtension, error)
	// [Preview API]
	GetInstalledExtensions(context.Context, GetInstalledExtensionsArgs) (*[]InstalledExtension, error)
	// [Preview API] Query for contribution nodes and provider details according the parameters in the passed in query object.
	QueryContributionNodes(context.Context, QueryContributionNodesArgs) (*ContributionNodeQueryResult, error)
	// [Preview API]
	QueryDataProviders(context.Context, QueryDataProvidersArgs) (*DataProviderResult, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// [Preview API]
func (client *ClientImpl) GetInstalledExtensionByName(ctx context.Context, args GetInstalledExtensionByNameArgs) (*InstalledExtension, error) {
	routeValues := make(map[string]string)
	if args.PublisherName == nil || *args.PublisherName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PublisherName"}
	}
	routeValues["publisherName"] = *args.PublisherName
	if args.ExtensionName == nil || *args.ExtensionName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ExtensionName"}
	}
	routeValues["extensionName"] = *args.ExtensionName

	queryParams := url.Values{}
	if args.AssetTypes != nil {
		listAsString := strings.Join((*args.AssetTypes)[:], ":")
		queryParams.Add("assetTypes", listAsString)
	}
	locationId, _ := uuid.Parse("3e2f6668-0798-4dcb-b592-bfe2fa57fde2")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue InstalledExtension
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetInstalledExtensionByName function
type GetInstalledExtensionByNameArgs struct {
	// (required)
	PublisherName *string
	// (required)
	ExtensionName *string
	// (optional)
	AssetTypes *[]string
}

// [Preview API]
func (client *ClientImpl) GetInstalledExtensions(ctx context.Context, args GetInstalledExtensionsArgs) (*[]InstalledExtension, error) {
	queryParams := url.Values{}
	if args.ContributionIds != nil {
		listAsString := strings.Join((*args.ContributionIds)[:], ";")
		queryParams.Add("contributionIds", listAsString)
	}
	if args.IncludeDisabledApps != nil {
		queryParams.Add("includeDisabledApps", strconv.FormatBool(*args.IncludeDisabledApps))
	}
	if args.AssetTypes != nil {
		listAsString := strings.Join((*args.AssetTypes)[:], ":")
		queryParams.Add("assetTypes", listAsString)
	}
	locationId, _ := uuid.Parse("2648442b-fd63-4b9a-902f-0c913510f139")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []InstalledExtension
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetInstalledExtensions function
type GetInstalledExtensionsArgs struct {
	// (optional)
	ContributionIds *[]string
	// (optional)
	IncludeDisabledApps *bool
	// (optional)
	AssetTypes *[]string
}

// [Preview API] Query for contribution nodes and provider details according the parameters in the passed in query object.
func (client *ClientImpl) QueryContributionNodes(ctx context.Context, args QueryContributionNodesArgs) (*ContributionNodeQueryResult, error) {
	if args.Query == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Query"}
	}
	body, marshalErr := json.Marshal(*args.Query)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("db7f2146-2309-4cee-b39c-c767777a1c55")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue ContributionNodeQueryResult
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the QueryContributionNodes function
type QueryContributionNodesArgs struct {
	// (required)
	Query *ContributionNodeQuery
}

// [Preview API]
func (client *ClientImpl) QueryDataProviders(ctx context.Context, args QueryDataProvidersArgs) (*DataProviderResult, error) {
	if args.Query == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Query"}
	}
	routeValues := make(map[string]string)
	if args.ScopeName != nil && *args.ScopeName != "" {
		routeValues["scopeName"] = *args.ScopeName
	}
	if args.ScopeValue != nil && *args.ScopeValue != "" {
		routeValues["scopeValue"] = *args.ScopeValue
	}

	body, marshalErr := json.Marshal(*args.Query)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("738368db-35ee-4b85-9f94-77ed34af2b0d")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue DataProviderResult
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the QueryDataProviders function
type QueryDataProvidersArgs struct {
	// (required)
	Query *DataProviderQuery
	// (optional)
	ScopeName *string
	// (optional)
	ScopeValue *string
}
