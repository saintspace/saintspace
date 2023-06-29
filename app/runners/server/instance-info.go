package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetInstanceInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Here's your info about your SaintSpace instance",
	})
}
