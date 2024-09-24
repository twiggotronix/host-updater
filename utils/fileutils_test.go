package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckFileIsWrittableSuccess(t *testing.T) {
	mockOpener := MockFileOpener{
		OpenFileFunc: func(name string, flag int, perm os.FileMode) (*os.File, error) {
			if name == "test.txt" {
				return &os.File{}, nil
			}
			return nil, errors.New("file not found")
		},
	}
	fileUtilsFactoy := FileUtilsFactoy{}
	fileUtils := fileUtilsFactoy.GetFileUtils()
	fileUtils.FileOpener = mockOpener

	result := fileUtils.CheckFileIsWrittable("test.txt")

	assert.True(t, result)
}
func TestCheckFileIsWrittableFail(t *testing.T) {
	mockOpener := MockFileOpener{
		OpenFileFunc: func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return nil, errors.New("file not found")
		},
	}
	fileUtilsFactoy := FileUtilsFactoy{}
	fileUtils := fileUtilsFactoy.GetFileUtils()
	fileUtils.FileOpener = mockOpener

	result := fileUtils.CheckFileIsWrittable("test.txt")

	assert.False(t, result)
}

func TestSetNewIp(t *testing.T) {
	var writtenContent string
	mockOpener := MockFileOpener{
		OpenFileFunc: func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return &os.File{}, nil
		},
		ReadFileFunc: func(name string) ([]byte, error) {
			return []byte("# some stuff at the beginning\r\n123.123.2.1 testhost [flag]\r\n192.168.1.1 nottobechanged"), nil
		},
		WriteFileFunc: func(name string, data []byte, fileMode fs.FileMode) error {
			fmt.Println(string(data))
			writtenContent = string(data)
			return nil
		},
	}
	fileUtilsFactoy := FileUtilsFactoy{}
	fileUtils := fileUtilsFactoy.GetFileUtils()
	fileUtils.FileOpener = mockOpener

	err := fileUtils.SetNewIp("test.txt", "[flag]", "192.168.0.1", false)

	assert.Nil(t, err)
	assert.Equal(t, "# some stuff at the beginning\r\n192.168.0.1 testhost [flag]\r\n192.168.1.1 nottobechanged", writtenContent)
}

func TestSetNewIpFileErrorFileNotWrittable(t *testing.T) {
	mockOpener := MockFileOpener{
		OpenFileFunc: func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return nil, errors.New("file not found")
		},
	}
	fileUtilsFactoy := FileUtilsFactoy{}
	fileUtils := fileUtilsFactoy.GetFileUtils()
	fileUtils.FileOpener = mockOpener

	err := fileUtils.SetNewIp("test.txt", "[flag]", "192.168.0.1", false)

	assert.Equal(t, "test.txt is not writtable", err.Message)
}
func TestSetNewIpFileErrorFileNotReadable(t *testing.T) {
	mockOpener := MockFileOpener{
		OpenFileFunc: func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return &os.File{}, nil
		},
		ReadFileFunc: func(name string) ([]byte, error) {
			return nil, errors.New("oh no...")
		},
	}
	fileUtilsFactoy := FileUtilsFactoy{}
	fileUtils := fileUtilsFactoy.GetFileUtils()
	fileUtils.FileOpener = mockOpener

	err := fileUtils.SetNewIp("test.txt", "[flag]", "192.168.0.1", false)

	assert.Equal(t, "error while reading to file : test.txt oh no...", err.Message)
}
