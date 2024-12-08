package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	setVars(env)
	commandPath := cmd[2]
	args := cmd[3:]
	command := exec.Command(commandPath, args...)
	command.Env = os.Environ()
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			os.Exit(exitError.ExitCode())
		}
		fmt.Fprintf(os.Stderr, "Error running command: %v\n", err)
		os.Exit(1)
	}
	return
}

func setVars(env Environment) {
	for key := range env {
		item := env[key]
		if item.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				return
			}
		}
		if item.Value != "" {
			err := os.Setenv(key, item.Value)
			if err != nil {
				return
			}
		}
	}
}
