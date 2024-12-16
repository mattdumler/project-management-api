package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mattdumler/project-management/internal/handlers"
)

func RegisterProjectsController(router *gin.RouterGroup) {
	projects := router.Group("/projects")
	{
		projects.POST("/", handlers.CreateProject)
	}
}
