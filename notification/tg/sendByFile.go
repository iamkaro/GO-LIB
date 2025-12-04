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
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func SendByFile(method string, data Data, files Files) (status string, res string) {
	if !ready() {
		hasError(http.ErrNoLocation)
		return
	}
	var (
		b   bytes.Buffer
		fw  io.Writer
		fr  *os.File
		err error
	)

	var w = multipart.NewWriter(&b)
	for field, value := range data {
		if hasError(w.WriteField(field, value)) {
			return
		}
	}
	for field, file := range files {
		fw, err = w.CreateFormFile(field, file.Name)
		if hasError(err) {
			return
		}
		fr, err = os.Open(file.Path)
		if hasError(err) {
			return
		}
		_, err = io.Copy(fw, fr)
		if hasError(err) {
			return
		}
	}
	if hasError(w.Close()) {
		return
	}
	req, err := http.NewRequest("POST", api(method), &b)
	if hasError(err) {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if hasError(err) {
		return
	}
	content, err := ioutil.ReadAll(resp.Body)
	if hasError(err) {
		return
	}
	return resp.Status, string(content)
}
