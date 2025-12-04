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

func (it *Time) Equal(t *Time) bool  { return it.Compare(t) == 0 }
func (it *Time) After(t *Time) bool  { return it.Compare(t) > 0 }
func (it *Time) Before(t *Time) bool { return it.Compare(t) < 0 }
func (it *Time) Compare(by *Time) int8 {
	var (
		t1 = it.UnixNano() - it.zone.nano()
		t2 = by.UnixNano() - by.zone.nano()
	)
	if t1 < t2 {
		return -1
	}
	if t1 > t2 {
		return +1
	}
	return 0
}
