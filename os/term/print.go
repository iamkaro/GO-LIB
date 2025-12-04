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

import "fmt"

func Print(v ...any)                 { fmt.Print(v...) }
func Printf(format string, v ...any) { fmt.Printf(format, v...) }
func Println(v ...any)               { fmt.Println(v...) }

func Light() *Style { return NewStyle().Light() }

func Black() *Style  { return NewStyle().Black() }
func Red() *Style    { return NewStyle().Red() }
func Green() *Style  { return NewStyle().Green() }
func Brown() *Style  { return NewStyle().Brown() }
func Blue() *Style   { return NewStyle().Blue() }
func Purple() *Style { return NewStyle().Purple() }
func Cyan() *Style   { return NewStyle().Cyan() }
func Gray() *Style   { return NewStyle().Gray() }

func BGBlack() *Style  { return NewStyle().BGBlack() }
func BGRed() *Style    { return NewStyle().BGRed() }
func BGGreen() *Style  { return NewStyle().BGGreen() }
func BGBrown() *Style  { return NewStyle().BGBrown() }
func BGBlue() *Style   { return NewStyle().BGBlue() }
func BGPurple() *Style { return NewStyle().BGPurple() }
func BGCyan() *Style   { return NewStyle().BGCyan() }
func BGGray() *Style   { return NewStyle().BGGray() }

func NewStyle() *Style {
	return &Style{
		count:      0,
		light:      false,
		brush:      0,
		background: 0,
	}
}

type (
	Style struct {
		count             int
		light             bool
		brush, background color
	}
)

func (it *Style) Reset()     { it.count = 0 }
func (it *Style) Count() int { return it.count }
func (it *Style) add(v string) string {
	it.count += len(v)
	return v
}

func (it *Style) n(n int, err error) *Style { return it }
func (it *Style) get(text string) string    { return get(it.light, it.add(text), it.brush, it.background) }
func (it *Style) Print(v ...any) *Style     { return it.n(fmt.Print(it.get(fmt.Sprint(v...)))) }
func (it *Style) Println(v ...any) *Style   { return it.n(fmt.Println(it.get(fmt.Sprint(v...)))) }
func (it *Style) Printf(format string, v ...any) *Style {
	return it.n(fmt.Print(it.get(fmt.Sprintf(format, v...))))
}

func (it *Style) Dark() *Style  { it.light = false; return it }
func (it *Style) Light() *Style { it.light = true; return it }

func (it *Style) Black() *Style  { it.brush = black; return it }
func (it *Style) Red() *Style    { it.brush = red; return it }
func (it *Style) Green() *Style  { it.brush = green; return it }
func (it *Style) Brown() *Style  { it.brush = brown; return it }
func (it *Style) Blue() *Style   { it.brush = blue; return it }
func (it *Style) Purple() *Style { it.brush = purple; return it }
func (it *Style) Cyan() *Style   { it.brush = cyan; return it }
func (it *Style) Gray() *Style   { it.brush = gray; return it }

func (it *Style) BGBlack() *Style  { it.background = black; return it }
func (it *Style) BGRed() *Style    { it.background = red; return it }
func (it *Style) BGGreen() *Style  { it.background = green; return it }
func (it *Style) BGBrown() *Style  { it.background = brown; return it }
func (it *Style) BGBlue() *Style   { it.background = blue; return it }
func (it *Style) BGPurple() *Style { it.background = purple; return it }
func (it *Style) BGCyan() *Style   { it.background = cyan; return it }
func (it *Style) BGGray() *Style   { it.background = gray; return it }
