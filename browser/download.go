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

package browser

import (
	"io"
	"net/http"
	"os"
)

func Download(url, saveto string) (ok bool) {
	var (
		resp, _ = http.Get(url)
		file, _ = os.Create(saveto)
	)
	go func() {
		_ = resp.Body.Close()
		_ = file.Close()
	}()

	if (resp == nil) || (file == nil) {
		return false
	}

	var n, _ = io.Copy(file, resp.Body)
	return n == resp.ContentLength
}
