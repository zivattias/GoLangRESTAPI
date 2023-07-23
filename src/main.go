package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}
