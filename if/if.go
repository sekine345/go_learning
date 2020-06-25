package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 { // 条件式のみで良い
		return sqrt(-x) + "i" //関数内からその関数自身を呼べる
	}
	return fmt.Sprint(math.Sqrt(x)) // フォーマットした結果を文字列で返す。
}

func main() {
	fmt.Println(sqrt(-2), "\n", sqrt(4))
}
