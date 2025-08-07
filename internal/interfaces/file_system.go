package interfaces

type FileSystem interface {
	Read(filename string) ([]byte, error)
	Write(filename string, content []byte) error
}
