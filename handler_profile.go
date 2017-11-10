package main

import (
	comm "gin-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ProfileHandler() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	p := comm.TheProfile()
	p.Declare(10*time.Millisecond, "10 ms")
	p.Declare(20*time.Millisecond, "20 ms")
	p.Declare(40*time.Millisecond, "40 ms")
	p.Declare(80*time.Millisecond, "80 ms")
	p.Declare(160*time.Millisecond, "160 ms")
	e.GET("/profile", func(c *gin.Context) {
		t := c.DefaultQuery("type", "JSON")
		comm.ObjectWrapper(c, t)(200, gin.H{
			"profile": p.String(),
		})

	})

	return e
}
