package utils

import "os"

type OSFileSystem struct {
}

func NewOSFileSystem() *OSFileSystem {
	return &OSFileSystem{}
}

func Read(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func Write(filename string, content []byte) error {
	return os.WriteFile(filename, content, os.ModePerm)
}
