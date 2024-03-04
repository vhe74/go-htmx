package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
}

func NewServer(Address string) (Server, error) {
	return Server{listenAddr: Address}, nil
}

func (s *Server) Start() error {
	log.Println("Starting server Listening on ", s.listenAddr)

	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")

	router.GET("/", handleIndex)

	router.GET("/ping", handlePing)

	router.GET("/hello", handleHello)
	router.GET("/hello/:who", handleHello)

	router.GET("/htmx/page", handleHtmxPage)
	router.POST("/htmx/clicked", handleHtmxClicked)

	router.Static("/static", "./static")

	router.Run(s.listenAddr)
	return nil
}
