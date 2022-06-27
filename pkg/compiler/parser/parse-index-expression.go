package parser

import (
	"github.com/shoriwe/gplasma/pkg/compiler/ast"
	"github.com/shoriwe/gplasma/pkg/compiler/lexer"
)

func (parser *Parser) parseIndexExpression(expression ast.IExpression) (*ast.IndexExpression, error) {
	tokenizationError := parser.next()
	if tokenizationError != nil {
		return nil, tokenizationError
	}
	newLinesRemoveError := parser.removeNewLines()
	if newLinesRemoveError != nil {
		return nil, newLinesRemoveError
	}
	// var rightIndex ast.Node

	index, parsingError := parser.parseBinaryExpression(0)
	if parsingError != nil {
		return nil, parsingError
	}
	if _, ok := index.(ast.IExpression); !ok {
		return nil, parser.expectingExpressionError(IndexExpression)
	}
	newLinesRemoveError = parser.removeNewLines()
	if newLinesRemoveError != nil {
		return nil, newLinesRemoveError
	}
	if !parser.matchDirectValue(lexer.CloseSquareBracket) {
		return nil, parser.newSyntaxError(IndexExpression)
	}
	tokenizationError = parser.next()
	if tokenizationError != nil {
		return nil, tokenizationError
	}
	return &ast.IndexExpression{
		Source: expression,
		Index:  index.(ast.IExpression),
	}, nil
}
