package common

import (
	"github.com/gin-gonic/gin"
	ss "strings"
)

func ObjectWrapper(c *gin.Context, t string) func(int, interface{}) {
	switch ss.ToUpper(t) {
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
