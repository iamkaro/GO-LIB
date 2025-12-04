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

// Pow   b ^ a
func Pow[AnyTypeNumber Number, IntegerOnly Integer](b AnyTypeNumber, a IntegerOnly) AnyTypeNumber {
	if a > 0 {
		return b * Pow(b, a-1)
	}
	return 1
}
