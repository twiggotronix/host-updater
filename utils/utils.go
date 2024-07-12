package utils

import (
	"bytes"

	"github.com/twiggotronix/host-updater/constants"
)

func NormalizeNewlines(d []byte) []byte {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}

func ConvertToOsNewLines(content string) string {
	newContent := bytes.Replace([]byte(content), []byte{10}, []byte(constants.LineBreak), -1)
	return string(newContent)
}
