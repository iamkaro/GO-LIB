/*! 
 * I am Karo  😊👍 
 * 
 * Contact me:  
 *     https://www.karo.link/ 
 *     https://github.com/iamkaro
 *     https://www.linkedin.com/in/iamkaro  
 * 
 * My own GO-based library 
 * https://github.com/iamkaro/GO-LIB.git
 * Copyright © 2021 developed. 
 */

package iprange

import (
	"bytes"
	"fmt"
	"math/big"
	"net"
)

type v4Range struct {
	start net.IP
	end   net.IP
}

// String returns the string form of the range.
func (r v4Range) String() string {
	return fmt.Sprintf("%s-%s", r.start, r.end)
}

// Family returns the range address family.
func (r v4Range) Family() Family { return V4Family }

// Contains reports whether the range includes IP.
func (r v4Range) Contains(ip net.IP) bool {
	return bytes.Compare(ip, r.start) >= 0 && bytes.Compare(ip, r.end) <= 0
}

// Size reports the number of IP addresses in the range.
func (r v4Range) Size() *big.Int {
	return big.NewInt(v4ToInt(r.end) - v4ToInt(r.start) + 1)
}

func v4ToInt(ip net.IP) int64 {
	ip = ip.To4()
	return (int64(ip[0]) << 24) | (int64(ip[1]) << 16) | (int64(ip[2]) << 8) | (int64(ip[3]))
}

