// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azureDevOps

import (
	"github.com/google/uuid"
	"strconv"
)

// Information about the location of a REST API resource
type ApiResourceLocation struct {
	// Area name for this resource
	Area *string `json:"area,omitempty"`
	// Unique Identifier for this location
	Id *uuid.UUID `json:"id,omitempty"`
	// Maximum api version that this resource supports (current server version for this resource)
	MaxVersion *string `json:"maxVersion,omitempty"`
	// Minimum api version that this resource supports
	MinVersion *string `json:"minVersion,omitempty"`
	// The latest version of this resource location that is in "Release" (non-preview) mode
	ReleasedVersion *string `json:"releasedVersion,omitempty"`
	// Resource name
	ResourceName *string `json:"resourceName,omitempty"`
	// The current resource version supported by this resource location
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	// This location's route template (templated relative path)
	RouteTemplate *string `json:"routeTemplate,omitempty"`
}

type WrappedImproperError struct {
	Count *int           `json:"count,omitEmpty"`
	Value *ImproperError `json:"value,omitEmpty"`
}

type ImproperError struct {
	Message *string `json:"Message,omitEmpty"`
}

type KeyValuePair struct {
	Key   *interface{} `json:"key,omitEmpty"`
	Value *interface{} `json:"value,omitEmpty"`
}

type ResourceAreaInfo struct {
	Id          *uuid.UUID `json:"id,omitempty"`
	LocationUrl *string    `json:"locationUrl,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

type ServerSystemError struct {
	ClassName      *string            `json:"className,omitempty"`
	InnerException *ServerSystemError `json:"innerException,omitempty"`
	Message        *string            `json:"message,omitempty"`
}

func (e ServerSystemError) Error() string {
	return *e.Message
}

type VssJsonCollectionWrapper struct {
	Count *int           `json:"count"`
	Value *[]interface{} `json:"value"`
}

type WrappedError struct {
	ExceptionId      *string                 `json:"$id,omitempty"`
	InnerError       *WrappedError           `json:"innerException,omitempty"`
	Message          *string                 `json:"message,omitempty"`
	TypeName         *string                 `json:"typeName,omitempty"`
	TypeKey          *string                 `json:"typeKey,omitempty"`
	ErrorCode        *int                    `json:"errorCode,omitempty"`
	EventId          *int                    `json:"eventId,omitempty"`
	CustomProperties *map[string]interface{} `json:"customProperties,omitempty"`
	StatusCode       *int
}

func (e WrappedError) Error() string {
	if e.Message == nil {
		if e.StatusCode != nil {
			return "REST call returned status code " + strconv.Itoa(*e.StatusCode)
		}
		return ""
	}
	return *e.Message
}
