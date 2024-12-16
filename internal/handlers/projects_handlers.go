package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProject(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"title": "Project One",
	})
}
