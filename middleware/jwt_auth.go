package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"web_app/utils"

	"github.com/gin-gonic/gin"
)

// JwtAuth jwt权限验证
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		fmt.Println(url)
		if url == "/login" || strings.HasPrefix(url, "/swagger") {
			c.Next()
		} else {
			user, err := utils.ParseToken(c)
			if err != nil {
				c.JSON(http.StatusOK, utils.ResOvertime())
				c.Abort()
				return
			}
			if user.ExpiresAt <= time.Now().Unix() {
				c.JSON(http.StatusOK, utils.ResOvertime())
				c.Abort()
				return
			}
			// c.Set("user_id", user.ID)
			c.Next()
		}
	}
}
