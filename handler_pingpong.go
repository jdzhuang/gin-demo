package main

import (
	comm "gin-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingPongHandler() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	//p := comm.TheProfile()
	e.GET("/ping", func(c *gin.Context) {
		p := comm.TheProfile() //confirm singleton
		p.Inc("count", 1)
		t := c.DefaultQuery("type", "JSON")
		comm.ObjectWrapper(c, t)(200, gin.H{
			"message": "pong",
		})

	})

	return e
}
