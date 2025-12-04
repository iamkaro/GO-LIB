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

package logger

import (
	"github.com/iamkaro/GO-LIB/text"
)

func NotError(err error) bool { return !errorByPrefix(err, "error:") }
func Error(err error) bool    { return errorByPrefix(err, "error:") }

func errorByPrefix(err error, prefix ...any) bool {
	if err == nil {
		return false
	}
	Log(text.StringOf(prefix...), err)
	return true
}
