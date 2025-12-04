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
	"math"

	"github.com/iamkaro/GO-LIB/check"
	"github.com/iamkaro/GO-LIB/number"
)

type (
	Paths []Path
	Float struct{ paths Paths }
	Int   struct{ paths Paths }
)

func (it Paths) Int() *Int     { return &Int{paths: it} }
func (it Paths) Float() *Float { return &Float{paths: it} }
func (it Paths) Walking(handle func(path Path)) {
	for _, path := range it {
		handle(path)
	}
}

func (it *Int) Max(ifEmpty int64) int64 {
	var xp = number.Max[int64]()
	it.paths.Walking(func(path Path) { xp.Value(path.Name().Int(math.MinInt64)) })
	return check.Check(xp.Out(), math.MinInt64, ifEmpty)
}
func (it *Int) Min(ifEmpty int64) int64 {
	var xp = number.Min[int64]()
	it.paths.Walking(func(path Path) { xp.Value(path.Name().Int(math.MaxInt64)) })
	return check.Check(xp.Out(), math.MaxInt64, ifEmpty)
}

func (it *Float) Max(ifEmpty float64) float64 {
	var xp = number.Max[float64]()
	it.paths.Walking(func(path Path) { xp.Value(path.Name().Float(-math.MaxFloat64)) })
	return check.Check(xp.Out(), -math.MaxFloat64, ifEmpty)
}
func (it *Float) Min(ifEmpty float64) float64 {
	var xp = number.Min[float64]()
	it.paths.Walking(func(path Path) { xp.Value(path.Name().Float(math.MaxFloat64)) })
	return check.Check(xp.Out(), math.MaxFloat64, ifEmpty)
}
