package greetings

import (
    "testing"
    "regexp"
)

// TestHelloNameは、名前を指定してgreetings.Helloを呼び出し、
// 有効な戻り値があるかどうかをチェックする。
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmptyは空の文字列でgreetings.Helloを呼び出し、
// エラーをチェックする。
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}