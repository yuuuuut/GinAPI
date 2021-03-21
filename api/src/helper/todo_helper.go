package helper

import "github.com/gin-gonic/gin"

func GetQueryIndex(c *gin.Context, t string) string {
	str, ok := c.GetQuery(t)
	if !ok {
		if t == "offset" {
			str = "0"
		} else if t == "limit" {
			str = "3"
		}
	}

	return str
}
