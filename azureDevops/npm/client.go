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
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
)

var ResourceAreaId, _ = uuid.Parse("4c83cfc1-f33a-477e-a789-29d38ffca52e")

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API]
// feedId (required)
// packageScope (required)
// unscopedPackageName (required)
// packageVersion (required)
func (client Client) GetContentScopedPackage(feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("09a4eafd-123a-495c-979c-0eda7bdb9a14")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get an unscoped npm package.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetContentUnscopedPackage(feedId *string, packageName *string, packageVersion *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("75caa482-cb1e-47cd-9f2c-c048a4b7a43e")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update several packages from a single feed in a single request. The updates to the packages do not happen atomically.
// batchRequest (required): Information about the packages to update, the operation to perform, and its associated data.
// feedId (required): Name or ID of the feed.
func (client Client) UpdatePackages(batchRequest *NpmPackagesBatchRequest, feedId *string) error {
    if batchRequest == nil {
        return errors.New("batchRequest is a required parameter")
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*batchRequest)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("06f34005-bbb2-41f4-88f5-23e03a99bb12")
    _, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the Readme for a package version with an npm scope.
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope\name)
// unscopedPackageName (required): Name of the package (the 'name' part of @scope\name)
// packageVersion (required): Version of the package.
func (client Client) GetReadmeScopedPackage(feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("6d4db777-7e4a-43b2-afad-779a1d197301")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the Readme for a package version that has no npm scope.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetReadmeUnscopedPackage(feedId *string, packageName *string, packageVersion *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("1099a396-b310-41d4-a4b6-33d134ce3fcf")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "text/plain", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a package version with an npm scope from the recycle bin.
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) DeleteScopedPackageVersionFromRecycleBin(feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) error {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a scoped package version in the recycle bin.
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name)
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) GetScopedPackageVersionMetadataFromRecycleBin(feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*NpmPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NpmPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version with an npm scope from the recycle bin to its feed.
// packageVersionDetails (required)
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) RestoreScopedPackageVersionFromRecycleBin(packageVersionDetails *NpmRecycleBinPackageVersionDetails, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) error {
    if packageVersionDetails == nil {
        return errors.New("packageVersionDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("220f45eb-94a5-432c-902a-5b8c6372e415")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Delete a package version without an npm scope from the recycle bin.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) DeletePackageVersionFromRecycleBin(feedId *string, packageName *string, packageVersion *string) error {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about an unscoped package version in the recycle bin.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetPackageVersionMetadataFromRecycleBin(feedId *string, packageName *string, packageVersion *string) (*NpmPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NpmPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version without an npm scope from the recycle bin to its feed.
// packageVersionDetails (required)
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) RestorePackageVersionFromRecycleBin(packageVersionDetails *NpmRecycleBinPackageVersionDetails, feedId *string, packageName *string, packageVersion *string) error {
    if packageVersionDetails == nil {
        return errors.New("packageVersionDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("63a4f31f-e92b-4ee4-bf92-22d485e73bef")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a scoped package version (such as @scope/name).
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) GetScopedPackageInfo(feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Unpublish a scoped package version (such as @scope/name).
// feedId (required): Name or ID of the feed.
// packageScope (required): Scope of the package (the 'scope' part of @scope/name).
// unscopedPackageName (required): Name of the package (the 'name' part of @scope/name).
// packageVersion (required): Version of the package.
func (client Client) UnpublishScopedPackage(feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// packageVersionDetails (required)
// feedId (required)
// packageScope (required)
// unscopedPackageName (required)
// packageVersion (required)
func (client Client) UpdateScopedPackage(packageVersionDetails *PackageVersionDetails, feedId *string, packageScope *string, unscopedPackageName *string, packageVersion *string) (*Package, error) {
    if packageVersionDetails == nil {
        return nil, errors.New("packageVersionDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageScope == nil || *packageScope == "" {
        return nil, errors.New("packageScope is a required parameter")
    }
    routeValues["packageScope"] = *packageScope
    if unscopedPackageName == nil || *unscopedPackageName == "" {
        return nil, errors.New("unscopedPackageName is a required parameter")
    }
    routeValues["unscopedPackageName"] = *unscopedPackageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e93d9ec3-4022-401e-96b0-83ea5d911e09")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get information about an unscoped package version.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetPackageInfo(feedId *string, packageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Unpublish an unscoped package version.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) UnpublishPackage(feedId *string, packageName *string, packageVersion *string) (*Package, error) {
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// packageVersionDetails (required)
// feedId (required)
// packageName (required)
// packageVersion (required)
func (client Client) UpdatePackage(packageVersionDetails *PackageVersionDetails, feedId *string, packageName *string, packageVersion *string) (*Package, error) {
    if packageVersionDetails == nil {
        return nil, errors.New("packageVersionDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if packageName == nil || *packageName == "" {
        return nil, errors.New("packageName is a required parameter")
    }
    routeValues["packageName"] = *packageName
    if packageVersion == nil || *packageVersion == "" {
        return nil, errors.New("packageVersion is a required parameter")
    }
    routeValues["packageVersion"] = *packageVersion

    body, marshalErr := json.Marshal(*packageVersionDetails)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ed579d62-67c9-4271-be66-9b029af5bcf9")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

