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

// Hellosは、名前を付けられた人それぞれに挨拶メッセージを関連付けたマップを返す。
func Hellos(names []string) (map[string]string, error) {
    // 名前とメッセージを関連付けるためのマップを作成
    messages := make(map[string]string)

    // 受信した名前のスライスをループし、Hello関数を呼び出してそれぞれの名前のメッセージを取得する。
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        // マップの中で、取得したメッセージを名前に関連付ける。
        messages[name] = message
    }

    return messages, nil
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
