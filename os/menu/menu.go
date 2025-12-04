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

package menu

import (
	"fmt"

	"github.com/iamkaro/GO-LIB/os/term"
	"github.com/iamkaro/GO-LIB/text"
)

var (
	commands = make([]*command, 0)
)

func AddCommand(cmd string, act func(), description string) {
	commands = append(commands, &command{
		Command:     cmd,
		Description: description,
		Action:      act,
	})
}

func Start(failureAction func()) {
	var list = map[string]*command{}
	for _, act := range commands {
		list[act.Command] = act
	}
	for {
		var v, ok = list[term.GetWord("menu> ")]
		if ok {
			v.Action()
			continue
		}
		if failureAction != nil {
			failureAction()
		}
	}
}

func Print() {
	for _, cmd := range commands {
		fmt.Printf(" %s %s\n",
			text.FixLength(cmd.Command, 12), cmd.Description)
	}
}

type (
	command struct {
		Command     string
		Description string
		Action      func()
	}
)
