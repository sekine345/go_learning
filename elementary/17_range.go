// range
// 配列、slice, map のfor文で使う


package main

import "fmt"


func main() {
    fmt.Println("--- slice")
    s := []int{1, 3, 5}
    for i, v := range s {  // index, value
        fmt.Println(i, v)
    }

    fmt.Println("--- slice, withot index")
    for _, v := range s {
        fmt.Println(v)
    }

    fmt.Println("--- map")
    m := map[string]int{"a":1, "b":2}
    for k, v := range m {  // key, value
        fmt.Println(k, v)
    }
}
