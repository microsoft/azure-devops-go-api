// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package feed

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

var ResourceAreaId, _ = uuid.Parse("7ab4e64e-c4d8-4f50-ae73-5ef2e21642a5")

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

// [Preview API] Generate a SVG badge for the latest version of a package.  The generated SVG is typically used as the image in an HTML link which takes users to the feed containing the package to accelerate discovery and consumption.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): Id of the package (GUID Id, not name).
// project (optional): Project ID or project name
func (client Client) GetBadge(ctx context.Context, feedId *string, packageId *uuid.UUID, project *string) (*string, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = (*packageId).String()

    locationId, _ := uuid.Parse("61d885fd-10f3-4a55-82b6-476d866b673f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Query a feed to determine its current state.
// ctx
// feedId (required): Name or ID of the feed.
// project (optional): Project ID or project name
func (client Client) GetFeedChange(ctx context.Context, feedId *string, project *string) (*FeedChange, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    locationId, _ := uuid.Parse("29ba2dad-389a-4661-b5d3-de76397ca05b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeedChange
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Query to determine which feeds have changed since the last call, tracked through the provided continuationToken. Only changes to a feed itself are returned and impact the continuationToken, not additions or alterations to packages within the feeds.
// ctx
// project (optional): Project ID or project name
// includeDeleted (optional): If true, get changes for all feeds including deleted feeds. The default value is false.
// continuationToken (optional): A continuation token which acts as a bookmark to a previously retrieved change. This token allows the user to continue retrieving changes in batches, picking up where the previous batch left off. If specified, all the changes that occur strictly after the token will be returned. If not specified or 0, iteration will start with the first change.
// batchSize (optional): Number of package changes to fetch. The default value is 1000. The maximum value is 2000.
func (client Client) GetFeedChanges(ctx context.Context, project *string, includeDeleted *bool, continuationToken *uint64, batchSize *int) (*FeedChangesResponse, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if continuationToken != nil {
        queryParams.Add("continuationToken", strconv.FormatUint(*continuationToken, 10))
    }
    if batchSize != nil {
        queryParams.Add("batchSize", strconv.Itoa(*batchSize))
    }
    locationId, _ := uuid.Parse("29ba2dad-389a-4661-b5d3-de76397ca05b")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeedChangesResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a feed, a container for various package types.
// ctx
// feed (required): A JSON object containing both required and optional attributes for the feed. Name is the only required value.
// project (optional): Project ID or project name
func (client Client) CreateFeed(ctx context.Context, feed *Feed, project *string) (*Feed, error) {
    if feed == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "feed"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    body, marshalErr := json.Marshal(*feed)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("c65009a7-474a-4ad1-8b42-7d852107ef8c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Feed
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Remove a feed and all its packages. The action does not result in packages moving to the RecycleBin and is not reversible.
// ctx
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
func (client Client) DeleteFeed(ctx context.Context, feedId *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    locationId, _ := uuid.Parse("c65009a7-474a-4ad1-8b42-7d852107ef8c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the settings for a specific feed.
// ctx
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
// includeDeletedUpstreams (optional): Include upstreams that have been deleted in the response.
func (client Client) GetFeed(ctx context.Context, feedId *string, project *string, includeDeletedUpstreams *bool) (*Feed, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    queryParams := url.Values{}
    if includeDeletedUpstreams != nil {
        queryParams.Add("includeDeletedUpstreams", strconv.FormatBool(*includeDeletedUpstreams))
    }
    locationId, _ := uuid.Parse("c65009a7-474a-4ad1-8b42-7d852107ef8c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Feed
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all feeds in an account where you have the provided role access.
// ctx
// project (optional): Project ID or project name
// feedRole (optional): Filter by this role, either Administrator(4), Contributor(3), or Reader(2) level permissions.
// includeDeletedUpstreams (optional): Include upstreams that have been deleted in the response.
func (client Client) GetFeeds(ctx context.Context, project *string, feedRole *string, includeDeletedUpstreams *bool) (*[]Feed, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if feedRole != nil {
        queryParams.Add("feedRole", *feedRole)
    }
    if includeDeletedUpstreams != nil {
        queryParams.Add("includeDeletedUpstreams", strconv.FormatBool(*includeDeletedUpstreams))
    }
    locationId, _ := uuid.Parse("c65009a7-474a-4ad1-8b42-7d852107ef8c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Feed
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Change the attributes of a feed.
// ctx
// feed (required): A JSON object containing the feed settings to be updated.
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
func (client Client) UpdateFeed(ctx context.Context, feed *FeedUpdate, feedId *string, project *string) (*Feed, error) {
    if feed == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "feed"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*feed)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("c65009a7-474a-4ad1-8b42-7d852107ef8c")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Feed
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all service-wide feed creation and administration permissions.
// ctx
// includeIds (optional): Set to true to add IdentityIds to the permission objects.
func (client Client) GetGlobalPermissions(ctx context.Context, includeIds *bool) (*[]GlobalPermission, error) {
    queryParams := url.Values{}
    if includeIds != nil {
        queryParams.Add("includeIds", strconv.FormatBool(*includeIds))
    }
    locationId, _ := uuid.Parse("a74419ef-b477-43df-8758-3cd1cd5f56c6")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GlobalPermission
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Set service-wide permissions that govern feed creation and administration.
// ctx
// globalPermissions (required): New permissions for the organization.
func (client Client) SetGlobalPermissions(ctx context.Context, globalPermissions *[]GlobalPermission) (*[]GlobalPermission, error) {
    if globalPermissions == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "globalPermissions"}
    }
    body, marshalErr := json.Marshal(*globalPermissions)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a74419ef-b477-43df-8758-3cd1cd5f56c6")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []GlobalPermission
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a batch of package changes made to a feed.  The changes returned are 'most recent change' so if an Add is followed by an Update before you begin enumerating, you'll only see one change in the batch.  While consuming batches using the continuation token, you may see changes to the same package version multiple times if they are happening as you enumerate.
// ctx
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
// continuationToken (optional): A continuation token which acts as a bookmark to a previously retrieved change. This token allows the user to continue retrieving changes in batches, picking up where the previous batch left off. If specified, all the changes that occur strictly after the token will be returned. If not specified or 0, iteration will start with the first change.
// batchSize (optional): Number of package changes to fetch. The default value is 1000. The maximum value is 2000.
func (client Client) GetPackageChanges(ctx context.Context, feedId *string, project *string, continuationToken *uint64, batchSize *int) (*PackageChangesResponse, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    queryParams := url.Values{}
    if continuationToken != nil {
        queryParams.Add("continuationToken", strconv.FormatUint(*continuationToken, 10))
    }
    if batchSize != nil {
        queryParams.Add("batchSize", strconv.Itoa(*batchSize))
    }
    locationId, _ := uuid.Parse("323a0631-d083-4005-85ae-035114dfb681")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PackageChangesResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// packageIdQuery (required)
// feedId (required)
// project (optional): Project ID or project name
func (client Client) QueryPackageMetrics(ctx context.Context, packageIdQuery *PackageMetricsQuery, feedId *string, project *string) (*[]PackageMetrics, error) {
    if packageIdQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageIdQuery"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*packageIdQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bddc9b3c-8a59-4a9f-9b40-ee1dcaa2cc0d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PackageMetrics
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get details about a specific package.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): The package Id (GUID Id, not the package name).
// project (optional): Project ID or project name
// includeAllVersions (optional): True to return all versions of the package in the response. Default is false (latest version only).
// includeUrls (optional): True to return REST Urls with the response. Default is True.
// isListed (optional): Only applicable for NuGet packages, setting it for other package types will result in a 404. If false, delisted package versions will be returned. Use this to filter the response when includeAllVersions is set to true. Default is unset (do not return delisted packages).
// isRelease (optional): Only applicable for Nuget packages. Use this to filter the response when includeAllVersions is set to true.  Default is True (only return packages without prerelease versioning).
// includeDeleted (optional): Return deleted or unpublished versions of packages in the response. Default is False.
// includeDescription (optional): Return the description for every version of each package in the response. Default is False.
func (client Client) GetPackage(ctx context.Context, feedId *string, packageId *string, project *string, includeAllVersions *bool, includeUrls *bool, isListed *bool, isRelease *bool, includeDeleted *bool, includeDescription *bool) (*Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil || *packageId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = *packageId

    queryParams := url.Values{}
    if includeAllVersions != nil {
        queryParams.Add("includeAllVersions", strconv.FormatBool(*includeAllVersions))
    }
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    if isListed != nil {
        queryParams.Add("isListed", strconv.FormatBool(*isListed))
    }
    if isRelease != nil {
        queryParams.Add("isRelease", strconv.FormatBool(*isRelease))
    }
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if includeDescription != nil {
        queryParams.Add("includeDescription", strconv.FormatBool(*includeDescription))
    }
    locationId, _ := uuid.Parse("7a20d846-c929-4acc-9ea2-0d5a7df1b197")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get details about all of the packages in the feed. Use the various filters to include or exclude information from the result set.
// ctx
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
// protocolType (optional): One of the supported artifact package types.
// packageNameQuery (optional): Filter to packages that contain the provided string. Characters in the string must conform to the package name constraints.
// normalizedPackageName (optional): [Obsolete] Used for legacy scenarios and may be removed in future versions.
// includeUrls (optional): True to return REST Urls with the response. Default is True.
// includeAllVersions (optional): True to return all versions of the package in the response. Default is false (latest version only).
// isListed (optional): Only applicable for NuGet packages, setting it for other package types will result in a 404. If false, delisted package versions will be returned. Use this to filter the response when includeAllVersions is set to true. Default is unset (do not return delisted packages).
// getTopPackageVersions (optional): Changes the behavior of $top and $skip to return all versions of each package up to $top. Must be used in conjunction with includeAllVersions=true
// isRelease (optional): Only applicable for Nuget packages. Use this to filter the response when includeAllVersions is set to true. Default is True (only return packages without prerelease versioning).
// includeDescription (optional): Return the description for every version of each package in the response. Default is False.
// top (optional): Get the top N packages (or package versions where getTopPackageVersions=true)
// skip (optional): Skip the first N packages (or package versions where getTopPackageVersions=true)
// includeDeleted (optional): Return deleted or unpublished versions of packages in the response. Default is False.
// isCached (optional): [Obsolete] Used for legacy scenarios and may be removed in future versions.
// directUpstreamId (optional): Filter results to return packages from a specific upstream.
func (client Client) GetPackages(ctx context.Context, feedId *string, project *string, protocolType *string, packageNameQuery *string, normalizedPackageName *string, includeUrls *bool, includeAllVersions *bool, isListed *bool, getTopPackageVersions *bool, isRelease *bool, includeDescription *bool, top *int, skip *int, includeDeleted *bool, isCached *bool, directUpstreamId *uuid.UUID) (*[]Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    queryParams := url.Values{}
    if protocolType != nil {
        queryParams.Add("protocolType", *protocolType)
    }
    if packageNameQuery != nil {
        queryParams.Add("packageNameQuery", *packageNameQuery)
    }
    if normalizedPackageName != nil {
        queryParams.Add("normalizedPackageName", *normalizedPackageName)
    }
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    if includeAllVersions != nil {
        queryParams.Add("includeAllVersions", strconv.FormatBool(*includeAllVersions))
    }
    if isListed != nil {
        queryParams.Add("isListed", strconv.FormatBool(*isListed))
    }
    if getTopPackageVersions != nil {
        queryParams.Add("getTopPackageVersions", strconv.FormatBool(*getTopPackageVersions))
    }
    if isRelease != nil {
        queryParams.Add("isRelease", strconv.FormatBool(*isRelease))
    }
    if includeDescription != nil {
        queryParams.Add("includeDescription", strconv.FormatBool(*includeDescription))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if includeDeleted != nil {
        queryParams.Add("includeDeleted", strconv.FormatBool(*includeDeleted))
    }
    if isCached != nil {
        queryParams.Add("isCached", strconv.FormatBool(*isCached))
    }
    if directUpstreamId != nil {
        queryParams.Add("directUpstreamId", (*directUpstreamId).String())
    }
    locationId, _ := uuid.Parse("7a20d846-c929-4acc-9ea2-0d5a7df1b197")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Package
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the permissions for a feed.
// ctx
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
// includeIds (optional): True to include user Ids in the response.  Default is false.
// excludeInheritedPermissions (optional): True to only return explicitly set permissions on the feed.  Default is false.
// identityDescriptor (optional): Filter permissions to the provided identity.
func (client Client) GetFeedPermissions(ctx context.Context, feedId *string, project *string, includeIds *bool, excludeInheritedPermissions *bool, identityDescriptor *string) (*[]FeedPermission, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    queryParams := url.Values{}
    if includeIds != nil {
        queryParams.Add("includeIds", strconv.FormatBool(*includeIds))
    }
    if excludeInheritedPermissions != nil {
        queryParams.Add("excludeInheritedPermissions", strconv.FormatBool(*excludeInheritedPermissions))
    }
    if identityDescriptor != nil {
        queryParams.Add("identityDescriptor", *identityDescriptor)
    }
    locationId, _ := uuid.Parse("be8c1476-86a7-44ed-b19d-aec0e9275cd8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []FeedPermission
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update the permissions on a feed.
// ctx
// feedPermission (required): Permissions to set.
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
func (client Client) SetFeedPermissions(ctx context.Context, feedPermission *[]FeedPermission, feedId *string, project *string) (*[]FeedPermission, error) {
    if feedPermission == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "feedPermission"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*feedPermission)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("be8c1476-86a7-44ed-b19d-aec0e9275cd8")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []FeedPermission
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Gets provenance for a package version.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): Id of the package (GUID Id, not name).
// packageVersionId (required): Id of the package version (GUID Id, not name).
// project (optional): Project ID or project name
func (client Client) GetPackageVersionProvenance(ctx context.Context, feedId *string, packageId *uuid.UUID, packageVersionId *uuid.UUID, project *string) (*PackageVersionProvenance, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = (*packageId).String()
    if packageVersionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageVersionId"} 
    }
    routeValues["packageVersionId"] = (*packageVersionId).String()

    locationId, _ := uuid.Parse("0aaeabd4-85cd-4686-8a77-8d31c15690b8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PackageVersionProvenance
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get information about a package and all its versions within the recycle bin.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): The package Id (GUID Id, not the package name).
// project (optional): Project ID or project name
// includeUrls (optional): True to return REST Urls with the response.  Default is True.
func (client Client) GetRecycleBinPackage(ctx context.Context, feedId *string, packageId *uuid.UUID, project *string, includeUrls *bool) (*Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = (*packageId).String()

    queryParams := url.Values{}
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    locationId, _ := uuid.Parse("2704e72c-f541-4141-99be-2004b50b05fa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Query for packages within the recycle bin.
// ctx
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
// protocolType (optional): Type of package (e.g. NuGet, npm, ...).
// packageNameQuery (optional): Filter to packages matching this name.
// includeUrls (optional): True to return REST Urls with the response.  Default is True.
// top (optional): Get the top N packages.
// skip (optional): Skip the first N packages.
// includeAllVersions (optional): True to return all versions of the package in the response.  Default is false (latest version only).
func (client Client) GetRecycleBinPackages(ctx context.Context, feedId *string, project *string, protocolType *string, packageNameQuery *string, includeUrls *bool, top *int, skip *int, includeAllVersions *bool) (*[]Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    queryParams := url.Values{}
    if protocolType != nil {
        queryParams.Add("protocolType", *protocolType)
    }
    if packageNameQuery != nil {
        queryParams.Add("packageNameQuery", *packageNameQuery)
    }
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    if top != nil {
        queryParams.Add("$top", strconv.Itoa(*top))
    }
    if skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*skip))
    }
    if includeAllVersions != nil {
        queryParams.Add("includeAllVersions", strconv.FormatBool(*includeAllVersions))
    }
    locationId, _ := uuid.Parse("2704e72c-f541-4141-99be-2004b50b05fa")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Package
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get information about a package version within the recycle bin.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): The package Id (GUID Id, not the package name).
// packageVersionId (required): The package version Id 9guid Id, not the version string).
// project (optional): Project ID or project name
// includeUrls (optional): True to return REST Urls with the response.  Default is True.
func (client Client) GetRecycleBinPackageVersion(ctx context.Context, feedId *string, packageId *uuid.UUID, packageVersionId *uuid.UUID, project *string, includeUrls *bool) (*RecycleBinPackageVersion, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = (*packageId).String()
    if packageVersionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageVersionId"} 
    }
    routeValues["packageVersionId"] = (*packageVersionId).String()

    queryParams := url.Values{}
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    locationId, _ := uuid.Parse("aceb4be7-8737-4820-834c-4c549e10fdc7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue RecycleBinPackageVersion
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of package versions within the recycle bin.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): The package Id (GUID Id, not the package name).
// project (optional): Project ID or project name
// includeUrls (optional): True to return REST Urls with the response.  Default is True.
func (client Client) GetRecycleBinPackageVersions(ctx context.Context, feedId *string, packageId *uuid.UUID, project *string, includeUrls *bool) (*[]RecycleBinPackageVersion, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = (*packageId).String()

    queryParams := url.Values{}
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    locationId, _ := uuid.Parse("aceb4be7-8737-4820-834c-4c549e10fdc7")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []RecycleBinPackageVersion
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete the retention policy for a feed.
// ctx
// feedId (required): Name or ID of the feed.
// project (optional): Project ID or project name
func (client Client) DeleteFeedRetentionPolicies(ctx context.Context, feedId *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    locationId, _ := uuid.Parse("ed52a011-0112-45b5-9f9e-e14efffb3193")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get the retention policy for a feed.
// ctx
// feedId (required): Name or ID of the feed.
// project (optional): Project ID or project name
func (client Client) GetFeedRetentionPolicies(ctx context.Context, feedId *string, project *string) (*FeedRetentionPolicy, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    locationId, _ := uuid.Parse("ed52a011-0112-45b5-9f9e-e14efffb3193")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeedRetentionPolicy
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Set the retention policy for a feed.
// ctx
// policy (required): Feed retention policy.
// feedId (required): Name or ID of the feed.
// project (optional): Project ID or project name
func (client Client) SetFeedRetentionPolicies(ctx context.Context, policy *FeedRetentionPolicy, feedId *string, project *string) (*FeedRetentionPolicy, error) {
    if policy == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "policy"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*policy)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ed52a011-0112-45b5-9f9e-e14efffb3193")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeedRetentionPolicy
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// packageVersionIdQuery (required)
// feedId (required)
// packageId (required)
// project (optional): Project ID or project name
func (client Client) QueryPackageVersionMetrics(ctx context.Context, packageVersionIdQuery *PackageVersionMetricsQuery, feedId *string, packageId *uuid.UUID, project *string) (*[]PackageVersionMetrics, error) {
    if packageVersionIdQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageVersionIdQuery"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = (*packageId).String()

    body, marshalErr := json.Marshal(*packageVersionIdQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("e6ae8caa-b6a8-4809-b840-91b2a42c19ad")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PackageVersionMetrics
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get details about a specific package version.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): Id of the package (GUID Id, not name).
// packageVersionId (required): Id of the package version (GUID Id, not name).
// project (optional): Project ID or project name
// includeUrls (optional): True to include urls for each version. Default is true.
// isListed (optional): Only applicable for NuGet packages. If false, delisted package versions will be returned.
// isDeleted (optional): This does not have any effect on the requested package version, for other versions returned specifies whether to return only deleted or non-deleted versions of packages in the response. Default is unset (return all versions).
func (client Client) GetPackageVersion(ctx context.Context, feedId *string, packageId *string, packageVersionId *string, project *string, includeUrls *bool, isListed *bool, isDeleted *bool) (*PackageVersion, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil || *packageId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = *packageId
    if packageVersionId == nil || *packageVersionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersionId"} 
    }
    routeValues["packageVersionId"] = *packageVersionId

    queryParams := url.Values{}
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    if isListed != nil {
        queryParams.Add("isListed", strconv.FormatBool(*isListed))
    }
    if isDeleted != nil {
        queryParams.Add("isDeleted", strconv.FormatBool(*isDeleted))
    }
    locationId, _ := uuid.Parse("3b331909-6a86-44cc-b9ec-c1834c35498f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PackageVersion
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of package versions, optionally filtering by state.
// ctx
// feedId (required): Name or Id of the feed.
// packageId (required): Id of the package (GUID Id, not name).
// project (optional): Project ID or project name
// includeUrls (optional): True to include urls for each version. Default is true.
// isListed (optional): Only applicable for NuGet packages. If false, delisted package versions will be returned.
// isDeleted (optional): If set specifies whether to return only deleted or non-deleted versions of packages in the response. Default is unset (return all versions).
func (client Client) GetPackageVersions(ctx context.Context, feedId *string, packageId *string, project *string, includeUrls *bool, isListed *bool, isDeleted *bool) (*[]PackageVersion, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if packageId == nil || *packageId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageId"} 
    }
    routeValues["packageId"] = *packageId

    queryParams := url.Values{}
    if includeUrls != nil {
        queryParams.Add("includeUrls", strconv.FormatBool(*includeUrls))
    }
    if isListed != nil {
        queryParams.Add("isListed", strconv.FormatBool(*isListed))
    }
    if isDeleted != nil {
        queryParams.Add("isDeleted", strconv.FormatBool(*isDeleted))
    }
    locationId, _ := uuid.Parse("3b331909-6a86-44cc-b9ec-c1834c35498f")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []PackageVersion
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a new view on the referenced feed.
// ctx
// view (required): View to be created.
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
func (client Client) CreateFeedView(ctx context.Context, view *FeedView, feedId *string, project *string) (*FeedView, error) {
    if view == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "view"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    body, marshalErr := json.Marshal(*view)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("42a8502a-6785-41bc-8c16-89477d930877")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeedView
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a feed view.
// ctx
// feedId (required): Name or Id of the feed.
// viewId (required): Name or Id of the view.
// project (optional): Project ID or project name
func (client Client) DeleteFeedView(ctx context.Context, feedId *string, viewId *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if viewId == nil || *viewId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "viewId"} 
    }
    routeValues["viewId"] = *viewId

    locationId, _ := uuid.Parse("42a8502a-6785-41bc-8c16-89477d930877")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a view by Id.
// ctx
// feedId (required): Name or Id of the feed.
// viewId (required): Name or Id of the view.
// project (optional): Project ID or project name
func (client Client) GetFeedView(ctx context.Context, feedId *string, viewId *string, project *string) (*FeedView, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if viewId == nil || *viewId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "viewId"} 
    }
    routeValues["viewId"] = *viewId

    locationId, _ := uuid.Parse("42a8502a-6785-41bc-8c16-89477d930877")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeedView
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all views for a feed.
// ctx
// feedId (required): Name or Id of the feed.
// project (optional): Project ID or project name
func (client Client) GetFeedViews(ctx context.Context, feedId *string, project *string) (*[]FeedView, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId

    locationId, _ := uuid.Parse("42a8502a-6785-41bc-8c16-89477d930877")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []FeedView
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update a view.
// ctx
// view (required): New settings to apply to the specified view.
// feedId (required): Name or Id of the feed.
// viewId (required): Name or Id of the view.
// project (optional): Project ID or project name
func (client Client) UpdateFeedView(ctx context.Context, view *FeedView, feedId *string, viewId *string, project *string) (*FeedView, error) {
    if view == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "view"}
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *feedId
    if viewId == nil || *viewId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "viewId"} 
    }
    routeValues["viewId"] = *viewId

    body, marshalErr := json.Marshal(*view)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("42a8502a-6785-41bc-8c16-89477d930877")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue FeedView
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

