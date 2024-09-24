package network

import (
	"fmt"
	"net"
	"strings"
)

type Intf struct {
	Name string
	Addr string
}

type LocalAddresses struct {
	NetNetworkInfo NetworkInfo
}

type LocalAddressesFactoy struct {
	LocalAddresses *LocalAddresses
}

func (localAddressesFactoy LocalAddressesFactoy) GetLocalAddresses() LocalAddresses {
	if localAddressesFactoy.LocalAddresses == nil {
		localAddressesFactoy.LocalAddresses = &LocalAddresses{
			NetNetworkInfo: &NetNetworkInfo{},
		}
	}

	return *localAddressesFactoy.LocalAddresses
}

func filterVirtualAndLoopbackNetworks(iface net.Interface) bool {
	if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagRunning == 0 {
		return true
	}
	// Common virtual interface name patterns
	virtualNames := []string{"veth", "docker", "virbr", "lo", "vboxnet", "vEthernet"}

	for _, name := range virtualNames {
		if strings.HasPrefix(iface.Name, name) {
			return true
		}
	}
	// Exclude VirtualBox Host-Only Adapter by exact name match
	if iface.Name == "VirtualBox Host-Only Ethernet Adapter" || strings.HasPrefix(iface.Name, "vboxnet") {
		return true
	}
	// Check if the interface's MAC address is zeroed out (often for virtual adapters)
	ifaceMAC := iface.HardwareAddr.String()
	if ifaceMAC == "" || ifaceMAC == "00:00:00:00:00:00" {
		return true
	}

	// Optionally, filter out by common VirtualBox MAC address ranges (for VirtualBox)
	// VirtualBox often uses MAC addresses starting with 08:00:27
	if strings.HasPrefix(strings.ToLower(ifaceMAC), "0a:00:27") {
		return true
	}

	return false
}

func ToIpAdresses(networkInterfaces []Intf) []string {
	interfaceAdresses := []string{}
	for _, i := range networkInterfaces {
		interfaceAdresses = append(interfaceAdresses, i.Addr)
	}

	return interfaceAdresses
}

func (localAddresses LocalAddresses) GetNetworkInterfaces() ([]Intf, error) {
	ifaces, err := localAddresses.NetNetworkInfo.Interfaces()
	if err != nil {
		return nil, err
	}
	interfaces := []Intf{}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(fmt.Errorf("localAddresses: %+v", err.Error()))
			continue
		}
		if filterVirtualAndLoopbackNetworks(i) {
			continue
		}
		fmt.Printf("  Flags: %s\n", i.Flags)
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				{
					fmt.Println(i.Name, i.HardwareAddr.String())
					if v.IP.DefaultMask() != nil {
						intf := Intf{
							Name: i.Name,
							Addr: cleanupIp(v.String()),
						}
						interfaces = append(interfaces, intf)
					}
				}
			}
		}
	}
	return interfaces, nil
}

func cleanupIp(ip string) string {
	index := strings.Index(ip, "/")
	if index != -1 {
		ip = ip[:index]
	}

	return ip
}
