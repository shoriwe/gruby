package vm

import (
	"fmt"
	"github.com/shoriwe/gplasma/pkg/lexer"
	"sync"
)

const (
	ValueId TypeId = iota
	StringId
	BytesId
	BoolId
	NoneId
	IntId
	FloatId
	ArrayId
	TupleId
	HashId
	BuiltInFunctionId
	FunctionId
	BuiltInClassId
	ClassId
)

type (
	TypeId   int
	Callback func(argument ...*Value) (*Value, error)
	FuncInfo struct {
		Arguments []string
		Bytecode  []byte
	}
	ClassInfo struct {
		prepared bool
		Bases    []*Value
		Bytecode []byte
	}
	Value struct {
		class  *Value
		typeId TypeId
		mutex  *sync.Mutex
		v      any
		vtable *Symbols
	}
)

func (plasma *Plasma) valueClass() *Value {
	class := plasma.NewValue(plasma.rootSymbols, BuiltInClassId, plasma.class)
	class.SetAny(func(argument ...*Value) (*Value, error) {
		return plasma.NewValue(plasma.rootSymbols, ValueId, plasma.value), nil
	})
	return class
}

func (value *Value) GetClass() *Value {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.class
}

func (value *Value) TypeId() TypeId {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.typeId
}

func (value *Value) VirtualTable() *Symbols {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.vtable
}

func (value *Value) SetAny(v any) {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	value.v = v
}

func (value *Value) GetHash() *Hash {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.(*Hash)
}

func (value *Value) GetCallback() Callback {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.(Callback)
}

func (value *Value) GetValues() []*Value {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.([]*Value)
}

func (value *Value) GetFuncInfo() FuncInfo {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.(FuncInfo)
}

func (value *Value) GetClassInfo() *ClassInfo {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.(*ClassInfo)
}

func (value *Value) GetBytes() []byte {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.([]byte)
}

func (value *Value) GetBool() bool {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.(bool)
}

func (value *Value) GetInt64() int64 {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.(int64)
}

func (value *Value) GetFloat64() float64 {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v.(float64)
}

func (value *Value) GetAny() any {
	value.mutex.Lock()
	defer value.mutex.Unlock()
	return value.v
}

func (value *Value) Set(symbol string, v *Value) {
	value.vtable.Set(symbol, v)
}

func (value *Value) Get(symbol string) (*Value, error) {
	return value.vtable.Get(symbol)
}

func (value *Value) Del(symbol string) error {
	return value.vtable.Del(symbol)
}

func (value *Value) Bool() bool {
	switch value.typeId {
	case ValueId:
		return true
	case StringId, BytesId:
		return len(value.GetBytes()) > 0
	case BoolId:
		return value.GetBool()
	case NoneId:
		return false
	case IntId:
		return value.GetInt64() != 0
	case FloatId:
		return value.GetFloat64() != 0
	case ArrayId, TupleId:
		return len(value.GetValues()) > 0
	case HashId:
		return value.GetHash().Size() > 0
	case BuiltInFunctionId:
		return true
	case FunctionId:
		return true
	case BuiltInClassId:
		return true
	case ClassId:
		return true
	}
	return false
}

func (value *Value) String() string {
	switch value.typeId {
	case ValueId:
		return "?Value"
	case StringId, BytesId:
		return string(value.GetBytes())
	case BoolId:
		if value.GetBool() {
			return lexer.TrueString
		}
		return lexer.FalseString
	case NoneId:
		return lexer.NoneString
	case IntId:
		return fmt.Sprintf("%d", value.GetInt64())
	case FloatId:
		return fmt.Sprintf("%f", value.GetFloat64())
	case ArrayId:
		return "[...]"
	case TupleId:
		return "(...)"
	case HashId:
		return "{...}"
	case BuiltInFunctionId:
		return "?BuiltInFunction"
	case FunctionId:
		return "?Function"
	case BuiltInClassId:
		return "?BuiltInClass"
	case ClassId:
		return "?Class"
	}
	return ""
}

func (value *Value) Contents() []byte {
	switch value.typeId {
	case ValueId:
		return nil
	case StringId, BytesId:
		return value.GetBytes()
	case BoolId:
		return nil
	case NoneId:
		return nil
	case IntId:
		return nil
	case FloatId:
		return nil
	case ArrayId:
		return nil
	case TupleId:
		return nil
	case HashId:
		return nil
	case BuiltInFunctionId:
		return nil
	case FunctionId:
		return nil
	case BuiltInClassId:
		return nil
	case ClassId:
		return nil
	}
	return nil
}

func (value *Value) Int() int64 {
	switch value.typeId {
	case ValueId:
		return 0
	case StringId:
		return 0
	case BytesId:
		return 0
	case BoolId:
		if value.GetBool() {
			return 1
		}
		return 0
	case NoneId:
		return 0
	case IntId:
		return value.GetInt64()
	case FloatId:
		return int64(value.GetFloat64())
	case ArrayId:
		return 0
	case TupleId:
		return 0
	case HashId:
		return 0
	case BuiltInFunctionId:
		return 0
	case FunctionId:
		return 0
	case BuiltInClassId:
		return 0
	case ClassId:
		return 0
	}
	return 0
}

func (value *Value) Float() float64 {
	switch value.typeId {
	case ValueId:
		return 0
	case StringId:
		return 0
	case BytesId:
		return 0
	case BoolId:
		if value.GetBool() {
			return 1
		}
		return 0
	case NoneId:
		return 0
	case IntId:
		return float64(value.GetInt64())
	case FloatId:
		return value.GetFloat64()
	case ArrayId:
		return 0
	case TupleId:
		return 0
	case HashId:
		return 0
	case BuiltInFunctionId:
		return 0
	case FunctionId:
		return 0
	case BuiltInClassId:
		return 0
	case ClassId:
		return 0
	}
	return 0
}

func (value *Value) Values() []*Value {
	switch value.typeId {
	case ValueId:
		return nil
	case StringId:
		return nil
	case BytesId:
		return nil
	case BoolId:
		return nil
	case NoneId:
		return nil
	case IntId:
		return nil
	case FloatId:
		return nil
	case ArrayId, TupleId:
		return value.GetValues()
	case HashId:
		return nil
	case BuiltInFunctionId:
		return nil
	case FunctionId:
		return nil
	case BuiltInClassId:
		return nil
	case ClassId:
		return nil
	}
	return nil
}

func (value *Value) Call(argument ...*Value) (*Value, error) {
	return value.GetCallback()(argument...)
}

/*
NewValue magic functions
TODO Not                __not__
TODO Positive           __positive__
TODO Negative           __negative__
TODO NegateBits         __negate_its__
TODO And                __and__
TODO Or                 __or__
TODO Xor                __xor__
TODO In                 __in__
TODO Is                 __is__
TODO Implements         __implements__
TODO Equals             __equals__
TODO NotEqual           __not_equal__
TODO GreaterThan        __greater_than__
TODO GreaterOrEqualThan __greater_or_equal
TODO LessThan           __less_than__
TODO LessOrEqualThan    __less_or_equal_th
TODO Add                __add__
TODO Sub                __sub__
TODO Mul                __mul__
TODO Div                __div__
TODO FloorDiv           __floor_div__
TODO Modulus            __mod__
TODO PowerOf            __pow__
TODO Length             __len__
TODO Bool               __bool__
TODO String             __string__
TODO Int                __int__
TODO Float              __float__
TODO Bytes              __bytes__
TODO Array              __array__
TODO Tuple              __tuple__
TODO Get                __get__
TODO Set                __set__
TODO Del                __del__
TODO Call               __call__
TODO Class              __class__
TODO Copy               __copy__
TODO Iter               __iter__
*/
func (plasma *Plasma) NewValue(parent *Symbols, typeId TypeId, class *Value) *Value {
	return &Value{
		class:  class,
		typeId: typeId,
		mutex:  &sync.Mutex{},
		v:      nil,
		vtable: NewSymbols(parent),
	}
}
