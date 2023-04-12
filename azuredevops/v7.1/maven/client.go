// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package maven

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

var ResourceAreaId, _ = uuid.Parse("6f7f8c07-ff36-473c-bcf3-bd6cc9b6c066")

type Client interface {
	// [Preview API] Permanently delete a package from a feed's recycle bin.
	DeletePackageVersionFromRecycleBin(context.Context, DeletePackageVersionFromRecycleBinArgs) error
	// [Preview API] Fulfills Maven package file download requests by either returning the URL of the requested package file or, in the case of Azure DevOps Server (OnPrem), returning the content as a stream.
	DownloadPackage(context.Context, DownloadPackageArgs) (io.ReadCloser, error)
	// [Preview API] Get information about a package version.
	GetPackageVersion(context.Context, GetPackageVersionArgs) (*Package, error)
	// [Preview API] Get information about a package version in the recycle bin.
	GetPackageVersionMetadataFromRecycleBin(context.Context, GetPackageVersionMetadataFromRecycleBinArgs) (*MavenPackageVersionDeletionState, error)
	// [Preview API] Get the upstreaming behavior of a package within the context of a feed
	GetUpstreamingBehavior(context.Context, GetUpstreamingBehaviorArgs) (*packagingshared.UpstreamingBehavior, error)
	// [Preview API] Delete a package version from the feed and move it to the feed's recycle bin.
	PackageDelete(context.Context, PackageDeleteArgs) error
	// [Preview API] Restore a package version from the recycle bin to its associated feed.
	RestorePackageVersionFromRecycleBin(context.Context, RestorePackageVersionFromRecycleBinArgs) error
	// [Preview API] Set the upstreaming behavior of a package within the context of a feed
	SetUpstreamingBehavior(context.Context, SetUpstreamingBehaviorArgs) error
	// [Preview API] Update state for a package version.
	UpdatePackageVersion(context.Context, UpdatePackageVersionArgs) error
	// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
	UpdatePackageVersions(context.Context, UpdatePackageVersionsArgs) error
	// [Preview API] Delete or restore several package versions from the recycle bin.
	UpdateRecycleBinPackages(context.Context, UpdateRecycleBinPackagesArgs) error
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

// [Preview API] Permanently delete a package from a feed's recycle bin.
func (client *ClientImpl) DeletePackageVersionFromRecycleBin(ctx context.Context, args DeletePackageVersionFromRecycleBinArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.Feed == nil || *args.Feed == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId
	if args.Version == nil || *args.Version == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Version"}
	}
	routeValues["version"] = *args.Version

	locationId, _ := uuid.Parse("f67e10eb-1254-4953-add7-d49b83a16c9f")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePackageVersionFromRecycleBin function
type DeletePackageVersionFromRecycleBinArgs struct {
	// (required) Name or ID of the feed.
	Feed *string
	// (required) Group ID of the package.
	GroupId *string
	// (required) Artifact ID of the package.
	ArtifactId *string
	// (required) Version of the package.
	Version *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Fulfills Maven package file download requests by either returning the URL of the requested package file or, in the case of Azure DevOps Server (OnPrem), returning the content as a stream.
func (client *ClientImpl) DownloadPackage(ctx context.Context, args DownloadPackageArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId
	if args.Version == nil || *args.Version == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Version"}
	}
	routeValues["version"] = *args.Version
	if args.FileName == nil || *args.FileName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FileName"}
	}
	routeValues["fileName"] = *args.FileName

	locationId, _ := uuid.Parse("c338d4b5-d30a-47e2-95b7-f157ef558833")
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
	// (required) GroupId of the maven package
	GroupId *string
	// (required) ArtifactId of the maven package
	ArtifactId *string
	// (required) Version of the package
	Version *string
	// (required) File name to download
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
	if args.Feed == nil || *args.Feed == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId
	if args.Version == nil || *args.Version == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Version"}
	}
	routeValues["version"] = *args.Version

	queryParams := url.Values{}
	if args.ShowDeleted != nil {
		queryParams.Add("showDeleted", strconv.FormatBool(*args.ShowDeleted))
	}
	locationId, _ := uuid.Parse("180ed967-377a-4112-986b-607adb14ded4")
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
	Feed *string
	// (required) Group ID of the package.
	GroupId *string
	// (required) Artifact ID of the package.
	ArtifactId *string
	// (required) Version of the package.
	Version *string
	// (optional) Project ID or project name
	Project *string
	// (optional) True to show information for deleted packages.
	ShowDeleted *bool
}

// [Preview API] Get information about a package version in the recycle bin.
func (client *ClientImpl) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetPackageVersionMetadataFromRecycleBinArgs) (*MavenPackageVersionDeletionState, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.Feed == nil || *args.Feed == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId
	if args.Version == nil || *args.Version == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Version"}
	}
	routeValues["version"] = *args.Version

	locationId, _ := uuid.Parse("f67e10eb-1254-4953-add7-d49b83a16c9f")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue MavenPackageVersionDeletionState
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPackageVersionMetadataFromRecycleBin function
type GetPackageVersionMetadataFromRecycleBinArgs struct {
	// (required) Name or ID of the feed.
	Feed *string
	// (required) Group ID of the package.
	GroupId *string
	// (required) Artifact ID of the package.
	ArtifactId *string
	// (required) Version of the package.
	Version *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get the upstreaming behavior of a package within the context of a feed
func (client *ClientImpl) GetUpstreamingBehavior(ctx context.Context, args GetUpstreamingBehaviorArgs) (*packagingshared.UpstreamingBehavior, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.Feed == nil || *args.Feed == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId

	locationId, _ := uuid.Parse("fba7ba8c-d1f5-4aeb-8f5d-f017a7d5e719")
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
	Feed *string
	// (required) The group id of the package
	GroupId *string
	// (required) The artifact id of the package
	ArtifactId *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete a package version from the feed and move it to the feed's recycle bin.
func (client *ClientImpl) PackageDelete(ctx context.Context, args PackageDeleteArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.Feed == nil || *args.Feed == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId
	if args.Version == nil || *args.Version == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Version"}
	}
	routeValues["version"] = *args.Version

	locationId, _ := uuid.Parse("180ed967-377a-4112-986b-607adb14ded4")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the PackageDelete function
type PackageDeleteArgs struct {
	// (required) Name or ID of the feed.
	Feed *string
	// (required) Group ID of the package.
	GroupId *string
	// (required) Artifact ID of the package.
	ArtifactId *string
	// (required) Version of the package.
	Version *string
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
	if args.Feed == nil || *args.Feed == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId
	if args.Version == nil || *args.Version == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Version"}
	}
	routeValues["version"] = *args.Version

	body, marshalErr := json.Marshal(*args.PackageVersionDetails)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("f67e10eb-1254-4953-add7-d49b83a16c9f")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RestorePackageVersionFromRecycleBin function
type RestorePackageVersionFromRecycleBinArgs struct {
	// (required) Set the 'Deleted' property to false to restore the package.
	PackageVersionDetails *MavenRecycleBinPackageVersionDetails
	// (required) Name or ID of the feed.
	Feed *string
	// (required) Group ID of the package.
	GroupId *string
	// (required) Artifact ID of the package.
	ArtifactId *string
	// (required) Version of the package.
	Version *string
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
	if args.Feed == nil || *args.Feed == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId

	body, marshalErr := json.Marshal(*args.Behavior)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("fba7ba8c-d1f5-4aeb-8f5d-f017a7d5e719")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the SetUpstreamingBehavior function
type SetUpstreamingBehaviorArgs struct {
	// (required) The name or id of the feed
	Feed *string
	// (required)
	GroupId *string
	// (required)
	ArtifactId *string
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
	if args.Feed == nil || *args.Feed == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed
	if args.GroupId == nil || *args.GroupId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.GroupId"}
	}
	routeValues["groupId"] = *args.GroupId
	if args.ArtifactId == nil || *args.ArtifactId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ArtifactId"}
	}
	routeValues["artifactId"] = *args.ArtifactId
	if args.Version == nil || *args.Version == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Version"}
	}
	routeValues["version"] = *args.Version

	body, marshalErr := json.Marshal(*args.PackageVersionDetails)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("180ed967-377a-4112-986b-607adb14ded4")
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
	Feed *string
	// (required) Group ID of the package.
	GroupId *string
	// (required) Artifact ID of the package.
	ArtifactId *string
	// (required) Version of the package.
	Version *string
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
	locationId, _ := uuid.Parse("b7c586b0-d947-4d35-811a-f1161de80e6c")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePackageVersions function
type UpdatePackageVersionsArgs struct {
	// (required) Information about the packages to update, the operation to perform, and its associated data.
	BatchRequest *MavenPackagesBatchRequest
	// (required) Feed which contains the packages to update.
	FeedId *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete or restore several package versions from the recycle bin.
func (client *ClientImpl) UpdateRecycleBinPackages(ctx context.Context, args UpdateRecycleBinPackagesArgs) error {
	if args.BatchRequest == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.BatchRequest"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.Feed == nil || *args.Feed == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Feed"}
	}
	routeValues["feed"] = *args.Feed

	body, marshalErr := json.Marshal(*args.BatchRequest)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("5dd6f547-c76f-4d9d-b2ec-4720feda641f")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdateRecycleBinPackages function
type UpdateRecycleBinPackagesArgs struct {
	// (required) Information about the packages to update, the operation to perform, and its associated data.
	BatchRequest *MavenPackagesBatchRequest
	// (required)
	Feed *string
	// (optional) Project ID or project name
	Project *string
}
