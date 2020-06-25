package main

import "fmt"

func main() {
	sum := 1
	fmt.Println(sum)
	for sum < 1000 { //初期化と後処理ステートメントの記述は任意、終了条件式は必須、while 文と同義
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}
