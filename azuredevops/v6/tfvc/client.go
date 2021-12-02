// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package tfvc

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/git"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("8aa40520-446d-40e6-89f6-9c9f9ce44c48")

type Client interface {
	// Create a new changeset.
	CreateChangeset(context.Context, CreateChangesetArgs) (*git.TfvcChangesetRef, error)
	// Returns changesets for a given list of changeset Ids.
	GetBatchedChangesets(context.Context, GetBatchedChangesetsArgs) (*[]git.TfvcChangesetRef, error)
	// Get a single branch hierarchy at the given path with parents or children as specified.
	GetBranch(context.Context, GetBranchArgs) (*git.TfvcBranch, error)
	// Get a collection of branch roots -- first-level children, branches with no parents.
	GetBranches(context.Context, GetBranchesArgs) (*[]git.TfvcBranch, error)
	// Get branch hierarchies below the specified scopePath
	GetBranchRefs(context.Context, GetBranchRefsArgs) (*[]git.TfvcBranchRef, error)
	// Retrieve a Tfvc Changeset
	GetChangeset(context.Context, GetChangesetArgs) (*git.TfvcChangeset, error)
	// Retrieve Tfvc changes for a given changeset.
	GetChangesetChanges(context.Context, GetChangesetChangesArgs) (*GetChangesetChangesResponseValue, error)
	// Retrieve Tfvc Changesets
	GetChangesets(context.Context, GetChangesetsArgs) (*[]git.TfvcChangesetRef, error)
	// Retrieves the work items associated with a particular changeset.
	GetChangesetWorkItems(context.Context, GetChangesetWorkItemsArgs) (*[]git.AssociatedWorkItem, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
	GetItem(context.Context, GetItemArgs) (*git.TfvcItem, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
	GetItemContent(context.Context, GetItemContentArgs) (io.ReadCloser, error)
	// Get a list of Tfvc items
	GetItems(context.Context, GetItemsArgs) (*[]git.TfvcItem, error)
	// Post for retrieving a set of items given a list of paths or a long path. Allows for specifying the recursionLevel and version descriptors for each path.
	GetItemsBatch(context.Context, GetItemsBatchArgs) (*[][]git.TfvcItem, error)
	// Post for retrieving a set of items given a list of paths or a long path. Allows for specifying the recursionLevel and version descriptors for each path.
	GetItemsBatchZip(context.Context, GetItemsBatchZipArgs) (io.ReadCloser, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
	GetItemText(context.Context, GetItemTextArgs) (io.ReadCloser, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
	GetItemZip(context.Context, GetItemZipArgs) (io.ReadCloser, error)
	// Get a single deep label.
	GetLabel(context.Context, GetLabelArgs) (*git.TfvcLabel, error)
	// Get items under a label.
	GetLabelItems(context.Context, GetLabelItemsArgs) (*[]git.TfvcItem, error)
	// Get a collection of shallow label references.
	GetLabels(context.Context, GetLabelsArgs) (*[]git.TfvcLabelRef, error)
	// Get a single deep shelveset.
	GetShelveset(context.Context, GetShelvesetArgs) (*git.TfvcShelveset, error)
	// Get changes included in a shelveset.
	GetShelvesetChanges(context.Context, GetShelvesetChangesArgs) (*[]git.TfvcChange, error)
	// Return a collection of shallow shelveset references.
	GetShelvesets(context.Context, GetShelvesetsArgs) (*[]git.TfvcShelvesetRef, error)
	// Get work items associated with a shelveset.
	GetShelvesetWorkItems(context.Context, GetShelvesetWorkItemsArgs) (*[]git.AssociatedWorkItem, error)
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

// Create a new changeset.
func (client *ClientImpl) CreateChangeset(ctx context.Context, args CreateChangesetArgs) (*git.TfvcChangesetRef, error) {
	if args.Changeset == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Changeset"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.Changeset)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("0bc8f0a4-6bfb-42a9-ba84-139da7b99c49")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue git.TfvcChangesetRef
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateChangeset function
type CreateChangesetArgs struct {
	// (required)
	Changeset *git.TfvcChangeset
	// (optional) Project ID or project name
	Project *string
}

// Returns changesets for a given list of changeset Ids.
func (client *ClientImpl) GetBatchedChangesets(ctx context.Context, args GetBatchedChangesetsArgs) (*[]git.TfvcChangesetRef, error) {
	if args.ChangesetsRequestData == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ChangesetsRequestData"}
	}
	body, marshalErr := json.Marshal(*args.ChangesetsRequestData)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("b7e7c173-803c-4fea-9ec8-31ee35c5502a")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcChangesetRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBatchedChangesets function
type GetBatchedChangesetsArgs struct {
	// (required) List of changeset IDs.
	ChangesetsRequestData *git.TfvcChangesetsRequestData
}

// Get a single branch hierarchy at the given path with parents or children as specified.
func (client *ClientImpl) GetBranch(ctx context.Context, args GetBranchArgs) (*git.TfvcBranch, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.IncludeParent != nil {
		queryParams.Add("includeParent", strconv.FormatBool(*args.IncludeParent))
	}
	if args.IncludeChildren != nil {
		queryParams.Add("includeChildren", strconv.FormatBool(*args.IncludeChildren))
	}
	locationId, _ := uuid.Parse("bc1f417e-239d-42e7-85e1-76e80cb2d6eb")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue git.TfvcBranch
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBranch function
type GetBranchArgs struct {
	// (required) Full path to the branch.  Default: $/ Examples: $/, $/MyProject, $/MyProject/SomeFolder.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Return the parent branch, if there is one. Default: False
	IncludeParent *bool
	// (optional) Return child branches, if there are any. Default: False
	IncludeChildren *bool
}

// Get a collection of branch roots -- first-level children, branches with no parents.
func (client *ClientImpl) GetBranches(ctx context.Context, args GetBranchesArgs) (*[]git.TfvcBranch, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.IncludeParent != nil {
		queryParams.Add("includeParent", strconv.FormatBool(*args.IncludeParent))
	}
	if args.IncludeChildren != nil {
		queryParams.Add("includeChildren", strconv.FormatBool(*args.IncludeChildren))
	}
	if args.IncludeDeleted != nil {
		queryParams.Add("includeDeleted", strconv.FormatBool(*args.IncludeDeleted))
	}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	locationId, _ := uuid.Parse("bc1f417e-239d-42e7-85e1-76e80cb2d6eb")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcBranch
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBranches function
type GetBranchesArgs struct {
	// (optional) Project ID or project name
	Project *string
	// (optional) Return the parent branch, if there is one. Default: False
	IncludeParent *bool
	// (optional) Return the child branches for each root branch. Default: False
	IncludeChildren *bool
	// (optional) Return deleted branches. Default: False
	IncludeDeleted *bool
	// (optional) Return links. Default: False
	IncludeLinks *bool
}

// Get branch hierarchies below the specified scopePath
func (client *ClientImpl) GetBranchRefs(ctx context.Context, args GetBranchRefsArgs) (*[]git.TfvcBranchRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.ScopePath == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "scopePath"}
	}
	queryParams.Add("scopePath", *args.ScopePath)
	if args.IncludeDeleted != nil {
		queryParams.Add("includeDeleted", strconv.FormatBool(*args.IncludeDeleted))
	}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	locationId, _ := uuid.Parse("bc1f417e-239d-42e7-85e1-76e80cb2d6eb")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcBranchRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBranchRefs function
type GetBranchRefsArgs struct {
	// (required) Full path to the branch.  Default: $/ Examples: $/, $/MyProject, $/MyProject/SomeFolder.
	ScopePath *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Return deleted branches. Default: False
	IncludeDeleted *bool
	// (optional) Return links. Default: False
	IncludeLinks *bool
}

// Retrieve a Tfvc Changeset
func (client *ClientImpl) GetChangeset(ctx context.Context, args GetChangesetArgs) (*git.TfvcChangeset, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.Id == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Id"}
	}
	routeValues["id"] = strconv.Itoa(*args.Id)

	queryParams := url.Values{}
	if args.MaxChangeCount != nil {
		queryParams.Add("maxChangeCount", strconv.Itoa(*args.MaxChangeCount))
	}
	if args.IncludeDetails != nil {
		queryParams.Add("includeDetails", strconv.FormatBool(*args.IncludeDetails))
	}
	if args.IncludeWorkItems != nil {
		queryParams.Add("includeWorkItems", strconv.FormatBool(*args.IncludeWorkItems))
	}
	if args.MaxCommentLength != nil {
		queryParams.Add("maxCommentLength", strconv.Itoa(*args.MaxCommentLength))
	}
	if args.IncludeSourceRename != nil {
		queryParams.Add("includeSourceRename", strconv.FormatBool(*args.IncludeSourceRename))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Orderby != nil {
		queryParams.Add("$orderby", *args.Orderby)
	}
	if args.SearchCriteria != nil {
		if args.SearchCriteria.ItemPath != nil {
			queryParams.Add("searchCriteria.itemPath", *args.SearchCriteria.ItemPath)
		}
		if args.SearchCriteria.Author != nil {
			queryParams.Add("searchCriteria.author", *args.SearchCriteria.Author)
		}
		if args.SearchCriteria.FromDate != nil {
			queryParams.Add("searchCriteria.fromDate", *args.SearchCriteria.FromDate)
		}
		if args.SearchCriteria.ToDate != nil {
			queryParams.Add("searchCriteria.toDate", *args.SearchCriteria.ToDate)
		}
		if args.SearchCriteria.FromId != nil {
			queryParams.Add("searchCriteria.fromId", strconv.Itoa(*args.SearchCriteria.FromId))
		}
		if args.SearchCriteria.ToId != nil {
			queryParams.Add("searchCriteria.toId", strconv.Itoa(*args.SearchCriteria.ToId))
		}
		if args.SearchCriteria.FollowRenames != nil {
			queryParams.Add("searchCriteria.followRenames", strconv.FormatBool(*args.SearchCriteria.FollowRenames))
		}
		if args.SearchCriteria.IncludeLinks != nil {
			queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*args.SearchCriteria.IncludeLinks))
		}
	}
	locationId, _ := uuid.Parse("0bc8f0a4-6bfb-42a9-ba84-139da7b99c49")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue git.TfvcChangeset
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetChangeset function
type GetChangesetArgs struct {
	// (required) Changeset Id to retrieve.
	Id *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Number of changes to return (maximum 100 changes) Default: 0
	MaxChangeCount *int
	// (optional) Include policy details and check-in notes in the response. Default: false
	IncludeDetails *bool
	// (optional) Include workitems. Default: false
	IncludeWorkItems *bool
	// (optional) Include details about associated work items in the response. Default: null
	MaxCommentLength *int
	// (optional) Include renames.  Default: false
	IncludeSourceRename *bool
	// (optional) Number of results to skip. Default: null
	Skip *int
	// (optional) The maximum number of results to return. Default: null
	Top *int
	// (optional) Results are sorted by ID in descending order by default. Use id asc to sort by ID in ascending order.
	Orderby *string
	// (optional) Following criteria available (.itemPath, .version, .versionType, .versionOption, .author, .fromId, .toId, .fromDate, .toDate) Default: null
	SearchCriteria *git.TfvcChangesetSearchCriteria
}

// Retrieve Tfvc changes for a given changeset.
func (client *ClientImpl) GetChangesetChanges(ctx context.Context, args GetChangesetChangesArgs) (*GetChangesetChangesResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Id != nil {
		routeValues["id"] = strconv.Itoa(*args.Id)
	}

	queryParams := url.Values{}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	locationId, _ := uuid.Parse("f32b86f2-15b9-4fe6-81b1-6f8938617ee5")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetChangesetChangesResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetChangesetChanges function
type GetChangesetChangesArgs struct {
	// (optional) ID of the changeset. Default: null
	Id *int
	// (optional) Number of results to skip. Default: null
	Skip *int
	// (optional) The maximum number of results to return. Default: null
	Top *int
}

// Return type for the GetChangesetChanges function
type GetChangesetChangesResponseValue struct {
	Value             []git.TfvcChange
	ContinuationToken string
}

// Retrieve Tfvc Changesets
func (client *ClientImpl) GetChangesets(ctx context.Context, args GetChangesetsArgs) (*[]git.TfvcChangesetRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.MaxCommentLength != nil {
		queryParams.Add("maxCommentLength", strconv.Itoa(*args.MaxCommentLength))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Orderby != nil {
		queryParams.Add("$orderby", *args.Orderby)
	}
	if args.SearchCriteria != nil {
		if args.SearchCriteria.ItemPath != nil {
			queryParams.Add("searchCriteria.itemPath", *args.SearchCriteria.ItemPath)
		}
		if args.SearchCriteria.Author != nil {
			queryParams.Add("searchCriteria.author", *args.SearchCriteria.Author)
		}
		if args.SearchCriteria.FromDate != nil {
			queryParams.Add("searchCriteria.fromDate", *args.SearchCriteria.FromDate)
		}
		if args.SearchCriteria.ToDate != nil {
			queryParams.Add("searchCriteria.toDate", *args.SearchCriteria.ToDate)
		}
		if args.SearchCriteria.FromId != nil {
			queryParams.Add("searchCriteria.fromId", strconv.Itoa(*args.SearchCriteria.FromId))
		}
		if args.SearchCriteria.ToId != nil {
			queryParams.Add("searchCriteria.toId", strconv.Itoa(*args.SearchCriteria.ToId))
		}
		if args.SearchCriteria.FollowRenames != nil {
			queryParams.Add("searchCriteria.followRenames", strconv.FormatBool(*args.SearchCriteria.FollowRenames))
		}
		if args.SearchCriteria.IncludeLinks != nil {
			queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*args.SearchCriteria.IncludeLinks))
		}
	}
	locationId, _ := uuid.Parse("0bc8f0a4-6bfb-42a9-ba84-139da7b99c49")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcChangesetRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetChangesets function
type GetChangesetsArgs struct {
	// (optional) Project ID or project name
	Project *string
	// (optional) Include details about associated work items in the response. Default: null
	MaxCommentLength *int
	// (optional) Number of results to skip. Default: null
	Skip *int
	// (optional) The maximum number of results to return. Default: null
	Top *int
	// (optional) Results are sorted by ID in descending order by default. Use id asc to sort by ID in ascending order.
	Orderby *string
	// (optional) Following criteria available (.itemPath, .version, .versionType, .versionOption, .author, .fromId, .toId, .fromDate, .toDate) Default: null
	SearchCriteria *git.TfvcChangesetSearchCriteria
}

// Retrieves the work items associated with a particular changeset.
func (client *ClientImpl) GetChangesetWorkItems(ctx context.Context, args GetChangesetWorkItemsArgs) (*[]git.AssociatedWorkItem, error) {
	routeValues := make(map[string]string)
	if args.Id != nil {
		routeValues["id"] = strconv.Itoa(*args.Id)
	}

	locationId, _ := uuid.Parse("64ae0bea-1d71-47c9-a9e5-fe73f5ea0ff4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.AssociatedWorkItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetChangesetWorkItems function
type GetChangesetWorkItemsArgs struct {
	// (optional) ID of the changeset.
	Id *int
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
func (client *ClientImpl) GetItem(ctx context.Context, args GetItemArgs) (*git.TfvcItem, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionOption != nil {
			queryParams.Add("versionDescriptor.versionOption", string(*args.VersionDescriptor.VersionOption))
		}
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue git.TfvcItem
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetItem function
type GetItemArgs struct {
	// (required) Version control path of an individual item to return.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) file name of item returned.
	FileName *string
	// (optional) If true, create a downloadable attachment.
	Download *bool
	// (optional) Version control path of a folder to return multiple items.
	ScopePath *string
	// (optional) None (just the item), or OneLevel (contents of a folder).
	RecursionLevel *git.VersionControlRecursionType
	// (optional) Version descriptor.  Default is null.
	VersionDescriptor *git.TfvcVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
func (client *ClientImpl) GetItemContent(ctx context.Context, args GetItemContentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionOption != nil {
			queryParams.Add("versionDescriptor.versionOption", string(*args.VersionDescriptor.VersionOption))
		}
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetItemContent function
type GetItemContentArgs struct {
	// (required) Version control path of an individual item to return.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) file name of item returned.
	FileName *string
	// (optional) If true, create a downloadable attachment.
	Download *bool
	// (optional) Version control path of a folder to return multiple items.
	ScopePath *string
	// (optional) None (just the item), or OneLevel (contents of a folder).
	RecursionLevel *git.VersionControlRecursionType
	// (optional) Version descriptor.  Default is null.
	VersionDescriptor *git.TfvcVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
}

// Get a list of Tfvc items
func (client *ClientImpl) GetItems(ctx context.Context, args GetItemsArgs) (*[]git.TfvcItem, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionOption != nil {
			queryParams.Add("versionDescriptor.versionOption", string(*args.VersionDescriptor.VersionOption))
		}
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
	}
	locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetItems function
type GetItemsArgs struct {
	// (optional) Project ID or project name
	Project *string
	// (optional) Version control path of a folder to return multiple items.
	ScopePath *string
	// (optional) None (just the item), or OneLevel (contents of a folder).
	RecursionLevel *git.VersionControlRecursionType
	// (optional) True to include links.
	IncludeLinks *bool
	// (optional)
	VersionDescriptor *git.TfvcVersionDescriptor
}

// Post for retrieving a set of items given a list of paths or a long path. Allows for specifying the recursionLevel and version descriptors for each path.
func (client *ClientImpl) GetItemsBatch(ctx context.Context, args GetItemsBatchArgs) (*[][]git.TfvcItem, error) {
	if args.ItemRequestData == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ItemRequestData"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.ItemRequestData)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("fe6f827b-5f64-480f-b8af-1eca3b80e833")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue [][]git.TfvcItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetItemsBatch function
type GetItemsBatchArgs struct {
	// (required)
	ItemRequestData *git.TfvcItemRequestData
	// (optional) Project ID or project name
	Project *string
}

// Post for retrieving a set of items given a list of paths or a long path. Allows for specifying the recursionLevel and version descriptors for each path.
func (client *ClientImpl) GetItemsBatchZip(ctx context.Context, args GetItemsBatchZipArgs) (io.ReadCloser, error) {
	if args.ItemRequestData == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ItemRequestData"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	body, marshalErr := json.Marshal(*args.ItemRequestData)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("fe6f827b-5f64-480f-b8af-1eca3b80e833")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "6.0", routeValues, nil, bytes.NewReader(body), "application/json", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetItemsBatchZip function
type GetItemsBatchZipArgs struct {
	// (required)
	ItemRequestData *git.TfvcItemRequestData
	// (optional) Project ID or project name
	Project *string
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
func (client *ClientImpl) GetItemText(ctx context.Context, args GetItemTextArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionOption != nil {
			queryParams.Add("versionDescriptor.versionOption", string(*args.VersionDescriptor.VersionOption))
		}
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "text/plain", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetItemText function
type GetItemTextArgs struct {
	// (required) Version control path of an individual item to return.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) file name of item returned.
	FileName *string
	// (optional) If true, create a downloadable attachment.
	Download *bool
	// (optional) Version control path of a folder to return multiple items.
	ScopePath *string
	// (optional) None (just the item), or OneLevel (contents of a folder).
	RecursionLevel *git.VersionControlRecursionType
	// (optional) Version descriptor.  Default is null.
	VersionDescriptor *git.TfvcVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
func (client *ClientImpl) GetItemZip(ctx context.Context, args GetItemZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionOption != nil {
			queryParams.Add("versionDescriptor.versionOption", string(*args.VersionDescriptor.VersionOption))
		}
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	locationId, _ := uuid.Parse("ba9fc436-9a38-4578-89d6-e4f3241f5040")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetItemZip function
type GetItemZipArgs struct {
	// (required) Version control path of an individual item to return.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) file name of item returned.
	FileName *string
	// (optional) If true, create a downloadable attachment.
	Download *bool
	// (optional) Version control path of a folder to return multiple items.
	ScopePath *string
	// (optional) None (just the item), or OneLevel (contents of a folder).
	RecursionLevel *git.VersionControlRecursionType
	// (optional) Version descriptor.  Default is null.
	VersionDescriptor *git.TfvcVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
}

// Get a single deep label.
func (client *ClientImpl) GetLabel(ctx context.Context, args GetLabelArgs) (*git.TfvcLabel, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.LabelId == nil || *args.LabelId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.LabelId"}
	}
	routeValues["labelId"] = *args.LabelId

	queryParams := url.Values{}
	if args.RequestData == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "requestData"}
	}
	if args.RequestData.LabelScope != nil {
		queryParams.Add("requestData.labelScope", *args.RequestData.LabelScope)
	}
	if args.RequestData.Name != nil {
		queryParams.Add("requestData.name", *args.RequestData.Name)
	}
	if args.RequestData.Owner != nil {
		queryParams.Add("requestData.owner", *args.RequestData.Owner)
	}
	if args.RequestData.ItemLabelFilter != nil {
		queryParams.Add("requestData.itemLabelFilter", *args.RequestData.ItemLabelFilter)
	}
	if args.RequestData.MaxItemCount != nil {
		queryParams.Add("requestData.maxItemCount", strconv.Itoa(*args.RequestData.MaxItemCount))
	}
	if args.RequestData.IncludeLinks != nil {
		queryParams.Add("requestData.includeLinks", strconv.FormatBool(*args.RequestData.IncludeLinks))
	}
	locationId, _ := uuid.Parse("a5d9bd7f-b661-4d0e-b9be-d9c16affae54")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue git.TfvcLabel
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLabel function
type GetLabelArgs struct {
	// (required) Unique identifier of label
	LabelId *string
	// (required) maxItemCount
	RequestData *git.TfvcLabelRequestData
	// (optional) Project ID or project name
	Project *string
}

// Get items under a label.
func (client *ClientImpl) GetLabelItems(ctx context.Context, args GetLabelItemsArgs) (*[]git.TfvcItem, error) {
	routeValues := make(map[string]string)
	if args.LabelId == nil || *args.LabelId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.LabelId"}
	}
	routeValues["labelId"] = *args.LabelId

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("06166e34-de17-4b60-8cd1-23182a346fda")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLabelItems function
type GetLabelItemsArgs struct {
	// (required) Unique identifier of label
	LabelId *string
	// (optional) Max number of items to return
	Top *int
	// (optional) Number of items to skip
	Skip *int
}

// Get a collection of shallow label references.
func (client *ClientImpl) GetLabels(ctx context.Context, args GetLabelsArgs) (*[]git.TfvcLabelRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.RequestData == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "requestData"}
	}
	if args.RequestData.LabelScope != nil {
		queryParams.Add("requestData.labelScope", *args.RequestData.LabelScope)
	}
	if args.RequestData.Name != nil {
		queryParams.Add("requestData.name", *args.RequestData.Name)
	}
	if args.RequestData.Owner != nil {
		queryParams.Add("requestData.owner", *args.RequestData.Owner)
	}
	if args.RequestData.ItemLabelFilter != nil {
		queryParams.Add("requestData.itemLabelFilter", *args.RequestData.ItemLabelFilter)
	}
	if args.RequestData.MaxItemCount != nil {
		queryParams.Add("requestData.maxItemCount", strconv.Itoa(*args.RequestData.MaxItemCount))
	}
	if args.RequestData.IncludeLinks != nil {
		queryParams.Add("requestData.includeLinks", strconv.FormatBool(*args.RequestData.IncludeLinks))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("a5d9bd7f-b661-4d0e-b9be-d9c16affae54")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcLabelRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLabels function
type GetLabelsArgs struct {
	// (required) labelScope, name, owner, and itemLabelFilter
	RequestData *git.TfvcLabelRequestData
	// (optional) Project ID or project name
	Project *string
	// (optional) Max number of labels to return, defaults to 100 when undefined
	Top *int
	// (optional) Number of labels to skip
	Skip *int
}

// Get a single deep shelveset.
func (client *ClientImpl) GetShelveset(ctx context.Context, args GetShelvesetArgs) (*git.TfvcShelveset, error) {
	queryParams := url.Values{}
	if args.ShelvesetId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "shelvesetId"}
	}
	queryParams.Add("shelvesetId", *args.ShelvesetId)
	if args.RequestData != nil {
		if args.RequestData.Name != nil {
			queryParams.Add("requestData.name", *args.RequestData.Name)
		}
		if args.RequestData.Owner != nil {
			queryParams.Add("requestData.owner", *args.RequestData.Owner)
		}
		if args.RequestData.MaxCommentLength != nil {
			queryParams.Add("requestData.maxCommentLength", strconv.Itoa(*args.RequestData.MaxCommentLength))
		}
		if args.RequestData.MaxChangeCount != nil {
			queryParams.Add("requestData.maxChangeCount", strconv.Itoa(*args.RequestData.MaxChangeCount))
		}
		if args.RequestData.IncludeDetails != nil {
			queryParams.Add("requestData.includeDetails", strconv.FormatBool(*args.RequestData.IncludeDetails))
		}
		if args.RequestData.IncludeWorkItems != nil {
			queryParams.Add("requestData.includeWorkItems", strconv.FormatBool(*args.RequestData.IncludeWorkItems))
		}
		if args.RequestData.IncludeLinks != nil {
			queryParams.Add("requestData.includeLinks", strconv.FormatBool(*args.RequestData.IncludeLinks))
		}
	}
	locationId, _ := uuid.Parse("e36d44fb-e907-4b0a-b194-f83f1ed32ad3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue git.TfvcShelveset
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetShelveset function
type GetShelvesetArgs struct {
	// (required) Shelveset's unique ID
	ShelvesetId *string
	// (optional) includeDetails, includeWorkItems, maxChangeCount, and maxCommentLength
	RequestData *git.TfvcShelvesetRequestData
}

// Get changes included in a shelveset.
func (client *ClientImpl) GetShelvesetChanges(ctx context.Context, args GetShelvesetChangesArgs) (*[]git.TfvcChange, error) {
	queryParams := url.Values{}
	if args.ShelvesetId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "shelvesetId"}
	}
	queryParams.Add("shelvesetId", *args.ShelvesetId)
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("dbaf075b-0445-4c34-9e5b-82292f856522")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcChange
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetShelvesetChanges function
type GetShelvesetChangesArgs struct {
	// (required) Shelveset's unique ID
	ShelvesetId *string
	// (optional) Max number of changes to return
	Top *int
	// (optional) Number of changes to skip
	Skip *int
}

// Return a collection of shallow shelveset references.
func (client *ClientImpl) GetShelvesets(ctx context.Context, args GetShelvesetsArgs) (*[]git.TfvcShelvesetRef, error) {
	queryParams := url.Values{}
	if args.RequestData != nil {
		if args.RequestData.Name != nil {
			queryParams.Add("requestData.name", *args.RequestData.Name)
		}
		if args.RequestData.Owner != nil {
			queryParams.Add("requestData.owner", *args.RequestData.Owner)
		}
		if args.RequestData.MaxCommentLength != nil {
			queryParams.Add("requestData.maxCommentLength", strconv.Itoa(*args.RequestData.MaxCommentLength))
		}
		if args.RequestData.MaxChangeCount != nil {
			queryParams.Add("requestData.maxChangeCount", strconv.Itoa(*args.RequestData.MaxChangeCount))
		}
		if args.RequestData.IncludeDetails != nil {
			queryParams.Add("requestData.includeDetails", strconv.FormatBool(*args.RequestData.IncludeDetails))
		}
		if args.RequestData.IncludeWorkItems != nil {
			queryParams.Add("requestData.includeWorkItems", strconv.FormatBool(*args.RequestData.IncludeWorkItems))
		}
		if args.RequestData.IncludeLinks != nil {
			queryParams.Add("requestData.includeLinks", strconv.FormatBool(*args.RequestData.IncludeLinks))
		}
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("e36d44fb-e907-4b0a-b194-f83f1ed32ad3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.TfvcShelvesetRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetShelvesets function
type GetShelvesetsArgs struct {
	// (optional) name, owner, and maxCommentLength
	RequestData *git.TfvcShelvesetRequestData
	// (optional) Max number of shelvesets to return
	Top *int
	// (optional) Number of shelvesets to skip
	Skip *int
}

// Get work items associated with a shelveset.
func (client *ClientImpl) GetShelvesetWorkItems(ctx context.Context, args GetShelvesetWorkItemsArgs) (*[]git.AssociatedWorkItem, error) {
	queryParams := url.Values{}
	if args.ShelvesetId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "shelvesetId"}
	}
	queryParams.Add("shelvesetId", *args.ShelvesetId)
	locationId, _ := uuid.Parse("a7a0c1c1-373e-425a-b031-a519474d743d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "6.0", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []git.AssociatedWorkItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetShelvesetWorkItems function
type GetShelvesetWorkItemsArgs struct {
	// (required) Shelveset's unique ID
	ShelvesetId *string
}
