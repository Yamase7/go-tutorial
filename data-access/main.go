package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func main() {
	// 接続プロパティの宣言
	user := os.Getenv("DBUSER")
	passwd := os.Getenv("DBPASS")
	host := "127.0.0.1"
	port := "5432"
	dbName := "recordings"

	// 接続プロパティの代入を空白区切りで結合
	keyValue := "user=" + user + " " +
		"password=" + passwd + " " +
		"host=" + host + " " +
		"port=" + port + " " +
		"database=" + dbName

	// データベースハンドルの取得
	var err error
	db, err = sql.Open("pgx", keyValue)
	if err != nil {
		log.Fatal(err)
	}

	// データベースへの接続を閉じる
	defer db.Close()

	// 接続確認
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
