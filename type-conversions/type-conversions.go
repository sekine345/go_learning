package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) //type conv
	//var f float64 = math.Sqrt(x*x + y*y) //type conv
	// -> cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
	var z uint = uint(f) //type conv
	//var z uint = f //type conv
	// -> cannot use f (type float64) as type uint in assignment
	fmt.Printf("x, Type: %T, value: %v\n", x, x)
	fmt.Printf("y, Type: %T, value: %v\n", y, y)
	fmt.Printf("f, Type: %T, value: %v\n", f, f)
	fmt.Printf("z, Type: %T, value: %v\n", z, z)
}
