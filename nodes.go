package bluebird

type comparatorNode struct {
	placeholder *template
	comparators []comparator
	template    string
}

func (node *comparatorNode) parse() {

}

type recursiveNode struct {
	placeholder *template
	condition   string
	template    string
}

type includeNode struct {
	placeholder *template
	path        string
}
