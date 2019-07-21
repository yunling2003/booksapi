package middleware

import (
	"booksapi/src/controllers"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := controllers.GetHeaderToken(c)
		newToken := controllers.RefreshTokenForUserOrVisitor(token)
		controllers.SetHeaderToken(c, newToken)
	}
}