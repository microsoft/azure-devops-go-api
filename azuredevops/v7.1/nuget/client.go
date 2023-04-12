// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package nuget

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

var ResourceAreaId, _ = uuid.Parse("b3be7473-68ea-4a81-bfc7-9530baaa19ad")

type Client interface {
	// [Preview API] Send a package version from the feed to its paired recycle bin.
	DeletePackageVersion(context.Context, DeletePackageVersionArgs) (*Package, error)
	// [Preview API] Delete a package version from a feed's recycle bin.
	DeletePackageVersionFromRecycleBin(context.Context, DeletePackageVersionFromRecycleBinArgs) error
	// [Preview API] Download a package version directly.
	DownloadPackage(context.Context, DownloadPackageArgs) (io.ReadCloser, error)
	// [Preview API] Get information about a package version.
	GetPackageVersion(context.Context, GetPackageVersionArgs) (*Package, error)
	// [Preview API] View a package version's deletion/recycled status
	GetPackageVersionMetadataFromRecycleBin(context.Context, GetPackageVersionMetadataFromRecycleBinArgs) (*NuGetPackageVersionDeletionState, error)
	// [Preview API] Get the upstreaming behavior of a package within the context of a feed
	GetUpstreamingBehavior(context.Context, GetUpstreamingBehaviorArgs) (*packagingshared.UpstreamingBehavior, error)
	// [Preview API] Restore a package version from a feed's recycle bin back into the active feed.
	RestorePackageVersionFromRecycleBin(context.Context, RestorePackageVersionFromRecycleBinArgs) error
	// [Preview API] Set the upstreaming behavior of a package within the context of a feed
	SetUpstreamingBehavior(context.Context, SetUpstreamingBehaviorArgs) error
	// [Preview API] Set mutable state on a package version.
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

// [Preview API] Send a package version from the feed to its paired recycle bin.
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

	locationId, _ := uuid.Parse("36c9353b-e250-4c57-b040-513c186c3905")
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
	// (required) Name of the package to delete.
	PackageName *string
	// (required) Version of the package to delete.
	PackageVersion *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete a package version from a feed's recycle bin.
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

	locationId, _ := uuid.Parse("07e88775-e3cb-4408-bbe1-628e036fac8c")
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

// [Preview API] Download a package version directly.
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

	queryParams := url.Values{}
	if args.SourceProtocolVersion != nil {
		queryParams.Add("sourceProtocolVersion", *args.SourceProtocolVersion)
	}
	locationId, _ := uuid.Parse("6ea81b8c-7386-490b-a71f-6cf23c80b388")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
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
	// (optional) Project ID or project name
	Project *string
	// (optional) Unused
	SourceProtocolVersion *string
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
	locationId, _ := uuid.Parse("36c9353b-e250-4c57-b040-513c186c3905")
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
	// (optional) True to include deleted packages in the response.
	ShowDeleted *bool
}

// [Preview API] View a package version's deletion/recycled status
func (client *ClientImpl) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetPackageVersionMetadataFromRecycleBinArgs) (*NuGetPackageVersionDeletionState, error) {
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

	locationId, _ := uuid.Parse("07e88775-e3cb-4408-bbe1-628e036fac8c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue NuGetPackageVersionDeletionState
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

	locationId, _ := uuid.Parse("b41eec47-6472-4efa-bcd5-a2c5607b66ec")
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

// [Preview API] Restore a package version from a feed's recycle bin back into the active feed.
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
	locationId, _ := uuid.Parse("07e88775-e3cb-4408-bbe1-628e036fac8c")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RestorePackageVersionFromRecycleBin function
type RestorePackageVersionFromRecycleBinArgs struct {
	// (required) Set the 'Deleted' member to 'false' to apply the restore operation
	PackageVersionDetails *NuGetRecycleBinPackageVersionDetails
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
	locationId, _ := uuid.Parse("b41eec47-6472-4efa-bcd5-a2c5607b66ec")
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

// [Preview API] Set mutable state on a package version.
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
	locationId, _ := uuid.Parse("36c9353b-e250-4c57-b040-513c186c3905")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePackageVersion function
type UpdatePackageVersionArgs struct {
	// (required) New state to apply to the referenced package.
	PackageVersionDetails *PackageVersionDetails
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package to update.
	PackageName *string
	// (required) Version of the package to update.
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
	locationId, _ := uuid.Parse("00c58ea7-d55f-49de-b59f-983533ae11dc")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePackageVersions function
type UpdatePackageVersionsArgs struct {
	// (required) Information about the packages to update, the operation to perform, and its associated data.
	BatchRequest *NuGetPackagesBatchRequest
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
	locationId, _ := uuid.Parse("6479ac16-32f4-40f7-aa96-9414de861352")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdateRecycleBinPackageVersions function
type UpdateRecycleBinPackageVersionsArgs struct {
	// (required) Information about the packages to update, the operation to perform, and its associated data. <c>Operation</c> must be <c>PermanentDelete</c> or <c>RestoreToFeed</c>
	BatchRequest *NuGetPackagesBatchRequest
	// (required) Name or ID of the feed.
	FeedId *string
	// (optional) Project ID or project name
	Project *string
}
