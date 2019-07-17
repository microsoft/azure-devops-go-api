// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package release

import (
    "github.com/google/uuid"
    "math/big"
    "time"
)

type AgentArtifactDefinition struct {
    // Gets or sets the artifact definition alias.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the artifact type.
    ArtifactType *AgentArtifactType `json:"artifactType,omitempty"`
    // Gets or sets the artifact definition details.
    Details *string `json:"details,omitempty"`
    // Gets or sets the name of artifact definition.
    Name *string `json:"name,omitempty"`
    // Gets or sets the version of artifact definition.
    Version *string `json:"version,omitempty"`
}

type AgentArtifactType string

type AgentBasedDeployPhase struct {
    // Gets and sets the agent job deployment input
    DeploymentInput *AgentDeploymentInput `json:"deploymentInput,omitempty"`
}

type AgentDeploymentInput struct {
    // Gets or sets the job condition.
    Condition *string `json:"condition,omitempty"`
    // Gets or sets the job cancel timeout in minutes for deployment which are cancelled by user for this release environment.
    JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
    // Gets or sets the override inputs.
    OverrideInputs *map[string]string `json:"overrideInputs,omitempty"`
    // Gets or sets the job execution timeout in minutes for deployment which are queued against this release environment.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
    // Artifacts that downloaded during job execution.
    ArtifactsDownloadInput *ArtifactsDownloadInput `json:"artifactsDownloadInput,omitempty"`
    // List demands that needs to meet to execute the job.
    Demands *[]interface{} `json:"demands,omitempty"`
    // Indicates whether to include access token in deployment job or not.
    EnableAccessToken *bool `json:"enableAccessToken,omitempty"`
    // Id of the pool on which job get executed.
    QueueId *int `json:"queueId,omitempty"`
    // Indicates whether artifacts downloaded while job execution or not.
    SkipArtifactsDownload *bool `json:"skipArtifactsDownload,omitempty"`
    // Specification for an agent on which a job gets executed.
    AgentSpecification *AgentSpecification `json:"agentSpecification,omitempty"`
    // Gets or sets the image ID.
    ImageId *int `json:"imageId,omitempty"`
    // Gets or sets the parallel execution input.
    ParallelExecution *ExecutionInput `json:"parallelExecution,omitempty"`
}

// Represents a reference to an agent queue.
type AgentPoolQueueReference struct {
    // An alias to be used when referencing the resource.
    Alias *string `json:"alias,omitempty"`
    // The ID of the queue.
    Id *int `json:"id,omitempty"`
}

// Specification of the agent defined by the pool provider.
type AgentSpecification struct {
    // Agent specification unique identifier.
    Identifier *string `json:"identifier,omitempty"`
}

type ApprovalExecutionOrder string

type ApprovalFilters string

type ApprovalOptions struct {
    // Specify whether the approval can be skipped if the same approver approved the previous stage.
    AutoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped *bool `json:"autoTriggeredAndPreviousEnvironmentApprovedCanBeSkipped,omitempty"`
    // Specify whether revalidate identity of approver before completing the approval.
    EnforceIdentityRevalidation *bool `json:"enforceIdentityRevalidation,omitempty"`
    // Approvals execution order.
    ExecutionOrder *ApprovalExecutionOrder `json:"executionOrder,omitempty"`
    // Specify whether the user requesting a release or deployment should allow to approver.
    ReleaseCreatorCanBeApprover *bool `json:"releaseCreatorCanBeApprover,omitempty"`
    // The number of approvals required to move release forward. '0' means all approvals required.
    RequiredApproverCount *int `json:"requiredApproverCount,omitempty"`
    // Approval timeout. Approval default timeout is 30 days. Maximum allowed timeout is 365 days. '0' means default timeout i.e 30 days.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
}

type ApprovalStatus string

type ApprovalType string

type Artifact struct {
    // Gets or sets alias.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets definition reference. e.g. {"project":{"id":"fed755ea-49c5-4399-acea-fd5b5aa90a6c","name":"myProject"},"definition":{"id":"1","name":"mybuildDefinition"},"connection":{"id":"1","name":"myConnection"}}.
    DefinitionReference *map[string]ArtifactSourceReference `json:"definitionReference,omitempty"`
    // Indicates whether artifact is primary or not.
    IsPrimary *bool `json:"isPrimary,omitempty"`
    // Indicates whether artifact is retained by release or not.
    IsRetained *bool `json:"isRetained,omitempty"`
    // Deprecated: This property is deprecated use Alias instead and remove all its references
    SourceId *string `json:"sourceId,omitempty"`
    // Gets or sets type. It can have value as 'Build', 'Jenkins', 'GitHub', 'Nuget', 'Team Build (external)', 'ExternalTFSBuild', 'Git', 'TFVC', 'ExternalTfsXamlBuild'.
    Type *string `json:"type,omitempty"`
}

type ArtifactContributionDefinition struct {
    ArtifactTriggerConfiguration *ArtifactTriggerConfiguration `json:"artifactTriggerConfiguration,omitempty"`
    ArtifactType *string `json:"artifactType,omitempty"`
    ArtifactTypeStreamMapping *map[string]string `json:"artifactTypeStreamMapping,omitempty"`
    BrowsableArtifactTypeMapping *map[string]string `json:"browsableArtifactTypeMapping,omitempty"`
    DataSourceBindings *[]DataSourceBinding `json:"dataSourceBindings,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    DownloadTaskId *string `json:"downloadTaskId,omitempty"`
    EndpointTypeId *string `json:"endpointTypeId,omitempty"`
    InputDescriptors *[]InputDescriptor `json:"inputDescriptors,omitempty"`
    IsCommitsTraceabilitySupported *bool `json:"isCommitsTraceabilitySupported,omitempty"`
    IsWorkitemsTraceabilitySupported *bool `json:"isWorkitemsTraceabilitySupported,omitempty"`
    Name *string `json:"name,omitempty"`
    TaskInputMapping *map[string]string `json:"taskInputMapping,omitempty"`
    UniqueSourceIdentifier *string `json:"uniqueSourceIdentifier,omitempty"`
}

type ArtifactDownloadInputBase struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the name of artifact definition. Valid values are 'Skip', 'Selective', 'All'.
    ArtifactDownloadMode *string `json:"artifactDownloadMode,omitempty"`
    // Gets or sets the artifact items of the input.
    ArtifactItems *[]string `json:"artifactItems,omitempty"`
    // Gets or sets the type of artifact.
    ArtifactType *string `json:"artifactType,omitempty"`
}

type ArtifactFilter struct {
    // Gets or sets whether a release should be created on build tagging.
    CreateReleaseOnBuildTagging *bool `json:"createReleaseOnBuildTagging,omitempty"`
    // Gets or sets the branch for the filter.
    SourceBranch *string `json:"sourceBranch,omitempty"`
    // Gets or sets the list of tags for the filter.
    Tags *[]string `json:"tags,omitempty"`
    // Gets or sets whether filter should default to build definition branch.
    UseBuildDefinitionBranch *bool `json:"useBuildDefinitionBranch,omitempty"`
}

type ArtifactInstanceData struct {
    AccountName *string `json:"accountName,omitempty"`
    AuthenticationToken *string `json:"authenticationToken,omitempty"`
    TfsUrl *string `json:"tfsUrl,omitempty"`
    Version *string `json:"version,omitempty"`
}

type ArtifactMetadata struct {
    // Sets alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Sets instance reference of artifact. e.g. for build artifact it is build number.
    InstanceReference *BuildVersion `json:"instanceReference,omitempty"`
}

type ArtifactProvider struct {
    // Gets or sets the id of artifact provider.
    Id *int `json:"id,omitempty"`
    // Gets or sets the name of artifact provider.
    Name *string `json:"name,omitempty"`
    // Gets or sets the link of artifact provider.
    SourceUri *string `json:"sourceUri,omitempty"`
    // Gets or sets the version of artifact provider.
    Version *string `json:"version,omitempty"`
}

type ArtifactsDownloadInput struct {
    DownloadInputs *[]ArtifactDownloadInputBase `json:"downloadInputs,omitempty"`
}

type ArtifactSourceId struct {
    // Gets or sets the artifact type of artifact source.
    ArtifactTypeId *string `json:"artifactTypeId,omitempty"`
    // Gets or sets the list of sourceIdInput of artifact source.
    SourceIdInputs *[]SourceIdInput `json:"sourceIdInputs,omitempty"`
}

type ArtifactSourceIdsQueryResult struct {
    // Gets or sets the list of artifactsourceIds.
    ArtifactSourceIds *[]ArtifactSourceId `json:"artifactSourceIds,omitempty"`
}

type ArtifactSourceReference struct {
    // ID of the artifact source.
    Id *string `json:"id,omitempty"`
    // Name of the artifact source.
    Name *string `json:"name,omitempty"`
}

type ArtifactSourceTrigger struct {
    // Artifact source alias for Artifact Source trigger type
    ArtifactAlias *string `json:"artifactAlias,omitempty"`
    TriggerConditions *[]ArtifactFilter `json:"triggerConditions,omitempty"`
}

type ArtifactTriggerConfiguration struct {
    // Gets or sets the whether trigger is supported or not.
    IsTriggerSupported *bool `json:"isTriggerSupported,omitempty"`
    // Gets or sets the whether trigger is supported only on hosted environment.
    IsTriggerSupportedOnlyInHosted *bool `json:"isTriggerSupportedOnlyInHosted,omitempty"`
    // Gets or sets the whether webhook is supported at server level.
    IsWebhookSupportedAtServerLevel *bool `json:"isWebhookSupportedAtServerLevel,omitempty"`
    // Gets or sets the payload hash header name for the artifact trigger configuration.
    PayloadHashHeaderName *string `json:"payloadHashHeaderName,omitempty"`
    // Gets or sets the resources for artifact trigger configuration.
    Resources *map[string]string `json:"resources,omitempty"`
    // Gets or sets the webhook payload mapping for artifact trigger configuration.
    WebhookPayloadMapping *map[string]string `json:"webhookPayloadMapping,omitempty"`
}

type ArtifactTypeDefinition struct {
    // Gets or sets the artifact trigger configuration of artifact type definition.
    ArtifactTriggerConfiguration *ArtifactTriggerConfiguration `json:"artifactTriggerConfiguration,omitempty"`
    // Gets or sets the artifact type of artifact type definition. Valid values are 'Build', 'Package', 'Source' or 'ContainerImage'.
    ArtifactType *string `json:"artifactType,omitempty"`
    // Gets or sets the display name of artifact type definition.
    DisplayName *string `json:"displayName,omitempty"`
    // Gets or sets the endpoint type id of artifact type definition.
    EndpointTypeId *string `json:"endpointTypeId,omitempty"`
    // Gets or sets the input descriptors of artifact type definition.
    InputDescriptors *[]InputDescriptor `json:"inputDescriptors,omitempty"`
    // Gets or sets the name of artifact type definition.
    Name *string `json:"name,omitempty"`
    // Gets or sets the unique source identifier of artifact type definition.
    UniqueSourceIdentifier *string `json:"uniqueSourceIdentifier,omitempty"`
}

type ArtifactVersion struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the default version of artifact.
    DefaultVersion *BuildVersion `json:"defaultVersion,omitempty"`
    // Gets or sets the error message encountered during querying of versions for artifact.
    ErrorMessage *string `json:"errorMessage,omitempty"`
    // Deprecated: This property is deprecated use Alias instead and remove all its references
    SourceId *string `json:"sourceId,omitempty"`
    // Gets or sets the list of build versions of artifact.
    Versions *[]BuildVersion `json:"versions,omitempty"`
}

type ArtifactVersionQueryResult struct {
    // Gets or sets the list for artifact versions of artifact version query result.
    ArtifactVersions *[]ArtifactVersion `json:"artifactVersions,omitempty"`
}

type AuditAction string

type AuthorizationHeader struct {
    Name *string `json:"name,omitempty"`
    Value *string `json:"value,omitempty"`
}

type AuthorizationHeaderFor string

type AutoTriggerIssue struct {
    Issue *Issue `json:"issue,omitempty"`
    IssueSource *IssueSource `json:"issueSource,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
    ReleaseDefinitionReference *ReleaseDefinitionShallowReference `json:"releaseDefinitionReference,omitempty"`
    ReleaseTriggerType *ReleaseTriggerType `json:"releaseTriggerType,omitempty"`
}

type AzureKeyVaultVariableGroupProviderData struct {
    // Gets or sets last refreshed time.
    LastRefreshedOn *time.Time `json:"lastRefreshedOn,omitempty"`
    // Gets or sets the service endpoint ID.
    ServiceEndpointId *uuid.UUID `json:"serviceEndpointId,omitempty"`
    // Gets or sets the vault name.
    Vault *string `json:"vault,omitempty"`
}

type AzureKeyVaultVariableValue struct {
    // Gets or sets as the variable is secret or not.
    IsSecret *bool `json:"isSecret,omitempty"`
    // Gets or sets the value.
    Value *string `json:"value,omitempty"`
    // Gets or sets the content type of key vault variable value.
    ContentType *string `json:"contentType,omitempty"`
    // Indicates the vault variable value enabled or not.
    Enabled *bool `json:"enabled,omitempty"`
    // Gets or sets the expire time of key vault variable value.
    Expires *time.Time `json:"expires,omitempty"`
}

type BaseDeploymentInput struct {
    // Gets or sets the job condition.
    Condition *string `json:"condition,omitempty"`
    // Gets or sets the job cancel timeout in minutes for deployment which are cancelled by user for this release environment.
    JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
    // Gets or sets the override inputs.
    OverrideInputs *map[string]string `json:"overrideInputs,omitempty"`
    // Gets or sets the job execution timeout in minutes for deployment which are queued against this release environment.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
}

type BuildArtifactDownloadInput struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the name of artifact definition. Valid values are 'Skip', 'Selective', 'All'.
    ArtifactDownloadMode *string `json:"artifactDownloadMode,omitempty"`
    // Gets or sets the artifact items of the input.
    ArtifactItems *[]string `json:"artifactItems,omitempty"`
    // Gets or sets the type of artifact.
    ArtifactType *string `json:"artifactType,omitempty"`
}

type BuildVersion struct {
    // Gets or sets the commit message for the artifact.
    CommitMessage *string `json:"commitMessage,omitempty"`
    // Gets or sets the definition id.
    DefinitionId *string `json:"definitionId,omitempty"`
    // Gets or sets the definition name.
    DefinitionName *string `json:"definitionName,omitempty"`
    // Gets or sets the build id.
    Id *string `json:"id,omitempty"`
    // Gets or sets if the artifact supports multiple definitions.
    IsMultiDefinitionType *bool `json:"isMultiDefinitionType,omitempty"`
    // Gets or sets the build number.
    Name *string `json:"name,omitempty"`
    // Gets or sets the source branch for the artifact.
    SourceBranch *string `json:"sourceBranch,omitempty"`
    // Gets or sets the source pull request version for the artifact.
    SourcePullRequestVersion *SourcePullRequestVersion `json:"sourcePullRequestVersion,omitempty"`
    // Gets or sets the repository id for the artifact.
    SourceRepositoryId *string `json:"sourceRepositoryId,omitempty"`
    // Gets or sets the repository type for the artifact.
    SourceRepositoryType *string `json:"sourceRepositoryType,omitempty"`
    // Gets or sets the source version for the artifact.
    SourceVersion *string `json:"sourceVersion,omitempty"`
}

// Represents a change associated with a build.
type Change struct {
    // The author of the change.
    Author *IdentityRef `json:"author,omitempty"`
    // The type of source. "TfsVersionControl", "TfsGit", etc.
    ChangeType *string `json:"changeType,omitempty"`
    // The location of a user-friendly representation of the resource.
    DisplayUri *string `json:"displayUri,omitempty"`
    // Something that identifies the change. For a commit, this would be the SHA1. For a TFVC changeset, this would be the changeset id.
    Id *string `json:"id,omitempty"`
    // The location of the full representation of the resource.
    Location *string `json:"location,omitempty"`
    // A description of the change. This might be a commit message or changeset description.
    Message *string `json:"message,omitempty"`
    // The person or process that pushed the change.
    PushedBy *IdentityRef `json:"pushedBy,omitempty"`
    // Deprecated: Use PushedBy instead
    Pusher *string `json:"pusher,omitempty"`
    // A timestamp for the change.
    Timestamp *time.Time `json:"timestamp,omitempty"`
}

type CodeRepositoryReference struct {
    // Gets and sets the repository references.
    RepositoryReference *map[string]ReleaseManagementInputValue `json:"repositoryReference,omitempty"`
    // It can have value as ‘GitHub’, ‘Vsts’.
    SystemType *PullRequestSystemType `json:"systemType,omitempty"`
}

type ComplianceSettings struct {
    // Scan the release definition for secrets
    CheckForCredentialsAndOtherSecrets *bool `json:"checkForCredentialsAndOtherSecrets,omitempty"`
}

type Condition struct {
    // Gets or sets the condition type.
    ConditionType *ConditionType `json:"conditionType,omitempty"`
    // Gets or sets the name of the condition. e.g. 'ReleaseStarted'.
    Name *string `json:"name,omitempty"`
    // Gets or set value of the condition.
    Value *string `json:"value,omitempty"`
}

type ConditionType string

type ConfigurationVariableValue struct {
    // Gets and sets if a variable can be overridden at deployment time or not.
    AllowOverride *bool `json:"allowOverride,omitempty"`
    // Gets or sets as variable is secret or not.
    IsSecret *bool `json:"isSecret,omitempty"`
    // Gets and sets value of the configuration variable.
    Value *string `json:"value,omitempty"`
}

type Consumer struct {
    // ID of the consumer.
    ConsumerId *int `json:"consumerId,omitempty"`
    // Name of the consumer.
    ConsumerName *string `json:"consumerName,omitempty"`
}

type ContainerImageTrigger struct {
    // Alias of the trigger.
    Alias *string `json:"alias,omitempty"`
    // List tag filters applied while trigger.
    TagFilters *[]TagFilter `json:"tagFilters,omitempty"`
}

type ContinuousDeploymentTriggerIssue struct {
    Issue *Issue `json:"issue,omitempty"`
    IssueSource *IssueSource `json:"issueSource,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
    ReleaseDefinitionReference *ReleaseDefinitionShallowReference `json:"releaseDefinitionReference,omitempty"`
    ReleaseTriggerType *ReleaseTriggerType `json:"releaseTriggerType,omitempty"`
    // Artifact type.
    ArtifactType *string `json:"artifactType,omitempty"`
    // ArtifactVersion ID.
    ArtifactVersionId *string `json:"artifactVersionId,omitempty"`
    // Artifact source ID.
    SourceId *string `json:"sourceId,omitempty"`
}

type ControlOptions struct {
    // Always run the job.
    AlwaysRun *bool `json:"alwaysRun,omitempty"`
    // Indicates whether to continuer job on error or not.
    ContinueOnError *bool `json:"continueOnError,omitempty"`
    // Indicates the job enabled or not.
    Enabled *bool `json:"enabled,omitempty"`
}

type CustomArtifactDownloadInput struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the name of artifact definition. Valid values are 'Skip', 'Selective', 'All'.
    ArtifactDownloadMode *string `json:"artifactDownloadMode,omitempty"`
    // Gets or sets the artifact items of the input.
    ArtifactItems *[]string `json:"artifactItems,omitempty"`
    // Gets or sets the type of artifact.
    ArtifactType *string `json:"artifactType,omitempty"`
}

type DataSourceBinding struct {
    // Pagination format supported by this data source(ContinuationToken/SkipTop).
    CallbackContextTemplate *string `json:"callbackContextTemplate,omitempty"`
    // Subsequent calls needed?
    CallBackRequiredTemplate *string `json:"callBackRequiredTemplate,omitempty"`
    // Name of the datasource.
    DataSourceName *string `json:"dataSourceName,omitempty"`
    // Endpoint ID of the datasource.
    EndpointId *string `json:"endpointId,omitempty"`
    // Endpoint URL of the datasource.
    EndpointUrl *string `json:"endpointUrl,omitempty"`
    // Defines the initial value of the query params
    InitialContextTemplate *string `json:"initialContextTemplate,omitempty"`
    // Parameters of the datasource.
    Parameters *map[string]string `json:"parameters,omitempty"`
    // Gets or sets http request body
    RequestContent *string `json:"requestContent,omitempty"`
    // Gets or sets http request verb
    RequestVerb *string `json:"requestVerb,omitempty"`
    // Result selector applied on output of datasource result, for example jsonpath:$.value[?(@.properties.isEnabled == true)].
    ResultSelector *string `json:"resultSelector,omitempty"`
    // Format of the return results, for example. { "Value" : "{{{id}}}", "DisplayValue" : "{{{name}}}" }.
    ResultTemplate *string `json:"resultTemplate,omitempty"`
    // Target of the datasource.
    Target *string `json:"target,omitempty"`
}

// Represents binding of data source for the service endpoint request.
type DataSourceBindingBase struct {
    // Pagination format supported by this data source(ContinuationToken/SkipTop).
    CallbackContextTemplate *string `json:"callbackContextTemplate,omitempty"`
    // Subsequent calls needed?
    CallbackRequiredTemplate *string `json:"callbackRequiredTemplate,omitempty"`
    // Gets or sets the name of the data source.
    DataSourceName *string `json:"dataSourceName,omitempty"`
    // Gets or sets the endpoint Id.
    EndpointId *string `json:"endpointId,omitempty"`
    // Gets or sets the url of the service endpoint.
    EndpointUrl *string `json:"endpointUrl,omitempty"`
    // Gets or sets the authorization headers.
    Headers *[]AuthorizationHeader `json:"headers,omitempty"`
    // Defines the initial value of the query params
    InitialContextTemplate *string `json:"initialContextTemplate,omitempty"`
    // Gets or sets the parameters for the data source.
    Parameters *map[string]string `json:"parameters,omitempty"`
    // Gets or sets http request body
    RequestContent *string `json:"requestContent,omitempty"`
    // Gets or sets http request verb
    RequestVerb *string `json:"requestVerb,omitempty"`
    // Gets or sets the result selector.
    ResultSelector *string `json:"resultSelector,omitempty"`
    // Gets or sets the result template.
    ResultTemplate *string `json:"resultTemplate,omitempty"`
    // Gets or sets the target of the data source.
    Target *string `json:"target,omitempty"`
}

type DefinitionEnvironmentReference struct {
    // Definition environment ID.
    DefinitionEnvironmentId *int `json:"definitionEnvironmentId,omitempty"`
    // Definition environment name.
    DefinitionEnvironmentName *string `json:"definitionEnvironmentName,omitempty"`
    // ReleaseDefinition ID.
    ReleaseDefinitionId *int `json:"releaseDefinitionId,omitempty"`
    // ReleaseDefinition name.
    ReleaseDefinitionName *string `json:"releaseDefinitionName,omitempty"`
}

type Deployment struct {
    // Deprecated: Use ReleaseReference instead.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets attempt number.
    Attempt *int `json:"attempt,omitempty"`
    // Gets the date on which deployment is complete.
    CompletedOn *time.Time `json:"completedOn,omitempty"`
    // Gets the list of condition associated with deployment.
    Conditions *[]Condition `json:"conditions,omitempty"`
    // Gets release definition environment id.
    DefinitionEnvironmentId *int `json:"definitionEnvironmentId,omitempty"`
    // Gets status of the deployment.
    DeploymentStatus *DeploymentStatus `json:"deploymentStatus,omitempty"`
    // Gets the unique identifier for deployment.
    Id *int `json:"id,omitempty"`
    // Gets the identity who last modified the deployment.
    LastModifiedBy *IdentityRef `json:"lastModifiedBy,omitempty"`
    // Gets the date on which deployment is last modified.
    LastModifiedOn *time.Time `json:"lastModifiedOn,omitempty"`
    // Gets operation status of deployment.
    OperationStatus *DeploymentOperationStatus `json:"operationStatus,omitempty"`
    // Gets list of PostDeployApprovals.
    PostDeployApprovals *[]ReleaseApproval `json:"postDeployApprovals,omitempty"`
    // Gets list of PreDeployApprovals.
    PreDeployApprovals *[]ReleaseApproval `json:"preDeployApprovals,omitempty"`
    // Gets or sets project reference.
    ProjectReference *ProjectReference `json:"projectReference,omitempty"`
    // Gets the date on which deployment is queued.
    QueuedOn *time.Time `json:"queuedOn,omitempty"`
    // Gets reason of deployment.
    Reason *DeploymentReason `json:"reason,omitempty"`
    // Gets the reference of release.
    Release *ReleaseReference `json:"release,omitempty"`
    // Gets releaseDefinitionReference which specifies the reference of the release definition to which the deployment is associated.
    ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`
    // Gets releaseEnvironmentReference which specifies the reference of the release environment to which the deployment is associated.
    ReleaseEnvironment *ReleaseEnvironmentShallowReference `json:"releaseEnvironment,omitempty"`
    // Gets the identity who requested.
    RequestedBy *IdentityRef `json:"requestedBy,omitempty"`
    // Gets the identity for whom deployment is requested.
    RequestedFor *IdentityRef `json:"requestedFor,omitempty"`
    // Gets the date on which deployment is scheduled.
    ScheduledDeploymentTime *time.Time `json:"scheduledDeploymentTime,omitempty"`
    // Gets the date on which deployment is started.
    StartedOn *time.Time `json:"startedOn,omitempty"`
}

type DeploymentApprovalCompletedEvent struct {
    Approval *ReleaseApproval `json:"approval,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
    Release *Release `json:"release,omitempty"`
}

type DeploymentApprovalPendingEvent struct {
    Approval *ReleaseApproval `json:"approval,omitempty"`
    ApprovalOptions *ApprovalOptions `json:"approvalOptions,omitempty"`
    CompletedApprovals *[]ReleaseApproval `json:"completedApprovals,omitempty"`
    Data *map[string]interface{} `json:"data,omitempty"`
    Deployment *Deployment `json:"deployment,omitempty"`
    IsMultipleRankApproval *bool `json:"isMultipleRankApproval,omitempty"`
    PendingApprovals *[]ReleaseApproval `json:"pendingApprovals,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
    Release *Release `json:"release,omitempty"`
}

type DeploymentAttempt struct {
    // Deployment attempt.
    Attempt *int `json:"attempt,omitempty"`
    // ID of the deployment.
    DeploymentId *int `json:"deploymentId,omitempty"`
    // Deprecated: Instead use Issues which contains both errors and warnings related to deployment
    ErrorLog *string `json:"errorLog,omitempty"`
    // Specifies whether deployment has started or not.
    HasStarted *bool `json:"hasStarted,omitempty"`
    // ID of deployment.
    Id *int `json:"id,omitempty"`
    // All the issues related to the deployment.
    Issues *[]Issue `json:"issues,omitempty"`
    // Deprecated: Use ReleaseDeployPhase.DeploymentJobs.Job instead.
    Job *ReleaseTask `json:"job,omitempty"`
    // Identity who last modified this deployment.
    LastModifiedBy *IdentityRef `json:"lastModifiedBy,omitempty"`
    // Time when this deployment last modified.
    LastModifiedOn *time.Time `json:"lastModifiedOn,omitempty"`
    // Deployment operation status.
    OperationStatus *DeploymentOperationStatus `json:"operationStatus,omitempty"`
    // Post deployment gates that executed in this deployment.
    PostDeploymentGates *ReleaseGates `json:"postDeploymentGates,omitempty"`
    // Pre deployment gates that executed in this deployment.
    PreDeploymentGates *ReleaseGates `json:"preDeploymentGates,omitempty"`
    // When this deployment queued on.
    QueuedOn *time.Time `json:"queuedOn,omitempty"`
    // Reason for the deployment.
    Reason *DeploymentReason `json:"reason,omitempty"`
    // List of release deployphases executed in this deployment.
    ReleaseDeployPhases *[]ReleaseDeployPhase `json:"releaseDeployPhases,omitempty"`
    // Identity who requested this deployment.
    RequestedBy *IdentityRef `json:"requestedBy,omitempty"`
    // Identity for this deployment requested.
    RequestedFor *IdentityRef `json:"requestedFor,omitempty"`
    // Deprecated: Use ReleaseDeployPhase.RunPlanId instead.
    RunPlanId *uuid.UUID `json:"runPlanId,omitempty"`
    // status of the deployment.
    Status *DeploymentStatus `json:"status,omitempty"`
    // Deprecated: Use ReleaseDeployPhase.DeploymentJobs.Tasks instead.
    Tasks *[]ReleaseTask `json:"tasks,omitempty"`
}

type DeploymentAuthorizationInfo struct {
    // Authorization header type, typically either RevalidateApproverIdentity or OnBehalfOf.
    AuthorizationHeaderFor *AuthorizationHeaderFor `json:"authorizationHeaderFor,omitempty"`
    // List of resources.
    Resources *[]string `json:"resources,omitempty"`
    // ID of the tenant.
    TenantId *string `json:"tenantId,omitempty"`
    // Access token key.
    VstsAccessTokenKey *string `json:"vstsAccessTokenKey,omitempty"`
}

type DeploymentAuthorizationOwner string

type DeploymentCompletedEvent struct {
    Comment *string `json:"comment,omitempty"`
    Data *map[string]interface{} `json:"data,omitempty"`
    Deployment *Deployment `json:"deployment,omitempty"`
    Environment *ReleaseEnvironment `json:"environment,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
}

type DeploymentExpands string

type DeploymentInput struct {
    // Gets or sets the job condition.
    Condition *string `json:"condition,omitempty"`
    // Gets or sets the job cancel timeout in minutes for deployment which are cancelled by user for this release environment.
    JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
    // Gets or sets the override inputs.
    OverrideInputs *map[string]string `json:"overrideInputs,omitempty"`
    // Gets or sets the job execution timeout in minutes for deployment which are queued against this release environment.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
    // Artifacts that downloaded during job execution.
    ArtifactsDownloadInput *ArtifactsDownloadInput `json:"artifactsDownloadInput,omitempty"`
    // List demands that needs to meet to execute the job.
    Demands *[]interface{} `json:"demands,omitempty"`
    // Indicates whether to include access token in deployment job or not.
    EnableAccessToken *bool `json:"enableAccessToken,omitempty"`
    // Id of the pool on which job get executed.
    QueueId *int `json:"queueId,omitempty"`
    // Indicates whether artifacts downloaded while job execution or not.
    SkipArtifactsDownload *bool `json:"skipArtifactsDownload,omitempty"`
}

type DeploymentJob struct {
    // Parent task of all executed tasks.
    Job *ReleaseTask `json:"job,omitempty"`
    // List of  executed tasks with in job.
    Tasks *[]ReleaseTask `json:"tasks,omitempty"`
}

type DeploymentManualInterventionPendingEvent struct {
    Deployment *Deployment `json:"deployment,omitempty"`
    EmailRecipients *[]uuid.UUID `json:"emailRecipients,omitempty"`
    EnvironmentOwner *IdentityRef `json:"environmentOwner,omitempty"`
    ManualIntervention *ManualIntervention `json:"manualIntervention,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
    Release *Release `json:"release,omitempty"`
}

type DeploymentOperationStatus string

type DeploymentQueryParameters struct {
    // Query deployments based specified artifact source id.
    ArtifactSourceId *string `json:"artifactSourceId,omitempty"`
    // Query deployments based specified artifact type id.
    ArtifactTypeId *string `json:"artifactTypeId,omitempty"`
    // Query deployments based specified artifact versions.
    ArtifactVersions *[]string `json:"artifactVersions,omitempty"`
    // Query deployments number of deployments per environment.
    DeploymentsPerEnvironment *int `json:"deploymentsPerEnvironment,omitempty"`
    // Query deployment based on deployment status.
    DeploymentStatus *DeploymentStatus `json:"deploymentStatus,omitempty"`
    // Query deployments of specified environments.
    Environments *[]DefinitionEnvironmentReference `json:"environments,omitempty"`
    // Query deployments based specified expands.
    Expands *DeploymentExpands `json:"expands,omitempty"`
    // Specify deleted deployments should return or not.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Deprecated: 
    LatestDeploymentsOnly *bool `json:"latestDeploymentsOnly,omitempty"`
    // Deprecated: 
    MaxDeploymentsPerEnvironment *int `json:"maxDeploymentsPerEnvironment,omitempty"`
    // Deprecated: 
    MaxModifiedTime *time.Time `json:"maxModifiedTime,omitempty"`
    // Deprecated: 
    MinModifiedTime *time.Time `json:"minModifiedTime,omitempty"`
    // Query deployment based on deployment operation status.
    OperationStatus *DeploymentOperationStatus `json:"operationStatus,omitempty"`
    // Deprecated: 
    QueryOrder *ReleaseQueryOrder `json:"queryOrder,omitempty"`
    // Query deployments based query type.
    QueryType *DeploymentsQueryType `json:"queryType,omitempty"`
    // Query deployments based specified source branch.
    SourceBranch *string `json:"sourceBranch,omitempty"`
}

type DeploymentReason string

type DeploymentsQueryType string

type DeploymentStartedEvent struct {
    Environment *ReleaseEnvironment `json:"environment,omitempty"`
    Project *ProjectReference `json:"project,omitempty"`
    Release *Release `json:"release,omitempty"`
}

type DeploymentStatus string

type DeployPhaseStatus string

type DeployPhaseTypes string

type EmailRecipients struct {
    // List of email addresses.
    EmailAddresses *[]string `json:"emailAddresses,omitempty"`
    // List of TFS IDs guids.
    TfsIds *[]uuid.UUID `json:"tfsIds,omitempty"`
}

// Defines policy on environment queuing at Release Management side queue. We will send to Environment Runner [creating pre-deploy and other steps] only when the policies mentioned are satisfied.
type EnvironmentExecutionPolicy struct {
    // This policy decides, how many environments would be with Environment Runner.
    ConcurrencyCount *int `json:"concurrencyCount,omitempty"`
    // Queue depth in the EnvironmentQueue table, this table keeps the environment entries till Environment Runner is free [as per it's policy] to take another environment for running.
    QueueDepthCount *int `json:"queueDepthCount,omitempty"`
}

type EnvironmentOptions struct {
    // Gets and sets as the auto link workitems or not.
    AutoLinkWorkItems *bool `json:"autoLinkWorkItems,omitempty"`
    // Gets and sets as the badge enabled or not.
    BadgeEnabled *bool `json:"badgeEnabled,omitempty"`
    // Deprecated: Use Notifications instead.
    EmailNotificationType *string `json:"emailNotificationType,omitempty"`
    // Deprecated: Use Notifications instead.
    EmailRecipients *string `json:"emailRecipients,omitempty"`
    // Deprecated: Use DeploymentInput.EnableAccessToken instead.
    EnableAccessToken *bool `json:"enableAccessToken,omitempty"`
    // Gets and sets as the publish deployment status or not.
    PublishDeploymentStatus *bool `json:"publishDeploymentStatus,omitempty"`
    // Gets and sets as the.pull request deployment enabled or not.
    PullRequestDeploymentEnabled *bool `json:"pullRequestDeploymentEnabled,omitempty"`
    // Deprecated: Use DeploymentInput.SkipArtifactsDownload instead.
    SkipArtifactsDownload *bool `json:"skipArtifactsDownload,omitempty"`
    // Deprecated: Use DeploymentInput.TimeoutInMinutes instead.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
}

type EnvironmentRetentionPolicy struct {
    // Gets and sets the number of days to keep environment.
    DaysToKeep *int `json:"daysToKeep,omitempty"`
    // Gets and sets the number of releases to keep.
    ReleasesToKeep *int `json:"releasesToKeep,omitempty"`
    // Gets and sets as the build to be retained or not.
    RetainBuild *bool `json:"retainBuild,omitempty"`
}

type EnvironmentStatus string

type EnvironmentTrigger struct {
    // Definition environment ID on which this trigger applicable.
    DefinitionEnvironmentId *int `json:"definitionEnvironmentId,omitempty"`
    // ReleaseDefinition ID on which this trigger applicable.
    ReleaseDefinitionId *int `json:"releaseDefinitionId,omitempty"`
    // Gets or sets the trigger content.
    TriggerContent *string `json:"triggerContent,omitempty"`
    // Gets or sets the trigger type.
    TriggerType *EnvironmentTriggerType `json:"triggerType,omitempty"`
}

type EnvironmentTriggerContent struct {
    // Gets or sets action.
    Action *string `json:"action,omitempty"`
    // Gets or sets list of event types.
    EventTypes *[]string `json:"eventTypes,omitempty"`
}

type EnvironmentTriggerType string

type ExecutionInput struct {
    // Parallel execution type, for example MultiConfiguration or MultiMachine.
    ParallelExecutionType *ParallelExecutionTypes `json:"parallelExecutionType,omitempty"`
}

// Class to represent favorite entry.
type FavoriteItem struct {
    // Application specific data for the entry.
    Data *string `json:"data,omitempty"`
    // Unique Id of the the entry.
    Id *uuid.UUID `json:"id,omitempty"`
    // Display text for favorite entry.
    Name *string `json:"name,omitempty"`
    // Application specific favorite entry type. Empty or Null represents that Favorite item is a Folder.
    Type *string `json:"type,omitempty"`
}

type Folder struct {
    // Identity who created this folder.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // Time when this folder created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Description of the folder.
    Description *string `json:"description,omitempty"`
    // Identity who last changed this folder.
    LastChangedBy *IdentityRef `json:"lastChangedBy,omitempty"`
    // Time when this folder last changed.
    LastChangedDate *time.Time `json:"lastChangedDate,omitempty"`
    // path of the folder.
    Path *string `json:"path,omitempty"`
}

type FolderPathQueryOrder string

type GatesDeploymentInput struct {
    // Gets or sets the job condition.
    Condition *string `json:"condition,omitempty"`
    // Gets or sets the job cancel timeout in minutes for deployment which are cancelled by user for this release environment.
    JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
    // Gets or sets the override inputs.
    OverrideInputs *map[string]string `json:"overrideInputs,omitempty"`
    // Gets or sets the job execution timeout in minutes for deployment which are queued against this release environment.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
    // Gates minimum success duration.
    MinimumSuccessDuration *int `json:"minimumSuccessDuration,omitempty"`
    // Gates sampling interval.
    SamplingInterval *int `json:"samplingInterval,omitempty"`
    // Gates stabilization time.
    StabilizationTime *int `json:"stabilizationTime,omitempty"`
}

type GatesDeployPhase struct {
    // Gets and sets the gate job input.
    DeploymentInput *GatesDeploymentInput `json:"deploymentInput,omitempty"`
}

type GateStatus string

type GateUpdateMetadata struct {
    // Comment.
    Comment *string `json:"comment,omitempty"`
    // Name of gate to be ignored.
    GatesToIgnore *[]string `json:"gatesToIgnore,omitempty"`
}

type GitArtifactDownloadInput struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the name of artifact definition. Valid values are 'Skip', 'Selective', 'All'.
    ArtifactDownloadMode *string `json:"artifactDownloadMode,omitempty"`
    // Gets or sets the artifact items of the input.
    ArtifactItems *[]string `json:"artifactItems,omitempty"`
    // Gets or sets the type of artifact.
    ArtifactType *string `json:"artifactType,omitempty"`
}

type GitHubArtifactDownloadInput struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the name of artifact definition. Valid values are 'Skip', 'Selective', 'All'.
    ArtifactDownloadMode *string `json:"artifactDownloadMode,omitempty"`
    // Gets or sets the artifact items of the input.
    ArtifactItems *[]string `json:"artifactItems,omitempty"`
    // Gets or sets the type of artifact.
    ArtifactType *string `json:"artifactType,omitempty"`
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

type IgnoredGate struct {
    // Gets the date on which gate is last ignored.
    LastModifiedOn *time.Time `json:"lastModifiedOn,omitempty"`
    // Name of gate ignored.
    Name *string `json:"name,omitempty"`
}

// Enumerates data types that are supported as subscription input values.
type InputDataType string

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

// Mode in which a subscription input should be entered (in a UI)
type InputMode string

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

type Issue struct {
    // Issue data.
    Data *map[string]string `json:"data,omitempty"`
    // Issue type, for example error, warning or info.
    IssueType *string `json:"issueType,omitempty"`
    // Issue message.
    Message *string `json:"message,omitempty"`
}

type IssueSource string

type JenkinsArtifactDownloadInput struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the name of artifact definition. Valid values are 'Skip', 'Selective', 'All'.
    ArtifactDownloadMode *string `json:"artifactDownloadMode,omitempty"`
    // Gets or sets the artifact items of the input.
    ArtifactItems *[]string `json:"artifactItems,omitempty"`
    // Gets or sets the type of artifact.
    ArtifactType *string `json:"artifactType,omitempty"`
}

type MachineGroupBasedDeployPhase struct {
    // Gets and sets the deployment group job input
    DeploymentInput *MachineGroupDeploymentInput `json:"deploymentInput,omitempty"`
}

type MachineGroupDeploymentInput struct {
    // Gets or sets the job condition.
    Condition *string `json:"condition,omitempty"`
    // Gets or sets the job cancel timeout in minutes for deployment which are cancelled by user for this release environment.
    JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
    // Gets or sets the override inputs.
    OverrideInputs *map[string]string `json:"overrideInputs,omitempty"`
    // Gets or sets the job execution timeout in minutes for deployment which are queued against this release environment.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
    // Artifacts that downloaded during job execution.
    ArtifactsDownloadInput *ArtifactsDownloadInput `json:"artifactsDownloadInput,omitempty"`
    // List demands that needs to meet to execute the job.
    Demands *[]interface{} `json:"demands,omitempty"`
    // Indicates whether to include access token in deployment job or not.
    EnableAccessToken *bool `json:"enableAccessToken,omitempty"`
    // Id of the pool on which job get executed.
    QueueId *int `json:"queueId,omitempty"`
    // Indicates whether artifacts downloaded while job execution or not.
    SkipArtifactsDownload *bool `json:"skipArtifactsDownload,omitempty"`
    // Deployment group health option.
    DeploymentHealthOption *string `json:"deploymentHealthOption,omitempty"`
    // Minimum percentage of the targets guaranteed to be healthy.
    HealthPercent *int `json:"healthPercent,omitempty"`
    // Deployment target tag filter.
    Tags *[]string `json:"tags,omitempty"`
}

type MailMessage struct {
    // Body of mail.
    Body *string `json:"body,omitempty"`
    // Mail CC recipients.
    CC *EmailRecipients `json:"cc,omitempty"`
    // Reply to.
    InReplyTo *string `json:"inReplyTo,omitempty"`
    // Message ID of the mail.
    MessageId *string `json:"messageId,omitempty"`
    // Data when should be replied to mail.
    ReplyBy *time.Time `json:"replyBy,omitempty"`
    // Reply to Email recipients.
    ReplyTo *EmailRecipients `json:"replyTo,omitempty"`
    // List of mail section types.
    Sections *[]MailSectionType `json:"sections,omitempty"`
    // Mail sender type.
    SenderType *SenderType `json:"senderType,omitempty"`
    // Subject of the mail.
    Subject *string `json:"subject,omitempty"`
    // Mail To recipients.
    To *EmailRecipients `json:"to,omitempty"`
}

type MailSectionType string

type ManualIntervention struct {
    // Gets or sets the identity who should approve.
    Approver *IdentityRef `json:"approver,omitempty"`
    // Gets or sets comments for approval.
    Comments *string `json:"comments,omitempty"`
    // Gets date on which it got created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Gets the unique identifier for manual intervention.
    Id *int `json:"id,omitempty"`
    // Gets or sets instructions for approval.
    Instructions *string `json:"instructions,omitempty"`
    // Gets date on which it got modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Gets or sets the name.
    Name *string `json:"name,omitempty"`
    // Gets releaseReference for manual intervention.
    Release *ReleaseShallowReference `json:"release,omitempty"`
    // Gets releaseDefinitionReference for manual intervention.
    ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`
    // Gets releaseEnvironmentReference for manual intervention.
    ReleaseEnvironment *ReleaseEnvironmentShallowReference `json:"releaseEnvironment,omitempty"`
    // Gets or sets the status of the manual intervention.
    Status *ManualInterventionStatus `json:"status,omitempty"`
    // Get task instance identifier.
    TaskInstanceId *uuid.UUID `json:"taskInstanceId,omitempty"`
    // Gets url to access the manual intervention.
    Url *string `json:"url,omitempty"`
}

// Describes manual intervention status
type ManualInterventionStatus string

type ManualInterventionUpdateMetadata struct {
    // Sets the comment for manual intervention update.
    Comment *string `json:"comment,omitempty"`
    // Sets the status of the manual intervention.
    Status *ManualInterventionStatus `json:"status,omitempty"`
}

type MappingDetails struct {
    Mappings *map[string]InputValue `json:"mappings,omitempty"`
}

type Metric struct {
    // Name of the Metric.
    Name *string `json:"name,omitempty"`
    // Value of the Metric.
    Value *int `json:"value,omitempty"`
}

type MultiConfigInput struct {
    // Parallel execution type, for example MultiConfiguration or MultiMachine.
    ParallelExecutionType *ParallelExecutionTypes `json:"parallelExecutionType,omitempty"`
    // Indicate whether continue execution of deployment on error or not.
    ContinueOnError *bool `json:"continueOnError,omitempty"`
    // Maximum number of agents used while parallel execution.
    MaxNumberOfAgents *int `json:"maxNumberOfAgents,omitempty"`
    // Multipliers for parallel execution of deployment, for example x86,x64.
    Multipliers *string `json:"multipliers,omitempty"`
}

type MultiMachineInput struct {
    // Parallel execution type, for example MultiConfiguration or MultiMachine.
    ParallelExecutionType *ParallelExecutionTypes `json:"parallelExecutionType,omitempty"`
    // Indicate whether continue execution of deployment on error or not.
    ContinueOnError *bool `json:"continueOnError,omitempty"`
    // Maximum number of agents used while parallel execution.
    MaxNumberOfAgents *int `json:"maxNumberOfAgents,omitempty"`
}

type PackageTrigger struct {
    // Package trigger alias.
    Alias *string `json:"alias,omitempty"`
}

type ParallelExecutionInputBase struct {
    // Parallel execution type, for example MultiConfiguration or MultiMachine.
    ParallelExecutionType *ParallelExecutionTypes `json:"parallelExecutionType,omitempty"`
    // Indicate whether continue execution of deployment on error or not.
    ContinueOnError *bool `json:"continueOnError,omitempty"`
    // Maximum number of agents used while parallel execution.
    MaxNumberOfAgents *int `json:"maxNumberOfAgents,omitempty"`
}

type ParallelExecutionTypes string

type PipelineProcess struct {
    // Pipeline process type.
    Type *PipelineProcessTypes `json:"type,omitempty"`
}

type PipelineProcessTypes string

type ProcessParameters struct {
    DataSourceBindings *[]DataSourceBindingBase `json:"dataSourceBindings,omitempty"`
    Inputs *[]TaskInputDefinitionBase `json:"inputs,omitempty"`
    SourceDefinitions *[]TaskSourceDefinitionBase `json:"sourceDefinitions,omitempty"`
}

type ProjectReference struct {
    // Gets the unique identifier of this field.
    Id *uuid.UUID `json:"id,omitempty"`
    // Gets name of project.
    Name *string `json:"name,omitempty"`
}

type PropertySelector struct {
    // List of properties.
    Properties *[]string `json:"properties,omitempty"`
    // Property selector type.
    SelectorType *PropertySelectorType `json:"selectorType,omitempty"`
}

type PropertySelectorType string

type PullRequestConfiguration struct {
    // Code repository reference.
    CodeRepositoryReference *CodeRepositoryReference `json:"codeRepositoryReference,omitempty"`
    // In case of Source based artifacts, Code reference will be present in Artifact details.
    UseArtifactReference *bool `json:"useArtifactReference,omitempty"`
}

type PullRequestFilter struct {
    // List of tags.
    Tags *[]string `json:"tags,omitempty"`
    // Target branch of pull request.
    TargetBranch *string `json:"targetBranch,omitempty"`
}

type PullRequestSystemType string

type PullRequestTrigger struct {
    // Artifact alias trigger is linked to.
    ArtifactAlias *string `json:"artifactAlias,omitempty"`
    // Code reference details of pull request.
    PullRequestConfiguration *PullRequestConfiguration `json:"pullRequestConfiguration,omitempty"`
    // Policy name using which status will be published to pull request.
    StatusPolicyName *string `json:"statusPolicyName,omitempty"`
    // List of filters applied while trigger.
    TriggerConditions *[]PullRequestFilter `json:"triggerConditions,omitempty"`
}

type QueuedReleaseData struct {
    // Project ID of the release.
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    // Release queue position.
    QueuePosition *int `json:"queuePosition,omitempty"`
    // Queued release ID.
    ReleaseId *int `json:"releaseId,omitempty"`
}

type RealtimeReleaseDefinitionEvent struct {
    DefinitionId *int `json:"definitionId,omitempty"`
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
}

type RealtimeReleaseEvent struct {
    EnvironmentId *int `json:"environmentId,omitempty"`
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    ReleaseId *int `json:"releaseId,omitempty"`
}

// The class to represent a collection of REST reference links.
type ReferenceLinks struct {
    // The readonly view of the links.  Because Reference links are readonly, we only want to expose them as read only.
    Links *map[string]interface{} `json:"links,omitempty"`
}

type Release struct {
    // Gets links to access the release.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets or sets the list of artifacts.
    Artifacts *[]Artifact `json:"artifacts,omitempty"`
    // Gets or sets comment.
    Comment *string `json:"comment,omitempty"`
    // Gets or sets the identity who created.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // Gets date on which it got created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Gets revision number of definition snapshot.
    DefinitionSnapshotRevision *int `json:"definitionSnapshotRevision,omitempty"`
    // Gets or sets description of release.
    Description *string `json:"description,omitempty"`
    // Gets list of environments.
    Environments *[]ReleaseEnvironment `json:"environments,omitempty"`
    // Gets the unique identifier of this field.
    Id *int `json:"id,omitempty"`
    // Whether to exclude the release from retention policies.
    KeepForever *bool `json:"keepForever,omitempty"`
    // Gets logs container url.
    LogsContainerUrl *string `json:"logsContainerUrl,omitempty"`
    // Gets or sets the identity who modified.
    ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`
    // Gets date on which it got modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Gets name.
    Name *string `json:"name,omitempty"`
    // Gets pool name.
    PoolName *string `json:"poolName,omitempty"`
    // Gets or sets project reference.
    ProjectReference *ProjectReference `json:"projectReference,omitempty"`
    Properties *interface{} `json:"properties,omitempty"`
    // Gets reason of release.
    Reason *ReleaseReason `json:"reason,omitempty"`
    // Gets releaseDefinitionReference which specifies the reference of the release definition to which this release is associated.
    ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`
    // Gets or sets the release definition revision.
    ReleaseDefinitionRevision *int `json:"releaseDefinitionRevision,omitempty"`
    // Gets release name format.
    ReleaseNameFormat *string `json:"releaseNameFormat,omitempty"`
    // Gets status.
    Status *ReleaseStatus `json:"status,omitempty"`
    // Gets or sets list of tags.
    Tags *[]string `json:"tags,omitempty"`
    TriggeringArtifactAlias *string `json:"triggeringArtifactAlias,omitempty"`
    // Deprecated: Use Links instead.
    Url *string `json:"url,omitempty"`
    // Gets the list of variable groups.
    VariableGroups *[]VariableGroup `json:"variableGroups,omitempty"`
    // Gets or sets the dictionary of variables.
    Variables *map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

type ReleaseAbandonedEvent struct {
    Project *ProjectReference `json:"project,omitempty"`
    Release *Release `json:"release,omitempty"`
}

type ReleaseApproval struct {
    // Gets or sets the type of approval.
    ApprovalType *ApprovalType `json:"approvalType,omitempty"`
    // Gets the identity who approved.
    ApprovedBy *IdentityRef `json:"approvedBy,omitempty"`
    // Gets or sets the identity who should approve.
    Approver *IdentityRef `json:"approver,omitempty"`
    // Gets or sets attempt which specifies as which deployment attempt it belongs.
    Attempt *int `json:"attempt,omitempty"`
    // Gets or sets comments for approval.
    Comments *string `json:"comments,omitempty"`
    // Gets date on which it got created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Gets history which specifies all approvals associated with this approval.
    History *[]ReleaseApprovalHistory `json:"history,omitempty"`
    // Gets the unique identifier of this field.
    Id *int `json:"id,omitempty"`
    // Gets or sets as approval is automated or not.
    IsAutomated *bool `json:"isAutomated,omitempty"`
    // Deprecated: Use Notifications instead.
    IsNotificationOn *bool `json:"isNotificationOn,omitempty"`
    // Gets date on which it got modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Gets or sets rank which specifies the order of the approval. e.g. Same rank denotes parallel approval.
    Rank *int `json:"rank,omitempty"`
    // Gets releaseReference which specifies the reference of the release to which this approval is associated.
    Release *ReleaseShallowReference `json:"release,omitempty"`
    // Gets releaseDefinitionReference which specifies the reference of the release definition to which this approval is associated.
    ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`
    // Gets releaseEnvironmentReference which specifies the reference of the release environment to which this approval is associated.
    ReleaseEnvironment *ReleaseEnvironmentShallowReference `json:"releaseEnvironment,omitempty"`
    // Gets the revision number.
    Revision *int `json:"revision,omitempty"`
    // Gets or sets the status of the approval.
    Status *ApprovalStatus `json:"status,omitempty"`
    // Deprecated: Use Attempt instead.
    TrialNumber *int `json:"trialNumber,omitempty"`
    // Gets url to access the approval.
    Url *string `json:"url,omitempty"`
}

type ReleaseApprovalHistory struct {
    // Identity of the approver.
    Approver *IdentityRef `json:"approver,omitempty"`
    // Identity of the object who changed approval.
    ChangedBy *IdentityRef `json:"changedBy,omitempty"`
    // Approval history comments.
    Comments *string `json:"comments,omitempty"`
    // Time when this approval created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Time when this approval modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Approval history revision.
    Revision *int `json:"revision,omitempty"`
}

type ReleaseApprovalPendingEvent struct {
    Approval *ReleaseApproval `json:"approval,omitempty"`
    ApprovalOptions *ApprovalOptions `json:"approvalOptions,omitempty"`
    CompletedApprovals *[]ReleaseApproval `json:"completedApprovals,omitempty"`
    DefinitionName *string `json:"definitionName,omitempty"`
    Deployment *Deployment `json:"deployment,omitempty"`
    EnvironmentId *int `json:"environmentId,omitempty"`
    EnvironmentName *string `json:"environmentName,omitempty"`
    Environments *[]ReleaseEnvironment `json:"environments,omitempty"`
    IsMultipleRankApproval *bool `json:"isMultipleRankApproval,omitempty"`
    PendingApprovals *[]ReleaseApproval `json:"pendingApprovals,omitempty"`
    ReleaseCreator *string `json:"releaseCreator,omitempty"`
    ReleaseName *string `json:"releaseName,omitempty"`
    Title *string `json:"title,omitempty"`
    WebAccessUri *string `json:"webAccessUri,omitempty"`
}

type ReleaseArtifact struct {
    // Gets or sets the artifact provider of ReleaseArtifact.
    ArtifactProvider *ArtifactProvider `json:"artifactProvider,omitempty"`
    // Gets or sets the artifact type of ReleaseArtifact.
    ArtifactType *string `json:"artifactType,omitempty"`
    // Gets or sets the definition json of ReleaseArtifact.
    DefinitionData *string `json:"definitionData,omitempty"`
    // Gets or sets the definition id of ReleaseArtifact.
    DefinitionId *int `json:"definitionId,omitempty"`
    // Gets or sets the description of ReleaseArtifact.
    Description *string `json:"description,omitempty"`
    // Gets or sets the id of ReleaseArtifact.
    Id *int `json:"id,omitempty"`
    // Gets or sets the name of ReleaseArtifact.
    Name *string `json:"name,omitempty"`
    // Gets or sets the release id.
    ReleaseId *int `json:"releaseId,omitempty"`
}

type ReleaseCondition struct {
    // Gets or sets the condition type.
    ConditionType *ConditionType `json:"conditionType,omitempty"`
    // Gets or sets the name of the condition. e.g. 'ReleaseStarted'.
    Name *string `json:"name,omitempty"`
    // Gets or set value of the condition.
    Value *string `json:"value,omitempty"`
    // The release condition result.
    Result *bool `json:"result,omitempty"`
}

type ReleaseCreatedEvent struct {
    Project *ProjectReference `json:"project,omitempty"`
    Release *Release `json:"release,omitempty"`
}

type ReleaseDefinition struct {
    // Gets the links to related resources, APIs, and views for the release definition.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets the unique identifier of release definition.
    Id *int `json:"id,omitempty"`
    // Gets or sets the name of the release definition.
    Name *string `json:"name,omitempty"`
    // Gets or sets the path of the release definition.
    Path *string `json:"path,omitempty"`
    // Gets or sets project reference.
    ProjectReference *ProjectReference `json:"projectReference,omitempty"`
    // Gets the REST API url to access the release definition.
    Url *string `json:"url,omitempty"`
    // Gets or sets the list of artifacts.
    Artifacts *[]Artifact `json:"artifacts,omitempty"`
    // Gets or sets comment.
    Comment *string `json:"comment,omitempty"`
    // Gets or sets the identity who created.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // Gets date on which it got created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Gets or sets the description.
    Description *string `json:"description,omitempty"`
    // Gets or sets the list of environments.
    Environments *[]ReleaseDefinitionEnvironment `json:"environments,omitempty"`
    // Whether release definition is deleted.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Gets the reference of last release.
    LastRelease *ReleaseReference `json:"lastRelease,omitempty"`
    // Gets or sets the identity who modified.
    ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`
    // Gets date on which it got modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Gets or sets properties.
    Properties *interface{} `json:"properties,omitempty"`
    // Gets or sets the release name format.
    ReleaseNameFormat *string `json:"releaseNameFormat,omitempty"`
    // Deprecated: Retention policy at Release Definition level is deprecated. Use the Retention Policy at environment and API version 3.0-preview.2 or later
    RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
    // Gets the revision number.
    Revision *int `json:"revision,omitempty"`
    // Gets or sets source of release definition.
    Source *ReleaseDefinitionSource `json:"source,omitempty"`
    // Gets or sets list of tags.
    Tags *[]string `json:"tags,omitempty"`
    // Gets or sets the list of triggers.
    Triggers *[]interface{} `json:"triggers,omitempty"`
    // Gets or sets the list of variable groups.
    VariableGroups *[]int `json:"variableGroups,omitempty"`
    // Gets or sets the dictionary of variables.
    Variables *map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

type ReleaseDefinitionApprovals struct {
    // Gets or sets the approval options.
    ApprovalOptions *ApprovalOptions `json:"approvalOptions,omitempty"`
    // Gets or sets the approvals.
    Approvals *[]ReleaseDefinitionApprovalStep `json:"approvals,omitempty"`
}

type ReleaseDefinitionApprovalStep struct {
    // ID of the approval or deploy step.
    Id *int `json:"id,omitempty"`
    // Gets and sets the approver.
    Approver *IdentityRef `json:"approver,omitempty"`
    // Indicates whether the approval automated.
    IsAutomated *bool `json:"isAutomated,omitempty"`
    // Indicates whether the approval notification set.
    IsNotificationOn *bool `json:"isNotificationOn,omitempty"`
    // Gets or sets the rank of approval step.
    Rank *int `json:"rank,omitempty"`
}

type ReleaseDefinitionDeployStep struct {
    // ID of the approval or deploy step.
    Id *int `json:"id,omitempty"`
    // The list of steps for this definition.
    Tasks *[]WorkflowTask `json:"tasks,omitempty"`
}

type ReleaseDefinitionEnvironment struct {
    // Gets or sets the BadgeUrl. BadgeUrl will be used when Badge will be enabled in Release Definition Environment.
    BadgeUrl *string `json:"badgeUrl,omitempty"`
    // Gets or sets the environment conditions.
    Conditions *[]Condition `json:"conditions,omitempty"`
    // Gets or sets the current release reference.
    CurrentRelease *ReleaseShallowReference `json:"currentRelease,omitempty"`
    // Gets or sets the demands.
    Demands *[]interface{} `json:"demands,omitempty"`
    // Gets or sets the deploy phases of environment.
    DeployPhases *[]interface{} `json:"deployPhases,omitempty"`
    // Gets or sets the deploystep.
    DeployStep *ReleaseDefinitionDeployStep `json:"deployStep,omitempty"`
    // Gets or sets the environment options.
    EnvironmentOptions *EnvironmentOptions `json:"environmentOptions,omitempty"`
    // Gets or sets the triggers on environment.
    EnvironmentTriggers *[]EnvironmentTrigger `json:"environmentTriggers,omitempty"`
    // Gets or sets the environment execution policy.
    ExecutionPolicy *EnvironmentExecutionPolicy `json:"executionPolicy,omitempty"`
    // Gets and sets the ID of the ReleaseDefinitionEnvironment.
    Id *int `json:"id,omitempty"`
    // Gets and sets the name of the ReleaseDefinitionEnvironment.
    Name *string `json:"name,omitempty"`
    // Gets and sets the Owner of the ReleaseDefinitionEnvironment.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Gets or sets the post deployment approvals.
    PostDeployApprovals *ReleaseDefinitionApprovals `json:"postDeployApprovals,omitempty"`
    // Gets or sets the post deployment gates.
    PostDeploymentGates *ReleaseDefinitionGatesStep `json:"postDeploymentGates,omitempty"`
    // Gets or sets the pre deployment approvals.
    PreDeployApprovals *ReleaseDefinitionApprovals `json:"preDeployApprovals,omitempty"`
    // Gets or sets the pre deployment gates.
    PreDeploymentGates *ReleaseDefinitionGatesStep `json:"preDeploymentGates,omitempty"`
    // Gets or sets the environment process parameters.
    ProcessParameters *ProcessParameters `json:"processParameters,omitempty"`
    // Gets or sets the properties on environment.
    Properties *interface{} `json:"properties,omitempty"`
    // Gets or sets the queue ID.
    QueueId *int `json:"queueId,omitempty"`
    // Gets and sets the rank of the ReleaseDefinitionEnvironment.
    Rank *int `json:"rank,omitempty"`
    // Gets or sets the environment retention policy.
    RetentionPolicy *EnvironmentRetentionPolicy `json:"retentionPolicy,omitempty"`
    // Deprecated: This property is deprecated, use EnvironmentOptions instead.
    RunOptions *map[string]string `json:"runOptions,omitempty"`
    // Gets or sets the schedules
    Schedules *[]ReleaseSchedule `json:"schedules,omitempty"`
    // Gets or sets the variable groups.
    VariableGroups *[]int `json:"variableGroups,omitempty"`
    // Gets and sets the variables.
    Variables *map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

type ReleaseDefinitionEnvironmentStep struct {
    // ID of the approval or deploy step.
    Id *int `json:"id,omitempty"`
}

type ReleaseDefinitionEnvironmentSummary struct {
    // ID of ReleaseDefinition environment summary.
    Id *int `json:"id,omitempty"`
    // List of release shallow reference deployed using this ReleaseDefinition.
    LastReleases *[]ReleaseShallowReference `json:"lastReleases,omitempty"`
    // Name of ReleaseDefinition environment summary.
    Name *string `json:"name,omitempty"`
}

type ReleaseDefinitionEnvironmentTemplate struct {
    // Indicates whether template can be deleted or not.
    CanDelete *bool `json:"canDelete,omitempty"`
    // Category of the ReleaseDefinition environment template.
    Category *string `json:"category,omitempty"`
    // Description of the ReleaseDefinition environment template.
    Description *string `json:"description,omitempty"`
    // ReleaseDefinition environment data which used to create this template.
    Environment *ReleaseDefinitionEnvironment `json:"environment,omitempty"`
    // ID of the task which used to display icon used for this template.
    IconTaskId *uuid.UUID `json:"iconTaskId,omitempty"`
    // Icon uri of the template.
    IconUri *string `json:"iconUri,omitempty"`
    // ID of the ReleaseDefinition environment template.
    Id *uuid.UUID `json:"id,omitempty"`
    // Indicates whether template deleted or not.
    IsDeleted *bool `json:"isDeleted,omitempty"`
    // Name of the ReleaseDefinition environment template.
    Name *string `json:"name,omitempty"`
}

type ReleaseDefinitionExpands string

type ReleaseDefinitionGate struct {
    // Gets or sets the gates workflow.
    Tasks *[]WorkflowTask `json:"tasks,omitempty"`
}

type ReleaseDefinitionGatesOptions struct {
    // Gets or sets as the gates enabled or not.
    IsEnabled *bool `json:"isEnabled,omitempty"`
    // Gets or sets the minimum duration for steady results after a successful gates evaluation.
    MinimumSuccessDuration *int `json:"minimumSuccessDuration,omitempty"`
    // Gets or sets the time between re-evaluation of gates.
    SamplingInterval *int `json:"samplingInterval,omitempty"`
    // Gets or sets the delay before evaluation.
    StabilizationTime *int `json:"stabilizationTime,omitempty"`
    // Gets or sets the timeout after which gates fail.
    Timeout *int `json:"timeout,omitempty"`
}

type ReleaseDefinitionGatesStep struct {
    // Gets or sets the gates.
    Gates *[]ReleaseDefinitionGate `json:"gates,omitempty"`
    // Gets or sets the gate options.
    GatesOptions *ReleaseDefinitionGatesOptions `json:"gatesOptions,omitempty"`
    // ID of the ReleaseDefinitionGateStep.
    Id *int `json:"id,omitempty"`
}

type ReleaseDefinitionQueryOrder string

type ReleaseDefinitionRevision struct {
    // Gets api-version for revision object.
    ApiVersion *string `json:"apiVersion,omitempty"`
    // Gets the identity who did change.
    ChangedBy *IdentityRef `json:"changedBy,omitempty"`
    // Gets date on which ReleaseDefinition changed.
    ChangedDate *time.Time `json:"changedDate,omitempty"`
    // Gets type of change.
    ChangeType *AuditAction `json:"changeType,omitempty"`
    // Gets comments for revision.
    Comment *string `json:"comment,omitempty"`
    // Get id of the definition.
    DefinitionId *int `json:"definitionId,omitempty"`
    // Gets definition URL.
    DefinitionUrl *string `json:"definitionUrl,omitempty"`
    // Get revision number of the definition.
    Revision *int `json:"revision,omitempty"`
}

type ReleaseDefinitionShallowReference struct {
    // Gets the links to related resources, APIs, and views for the release definition.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets the unique identifier of release definition.
    Id *int `json:"id,omitempty"`
    // Gets or sets the name of the release definition.
    Name *string `json:"name,omitempty"`
    // Gets or sets the path of the release definition.
    Path *string `json:"path,omitempty"`
    // Gets or sets project reference.
    ProjectReference *ProjectReference `json:"projectReference,omitempty"`
    // Gets the REST API url to access the release definition.
    Url *string `json:"url,omitempty"`
}

type ReleaseDefinitionSource string

type ReleaseDefinitionSummary struct {
    // List of Release Definition environment summary.
    Environments *[]ReleaseDefinitionEnvironmentSummary `json:"environments,omitempty"`
    // Release Definition reference.
    ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`
    // List of releases deployed using this Release Definition.
    Releases *[]Release `json:"releases,omitempty"`
}

type ReleaseDefinitionUndeleteParameter struct {
    // Gets or sets comment.
    Comment *string `json:"comment,omitempty"`
}

type ReleaseDeployPhase struct {
    // Deployment jobs of the phase.
    DeploymentJobs *[]DeploymentJob `json:"deploymentJobs,omitempty"`
    // Phase execution error logs.
    ErrorLog *string `json:"errorLog,omitempty"`
    // Deprecated: 
    Id *int `json:"id,omitempty"`
    // List of manual intervention tasks execution information in phase.
    ManualInterventions *[]ManualIntervention `json:"manualInterventions,omitempty"`
    // Name of the phase.
    Name *string `json:"name,omitempty"`
    // ID of the phase.
    PhaseId *string `json:"phaseId,omitempty"`
    // Type of the phase.
    PhaseType *DeployPhaseTypes `json:"phaseType,omitempty"`
    // Rank of the phase.
    Rank *int `json:"rank,omitempty"`
    // Run Plan ID of the phase.
    RunPlanId *uuid.UUID `json:"runPlanId,omitempty"`
    // Phase start time.
    StartedOn *time.Time `json:"startedOn,omitempty"`
    // Status of the phase.
    Status *DeployPhaseStatus `json:"status,omitempty"`
}

type ReleaseEnvironment struct {
    // Gets list of conditions.
    Conditions *[]ReleaseCondition `json:"conditions,omitempty"`
    // Gets date on which it got created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Gets definition environment id.
    DefinitionEnvironmentId *int `json:"definitionEnvironmentId,omitempty"`
    // Deprecated: Use DeploymentInput.Demands instead.
    Demands *[]interface{} `json:"demands,omitempty"`
    // Gets list of deploy phases snapshot.
    DeployPhasesSnapshot *[]interface{} `json:"deployPhasesSnapshot,omitempty"`
    // Gets deploy steps.
    DeploySteps *[]DeploymentAttempt `json:"deploySteps,omitempty"`
    // Gets environment options.
    EnvironmentOptions *EnvironmentOptions `json:"environmentOptions,omitempty"`
    // Gets the unique identifier of this field.
    Id *int `json:"id,omitempty"`
    // Gets date on which it got modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Gets name.
    Name *string `json:"name,omitempty"`
    // Gets next scheduled UTC time.
    NextScheduledUtcTime *time.Time `json:"nextScheduledUtcTime,omitempty"`
    // Gets the identity who is owner for release environment.
    Owner *IdentityRef `json:"owner,omitempty"`
    // Gets list of post deploy approvals snapshot.
    PostApprovalsSnapshot *ReleaseDefinitionApprovals `json:"postApprovalsSnapshot,omitempty"`
    // Gets list of post deploy approvals.
    PostDeployApprovals *[]ReleaseApproval `json:"postDeployApprovals,omitempty"`
    // Post deployment gates snapshot data.
    PostDeploymentGatesSnapshot *ReleaseDefinitionGatesStep `json:"postDeploymentGatesSnapshot,omitempty"`
    // Gets list of pre deploy approvals snapshot.
    PreApprovalsSnapshot *ReleaseDefinitionApprovals `json:"preApprovalsSnapshot,omitempty"`
    // Gets list of pre deploy approvals.
    PreDeployApprovals *[]ReleaseApproval `json:"preDeployApprovals,omitempty"`
    // Pre deployment gates snapshot data.
    PreDeploymentGatesSnapshot *ReleaseDefinitionGatesStep `json:"preDeploymentGatesSnapshot,omitempty"`
    // Gets process parameters.
    ProcessParameters *ProcessParameters `json:"processParameters,omitempty"`
    // Deprecated: Use DeploymentInput.QueueId instead.
    QueueId *int `json:"queueId,omitempty"`
    // Gets rank.
    Rank *int `json:"rank,omitempty"`
    // Gets release reference which specifies the reference of the release to which this release environment is associated.
    Release *ReleaseShallowReference `json:"release,omitempty"`
    // Gets the identity who created release.
    ReleaseCreatedBy *IdentityRef `json:"releaseCreatedBy,omitempty"`
    // Gets releaseDefinitionReference which specifies the reference of the release definition to which this release environment is associated.
    ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`
    // Deprecated: Use Release object Description instead.
    ReleaseDescription *string `json:"releaseDescription,omitempty"`
    // Gets release id.
    ReleaseId *int `json:"releaseId,omitempty"`
    // Gets schedule deployment time of release environment.
    ScheduledDeploymentTime *time.Time `json:"scheduledDeploymentTime,omitempty"`
    // Gets list of schedules.
    Schedules *[]ReleaseSchedule `json:"schedules,omitempty"`
    // Gets environment status.
    Status *EnvironmentStatus `json:"status,omitempty"`
    // Gets time to deploy.
    TimeToDeploy *float64 `json:"timeToDeploy,omitempty"`
    // Gets trigger reason.
    TriggerReason *string `json:"triggerReason,omitempty"`
    // Gets the list of variable groups.
    VariableGroups *[]VariableGroup `json:"variableGroups,omitempty"`
    // Gets the dictionary of variables.
    Variables *map[string]ConfigurationVariableValue `json:"variables,omitempty"`
    // Deprecated: Use DeployPhase.WorkflowTasks instead.
    WorkflowTasks *[]WorkflowTask `json:"workflowTasks,omitempty"`
}

type ReleaseEnvironmentCompletedEvent struct {
    CreatedByName *string `json:"createdByName,omitempty"`
    DefinitionId *int `json:"definitionId,omitempty"`
    DefinitionName *string `json:"definitionName,omitempty"`
    Environment *ReleaseEnvironment `json:"environment,omitempty"`
    EnvironmentId *int `json:"environmentId,omitempty"`
    ProjectName *string `json:"projectName,omitempty"`
    Reason *DeploymentReason `json:"reason,omitempty"`
    ReleaseCreatedBy *IdentityRef `json:"releaseCreatedBy,omitempty"`
    ReleaseLogsUri *string `json:"releaseLogsUri,omitempty"`
    ReleaseName *string `json:"releaseName,omitempty"`
    Status *string `json:"status,omitempty"`
    Title *string `json:"title,omitempty"`
    WebAccessUri *string `json:"webAccessUri,omitempty"`
}

type ReleaseEnvironmentShallowReference struct {
    // Gets the links to related resources, APIs, and views for the release environment.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets the unique identifier of release environment.
    Id *int `json:"id,omitempty"`
    // Gets or sets the name of the release environment.
    Name *string `json:"name,omitempty"`
    // Gets the REST API url to access the release environment.
    Url *string `json:"url,omitempty"`
}

type ReleaseEnvironmentStatusUpdatedEvent struct {
    DefinitionId *int `json:"definitionId,omitempty"`
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    EnvironmentId *int `json:"environmentId,omitempty"`
    EnvironmentStatus *EnvironmentStatus `json:"environmentStatus,omitempty"`
    LatestDeploymentOperationStatus *DeploymentOperationStatus `json:"latestDeploymentOperationStatus,omitempty"`
    LatestDeploymentStatus *DeploymentStatus `json:"latestDeploymentStatus,omitempty"`
    ReleaseId *int `json:"releaseId,omitempty"`
}

type ReleaseEnvironmentUpdateMetadata struct {
    // Gets or sets comment.
    Comment *string `json:"comment,omitempty"`
    // Gets or sets scheduled deployment time.
    ScheduledDeploymentTime *time.Time `json:"scheduledDeploymentTime,omitempty"`
    // Gets or sets status of environment.
    Status *EnvironmentStatus `json:"status,omitempty"`
    // Sets list of environment variables to be overridden at deployment time.
    Variables *map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

type ReleaseExpands string

type ReleaseGates struct {
    // Contains the gates job details of each evaluation.
    DeploymentJobs *[]DeploymentJob `json:"deploymentJobs,omitempty"`
    // ID of release gates.
    Id *int `json:"id,omitempty"`
    // List of ignored gates.
    IgnoredGates *[]IgnoredGate `json:"ignoredGates,omitempty"`
    // Gates last modified time.
    LastModifiedOn *time.Time `json:"lastModifiedOn,omitempty"`
    // Run plan ID of the gates.
    RunPlanId *uuid.UUID `json:"runPlanId,omitempty"`
    // Gates stabilization completed date and time.
    StabilizationCompletedOn *time.Time `json:"stabilizationCompletedOn,omitempty"`
    // Gates evaluation started time.
    StartedOn *time.Time `json:"startedOn,omitempty"`
    // Status of release gates.
    Status *GateStatus `json:"status,omitempty"`
    // Date and time at which all gates executed successfully.
    SucceedingSince *time.Time `json:"succeedingSince,omitempty"`
}

type ReleaseGatesPhase struct {
    // Deployment jobs of the phase.
    DeploymentJobs *[]DeploymentJob `json:"deploymentJobs,omitempty"`
    // Phase execution error logs.
    ErrorLog *string `json:"errorLog,omitempty"`
    // ID of the phase.
    Id *int `json:"id,omitempty"`
    // List of manual intervention tasks execution information in phase.
    ManualInterventions *[]ManualIntervention `json:"manualInterventions,omitempty"`
    // Name of the phase.
    Name *string `json:"name,omitempty"`
    // ID of the phase.
    PhaseId *string `json:"phaseId,omitempty"`
    // Type of the phase.
    PhaseType *DeployPhaseTypes `json:"phaseType,omitempty"`
    // Rank of the phase.
    Rank *int `json:"rank,omitempty"`
    // Run Plan ID of the phase.
    RunPlanId *uuid.UUID `json:"runPlanId,omitempty"`
    // Phase start time.
    StartedOn *time.Time `json:"startedOn,omitempty"`
    // Status of the phase.
    Status *DeployPhaseStatus `json:"status,omitempty"`
    // List of ignored gates.
    IgnoredGates *[]IgnoredGate `json:"ignoredGates,omitempty"`
    // Date and time at which stabilization of gates completed.
    StabilizationCompletedOn *time.Time `json:"stabilizationCompletedOn,omitempty"`
    // Date and time at which all gates executed successfully.
    SucceedingSince *time.Time `json:"succeedingSince,omitempty"`
}

type ReleaseManagementInputValue struct {
    // The text to show for the display of this value.
    DisplayValue *string `json:"displayValue,omitempty"`
    // The value to store for this input.
    Value *string `json:"value,omitempty"`
}

type ReleaseNotCreatedEvent struct {
    DefinitionReference *ReleaseDefinitionShallowReference `json:"definitionReference,omitempty"`
    Message *string `json:"message,omitempty"`
    ReleaseReason *ReleaseReason `json:"releaseReason,omitempty"`
    RequestedBy *IdentityRef `json:"requestedBy,omitempty"`
}

type ReleaseQueryOrder string

type ReleaseReason string

type ReleaseReference struct {
    // Gets links to access the release.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets list of artifacts.
    Artifacts *[]Artifact `json:"artifacts,omitempty"`
    // Gets the identity who created release.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // Gets date on when this release created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Gets description.
    Description *string `json:"description,omitempty"`
    // ID of the Release.
    Id *int `json:"id,omitempty"`
    // Gets the identity who modified release.
    ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`
    // Gets name of release.
    Name *string `json:"name,omitempty"`
    // Gets reason for release.
    Reason *ReleaseReason `json:"reason,omitempty"`
    // Gets release definition shallow reference.
    ReleaseDefinition *ReleaseDefinitionShallowReference `json:"releaseDefinition,omitempty"`
    // Deprecated: Use Links instead
    Url *string `json:"url,omitempty"`
    // Deprecated: Use Links instead
    WebAccessUri *string `json:"webAccessUri,omitempty"`
}

type ReleaseRevision struct {
    // Gets or sets the identity who changed.
    ChangedBy *IdentityRef `json:"changedBy,omitempty"`
    // Change date of the revision.
    ChangedDate *time.Time `json:"changedDate,omitempty"`
    // Change details of the revision.
    ChangeDetails *string `json:"changeDetails,omitempty"`
    // Change details of the revision. Typically ChangeDetails values are Add and Update.
    ChangeType *string `json:"changeType,omitempty"`
    // Comment of the revision.
    Comment *string `json:"comment,omitempty"`
    // Release ID of which this revision belongs.
    DefinitionSnapshotRevision *int `json:"definitionSnapshotRevision,omitempty"`
    // Gets or sets the release ID of which this revision belongs.
    ReleaseId *int `json:"releaseId,omitempty"`
}

type ReleaseSchedule struct {
    // Days of the week to release.
    DaysToRelease *ScheduleDays `json:"daysToRelease,omitempty"`
    // Team Foundation Job Definition Job Id.
    JobId *uuid.UUID `json:"jobId,omitempty"`
    // Flag to determine if this schedule should only release if the associated artifact has been changed or release definition changed.
    ScheduleOnlyWithChanges *bool `json:"scheduleOnlyWithChanges,omitempty"`
    // Local time zone hour to start.
    StartHours *int `json:"startHours,omitempty"`
    // Local time zone minute to start.
    StartMinutes *int `json:"startMinutes,omitempty"`
    // Time zone Id of release schedule, such as 'UTC'.
    TimeZoneId *string `json:"timeZoneId,omitempty"`
}

type ReleaseSettings struct {
    // Release Compliance settings.
    ComplianceSettings *ComplianceSettings `json:"complianceSettings,omitempty"`
    // Release retention settings.
    RetentionSettings *RetentionSettings `json:"retentionSettings,omitempty"`
}

type ReleaseShallowReference struct {
    // Gets the links to related resources, APIs, and views for the release.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Gets the unique identifier of release.
    Id *int `json:"id,omitempty"`
    // Gets or sets the name of the release.
    Name *string `json:"name,omitempty"`
    // Gets the REST API url to access the release.
    Url *string `json:"url,omitempty"`
}

type ReleaseStartEnvironmentMetadata struct {
    // Sets release definition environment id.
    DefinitionEnvironmentId *int `json:"definitionEnvironmentId,omitempty"`
    // Sets list of environments variables to be overridden at deployment time.
    Variables *map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

type ReleaseStartMetadata struct {
    // Sets list of artifact to create a release.
    Artifacts *[]ArtifactMetadata `json:"artifacts,omitempty"`
    // Sets definition Id to create a release.
    DefinitionId *int `json:"definitionId,omitempty"`
    // Sets description to create a release.
    Description *string `json:"description,omitempty"`
    // Sets list of environments meta data.
    EnvironmentsMetadata *[]ReleaseStartEnvironmentMetadata `json:"environmentsMetadata,omitempty"`
    // Sets 'true' to create release in draft mode, 'false' otherwise.
    IsDraft *bool `json:"isDraft,omitempty"`
    // Sets list of environments to manual as condition.
    ManualEnvironments *[]string `json:"manualEnvironments,omitempty"`
    Properties *interface{} `json:"properties,omitempty"`
    // Sets reason to create a release.
    Reason *ReleaseReason `json:"reason,omitempty"`
    // Sets list of release variables to be overridden at deployment time.
    Variables *map[string]ConfigurationVariableValue `json:"variables,omitempty"`
}

type ReleaseStatus string

type ReleaseTask struct {
    // Agent name on which task executed.
    AgentName *string `json:"agentName,omitempty"`
    // Deprecated: Use FinishTime instead
    DateEnded *time.Time `json:"dateEnded,omitempty"`
    // Deprecated: Use StartTime instead.
    DateStarted *time.Time `json:"dateStarted,omitempty"`
    // Finish time of the release task.
    FinishTime *time.Time `json:"finishTime,omitempty"`
    // ID of the release task.
    Id *int `json:"id,omitempty"`
    // List of issues occurred while execution of task.
    Issues *[]Issue `json:"issues,omitempty"`
    // Number of lines log release task has.
    LineCount *uint64 `json:"lineCount,omitempty"`
    // Log URL of the task.
    LogUrl *string `json:"logUrl,omitempty"`
    // Name of the task.
    Name *string `json:"name,omitempty"`
    // Task execution complete precent.
    PercentComplete *int `json:"percentComplete,omitempty"`
    // Rank of the release task.
    Rank *int `json:"rank,omitempty"`
    // Result code of the task.
    ResultCode *string `json:"resultCode,omitempty"`
    // ID of the release task.
    StartTime *time.Time `json:"startTime,omitempty"`
    // Status of release task.
    Status *TaskStatus `json:"status,omitempty"`
    // Workflow task reference.
    Task *WorkflowTaskReference `json:"task,omitempty"`
    // Timeline record ID of the release task.
    TimelineRecordId *uuid.UUID `json:"timelineRecordId,omitempty"`
}

type ReleaseTaskAttachment struct {
    // Reference links of task.
    Links *ReferenceLinks `json:"_links,omitempty"`
    // Data and time when it created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Identity who modified.
    ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`
    // Data and time when modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Name of the task attachment.
    Name *string `json:"name,omitempty"`
    // Record ID of the task.
    RecordId *uuid.UUID `json:"recordId,omitempty"`
    // Timeline ID of the task.
    TimelineId *uuid.UUID `json:"timelineId,omitempty"`
    // Type of task attachment.
    Type *string `json:"type,omitempty"`
}

type ReleaseTaskLogUpdatedEvent struct {
    EnvironmentId *int `json:"environmentId,omitempty"`
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    ReleaseId *int `json:"releaseId,omitempty"`
    Lines *[]string `json:"lines,omitempty"`
    StepRecordId *uuid.UUID `json:"stepRecordId,omitempty"`
    TimelineRecordId *uuid.UUID `json:"timelineRecordId,omitempty"`
}

type ReleaseTasksUpdatedEvent struct {
    EnvironmentId *int `json:"environmentId,omitempty"`
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    ReleaseId *int `json:"releaseId,omitempty"`
    Job *ReleaseTask `json:"job,omitempty"`
    PlanId *uuid.UUID `json:"planId,omitempty"`
    // Deprecated: 
    ReleaseDeployPhaseId *int `json:"releaseDeployPhaseId,omitempty"`
    ReleaseStepId *int `json:"releaseStepId,omitempty"`
    Tasks *[]ReleaseTask `json:"tasks,omitempty"`
}

type ReleaseTriggerType string

type ReleaseUpdatedEvent struct {
    EnvironmentId *int `json:"environmentId,omitempty"`
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
    ReleaseId *int `json:"releaseId,omitempty"`
    Release *Release `json:"release,omitempty"`
}

type ReleaseUpdateMetadata struct {
    // Sets comment for release.
    Comment *string `json:"comment,omitempty"`
    // Set 'true' to exclude the release from retention policies.
    KeepForever *bool `json:"keepForever,omitempty"`
    // Sets list of manual environments.
    ManualEnvironments *[]string `json:"manualEnvironments,omitempty"`
    // Sets name of the release.
    Name *string `json:"name,omitempty"`
    // Sets status of the release.
    Status *ReleaseStatus `json:"status,omitempty"`
}

type ReleaseWorkItemRef struct {
    // Deprecated: 
    Assignee *string `json:"assignee,omitempty"`
    // Gets or sets the ID.
    Id *string `json:"id,omitempty"`
    // Gets or sets the state.
    State *string `json:"state,omitempty"`
    // Gets or sets the title.
    Title *string `json:"title,omitempty"`
    // Gets or sets the type.
    Type *string `json:"type,omitempty"`
    // Gets or sets the workitem url.
    Url *string `json:"url,omitempty"`
}

// Represents a reference to a resource.
type ResourceReference struct {
    // An alias to be used when referencing the resource.
    Alias *string `json:"alias,omitempty"`
}

type RetentionPolicy struct {
    // Indicates the number of days to keep deployment.
    DaysToKeep *int `json:"daysToKeep,omitempty"`
}

type RetentionSettings struct {
    // Number of days to keep deleted releases.
    DaysToKeepDeletedReleases *int `json:"daysToKeepDeletedReleases,omitempty"`
    // Specifies the default environment retention policy.
    DefaultEnvironmentRetentionPolicy *EnvironmentRetentionPolicy `json:"defaultEnvironmentRetentionPolicy,omitempty"`
    // Specifies the maximum environment retention policy.
    MaximumEnvironmentRetentionPolicy *EnvironmentRetentionPolicy `json:"maximumEnvironmentRetentionPolicy,omitempty"`
}

type RunOnServerDeployPhase struct {
    // Gets and sets the agentless job input.
    DeploymentInput *ServerDeploymentInput `json:"deploymentInput,omitempty"`
}

type ScheduleDays string

type ScheduledReleaseTrigger struct {
    // Release schedule for Scheduled Release trigger type.
    Schedule *ReleaseSchedule `json:"schedule,omitempty"`
}

type SenderType string

type ServerDeploymentInput struct {
    // Gets or sets the job condition.
    Condition *string `json:"condition,omitempty"`
    // Gets or sets the job cancel timeout in minutes for deployment which are cancelled by user for this release environment.
    JobCancelTimeoutInMinutes *int `json:"jobCancelTimeoutInMinutes,omitempty"`
    // Gets or sets the override inputs.
    OverrideInputs *map[string]string `json:"overrideInputs,omitempty"`
    // Gets or sets the job execution timeout in minutes for deployment which are queued against this release environment.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
    // Gets or sets the parallel execution input.
    ParallelExecution *ExecutionInput `json:"parallelExecution,omitempty"`
}

// Represents a reference to a service endpoint.
type ServiceEndpointReference struct {
    // An alias to be used when referencing the resource.
    Alias *string `json:"alias,omitempty"`
    // The ID of the service endpoint.
    Id *uuid.UUID `json:"id,omitempty"`
}

type SingleReleaseExpands string

type SourceIdInput struct {
    // ID of source.
    Id *string `json:"id,omitempty"`
    // Name of the source.
    Name *string `json:"name,omitempty"`
}

type SourcePullRequestVersion struct {
    // Pull Request Iteration Id for which the release will publish status.
    IterationId *string `json:"iterationId,omitempty"`
    // Pull Request Id for which the release will publish status.
    PullRequestId *string `json:"pullRequestId,omitempty"`
    // Date and time of the pull request merge creation. It is required to keep timeline record of Releases created by pull request.
    PullRequestMergedAt *time.Time `json:"pullRequestMergedAt,omitempty"`
    // Source branch of the Pull Request.
    SourceBranch *string `json:"sourceBranch,omitempty"`
    // Source branch commit Id of the Pull Request for which the release will publish status.
    SourceBranchCommitId *string `json:"sourceBranchCommitId,omitempty"`
    // Target branch of the Pull Request.
    TargetBranch *string `json:"targetBranch,omitempty"`
}

type SourceRepoTrigger struct {
    // Alias of the source repo trigger.
    Alias *string `json:"alias,omitempty"`
    BranchFilters *[]string `json:"branchFilters,omitempty"`
}

type SummaryMailSection struct {
    // Html content of summary mail.
    HtmlContent *string `json:"htmlContent,omitempty"`
    // Rank of the summary mail.
    Rank *int `json:"rank,omitempty"`
    // Summary mail section type. MailSectionType has section types.
    SectionType *MailSectionType `json:"sectionType,omitempty"`
    // Title of the summary mail.
    Title *string `json:"title,omitempty"`
}

type TagFilter struct {
    // Gets or sets the tag filter pattern.
    Pattern *string `json:"pattern,omitempty"`
}

type TaskInputDefinitionBase struct {
    Aliases *[]string `json:"aliases,omitempty"`
    DefaultValue *string `json:"defaultValue,omitempty"`
    GroupName *string `json:"groupName,omitempty"`
    HelpMarkDown *string `json:"helpMarkDown,omitempty"`
    Label *string `json:"label,omitempty"`
    Name *string `json:"name,omitempty"`
    Options *map[string]string `json:"options,omitempty"`
    Properties *map[string]string `json:"properties,omitempty"`
    Required *bool `json:"required,omitempty"`
    Type *string `json:"type,omitempty"`
    Validation *TaskInputValidation `json:"validation,omitempty"`
    VisibleRule *string `json:"visibleRule,omitempty"`
}

type TaskInputValidation struct {
    // Conditional expression
    Expression *string `json:"expression,omitempty"`
    // Message explaining how user can correct if validation fails
    Message *string `json:"message,omitempty"`
}

type TaskOrchestrationPlanGroupReference struct {
    // Gets or sets the plan group.
    PlanGroup *string `json:"planGroup,omitempty"`
    // ID of the Project.
    ProjectId *uuid.UUID `json:"projectId,omitempty"`
}

type TaskOrchestrationPlanGroupsStartedEvent struct {
    PlanGroups *[]TaskOrchestrationPlanGroupReference `json:"planGroups,omitempty"`
}

type TaskSourceDefinitionBase struct {
    AuthKey *string `json:"authKey,omitempty"`
    Endpoint *string `json:"endpoint,omitempty"`
    KeySelector *string `json:"keySelector,omitempty"`
    Selector *string `json:"selector,omitempty"`
    Target *string `json:"target,omitempty"`
}

type TaskStatus string

type TfvcArtifactDownloadInput struct {
    // Gets or sets the alias of artifact.
    Alias *string `json:"alias,omitempty"`
    // Gets or sets the name of artifact definition. Valid values are 'Skip', 'Selective', 'All'.
    ArtifactDownloadMode *string `json:"artifactDownloadMode,omitempty"`
    // Gets or sets the artifact items of the input.
    ArtifactItems *[]string `json:"artifactItems,omitempty"`
    // Gets or sets the type of artifact.
    ArtifactType *string `json:"artifactType,omitempty"`
}

type TimeZone struct {
    // Display name of the time zone.
    DisplayName *string `json:"displayName,omitempty"`
    // Id of the time zone.
    Id *string `json:"id,omitempty"`
}

type TimeZoneList struct {
    // UTC timezone.
    UtcTimeZone *TimeZone `json:"utcTimeZone,omitempty"`
    // List of valid timezones.
    ValidTimeZones *[]TimeZone `json:"validTimeZones,omitempty"`
}

type VariableGroup struct {
    // Gets or sets the identity who created.
    CreatedBy *IdentityRef `json:"createdBy,omitempty"`
    // Gets date on which it got created.
    CreatedOn *time.Time `json:"createdOn,omitempty"`
    // Gets or sets description.
    Description *string `json:"description,omitempty"`
    // Gets the unique identifier of this field.
    Id *int `json:"id,omitempty"`
    // Denotes if a variable group is shared with other project or not.
    IsShared *bool `json:"isShared,omitempty"`
    // Gets or sets the identity who modified.
    ModifiedBy *IdentityRef `json:"modifiedBy,omitempty"`
    // Gets date on which it got modified.
    ModifiedOn *time.Time `json:"modifiedOn,omitempty"`
    // Gets or sets name.
    Name *string `json:"name,omitempty"`
    // Gets or sets provider data.
    ProviderData *VariableGroupProviderData `json:"providerData,omitempty"`
    // Gets or sets type.
    Type *string `json:"type,omitempty"`
    // Gets and sets the dictionary of variables.
    Variables *map[string]VariableValue `json:"variables,omitempty"`
}

type VariableGroupActionFilter string

type VariableGroupProviderData struct {
}

type VariableValue struct {
    // Gets or sets as the variable is secret or not.
    IsSecret *bool `json:"isSecret,omitempty"`
    // Gets or sets the value.
    Value *string `json:"value,omitempty"`
}

type WorkflowTask struct {
    // Gets or sets as the task always run or not.
    AlwaysRun *bool `json:"alwaysRun,omitempty"`
    // Gets or sets the task condition.
    Condition *string `json:"condition,omitempty"`
    // Gets or sets as the task continue run on error or not.
    ContinueOnError *bool `json:"continueOnError,omitempty"`
    // Gets or sets the task definition type. Example:- 'Agent', DeploymentGroup', 'Server' or 'ServerGate'.
    DefinitionType *string `json:"definitionType,omitempty"`
    // Gets or sets as the task enabled or not.
    Enabled *bool `json:"enabled,omitempty"`
    // Gets or sets the task environment variables.
    Environment *map[string]string `json:"environment,omitempty"`
    // Gets or sets the task inputs.
    Inputs *map[string]string `json:"inputs,omitempty"`
    // Gets or sets the name of the task.
    Name *string `json:"name,omitempty"`
    // Gets or sets the task override inputs.
    OverrideInputs *map[string]string `json:"overrideInputs,omitempty"`
    // Gets or sets the reference name of the task.
    RefName *string `json:"refName,omitempty"`
    // Gets or sets the ID of the task.
    TaskId *uuid.UUID `json:"taskId,omitempty"`
    // Gets or sets the task timeout.
    TimeoutInMinutes *int `json:"timeoutInMinutes,omitempty"`
    // Gets or sets the version of the task.
    Version *string `json:"version,omitempty"`
}

type WorkflowTaskReference struct {
    // Task identifier.
    Id *uuid.UUID `json:"id,omitempty"`
    // Name of the task.
    Name *string `json:"name,omitempty"`
    // Version of the task.
    Version *string `json:"version,omitempty"`
}

type YamlFileSource struct {
    // Gets or sets definition reference. e.g. {"project":{"id":"fed755ea-49c5-4399-acea-fd5b5aa90a6c","name":"myProject"},"definition":{"id":"1","name":"mybuildDefinition"},"connection":{"id":"1","name":"myConnection"}}
    SourceReference *map[string]YamlSourceReference `json:"sourceReference,omitempty"`
    Type *YamlFileSourceTypes `json:"type,omitempty"`
}

type YamlFileSourceTypes string

type YamlPipelineProcess struct {
    // Pipeline process type.
    Type *PipelineProcessTypes `json:"type,omitempty"`
    Errors *[]string `json:"errors,omitempty"`
    Filename *string `json:"filename,omitempty"`
    FileSource *YamlFileSource `json:"fileSource,omitempty"`
    Resources *YamlPipelineProcessResources `json:"resources,omitempty"`
}

type YamlPipelineProcessResources struct {
    Endpoints *[]ServiceEndpointReference `json:"endpoints,omitempty"`
    Queues *[]AgentPoolQueueReference `json:"queues,omitempty"`
}

type YamlSourceReference struct {
    Id *string `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}
