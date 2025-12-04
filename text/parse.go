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

package text

import (
	"github.com/iamkaro/GO-LIB/number"
	"strconv"
)

func ParseBool(text string) bool {
	var b, _ = strconv.ParseBool(text)
	return b
}

func Parse[T number.Number](text string, failure T) T {
	switch any(T(0)).(type) {

	case int64, int:
		return T(out(strconv.ParseInt(text, 10, 64))(int64(failure)))
	case int32:
		return T(out(strconv.ParseInt(text, 10, 32))(int64(failure)))
	case int16:
		return T(out(strconv.ParseInt(text, 10, 16))(int64(failure)))
	case int8:
		return T(out(strconv.ParseInt(text, 10, 8))(int64(failure)))

	case uint64, uint:
		return T(out(strconv.ParseUint(text, 10, 64))(uint64(failure)))
	case uint32:
		return T(out(strconv.ParseUint(text, 10, 32))(uint64(failure)))
	case uint16:
		return T(out(strconv.ParseUint(text, 10, 16))(uint64(failure)))
	case uint8:
		return T(out(strconv.ParseUint(text, 10, 8))(uint64(failure)))

	case float64:
		return T(out(strconv.ParseFloat(text, 64))(float64(failure)))
	case float32:
		return T(out(strconv.ParseFloat(text, 32))(float64(failure)))
	}
	return failure
}

func ParseComplex[T number.Complex](text string, failure T) T {
	switch any(T(0)).(type) {
	case complex64:
		return T(out(strconv.ParseComplex(text, 64))(complex128(failure)))
	case complex128:
		return T(out(strconv.ParseComplex(text, 128))(complex128(failure)))
	}
	return failure
}

func out[T any](value T, err error) func(failure T) T {
	return func(failure T) T {
		if err == nil {
			return value
		}
		return failure
	}
}
