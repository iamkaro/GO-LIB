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
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

const (
	agent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
)

func New(timeout time.Duration) *Client {
	var jar, _ = cookiejar.New(nil)
	return &Client{
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           jar,
			Timeout:       timeout,
		},
	}
}
func NewWithoutCookies(timeout time.Duration) *Client {
	return &Client{
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       timeout,
		},
	}
}

type (
	Client struct {
		client *http.Client
	}
	Cookies []*http.Cookie
)

func (it *Client) CheckRedirect(handle func(req *http.Request, via []*http.Request) error) {
	it.client.CheckRedirect = handle
}

func (it *Client) Cookie(site string) Cookies {
	var URL, _ = url.Parse(site)
	return it.client.Jar.Cookies(URL)
}
func (it *Client) SetCookie(site string, cookies Cookies) {
	var URL, _ = url.Parse(site)
	it.client.Jar.SetCookies(URL, cookies)
}

func (it *Client) request(method, url, contentType string, body io.Reader) (string, *http.Response) {

	requests.OK()
	requests.Increase()
	defer requests.Decrease()

	var req, _ = http.NewRequest(method, url, body)
	req.Header.Set("User-Agent", agent)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	var res, _ = it.client.Do(req)
	if (res) == nil {
		return "", nil
	}

	var data, _ = ioutil.ReadAll(res.Body)
	return string(data), res
}

func (it *Client) Get(url string) (string, *http.Response) {
	return it.request(http.MethodGet, url, "", nil)
}

func (it *Client) PostForm(url string, data url.Values) (string, *http.Response) {
	return it.request(http.MethodPost, url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()))
}

func (it *Client) Post(url string, contentType string, body io.Reader) (string, *http.Response) {
	return it.request(http.MethodPost, url, contentType, body)
}
