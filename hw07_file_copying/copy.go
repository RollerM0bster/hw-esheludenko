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
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var file *os.File
	err := checkPaths(fromPath, toPath)
	if err != nil {
		return err
	}
	file, _ = os.Open(fromPath)
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}
	fileSize := fileInfo.Size()
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 || limit > fileSize-offset {
		limit = fileSize - offset
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

	progressBar := pb.Full.Start64(limit)
	defer progressBar.Finish()

	progressReader := progressBar.NewProxyReader(io.LimitReader(file, limit))

	_, err = io.Copy(writeTo, progressReader)

	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

func checkPaths(fromPath, toPath string) error {
	src, err := os.Stat(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	dst, err := os.Stat(toPath)
	if err == nil {
		if os.SameFile(src, dst) {
			return errors.New("source and destination are the same")
		}
	}
	return nil
}
