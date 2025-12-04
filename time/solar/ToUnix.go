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

func ToUnix(date *Solar) (time int64) {
	time = 0
	if date == nil {
		return
	}
	for Y := 1348; Y < date.Year; Y++ {
		if LeapYear(Y) {
			time += 366
		} else {
			time += 365
		}
	}
	var M = 10
	for ; M < date.Moon; M++ {
		if M < 7 {
			time += 31
		} else if M < 12 {
			time += 30
		}
	}
	for M > date.Moon {
		M -= 1
		if M < 7 {
			time -= 31
		} else if M < 12 {
			time -= 30
		}
	}
	time += int64(date.Day) - 11
	time = time * 86400
	return
}

