//go:build windows
// +build windows

package network

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func (o *NetUtils) IsWiFiInterface(networkName string) bool {
	cmd := exec.Command("netsh", "wlan", "show", "interfaces")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing netsh command:", err)
		return false
	}

	// Check if the interface appears in the netsh wlan show interfaces output
	return strings.Contains(out.String(), networkName)
}
