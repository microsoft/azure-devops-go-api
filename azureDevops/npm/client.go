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
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
)

var ResourceAreaId, _ = uuid.Parse("4c83cfc1-f33a-477e-a789-29d38ffca52e")

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API]
// ctx
// feedId (required)
// packageScope (required)
// unscopedPackageName (required)
// packageVersion (required)
func (client Client) GetContentScopedPackage(ctx context.Context, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("09a4eafd-123a-495c-979c-0eda7bdb9a14")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Get an unscoped npm package.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetContentUnscopedPackage(ctx context.Context, feedId *string, packageName *string, packageVersion *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("75caa482-cb1e-47cd-9f2c-c048a4b7a43e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
// ctx
// batchRequest (required): Information about the packages to update, the operation to perform, and its associated data.
// feedId (required): Name or ID of the feed.
func (client Client) UpdatePackages(ctx context.Context, batchRequest *NpmPackagesBatchRequest, feedId *string) error {
    if batchRequest == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "batchRequest"}
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*batchRequest)
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

// [Preview API] Get the Readme for a package version with an npm scope.
// ctx
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope\name)
// unscopedPackageName (required): Name of the package (the 'name' part of @scope\name)
// packageVersion (required): Version of the package.
func (client Client) GetReadmeScopedPackage(ctx context.Context, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("6d4db777-7e4a-43b2-afad-779a1d197301")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Get the Readme for a package version that has no npm scope.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetReadmeUnscopedPackage(ctx context.Context, feedId *string, packageName *string, packageVersion *string) (interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("1099a396-b310-41d4-a4b6-33d134ce3fcf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, responseValue)
    return responseValue, err
}

// [Preview API] Delete a package version with an npm scope from the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) DeleteScopedPackageVersionFromRecycleBin(ctx context.Context, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) error {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a scoped package version in the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name)
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) GetScopedPackageVersionMetadataFromRecycleBin(ctx context.Context, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*NpmPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NpmPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version with an npm scope from the recycle bin to its feed.
// ctx
// packageVersionDetails (required)
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) RestoreScopedPackageVersionFromRecycleBin(ctx context.Context, packageVersionDetails *NpmRecycleBinPackageVersionDetails, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) error {
    if packageVersionDetails == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
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

// [Preview API] Delete a package version without an npm scope from the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) DeletePackageVersionFromRecycleBin(ctx context.Context, feedId *string, packageName *string, packageVersion *string) error {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about an unscoped package version in the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, feedId *string, packageName *string, packageVersion *string) (*NpmPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NpmPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version without an npm scope from the recycle bin to its feed.
// ctx
// packageVersionDetails (required)
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) RestorePackageVersionFromRecycleBin(ctx context.Context, packageVersionDetails *NpmRecycleBinPackageVersionDetails, feedId *string, packageName *string, packageVersion *string) error {
    if packageVersionDetails == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
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

// [Preview API] Get information about a scoped package version (such as @scope/name).
// ctx
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) GetScopedPackageInfo(ctx context.Context, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Unpublish a scoped package version (such as @scope/name).
// ctx
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) UnpublishScopedPackage(ctx context.Context, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// packageVersionDetails (required)
// feedId (required)
// packageScope (required)
// unscopedPackageName (required)
// packageVersion (required)
func (client Client) UpdateScopedPackage(ctx context.Context, packageVersionDetails *PackageVersionDetails, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*Package, error) {
    if packageVersionDetails == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageScope"} 
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "unscopedPackageName"} 
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
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

// [Preview API] Get information about an unscoped package version.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetPackageInfo(ctx context.Context, feedId *string, packageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Unpublish an unscoped package version.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) UnpublishPackage(ctx context.Context, feedId *string, packageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// packageVersionDetails (required)
// feedId (required)
// packageName (required)
// packageVersion (required)
func (client Client) UpdatePackage(ctx context.Context, packageVersionDetails *PackageVersionDetails, feedId *string, packageName *string, packageVersion *string) (*Package, error) {
    if packageVersionDetails == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageVersionDetails"}
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
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

