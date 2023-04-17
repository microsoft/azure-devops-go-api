// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package featureavailability

type FeatureFlag struct {
	Description    *string `json:"description,omitempty"`
	EffectiveState *string `json:"effectiveState,omitempty"`
	ExplicitState  *string `json:"explicitState,omitempty"`
	Name           *string `json:"name,omitempty"`
	Uri            *string `json:"uri,omitempty"`
}

// This is passed to the FeatureFlagController to edit the status of a feature flag
type FeatureFlagPatch struct {
	State *string `json:"state,omitempty"`
}
