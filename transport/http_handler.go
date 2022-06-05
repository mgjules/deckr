package transport

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//go:embed v1/deck_service.swagger.json
var docsFS embed.FS

func (HTTPServer) handleHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, Success{Message: "I'm alive!"})
	}
}

func (s *HTTPServer) handleVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, s.build)
	}
}

func (HTTPServer) handleSwagger() gin.HandlerFunc {
	url := ginSwagger.URL("/docs/v1/deck_service.swagger.json")

	return ginSwagger.WrapHandler(swaggerFiles.Handler, url)
}
