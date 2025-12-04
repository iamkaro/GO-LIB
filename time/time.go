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
	"time"
)

type (
	Time struct {
		zone Zone
		time time.Time
	}
)

func (it *Time) Year() int                                   { return it.time.Year() }
func (it *Time) Month() time.Month                           { return it.time.Month() }
func (it *Time) Day() int                                    { return it.time.Day() }
func (it *Time) Date() (year int, month time.Month, day int) { return it.time.Date() }

func (it *Time) Hour() int                   { return it.time.Hour() }
func (it *Time) Minute() int                 { return it.time.Minute() }
func (it *Time) Second() int                 { return it.time.Second() }
func (it *Time) Clock() (hour, min, sec int) { return it.time.Clock() }

func (it *Time) Nanosecond() int       { return it.time.Nanosecond() }
func (it *Time) Weekday() time.Weekday { return it.time.Weekday() }
func (it *Time) YearDay() int          { return it.time.YearDay() }

func (it *Time) String() string { return it.Timestamp() }
