// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package nuGet

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("b3be7473-68ea-4a81-bfc7-9530baaa19ad")

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

// [Preview API] Download a package version directly.  This API is intended for manual UI download options, not for programmatic access and scripting.  You may be heavily throttled if accessing this api for scripting purposes.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
// sourceProtocolVersion (optional): Unused
func (client Client) DownloadPackage(feedId *string, packageName *string, packageVersion *string, project *string, sourceProtocolVersion *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
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

    queryParams := url.Values{}
    if sourceProtocolVersion != nil {
        queryParams.Add("sourceProtocolVersion", *sourceProtocolVersion)
    }
    locationId, _ := uuid.Parse("6ea81b8c-7386-490b-a71f-6cf23c80b388")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
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
// project (optional): Project ID or project name
func (client Client) UpdatePackageVersions(batchRequest *NuGetPackagesBatchRequest, feedId *string, project *string) error {
    if batchRequest == nil {
        return errors.New("batchRequest is a required parameter")
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*batchRequest)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("00c58ea7-d55f-49de-b59f-983533ae11dc")
    _, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Delete a package version from a feed's recycle bin.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) DeletePackageVersionFromRecycleBin(feedId *string, packageName *string, packageVersion *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
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

    locationId, _ := uuid.Parse("07e88775-e3cb-4408-bbe1-628e036fac8c")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] View a package version's deletion/recycled status
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) GetPackageVersionMetadataFromRecycleBin(feedId *string, packageName *string, packageVersion *string, project *string) (*NuGetPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
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

    locationId, _ := uuid.Parse("07e88775-e3cb-4408-bbe1-628e036fac8c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NuGetPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version from a feed's recycle bin back into the active feed.
// packageVersionDetails (required): Set the 'Deleted' member to 'false' to apply the restore operation
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) RestorePackageVersionFromRecycleBin(packageVersionDetails *NuGetRecycleBinPackageVersionDetails, feedId *string, packageName *string, packageVersion *string, project *string) error {
    if packageVersionDetails == nil {
        return errors.New("packageVersionDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
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
    locationId, _ := uuid.Parse("07e88775-e3cb-4408-bbe1-628e036fac8c")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Send a package version from the feed to its paired recycle bin.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package to delete.
// packageVersion (required): Version of the package to delete.
// project (optional): Project ID or project name
func (client Client) DeletePackageVersion(feedId *string, packageName *string, packageVersion *string, project *string) (*Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
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

    locationId, _ := uuid.Parse("36c9353b-e250-4c57-b040-513c186c3905")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get information about a package version.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// project (optional): Project ID or project name
// showDeleted (optional): True to include deleted packages in the response.
func (client Client) GetPackageVersion(feedId *string, packageName *string, packageVersion *string, project *string, showDeleted *bool) (*Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
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

    queryParams := url.Values{}
    if showDeleted != nil {
        queryParams.Add("showDeleted", strconv.FormatBool(*showDeleted))
    }
    locationId, _ := uuid.Parse("36c9353b-e250-4c57-b040-513c186c3905")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Set mutable state on a package version.
// packageVersionDetails (required): New state to apply to the referenced package.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package to update.
// packageVersion (required): Version of the package to update.
// project (optional): Project ID or project name
func (client Client) UpdatePackageVersion(packageVersionDetails *PackageVersionDetails, feedId *string, packageName *string, packageVersion *string, project *string) error {
    if packageVersionDetails == nil {
        return errors.New("packageVersionDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
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
    locationId, _ := uuid.Parse("36c9353b-e250-4c57-b040-513c186c3905")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

