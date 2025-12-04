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
	"github.com/iamkaro/GO-LIB/text"
	"os"
)

func NewPath(name ...any) Path {
	return Path(text.ToString("/", name...))
}

type (
	Path   string
	Walker func(path Path, info os.FileInfo)
)

func (it Path) String() string     { return string(it) }
func (it Path) IsFile() bool       { return IsFile(it) }
func (it Path) IsFolder() bool     { return IsFolder(it) }
func (it Path) State() os.FileInfo { return State(it) }
func (it Path) Exist() bool        { return Exist(it) }
func (it Path) NotExist() bool     { return NotExist(it) }

func (it Path) FileWriter() *Writer                          { return FileWriter(it) }
func (it Path) FileReader() *Reader                          { return FileReader(it) }
func (it Path) Move(to Path) bool                            { return Move(it, to) }
func (it Path) Delete() bool                                 { return Delete(it) }
func (it Path) DeleteFull() bool                             { return DeleteFull(it) }
func (it Path) CreateFolder(perm os.FileMode, all bool) bool { return CreateFolder(it, perm, all) }
func (it Path) Glob(pattern string) Paths                    { return Glob(it.Sub(pattern)) }
func (it Path) Scan(handle Walker) bool                      { return Scan(it, handle) }

func (it Path) Sub(name ...any) Path {
	return Path(text.ToString("/", it, text.ToString("/", name...)))
}
