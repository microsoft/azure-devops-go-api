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
    "context"
    "encoding/json"
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

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// [Preview API] Delete a package version from the recycle bin.
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

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a package version in the recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, feedId *string, packageName *string, packageVersion *string) (*UPackPackageVersionDeletionState, error) {
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

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version from the recycle bin to its associated feed.
// ctx
// packageVersionDetails (required): Set the 'Deleted' property to 'false' to restore the package.
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) RestorePackageVersionFromRecycleBin(ctx context.Context, packageVersionDetails *UPackRecycleBinPackageVersionDetails, feedId *string, packageName *string, packageVersion *string) error {
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
    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Delete a package version from a feed's recycle bin.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) DeletePackageVersion(ctx context.Context, feedId *string, packageName *string, packageVersion *string) (*Package, error) {
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

    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Show information about a package version.
// ctx
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
// showDeleted (optional): True to show information for deleted versions
func (client Client) GetPackageVersion(ctx context.Context, feedId *string, packageName *string, packageVersion *string, showDeleted *bool) (*Package, error) {
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

    queryParams := url.Values{}
    if showDeleted != nil {
        queryParams.Add("showDeleted", strconv.FormatBool(*showDeleted))
    }
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update information for a package version.
// ctx
// packageVersionDetails (required)
// feedId (required): Name or ID of the feed.
// packageName (required): Name of the package.
// packageVersion (required): Version of the package.
func (client Client) UpdatePackageVersion(ctx context.Context, packageVersionDetails *PackageVersionDetails, feedId *string, packageName *string, packageVersion *string) error {
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
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
    _, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

