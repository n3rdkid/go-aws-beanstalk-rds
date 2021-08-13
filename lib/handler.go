package lib

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RequestHandler -> function
type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler -> creates a new request handler
func NewRequestHandler() RequestHandler {
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello 世界!",
		})
	})
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	return RequestHandler{Gin: engine}
}
