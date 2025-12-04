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

package value

import "sync"

func New[T any](value T) *Value[T] {
	return &Value[T]{
		value: value,
		m:     sync.Mutex{},
	}
}

type (
	Value[T any] struct {
		value T
		m     sync.Mutex
	}
)

func (it *Value[T]) Get() T {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	return it.value
}

func (it *Value[T]) Set(value T) T {
	it.m.Lock()
	defer it.m.Unlock()
	/*----------*/
	it.value = value
	return it.value
}
