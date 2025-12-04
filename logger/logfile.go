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
	"os"

	"github.com/iamkaro/GO-LIB/time"
)

var (
	logfile *os.File = nil
	Time             = time.Iran
)

func logToFile(log string) {
	if logfile != nil {
		_, _ = logfile.WriteString(Time.TimeNow().Timestamp() + "  " + log + "\n")
	}
}

func ToFile(path string) {
	if logfile == nil {
		var file, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if (file == nil) || (err != nil) {
			Log("error: fs.FileWriter, open file error.  (", err, ")")
			return
		}
		logfile = file
	}
}

func ToFileCancel() {
	if logfile != nil {
		_ = logfile.Close()
		logfile = nil
	}
}
