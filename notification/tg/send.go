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

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func Send(method string, data Data) (status string, res string) {
	if !ready() {
		hasError(http.ErrNoLocation)
		return
	}
	var values = url.Values{}
	for k, v := range data {

		values[k] = []string{v}
	}
	resp, err := http.PostForm(api(method), values)
	if hasError(err) {
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if hasError(err) {
		return
	}
	return resp.Status, string(content)
}
