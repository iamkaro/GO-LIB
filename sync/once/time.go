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

package once

import "time"

func timeNow() time.Duration {
	return time.Duration(time.Now().UnixNano())
}

type (
	Millisecond time.Duration
)

func (it Millisecond) Duration() time.Duration {
	return time.Duration(it) * time.Millisecond
}
