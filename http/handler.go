package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgjules/deckr/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// handleHealthCheck godoc
// @Summary      Health Check
// @Description  checks if server is running
// @Tags         core
// @Produce      json
// @Success      200  {string}	I'm alive!
// @Router       / [get]
func (Server) handleHealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, Success{Message: "I'm alive!"})
	}
}

// handleVersion godoc
// @Summary      Health Check
// @Description  checks the server's version
// @Tags         core
// @Produce      json
// @Success      200  {object}  build.Info
// @Router       /version [get]
func (s *Server) handleVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, s.build)
	}
}

func (Server) handleSwagger() gin.HandlerFunc {
	docs.SwaggerInfo.BasePath = "/"

	url := ginSwagger.URL("/swagger/doc.json")

	return ginSwagger.WrapHandler(swaggerFiles.Handler, url)
}
