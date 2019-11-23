// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package search

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/searchshared"
	"net/http"
)

var ResourceAreaId, _ = uuid.Parse("ea48a0a1-269c-42d8-b8ad-ddc8fcdcf578")

type Client interface {
	// [Preview API] Provides a set of results for the search text.
	FetchCodeSearchResults(context.Context, FetchCodeSearchResultsArgs) (*CodeSearchResponse, error)
	// [Preview API] Provides a set of results for the search text.
	FetchPackageSearchResults(context.Context, FetchPackageSearchResultsArgs) (*searchshared.PackageSearchResponse, error)
	// [Preview API] Provides a set of results for the search text.
	FetchScrollCodeSearchResults(context.Context, FetchScrollCodeSearchResultsArgs) (*CodeSearchResponse, error)
	// [Preview API] Provides a set of results for the search request.
	FetchWikiSearchResults(context.Context, FetchWikiSearchResultsArgs) (*searchshared.WikiSearchResponse, error)
	// [Preview API] Provides a set of results for the search text.
	FetchWorkItemSearchResults(context.Context, FetchWorkItemSearchResultsArgs) (*WorkItemSearchResponse, error)
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

// [Preview API] Provides a set of results for the search text.
func (client *ClientImpl) FetchCodeSearchResults(ctx context.Context, args FetchCodeSearchResultsArgs) (*CodeSearchResponse, error) {
	if args.Request == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Request"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.Request)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e7f29993-5b82-4fca-9386-f5cfe683d524")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CodeSearchResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the FetchCodeSearchResults function
type FetchCodeSearchResultsArgs struct {
	// (required) The Code Search Request.
	Request *CodeSearchRequest
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Provides a set of results for the search text.
func (client *ClientImpl) FetchPackageSearchResults(ctx context.Context, args FetchPackageSearchResultsArgs) (*searchshared.PackageSearchResponse, error) {
	if args.Request == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Request"}
	}
	body, marshalErr := json.Marshal(*args.Request)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("f62ada48-eedc-4c8e-93f0-de870e4ecce0")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseBodyValue searchshared.PackageSearchResponseContent
	err = client.Client.UnmarshalBody(resp, &responseBodyValue)

	var responseValue *searchshared.PackageSearchResponse
	if err == nil {
		responseValue = &searchshared.PackageSearchResponse{
			Content:    &responseBodyValue,
			ActivityId: &[]string{resp.Header.Get("ActivityId")},
		}
	}

	return responseValue, err
}

// Arguments for the FetchPackageSearchResults function
type FetchPackageSearchResultsArgs struct {
	// (required) The Package Search Request.
	Request *searchshared.PackageSearchRequest
}

// [Preview API] Provides a set of results for the search text.
func (client *ClientImpl) FetchScrollCodeSearchResults(ctx context.Context, args FetchScrollCodeSearchResultsArgs) (*CodeSearchResponse, error) {
	if args.Request == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Request"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.Request)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("852dac94-e8f7-45a2-9910-927ae35766a2")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue CodeSearchResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the FetchScrollCodeSearchResults function
type FetchScrollCodeSearchResultsArgs struct {
	// (required) The Code Search Request.
	Request *searchshared.ScrollSearchRequest
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Provides a set of results for the search request.
func (client *ClientImpl) FetchWikiSearchResults(ctx context.Context, args FetchWikiSearchResultsArgs) (*searchshared.WikiSearchResponse, error) {
	if args.Request == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Request"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.Request)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e90e7664-7049-4100-9a86-66b161d81080")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue searchshared.WikiSearchResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the FetchWikiSearchResults function
type FetchWikiSearchResultsArgs struct {
	// (required) The Wiki Search Request.
	Request *searchshared.WikiSearchRequest
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Provides a set of results for the search text.
func (client *ClientImpl) FetchWorkItemSearchResults(ctx context.Context, args FetchWorkItemSearchResultsArgs) (*WorkItemSearchResponse, error) {
	if args.Request == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Request"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.Request)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("73b2c9e2-ff9e-4447-8cda-5f5b21ff7cae")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue WorkItemSearchResponse
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the FetchWorkItemSearchResults function
type FetchWorkItemSearchResultsArgs struct {
	// (required) The Work Item Search Request.
	Request *WorkItemSearchRequest
	// (optional) Project ID or project name
	Project *string
}
