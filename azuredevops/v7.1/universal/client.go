// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package universal

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azuredevops/v7.1"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("d397749b-f115-4027-b6dd-77a65dd10d21")

type Client interface {
    // [Preview API] Delete a package version from a feed's recycle bin.
DeletePackageVersion(context.Context, DeletePackageVersionArgs) (*Package, error)
    // [Preview API] Delete a package version from the recycle bin.
DeletePackageVersionFromRecycleBin(context.Context, DeletePackageVersionFromRecycleBinArgs) error
    // [Preview API] Show information about a package version.
GetPackageVersion(context.Context, GetPackageVersionArgs) (*Package, error)
    // [Preview API] Get information about a package version in the recycle bin.
GetPackageVersionMetadataFromRecycleBin(context.Context, GetPackageVersionMetadataFromRecycleBinArgs) (*UPackPackageVersionDeletionState, error)
    // [Preview API] Restore a package version from the recycle bin to its associated feed.
RestorePackageVersionFromRecycleBin(context.Context, RestorePackageVersionFromRecycleBinArgs) error
    // [Preview API] Update information for a package version.
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

// [Preview API] Delete a package version from a feed's recycle bin.
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

    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
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

// [Preview API] Delete a package version from the recycle bin.
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

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
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

// [Preview API] Show information about a package version.
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
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
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
    // (optional) True to show information for deleted versions
ShowDeleted *bool
}

// [Preview API] Get information about a package version in the recycle bin.
func (client *ClientImpl) GetPackageVersionMetadataFromRecycleBin(ctx context.Context, args GetPackageVersionMetadataFromRecycleBinArgs) (*UPackPackageVersionDeletionState, error) {
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

    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "7.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackPackageVersionDeletionState
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
    locationId, _ := uuid.Parse("3ba455ae-31e6-409e-849f-56c66888d004")
_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the RestorePackageVersionFromRecycleBin function
type RestorePackageVersionFromRecycleBinArgs struct {
    // (required) Set the 'Deleted' property to 'false' to restore the package.
PackageVersionDetails *UPackRecycleBinPackageVersionDetails
    // (required) Name or ID of the feed.
FeedId *string
    // (required) Name of the package.
PackageName *string
    // (required) Version of the package.
PackageVersion *string
    // (optional) Project ID or project name
Project *string
}

// [Preview API] Update information for a package version.
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
    locationId, _ := uuid.Parse("72f61ca4-e07c-4eca-be75-6c0b2f3f4051")
_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the UpdatePackageVersion function
type UpdatePackageVersionArgs struct {
    // (required)
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
    locationId, _ := uuid.Parse("c17e81ae-4caa-4d8b-a431-6b329e890281")
_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the UpdatePackageVersions function
type UpdatePackageVersionsArgs struct {
    // (required) Information about the packages to update, the operation to perform, and its associated data.
BatchRequest *UPackPackagesBatchRequest
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
    locationId, _ := uuid.Parse("12f73313-0937-4114-bb9f-4e9e720fdc78")
_, err := client.Client.Send(ctx, http.MethodPost, locationId, "7.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the UpdateRecycleBinPackageVersions function
type UpdateRecycleBinPackageVersionsArgs struct {
    // (required) Information about the packages to update, the operation to perform, and its associated data. <c>Operation</c> must be <c>PermanentDelete</c> or <c>RestoreToFeed</c>
BatchRequest *UPackPackagesBatchRequest
    // (required) Feed which contains the packages to update.
FeedId *string
    // (optional) Project ID or project name
Project *string
}

