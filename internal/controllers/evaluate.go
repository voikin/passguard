package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) EvaluatePasswordHandler(ctx *gin.Context) {
	password := ctx.Query("password")
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	score := c.useCases.EvaluatePassword(password)
	ctx.JSON(http.StatusOK, gin.H{"score": score})
}
