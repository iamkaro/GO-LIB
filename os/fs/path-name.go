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
	"strings"

	"github.com/iamkaro/GO-LIB/text"
)

func (it Path) Name() Name {
	var list = strings.Split(it.String(), "/")
	return Name(list[len(list)-1])
}

func (it Path) Direction() Path {
	var list = strings.Split(it.String(), "/")
	return Path(text.ToString("/", list[0:len(list)-1]))
}

func (it Path) Names() Names {
	var (
		list  = strings.Split(it.String(), "/")
		names = make(Names, len(list))
	)
	for i, name := range list {
		names[i] = Name(name)
	}
	return names
}

type (
	Names []Name
	Name  string
)

func (it Name) String() string                { return string(it) }
func (it Name) Int(failure int64) int64       { return text.Parse[int64](string(it), failure) }
func (it Name) Float(failure float64) float64 { return text.Parse[float64](string(it), failure) }
func (it Name) Sub(name ...any) Path {
	return Path(text.ToString("/", it, text.ToString("/", name...)))
}
