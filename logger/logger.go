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

package logger

import (
	"fmt"

	"github.com/iamkaro/GO-LIB/text"
)

func NewLogger(name ...any) Logger {
	return Logger(text.StringOf(name...))
}

type (
	Logger string
)

func (it Logger) SubLogger(name ...any) Logger   { return NewLogger(it, " ", text.StringOf(name...)) }
func (it Logger) Log(v ...any)                   { Log(it, " ", text.ToString(" ", v...)) }
func (it Logger) Panic(v ...any)                 { Panic(it, " ", text.ToString(" ", v...)) }
func (it Logger) Logf(format string, v ...any)   { Log(it, " ", fmt.Sprintf(format, v...)) }
func (it Logger) Panicf(format string, v ...any) { Panic(it, " ", fmt.Sprintf(format, v...)) }
func (it Logger) NotError(err error) bool        { return !errorByPrefix(err, "error,", it, ":") }
func (it Logger) Error(err error) bool           { return errorByPrefix(err, "error,", it, ":") }
