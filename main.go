package main

import (
	"github.com/gin-gonic/gin"
	ss "strings"
)

func get_wrapper(c *gin.Context, t string) func(int, interface{}) {
	switch t {
	case "XML":
		return c.XML
	case "YAML":
		return c.YAML
	case "INDENTEDJSON":
		return c.IndentedJSON
	default:
		return c.JSON
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		t := c.DefaultQuery("type", "JSON")
		t = ss.ToUpper(t)
		get_wrapper(c, t)(200, gin.H{
			"message": "pong",
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
