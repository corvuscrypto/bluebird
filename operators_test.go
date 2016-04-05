package bluebird

import "testing"

var testIfStatements = []string{
	"7 == 8",
	"7 == 8 || 8 >= 7",
	"7 == 8 && 8 <= 9",
	"7 == 8 || (8 >= 9 || 7 == 7 )",
	"7 == 8 || (8 <= 9 && 7 == 7 )",
	"7 == 8 || (8 <= 9 && 7 == 8 )",
	"7 == 7 || (8 <= 9 && 7 == 8 )",
}

var testIfStatementsResults = []bool{
	false,
	true,
	false,
	true,
	true,
	false,
	true,
}

func TestIfOperators(T *testing.T) {
	for i, v := range testIfStatements {
		result := parseIf(v)
		answer := testIfStatementsResults[i]
		if result != answer {
			T.Errorf("%s: expected %t, got %t", v, answer, result)
		}
	}
}
