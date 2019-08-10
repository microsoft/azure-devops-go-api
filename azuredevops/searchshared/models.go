// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package searchshared

import (
	"github.com/google/uuid"
)

// Defines the details of the collection.
type Collection struct {
	// Name of the collection.
	Name *string `json:"name,omitempty"`
}

// Base contract for search request types without scroll support.
type EntitySearchRequest struct {
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
	// Options for sorting search results. If set to null, the results will be returned sorted by relevance. If more than one sort option is provided, the results are sorted in the order specified in the OrderBy.
	OrderBy *[]SortOption `json:"$orderBy,omitempty"`
	// Number of results to be skipped.
	Skip *int `json:"$skip,omitempty"`
	// Number of results to be returned.
	Top *int `json:"$top,omitempty"`
	// Flag to opt for faceting in the result. Default behavior is false.
	IncludeFacets *bool `json:"includeFacets,omitempty"`
}

// Base class for search request types.
type EntitySearchRequestBase struct {
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
}

// Defines the base contract for search response.
type EntitySearchResponse struct {
	// A dictionary storing an array of <code>Filter</code> object against each facet.
	Facets *map[string][]Filter `json:"facets,omitempty"`
	// Numeric code indicating any additional information: 0 - Ok, 1 - Account is being reindexed, 2 - Account indexing has not started, 3 - Invalid Request, 4 - Prefix wildcard query not supported, 5 - MultiWords with code facet not supported, 6 - Account is being onboarded, 7 - Account is being onboarded or reindexed, 8 - Top value trimmed to maxresult allowed 9 - Branches are being indexed, 10 - Faceting not enabled, 11 - Work items not accessible, 19 - Phrase queries with code type filters not supported, 20 - Wildcard queries with code type filters not supported. Any other info code is used for internal purpose.
	InfoCode *int `json:"infoCode,omitempty"`
}

// Defines the details of a feed.
type FeedInfo struct {
	// Id of the collection.
	CollectionId *string `json:"collectionId,omitempty"`
	// Name of the collection.
	CollectionName *string `json:"collectionName,omitempty"`
	// Id of the feed.
	FeedId *string `json:"feedId,omitempty"`
	// Name of the feed.
	FeedName *string `json:"feedName,omitempty"`
	// Latest matched version of package in this Feed.
	LatestMatchedVersion *string `json:"latestMatchedVersion,omitempty"`
	// Latest version of package in this Feed.
	LatestVersion *string `json:"latestVersion,omitempty"`
	// Url of package in this Feed.
	PackageUrl *string `json:"packageUrl,omitempty"`
	// List of views which contain the matched package.
	Views *[]string `json:"views,omitempty"`
}

// Describes a filter bucket item representing the total matches of search result, name and id.
type Filter struct {
	// Id of the filter bucket.
	Id *string `json:"id,omitempty"`
	// Name of the filter bucket.
	Name *string `json:"name,omitempty"`
	// Count of matches in the filter bucket.
	ResultCount *int `json:"resultCount,omitempty"`
}

// Describes the position of a piece of text in a document.
type Hit struct {
	// Gets or sets the start character offset of a piece of text.
	CharOffset *int `json:"charOffset,omitempty"`
	// Gets or sets the length of a piece of text.
	Length *int `json:"length,omitempty"`
}

// Standard info codes that we return from Query Pipeline to UX/Client as part of REST contracts
type InfoCodes string

type infoCodesValuesType struct {
	Ok                                       InfoCodes
	AccountIsBeingReindexed                  InfoCodes
	IndexingNotStarted                       InfoCodes
	InvalidRequest                           InfoCodes
	PrefixWildcardQueryNotSupported          InfoCodes
	MultiWordWithCodeFacetNotSupported       InfoCodes
	AccountIsBeingOnboarded                  InfoCodes
	TakeResultValueTrimmedToMaxResultAllowed InfoCodes
	BranchesAreBeingIndexed                  InfoCodes
	FacetingNotEnabledAtScaleUnit            InfoCodes
	WorkItemsNotAccessible                   InfoCodes
	EmptyQueryNotSupported                   InfoCodes
	OnlyWildcardQueryNotSupported            InfoCodes
	ZeroResultsWithWildcard                  InfoCodes
	ZeroResultsWithFilter                    InfoCodes
	ZeroResultsWithWildcardAndFilter         InfoCodes
	ZeroResultsWithNoWildcardNoFilter        InfoCodes
	PartialResultsDueToSearchRequestTimeout  InfoCodes
	PhraseQueriesWithCEFacetsNotSupported    InfoCodes
	WildcardQueriesWithCEFacetsNotSupported  InfoCodes
	ClearedScrollSearchRequestParam          InfoCodes
	InvalidScrollSearchRequestParam          InfoCodes
	StopWarmerRequests                       InfoCodes
}

var InfoCodesValues = infoCodesValuesType{
	// Everything ok with the result.
	Ok: "ok",
	// Account is being re-indexed. Do not use this for fault-in scenarios; use AccountIsBeingOnboarded instead.
	AccountIsBeingReindexed: "accountIsBeingReindexed",
	// Indexing is not started yet for the collection
	IndexingNotStarted: "indexingNotStarted",
	// Invalid request
	InvalidRequest: "invalidRequest",
	// Search text containing prefix wildcard code term is not supported.
	PrefixWildcardQueryNotSupported: "prefixWildcardQueryNotSupported",
	// Multi Word Search text with code facet is not supported.
	MultiWordWithCodeFacetNotSupported: "multiWordWithCodeFacetNotSupported",
	// Account is being onboarded. This is similar to AccountIsBeingReindexed except that this is used only when the collection is faulted-in for the first time in search service.
	AccountIsBeingOnboarded: "accountIsBeingOnboarded",
	// $top Value is more than the value allowed in one fetch. $top is truncated as specified value exceeds the limit.
	TakeResultValueTrimmedToMaxResultAllowed: "takeResultValueTrimmedToMaxResultAllowed",
	// One or more branches in the collection are being indexed.
	BranchesAreBeingIndexed: "branchesAreBeingIndexed",
	// IncludeFacets is true but facets support is disabled for this deployment.
	FacetingNotEnabledAtScaleUnit: "facetingNotEnabledAtScaleUnit",
	// User has no permissions.
	WorkItemsNotAccessible: "workItemsNotAccessible",
	// When Search query contains only operators (i.e. [, ], (, ), :, ", ?, *) it is converted to empty expression and no results are fetched for them. [Todo:ManasaP] [Task 1163307] Once we have suggestions supported as part of the response, remove this info code.
	EmptyQueryNotSupported: "emptyQueryNotSupported",
	// When Search query contains only wildcard chars (ex: ***, *?, ??)
	OnlyWildcardQueryNotSupported: "onlyWildcardQueryNotSupported",
	// When Search query fetches zero results with wildcard in it
	ZeroResultsWithWildcard: "zeroResultsWithWildcard",
	// When Search query fetches zero results and has filters
	ZeroResultsWithFilter: "zeroResultsWithFilter",
	// When Search query fetches zero results and has wildcard and filter
	ZeroResultsWithWildcardAndFilter: "zeroResultsWithWildcardAndFilter",
	// When Search query fetches zero results and has no wildcard or filter
	ZeroResultsWithNoWildcardNoFilter: "zeroResultsWithNoWildcardNoFilter",
	// When Search request times out and hence potentially gives partial results
	PartialResultsDueToSearchRequestTimeout: "partialResultsDueToSearchRequestTimeout",
	// Phrase queries with code facet is not supported.
	PhraseQueriesWithCEFacetsNotSupported: "phraseQueriesWithCEFacetsNotSupported",
	// Wildcard queries with code facet is not supported.
	WildcardQueriesWithCEFacetsNotSupported: "wildcardQueriesWithCEFacetsNotSupported",
	// When Scroll Search Request returns count of zero we will clear the scroll.
	ClearedScrollSearchRequestParam: "clearedScrollSearchRequestParam",
	// When Scroll Search Request has an invalid scroll id value.
	InvalidScrollSearchRequestParam: "invalidScrollSearchRequestParam",
	// For Stopping warmer requests
	StopWarmerRequests: "stopWarmerRequests",
}

// Defines the matched terms in the field of the package result.
type PackageHit struct {
	// Reference name of the highlighted field.
	FieldReferenceName *string `json:"fieldReferenceName,omitempty"`
	// Matched/highlighted snippets of the field.
	Highlights *[]string `json:"highlights,omitempty"`
}

// Defines the package result that matched a package search request.
type PackageResult struct {
	// Description of the package.
	Description *string `json:"description,omitempty"`
	// List of feeds which contain the matching package.
	Feeds *[]FeedInfo `json:"feeds,omitempty"`
	// List of highlighted fields for the match.
	Hits *[]PackageHit `json:"hits,omitempty"`
	// Id of the package.
	Id *string `json:"id,omitempty"`
	// Name of the package.
	Name *string `json:"name,omitempty"`
	// Type of the package.
	ProtocolType *string `json:"protocolType,omitempty"`
}

// Defines a package search request.
type PackageSearchRequest struct {
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
	// Options for sorting search results. If set to null, the results will be returned sorted by relevance. If more than one sort option is provided, the results are sorted in the order specified in the OrderBy.
	OrderBy *[]SortOption `json:"$orderBy,omitempty"`
	// Number of results to be skipped.
	Skip *int `json:"$skip,omitempty"`
	// Number of results to be returned.
	Top *int `json:"$top,omitempty"`
	// Flag to opt for faceting in the result. Default behavior is false.
	IncludeFacets *bool `json:"includeFacets,omitempty"`
}

type PackageSearchResponse struct {
	ActivityId *[]string                     `json:"activityId,omitempty"`
	Content    *PackageSearchResponseContent `json:"content,omitempty"`
}

// Defines a response item that is returned for a package search request.
type PackageSearchResponseContent struct {
	// A dictionary storing an array of <code>Filter</code> object against each facet.
	Facets *map[string][]Filter `json:"facets,omitempty"`
	// Numeric code indicating any additional information: 0 - Ok, 1 - Account is being reindexed, 2 - Account indexing has not started, 3 - Invalid Request, 4 - Prefix wildcard query not supported, 5 - MultiWords with code facet not supported, 6 - Account is being onboarded, 7 - Account is being onboarded or reindexed, 8 - Top value trimmed to maxresult allowed 9 - Branches are being indexed, 10 - Faceting not enabled, 11 - Work items not accessible, 19 - Phrase queries with code type filters not supported, 20 - Wildcard queries with code type filters not supported. Any other info code is used for internal purpose.
	InfoCode *int `json:"infoCode,omitempty"`
	// Total number of matched packages.
	Count *int `json:"count,omitempty"`
	// List of matched packages.
	Results *[]PackageResult `json:"results,omitempty"`
}

// Defines the details of the project.
type ProjectReference struct {
	// ID of the project.
	Id *uuid.UUID `json:"id,omitempty"`
	// Name of the project.
	Name *string `json:"name,omitempty"`
	// Visibility of the project.
	Visibility *string `json:"visibility,omitempty"`
}

// Defines the details of the repository.
type Repository struct {
	// Id of the repository.
	Id *string `json:"id,omitempty"`
	// Name of the repository.
	Name *string `json:"name,omitempty"`
	// Version control type of the result file.
	Type *VersionControlType `json:"type,omitempty"`
}

// Defines a scroll code search request.
type ScrollSearchRequest struct {
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
	// Scroll Id for scroll search query.
	ScrollId *string `json:"$scrollId,omitempty"`
	// Size of data to return for scroll search query. Min value is 201.
	ScrollSize *int `json:"$scrollSize,omitempty"`
}

// Defines how to sort the result.
type SortOption struct {
	// Field name on which sorting should be done.
	Field *string `json:"field,omitempty"`
	// Order (ASC/DESC) in which the results should be sorted.
	SortOrder *string `json:"sortOrder,omitempty"`
}

type SortOrder string

type sortOrderValuesType struct {
	Undefined  SortOrder
	Ascending  SortOrder
	Descending SortOrder
}

var SortOrderValues = sortOrderValuesType{
	Undefined:  "undefined",
	Ascending:  "ascending",
	Descending: "descending",
}

// Describes the details pertaining to a version of the result file.
type Version struct {
	// Name of the branch.
	BranchName *string `json:"branchName,omitempty"`
	// ChangeId in the given branch associated with this match.
	ChangeId *string `json:"changeId,omitempty"`
}

// Version control of the repository.
type VersionControlType string

type versionControlTypeValuesType struct {
	Git    VersionControlType
	Tfvc   VersionControlType
	Custom VersionControlType
}

var VersionControlTypeValues = versionControlTypeValuesType{
	Git:  "git",
	Tfvc: "tfvc",
	// For internal use.
	Custom: "custom",
}

// Defines the details of wiki.
type Wiki struct {
	// Id of the wiki.
	Id *string `json:"id,omitempty"`
	// Mapped path for the wiki.
	MappedPath *string `json:"mappedPath,omitempty"`
	// Name of the wiki.
	Name *string `json:"name,omitempty"`
	// Version for wiki.
	Version *string `json:"version,omitempty"`
}

// Defines the matched terms in the field of the wiki result.
type WikiHit struct {
	// Reference name of the highlighted field.
	FieldReferenceName *string `json:"fieldReferenceName,omitempty"`
	// Matched/highlighted snippets of the field.
	Highlights *[]string `json:"highlights,omitempty"`
}

// Defines the wiki result that matched a wiki search request.
type WikiResult struct {
	// Collection of the result file.
	Collection *Collection `json:"collection,omitempty"`
	// ContentId of the result file.
	ContentId *string `json:"contentId,omitempty"`
	// Name of the result file.
	FileName *string `json:"fileName,omitempty"`
	// Highlighted snippets of fields that match the search request. The list is sorted by relevance of the snippets.
	Hits *[]WikiHit `json:"hits,omitempty"`
	// Path at which result file is present.
	Path *string `json:"path,omitempty"`
	// Project details of the wiki document.
	Project *ProjectReference `json:"project,omitempty"`
	// Wiki information for the result.
	Wiki *Wiki `json:"wiki,omitempty"`
}

// Defines a wiki search request.
type WikiSearchRequest struct {
	// Filters to be applied. Set it to null if there are no filters to be applied.
	Filters *map[string][]string `json:"filters,omitempty"`
	// The search text.
	SearchText *string `json:"searchText,omitempty"`
	// Options for sorting search results. If set to null, the results will be returned sorted by relevance. If more than one sort option is provided, the results are sorted in the order specified in the OrderBy.
	OrderBy *[]SortOption `json:"$orderBy,omitempty"`
	// Number of results to be skipped.
	Skip *int `json:"$skip,omitempty"`
	// Number of results to be returned.
	Top *int `json:"$top,omitempty"`
	// Flag to opt for faceting in the result. Default behavior is false.
	IncludeFacets *bool `json:"includeFacets,omitempty"`
}

// Defines a wiki search response item.
type WikiSearchResponse struct {
	// A dictionary storing an array of <code>Filter</code> object against each facet.
	Facets *map[string][]Filter `json:"facets,omitempty"`
	// Numeric code indicating any additional information: 0 - Ok, 1 - Account is being reindexed, 2 - Account indexing has not started, 3 - Invalid Request, 4 - Prefix wildcard query not supported, 5 - MultiWords with code facet not supported, 6 - Account is being onboarded, 7 - Account is being onboarded or reindexed, 8 - Top value trimmed to maxresult allowed 9 - Branches are being indexed, 10 - Faceting not enabled, 11 - Work items not accessible, 19 - Phrase queries with code type filters not supported, 20 - Wildcard queries with code type filters not supported. Any other info code is used for internal purpose.
	InfoCode *int `json:"infoCode,omitempty"`
	// Total number of matched wiki documents.
	Count *int `json:"count,omitempty"`
	// List of top matched wiki documents.
	Results *[]WikiResult `json:"results,omitempty"`
}
