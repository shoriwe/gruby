package parser

import "github.com/shoriwe/gplasma/pkg/compiler/ast"

func (parser *Parser) parsePassStatement() (*ast.PassStatement, error) {
	tokenizingError := parser.next()
	if tokenizingError != nil {
		return nil, tokenizingError
	}
	return &ast.PassStatement{}, nil
}
