// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package profileregions

type GeoRegion struct {
	RegionCode *string `json:"regionCode,omitempty"`
}

// Country/region information
type ProfileRegion struct {
	// The two-letter code defined in ISO 3166 for the country/region.
	Code *string `json:"code,omitempty"`
	// Localized country/region name
	Name *string `json:"name,omitempty"`
}

// Container of country/region information
type ProfileRegions struct {
	// List of country/region code with contact consent requirement type of notice
	NoticeContactConsentRequirementRegions *[]string `json:"noticeContactConsentRequirementRegions,omitempty"`
	// List of country/region code with contact consent requirement type of opt-out
	OptOutContactConsentRequirementRegions *[]string `json:"optOutContactConsentRequirementRegions,omitempty"`
	// List of country/regions
	Regions *[]ProfileRegion `json:"regions,omitempty"`
}
