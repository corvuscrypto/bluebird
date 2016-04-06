package bluebird

import (
	"fmt"
	"strings"
)

type Node interface {
	eval() interface{}
}

type IfNode struct {
	op    ExprToken
	left  *Node
	right *Node
}

type ValueNode struct {
	value interface{}
}

type ExprToken int

const (
	AND ExprToken = iota + 1
	OR
	EQ
	LT
	GT
	LTE
	GTE

	NOT

	SUB
	ADD
	DIV
	MUL
	MOD

	BEGINSCOPE
	ENDSCOPE
)

var tokenMap = map[string]ExprToken{
	"+":  ADD,
	"-":  SUB,
	"*":  MUL,
	"/":  DIV,
	"%":  MOD,
	"!":  NOT,
	"&&": AND,
	"||": OR,
	"==": EQ,
	"<":  LT,
	">":  GT,
	"<=": LTE,
	">=": GTE,
	"(":  BEGINSCOPE,
	")":  ENDSCOPE,
}

var boundaryChars = map[rune]bool{
	'&': true,
	'|': true,
	'<': true,
	'>': true,
	'=': true,
	'(': true,
	')': true,
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'%': true,
	' ': true,
}

//iterate through the characters in the string and parse as we go
func parseExpr(stmt string) {
	tokens := tokenizeExpr(stmt)
	for _, a := range tokens {
		fmt.Println(a, getPrecedence(a))
	}

}

func getPrecedence(tkn string) int {
	expr := tokenMap[tkn]
	if expr == 0 {
		return 0
	} else if expr <= GTE {
		return 1
	} else {
		return int(expr)
	}
}

func tokenizeExpr(stmt string) []string {
	result := []string{}
	flag := true
	//this will be used to hold some temp string data
	temp := ""
	//iterate yo!
	for _, char := range stmt {
		if boundaryChars[char] == flag || char == '(' || char == ')' || char == ' ' {
			if temp != " " && temp != "" {
				result = append(result, strings.Trim(temp, " "))
			}
			temp = string(char)
			flag = !isReserved(char)
		} else {
			temp += string(char)
		}
	}
	result = append(result, strings.Trim(temp, " "))
	return result
}

func isReserved(ch rune) bool {
	return boundaryChars[ch]
}
