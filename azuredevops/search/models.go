// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package search

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/searchshared"
)

// Defines the code result containing information of the searched files and its metadata.
type CodeResult struct {
	// Collection of the result file.
	Collection *searchshared.Collection `json:"collection,omitempty"`
	// ContentId of the result file.
	ContentId *string `json:"contentId,omitempty"`
	// Name of the result file.
	FileName *string `json:"fileName,omitempty"`
	// Dictionary of field to hit offsets in the result file. Key identifies the area in which hits were found, for ex: file content/file name etc.
	Matches *map[string][]searchshared.Hit `json:"matches,omitempty"`
	// Path at which result file is present.
	Path *string `json:"path,omitempty"`
	// Project of the result file.
	Project *Project `json:"project,omitempty"`
	// Repository of the result file.
	Repository *searchshared.Repository `json:"repository,omitempty"`
	// Versions of the result file.
	Versions *[]searchshared.Version `json:"versions,omitempty"`
}

// Defines a code search request.
type CodeSearchRequest struct {
}

// Defines a code search response item.
type CodeSearchResponse struct {
	// Total number of matched files.
	Count *int `json:"count,omitempty"`
	// List of matched files.
	Results *[]CodeResult `json:"results,omitempty"`
}

// Defines the details of the project.
type Project struct {
	// Id of the project.
	Id *uuid.UUID `json:"id,omitempty"`
	// Name of the project.
	Name *string `json:"name,omitempty"`
}

// Defines the matched terms in the field of the work item result.
type WorkItemHit struct {
	// Reference name of the highlighted field.
	FieldReferenceName *string `json:"fieldReferenceName,omitempty"`
	// Matched/highlighted snippets of the field.
	Highlights *[]string `json:"highlights,omitempty"`
}

// Defines the work item result that matched a work item search request.
type WorkItemResult struct {
	// A standard set of work item fields and their values.
	Fields *map[string]string `json:"fields,omitempty"`
	// Highlighted snippets of fields that match the search request. The list is sorted by relevance of the snippets.
	Hits *[]WorkItemHit `json:"hits,omitempty"`
	// Project details of the work item.
	Project *Project `json:"project,omitempty"`
	// Reference to the work item.
	Url *string `json:"url,omitempty"`
}

// Defines a work item search request.
type WorkItemSearchRequest struct {
}

// Defines a response item that is returned for a work item search request.
type WorkItemSearchResponse struct {
	// Total number of matched work items.
	Count *int `json:"count,omitempty"`
	// List of top matched work items.
	Results *[]WorkItemResult `json:"results,omitempty"`
}
