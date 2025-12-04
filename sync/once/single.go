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

package once

import (
	"sync"
)

func NewSingle() Single {
	return Single{
		m: sync.Mutex{},
	}
}

type (
	Single struct {
		m sync.Mutex
	}
)

func (it *Single) Run(task func()) {
	it.m.Lock()
	task()
	it.m.Unlock()
}

func (it *Single) TryRun(task func()) bool {
	if it.m.TryLock() {
		task()
		it.m.Unlock()
		return true
	}
	return false
}
