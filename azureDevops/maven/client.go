// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package maven

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

var ResourceAreaId, _ = uuid.Parse("6f7f8c07-ff36-473c-bcf3-bd6cc9b6c066")

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

// [Preview API] Fulfills Maven package file download requests by either returning the URL of the requested package file or, in the case of Azure DevOps Server (OnPrem), returning the content as a stream.
// feedId (required): Name or ID of the feed.
// groupId (required): GroupId of the maven package
// artifactId (required): ArtifactId of the maven package
// version (required): Version of the package
// fileName (required): File name to download
// project (optional): Project ID or project name
func (client Client) DownloadPackage(feedId *string, groupId *string, artifactId *string, version *string, fileName *string, project *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feedId == nil || *feedId == "" {
        return nil, errors.New("feedId is a required parameter")
    }
    routeValues["feedId"] = *feedId
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if artifactId == nil || *artifactId == "" {
        return nil, errors.New("artifactId is a required parameter")
    }
    routeValues["artifactId"] = *artifactId
    if version == nil || *version == "" {
        return nil, errors.New("version is a required parameter")
    }
    routeValues["version"] = *version
    if fileName == nil || *fileName == "" {
        return nil, errors.New("fileName is a required parameter")
    }
    routeValues["fileName"] = *fileName

    locationId, _ := uuid.Parse("c338d4b5-d30a-47e2-95b7-f157ef558833")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Permanently delete a package from a feed's recycle bin.
// feed (required): Name or ID of the feed.
// groupId (required): Group ID of the package.
// artifactId (required): Artifact ID of the package.
// version (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) DeletePackageVersionFromRecycleBin(feed *string, groupId *string, artifactId *string, version *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feed == nil || *feed == "" {
        return errors.New("feed is a required parameter")
    }
    routeValues["feed"] = *feed
    if groupId == nil || *groupId == "" {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if artifactId == nil || *artifactId == "" {
        return errors.New("artifactId is a required parameter")
    }
    routeValues["artifactId"] = *artifactId
    if version == nil || *version == "" {
        return errors.New("version is a required parameter")
    }
    routeValues["version"] = *version

    locationId, _ := uuid.Parse("f67e10eb-1254-4953-add7-d49b83a16c9f")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a package version in the recycle bin.
// feed (required): Name or ID of the feed.
// groupId (required): Group ID of the package.
// artifactId (required): Artifact ID of the package.
// version (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) GetPackageVersionMetadataFromRecycleBin(feed *string, groupId *string, artifactId *string, version *string, project *string) (*MavenPackageVersionDeletionState, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feed == nil || *feed == "" {
        return nil, errors.New("feed is a required parameter")
    }
    routeValues["feed"] = *feed
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if artifactId == nil || *artifactId == "" {
        return nil, errors.New("artifactId is a required parameter")
    }
    routeValues["artifactId"] = *artifactId
    if version == nil || *version == "" {
        return nil, errors.New("version is a required parameter")
    }
    routeValues["version"] = *version

    locationId, _ := uuid.Parse("f67e10eb-1254-4953-add7-d49b83a16c9f")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue MavenPackageVersionDeletionState
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Restore a package version from the recycle bin to its associated feed.
// packageVersionDetails (required): Set the 'Deleted' property to false to restore the package.
// feed (required): Name or ID of the feed.
// groupId (required): Group ID of the package.
// artifactId (required): Artifact ID of the package.
// version (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) RestorePackageVersionFromRecycleBin(packageVersionDetails *MavenRecycleBinPackageVersionDetails, feed *string, groupId *string, artifactId *string, version *string, project *string) error {
    if packageVersionDetails == nil {
        return errors.New("packageVersionDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feed == nil || *feed == "" {
        return errors.New("feed is a required parameter")
    }
    routeValues["feed"] = *feed
    if groupId == nil || *groupId == "" {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if artifactId == nil || *artifactId == "" {
        return errors.New("artifactId is a required parameter")
    }
    routeValues["artifactId"] = *artifactId
    if version == nil || *version == "" {
        return errors.New("version is a required parameter")
    }
    routeValues["version"] = *version

    body, marshalErr := json.Marshal(*packageVersionDetails)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("f67e10eb-1254-4953-add7-d49b83a16c9f")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get information about a package version.
// feed (required): Name or ID of the feed.
// groupId (required): Group ID of the package.
// artifactId (required): Artifact ID of the package.
// version (required): Version of the package.
// project (optional): Project ID or project name
// showDeleted (optional): True to show information for deleted packages.
func (client Client) GetPackageVersion(feed *string, groupId *string, artifactId *string, version *string, project *string, showDeleted *bool) (*Package, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feed == nil || *feed == "" {
        return nil, errors.New("feed is a required parameter")
    }
    routeValues["feed"] = *feed
    if groupId == nil || *groupId == "" {
        return nil, errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if artifactId == nil || *artifactId == "" {
        return nil, errors.New("artifactId is a required parameter")
    }
    routeValues["artifactId"] = *artifactId
    if version == nil || *version == "" {
        return nil, errors.New("version is a required parameter")
    }
    routeValues["version"] = *version

    queryParams := url.Values{}
    if showDeleted != nil {
        queryParams.Add("showDeleted", strconv.FormatBool(*showDeleted))
    }
    locationId, _ := uuid.Parse("180ed967-377a-4112-986b-607adb14ded4")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Package
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a package version from the feed and move it to the feed's recycle bin.
// feed (required): Name or ID of the feed.
// groupId (required): Group ID of the package.
// artifactId (required): Artifact ID of the package.
// version (required): Version of the package.
// project (optional): Project ID or project name
func (client Client) PackageDelete(feed *string, groupId *string, artifactId *string, version *string, project *string) error {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if feed == nil || *feed == "" {
        return errors.New("feed is a required parameter")
    }
    routeValues["feed"] = *feed
    if groupId == nil || *groupId == "" {
        return errors.New("groupId is a required parameter")
    }
    routeValues["groupId"] = *groupId
    if artifactId == nil || *artifactId == "" {
        return errors.New("artifactId is a required parameter")
    }
    routeValues["artifactId"] = *artifactId
    if version == nil || *version == "" {
        return errors.New("version is a required parameter")
    }
    routeValues["version"] = *version

    locationId, _ := uuid.Parse("180ed967-377a-4112-986b-607adb14ded4")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

