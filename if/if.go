package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i" //文字列に変換
	}
	return fmt.Sprint(math.Sqrt(x)) //文字列に変換
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Printf("%v, %T\n", sqrt(3), sqrt(3))
	fmt.Printf("%v, %T\n", sqrt(-9), sqrt(-9))
}
