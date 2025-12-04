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

func Max[T Number]() MAX[T] {
	switch any(T(0)).(type) {
	case int, int64, int32, int16, int8:
		return &maxInt[T]{max: math.MinInt64, v: 0}
	case uint, uint64, uint32, uint16, uint8:
		return &maxUInt[T]{max: 0, v: 0}
	case float64, float32:
		return &maxFloat[T]{max: -(math.MaxFloat64), v: 0}
	}
	return nil
}

type (
	MAX[T Number] interface {
		Value(values ...T) MAX[T]
		Out() T
	}
	maxInt[T Number]   struct{ max, v int64 }
	maxUInt[T Number]  struct{ max, v uint64 }
	maxFloat[T Number] struct{ max, v float64 }
)

func (it *maxInt[T]) Out() T { return T(it.max) }
func (it *maxInt[T]) Value(values ...T) MAX[T] {
	if len(values) > 0 {
		for _, value := range values {
			if it.v = int64(value); it.max < it.v {
				it.max = it.v
			}
		}
	}
	return it
}

func (it *maxUInt[T]) Out() T { return T(it.max) }
func (it *maxUInt[T]) Value(values ...T) MAX[T] {
	if len(values) > 0 {
		for _, value := range values {
			if it.v = uint64(value); it.max < it.v {
				it.max = it.v
			}
		}
	}
	return it
}

func (it *maxFloat[T]) Out() T { return T(it.max) }
func (it *maxFloat[T]) Value(values ...T) MAX[T] {
	if len(values) > 0 {
		for _, value := range values {
			if it.v = float64(value); it.max < it.v {
				it.max = it.v
			}
		}
	}
	return it
}
