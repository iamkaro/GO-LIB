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

func (it *Time) Zone() Zone { return it.zone }
func (it *Time) AS(zone Zone) *Time {
	return zone.ParseUnixNano(it.UnixNano() - it.zone.nano() + zone.nano())
}
