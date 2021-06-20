package vm

import "math/big"

type KeyValue struct {
	Key   Value
	Value Value
}

type HashTable struct {
	*Object
}

func (p *Plasma) NewHashTable(isBuiltIn bool, parent *SymbolTable, entries map[int64][]*KeyValue, entriesLength int) *HashTable {
	hashTable := &HashTable{
		Object: p.NewObject(false, HashName, nil, parent),
	}
	hashTable.SetKeyValues(entries)
	hashTable.SetLength(entriesLength)
	p.HashTableInitialize(isBuiltIn)(hashTable)
	hashTable.Set(Self, hashTable)
	return hashTable
}

func (p *Plasma) HashTableInitialize(isBuiltIn bool) ConstructorCallBack {
	return func(object Value) *Object {
		object.Set(Equals,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 1,
					func(self Value, arguments ...Value) (Value, *Object) {

						rawRight := arguments[0]
						if _, ok := rawRight.(*HashTable); !ok {
							return p.NewBool(false, p.PeekSymbolTable(), false), nil
						}
						right := rawRight.(*HashTable)
						if self.GetLength() != right.Length {
							return p.NewBool(false, p.PeekSymbolTable(), false), nil
						}
						rightIndex, getError := right.Get(Index)
						if getError != nil {
							return nil, p.NewObjectWithNameNotFoundError(Index)
						}
						if _, ok := rightIndex.(*Function); !ok {
							return nil, p.NewInvalidTypeError(rightIndex.TypeName(), FunctionName)
						}
						for key, leftValue := range self.GetKeyValues() {
							// Check if other has the key
							rightValue, ok := right.KeyValues[key]
							if !ok {
								return p.NewBool(false, p.PeekSymbolTable(), false), nil
							}
							// Check if the each entry one has the same length
							if len(leftValue) != len(rightValue) {
								return p.NewBool(false, p.PeekSymbolTable(), false), nil
							}
							// Start comparing the entries
							for _, entry := range leftValue {
								_, indexingError := p.CallFunction(rightIndex.(*Function), p.PeekSymbolTable(), entry.Key)
								if indexingError != nil {
									return p.NewBool(false, p.PeekSymbolTable(), false), nil
								}
							}
						}
						return p.NewBool(false, p.PeekSymbolTable(), true), nil
					},
				),
			),
		)
		object.Set(RightEquals,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 1,
					func(self Value, arguments ...Value) (Value, *Object) {

						rawLeft := arguments[0]
						if _, ok := rawLeft.(*HashTable); !ok {
							return p.NewBool(false, p.PeekSymbolTable(), false), nil
						}
						left := rawLeft.(*HashTable)
						if self.GetLength() != left.Length {
							return p.NewBool(false, p.PeekSymbolTable(), false), nil
						}
						leftIndex, getError := left.Get(Index)
						if getError != nil {
							return nil, p.NewObjectWithNameNotFoundError(Index)
						}
						if _, ok := leftIndex.(*Function); !ok {
							return nil, p.NewInvalidTypeError(leftIndex.TypeName(), FunctionName)
						}
						for key, leftValue := range left.KeyValues {
							// Check if other has the key
							rightValue, ok := self.GetKeyValues()[key]
							if !ok {
								return p.NewBool(false, p.PeekSymbolTable(), false), nil
							}
							// Check if the each entry one has the same length
							if len(leftValue) != len(rightValue) {
								return p.NewBool(false, p.PeekSymbolTable(), false), nil
							}
							// Start comparing the entries
							for _, entry := range leftValue {
								_, indexingError := p.CallFunction(leftIndex.(*Function), p.PeekSymbolTable(), entry.Key)
								if indexingError != nil {
									return p.NewBool(false, p.PeekSymbolTable(), false), nil
								}
							}
						}
						return p.NewBool(false, p.PeekSymbolTable(), true), nil
					},
				),
			),
		)
		object.Set(NotEquals,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 1,
					func(self Value, arguments ...Value) (Value, *Object) {
						rawRight := arguments[0]
						if _, ok := rawRight.(*HashTable); !ok {
							return p.NewBool(false, p.PeekSymbolTable(), true), nil
						}
						right := rawRight.(*HashTable)
						if self.GetLength() != right.Length {
							return p.NewBool(false, p.PeekSymbolTable(), true), nil
						}
						rightIndex, getError := right.Get(Index)
						if getError != nil {
							return nil, p.NewObjectWithNameNotFoundError(Index)
						}
						if _, ok := rightIndex.(*Function); !ok {
							return nil, p.NewInvalidTypeError(rightIndex.TypeName(), FunctionName)
						}
						for key, leftValue := range self.GetKeyValues() {
							// Check if other has the key
							rightValue, ok := right.KeyValues[key]
							if !ok {
								return p.NewBool(false, p.PeekSymbolTable(), true), nil
							}
							// Check if the each entry one has the same length
							if len(leftValue) != len(rightValue) {
								return p.NewBool(false, p.PeekSymbolTable(), true), nil
							}
							// Start comparing the entries
							for _, entry := range leftValue {
								_, indexingError := p.CallFunction(rightIndex.(*Function), p.PeekSymbolTable(), entry.Key)
								if indexingError != nil {
									return p.NewBool(false, p.PeekSymbolTable(), true), nil
								}
							}
						}
						return p.NewBool(false, p.PeekSymbolTable(), false), nil
					},
				),
			),
		)
		object.Set(RightNotEquals,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 1,
					func(self Value, arguments ...Value) (Value, *Object) {
						rawLeft := arguments[0]
						if _, ok := rawLeft.(*HashTable); !ok {
							return p.NewBool(false, p.PeekSymbolTable(), true), nil
						}
						left := rawLeft.(*HashTable)
						if self.GetLength() != left.Length {
							return p.NewBool(false, p.PeekSymbolTable(), true), nil
						}
						leftIndex, getError := left.Get(Index)
						if getError != nil {
							return nil, p.NewObjectWithNameNotFoundError(Index)
						}
						if _, ok := leftIndex.(*Function); !ok {
							return nil, p.NewInvalidTypeError(leftIndex.TypeName(), FunctionName)
						}
						for key, leftValue := range left.KeyValues {
							// Check if other has the key
							rightValue, ok := self.GetKeyValues()[key]
							if !ok {
								return p.NewBool(false, p.PeekSymbolTable(), true), nil
							}
							// Check if the each entry one has the same length
							if len(leftValue) != len(rightValue) {
								return p.NewBool(false, p.PeekSymbolTable(), true), nil
							}
							// Start comparing the entries
							for _, entry := range leftValue {
								_, indexingError := p.CallFunction(leftIndex.(*Function), p.PeekSymbolTable(), entry.Key)
								if indexingError != nil {
									return p.NewBool(false, p.PeekSymbolTable(), true), nil
								}
							}
						}
						return p.NewBool(false, p.PeekSymbolTable(), false), nil
					},
				),
			),
		)

		object.Set(Hash,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 0,
					func(_ Value, _ ...Value) (Value, *Object) {
						panic("Implement me!!!")
					},
				),
			),
		)
		object.Set(Copy,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 0,
					func(_ Value, _ ...Value) (Value, *Object) {
						return nil, p.NewUnhashableTypeError(object.GetClass())
					},
				),
			),
		)
		object.Set(Index,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 1,
					func(self Value, arguments ...Value) (Value, *Object) {
						indexObject := arguments[0]
						indexObjectHash, getError := indexObject.Get(Hash)
						if getError != nil {
							return nil, p.NewObjectWithNameNotFoundError(Hash)
						}
						if _, ok := indexObjectHash.(*Function); !ok {
							return nil, p.NewInvalidTypeError(indexObjectHash.TypeName(), FunctionName)
						}
						indexHash, callError := p.CallFunction(indexObjectHash.(*Function), indexObject.SymbolTable())
						if callError != nil {
							return nil, callError
						}
						if _, ok := indexHash.(*Integer); !ok {
							return nil, p.NewInvalidTypeError(indexHash.TypeName(), IntegerName)
						}
						keyValues, found := self.GetKeyValues()[indexHash.GetInteger().Int64()]
						if !found {
							return nil, p.NewKeyNotFoundError(indexObject)
						}
						var indexObjectEquals Value
						indexObjectEquals, getError = indexObject.Get(Equals)
						if _, ok := indexObjectEquals.(*Function); !ok {
							return nil, p.NewInvalidTypeError(indexObjectEquals.TypeName(), FunctionName)
						}
						var equals Value
						for _, keyValue := range keyValues {
							equals, callError = p.CallFunction(indexObjectEquals.(*Function), indexObject.SymbolTable(), keyValue.Key)
							if callError != nil {
								return nil, callError
							}
							if _, ok := equals.(*Bool); !ok {
								return nil, p.NewInvalidTypeError(equals.TypeName(), BoolName)
							}
							if equals.GetBool() {
								return keyValue.Value, nil
							}
						}
						return nil, p.NewKeyNotFoundError(indexObject)
					},
				),
			),
		)
		object.Set(Assign,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 2,
					func(self Value, arguments ...Value) (Value, *Object) {
						indexObject := arguments[0]
						newValue := arguments[1]
						indexObjectHash, getError := indexObject.Get(Hash)
						if getError != nil {
							return nil, p.NewObjectWithNameNotFoundError(Hash)
						}
						if _, ok := indexObjectHash.(*Function); !ok {
							return nil, p.NewInvalidTypeError(indexObjectHash.TypeName(), FunctionName)
						}
						indexHash, callError := p.CallFunction(indexObjectHash.(*Function), indexObject.SymbolTable())
						if callError != nil {
							return nil, callError
						}
						if _, ok := indexHash.(*Integer); !ok {
							return nil, p.NewInvalidTypeError(indexHash.TypeName(), IntegerName)
						}
						keyValues, found := self.GetKeyValues()[indexHash.GetInteger().Int64()]
						if found {
							self.AddKeyValue(indexHash.GetInteger().Int64(), &KeyValue{
								Key:   indexObject,
								Value: newValue,
							})
							return p.NewNone(), nil
						}
						var indexObjectEquals Value
						indexObjectEquals, getError = indexObject.Get(Equals)
						if _, ok := indexObjectEquals.(*Function); !ok {
							return nil, p.NewInvalidTypeError(indexObjectEquals.TypeName(), FunctionName)
						}
						var equals Value
						for index, keyValue := range keyValues {
							equals, callError = p.CallFunction(indexObjectEquals.(*Function), indexObject.SymbolTable(), keyValue.Key)
							if callError != nil {
								return nil, callError
							}
							if _, ok := equals.(*Bool); !ok {
								return nil, p.NewInvalidTypeError(equals.TypeName(), BoolName)
							}
							if equals.GetBool() {
								self.GetKeyValues()[indexHash.GetInteger().Int64()][index].Value = newValue
								return p.NewNone(), nil
							}
						}
						self.IncreaseLength()
						self.GetKeyValues()[indexHash.GetInteger().Int64()] = append(
							self.GetKeyValues()[indexHash.GetInteger().Int64()],
							&KeyValue{
								Key:   indexObject,
								Value: newValue,
							},
						)
						return p.NewNone(), nil
					},
				),
			),
		)
		object.Set(Iter,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 0,
					func(self Value, _ ...Value) (Value, *Object) {
						toTuple, getError := self.Get(ToTuple)
						if getError != nil {
							return nil, p.NewObjectWithNameNotFoundError(ToTuple)
						}
						if _, ok := toTuple.(*Function); !ok {
							return nil, p.NewInvalidTypeError(toTuple.TypeName(), FunctionName)
						}
						hashKeys, callError := p.CallFunction(toTuple.(*Function), self.SymbolTable())
						if callError != nil {
							return nil, callError
						}
						iterator := p.NewIterator(false, p.PeekSymbolTable())
						iterator.SetInteger(big.NewInt(0)) // This is the index
						iterator.SetContent(hashKeys.GetContent())
						iterator.SetLength(len(hashKeys.GetContent()))
						iterator.Set(HasNext,
							p.NewFunction(isBuiltIn, iterator.SymbolTable(),
								NewBuiltInClassFunction(iterator,
									0,
									func(funcSelf Value, _ ...Value) (Value, *Object) {
										return p.NewBool(false, p.PeekSymbolTable(), funcSelf.GetInteger().Cmp(big.NewInt(int64(funcSelf.GetLength()))) == -1), nil
									},
								),
							),
						)
						iterator.Set(Next,
							p.NewFunction(isBuiltIn, iterator.SymbolTable(),
								NewBuiltInClassFunction(iterator,
									0,
									func(funcSelf Value, _ ...Value) (Value, *Object) {
										value := funcSelf.GetContent()[int(funcSelf.GetInteger().Int64())]
										funcSelf.SetInteger(new(big.Int).Add(funcSelf.GetInteger(), big.NewInt(1)))
										return value, nil
									},
								),
							),
						)
						return iterator, nil
					},
				),
			),
		)

		object.Set(ToString,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 0,
					func(self Value, _ ...Value) (Value, *Object) {
						result := "{"
						var (
							keyString     Value
							valueToString Value
							valueString   Value
							callError     *Object
						)
						for _, keyValues := range self.GetKeyValues() {
							for _, keyValue := range keyValues {
								keyToString, getError := keyValue.Key.Get(ToString)
								if getError != nil {
									return nil, p.NewObjectWithNameNotFoundError(ToString)
								}
								keyString, callError = p.CallFunction(keyToString.(*Function), keyValue.Key.SymbolTable())
								if callError != nil {
									return nil, callError
								}
								result += keyString.GetString()
								valueToString, getError = keyValue.Value.Get(ToString)
								if getError != nil {
									return nil, p.NewObjectWithNameNotFoundError(ToString)
								}
								valueString, callError = p.CallFunction(valueToString.(*Function), keyValue.Value.SymbolTable())
								if callError != nil {
									return nil, callError
								}
								result += ": " + valueString.GetString() + ", "
							}
						}
						if len(result) > 1 {
							result = result[:len(result)-2]
						}
						return p.NewString(false, p.PeekSymbolTable(), result+"}"), nil
					},
				),
			),
		)
		object.Set(ToBool,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 0,
					func(self Value, _ ...Value) (Value, *Object) {
						if self.GetLength() > 0 {
							return p.NewBool(false, p.PeekSymbolTable(), true), nil
						}
						return p.NewBool(false, p.PeekSymbolTable(), false), nil
					},
				),
			),
		)
		object.Set(ToArray,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 0,
					func(self Value, _ ...Value) (Value, *Object) {
						var keys []Value
						for _, keyValues := range self.GetKeyValues() {
							for _, keyValue := range keyValues {
								keys = append(keys, keyValue.Key)
							}
						}
						return p.NewArray(false, p.PeekSymbolTable(), keys), nil
					},
				),
			),
		)
		object.Set(ToTuple,
			p.NewFunction(isBuiltIn, object.SymbolTable(),
				NewBuiltInClassFunction(object, 0,
					func(self Value, _ ...Value) (Value, *Object) {
						var keys []Value
						for _, keyValues := range self.GetKeyValues() {
							for _, keyValue := range keyValues {
								keys = append(keys, keyValue.Key)
							}
						}
						return p.NewTuple(false, p.PeekSymbolTable(), keys), nil
					},
				),
			),
		)
		return nil
	}
}
