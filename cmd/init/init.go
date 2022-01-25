package main

import (
	"fmt"
	"os"

	"github.com/creepinson/jsos/pkg/core/console"
	"github.com/creepinson/jsos/pkg/core/eventloop"
	"github.com/creepinson/jsos/pkg/core/require"
	"github.com/creepinson/jsos/pkg/std"
	"github.com/dop251/goja"
)

func main() {
	registry := new(require.Registry) // this can be shared by multiple runtimes

	if len(os.Args) < 2 {
		fmt.Println("No file specified")
	}

	file := os.Args[1]
	// read file
	f, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	loop := eventloop.NewEventLoop()

	program, err := goja.Compile(file, string(f), false)
	if err != nil {
		fmt.Printf("Compilation error: %s\n", err)
	}

	loop.Run(func(vm *goja.Runtime) {
		registry.Enable(vm)
		registry.RegisterNativeModule("std/sys", std.SystemModule())
		c := console.NewConsole(vm)
		vm.Set("console", c)
		// vm.Set("sys", std.SystemModule())

		_, err := vm.RunProgram(program)

		if err != nil {
			fmt.Printf("Runtime error: %s\n", err)
		}
	})

	// get all .js files in /lib using glob
	// for each file, require it
	// files, err := filepath.Glob("/lib/**/*.js")

	// if err != nil {
	// 	panic(err)
	// }

	// for _, file := range files {
	// 	m, err := req.Require("m.js")
	// 	_, _ = m, err
	// }

}
