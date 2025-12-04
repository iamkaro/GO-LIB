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

import (
	"bufio"
	"fmt"
	"os"
)

func GetWord(name string) string {
	fmt.Print(name)
	var code string
	_, _ = fmt.Scanln(&code)
	return code
}

func GetLine(name string) string {
	fmt.Print(name)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
