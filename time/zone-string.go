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

package time

import "fmt"

const (
	minute int16 = 1
	hour         = minute * 60
)

func (it Zone) FromUTCInMinutes() int { return int(it) }
func (it Zone) String() string {
	var (
		def      = int16(it)
		negative = def < 0
	)
	if negative {
		def = -def
	}
	var h = (def) / hour
	var m = (def % hour) / minute

	if negative {
		return fmt.Sprintf("UTC-(%d:%d)", h, m)
	}
	return fmt.Sprintf("UTC+(%d:%d)", h, m)
}
