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

func ToSolar(unixTime int64) (date *Solar) {
	date = &Solar{Year: 1348, Moon: 10, Day: 11}

	unixTime = unixTime / 86400

	for unixTime > 365 {
		if LeapYear(date.Year) {
			unixTime -= 1
		}
		date.Year += 1
		unixTime -= 365
	}

	for unixTime > 30 {
		if date.Moon < 7 {
			unixTime -= 31
		} else if date.Moon < 12 {
			unixTime -= 30
		} else {
			if LeapYear(date.Year) {
				unixTime -= 30
			} else {
				unixTime -= 29
			}
			date.Year += 1
			date.Moon = 0
		}
		date.Moon += 1
	}

	t := 29
	if LeapYear(date.Year) {
		t = 30
	}

	for unixTime > 0 {
		if date.Moon < 7 {
			if date.Day < 31 {
				date.Day += 1
			} else {
				date.Day = 1
				date.Moon += 1
			}
		} else if date.Moon < 12 {
			if date.Day < 30 {
				date.Day += 1
			} else {
				date.Day = 1
				date.Moon += 1
			}
		} else if date.Day < t {
			date.Day += 1
		} else {
			date.Day = 1
			date.Moon = 1
			date.Year += 1
		}

		unixTime -= 1
	}

	return
}
