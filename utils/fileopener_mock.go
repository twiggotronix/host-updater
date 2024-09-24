package utils

import (
	"errors"
	"io/fs"
	"os"
)

type MockFileOpener struct {
	OpenFileFunc  func(name string, flag int, perm os.FileMode) (*os.File, error)
	ReadFileFunc  func(name string) ([]byte, error)
	WriteFileFunc func(name string, data []byte, fileMode fs.FileMode) error
}

func (m MockFileOpener) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if m.OpenFileFunc != nil {
		return m.OpenFileFunc(name, flag, perm)
	}
	return nil, errors.New("not implemented")
}

func (m MockFileOpener) ReadFile(name string) ([]byte, error) {
	if m.OpenFileFunc != nil {
		return m.ReadFileFunc(name)
	}
	return nil, errors.New("not implemented")
}

func (m MockFileOpener) WriteFile(name string, data []byte, fileMode fs.FileMode) error {
	if m.OpenFileFunc != nil {
		return m.WriteFileFunc(name, data, fileMode)
	}
	return errors.New("not implemented")
}
