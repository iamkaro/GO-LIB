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

package browser

import (
	"sync"
	"time"
)

var (
	maxRequests = 7
	requests    = parallel{
		Mutex: sync.Mutex{},
		Count: 0,
	}
)

func MaxRequests(count int) {
	maxRequests = count
}

type (
	parallel struct {
		sync.Mutex
		Count int
	}
)

func (it *parallel) OK() {
	for it.Count >= maxRequests {
		time.Sleep(100 * time.Millisecond)
	}
}

func (it *parallel) Increase() {
	it.Lock()
	it.Count += 1
	it.Unlock()
}

func (it *parallel) Decrease() {
	it.Lock()
	it.Count -= 1
	it.Unlock()
}
