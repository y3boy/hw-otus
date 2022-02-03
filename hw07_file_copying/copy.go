package main

import (
	"errors"
	"io"
	"os"
)

// var : diffrent type of errors
var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrNoneValideOperation   = errors.New("none valide operation")
)

// GetFileSize is simple funnction
func GetFileSize(fromPath string) (int64, error) {
	fi, err := os.Stat(fromPath)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// Copy function
func Copy(fromPath, toPath string, offset, limit int64) error {
	if fileSize, err := GetFileSize(fromPath); offset > fileSize || limit > fileSize || offset < 0 || limit < 0{
		return ErrOffsetExceedsFileSize
	} else if err != nil {
		return ErrUnsupportedFile
	}
	
	source, err := os.OpenFile(fromPath, os.O_RDWR, 0755)
	defer source.Close()
	if err != nil {
		return ErrUnsupportedFile
	}

	fileToWrite, _ := os.OpenFile(toPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	source.Seek(offset, 0)
	defer fileToWrite.Close()
	if limit == 0 {
			io.Copy(fileToWrite, source)
	} else {
		io.CopyN(fileToWrite, source, limit)
	}
	return nil
}
