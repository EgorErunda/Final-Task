package calculator_test

import (
	"testing"

	"github.com/EgorErunda/FinalTaskSprint_1/pkg/calculator"
)

func TestCalc(t *testing.T) {
	PositiveCases := []struct {
		name       string
		expression string
		result     float64
	}{
		{
			name:       "easy",
			expression: "2+2",
			result:     4,
		},
		{
			name:       "check work of parenthesis",
			expression: "(3+4)*5",
			result:     35,
		},
		{
			name:       "importance of multiplication",
			expression: "3+4*5",
			result:     23,
		},
		{
			name:       "check /",
			expression: "2/8",
			result:     0.25,
		},
	}

	for _, testCase := range PositiveCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculator.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.result {
				t.Fatalf("%f should be equal %f", val, testCase.result)
			}
		})
	}

	NegativeCases := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "easy",
			expression: "1+1*",
		},
		{
			name:       "priority",
			expression: "2+2**2",
		},
		{
			name:       "priority",
			expression: "((2+2-*(2",
		},
		{
			name:       "/",
			expression: "",
		},
	}

	for _, testCase := range NegativeCases {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculator.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is incorrect but result  %f was get", testCase.expression, val)
			}
		})
	}
}
