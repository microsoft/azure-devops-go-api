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
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/searchshared"
)

// Defines the Board result that matched a Board search request.
type BoardResult struct {
	// Board Type of the board document.
	Boardtype *string `json:"boardtype,omitempty"`
	// Collection details of the baord document.
	Collection *searchshared.Collection `json:"collection,omitempty"`
	// Project details of the board document.
	Project *Project `json:"project,omitempty"`
	// Team details of the board document.
	Team *Team `json:"team,omitempty"`
}

// Defines a Board search request.
type BoardSearchRequest struct {
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
	// Options for sorting search results. If set to null, the results will be returned sorted by relevance. If more than one sort option is provided, the results are sorted in the order specified in the OrderBy.
	OrderBy *[]searchshared.SortOption `json:"$orderBy,omitempty"`
	// Number of results to be skipped.
	Skip *int `json:"$skip,omitempty"`
	// Number of results to be returned.
	Top *int `json:"$top,omitempty"`
	// Flag to opt for faceting in the result. Default behavior is false.
	IncludeFacets *bool `json:"includeFacets,omitempty"`
}

// Defines a Board search response item.
type BoardSearchResponse struct {
	// A dictionary storing an array of <code>Filter</code> object against each facet.
	Facets *map[string][]searchshared.Filter `json:"facets,omitempty"`
	// Numeric code indicating any additional information: 0 - Ok, 1 - Account is being reindexed, 2 - Account indexing has not started, 3 - Invalid Request, 4 - Prefix wildcard query not supported, 5 - MultiWords with code facet not supported, 6 - Account is being onboarded, 7 - Account is being onboarded or reindexed, 8 - Top value trimmed to maxresult allowed 9 - Branches are being indexed, 10 - Faceting not enabled, 11 - Work items not accessible, 19 - Phrase queries with code type filters not supported, 20 - Wildcard queries with code type filters not supported. Any other info code is used for internal purpose.
	InfoCode *int `json:"infoCode,omitempty"`
	// Total number of matched Board documents.
	Count *int `json:"count,omitempty"`
	// List of top matched Board documents.
	Results *[]BoardResult `json:"results,omitempty"`
}

// Information about the configured branch.
type BranchInfo struct {
	// Name of the indexed branch
	Name *string `json:"name,omitempty"`
}

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
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
	// Options for sorting search results. If set to null, the results will be returned sorted by relevance. If more than one sort option is provided, the results are sorted in the order specified in the OrderBy.
	OrderBy *[]searchshared.SortOption `json:"$orderBy,omitempty"`
	// Number of results to be skipped.
	Skip *int `json:"$skip,omitempty"`
	// Number of results to be returned.
	Top *int `json:"$top,omitempty"`
	// Flag to opt for faceting in the result. Default behavior is false.
	IncludeFacets *bool `json:"includeFacets,omitempty"`
	// Flag to opt for including matched code snippet in the result. Default behavior is false.
	IncludeSnippet *bool `json:"includeSnippet,omitempty"`
}

// Defines a code search response item.
type CodeSearchResponse struct {
	// A dictionary storing an array of <code>Filter</code> object against each facet.
	Facets *map[string][]searchshared.Filter `json:"facets,omitempty"`
	// Numeric code indicating any additional information: 0 - Ok, 1 - Account is being reindexed, 2 - Account indexing has not started, 3 - Invalid Request, 4 - Prefix wildcard query not supported, 5 - MultiWords with code facet not supported, 6 - Account is being onboarded, 7 - Account is being onboarded or reindexed, 8 - Top value trimmed to maxresult allowed 9 - Branches are being indexed, 10 - Faceting not enabled, 11 - Work items not accessible, 19 - Phrase queries with code type filters not supported, 20 - Wildcard queries with code type filters not supported. Any other info code is used for internal purpose.
	InfoCode *int `json:"infoCode,omitempty"`
	// Total number of matched files.
	Count *int `json:"count,omitempty"`
	// List of matched files.
	Results *[]CodeResult `json:"results,omitempty"`
}

// Defines the custom repository status.
type CustomRepositoryStatusResponse struct {
	// Repository Id.
	Id *uuid.UUID `json:"id,omitempty"`
	// List of indexed top level folders info.
	IndexedTopLevelFolders *[]DepotInfo `json:"indexedTopLevelFolders,omitempty"`
	// Repository Name.
	Name *string `json:"name,omitempty"`
}

// Information about the custom repository indexing freshness for configured branches and depots.
type DepotInfo struct {
	// Name of the indexed top level folder (depot).
	Name *string `json:"name,omitempty"`
}

// Defines the details of the project.
type Project struct {
	// Id of the project.
	Id *uuid.UUID `json:"id,omitempty"`
	// Name of the project.
	Name *string `json:"name,omitempty"`
}

// Defines the repository status.
type RepositoryStatusResponse struct {
	// Repository Id.
	Id *uuid.UUID `json:"id,omitempty"`
	// List of Indexed branches info.
	IndexedBranches *[]BranchInfo `json:"indexedBranches,omitempty"`
	// Repository Name.
	Name *string `json:"name,omitempty"`
}

// Defines the details of the team.
type Team struct {
	// Id of the team.
	Id *uuid.UUID `json:"id,omitempty"`
	// Name of the Team.
	Name *string `json:"name,omitempty"`
}

// Defines the TFVC repository status.
type TfvcRepositoryStatusResponse struct {
	// Repository Id.
	Id *uuid.UUID `json:"id,omitempty"`
	// List of Indexing Information for TFVC repository
	IndexingInformation *[]BranchInfo `json:"indexingInformation,omitempty"`
	// Repository Name.
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
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
	// Options for sorting search results. If set to null, the results will be returned sorted by relevance. If more than one sort option is provided, the results are sorted in the order specified in the OrderBy.
	OrderBy *[]searchshared.SortOption `json:"$orderBy,omitempty"`
	// Number of results to be skipped.
	Skip *int `json:"$skip,omitempty"`
	// Number of results to be returned.
	Top *int `json:"$top,omitempty"`
	// Flag to opt for faceting in the result. Default behavior is false.
	IncludeFacets *bool `json:"includeFacets,omitempty"`
}

// Defines a response item that is returned for a work item search request.
type WorkItemSearchResponse struct {
	// A dictionary storing an array of <code>Filter</code> object against each facet.
	Facets *map[string][]searchshared.Filter `json:"facets,omitempty"`
	// Numeric code indicating any additional information: 0 - Ok, 1 - Account is being reindexed, 2 - Account indexing has not started, 3 - Invalid Request, 4 - Prefix wildcard query not supported, 5 - MultiWords with code facet not supported, 6 - Account is being onboarded, 7 - Account is being onboarded or reindexed, 8 - Top value trimmed to maxresult allowed 9 - Branches are being indexed, 10 - Faceting not enabled, 11 - Work items not accessible, 19 - Phrase queries with code type filters not supported, 20 - Wildcard queries with code type filters not supported. Any other info code is used for internal purpose.
	InfoCode *int `json:"infoCode,omitempty"`
	// Total number of matched work items.
	Count *int `json:"count,omitempty"`
	// List of top matched work items.
	Results *[]WorkItemResult `json:"results,omitempty"`
}
