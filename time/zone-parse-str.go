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

import (
	"github.com/iamkaro/GO-LIB/text"
	"strings"
)

func (it Zone) ParseString(datetime string) *Time {
	var (
		Year, Month, Day, Hour, Minute, Second = 1, 1, 1, 0, 0, 0
		vs                                     []string
		date, time                             = "", ""
	)
	datetime = strings.TrimSpace(datetime)
	/*----------*/
	vs = strings.Split(datetime, " ")
	if (len(vs) != 1) && (len(vs) != 2) {
		panic("error, Zone.ParseString(): datetime format error!")
		return nil
	}
	if date = vs[0]; len(vs) == 2 {
		time = vs[1]
	}
	/*----------*/
	vs = strings.Split(date, "-")
	if len(vs) != 3 {
		panic("error, Zone.ParseString(): date format error!")
		return nil
	}
	Year = text.Parse(vs[0], 0)
	if (Year < 1000) || (3000 < Year) {
		panic("error, Zone.ParseString(): date.Year error!")
		return nil
	}
	Month = text.Parse(vs[1], 0)
	if (Month < 1) || (12 < Month) {
		panic("error, Zone.ParseString(): date.Month error!")
		return nil
	}
	Day = text.Parse(vs[2], 0)
	if (Day < 1) || (31 < Day) {
		panic("error, Zone.ParseString(): date.Day error!")
		return nil
	}
	/*----------*/
	if time != "" {
		vs = strings.Split(time, ":")
		if (len(vs) != 2) && (len(vs) != 3) {
			panic("error, Zone.ParseString(): time format error!")
			return nil
		}
		Hour = text.Parse(vs[0], 0)
		if (Hour < 0) || (23 < Hour) {
			panic("error, Zone.ParseString(): time.Hour error!")
			return nil
		}
		Minute = text.Parse(vs[1], 0)
		if (Minute < 0) || (59 < Minute) {
			panic("error, Zone.ParseString(): time.Minute error!")
			return nil
		}
		if len(vs) == 3 {
			Second = text.Parse(vs[2], 0)
			if (Second < 0) || (59 < Second) {
				panic("error, Zone.ParseString(): time.Second error!")
				return nil
			}
		}
	}
	/*----------*/

	return it.Parse(Year, Month, Day, Hour, Minute, Second)
}
