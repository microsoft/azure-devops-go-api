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
    "github.com/microsoft/azure-devops-go-api/azureDevOps"
    "io"
    "net/http"
)

var ResourceAreaId, _ = uuid.Parse("4c83cfc1-f33a-477e-a789-29d38ffca52e")

type Client struct {
    Client azureDevOps.Client
}

func NewClient(ctx context.Context, connection *azureDevOps.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API]
func (client *Client) GetContentScopedPackage(ctx context.Context, args GetContentScopedPackageArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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
func (client *Client) GetContentUnscopedPackage(ctx context.Context, args GetContentUnscopedPackageArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
func (client *Client) UpdatePackages(ctx context.Context, args UpdatePackagesArgs) error {
    if args.BatchRequest == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "batchRequest"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
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

// [Preview API] Get the Readme for a package version with an npm scope.
func (client *Client) GetReadmeScopedPackage(ctx context.Context, args GetReadmeScopedPackageArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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
func (client *Client) GetReadmeUnscopedPackage(ctx context.Context, args GetReadmeUnscopedPackageArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Delete a package version with an npm scope from the recycle bin.
func (client *Client) DeleteScopedPackageVersionFromRecycleBin(ctx context.Context, args DeleteScopedPackageVersionFromRecycleBinArgs) error {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Get information about a scoped package version in the recycle bin.
func (client *Client) GetScopedPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetScopedPackageVersionMetadataFromRecycleBinArgs) (*NpmPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Restore a package version with an npm scope from the recycle bin to its feed.
func (client *Client) RestoreScopedPackageVersionFromRecycleBin(ctx context.Context, args RestoreScopedPackageVersionFromRecycleBinArgs) error {
    if args.PackageVersionDetails == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Delete a package version without an npm scope from the recycle bin.
func (client *Client) DeletePackageVersionFromRecycleBin(ctx context.Context, args DeletePackageVersionFromRecycleBinArgs) error {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Get information about an unscoped package version in the recycle bin.
func (client *Client) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetPackageVersionMetadataFromRecycleBinArgs) (*NpmPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Restore a package version without an npm scope from the recycle bin to its feed.
func (client *Client) RestorePackageVersionFromRecycleBin(ctx context.Context, args RestorePackageVersionFromRecycleBinArgs) error {
    if args.PackageVersionDetails == nil {
        return &azureDevOps.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Get information about a scoped package version (such as @scope/name).
func (client *Client) GetScopedPackageInfo(ctx context.Context, args GetScopedPackageInfoArgs) (*Package, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Unpublish a scoped package version (such as @scope/name).
func (client *Client) UnpublishScopedPackage(ctx context.Context, args UnpublishScopedPackageArgs) (*Package, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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
func (client *Client) UpdateScopedPackage(ctx context.Context, args UpdateScopedPackageArgs) (*Package, error) {
    if args.PackageVersionDetails == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageScope == nil || *args.PackageScope == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *args.PackageScope
    if args.UnscopedPackageName == nil || *args.UnscopedPackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *args.UnscopedPackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Get information about an unscoped package version.
func (client *Client) GetPackageInfo(ctx context.Context, args GetPackageInfoArgs) (*Package, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API] Unpublish an unscoped package version.
func (client *Client) UnpublishPackage(ctx context.Context, args UnpublishPackageArgs) (*Package, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

// [Preview API]
func (client *Client) UpdatePackage(ctx context.Context, args UpdatePackageArgs) (*Package, error) {
    if args.PackageVersionDetails == nil {
        return nil, &azureDevOps.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevOps.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
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

