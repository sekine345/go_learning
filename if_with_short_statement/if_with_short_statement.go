package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//if v := math.Pow(x, n); v < lim { // vを数式で定義したとき; もし v < lim なら
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20), //改行したならカンマを付ける
	)
}
