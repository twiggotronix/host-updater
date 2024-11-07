package network

type NetworkUtils interface {
	IsWiFiInterface(networkName string) bool
}

type NetUtils struct {
}
