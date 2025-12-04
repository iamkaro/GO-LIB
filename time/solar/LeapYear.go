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

package solar

func LeapYear(year int) (yes bool) {
	var ky = 1342
	for jump := 0; ky < year; ky += 4 {
		jump += 1
		if jump%8 == 0 {
			ky++
		}
	}
	return ky == year
}

