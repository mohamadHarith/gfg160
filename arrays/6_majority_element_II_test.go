package main

import (
	"reflect"
	"testing"
)

type TestCaseMajorityElement struct {
	name           string
	input          []int
	expectedOutput []int
}

func TestMajorityElement(t *testing.T) {
	tcs := []TestCaseMajorityElement{
		{
			name:           "TC01",
			input:          []int{2, 1, 5, 5, 5, 5, 6, 6, 6, 6, 6},
			expectedOutput: []int{5, 6},
		},
		{
			name:           "TC02",
			input:          []int{1, 2, 3, 4, 5},
			expectedOutput: []int{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			resp := MajorityElement(tc.input)
			if len(tc.expectedOutput) != len(resp) {
				t.Fatalf("expected %v but got %v", tc.expectedOutput, resp)
			}
			if !reflect.DeepEqual(resp, tc.expectedOutput) && len(tc.expectedOutput) != 0 {
				t.Fatalf("expected %v but got %v", tc.expectedOutput, resp)
			}
		})
	}
}
