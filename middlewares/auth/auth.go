package auth

import (
	"crypto/sha256"
	"crypto/subtle"
	"main/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {

	username, password, ok := c.Request.BasicAuth()
	if ok {

		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))

		usernameDB, passwordDB := user.GetLoginInfo(username)

		expectedUsernameHash := sha256.Sum256([]byte(usernameDB))
		expectedPasswordHash := sha256.Sum256([]byte(passwordDB))

		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

		if usernameMatch && passwordMatch {
			c.Next()
			return
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{})
		}
	}
	c.Abort()
	c.JSON(http.StatusUnauthorized, gin.H{})
}
