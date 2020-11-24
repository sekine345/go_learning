package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

/*
if ステートメントは、 for のように、条件の前に、評価するための簡単なステートメントを書くことができます。

    if [A]; [B] {
		[処理]
	}
[A] で定義した変数を、[B]で評価し、[B]の条件にあっていたら [処理]を行う。
[A] で定義した変数は if スコープ内でのみ使える。
*/
