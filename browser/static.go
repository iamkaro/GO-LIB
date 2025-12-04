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
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

func Get(url string, headers Headers) (ok bool, header http.Header, response string) {
	var req = newGet(url, headers)
	if (req) == nil {
		return false, nil, ""
	}
	return req.do()
}

func Post(url string, headers Headers, fields Values, fieldFiles Files) (
	ok bool, header http.Header, response string) {
	var req, f = newPost(url, headers)
	if (req == nil) ||
		(f.fields(fields)) ||
		(f.files(fieldFiles)) ||
		(f.writer.Close() == nil) {
		return false, nil, ""
	}
	return req.do()
}

func newGet(url string, headers Headers) *request {
	var get, _ = http.NewRequest("GET", url, nil)
	if (get) == nil {
		return nil
	}
	if headers != nil {
		for name, value := range headers {
			get.Header.Add(name, value)
		}
	}
	return &request{request: get}
}

func newPost(url string, headers Headers) (*request, *form) {
	var (
		body    = &bytes.Buffer{}
		writer  = multipart.NewWriter(body)
		post, _ = http.NewRequest("POST", url, body)
	)
	if (post) == nil {
		return nil, nil
	}
	if headers != nil {
		for name, value := range headers {
			post.Header.Add(name, value)
		}
	}
	post.Header.Add("Content-Type", writer.FormDataContentType())
	return &request{request: post}, &form{writer: writer}
}

var (
	client = http.Client{Timeout: 7 * time.Minute}
)

type (
	request struct{ request *http.Request }
	form    struct{ writer *multipart.Writer }
	Values  map[string]string
	Headers map[string]string
	File    struct {
		Name   string
		Reader io.Reader
	}
	Files map[string]*File
)

func (it *request) do() (ok bool, header http.Header, response string) {
	var (
		resp, err = client.Do(it.request)
		respB     []byte
	)
	if (resp == nil) || (err != nil) {
		return false, nil, ""
	}
	respB, err = ioutil.ReadAll(resp.Body)
	return err == nil, resp.Header, string(respB)
}

func (it *form) fields(fields Values) (ok bool) {
	if fields == nil {
		return true
	}
	for name, value := range fields {
		if it.writer.WriteField(name, value) != nil {
			return false
		}
	}
	return true
}
func (it *form) files(files Files) (ok bool) {
	if files == nil {
		return true
	}
	var (
		writer io.Writer
		err    error
	)
	for name, file := range files {
		writer, err = it.writer.CreateFormFile(name, file.Name)
		if err != nil {
			return false
		}
		_, err = io.Copy(writer, file.Reader)
		if err != nil {
			return false
		}
	}
	return true
}
