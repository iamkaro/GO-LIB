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

package fs

import "os"

func IsFile(path Path) bool   { var stt = State(path); return (stt != nil) && (!stt.IsDir()) }
func IsFolder(path Path) bool { var stt = State(path); return (stt != nil) && (stt.IsDir()) }

func State(path Path) os.FileInfo {
	if Exist(path) {
		var state, _ = os.Stat(path.String())
		return state
	}
	return nil
}

func Exist(path Path) bool { return !NotExist(path) }
func NotExist(path Path) bool {
	var state, err = os.Stat(path.String())
	return os.IsNotExist(err) || (state == nil)
}
