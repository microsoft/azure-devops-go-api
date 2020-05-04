// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Time struct {
	Time time.Time
}

func tryUnmarshalJSON(b []byte) (time.Time, error) {
	t := time.Time{}
	err := json.Unmarshal(b, &t)

	// ignore errors for 0001-01-01T00:00:00 dates. The Azure DevOps service
	// returns default dates in a format that is invalid for a time.Time. The
	// correct value would have a 'z' at the end to represent utc. We are going
	// to ignore this error, and set the value to the default time.Time value.
	// https://github.com/microsoft/azure-devops-go-api/issues/17
	if err != nil {
		if parseError, ok := err.(*time.ParseError); ok && parseError.Value == "\"0001-01-01T00:00:00\"" {
			err = nil
		}
	}

	return t, err
}

func makeTimeParseFunction(layout string) func(b []byte) (time.Time, error) {
	return func(b []byte) (time.Time, error) {
		s := strings.Trim(string(b), "\"")
		if s == "null" {
			return time.Time{}, nil
		}
		return time.Parse(layout, s)
	}
}

// UnmarshalJSON attempts a variety of time marshalling strategies. The first one that works
// will be used. If none succeede, then an error is returned.
func (t *Time) UnmarshalJSON(b []byte) error {

	// The dates returned by the Azure DevOps services do not conform to any predictable
	// format or standard. The strategy list here tries to account for as many permutations
	// as possible by applying the following strategies:
	//
	//	(1) Default Go time deserialization
	//	(2) Attempt each of the time formats defined as constants in the time library.
	//		These are documented here: https://golang.org/pkg/time/#pkg-constants
	//	(3) Special case times found to be returned by the Azure DevOps services
	//
	// The first "working" strategy is used.
	strategies := []func([]byte) (time.Time, error){
		tryUnmarshalJSON,
		makeTimeParseFunction(time.ANSIC),
		makeTimeParseFunction(time.UnixDate),
		makeTimeParseFunction(time.RubyDate),
		makeTimeParseFunction(time.RFC822),
		makeTimeParseFunction(time.RFC822Z),
		makeTimeParseFunction(time.RFC850),
		makeTimeParseFunction(time.RFC1123),
		makeTimeParseFunction(time.RFC1123Z),
		makeTimeParseFunction(time.RFC3339),
		makeTimeParseFunction(time.RFC3339Nano),
		makeTimeParseFunction(time.Stamp),
		makeTimeParseFunction(time.StampMilli),
		makeTimeParseFunction(time.StampMicro),
		makeTimeParseFunction(time.StampNano),
		makeTimeParseFunction("2006-01-02T15:04:05.9999999"), // https://github.com/microsoft/azure-devops-go-api/issues/59
	}

	for _, strategy := range strategies {
		parsedTime, err := strategy(b)
		if err == nil {
			t.Time = parsedTime
			return nil
		}
	}

	return fmt.Errorf("Unable to deserialize time (%s). %d strategies were attempted", string(b), len(strategies))
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

func (t Time) String() string {
	return t.Time.String()
}

func (t Time) Equal(u Time) bool {
	return t.Time.Equal(u.Time)
}
