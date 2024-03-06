package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {

	fromFile, err := os.OpenFile(fromPath, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}
	defer fromFile.Close()

	toFile, err := os.Create(toPath)
	if err != nil { // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	defer toFile.Close()

	io.CopyN(toFile, fromFile, offset)
	return nil
}
