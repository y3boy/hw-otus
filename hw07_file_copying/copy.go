package main

import (
	"errors"
	"io"
	"os"
)

// var : different type of errors.
var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

// GetFileSize is simple funnction.
func GetFileSize(fromPath string) (int64, error) {
	fi, err := os.Stat(fromPath)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// Copy function.
func Copy(fromPath, toPath string, offset, limit int64) error {
	if fileSize, err := GetFileSize(fromPath); offset > fileSize || offset < 0 || limit < 0 {
		return ErrOffsetExceedsFileSize
	} else if err != nil {
		return ErrUnsupportedFile
	}

	source, err := os.OpenFile(fromPath, os.O_RDWR, 0o755)
	if err != nil {
		return ErrUnsupportedFile
	}

	fileToWrite, _ := os.OpenFile(toPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	source.Seek(offset, 0)

	if limit == 0 {
		io.Copy(fileToWrite, source)
	} else {
		io.CopyN(fileToWrite, source, limit)
	}

	source.Close()
	fileToWrite.Close()
	return nil
}
