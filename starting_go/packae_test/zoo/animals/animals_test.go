// test: goはファイル名に"_test.go"とあれば、テスト用のパッケージであると認識する
// Testで始まる関数がテストの単位

/* [terminal in zoo]
$ go test ./animals  // 全ての結果のみ表示
ok  	zoo/animals	0.250s
$
$ go test -v ./animals  // 実行結果の詳細を表示
=== RUN   TestElephantFeed
--- PASS: TestElephantFeed (0.00s)
=== RUN   TestMonkeyFeed
--- PASS: TestMonkeyFeed (0.00s)
=== RUN   TestRabbitFeed
--- PASS: TestRabbitFeed (0.00s)
PASS
ok  	zoo/animals	0.162s
*/


package animals

import (
    "testing"
)


func TestElephantFeed(t *testing.T) {
    expect := "Grass"
    result := ElephantFeed()  // "animals.ElephantFeed()"としなくとも実行できる

    if expect != result {
        t.Errorf("%s != %s", expect, result)
    }
}

func TestMonkeyFeed(t *testing.T) {
    expect := "Banana"
    result := MonkeyFeed()

    if expect != result {
        t.Errorf("%s != %s", expect, result)
    }
}

func TestRabbitFeed(t *testing.T) {
    expect := "Carrot"
    result := RabbitFeed()

    if expect != result {
        t.Errorf("%s != %s", expect, result)
    }
}
