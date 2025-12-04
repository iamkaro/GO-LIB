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

const (
	UTC  Zone = 0    // UTC
	Iran Zone = 210  // Asia/Tehran
	USA  Zone = -300 // America/New_York
	EET  Zone = 120  // EET
)

type (
	Zone int16
)

func utc() int64                      { return time.Now().In(time.UTC).UnixNano() }
func parse(sec, nsec int64) time.Time { return time.Unix(sec, nsec).In(time.UTC) }

func (it Zone) nano() int64               { return int64(it) * 60e9 }
func (it Zone) TimeNow() *Time            { return it.ParseUnixNano(utc() + it.nano()) }
func (it Zone) ParseUnix(sec int64) *Time { return it.ParseUnixNano(sec * 1e9) }
func (it Zone) ParseUnixNano(nsec int64) *Time {
	return &Time{
		time: parse(nsec/1e9, nsec%1e9),
		zone: it,
	}
}
