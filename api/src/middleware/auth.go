package middleware

import (
	"context"

	//"strings"

	"github.com/gin-gonic/gin"
	"github.com/yuuuuut/gin-api/src/util"
)

// FirebaseAuth Middleware は送られてきたTokenが有効なTokenかチェックします
func FirebaseAuth(c *gin.Context) {
	auth := util.GetFirebase()

	token, err := util.GetVerifyIDToken(auth)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	/*
		authorizationToken := c.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
		if idToken == "" {
			c.JSON(400, gin.H{"error": "Id Token Not Avaliable"})
			c.Abort()
			return
		}


		token, err := auth.VerifyIDToken(context.Background(), idToken)
	*/

	decoded, err := auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Set("currentUserId", decoded.UID)
	c.Next()
}
