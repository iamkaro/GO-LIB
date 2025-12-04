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

func SubString(text string, start, length int) string {
	if (start < 0) || (length <= 0) || (start >= len(text)) {
		return ""
	}
	var end = start + length
	if (end) < len(text) {
		return string([]byte(text)[start:end])
	}
	return string([]byte(text)[start:])
}
