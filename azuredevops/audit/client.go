// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package audit

import (
	"context"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("94ff054d-5ee1-413d-9341-3f4a7827de2e")

type Client interface {
	// [Preview API] Downloads audit log entries.
	DownloadLog(context.Context, DownloadLogArgs) (io.ReadCloser, error)
	// [Preview API] Queries audit log entries
	QueryLog(context.Context, QueryLogArgs) (*AuditLogQueryResult, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// [Preview API] Downloads audit log entries.
func (client *ClientImpl) DownloadLog(ctx context.Context, args DownloadLogArgs) (io.ReadCloser, error) {
	queryParams := url.Values{}
	if args.Format == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "format"}
	}
	queryParams.Add("format", *args.Format)
	if args.StartTime != nil {
		queryParams.Add("startTime", (*args.StartTime).AsQueryParameter())
	}
	if args.EndTime != nil {
		queryParams.Add("endTime", (*args.EndTime).AsQueryParameter())
	}
	locationId, _ := uuid.Parse("b7b98a76-04e8-4f4d-ac72-9d46492caaac")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the DownloadLog function
type DownloadLogArgs struct {
	// (required) File format for download. Can be "json" or "csv".
	Format *string
	// (optional) Start time of download window. Optional
	StartTime *azuredevops.Time
	// (optional) End time of download window. Optional
	EndTime *azuredevops.Time
}

// [Preview API] Queries audit log entries
func (client *ClientImpl) QueryLog(ctx context.Context, args QueryLogArgs) (*AuditLogQueryResult, error) {
	queryParams := url.Values{}
	if args.StartTime != nil {
		queryParams.Add("startTime", (*args.StartTime).AsQueryParameter())
	}
	if args.EndTime != nil {
		queryParams.Add("endTime", (*args.EndTime).AsQueryParameter())
	}
	if args.BatchSize != nil {
		queryParams.Add("batchSize", strconv.Itoa(*args.BatchSize))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	if args.SkipAggregation != nil {
		queryParams.Add("skipAggregation", strconv.FormatBool(*args.SkipAggregation))
	}
	locationId, _ := uuid.Parse("4e5fa14f-7097-4b73-9c85-00abc7353c61")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue AuditLogQueryResult
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the QueryLog function
type QueryLogArgs struct {
	// (optional) Start time of download window. Optional
	StartTime *azuredevops.Time
	// (optional) End time of download window. Optional
	EndTime *azuredevops.Time
	// (optional) Max number of results to return. Optional
	BatchSize *int
	// (optional) Token used for returning next set of results from previous query. Optional
	ContinuationToken *string
	// (optional) Skips aggregating events and leaves them as individual entries instead. By default events are aggregated. Event types that are aggregated: AuditLog.AccessLog.
	SkipAggregation *bool
}
