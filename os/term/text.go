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

package term

const (
	black  color = '0'
	red    color = '1'
	green  color = '2'
	brown  color = '3'
	blue   color = '4'
	purple color = '5'
	cyan   color = '6'
	gray   color = '7'
)

var reset = m(false, 0, 0)

func get(light bool, text string, brush, background color) string {
	return m(light, brush, background) + text + reset
}

func m(light bool, text, background color) string {
	var e = []byte{Escape, '[', value[byte](light, '1', '0')}
	if text > 0 {
		e = append(e, ';', '3', byte(text))
	}
	if background > 0 {
		e = append(e, ';', '4', byte(background))
	}
	return string(e) + "m"
}

type (
	color byte
)
