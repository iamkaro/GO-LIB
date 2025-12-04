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

package term

/*
#include <stdio.h>
#include <signal.h>

void sigintHandler(int sig_num) {
	signal(SIGINT, sigintHandler);
	fflush(stdout);
}
void disableCtrlC() {
	signal(SIGINT, sigintHandler);
}
*/
import "C"

func DisableCtrlC() {
	C.disableCtrlC()
}
