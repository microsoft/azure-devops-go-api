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
