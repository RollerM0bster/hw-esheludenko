package main

import (
	"errors"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	srcFile, err := os.CreateTemp("", "srcFile*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary source file: %v", err)
	}
	defer os.Remove(srcFile.Name())
	srcContent := []byte("Hello, World!")
	srcFile.Write(srcContent)
	srcFile.Close()

	tests := []struct {
		name        string
		fromPath    string
		toPath      string
		offset      int64
		limit       int64
		wantErr     error
		wantContent string
	}{
		{
			name:        "Normal Copy",
			fromPath:    srcFile.Name(),
			toPath:      "dst_normal.txt",
			offset:      0,
			limit:       5,
			wantErr:     nil,
			wantContent: "Hello",
		},
		{
			name:     "Offset Exceeds File Size",
			fromPath: srcFile.Name(),
			toPath:   "dst_offset_exceeds.txt",
			offset:   int64(len(srcContent) + 1),
			limit:    5,
			wantErr:  ErrOffsetExceedsFileSize,
		},
		{
			name:        "Limit Larger Than Remaining Data",
			fromPath:    srcFile.Name(),
			toPath:      "dst_limit_large.txt",
			offset:      7,
			limit:       50,
			wantErr:     nil,
			wantContent: "World!",
		},
		{
			name:        "Zero Limit (Copy All from Offset)",
			fromPath:    srcFile.Name(),
			toPath:      "dst_zero_limit.txt",
			offset:      7,
			limit:       0,
			wantErr:     nil,
			wantContent: "World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.fromPath, tt.toPath, tt.offset, tt.limit)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Copy() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == nil {
				content, err := os.ReadFile(tt.toPath)
				if err != nil {
					t.Fatalf("Failed to read destination file: %v", err)
				}
				if string(content) != tt.wantContent {
					t.Errorf("File content = %q, wantContent = %q", string(content), tt.wantContent)
				}
				os.Remove(tt.toPath)
			}
		})
	}
}
