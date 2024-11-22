package main

import "testing"

type TestCaseStockBuyAndSellI struct {
	name           string
	input          []int
	expectedOutput int
}

func TestMaxProfitI(t *testing.T) {
	tcs := []TestCaseStockBuyAndSellI{
		{
			name:           "TC01",
			input:          []int{7, 10, 1, 3, 6, 9, 2},
			expectedOutput: 8,
		},
		{
			name:           "TC02",
			input:          []int{7, 6, 4, 3, 1},
			expectedOutput: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			resp := MaxProfitI(tc.input)
			if tc.expectedOutput != resp {
				t.Fatalf("expected %v but got %v", tc.expectedOutput, resp)
			}
		})
	}
}
