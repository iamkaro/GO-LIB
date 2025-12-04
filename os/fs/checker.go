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
	"io"

	"github.com/iamkaro/GO-LIB/logger"
)

func Check(n int, err error) *Checker {
	return &Checker{n: n, err: err}
}

type (
	Checker struct {
		n   int
		err error
	}
)

func (it *Checker) Length() int          { return it.n }
func (it *Checker) OK() bool             { return !it.Error() }
func (it *Checker) EOF() bool            { return it.err == io.EOF }
func (it *Checker) Full(length int) bool { return it.OK() && (length == it.n) }
func (it *Checker) Error() bool {
	if it.EOF() && (it.n > 0) {
		return false
	}
	return logger.Error(it.err)
}
