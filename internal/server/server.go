package server

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	router := gin.Default()

	// register any middleware

	return router
}
