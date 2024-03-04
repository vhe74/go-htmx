package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
