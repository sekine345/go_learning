// 演算
/*
数値: +, -, *, /, %
文字列: +
論理値: AND(&&), OR(||), NOT(!)
*/


package main

import "fmt"


func main(){
    //var x int
    //x = 10 % 3 //あまり
    //x += 3 // x = x + 3  // -, * も同様
    //x++  // x = x + 1
    //fmt.Println(x)

    //var s string
    //s = "hello " + "world."
    //fmt.Println(s)

    a := true
    b := false
    fmt.Println(a && b)  // true AND false
    fmt.Println(a || b)  // true OR false
    fmt.Println(!a)  // NOR true
}
