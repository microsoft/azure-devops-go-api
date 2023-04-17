package azuredevops

import (
	"testing"
)

func TestNegotiateRequestVersion_NonReleasedApi(t *testing.T) {
	// key is requested version, value is expected negotiated version
	tests := make(map[string]string)
	tests[""] = ""
	tests["2.2-preview.4"] = "2.2-preview.3"
	tests["2.2-preview"] = "2.2-preview"
	tests["2.2-preview.3"] = "2.2-preview.3"
	tests["2.2-preview.2"] = "2.2-preview.2"
	tests["2.1-preview.4"] = "2.1-preview.3"
	tests["2.1-preview.3"] = "2.1-preview.3"
	tests["1.0-preview.3"] = "1.0-preview.3"
	tests["2.2"] = "2.2-preview"
	tests["1.0"] = "1.0-preview"
	tests["2.2-preview"] = "2.2-preview"
	tests["1.0-preview"] = "1.0-preview"
	tests["2.3-preview"] = "2.2-preview"
	tests["3.0-preview"] = "2.2-preview"

	for requested, expected := range tests {
		compareVersions(t, unreleasedApiResourceLocation, requested, expected)
	}
}

func TestNegotiateRequestVersion_ReleasedApi(t *testing.T) {
	// key is requested version, value is expected negotiated version
	tests := make(map[string]string)
	tests[""] = ""
	tests["2.2-preview.4"] = "2.2-preview.3"
	tests["2.2-preview"] = "2.2-preview"
	tests["2.2-preview.3"] = "2.2-preview.3"
	tests["2.2-preview.2"] = "2.2-preview.2"
	tests["2.1-preview.4"] = "2.1-preview.3"
	tests["2.1-preview.3"] = "2.1-preview.3"
	tests["1.0-preview.3"] = "1.0-preview.3"
	tests["2.2-preview"] = "2.2-preview"
	tests["1.0-preview"] = "1.0-preview"
	tests["2.3-preview"] = "2.2-preview"
	tests["3.0-preview"] = "2.2-preview"
	tests["3.0"] = "2.2-preview"
	tests["2.3"] = "2.2-preview"
	tests["2.2"] = "2.2-preview"
	tests["2.1"] = "2.1"
	tests["2.0"] = "2.0"
	tests["1.0"] = "1.0"

	for requested, expected := range tests {
		compareVersions(t, releasedApiResourceLocation, requested, expected)
	}
}

func compareVersions(t *testing.T, apiResourceLocation *ApiResourceLocation, requestedVersion string, expectedVersion string) {
	negotiatedVersion, err := negotiateRequestVersion(apiResourceLocation, requestedVersion)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if negotiatedVersion != expectedVersion {
		t.Errorf("Negotiated version did not match expected. Requested: %v  Expected: %v  Actual: %v", requestedVersion, expectedVersion, negotiatedVersion)
	}
}

var resourceVersionThree = 3
var version00 = "0.0"
var version10 = "1.0"
var version20 = "2.0"
var version21 = "2.1"
var version22 = "2.2"

var unreleasedApiResourceLocation = &ApiResourceLocation{
	ResourceVersion: &resourceVersionThree,
	MinVersion:      &version10,
	MaxVersion:      &version22,
	ReleasedVersion: &version00,
}

var releasedApiResourceLocation = &ApiResourceLocation{
	ResourceVersion: &resourceVersionThree,
	MinVersion:      &version20,
	MaxVersion:      &version22,
	ReleasedVersion: &version21,
}
