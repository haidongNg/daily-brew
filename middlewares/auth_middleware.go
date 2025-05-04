package middlewares

import (
	"daily-brew/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			c.Abort()
			return
		}

		// Đảm bảo đúng định dạng "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.VerifyAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		/// Lưu claims vào context để các route sau có thể dùng
		c.Set("memberId", claims.MemberID)
		c.Set("role", claims.Role)
		c.Set("refreshTokenId", claims.RefreshTokenId)
		c.Next()
	}
}
