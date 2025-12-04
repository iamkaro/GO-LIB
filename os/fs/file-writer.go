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

	"github.com/iamkaro/GO-LIB/logger"
)

func FileWriter(path Path) *Writer {
	var file, err = os.OpenFile(path.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if (file == nil) || (err != nil) {
		logger.Log("error: fs.FileWriter, open file error.  (", err, ")")
		return nil
	}
	return &Writer{file: file}
}

type Writer struct{ file *os.File }

func (it *Writer) Name() string                                { return it.file.Name() }
func (it *Writer) Write(b []byte) (n int, err error)           { return it.file.Write(b) }
func (it *Writer) WriteLine(line string) (n int, err error)    { return it.file.WriteString(line + "\n") }
func (it *Writer) WriteString(value string) (n int, err error) { return it.file.WriteString(value) }
func (it *Writer) Close() error                                { return it.file.Close() }
func (it *Writer) State() (state os.FileInfo)                  { state, _ = it.file.Stat(); return }
