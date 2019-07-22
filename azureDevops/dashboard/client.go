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
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("31c84e0a-3ece-48fd-a29d-100849af99ba")

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

// [Preview API] Create the supplied dashboard.
// ctx
// dashboard (required): The initial state of the dashboard
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) CreateDashboard(ctx context.Context, dashboard *Dashboard, project *string, team *string) (*Dashboard, error) {
    if dashboard == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboard"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
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
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete a dashboard given its ID. This also deletes the widgets associated with this dashboard.
// ctx
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard to delete.
// team (optional): Team ID or team name
func (client Client) DeleteDashboard(ctx context.Context, project *string, dashboardId *uuid.UUID, team *string) error {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// [Preview API] Get a dashboard by its ID.
// ctx
// project (required): Project ID or project name
// dashboardId (required)
// team (optional): Team ID or team name
func (client Client) GetDashboard(ctx context.Context, project *string, dashboardId *uuid.UUID, team *string) (*Dashboard, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get a list of dashboards.
// ctx
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) GetDashboards(ctx context.Context, project *string, team *string) (*DashboardGroup, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }

    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DashboardGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Replace configuration for the specified dashboard. Replaces Widget list on Dashboard, only if property is supplied.
// ctx
// dashboard (required): The Configuration of the dashboard to replace.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard to replace.
// team (optional): Team ID or team name
func (client Client) ReplaceDashboard(ctx context.Context, dashboard *Dashboard, project *string, dashboardId *uuid.UUID, team *string) (*Dashboard, error) {
    if dashboard == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboard"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    body, marshalErr := json.Marshal(*dashboard)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("454b3e51-2e6e-48d4-ad81-978154089351")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update the name and position of dashboards in the supplied group, and remove omitted dashboards. Does not modify dashboard content.
// ctx
// group (required)
// project (required): Project ID or project name
// team (optional): Team ID or team name
func (client Client) ReplaceDashboards(ctx context.Context, group *DashboardGroup, project *string, team *string) (*DashboardGroup, error) {
    if group == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "group"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
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
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue DashboardGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Create a widget on the specified dashboard.
// ctx
// widget (required): State of the widget to add
// project (required): Project ID or project name
// dashboardId (required): ID of dashboard the widget will be added to.
// team (optional): Team ID or team name
func (client Client) CreateWidget(ctx context.Context, widget *Widget, project *string, dashboardId *uuid.UUID, team *string) (*Widget, error) {
    if widget == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widget"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    body, marshalErr := json.Marshal(*widget)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Delete the specified widget.
// ctx
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to update.
// team (optional): Team ID or team name
func (client Client) DeleteWidget(ctx context.Context, project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Dashboard, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widgetId"} 
    }
    routeValues["widgetId"] = (*widgetId).String()

    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Dashboard
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get the current state of the specified widget.
// ctx
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to read.
// team (optional): Team ID or team name
func (client Client) GetWidget(ctx context.Context, project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Widget, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widgetId"} 
    }
    routeValues["widgetId"] = (*widgetId).String()

    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get widgets contained on the specified dashboard.
// ctx
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard to read.
// team (optional): Team ID or team name
// eTag (optional): Dashboard Widgets Version
func (client Client) GetWidgets(ctx context.Context, project *string, dashboardId *uuid.UUID, team *string, eTag *string) (*WidgetsVersionedList, error) {
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()

    additionalHeaders := make(map[string]string)
    if eTag != nil {
        additionalHeaders["ETag"] = *eTag
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", routeValues, nil, nil, "", "application/json", additionalHeaders)
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
// ctx
// widget (required): State to be written for the widget.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to update.
// team (optional): Team ID or team name
func (client Client) ReplaceWidget(ctx context.Context, widget *Widget, project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Widget, error) {
    if widget == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widget"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widgetId"} 
    }
    routeValues["widgetId"] = (*widgetId).String()

    body, marshalErr := json.Marshal(*widget)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Replace the widgets on specified dashboard with the supplied widgets.
// ctx
// widgets (required): Revised state of widgets to store for the dashboard.
// project (required): Project ID or project name
// dashboardId (required): ID of the Dashboard to modify.
// team (optional): Team ID or team name
// eTag (optional): Dashboard Widgets Version
func (client Client) ReplaceWidgets(ctx context.Context, widgets *[]Widget, project *string, dashboardId *uuid.UUID, team *string, eTag *string) (*WidgetsVersionedList, error) {
    if widgets == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widgets"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
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
    resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", additionalHeaders)
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
// ctx
// widget (required): Description of the widget changes to apply. All non-null fields will be replaced.
// project (required): Project ID or project name
// dashboardId (required): ID of the dashboard containing the widget.
// widgetId (required): ID of the widget to update.
// team (optional): Team ID or team name
func (client Client) UpdateWidget(ctx context.Context, widget *Widget, project *string, dashboardId *uuid.UUID, widgetId *uuid.UUID, team *string) (*Widget, error) {
    if widget == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widget"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
    }
    routeValues["dashboardId"] = (*dashboardId).String()
    if widgetId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widgetId"} 
    }
    routeValues["widgetId"] = (*widgetId).String()

    body, marshalErr := json.Marshal(*widget)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bdcff53a-8355-4172-a00a-40497ea23afc")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Widget
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Update the supplied widgets on the dashboard using supplied state. State of existing Widgets not passed in the widget list is preserved.
// ctx
// widgets (required): The set of widget states to update on the dashboard.
// project (required): Project ID or project name
// dashboardId (required): ID of the Dashboard to modify.
// team (optional): Team ID or team name
// eTag (optional): Dashboard Widgets Version
func (client Client) UpdateWidgets(ctx context.Context, widgets *[]Widget, project *string, dashboardId *uuid.UUID, team *string, eTag *string) (*WidgetsVersionedList, error) {
    if widgets == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "widgets"}
    }
    routeValues := make(map[string]string)
    if project == nil || *project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *project
    if team != nil && *team != "" {
        routeValues["team"] = *team
    }
    if dashboardId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "dashboardId"} 
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
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", additionalHeaders)
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
// ctx
// contributionId (required): The ID of Contribution for the Widget
// project (optional): Project ID or project name
func (client Client) GetWidgetMetadata(ctx context.Context, contributionId *string, project *string) (*WidgetMetadataResponse, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }
    if contributionId == nil || *contributionId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "contributionId"} 
    }
    routeValues["contributionId"] = *contributionId

    locationId, _ := uuid.Parse("6b3628d3-e96f-4fc7-b176-50240b03b515")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WidgetMetadataResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// [Preview API] Get all available widget metadata in alphabetical order.
// ctx
// scope (required)
// project (optional): Project ID or project name
func (client Client) GetWidgetTypes(ctx context.Context, scope *WidgetScope, project *string) (*WidgetTypesResponse, error) {
    routeValues := make(map[string]string)
    if project != nil && *project != "" {
        routeValues["project"] = *project
    }

    queryParams := url.Values{}
    if scope == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "scope"}
    }
    queryParams.Add("$scope", string(*scope))
    locationId, _ := uuid.Parse("6b3628d3-e96f-4fc7-b176-50240b03b515")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue WidgetTypesResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

