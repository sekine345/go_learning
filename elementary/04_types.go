// 変数型
/*
string: "foge"
int: 10  // int8, int32などもある。基本はintで良い
float64: 12.3  // float8なども
bool: false, true
nil

Default:
- var string -> ""  // 空文字
- var int    -> 0
- var bool   -> false

書式付きの出力: fmt.Printf
- %d: 整数値
- %f: 小数値
- %s: 文字列
- %t: 真偽値
*/


package main

import "fmt"


func main(){
    a := 10
    b := 12.3
    c := "hoge"
    var d bool
    fmt.Printf("a: %d, b: %f, c: %s, d: %t\n", a, b, c, d)
}
