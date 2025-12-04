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
	"fmt"
	"math/big"
	"net"
)

const (
	// V4Family is IPv4 address-family.
	V4Family Family = iota
	// V6Family is IPv6 address-family.
	V6Family
)

// New returns new IP Range.
// If it is not a valid range (start and end IPs have different address-families, or start > end),
// New returns nil.
func New(start, end net.IP) Range {
	if isV4RangeValid(start, end) {
		return v4Range{start: start, end: end}
	}
	if isV6RangeValid(start, end) {
		return v6Range{start: start, end: end}
	}
	return nil
}

type (
	// Family represents IP Range address-family.
	Family uint8

	// Range represents an IP range.
	Range interface {
		Family() Family
		Contains(ip net.IP) bool
		Size() *big.Int
		fmt.Stringer
	}
)

