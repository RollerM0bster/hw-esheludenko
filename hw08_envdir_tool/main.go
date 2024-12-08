package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	envs, err := ReadDir(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	RunCmd(args, envs)
	os.Exit(0)
}
