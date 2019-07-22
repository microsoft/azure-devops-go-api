// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package serviceHooks

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

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) *Client {
    client := connection.GetClientByUrl(connection.BaseUrl)
    return &Client {
        Client: *client,
    }
}

// Get details about a specific consumer action.
// ctx
// consumerId (required): ID for a consumer.
// consumerActionId (required): ID for a consumerActionId.
// publisherId (optional)
func (client Client) GetConsumerAction(ctx context.Context, consumerId *string, consumerActionId *string, publisherId *string) (*ConsumerAction, error) {
    routeValues := make(map[string]string)
    if consumerId == nil || *consumerId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "consumerId"} 
    }
    routeValues["consumerId"] = *consumerId
    if consumerActionId == nil || *consumerActionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "consumerActionId"} 
    }
    routeValues["consumerActionId"] = *consumerActionId

    queryParams := url.Values{}
    if publisherId != nil {
        queryParams.Add("publisherId", *publisherId)
    }
    locationId, _ := uuid.Parse("c3428e90-7a69-4194-8ed8-0f153185ee0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ConsumerAction
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of consumer actions for a specific consumer.
// ctx
// consumerId (required): ID for a consumer.
// publisherId (optional)
func (client Client) ListConsumerActions(ctx context.Context, consumerId *string, publisherId *string) (*[]ConsumerAction, error) {
    routeValues := make(map[string]string)
    if consumerId == nil || *consumerId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "consumerId"} 
    }
    routeValues["consumerId"] = *consumerId

    queryParams := url.Values{}
    if publisherId != nil {
        queryParams.Add("publisherId", *publisherId)
    }
    locationId, _ := uuid.Parse("c3428e90-7a69-4194-8ed8-0f153185ee0d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ConsumerAction
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a specific consumer service. Optionally filter out consumer actions that do not support any event types for the specified publisher.
// ctx
// consumerId (required): ID for a consumer.
// publisherId (optional)
func (client Client) GetConsumer(ctx context.Context, consumerId *string, publisherId *string) (*Consumer, error) {
    routeValues := make(map[string]string)
    if consumerId == nil || *consumerId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "consumerId"} 
    }
    routeValues["consumerId"] = *consumerId

    queryParams := url.Values{}
    if publisherId != nil {
        queryParams.Add("publisherId", *publisherId)
    }
    locationId, _ := uuid.Parse("4301c514-5f34-4f5d-a145-f0ea7b5b7d19")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Consumer
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of available service hook consumer services. Optionally filter by consumers that support at least one event type from the specific publisher.
// ctx
// publisherId (optional)
func (client Client) ListConsumers(ctx context.Context, publisherId *string) (*[]Consumer, error) {
    queryParams := url.Values{}
    if publisherId != nil {
        queryParams.Add("publisherId", *publisherId)
    }
    locationId, _ := uuid.Parse("4301c514-5f34-4f5d-a145-f0ea7b5b7d19")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Consumer
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
// ctx
// subscriptionId (required)
func (client Client) GetSubscriptionDiagnostics(ctx context.Context, subscriptionId *string) (*SubscriptionDiagnostics, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil || *subscriptionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = *subscriptionId

    locationId, _ := uuid.Parse("3b36bcb5-02ad-43c6-bbfa-6dfc6f8e9d68")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionDiagnostics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API]
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
    locationId, _ := uuid.Parse("3b36bcb5-02ad-43c6-bbfa-6dfc6f8e9d68")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionDiagnostics
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a specific event type.
// ctx
// publisherId (required): ID for a publisher.
// eventTypeId (required)
func (client Client) GetEventType(ctx context.Context, publisherId *string, eventTypeId *string) (*EventTypeDescriptor, error) {
    routeValues := make(map[string]string)
    if publisherId == nil || *publisherId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherId"} 
    }
    routeValues["publisherId"] = *publisherId
    if eventTypeId == nil || *eventTypeId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "eventTypeId"} 
    }
    routeValues["eventTypeId"] = *eventTypeId

    locationId, _ := uuid.Parse("db4777cd-8e08-4a84-8ba3-c974ea033718")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue EventTypeDescriptor
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get the event types for a specific publisher.
// ctx
// publisherId (required): ID for a publisher.
func (client Client) ListEventTypes(ctx context.Context, publisherId *string) (*[]EventTypeDescriptor, error) {
    routeValues := make(map[string]string)
    if publisherId == nil || *publisherId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherId"} 
    }
    routeValues["publisherId"] = *publisherId

    locationId, _ := uuid.Parse("db4777cd-8e08-4a84-8ba3-c974ea033718")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []EventTypeDescriptor
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Get a specific notification for a subscription.
// ctx
// subscriptionId (required): ID for a subscription.
// notificationId (required)
func (client Client) GetNotification(ctx context.Context, subscriptionId *uuid.UUID, notificationId *int) (*Notification, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = (*subscriptionId).String()
    if notificationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "notificationId"} 
    }
    routeValues["notificationId"] = strconv.Itoa(*notificationId)

    locationId, _ := uuid.Parse("0c62d343-21b0-4732-997b-017fde84dc28")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Notification
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of notifications for a specific subscription. A notification includes details about the event, the request to and the response from the consumer service.
// ctx
// subscriptionId (required): ID for a subscription.
// maxResults (optional): Maximum number of notifications to return. Default is **100**.
// status (optional): Get only notifications with this status.
// result (optional): Get only notifications with this result type.
func (client Client) GetNotifications(ctx context.Context, subscriptionId *uuid.UUID, maxResults *int, status *NotificationStatus, result *NotificationResult) (*[]Notification, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = (*subscriptionId).String()

    queryParams := url.Values{}
    if maxResults != nil {
        queryParams.Add("maxResults", strconv.Itoa(*maxResults))
    }
    if status != nil {
        queryParams.Add("status", string(*status))
    }
    if result != nil {
        queryParams.Add("result", string(*result))
    }
    locationId, _ := uuid.Parse("0c62d343-21b0-4732-997b-017fde84dc28")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Notification
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Query for notifications. A notification includes details about the event, the request to and the response from the consumer service.
// ctx
// query (required)
func (client Client) QueryNotifications(ctx context.Context, query *NotificationsQuery) (*NotificationsQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1a57562f-160a-4b5c-9185-905e95b39d36")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue NotificationsQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// ctx
// inputValuesQuery (required)
// publisherId (required)
func (client Client) QueryInputValues(ctx context.Context, inputValuesQuery *InputValuesQuery, publisherId *string) (*InputValuesQuery, error) {
    if inputValuesQuery == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "inputValuesQuery"}
    }
    routeValues := make(map[string]string)
    if publisherId == nil || *publisherId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherId"} 
    }
    routeValues["publisherId"] = *publisherId

    body, marshalErr := json.Marshal(*inputValuesQuery)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d815d352-a566-4dc1-a3e3-fd245acf688c")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue InputValuesQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a specific service hooks publisher.
// ctx
// publisherId (required): ID for a publisher.
func (client Client) GetPublisher(ctx context.Context, publisherId *string) (*Publisher, error) {
    routeValues := make(map[string]string)
    if publisherId == nil || *publisherId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "publisherId"} 
    }
    routeValues["publisherId"] = *publisherId

    locationId, _ := uuid.Parse("1e83a210-5b53-43bc-90f0-d476a4e5d731")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Publisher
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of publishers.
// ctx
func (client Client) ListPublishers(ctx context.Context, ) (*[]Publisher, error) {
    locationId, _ := uuid.Parse("1e83a210-5b53-43bc-90f0-d476a4e5d731")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Publisher
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Query for service hook publishers.
// ctx
// query (required)
func (client Client) QueryPublishers(ctx context.Context, query *PublishersQuery) (*PublishersQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("99b44a8a-65a8-4670-8f3e-e7f7842cce64")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue PublishersQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Create a subscription.
// ctx
// subscription (required): Subscription to be created.
func (client Client) CreateSubscription(ctx context.Context, subscription *Subscription) (*Subscription, error) {
    if subscription == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscription"}
    }
    body, marshalErr := json.Marshal(*subscription)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("fc50d02a-849f-41fb-8af1-0a5216103269")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Subscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Delete a specific service hooks subscription.
// ctx
// subscriptionId (required): ID for a subscription.
func (client Client) DeleteSubscription(ctx context.Context, subscriptionId *uuid.UUID) error {
    routeValues := make(map[string]string)
    if subscriptionId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = (*subscriptionId).String()

    locationId, _ := uuid.Parse("fc50d02a-849f-41fb-8af1-0a5216103269")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Get a specific service hooks subscription.
// ctx
// subscriptionId (required): ID for a subscription.
func (client Client) GetSubscription(ctx context.Context, subscriptionId *uuid.UUID) (*Subscription, error) {
    routeValues := make(map[string]string)
    if subscriptionId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscriptionId"} 
    }
    routeValues["subscriptionId"] = (*subscriptionId).String()

    locationId, _ := uuid.Parse("fc50d02a-849f-41fb-8af1-0a5216103269")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Subscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Get a list of subscriptions.
// ctx
// publisherId (optional): ID for a subscription.
// eventType (optional): The event type to filter on (if any).
// consumerId (optional): ID for a consumer.
// consumerActionId (optional): ID for a consumerActionId.
func (client Client) ListSubscriptions(ctx context.Context, publisherId *string, eventType *string, consumerId *string, consumerActionId *string) (*[]Subscription, error) {
    queryParams := url.Values{}
    if publisherId != nil {
        queryParams.Add("publisherId", *publisherId)
    }
    if eventType != nil {
        queryParams.Add("eventType", *eventType)
    }
    if consumerId != nil {
        queryParams.Add("consumerId", *consumerId)
    }
    if consumerActionId != nil {
        queryParams.Add("consumerActionId", *consumerActionId)
    }
    locationId, _ := uuid.Parse("fc50d02a-849f-41fb-8af1-0a5216103269")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Subscription
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Update a subscription. <param name="subscriptionId">ID for a subscription that you wish to update.</param>
// ctx
// subscription (required)
// subscriptionId (optional)
func (client Client) ReplaceSubscription(ctx context.Context, subscription *Subscription, subscriptionId *uuid.UUID) (*Subscription, error) {
    if subscription == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "subscription"}
    }
    routeValues := make(map[string]string)
    if subscriptionId != nil {
        routeValues["subscriptionId"] = (*subscriptionId).String()
    }

    body, marshalErr := json.Marshal(*subscription)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("fc50d02a-849f-41fb-8af1-0a5216103269")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Subscription
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Query for service hook subscriptions.
// ctx
// query (required)
func (client Client) CreateSubscriptionsQuery(ctx context.Context, query *SubscriptionsQuery) (*SubscriptionsQuery, error) {
    if query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    body, marshalErr := json.Marshal(*query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("c7c3c1cf-9e05-4c0d-a425-a0f922c2c6ed")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SubscriptionsQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Sends a test notification. This is useful for verifying the configuration of an updated or new service hooks subscription.
// ctx
// testNotification (required)
// useRealData (optional): Only allow testing with real data in existing subscriptions.
func (client Client) CreateTestNotification(ctx context.Context, testNotification *Notification, useRealData *bool) (*Notification, error) {
    if testNotification == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testNotification"}
    }
    queryParams := url.Values{}
    if useRealData != nil {
        queryParams.Add("useRealData", strconv.FormatBool(*useRealData))
    }
    body, marshalErr := json.Marshal(*testNotification)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1139462c-7e27-4524-a997-31b9b73551fe")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Notification
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

