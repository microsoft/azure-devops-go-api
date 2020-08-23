// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import (
	"encoding/json"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestModels_Unmarshal_Time(t *testing.T) {
	text := []byte("{\"id\":\"d221ad31-3a7b-52c0-b71d-b255b1ff63ba\",\"time1\":\"0001-01-01T00:00:00\",\"time2\":\"2019-09-01T00:07:26Z\",\"time3\":\"2020-05-16T20:55:32.0116793\",\"int\":10,\"string\":\"test string\"}")
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

	parsedTime, err := time.Parse(time.RFC3339Nano, "2019-09-01T00:07:26Z")
	if err != nil {
		t.Errorf(err.Error())
	}
	if testModel.Time2.Time != parsedTime {
		t.Errorf("Expected time: %v  Actual time: %v", parsedTime, testModel.Time2.Time)
	}

	// Test workaround for issue #59 https://github.com/microsoft/azure-devops-go-api/issues/59
	parsedTime59, err := time.Parse("2006-01-02T15:04:05.999999999", "2020-05-16T20:55:32.0116793")
	if testModel.Time3.Time != parsedTime59 {
		t.Errorf("Expected time: %v  Actual time: %v", parsedTime59, testModel.Time3.Time)
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

func TestModels_Time_AsQueryParameter(t *testing.T) {
	tt := &Time{
		// NOTE: set nsec to 0 to avoid precision loss
		Time: time.Date(2006, 1, 2, 3, 4, 5, 0, time.UTC),
	}
	queryParameterValue := tt.AsQueryParameter()
	parsedTime, err := time.Parse(time.RFC3339Nano, queryParameterValue)
	if err != nil {
		t.Errorf("invalid query parameter %s value for time: %s", tt.Time, queryParameterValue)
	}
	if !tt.Time.Equal(parsedTime) {
		t.Errorf("should convert back to original value: %s %s", tt.Time, parsedTime)
	}
}

type TestModel struct {
	Id     *uuid.UUID `json:"id,omitempty"`
	Time1  *Time      `json:"time1,omitempty"`
	Time2  *Time      `json:"time2,omitempty"`
	Time3  *Time      `json:"time3,omitempty"`
	Int    *uint64    `json:"int,omitempty"`
	String *string    `json:"string,omitempty"`
}
