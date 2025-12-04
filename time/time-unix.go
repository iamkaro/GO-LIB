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

func (it *Time) Unix() int64      { return it.time.Unix() }
func (it *Time) UnixMilli() int64 { return it.time.UnixMilli() }
func (it *Time) UnixMicro() int64 { return it.time.UnixMicro() }
func (it *Time) UnixNano() int64  { return it.time.UnixNano() }
