package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mattdumler/project-management/internal/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/1.0")
	controllers.RegisterProjectsController(v1)
}
