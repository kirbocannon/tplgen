package tplgen

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Router struct {
	Hostname string
	Serial string
	LoopbackIPAddress string
}


func CreateRouterConfiguration(router Router) Router {
	return router
}

func ConvertIpCidrToIpNetmask(ip string) string {
	ipv4Addr, _, err := net.ParseCIDR(ip)
	if err != nil {
		log.Fatal(err)
	}

	// if cidr notation, convert to netmask equivalent
	if strings.Contains(ip, "/") {
		cidr, _ := strconv.ParseUint(strings.Split(ip, "/")[1], 10, 32)
		// 0xFFFFFFFF is hexadecimal constant
		// << (left shift) is multiplication >> (right shift) is division), only works for nums 2 to the power
		mask := (0xFFFFFFFF << (32 - cidr)) & 0xFFFFFFFF
		var dmask uint64
		dmask = 32
		localmask := make([]int, 0, 4)
		for i := 1; i <= 4; i++ {
			tmp := mask >> (dmask - 8) & 0xFF
			localmask = append(localmask, tmp)
			dmask -= 8
		}
		netmaskList := make([]string, 0, 4)
		for i := 0; i <= 3; i++ {
			octet := strconv.Itoa(localmask[i])
			netmaskList = append(netmaskList, octet)
		}
		netmask := strings.Join(netmaskList, ".")
		return fmt.Sprintf("%s %s", ipv4Addr, netmask)
	}
	return ip
}

func RemoveCidr(s string) string {
	if strings.Contains(s, "/") {
		return strings.Split(s, "/")[0]
	}
	return s
}


