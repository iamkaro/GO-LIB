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

var (
	botURL = ""
)

func ready() bool              { return botURL != "" }
func api(method string) string { return botURL + method }

func SetBot(token string) {
	botURL = "https://api.telegram.org/bot" + token + "/"
}
