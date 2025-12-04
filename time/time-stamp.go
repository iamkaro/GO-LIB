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

func (it *Time) DateOnly() string       { return it.time.Format("2006-01-02") }
func (it *Time) TimeOnly() string       { return it.time.Format("15:04:05") }
func (it *Time) Timestamp() string      { return it.time.Format("2006-01-02 15:04:05") }
func (it *Time) TimestampMilli() string { return it.time.Format("2006-01-02 15:04:05.000") }
