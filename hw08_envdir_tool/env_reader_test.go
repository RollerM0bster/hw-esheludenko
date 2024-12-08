package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir := t.TempDir()

	testFiles := map[string]string{
		"BAR":        "123\n",
		"EMPTY":      "",
		"FOO":        "foo\t\t\r",
		"INVALID=22": "skip",
		"NULL_BYTE":  "filestart\x00filefinish",
	}

	expected := map[string]EnvValue{
		"BAR":       {Value: "123", NeedRemove: true},
		"EMPTY":     {Value: "", NeedRemove: false},
		"FOO":       {Value: "foo", NeedRemove: true},
		"NULL_BYTE": {Value: "filestart\nfilefinish", NeedRemove: false},
	}
	for name, content := range testFiles {
		err := os.WriteFile(
			filepath.Join(dir, name),
			[]byte(content), 0644)
		if err != nil {
			t.Fatal(err)
		}
		os.Setenv("BAR", "123")
		os.Setenv("FOO", "456")
	}

	os.Setenv("FOO", "some_env_var")
	defer os.Unsetenv("FOO")

	env, err := ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for key, expectedValue := range expected {
		value, exists := env[key]
		if !exists {
			t.Errorf("Expected key %s not found in the result", key)
			continue
		}
		if value.Value != expectedValue.Value {
			t.Errorf("For key %s, expected value %q, got %q", key, expectedValue.Value, value.Value)
		}
		if value.NeedRemove != expectedValue.NeedRemove {
			t.Errorf("For key %s, expected NeedRemove %v, got %v", key, expectedValue.NeedRemove, value.NeedRemove)
		}
	}
	if _, exists := env["INVALID=KEY"]; exists {
		t.Errorf("Expected invalid key 'INVALID=KEY' to be skipped")
	}
}
