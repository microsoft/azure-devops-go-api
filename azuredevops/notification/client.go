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
    "github.com/microsoft/azure-devops-go-api/azuredevops"
    "net/http"
    "net/url"
    "strings"
    "time"
)

type Client struct {
    Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client{
        Client: *client,
    }
}

// List diagnostic logs this service.
func (client *Client) ListLogs(ctx context.Context, args ListLogsArgs) (*[]INotificationDiagnosticLog, error) {
    routeValues := make(map[string]string)
    if args.Source == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "source"} 
    }
    routeValues["source"] = (*args.Source).String()
    if args.EntryId != nil {
        routeValues["entryId"] = (*args.EntryId).String()
    }

    queryParams := url.Values{}
    if args.StartTime != nil {
        queryParams.Add("startTime", (*args.StartTime).String())
    }
    if args.EndTime != nil {
        queryParams.Add("endTime", (*args.EndTime).String())
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

// Arguments for the ListLogs function
type ListLogsArgs struct {
    // (required)
    Source *uuid.UUID
    // (optional)
    EntryId *uuid.UUID
    // (optional)
    StartTime *time.Time
    // (optional)
    EndTime *time.Time
}

func (client *Client) GetSubscriptionDiagnostics(ctx context.Context, args GetSubscriptionDiagnosticsArgs) (*SubscriptionDiagnostics, error) {
    routeValues := make(map[string]string)
    if args.SubscriptionId == nil || *args.SubscriptionId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *args.SubscriptionId

    locationId, _ := uuid.Parse("20f1929d-4be7-4c2e-a74e-d47640ff3418")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionDiagnostics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetSubscriptionDiagnostics function
type GetSubscriptionDiagnosticsArgs struct {
    // (required)
    SubscriptionId *string
}

func (client *Client) UpdateSubscriptionDiagnostics(ctx context.Context, args UpdateSubscriptionDiagnosticsArgs) (*SubscriptionDiagnostics, error) {
    if args.UpdateParameters == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    routeValues := make(map[string]string)
    if args.SubscriptionId == nil || *args.SubscriptionId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *args.SubscriptionId

    body, marshalErr := json.Marshal(*args.UpdateParameters)
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

// Arguments for the UpdateSubscriptionDiagnostics function
type UpdateSubscriptionDiagnosticsArgs struct {
    // (required)
    UpdateParameters *UpdateSubscripitonDiagnosticsParameters
    // (required)
    SubscriptionId *string
}

// Get a specific event type.
func (client *Client) GetEventType(ctx context.Context, args GetEventTypeArgs) (*NotificationEventType, error) {
    routeValues := make(map[string]string)
    if args.EventType == nil || *args.EventType == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "eventType"} 
    }
    routeValues["eventType"] = *args.EventType

    locationId, _ := uuid.Parse("cc84fb5f-6247-4c7a-aeae-e5a3c3fddb21")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationEventType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetEventType function
type GetEventTypeArgs struct {
    // (required)
    EventType *string
}

// List available event types for this service. Optionally filter by only event types for the specified publisher.
func (client *Client) ListEventTypes(ctx context.Context, args ListEventTypesArgs) (*[]NotificationEventType, error) {
    queryParams := url.Values{}
    if args.PublisherId != nil {
        queryParams.Add("publisherId", *args.PublisherId)
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

// Arguments for the ListEventTypes function
type ListEventTypesArgs struct {
    // (optional) Limit to event types for this publisher
    PublisherId *string
}

func (client *Client) GetSettings(ctx context.Context, args GetSettingsArgs) (*NotificationAdminSettings, error) {
    locationId, _ := uuid.Parse("cbe076d8-2803-45ff-8d8d-44653686ea2a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationAdminSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetSettings function
type GetSettingsArgs struct {
}

func (client *Client) UpdateSettings(ctx context.Context, args UpdateSettingsArgs) (*NotificationAdminSettings, error) {
    if args.UpdateParameters == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    body, marshalErr := json.Marshal(*args.UpdateParameters)
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

// Arguments for the UpdateSettings function
type UpdateSettingsArgs struct {
    // (required)
    UpdateParameters *NotificationAdminSettingsUpdateParameters
}

func (client *Client) GetSubscriber(ctx context.Context, args GetSubscriberArgs) (*NotificationSubscriber, error) {
    routeValues := make(map[string]string)
    if args.SubscriberId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "subscriberId"} 
    }
    routeValues["subscriberId"] = (*args.SubscriberId).String()

    locationId, _ := uuid.Parse("4d5caff1-25ba-430b-b808-7a1f352cc197")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationSubscriber
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetSubscriber function
type GetSubscriberArgs struct {
    // (required)
    SubscriberId *uuid.UUID
}

func (client *Client) UpdateSubscriber(ctx context.Context, args UpdateSubscriberArgs) (*NotificationSubscriber, error) {
    if args.UpdateParameters == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    routeValues := make(map[string]string)
    if args.SubscriberId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "subscriberId"} 
    }
    routeValues["subscriberId"] = (*args.SubscriberId).String()

    body, marshalErr := json.Marshal(*args.UpdateParameters)
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

// Arguments for the UpdateSubscriber function
type UpdateSubscriberArgs struct {
    // (required)
    UpdateParameters *NotificationSubscriberUpdateParameters
    // (required)
    SubscriberId *uuid.UUID
}

// Query for subscriptions. A subscription is returned if it matches one or more of the specified conditions.
func (client *Client) QuerySubscriptions(ctx context.Context, args QuerySubscriptionsArgs) (*[]NotificationSubscription, error) {
    if args.SubscriptionQuery == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "subscriptionQuery"}
    }
    body, marshalErr := json.Marshal(*args.SubscriptionQuery)
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

// Arguments for the QuerySubscriptions function
type QuerySubscriptionsArgs struct {
    // (required)
    SubscriptionQuery *SubscriptionQuery
}

// Create a new subscription.
func (client *Client) CreateSubscription(ctx context.Context, args CreateSubscriptionArgs) (*NotificationSubscription, error) {
    if args.CreateParameters == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "createParameters"}
    }
    body, marshalErr := json.Marshal(*args.CreateParameters)
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

// Arguments for the CreateSubscription function
type CreateSubscriptionArgs struct {
    // (required)
    CreateParameters *NotificationSubscriptionCreateParameters
}

// Delete a subscription.
func (client *Client) DeleteSubscription(ctx context.Context, args DeleteSubscriptionArgs) error {
    routeValues := make(map[string]string)
    if args.SubscriptionId == nil || *args.SubscriptionId == "" {
        return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *args.SubscriptionId

    locationId, _ := uuid.Parse("70f911d6-abac-488c-85b3-a206bf57e165")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteSubscription function
type DeleteSubscriptionArgs struct {
    // (required)
    SubscriptionId *string
}

// Get a notification subscription by its ID.
func (client *Client) GetSubscription(ctx context.Context, args GetSubscriptionArgs) (*NotificationSubscription, error) {
    routeValues := make(map[string]string)
    if args.SubscriptionId == nil || *args.SubscriptionId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *args.SubscriptionId

    queryParams := url.Values{}
    if args.QueryFlags != nil {
        queryParams.Add("queryFlags", string(*args.QueryFlags))
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

// Arguments for the GetSubscription function
type GetSubscriptionArgs struct {
    // (required)
    SubscriptionId *string
    // (optional)
    QueryFlags *SubscriptionQueryFlags
}

// Get a list of notification subscriptions, either by subscription IDs or by all subscriptions for a given user or group.
func (client *Client) ListSubscriptions(ctx context.Context, args ListSubscriptionsArgs) (*[]NotificationSubscription, error) {
    queryParams := url.Values{}
    if args.TargetId != nil {
        queryParams.Add("targetId", (*args.TargetId).String())
    }
    if args.Ids != nil {
        listAsString := strings.Join((*args.Ids)[:], ",")
        queryParams.Add("ids", listAsString)
    }
    if args.QueryFlags != nil {
        queryParams.Add("queryFlags", string(*args.QueryFlags))
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

// Arguments for the ListSubscriptions function
type ListSubscriptionsArgs struct {
    // (optional) User or Group ID
    TargetId *uuid.UUID
    // (optional) List of subscription IDs
    Ids *[]string
    // (optional)
    QueryFlags *SubscriptionQueryFlags
}

// Update an existing subscription. Depending on the type of subscription and permissions, the caller can update the description, filter settings, channel (delivery) settings and more.
func (client *Client) UpdateSubscription(ctx context.Context, args UpdateSubscriptionArgs) (*NotificationSubscription, error) {
    if args.UpdateParameters == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "updateParameters"}
    }
    routeValues := make(map[string]string)
    if args.SubscriptionId == nil || *args.SubscriptionId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *args.SubscriptionId

    body, marshalErr := json.Marshal(*args.UpdateParameters)
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

// Arguments for the UpdateSubscription function
type UpdateSubscriptionArgs struct {
    // (required)
    UpdateParameters *NotificationSubscriptionUpdateParameters
    // (required)
    SubscriptionId *string
}

// Get available subscription templates.
func (client *Client) GetSubscriptionTemplates(ctx context.Context, args GetSubscriptionTemplatesArgs) (*[]NotificationSubscriptionTemplate, error) {
    locationId, _ := uuid.Parse("fa5d24ba-7484-4f3d-888d-4ec6b1974082")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []NotificationSubscriptionTemplate
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetSubscriptionTemplates function
type GetSubscriptionTemplatesArgs struct {
}

// Update the specified user's settings for the specified subscription. This API is typically used to opt in or out of a shared subscription. User settings can only be applied to shared subscriptions, like team subscriptions or default subscriptions.
func (client *Client) UpdateSubscriptionUserSettings(ctx context.Context, args UpdateSubscriptionUserSettingsArgs) (*SubscriptionUserSettings, error) {
    if args.UserSettings == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "userSettings"}
    }
    routeValues := make(map[string]string)
    if args.SubscriptionId == nil || *args.SubscriptionId == "" {
        return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *args.SubscriptionId
    if args.UserId == nil {
        return nil, &azuredevops.ArgumentNilError{ArgumentName: "userId"} 
    }
    routeValues["userId"] = (*args.UserId).String()

    body, marshalErr := json.Marshal(*args.UserSettings)
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

// Arguments for the UpdateSubscriptionUserSettings function
type UpdateSubscriptionUserSettingsArgs struct {
    // (required)
    UserSettings *SubscriptionUserSettings
    // (required)
    SubscriptionId *string
    // (required) ID of the user
    UserId *uuid.UUID
}

