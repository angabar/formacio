package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	calculate("What is 3 plus 5")
}

func calculate(expression string) int32 {
	expression = strings.Replace(expression, "What is", "", -1)
	expression = strings.Replace(expression, "plus", "+", -1)
	expression = strings.Replace(expression, "minus", "-", -1)
	expression = strings.Replace(expression, "multiplied", "*", -1)
	expression = strings.Replace(expression, "divided", "/", -1)

	tr, err := parser.ParseExpr(expression)

	if err != nil {
		fmt.Println(err)
	}

	fs := token.NewFileSet()
	ast.Print(fs, tr)

	return 3
}
