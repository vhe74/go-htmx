package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Greet struct {
	Name string
}

func handleHello(c *gin.Context) {
	who := c.Param("who")
	if who == "" {
		c.HTML(http.StatusOK, "hello.html", nil)
	} else {
		g := Greet{Name: who}
		c.HTML(http.StatusOK, "hello.html", g)
	}
}
