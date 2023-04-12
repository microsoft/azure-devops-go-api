// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package pypiapi

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7.1"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7.1/packagingshared"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("92f0314b-06c5-46e0-abe7-15fd9d13276a")

type Client interface {
	// [Preview API] Delete a package version, moving it to the recycle bin.
	DeletePackageVersion(context.Context, DeletePackageVersionArgs) (*Package, error)
	// [Preview API] Delete a package version from the feed, moving it to the recycle bin.
	DeletePackageVersionFromRecycleBin(context.Context, DeletePackageVersionFromRecycleBinArgs) error
	// [Preview API] Download a python package file directly. This API is intended for manual UI download options, not for programmatic access and scripting.
	DownloadPackage(context.Context, DownloadPackageArgs) (io.ReadCloser, error)
	// [Preview API] Get information about a package version.
	GetPackageVersion(context.Context, GetPackageVersionArgs) (*Package, error)
	// [Preview API] Get information about a package version in the recycle bin.
	GetPackageVersionMetadataFromRecycleBin(context.Context, GetPackageVersionMetadataFromRecycleBinArgs) (*PyPiPackageVersionDeletionState, error)
	// [Preview API] Get the upstreaming behavior of a package within the context of a feed
	GetUpstreamingBehavior(context.Context, GetUpstreamingBehaviorArgs) (*packagingshared.UpstreamingBehavior, error)
	// [Preview API] Restore a package version from the recycle bin to its associated feed.
	RestorePackageVersionFromRecycleBin(context.Context, RestorePackageVersionFromRecycleBinArgs) error
	// [Preview API] Set the upstreaming behavior of a package within the context of a feed
	SetUpstreamingBehavior(context.Context, SetUpstreamingBehaviorArgs) error
	// [Preview API] Update state for a package version.
	UpdatePackageVersion(context.Context, UpdatePackageVersionArgs) error
	// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
	UpdatePackageVersions(context.Context, UpdatePackageVersionsArgs) error
	// [Preview API] Delete or restore several package versions from the recycle bin.
	UpdateRecycleBinPackageVersions(context.Context, UpdateRecycleBinPackageVersionsArgs) error
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

// [Preview API] Delete a package version, moving it to the recycle bin.
func (client *ClientImpl) DeletePackageVersion(ctx context.Context, args DeletePackageVersionArgs) (*Package, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("d146ac7e-9e3f-4448-b956-f9bb3bdf9b2e")
	resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the DeletePackageVersion function
type DeletePackageVersionArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete a package version from the feed, moving it to the recycle bin.
func (client *ClientImpl) DeletePackageVersionFromRecycleBin(ctx context.Context, args DeletePackageVersionFromRecycleBinArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("07143752-3d94-45fd-86c2-0c77ed87847b")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePackageVersionFromRecycleBin function
type DeletePackageVersionFromRecycleBinArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Download a python package file directly. This API is intended for manual UI download options, not for programmatic access and scripting.
func (client *ClientImpl) DownloadPackage(ctx context.Context, args DownloadPackageArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion
	if args.FileName == nil || *args.FileName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FileName"}
	}
	routeValues["fileName"] = *args.FileName

	locationId, _ := uuid.Parse("97218bae-a64d-4381-9257-b5b7951f0b98")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the DownloadPackage function
type DownloadPackageArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
	// (required) Name of the file in the package
	FileName *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get information about a package version.
func (client *ClientImpl) GetPackageVersion(ctx context.Context, args GetPackageVersionArgs) (*Package, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	queryParams := url.Values{}
	if args.ShowDeleted != nil {
		queryParams.Add("showDeleted", strconv.FormatBool(*args.ShowDeleted))
	}
	locationId, _ := uuid.Parse("d146ac7e-9e3f-4448-b956-f9bb3bdf9b2e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPackageVersion function
type GetPackageVersionArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
	// (optional) Project ID or project name
	Project *string
	// (optional) True to show information for deleted package versions.
	ShowDeleted *bool
}

// [Preview API] Get information about a package version in the recycle bin.
func (client *ClientImpl) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetPackageVersionMetadataFromRecycleBinArgs) (*PyPiPackageVersionDeletionState, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("07143752-3d94-45fd-86c2-0c77ed87847b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue PyPiPackageVersionDeletionState
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPackageVersionMetadataFromRecycleBin function
type GetPackageVersionMetadataFromRecycleBinArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get the upstreaming behavior of a package within the context of a feed
func (client *ClientImpl) GetUpstreamingBehavior(ctx context.Context, args GetUpstreamingBehaviorArgs) (*packagingshared.UpstreamingBehavior, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName

	locationId, _ := uuid.Parse("21b8c9a7-7080-45be-a5ba-e50bb4c18130")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue packagingshared.UpstreamingBehavior
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetUpstreamingBehavior function
type GetUpstreamingBehaviorArgs struct {
	// (required) The name or id of the feed
	FeedId *string
	// (required) The name of the package
	PackageName *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Restore a package version from the recycle bin to its associated feed.
func (client *ClientImpl) RestorePackageVersionFromRecycleBin(ctx context.Context, args RestorePackageVersionFromRecycleBinArgs) error {
	if args.PackageVersionDetails == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PackageVersionDetails"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	body, marshalErr := json.Marshal(*args.PackageVersionDetails)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("07143752-3d94-45fd-86c2-0c77ed87847b")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RestorePackageVersionFromRecycleBin function
type RestorePackageVersionFromRecycleBinArgs struct {
	// (required) Set the 'Deleted' state to 'false' to restore the package to its feed.
	PackageVersionDetails *PyPiRecycleBinPackageVersionDetails
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Set the upstreaming behavior of a package within the context of a feed
func (client *ClientImpl) SetUpstreamingBehavior(ctx context.Context, args SetUpstreamingBehaviorArgs) error {
	if args.Behavior == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.Behavior"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName

	body, marshalErr := json.Marshal(*args.Behavior)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("21b8c9a7-7080-45be-a5ba-e50bb4c18130")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the SetUpstreamingBehavior function
type SetUpstreamingBehaviorArgs struct {
	// (required) The name or id of the feed
	FeedId *string
	// (required) The name of the package
	PackageName *string
	// (required) The behavior to apply to the package within the scope of the feed
	Behavior *packagingshared.UpstreamingBehavior
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Update state for a package version.
func (client *ClientImpl) UpdatePackageVersion(ctx context.Context, args UpdatePackageVersionArgs) error {
	if args.PackageVersionDetails == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PackageVersionDetails"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageName == nil || *args.PackageName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageName"}
	}
	routeValues["packageName"] = *args.PackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	body, marshalErr := json.Marshal(*args.PackageVersionDetails)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("d146ac7e-9e3f-4448-b956-f9bb3bdf9b2e")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePackageVersion function
type UpdatePackageVersionArgs struct {
	// (required) Details to be updated.
	PackageVersionDetails *PackageVersionDetails
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
func (client *ClientImpl) UpdatePackageVersions(ctx context.Context, args UpdatePackageVersionsArgs) error {
	if args.BatchRequest == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.BatchRequest"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId

	body, marshalErr := json.Marshal(*args.BatchRequest)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("4e53d561-70c1-4c98-b937-0f22acb27b0b")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePackageVersions function
type UpdatePackageVersionsArgs struct {
	// (required) Information about the packages to update, the operation to perform, and its associated data.
	BatchRequest *PyPiPackagesBatchRequest
	// (required) Name or ID of the feed.
	FeedId *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete or restore several package versions from the recycle bin.
func (client *ClientImpl) UpdateRecycleBinPackageVersions(ctx context.Context, args UpdateRecycleBinPackageVersionsArgs) error {
	if args.BatchRequest == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.BatchRequest"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId

	body, marshalErr := json.Marshal(*args.BatchRequest)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("d2d89918-c69e-4ef4-b357-1c3ccb4d28d2")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdateRecycleBinPackageVersions function
type UpdateRecycleBinPackageVersionsArgs struct {
	// (required) Information about the packages to update, the operation to perform, and its associated data. <c>Operation</c> must be <c>PermanentDelete</c> or <c>RestoreToFeed</c>
	BatchRequest *PyPiPackagesBatchRequest
	// (required) Feed which contains the packages to update.
	FeedId *string
	// (optional) Project ID or project name
	Project *string
}
