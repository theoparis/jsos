package std

import (
	"os/exec"
	"syscall"

	"github.com/creepinson/jsos/pkg/core/require"
	"github.com/dop251/goja"
)

func sysMount(call goja.FunctionCall) goja.Value {
	if len(call.Arguments) < 2 {
		return goja.Undefined()
	}

	fsType := call.Arguments[0].String()
	mountPoint := call.Arguments[0].String()
	mountPath := call.Arguments[1].String()

	err := syscall.Mount(fsType, mountPoint, mountPath, 0, "")

	if err != nil {
		panic(err)
	}

	return goja.Undefined()
}

func SystemModule() require.ModuleLoader {
	return func(runtime *goja.Runtime, module *goja.Object) {
		ex := module.Get("exports").(*goja.Object)
		ex.Set("mount", sysMount)
		ex.Set("exec", func(call goja.FunctionCall) goja.Value {
			if len(call.Arguments) < 1 {
				return goja.Undefined()
			}

			cmd := call.Arguments[0].String()
			args := []string{}

			for _, arg := range call.Arguments[1:] {
				args = append(args, arg.String())
			}

			command := exec.Command(cmd, args...)

			// return stdout, stderr
			stdout, err := command.Output()

			if err != nil {
				panic(err)
			}

			return runtime.ToValue(string(stdout))
		})
	}
}
