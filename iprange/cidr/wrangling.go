package cidr

import (
	"fmt"
	"math/big"
	"net"
)

func ipToInt(ip net.IP) (*big.Int, int) {
	val := &big.Int{}
	val.SetBytes(ip)
	if len(ip) == net.IPv4len {
		return val, 32
	}
	if len(ip) == net.IPv6len {
		return val, 128
	}
	fmt.Printf("Unsupported address length %d \n", len(ip))
	return val, 0
}

func intToIP(ipInt *big.Int, bits int) net.IP {
	ipBytes := ipInt.Bytes()
	ret := make([]byte, bits/8)
	for i := 1; (i <= len(ipBytes)) && (i <= len(ret)); i++ {
		ret[len(ret)-i] = ipBytes[len(ipBytes)-i]
	}
	return ret
}

