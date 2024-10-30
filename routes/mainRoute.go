package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainRoute(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Politrack"}) })

}
