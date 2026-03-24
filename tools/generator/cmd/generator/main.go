package main

import (
	"os"

	"github.com/ccpgames/terraform-provider-awx/tools/generator/cmd/generator/cmd"
)

func main() {
	var err error
	if err != cmd.Execute() {
		os.Exit(1)
	}
	os.Exit(0)
}
