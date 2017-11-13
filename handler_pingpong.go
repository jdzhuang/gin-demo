package main

import (
	comm "gin-demo/common"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func pingpongHandler() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	//p := comm.TheProfile()
	e.GET("/ping", func(c *gin.Context) {
		p := comm.TheProfile() //confirm singleton
		defer p.Since(time.Now())
		p.Inc("count", 1)
		t := c.DefaultQuery("type", "JSON")
		comm.ObjectWrapper(c, t)(200, gin.H{
			"message": "pong",
		})

	})

	e.GET("/slow", func(c *gin.Context) {
		p := comm.TheProfile() //confirm singleton
		defer p.Since(time.Now())
		p.Inc("count", 1)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		t := c.DefaultQuery("type", "JSON")
		comm.ObjectWrapper(c, t)(200, gin.H{
			"message": "pong",
		})

	})

	return e
}
