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

package tg

import "github.com/iamkaro/GO-LIB/logger"

type (
	Data  map[string]string
	Files map[string]*struct{ Name, Path string }
)

func hasError(err error) bool {
	return logger.Error(err)
}
