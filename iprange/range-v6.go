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

type v6Range struct {
	start net.IP
	end   net.IP
}

// String returns the string form of the range.
func (r v6Range) String() string {
	return fmt.Sprintf("%s-%s", r.start, r.end)
}

// Family returns the range address family.
func (r v6Range) Family() Family { return V6Family }

// Contains reports whether the range includes IP.
func (r v6Range) Contains(ip net.IP) bool {
	return bytes.Compare(ip, r.start) >= 0 && bytes.Compare(ip, r.end) <= 0
}

// Size reports the number of IP addresses in the range.
func (r v6Range) Size() *big.Int {
	size := big.NewInt(0)
	size.Add(size, big.NewInt(0).SetBytes(r.end))
	size.Sub(size, big.NewInt(0).SetBytes(r.start))
	size.Add(size, big.NewInt(1))
	return size
}

