package middleware

import "github.com/gin-gonic/gin"

func SetCookie() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.SetCookie("gin_admin_real","1234567890",60,"/","localhost",false,true)
		context.Next()
	}
}
