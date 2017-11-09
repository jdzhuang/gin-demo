package main

import (
	comm "gin-demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProfileHandler() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	p := comm.TheProfile()
	e.GET("/profile", func(c *gin.Context) {
		t := c.DefaultQuery("type", "JSON")
		comm.ObjectWrapper(c, t)(200, gin.H{
			"count": p.Get("count"),
		})

	})

	return e
}
