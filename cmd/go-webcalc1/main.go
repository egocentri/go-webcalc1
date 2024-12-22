package main

import (
	"github.com/gin-gonic/gin"
	"github.com/egocentri/go-webcalc1/internal/handlers"
)

func main() {
	a := gin.Default()

	a.POST("/api/v1/calculate", handlers.CalculateExpression)

	a.Run(":8080")
}
