package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	fmt.Println(x, y)
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Println(f)
	fmt.Printf("%T, %v \n", f, f)
	var z uint = uint(f) // uint: 符号なし整数型
	fmt.Println(x, y, z)
}
