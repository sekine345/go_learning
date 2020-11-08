package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	o := runtime.GOOS
	fmt.Println(o)
	switch os := runtime.GOOS; os { // osを定義したとき、osが...
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%s.", os)
	}
	fmt.Println("..fin") //switchのあとは普通に処理
}
