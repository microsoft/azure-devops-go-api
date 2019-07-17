// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package uPackApi

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

var ResourceAreaId, _ = uuid.Parse("d397749b-f115-4027-b6dd-77a65dd10d21")

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

// [Preview API] Delete a package version from the recycle bin.
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

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a package version in the recycle bin.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetPackageVersionMetadataFromRecycleBin(feedId *string, packageName *string, packageVersion *string) (*UPackPackageVersionDeletionState, error) {
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

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version from the recycle bin to its associated feed.
// packageVersionDetails (required): Set the 'Deleted' property to 'false' to restore the package.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) RestorePackageVersionFromRecycleBin(packageVersionDetails *UPackRecycleBinPackageVersionDetails, feedId *string, packageName *string, packageVersion *string) error {
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
    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Delete a package version from a feed's recycle bin.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) DeletePackageVersion(feedId *string, packageName *string, packageVersion *string) (*Package, error) {
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

    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Show information about a package version.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// showDeleted (optional): True to show information for deleted versions
func (client Client) GetPackageVersion(feedId *string, packageName *string, packageVersion *string, showDeleted *bool) (*Package, error) {
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

    queryParams := url.Values{}
    if showDeleted != nil {
        queryParams.Add("showDeleted", strconv.FormatBool(*showDeleted))
    }
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update information for a package version.
// packageVersionDetails (required)
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) UpdatePackageVersion(packageVersionDetails *PackageVersionDetails, feedId *string, packageName *string, packageVersion *string) error {
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
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

