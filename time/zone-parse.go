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
	"fmt"
	"time"
)

func (it Zone) Parse(year, month, day, hour, minute, second int) *Time {
	var (
		t, err = time.ParseInLocation("2006-01-02 15:04:05",
			fmt.Sprintf("%d-%s-%s %s:%s:%s",
				year, s(month), s(day), s(hour), s(minute), s(second)),
			time.UTC)
	)
	if err == nil {
		return &Time{time: t, zone: it}
	}
	fmt.Println("error,time.Parse:", err)
	return nil
}

func s(i int) string {
	if i < 10 {
		return fmt.Sprintf("0%d", i)
	}
	return fmt.Sprintf("%d", i)
}
