package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"school-bbs/util"
)

func AuthClubAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := util.ParseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			return
		}
		if userClaim == nil || userClaim.IsAdmin != 2 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "你不是社团管理员",
			})
			return
		}
		c.Next()
	}
}
