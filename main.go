package main

import (
	"api/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.AppRoutes(app)

	app.Run("localhost:3001")
}

// test curl http://localhost:3001/v1
// => {"teste":"OK!"}
