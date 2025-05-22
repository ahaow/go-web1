package controllers

import (
	"github.com/gin-gonic/gin"
	"go-web1/models"
	"net/http"
)

func Register(ctx *gin.Context) {
	var users models.User

	if err := ctx.ShouldBindJSON(users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
