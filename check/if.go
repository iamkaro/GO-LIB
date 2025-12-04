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

package check

func IF(condition bool, then func()) IfElse {
	return ifElse(false).IF(condition, then)
}

type (
	ifElse bool
	IfElse interface {
		IF(condition bool, then func()) IfElse
		Else(then func())
	}
)

func (it ifElse) IF(condition bool, then func()) IfElse {
	if !it {
		if condition {
			then()
		}
	}
	return it || ifElse(condition)
}

func (it ifElse) Else(then func()) {
	if !it {
		then()
	}
}
