package main

import (
	"errors"
	"fmt"
	"net"
)

// externalIP returns the host's external network address.
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if skipNetInterface(iface) {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ip := addrToIP(addr)
			if ip == nil || ip.IsLoopback() {
				continue
			}

			ip = ip.To4()
			if ip == nil {
				continue
			}

			return ip.String(), nil
		}
	}

	return "", errors.New("make sure you are connected to the network")
}

// skipNetInterface skip the interface of down and loopback interface.
func skipNetInterface(iface net.Interface) bool {
	if iface.Flags&net.FlagUp == 0 {
		return true
	}

	if iface.Flags&net.FlagLoopback != 0 {
		return true
	}

	return false
}

func addrToIP(addr net.Addr) net.IP {
	var ip net.IP

	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}

	return ip
}

func main() {
	ip, _ := externalIP()

	fmt.Println(ip)
}
