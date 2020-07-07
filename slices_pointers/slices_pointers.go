package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"      //b[0]を書き換え、もとの配列 names のスライスを編集することで、もとの配列 names も変更される
	fmt.Println(a, b) //b -> names -> a と、他のスライスも編集される。バグに注意
	fmt.Println(names)
}
