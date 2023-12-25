package middlewares

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userEmail := session.Get("userEmail")
		if userEmail == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login-required/")
			c.Abort()
			return
		}
		c.Next()
	}
}
