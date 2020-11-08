package main

import "fmt"

func main() {
	prime := [6]int{2, 3, 5, 7, 11, 13} //配列を作成

	var s []int = prime[1:4] //int の配列を、他の配列の1番目から4番目まで取得
	fmt.Println(s)
	fmt.Println(prime[1:4]) //同義
}
