package main

import "fmt"

func main() {
	for i := 0; ; i++ { // 終了条件がないと無限ループ
		fmt.Println(i)
	}
}
