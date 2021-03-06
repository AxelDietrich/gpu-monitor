package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/ping", ping)
  r.Run()
}

func ping(c *gin.Context) { 
  c.String(http.StatusOK, "pong")
}
