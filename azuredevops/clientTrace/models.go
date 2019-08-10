// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package clienttrace

type ClientTraceEvent struct {
	Area          *string                 `json:"area,omitempty"`
	Component     *string                 `json:"component,omitempty"`
	ExceptionType *string                 `json:"exceptionType,omitempty"`
	Feature       *string                 `json:"feature,omitempty"`
	Level         *Level                  `json:"level,omitempty"`
	Message       *string                 `json:"message,omitempty"`
	Method        *string                 `json:"method,omitempty"`
	Properties    *map[string]interface{} `json:"properties,omitempty"`
}

type Level string

type levelValuesType struct {
	Off     Level
	Error   Level
	Warning Level
	Info    Level
	Verbose Level
}

var LevelValues = levelValuesType{
	Off:     "off",
	Error:   "error",
	Warning: "warning",
	Info:    "info",
	Verbose: "verbose",
}
