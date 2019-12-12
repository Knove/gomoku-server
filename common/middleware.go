package common

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/*
CORSMiddleware 跨域中间件

*/
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

/*
JWTMiddleware JWT 中间件

*/
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint64
		var data interface{}
		var token string

		code = OK
		tokenArray := c.Request.Header["Token"]
		if len(tokenArray) > 0 {
			token = tokenArray[0]
		}

		if "/user/login" == c.FullPath() {
			c.Next()
			return
		}

		if token == "" {
			code = Unauthorized
		} else {
			_, err := ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = Unauthorized
				default:
					code = Unauthorized
				}
			}
		}

		if code != OK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  GetErrorMessage(code, ""),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
