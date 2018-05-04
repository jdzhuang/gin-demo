package main

import (
	"fmt"
	comm "gin-demo/common"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func pingpongHandler() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	p := comm.TheProfile()
	e.GET("/ping", func(c *gin.Context) {
		fmt.Fprintf(os.Stdout, "/ping? in.\n")
		defer p.Since(time.Now())
		p.Inc("count", 1)
		t := c.DefaultQuery("type", "JSON")
		fmt.Fprintf(os.Stdout, "/ping? done.\n")
		comm.ObjectWrapper(c, t)(200, gin.H{
			"message": "pong",
		})

	})

	e.GET("/slow", func(c *gin.Context) {
		fmt.Fprintf(os.Stdout, "/slow? in.\n")
		defer p.Since(time.Now())
		p.Inc("count", 1)
		fmt.Fprintf(os.Stdout, "/slow?\n")
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		t := c.DefaultQuery("type", "JSON")
		fmt.Fprintf(os.Stdout, "/slow? done.\n")
		comm.ObjectWrapper(c, t)(200, gin.H{
			"message": "pong",
		})

	})

	return e
}
