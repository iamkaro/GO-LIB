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

package base64

import "encoding/base64"

func Encode(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func Decode(in string) string {
	var txt, _ = base64.StdEncoding.DecodeString(in)
	return string(txt)
}

