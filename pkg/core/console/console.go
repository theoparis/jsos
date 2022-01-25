package console

import (
	"fmt"

	"github.com/dop251/goja"
)

type Console struct{}

func NewConsole(vm *goja.Runtime) *goja.Object {
	c := &Console{}
	obj := vm.NewObject()
	obj.Set("log", c.log)

	return obj
}

func (c *Console) log(call goja.FunctionCall) goja.Value {
	for _, arg := range call.Arguments {
		fmt.Println(arg.String())
	}

	return goja.Undefined()
}
