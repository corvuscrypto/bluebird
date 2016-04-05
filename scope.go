package bluebird

type scopeToken string
type scopeStack struct {
	stack []interface{}
}

const (
	scopeBeginToken scopeToken = "("
	scopeEndToken   scopeToken = ")"
)

func (s *scopeStack) begin(v interface{}) {
	s.stack = append(s.stack, v)
}

func (s *scopeStack) end() {
	s.stack = s.stack[:len(s.stack)-1]
}
