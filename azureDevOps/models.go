// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azureDevOps

import (
	"github.com/google/uuid"
	"strconv"
)

type ApiResourceLocation struct {
	Id *uuid.UUID `json:"id,omitempty"`
	Area *string `json:"area,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	RouteTemplate *string `json:"routeTemplate,omitempty"`
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	MinVersion *string `json:"minVersion,omitempty"`
	MaxVersion *string `json:"maxVersion,omitempty"`
	ReleasedVersion *string `json:"releasedVersion,omitempty"`
}

type WrappedImproperError struct {
	Count *int `json:"count,omitEmpty"`
	Value *ImproperError `json:"value,omitEmpty"`
}

type ImproperError struct {
	Message *string `json:"Message,omitEmpty"`
}

type KeyValuePair struct {
	Key *interface{} `json:"key,omitEmpty"`
	Value *interface{} `json:"value,omitEmpty"`
}

type ResourceAreaInfo struct {
	Id *uuid.UUID `json:"id,omitempty"`
	LocationUrl *string `json:"locationUrl,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ServerSystemError struct {
	ClassName *string `json:"className,omitempty"`
	InnerException *ServerSystemError `json:"innerException,omitempty"`
	Message *string `json:"message,omitempty"`
}

func (e ServerSystemError) Error() string {
	return *e.Message
}

type VssJsonCollectionWrapper struct {
	Count *int `json:"count"`
	Value *[]interface{} `json:"value"`
}

type WrappedError struct {
	ExceptionId *string                      `json:"$id,omitempty"`
	InnerError *WrappedError                 `json:"innerException,omitempty"`
	Message *string                          `json:"message,omitempty"`
	TypeName *string                         `json:"typeName,omitempty"`
	TypeKey *string                          `json:"typeKey,omitempty"`
	ErrorCode *int                           `json:"errorCode,omitempty"`
	EventId *int                             `json:"eventId,omitempty"`
	CustomProperties *map[string]interface{} `json:"customProperties,omitempty"`
	StatusCode *int
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

