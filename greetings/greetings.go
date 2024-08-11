package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello関数は指定された人物への挨拶を返す。
func Hello(name string) (string, error) {
    // 名前が与えられていない場合はメッセージとともにエラーを返す。
    if name == "" {
        return "", errors.New("empty name")
    }

    // ランダムなフォーマットでメッセージを作成する。
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// randomFormat は、挨拶メッセージのセットを返します。
// 返されるメッセージはランダムに選択されます。
func randomFormat() string {
    // メッセージ候補を集めたスライスを作成
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // formatsのスライスにランダムなインデックスを指定して、
    // ランダムに選択されたメッセージ・フォーマットを返す。
    return formats[rand.Intn(len(formats))]
}
