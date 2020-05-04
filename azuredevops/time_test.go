// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

type TimeTestStruct struct {
	T Time
}

func getJSONBytes(timeStr string) []byte {
	return []byte(fmt.Sprintf("{ \"T\": \"%s\" }", timeStr))
}

func TestTime_MarshallingRoundTripForSupportedCases(t *testing.T) {
	testCases := []struct {
		name    string
		timeStr string
	}{
		{
			"ANSIC",
			"Mon Jan 2 15:04:05 2006",
		}, {
			"UnixDate",
			"Mon Jan 2 15:04:05 MST 2006",
		}, {
			"RubyDate",
			"Mon Jan 02 15:04:05 -0700 2006",
		}, {
			"RFC822",
			"02 Jan 06 15:04 MST",
		}, {
			"RFC822Z",
			"02 Jan 06 15:04 -0700",
		}, {
			"RFC850",
			"Monday, 02-Jan-06 15:04:05 MST",
		}, {
			"RFC1123",
			"Mon, 02 Jan 2006 15:04:05 MST",
		}, {
			"RFC1123Z",
			"Mon, 02 Jan 2006 15:04:05 -0700",
		}, {
			"RFC3339_1",
			"2006-01-02T15:04:05Z",
		}, {
			"RFC3339_2",
			"2006-01-02T15:04:05+07:00",
		}, {
			"RFC3339Nano_1",
			"2006-01-02T15:04:05.999999999Z",
		}, {
			"RFC3339Nano_2",
			"2006-01-02T15:04:05.999999999+07:00",
		}, {
			"Stamp",
			"Jan 2 15:04:05",
		}, {
			"StampMilli",
			"Jan 2 15:04:05.000",
		}, {
			"StampMicro",
			"Jan 2 15:04:05.000000",
		}, {
			"StampNano",
			"Jan 2 15:04:05.000000000",
		}, {
			"GitHubIssue59",
			"2006-01-02T15:04:05.9999999",
		},
	}

	for _, testCase := range testCases {
		expectedTime, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05+00:00")

		t.Run(testCase.name, func(t *testing.T) {
			s := TimeTestStruct{}
			err := json.Unmarshal(getJSONBytes(testCase.timeStr), &s)
			if err != nil {
				t.Errorf("Unexpectedly could not unmarshal %s into a valid time struct: %+v", testCase.timeStr, err)
			}

			// Not all timestamps have the same year and timezone information. So we can pick a subset of properties to test
			if expectedTime.Day() != s.T.Time.Day() {
				t.Errorf("Expected day %+v but got %+v", expectedTime.Day(), s.T.Time.Day())
			}

			if expectedTime.Hour() != s.T.Time.Hour() {
				t.Errorf("Expected hour %+v but got %+v", expectedTime.Hour(), s.T.Time.Hour())
			}

			if expectedTime.Minute() != s.T.Time.Minute() {
				t.Errorf("Expected minute %+v but got %+v", expectedTime.Minute(), s.T.Time.Minute())
			}
		})
	}
}

func TestModels_Unmarshal_Time(t *testing.T) {
	text := []byte("{\"id\":\"d221ad31-3a7b-52c0-b71d-b255b1ff63ba\",\"time1\":\"0001-01-01T00:00:00\",\"time2\":\"2019-09-01T00:07:26Z\",\"int\":10,\"string\":\"test string\"}")
	testModel := TestModel{}

	testModel.Time1 = &Time{}
	testModel.Time1.Time = time.Now() // this ensures we test the value is set back to default when issue #17 is hit.

	err := json.Unmarshal(text, &testModel)
	if err != nil {
		t.Errorf("Error occurred during deserialization: %v", err)
	}
	if (testModel.Time1.Time != time.Time{}) {
		t.Errorf("Expecting deserialized time to equal default time.  Actual time: %v", testModel.Time1)
	}

	parsedTime, err := time.Parse(time.RFC3339, "2019-09-01T00:07:26Z")
	if err != nil {
		t.Errorf(err.Error())
	}
	if testModel.Time2.Time != parsedTime {
		t.Errorf("Expected time: %v  Actual time: %v", parsedTime, testModel.Time1.Time)
	}
}

func TestModels_Marshal_Unmarshal_Time(t *testing.T) {
	testModel1 := TestModel{}
	testModel1.Time1 = &Time{}
	testModel1.Time1.Time = time.Now()
	b, err := json.Marshal(testModel1)
	if err != nil {
		t.Errorf(err.Error())
	}

	testModel2 := TestModel{}
	err = json.Unmarshal(b, &testModel2)
	if err != nil {
		t.Errorf(err.Error())
	}

	if testModel1.Time1 != testModel1.Time1 {
		t.Errorf("Expected time: %v  Actual time: %v", testModel1.Time1, testModel1.Time2)
	}

	if testModel1.Time1.Time != testModel1.Time1.Time {
		t.Errorf("Expected time: %v  Actual time: %v", testModel1.Time1.Time, testModel1.Time2.Time)
	}
}

type TestModel struct {
	Id     *uuid.UUID `json:"id,omitempty"`
	Time1  *Time      `json:"time1,omitempty"`
	Time2  *Time      `json:"time2,omitempty"`
	Int    *uint64    `json:"int,omitempty"`
	String *string    `json:"string,omitempty"`
}
