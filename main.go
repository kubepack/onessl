package main

import (
	"os"

	logs "github.com/appscode/go/log/golog"
	"github.com/kubepack/onessl/cmds"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	_, plugin := os.LookupEnv("KUBECTL_PLUGINS_CALLER")
	if err := cmds.NewRootCmd(Version, plugin).Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
