package bluebird

import "testing"

func TestExpression(T *testing.T) {
	parseExpr("((test || (test3))) && test2 * 3")
}
