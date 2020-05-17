// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package provenance

type SessionRequest struct {
	// Generic property bag to store data about the session
	Data *map[string]string `json:"data,omitempty"`
	// The feed name or id for the session
	Feed *string `json:"feed,omitempty"`
	// The type of session If a known value is provided, the Data dictionary will be validated for the presence of properties required by that type
	Source *string `json:"source,omitempty"`
}

type SessionResponse struct {
	// The unique identifier for the session
	SessionId *string `json:"sessionId,omitempty"`
	// The name for the session
	SessionName *string `json:"sessionName,omitempty"`
}
