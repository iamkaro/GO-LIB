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

package storage

import "sync"

func NewBytes(length int) *Bytes {
	return &Bytes{
		storage: make([]byte, length),
		index:   0,
		m:       sync.Mutex{},
	}
}

type (
	Bytes struct {
		storage []byte
		index   int
		m       sync.Mutex
	}
)

func (it *Bytes) IsFull() bool { return !(it.IsNotFull()) }
func (it *Bytes) IsNotFull() bool {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	return it.index < len(it.storage)
}

func (it *Bytes) SaveByte(b ...byte) { it.Save(b) }
func (it *Bytes) Save(data []byte) {
	if it.IsNotFull() {
		it.m.Lock()
		defer it.m.Unlock()
		/*----------*/
		it.index += copy(it.storage[it.index:], data)
	}
}

func (it *Bytes) Count() int {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	return it.index
}

func (it *Bytes) Get() []byte {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	return it.storage[0:it.index]
}

func (it *Bytes) GetAndReset() []byte {
	it.m.Lock()
	defer func() {
		it.index = 0
		it.m.Unlock()
	}()
	/*----------*/
	return it.storage[0:it.index]
}
