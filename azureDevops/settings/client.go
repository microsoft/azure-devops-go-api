// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package settings

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
)

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// [Preview API] Get all setting entries for the given user/all-users scope
// userScope (required): User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
// key (optional): Optional key under which to filter all the entries
func (client Client) GetEntries(userScope *string, key *string) (*map[string]interface{}, error) {
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return nil, errors.New("userScope is a required parameter")
    }
    routeValues["userScope"] = *userScope
    if key != nil && *key != "" {
        routeValues["key"] = *key
    }

    locationId, _ := uuid.Parse("cd006711-163d-4cd4-a597-b05bad2556ff")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]interface{}
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Remove the entry or entries under the specified path
// userScope (required): User-Scope at which to remove the value. Should be "me" for the current user or "host" for all users.
// key (required): Root key of the entry or entries to remove
func (client Client) RemoveEntries(userScope *string, key *string) error {
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return errors.New("userScope is a required parameter")
    }
    routeValues["userScope"] = *userScope
    if key == nil || *key == "" {
        return errors.New("key is a required parameter")
    }
    routeValues["key"] = *key

    locationId, _ := uuid.Parse("cd006711-163d-4cd4-a597-b05bad2556ff")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Set the specified setting entry values for the given user/all-users scope
// entries (required): The entries to set
// userScope (required): User-Scope at which to set the values. Should be "me" for the current user or "host" for all users.
func (client Client) SetEntries(entries *map[string]interface{}, userScope *string) error {
    if entries == nil {
        return errors.New("entries is a required parameter")
    }
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return errors.New("userScope is a required parameter")
    }
    routeValues["userScope"] = *userScope

    body, marshalErr := json.Marshal(*entries)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("cd006711-163d-4cd4-a597-b05bad2556ff")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get all setting entries for the given named scope
// userScope (required): User-Scope at which to get the value. Should be "me" for the current user or "host" for all users.
// scopeName (required): Scope at which to get the setting for (e.g. "project" or "team")
// scopeValue (required): Value of the scope (e.g. the project or team id)
// key (optional): Optional key under which to filter all the entries
func (client Client) GetEntriesForScope(userScope *string, scopeName *string, scopeValue *string, key *string) (*map[string]interface{}, error) {
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return nil, errors.New("userScope is a required parameter")
    }
    routeValues["userScope"] = *userScope
    if scopeName == nil || *scopeName == "" {
        return nil, errors.New("scopeName is a required parameter")
    }
    routeValues["scopeName"] = *scopeName
    if scopeValue == nil || *scopeValue == "" {
        return nil, errors.New("scopeValue is a required parameter")
    }
    routeValues["scopeValue"] = *scopeValue
    if key != nil && *key != "" {
        routeValues["key"] = *key
    }

    locationId, _ := uuid.Parse("4cbaafaf-e8af-4570-98d1-79ee99c56327")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue map[string]interface{}
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Remove the entry or entries under the specified path
// userScope (required): User-Scope at which to remove the value. Should be "me" for the current user or "host" for all users.
// scopeName (required): Scope at which to get the setting for (e.g. "project" or "team")
// scopeValue (required): Value of the scope (e.g. the project or team id)
// key (required): Root key of the entry or entries to remove
func (client Client) RemoveEntriesForScope(userScope *string, scopeName *string, scopeValue *string, key *string) error {
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return errors.New("userScope is a required parameter")
    }
    routeValues["userScope"] = *userScope
    if scopeName == nil || *scopeName == "" {
        return errors.New("scopeName is a required parameter")
    }
    routeValues["scopeName"] = *scopeName
    if scopeValue == nil || *scopeValue == "" {
        return errors.New("scopeValue is a required parameter")
    }
    routeValues["scopeValue"] = *scopeValue
    if key == nil || *key == "" {
        return errors.New("key is a required parameter")
    }
    routeValues["key"] = *key

    locationId, _ := uuid.Parse("4cbaafaf-e8af-4570-98d1-79ee99c56327")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Set the specified entries for the given named scope
// entries (required): The entries to set
// userScope (required): User-Scope at which to set the values. Should be "me" for the current user or "host" for all users.
// scopeName (required): Scope at which to set the settings on (e.g. "project" or "team")
// scopeValue (required): Value of the scope (e.g. the project or team id)
func (client Client) SetEntriesForScope(entries *map[string]interface{}, userScope *string, scopeName *string, scopeValue *string) error {
    if entries == nil {
        return errors.New("entries is a required parameter")
    }
    routeValues := make(map[string]string)
    if userScope == nil || *userScope == "" {
        return errors.New("userScope is a required parameter")
    }
    routeValues["userScope"] = *userScope
    if scopeName == nil || *scopeName == "" {
        return errors.New("scopeName is a required parameter")
    }
    routeValues["scopeName"] = *scopeName
    if scopeValue == nil || *scopeValue == "" {
        return errors.New("scopeValue is a required parameter")
    }
    routeValues["scopeValue"] = *scopeValue

    body, marshalErr := json.Marshal(*entries)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4cbaafaf-e8af-4570-98d1-79ee99c56327")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

