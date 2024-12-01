package main

import "testing"

func TestRunCmd(t *testing.T) {
	env := Environment{
		"FOO":  EnvValue{Value: "bar", NeedRemove: false},
		"TEST": EnvValue{Value: "value", NeedRemove: true},
	}

	t.Run("echo", func(t *testing.T) {
		cmd := []string{"./go-envdir", "hw08_envdir_tool/testdata/env", "echo", "123", "456"}
		returnCode := RunCmd(cmd, env)
		if returnCode == 0 {
			t.Errorf("Unexpected return code: %d", returnCode)
		}
	})
}
