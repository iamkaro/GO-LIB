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

import (
	"fmt"
)

// ANSI escape codes

func Clear()            { fmt.Printf("%c[2J", Escape) }
func ClearLineToRight() { fmt.Printf("%c[K", Escape) }

func GotoX(x int)                     { fmt.Printf("%c[%dG", Escape, x) }
func GotoXY(x, y int)                 { fmt.Printf("%c[%d;%dH", Escape, y, x) }
func GotoBeginningLineUp(lines int)   { fmt.Printf("%c[%dF", Escape, lines) }
func GotoBeginningLineDown(lines int) { fmt.Printf("%c[%dE", Escape, lines) }

func MoveUP(count int)    { move('A', count) }
func MoveDown(count int)  { move('B', count) }
func MoveRight(count int) { move('C', count) }
func MoveLeft(count int)  { move('D', count) }
func move(code byte, count int) {
	if count > 1 {
		fmt.Printf("%c[%d;%c", Escape, count, code)
	}
	fmt.Printf("%c[%c", Escape, code)
}
