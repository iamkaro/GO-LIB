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
	"time"
)

func NewThread() *Thread {
	return &Thread{
		on:    false,
		mutex: sync.Mutex{},
		block: nil,
	}
}

type (
	Thread struct {
		on    bool
		mutex sync.Mutex
		block *block
	}
	block struct {
		next *block
		Func Func
		Time time.Duration
	}
	Func func()
)

func (it *Thread) Empty() {
	it.mutex.Lock()
	it.block = nil
	it.mutex.Unlock()
}

func (it *Thread) RunNow(task Func) {
	it.mutex.Lock()
	it.block = &block{next: it.block, Func: task, Time: 1}
	it.mutex.Unlock()
}

func (it *Thread) Run(task Func, after Millisecond) {
	it.mutex.Lock()
	defer it.mutex.Unlock()
	/*----------*/
	var (
		Timeout = after.Duration() + timeNow()
		b       = &block{next: nil, Func: task, Time: Timeout}
	)
	if it.block == nil {
		it.block = b
		return
	}
	var i, o = it.block, it.block
	for (i.next != nil) && (i.next.Time <= Timeout) {
		o = i
		i = i.next
	}
	if i.Time <= Timeout {
		b.next = i.next
		i.next = b
		return
	}
	if it.block == i {
		b.next = i
		it.block = b
		return
	}
	b.next = i
	o.next = b
}

func (it *Thread) Stop() { it.on = false }
func (it *Thread) Start() {
	var b *block = nil
	it.on = true
	for it.on {

		it.mutex.Lock()
		b = it.block
		if b != nil {
			if (b.Time - timeNow()) <= 0 {
				it.block = b.next
			} else {
				b = nil
			}
		}
		it.mutex.Unlock()

		if b != nil {
			if b.Func != nil {
				b.Func()
			}
			continue
		}

		time.Sleep(time.Millisecond * 100)
	}
}
