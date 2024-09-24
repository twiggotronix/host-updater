package utils

import (
	"io/fs"
	"os"
)

type FileOpener interface {
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, fileMode fs.FileMode) error
}

type OSFileOpener struct{}

func (o OSFileOpener) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (o OSFileOpener) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (o OSFileOpener) WriteFile(name string, data []byte, fileMode fs.FileMode) error {
	return os.WriteFile(name, data, fileMode)
}
