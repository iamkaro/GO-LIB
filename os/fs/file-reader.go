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
	"bufio"
	"io"
	"os"

	"github.com/iamkaro/GO-LIB/logger"
)

func FileReader(path Path) *Reader {
	var file, err = os.Open(path.String())
	if (file == nil) || (err != nil) {
		logger.Log("error: fs.FileReader, open file error.  (", err, ")")
		return nil
	}
	return &Reader{file: file, scanner: nil, eof: false}
}

type (
	Reader struct {
		file    *os.File
		scanner *bufio.Scanner
		eof     bool
	}
)

func (it *Reader) Name() string                        { return it.file.Name() }
func (it *Reader) Read(b []byte) (n int, err error)    { return it.check(it.file.Read(b)) }
func (it *Reader) Close() error                        { return it.file.Close() }
func (it *Reader) State() (state os.FileInfo)          { state, _ = it.file.Stat(); return }
func (it *Reader) check(n int, err error) (int, error) { it.eof = err == io.EOF; return n, err }
func (it *Reader) EOF() bool                           { return it.eof }
func (it *Reader) NotEOF() bool                        { return !it.eof }

func (it *Reader) ReadLine() string {
	if it.scanner == nil {
		it.scanner = bufio.NewScanner(it.file)
	}
	if it.scanner.Scan() {
		return it.scanner.Text()
	}
	it.eof = true
	return ""
}
