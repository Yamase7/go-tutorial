package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

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

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)
}

// 指定したアーティスト名を持つアルバムを検索します。
func albumsByArtist(name string) ([]Album, error) {
	// 返された行のデータを保持するスライス。
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// 行をループし、スキャンを使用して列データを構造体フィールドに割り当てます。
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}
