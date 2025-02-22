package geoip

import "net"

func Check(ip string) bool {
	o := net.ParseIP(ip)
	return theTrie.Contains(o)
}
