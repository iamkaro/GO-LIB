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
#include <termios.h>
#include <unistd.h>
#include <stdio.h>

int getch() {
    struct termios oldt, newt;

    // Get current terminal attributes
    tcgetattr(STDIN_FILENO, &oldt);

    // Create new termios settings
    newt = oldt;
    newt.c_lflag &= ~(ICANON | ECHO); // Disable canonical mode and echoing

    // Apply new termios settings
    tcsetattr(STDIN_FILENO, TCSANOW, &newt);

	int ch;
    // Read a single character
    // read(STDIN_FILENO, &ch, 1);
	ch = getchar();

    // Restore original terminal attributes
    tcsetattr(STDIN_FILENO, TCSANOW, &oldt);

    return ch;
}
*/
import "C"

func GetChar() byte {
	return byte(C.getch())
}
