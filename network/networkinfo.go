package network

import (
	"net"
)

type NetworkInfo interface {
	Interfaces() ([]net.Interface, error)
}

type NetNetworkInfo struct{}

func (o NetNetworkInfo) Interfaces() ([]net.Interface, error) {
	return net.Interfaces()
}
