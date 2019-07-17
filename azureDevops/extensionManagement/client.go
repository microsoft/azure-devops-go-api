// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package extensionManagement

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
    "strings"
)

var ResourceAreaId, _ = uuid.Parse("6c2b0933-3600-42ae-bf8b-93d4f7e83594")

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

// [Preview API] List the installed extensions in the account / project collection.
// includeDisabledExtensions (optional): If true (the default), include disabled extensions in the results.
// includeErrors (optional): If true, include installed extensions with errors.
// assetTypes (optional)
// includeInstallationIssues (optional)
func (client Client) GetInstalledExtensions(includeDisabledExtensions *bool, includeErrors *bool, assetTypes *[]string, includeInstallationIssues *bool) (*[]InstalledExtension, error) {
    queryParams := url.Values{}
    if includeDisabledExtensions != nil {
        queryParams.Add("includeDisabledExtensions", strconv.FormatBool(*includeDisabledExtensions))
    }
    if includeErrors != nil {
        queryParams.Add("includeErrors", strconv.FormatBool(*includeErrors))
    }
    if assetTypes != nil {
        listAsString := strings.Join((*assetTypes)[:], ":")
        queryParams.Add("assetTypes", listAsString)
    }
    if includeInstallationIssues != nil {
        queryParams.Add("includeInstallationIssues", strconv.FormatBool(*includeInstallationIssues))
    }
    locationId, _ := uuid.Parse("275424d0-c844-4fe2-bda6-04933a1357d8")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []InstalledExtension
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update an installed extension. Typically this API is used to enable or disable an extension.
// extension (required)
func (client Client) UpdateInstalledExtension(extension *InstalledExtension) (*InstalledExtension, error) {
    if extension == nil {
        return nil, errors.New("extension is a required parameter")
    }
    body, marshalErr := json.Marshal(*extension)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("275424d0-c844-4fe2-bda6-04933a1357d8")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue InstalledExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get an installed extension by its publisher and extension name.
// publisherName (required): Name of the publisher. Example: "fabrikam".
// extensionName (required): Name of the extension. Example: "ops-tools".
// assetTypes (optional)
func (client Client) GetInstalledExtensionByName(publisherName *string, extensionName *string, assetTypes *[]string) (*InstalledExtension, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, errors.New("publisherName is a required parameter")
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, errors.New("extensionName is a required parameter")
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if assetTypes != nil {
        listAsString := strings.Join((*assetTypes)[:], ":")
        queryParams.Add("assetTypes", listAsString)
    }
    locationId, _ := uuid.Parse("fb0da285-f23e-4b56-8b53-3ef5f9f6de66")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue InstalledExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Install the specified extension into the account / project collection.
// publisherName (required): Name of the publisher. Example: "fabrikam".
// extensionName (required): Name of the extension. Example: "ops-tools".
// version (optional)
func (client Client) InstallExtensionByName(publisherName *string, extensionName *string, version *string) (*InstalledExtension, error) {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return nil, errors.New("publisherName is a required parameter")
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return nil, errors.New("extensionName is a required parameter")
    }
    routeValues["extensionName"] = *extensionName
    if version != nil && *version != "" {
        routeValues["version"] = *version
    }

    locationId, _ := uuid.Parse("fb0da285-f23e-4b56-8b53-3ef5f9f6de66")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue InstalledExtension
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Uninstall the specified extension from the account / project collection.
// publisherName (required): Name of the publisher. Example: "fabrikam".
// extensionName (required): Name of the extension. Example: "ops-tools".
// reason (optional)
// reasonCode (optional)
func (client Client) UninstallExtensionByName(publisherName *string, extensionName *string, reason *string, reasonCode *string) error {
    routeValues := make(map[string]string)
    if publisherName == nil || *publisherName == "" {
        return errors.New("publisherName is a required parameter")
    }
    routeValues["publisherName"] = *publisherName
    if extensionName == nil || *extensionName == "" {
        return errors.New("extensionName is a required parameter")
    }
    routeValues["extensionName"] = *extensionName

    queryParams := url.Values{}
    if reason != nil {
        queryParams.Add("reason", *reason)
    }
    if reasonCode != nil {
        queryParams.Add("reasonCode", *reasonCode)
    }
    locationId, _ := uuid.Parse("fb0da285-f23e-4b56-8b53-3ef5f9f6de66")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

