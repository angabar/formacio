package main

import (
	localRouter "example.com/gin-basic/localRouter/album"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	localRouter.SetRouter(router)

	router.Run("localhost:8080")
}
