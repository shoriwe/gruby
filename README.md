# plasma

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shoriwe/gplasma)
[![Go Report Card](https://goreportcard.com/badge/github.com/shoriwe/gplasma)](https://goreportcard.com/report/github.com/shoriwe/gplasma)

Plasma is an embeddable scripting ruby like language.

<img src="logos/plasma-logos.jpeg" alt="logo" style="zoom:50%;" />

## Features

- Simple extensibility: Easy to add new go bindings
- Zero dependency: The language used as a library doesn't depend on any external project.
- Thread safe: the virtual machine and all the objects created during runtime are thread safe
- Rich syntax: Generators, defer, special boolean operators and more (check documentation for more details)
- Bytecode VM backend: the language compiles to a custom bytecode that can then stored and preloaded in the machine
  without recompiling scripts.
- Stop vm execution: Plasma let you stop at any the time the execution of the VM.

## Documentation

You can find documentation in:

- [Official documentation](https://shoriwe.github.io/plasma/index.html)
- [pkg.go.dev](https://pkg.go.dev/github.com/shoriwe/gplasma)

## Install interpreter

```shell
go install github.com/shoriwe/gplasma/cmd/plasma@latest
```

## Preview

## REPL

You can start a REPL with:

```shell
plasma
```

![Repl demo](demos/repl-demo.gif)

## Embedding and creating Go bindings

```go
package main

import (
	"github.com/shoriwe/gplasma/pkg/vm"
	"os"
)

const myScript = `
args = get_args()
if args.__len__() > 1
    println(args.__string__())
else
    println("No")
end
`

func main() {
	plasma := vm.NewVM(os.Stdin, os.Stdout, os.Stderr)
	plasma.Load("get_args", func(plasma *vm.Plasma) *vm.Value {
		return plasma.NewBuiltInFunction(plasma.Symbols(),
			func(argument ...*vm.Value) (*vm.Value, error) {
				tupleValues := make([]*vm.Value, 0, len(os.Args))
				for _, cmdArg := range os.Args {
					tupleValues = append(tupleValues, plasma.NewString([]byte(cmdArg)))
				}
				return plasma.NewTuple(tupleValues), nil
			})
	})
	_, errorChannel, _ := plasma.ExecuteString(myScript)
	err := <-errorChannel
	if err != nil {
		panic(err)
	}
}
```

## Contributing

To contribute to this project please follow the [contribution guidelines](CONTRIBUTING.md) and
the [code of conduct](CODE_OF_CONDUCT.md)