// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package serviceHooks

import (
    "github.com/google/uuid"
    "math/big"
    "time"
)

// Enumerates consumer authentication types.
type AuthenticationType string

type authenticationTypeValuesType struct {
    None AuthenticationType
    OAuth AuthenticationType
    External AuthenticationType
}

var AuthenticationTypeValues = authenticationTypeValuesType{
    // No authentication is required.
    None: "none",
    // OAuth authentication.
    OAuth: "oAuth",
    // Externally-configured authentication.
    External: "external",
}

// Defines the data contract of a consumer.
type Consumer struct {
    // Reference Links
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets this consumer's actions.
    Actions *[]ConsumerAction `json:"actions,omitempty"`
    // Gets or sets this consumer's authentication type.
    AuthenticationType *AuthenticationType `json:"authenticationType,omitempty"`
    // Gets or sets this consumer's localized description.
    Description *string `json:"description,omitempty"`
    // Non-null only if subscriptions for this consumer are configured externally.
    ExternalConfiguration *ExternalConfigurationDescriptor `json:"externalConfiguration,omitempty"`
    // Gets or sets this consumer's identifier.
    Id *string `json:"id,omitempty"`
    // Gets or sets this consumer's image URL, if any.
    ImageUrl *string `json:"imageUrl,omitempty"`
    // Gets or sets this consumer's information URL, if any.
    InformationUrl *string `json:"informationUrl,omitempty"`
    // Gets or sets this consumer's input descriptors.
    InputDescriptors *[]InputDescriptor `json:"inputDescriptors,omitempty"`
    // Gets or sets this consumer's localized name.
    Name *string `json:"name,omitempty"`
    // The url for this resource
    Url *string `json:"url,omitempty"`
}

// Defines the data contract of a consumer action.
type ConsumerAction struct {
    // Reference Links
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets or sets the flag indicating if resource version can be overridden when creating or editing a subscription.
    AllowResourceVersionOverride *bool `json:"allowResourceVersionOverride,omitempty"`
    // Gets or sets the identifier of the consumer to which this action belongs.
    ConsumerId *string `json:"consumerId,omitempty"`
    // Gets or sets this action's localized description.
    Description *string `json:"description,omitempty"`
    // Gets or sets this action's identifier.
    Id *string `json:"id,omitempty"`
    // Gets or sets this action's input descriptors.
    InputDescriptors *[]InputDescriptor `json:"inputDescriptors,omitempty"`
    // Gets or sets this action's localized name.
    Name *string `json:"name,omitempty"`
    // Gets or sets this action's supported event identifiers.
    SupportedEventTypes *[]string `json:"supportedEventTypes,omitempty"`
    // Gets or sets this action's supported resource versions.
    SupportedResourceVersions *map[string][]string `json:"supportedResourceVersions,omitempty"`
    // The url for this resource
    Url *string `json:"url,omitempty"`
}

// Encapsulates the properties of an event.
type Event struct {
    // Gets or sets the UTC-based date and time that this event was created.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Gets or sets the detailed message associated with this event.
    DetailedMessage *FormattedEventMessage `json:"detailedMessage,omitempty"`
    // Gets or sets the type of this event.
    EventType *string `json:"eventType,omitempty"`
    // Gets or sets the unique identifier of this event.
    Id *uuid.UUID `json:"id,omitempty"`
    // Gets or sets the (brief) message associated with this event.
    Message *FormattedEventMessage `json:"message,omitempty"`
    // Gets or sets the identifier of the publisher that raised this event.
    PublisherId *string `json:"publisherId,omitempty"`
    // Gets or sets the data associated with this event.
    Resource *interface{} `json:"resource,omitempty"`
    // Gets or sets the resource containers.
    ResourceContainers *map[string]ResourceContainer `json:"resourceContainers,omitempty"`
    // Gets or sets the version of the data associated with this event.
    ResourceVersion *string `json:"resourceVersion,omitempty"`
    // Gets or sets the Session Token that can be used in further interactions
    SessionToken *SessionToken `json:"sessionToken,omitempty"`
}

// Describes a type of event
type EventTypeDescriptor struct {
    // A localized description of the event type
    Description *string `json:"description,omitempty"`
    // A unique id for the event type
    Id *string `json:"id,omitempty"`
    // Event-specific inputs
    InputDescriptors *[]InputDescriptor `json:"inputDescriptors,omitempty"`
    // A localized friendly name for the event type
    Name *string `json:"name,omitempty"`
    // A unique id for the publisher of this event type
    PublisherId *string `json:"publisherId,omitempty"`
    // Supported versions for the event's resource payloads.
    SupportedResourceVersions *[]string `json:"supportedResourceVersions,omitempty"`
    // The url for this resource
    Url *string `json:"url,omitempty"`
}

// Describes how to configure a subscription that is managed externally.
type ExternalConfigurationDescriptor struct {
    // Url of the site to create this type of subscription.
    CreateSubscriptionUrl *string `json:"createSubscriptionUrl,omitempty"`
    // The name of an input property that contains the URL to edit a subscription.
    EditSubscriptionPropertyName *string `json:"editSubscriptionPropertyName,omitempty"`
    // True if the external configuration applies only to hosted.
    HostedOnly *bool `json:"hostedOnly,omitempty"`
}

// Provides different formats of an event message
type FormattedEventMessage struct {
    // Gets or sets the html format of the message
    Html *string `json:"html,omitempty"`
    // Gets or sets the markdown format of the message
    Markdown *string `json:"markdown,omitempty"`
    // Gets or sets the raw text of the message
    Text *string `json:"text,omitempty"`
}

type GraphSubjectBase struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
}

type IdentityRef struct {
    // This field contains zero or more interesting links about the graph subject. These links may be invoked to obtain additional relationships or more detailed information about this graph subject.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // The descriptor is the primary way to reference the graph subject while the system is running. This field will uniquely identify the same graph subject across both Accounts and Organizations.
    Descriptor *string `json:"descriptor,omitempty"`
    // This is the non-unique display name of the graph subject. To change this field, you must alter its value in the source provider.
    DisplayName *string `json:"displayName,omitempty"`
    // This url is the full route to the source resource of this graph subject.
    Url *string `json:"url,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph user referenced in the "self" entry of the IdentityRef "_links" dictionary
    DirectoryAlias *string `json:"directoryAlias,omitempty"`
    Id *string `json:"id,omitempty"`
    // Deprecated - Available in the "avatar" entry of the IdentityRef "_links" dictionary
    ImageUrl *string `json:"imageUrl,omitempty"`
    // Deprecated - Can be retrieved by querying the Graph membership state referenced in the "membershipState" entry of the GraphUser "_links" dictionary
    Inactive *bool `json:"inactive,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsAadUserType/Descriptor.IsAadGroupType)
    IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
    // Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsGroupType)
    IsContainer *bool `json:"isContainer,omitempty"`
    IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
    // Deprecated - not in use in most preexisting implementations of ToIdentityRef
    ProfileUrl *string `json:"profileUrl,omitempty"`
    // Deprecated - use Domain+PrincipalName instead
    UniqueName *string `json:"uniqueName,omitempty"`
}

// Enumerates data types that are supported as subscription input values.
type InputDataType string

type inputDataTypeValuesType struct {
    None InputDataType
    String InputDataType
    Number InputDataType
    Boolean InputDataType
    Guid InputDataType
    Uri InputDataType
}

var InputDataTypeValues = inputDataTypeValuesType{
    // No data type is specified.
    None: "none",
    // Represents a textual value.
    String: "string",
    // Represents a numeric value.
    Number: "number",
    // Represents a value of true or false.
    Boolean: "boolean",
    // Represents a Guid.
    Guid: "guid",
    // Represents a URI.
    Uri: "uri",
}

// Describes an input for subscriptions.
type InputDescriptor struct {
    // The ids of all inputs that the value of this input is dependent on.
    DependencyInputIds *[]string `json:"dependencyInputIds,omitempty"`
    // Description of what this input is used for
    Description *string `json:"description,omitempty"`
    // The group localized name to which this input belongs and can be shown as a header for the container that will include all the inputs in the group.
    GroupName *string `json:"groupName,omitempty"`
    // If true, the value information for this input is dynamic and should be fetched when the value of dependency inputs change.
    HasDynamicValueInformation *bool `json:"hasDynamicValueInformation,omitempty"`
    // Identifier for the subscription input
    Id *string `json:"id,omitempty"`
    // Mode in which the value of this input should be entered
    InputMode *InputMode `json:"inputMode,omitempty"`
    // Gets whether this input is confidential, such as for a password or application key
    IsConfidential *bool `json:"isConfidential,omitempty"`
    // Localized name which can be shown as a label for the subscription input
    Name *string `json:"name,omitempty"`
    // Custom properties for the input which can be used by the service provider
    Properties *map[string]interface{} `json:"properties,omitempty"`
    // Underlying data type for the input value. When this value is specified, InputMode, Validation and Values are optional.
    Type *string `json:"type,omitempty"`
    // Gets whether this input is included in the default generated action description.
    UseInDefaultDescription *bool `json:"useInDefaultDescription,omitempty"`
    // Information to use to validate this input's value
    Validation *InputValidation `json:"validation,omitempty"`
    // A hint for input value. It can be used in the UI as the input placeholder.
    ValueHint *string `json:"valueHint,omitempty"`
    // Information about possible values for this input
    Values *InputValues `json:"values,omitempty"`
}

// Defines a filter for subscription inputs. The filter matches a set of inputs if any (one or more) of the groups evaluates to true.
type InputFilter struct {
    // Groups of input filter expressions. This filter matches a set of inputs if any (one or more) of the groups evaluates to true.
    Conditions *[]InputFilterCondition `json:"conditions,omitempty"`
}

// An expression which can be applied to filter a list of subscription inputs
type InputFilterCondition struct {
    // Whether or not to do a case sensitive match
    CaseSensitive *bool `json:"caseSensitive,omitempty"`
    // The Id of the input to filter on
    InputId *string `json:"inputId,omitempty"`
    // The "expected" input value to compare with the actual input value
    InputValue *string `json:"inputValue,omitempty"`
    // The operator applied between the expected and actual input value
    Operator *InputFilterOperator `json:"operator,omitempty"`
}

type InputFilterOperator string

type inputFilterOperatorValuesType struct {
    Equals InputFilterOperator
    NotEquals InputFilterOperator
}

var InputFilterOperatorValues = inputFilterOperatorValuesType{
    Equals: "equals",
    NotEquals: "notEquals",
}

// Mode in which a subscription input should be entered (in a UI)
type InputMode string

type inputModeValuesType struct {
    None InputMode
    TextBox InputMode
    PasswordBox InputMode
    Combo InputMode
    RadioButtons InputMode
    CheckBox InputMode
    TextArea InputMode
}

var InputModeValues = inputModeValuesType{
    // This input should not be shown in the UI
    None: "none",
    // An input text box should be shown
    TextBox: "textBox",
    // An password input box should be shown
    PasswordBox: "passwordBox",
    // A select/combo control should be shown
    Combo: "combo",
    // Radio buttons should be shown
    RadioButtons: "radioButtons",
    // Checkbox should be shown(for true/false values)
    CheckBox: "checkBox",
    // A multi-line text area should be shown
    TextArea: "textArea",
}

// Describes what values are valid for a subscription input
type InputValidation struct {
    // Gets or sets the data data type to validate.
    DataType *InputDataType `json:"dataType,omitempty"`
    // Gets or sets if this is a required field.
    IsRequired *bool `json:"isRequired,omitempty"`
    // Gets or sets the maximum length of this descriptor.
    MaxLength *int `json:"maxLength,omitempty"`
    // Gets or sets the minimum value for this descriptor.
    MaxValue *big.Float `json:"maxValue,omitempty"`
    // Gets or sets the minimum length of this descriptor.
    MinLength *int `json:"minLength,omitempty"`
    // Gets or sets the minimum value for this descriptor.
    MinValue *big.Float `json:"minValue,omitempty"`
    // Gets or sets the pattern to validate.
    Pattern *string `json:"pattern,omitempty"`
    // Gets or sets the error on pattern mismatch.
    PatternMismatchErrorMessage *string `json:"patternMismatchErrorMessage,omitempty"`
}

// Information about a single value for an input
type InputValue struct {
    // Any other data about this input
    Data *map[string]interface{} `json:"data,omitempty"`
    // The text to show for the display of this value
    DisplayValue *string `json:"displayValue,omitempty"`
    // The value to store for this input
    Value *string `json:"value,omitempty"`
}

// Information about the possible/allowed values for a given subscription input
type InputValues struct {
    // The default value to use for this input
    DefaultValue *string `json:"defaultValue,omitempty"`
    // Errors encountered while computing dynamic values.
    Error *InputValuesError `json:"error,omitempty"`
    // The id of the input
    InputId *string `json:"inputId,omitempty"`
    // Should this input be disabled
    IsDisabled *bool `json:"isDisabled,omitempty"`
    // Should the value be restricted to one of the values in the PossibleValues (True) or are the values in PossibleValues just a suggestion (False)
    IsLimitedToPossibleValues *bool `json:"isLimitedToPossibleValues,omitempty"`
    // Should this input be made read-only
    IsReadOnly *bool `json:"isReadOnly,omitempty"`
    // Possible values that this input can take
    PossibleValues *[]InputValue `json:"possibleValues,omitempty"`
}

// Error information related to a subscription input value.
type InputValuesError struct {
    // The error message.
    Message *string `json:"message,omitempty"`
}

type InputValuesQuery struct {
    CurrentValues *map[string]string `json:"currentValues,omitempty"`
    // The input values to return on input, and the result from the consumer on output.
    InputValues *[]InputValues `json:"inputValues,omitempty"`
    // Subscription containing information about the publisher/consumer and the current input values
    Resource *interface{} `json:"resource,omitempty"`
}

// Defines the data contract of the result of processing an event for a subscription.
type Notification struct {
    // Gets or sets date and time that this result was created.
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    // Details about this notification (if available)
    Details *NotificationDetails `json:"details,omitempty"`
    // The event id associated with this notification
    EventId *uuid.UUID `json:"eventId,omitempty"`
    // The notification id
    Id *int `json:"id,omitempty"`
    // Gets or sets date and time that this result was last modified.
    ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
    // Result of the notification
    Result *NotificationResult `json:"result,omitempty"`
    // Status of the notification
    Status *NotificationStatus `json:"status,omitempty"`
    // The subscriber Id  associated with this notification. This is the last identity who touched in the subscription. In case of test notifications it can be the tester if the subscription is not created yet.
    SubscriberId *uuid.UUID `json:"subscriberId,omitempty"`
    // The subscription id associated with this notification
    SubscriptionId *uuid.UUID `json:"subscriptionId,omitempty"`
}

// Defines the data contract of notification details.
type NotificationDetails struct {
    // Gets or sets the time that this notification was completed (response received from the consumer)
    CompletedDate *time.Time `json:"completedDate,omitempty"`
    // Gets or sets this notification detail's consumer action identifier.
    ConsumerActionId *string `json:"consumerActionId,omitempty"`
    // Gets or sets this notification detail's consumer identifier.
    ConsumerId *string `json:"consumerId,omitempty"`
    // Gets or sets this notification detail's consumer inputs.
    ConsumerInputs *map[string]string `json:"consumerInputs,omitempty"`
    // Gets or sets the time that this notification was dequeued for processing
    DequeuedDate *time.Time `json:"dequeuedDate,omitempty"`
    // Gets or sets this notification detail's error detail.
    ErrorDetail *string `json:"errorDetail,omitempty"`
    // Gets or sets this notification detail's error message.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // Gets or sets this notification detail's event content.
    Event *Event `json:"event,omitempty"`
    // Gets or sets this notification detail's event type.
    EventType *string `json:"eventType,omitempty"`
    // Gets or sets the time that this notification was finished processing (just before the request is sent to the consumer)
    ProcessedDate *time.Time `json:"processedDate,omitempty"`
    // Gets or sets this notification detail's publisher identifier.
    PublisherId *string `json:"publisherId,omitempty"`
    // Gets or sets this notification detail's publisher inputs.
    PublisherInputs *map[string]string `json:"publisherInputs,omitempty"`
    // Gets or sets the time that this notification was queued (created)
    QueuedDate *time.Time `json:"queuedDate,omitempty"`
    // Gets or sets this notification detail's request.
    Request *string `json:"request,omitempty"`
    // Number of requests attempted to be sent to the consumer
    RequestAttempts *int `json:"requestAttempts,omitempty"`
    // Duration of the request to the consumer in seconds
    RequestDuration *float64 `json:"requestDuration,omitempty"`
    // Gets or sets this notification detail's response.
    Response *string `json:"response,omitempty"`
}

// Enumerates possible result types of a notification.
type NotificationResult string

type notificationResultValuesType struct {
    Pending NotificationResult
    Succeeded NotificationResult
    Failed NotificationResult
    Filtered NotificationResult
}

var NotificationResultValues = notificationResultValuesType{
    // The notification has not yet completed
    Pending: "pending",
    // The notification was sent successfully
    Succeeded: "succeeded",
    // The notification failed to be sent successfully to the consumer
    Failed: "failed",
    // The notification was filtered by the Delivery Job
    Filtered: "filtered",
}

// Summary of a particular result and count.
type NotificationResultsSummaryDetail struct {
    // Count of notification sent out with a matching result.
    NotificationCount *int `json:"notificationCount,omitempty"`
    // Result of the notification
    Result *NotificationResult `json:"result,omitempty"`
}

// Defines a query for service hook notifications.
type NotificationsQuery struct {
    // The subscriptions associated with the notifications returned from the query
    AssociatedSubscriptions *[]Subscription `json:"associatedSubscriptions,omitempty"`
    // If true, we will return all notification history for the query provided; otherwise, the summary is returned.
    IncludeDetails *bool `json:"includeDetails,omitempty"`
    // Optional maximum date at which the notification was created
    MaxCreatedDate *time.Time `json:"maxCreatedDate,omitempty"`
    // Optional maximum number of overall results to include
    MaxResults *int `json:"maxResults,omitempty"`
    // Optional maximum number of results for each subscription. Only takes effect when a list of subscription ids is supplied in the query.
    MaxResultsPerSubscription *int `json:"maxResultsPerSubscription,omitempty"`
    // Optional minimum date at which the notification was created
    MinCreatedDate *time.Time `json:"minCreatedDate,omitempty"`
    // Optional publisher id to restrict the results to
    PublisherId *string `json:"publisherId,omitempty"`
    // Results from the query
    Results *[]Notification `json:"results,omitempty"`
    // Optional notification result type to filter results to
    ResultType *NotificationResult `json:"resultType,omitempty"`
    // Optional notification status to filter results to
    Status *NotificationStatus `json:"status,omitempty"`
    // Optional list of subscription ids to restrict the results to
    SubscriptionIds *[]uuid.UUID `json:"subscriptionIds,omitempty"`
    // Summary of notifications - the count of each result type (success, fail, ..).
    Summary *[]NotificationSummary `json:"summary,omitempty"`
}

// Enumerates possible status' of a notification.
type NotificationStatus string

type notificationStatusValuesType struct {
    Queued NotificationStatus
    Processing NotificationStatus
    RequestInProgress NotificationStatus
    Completed NotificationStatus
}

var NotificationStatusValues = notificationStatusValuesType{
    // The notification has been queued
    Queued: "queued",
    // The notification has been dequeued and has begun processing.
    Processing: "processing",
    // The consumer action has processed the notification. The request is in progress.
    RequestInProgress: "requestInProgress",
    // The request completed
    Completed: "completed",
}

// Summary of the notifications for a subscription.
type NotificationSummary struct {
    // The notification results for this particular subscription.
    Results *[]NotificationResultsSummaryDetail `json:"results,omitempty"`
    // The subscription id associated with this notification
    SubscriptionId *uuid.UUID `json:"subscriptionId,omitempty"`
}

// Defines the data contract of an event publisher.
type Publisher struct {
    // Reference Links
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets this publisher's localized description.
    Description *string `json:"description,omitempty"`
    // Gets this publisher's identifier.
    Id *string `json:"id,omitempty"`
    // Publisher-specific inputs
    InputDescriptors *[]InputDescriptor `json:"inputDescriptors,omitempty"`
    // Gets this publisher's localized name.
    Name *string `json:"name,omitempty"`
    // The service instance type of the first party publisher.
    ServiceInstanceType *string `json:"serviceInstanceType,omitempty"`
    // Gets this publisher's supported event types.
    SupportedEvents *[]EventTypeDescriptor `json:"supportedEvents,omitempty"`
    // The url for this resource
    Url *string `json:"url,omitempty"`
}

// Wrapper around an event which is being published
type PublisherEvent struct {
    // Add key/value pairs which will be stored with a published notification in the SH service DB.  This key/value pairs are for diagnostic purposes only and will have not effect on the delivery of a notificaton.
    Diagnostics *map[string]string `json:"diagnostics,omitempty"`
    // The event being published
    Event *Event `json:"event,omitempty"`
    // Gets or sets flag for filtered events
    IsFilteredEvent *bool `json:"isFilteredEvent,omitempty"`
    // Additional data that needs to be sent as part of notification to complement the Resource data in the Event
    NotificationData *map[string]string `json:"notificationData,omitempty"`
    // Gets or sets the array of older supported resource versions.
    OtherResourceVersions *[]VersionedResource `json:"otherResourceVersions,omitempty"`
    // Optional publisher-input filters which restricts the set of subscriptions which are triggered by the event
    PublisherInputFilters *[]InputFilter `json:"publisherInputFilters,omitempty"`
    // Gets or sets matched hooks subscription which caused this event.
    Subscription *Subscription `json:"subscription,omitempty"`
}

// Defines a query for service hook publishers.
type PublishersQuery struct {
    // Optional list of publisher ids to restrict the results to
    PublisherIds *[]string `json:"publisherIds,omitempty"`
    // Filter for publisher inputs
    PublisherInputs *map[string]string `json:"publisherInputs,omitempty"`
    // Results from the query
    Results *[]Publisher `json:"results,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

// The base class for all resource containers, i.e. Account, Collection, Project
type ResourceContainer struct {
    // Gets or sets the container's base URL, i.e. the URL of the host (collection, application, or deployment) containing the container resource.
    BaseUrl *string `json:"baseUrl,omitempty"`
    // Gets or sets the container's specific Id.
    Id *uuid.UUID `json:"id,omitempty"`
    // Gets or sets the container's name.
    Name *string `json:"name,omitempty"`
    // Gets or sets the container's REST API URL.
    Url *string `json:"url,omitempty"`
}

// Represents a session token to be attached in Events for Consumer actions that need it.
type SessionToken struct {
    // The error message in case of error
    Error *string `json:"error,omitempty"`
    // The access token
    Token *string `json:"token,omitempty"`
    // The expiration date in UTC
    ValidTo *time.Time `json:"validTo,omitempty"`
}

// Encapsulates an event subscription.
type Subscription struct {
    // Reference Links
    Links *ReferenceLinks `json:"_links,omitempty"`
    ActionDescription *string `json:"actionDescription,omitempty"`
    ConsumerActionId *string `json:"consumerActionId,omitempty"`
    ConsumerId *string `json:"consumerId,omitempty"`
    // Consumer input values
    ConsumerInputs *map[string]string `json:"consumerInputs,omitempty"`
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    CreatedDate *time.Time `json:"createdDate,omitempty"`
    EventDescription *string `json:"eventDescription,omitempty"`
    EventType *string `json:"eventType,omitempty"`
    Id *uuid.UUID `json:"id,omitempty"`
    ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`
    ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
    ProbationRetries *byte `json:"probationRetries,omitempty"`
    PublisherId *string `json:"publisherId,omitempty"`
    // Publisher input values
    PublisherInputs *map[string]string `json:"publisherInputs,omitempty"`
    ResourceVersion *string `json:"resourceVersion,omitempty"`
    Status *SubscriptionStatus `json:"status,omitempty"`
    Subscriber *IdentityRef `json:"subscriber,omitempty"`
    Url *string `json:"url,omitempty"`
}

type SubscriptionDiagnostics struct {
    DeliveryResults *SubscriptionTracing `json:"deliveryResults,omitempty"`
    DeliveryTracing *SubscriptionTracing `json:"deliveryTracing,omitempty"`
    EvaluationTracing *SubscriptionTracing `json:"evaluationTracing,omitempty"`
}

// The scope to which a subscription input applies
type SubscriptionInputScope string

type subscriptionInputScopeValuesType struct {
    Publisher SubscriptionInputScope
    Consumer SubscriptionInputScope
}

var SubscriptionInputScopeValues = subscriptionInputScopeValuesType{
    // An input defined and consumed by a Publisher or Publisher Event Type
    Publisher: "publisher",
    // An input defined and consumed by a Consumer or Consumer Action
    Consumer: "consumer",
}

// Query for obtaining information about the possible/allowed values for one or more subscription inputs
type SubscriptionInputValuesQuery struct {
    // The input values to return on input, and the result from the consumer on output.
    InputValues *[]InputValues `json:"inputValues,omitempty"`
    // The scope at which the properties to query belong
    Scope *SubscriptionInputScope `json:"scope,omitempty"`
    // Subscription containing information about the publisher/consumer and the current input values
    Subscription *Subscription `json:"subscription,omitempty"`
}

// Defines a query for service hook subscriptions.
type SubscriptionsQuery struct {
    // Optional consumer action id to restrict the results to (null for any)
    ConsumerActionId *string `json:"consumerActionId,omitempty"`
    // Optional consumer id to restrict the results to (null for any)
    ConsumerId *string `json:"consumerId,omitempty"`
    // Filter for subscription consumer inputs
    ConsumerInputFilters *[]InputFilter `json:"consumerInputFilters,omitempty"`
    // Optional event type id to restrict the results to (null for any)
    EventType *string `json:"eventType,omitempty"`
    // Optional publisher id to restrict the results to (null for any)
    PublisherId *string `json:"publisherId,omitempty"`
    // Filter for subscription publisher inputs
    PublisherInputFilters *[]InputFilter `json:"publisherInputFilters,omitempty"`
    // Results from the query
    Results *[]Subscription `json:"results,omitempty"`
    // Optional subscriber filter.
    SubscriberId *uuid.UUID `json:"subscriberId,omitempty"`
}

// Enumerates possible states of a subscription.
type SubscriptionStatus string

type subscriptionStatusValuesType struct {
    Enabled SubscriptionStatus
    OnProbation SubscriptionStatus
    DisabledByUser SubscriptionStatus
    DisabledBySystem SubscriptionStatus
    DisabledByInactiveIdentity SubscriptionStatus
}

var SubscriptionStatusValues = subscriptionStatusValuesType{
    // The subscription is enabled.
    Enabled: "enabled",
    // The subscription is temporarily on probation by the system.
    OnProbation: "onProbation",
    // The subscription is disabled by a user.
    DisabledByUser: "disabledByUser",
    // The subscription is disabled by the system.
    DisabledBySystem: "disabledBySystem",
    // The subscription is disabled because the owner is inactive or is missing permissions.
    DisabledByInactiveIdentity: "disabledByInactiveIdentity",
}

type SubscriptionTracing struct {
    Enabled *bool `json:"enabled,omitempty"`
    // Trace until the specified end date.
    EndDate *time.Time `json:"endDate,omitempty"`
    // The maximum number of result details to trace.
    MaxTracedEntries *int `json:"maxTracedEntries,omitempty"`
    // The date and time tracing started.
    StartDate *time.Time `json:"startDate,omitempty"`
    // Trace until remaining count reaches 0.
    TracedEntries *int `json:"tracedEntries,omitempty"`
}

type UpdateSubscripitonDiagnosticsParameters struct {
    DeliveryResults *UpdateSubscripitonTracingParameters `json:"deliveryResults,omitempty"`
    DeliveryTracing *UpdateSubscripitonTracingParameters `json:"deliveryTracing,omitempty"`
    EvaluationTracing *UpdateSubscripitonTracingParameters `json:"evaluationTracing,omitempty"`
}

type UpdateSubscripitonTracingParameters struct {
    Enabled *bool `json:"enabled,omitempty"`
}

// Encapsulates the resource version and its data or reference to the compatible version. Only one of the two last fields should be not null.
type VersionedResource struct {
    // Gets or sets the reference to the compatible version.
    CompatibleWith *string `json:"compatibleWith,omitempty"`
    // Gets or sets the resource data.
    Resource *interface{} `json:"resource,omitempty"`
    // Gets or sets the version of the resource data.
    ResourceVersion *string `json:"resourceVersion,omitempty"`
}
