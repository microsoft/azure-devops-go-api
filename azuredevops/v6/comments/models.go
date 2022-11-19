// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package comments

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/webapi"
)

// Comment on an artifact like Work Item or Wiki, etc.
type Comment struct {
	Url *string `json:"url,omitempty"`
	// The id of the artifact this comment belongs to
	ArtifactId *string `json:"artifactId,omitempty"`
	// IdentityRef of the creator of the comment.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// The creation date of the comment.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The id assigned to the comment.
	Id *int `json:"id,omitempty"`
	// Indicates if the comment has been deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// The mentions of the comment.
	Mentions *[]CommentMention `json:"mentions,omitempty"`
	// IdentityRef of the user who last modified the comment.
	ModifiedBy *webapi.IdentityRef `json:"modifiedBy,omitempty"`
	// The last modification date of the comment.
	ModifiedDate *azuredevops.Time `json:"modifiedDate,omitempty"`
	// The comment id of the parent comment, if any
	ParentId *int `json:"parentId,omitempty"`
	// The reactions on the comment.
	Reactions *[]CommentReaction `json:"reactions,omitempty"`
	// The rendered text of the comment
	RenderedText *string `json:"renderedText,omitempty"`
	// Replies for this comment
	Replies *CommentList `json:"replies,omitempty"`
	// Indicates the current state of the comment
	State *CommentState `json:"state,omitempty"`
	// The plaintext/markdown version of the comment
	Text *string `json:"text,omitempty"`
	// The current version of the comment
	Version *int `json:"version,omitempty"`
}

// Represents an attachment to a comment.
type CommentAttachment struct {
	Url *string `json:"url,omitempty"`
	// IdentityRef of the creator of the attachment.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// The creation date of the attachment.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// Unique Id of the attachment.
	Id *uuid.UUID `json:"id,omitempty"`
}

// Represents a request to create a work item comment.
type CommentCreateParameters struct {
	// Optional CommentId of the parent in order to add a reply for an existing comment
	ParentId *int    `json:"parentId,omitempty"`
	Text     *string `json:"text,omitempty"`
}

// [Flags] Specifies the additional data retrieval options for comments.
type CommentExpandOptions string

type commentExpandOptionsValuesType struct {
	None             CommentExpandOptions
	Reactions        CommentExpandOptions
	RenderedText     CommentExpandOptions
	RenderedTextOnly CommentExpandOptions
	Children         CommentExpandOptions
	All              CommentExpandOptions
}

var CommentExpandOptionsValues = commentExpandOptionsValuesType{
	// Include comments only, no mentions, reactions or rendered text
	None: "none",
	// Include comment reactions
	Reactions: "reactions",
	// Include the rendered text (html) in addition to markdown text
	RenderedText:     "renderedText",
	RenderedTextOnly: "renderedTextOnly",
	// If specified, then responses will be expanded in the results
	Children: "children",
	// Expand everything including Reactions, Mentions and also include RenderedText (HTML) for markdown comments
	All: "all",
}

// Format of the comment. Ex. Markdown, Html.
type CommentFormat string

type commentFormatValuesType struct {
	Markdown CommentFormat
	Html     CommentFormat
}

var CommentFormatValues = commentFormatValuesType{
	Markdown: "markdown",
	Html:     "html",
}

// Represents a list of comments.
type CommentList struct {
	Url *string `json:"url,omitempty"`
	// List of comments in the current batch.
	Comments *[]Comment `json:"comments,omitempty"`
	// A string token that can be used to retrieving next page of comments if available. Otherwise null.
	ContinuationToken *string `json:"continuationToken,omitempty"`
	// The count of comments in the current batch.
	Count *int `json:"count,omitempty"`
	// Uri to the next page of comments if it is available. Otherwise null.
	NextPage *string `json:"nextPage,omitempty"`
	// Total count of comments on a work item.
	TotalCount *int `json:"totalCount,omitempty"`
}

// Contains information about various artifacts mentioned in the comment
type CommentMention struct {
	Url *string `json:"url,omitempty"`
	// Id of the artifact this mention belongs to
	ArtifactId *string `json:"artifactId,omitempty"`
	// Id of the comment associated with this mention. Nullable to support legacy mentions which can potentially have null commentId
	CommentId *int `json:"commentId,omitempty"`
	// Value of the mentioned artifact. Expected Value varies by CommentMentionType: Person:         VSID associated with the identity Work Item:      ID of the work item Pull Request:   ID of the Pull Request
	MentionedArtifact *string `json:"mentionedArtifact,omitempty"`
	// The context which represent where this mentioned was parsed from
	Type *CommentMentionType `json:"type,omitempty"`
}

type CommentMentionType string

type commentMentionTypeValuesType struct {
	Person      CommentMentionType
	WorkItem    CommentMentionType
	PullRequest CommentMentionType
}

var CommentMentionTypeValues = commentMentionTypeValuesType{
	// An identity was mentioned by using the format @{VSID}
	Person: "person",
	// A work item was mentioned by using the format #{Work Item ID}
	WorkItem: "workItem",
	// A Pull Request was mentioned by using the format !{PR Number}
	PullRequest: "pullRequest",
}

// Contains information about comment reaction for a particular reaction type.
type CommentReaction struct {
	Url *string `json:"url,omitempty"`
	// The id of the comment this reaction belongs to.
	CommentId *int `json:"commentId,omitempty"`
	// Total number of reactions for the CommentReactionType.
	Count *int `json:"count,omitempty"`
	// Flag to indicate if the current user has engaged on this particular EngagementType (e.g. if they liked the associated comment).
	IsCurrentUserEngaged *bool `json:"isCurrentUserEngaged,omitempty"`
	// Type of the reaction.
	Type *CommentReactionType `json:"type,omitempty"`
}

// Represents different reaction types for a comment
type CommentReactionType string

type commentReactionTypeValuesType struct {
	Like     CommentReactionType
	Dislike  CommentReactionType
	Heart    CommentReactionType
	Hooray   CommentReactionType
	Smile    CommentReactionType
	Confused CommentReactionType
}

var CommentReactionTypeValues = commentReactionTypeValuesType{
	Like:     "like",
	Dislike:  "dislike",
	Heart:    "heart",
	Hooray:   "hooray",
	Smile:    "smile",
	Confused: "confused",
}

// Base class for comment resource references
type CommentResourceReference struct {
	Url *string `json:"url,omitempty"`
}

type CommentSortOrder string

type commentSortOrderValuesType struct {
	Asc  CommentSortOrder
	Desc CommentSortOrder
}

var CommentSortOrderValues = commentSortOrderValuesType{
	// The results will be sorted in Ascending order.
	Asc: "asc",
	// The results will be sorted in Descending order.
	Desc: "desc",
}

// Represents the possible comment states.
type CommentState string

type commentStateValuesType struct {
	Active   CommentState
	Resolved CommentState
	Closed   CommentState
}

var CommentStateValues = commentStateValuesType{
	Active:   "active",
	Resolved: "resolved",
	Closed:   "closed",
}

// Represents a request to update a comment.
type CommentUpdateParameters struct {
	// Set the current state of the comment
	State *CommentState `json:"state,omitempty"`
	// The updated text of the comment
	Text *string `json:"text,omitempty"`
}

// Represents a specific version of a comment on a work item.
type CommentVersion struct {
	Url *string `json:"url,omitempty"`
	// IdentityRef of the creator of the comment.
	CreatedBy *webapi.IdentityRef `json:"createdBy,omitempty"`
	// The creation date of the comment.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// The id assigned to the comment.
	Id *int `json:"id,omitempty"`
	// Indicates if the comment has been deleted at this version.
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// IdentityRef of the user who modified the comment at this version.
	ModifiedBy *webapi.IdentityRef `json:"modifiedBy,omitempty"`
	// The modification date of the comment for this version.
	ModifiedDate *azuredevops.Time `json:"modifiedDate,omitempty"`
	// The rendered content of the comment at this version.
	RenderedText *string `json:"renderedText,omitempty"`
	// Indicates the current state of the comment
	State *CommentState `json:"state,omitempty"`
	// The text of the comment at this version.
	Text *string `json:"text,omitempty"`
	// The version number.
	Version *int `json:"version,omitempty"`
}
