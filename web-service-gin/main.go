package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// レコードアルバムを表すデータ
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// スライスでデータを用意する。
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// すべてのアルバムのリストをJSONで返す
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// 受け取ったJSONからアルバムを追加する
func postAlbums(c *gin.Context) {
	var newAlbum album

	// 受け取ったJSONをnewAlbumにバインドする
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// 新しいアルバムをスライスに追加する
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
