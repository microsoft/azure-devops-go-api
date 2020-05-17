// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package npm

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"io"
	"net/http"
)

var ResourceAreaId, _ = uuid.Parse("4c83cfc1-f33a-477e-a789-29d38ffca52e")

type Client interface {
	// [Preview API] Delete a package version without an npm scope from the recycle bin.
	DeletePackageVersionFromRecycleBin(context.Context, DeletePackageVersionFromRecycleBinArgs) error
	// [Preview API] Delete a package version with an npm scope from the recycle bin.
	DeleteScopedPackageVersionFromRecycleBin(context.Context, DeleteScopedPackageVersionFromRecycleBinArgs) error
	// [Preview API]
	GetContentScopedPackage(context.Context, GetContentScopedPackageArgs) (io.ReadCloser, error)
	// [Preview API] Get an unscoped npm package.
	GetContentUnscopedPackage(context.Context, GetContentUnscopedPackageArgs) (io.ReadCloser, error)
	// [Preview API] Get information about an unscoped package version.
	GetPackageInfo(context.Context, GetPackageInfoArgs) (*Package, error)
	// [Preview API] Get information about an unscoped package version in the recycle bin.
	GetPackageVersionMetadataFromRecycleBin(context.Context, GetPackageVersionMetadataFromRecycleBinArgs) (*NpmPackageVersionDeletionState, error)
	// [Preview API] Get the Readme for a package version with an npm scope.
	GetReadmeScopedPackage(context.Context, GetReadmeScopedPackageArgs) (io.ReadCloser, error)
	// [Preview API] Get the Readme for a package version that has no npm scope.
	GetReadmeUnscopedPackage(context.Context, GetReadmeUnscopedPackageArgs) (io.ReadCloser, error)
	// [Preview API] Get information about a scoped package version (such as @scope/name).
	GetScopedPackageInfo(context.Context, GetScopedPackageInfoArgs) (*Package, error)
	// [Preview API] Get information about a scoped package version in the recycle bin.
	GetScopedPackageVersionMetadataFromRecycleBin(context.Context, GetScopedPackageVersionMetadataFromRecycleBinArgs) (*NpmPackageVersionDeletionState, error)
	// [Preview API] Restore a package version without an npm scope from the recycle bin to its feed.
	RestorePackageVersionFromRecycleBin(context.Context, RestorePackageVersionFromRecycleBinArgs) error
	// [Preview API] Restore a package version with an npm scope from the recycle bin to its feed.
	RestoreScopedPackageVersionFromRecycleBin(context.Context, RestoreScopedPackageVersionFromRecycleBinArgs) error
	// [Preview API] Unpublish an unscoped package version.
	UnpublishPackage(context.Context, UnpublishPackageArgs) (*Package, error)
	// [Preview API] Unpublish a scoped package version (such as @scope/name).
	UnpublishScopedPackage(context.Context, UnpublishScopedPackageArgs) (*Package, error)
	// [Preview API]
	UpdatePackage(context.Context, UpdatePackageArgs) (*Package, error)
	// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
	UpdatePackages(context.Context, UpdatePackagesArgs) error
	// [Preview API]
	UpdateScopedPackage(context.Context, UpdateScopedPackageArgs) (*Package, error)
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

// [Preview API] Delete a package version without an npm scope from the recycle bin.
func (client *ClientImpl) DeletePackageVersionFromRecycleBin(ctx context.Context, args DeletePackageVersionFromRecycleBinArgs) error {
	routeValues := make(map[string]string)
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

	locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
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
}

// [Preview API] Delete a package version with an npm scope from the recycle bin.
func (client *ClientImpl) DeleteScopedPackageVersionFromRecycleBin(ctx context.Context, args DeleteScopedPackageVersionFromRecycleBinArgs) error {
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteScopedPackageVersionFromRecycleBin function
type DeleteScopedPackageVersionFromRecycleBinArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Scope of the package (the 'scope' part of @scope/name).
	PackageScope *string
	// (required) Name of the package (the 'name' part of @scope/name).
	UnscopedPackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API]
func (client *ClientImpl) GetContentScopedPackage(ctx context.Context, args GetContentScopedPackageArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("09a4eafd-123a-495c-979c-0eda7bdb9a14")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetContentScopedPackage function
type GetContentScopedPackageArgs struct {
	// (required)
	FeedId *string
	// (required)
	PackageScope *string
	// (required)
	UnscopedPackageName *string
	// (required)
	PackageVersion *string
}

// [Preview API] Get an unscoped npm package.
func (client *ClientImpl) GetContentUnscopedPackage(ctx context.Context, args GetContentUnscopedPackageArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
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

	locationId, _ := uuid.Parse("75caa482-cb1e-47cd-9f2c-c048a4b7a43e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetContentUnscopedPackage function
type GetContentUnscopedPackageArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Get information about an unscoped package version.
func (client *ClientImpl) GetPackageInfo(ctx context.Context, args GetPackageInfoArgs) (*Package, error) {
	routeValues := make(map[string]string)
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

	locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPackageInfo function
type GetPackageInfoArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Get information about an unscoped package version in the recycle bin.
func (client *ClientImpl) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetPackageVersionMetadataFromRecycleBinArgs) (*NpmPackageVersionDeletionState, error) {
	routeValues := make(map[string]string)
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

	locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue NpmPackageVersionDeletionState
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
}

// [Preview API] Get the Readme for a package version with an npm scope.
func (client *ClientImpl) GetReadmeScopedPackage(ctx context.Context, args GetReadmeScopedPackageArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("6d4db777-7e4a-43b2-afad-779a1d197301")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetReadmeScopedPackage function
type GetReadmeScopedPackageArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Scope of the package (the 'scope' part of @scope\name)
	PackageScope *string
	// (required) Name of the package (the 'name' part of @scope\name)
	UnscopedPackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Get the Readme for a package version that has no npm scope.
func (client *ClientImpl) GetReadmeUnscopedPackage(ctx context.Context, args GetReadmeUnscopedPackageArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
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

	locationId, _ := uuid.Parse("1099a396-b310-41d4-a4b6-33d134ce3fcf")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetReadmeUnscopedPackage function
type GetReadmeUnscopedPackageArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Get information about a scoped package version (such as @scope/name).
func (client *ClientImpl) GetScopedPackageInfo(ctx context.Context, args GetScopedPackageInfoArgs) (*Package, error) {
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetScopedPackageInfo function
type GetScopedPackageInfoArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Scope of the package (the 'scope' part of @scope/name).
	PackageScope *string
	// (required) Name of the package (the 'name' part of @scope/name).
	UnscopedPackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Get information about a scoped package version in the recycle bin.
func (client *ClientImpl) GetScopedPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetScopedPackageVersionMetadataFromRecycleBinArgs) (*NpmPackageVersionDeletionState, error) {
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue NpmPackageVersionDeletionState
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetScopedPackageVersionMetadataFromRecycleBin function
type GetScopedPackageVersionMetadataFromRecycleBinArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Scope of the package (the 'scope' part of @scope/name)
	PackageScope *string
	// (required) Name of the package (the 'name' part of @scope/name).
	UnscopedPackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Restore a package version without an npm scope from the recycle bin to its feed.
func (client *ClientImpl) RestorePackageVersionFromRecycleBin(ctx context.Context, args RestorePackageVersionFromRecycleBinArgs) error {
	if args.PackageVersionDetails == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PackageVersionDetails"}
	}
	routeValues := make(map[string]string)
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
	locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RestorePackageVersionFromRecycleBin function
type RestorePackageVersionFromRecycleBinArgs struct {
	// (required)
	PackageVersionDetails *NpmRecycleBinPackageVersionDetails
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Restore a package version with an npm scope from the recycle bin to its feed.
func (client *ClientImpl) RestoreScopedPackageVersionFromRecycleBin(ctx context.Context, args RestoreScopedPackageVersionFromRecycleBinArgs) error {
	if args.PackageVersionDetails == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PackageVersionDetails"}
	}
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	body, marshalErr := json.Marshal(*args.PackageVersionDetails)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the RestoreScopedPackageVersionFromRecycleBin function
type RestoreScopedPackageVersionFromRecycleBinArgs struct {
	// (required)
	PackageVersionDetails *NpmRecycleBinPackageVersionDetails
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Scope of the package (the 'scope' part of @scope/name).
	PackageScope *string
	// (required) Name of the package (the 'name' part of @scope/name).
	UnscopedPackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Unpublish an unscoped package version.
func (client *ClientImpl) UnpublishPackage(ctx context.Context, args UnpublishPackageArgs) (*Package, error) {
	routeValues := make(map[string]string)
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

	locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
	resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UnpublishPackage function
type UnpublishPackageArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Name of the package.
	PackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API] Unpublish a scoped package version (such as @scope/name).
func (client *ClientImpl) UnpublishScopedPackage(ctx context.Context, args UnpublishScopedPackageArgs) (*Package, error) {
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
	resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UnpublishScopedPackage function
type UnpublishScopedPackageArgs struct {
	// (required) Name or ID of the feed.
	FeedId *string
	// (required) Scope of the package (the 'scope' part of @scope/name).
	PackageScope *string
	// (required) Name of the package (the 'name' part of @scope/name).
	UnscopedPackageName *string
	// (required) Version of the package.
	PackageVersion *string
}

// [Preview API]
func (client *ClientImpl) UpdatePackage(ctx context.Context, args UpdatePackageArgs) (*Package, error) {
	if args.PackageVersionDetails == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PackageVersionDetails"}
	}
	routeValues := make(map[string]string)
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

	body, marshalErr := json.Marshal(*args.PackageVersionDetails)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdatePackage function
type UpdatePackageArgs struct {
	// (required)
	PackageVersionDetails *PackageVersionDetails
	// (required)
	FeedId *string
	// (required)
	PackageName *string
	// (required)
	PackageVersion *string
}

// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
func (client *ClientImpl) UpdatePackages(ctx context.Context, args UpdatePackagesArgs) error {
	if args.BatchRequest == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.BatchRequest"}
	}
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId

	body, marshalErr := json.Marshal(*args.BatchRequest)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("06f34005-bbb2-41f4-88f5-23e03a99bb12")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePackages function
type UpdatePackagesArgs struct {
	// (required) Information about the packages to update, the operation to perform, and its associated data.
	BatchRequest *NpmPackagesBatchRequest
	// (required) Name or ID of the feed.
	FeedId *string
}

// [Preview API]
func (client *ClientImpl) UpdateScopedPackage(ctx context.Context, args UpdateScopedPackageArgs) (*Package, error) {
	if args.PackageVersionDetails == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PackageVersionDetails"}
	}
	routeValues := make(map[string]string)
	if args.FeedId == nil || *args.FeedId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FeedId"}
	}
	routeValues["feedId"] = *args.FeedId
	if args.PackageScope == nil || *args.PackageScope == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageScope"}
	}
	routeValues["packageScope"] = *args.PackageScope
	if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.UnscopedPackageName"}
	}
	routeValues["unscopedPackageName"] = *args.UnscopedPackageName
	if args.PackageVersion == nil || *args.PackageVersion == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.PackageVersion"}
	}
	routeValues["packageVersion"] = *args.PackageVersion

	body, marshalErr := json.Marshal(*args.PackageVersionDetails)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Package
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateScopedPackage function
type UpdateScopedPackageArgs struct {
	// (required)
	PackageVersionDetails *PackageVersionDetails
	// (required)
	FeedId *string
	// (required)
	PackageScope *string
	// (required)
	UnscopedPackageName *string
	// (required)
	PackageVersion *string
}
