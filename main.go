package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	RegisterUserRouter(router.Group("/api"))
	RegisterBindRoutes(router.Group("/testing"))
	router.Use(ErrorHandler())
	router.GET("/", get)
	router.GET("/me", me)
	router.GET("/welcome", welcome)
	router.GET("/name/:name/*action", pathParam)
	router.GET("/name/:name", pathParam)
	test := router.Group("/test")
	test.GET("", func(c *gin.Context) {
		c.Request.URL.Path = "/final"
		router.HandleContext(c)
	})
	router.POST("/submit", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/result")
	})
	router.GET("/result", func(c *gin.Context) {
		c.String(http.StatusOK, "Redirected here!")
	})
	router.GET("/final", final)
	protected := router.Group("/api")
	protected.Use(Auth())
	{
		protected.GET("/users", func(ctx *gin.Context) {
			name, ok := ctx.Get("name")
			if !ok {
				_ = ctx.Error(ErrNotFound)
				return
			}
			age, ok := ctx.Get("me")
			if !ok {
				fmt.Println("User age not in context")
				_ = ctx.Error(ErrNotFound)
				return
			}

			fmt.Println(name, age)
			ctx.JSON(200, gin.H{"name": name, "age": age})
		})
	}
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"name": "Obiefuna Marcel",
		})
	})
	router.Run()
}
