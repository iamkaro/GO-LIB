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
	"math/big"
	"net"
	"strings"
)

// NewPool create a new collection.
func NewPool() *Pool {
	return &Pool{
		list:  map[int]Range{},
		index: 0,
	}
}

type (
	// Pool is a collection of IP Ranges.
	Pool struct {
		list  map[int]Range
		index int
	}
)

func (p *Pool) Add(r Range) {
	p.list[p.index] = r
	p.index += 1
}
func (p *Pool) Remove(r Range) {
	for i, v := range p.list {
		if r == v {
			delete(p.list, i)
			return
		}
	}
}

// String returns the string form of the pool.
func (p *Pool) String() string {
	var b strings.Builder
	for _, r := range p.list {
		b.WriteString(r.String() + " ")
	}
	return strings.TrimSpace(b.String())
}

// Size reports the number of IP addresses in the pool.
func (p *Pool) Size() *big.Int {
	size := big.NewInt(0)
	for _, r := range p.list {
		size.Add(size, r.Size())
	}
	return size
}

// Contains reports whether the pool includes IP.
func (p *Pool) Contains(ip net.IP) bool {

	if ip == nil {
		return false
	}

	for _, r := range p.list {
		if r.Contains(ip) {
			return true
		}
	}

	return false
}

