package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 1; i < 10; i++ {
		defer fmt.Println(i) //last_in_first_out, あとに入ったものが、先に出てくる
	}

	fmt.Println("done")
}
