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

package list

import "sync"

func NewArray[T any](items ...T) *Array[T] {
	return &Array[T]{
		list: items,
		m:    sync.Mutex{},
	}
}

type (
	Array[T any] struct {
		list []T
		m    sync.Mutex
	}
)

func (it *Array[T]) Add(item T) {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	it.list = append(it.list, item)
}

func (it *Array[T]) Len() int {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	return len(it.list)
}

func (it *Array[T]) Get(index int) (item T) {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	if len(it.list) > index {
		return it.list[index]
	}
	return
}

func (it *Array[T]) Foreach(handle func(k int, v T)) {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	for k, v := range it.list {
		handle(k, v)
	}
}

func (it *Array[T]) Remove(index int) {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	if (index < 0) || (len(it.list) <= index) {
		return
	}
	for index += 1; index < len(it.list); index += 1 {
		it.list[index-1] = it.list[index]
	}
	it.list = it.list[0 : len(it.list)-1]
}
