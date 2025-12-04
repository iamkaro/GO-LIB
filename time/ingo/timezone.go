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

package ingo

import (
	"log"
	"time"
)

const (
	UTC   Timezone = "UTC"
	Local Timezone = "Local"
	Iran  Timezone = "Asia/Tehran"
	USA   Timezone = "America/New_York"
	EET   Timezone = "EET"
)

type (
	Timezone string
)

func (it Timezone) location() *time.Location {
	var loc, err = time.LoadLocation(string(it))
	if (loc == nil) || (err != nil) {
		log.Println("error: Timezone(", it, "), ", err)
		return time.UTC
	}
	return loc
}
