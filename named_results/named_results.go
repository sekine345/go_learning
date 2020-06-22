package main

import "fmt"

func split(sum int) (x, y int) { // x と y を int で返す
	x = sum * 4 / 9
	y = sum - x
	return // 最初に定義した戻り地 x, y を、何も書かずに返す
}

func main() {
	fmt.Println(split(17))
}
