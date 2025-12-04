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

package self

import "sync"

func New() Self {
	return Self{
		closed: false,
		m:      sync.Mutex{},
	}
}

type (
	Self struct {
		closed bool
		m      sync.Mutex
	}
)

func (it *Self) NotClosed() bool { return !(it.Closed()) }
func (it *Self) Closed() bool {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	return it.closed
}

func (it *Self) Close() {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	it.closed = true
}
