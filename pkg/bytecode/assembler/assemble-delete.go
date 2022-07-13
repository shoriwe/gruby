package assembler

import (
	"fmt"
	"github.com/shoriwe/gplasma/pkg/ast3"
	"github.com/shoriwe/gplasma/pkg/bytecode/opcodes"
	"github.com/shoriwe/gplasma/pkg/common"
	"reflect"
)

func (a *assembler) Delete(del *ast3.Delete) []byte {
	var result []byte
	switch x := del.X.(type) {
	case *ast3.Identifier:
		result = append(result, opcodes.DeleteIdentifier)
		result = append(result, common.IntToBytes(len(x.Symbol))...)
		result = append(result, []byte(x.Symbol)...)
	case *ast3.Selector:
		result = append(result, a.Expression(x.X)...)
		result = append(result, opcodes.Push)
		result = append(result, opcodes.DeleteSelector)
		result = append(result, common.IntToBytes(len(x.Identifier.Symbol))...)
		result = append(result, []byte(x.Identifier.Symbol)...)
	default:
		panic(fmt.Sprintf("invalid type of delete target %s", reflect.TypeOf(x).String()))
	}
	return result
}
