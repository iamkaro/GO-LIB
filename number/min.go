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

package number

import "math"

func Min[T Number]() MIN[T] {
	switch any(T(0)).(type) {
	case int, int64, int32, int16, int8:
		return &minInt[T]{min: math.MaxInt64, v: 0}
	case uint, uint64, uint32, uint16, uint8:
		return &minUInt[T]{min: math.MaxUint64, v: 0}
	case float64, float32:
		return &minFloat[T]{min: math.MaxFloat64, v: 0}
	}
	return nil
}

type (
	MIN[T Number] interface {
		Value(values ...T) MIN[T]
		Out() T
	}
	minInt[T Number]   struct{ min, v int64 }
	minUInt[T Number]  struct{ min, v uint64 }
	minFloat[T Number] struct{ min, v float64 }
)

func (it *minInt[T]) Out() T { return T(it.min) }
func (it *minInt[T]) Value(values ...T) MIN[T] {
	if len(values) > 0 {
		for _, value := range values {
			if it.v = int64(value); it.min > it.v {
				it.min = it.v
			}
		}
	}
	return it
}

func (it *minUInt[T]) Out() T { return T(it.min) }
func (it *minUInt[T]) Value(values ...T) MIN[T] {
	if len(values) > 0 {
		for _, value := range values {
			if it.v = uint64(value); it.min > it.v {
				it.min = it.v
			}
		}
	}
	return it
}

func (it *minFloat[T]) Out() T { return T(it.min) }
func (it *minFloat[T]) Value(values ...T) MIN[T] {
	if len(values) > 0 {
		for _, value := range values {
			if it.v = float64(value); it.min > it.v {
				it.min = it.v
			}
		}
	}
	return it
}
