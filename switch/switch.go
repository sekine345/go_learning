package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // os を定義したとき、os の値により処理を変更するswitch文
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}

/*
switch ステートメントは if - else ステートメントのシーケンスを短く書く方法です。
*/
