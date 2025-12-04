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

func Delete(filename Path) bool {
	return NotExist(filename) || (os.Remove(filename.String()) == nil)
}

func DeleteFull(filename Path) bool {
	return NotExist(filename) || (os.RemoveAll(filename.String()) == nil)
}

func Move(from, to Path) bool {
	return Exist(from) && NotExist(to) &&
		(os.Rename(from.String(), to.String()) == nil)
}
