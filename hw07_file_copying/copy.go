package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotExists         = errors.New("file does not exist")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var file *os.File
	file, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrFileNotExists
		}
		return ErrUnsupportedFile
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	fileSize := fileInfo.Size()
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 || limit > fileSize {
		limit = fileSize
	}
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}
	writeTo, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer writeTo.Close()

	progressBar := pb.New(int(limit)).Start()
	defer progressBar.Finish()

	progressReader := progressBar.NewProxyReader(file)

	_, err = io.CopyN(writeTo, progressReader, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	defer file.Close()
	return nil
}
