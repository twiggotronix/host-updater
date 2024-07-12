package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func checkFileIsWrittable(fileName string) bool {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

func SetNewIp(fileName string, flag string, selectedAddress string, dryRun bool) {
	if !checkFileIsWrittable(fileName) {
		fmt.Printf("%s is not writtable\n", fileName)
		if !dryRun {
			return
		}
	}
	biteContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error while reading to file : %s %s", fileName, err)
		return
	}
	biteContent = NormalizeNewlines(biteContent)
	eofSlice := strings.Split(string(biteContent), "\n")
	re := regexp.MustCompile(`(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)\.(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)\.(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)\.(25[0–5]|2[0-5][0-9]?|1[0-9]?[0-9]?|[0-9][0-9]?)`)
	newContent := ""
	for _, line := range eofSlice {
		if strings.Contains(line, flag) {
			newContent = fmt.Sprintf("%s%s%s", newContent, "\n", re.ReplaceAll([]byte(line), []byte(selectedAddress)))
		} else {
			newContent = fmt.Sprintf("%s%s%s", newContent, "\n", line)
		}
	}
	newContent = ConvertToOsNewLines(newContent)
	if !dryRun {
		writeFileError := os.WriteFile(fileName, []byte(newContent), os.ModeCharDevice)
		if writeFileError != nil {
			fmt.Printf("error while writting to file : %s %s", fileName, writeFileError)
		}
	}
}
