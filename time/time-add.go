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

import "time"

func (it *Time) new(time time.Time) *Time  { return &Time{time: time, zone: it.zone} }
func (it *Time) Add(d time.Duration) *Time { return it.new(it.time.Add(d)) }
func (it *Time) AddDate(years, months, days int) *Time {
	return it.new(it.time.AddDate(years, months, days))
}

func (it *Time) NextDay() *Time   { return it.AddDate(0, 0, 1) }
func (it *Time) NextMonth() *Time { return it.AddDate(0, 1, 0) }
func (it *Time) NextYear() *Time  { return it.AddDate(1, 0, 0) }

func (it *Time) BackDay() *Time   { return it.AddDate(0, 0, -1) }
func (it *Time) BackMonth() *Time { return it.AddDate(0, -1, 0) }
func (it *Time) BackYear() *Time  { return it.AddDate(-1, 0, 0) }
