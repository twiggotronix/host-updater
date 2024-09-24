package network

import (
	"errors"
	"net"
)

type MockNetworkInfo struct {
	InterfacesFunc func() ([]net.Interface, error)
}

func (m MockNetworkInfo) Interfaces() ([]net.Interface, error) {
	if m.InterfacesFunc != nil {
		return m.InterfacesFunc()
	}
	return nil, errors.New("not implemented")
}
