package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name           string
	input          []uint32
	expectedOutput []uint32
}

func TestNextPermutation(t *testing.T) {
	tcs := []TestCase{
		{
			name:           "TC01",
			input:          []uint32{2, 4, 1, 7, 5, 0},
			expectedOutput: []uint32{2, 4, 5, 0, 1, 7},
		},
		{
			name:           "TC02",
			input:          []uint32{3, 2, 1},
			expectedOutput: []uint32{1, 2, 3},
		},
		{
			name:           "TC03",
			input:          []uint32{3, 4, 2, 5, 1},
			expectedOutput: []uint32{3, 4, 5, 1, 2},
		},
		{
			name:           "TC04",
			input:          []uint32{1, 3, 2},
			expectedOutput: []uint32{2, 1, 3},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			NextPermutation(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expectedOutput) {
				t.Fatalf("expected %v but got %v", tc.expectedOutput, tc.input)
			}
		})
	}
}
