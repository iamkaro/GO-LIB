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

import (
	"os"
	"path/filepath"
)

func CreateFolder(path Path, perm os.FileMode, all bool) bool {
	if Exist(path) {
		return IsFolder(path)
	}
	if all {
		return os.MkdirAll(path.String(), perm) == nil
	}
	return os.Mkdir(path.String(), perm) == nil
}

func Glob(pattern Path) Paths {
	var (
		list, _ = filepath.Glob(pattern.String())
		paths   = make(Paths, len(list))
	)
	for i, p := range list {
		paths[i] = Path(p)
	}
	return paths
}

func Scan(folder Path, handle Walker) bool {
	if IsFolder(folder) {
		var (
			walk = func(path string, info os.FileInfo, err error) error {
				handle(Path(path), info)
				return nil
			}
			err = filepath.Walk(folder.String(), walk)
		)
		return err == nil
	}
	return false
}
