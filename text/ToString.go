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
	"fmt"
)

func StringOf(values ...any) string { return ToString("", values...) }
func ToString(join string, values ...any) string {
	var value = ""
	for i, v := range values {
		if i > 0 {
			value += join
		}
		value += fmt.Sprint(v)
	}
	return value
}
