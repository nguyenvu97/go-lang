package middleweres

import (
	"awesomeProject/src/jwt"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)


func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.URL.Path
		log.Println(" uri request: ", uri)
		
		jwtToken, err := jwt.ExtractClaim(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized", "description": ""})
			return
		}

		ctx := context.WithValue(c.Request.Context(), "subjectUUID", jwtToken.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
