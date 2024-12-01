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
		"BAR":       {Value: "123", NeedRemove: false},
		"EMPTY":     {Value: "", NeedRemove: false},
		"FOO":       {Value: "foo", NeedRemove: false},
		"NULL_BYTE": {Value: "filestart\nfilefinish", NeedRemove: false},
	}
	for name, content := range testFiles {
		err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0644)
		if err != nil {
			t.Fatal(err)
		}
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
		if key == "FOO" && !value.NeedRemove {
			t.Errorf("Expected NeedRemove to be true for key %s, got false", key)
		}
	}
	if _, exists := env["INVALID=KEY"]; exists {
		t.Errorf("Expected invalid key 'INVALID=KEY' to be skipped")
	}
}
