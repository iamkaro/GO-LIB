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

package ingo

import (
	"fmt"
	"time"
)

const (
	minute = 1
	hour   = minute * 60
	day    = hour * 24
)

func (it Timezone) FromUTCInMinutes() Difference {
	var (
		t   = time.Unix(day*3, 0)
		utc = t.In(UTC.location())
		its = t.In(it.location())
		def = 0
	)

	def += (its.YearDay() - utc.YearDay()) * day
	def += (its.Hour() - utc.Hour()) * hour
	def += (its.Minute() - utc.Minute()) * minute

	return Difference(def)
}

type (
	Difference int
)

func (it Difference) ToString() string {
	var (
		def      = int(it)
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
