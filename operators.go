package bluebird

type operator string
type boolOpType bool

const (
	eq  operator = "=="
	neq operator = "!="
	lt  operator = "<"
	lte operator = "<="
	gt  operator = ">"
	gte operator = ">="
)

const (
	and boolOpType = false
	or  boolOpType = true
)

type ifExpression interface {
	eval() bool
}

type boolOp struct {
	boolType    boolOpType
	expressions []ifExpression
}

func (b *boolOp) eval() bool {
	var result bool
	if b.boolType == and {
		result = true
		for _, c := range b.expressions {
			result = c.eval() && result
			if !result {
				return false
			}
		}
	} else {
		result = false
		for _, c := range b.expressions {
			result = c.eval() || result
			if result {
				return true
			}
		}
	}
	return result
}

type comparator struct {
	left  interface{}
	right interface{}
	op    operator
}

func (c *comparator) eval() bool {

	return false
}

func parseIf(stmt string) bool {
	return false
}
