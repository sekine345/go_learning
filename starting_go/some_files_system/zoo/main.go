// 複数ファイルで一つのシステムを実装するやり方
/*
goファイルはどれもパッケージに属していなければいけない
パッケージ名は何でも良いが、管理のしやすさから格納するディレクトリと同じ名前を推奨
パッケージに属した関数を呼ぶには、import でパスを指定する
*/

/* 参考書のGoのバージョンが低かったため、最新版での実行方法をメモ
1. 以下コマンドを実行
```sh
$ cd zoo
$ go mod init zoo
go: creating new go.mod: module zoo
go: to add module requirements and sums:
	go mod tidy
```

2. importに記載する自作モジュールは、zooからのパスを記載
*/



package main

import (
    "fmt"
    "zoo/animals"
)


func main() {
    fmt.Println(animals.ElephantFeed())
    fmt.Println(animals.MonkeyFeed())
    fmt.Println(animals.RabbitFeed())
}

