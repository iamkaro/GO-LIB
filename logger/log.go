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
	"io"

	"github.com/iamkaro/GO-LIB/text"
)

var (
	Handle     func(log string) = nil
	Output     io.Writer        = nil
	PrintToSTD                  = true
)

func Panicf(format string, v ...any) { Panic(fmt.Sprintf(format, v...)) }
func Panic(v ...any)                 { log(v, func(log string) { panic(log) }) }

func Logf(format string, v ...any) { Log(fmt.Sprintf(format, v...)) }
func Log(v ...any) {
	log(v, func(log string) {
		if PrintToSTD {
			fmt.Println(log)
		}
	})
}

func log(values []any, log func(log string)) {
	var out = text.ToString(" ", values...)
	/*----------*/
	if Output != nil {
		_, _ = Output.Write([]byte(out))
	}
	/*----------*/
	if Handle != nil {
		Handle(out)
	}
	/*----------*/
	logToFile(out)
	log(out)
}
