// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package location

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/identity"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6/webapi"
)

type AccessMapping struct {
	AccessPoint *string `json:"accessPoint,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Moniker     *string `json:"moniker,omitempty"`
	// The service which owns this access mapping e.g. TFS, ELS, etc.
	ServiceOwner *uuid.UUID `json:"serviceOwner,omitempty"`
	// Part of the access mapping which applies context after the access point of the server.
	VirtualDirectory *string `json:"virtualDirectory,omitempty"`
}

// Data transfer class that holds information needed to set up a connection with a VSS server.
type ConnectionData struct {
	// The Id of the authenticated user who made this request. More information about the user can be obtained by passing this Id to the Identity service
	AuthenticatedUser *identity.Identity `json:"authenticatedUser,omitempty"`
	// The Id of the authorized user who made this request. More information about the user can be obtained by passing this Id to the Identity service
	AuthorizedUser *identity.Identity `json:"authorizedUser,omitempty"`
	// The id for the server.
	DeploymentId *uuid.UUID `json:"deploymentId,omitempty"`
	// The type for the server Hosted/OnPremises.
	DeploymentType *webapi.DeploymentFlags `json:"deploymentType,omitempty"`
	// The instance id for this host.
	InstanceId *uuid.UUID `json:"instanceId,omitempty"`
	// The last user access for this instance.  Null if not requested specifically.
	LastUserAccess *azuredevops.Time `json:"lastUserAccess,omitempty"`
	// Data that the location service holds.
	LocationServiceData *LocationServiceData `json:"locationServiceData,omitempty"`
	// The virtual directory of the host we are talking to.
	WebApplicationRelativeDirectory *string `json:"webApplicationRelativeDirectory,omitempty"`
}

type InheritLevel string

type inheritLevelValuesType struct {
	None       InheritLevel
	Deployment InheritLevel
	Account    InheritLevel
	Collection InheritLevel
	All        InheritLevel
}

var InheritLevelValues = inheritLevelValuesType{
	None:       "none",
	Deployment: "deployment",
	Account:    "account",
	Collection: "collection",
	All:        "all",
}

type LocationMapping struct {
	AccessMappingMoniker *string `json:"accessMappingMoniker,omitempty"`
	Location             *string `json:"location,omitempty"`
}

// Data transfer class used to transfer data about the location service data over the web service.
type LocationServiceData struct {
	// Data about the access mappings contained by this location service.
	AccessMappings *[]AccessMapping `json:"accessMappings,omitempty"`
	// Data that the location service holds.
	ClientCacheFresh *bool `json:"clientCacheFresh,omitempty"`
	// The time to live on the location service cache.
	ClientCacheTimeToLive *int `json:"clientCacheTimeToLive,omitempty"`
	// The default access mapping moniker for the server.
	DefaultAccessMappingMoniker *string `json:"defaultAccessMappingMoniker,omitempty"`
	// The obsolete id for the last change that took place on the server (use LastChangeId64).
	LastChangeId *int `json:"lastChangeId,omitempty"`
	// The non-truncated 64-bit id for the last change that took place on the server.
	LastChangeId64 *uint64 `json:"lastChangeId64,omitempty"`
	// Data about the service definitions contained by this location service.
	ServiceDefinitions *[]ServiceDefinition `json:"serviceDefinitions,omitempty"`
	// The identifier of the deployment which is hosting this location data (e.g. SPS, TFS, ELS, Napa, etc.)
	ServiceOwner *uuid.UUID `json:"serviceOwner,omitempty"`
}

type RelativeToSetting string

type relativeToSettingValuesType struct {
	Context        RelativeToSetting
	WebApplication RelativeToSetting
	FullyQualified RelativeToSetting
}

var RelativeToSettingValues = relativeToSettingValuesType{
	Context:        "context",
	WebApplication: "webApplication",
	FullyQualified: "fullyQualified",
}

type ResourceAreaInfo struct {
	Id          *uuid.UUID `json:"id,omitempty"`
	LocationUrl *string    `json:"locationUrl,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type ServiceDefinition struct {
	Description      *string            `json:"description,omitempty"`
	DisplayName      *string            `json:"displayName,omitempty"`
	Identifier       *uuid.UUID         `json:"identifier,omitempty"`
	InheritLevel     *InheritLevel      `json:"inheritLevel,omitempty"`
	LocationMappings *[]LocationMapping `json:"locationMappings,omitempty"`
	// Maximum api version that this resource supports (current server version for this resource). Copied from <c>ApiResourceLocation</c>.
	MaxVersion *string `json:"maxVersion,omitempty"`
	// Minimum api version that this resource supports. Copied from <c>ApiResourceLocation</c>.
	MinVersion        *string            `json:"minVersion,omitempty"`
	ParentIdentifier  *uuid.UUID         `json:"parentIdentifier,omitempty"`
	ParentServiceType *string            `json:"parentServiceType,omitempty"`
	Properties        interface{}        `json:"properties,omitempty"`
	RelativePath      *string            `json:"relativePath,omitempty"`
	RelativeToSetting *RelativeToSetting `json:"relativeToSetting,omitempty"`
	// The latest version of this resource location that is in "Release" (non-preview) mode. Copied from <c>ApiResourceLocation</c>.
	ReleasedVersion *string `json:"releasedVersion,omitempty"`
	// The current resource version supported by this resource location. Copied from <c>ApiResourceLocation</c>.
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	// The service which owns this definition e.g. TFS, ELS, etc.
	ServiceOwner *uuid.UUID     `json:"serviceOwner,omitempty"`
	ServiceType  *string        `json:"serviceType,omitempty"`
	Status       *ServiceStatus `json:"status,omitempty"`
	ToolId       *string        `json:"toolId,omitempty"`
}

type ServiceStatus string

type serviceStatusValuesType struct {
	Assigned ServiceStatus
	Active   ServiceStatus
	Moving   ServiceStatus
}

var ServiceStatusValues = serviceStatusValuesType{
	Assigned: "assigned",
	Active:   "active",
	Moving:   "moving",
}
