package parser

import (
	"fmt"
	"github.com/shoriwe/gruby/pkg/compiler/ast"
	"github.com/shoriwe/gruby/pkg/compiler/lexer"
	"testing"
)

func walker(node ast.Node, deep int) {
	switch node.(type) {
	case *ast.Program:
		for _, child := range node.(*ast.Program).Body {
			walker(child, deep+1)
		}
	case *ast.BinaryExpression:
		walker(node.(*ast.BinaryExpression).LeftHandSide, deep+1)
		fmt.Print(node.(*ast.BinaryExpression).Operator)
		walker(node.(*ast.BinaryExpression).RightHandSide, deep+1)
	case *ast.BasicLiteralExpression:
		fmt.Print(node.(*ast.BasicLiteralExpression).String)
	case *ast.UnaryExpression:
		fmt.Print(node.(*ast.UnaryExpression).Operator)
		walker(node.(*ast.UnaryExpression).X, deep+1)
	case *ast.SelectorExpression:
		walker(node.(*ast.SelectorExpression).X, deep+1)
		fmt.Print("." + node.(*ast.SelectorExpression).Identifier.String)
	case *ast.Identifier:
		fmt.Print(node.(*ast.Identifier).String)
	case *ast.MethodInvocationExpression:
		walker(node.(*ast.MethodInvocationExpression).Function, deep+1)
		fmt.Print("(")
		isFirst := true
		for _, child := range node.(*ast.MethodInvocationExpression).Arguments {
			if isFirst {
				isFirst = false
			} else {
				fmt.Print(", ")
			}
			walker(child, deep+1)
		}
		fmt.Print(")")
	case *ast.IndexExpression:
		walker(node.(*ast.IndexExpression).Source, deep+1)
		fmt.Print("[")

		walker(node.(*ast.IndexExpression).Index[0], deep+1)
		if node.(*ast.IndexExpression).Index[1] != nil {
			fmt.Print(":")
			walker(node.(*ast.IndexExpression).Index[1], deep+1)
		}
		fmt.Print("]")
	case *ast.LambdaExpression:
		fmt.Print("lambda")
		isFirst := true
		for _, argument := range node.(*ast.LambdaExpression).Arguments {
			if isFirst {
				fmt.Print(" ")
				isFirst = false
			} else {
				fmt.Print(", ")
			}
			walker(argument, deep+1)
		}
		fmt.Print(": ")
		walker(node.(*ast.LambdaExpression).Code, deep+1)
	}
}

func walk(node ast.Node) {
	walker(node, 0)
	fmt.Println("")
}

func test(t *testing.T, samples []string) {
	for _, sample := range samples {
		lex := lexer.NewLexer(sample)
		parser, parserCreationError := NewParser(lex)
		if parserCreationError != nil {
			t.Error(parserCreationError)
			return
		}
		program, parsingError := parser.Parse()
		if parsingError != nil {
			t.Error(parsingError)
			return
		}
		walk(program)
	}
}

var basicSamples = []string{
	"1 + 2 * 3",
	"1 * 2 + 3",
	"1 >= 2 == 3 - 4 + 5 - 6 / 7 / 8 ** 9 - 10",
	"5--5",
	"hello.world.carro",
	"1.4.hello.world()",
	"hello(1)",
	"'Hello world'.index(str(12345)[0])",
	"lambda x, y, z: print(x, y - z)",
}

func TestParseBasic(t *testing.T) {
	test(t, basicSamples)
}
