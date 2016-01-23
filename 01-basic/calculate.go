package main

import (
	"math"
	"strings"
)

import "github.com/gopherjs/gopherjs/js"

func calculate(i int) int {
	return 1000 - i
}

func othercalc(i int) int {
	return i*2 + 555
}

func main() {
	// expose functions main.calculate and main.othercalc
	// to javascript as my mypackage.mycalculate and
	// mypackage.myothercalc
	js.Global.Set("mypackage", map[string]interface{}{
		"mycalculate": calculate,
		"myothercalc": othercalc,
		"pow":         math.Pow10,
		"upper":       strings.ToUpper,
	})
}
