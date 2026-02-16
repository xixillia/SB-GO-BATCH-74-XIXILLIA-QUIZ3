package middlewares

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
)

func BasicAuth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid basic auth"})
			c.Abort()
			return
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid base64 encoding"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials format"})
			c.Abort()
			return
		}

		username, password := credentials[0], credentials[1]

		var basicPw string
		query := "SELECT password FROM Users WHERE username = $1"
		err = db.QueryRow(query, username).Scan(&basicPw)
		if err != nil {
			log.Printf("DB Query Error for username '%s': %v", username, err)
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
				c.Abort()
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
				c.Abort()
				return
			}
		}

		if basicPw != password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			c.Abort()
			return
		}

		c.Next()
	}
}
