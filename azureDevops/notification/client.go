// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package notification

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strings"
    "time"
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

// List diagnostic logs this service.
// source (required)
// entryId (optional)
// startTime (optional)
// endTime (optional)
func (client Client) ListLogs(source *uuid.UUID, entryId *uuid.UUID, startTime *time.Time, endTime *time.Time) (*[]INotificationDiagnosticLog, error) {
    routeValues := make(map[string]string)
    if source == nil {
        return nil, errors.New("source is a required parameter")
    }
    routeValues["source"] = (*source).String()
    if entryId != nil {
        routeValues["entryId"] = (*entryId).String()
    }

    queryParams := url.Values{}
    if startTime != nil {
        queryParams.Add("startTime", (*startTime).String())
    }
    if endTime != nil {
        queryParams.Add("endTime", (*endTime).String())
    }
    locationId, _ := uuid.Parse("991842f3-eb16-4aea-ac81-81353ef2b75c")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []INotificationDiagnosticLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// subscriptionId (required)
func (client Client) GetSubscriptionDiagnostics(subscriptionId *string) (*SubscriptionDiagnostics, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, errors.New("subscriptionId is a required parameter")
    }
    routeValues["subscriptionId"] = *subscriptionId

    locationId, _ := uuid.Parse("20f1929d-4be7-4c2e-a74e-d47640ff3418")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionDiagnostics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// updateParameters (required)
// subscriptionId (required)
func (client Client) UpdateSubscriptionDiagnostics(updateParameters *UpdateSubscripitonDiagnosticsParameters, subscriptionId *string) (*SubscriptionDiagnostics, error) {
    if updateParameters == nil {
        return nil, errors.New("updateParameters is a required parameter")
    }
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, errors.New("subscriptionId is a required parameter")
    }
    routeValues["subscriptionId"] = *subscriptionId

    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("20f1929d-4be7-4c2e-a74e-d47640ff3418")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionDiagnostics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a specific event type.
// eventType (required)
func (client Client) GetEventType(eventType *string) (*NotificationEventType, error) {
    routeValues := make(map[string]string)
    if eventType == nil || *eventType == "" {
        return nil, errors.New("eventType is a required parameter")
    }
    routeValues["eventType"] = *eventType

    locationId, _ := uuid.Parse("cc84fb5f-6247-4c7a-aeae-e5a3c3fddb21")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationEventType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// List available event types for this service. Optionally filter by only event types for the specified publisher.
// publisherId (optional): Limit to event types for this publisher
func (client Client) ListEventTypes(publisherId *string) (*[]NotificationEventType, error) {
    queryParams := url.Values{}
    if publisherId != nil {
        queryParams.Add("publisherId", *publisherId)
    }
    locationId, _ := uuid.Parse("cc84fb5f-6247-4c7a-aeae-e5a3c3fddb21")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationEventType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

func (client Client) GetSettings() (*NotificationAdminSettings, error) {
    locationId, _ := uuid.Parse("cbe076d8-2803-45ff-8d8d-44653686ea2a")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationAdminSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// updateParameters (required)
func (client Client) UpdateSettings(updateParameters *NotificationAdminSettingsUpdateParameters) (*NotificationAdminSettings, error) {
    if updateParameters == nil {
        return nil, errors.New("updateParameters is a required parameter")
    }
    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cbe076d8-2803-45ff-8d8d-44653686ea2a")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationAdminSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// subscriberId (required)
func (client Client) GetSubscriber(subscriberId *uuid.UUID) (*NotificationSubscriber, error) {
    routeValues := make(map[string]string)
    if subscriberId == nil {
        return nil, errors.New("subscriberId is a required parameter")
    }
    routeValues["subscriberId"] = (*subscriberId).String()

    locationId, _ := uuid.Parse("4d5caff1-25ba-430b-b808-7a1f352cc197")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscriber
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// updateParameters (required)
// subscriberId (required)
func (client Client) UpdateSubscriber(updateParameters *NotificationSubscriberUpdateParameters, subscriberId *uuid.UUID) (*NotificationSubscriber, error) {
    if updateParameters == nil {
        return nil, errors.New("updateParameters is a required parameter")
    }
    routeValues := make(map[string]string)
    if subscriberId == nil {
        return nil, errors.New("subscriberId is a required parameter")
    }
    routeValues["subscriberId"] = (*subscriberId).String()

    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4d5caff1-25ba-430b-b808-7a1f352cc197")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscriber
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Query for subscriptions. A subscription is returned if it matches one or more of the specified conditions.
// subscriptionQuery (required)
func (client Client) QuerySubscriptions(subscriptionQuery *SubscriptionQuery) (*[]NotificationSubscription, error) {
    if subscriptionQuery == nil {
        return nil, errors.New("subscriptionQuery is a required parameter")
    }
    body, marshalErr := json.Marshal(*subscriptionQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6864db85-08c0-4006-8e8e-cc1bebe31675")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationSubscription
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a new subscription.
// createParameters (required)
func (client Client) CreateSubscription(createParameters *NotificationSubscriptionCreateParameters) (*NotificationSubscription, error) {
    if createParameters == nil {
        return nil, errors.New("createParameters is a required parameter")
    }
    body, marshalErr := json.Marshal(*createParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a subscription.
// subscriptionId (required)
func (client Client) DeleteSubscription(subscriptionId *string) error {
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return errors.New("subscriptionId is a required parameter")
    }
    routeValues["subscriptionId"] = *subscriptionId

    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a notification subscription by its ID.
// subscriptionId (required)
// queryFlags (optional)
func (client Client) GetSubscription(subscriptionId *string, queryFlags *string) (*NotificationSubscription, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, errors.New("subscriptionId is a required parameter")
    }
    routeValues["subscriptionId"] = *subscriptionId

    queryParams := url.Values{}
    if queryFlags != nil {
        queryParams.Add("queryFlags", *queryFlags)
    }
    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of notification subscriptions, either by subscription IDs or by all subscriptions for a given user or group.
// targetId (optional): User or Group ID
// ids (optional): List of subscription IDs
// queryFlags (optional)
func (client Client) ListSubscriptions(targetId *uuid.UUID, ids *[]string, queryFlags *string) (*[]NotificationSubscription, error) {
    queryParams := url.Values{}
    if targetId != nil {
        queryParams.Add("targetId", (*targetId).String())
    }
    if ids != nil {
        listAsString := strings.Join((*ids)[:], ",")
        queryParams.Add("ids", listAsString)
    }
    if queryFlags != nil {
        queryParams.Add("queryFlags", *queryFlags)
    }
    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationSubscription
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update an existing subscription. Depending on the type of subscription and permissions, the caller can update the description, filter settings, channel (delivery) settings and more.
// updateParameters (required)
// subscriptionId (required)
func (client Client) UpdateSubscription(updateParameters *NotificationSubscriptionUpdateParameters, subscriptionId *string) (*NotificationSubscription, error) {
    if updateParameters == nil {
        return nil, errors.New("updateParameters is a required parameter")
    }
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, errors.New("subscriptionId is a required parameter")
    }
    routeValues["subscriptionId"] = *subscriptionId

    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get available subscription templates.
func (client Client) GetSubscriptionTemplates() (*[]NotificationSubscriptionTemplate, error) {
    locationId, _ := uuid.Parse("fa5d24ba-7484-4f3d-888d-4ec6b1974082")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationSubscriptionTemplate
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update the specified user's settings for the specified subscription. This API is typically used to opt in or out of a shared subscription. User settings can only be applied to shared subscriptions, like team subscriptions or default subscriptions.
// userSettings (required)
// subscriptionId (required)
// userId (required): ID of the user
func (client Client) UpdateSubscriptionUserSettings(userSettings *SubscriptionUserSettings, subscriptionId *string, userId *uuid.UUID) (*SubscriptionUserSettings, error) {
    if userSettings == nil {
        return nil, errors.New("userSettings is a required parameter")
    }
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, errors.New("subscriptionId is a required parameter")
    }
    routeValues["subscriptionId"] = *subscriptionId
    if userId == nil {
        return nil, errors.New("userId is a required parameter")
    }
    routeValues["userId"] = (*userId).String()

    body, marshalErr := json.Marshal(*userSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ed5a3dff-aeb5-41b1-b4f7-89e66e58b62e")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionUserSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

