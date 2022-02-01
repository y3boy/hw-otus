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
	if fileSize, err := GetFileSize(fromPath); offset > fileSize || limit > fileSize {
		return ErrOffsetExceedsFileSize
	} else if err != nil {
		return ErrUnsupportedFile
	}
	
	source, err := os.Open(fromPath)
	defer source.Close()
	if err != nil {
		return ErrUnsupportedFile
	}

	fileToWrite, _ := os.Create(toPath)
	defer fileToWrite.Close()
	if limit == 0 {
			io.Copy(fileToWrite, source)
	} 
	io.CopyN(fileToWrite, source, limit)
	
	
	return nil
}
