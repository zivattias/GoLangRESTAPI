package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums/send", sendAlbumsByEmail)

	router.Run("localhost:8080")
}
