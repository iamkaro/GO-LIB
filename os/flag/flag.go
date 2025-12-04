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

package flag

import (
	"flag"
	"io"
	"os"
)

var (
	flags = create()
)

func create() (out *flag.FlagSet) {
	out = flag.NewFlagSet("", flag.ContinueOnError)
	out.SetOutput(io.Discard)
	return
}

func Reset()                       { flags = create() }
func Parse() bool                  { return ParseFrom(os.Args[1:]) }
func ParseFrom(from []string) bool { return flags.Parse(from) == nil }

func Print() {
	flags.SetOutput(os.Stdout)
	flags.PrintDefaults()
	flags.SetOutput(io.Discard)
}

func Get[T bool | int | uint | int64 | uint64 | string](name string, ifNull T, usage string) T {
	var null = any(ifNull)
	switch null.(type) {
	case bool:
		return *any(flags.Bool(name, null.(bool), usage)).(*T)
	case int:
		return *any(flags.Int(name, null.(int), usage)).(*T)
	case uint:
		return *any(flags.Uint(name, null.(uint), usage)).(*T)
	case int64:
		return *any(flags.Int64(name, null.(int64), usage)).(*T)
	case uint64:
		return *any(flags.Uint64(name, null.(uint64), usage)).(*T)
	case string:
		return *any(flags.String(name, null.(string), usage)).(*T)
	}
	return ifNull
}
