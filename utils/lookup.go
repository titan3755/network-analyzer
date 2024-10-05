package utils

import (
	"net"
)

func LookupHostIP(host string) ([]string, error) {
	ips, err := net.LookupHost(host)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

func LookupAddrHost(ip string) ([]string, error) {
	hosts, err := net.LookupAddr(ip)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}