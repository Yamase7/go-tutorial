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
	router.GET("/albums/:id", getAlbumByID)
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

// 受信したidパラメーターとIDフィールドの値が一致するアルバムを探し、そのアルバムを応答として返す
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// アルバムのリストをループしてIDフィールドの値がパラメーターと一致するアルバムを探す
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
