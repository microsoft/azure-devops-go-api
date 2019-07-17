// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package dashboard

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("31c84e0a-3ece-48fd-a29d-100849af99ba")

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

// [Preview API] Create the supplied dashboard.
// dashboard (required): The initial state of the dashboard
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) CreateDashboard(dashboard *Dashboard, project *string, team *string) (*Dashboard, error) {
    if dashboard == nil {
        return nil, errors.New("dashboard is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    body, marshalErr := json.Marshal(*dashboard)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a dashboard given its ID. This also deletes the widgets associated with this dashboard.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard to delete.
// team (optional): Team ID or team name
func (client Client) DeleteDashboard(project *string, dashboardId *uuid.UUID, team *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    _, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a dashboard by its ID.
// project (required): Project ID or project name
// dashboardId (required)
// team (optional): Team ID or team name
func (client Client) GetDashboard(project *string, dashboardId *uuid.UUID, team *string) (*Dashboard, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of dashboards.
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) GetDashboards(project *string, team *string) (*DashboardGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DashboardGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Replace configuration for the specified dashboard. Replaces Widget list on Dashboard, only if property is supplied.
// dashboard (required): The Configuration of the dashboard to replace.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard to replace.
// team (optional): Team ID or team name
func (client Client) ReplaceDashboard(dashboard *Dashboard, project *string, dashboardId *uuid.UUID, team *string) (*Dashboard, error) {
    if dashboard == nil {
        return nil, errors.New("dashboard is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    body, marshalErr := json.Marshal(*dashboard)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update the name and position of dashboards in the supplied group, and remove omitted dashboards. Does not modify dashboard content.
// group (required)
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) ReplaceDashboards(group *DashboardGroup, project *string, team *string) (*DashboardGroup, error) {
    if group == nil {
        return nil, errors.New("group is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DashboardGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a widget on the specified dashboard.
// widget (required): State of the widget to add
// project (required): Project ID or project name
// dashboardId (required): ID of dashboard the widget will be added to.
// team (optional): Team ID or team name
func (client Client) CreateWidget(widget *Widget, project *string, dashboardId *uuid.UUID, team *string) (*Widget, error) {
    if widget == nil {
        return nil, errors.New("widget is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    body, marshalErr := json.Marshal(*widget)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete the specified widget.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to update.
// team (optional): Team ID or team name
func (client Client) DeleteWidget(project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Dashboard, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, errors.New("widgetId is a required parameter")
    }
    routeValues["widgetId"] = (*widgetId).String()

    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the current state of the specified widget.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to read.
// team (optional): Team ID or team name
func (client Client) GetWidget(project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Widget, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, errors.New("widgetId is a required parameter")
    }
    routeValues["widgetId"] = (*widgetId).String()

    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get widgets contained on the specified dashboard.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard to read.
// team (optional): Team ID or team name
// eTag (optional): Dashboard Widgets Version
func (client Client) GetWidgets(project *string, dashboardId *uuid.UUID, team *string, eTag *string) (*WidgetsVersionedList, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    additionalHeaders := make(map[string]string)
    if eTag != nil {
        additionalHeaders["ETag"] = *eTag
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseBodyValue []Widget
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WidgetsVersionedList
    if err == nil {
        responseValue = &WidgetsVersionedList {
            Widgets: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// [Preview API] Override the  state of the specified widget.
// widget (required): State to be written for the widget.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to update.
// team (optional): Team ID or team name
func (client Client) ReplaceWidget(widget *Widget, project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Widget, error) {
    if widget == nil {
        return nil, errors.New("widget is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, errors.New("widgetId is a required parameter")
    }
    routeValues["widgetId"] = (*widgetId).String()

    body, marshalErr := json.Marshal(*widget)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Replace the widgets on specified dashboard with the supplied widgets.
// widgets (required): Revised state of widgets to store for the dashboard.
// project (required): Project ID or project name
// dashboardId (required): ID of the Dashboard to modify.
// team (optional): Team ID or team name
// eTag (optional): Dashboard Widgets Version
func (client Client) ReplaceWidgets(widgets *[]Widget, project *string, dashboardId *uuid.UUID, team *string, eTag *string) (*WidgetsVersionedList, error) {
    if widgets == nil {
        return nil, errors.New("widgets is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    additionalHeaders := make(map[string]string)
    if eTag != nil {
        additionalHeaders["ETag"] = *eTag
    }
    body, marshalErr := json.Marshal(*widgets)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseBodyValue []Widget
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WidgetsVersionedList
    if err == nil {
        responseValue = &WidgetsVersionedList {
            Widgets: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// [Preview API] Perform a partial update of the specified widget.
// widget (required): Description of the widget changes to apply. All non-null fields will be replaced.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to update.
// team (optional): Team ID or team name
func (client Client) UpdateWidget(widget *Widget, project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Widget, error) {
    if widget == nil {
        return nil, errors.New("widget is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, errors.New("widgetId is a required parameter")
    }
    routeValues["widgetId"] = (*widgetId).String()

    body, marshalErr := json.Marshal(*widget)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update the supplied widgets on the dashboard using supplied state. State of existing Widgets not passed in the widget list is preserved.
// widgets (required): The set of widget states to update on the dashboard.
// project (required): Project ID or project name
// dashboardId (required): ID of the Dashboard to modify.
// team (optional): Team ID or team name
// eTag (optional): Dashboard Widgets Version
func (client Client) UpdateWidgets(widgets *[]Widget, project *string, dashboardId *uuid.UUID, team *string, eTag *string) (*WidgetsVersionedList, error) {
    if widgets == nil {
        return nil, errors.New("widgets is a required parameter")
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, errors.New("project is a required parameter")
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, errors.New("dashboardId is a required parameter")
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    additionalHeaders := make(map[string]string)
    if eTag != nil {
        additionalHeaders["ETag"] = *eTag
    }
    body, marshalErr := json.Marshal(*widgets)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", additionalHeaders)
    if err != nil {
        return nil, err
    }

    var responseBodyValue []Widget
    err = client.Client.UnmarshalBody(resp, &responseBodyValue)

    var responseValue *WidgetsVersionedList
    if err == nil {
        responseValue = &WidgetsVersionedList {
            Widgets: &responseBodyValue,
            ETag: &[]string{ resp.Header.Get("ETag") },
        }
    }

    return responseValue, err
}

// [Preview API] Get the widget metadata satisfying the specified contribution ID.
// contributionId (required): The ID of Contribution for the Widget
// project (optional): Project ID or project name
func (client Client) GetWidgetMetadata(contributionId *string, project *string) (*WidgetMetadataResponse, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if contributionId == nil || *contributionId == "" {
        return nil, errors.New("contributionId is a required parameter")
    }
    routeValues["contributionId"] = *contributionId

    locationId, _ := uuid.Parse("6b3628d3-e96f-4fc7-b176-50240b03b515")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WidgetMetadataResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all available widget metadata in alphabetical order.
// scope (required)
// project (optional): Project ID or project name
func (client Client) GetWidgetTypes(scope *string, project *string) (*WidgetTypesResponse, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if scope == nil {
        return nil, errors.New("scope is a required parameter")
    }
    queryParams.Add("$scope", *scope)
    locationId, _ := uuid.Parse("6b3628d3-e96f-4fc7-b176-50240b03b515")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WidgetTypesResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

