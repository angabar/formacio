package localRouter

import (
	controller "example.com/gin-basic/controller/album"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	router.GET("/albums", controller.GetAlbums)
}
