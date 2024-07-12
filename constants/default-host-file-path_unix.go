//go:build !windows
// +build !windows

package constants

import (
	"path/filepath"
)

var DefaultHostFilePath = filepath.Join("etc")
