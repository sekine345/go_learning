// for

package main

import "fmt"


func main() {
    fmt.Println("-----")
    for i := 0; i < 10; i++ {
        if i == 3 {
            // break
            continue
        }
        fmt.Println(i)
    }

    fmt.Println("-----")
    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }

    fmt.Println("-----")
    // while
    i = 0
    for {  // 無限ループ
        fmt.Println(i)
        i++
        if i == 3 { break }
    }

    fmt.Println("-----")
    // スライスのfor文の使い方、rangeを使わない場合
    a := []int{1,3,5}
    for i := 0; i < len(a); i++ {
        fmt.Println(a[i])
    }

    fmt.Println("-----")
    for i := 1; i < 10; i++ {
        fmt.Println(i) // 1, 3, 5, 9
        i++
    }
}
