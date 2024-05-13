package middlewares

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "final-task/helpers"
    "strings"
)

// AuthMiddleware berfungsi untuk melakukan otentikasi pada setiap permintaan yang masuk.
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Akses ditolak. Tidak ada token yang diberikan."})
            c.Abort()
            return
        }

        token = strings.TrimPrefix(token, "Bearer ")
        claims, err := helpers.VerifyToken(token)
        
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}
