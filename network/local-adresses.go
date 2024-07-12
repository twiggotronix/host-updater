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

func validateNetworkName(name string) bool {
	return !strings.HasPrefix(name, "Loopback") && !strings.HasPrefix(name, "vEthernet")
}

func ToIpAdresses(networkInterfaces []Intf) []string {
	interfaceAdresses := []string{}
	for _, i := range networkInterfaces {
		interfaceAdresses = append(interfaceAdresses, i.Addr)
	}

	return interfaceAdresses
}
func GetNetworkInterfaces() ([]Intf, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	interfaces := []Intf{}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println(fmt.Errorf("localAddresses: %+v", err.Error()))
			continue
		}

		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				{
					if v.IP.DefaultMask() != nil && validateNetworkName(i.Name) {
						intf := Intf{
							Name: i.Name,
							Addr: v.String(),
						}
						interfaces = append(interfaces, intf)
					}
				}
			}
		}
	}
	return interfaces, nil
}
