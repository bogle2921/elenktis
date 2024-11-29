package middleware

// Can add multiple middleware functions in endpoints, executed from left to right in function parameters
import (
	"net/http"

	"github.com/bogle2921/elenktis/api/internal/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
