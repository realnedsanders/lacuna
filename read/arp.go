/*
 * @File: arp.go
 * @Date: 2019-05-31 03:02:03
 * @OA:   antonioe
 * @CA:   Antonio Escalera
 * @Time: 2019-06-07 12:53:49
 * @Mail: antonioe@wolfram.com
 * @Copy: Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>
 */

package read

import (
	"bufio"
	"net"
	"os"
	"strings"
)

type ArpData struct {
	Ipaddr  net.IP
	Macaddr string
	Iface   string
}

func ReadArpDataIntoSlice(path string, debug bool) ([]net.IP, error) {
	n := []net.IP{}
	arp, err := ReadArpDataIntoStruct(path)
	if err != nil {
		return n, err
	}
	for _, v := range arp {
		n = append(n, v.Ipaddr)
	}
	r := unique(n)
	return r, nil
}

func ReadArpDataIntoStruct(path string) ([]ArpData, error) {
	n := []ArpData{}
	recs, err := readLines(path)
	if err != nil {
		return n, err
	}
	a := make([]ArpData, len(recs)-1)
	for idx, val := range recs {
		if idx == 0 {
		} else {
			vf := strings.Fields(val)
			if len(vf) == 0 {
				break
			}
			a[idx-1].Macaddr = vf[3]
			a[idx-1].Ipaddr = net.ParseIP(vf[0])
			a[idx-1].Iface = vf[5]
		}
	}
	return a, nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
