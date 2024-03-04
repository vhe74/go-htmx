package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHtmxPage(c *gin.Context) {
	c.HTML(http.StatusOK, "htmx.html", nil)
}

func handleHtmxClicked(c *gin.Context) {
	c.HTML(http.StatusOK, "htmx_partial_clicked.html", nil)
}
