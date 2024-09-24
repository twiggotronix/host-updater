package network

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNetworkInterfaces(t *testing.T) {
	localAddressesFactoy := LocalAddressesFactoy{}
	localAddresses := localAddressesFactoy.GetLocalAddresses()
	mockNetworkInfo := MockNetworkInfo{
		InterfacesFunc: func() ([]net.Interface, error) {
			ifaceArray := []net.Interface{
				{
					Index:        1,
					MTU:          1500,
					Name:         "eth0",
					HardwareAddr: net.HardwareAddr{0x00, 0x14, 0x22, 0x01, 0x23, 0x45},
					Flags:        net.FlagUp | net.FlagRunning,
				},
				{
					Index:        1,
					MTU:          1500,
					Name:         "wlan0",
					HardwareAddr: net.HardwareAddr{0x00, 0x14, 0x22, 0x01, 0x23, 0x45},
					Flags:        net.FlagUp | net.FlagRunning,
				},
			}
			return ifaceArray, nil
		},
	}
	localAddresses.NetNetworkInfo = mockNetworkInfo
	result, err := localAddresses.GetNetworkInterfaces()
	assert.Nil(t, err)
	assert.Equal(t, []Intf{{Name: "eth0", Addr: "127.0.0.1"}, {Name: "wlan0", Addr: "127.0.0.1"}}, result)
}

func TestToIpAdresses(t *testing.T) {
	networkInterfaces := []Intf{
		{Name: "test", Addr: "192.168.1.5"},
		{Name: "test2", Addr: "192.168.1.6"},
	}
	result := ToIpAdresses(networkInterfaces)

	assert.Equal(t, []string{"192.168.1.5", "192.168.1.6"}, result)
}
