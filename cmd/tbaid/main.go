package main

import (
	"github.com/ktoyod/tbaid/pkg/cmd"
)

func main() {
	command := cmd.NewTbaidCommand()
	command.Execute()
}
