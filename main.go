package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.Default()

	app.GET("/v1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"teste": "OK!",
		})
	})

	app.Run("localhost:3001")
}

// test curl http://localhost:3001/v1
// => {"teste":"OK!"}
