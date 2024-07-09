package middleware

import (
	"errors"
	"strconv"
	"time"

	"github.com/congwa/gin-start/global"
	"github.com/congwa/gin-start/model/common/response"
	"github.com/congwa/gin-start/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("缺失token", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, utils.ErrTokenExpired) || errors.Is(err, utils.ErrTokenNotValidYet) {
				response.NoAuth("token已过期", c)
				c.Abort()
				return
			}
			response.NoAuth("token解析失败", c)
			utils.ClearToken(c)
			c.Abort()
			return
		}

		// 存入上下文
		c.Set("claims", claims)

		if claims.ExpiresAt.Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.Config.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds()))
		}
		c.Next()
	}
}
