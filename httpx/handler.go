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

package httpx

import "net/http"

type (
	Request  struct{ http.Request }
	Response struct{ http.ResponseWriter }
	Handler  func(w *Response, r *Request)
)

func (it *Response) WriteString(in string)            { _, _ = it.Write([]byte(in)) }
func (it *Response) EWriteString(in string) (e error) { _, e = it.Write([]byte(in)); return }
