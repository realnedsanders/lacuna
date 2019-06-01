/*
 * @File: find.go
 * @Date: 2019-05-30 17:32:24
 * @OA:   antonioe
 * @CA:   antonioe
 * @Time: 2019-06-01 01:52:13
 * @Mail: antonioe@wolfram.com
 * @Copy: Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>
 */

package find

import (
	"net"
)

type NetworkHosts map[string][]net.IP

//type NetworkHosts struct {
//	Arp      []net.IP
//	Arpwatch []net.IP
//	Arpwitch []net.IP
//	Dns      []net.IP
//	Ping     []net.IP
//}

// returns IP addresses in subnet.
// Accepts subnet and struct of arrays of net.IP
func FreeIPs(s *net.IPNet, n NetworkHosts) ([]net.IP, error) {
	ips, err := HostsInSubnet(s)
	if err != nil {
		return []net.IP{}, err
	}
	for _, val := range n {
		ips = removeIPs(ips, val)
	}
	return ips, nil
}

// removes IPs in s from n
func removeIPs(n []net.IP, s []net.IP) []net.IP {
	for idx, val := range n {
		for _, v := range s {
			if val.Equal(v) {
				n = rem(n, idx)
			}
		}
	}
	return n
}

// removes net.IP at index i from s and returns the new slice
func rem(s []net.IP, i int) []net.IP {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil // or the zero value of T
	s = s[:len(s)-1]

	return s
	//return append(s[:i], s[i+1:]...)
}

// returns an array of hosts inside a given subnet
func HostsInSubnet(s *net.IPNet) ([]net.IP, error) {
	ip, ipnet, err := net.ParseCIDR(s.String())
	if err != nil {
		return []net.IP{}, err
	}
	var ips []net.IP
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); ip = inc(ip) {
		ips = append(ips, ip)
		//fmt.Printf("%v", ips)
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

// increments the ip address by 1
func inc(ip net.IP) net.IP {
	incIP := make([]byte, len(ip))
	copy(incIP, ip)
	for j := len(incIP) - 1; j >= 0; j-- {
		incIP[j]++
		if incIP[j] > 0 {
			break
		}
	}
	return incIP
}