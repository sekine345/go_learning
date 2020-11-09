package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		fmt.Println(math.pi)
		-> cannot refer to unexported name math.pi
	*/
	fmt.Println(math.Pi)
	fmt.Println(math.E)
}

/*
最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポート（公開）された名前( exported name )
エクスポートされていない名前(小文字ではじまる名前)は、外部のパッケージからアクセスすることはできない
*/
