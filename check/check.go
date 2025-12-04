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

package check

func Check[T comparable](value, empty, ifEmpty T) T {
	if value == empty {
		return ifEmpty
	}
	return value
}
