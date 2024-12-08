package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	res := make(Environment)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if strings.Contains(name, "=") {
			fmt.Fprintf(os.Stderr, "Skipping invalid file name %s\n", name)
			continue
		}
		path := filepath.Join(dir, name)
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		if info.Size() == 0 {
			res[name] = EnvValue{Value: "", NeedRemove: true}
		}
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %s\n", name, err)
			os.Exit(1)
		}
		lines := strings.Split(string(content), "\n")
		firstLine := lines[0]
		content = bytes.Replace([]byte(firstLine), []byte{0x00}, []byte("\n"), 1)
		strVal := strings.TrimRight(string(content), " \t\r")
		res[name] = EnvValue{Value: strVal, NeedRemove: false}
		_, exists := os.LookupEnv(name)
		if exists {
			temp := res[name]
			temp.NeedRemove = true
			res[name] = temp
		}
	}
	return res, nil
}
