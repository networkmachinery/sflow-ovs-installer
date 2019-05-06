package main

import (
	"os"

	"github.com/networkmachinery/sflow-ovs-installer/pkg/cmd"
)

func main() {
	root := cmd.NewCmdSFlowInstaller(os.Args[1:])
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
