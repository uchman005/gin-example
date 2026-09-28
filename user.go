package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	{
		user.GET("/", func(ctx *gin.Context) {
			_= ctx.Error(ErrorTooManyRequest)
		})
	}
}
