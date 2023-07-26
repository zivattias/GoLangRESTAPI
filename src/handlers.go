package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbum adds an album from JSON received in the request body
func postAlbum(c *gin.Context) {
	var newAlbum album
	// Calling BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	fmt.Printf("New album struct: %+v\n", newAlbum)

	// Adding new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumById locates the album whose ID value matches the request param ID
// sent by the client, returns that album as a response
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID matches the param
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("album %s not found", id)})
}

func sendAlbumsByEmail(c *gin.Context) {
	if sent, err := sendEmail(c); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, sent)
	}
}
