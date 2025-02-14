package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/egocentri/go-webcalc1/side/mod"
	"github.com/egocentri/go-webcalc1/side/services"
)

func CalculateExpression(c *gin.Context) {
	var req models.CalculateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Expression is not valid",
		})
		return
	}

	result, err := services.Calc(req.Expression)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
