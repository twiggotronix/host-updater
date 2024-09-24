package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type FileUtils struct {
	FileOpener FileOpener
}

type FileUtilsFactoy struct {
	fileUtilsInstance *FileUtils
}

type FileUtilsError struct {
	Message string
}

// Implement the error interface
func (e *FileUtilsError) Error() string {
	return e.Message
}

func (fileUtilsFactoy FileUtilsFactoy) GetFileUtils() FileUtils {
	if fileUtilsFactoy.fileUtilsInstance == nil {
		fileUtilsFactoy.fileUtilsInstance = &FileUtils{
			FileOpener: OSFileOpener{},
		}
	}

	return *fileUtilsFactoy.fileUtilsInstance
}

func (fileUtils FileUtils) CheckFileIsWrittable(fileName string) bool {
	f, err := fileUtils.FileOpener.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

func (fileUtils FileUtils) SetNewIp(fileName string, flag string, selectedAddress string, dryRun bool) *FileUtilsError {
	if !fileUtils.CheckFileIsWrittable(fileName) {
		fmt.Printf("%s is not writtable\n", fileName)
		if !dryRun {
			return &FileUtilsError{Message: fmt.Sprintf("%s is not writtable", fileName)}
		}
	}
	biteContent, err := fileUtils.FileOpener.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error while reading to file : %s %s", fileName, err)
		return &FileUtilsError{Message: fmt.Sprintf("error while reading to file : %s %s", fileName, err)}
	}
	biteContent = NormalizeNewlines(biteContent)
	eofSlice := strings.Split(string(biteContent), "\n")
	re := regexp.MustCompile(`(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)\.(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)\.(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)\.(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)`)
	newContent := ""
	for _, line := range eofSlice {
		if newContent != "" {
			newContent = fmt.Sprintf("%s%s", newContent, "\n")
		}
		if strings.Contains(line, flag) {
			newContent = fmt.Sprintf("%s%s", newContent, re.ReplaceAll([]byte(line), []byte(selectedAddress)))
		} else {
			newContent = fmt.Sprintf("%s%s", newContent, line)
		}
	}
	newContent = ConvertToOsNewLines(newContent)
	if !dryRun {
		writeFileError := fileUtils.FileOpener.WriteFile(fileName, []byte(newContent), os.ModeCharDevice)
		if writeFileError != nil {
			fmt.Printf("error while writting to file : %s %s", fileName, writeFileError)
		}
	}
	return nil
}
