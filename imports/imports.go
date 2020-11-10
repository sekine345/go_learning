package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

/*
# print系
- 接頭辞 F
    書き込み先を指定
- 接頭辞 S
    結果を文字列で返す
- 接尾辞 f
    フォーマットを指定
- 接尾辞 ln
    オペランド間にスペース、最後に改行を追加
*/

/*
# printf のフォーマット
- %t
    bool
- %d
    整数
- %g
    小数、複素数
- %s
    文字列
- %p
    チャネル、ポインタ
*/
