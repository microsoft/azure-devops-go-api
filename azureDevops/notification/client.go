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
    "context"
    "encoding/json"
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

func NewClient(ctx context.Context, connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// List diagnostic logs this service.
// ctx
// source (required)
// entryId (optional)
// startTime (optional)
// endTime (optional)
func (client Client) ListLogs(ctx context.Context, source *uuid.UUID, entryId *uuid.UUID, startTime *time.Time, endTime *time.Time) (*[]INotificationDiagnosticLog, error) {
    routeValues := make(map[string]string)
    if source == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "source"} 
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
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []INotificationDiagnosticLog
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// subscriptionId (required)
func (client Client) GetSubscriptionDiagnostics(ctx context.Context, subscriptionId *string) (*SubscriptionDiagnostics, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *subscriptionId

    locationId, _ := uuid.Parse("20f1929d-4be7-4c2e-a74e-d47640ff3418")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionDiagnostics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// updateParameters (required)
// subscriptionId (required)
func (client Client) UpdateSubscriptionDiagnostics(ctx context.Context, updateParameters *UpdateSubscripitonDiagnosticsParameters, subscriptionId *string) (*SubscriptionDiagnostics, error) {
    if updateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *subscriptionId

    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("20f1929d-4be7-4c2e-a74e-d47640ff3418")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionDiagnostics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a specific event type.
// ctx
// eventType (required)
func (client Client) GetEventType(ctx context.Context, eventType *string) (*NotificationEventType, error) {
    routeValues := make(map[string]string)
    if eventType == nil || *eventType == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "eventType"} 
    }
    routeValues["eventType"] = *eventType

    locationId, _ := uuid.Parse("cc84fb5f-6247-4c7a-aeae-e5a3c3fddb21")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationEventType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// List available event types for this service. Optionally filter by only event types for the specified publisher.
// ctx
// publisherId (optional): Limit to event types for this publisher
func (client Client) ListEventTypes(ctx context.Context, publisherId *string) (*[]NotificationEventType, error) {
    queryParams := url.Values{}
    if publisherId != nil {
        queryParams.Add("publisherId", *publisherId)
    }
    locationId, _ := uuid.Parse("cc84fb5f-6247-4c7a-aeae-e5a3c3fddb21")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationEventType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
func (client Client) GetSettings(ctx context.Context, ) (*NotificationAdminSettings, error) {
    locationId, _ := uuid.Parse("cbe076d8-2803-45ff-8d8d-44653686ea2a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationAdminSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// updateParameters (required)
func (client Client) UpdateSettings(ctx context.Context, updateParameters *NotificationAdminSettingsUpdateParameters) (*NotificationAdminSettings, error) {
    if updateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cbe076d8-2803-45ff-8d8d-44653686ea2a")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationAdminSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// subscriberId (required)
func (client Client) GetSubscriber(ctx context.Context, subscriberId *uuid.UUID) (*NotificationSubscriber, error) {
    routeValues := make(map[string]string)
    if subscriberId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscriberId"} 
    }
    routeValues["subscriberId"] = (*subscriberId).String()

    locationId, _ := uuid.Parse("4d5caff1-25ba-430b-b808-7a1f352cc197")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscriber
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// updateParameters (required)
// subscriberId (required)
func (client Client) UpdateSubscriber(ctx context.Context, updateParameters *NotificationSubscriberUpdateParameters, subscriberId *uuid.UUID) (*NotificationSubscriber, error) {
    if updateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    routeValues := make(map[string]string)
    if subscriberId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscriberId"} 
    }
    routeValues["subscriberId"] = (*subscriberId).String()

    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4d5caff1-25ba-430b-b808-7a1f352cc197")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscriber
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Query for subscriptions. A subscription is returned if it matches one or more of the specified conditions.
// ctx
// subscriptionQuery (required)
func (client Client) QuerySubscriptions(ctx context.Context, subscriptionQuery *SubscriptionQuery) (*[]NotificationSubscription, error) {
    if subscriptionQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscriptionQuery"}
    }
    body, marshalErr := json.Marshal(*subscriptionQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("6864db85-08c0-4006-8e8e-cc1bebe31675")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationSubscription
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Create a new subscription.
// ctx
// createParameters (required)
func (client Client) CreateSubscription(ctx context.Context, createParameters *NotificationSubscriptionCreateParameters) (*NotificationSubscription, error) {
    if createParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "createParameters"}
    }
    body, marshalErr := json.Marshal(*createParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a subscription.
// ctx
// subscriptionId (required)
func (client Client) DeleteSubscription(ctx context.Context, subscriptionId *string) error {
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *subscriptionId

    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a notification subscription by its ID.
// ctx
// subscriptionId (required)
// queryFlags (optional)
func (client Client) GetSubscription(ctx context.Context, subscriptionId *string, queryFlags *SubscriptionQueryFlags) (*NotificationSubscription, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *subscriptionId

    queryParams := url.Values{}
    if queryFlags != nil {
        queryParams.Add("queryFlags", *queryFlags)
    }
    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of notification subscriptions, either by subscription IDs or by all subscriptions for a given user or group.
// ctx
// targetId (optional): User or Group ID
// ids (optional): List of subscription IDs
// queryFlags (optional)
func (client Client) ListSubscriptions(ctx context.Context, targetId *uuid.UUID, ids *[]string, queryFlags *SubscriptionQueryFlags) (*[]NotificationSubscription, error) {
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
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationSubscription
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update an existing subscription. Depending on the type of subscription and permissions, the caller can update the description, filter settings, channel (delivery) settings and more.
// ctx
// updateParameters (required)
// subscriptionId (required)
func (client Client) UpdateSubscription(ctx context.Context, updateParameters *NotificationSubscriptionUpdateParameters, subscriptionId *string) (*NotificationSubscription, error) {
    if updateParameters == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *subscriptionId

    body, marshalErr := json.Marshal(*updateParameters)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get available subscription templates.
// ctx
func (client Client) GetSubscriptionTemplates(ctx context.Context, ) (*[]NotificationSubscriptionTemplate, error) {
    locationId, _ := uuid.Parse("fa5d24ba-7484-4f3d-888d-4ec6b1974082")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationSubscriptionTemplate
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update the specified user's settings for the specified subscription. This API is typically used to opt in or out of a shared subscription. User settings can only be applied to shared subscriptions, like team subscriptions or default subscriptions.
// ctx
// userSettings (required)
// subscriptionId (required)
// userId (required): ID of the user
func (client Client) UpdateSubscriptionUserSettings(ctx context.Context, userSettings *SubscriptionUserSettings, subscriptionId *string, userId *uuid.UUID) (*SubscriptionUserSettings, error) {
    if userSettings == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "userSettings"}
    }
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *subscriptionId
    if userId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "userId"} 
    }
    routeValues["userId"] = (*userId).String()

    body, marshalErr := json.Marshal(*userSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ed5a3dff-aeb5-41b1-b4f7-89e66e58b62e")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionUserSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

