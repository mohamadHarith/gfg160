package main

import (
	"testing"
)

type TestCaseStockBuyAndSellII struct {
	name           string
	input          []int
	expectedOutput int
}

func TestMaxProfitII(t *testing.T) {
	tcs := []TestCaseStockBuyAndSellII{
		{
			name:           "TC01",
			input:          []int{100, 180, 260, 310, 40, 535, 695},
			expectedOutput: 865,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			resp := MaxProfitII(tc.input)
			if tc.expectedOutput != resp {
				t.Fatalf("expected %v but got %v", tc.expectedOutput, resp)
			}
		})
	}
}
