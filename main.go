package main

import (
	"gee_web/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// 业务代码
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"))
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// 业务代码
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"))
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you are at %s\n", c.Query("name"))
	})

	r.POST("/login", func(c *gee.Context) {
		// 业务代码
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
