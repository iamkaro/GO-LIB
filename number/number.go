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

type (
	Number interface {
		Int | UInt | Float
	}
	Integer interface {
		Int | UInt
	}
	Int interface {
		int | int64 | int32 | int16 | int8
	}
	UInt interface {
		uint | uint64 | uint32 | uint16 | uint8
	}
	Float interface {
		float64 | float32
	}
	Complex interface {
		complex128 | complex64
	}
)
