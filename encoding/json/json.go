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

package json

import (
	"encoding/json"
)

func Encode(in interface{}) string { return string(BEncode(in)) }
func BEncode(in interface{}) []byte {
	var out, _ = json.Marshal(in)
	if (out) == nil {
		return []byte{}
	}
	return out
}

func Decode(in string, out interface{}) bool  { return BDecode([]byte(in), out) }
func BDecode(in []byte, out interface{}) bool { return json.Unmarshal(in, out) == nil }
