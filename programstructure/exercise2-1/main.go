package main

import (
	"fmt"

	"book/programstructure/exercise2-1/tempconv"
)

func main() {
	var c tempconv.Celsius = 126.85
	var f tempconv.Fahrenheit = 32.0
	var k tempconv.Kelvin = 300.0
	fmt.Println("Temp:\t", tempconv.CToK(c))
	fmt.Println("Temp:\t", tempconv.FToC(f))
	fmt.Println("Temp:\t", tempconv.KToC(k))
}
