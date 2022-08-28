package controller

import (
	"net/http"

	model "example.com/gin-basic/model/album"
	"github.com/gin-gonic/gin"
)

var albums = []model.Album{
	{ID: "1", Title: "Test1", Artist: "jaja", Price: 2.5},
	{ID: "2", Title: "Test2", Artist: "jaja", Price: 12.2},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
