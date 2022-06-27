package parser

import (
	"github.com/shoriwe/gplasma/pkg/compiler/ast"
	"github.com/shoriwe/gplasma/pkg/compiler/lexer"
)

func (parser *Parser) parseYieldStatement() (*ast.YieldStatement, error) {
	tokenizingError := parser.next()
	if tokenizingError != nil {
		return nil, tokenizingError
	}
	newLinesRemoveError := parser.removeNewLines()
	if newLinesRemoveError != nil {
		return nil, newLinesRemoveError
	}
	var results []ast.IExpression
	for parser.hasNext() {
		if parser.matchKind(lexer.Separator) || parser.matchKind(lexer.EOF) {
			break
		}

		result, parsingError := parser.parseBinaryExpression(0)
		if parsingError != nil {
			return nil, parsingError
		}
		if _, ok := result.(ast.IExpression); !ok {
			return nil, parser.expectingExpressionError(YieldStatement)
		}
		results = append(results, result.(ast.IExpression))
		if parser.matchDirectValue(lexer.Comma) {
			tokenizingError = parser.next()
			if tokenizingError != nil {
				return nil, tokenizingError
			}
		} else if !(parser.matchKind(lexer.Separator) || parser.matchKind(lexer.EOF)) {
			return nil, parser.newSyntaxError(YieldStatement)
		}
	}
	return &ast.YieldStatement{
		Results: results,
	}, nil
}
