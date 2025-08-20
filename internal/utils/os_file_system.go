package utils

import "os"

type OSFileSystem struct {
}

func NewOSFileSystem() *OSFileSystem {
	return &OSFileSystem{}
}

func (fs *OSFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (fs *OSFileSystem) WriteFile(filename string, content []byte) error {
	return os.WriteFile(filename, content, os.ModePerm)
}
